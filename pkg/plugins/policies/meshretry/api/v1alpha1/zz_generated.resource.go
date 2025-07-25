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
	MeshRetryResourceTypeDescriptor.Validator = validator
	MeshRetryResourceTypeDescriptor.StructuralSchema = structuralSchema
}

const (
	MeshRetryType model.ResourceType = "MeshRetry"
)

var _ model.Resource = &MeshRetryResource{}

type MeshRetryResource struct {
	Meta model.ResourceMeta
	Spec *MeshRetry
}

func NewMeshRetryResource() *MeshRetryResource {
	return &MeshRetryResource{
		Spec: &MeshRetry{},
	}
}

func (t *MeshRetryResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *MeshRetryResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *MeshRetryResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *MeshRetryResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*MeshRetry)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &MeshRetry{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *MeshRetryResource) GetStatus() model.ResourceStatus {
	return nil
}

func (t *MeshRetryResource) SetStatus(_ model.ResourceStatus) error {
	return errors.New("status not supported")
}

func (t *MeshRetryResource) Descriptor() model.ResourceTypeDescriptor {
	return MeshRetryResourceTypeDescriptor
}

func (t *MeshRetryResource) Validate() error {
	if v, ok := interface{}(t).(interface{ validate() error }); !ok {
		return nil
	} else {
		return v.validate()
	}
}

var _ model.ResourceList = &MeshRetryResourceList{}

type MeshRetryResourceList struct {
	Items      []*MeshRetryResource
	Pagination model.Pagination
}

func (l *MeshRetryResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *MeshRetryResourceList) GetItemType() model.ResourceType {
	return MeshRetryType
}

func (l *MeshRetryResourceList) NewItem() model.Resource {
	return NewMeshRetryResource()
}

func (l *MeshRetryResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*MeshRetryResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*MeshRetryResource)(nil), r)
	}
}

func (l *MeshRetryResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *MeshRetryResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var MeshRetryResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                         MeshRetryType,
	Resource:                     NewMeshRetryResource(),
	ResourceList:                 &MeshRetryResourceList{},
	Scope:                        model.ScopeMesh,
	KDSFlags:                     model.GlobalToZonesFlag | model.ZoneToGlobalFlag | model.SyncedAcrossZonesFlag,
	WsPath:                       "meshretries",
	KumactlArg:                   "meshretry",
	KumactlListArg:               "meshretries",
	AllowToInspect:               true,
	IsPolicy:                     true,
	IsDestination:                false,
	IsExperimental:               false,
	SingularDisplayName:          "Mesh Retry",
	PluralDisplayName:            "Mesh Retries",
	IsPluginOriginated:           true,
	IsTargetRefBased:             true,
	HasToTargetRef:               true,
	HasFromTargetRef:             false,
	HasRulesTargetRef:            false,
	HasStatus:                    false,
	AllowedOnSystemNamespaceOnly: false,
	IsReferenceableInTo:          false,
	ShortName:                    "mr",
	IsFromAsRules:                false,
}
