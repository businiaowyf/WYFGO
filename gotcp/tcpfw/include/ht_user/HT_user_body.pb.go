// Code generated by protoc-gen-go.
// source: HT_user_body.proto
// DO NOT EDIT!

/*
Package ht_user is a generated protocol buffer package.

It is generated from these files:
	HT_user_body.proto

It has these top-level messages:
	UserInfoBody
	ErrorInfoBody
	ReqBody
	RspBody
*/
package ht_user

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

type ProtoCmd int32

const (
	ProtoCmd_CMD_GET_USERINFO  ProtoCmd = 1
	ProtoCmd_CMD_MOD_USERINFO  ProtoCmd = 2
	ProtoCmd_CMD_GET_BLACKLIST ProtoCmd = 3
)

var ProtoCmd_name = map[int32]string{
	1: "CMD_GET_USERINFO",
	2: "CMD_MOD_USERINFO",
	3: "CMD_GET_BLACKLIST",
}
var ProtoCmd_value = map[string]int32{
	"CMD_GET_USERINFO":  1,
	"CMD_MOD_USERINFO":  2,
	"CMD_GET_BLACKLIST": 3,
}

func (x ProtoCmd) Enum() *ProtoCmd {
	p := new(ProtoCmd)
	*p = x
	return p
}
func (x ProtoCmd) String() string {
	return proto.EnumName(ProtoCmd_name, int32(x))
}
func (x *ProtoCmd) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProtoCmd_value, data, "ProtoCmd")
	if err != nil {
		return err
	}
	*x = ProtoCmd(value)
	return nil
}
func (ProtoCmd) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ResultCode int32

const (
	ResultCode_RET_SUCCESS          ResultCode = 0
	ResultCode_RET_NOTFOUND         ResultCode = 1
	ResultCode_RET_CACHE_ERR        ResultCode = 2
	ResultCode_RET_INTERNAL_ERR     ResultCode = 3
	ResultCode_RET_SESS_TIMEOUT_ERR ResultCode = 4
)

var ResultCode_name = map[int32]string{
	0: "RET_SUCCESS",
	1: "RET_NOTFOUND",
	2: "RET_CACHE_ERR",
	3: "RET_INTERNAL_ERR",
	4: "RET_SESS_TIMEOUT_ERR",
}
var ResultCode_value = map[string]int32{
	"RET_SUCCESS":          0,
	"RET_NOTFOUND":         1,
	"RET_CACHE_ERR":        2,
	"RET_INTERNAL_ERR":     3,
	"RET_SESS_TIMEOUT_ERR": 4,
}

func (x ResultCode) Enum() *ResultCode {
	p := new(ResultCode)
	*p = x
	return p
}
func (x ResultCode) String() string {
	return proto.EnumName(ResultCode_name, int32(x))
}
func (x *ResultCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ResultCode_value, data, "ResultCode")
	if err != nil {
		return err
	}
	*x = ResultCode(value)
	return nil
}
func (ResultCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type UserInfoBody struct {
	UserID           *uint32  `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	UserName         []byte   `protobuf:"bytes,2,opt,name=UserName" json:"UserName,omitempty"`
	Sex              *uint32  `protobuf:"varint,3,opt,name=Sex" json:"Sex,omitempty"`
	Birthday         *uint64  `protobuf:"varint,4,opt,name=Birthday" json:"Birthday,omitempty"`
	National         []byte   `protobuf:"bytes,5,opt,name=National" json:"National,omitempty"`
	NativeLang       *uint32  `protobuf:"varint,6,opt,name=NativeLang" json:"NativeLang,omitempty"`
	LearnLang1       *uint32  `protobuf:"varint,7,opt,name=LearnLang1" json:"LearnLang1,omitempty"`
	LearnLang2       *uint32  `protobuf:"varint,8,opt,name=LearnLang2" json:"LearnLang2,omitempty"`
	LearnLang3       *uint32  `protobuf:"varint,9,opt,name=LearnLang3" json:"LearnLang3,omitempty"`
	TeachLang2       *uint32  `protobuf:"varint,10,opt,name=TeachLang2" json:"TeachLang2,omitempty"`
	TeachLang3       *uint32  `protobuf:"varint,11,opt,name=TeachLang3" json:"TeachLang3,omitempty"`
	BlackidList      []uint32 `protobuf:"varint,12,rep,name=blackid_list" json:"blackid_list,omitempty"`
	FriendidList     []uint32 `protobuf:"varint,13,rep,name=friendid_list" json:"friendid_list,omitempty"`
	FansList         []uint32 `protobuf:"varint,14,rep,name=fans_list" json:"fans_list,omitempty"`
	Nickname         []byte   `protobuf:"bytes,15,opt,name=nickname" json:"nickname,omitempty"`
	VipExpireTs      *uint32  `protobuf:"varint,16,opt,name=vip_expire_ts" json:"vip_expire_ts,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *UserInfoBody) Reset()                    { *m = UserInfoBody{} }
func (m *UserInfoBody) String() string            { return proto.CompactTextString(m) }
func (*UserInfoBody) ProtoMessage()               {}
func (*UserInfoBody) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserInfoBody) GetUserID() uint32 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *UserInfoBody) GetUserName() []byte {
	if m != nil {
		return m.UserName
	}
	return nil
}

