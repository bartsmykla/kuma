// Generated by tools/policy-gen.
// Run "make generate" to update this file.

// nolint:whitespace
package v1alpha1

import (
	_ "embed"
	"errors"
	"fmt"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/schema"
	"k8s.io/kube-openapi/pkg/validation/strfmt"
	"k8s.io/kube-openapi/pkg/validation/validate"
	"sigs.k8s.io/yaml"

	"github.com/kumahq/kuma/pkg/core/resources/model"
)

//go:embed schema.yaml
var rawSchema []byte

func init() {
	var structuralSchema *schema.Structural
	var v1JsonSchemaProps *apiextensionsv1.JSONSchemaProps
	var validator *validate.SchemaValidator
	if rawSchema != nil {
		if err := yaml.Unmarshal(rawSchema, &v1JsonSchemaProps); err != nil {
			panic(err)
		}
		var jsonSchemaProps apiextensions.JSONSchemaProps
		err := apiextensionsv1.Convert_v1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(v1JsonSchemaProps, &jsonSchemaProps, nil)
		if err != nil {
			panic(err)
		}
		structuralSchema, err = schema.NewStructural(&jsonSchemaProps)
		if err != nil {
			panic(err)
		}
		schemaObject := structuralSchema.ToKubeOpenAPI()
		validator = validate.NewSchemaValidator(schemaObject, nil, "", strfmt.Default)
	}
	rawSchema = nil
	MeshFaultInjectionResourceTypeDescriptor.Validator = validator
	MeshFaultInjectionResourceTypeDescriptor.StructuralSchema = structuralSchema
}

const (
	MeshFaultInjectionType model.ResourceType = "MeshFaultInjection"
)

var _ model.Resource = &MeshFaultInjectionResource{}

type MeshFaultInjectionResource struct {
	Meta model.ResourceMeta
	Spec *MeshFaultInjection
}

func NewMeshFaultInjectionResource() *MeshFaultInjectionResource {
	return &MeshFaultInjectionResource{
		Spec: &MeshFaultInjection{},
	}
}

func (t *MeshFaultInjectionResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *MeshFaultInjectionResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *MeshFaultInjectionResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *MeshFaultInjectionResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*MeshFaultInjection)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &MeshFaultInjection{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *MeshFaultInjectionResource) GetStatus() model.ResourceStatus {
	return nil
}

func (t *MeshFaultInjectionResource) SetStatus(_ model.ResourceStatus) error {
	return errors.New("status not supported")
}

func (t *MeshFaultInjectionResource) Descriptor() model.ResourceTypeDescriptor {
	return MeshFaultInjectionResourceTypeDescriptor
}

func (t *MeshFaultInjectionResource) Validate() error {
	if v, ok := interface{}(t).(interface{ validate() error }); !ok {
		return nil
	} else {
		return v.validate()
	}
}

var _ model.ResourceList = &MeshFaultInjectionResourceList{}

type MeshFaultInjectionResourceList struct {
	Items      []*MeshFaultInjectionResource
	Pagination model.Pagination
}

func (l *MeshFaultInjectionResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *MeshFaultInjectionResourceList) GetItemType() model.ResourceType {
	return MeshFaultInjectionType
}

func (l *MeshFaultInjectionResourceList) NewItem() model.Resource {
	return NewMeshFaultInjectionResource()
}

func (l *MeshFaultInjectionResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*MeshFaultInjectionResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*MeshFaultInjectionResource)(nil), r)
	}
}

func (l *MeshFaultInjectionResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *MeshFaultInjectionResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var MeshFaultInjectionResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                         MeshFaultInjectionType,
	Resource:                     NewMeshFaultInjectionResource(),
	ResourceList:                 &MeshFaultInjectionResourceList{},
	Scope:                        model.ScopeMesh,
	KDSFlags:                     model.GlobalToZonesFlag | model.ZoneToGlobalFlag | model.SyncedAcrossZonesFlag,
	WsPath:                       "meshfaultinjections",
	KumactlArg:                   "meshfaultinjection",
	KumactlListArg:               "meshfaultinjections",
	AllowToInspect:               true,
	IsPolicy:                     true,
	IsDestination:                false,
	IsExperimental:               false,
	SingularDisplayName:          "Mesh Fault Injection",
	PluralDisplayName:            "Mesh Fault Injections",
	IsPluginOriginated:           true,
	IsTargetRefBased:             true,
	HasToTargetRef:               true,
	HasFromTargetRef:             true,
	HasRulesTargetRef:            false,
	HasStatus:                    false,
	AllowedOnSystemNamespaceOnly: false,
	IsReferenceableInTo:          false,
	ShortName:                    "mfi",
	IsFromAsRules:                false,
}
