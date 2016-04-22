// Code generated by protoc-gen-go.
// source: types/types.proto
// DO NOT EDIT!

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	types/types.proto

It has these top-level messages:
	Request
	Response
	Validator
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MessageType int32

const (
	MessageType_NullMessage MessageType = 0
	MessageType_Echo        MessageType = 1
	MessageType_Flush       MessageType = 2
	MessageType_Info        MessageType = 3
	MessageType_SetOption   MessageType = 4
	MessageType_Exception   MessageType = 5
	MessageType_AppendTx    MessageType = 17
	MessageType_CheckTx     MessageType = 18
	MessageType_Commit      MessageType = 19
	MessageType_Query       MessageType = 20
	MessageType_InitChain   MessageType = 21
	MessageType_BeginBlock  MessageType = 22
	MessageType_EndBlock    MessageType = 23
)

var MessageType_name = map[int32]string{
	0:  "NullMessage",
	1:  "Echo",
	2:  "Flush",
	3:  "Info",
	4:  "SetOption",
	5:  "Exception",
	17: "AppendTx",
	18: "CheckTx",
	19: "Commit",
	20: "Query",
	21: "InitChain",
	22: "BeginBlock",
	23: "EndBlock",
}
var MessageType_value = map[string]int32{
	"NullMessage": 0,
	"Echo":        1,
	"Flush":       2,
	"Info":        3,
	"SetOption":   4,
	"Exception":   5,
	"AppendTx":    17,
	"CheckTx":     18,
	"Commit":      19,
	"Query":       20,
	"InitChain":   21,
	"BeginBlock":  22,
	"EndBlock":    23,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CodeType int32

const (
	CodeType_OK CodeType = 0
	// General response codes, 0 ~ 99
	CodeType_InternalError     CodeType = 1
	CodeType_EncodingError     CodeType = 2
	CodeType_BadNonce          CodeType = 3
	CodeType_Unauthorized      CodeType = 4
	CodeType_InsufficientFunds CodeType = 5
	CodeType_UnknownRequest    CodeType = 6
	// Reserved for basecoin, 100 ~ 199
	CodeType_BaseDuplicateAddress     CodeType = 101
	CodeType_BaseEncodingError        CodeType = 102
	CodeType_BaseInsufficientFees     CodeType = 103
	CodeType_BaseInsufficientFunds    CodeType = 104
	CodeType_BaseInsufficientGasPrice CodeType = 105
	CodeType_BaseInvalidInput         CodeType = 106
	CodeType_BaseInvalidOutput        CodeType = 107
	CodeType_BaseInvalidPubKey        CodeType = 108
	CodeType_BaseInvalidSequence      CodeType = 109
	CodeType_BaseInvalidSignature     CodeType = 110
	CodeType_BaseUnknownAddress       CodeType = 111
	CodeType_BaseUnknownPubKey        CodeType = 112
	CodeType_BaseUnknownPlugin        CodeType = 113
	// Reserved for governance, 200 ~ 299
	CodeType_GovUnknownEntity      CodeType = 201
	CodeType_GovUnknownGroup       CodeType = 202
	CodeType_GovUnknownProposal    CodeType = 203
	CodeType_GovDuplicateGroup     CodeType = 204
	CodeType_GovDuplicateMember    CodeType = 205
	CodeType_GovDuplicateProposal  CodeType = 206
	CodeType_GovDuplicateVote      CodeType = 207
	CodeType_GovInvalidMember      CodeType = 208
	CodeType_GovInvalidVote        CodeType = 209
	CodeType_GovInvalidVotingPower CodeType = 210
)

var CodeType_name = map[int32]string{
	0:   "OK",
	1:   "InternalError",
	2:   "EncodingError",
	3:   "BadNonce",
	4:   "Unauthorized",
	5:   "InsufficientFunds",
	6:   "UnknownRequest",
	101: "BaseDuplicateAddress",
	102: "BaseEncodingError",
	103: "BaseInsufficientFees",
	104: "BaseInsufficientFunds",
	105: "BaseInsufficientGasPrice",
	106: "BaseInvalidInput",
	107: "BaseInvalidOutput",
	108: "BaseInvalidPubKey",
	109: "BaseInvalidSequence",
	110: "BaseInvalidSignature",
	111: "BaseUnknownAddress",
	112: "BaseUnknownPubKey",
	113: "BaseUnknownPlugin",
	201: "GovUnknownEntity",
	202: "GovUnknownGroup",
	203: "GovUnknownProposal",
	204: "GovDuplicateGroup",
	205: "GovDuplicateMember",
	206: "GovDuplicateProposal",
	207: "GovDuplicateVote",
	208: "GovInvalidMember",
	209: "GovInvalidVote",
	210: "GovInvalidVotingPower",
}
var CodeType_value = map[string]int32{
	"OK":                       0,
	"InternalError":            1,
	"EncodingError":            2,
	"BadNonce":                 3,
	"Unauthorized":             4,
	"InsufficientFunds":        5,
	"UnknownRequest":           6,
	"BaseDuplicateAddress":     101,
	"BaseEncodingError":        102,
	"BaseInsufficientFees":     103,
	"BaseInsufficientFunds":    104,
	"BaseInsufficientGasPrice": 105,
	"BaseInvalidInput":         106,
	"BaseInvalidOutput":        107,
	"BaseInvalidPubKey":        108,
	"BaseInvalidSequence":      109,
	"BaseInvalidSignature":     110,
	"BaseUnknownAddress":       111,
	"BaseUnknownPubKey":        112,
	"BaseUnknownPlugin":        113,
	"GovUnknownEntity":         201,
	"GovUnknownGroup":          202,
	"GovUnknownProposal":       203,
	"GovDuplicateGroup":        204,
	"GovDuplicateMember":       205,
	"GovDuplicateProposal":     206,
	"GovDuplicateVote":         207,
	"GovInvalidMember":         208,
	"GovInvalidVote":           209,
	"GovInvalidVotingPower":    210,
}

func (x CodeType) String() string {
	return proto.EnumName(CodeType_name, int32(x))
}
func (CodeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Request struct {
	Type       MessageType  `protobuf:"varint,1,opt,name=type,enum=types.MessageType" json:"type,omitempty"`
	Data       []byte       `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Key        string       `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Value      string       `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	Validators []*Validator `protobuf:"bytes,5,rep,name=validators" json:"validators,omitempty"`
	Height     uint64       `protobuf:"varint,6,opt,name=height" json:"height,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetValidators() []*Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

type Response struct {
	Type       MessageType  `protobuf:"varint,1,opt,name=type,enum=types.MessageType" json:"type,omitempty"`
	Data       []byte       `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Code       CodeType     `protobuf:"varint,3,opt,name=code,enum=types.CodeType" json:"code,omitempty"`
	Error      string       `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
	Log        string       `protobuf:"bytes,5,opt,name=log" json:"log,omitempty"`
	Validators []*Validator `protobuf:"bytes,6,rep,name=validators" json:"validators,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetValidators() []*Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

type Validator struct {
	PubKey []byte `protobuf:"bytes,1,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Power  uint64 `protobuf:"varint,2,opt,name=power" json:"power,omitempty"`
}

func (m *Validator) Reset()                    { *m = Validator{} }
func (m *Validator) String() string            { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()               {}
func (*Validator) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Request)(nil), "types.Request")
	proto.RegisterType((*Response)(nil), "types.Response")
	proto.RegisterType((*Validator)(nil), "types.Validator")
	proto.RegisterEnum("types.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("types.CodeType", CodeType_name, CodeType_value)
}

var fileDescriptor0 = []byte{
	// 729 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x94, 0x4b, 0x4f, 0xeb, 0x46,
	0x14, 0xc7, 0xeb, 0xc4, 0x09, 0xc9, 0x09, 0x84, 0xc9, 0x90, 0x80, 0x5b, 0x75, 0x81, 0xa8, 0x54,
	0x21, 0x16, 0xb4, 0xa2, 0xab, 0x2e, 0x49, 0x1a, 0x50, 0x84, 0x80, 0xd4, 0x3c, 0xf6, 0xc6, 0x3e,
	0x89, 0xa7, 0x71, 0x66, 0x8c, 0x3d, 0x06, 0xd2, 0xef, 0x55, 0x55, 0xba, 0xfb, 0xbb, 0xb8, 0xef,
	0xc7, 0x27, 0xba, 0x33, 0x7e, 0x24, 0x21, 0x77, 0x71, 0x17, 0x77, 0x13, 0xcd, 0xf9, 0xfd, 0x67,
	0xce, 0x39, 0xff, 0x33, 0x13, 0x43, 0x4b, 0xce, 0x42, 0x8c, 0x7f, 0x4b, 0x7f, 0x0f, 0xc3, 0x48,
	0x48, 0x41, 0x2b, 0x69, 0xb0, 0xf7, 0xbf, 0x01, 0x6b, 0x36, 0xde, 0x27, 0x18, 0x4b, 0xfa, 0x2b,
	0x98, 0x1a, 0x5a, 0xc6, 0xae, 0xb1, 0xdf, 0x3c, 0xa2, 0x87, 0xd9, 0xf6, 0x73, 0x8c, 0x63, 0x67,
	0x8c, 0xd7, 0x2a, 0xb0, 0x53, 0x9d, 0x52, 0x30, 0x3d, 0x47, 0x3a, 0x56, 0x49, 0xed, 0x5b, 0xb7,
	0xd3, 0x35, 0x25, 0x50, 0x9e, 0xe0, 0xcc, 0x2a, 0x2b, 0x54, 0xb7, 0xf5, 0x92, 0xb6, 0xa1, 0xf2,
	0xe0, 0x04, 0x09, 0x5a, 0x66, 0xca, 0xb2, 0x80, 0xfe, 0x0e, 0xa0, 0x16, 0x4c, 0x9d, 0x11, 0x51,
	0x6c, 0x55, 0x76, 0xcb, 0xfb, 0x8d, 0x23, 0x92, 0x57, 0xba, 0x2d, 0x04, 0x7b, 0x69, 0x0f, 0xdd,
	0x86, 0xaa, 0x8f, 0x6c, 0xec, 0x4b, 0xab, 0xaa, 0x12, 0x99, 0x76, 0x1e, 0xed, 0xbd, 0x34, 0xa0,
	0x66, 0x63, 0x1c, 0x0a, 0x1e, 0xe3, 0x77, 0xb5, 0xfe, 0x0b, 0x98, 0xae, 0xf0, 0x30, 0xed, 0xbd,
	0x79, 0xb4, 0x99, 0x9f, 0xed, 0x29, 0x94, 0x1d, 0xd4, 0xa2, 0x76, 0x83, 0x51, 0x24, 0xa2, 0xc2,
	0x4d, 0x1a, 0x68, 0xd7, 0x81, 0x18, 0x2b, 0x1b, 0xa9, 0x6b, 0xb5, 0x5c, 0xf1, 0x57, 0xfd, 0xb6,
	0xbf, 0xbd, 0x3f, 0xa1, 0x3e, 0x17, 0xb4, 0xd9, 0x30, 0xb9, 0x3b, 0x53, 0x93, 0x34, 0xd2, 0x0e,
	0xf3, 0x48, 0x97, 0x0f, 0xc5, 0x23, 0x46, 0x69, 0xe3, 0xa6, 0x9d, 0x05, 0x07, 0x2f, 0x0c, 0x68,
	0x2c, 0x79, 0xa4, 0x9b, 0xd0, 0xb8, 0x48, 0x82, 0x20, 0x47, 0xe4, 0x07, 0x5a, 0x03, 0xb3, 0xef,
	0xfa, 0x82, 0x18, 0xb4, 0x0e, 0x95, 0x93, 0x20, 0x89, 0x7d, 0x52, 0xd2, 0x70, 0xc0, 0x47, 0x82,
	0x94, 0xe9, 0x06, 0xd4, 0xaf, 0x50, 0x5e, 0x86, 0x92, 0x09, 0x4e, 0x4c, 0x1d, 0xf6, 0x9f, 0x5c,
	0xcc, 0xc2, 0x0a, 0x5d, 0x87, 0xda, 0x71, 0x18, 0x22, 0xf7, 0xae, 0x9f, 0x48, 0x8b, 0x36, 0x60,
	0xad, 0xe7, 0xa3, 0x3b, 0x51, 0x81, 0x9a, 0x22, 0x54, 0x7b, 0x62, 0x3a, 0x65, 0x92, 0x6c, 0xe9,
	0xcc, 0x7f, 0x27, 0x18, 0xcd, 0x48, 0x5b, 0x27, 0x18, 0x70, 0x26, 0x7b, 0xbe, 0xc3, 0x38, 0xe9,
	0xd0, 0x26, 0x40, 0x17, 0xc7, 0x8c, 0x77, 0x03, 0xe1, 0x4e, 0xc8, 0xb6, 0x4e, 0xd8, 0xe7, 0x5e,
	0x16, 0xed, 0x1c, 0xfc, 0x57, 0x81, 0x5a, 0x31, 0x64, 0x5a, 0x85, 0xd2, 0xe5, 0x99, 0x6a, 0xb8,
	0x05, 0x1b, 0x03, 0x2e, 0x31, 0xe2, 0x4e, 0xd0, 0xd7, 0x13, 0x56, 0x9d, 0x2b, 0xd4, 0xe7, 0xea,
	0x0e, 0x18, 0x1f, 0x67, 0xa8, 0xa4, 0x13, 0x75, 0x1d, 0xef, 0x42, 0x70, 0x17, 0x95, 0x0b, 0x02,
	0xeb, 0x37, 0xdc, 0x49, 0xa4, 0x2f, 0x22, 0xf6, 0x2f, 0x7a, 0xca, 0x48, 0x07, 0x5a, 0x03, 0x1e,
	0x27, 0xa3, 0x11, 0x73, 0x19, 0x72, 0x79, 0x92, 0x70, 0x2f, 0x56, 0x86, 0x28, 0x34, 0x6f, 0xf8,
	0x84, 0x8b, 0x47, 0x9e, 0xbf, 0x78, 0x52, 0xa5, 0x16, 0xb4, 0xbb, 0x4e, 0x8c, 0x7f, 0x25, 0x61,
	0xc0, 0x5c, 0x47, 0xe2, 0xb1, 0xe7, 0x45, 0x6a, 0x7c, 0x04, 0x75, 0x12, 0xad, 0x3c, 0xaf, 0x3d,
	0x2a, 0x0e, 0x3c, 0xcb, 0x8f, 0x18, 0x93, 0x31, 0xfd, 0x11, 0x3a, 0x5f, 0x29, 0x69, 0x65, 0x9f,
	0xfe, 0x0c, 0xd6, 0xaa, 0x74, 0xea, 0xc4, 0xc3, 0x88, 0x29, 0x03, 0x4c, 0x5d, 0x2e, 0xc9, 0xd4,
	0xf4, 0x55, 0x0c, 0x78, 0x98, 0x48, 0xf2, 0x4f, 0x51, 0x3f, 0xa7, 0x97, 0x89, 0xd4, 0x78, 0xb2,
	0x82, 0x87, 0xe9, 0xf3, 0x20, 0x01, 0xdd, 0x81, 0xad, 0x25, 0x7c, 0xa5, 0xfd, 0xe9, 0xe9, 0x4c,
	0x17, 0xfd, 0x66, 0x02, 0x1b, 0x73, 0x47, 0x26, 0x11, 0x12, 0xae, 0xde, 0x1a, 0xd5, 0x4a, 0x3e,
	0x92, 0xc2, 0xb8, 0x28, 0x2a, 0xe4, 0x3c, 0xaf, 0x10, 0xae, 0xe2, 0x20, 0x51, 0x37, 0x4b, 0xee,
	0x15, 0x26, 0xa7, 0xe2, 0x21, 0xa7, 0x7d, 0x2e, 0x99, 0x9c, 0x91, 0x57, 0x86, 0xf2, 0xb4, 0xb9,
	0xc0, 0xa7, 0x91, 0x48, 0x42, 0xf2, 0xda, 0x50, 0x5d, 0xd2, 0x05, 0x1d, 0x46, 0x22, 0x14, 0xb1,
	0x13, 0x90, 0x37, 0x86, 0xea, 0xa5, 0xa5, 0x84, 0xf9, 0x2d, 0x64, 0x07, 0xde, 0x16, 0x07, 0xe6,
	0xfc, 0x1c, 0xa7, 0x77, 0x18, 0x91, 0x77, 0x86, 0x1a, 0x76, 0x7b, 0x59, 0x98, 0xe7, 0x7a, 0x6f,
	0xe4, 0x1d, 0xcd, 0xa5, 0x5b, 0x21, 0x91, 0x7c, 0x28, 0x70, 0x3e, 0x87, 0x3c, 0xd1, 0x47, 0x83,
	0x6e, 0x41, 0x73, 0x81, 0xd3, 0xbd, 0x9f, 0x0c, 0xfa, 0x13, 0x74, 0x9e, 0x41, 0x75, 0xff, 0x43,
	0xfd, 0x8f, 0x23, 0x9f, 0x8d, 0xbb, 0x6a, 0xfa, 0xfd, 0xfc, 0xe3, 0x4b, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc9, 0xcf, 0x96, 0x2d, 0x54, 0x05, 0x00, 0x00,
}
