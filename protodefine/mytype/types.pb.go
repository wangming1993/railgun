// Code generated by protoc-gen-go.
// source: railgun/protodefine/mytype/types.proto
// DO NOT EDIT!

/*
Package bs_types is a generated protocol buffer package.

It is generated from these files:
	railgun/protodefine/mytype/types.proto

It has these top-level messages:
	BaseInfo
	BaseUserInfo
	UserSessionInfo
	FundItem
*/
package bs_types

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

type CMDKindId int32

const (
	CMDKindId_NullType              CMDKindId = 0
	CMDKindId_IDKindLogger          CMDKindId = 8
	CMDKindId_IDKindNetTCP          CMDKindId = 9
	CMDKindId_IDKindRouter          CMDKindId = 10
	CMDKindId_IDKindAppFrame        CMDKindId = 11
	CMDKindId_IDKindGate            CMDKindId = 12
	CMDKindId_IDKindClient          CMDKindId = 13
	CMDKindId_IDKindMatchClient     CMDKindId = 14
	CMDKindId_IDKindMCMS            CMDKindId = 15
	CMDKindId_IDKindTableLogic      CMDKindId = 16
	CMDKindId_IDKindFund            CMDKindId = 17
	CMDKindId_IDKindHallClient      CMDKindId = 18
	CMDKindId_IDKindMatchPhase      CMDKindId = 19
	CMDKindId_IDKindRankList        CMDKindId = 20
	CMDKindId_IDKindList            CMDKindId = 21
	CMDKindId_IDKindMatchDB         CMDKindId = 22
	CMDKindId_IDKindPrivateInternal CMDKindId = 256
	CMDKindId_IDKindGameStart       CMDKindId = 288
	CMDKindId_IDLast                CMDKindId = 4094
)

var CMDKindId_name = map[int32]string{
	0:    "NullType",
	8:    "IDKindLogger",
	9:    "IDKindNetTCP",
	10:   "IDKindRouter",
	11:   "IDKindAppFrame",
	12:   "IDKindGate",
	13:   "IDKindClient",
	14:   "IDKindMatchClient",
	15:   "IDKindMCMS",
	16:   "IDKindTableLogic",
	17:   "IDKindFund",
	18:   "IDKindHallClient",
	19:   "IDKindMatchPhase",
	20:   "IDKindRankList",
	21:   "IDKindList",
	22:   "IDKindMatchDB",
	256:  "IDKindPrivateInternal",
	288:  "IDKindGameStart",
	4094: "IDLast",
}
var CMDKindId_value = map[string]int32{
	"NullType":              0,
	"IDKindLogger":          8,
	"IDKindNetTCP":          9,
	"IDKindRouter":          10,
	"IDKindAppFrame":        11,
	"IDKindGate":            12,
	"IDKindClient":          13,
	"IDKindMatchClient":     14,
	"IDKindMCMS":            15,
	"IDKindTableLogic":      16,
	"IDKindFund":            17,
	"IDKindHallClient":      18,
	"IDKindMatchPhase":      19,
	"IDKindRankList":        20,
	"IDKindList":            21,
	"IDKindMatchDB":         22,
	"IDKindPrivateInternal": 256,
	"IDKindGameStart":       288,
	"IDLast":                4094,
}

func (x CMDKindId) String() string {
	return proto.EnumName(CMDKindId_name, int32(x))
}
func (CMDKindId) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EnumAppId int32

const (
	EnumAppId_UnknowId    EnumAppId = 0
	EnumAppId_Send2All    EnumAppId = 1
	EnumAppId_Send2AnyOne EnumAppId = 2
)

var EnumAppId_name = map[int32]string{
	0: "UnknowId",
	1: "Send2All",
	2: "Send2AnyOne",
}
var EnumAppId_value = map[string]int32{
	"UnknowId":    0,
	"Send2All":    1,
	"Send2AnyOne": 2,
}

func (x EnumAppId) String() string {
	return proto.EnumName(EnumAppId_name, int32(x))
}
func (EnumAppId) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type EnumAppType int32

