// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: parca/share/v1alpha1/share.proto

package sharev1alpha1

import (
	v1alpha1 "github.com/parca-dev/parca/gen/proto/go/parca/query/v1alpha1"
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

// UploadRequest represents the request with profile bytes and description.
type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pprof bytes of the profile to be uploaded.
	Profile []byte `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	// Description of the profile.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_share_v1alpha1_share_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parca_share_v1alpha1_share_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_parca_share_v1alpha1_share_proto_rawDescGZIP(), []int{0}
}

func (x *UploadRequest) GetProfile() []byte {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *UploadRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// UploadResponse represents the response with the link that can be used to access the profile.
type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the uploaded profile.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// link that can be used to access the profile.
	Link string `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_share_v1alpha1_share_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parca_share_v1alpha1_share_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_parca_share_v1alpha1_share_proto_rawDescGZIP(), []int{1}
}

func (x *UploadResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UploadResponse) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

// QueryRequest represents the request with the id of the profile to be queried.
type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the profile to be queried.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Type of the profile to be queried.
	ProfileType *string `protobuf:"bytes,2,opt,name=profile_type,json=profileType,proto3,oneof" json:"profile_type,omitempty"`
	// report_type is the type of report to return
	ReportType v1alpha1.QueryRequest_ReportType `protobuf:"varint,3,opt,name=report_type,json=reportType,proto3,enum=parca.query.v1alpha1.QueryRequest_ReportType" json:"report_type,omitempty"`
	// filter_query is the query string to filter the profile samples
	FilterQuery *string `protobuf:"bytes,4,opt,name=filter_query,json=filterQuery,proto3,oneof" json:"filter_query,omitempty"`
	// node_trim_threshold is the threshold % where the nodes with Value less than this will be removed from the report
	NodeTrimThreshold *float32 `protobuf:"fixed32,5,opt,name=node_trim_threshold,json=nodeTrimThreshold,proto3,oneof" json:"node_trim_threshold,omitempty"`
	// which runtime frames to filter out, often interpreter frames like python or ruby are not super useful by default
	RuntimeFilter *v1alpha1.RuntimeFilter `protobuf:"bytes,6,opt,name=runtime_filter,json=runtimeFilter,proto3,oneof" json:"runtime_filter,omitempty"`
	// group_by indicates the fields to group by
	GroupBy *v1alpha1.GroupBy `protobuf:"bytes,7,opt,name=group_by,json=groupBy,proto3,oneof" json:"group_by,omitempty"`
	// invert_call_stack inverts the call stacks in the flamegraph
	InvertCallStack *bool `protobuf:"varint,9,opt,name=invert_call_stack,json=invertCallStack,proto3,oneof" json:"invert_call_stack,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_share_v1alpha1_share_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parca_share_v1alpha1_share_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_parca_share_v1alpha1_share_proto_rawDescGZIP(), []int{2}
}

func (x *QueryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *QueryRequest) GetProfileType() string {
	if x != nil && x.ProfileType != nil {
		return *x.ProfileType
	}
	return ""
}

func (x *QueryRequest) GetReportType() v1alpha1.QueryRequest_ReportType {
	if x != nil {
		return x.ReportType
	}
	return v1alpha1.QueryRequest_ReportType(0)
}

func (x *QueryRequest) GetFilterQuery() string {
	if x != nil && x.FilterQuery != nil {
		return *x.FilterQuery
	}
	return ""
}

func (x *QueryRequest) GetNodeTrimThreshold() float32 {
	if x != nil && x.NodeTrimThreshold != nil {
		return *x.NodeTrimThreshold
	}
	return 0
}

func (x *QueryRequest) GetRuntimeFilter() *v1alpha1.RuntimeFilter {
	if x != nil {
		return x.RuntimeFilter
	}
	return nil
}

func (x *QueryRequest) GetGroupBy() *v1alpha1.GroupBy {
	if x != nil {
		return x.GroupBy
	}
	return nil
}

func (x *QueryRequest) GetInvertCallStack() bool {
	if x != nil && x.InvertCallStack != nil {
		return *x.InvertCallStack
	}
	return false
}

// ProfileTypesRequest represents the profile types request with the id of the profile to be queried.
type ProfileTypesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the profile's types to be queried.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProfileTypesRequest) Reset() {
	*x = ProfileTypesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_share_v1alpha1_share_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileTypesRequest) ProtoMessage() {}

func (x *ProfileTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parca_share_v1alpha1_share_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileTypesRequest.ProtoReflect.Descriptor instead.
func (*ProfileTypesRequest) Descriptor() ([]byte, []int) {
	return file_parca_share_v1alpha1_share_proto_rawDescGZIP(), []int{3}
}

func (x *ProfileTypesRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// ProfileTypesResponse represents the response with the list of available profile types.
type ProfileTypesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of available profile types.
	Types []*v1alpha1.ProfileType `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty"`
	// description of the profile uploaded.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ProfileTypesResponse) Reset() {
	*x = ProfileTypesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_share_v1alpha1_share_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileTypesResponse) ProtoMessage() {}

