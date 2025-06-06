// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/payment/payment.proto

/*
Package go_micro_service_payment is a generated protocol buffer package.

It is generated from these files:
	proto/payment/payment.proto

It has these top-level messages:
	PaymentInfo
	PaymentID
	Response
	All
	PaymentAll
*/
package go_micro_service_payment

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

type PaymentInfo struct {
	Id            int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	PaymentName   string `protobuf:"bytes,2,opt,name=payment_name,json=paymentName" json:"payment_name,omitempty"`
	PaymentId     string `protobuf:"bytes,3,opt,name=payment_id,json=paymentId" json:"payment_id,omitempty"`
	PaymentStatus bool   `protobuf:"varint,4,opt,name=payment_status,json=paymentStatus" json:"payment_status,omitempty"`
	PaymentImage  string `protobuf:"bytes,5,opt,name=PaymentImage" json:"PaymentImage,omitempty"`
}

func (m *PaymentInfo) Reset()                    { *m = PaymentInfo{} }
func (m *PaymentInfo) String() string            { return proto.CompactTextString(m) }
func (*PaymentInfo) ProtoMessage()               {}
func (*PaymentInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PaymentInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PaymentInfo) GetPaymentName() string {
	if m != nil {
		return m.PaymentName
	}
	return ""
}

func (m *PaymentInfo) GetPaymentId() string {
	if m != nil {
		return m.PaymentId
	}
	return ""
}

func (m *PaymentInfo) GetPaymentStatus() bool {
	if m != nil {
		return m.PaymentStatus
	}
	return false
}

func (m *PaymentInfo) GetPaymentImage() string {
	if m != nil {
		return m.PaymentImage
	}
	return ""
}

type PaymentID struct {
	PaymentId int64 `protobuf:"varint,1,opt,name=payment_id,json=paymentId" json:"payment_id,omitempty"`
}

func (m *PaymentID) Reset()                    { *m = PaymentID{} }
func (m *PaymentID) String() string            { return proto.CompactTextString(m) }
func (*PaymentID) ProtoMessage()               {}
func (*PaymentID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PaymentID) GetPaymentId() int64 {
	if m != nil {
		return m.PaymentId
	}
	return 0
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type All struct {
}

func (m *All) Reset()                    { *m = All{} }
func (m *All) String() string            { return proto.CompactTextString(m) }
func (*All) ProtoMessage()               {}
func (*All) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type PaymentAll struct {
	PaymentInfo []*PaymentInfo `protobuf:"bytes,1,rep,name=payment_info,json=paymentInfo" json:"payment_info,omitempty"`
}

func (m *PaymentAll) Reset()                    { *m = PaymentAll{} }
func (m *PaymentAll) String() string            { return proto.CompactTextString(m) }
func (*PaymentAll) ProtoMessage()               {}
func (*PaymentAll) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PaymentAll) GetPaymentInfo() []*PaymentInfo {
	if m != nil {
		return m.PaymentInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*PaymentInfo)(nil), "go.micro.service.payment.PaymentInfo")
	proto.RegisterType((*PaymentID)(nil), "go.micro.service.payment.PaymentID")
	proto.RegisterType((*Response)(nil), "go.micro.service.payment.Response")
	proto.RegisterType((*All)(nil), "go.micro.service.payment.All")
	proto.RegisterType((*PaymentAll)(nil), "go.micro.service.payment.PaymentAll")
}

