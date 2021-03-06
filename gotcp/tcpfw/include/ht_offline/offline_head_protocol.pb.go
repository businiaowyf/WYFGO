// Code generated by protoc-gen-go.
// source: offline_head_protocol.proto
// DO NOT EDIT!

/*
Package ht_offline is a generated protocol buffer package.

It is generated from these files:
	offline_head_protocol.proto
	offline_body_protocol.proto

It has these top-level messages:
	ReqBody
	RspBody
	Msg
	SaveOfflineMsgReqBody
	SaveOfflineMsgRspBody
	DeleteOfflineMsgReqBody
	QueryOfflineMsgReqBody
	QueryOfflineMsgRspBody
	QueryMsgIDReqBody
	QueryMsgIDRspBody
*/
package ht_offline

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

// 根据不同的cmd_type来使用不同的body
type CMD_TYPE int32

const (
	CMD_TYPE_CMD_SAVE_OFFLINE_MSG   CMD_TYPE = 1
	CMD_TYPE_CMD_DELETE_OFFLINE_MSG CMD_TYPE = 2
	CMD_TYPE_CMD_QUERY_MSG_ID       CMD_TYPE = 3
	CMD_TYPE_CMD_QUERY_OFFLINE_MSG  CMD_TYPE = 4
)

var CMD_TYPE_name = map[int32]string{
	1: "CMD_SAVE_OFFLINE_MSG",
	2: "CMD_DELETE_OFFLINE_MSG",
	3: "CMD_QUERY_MSG_ID",
	4: "CMD_QUERY_OFFLINE_MSG",
}
var CMD_TYPE_value = map[string]int32{
	"CMD_SAVE_OFFLINE_MSG":   1,
	"CMD_DELETE_OFFLINE_MSG": 2,
	"CMD_QUERY_MSG_ID":       3,
	"CMD_QUERY_OFFLINE_MSG":  4,
}