func (x *ProfileTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parca_share_v1alpha1_share_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileTypesResponse.ProtoReflect.Descriptor instead.
func (*ProfileTypesResponse) Descriptor() ([]byte, []int) {
	return file_parca_share_v1alpha1_share_proto_rawDescGZIP(), []int{4}
}

func (x *ProfileTypesResponse) GetTypes() []*v1alpha1.ProfileType {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *ProfileTypesResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// QueryResponse is the returned report for the given query.
type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// report is the generated report
	//
	// Types that are assignable to Report:
	//
	//	*QueryResponse_Flamegraph
	//	*QueryResponse_Pprof
	//	*QueryResponse_Top
	//	*QueryResponse_Callgraph
	//	*QueryResponse_FlamegraphArrow
	//	*QueryResponse_Source
	//	*QueryResponse_TableArrow
	Report isQueryResponse_Report `protobuf_oneof:"report"`
	// total is the total number of samples shown in the report.
	Total int64 `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	// filtered is the number of samples filtered out of the report.
	Filtered int64 `protobuf:"varint,6,opt,name=filtered,proto3" json:"filtered,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_share_v1alpha1_share_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parca_share_v1alpha1_share_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_parca_share_v1alpha1_share_proto_rawDescGZIP(), []int{5}
}

func (m *QueryResponse) GetReport() isQueryResponse_Report {
	if m != nil {
		return m.Report
	}
	return nil
}

func (x *QueryResponse) GetFlamegraph() *v1alpha1.Flamegraph {
	if x, ok := x.GetReport().(*QueryResponse_Flamegraph); ok {
		return x.Flamegraph
	}
	return nil
}

func (x *QueryResponse) GetPprof() []byte {
	if x, ok := x.GetReport().(*QueryResponse_Pprof); ok {
		return x.Pprof
	}
	return nil
}

func (x *QueryResponse) GetTop() *v1alpha1.Top {
	if x, ok := x.GetReport().(*QueryResponse_Top); ok {
		return x.Top
	}
	return nil
}

func (x *QueryResponse) GetCallgraph() *v1alpha1.Callgraph {
	if x, ok := x.GetReport().(*QueryResponse_Callgraph); ok {
		return x.Callgraph
	}
	return nil
}

func (x *QueryResponse) GetFlamegraphArrow() *v1alpha1.FlamegraphArrow {
	if x, ok := x.GetReport().(*QueryResponse_FlamegraphArrow); ok {
		return x.FlamegraphArrow
	}
	return nil
}

func (x *QueryResponse) GetSource() *v1alpha1.Source {
	if x, ok := x.GetReport().(*QueryResponse_Source); ok {
		return x.Source
	}
	return nil
}

func (x *QueryResponse) GetTableArrow() *v1alpha1.TableArrow {
	if x, ok := x.GetReport().(*QueryResponse_TableArrow); ok {
		return x.TableArrow
	}
	return nil
}

func (x *QueryResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *QueryResponse) GetFiltered() int64 {
	if x != nil {
		return x.Filtered
	}
	return 0
}

type isQueryResponse_Report interface {
	isQueryResponse_Report()
}

type QueryResponse_Flamegraph struct {
	// flamegraph is a flamegraph representation of the report
	Flamegraph *v1alpha1.Flamegraph `protobuf:"bytes,1,opt,name=flamegraph,proto3,oneof"`
}

type QueryResponse_Pprof struct {
	// pprof is a pprof profile as compressed bytes
	Pprof []byte `protobuf:"bytes,2,opt,name=pprof,proto3,oneof"`
}

type QueryResponse_Top struct {
	// top is a top list representation of the report
	Top *v1alpha1.Top `protobuf:"bytes,3,opt,name=top,proto3,oneof"`
}

type QueryResponse_Callgraph struct {
	// callgraph is a callgraph nodes and edges representation of the report
	Callgraph *v1alpha1.Callgraph `protobuf:"bytes,4,opt,name=callgraph,proto3,oneof"`
}

type QueryResponse_FlamegraphArrow struct {
	// flamegraph_arrow is a flamegraph encoded as a arrow record
	FlamegraphArrow *v1alpha1.FlamegraphArrow `protobuf:"bytes,7,opt,name=flamegraph_arrow,json=flamegraphArrow,proto3,oneof"`
}

type QueryResponse_Source struct {
	// source is the source report type result
	Source *v1alpha1.Source `protobuf:"bytes,8,opt,name=source,proto3,oneof"`
}

type QueryResponse_TableArrow struct {
	// table_arrow is a table encoded as a arrow record
	TableArrow *v1alpha1.TableArrow `protobuf:"bytes,9,opt,name=table_arrow,json=tableArrow,proto3,oneof"`
}

func (*QueryResponse_Flamegraph) isQueryResponse_Report() {}

func (*QueryResponse_Pprof) isQueryResponse_Report() {}

func (*QueryResponse_Top) isQueryResponse_Report() {}

func (*QueryResponse_Callgraph) isQueryResponse_Report() {}

func (*QueryResponse_FlamegraphArrow) isQueryResponse_Report() {}

func (*QueryResponse_Source) isQueryResponse_Report() {}

func (*QueryResponse_TableArrow) isQueryResponse_Report() {}

var File_parca_share_v1alpha1_share_proto protoreflect.FileDescriptor

var file_parca_share_v1alpha1_share_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x20, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x0d, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x34, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0xa4, 0x04,
	0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4e, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x70, 0x61,
	0x72, 0x63, 0x61, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x33,
	0x0a, 0x13, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x72, 0x69, 0x6d, 0x5f, 0x74, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x48, 0x02, 0x52, 0x11, 0x6e,
	0x6f, 0x64, 0x65, 0x54, 0x72, 0x69, 0x6d, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x4f, 0x0a, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x61,
	0x72, 0x63, 0x61, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x48, 0x03, 0x52, 0x0d, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x42, 0x79, 0x48, 0x04, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79,
	0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x11, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x5f, 0x63, 0x61,
	0x6c, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x48, 0x05,
	0x52, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x74, 0x72, 0x69, 0x6d, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x42, 0x11,
	0x0a, 0x0f, 0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x42, 0x14,
	0x0a, 0x12, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x22, 0x25, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x71, 0x0a, 0x14, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe8,
	0x03, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x42, 0x0a, 0x0a, 0x66, 0x6c, 0x61, 0x6d, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x6c, 0x61, 0x6d,
	0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x6c, 0x61, 0x6d, 0x65, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x12, 0x16, 0x0a, 0x05, 0x70, 0x70, 0x72, 0x6f, 0x66, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x70, 0x70, 0x72, 0x6f, 0x66, 0x12, 0x2d, 0x0a, 0x03,
	0x74, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x61, 0x72, 0x63,
	0x61, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x54, 0x6f, 0x70, 0x48, 0x00, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x3f, 0x0a, 0x09, 0x63,
	0x61, 0x6c, 0x6c, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x67, 0x72, 0x61, 0x70, 0x68, 0x48,
	0x00, 0x52, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x67, 0x72, 0x61, 0x70, 0x68, 0x12, 0x52, 0x0a, 0x10,
	0x66, 0x6c, 0x61, 0x6d, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x61, 0x72, 0x72, 0x6f, 0x77,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x6c,
	0x61, 0x6d, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x48, 0x00, 0x52,
	0x0f, 0x66, 0x6c, 0x61, 0x6d, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x41, 0x72, 0x72, 0x6f, 0x77,
	0x12, 0x36, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x48,
	0x00, 0x52, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x42,
	0x08, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x32, 0xa2, 0x02, 0x0a, 0x0c, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x06, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x61, 0x72, 0x63,
	0x61, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x52, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x22, 0x2e, 0x70, 0x61, 0x72,
	0x63, 0x61, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x29, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xe4,
	0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2d, 0x64, 0x65, 0x76, 0x2f,
	0x70, 0x61, 0x72, 0x63, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x53, 0x58, 0xaa, 0x02, 0x14, 0x50, 0x61,
	0x72, 0x63, 0x61, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xca, 0x02, 0x14, 0x50, 0x61, 0x72, 0x63, 0x61, 0x5c, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x20, 0x50, 0x61, 0x72, 0x63,
	0x61, 0x5c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x50,
	0x61, 0x72, 0x63, 0x61, 0x3a, 0x3a, 0x53, 0x68, 0x61, 0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_parca_share_v1alpha1_share_proto_rawDescOnce sync.Once
	file_parca_share_v1alpha1_share_proto_rawDescData = file_parca_share_v1alpha1_share_proto_rawDesc
)