func init() { proto.RegisterFile("proto/payment/payment.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xef, 0x4a, 0xf3, 0x30,
	0x14, 0xc6, 0xdb, 0xf5, 0xdd, 0xeb, 0x7a, 0xf6, 0x47, 0xcd, 0xa7, 0x30, 0x1d, 0xd4, 0xe8, 0xa0,
	0xf8, 0xa1, 0xc2, 0xbc, 0x82, 0x4a, 0x11, 0xf7, 0x45, 0xa4, 0xa2, 0x82, 0x88, 0xa3, 0x2e, 0xd9,
	0x08, 0xa4, 0x4d, 0x59, 0xab, 0xb0, 0x5b, 0xf2, 0x9e, 0xbc, 0x17, 0x69, 0x97, 0x6c, 0x2a, 0x94,
	0x75, 0x9f, 0x9a, 0xf3, 0x9c, 0xa7, 0xbf, 0x73, 0xf2, 0xd0, 0xc2, 0x51, 0xba, 0x90, 0xb9, 0xbc,
	0x48, 0xa3, 0x65, 0xcc, 0x92, 0x5c, 0x3f, 0xbd, 0x52, 0x45, 0x78, 0x2e, 0xbd, 0x98, 0x4f, 0x17,
	0xd2, 0xcb, 0xd8, 0xe2, 0x83, 0x4f, 0x99, 0xa7, 0xfa, 0xe4, 0xd3, 0x84, 0xf6, 0xdd, 0xea, 0x3c,
	0x4e, 0x66, 0x12, 0xf5, 0xa0, 0xc1, 0x29, 0x36, 0x1d, 0xd3, 0xb5, 0xc2, 0x06, 0xa7, 0xe8, 0x04,
	0x3a, 0xca, 0x3a, 0x49, 0xa2, 0x98, 0xe1, 0x86, 0x63, 0xba, 0x76, 0xd8, 0x56, 0xda, 0x6d, 0x14,
	0x33, 0x34, 0x00, 0xd0, 0x16, 0x4e, 0xb1, 0x55, 0x1a, 0x6c, 0xa5, 0x8c, 0x29, 0x1a, 0x42, 0x4f,
	0xb7, 0xb3, 0x3c, 0xca, 0xdf, 0x33, 0xfc, 0xcf, 0x31, 0xdd, 0x56, 0xd8, 0x55, 0xea, 0x7d, 0x29,
	0x22, 0x02, 0x1d, 0xbd, 0x47, 0x1c, 0xcd, 0x19, 0x6e, 0x96, 0x9c, 0x5f, 0x1a, 0x39, 0x07, 0x5b,
	0xd7, 0xc1, 0x9f, 0xb1, 0xab, 0x8d, 0x37, 0x63, 0xc9, 0x31, 0xb4, 0x42, 0x96, 0xa5, 0x32, 0xc9,
	0x18, 0x3a, 0x00, 0x2b, 0xce, 0xe6, 0xa5, 0xc7, 0x0e, 0x8b, 0x23, 0x69, 0x82, 0xe5, 0x0b, 0x41,
	0x1e, 0x01, 0x14, 0xd0, 0x17, 0x02, 0xdd, 0x6c, 0xee, 0xca, 0x93, 0x99, 0xc4, 0xa6, 0x63, 0xb9,
	0xed, 0xd1, 0xd0, 0xab, 0x0a, 0xcf, 0xfb, 0x11, 0xdc, 0x3a, 0x92, 0xa2, 0x18, 0x7d, 0x59, 0xb0,
	0xa7, 0x9a, 0xe8, 0x19, 0xc0, 0xa7, 0x54, 0x57, 0xf5, 0x68, 0xfd, 0xd3, 0xed, 0xb6, 0x80, 0x18,
	0xe8, 0x05, 0xba, 0x0f, 0x29, 0x8d, 0x72, 0xb6, 0x23, 0x9e, 0x54, 0xdb, 0x74, 0x68, 0xc4, 0x40,
	0xaf, 0x70, 0x18, 0x30, 0xc1, 0xd6, 0xf4, 0xab, 0xe5, 0x38, 0x40, 0x75, 0x36, 0xab, 0xc9, 0x9f,
	0xc0, 0xfe, 0x35, 0x4f, 0xe8, 0xce, 0xf4, 0x7a, 0x97, 0x24, 0x06, 0x7a, 0x82, 0x5e, 0x31, 0xc0,
	0x17, 0x42, 0xe7, 0x33, 0xa8, 0x7e, 0xd5, 0x17, 0xa2, 0x7f, 0xb6, 0x95, 0x5c, 0x7c, 0x35, 0xc6,
	0xdb, 0xff, 0xf2, 0xb7, 0xba, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xc9, 0x5d, 0x8e, 0x75,
	0x03, 0x00, 0x00,
}
