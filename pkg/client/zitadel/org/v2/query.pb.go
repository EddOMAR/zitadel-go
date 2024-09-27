// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: zitadel/org/v2/query.proto

package org

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v2 "github.com/zitadel/zitadel-go/v3/pkg/client/zitadel/object/v2"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrganizationFieldName int32

const (
	OrganizationFieldName_ORGANIZATION_FIELD_NAME_UNSPECIFIED OrganizationFieldName = 0
	OrganizationFieldName_ORGANIZATION_FIELD_NAME_NAME        OrganizationFieldName = 1
)

// Enum value maps for OrganizationFieldName.
var (
	OrganizationFieldName_name = map[int32]string{
		0: "ORGANIZATION_FIELD_NAME_UNSPECIFIED",
		1: "ORGANIZATION_FIELD_NAME_NAME",
	}
	OrganizationFieldName_value = map[string]int32{
		"ORGANIZATION_FIELD_NAME_UNSPECIFIED": 0,
		"ORGANIZATION_FIELD_NAME_NAME":        1,
	}
)

func (x OrganizationFieldName) Enum() *OrganizationFieldName {
	p := new(OrganizationFieldName)
	*p = x
	return p
}

func (x OrganizationFieldName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrganizationFieldName) Descriptor() protoreflect.EnumDescriptor {
	return file_zitadel_org_v2_query_proto_enumTypes[0].Descriptor()
}

func (OrganizationFieldName) Type() protoreflect.EnumType {
	return &file_zitadel_org_v2_query_proto_enumTypes[0]
}

func (x OrganizationFieldName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrganizationFieldName.Descriptor instead.
func (OrganizationFieldName) EnumDescriptor() ([]byte, []int) {
	return file_zitadel_org_v2_query_proto_rawDescGZIP(), []int{0}
}

type SearchQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Query:
	//
	//	*SearchQuery_NameQuery
	//	*SearchQuery_DomainQuery
	//	*SearchQuery_StateQuery
	//	*SearchQuery_IdQuery
	Query isSearchQuery_Query `protobuf_oneof:"query"`
}

func (x *SearchQuery) Reset() {
	*x = SearchQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_org_v2_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchQuery) ProtoMessage() {}

func (x *SearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_org_v2_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchQuery.ProtoReflect.Descriptor instead.
func (*SearchQuery) Descriptor() ([]byte, []int) {
	return file_zitadel_org_v2_query_proto_rawDescGZIP(), []int{0}
}

func (m *SearchQuery) GetQuery() isSearchQuery_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (x *SearchQuery) GetNameQuery() *OrganizationNameQuery {
	if x, ok := x.GetQuery().(*SearchQuery_NameQuery); ok {
		return x.NameQuery
	}
	return nil
}

func (x *SearchQuery) GetDomainQuery() *OrganizationDomainQuery {
	if x, ok := x.GetQuery().(*SearchQuery_DomainQuery); ok {
		return x.DomainQuery
	}
	return nil
}

func (x *SearchQuery) GetStateQuery() *OrganizationStateQuery {
	if x, ok := x.GetQuery().(*SearchQuery_StateQuery); ok {
		return x.StateQuery
	}
	return nil
}

func (x *SearchQuery) GetIdQuery() *OrganizationIDQuery {
	if x, ok := x.GetQuery().(*SearchQuery_IdQuery); ok {
		return x.IdQuery
	}
	return nil
}

type isSearchQuery_Query interface {
	isSearchQuery_Query()
}

type SearchQuery_NameQuery struct {
	NameQuery *OrganizationNameQuery `protobuf:"bytes,1,opt,name=name_query,json=nameQuery,proto3,oneof"`
}

type SearchQuery_DomainQuery struct {
	DomainQuery *OrganizationDomainQuery `protobuf:"bytes,2,opt,name=domain_query,json=domainQuery,proto3,oneof"`
}

type SearchQuery_StateQuery struct {
	StateQuery *OrganizationStateQuery `protobuf:"bytes,3,opt,name=state_query,json=stateQuery,proto3,oneof"`
}

type SearchQuery_IdQuery struct {
	IdQuery *OrganizationIDQuery `protobuf:"bytes,4,opt,name=id_query,json=idQuery,proto3,oneof"`
}

func (*SearchQuery_NameQuery) isSearchQuery_Query() {}

func (*SearchQuery_DomainQuery) isSearchQuery_Query() {}

func (*SearchQuery_StateQuery) isSearchQuery_Query() {}

func (*SearchQuery_IdQuery) isSearchQuery_Query() {}

type OrganizationNameQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the organization.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Defines which text equality method is used.
	Method v2.TextQueryMethod `protobuf:"varint,2,opt,name=method,proto3,enum=zitadel.object.v2.TextQueryMethod" json:"method,omitempty"`
}

