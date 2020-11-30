// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/etherboy-core/txmsg/txmsg.proto

/*
Package txmsg is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/etherboy-core/txmsg/txmsg.proto

It has these top-level messages:
	EtherboyCreateAccountTx
	EtherboyStateTx
	EtherboyAppState
	StateQueryParams
	StateQueryResult
*/
package txmsg

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type EtherboyCreateAccountTx struct {
	Version int32  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Owner   string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Data    []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *EtherboyCreateAccountTx) Reset()                    { *m = EtherboyCreateAccountTx{} }
func (m *EtherboyCreateAccountTx) String() string            { return proto.CompactTextString(m) }
func (*EtherboyCreateAccountTx) ProtoMessage()               {}
func (*EtherboyCreateAccountTx) Descriptor() ([]byte, []int) { return fileDescriptorTxmsg, []int{0} }

func (m *EtherboyCreateAccountTx) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *EtherboyCreateAccountTx) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *EtherboyCreateAccountTx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type EtherboyStateTx struct {
	Version int32  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Owner   string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Data    []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *EtherboyStateTx) Reset()                    { *m = EtherboyStateTx{} }
func (m *EtherboyStateTx) String() string            { return proto.CompactTextString(m) }
func (*EtherboyStateTx) ProtoMessage()               {}
func (*EtherboyStateTx) Descriptor() ([]byte, []int) { return fileDescriptorTxmsg, []int{1} }

func (m *EtherboyStateTx) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *EtherboyStateTx) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *EtherboyStateTx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type EtherboyAppState struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Blob    []byte `protobuf:"bytes,2,opt,name=blob,proto3" json:"blob,omitempty"`
}

func (m *EtherboyAppState) Reset()                    { *m = EtherboyAppState{} }
func (m *EtherboyAppState) String() string            { return proto.CompactTextString(m) }
func (*EtherboyAppState) ProtoMessage()               {}
func (*EtherboyAppState) Descriptor() ([]byte, []int) { return fileDescriptorTxmsg, []int{2} }

func (m *EtherboyAppState) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *EtherboyAppState) GetBlob() []byte {
	if m != nil {
		return m.Blob
	}
	return nil
}

type StateQueryParams struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *StateQueryParams) Reset()                    { *m = StateQueryParams{} }
func (m *StateQueryParams) String() string            { return proto.CompactTextString(m) }
func (*StateQueryParams) ProtoMessage()               {}
func (*StateQueryParams) Descriptor() ([]byte, []int) { return fileDescriptorTxmsg, []int{3} }

func (m *StateQueryParams) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type StateQueryResult struct {
	State []byte `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (m *StateQueryResult) Reset()                    { *m = StateQueryResult{} }
func (m *StateQueryResult) String() string            { return proto.CompactTextString(m) }
func (*StateQueryResult) ProtoMessage()               {}
func (*StateQueryResult) Descriptor() ([]byte, []int) { return fileDescriptorTxmsg, []int{4} }

func (m *StateQueryResult) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

func init() {
	proto.RegisterType((*EtherboyCreateAccountTx)(nil), "EtherboyCreateAccountTx")
	proto.RegisterType((*EtherboyStateTx)(nil), "EtherboyStateTx")
	proto.RegisterType((*EtherboyAppState)(nil), "EtherboyAppState")
	proto.RegisterType((*StateQueryParams)(nil), "StateQueryParams")
	proto.RegisterType((*StateQueryResult)(nil), "StateQueryResult")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/etherboy-core/txmsg/txmsg.proto", fileDescriptorTxmsg)
}

var fileDescriptorTxmsg = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x89, 0xba, 0x8a, 0x43, 0xc1, 0xa5, 0x08, 0xf6, 0x58, 0x7a, 0xea, 0xc5, 0xed, 0x41,
	0xf0, 0xec, 0x22, 0xde, 0x35, 0xea, 0xd1, 0x43, 0xd2, 0x0e, 0xbb, 0x8b, 0x6d, 0xa7, 0x4c, 0xa6,
	0xee, 0xee, 0xbf, 0x97, 0xa4, 0x0d, 0xea, 0xdd, 0x4b, 0x78, 0x5f, 0x78, 0xf9, 0x1e, 0x04, 0xee,
	0x37, 0x3b, 0xd9, 0x8e, 0x76, 0x55, 0x53, 0x57, 0xb5, 0x44, 0x5d, 0x8f, 0xb2, 0x27, 0xfe, 0xac,
	0x50, 0xb6, 0xc8, 0x96, 0x8e, 0xb7, 0x35, 0x31, 0x56, 0x72, 0xe8, 0xdc, 0x66, 0x3a, 0x57, 0x03,
	0x93, 0x50, 0xf1, 0x01, 0x37, 0x4f, 0x73, 0xe5, 0x91, 0xd1, 0x08, 0xae, 0xeb, 0x9a, 0xc6, 0x5e,
	0xde, 0x0e, 0x69, 0x06, 0x17, 0x5f, 0xc8, 0x6e, 0x47, 0x7d, 0xa6, 0x72, 0x55, 0x2e, 0x74, 0xc4,
	0xf4, 0x1a, 0x16, 0xb4, 0xef, 0x91, 0xb3, 0x93, 0x5c, 0x95, 0x97, 0x7a, 0x82, 0x34, 0x85, 0xb3,
	0xc6, 0x88, 0xc9, 0x4e, 0x73, 0x55, 0x26, 0x3a, 0xe4, 0xe2, 0x1d, 0xae, 0xa2, 0xfe, 0x55, 0x8c,
	0xe0, 0x3f, 0x69, 0x1f, 0x60, 0x19, 0xb5, 0xeb, 0x61, 0x08, 0x66, 0xef, 0x35, 0x4d, 0xc3, 0xe8,
	0x5c, 0xf0, 0x26, 0x3a, 0xa2, 0x37, 0xd8, 0x96, 0x6c, 0xd0, 0x26, 0x3a, 0xe4, 0xa2, 0x84, 0x65,
	0x78, 0xf6, 0x32, 0x22, 0x1f, 0x9f, 0x0d, 0x9b, 0xce, 0xfd, 0xec, 0xab, 0x5f, 0xfb, 0x7f, 0x9b,
	0x1a, 0xdd, 0xd8, 0x8a, 0x6f, 0x3a, 0x7f, 0x37, 0x2f, 0x4d, 0x60, 0xcf, 0xc3, 0x97, 0xde, 0x7d,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x14, 0xf4, 0xe5, 0x8c, 0x01, 0x00, 0x00,
}
