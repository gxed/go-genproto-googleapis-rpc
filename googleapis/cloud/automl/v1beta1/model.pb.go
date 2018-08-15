// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/model.proto

package automl // import "google.golang.org/genproto/googleapis/cloud/automl/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Deployment state of the model.
type Model_DeploymentState int32

const (
	// Should not be used, an un-set enum has this value by default.
	Model_DEPLOYMENT_STATE_UNSPECIFIED Model_DeploymentState = 0
	// Model is deployed.
	Model_DEPLOYED Model_DeploymentState = 1
	// Model is not deployed.
	Model_UNDEPLOYED Model_DeploymentState = 2
)

var Model_DeploymentState_name = map[int32]string{
	0: "DEPLOYMENT_STATE_UNSPECIFIED",
	1: "DEPLOYED",
	2: "UNDEPLOYED",
}
var Model_DeploymentState_value = map[string]int32{
	"DEPLOYMENT_STATE_UNSPECIFIED": 0,
	"DEPLOYED":                     1,
	"UNDEPLOYED":                   2,
}

func (x Model_DeploymentState) String() string {
	return proto.EnumName(Model_DeploymentState_name, int32(x))
}
func (Model_DeploymentState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_model_20a2aca5850ac7e3, []int{0, 0}
}

// API proto representing a trained machine learning model.
type Model struct {
	// Required.
	// The model metadata that is specific to the problem type.
	// Must match the metadata type of the dataset used to train the model.
	//
	// Types that are valid to be assigned to ModelMetadata:
	//	*Model_ImageClassificationModelMetadata
	//	*Model_TextClassificationModelMetadata
	//	*Model_TranslationModelMetadata
	ModelMetadata isModel_ModelMetadata `protobuf_oneof:"model_metadata"`
	// Output only.
	// Resource name of the model.
	// Format: `projects/{project_id}/locations/{location_id}/models/{model_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The name of the model to show in the interface. The name can be
	// up to 32 characters
	// long and can consist only of ASCII Latin letters A-Z and a-z, underscores
	// (_), and ASCII digits 0-9.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required.
	// The resource ID of the dataset used to create the model. The dataset must
	// come from the
	// same ancestor project and location.
	DatasetId string `protobuf:"bytes,3,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// Output only.
	// Timestamp when this model was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only.
	// Timestamp when this model was last updated.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,11,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Deployment state of the model.
	DeploymentState      Model_DeploymentState `protobuf:"varint,8,opt,name=deployment_state,json=deploymentState,proto3,enum=google.cloud.automl.v1beta1.Model_DeploymentState" json:"deployment_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Model) Reset()         { *m = Model{} }
func (m *Model) String() string { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()    {}
func (*Model) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_20a2aca5850ac7e3, []int{0}
}
func (m *Model) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Model.Unmarshal(m, b)
}
func (m *Model) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Model.Marshal(b, m, deterministic)
}
func (dst *Model) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Model.Merge(dst, src)
}
func (m *Model) XXX_Size() int {
	return xxx_messageInfo_Model.Size(m)
}
func (m *Model) XXX_DiscardUnknown() {
	xxx_messageInfo_Model.DiscardUnknown(m)
}

var xxx_messageInfo_Model proto.InternalMessageInfo

type isModel_ModelMetadata interface {
	isModel_ModelMetadata()
}

type Model_ImageClassificationModelMetadata struct {
	ImageClassificationModelMetadata *ImageClassificationModelMetadata `protobuf:"bytes,13,opt,name=image_classification_model_metadata,json=imageClassificationModelMetadata,proto3,oneof"`
}

type Model_TextClassificationModelMetadata struct {
	TextClassificationModelMetadata *TextClassificationModelMetadata `protobuf:"bytes,14,opt,name=text_classification_model_metadata,json=textClassificationModelMetadata,proto3,oneof"`
}

type Model_TranslationModelMetadata struct {
	TranslationModelMetadata *TranslationModelMetadata `protobuf:"bytes,15,opt,name=translation_model_metadata,json=translationModelMetadata,proto3,oneof"`
}

func (*Model_ImageClassificationModelMetadata) isModel_ModelMetadata() {}

func (*Model_TextClassificationModelMetadata) isModel_ModelMetadata() {}

func (*Model_TranslationModelMetadata) isModel_ModelMetadata() {}

func (m *Model) GetModelMetadata() isModel_ModelMetadata {
	if m != nil {
		return m.ModelMetadata
	}
	return nil
}

func (m *Model) GetImageClassificationModelMetadata() *ImageClassificationModelMetadata {
	if x, ok := m.GetModelMetadata().(*Model_ImageClassificationModelMetadata); ok {
		return x.ImageClassificationModelMetadata
	}
	return nil
}

