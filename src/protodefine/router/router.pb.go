// Code generated by protoc-gen-go.
// source: protodefine/router/router.proto
// DO NOT EDIT!

/*
Package bs_router is a generated protocol buffer package.

It is generated from these files:
	protodefine/router/router.proto

It has these top-level messages:
	RouterTransferData
	RegisterAppReq
	RegisterAppRsp
*/
package bs_router

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bs_types "protodefine/mytype"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CMDID_Router int32

const (
	CMDID_Router_IDUnknow         CMDID_Router = 0
	CMDID_Router_IDTransferData   CMDID_Router = 1
	CMDID_Router_IDRegisterAppReq CMDID_Router = 2
	CMDID_Router_IDRegisterAppRsp CMDID_Router = 3
	CMDID_Router_IDLast           CMDID_Router = 100
)

var CMDID_Router_name = map[int32]string{
	0:   "IDUnknow",
	1:   "IDTransferData",
	2:   "IDRegisterAppReq",
	3:   "IDRegisterAppRsp",
	100: "IDLast",
}
var CMDID_Router_value = map[string]int32{
	"IDUnknow":         0,
	"IDTransferData":   1,
	"IDRegisterAppReq": 2,
	"IDRegisterAppRsp": 3,
	"IDLast":           100,
}

func (x CMDID_Router) String() string {
	return proto.EnumName(CMDID_Router_name, int32(x))
}
func (CMDID_Router) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RouterTransferData_DataDirection int32

const (
	RouterTransferData_App2App    RouterTransferData_DataDirection = 0
	RouterTransferData_App2Client RouterTransferData_DataDirection = 1
	RouterTransferData_Client2App RouterTransferData_DataDirection = 2
)

var RouterTransferData_DataDirection_name = map[int32]string{
	0: "App2App",
	1: "App2Client",
	2: "Client2App",
}
var RouterTransferData_DataDirection_value = map[string]int32{
	"App2App":    0,
	"App2Client": 1,
	"Client2App": 2,
}