func file_parca_share_v1alpha1_share_proto_rawDescGZIP() []byte {
	file_parca_share_v1alpha1_share_proto_rawDescOnce.Do(func() {
		file_parca_share_v1alpha1_share_proto_rawDescData = protoimpl.X.CompressGZIP(file_parca_share_v1alpha1_share_proto_rawDescData)
	})
	return file_parca_share_v1alpha1_share_proto_rawDescData
}

var file_parca_share_v1alpha1_share_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_parca_share_v1alpha1_share_proto_goTypes = []interface{}{
	(*UploadRequest)(nil),                 // 0: parca.share.v1alpha1.UploadRequest
	(*UploadResponse)(nil),                // 1: parca.share.v1alpha1.UploadResponse
	(*QueryRequest)(nil),                  // 2: parca.share.v1alpha1.QueryRequest
	(*ProfileTypesRequest)(nil),           // 3: parca.share.v1alpha1.ProfileTypesRequest
	(*ProfileTypesResponse)(nil),          // 4: parca.share.v1alpha1.ProfileTypesResponse
	(*QueryResponse)(nil),                 // 5: parca.share.v1alpha1.QueryResponse
	(v1alpha1.QueryRequest_ReportType)(0), // 6: parca.query.v1alpha1.QueryRequest.ReportType
	(*v1alpha1.RuntimeFilter)(nil),        // 7: parca.query.v1alpha1.RuntimeFilter
	(*v1alpha1.GroupBy)(nil),              // 8: parca.query.v1alpha1.GroupBy
	(*v1alpha1.ProfileType)(nil),          // 9: parca.query.v1alpha1.ProfileType
	(*v1alpha1.Flamegraph)(nil),           // 10: parca.query.v1alpha1.Flamegraph
	(*v1alpha1.Top)(nil),                  // 11: parca.query.v1alpha1.Top
	(*v1alpha1.Callgraph)(nil),            // 12: parca.query.v1alpha1.Callgraph
	(*v1alpha1.FlamegraphArrow)(nil),      // 13: parca.query.v1alpha1.FlamegraphArrow
	(*v1alpha1.Source)(nil),               // 14: parca.query.v1alpha1.Source
	(*v1alpha1.TableArrow)(nil),           // 15: parca.query.v1alpha1.TableArrow
}
var file_parca_share_v1alpha1_share_proto_depIdxs = []int32{
	6,  // 0: parca.share.v1alpha1.QueryRequest.report_type:type_name -> parca.query.v1alpha1.QueryRequest.ReportType
	7,  // 1: parca.share.v1alpha1.QueryRequest.runtime_filter:type_name -> parca.query.v1alpha1.RuntimeFilter
	8,  // 2: parca.share.v1alpha1.QueryRequest.group_by:type_name -> parca.query.v1alpha1.GroupBy
	9,  // 3: parca.share.v1alpha1.ProfileTypesResponse.types:type_name -> parca.query.v1alpha1.ProfileType
	10, // 4: parca.share.v1alpha1.QueryResponse.flamegraph:type_name -> parca.query.v1alpha1.Flamegraph
	11, // 5: parca.share.v1alpha1.QueryResponse.top:type_name -> parca.query.v1alpha1.Top
	12, // 6: parca.share.v1alpha1.QueryResponse.callgraph:type_name -> parca.query.v1alpha1.Callgraph
	13, // 7: parca.share.v1alpha1.QueryResponse.flamegraph_arrow:type_name -> parca.query.v1alpha1.FlamegraphArrow
	14, // 8: parca.share.v1alpha1.QueryResponse.source:type_name -> parca.query.v1alpha1.Source
	15, // 9: parca.share.v1alpha1.QueryResponse.table_arrow:type_name -> parca.query.v1alpha1.TableArrow
	0,  // 10: parca.share.v1alpha1.ShareService.Upload:input_type -> parca.share.v1alpha1.UploadRequest
	2,  // 11: parca.share.v1alpha1.ShareService.Query:input_type -> parca.share.v1alpha1.QueryRequest
	3,  // 12: parca.share.v1alpha1.ShareService.ProfileTypes:input_type -> parca.share.v1alpha1.ProfileTypesRequest
	1,  // 13: parca.share.v1alpha1.ShareService.Upload:output_type -> parca.share.v1alpha1.UploadResponse
	5,  // 14: parca.share.v1alpha1.ShareService.Query:output_type -> parca.share.v1alpha1.QueryResponse
	4,  // 15: parca.share.v1alpha1.ShareService.ProfileTypes:output_type -> parca.share.v1alpha1.ProfileTypesResponse
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_parca_share_v1alpha1_share_proto_init() }
func file_parca_share_v1alpha1_share_proto_init() {
	if File_parca_share_v1alpha1_share_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_parca_share_v1alpha1_share_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_parca_share_v1alpha1_share_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResponse); i {
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
		file_parca_share_v1alpha1_share_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_parca_share_v1alpha1_share_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileTypesRequest); i {
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
		file_parca_share_v1alpha1_share_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileTypesResponse); i {
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
		file_parca_share_v1alpha1_share_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResponse); i {
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
	file_parca_share_v1alpha1_share_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_parca_share_v1alpha1_share_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*QueryResponse_Flamegraph)(nil),
		(*QueryResponse_Pprof)(nil),
		(*QueryResponse_Top)(nil),
		(*QueryResponse_Callgraph)(nil),
		(*QueryResponse_FlamegraphArrow)(nil),
		(*QueryResponse_Source)(nil),
		(*QueryResponse_TableArrow)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_parca_share_v1alpha1_share_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_parca_share_v1alpha1_share_proto_goTypes,
		DependencyIndexes: file_parca_share_v1alpha1_share_proto_depIdxs,
		MessageInfos:      file_parca_share_v1alpha1_share_proto_msgTypes,
	}.Build()
	File_parca_share_v1alpha1_share_proto = out.File
	file_parca_share_v1alpha1_share_proto_rawDesc = nil
	file_parca_share_v1alpha1_share_proto_goTypes = nil
	file_parca_share_v1alpha1_share_proto_depIdxs = nil
}