func (m *Model) GetTextClassificationModelMetadata() *TextClassificationModelMetadata {
	if x, ok := m.GetModelMetadata().(*Model_TextClassificationModelMetadata); ok {
		return x.TextClassificationModelMetadata
	}
	return nil
}

func (m *Model) GetTranslationModelMetadata() *TranslationModelMetadata {
	if x, ok := m.GetModelMetadata().(*Model_TranslationModelMetadata); ok {
		return x.TranslationModelMetadata
	}
	return nil
}

func (m *Model) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Model) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Model) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *Model) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Model) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Model) GetDeploymentState() Model_DeploymentState {
	if m != nil {
		return m.DeploymentState
	}
	return Model_DEPLOYMENT_STATE_UNSPECIFIED
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Model) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Model_OneofMarshaler, _Model_OneofUnmarshaler, _Model_OneofSizer, []interface{}{
		(*Model_ImageClassificationModelMetadata)(nil),
		(*Model_TextClassificationModelMetadata)(nil),
		(*Model_TranslationModelMetadata)(nil),
	}
}

func _Model_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Model)
	// model_metadata
	switch x := m.ModelMetadata.(type) {
	case *Model_ImageClassificationModelMetadata:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ImageClassificationModelMetadata); err != nil {
			return err
		}
	case *Model_TextClassificationModelMetadata:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TextClassificationModelMetadata); err != nil {
			return err
		}
	case *Model_TranslationModelMetadata:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TranslationModelMetadata); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Model.ModelMetadata has unexpected type %T", x)
	}
	return nil
}

func _Model_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Model)
	switch tag {
	case 13: // model_metadata.image_classification_model_metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ImageClassificationModelMetadata)
		err := b.DecodeMessage(msg)
		m.ModelMetadata = &Model_ImageClassificationModelMetadata{msg}
		return true, err
	case 14: // model_metadata.text_classification_model_metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TextClassificationModelMetadata)
		err := b.DecodeMessage(msg)
		m.ModelMetadata = &Model_TextClassificationModelMetadata{msg}
		return true, err
	case 15: // model_metadata.translation_model_metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TranslationModelMetadata)
		err := b.DecodeMessage(msg)
		m.ModelMetadata = &Model_TranslationModelMetadata{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Model_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Model)
	// model_metadata
	switch x := m.ModelMetadata.(type) {
	case *Model_ImageClassificationModelMetadata:
		s := proto.Size(x.ImageClassificationModelMetadata)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Model_TextClassificationModelMetadata:
		s := proto.Size(x.TextClassificationModelMetadata)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Model_TranslationModelMetadata:
		s := proto.Size(x.TranslationModelMetadata)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Model)(nil), "google.cloud.automl.v1beta1.Model")
	proto.RegisterEnum("google.cloud.automl.v1beta1.Model_DeploymentState", Model_DeploymentState_name, Model_DeploymentState_value)
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/model.proto", fileDescriptor_model_20a2aca5850ac7e3)
}

