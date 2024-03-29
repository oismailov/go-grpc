// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data_workflow.proto3

package workflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamSource int32

const (
	StreamSource_CSV       StreamSource = 0
	StreamSource_Streaming StreamSource = 1
	StreamSource_Airtable  StreamSource = 2
)

var StreamSource_name = map[int32]string{
	0: "CSV",
	1: "Streaming",
	2: "Airtable",
}
var StreamSource_value = map[string]int32{
	"CSV":       0,
	"Streaming": 1,
	"Airtable":  2,
}

func (x StreamSource) String() string {
	return proto.EnumName(StreamSource_name, int32(x))
}
func (StreamSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_data_workflow_a5435ca8839226cd, []int{0}
}

type DataType int32

const (
	DataType_Collect            DataType = 0
	DataType_Clean              DataType = 1
	DataType_CollectContactInfo DataType = 2
)

var DataType_name = map[int32]string{
	0: "Collect",
	1: "Clean",
	2: "CollectContactInfo",
}
var DataType_value = map[string]int32{
	"Collect":            0,
	"Clean":              1,
	"CollectContactInfo": 2,
}

func (x DataType) String() string {
	return proto.EnumName(DataType_name, int32(x))
}
func (DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_data_workflow_a5435ca8839226cd, []int{1}
}

