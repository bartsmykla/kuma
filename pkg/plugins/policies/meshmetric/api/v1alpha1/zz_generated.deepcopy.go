//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	commonv1alpha1 "github.com/kumahq/kuma/api/common/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backend) DeepCopyInto(out *Backend) {
	*out = *in
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(PrometheusBackend)
		(*in).DeepCopyInto(*out)
	}
	if in.OpenTelemetry != nil {
		in, out := &in.OpenTelemetry, &out.OpenTelemetry
		*out = new(OpenTelemetryBackend)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backend.
func (in *Backend) DeepCopy() *Backend {
	if in == nil {
		return nil
	}
	out := new(Backend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Conf) DeepCopyInto(out *Conf) {
	*out = *in
	if in.Sidecar != nil {
		in, out := &in.Sidecar, &out.Sidecar
		*out = new(Sidecar)
		(*in).DeepCopyInto(*out)
	}
	if in.Applications != nil {
		in, out := &in.Applications, &out.Applications
		*out = new([]Application)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Application, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = new([]Backend)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Backend, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conf.
func (in *Conf) DeepCopy() *Conf {
	if in == nil {
		return nil
	}
	out := new(Conf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshMetric) DeepCopyInto(out *MeshMetric) {
	*out = *in
	if in.TargetRef != nil {
		in, out := &in.TargetRef, &out.TargetRef
		*out = new(commonv1alpha1.TargetRef)
		(*in).DeepCopyInto(*out)
	}
	in.Default.DeepCopyInto(&out.Default)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshMetric.
func (in *MeshMetric) DeepCopy() *MeshMetric {
	if in == nil {
		return nil
	}
	out := new(MeshMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenTelemetryBackend) DeepCopyInto(out *OpenTelemetryBackend) {
	*out = *in
	if in.RefreshInterval != nil {
		in, out := &in.RefreshInterval, &out.RefreshInterval
		*out = new(v1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenTelemetryBackend.
func (in *OpenTelemetryBackend) DeepCopy() *OpenTelemetryBackend {
	if in == nil {
		return nil
	}
	out := new(OpenTelemetryBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Profile) DeepCopyInto(out *Profile) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Profile.
func (in *Profile) DeepCopy() *Profile {
	if in == nil {
		return nil
	}
	out := new(Profile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Profiles) DeepCopyInto(out *Profiles) {
	*out = *in
	if in.AppendProfiles != nil {
		in, out := &in.AppendProfiles, &out.AppendProfiles
		*out = new([]Profile)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Profile, len(*in))
			copy(*out, *in)
		}
	}
	if in.Exclude != nil {
		in, out := &in.Exclude, &out.Exclude
		*out = new([]Selector)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Selector, len(*in))
			copy(*out, *in)
		}
	}
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = new([]Selector)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Selector, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Profiles.
func (in *Profiles) DeepCopy() *Profiles {
	if in == nil {
		return nil
	}
	out := new(Profiles)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusBackend) DeepCopyInto(out *PrometheusBackend) {
	*out = *in
	if in.ClientId != nil {
		in, out := &in.ClientId, &out.ClientId
		*out = new(string)
		**out = **in
	}
	if in.Tls != nil {
		in, out := &in.Tls, &out.Tls
		*out = new(PrometheusTls)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusBackend.
func (in *PrometheusBackend) DeepCopy() *PrometheusBackend {
	if in == nil {
		return nil
	}
	out := new(PrometheusBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusTls) DeepCopyInto(out *PrometheusTls) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusTls.
func (in *PrometheusTls) DeepCopy() *PrometheusTls {
	if in == nil {
		return nil
	}
	out := new(PrometheusTls)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Selector) DeepCopyInto(out *Selector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Selector.
func (in *Selector) DeepCopy() *Selector {
	if in == nil {
		return nil
	}
	out := new(Selector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sidecar) DeepCopyInto(out *Sidecar) {
	*out = *in
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = new(Profiles)
		(*in).DeepCopyInto(*out)
	}
	if in.IncludeUnused != nil {
		in, out := &in.IncludeUnused, &out.IncludeUnused
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sidecar.
func (in *Sidecar) DeepCopy() *Sidecar {
	if in == nil {
		return nil
	}
	out := new(Sidecar)
	in.DeepCopyInto(out)
	return out
}