func (x *OrganizationNameQuery) Reset() {
	*x = OrganizationNameQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_org_v2_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationNameQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationNameQuery) ProtoMessage() {}

func (x *OrganizationNameQuery) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_org_v2_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationNameQuery.ProtoReflect.Descriptor instead.
func (*OrganizationNameQuery) Descriptor() ([]byte, []int) {
	return file_zitadel_org_v2_query_proto_rawDescGZIP(), []int{1}
}

func (x *OrganizationNameQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrganizationNameQuery) GetMethod() v2.TextQueryMethod {
	if x != nil {
		return x.Method
	}
	return v2.TextQueryMethod(0)
}

type OrganizationDomainQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Domain used in organization, not necessary primary domain.
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Defines which text equality method is used.
	Method v2.TextQueryMethod `protobuf:"varint,2,opt,name=method,proto3,enum=zitadel.object.v2.TextQueryMethod" json:"method,omitempty"`
}

func (x *OrganizationDomainQuery) Reset() {
	*x = OrganizationDomainQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_org_v2_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationDomainQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationDomainQuery) ProtoMessage() {}

func (x *OrganizationDomainQuery) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_org_v2_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationDomainQuery.ProtoReflect.Descriptor instead.
func (*OrganizationDomainQuery) Descriptor() ([]byte, []int) {
	return file_zitadel_org_v2_query_proto_rawDescGZIP(), []int{2}
}

func (x *OrganizationDomainQuery) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *OrganizationDomainQuery) GetMethod() v2.TextQueryMethod {
	if x != nil {
		return x.Method
	}
	return v2.TextQueryMethod(0)
}

type OrganizationStateQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Current state of the organization.
	State OrganizationState `protobuf:"varint,1,opt,name=state,proto3,enum=zitadel.org.v2.OrganizationState" json:"state,omitempty"`
}

func (x *OrganizationStateQuery) Reset() {
	*x = OrganizationStateQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_org_v2_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationStateQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationStateQuery) ProtoMessage() {}

func (x *OrganizationStateQuery) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_org_v2_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationStateQuery.ProtoReflect.Descriptor instead.
func (*OrganizationStateQuery) Descriptor() ([]byte, []int) {
	return file_zitadel_org_v2_query_proto_rawDescGZIP(), []int{3}
}

func (x *OrganizationStateQuery) GetState() OrganizationState {
	if x != nil {
		return x.State
	}
	return OrganizationState_ORGANIZATION_STATE_UNSPECIFIED
}

type OrganizationIDQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the organization.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrganizationIDQuery) Reset() {
	*x = OrganizationIDQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_org_v2_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationIDQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationIDQuery) ProtoMessage() {}

func (x *OrganizationIDQuery) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_org_v2_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationIDQuery.ProtoReflect.Descriptor instead.
func (*OrganizationIDQuery) Descriptor() ([]byte, []int) {
	return file_zitadel_org_v2_query_proto_rawDescGZIP(), []int{4}
}

func (x *OrganizationIDQuery) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_zitadel_org_v2_query_proto protoreflect.FileDescriptor

var file_zitadel_org_v2_query_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x76, 0x32,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x7a, 0x69,
	0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x6f, 0x72,
	0x67, 0x2f, 0x76, 0x32, 0x2f, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76,
	0x32, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe,
	0x02, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x46,
	0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x4c, 0x0a, 0x0c, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x7a,
	0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x49, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x7a, 0x69, 0x74, 0x61,
	0x64, 0x65, 0x6c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x40, 0x0a, 0x08, 0x69, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x76, 0x32, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x07, 0x69, 0x64, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x42, 0x0c, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22,
	0x99, 0x01, 0x0a, 0x15, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3a, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0x92, 0x41, 0x16, 0x4a, 0x0e, 0x22, 0x67,
	0x69, 0x67, 0x69, 0x2d, 0x67, 0x69, 0x72, 0x61, 0x66, 0x66, 0x65, 0x22, 0x78, 0xc8, 0x01, 0x80,
	0x01, 0x01, 0xe0, 0x41, 0x02, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xc8, 0x01, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0xa0, 0x01, 0x0a, 0x17,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3f, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0x92, 0x41, 0x17, 0x4a, 0x0f, 0x22, 0x63,
	0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x22, 0x78, 0xc8, 0x01,
	0x80, 0x01, 0x01, 0xe0, 0x41, 0x02, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xc8, 0x01,
	0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x44, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64,
	0x65, 0x6c, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x78,
	0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x5b,
	0x0a, 0x16, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x41, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65,
	0x6c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x52, 0x0a, 0x13, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x3b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b,
	0x92, 0x41, 0x1b, 0x4a, 0x13, 0x22, 0x36, 0x39, 0x36, 0x32, 0x39, 0x30, 0x32, 0x33, 0x39, 0x30,
	0x36, 0x34, 0x38, 0x38, 0x33, 0x33, 0x34, 0x22, 0x78, 0xc8, 0x01, 0x80, 0x01, 0x01, 0xe0, 0x41,
	0x02, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xc8, 0x01, 0x52, 0x02, 0x69, 0x64, 0x2a,
	0x62, 0x0a, 0x15, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x23, 0x4f, 0x52, 0x47, 0x41,
	0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x4e,
	0x41, 0x4d, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x20, 0x0a, 0x1c, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4e, 0x41, 0x4d,
	0x45, 0x10, 0x01, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65,
	0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x76,
	0x32, 0x3b, 0x6f, 0x72, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zitadel_org_v2_query_proto_rawDescOnce sync.Once
	file_zitadel_org_v2_query_proto_rawDescData = file_zitadel_org_v2_query_proto_rawDesc
)

