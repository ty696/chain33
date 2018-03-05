// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Reply struct {
	IsOk bool   `protobuf:"varint,1,opt,name=isOk" json:"isOk,omitempty"`
	Msg  []byte `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Reply) GetIsOk() bool {
	if m != nil {
		return m.IsOk
	}
	return false
}

func (m *Reply) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

type ReqString struct {
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *ReqString) Reset()                    { *m = ReqString{} }
func (m *ReqString) String() string            { return proto.CompactTextString(m) }
func (*ReqString) ProtoMessage()               {}
func (*ReqString) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *ReqString) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ReplyString struct {
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *ReplyString) Reset()                    { *m = ReplyString{} }
func (m *ReplyString) String() string            { return proto.CompactTextString(m) }
func (*ReplyString) ProtoMessage()               {}
func (*ReplyString) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ReplyString) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ReplyStrings struct {
	Datas []string `protobuf:"bytes,1,rep,name=datas" json:"datas,omitempty"`
}

func (m *ReplyStrings) Reset()                    { *m = ReplyStrings{} }
func (m *ReplyStrings) String() string            { return proto.CompactTextString(m) }
func (*ReplyStrings) ProtoMessage()               {}
func (*ReplyStrings) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *ReplyStrings) GetDatas() []string {
	if m != nil {
		return m.Datas
	}
	return nil
}

type ReqInt struct {
	Height int64 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
}

func (m *ReqInt) Reset()                    { *m = ReqInt{} }
func (m *ReqInt) String() string            { return proto.CompactTextString(m) }
func (*ReqInt) ProtoMessage()               {}
func (*ReqInt) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *ReqInt) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type Int64 struct {
	Data int64 `protobuf:"varint,1,opt,name=data" json:"data,omitempty"`
}

func (m *Int64) Reset()                    { *m = Int64{} }
func (m *Int64) String() string            { return proto.CompactTextString(m) }
func (*Int64) ProtoMessage()               {}
func (*Int64) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *Int64) GetData() int64 {
	if m != nil {
		return m.Data
	}
	return 0
}

type ReqHash struct {
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *ReqHash) Reset()                    { *m = ReqHash{} }
func (m *ReqHash) String() string            { return proto.CompactTextString(m) }
func (*ReqHash) ProtoMessage()               {}
func (*ReqHash) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *ReqHash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type ReplyHash struct {
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *ReplyHash) Reset()                    { *m = ReplyHash{} }
func (m *ReplyHash) String() string            { return proto.CompactTextString(m) }
func (*ReplyHash) ProtoMessage()               {}
func (*ReplyHash) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *ReplyHash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type ReqNil struct {
}

func (m *ReqNil) Reset()                    { *m = ReqNil{} }
func (m *ReqNil) String() string            { return proto.CompactTextString(m) }
func (*ReqNil) ProtoMessage()               {}
func (*ReqNil) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

type ReqHashes struct {
	Hashes [][]byte `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
}

