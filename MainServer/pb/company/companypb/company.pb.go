// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/company/companypb/company.proto

package companypb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Company struct {
	CompanyUUID          string   `protobuf:"bytes,1,opt,name=CompanyUUID,proto3" json:"CompanyUUID,omitempty"`
	CompanyName          string   `protobuf:"bytes,2,opt,name=CompanyName,proto3" json:"CompanyName,omitempty"`
	CompanyWebsite       string   `protobuf:"bytes,3,opt,name=CompanyWebsite,proto3" json:"CompanyWebsite,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Company) Reset()         { *m = Company{} }
func (m *Company) String() string { return proto.CompactTextString(m) }
func (*Company) ProtoMessage()    {}
func (*Company) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b07a7d68a935aae, []int{0}
}

func (m *Company) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Company.Unmarshal(m, b)
}
func (m *Company) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Company.Marshal(b, m, deterministic)
}
func (m *Company) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Company.Merge(m, src)
}
func (m *Company) XXX_Size() int {
	return xxx_messageInfo_Company.Size(m)
}
func (m *Company) XXX_DiscardUnknown() {
	xxx_messageInfo_Company.DiscardUnknown(m)
}

var xxx_messageInfo_Company proto.InternalMessageInfo

func (m *Company) GetCompanyUUID() string {
	if m != nil {
		return m.CompanyUUID
	}
	return ""
}

func (m *Company) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *Company) GetCompanyWebsite() string {
	if m != nil {
		return m.CompanyWebsite
	}
	return ""
}

type CompanyResponse struct {
	Companies            []*Company `protobuf:"bytes,1,rep,name=companies,proto3" json:"companies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CompanyResponse) Reset()         { *m = CompanyResponse{} }
func (m *CompanyResponse) String() string { return proto.CompactTextString(m) }
func (*CompanyResponse) ProtoMessage()    {}
func (*CompanyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b07a7d68a935aae, []int{1}
}

func (m *CompanyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompanyResponse.Unmarshal(m, b)
}
func (m *CompanyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompanyResponse.Marshal(b, m, deterministic)
}
func (m *CompanyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompanyResponse.Merge(m, src)
}
func (m *CompanyResponse) XXX_Size() int {
	return xxx_messageInfo_CompanyResponse.Size(m)
}
func (m *CompanyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CompanyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CompanyResponse proto.InternalMessageInfo

func (m *CompanyResponse) GetCompanies() []*Company {
	if m != nil {
		return m.Companies
	}
	return nil
}

func init() {
	proto.RegisterType((*Company)(nil), "Company")
	proto.RegisterType((*CompanyResponse)(nil), "CompanyResponse")
}

func init() {
	proto.RegisterFile("pb/company/companypb/company.proto", fileDescriptor_2b07a7d68a935aae)
}

var fileDescriptor_2b07a7d68a935aae = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x48, 0xd2, 0x4f,
	0xce, 0xcf, 0x2d, 0x48, 0xcc, 0xab, 0x84, 0xd1, 0x08, 0x11, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c,
	0xa5, 0x52, 0x2e, 0x76, 0x67, 0x88, 0x80, 0x90, 0x02, 0x17, 0x37, 0x94, 0x19, 0x1a, 0xea, 0xe9,
	0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x2c, 0x84, 0xa4, 0xc2, 0x2f, 0x31, 0x37, 0x55,
	0x82, 0x09, 0x45, 0x05, 0x48, 0x48, 0x48, 0x8d, 0x8b, 0x0f, 0xca, 0x0d, 0x4f, 0x4d, 0x2a, 0xce,
	0x2c, 0x49, 0x95, 0x60, 0x06, 0x2b, 0x42, 0x13, 0x55, 0xb2, 0xe4, 0xe2, 0x87, 0x8a, 0x04, 0xa5,
	0x16, 0x17, 0xe4, 0xe7, 0x15, 0x83, 0xb4, 0x72, 0x42, 0x9c, 0x96, 0x99, 0x5a, 0x2c, 0xc1, 0xa8,
	0xc0, 0xac, 0xc1, 0x6d, 0xc4, 0xa1, 0x07, 0x53, 0x84, 0x90, 0x72, 0xe2, 0x8e, 0xe2, 0x84, 0x7b,
	0x26, 0x89, 0x0d, 0xec, 0x0b, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x62, 0x27, 0xf0, 0xfe,
	0xeb, 0x00, 0x00, 0x00,
}