func (x CMD_TYPE) Enum() *CMD_TYPE {
	p := new(CMD_TYPE)
	*p = x
	return p
}
func (x CMD_TYPE) String() string {
	return proto.EnumName(CMD_TYPE_name, int32(x))
}
func (x *CMD_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CMD_TYPE_value, data, "CMD_TYPE")
	if err != nil {
		return err
	}
	*x = CMD_TYPE(value)
	return nil
}
func (CMD_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RET_CODE int32

const (
	RET_CODE_RET_SUCCESS          RET_CODE = 0
	RET_CODE_RET_BODY_ERR         RET_CODE = 1
	RET_CODE_RET_INTERNAL_ERR     RET_CODE = 2
	RET_CODE_RET_SESS_TIMEOUT_ERR RET_CODE = 3
	RET_CODE_RET_DB_ERR           RET_CODE = 4
	RET_CODE_RET_RECORD_NOT_EXIST RET_CODE = 5
)

var RET_CODE_name = map[int32]string{
	0: "RET_SUCCESS",
	1: "RET_BODY_ERR",
	2: "RET_INTERNAL_ERR",
	3: "RET_SESS_TIMEOUT_ERR",
	4: "RET_DB_ERR",
	5: "RET_RECORD_NOT_EXIST",
}
var RET_CODE_value = map[string]int32{
	"RET_SUCCESS":          0,
	"RET_BODY_ERR":         1,
	"RET_INTERNAL_ERR":     2,
	"RET_SESS_TIMEOUT_ERR": 3,
	"RET_DB_ERR":           4,
	"RET_RECORD_NOT_EXIST": 5,
}

func (x RET_CODE) Enum() *RET_CODE {
	p := new(RET_CODE)
	*p = x
	return p
}
func (x RET_CODE) String() string {
	return proto.EnumName(RET_CODE_name, int32(x))
}
func (x *RET_CODE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RET_CODE_value, data, "RET_CODE")
	if err != nil {
		return err
	}
	*x = RET_CODE(value)
	return nil
}
func (RET_CODE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type SYS_TYPE int32

const (
	SYS_TYPE_SYS_IM_SERVER   SYS_TYPE = 1
	SYS_TYPE_SYS_VOIP_SERVER SYS_TYPE = 2
	SYS_TYPE_SYS_MUC_SERVER  SYS_TYPE = 3
	SYS_TYPE_SYS_QIM_SERVER  SYS_TYPE = 4
	SYS_TYPE_SYS_TYPE_TOOL   SYS_TYPE = 20
)

var SYS_TYPE_name = map[int32]string{
	1:  "SYS_IM_SERVER",
	2:  "SYS_VOIP_SERVER",
	3:  "SYS_MUC_SERVER",
	4:  "SYS_QIM_SERVER",
	20: "SYS_TYPE_TOOL",
}
var SYS_TYPE_value = map[string]int32{
	"SYS_IM_SERVER":   1,
	"SYS_VOIP_SERVER": 2,
	"SYS_MUC_SERVER":  3,
	"SYS_QIM_SERVER":  4,
	"SYS_TYPE_TOOL":   20,
}

func (x SYS_TYPE) Enum() *SYS_TYPE {
	p := new(SYS_TYPE)
	*p = x
	return p
}
func (x SYS_TYPE) String() string {
	return proto.EnumName(SYS_TYPE_name, int32(x))
}
func (x *SYS_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SYS_TYPE_value, data, "SYS_TYPE")
	if err != nil {
		return err
	}
	*x = SYS_TYPE(value)
	return nil
}
func (SYS_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ReqBody struct {
	SaveOfflineMsgReqbody   *SaveOfflineMsgReqBody   `protobuf:"bytes,1,opt,name=save_offline_msg_reqbody" json:"save_offline_msg_reqbody,omitempty"`
	DeleteOfflineMsgReqbody *DeleteOfflineMsgReqBody `protobuf:"bytes,2,opt,name=delete_offline_msg_reqbody" json:"delete_offline_msg_reqbody,omitempty"`
	QueryMsgIdReqbody       *QueryMsgIDReqBody       `protobuf:"bytes,3,opt,name=query_msg_id_reqbody" json:"query_msg_id_reqbody,omitempty"`
	QueryOfflineMsgReqbody  *QueryOfflineMsgReqBody  `protobuf:"bytes,4,opt,name=query_offline_msg_reqbody" json:"query_offline_msg_reqbody,omitempty"`
	XXX_unrecognized        []byte                   `json:"-"`
}

func (m *ReqBody) Reset()                    { *m = ReqBody{} }
func (m *ReqBody) String() string            { return proto.CompactTextString(m) }
func (*ReqBody) ProtoMessage()               {}
func (*ReqBody) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReqBody) GetSaveOfflineMsgReqbody() *SaveOfflineMsgReqBody {
	if m != nil {
		return m.SaveOfflineMsgReqbody
	}
	return nil
}

func (m *ReqBody) GetDeleteOfflineMsgReqbody() *DeleteOfflineMsgReqBody {
	if m != nil {
		return m.DeleteOfflineMsgReqbody
	}
	return nil
}

func (m *ReqBody) GetQueryMsgIdReqbody() *QueryMsgIDReqBody {
	if m != nil {
		return m.QueryMsgIdReqbody
	}
	return nil
}

func (m *ReqBody) GetQueryOfflineMsgReqbody() *QueryOfflineMsgReqBody {
	if m != nil {
		return m.QueryOfflineMsgReqbody
	}
	return nil
}

