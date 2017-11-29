// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MasterClient.proto

/*
Package masterclient is a generated protocol buffer package.

It is generated from these files:
	MasterClient.proto

It has these top-level messages:
	MasterClient
*/
package masterclient

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

type MasterClient struct {
	ClientID    uint32 `protobuf:"varint,1,opt,name=clientID" json:"clientID,omitempty"`
	Application string `protobuf:"bytes,2,opt,name=application" json:"application,omitempty"`
	Dataset     []byte `protobuf:"bytes,3,opt,name=dataset,proto3" json:"dataset,omitempty"`
	Success     bool   `protobuf:"varint,4,opt,name=success" json:"success,omitempty"`
	Result      []byte `protobuf:"bytes,5,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *MasterClient) Reset()                    { *m = MasterClient{} }
func (m *MasterClient) String() string            { return proto.CompactTextString(m) }
func (*MasterClient) ProtoMessage()               {}
func (*MasterClient) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MasterClient) GetClientID() uint32 {
	if m != nil {
		return m.ClientID
	}
	return 0
}

func (m *MasterClient) GetApplication() string {
	if m != nil {
		return m.Application
	}
	return ""
}

func (m *MasterClient) GetDataset() []byte {
	if m != nil {
		return m.Dataset
	}
	return nil
}

func (m *MasterClient) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *MasterClient) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*MasterClient)(nil), "masterclient.MasterClient")
}

func init() { proto.RegisterFile("MasterClient.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xf2, 0x4d, 0x2c, 0x2e,
	0x49, 0x2d, 0x72, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0xc9, 0x05, 0x8b, 0x25, 0x83, 0xc5, 0x94, 0x66, 0x30, 0x72, 0xf1, 0x20, 0x2b, 0x12, 0x92, 0xe2,
	0xe2, 0x80, 0x48, 0x79, 0xba, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x06, 0xc1, 0xf9, 0x42, 0x0a,
	0x5c, 0xdc, 0x89, 0x05, 0x05, 0x39, 0x99, 0xc9, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x12, 0x4c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0xc8, 0x42, 0x42, 0x12, 0x5c, 0xec, 0x29, 0x89, 0x25, 0x89, 0xc5, 0xa9,
	0x25, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x30, 0x2e, 0x48, 0xa6, 0xb8, 0x34, 0x39, 0x39,
	0xb5, 0xb8, 0x58, 0x82, 0x45, 0x81, 0x51, 0x83, 0x23, 0x08, 0xc6, 0x15, 0x12, 0xe3, 0x62, 0x2b,
	0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x91, 0x60, 0x05, 0x6b, 0x81, 0xf2, 0x92, 0xd8, 0xc0, 0xee, 0x35,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x50, 0xf0, 0x1b, 0x3b, 0xc5, 0x00, 0x00, 0x00,
}
