// Code generated by protoc-gen-go.
// source: x-samplecompanytwo.proto
// DO NOT EDIT!

/*
Package samplecompanytwo is a generated protocol buffer package.

It is generated from these files:
	x-samplecompanytwo.proto

It has these top-level messages:
	SampleCompanyTwoBook
	SampleCompanyTwoShelve
*/
package samplecompanytwo

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

type SampleCompanyTwoBook struct {
	Code    int64 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message int64 `protobuf:"varint,2,opt,name=message" json:"message,omitempty"`
}

func (m *SampleCompanyTwoBook) Reset()                    { *m = SampleCompanyTwoBook{} }
func (m *SampleCompanyTwoBook) String() string            { return proto.CompactTextString(m) }
func (*SampleCompanyTwoBook) ProtoMessage()               {}
func (*SampleCompanyTwoBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SampleCompanyTwoBook) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SampleCompanyTwoBook) GetMessage() int64 {
	if m != nil {
		return m.Message
	}
	return 0
}

type SampleCompanyTwoShelve struct {
	Foo1 int64 `protobuf:"varint,1,opt,name=foo1" json:"foo1,omitempty"`
	Bar  int64 `protobuf:"varint,2,opt,name=bar" json:"bar,omitempty"`
}

func (m *SampleCompanyTwoShelve) Reset()                    { *m = SampleCompanyTwoShelve{} }
func (m *SampleCompanyTwoShelve) String() string            { return proto.CompactTextString(m) }
func (*SampleCompanyTwoShelve) ProtoMessage()               {}
func (*SampleCompanyTwoShelve) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SampleCompanyTwoShelve) GetFoo1() int64 {
	if m != nil {
		return m.Foo1
	}
	return 0
}

func (m *SampleCompanyTwoShelve) GetBar() int64 {
	if m != nil {
		return m.Bar
	}
	return 0
}

func init() {
	proto.RegisterType((*SampleCompanyTwoBook)(nil), "samplecompanytwo.SampleCompanyTwoBook")
	proto.RegisterType((*SampleCompanyTwoShelve)(nil), "samplecompanytwo.SampleCompanyTwoShelve")
}

func init() { proto.RegisterFile("x-samplecompanytwo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0xa8, 0xd0, 0x2d, 0x4e,
	0xcc, 0x2d, 0xc8, 0x49, 0x4d, 0xce, 0xcf, 0x2d, 0x48, 0xcc, 0xab, 0x2c, 0x29, 0xcf, 0xd7, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x40, 0x17, 0x57, 0x72, 0xe1, 0x12, 0x09, 0x06, 0x8b, 0x39,
	0x43, 0xc4, 0x42, 0xca, 0xf3, 0x9d, 0xf2, 0xf3, 0xb3, 0x85, 0x84, 0xb8, 0x58, 0x92, 0xf3, 0x53,
	0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0xc0, 0x6c, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4,
	0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x26, 0xb0, 0x30, 0x8c, 0xab, 0x64, 0xc7, 0x25, 0x86, 0x6e,
	0x4a, 0x70, 0x46, 0x6a, 0x4e, 0x59, 0x2a, 0xc8, 0x9c, 0xb4, 0xfc, 0x7c, 0x43, 0x98, 0x39, 0x20,
	0xb6, 0x90, 0x00, 0x17, 0x73, 0x52, 0x62, 0x11, 0xd4, 0x0c, 0x10, 0xd3, 0x49, 0xc2, 0x49, 0x24,
	0x2c, 0x35, 0x2f, 0x25, 0xbf, 0xc8, 0xb5, 0xa2, 0x24, 0x35, 0xaf, 0x38, 0x33, 0x3f, 0x2f, 0x00,
	0xe4, 0xde, 0x00, 0xc6, 0x24, 0x36, 0xb0, 0xc3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x83,
	0x29, 0x59, 0x7f, 0xd4, 0x00, 0x00, 0x00,
}