type RspBody struct {
	SaveOfflineMsgRspbody  *SaveOfflineMsgRspBody  `protobuf:"bytes,1,opt,name=save_offline_msg_rspbody" json:"save_offline_msg_rspbody,omitempty"`
	QueryMsgIdRspbody      *QueryMsgIDRspBody      `protobuf:"bytes,3,opt,name=query_msg_id_rspbody" json:"query_msg_id_rspbody,omitempty"`
	QueryOfflineMsgRspbody *QueryOfflineMsgRspBody `protobuf:"bytes,4,opt,name=query_offline_msg_rspbody" json:"query_offline_msg_rspbody,omitempty"`
	XXX_unrecognized       []byte                  `json:"-"`
}

func (m *RspBody) Reset()                    { *m = RspBody{} }
func (m *RspBody) String() string            { return proto.CompactTextString(m) }
func (*RspBody) ProtoMessage()               {}
func (*RspBody) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RspBody) GetSaveOfflineMsgRspbody() *SaveOfflineMsgRspBody {
	if m != nil {
		return m.SaveOfflineMsgRspbody
	}
	return nil
}

func (m *RspBody) GetQueryMsgIdRspbody() *QueryMsgIDRspBody {
	if m != nil {
		return m.QueryMsgIdRspbody
	}
	return nil
}

func (m *RspBody) GetQueryOfflineMsgRspbody() *QueryOfflineMsgRspBody {
	if m != nil {
		return m.QueryOfflineMsgRspbody
	}
	return nil
}

func init() {
	proto.RegisterType((*ReqBody)(nil), "ht.offline.ReqBody")
	proto.RegisterType((*RspBody)(nil), "ht.offline.RspBody")
	proto.RegisterEnum("ht.offline.CMD_TYPE", CMD_TYPE_name, CMD_TYPE_value)
	proto.RegisterEnum("ht.offline.RET_CODE", RET_CODE_name, RET_CODE_value)
	proto.RegisterEnum("ht.offline.SYS_TYPE", SYS_TYPE_name, SYS_TYPE_value)
}