var fileDescriptor_model_20a2aca5850ac7e3 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x97, 0x01, 0x63, 0x3b, 0x1d, 0x6d, 0xe5, 0xab, 0xa8, 0x0c, 0x35, 0x14, 0x09, 0x7a,
	0x43, 0xa2, 0x15, 0x71, 0x35, 0xb8, 0xd8, 0xd6, 0x20, 0x2a, 0xd1, 0x52, 0xda, 0xec, 0x02, 0x24,
	0x14, 0xb9, 0x89, 0x1b, 0x59, 0x4a, 0xe2, 0xa8, 0x39, 0x41, 0xdb, 0x33, 0x20, 0xf1, 0x2e, 0xbc,
	0x21, 0x8a, 0xed, 0x16, 0x5a, 0x31, 0xf7, 0xae, 0x3e, 0xe7, 0xfb, 0x8f, 0x7f, 0xff, 0x3d, 0x81,
	0x57, 0x89, 0x10, 0x49, 0xca, 0xbc, 0x28, 0x15, 0x55, 0xec, 0xd1, 0x0a, 0x45, 0x96, 0x7a, 0x3f,
	0xce, 0x17, 0x0c, 0xe9, 0xb9, 0x97, 0x89, 0x98, 0xa5, 0x6e, 0xb1, 0x12, 0x28, 0xc8, 0x53, 0x05,
	0xba, 0x12, 0x74, 0x15, 0xe8, 0x6a, 0xb0, 0x73, 0xa6, 0xa7, 0xd0, 0x82, 0x7b, 0x34, 0xcf, 0x05,
	0x52, 0xe4, 0x22, 0x2f, 0x95, 0xb4, 0x63, 0xbc, 0x83, 0x67, 0x34, 0x61, 0x1a, 0x7c, 0x69, 0x02,
	0x91, 0xdd, 0xa2, 0xe6, 0x5e, 0x1b, 0xb9, 0x15, 0xcd, 0xcb, 0x54, 0x1a, 0xd0, 0x78, 0x57, 0xe3,
	0xf2, 0xb4, 0xa8, 0x96, 0x1e, 0xf2, 0x8c, 0x95, 0x48, 0xb3, 0x42, 0x01, 0xbd, 0xdf, 0x47, 0xf0,
	0x68, 0x5c, 0xbf, 0x95, 0xfc, 0xb2, 0xe0, 0x85, 0x74, 0x14, 0x46, 0x29, 0x2d, 0x4b, 0xbe, 0xe4,
	0x91, 0x9c, 0x14, 0xca, 0x28, 0xc2, 0x8c, 0x21, 0x8d, 0x29, 0x52, 0xfb, 0x89, 0x63, 0xf5, 0x1b,
	0x83, 0xf7, 0xae, 0x21, 0x14, 0x77, 0x54, 0xcf, 0xb9, 0xde, 0x1a, 0x23, 0x2f, 0x19, 0xeb, 0x21,
	0x1f, 0x0f, 0x66, 0x0e, 0xdf, 0xc3, 0x90, 0x9f, 0x16, 0xf4, 0xea, 0x97, 0xef, 0xf1, 0xd3, 0x94,
	0x7e, 0xde, 0x19, 0xfd, 0x04, 0xec, 0x16, 0xcd, 0x76, 0xba, 0x68, 0x46, 0x48, 0x05, 0x9d, 0x7f,
	0xe2, 0xdd, 0x35, 0xd1, 0x92, 0x26, 0xde, 0x9a, 0x4d, 0xfc, 0x95, 0xef, 0xde, 0x6e, 0xe3, 0x3d,
	0x3d, 0x42, 0xe0, 0x61, 0x4e, 0x33, 0x66, 0x5b, 0x8e, 0xd5, 0x3f, 0x99, 0xc9, 0xdf, 0xe4, 0x39,
	0x9c, 0xc6, 0xbc, 0x2c, 0x52, 0x7a, 0x17, 0xca, 0xde, 0xa1, 0xec, 0x35, 0x74, 0x6d, 0x52, 0x23,
	0xcf, 0x00, 0x6a, 0x79, 0xc9, 0x30, 0xe4, 0xb1, 0xfd, 0x40, 0x02, 0x27, 0xba, 0x32, 0x8a, 0xc9,
	0x05, 0x34, 0xa2, 0x15, 0xa3, 0xc8, 0xc2, 0x7a, 0x1f, 0xec, 0xc7, 0xd2, 0x7d, 0x67, 0xed, 0x7e,
	0xbd, 0x2c, 0x6e, 0xb0, 0x5e, 0x96, 0x19, 0x28, 0xbc, 0x2e, 0xd4, 0xe2, 0xaa, 0x88, 0x37, 0xe2,
	0xc6, 0x7e, 0xb1, 0xc2, 0xa5, 0xf8, 0x3b, 0xb4, 0x63, 0x56, 0xa4, 0xe2, 0x2e, 0x63, 0x39, 0x86,
	0x25, 0x52, 0x64, 0xf6, 0xb1, 0x63, 0xf5, 0x9b, 0x83, 0x81, 0x31, 0x3c, 0x99, 0x8a, 0x3b, 0xdc,
	0x48, 0xe7, 0xb5, 0x72, 0xd6, 0x8a, 0xb7, 0x0b, 0xbd, 0x2f, 0xd0, 0xda, 0x61, 0x88, 0x03, 0x67,
	0x43, 0x7f, 0xfa, 0xe9, 0xf3, 0xd7, 0xb1, 0x3f, 0x09, 0xc2, 0x79, 0x70, 0x19, 0xf8, 0xe1, 0xcd,
	0x64, 0x3e, 0xf5, 0xaf, 0x47, 0x1f, 0x46, 0xfe, 0xb0, 0x7d, 0x40, 0x4e, 0xe1, 0x58, 0x11, 0xfe,
	0xb0, 0x6d, 0x91, 0x26, 0xc0, 0xcd, 0x64, 0x73, 0x3e, 0xbc, 0x6a, 0x43, 0x73, 0xfb, 0xcf, 0xbe,
	0x5a, 0x42, 0x37, 0x12, 0x99, 0xc9, 0xee, 0xd4, 0xfa, 0x76, 0xa9, 0xdb, 0x89, 0x48, 0x69, 0x9e,
	0xb8, 0x62, 0x95, 0x78, 0x09, 0xcb, 0x65, 0x3a, 0x9e, 0x6a, 0xd1, 0x82, 0x97, 0xff, 0xfd, 0x8e,
	0x2f, 0xd4, 0x71, 0x71, 0x24, 0xe9, 0x37, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x10, 0xe2,
	0xaa, 0xa9, 0x04, 0x00, 0x00,
}