func (m *UserInfoBody) GetSex() uint32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *UserInfoBody) GetBirthday() uint64 {
	if m != nil && m.Birthday != nil {
		return *m.Birthday
	}
	return 0
}

func (m *UserInfoBody) GetNational() []byte {
	if m != nil {
		return m.National
	}
	return nil
}

func (m *UserInfoBody) GetNativeLang() uint32 {
	if m != nil && m.NativeLang != nil {
		return *m.NativeLang
	}
	return 0
}

func (m *UserInfoBody) GetLearnLang1() uint32 {
	if m != nil && m.LearnLang1 != nil {
		return *m.LearnLang1
	}
	return 0
}

func (m *UserInfoBody) GetLearnLang2() uint32 {
	if m != nil && m.LearnLang2 != nil {
		return *m.LearnLang2
	}
	return 0
}

func (m *UserInfoBody) GetLearnLang3() uint32 {
	if m != nil && m.LearnLang3 != nil {
		return *m.LearnLang3
	}
	return 0
}

func (m *UserInfoBody) GetTeachLang2() uint32 {
	if m != nil && m.TeachLang2 != nil {
		return *m.TeachLang2
	}
	return 0
}

func (m *UserInfoBody) GetTeachLang3() uint32 {
	if m != nil && m.TeachLang3 != nil {
		return *m.TeachLang3
	}
	return 0
}

func (m *UserInfoBody) GetBlackidList() []uint32 {
	if m != nil {
		return m.BlackidList
	}
	return nil
}

func (m *UserInfoBody) GetFriendidList() []uint32 {
	if m != nil {
		return m.FriendidList
	}
	return nil
}

func (m *UserInfoBody) GetFansList() []uint32 {
	if m != nil {
		return m.FansList
	}
	return nil
}

func (m *UserInfoBody) GetNickname() []byte {
	if m != nil {
		return m.Nickname
	}
	return nil
}

func (m *UserInfoBody) GetVipExpireTs() uint32 {
	if m != nil && m.VipExpireTs != nil {
		return *m.VipExpireTs
	}
	return 0
}

type ErrorInfoBody struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg              []byte  `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ErrorInfoBody) Reset()                    { *m = ErrorInfoBody{} }
func (m *ErrorInfoBody) String() string            { return proto.CompactTextString(m) }
func (*ErrorInfoBody) ProtoMessage()               {}
func (*ErrorInfoBody) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ErrorInfoBody) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *ErrorInfoBody) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

type ReqBody struct {
	User             []*UserInfoBody `protobuf:"bytes,1,rep,name=user" json:"user,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ReqBody) Reset()                    { *m = ReqBody{} }
func (m *ReqBody) String() string            { return proto.CompactTextString(m) }
func (*ReqBody) ProtoMessage()               {}
func (*ReqBody) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReqBody) GetUser() []*UserInfoBody {
	if m != nil {
		return m.User
	}
	return nil
}

type RspBody struct {
	Error            *ErrorInfoBody  `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	User             []*UserInfoBody `protobuf:"bytes,2,rep,name=user" json:"user,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RspBody) Reset()                    { *m = RspBody{} }
func (m *RspBody) String() string            { return proto.CompactTextString(m) }
func (*RspBody) ProtoMessage()               {}
func (*RspBody) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RspBody) GetError() *ErrorInfoBody {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *RspBody) GetUser() []*UserInfoBody {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInfoBody)(nil), "ht.user.UserInfoBody")
	proto.RegisterType((*ErrorInfoBody)(nil), "ht.user.ErrorInfoBody")
	proto.RegisterType((*ReqBody)(nil), "ht.user.ReqBody")
	proto.RegisterType((*RspBody)(nil), "ht.user.RspBody")
	proto.RegisterEnum("ht.user.ProtoCmd", ProtoCmd_name, ProtoCmd_value)
	proto.RegisterEnum("ht.user.ResultCode", ResultCode_name, ResultCode_value)
}