func init() { proto.RegisterFile("offline_head_protocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x90, 0xdd, 0x8e, 0x12, 0x31,
	0x14, 0xc7, 0x9d, 0x01, 0x23, 0x39, 0xab, 0xbb, 0x63, 0x45, 0xc3, 0x62, 0x4c, 0x14, 0x6f, 0x0c,
	0x17, 0x5c, 0x78, 0xeb, 0xd5, 0x32, 0xed, 0x6e, 0x9a, 0x30, 0x74, 0x69, 0x0b, 0x91, 0xab, 0x66,
	0x95, 0xee, 0x47, 0xc2, 0x3a, 0xc0, 0xa0, 0xc9, 0x3e, 0x81, 0x0f, 0xe1, 0x5b, 0xf9, 0x44, 0xf6,
	0x94, 0xe9, 0x82, 0x82, 0x1b, 0xee, 0x4e, 0xff, 0x1f, 0xbf, 0xf6, 0x14, 0x5e, 0xe7, 0x97, 0x97,
	0xd3, 0x9b, 0x6f, 0xd6, 0x5c, 0xdb, 0x8b, 0x89, 0x99, 0x2d, 0xf2, 0x65, 0xfe, 0x35, 0x9f, 0x76,
	0xfc, 0x40, 0xe0, 0x7a, 0xd9, 0x29, 0xfd, 0xe6, 0x7d, 0xf0, 0x4b, 0x3e, 0xb9, 0xfb, 0x27, 0xd8,
	0xfa, 0x15, 0xc3, 0x13, 0x69, 0xe7, 0x5d, 0x67, 0x91, 0x14, 0x1a, 0xc5, 0xc5, 0x0f, 0x6b, 0x42,
	0xfe, 0xb6, 0xb8, 0x32, 0x0b, 0x3b, 0xc7, 0x5a, 0x23, 0x7a, 0x1b, 0x7d, 0x38, 0xf8, 0xf8, 0xae,
	0xb3, 0xe6, 0x76, 0x94, 0xcb, 0x8a, 0xd5, 0x9c, 0x15, 0x57, 0x01, 0x72, 0x06, 0xcd, 0x89, 0x9d,
	0xda, 0xe5, 0x6e, 0x4c, 0xec, 0x31, 0xef, 0x37, 0x31, 0xd4, 0xa7, 0xb7, 0x41, 0x9f, 0xa0, 0x3e,
	0xff, 0x6e, 0x17, 0x77, 0xbe, 0x7f, 0x33, 0xb9, 0x47, 0x54, 0x3c, 0xe2, 0xcd, 0x26, 0x62, 0x80,
	0x39, 0x57, 0xe5, 0x34, 0x94, 0x19, 0x1c, 0xaf, 0xca, 0xbb, 0x1e, 0x51, 0xf5, 0x84, 0xd6, 0x16,
	0x61, 0xeb, 0x0d, 0xad, 0xdf, 0x91, 0xfb, 0x9d, 0x62, 0xf6, 0xff, 0xdf, 0x29, 0x66, 0x7b, 0xfe,
	0x4e, 0x09, 0xd9, 0x5a, 0xaa, 0x04, 0x3c, 0xbc, 0x54, 0x59, 0xde, 0xbd, 0x54, 0x49, 0xd8, 0x63,
	0xa9, 0x15, 0xa6, 0x3d, 0x87, 0x5a, 0x9a, 0x51, 0xa3, 0xc7, 0xe7, 0x8c, 0x34, 0xa0, 0x8e, 0xb3,
	0x3a, 0x19, 0x31, 0x23, 0x4e, 0x4f, 0x7b, 0xbc, 0xcf, 0x4c, 0xa6, 0xce, 0x92, 0x88, 0x34, 0xe1,
	0x15, 0x3a, 0x94, 0xf5, 0x98, 0xfe, 0xdb, 0x8b, 0x49, 0x1d, 0x12, 0xf4, 0x06, 0x43, 0x26, 0xc7,
	0x28, 0x19, 0x4e, 0x93, 0x0a, 0x39, 0x86, 0x97, 0x6b, 0x75, 0xb3, 0x50, 0x6d, 0xff, 0x8c, 0xa0,
	0x26, 0x99, 0x36, 0xa9, 0xa0, 0x8c, 0x1c, 0xc1, 0x01, 0xce, 0x6a, 0x98, 0xa6, 0x4c, 0xa9, 0xe4,
	0x11, 0x49, 0xe0, 0x29, 0x0a, 0x5d, 0x41, 0xc7, 0x86, 0x49, 0xe9, 0x2e, 0x77, 0x17, 0xa0, 0xc2,
	0xfb, 0x9a, 0xc9, 0xfe, 0x49, 0xcf, 0xab, 0x31, 0x3e, 0xd6, 0x17, 0x5d, 0xcb, 0x68, 0x9e, 0x31,
	0x31, 0xd4, 0xde, 0xa9, 0x90, 0x43, 0x00, 0x74, 0x68, 0xd7, 0x9f, 0xab, 0x21, 0x29, 0x59, 0x2a,
	0x24, 0x35, 0x7d, 0xe1, 0x72, 0x9f, 0xb9, 0xd2, 0xc9, 0xe3, 0xf6, 0x2d, 0xd4, 0xd4, 0x58, 0xad,
	0x96, 0x7f, 0x0e, 0xcf, 0x70, 0xe6, 0x99, 0x43, 0xca, 0x11, 0xc3, 0x8b, 0x5f, 0xc0, 0x11, 0x4a,
	0x23, 0xc1, 0xcf, 0x83, 0x18, 0x13, 0x02, 0x87, 0x28, 0x66, 0xc3, 0x34, 0x68, 0x95, 0xa0, 0x0d,
	0xd6, 0xe5, 0x6a, 0xe0, 0x21, 0xdb, 0x68, 0x21, 0x7a, 0x49, 0xfd, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xfb, 0x09, 0x68, 0x5a, 0xa5, 0x03, 0x00, 0x00,
}
