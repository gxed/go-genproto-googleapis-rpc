// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/model_evaluation.proto

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

// Evaluation results of a model.
type ModelEvaluation struct {
	// Output only. Problem type specific evaluation metrics.
	//
	// Types that are valid to be assigned to Metrics:
	//	*ModelEvaluation_ClassificationEvaluationMetrics
	//	*ModelEvaluation_TranslationEvaluationMetrics
	Metrics isModelEvaluation_Metrics `protobuf_oneof:"metrics"`
	// Output only.
	// Resource name of the model evaluation.
	// Format:
	//
	// `projects/{project_id}/locations/{location_id}/models/{model_id}/modelEvaluations/{model_evaluation_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only.
	// The ID of the annotation spec that the model evaluation applies to. The
	// ID is empty for overall model evaluation.
	// NOTE: Currently there is no way to obtain the display_name of the
	// annotation spec from its ID. To see the display_names, review the model
	// evaluations in the UI.
	AnnotationSpecId string `protobuf:"bytes,2,opt,name=annotation_spec_id,json=annotationSpecId,proto3" json:"annotation_spec_id,omitempty"`
	// Output only.
	// Timestamp when this model evaluation was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The number of examples used for model evaluation.
	EvaluatedExampleCount int32    `protobuf:"varint,6,opt,name=evaluated_example_count,json=evaluatedExampleCount,proto3" json:"evaluated_example_count,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ModelEvaluation) Reset()         { *m = ModelEvaluation{} }
func (m *ModelEvaluation) String() string { return proto.CompactTextString(m) }
func (*ModelEvaluation) ProtoMessage()    {}
func (*ModelEvaluation) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_evaluation_a77b44488aa864eb, []int{0}
}
func (m *ModelEvaluation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelEvaluation.Unmarshal(m, b)
}
func (m *ModelEvaluation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelEvaluation.Marshal(b, m, deterministic)
}
func (dst *ModelEvaluation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelEvaluation.Merge(dst, src)
}
func (m *ModelEvaluation) XXX_Size() int {
	return xxx_messageInfo_ModelEvaluation.Size(m)
}
func (m *ModelEvaluation) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelEvaluation.DiscardUnknown(m)
}

var xxx_messageInfo_ModelEvaluation proto.InternalMessageInfo

type isModelEvaluation_Metrics interface {
	isModelEvaluation_Metrics()
}

type ModelEvaluation_ClassificationEvaluationMetrics struct {
	ClassificationEvaluationMetrics *ClassificationEvaluationMetrics `protobuf:"bytes,8,opt,name=classification_evaluation_metrics,json=classificationEvaluationMetrics,proto3,oneof"`
}

type ModelEvaluation_TranslationEvaluationMetrics struct {
	TranslationEvaluationMetrics *TranslationEvaluationMetrics `protobuf:"bytes,9,opt,name=translation_evaluation_metrics,json=translationEvaluationMetrics,proto3,oneof"`
}

func (*ModelEvaluation_ClassificationEvaluationMetrics) isModelEvaluation_Metrics() {}

func (*ModelEvaluation_TranslationEvaluationMetrics) isModelEvaluation_Metrics() {}

func (m *ModelEvaluation) GetMetrics() isModelEvaluation_Metrics {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *ModelEvaluation) GetClassificationEvaluationMetrics() *ClassificationEvaluationMetrics {
	if x, ok := m.GetMetrics().(*ModelEvaluation_ClassificationEvaluationMetrics); ok {
		return x.ClassificationEvaluationMetrics
	}
	return nil
}

func (m *ModelEvaluation) GetTranslationEvaluationMetrics() *TranslationEvaluationMetrics {
	if x, ok := m.GetMetrics().(*ModelEvaluation_TranslationEvaluationMetrics); ok {
		return x.TranslationEvaluationMetrics
	}
	return nil
}

func (m *ModelEvaluation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModelEvaluation) GetAnnotationSpecId() string {
	if m != nil {
		return m.AnnotationSpecId
	}
	return ""
}

func (m *ModelEvaluation) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ModelEvaluation) GetEvaluatedExampleCount() int32 {
	if m != nil {
		return m.EvaluatedExampleCount
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ModelEvaluation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ModelEvaluation_OneofMarshaler, _ModelEvaluation_OneofUnmarshaler, _ModelEvaluation_OneofSizer, []interface{}{
		(*ModelEvaluation_ClassificationEvaluationMetrics)(nil),
		(*ModelEvaluation_TranslationEvaluationMetrics)(nil),
	}
}

func _ModelEvaluation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ModelEvaluation)
	// metrics
	switch x := m.Metrics.(type) {
	case *ModelEvaluation_ClassificationEvaluationMetrics:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClassificationEvaluationMetrics); err != nil {
			return err
		}
	case *ModelEvaluation_TranslationEvaluationMetrics:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TranslationEvaluationMetrics); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ModelEvaluation.Metrics has unexpected type %T", x)
	}
	return nil
}

func _ModelEvaluation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ModelEvaluation)
	switch tag {
	case 8: // metrics.classification_evaluation_metrics
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClassificationEvaluationMetrics)
		err := b.DecodeMessage(msg)
		m.Metrics = &ModelEvaluation_ClassificationEvaluationMetrics{msg}
		return true, err
	case 9: // metrics.translation_evaluation_metrics
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TranslationEvaluationMetrics)
		err := b.DecodeMessage(msg)
		m.Metrics = &ModelEvaluation_TranslationEvaluationMetrics{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ModelEvaluation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ModelEvaluation)
	// metrics
	switch x := m.Metrics.(type) {
	case *ModelEvaluation_ClassificationEvaluationMetrics:
		s := proto.Size(x.ClassificationEvaluationMetrics)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ModelEvaluation_TranslationEvaluationMetrics:
		s := proto.Size(x.TranslationEvaluationMetrics)
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
	proto.RegisterType((*ModelEvaluation)(nil), "google.cloud.automl.v1beta1.ModelEvaluation")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/model_evaluation.proto", fileDescriptor_model_evaluation_a77b44488aa864eb)
}

var fileDescriptor_model_evaluation_a77b44488aa864eb = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4f, 0x6b, 0xdb, 0x30,
	0x1c, 0x9d, 0xb7, 0x25, 0x5b, 0x94, 0xc3, 0x86, 0x60, 0xcc, 0x78, 0x61, 0xc9, 0x76, 0xca, 0x61,
	0x93, 0x96, 0x0c, 0x06, 0x23, 0xbd, 0x34, 0x21, 0xd0, 0x1e, 0x02, 0xc5, 0xcd, 0xa9, 0x17, 0xa3,
	0xc8, 0x8a, 0x11, 0xe8, 0x8f, 0xb1, 0xe4, 0xd0, 0x6b, 0xaf, 0xfd, 0x9c, 0xfd, 0x20, 0xc5, 0x92,
	0x13, 0x3b, 0x10, 0x7c, 0xb3, 0xfc, 0xde, 0xfb, 0xfd, 0xde, 0xd3, 0x13, 0x98, 0x67, 0x5a, 0x67,
	0x82, 0x61, 0x2a, 0x74, 0x99, 0x62, 0x52, 0x5a, 0x2d, 0x05, 0x3e, 0xcc, 0x76, 0xcc, 0x92, 0x19,
	0x96, 0x3a, 0x65, 0x22, 0x61, 0x07, 0x22, 0x4a, 0x62, 0xb9, 0x56, 0x28, 0x2f, 0xb4, 0xd5, 0xf0,
	0x9b, 0xd7, 0x20, 0xa7, 0x41, 0x5e, 0x83, 0x6a, 0x4d, 0x34, 0xaa, 0x07, 0x92, 0x9c, 0x63, 0xa2,
	0x94, 0xb6, 0x4e, 0x69, 0xbc, 0x34, 0xfa, 0xd3, 0xb5, 0x8e, 0x0a, 0x62, 0x0c, 0xdf, 0x73, 0xda,
	0x5a, 0x16, 0xfd, 0xee, 0x52, 0xd8, 0x82, 0x28, 0x23, 0xda, 0xf4, 0x71, 0x4d, 0x77, 0xa7, 0x5d,
	0xb9, 0xc7, 0x96, 0x4b, 0x66, 0x2c, 0x91, 0xb9, 0x27, 0xfc, 0x7c, 0x79, 0x07, 0x3e, 0x6d, 0xaa,
	0x5c, 0xeb, 0x53, 0x2c, 0xf8, 0x1c, 0x80, 0x1f, 0xe7, 0xcb, 0x5b, 0xa1, 0x13, 0xc9, 0x6c, 0xc1,
	0xa9, 0x09, 0x3f, 0x4e, 0x82, 0xe9, 0x70, 0x7e, 0x85, 0x3a, 0xd2, 0xa3, 0xd5, 0xd9, 0x94, 0x66,
	0xc5, 0xc6, 0xcf, 0xb8, 0x79, 0x13, 0x8f, 0x69, 0x37, 0x05, 0x3e, 0x05, 0xe0, 0x7b, 0x2b, 0xd7,
	0x25, 0x27, 0x03, 0xe7, 0xe4, 0x7f, 0xa7, 0x93, 0x6d, 0x33, 0xe2, 0x92, 0x8d, 0x91, 0xed, 0xc0,
	0x21, 0x04, 0xef, 0x15, 0x91, 0x2c, 0x0c, 0x26, 0xc1, 0x74, 0x10, 0xbb, 0x6f, 0xf8, 0x0b, 0xc0,
	0xa6, 0xcf, 0xc4, 0xe4, 0x8c, 0x26, 0x3c, 0x0d, 0xdf, 0x3a, 0xc6, 0xe7, 0x06, 0xb9, 0xcf, 0x19,
	0xbd, 0x4d, 0xe1, 0x02, 0x0c, 0x69, 0xc1, 0x88, 0x65, 0x49, 0x55, 0x40, 0xd8, 0x73, 0x8e, 0xa3,
	0xa3, 0xe3, 0x63, 0x3b, 0x68, 0x7b, 0x6c, 0x27, 0x06, 0x9e, 0x5e, 0xfd, 0x80, 0xff, 0xc0, 0xd7,
	0x3a, 0x35, 0x4b, 0x13, 0xf6, 0x48, 0x64, 0x2e, 0x58, 0x42, 0x75, 0xa9, 0x6c, 0xd8, 0x9f, 0x04,
	0xd3, 0x5e, 0xfc, 0xe5, 0x04, 0xaf, 0x3d, 0xba, 0xaa, 0xc0, 0xe5, 0x00, 0x7c, 0xa8, 0xaf, 0x68,
	0xb9, 0x07, 0x63, 0xaa, 0x65, 0xd7, 0x0d, 0xdd, 0x05, 0x0f, 0xd7, 0x35, 0x9c, 0x69, 0x41, 0x54,
	0x86, 0x74, 0x91, 0xe1, 0x8c, 0x29, 0x67, 0x0e, 0x7b, 0x88, 0xe4, 0xdc, 0x5c, 0x7c, 0x7a, 0x0b,
	0x7f, 0xdc, 0xf5, 0x1d, 0xfb, 0xef, 0x6b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x78, 0x76, 0xc2,
	0x48, 0x03, 0x00, 0x00,
}