const (
	EnumAppType_UnknowType     EnumAppType = 0
	EnumAppType_Gate           EnumAppType = 5
	EnumAppType_Router         EnumAppType = 6
	EnumAppType_Login          EnumAppType = 7
	EnumAppType_Online         EnumAppType = 8
	EnumAppType_Fund           EnumAppType = 9
	EnumAppType_List           EnumAppType = 10
	EnumAppType_FreeMatch      EnumAppType = 12
	EnumAppType_Match          EnumAppType = 13
	EnumAppType_TableLogic     EnumAppType = 14
	EnumAppType_MatchPhase     EnumAppType = 15
	EnumAppType_RankList       EnumAppType = 16
	EnumAppType_MatchDB        EnumAppType = 18
	EnumAppType_Conect_To_Gold EnumAppType = 19
	EnumAppType_Tool           EnumAppType = 224
	EnumAppType_Last           EnumAppType = 254
)

var EnumAppType_name = map[int32]string{
	0:   "UnknowType",
	5:   "Gate",
	6:   "Router",
	7:   "Login",
	8:   "Online",
	9:   "Fund",
	10:  "List",
	12:  "FreeMatch",
	13:  "Match",
	14:  "TableLogic",
	15:  "MatchPhase",
	16:  "RankList",
	18:  "MatchDB",
	19:  "Conect_To_Gold",
	224: "Tool",
	254: "Last",
}
var EnumAppType_value = map[string]int32{
	"UnknowType":     0,
	"Gate":           5,
	"Router":         6,
	"Login":          7,
	"Online":         8,
	"Fund":           9,
	"List":           10,
	"FreeMatch":      12,
	"Match":          13,
	"TableLogic":     14,
	"MatchPhase":     15,
	"RankList":       16,
	"MatchDB":        18,
	"Conect_To_Gold": 19,
	"Tool":           224,
	"Last":           254,
}

