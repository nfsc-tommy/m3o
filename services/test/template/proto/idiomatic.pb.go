// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/m3o/m3o/services/test/template/proto/idiomatic.proto

package idiomatic

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_62f877917a7b488e, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_62f877917a7b488e, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_62f877917a7b488e, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62f877917a7b488e, []int{3}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62f877917a7b488e, []int{4}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_62f877917a7b488e, []int{5}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_62f877917a7b488e, []int{6}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "idiomatic.Message")
	proto.RegisterType((*Request)(nil), "idiomatic.Request")
	proto.RegisterType((*Response)(nil), "idiomatic.Response")
	proto.RegisterType((*StreamingRequest)(nil), "idiomatic.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "idiomatic.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "idiomatic.Ping")
	proto.RegisterType((*Pong)(nil), "idiomatic.Pong")
}

func init() {
	proto.RegisterFile("github.com/m3o/m3o/services/test/template/proto/idiomatic.proto", fileDescriptor_62f877917a7b488e)
}

var fileDescriptor_62f877917a7b488e = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xbd, 0x4e, 0x33, 0x31,
	0x10, 0x8c, 0x95, 0x7c, 0xf9, 0xd9, 0xe6, 0x0b, 0x06, 0x21, 0x94, 0x04, 0x84, 0x5c, 0x85, 0x26,
	0x0e, 0xa1, 0x40, 0x08, 0xd1, 0x40, 0x81, 0x28, 0x90, 0xd0, 0xd1, 0xd1, 0x39, 0xc7, 0xca, 0x58,
	0xc4, 0x76, 0xb8, 0x75, 0x90, 0x78, 0x36, 0x5e, 0x0e, 0x9d, 0xef, 0x2e, 0x5c, 0x50, 0x28, 0xe8,
	0x76, 0x67, 0x66, 0x47, 0x33, 0x36, 0x5c, 0x69, 0x13, 0x5e, 0x56, 0xf3, 0x49, 0xea, 0xad, 0xb4,
	0x26, 0xcd, 0xbc, 0x24, 0xcc, 0xde, 0x4d, 0x8a, 0x24, 0x03, 0x52, 0x90, 0x01, 0xed, 0x72, 0xa1,
	0x02, 0xca, 0x65, 0xe6, 0x83, 0x97, 0xe6, 0xd9, 0x78, 0xab, 0x82, 0x49, 0x27, 0x71, 0xe7, 0xbd,
	0x35, 0x20, 0x86, 0xd0, 0xb9, 0x47, 0x22, 0xa5, 0x91, 0xf7, 0xa1, 0x49, 0xea, 0xe3, 0x80, 0x1d,
	0xb3, 0x71, 0x2f, 0xc9, 0x47, 0x71, 0x08, 0x9d, 0x04, 0xdf, 0x56, 0x48, 0x81, 0x73, 0x68, 0x39,
	0x65, 0xb1, 0x64, 0xe3, 0x2c, 0x46, 0xd0, 0x4d, 0x90, 0x96, 0xde, 0x51, 0x3c, 0xb6, 0xa4, 0xab,
	0x63, 0x4b, 0x5a, 0x8c, 0xa1, 0xff, 0x18, 0x32, 0x54, 0xd6, 0x38, 0x5d, 0xb9, 0xec, 0xc1, 0xbf,
	0xd4, 0xaf, 0x5c, 0x88, 0xba, 0x66, 0x52, 0x2c, 0xe2, 0x04, 0x76, 0x6a, 0xca, 0xd2, 0x70, 0xbb,
	0xf4, 0x08, 0x5a, 0x0f, 0xc6, 0x69, 0xbe, 0x0f, 0x6d, 0x0a, 0x99, 0x7f, 0xc5, 0x92, 0x2e, 0xb7,
	0xc8, 0xfb, 0xdf, 0xf9, 0xd9, 0x27, 0x83, 0xde, 0x5d, 0x55, 0x9e, 0x9f, 0x42, 0xeb, 0x46, 0x2d,
	0x16, 0x9c, 0x4f, 0xbe, 0x5f, 0xa8, 0x8c, 0x3a, 0xd8, 0xdd, 0xc0, 0x8a, 0x50, 0xa2, 0xc1, 0x6f,
	0xa1, 0x5d, 0x64, 0xe5, 0xc3, 0x9a, 0xe0, 0x67, 0xd1, 0xc1, 0x68, 0x3b, 0x59, 0xd9, 0x4c, 0x19,
	0x9f, 0x41, 0x37, 0x6f, 0x12, 0xd3, 0xfe, 0xaf, 0xa9, 0x73, 0x70, 0xb0, 0x01, 0x78, 0xa7, 0x45,
	0x63, 0xcc, 0xa6, 0xec, 0xfa, 0xe2, 0xe9, 0xfc, 0x4f, 0x1f, 0x7f, 0xb9, 0x76, 0x99, 0xb7, 0x23,
	0x70, 0xf6, 0x15, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xd1, 0xb5, 0x3e, 0x3a, 0x02, 0x00, 0x00,
}
