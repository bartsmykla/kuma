package uninstall

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"

	kumactl_cmd "github.com/kumahq/kuma/app/kumactl/pkg/cmd"
	"github.com/kumahq/kuma/app/kumactl/pkg/install/k8s"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	typedbatchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"

	kuma_version "github.com/kumahq/kuma/pkg/version"
)

const (
	AppLabel          = "kuma.io/bpf-cleanup"
	BpfCleanupJobName = "kuma-bpf-cleanup"
	BpfCleanupImage   = "kumahq/kuma-init"
)

var (
	KumaBpfLabelSelector      = fmt.Sprintf("%s=%s", AppLabel, BpfCleanupJobName)
	KumaBpfCleanupAppSelector = metav1.ListOptions{LabelSelector: KumaBpfLabelSelector}
)

type CleanupJobProps struct {
	nodeName string
	phase    corev1.PodPhase
}

type CleanupJob struct {
	*sync.RWMutex
	jobClient   typedbatchv1.JobInterface
	podClient   typedcorev1.PodInterface
	startedJobs map[string]*CleanupJobProps
	stdout      io.Writer
	stderr      io.Writer
}

type ebpfArgs struct {
	BPFFsPath           string
	Timeout             time.Duration
	CleanupImageVersion string
	RemoveOnly          bool
	Namespace           string
}

func newUninstallEbpf(root *kumactl_cmd.RootContext) *cobra.Command {
	args := ebpfArgs{
		// default value that we inject in pod injector
		BPFFsPath:           root.InstallCpContext.Args.Ebpf_bpffs_path,
		Timeout:             120 * time.Second,
		CleanupImageVersion: kuma_version.Build.Version,
		RemoveOnly:          false,
		Namespace:           root.InstallCpContext.Args.Namespace,
	}

	cmd := &cobra.Command{
		Use:   "ebpf",
		Short: "Uninstall BPF files from the nodes",
		Long:  "Uninstall BPF files from the nodes by removing BPF programs from all the nodes",
		RunE: func(cmd *cobra.Command, _ []string) error {
			kubeClientConfig, err := k8s.DefaultClientConfig("", "")
			if err != nil {
				return errors.Wrap(err, "Could not detect Kubernetes configuration")
			}

			k8sClient, err := kubernetes.NewForConfig(kubeClientConfig)
			if err != nil {
				return errors.Wrap(err, "Could not create Kubernetes client")
			}

			ctx, cancel := context.WithTimeout(cmd.Context(), args.Timeout)
			defer cancel()

			nodes, err := k8sClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
			if err != nil {
				return errors.Wrap(err, "Failed obtaining nodes from Kubernetes cluster")
			}

			jobResource := CleanupJob{
				jobClient:   k8sClient.BatchV1().Jobs(args.Namespace),
				podClient:   k8sClient.CoreV1().Pods(args.Namespace),
				startedJobs: map[string]*CleanupJobProps{},
				stdout:      cmd.OutOrStdout(),
				stderr:      cmd.ErrOrStderr(),
				RWMutex:     &sync.RWMutex{},
			}

			if args.RemoveOnly {
				if err := jobResource.Cleanup(ctx); err != nil {
					return errors.Wrap(err, "Failed cleaning jobs")
				}

				return nil
			}

			for id, node := range nodes.Items {
				jobName := fmt.Sprintf("%s-%d", BpfCleanupJobName, id)
				jobSpec := genJobSpec(jobName, node.Name, &args)

				if _, err := jobResource.jobClient.Create(ctx, jobSpec, metav1.CreateOptions{}); err != nil {
					return errors.Wrap(err, "failed creating cleanup job")
				}

				jobResource.startedJobs[jobName] = &CleanupJobProps{
					nodeName: node.Name,
				}
			}

			watcher, err := jobResource.podClient.Watch(ctx, metav1.ListOptions{
				LabelSelector: KumaBpfLabelSelector,
				Watch:         true,
			})
			if err != nil {
				return errors.Wrap(err, "failed to create pod watcher")
			}

			errCh := make(chan error, 1)

			defer func() {
				if e := jobResource.Cleanup(ctx); e != nil {
					errCh <- e
				}
			}()

			go func() {
				errCh <- jobResource.Watch(ctx, watcher)
			}()

			select {
			case <-ctx.Done():
				return nil
			case err := <-errCh:
				return err
			}
		},
	}

	cmd.Flags().StringVar(&args.Namespace, "namespace", args.Namespace, "namespace where job is created")
	cmd.Flags().StringVar(&args.BPFFsPath, "bpffs-path", args.BPFFsPath, "path where bpf programs were installed")
	cmd.Flags().DurationVar(&args.Timeout, "timeout", args.Timeout, "timeout for whole process of removing left files")
	cmd.Flags().StringVar(&args.CleanupImageVersion, "cleanup-image-version", args.CleanupImageVersion, "version of cleanup ebpf job image")
	cmd.Flags().BoolVar(&args.RemoveOnly, "remove-only", args.RemoveOnly, "cleanup jobs and pods only")

	return cmd
}