func (x RouterTransferData_DataDirection) String() string {
	return proto.EnumName(RouterTransferData_DataDirection_name, int32(x))
}
func (RouterTransferData_DataDirection) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type RouterTransferData struct {
	Base          *bs_types.BaseInfo               `protobuf:"bytes,1,opt,name=base" json:"base,omitempty"`
	DestApptype   uint32                           `protobuf:"varint,2,opt,name=dest_apptype,json=destApptype" json:"dest_apptype,omitempty"`
	DestAppid     uint32                           `protobuf:"varint,3,opt,name=dest_appid,json=destAppid" json:"dest_appid,omitempty"`
	SrcApptype    uint32                           `protobuf:"varint,4,opt,name=src_apptype,json=srcApptype" json:"src_apptype,omitempty"`
	SrcAppid      uint32                           `protobuf:"varint,5,opt,name=src_appid,json=srcAppid" json:"src_appid,omitempty"`
	DataCmdKind   uint32                           `protobuf:"varint,6,opt,name=data_cmd_kind,json=dataCmdKind" json:"data_cmd_kind,omitempty"`
	DataCmdSubid  uint32                           `protobuf:"varint,7,opt,name=data_cmd_subid,json=dataCmdSubid" json:"data_cmd_subid,omitempty"`
	Data          []byte                           `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
	DataDirection RouterTransferData_DataDirection `protobuf:"varint,9,opt,name=data_direction,json=dataDirection,enum=bs.router.RouterTransferData_DataDirection" json:"data_direction,omitempty"`
	// 如果是App2Client或者Client2App，下面4项有值，如果是App2App则可以不填
	ClientRemoteAddress string `protobuf:"bytes,10,opt,name=client_remote_address,json=clientRemoteAddress" json:"client_remote_address,omitempty"`
	AttGateid           uint32 `protobuf:"varint,11,opt,name=att_gateid,json=attGateid" json:"att_gateid,omitempty"`
	AttUserid           uint64 `protobuf:"varint,12,opt,name=att_userid,json=attUserid" json:"att_userid,omitempty"`
	AttGateconnid       uint64 `protobuf:"varint,13,opt,name=att_gateconnid,json=attGateconnid" json:"att_gateconnid,omitempty"`
}

func (m *RouterTransferData) Reset()                    { *m = RouterTransferData{} }
func (m *RouterTransferData) String() string            { return proto.CompactTextString(m) }
func (*RouterTransferData) ProtoMessage()               {}
func (*RouterTransferData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RouterTransferData) GetBase() *bs_types.BaseInfo {
	if m != nil {
		return m.Base
	}
	return nil
}

type RegisterAppReq struct {
	Base      *bs_types.BaseInfo `protobuf:"bytes,1,opt,name=base" json:"base,omitempty"`
	AuthKey   string             `protobuf:"bytes,2,opt,name=auth_key,json=authKey" json:"auth_key,omitempty"`
	AttData   string             `protobuf:"bytes,3,opt,name=att_data,json=attData" json:"att_data,omitempty"`
	MyAddress string             `protobuf:"bytes,4,opt,name=my_address,json=myAddress" json:"my_address,omitempty"`
	AppType   uint32             `protobuf:"varint,5,opt,name=app_type,json=appType" json:"app_type,omitempty"`
	AppId     uint32             `protobuf:"varint,6,opt,name=app_id,json=appId" json:"app_id,omitempty"`
}

func (m *RegisterAppReq) Reset()                    { *m = RegisterAppReq{} }
func (m *RegisterAppReq) String() string            { return proto.CompactTextString(m) }
func (*RegisterAppReq) ProtoMessage()               {}
func (*RegisterAppReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterAppReq) GetBase() *bs_types.BaseInfo {
	if m != nil {
		return m.Base
	}
	return nil
}

type RegisterAppRsp struct {
	Base      *bs_types.BaseInfo `protobuf:"bytes,1,opt,name=base" json:"base,omitempty"`
	RegResult uint32             `protobuf:"varint,2,opt,name=reg_result,json=regResult" json:"reg_result,omitempty"`
}

func (m *RegisterAppRsp) Reset()                    { *m = RegisterAppRsp{} }
func (m *RegisterAppRsp) String() string            { return proto.CompactTextString(m) }
func (*RegisterAppRsp) ProtoMessage()               {}
func (*RegisterAppRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RegisterAppRsp) GetBase() *bs_types.BaseInfo {
	if m != nil {
		return m.Base
	}
	return nil
}

func init() {
	proto.RegisterType((*RouterTransferData)(nil), "bs.router.RouterTransferData")
	proto.RegisterType((*RegisterAppReq)(nil), "bs.router.RegisterAppReq")
	proto.RegisterType((*RegisterAppRsp)(nil), "bs.router.RegisterAppRsp")
	proto.RegisterEnum("bs.router.CMDID_Router", CMDID_Router_name, CMDID_Router_value)
	proto.RegisterEnum("bs.router.RouterTransferData_DataDirection", RouterTransferData_DataDirection_name, RouterTransferData_DataDirection_value)
}

func init() { proto.RegisterFile("protodefine/router/router.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x6f, 0xd3, 0x4c,
	0x10, 0xc6, 0xdf, 0x34, 0x6e, 0x12, 0x4f, 0x1c, 0xcb, 0xda, 0x97, 0x4a, 0x06, 0x04, 0x85, 0x08,
	0x10, 0x02, 0xc9, 0x95, 0xc2, 0x95, 0x4b, 0x89, 0x25, 0x64, 0x15, 0x2e, 0x4b, 0x2b, 0x8e, 0x96,
	0x93, 0xdd, 0xa4, 0x56, 0xea, 0x3f, 0x78, 0x37, 0x42, 0xfe, 0x74, 0x9c, 0xf8, 0x5e, 0xcc, 0xce,
	0xda, 0x51, 0xa3, 0x5e, 0x7a, 0x88, 0xe3, 0x79, 0x7e, 0xcf, 0xac, 0xc6, 0x8f, 0xc7, 0x70, 0x5e,
	0x37, 0x95, 0xae, 0x84, 0xdc, 0xe4, 0xa5, 0xbc, 0x68, 0xaa, 0xbd, 0x96, 0x4d, 0xf7, 0x17, 0x11,
	0x61, 0xee, 0x4a, 0x45, 0x56, 0x78, 0xf6, 0xf2, 0xbe, 0xb7, 0x68, 0x75, 0x5b, 0xcb, 0x0b, 0x73,
	0x51, 0xd6, 0x3a, 0xff, 0xeb, 0x00, 0xe3, 0x64, 0xbd, 0x6e, 0xb2, 0x52, 0x6d, 0x64, 0x13, 0x67,
	0x3a, 0x63, 0xef, 0xc0, 0x59, 0x65, 0x4a, 0x86, 0x83, 0x57, 0x83, 0xf7, 0xd3, 0x05, 0x8b, 0xf0,
	0x40, 0xdb, 0xf5, 0x05, 0xd5, 0xa4, 0xdc, 0x54, 0x9c, 0x38, 0x7b, 0x0d, 0x9e, 0x90, 0x4a, 0xa7,
	0x59, 0x5d, 0x1b, 0x1e, 0x9e, 0xa0, 0x7f, 0xc6, 0xa7, 0x46, 0xbb, 0xb4, 0x12, 0x7b, 0x01, 0xd0,
	0x5b, 0x72, 0x11, 0x0e, 0xc9, 0xe0, 0x76, 0x86, 0x5c, 0xb0, 0x73, 0x98, 0xaa, 0x66, 0x7d, 0x38,
	0xc0, 0x21, 0x0e, 0x28, 0xf5, 0xfd, 0xcf, 0xc1, 0xed, 0x0c, 0xd8, 0x7e, 0x4a, 0x78, 0x62, 0x31,
	0x76, 0xcf, 0x61, 0x26, 0x70, 0xde, 0x74, 0x5d, 0x88, 0x74, 0x97, 0x97, 0x22, 0x1c, 0x75, 0x03,
	0xa0, 0xb8, 0x2c, 0xc4, 0x15, 0x4a, 0xec, 0x0d, 0xf8, 0x07, 0x8f, 0xda, 0xaf, 0xf0, 0x94, 0x31,
	0x99, 0xbc, 0xce, 0xf4, 0xc3, 0x68, 0x8c, 0x81, 0x63, 0xea, 0x70, 0x82, 0xcc, 0xe3, 0x74, 0xcf,
	0x78, 0xd7, 0x29, 0xf2, 0x46, 0xae, 0x75, 0x5e, 0x95, 0xa1, 0x8b, 0xd4, 0x5f, 0x7c, 0x8c, 0x0e,
	0x01, 0x47, 0x0f, 0xc3, 0x8b, 0xcc, 0x25, 0xee, 0x5b, 0x38, 0x0d, 0x78, 0x28, 0xd9, 0x02, 0xce,
	0xd6, 0x77, 0xb9, 0x2c, 0x75, 0xda, 0xc8, 0xa2, 0xd2, 0x32, 0xcd, 0x84, 0x68, 0xa4, 0x52, 0x21,
	0xe0, 0xd1, 0x2e, 0xff, 0xdf, 0x42, 0x4e, 0xec, 0xd2, 0x22, 0x13, 0x61, 0xa6, 0x75, 0xba, 0xcd,
	0xb4, 0xc4, 0xe9, 0xa7, 0x36, 0x42, 0x54, 0xbe, 0x92, 0xd0, 0xe3, 0xbd, 0x92, 0x0d, 0x62, 0x0f,
	0xb1, 0x43, 0xf8, 0x86, 0x04, 0xf6, 0x16, 0xfc, 0xbe, 0x7b, 0x5d, 0x95, 0x25, 0x5a, 0x66, 0x64,
	0x99, 0x75, 0x27, 0x58, 0x71, 0xfe, 0x19, 0x66, 0x47, 0x83, 0xb3, 0x29, 0x8c, 0x31, 0xe4, 0x05,
	0xfe, 0x82, 0xff, 0x98, 0x0f, 0x60, 0x8a, 0x25, 0x4d, 0x17, 0x0c, 0x4c, 0x6d, 0xef, 0x89, 0x9f,
	0xcc, 0xff, 0x0c, 0xc0, 0xe7, 0x72, 0x9b, 0x2b, 0x0c, 0x03, 0x15, 0x2e, 0x7f, 0x3d, 0x7a, 0x87,
	0x9e, 0xc2, 0x24, 0xdb, 0xeb, 0xdb, 0x74, 0x27, 0x5b, 0xda, 0x1f, 0x97, 0x8f, 0x4d, 0x7d, 0x25,
	0x5b, 0x42, 0x38, 0x3a, 0xbd, 0x98, 0x61, 0x87, 0xb4, 0xa6, 0x0d, 0xc5, 0x87, 0x2e, 0xda, 0x43,
	0x78, 0x0e, 0x41, 0xb7, 0x68, 0xfb, 0xc8, 0x4c, 0x67, 0x5d, 0xa7, 0xb4, 0x53, 0x76, 0x69, 0xc6,
	0x58, 0x5f, 0x9b, 0x85, 0x3a, 0x83, 0x91, 0x41, 0x79, 0xbf, 0x2c, 0xa7, 0x58, 0x25, 0x62, 0xfe,
	0xf3, 0xf8, 0x01, 0x54, 0xfd, 0xe8, 0x07, 0xc0, 0x51, 0x1a, 0xb9, 0xc5, 0xf7, 0xa9, 0xf6, 0x77,
	0xba, 0xfb, 0x04, 0x5c, 0x54, 0x38, 0x09, 0x1f, 0x6e, 0xc1, 0x5b, 0x7e, 0x8f, 0x93, 0x38, 0xb5,
	0xab, 0xc2, 0x3c, 0x98, 0x24, 0xf1, 0x4d, 0xb9, 0x2b, 0xab, 0xdf, 0x18, 0x2c, 0x03, 0x3f, 0x89,
	0xef, 0xaf, 0x0f, 0x86, 0xfb, 0x04, 0x82, 0x24, 0x3e, 0x4e, 0x33, 0x38, 0x79, 0xa8, 0xaa, 0x3a,
	0x18, 0x32, 0x80, 0x51, 0x12, 0x7f, 0xcb, 0x94, 0x0e, 0xc4, 0x6a, 0x44, 0xdf, 0xf4, 0xa7, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x6d, 0x0b, 0x51, 0x21, 0x04, 0x00, 0x00,
}