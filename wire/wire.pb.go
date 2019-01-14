// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wire/wire.proto

package wire

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TracerState struct {
	TraceId              *string           `protobuf:"bytes,1,req,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	SpanId               *string           `protobuf:"bytes,2,req,name=span_id,json=spanId" json:"span_id,omitempty"`
	Sampled              *bool             `protobuf:"varint,3,req,name=sampled" json:"sampled,omitempty"`
	BaggageItems         map[string]string `protobuf:"bytes,4,rep,name=baggage_items,json=baggageItems" json:"baggage_items,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TracerState) Reset()         { *m = TracerState{} }
func (m *TracerState) String() string { return proto.CompactTextString(m) }
func (*TracerState) ProtoMessage()    {}
func (*TracerState) Descriptor() ([]byte, []int) {
	return fileDescriptor_wire_a709bc5ff387907d, []int{0}
}
func (m *TracerState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TracerState.Unmarshal(m, b)
}
func (m *TracerState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TracerState.Marshal(b, m, deterministic)
}
func (dst *TracerState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TracerState.Merge(dst, src)
}
func (m *TracerState) XXX_Size() int {
	return xxx_messageInfo_TracerState.Size(m)
}
func (m *TracerState) XXX_DiscardUnknown() {
	xxx_messageInfo_TracerState.DiscardUnknown(m)
}

var xxx_messageInfo_TracerState proto.InternalMessageInfo

func (m *TracerState) GetTraceId() string {
	if m != nil && m.TraceId != nil {
		return *m.TraceId
	}
	return ""
}

func (m *TracerState) GetSpanId() string {
	if m != nil && m.SpanId != nil {
		return *m.SpanId
	}
	return ""
}

func (m *TracerState) GetSampled() bool {
	if m != nil && m.Sampled != nil {
		return *m.Sampled
	}
	return false
}

func (m *TracerState) GetBaggageItems() map[string]string {
	if m != nil {
		return m.BaggageItems
	}
	return nil
}

func init() {
	proto.RegisterType((*TracerState)(nil), "wavefront_go.wire.TracerState")
	proto.RegisterMapType((map[string]string)(nil), "wavefront_go.wire.TracerState.BaggageItemsEntry")
}

func init() { proto.RegisterFile("wire/wire.proto", fileDescriptor_wire_a709bc5ff387907d) }

var fileDescriptor_wire_a709bc5ff387907d = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xcf, 0x2c, 0x4a,
	0xd5, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x82, 0xe5, 0x89, 0x65, 0xa9, 0x69,
	0x45, 0xf9, 0x79, 0x25, 0xf1, 0xe9, 0xf9, 0x7a, 0x20, 0x09, 0xa5, 0xcf, 0x8c, 0x5c, 0xdc, 0x21,
	0x45, 0x89, 0xc9, 0xa9, 0x45, 0xc1, 0x25, 0x89, 0x25, 0xa9, 0x42, 0x92, 0x5c, 0x1c, 0x25, 0x20,
	0x6e, 0x7c, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0x93, 0x06, 0x67, 0x10, 0x3b, 0x98, 0xef, 0x99, 0x22,
	0x24, 0xce, 0xc5, 0x5e, 0x5c, 0x90, 0x98, 0x07, 0x92, 0x61, 0x02, 0xcb, 0xb0, 0x81, 0xb8, 0x9e,
	0x29, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0x89, 0xb9, 0x05, 0x39, 0xa9, 0x29, 0x12, 0xcc, 0x0a, 0x4c,
	0x1a, 0x1c, 0x41, 0x30, 0xae, 0x50, 0x28, 0x17, 0x6f, 0x52, 0x62, 0x7a, 0x7a, 0x62, 0x7a, 0x6a,
	0x7c, 0x66, 0x49, 0x6a, 0x6e, 0xb1, 0x04, 0x8b, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x81, 0x1e, 0x86,
	0x43, 0xf4, 0x90, 0x1c, 0xa1, 0xe7, 0x04, 0xd1, 0xe3, 0x09, 0xd2, 0xe2, 0x9a, 0x57, 0x52, 0x54,
	0x19, 0xc4, 0x93, 0x84, 0x24, 0x24, 0x65, 0xcf, 0x25, 0x88, 0xa1, 0x44, 0x48, 0x80, 0x8b, 0x39,
	0x3b, 0xb5, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc4, 0x14, 0x12, 0xe1, 0x62, 0x2d,
	0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x02, 0x8b, 0x41, 0x38, 0x56, 0x4c, 0x16, 0x8c, 0x4e, 0x6c,
	0x51, 0x2c, 0x20, 0x4b, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x35, 0x84, 0x93, 0xbd, 0x22, 0x01,
	0x00, 0x00,
}