func genJobSpec(jobName, nodeName string, args *ebpfArgs) *batchv1.Job {
	return &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobName,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						AppLabel: BpfCleanupJobName,
					},
				},
				Spec: corev1.PodSpec{
					NodeName: nodeName,
					Containers: []corev1.Container{
						{
							Name:  BpfCleanupJobName,
							Image: fmt.Sprintf("%s:%s", BpfCleanupImage, args.CleanupImageVersion),
							Command: []string{
								"kumactl",
								"uninstall",
								"transparent-proxy",
								"--ebpf-enabled",
								"--ebpf-bpffs-path", args.BPFFsPath,
							},
							SecurityContext: &corev1.SecurityContext{
								Privileged: new(bool),
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "bpf-fs-path",
									MountPath: args.BPFFsPath,
								},
							},
						},
					},
					RestartPolicy: corev1.RestartPolicyNever,
					Volumes: []corev1.Volume{
						{
							Name: "bpf-fs-path",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: args.BPFFsPath,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *CleanupJob) getPodLogs(
	ctx context.Context,
	pod *corev1.Pod,
) (string, error) {
	podLogs, err := r.podClient.GetLogs(
		pod.Name,
		&corev1.PodLogOptions{},
	).Stream(ctx)
	if err != nil {
		return "", fmt.Errorf("getting %s pod logs failed with error: %s",
			pod.Name, err)
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return "", fmt.Errorf("copy data from podLogs to buf failed with error:"+
			" %s", err)
	}

	return buf.String(), nil
}

func (r *CleanupJob) Watch(ctx context.Context, watcher watch.Interface) error {
	eg, _ := errgroup.WithContext(ctx)

	eg.Go(func() error {
		r.Lock()
		defer r.Unlock()

		for event := range watcher.ResultChan() {
			if len(r.startedJobs) == 0 {
				break
			}

			pod, ok := event.Object.(*corev1.Pod)
			if !ok {
				return nil
			}

			jobName := pod.Labels["job-name"]

			job, ok := r.startedJobs[jobName]
			if !ok || pod.Status.Phase == job.phase {
				continue
			}

			switch pod.Status.Phase {
			case corev1.PodSucceeded:
				_, _ = fmt.Fprintf(r.stdout,
					"cleanup for node: %s finished successfully\n", job.nodeName)
				delete(r.startedJobs, jobName)
				continue
			case corev1.PodFailed:
				if logs, err := r.getPodLogs(ctx, pod); err != nil {
					_, _ = fmt.Fprintf(r.stdout,
						"cleanup for node: %s failed but couldn't get any logs",
						job.nodeName)
				} else {
					_, _ = fmt.Fprintf(r.stdout,
						"cleanup for node: %s failed: %s", job.nodeName, logs)
				}
				delete(r.startedJobs, jobName)
				continue
			}

			job.phase = pod.Status.Phase
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

func (r *CleanupJob) Cleanup(ctx context.Context) error {
	policy := metav1.DeletePropagationBackground
	deleteImmediately := metav1.DeleteOptions{
		GracePeriodSeconds: new(int64),
		PropagationPolicy:  &policy,
	}

	if err := r.jobClient.DeleteCollection(ctx, deleteImmediately, KumaBpfCleanupAppSelector); err != nil {
		return fmt.Errorf("failed to delete jobs %s", err)
	}

	if err := r.podClient.DeleteCollection(ctx, deleteImmediately, KumaBpfCleanupAppSelector); err != nil {
		return fmt.Errorf("failed to delete pods %s", err)
	}

	return nil
}