func (x EnumAppType) String() string {
	return proto.EnumName(EnumAppType_name, int32(x))
}
func (EnumAppType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type BaseUserInfo_UserType int32

const (
	BaseUserInfo_UNKNOW BaseUserInfo_UserType = 0
	BaseUserInfo_Normal BaseUserInfo_UserType = 1
	BaseUserInfo_Robot  BaseUserInfo_UserType = 10
)

var BaseUserInfo_UserType_name = map[int32]string{
	0:  "UNKNOW",
	1:  "Normal",
	10: "Robot",
}
var BaseUserInfo_UserType_value = map[string]int32{
	"UNKNOW": 0,
	"Normal": 1,
	"Robot":  10,
}

func (x BaseUserInfo_UserType) String() string {
	return proto.EnumName(BaseUserInfo_UserType_name, int32(x))
}
func (BaseUserInfo_UserType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type BaseInfo struct {
	KindId     uint32 `protobuf:"varint,1,opt,name=kind_id,json=kindId" json:"kind_id,omitempty"`
	SubId      uint32 `protobuf:"varint,2,opt,name=sub_id,json=subId" json:"sub_id,omitempty"`
	ConnId     uint64 `protobuf:"varint,3,opt,name=conn_id,json=connId" json:"conn_id,omitempty"`
	GateConnId uint64 `protobuf:"varint,4,opt,name=gate_conn_id,json=gateConnId" json:"gate_conn_id,omitempty"`
	RemoteAdd  string `protobuf:"bytes,5,opt,name=remote_add,json=remoteAdd" json:"remote_add,omitempty"`
	AttApptype uint32 `protobuf:"varint,6,opt,name=att_apptype,json=attApptype" json:"att_apptype,omitempty"`
	AttAppid   uint32 `protobuf:"varint,7,opt,name=att_appid,json=attAppid" json:"att_appid,omitempty"`
}

func (m *BaseInfo) Reset()                    { *m = BaseInfo{} }
func (m *BaseInfo) String() string            { return proto.CompactTextString(m) }
func (*BaseInfo) ProtoMessage()               {}
func (*BaseInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

//
// 基础用户信息
//
type BaseUserInfo struct {
	UserId       uint64                `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	GameId       uint64                `protobuf:"varint,2,opt,name=game_id,json=gameId" json:"game_id,omitempty"`
	Gender       uint32                `protobuf:"varint,3,opt,name=gender" json:"gender,omitempty"`
	VipLevel     uint32                `protobuf:"varint,4,opt,name=vip_level,json=vipLevel" json:"vip_level,omitempty"`
	VipEndday    uint32                `protobuf:"varint,5,opt,name=vip_endday,json=vipEndday" json:"vip_endday,omitempty"`
	MarketId     uint32                `protobuf:"varint,6,opt,name=market_id,json=marketId" json:"market_id,omitempty"`
	FaceId       uint32                `protobuf:"varint,7,opt,name=face_id,json=faceId" json:"face_id,omitempty"`
	CustomFace   string                `protobuf:"bytes,8,opt,name=custom_face,json=customFace" json:"custom_face,omitempty"`
	NickName     string                `protobuf:"bytes,9,opt,name=nick_name,json=nickName" json:"nick_name,omitempty"`
	ManagerRight uint32                `protobuf:"varint,10,opt,name=manager_right,json=managerRight" json:"manager_right,omitempty"`
	UserRight    uint32                `protobuf:"varint,11,opt,name=user_right,json=userRight" json:"user_right,omitempty"`
	UserType     BaseUserInfo_UserType `protobuf:"varint,12,opt,name=user_type,json=userType,enum=bs.types.BaseUserInfo_UserType" json:"user_type,omitempty"`
}

func (m *BaseUserInfo) Reset()                    { *m = BaseUserInfo{} }
func (m *BaseUserInfo) String() string            { return proto.CompactTextString(m) }
func (*BaseUserInfo) ProtoMessage()               {}
func (*BaseUserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

//
// 用户当前连接信息
//
type UserSessionInfo struct {
	GateId     uint32 `protobuf:"varint,1,opt,name=gate_id,json=gateId" json:"gate_id,omitempty"`
	GateConnId uint64 `protobuf:"varint,2,opt,name=gate_conn_id,json=gateConnId" json:"gate_conn_id,omitempty"`
	Client_IP  string `protobuf:"bytes,9,opt,name=client_IP,json=clientIP" json:"client_IP,omitempty"`
}

func (m *UserSessionInfo) Reset()                    { *m = UserSessionInfo{} }
func (m *UserSessionInfo) String() string            { return proto.CompactTextString(m) }
func (*UserSessionInfo) ProtoMessage()               {}
func (*UserSessionInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 财富类型项
type FundItem struct {
	Id    uint32 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Count uint64 `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	Name  string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
}

func (m *FundItem) Reset()                    { *m = FundItem{} }
func (m *FundItem) String() string            { return proto.CompactTextString(m) }
func (*FundItem) ProtoMessage()               {}
func (*FundItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*BaseInfo)(nil), "bs.types.BaseInfo")
	proto.RegisterType((*BaseUserInfo)(nil), "bs.types.BaseUserInfo")
	proto.RegisterType((*UserSessionInfo)(nil), "bs.types.UserSessionInfo")
	proto.RegisterType((*FundItem)(nil), "bs.types.FundItem")
	proto.RegisterEnum("bs.types.CMDKindId", CMDKindId_name, CMDKindId_value)
	proto.RegisterEnum("bs.types.EnumAppId", EnumAppId_name, EnumAppId_value)
	proto.RegisterEnum("bs.types.EnumAppType", EnumAppType_name, EnumAppType_value)
	proto.RegisterEnum("bs.types.BaseUserInfo_UserType", BaseUserInfo_UserType_name, BaseUserInfo_UserType_value)
}

func init() { proto.RegisterFile("railgun/protodefine/mytype/types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 861 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x54, 0xc1, 0x6e, 0xe3, 0x36,
	0x10, 0xad, 0xbd, 0xb6, 0x22, 0x8d, 0xed, 0x98, 0x61, 0x37, 0x5b, 0xa3, 0x45, 0x91, 0x85, 0x0b,
	0x14, 0x45, 0x80, 0x3a, 0xc0, 0xf6, 0xd2, 0x43, 0x2f, 0x59, 0x67, 0xb3, 0x35, 0x36, 0x71, 0x0c,
	0xc5, 0x41, 0x8f, 0x02, 0x2d, 0x31, 0x8e, 0x60, 0x89, 0x34, 0x24, 0x2a, 0x45, 0x6e, 0xfd, 0x94,
	0xfe, 0x46, 0x7f, 0xa3, 0xf7, 0x02, 0xfd, 0x88, 0x9e, 0x8b, 0xce, 0x90, 0xd2, 0x2a, 0x40, 0xf7,
	0x62, 0xf0, 0xbd, 0x19, 0x72, 0xe6, 0xbd, 0x19, 0x19, 0xbe, 0x2d, 0x44, 0x9a, 0x6d, 0x2b, 0x75,
	0xb6, 0x2f, 0xb4, 0xd1, 0x89, 0xbc, 0x4f, 0x95, 0x3c, 0xcb, 0x9f, 0xcc, 0xd3, 0x5e, 0x9e, 0xd1,
	0x4f, 0x39, 0xb3, 0x01, 0xee, 0x6f, 0xca, 0x99, 0xc5, 0xd3, 0x3f, 0x3b, 0xe0, 0xbf, 0x15, 0xa5,
	0x5c, 0xa8, 0x7b, 0xcd, 0xbf, 0x80, 0x83, 0x5d, 0xaa, 0x92, 0x28, 0x4d, 0x26, 0x9d, 0xd7, 0x9d,
	0xef, 0x46, 0xa1, 0x47, 0x70, 0x91, 0xf0, 0x63, 0xf0, 0xca, 0x6a, 0x43, 0x7c, 0xd7, 0xf2, 0x7d,
	0x44, 0x48, 0x63, 0x7e, 0xac, 0x95, 0x22, 0xfe, 0x05, 0xf2, 0xbd, 0xd0, 0x23, 0x88, 0x81, 0xd7,
	0x30, 0xdc, 0x0a, 0x23, 0xa3, 0x26, 0xda, 0xb3, 0x51, 0x20, 0x6e, 0xee, 0x32, 0xbe, 0x06, 0x28,
	0x64, 0xae, 0x31, 0x47, 0x24, 0xc9, 0xa4, 0x8f, 0xf1, 0x20, 0x0c, 0x1c, 0x73, 0x9e, 0x24, 0xfc,
	0x04, 0x06, 0xc2, 0x98, 0x48, 0xec, 0xf7, 0xd4, 0xe6, 0xc4, 0xb3, 0x55, 0x01, 0xa9, 0x73, 0xc7,
	0xf0, 0xaf, 0x20, 0xa8, 0x13, 0xf0, 0xf9, 0x03, 0x1b, 0xf6, 0x5d, 0x38, 0x4d, 0xa6, 0x7f, 0xbc,
	0x80, 0x21, 0x89, 0xba, 0x2b, 0x65, 0xd1, 0x08, 0xab, 0xf0, 0xdc, 0x08, 0xc3, 0x46, 0x09, 0x3a,
	0x05, 0x5b, 0x91, 0xcb, 0x46, 0x19, 0x06, 0x08, 0x62, 0xe0, 0x15, 0x78, 0x5b, 0xa9, 0x12, 0x59,
	0x58, 0x65, 0xe8, 0x84, 0x43, 0x54, 0xf7, 0x31, 0xdd, 0x47, 0x99, 0x7c, 0x94, 0x99, 0x95, 0x85,
	0x75, 0x91, 0xb8, 0x22, 0x4c, 0xa2, 0x28, 0x88, 0x99, 0x89, 0x78, 0xb2, 0xa2, 0x46, 0x21, 0xa5,
	0xbf, 0xb3, 0x04, 0xdd, 0xcd, 0x45, 0xb1, 0x93, 0x86, 0xca, 0x39, 0x49, 0xbe, 0x23, 0x5c, 0x27,
	0xf7, 0x22, 0xb6, 0x9d, 0x38, 0x39, 0x1e, 0xc1, 0x85, 0xb5, 0x22, 0xae, 0x4a, 0xa3, 0xf3, 0x88,
	0x88, 0x89, 0x6f, 0xad, 0x02, 0x47, 0x5d, 0x22, 0x43, 0xcf, 0xaa, 0x34, 0xde, 0x45, 0x0a, 0x3b,
	0x9f, 0x04, 0x36, 0xec, 0x13, 0xb1, 0x44, 0xcc, 0xbf, 0x81, 0x51, 0x2e, 0x94, 0xd8, 0xa2, 0xf8,
	0x22, 0xdd, 0x3e, 0x98, 0x09, 0xd8, 0xc7, 0x87, 0x35, 0x19, 0x12, 0x47, 0x7d, 0x5b, 0x7b, 0x5c,
	0xc6, 0xc0, 0xf5, 0x4d, 0x8c, 0x0b, 0xff, 0x04, 0x16, 0x44, 0x76, 0x14, 0x43, 0x8c, 0x1e, 0xbe,
	0x39, 0x99, 0x35, 0x1b, 0x34, 0x7b, 0x6e, 0xf4, 0x8c, 0x0e, 0x6b, 0xa4, 0x43, 0xbf, 0xaa, 0x4f,
	0xd3, 0xef, 0xc1, 0x6f, 0x58, 0x0e, 0xe0, 0xdd, 0x2d, 0x3f, 0x2c, 0x6f, 0x7e, 0x61, 0x9f, 0xd1,
	0x79, 0xa9, 0x8b, 0x5c, 0x64, 0xac, 0xc3, 0x03, 0xe8, 0x87, 0x7a, 0xa3, 0x0d, 0x83, 0x69, 0x0a,
	0x63, 0x4a, 0xbf, 0x95, 0x65, 0x99, 0x6a, 0xd5, 0x4c, 0xcf, 0x6e, 0x53, 0xbb, 0x96, 0x04, 0x3f,
	0xb1, 0x66, 0xdd, 0xff, 0xad, 0x19, 0x7a, 0x13, 0x67, 0xa9, 0x54, 0x26, 0x5a, 0xac, 0x1a, 0x6f,
	0x1c, 0xb1, 0x58, 0x4d, 0x2f, 0xc0, 0xbf, 0xac, 0x70, 0xbf, 0x8d, 0xcc, 0xf9, 0x21, 0x74, 0x3f,
	0x6e, 0x37, 0x9e, 0xf8, 0x4b, 0xe8, 0xc7, 0xba, 0x52, 0xa6, 0x5e, 0x6c, 0x07, 0x38, 0x87, 0x9e,
	0x75, 0xb9, 0x67, 0x5f, 0xb2, 0xe7, 0xd3, 0x7f, 0xba, 0x10, 0xcc, 0xaf, 0x2f, 0x3e, 0xb8, 0x2f,
	0x65, 0x08, 0xfe, 0xb2, 0xca, 0x32, 0x52, 0x8b, 0x1a, 0x19, 0x0c, 0x17, 0x36, 0x72, 0xa5, 0xb7,
	0xe8, 0x36, 0xf3, 0x5b, 0x66, 0x29, 0xcd, 0x7a, 0xbe, 0x62, 0x41, 0xcb, 0x84, 0xba, 0x32, 0x98,
	0x03, 0x58, 0xe5, 0xd0, 0x31, 0xb8, 0xcd, 0x97, 0x05, 0xd6, 0x60, 0x03, 0xec, 0x0f, 0x1c, 0xf7,
	0x1e, 0xc5, 0xb1, 0x61, 0x7b, 0x6b, 0x6e, 0xd5, 0xb0, 0x11, 0x7e, 0xa3, 0x47, 0x8e, 0xb9, 0x16,
	0x26, 0x7e, 0xa8, 0xe9, 0xc3, 0xf6, 0xe2, 0xf5, 0xfc, 0xfa, 0x96, 0x8d, 0x51, 0x18, 0x73, 0x78,
	0x2d, 0x36, 0x99, 0xc4, 0xbe, 0xd2, 0x98, 0xb1, 0x36, 0x8b, 0x0c, 0x61, 0x47, 0x6d, 0xd6, 0xcf,
	0x22, 0xcb, 0xea, 0xb7, 0x78, 0xcb, 0xda, 0x12, 0xab, 0x07, 0x9c, 0x3c, 0xfb, 0xbc, 0x6d, 0x37,
	0x14, 0x6a, 0x77, 0x95, 0x96, 0x86, 0xbd, 0x6c, 0xdf, 0xb3, 0xf8, 0x98, 0x1f, 0xc1, 0xe8, 0xd9,
	0xcd, 0x8b, 0xb7, 0xec, 0x15, 0xff, 0x12, 0x8e, 0x1d, 0xb5, 0x2a, 0xd2, 0x47, 0x9a, 0xa7, 0x42,
	0xf5, 0x0a, 0xd7, 0xe1, 0xb7, 0x2e, 0x16, 0x1a, 0x37, 0x6a, 0x73, 0x79, 0x6b, 0x44, 0x61, 0xd8,
	0xef, 0x5d, 0x3e, 0x00, 0x6f, 0x71, 0x71, 0x25, 0xf0, 0xc1, 0x7f, 0x4f, 0x4e, 0x7f, 0x84, 0xe0,
	0x9d, 0xaa, 0x72, 0xb4, 0xc8, 0xb9, 0x7e, 0xa7, 0x76, 0x4a, 0xff, 0xba, 0x48, 0xd0, 0x75, 0x44,
	0xb7, 0xf8, 0x0d, 0xbe, 0x39, 0xcf, 0x68, 0xb7, 0xc6, 0x30, 0x70, 0x48, 0x3d, 0xdd, 0x28, 0xc9,
	0xba, 0xa7, 0x7f, 0x75, 0x60, 0x50, 0x5f, 0xb5, 0x4b, 0x89, 0xbd, 0xba, 0xcb, 0xf5, 0xd0, 0x7c,
	0xe8, 0x59, 0x93, 0xfb, 0xb4, 0xa2, 0xf5, 0x50, 0x3c, 0x5a, 0x51, 0x32, 0x4b, 0xb1, 0x03, 0xa2,
	0x6f, 0x54, 0x86, 0x7f, 0xad, 0x38, 0x4f, 0x4c, 0xb6, 0x96, 0x05, 0x74, 0xb2, 0x62, 0x81, 0x8f,
	0x20, 0xb8, 0x2c, 0xa4, 0xb4, 0x52, 0x71, 0x54, 0x78, 0xd3, 0x1d, 0x47, 0x54, 0xea, 0x99, 0xed,
	0x76, 0x38, 0xcf, 0xac, 0x1c, 0x53, 0xe7, 0x1f, 0x4d, 0x64, 0xa8, 0xf7, 0xa0, 0xb1, 0x8b, 0x93,
	0xcb, 0xb8, 0xd3, 0x32, 0x36, 0xd1, 0x5a, 0x47, 0xef, 0x75, 0x96, 0xa0, 0xf3, 0x01, 0xf4, 0xd6,
	0x5a, 0x67, 0xec, 0x6f, 0xfa, 0x82, 0x7a, 0xce, 0x99, 0xce, 0xc6, 0xb3, 0xff, 0xf1, 0x3f, 0xfc,
	0x17, 0x00, 0x00, 0xff, 0xff, 0x88, 0x07, 0xd4, 0x88, 0x0d, 0x06, 0x00, 0x00,
}