func file_zitadel_org_v2_query_proto_rawDescGZIP() []byte {
	file_zitadel_org_v2_query_proto_rawDescOnce.Do(func() {
		file_zitadel_org_v2_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_zitadel_org_v2_query_proto_rawDescData)
	})
	return file_zitadel_org_v2_query_proto_rawDescData
}

var file_zitadel_org_v2_query_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_zitadel_org_v2_query_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_zitadel_org_v2_query_proto_goTypes = []interface{}{
	(OrganizationFieldName)(0),      // 0: zitadel.org.v2.OrganizationFieldName
	(*SearchQuery)(nil),             // 1: zitadel.org.v2.SearchQuery
	(*OrganizationNameQuery)(nil),   // 2: zitadel.org.v2.OrganizationNameQuery
	(*OrganizationDomainQuery)(nil), // 3: zitadel.org.v2.OrganizationDomainQuery
	(*OrganizationStateQuery)(nil),  // 4: zitadel.org.v2.OrganizationStateQuery
	(*OrganizationIDQuery)(nil),     // 5: zitadel.org.v2.OrganizationIDQuery
	(v2.TextQueryMethod)(0),         // 6: zitadel.object.v2.TextQueryMethod
	(OrganizationState)(0),          // 7: zitadel.org.v2.OrganizationState
}
var file_zitadel_org_v2_query_proto_depIdxs = []int32{
	2, // 0: zitadel.org.v2.SearchQuery.name_query:type_name -> zitadel.org.v2.OrganizationNameQuery
	3, // 1: zitadel.org.v2.SearchQuery.domain_query:type_name -> zitadel.org.v2.OrganizationDomainQuery
	4, // 2: zitadel.org.v2.SearchQuery.state_query:type_name -> zitadel.org.v2.OrganizationStateQuery
	5, // 3: zitadel.org.v2.SearchQuery.id_query:type_name -> zitadel.org.v2.OrganizationIDQuery
	6, // 4: zitadel.org.v2.OrganizationNameQuery.method:type_name -> zitadel.object.v2.TextQueryMethod
	6, // 5: zitadel.org.v2.OrganizationDomainQuery.method:type_name -> zitadel.object.v2.TextQueryMethod
	7, // 6: zitadel.org.v2.OrganizationStateQuery.state:type_name -> zitadel.org.v2.OrganizationState
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_zitadel_org_v2_query_proto_init() }
func file_zitadel_org_v2_query_proto_init() {
	if File_zitadel_org_v2_query_proto != nil {
		return
	}
	file_zitadel_org_v2_org_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_zitadel_org_v2_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_zitadel_org_v2_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationNameQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_zitadel_org_v2_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationDomainQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_zitadel_org_v2_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationStateQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_zitadel_org_v2_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationIDQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_zitadel_org_v2_query_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SearchQuery_NameQuery)(nil),
		(*SearchQuery_DomainQuery)(nil),
		(*SearchQuery_StateQuery)(nil),
		(*SearchQuery_IdQuery)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zitadel_org_v2_query_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zitadel_org_v2_query_proto_goTypes,
		DependencyIndexes: file_zitadel_org_v2_query_proto_depIdxs,
		EnumInfos:         file_zitadel_org_v2_query_proto_enumTypes,
		MessageInfos:      file_zitadel_org_v2_query_proto_msgTypes,
	}.Build()
	File_zitadel_org_v2_query_proto = out.File
	file_zitadel_org_v2_query_proto_rawDesc = nil
	file_zitadel_org_v2_query_proto_goTypes = nil
	file_zitadel_org_v2_query_proto_depIdxs = nil
}