func (m *ReqHashes) Reset()                    { *m = ReqHashes{} }
func (m *ReqHashes) String() string            { return proto.CompactTextString(m) }
func (*ReqHashes) ProtoMessage()               {}
func (*ReqHashes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *ReqHashes) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type ReplyHashes struct {
	Hashes [][]byte `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
}

func (m *ReplyHashes) Reset()                    { *m = ReplyHashes{} }
func (m *ReplyHashes) String() string            { return proto.CompactTextString(m) }
func (*ReplyHashes) ProtoMessage()               {}
func (*ReplyHashes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *ReplyHashes) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type KeyValue struct {
	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *KeyValue) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Reply)(nil), "types.Reply")
	proto.RegisterType((*ReqString)(nil), "types.ReqString")
	proto.RegisterType((*ReplyString)(nil), "types.ReplyString")
	proto.RegisterType((*ReplyStrings)(nil), "types.ReplyStrings")
	proto.RegisterType((*ReqInt)(nil), "types.ReqInt")
	proto.RegisterType((*Int64)(nil), "types.Int64")
	proto.RegisterType((*ReqHash)(nil), "types.ReqHash")
	proto.RegisterType((*ReplyHash)(nil), "types.ReplyHash")
	proto.RegisterType((*ReqNil)(nil), "types.ReqNil")
	proto.RegisterType((*ReqHashes)(nil), "types.ReqHashes")
	proto.RegisterType((*ReplyHashes)(nil), "types.ReplyHashes")
	proto.RegisterType((*KeyValue)(nil), "types.KeyValue")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0x4d, 0x4b, 0xc5, 0x30,
	0x10, 0xa4, 0xd6, 0xd6, 0x76, 0xed, 0x41, 0x82, 0x48, 0x41, 0xc4, 0x1a, 0x15, 0x7a, 0xd1, 0x83,
	0x8a, 0xbf, 0xc1, 0x87, 0xa0, 0xb0, 0x82, 0xf7, 0xa8, 0x4b, 0x53, 0x5e, 0x3f, 0x5f, 0xa2, 0x90,
	0x7f, 0x2f, 0xd9, 0xe6, 0xa0, 0x07, 0x7d, 0xb7, 0x99, 0xec, 0xec, 0xce, 0xee, 0x04, 0x8a, 0xf7,
	0xb1, 0xef, 0xc7, 0xe1, 0x7a, 0xda, 0x8c, 0x76, 0x14, 0x89, 0x75, 0x13, 0x19, 0x79, 0x05, 0x09,
	0xd2, 0xd4, 0x39, 0x21, 0x60, 0xb7, 0x35, 0xcf, 0xeb, 0x32, 0xaa, 0xa2, 0x3a, 0x43, 0xc6, 0xe2,
	0x00, 0xe2, 0xde, 0x34, 0xe5, 0x4e, 0x15, 0xd5, 0x05, 0x7a, 0x28, 0x4f, 0x21, 0x47, 0x9a, 0x5f,
	0xec, 0xa6, 0x1d, 0x1a, 0xdf, 0xf2, 0xa1, 0xac, 0xe2, 0x96, 0x1c, 0x19, 0xcb, 0x33, 0xd8, 0xe7,
	0x79, 0xff, 0x48, 0x2e, 0xa0, 0xf8, 0x21, 0x31, 0xe2, 0x10, 0x12, 0xff, 0x6e, 0xca, 0xa8, 0x8a,
	0xeb, 0x1c, 0x17, 0x22, 0x2b, 0x48, 0x91, 0xe6, 0xd5, 0x60, 0xc5, 0x11, 0xa4, 0x9a, 0xda, 0x46,
	0x5b, 0x9e, 0x12, 0x63, 0x60, 0xf2, 0x18, 0x92, 0xd5, 0x60, 0xef, 0xef, 0x7e, 0x99, 0xc4, 0xc1,
	0xe4, 0x04, 0xf6, 0x90, 0xe6, 0x07, 0x65, 0xb4, 0x2f, 0x6b, 0x65, 0x34, 0x97, 0x0b, 0x64, 0xbc,
	0xdc, 0x31, 0x75, 0xee, 0x4f, 0x41, 0xc6, 0xf6, 0x4f, 0x6d, 0x27, 0xcf, 0xf9, 0x64, 0x2f, 0x24,
	0xc3, 0xbb, 0x30, 0xe2, 0x65, 0x0b, 0x0c, 0x4c, 0x5e, 0x86, 0xb3, 0xb7, 0xc8, 0x6e, 0x20, 0x7b,
	0x24, 0xf7, 0xaa, 0xba, 0x4f, 0xf2, 0xe1, 0xae, 0xc9, 0x05, 0x53, 0x0f, 0x7d, 0x10, 0x5f, 0xbe,
	0x14, 0x02, 0x5f, 0xc8, 0x5b, 0xca, 0xff, 0x75, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x5a,
	0xda, 0x59, 0xbf, 0x01, 0x00, 0x00,
}