type DataRequest struct {
	Id                   int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 DataType          `protobuf:"varint,2,opt,name=type,proto3,enum=DataType" json:"type,omitempty"`
	Title                string            `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Instructions         string            `protobuf:"bytes,4,opt,name=instructions,proto3" json:"instructions,omitempty"`
	InstructionVideo     string            `protobuf:"bytes,5,opt,name=instructionVideo,proto3" json:"instructionVideo,omitempty"`
	Secrets              map[string]string `protobuf:"bytes,6,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DataRequest) Reset()         { *m = DataRequest{} }
func (m *DataRequest) String() string { return proto.CompactTextString(m) }
func (*DataRequest) ProtoMessage()    {}
func (*DataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_workflow_a5435ca8839226cd, []int{0}
}
func (m *DataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataRequest.Unmarshal(m, b)
}
func (m *DataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataRequest.Marshal(b, m, deterministic)
}
func (dst *DataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRequest.Merge(dst, src)
}
func (m *DataRequest) XXX_Size() int {
	return xxx_messageInfo_DataRequest.Size(m)
}
func (m *DataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataRequest proto.InternalMessageInfo

func (m *DataRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DataRequest) GetType() DataType {
	if m != nil {
		return m.Type
	}
	return DataType_Collect
}

func (m *DataRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *DataRequest) GetInstructions() string {
	if m != nil {
		return m.Instructions
	}
	return ""
}

func (m *DataRequest) GetInstructionVideo() string {
	if m != nil {
		return m.InstructionVideo
	}
	return ""
}

func (m *DataRequest) GetSecrets() map[string]string {
	if m != nil {
		return m.Secrets
	}
	return nil
}

type DataResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields               []*Field `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataResponse) Reset()         { *m = DataResponse{} }
func (m *DataResponse) String() string { return proto.CompactTextString(m) }
func (*DataResponse) ProtoMessage()    {}
func (*DataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_workflow_a5435ca8839226cd, []int{1}
}
func (m *DataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataResponse.Unmarshal(m, b)
}
func (m *DataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataResponse.Marshal(b, m, deterministic)
}
func (dst *DataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataResponse.Merge(dst, src)
}
func (m *DataResponse) XXX_Size() int {
	return xxx_messageInfo_DataResponse.Size(m)
}
func (m *DataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataResponse proto.InternalMessageInfo

func (m *DataResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DataResponse) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

type DataRecord struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Fields               []*Field `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataRecord) Reset()         { *m = DataRecord{} }
func (m *DataRecord) String() string { return proto.CompactTextString(m) }
func (*DataRecord) ProtoMessage()    {}
func (*DataRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_workflow_a5435ca8839226cd, []int{2}
}
func (m *DataRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataRecord.Unmarshal(m, b)
}
func (m *DataRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataRecord.Marshal(b, m, deterministic)
}
func (dst *DataRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRecord.Merge(dst, src)
}
func (m *DataRecord) XXX_Size() int {
	return xxx_messageInfo_DataRecord.Size(m)
}
func (m *DataRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRecord.DiscardUnknown(m)
}

var xxx_messageInfo_DataRecord proto.InternalMessageInfo

func (m *DataRecord) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DataRecord) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DataRecord) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterType((*DataRequest)(nil), "DataRequest")
	proto.RegisterMapType((map[string]string)(nil), "DataRequest.SecretsEntry")
	proto.RegisterType((*DataResponse)(nil), "DataResponse")
	proto.RegisterType((*DataRecord)(nil), "DataRecord")
	proto.RegisterEnum("StreamSource", StreamSource_name, StreamSource_value)
	proto.RegisterEnum("DataType", DataType_name, DataType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataWorkflowClient is the client API for DataWorkflow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataWorkflowClient interface {
	New(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DataResponse, error)
	Get(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DataResponse, error)
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (DataWorkflow_ListClient, error)
	Sync(ctx context.Context, in *DataRecord, opts ...grpc.CallOption) (*DataRecord, error)
}

type dataWorkflowClient struct {
	cc *grpc.ClientConn
}

func NewDataWorkflowClient(cc *grpc.ClientConn) DataWorkflowClient {
	return &dataWorkflowClient{cc}
}

func (c *dataWorkflowClient) New(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/DataWorkflow/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataWorkflowClient) Get(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/DataWorkflow/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataWorkflowClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (DataWorkflow_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataWorkflow_serviceDesc.Streams[0], "/DataWorkflow/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataWorkflowListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataWorkflow_ListClient interface {
	Recv() (*DataResponse, error)
	grpc.ClientStream
}

type dataWorkflowListClient struct {
	grpc.ClientStream
}

func (x *dataWorkflowListClient) Recv() (*DataResponse, error) {
	m := new(DataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataWorkflowClient) Sync(ctx context.Context, in *DataRecord, opts ...grpc.CallOption) (*DataRecord, error) {
	out := new(DataRecord)
	err := c.cc.Invoke(ctx, "/DataWorkflow/Sync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataWorkflowServer is the server API for DataWorkflow service.
type DataWorkflowServer interface {
	New(context.Context, *DataRequest) (*DataResponse, error)
	Get(context.Context, *DataRequest) (*DataResponse, error)
	List(*empty.Empty, DataWorkflow_ListServer) error
	Sync(context.Context, *DataRecord) (*DataRecord, error)
}

func RegisterDataWorkflowServer(s *grpc.Server, srv DataWorkflowServer) {
	s.RegisterService(&_DataWorkflow_serviceDesc, srv)
}

func _DataWorkflow_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataWorkflowServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataWorkflow/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataWorkflowServer).New(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataWorkflow_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataWorkflowServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataWorkflow/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataWorkflowServer).Get(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataWorkflow_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataWorkflowServer).List(m, &dataWorkflowListServer{stream})
}

type DataWorkflow_ListServer interface {
	Send(*DataResponse) error
	grpc.ServerStream
}

type dataWorkflowListServer struct {
	grpc.ServerStream
}

func (x *dataWorkflowListServer) Send(m *DataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DataWorkflow_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataWorkflowServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataWorkflow/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataWorkflowServer).Sync(ctx, req.(*DataRecord))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataWorkflow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DataWorkflow",
	HandlerType: (*DataWorkflowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _DataWorkflow_New_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DataWorkflow_Get_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _DataWorkflow_Sync_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _DataWorkflow_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "data_workflow.proto3",
}

func init() { proto.RegisterFile("data_workflow.proto3", fileDescriptor_data_workflow_a5435ca8839226cd) }

var fileDescriptor_data_workflow_a5435ca8839226cd = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x6e, 0x92, 0xfe, 0x6c, 0x4e, 0xd3, 0x25, 0x0c, 0x4b, 0x89, 0x15, 0xa5, 0x04, 0x2f, 0x4a,
	0x2f, 0x66, 0xb5, 0xeb, 0x85, 0xf4, 0x42, 0x90, 0xba, 0x8a, 0x20, 0x5e, 0xa4, 0xcb, 0x7a, 0x29,
	0xd3, 0xe4, 0xb4, 0x0c, 0x9b, 0xcd, 0xc4, 0xcc, 0x89, 0x25, 0xcf, 0xe1, 0x93, 0xf8, 0x86, 0x92,
	0x4c, 0x8a, 0x91, 0x0a, 0xde, 0xe5, 0xfb, 0x39, 0x93, 0x33, 0xdf, 0x7c, 0x70, 0x95, 0x08, 0x12,
	0xdf, 0x8e, 0xaa, 0x78, 0xd8, 0xa7, 0xea, 0xc8, 0xf3, 0x42, 0x91, 0xba, 0x99, 0x3d, 0x3d, 0x28,
	0x75, 0x48, 0xf1, 0xba, 0x81, 0xbb, 0x72, 0x7f, 0x8d, 0x8f, 0x39, 0x55, 0x46, 0x9d, 0x4d, 0xf6,
	0x12, 0xd3, 0x44, 0xb7, 0xde, 0xf0, 0xa7, 0x0d, 0xe3, 0xf7, 0x82, 0x44, 0x84, 0xdf, 0x4b, 0xd4,
	0xc4, 0x2e, 0xc1, 0x96, 0x49, 0x60, 0xcd, 0xad, 0x85, 0x13, 0xd9, 0x32, 0x61, 0xcf, 0xa0, 0x4f,
	0x55, 0x8e, 0x81, 0x3d, 0xb7, 0x16, 0x97, 0x2b, 0x97, 0xd7, 0xde, 0xbb, 0x2a, 0xc7, 0xa8, 0xa1,
	0xd9, 0x15, 0x0c, 0x48, 0x52, 0x8a, 0x81, 0x33, 0xb7, 0x16, 0x6e, 0x64, 0x00, 0x0b, 0xc1, 0x93,
	0x99, 0xa6, 0xa2, 0x8c, 0x49, 0xaa, 0x4c, 0x07, 0xfd, 0x46, 0xfc, 0x8b, 0x63, 0x4b, 0xf0, 0x3b,
	0xf8, 0x5e, 0x26, 0xa8, 0x82, 0x41, 0xe3, 0x3b, 0xe3, 0xd9, 0x0d, 0x8c, 0x34, 0xc6, 0x05, 0x92,
	0x0e, 0x86, 0x73, 0x67, 0x31, 0x5e, 0x3d, 0xe1, 0x9d, 0x9d, 0xf9, 0xd6, 0x68, 0xb7, 0x19, 0x15,
	0x55, 0x74, 0x72, 0xce, 0xd6, 0xe0, 0x75, 0x05, 0xe6, 0x83, 0xf3, 0x80, 0x55, 0x73, 0x35, 0x37,
	0xaa, 0x3f, 0xeb, 0xe5, 0x7f, 0x88, 0xb4, 0x34, 0x97, 0x73, 0x23, 0x03, 0xd6, 0xf6, 0x1b, 0x2b,
	0x7c, 0x0b, 0x9e, 0xf9, 0x81, 0xce, 0x55, 0xa6, 0xf1, 0x2c, 0x95, 0xe7, 0x30, 0x34, 0x31, 0x06,
	0x76, 0xb3, 0xcf, 0x90, 0x7f, 0xa8, 0x61, 0xd4, 0xb2, 0xe1, 0x1d, 0x80, 0x99, 0x8f, 0x55, 0x91,
	0x9c, 0x4d, 0x4f, 0x61, 0xa8, 0x49, 0x50, 0xa9, 0xdb, 0xd4, 0x5a, 0xd4, 0x39, 0xb5, 0xff, 0xaf,
	0x53, 0x97, 0xaf, 0xc1, 0xdb, 0x52, 0x81, 0xe2, 0x71, 0xab, 0xca, 0x22, 0x46, 0x36, 0x02, 0x67,
	0xb3, 0xbd, 0xf7, 0x7b, 0x6c, 0x02, 0xae, 0x11, 0x64, 0x76, 0xf0, 0x2d, 0xe6, 0xc1, 0xc5, 0x3b,
	0x59, 0x90, 0xd8, 0xa5, 0xe8, 0xdb, 0xcb, 0x35, 0x5c, 0x9c, 0x1e, 0x8d, 0x8d, 0x61, 0xb4, 0x51,
	0x69, 0x8a, 0x31, 0xf9, 0x3d, 0xe6, 0xc2, 0x60, 0x93, 0xa2, 0xc8, 0x7c, 0x8b, 0x4d, 0x81, 0xb5,
	0xfc, 0x46, 0x65, 0x24, 0x62, 0xfa, 0x94, 0xed, 0x95, 0x6f, 0xaf, 0x7e, 0x59, 0x26, 0x88, 0xaf,
	0x6d, 0xc3, 0xd8, 0x0b, 0x70, 0xbe, 0xe0, 0x91, 0x79, 0xdd, 0xfc, 0x67, 0x13, 0xde, 0x0d, 0x2b,
	0xec, 0xd5, 0xae, 0x8f, 0x48, 0xff, 0x73, 0xbd, 0x82, 0xfe, 0x67, 0xa9, 0x89, 0x4d, 0xb9, 0xe9,
	0x2b, 0x3f, 0xf5, 0x95, 0xdf, 0xd6, 0x7d, 0x3d, 0x1b, 0x78, 0x69, 0xb1, 0x10, 0xfa, 0xdb, 0x2a,
	0x8b, 0xd9, 0x98, 0xff, 0x89, 0x77, 0xd6, 0x05, 0x61, 0x6f, 0x37, 0x34, 0xcd, 0xfe, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x4b, 0x5d, 0x61, 0x6c, 0x1c, 0x03, 0x00, 0x00,
}