func init() { proto.RegisterFile("HT_user_body.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x6e, 0x9b, 0x4c,
	0x10, 0xc5, 0x3f, 0x1b, 0x12, 0x3b, 0x03, 0x24, 0x18, 0x25, 0x9f, 0xf6, 0x32, 0x72, 0x55, 0x29,
	0xf2, 0x05, 0x52, 0x9d, 0x27, 0xb0, 0x31, 0x69, 0x50, 0x31, 0x54, 0xfc, 0xb9, 0x5e, 0x6d, 0xcc,
	0x3a, 0x46, 0xb1, 0xc1, 0x5d, 0x48, 0x94, 0xbc, 0x5e, 0x9f, 0xac, 0xb3, 0xeb, 0xd0, 0xb4, 0xb2,
	0xd4, 0x3b, 0xf6, 0x37, 0x67, 0xce, 0x68, 0xce, 0x00, 0xce, 0x7d, 0x46, 0x9f, 0x1b, 0x2e, 0xe8,
	0x43, 0x5d, 0xbc, 0xb9, 0x7b, 0x51, 0xb7, 0xb5, 0x33, 0xd8, 0xb4, 0xae, 0x64, 0xe3, 0x9f, 0x7d,
	0x30, 0x73, 0xfc, 0x08, 0xaa, 0x75, 0x3d, 0xc7, 0xba, 0x73, 0x0e, 0xa7, 0xea, 0xbd, 0x20, 0xbd,
	0xeb, 0xde, 0x8d, 0xe5, 0xd8, 0x30, 0x94, 0xef, 0x88, 0xed, 0x38, 0xe9, 0x23, 0x31, 0x1d, 0x03,
	0xb4, 0x94, 0xbf, 0x12, 0xad, 0x2b, 0xcf, 0x4b, 0xd1, 0x6e, 0x0a, 0xf6, 0x46, 0x74, 0x24, 0xba,
	0x24, 0x11, 0x6b, 0xcb, 0xba, 0x62, 0x5b, 0x72, 0xa2, 0x1a, 0x1c, 0x00, 0x49, 0x5e, 0x78, 0xc8,
	0xaa, 0x47, 0x72, 0xaa, 0xfa, 0x90, 0x85, 0x9c, 0x89, 0x4a, 0xa2, 0x2f, 0x64, 0x70, 0xc4, 0xa6,
	0x64, 0x78, 0xc4, 0x6e, 0xc9, 0x59, 0xc7, 0x32, 0xce, 0x56, 0x9b, 0x83, 0x0e, 0x8e, 0xd8, 0x2d,
	0x31, 0x14, 0xbb, 0x04, 0xf3, 0x61, 0xcb, 0x56, 0x4f, 0x65, 0x41, 0xb7, 0x65, 0xd3, 0x12, 0xf3,
	0x5a, 0x43, 0x7a, 0x05, 0xd6, 0x5a, 0x94, 0xbc, 0x2a, 0x3a, 0x6c, 0x29, 0x3c, 0x82, 0xb3, 0x35,
	0xab, 0x9a, 0x03, 0x3a, 0x57, 0x08, 0x37, 0xa9, 0xca, 0xd5, 0x53, 0x25, 0x57, 0xbf, 0x50, 0x9b,
	0x60, 0xef, 0x4b, 0xb9, 0xa7, 0xfc, 0x75, 0x5f, 0x0a, 0x4e, 0xdb, 0x86, 0xd8, 0x72, 0xd0, 0x78,
	0x02, 0x96, 0x2f, 0x44, 0xfd, 0x11, 0xa2, 0x09, 0xfa, 0xaa, 0x2e, 0xf8, 0x7b, 0x84, 0x18, 0xd8,
	0xae, 0x79, 0x3c, 0xa4, 0x37, 0x76, 0x61, 0x90, 0xf0, 0x1f, 0x4a, 0xf5, 0x09, 0x74, 0x79, 0x03,
	0x54, 0x69, 0x37, 0xc6, 0xf4, 0xca, 0x7d, 0xbf, 0x89, 0xfb, 0xe7, 0x3d, 0xc6, 0x39, 0xea, 0x9b,
	0xbd, 0xd2, 0x7f, 0x86, 0x13, 0x2e, 0xc7, 0x28, 0x5b, 0x63, 0xfa, 0xff, 0xef, 0x86, 0xbf, 0x87,
	0x77, 0xb6, 0xfd, 0x7f, 0xd8, 0x4e, 0x96, 0x30, 0xfc, 0x2e, 0xff, 0x04, 0x6f, 0x57, 0x60, 0x4e,
	0xb6, 0xb7, 0x5c, 0xd0, 0xaf, 0x7e, 0x46, 0xf3, 0xd4, 0x4f, 0x82, 0xe8, 0x2e, 0xb6, 0x7b, 0x1d,
	0x5d, 0xc6, 0x8b, 0x0f, 0xda, 0xc7, 0x04, 0x46, 0x9d, 0x76, 0x1e, 0xce, 0xbc, 0x6f, 0x61, 0x90,
	0x66, 0xb6, 0x36, 0x11, 0x00, 0x09, 0x6f, 0x9e, 0xb7, 0xad, 0x87, 0x6b, 0x3b, 0x17, 0x60, 0x24,
	0x28, 0x48, 0x73, 0xcf, 0xf3, 0xd3, 0xd4, 0xfe, 0x0f, 0x93, 0x34, 0x25, 0x88, 0xe2, 0xec, 0x2e,
	0xce, 0xa3, 0x05, 0xba, 0x8f, 0xc0, 0x92, 0xc4, 0x9b, 0x79, 0xf7, 0x3e, 0xf5, 0x93, 0x04, 0xad,
	0x71, 0xa0, 0x44, 0x41, 0x94, 0xf9, 0x49, 0x34, 0x0b, 0x15, 0xd5, 0x1c, 0x02, 0x97, 0xca, 0x0b,
	0x8d, 0x68, 0x16, 0x2c, 0xfd, 0x38, 0xcf, 0x54, 0x45, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0xee,
	0xf7, 0x35, 0x2a, 0xd8, 0x02, 0x00, 0x00,
}
