// Code generated by protoc-gen-go. DO NOT EDIT.
// source: code_workflow.proto3

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

type CodeType int32

const (
	CodeType_RUBY_LAMBDA       CodeType = 0
	CodeType_NODEJS_LAMBDA     CodeType = 1
	CodeType_PYTHON_LAMBDA     CodeType = 2
	CodeType_REACTJS           CodeType = 3
	CodeType_REACTJS_COMPONENT CodeType = 4
)

var CodeType_name = map[int32]string{
	0: "RUBY_LAMBDA",
	1: "NODEJS_LAMBDA",
	2: "PYTHON_LAMBDA",
	3: "REACTJS",
	4: "REACTJS_COMPONENT",
}
var CodeType_value = map[string]int32{
	"RUBY_LAMBDA":       0,
	"NODEJS_LAMBDA":     1,
	"PYTHON_LAMBDA":     2,
	"REACTJS":           3,
	"REACTJS_COMPONENT": 4,
}

func (x CodeType) String() string {
	return proto.EnumName(CodeType_name, int32(x))
}
func (CodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_code_workflow_9478c0a876ac4e10, []int{0}
}

type CodeDecisionType int32

const (
	CodeDecisionType_APPROVE CodeDecisionType = 0
	CodeDecisionType_REJECT  CodeDecisionType = 1
)

var CodeDecisionType_name = map[int32]string{
	0: "APPROVE",
	1: "REJECT",
}
var CodeDecisionType_value = map[string]int32{
	"APPROVE": 0,
	"REJECT":  1,
}

func (x CodeDecisionType) String() string {
	return proto.EnumName(CodeDecisionType_name, int32(x))
}
func (CodeDecisionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_code_workflow_9478c0a876ac4e10, []int{1}
}

type CodeRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 CodeType `protobuf:"varint,2,opt,name=type,proto3,enum=CodeType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeRequest) Reset()         { *m = CodeRequest{} }
func (m *CodeRequest) String() string { return proto.CompactTextString(m) }
func (*CodeRequest) ProtoMessage()    {}
func (*CodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_workflow_9478c0a876ac4e10, []int{0}
}
func (m *CodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeRequest.Unmarshal(m, b)
}
func (m *CodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeRequest.Marshal(b, m, deterministic)
}
func (dst *CodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeRequest.Merge(dst, src)
}
func (m *CodeRequest) XXX_Size() int {
	return xxx_messageInfo_CodeRequest.Size(m)
}
func (m *CodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CodeRequest proto.InternalMessageInfo

func (m *CodeRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CodeRequest) GetType() CodeType {
	if m != nil {
		return m.Type
	}
	return CodeType_RUBY_LAMBDA
}

type CodeRequest_Github struct {
	AuthToken            string   `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	FullRepo             string   `protobuf:"bytes,2,opt,name=full_repo,json=fullRepo,proto3" json:"full_repo,omitempty"`
	IssueNumber          string   `protobuf:"bytes,3,opt,name=issue_number,json=issueNumber,proto3" json:"issue_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeRequest_Github) Reset()         { *m = CodeRequest_Github{} }
func (m *CodeRequest_Github) String() string { return proto.CompactTextString(m) }
func (*CodeRequest_Github) ProtoMessage()    {}
func (*CodeRequest_Github) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_workflow_9478c0a876ac4e10, []int{0, 0}
}
func (m *CodeRequest_Github) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeRequest_Github.Unmarshal(m, b)
}
func (m *CodeRequest_Github) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeRequest_Github.Marshal(b, m, deterministic)
}
func (dst *CodeRequest_Github) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeRequest_Github.Merge(dst, src)
}
func (m *CodeRequest_Github) XXX_Size() int {
	return xxx_messageInfo_CodeRequest_Github.Size(m)
}
func (m *CodeRequest_Github) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeRequest_Github.DiscardUnknown(m)
}

var xxx_messageInfo_CodeRequest_Github proto.InternalMessageInfo

func (m *CodeRequest_Github) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

func (m *CodeRequest_Github) GetFullRepo() string {
	if m != nil {
		return m.FullRepo
	}
	return ""
}

func (m *CodeRequest_Github) GetIssueNumber() string {
	if m != nil {
		return m.IssueNumber
	}
	return ""
}

type CodeResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeResponse) Reset()         { *m = CodeResponse{} }
func (m *CodeResponse) String() string { return proto.CompactTextString(m) }
func (*CodeResponse) ProtoMessage()    {}
func (*CodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_workflow_9478c0a876ac4e10, []int{1}
}
func (m *CodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeResponse.Unmarshal(m, b)
}
func (m *CodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeResponse.Marshal(b, m, deterministic)
}
func (dst *CodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeResponse.Merge(dst, src)
}
func (m *CodeResponse) XXX_Size() int {
	return xxx_messageInfo_CodeResponse.Size(m)
}
func (m *CodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CodeResponse proto.InternalMessageInfo

func (m *CodeResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CodeResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type CodeResponse_Github struct {
	PullRequest          string   `protobuf:"bytes,1,opt,name=pull_request,json=pullRequest,proto3" json:"pull_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeResponse_Github) Reset()         { *m = CodeResponse_Github{} }
func (m *CodeResponse_Github) String() string { return proto.CompactTextString(m) }
func (*CodeResponse_Github) ProtoMessage()    {}
func (*CodeResponse_Github) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_workflow_9478c0a876ac4e10, []int{1, 0}
}
func (m *CodeResponse_Github) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeResponse_Github.Unmarshal(m, b)
}
func (m *CodeResponse_Github) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeResponse_Github.Marshal(b, m, deterministic)
}
func (dst *CodeResponse_Github) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeResponse_Github.Merge(dst, src)
}
func (m *CodeResponse_Github) XXX_Size() int {
	return xxx_messageInfo_CodeResponse_Github.Size(m)
}
func (m *CodeResponse_Github) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeResponse_Github.DiscardUnknown(m)
}

var xxx_messageInfo_CodeResponse_Github proto.InternalMessageInfo

func (m *CodeResponse_Github) GetPullRequest() string {
	if m != nil {
		return m.PullRequest
	}
	return ""
}

type CodeDecision struct {
	Id                   int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Decison              CodeDecisionType `protobuf:"varint,2,opt,name=decison,proto3,enum=CodeDecisionType" json:"decison,omitempty"`
	Body                 string           `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CodeDecision) Reset()         { *m = CodeDecision{} }
func (m *CodeDecision) String() string { return proto.CompactTextString(m) }
func (*CodeDecision) ProtoMessage()    {}
func (*CodeDecision) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_workflow_9478c0a876ac4e10, []int{2}
}
func (m *CodeDecision) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeDecision.Unmarshal(m, b)
}
func (m *CodeDecision) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeDecision.Marshal(b, m, deterministic)
}
func (dst *CodeDecision) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeDecision.Merge(dst, src)
}
func (m *CodeDecision) XXX_Size() int {
	return xxx_messageInfo_CodeDecision.Size(m)
}
func (m *CodeDecision) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeDecision.DiscardUnknown(m)
}

var xxx_messageInfo_CodeDecision proto.InternalMessageInfo

func (m *CodeDecision) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CodeDecision) GetDecison() CodeDecisionType {
	if m != nil {
		return m.Decison
	}
	return CodeDecisionType_APPROVE
}

func (m *CodeDecision) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*CodeRequest)(nil), "CodeRequest")
	proto.RegisterType((*CodeRequest_Github)(nil), "CodeRequest.Github")
	proto.RegisterType((*CodeResponse)(nil), "CodeResponse")
	proto.RegisterType((*CodeResponse_Github)(nil), "CodeResponse.Github")
	proto.RegisterType((*CodeDecision)(nil), "CodeDecision")
	proto.RegisterEnum("CodeType", CodeType_name, CodeType_value)
	proto.RegisterEnum("CodeDecisionType", CodeDecisionType_name, CodeDecisionType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CodeWorkflowClient is the client API for CodeWorkflow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CodeWorkflowClient interface {
	New(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*CodeResponse, error)
	Get(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*CodeResponse, error)
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (CodeWorkflow_ListClient, error)
	Decide(ctx context.Context, in *CodeDecision, opts ...grpc.CallOption) (*empty.Empty, error)
}

type codeWorkflowClient struct {
	cc *grpc.ClientConn
}

func NewCodeWorkflowClient(cc *grpc.ClientConn) CodeWorkflowClient {
	return &codeWorkflowClient{cc}
}

func (c *codeWorkflowClient) New(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*CodeResponse, error) {
	out := new(CodeResponse)
	err := c.cc.Invoke(ctx, "/CodeWorkflow/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeWorkflowClient) Get(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*CodeResponse, error) {
	out := new(CodeResponse)
	err := c.cc.Invoke(ctx, "/CodeWorkflow/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeWorkflowClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (CodeWorkflow_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CodeWorkflow_serviceDesc.Streams[0], "/CodeWorkflow/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &codeWorkflowListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CodeWorkflow_ListClient interface {
	Recv() (*CodeResponse, error)
	grpc.ClientStream
}

type codeWorkflowListClient struct {
	grpc.ClientStream
}

func (x *codeWorkflowListClient) Recv() (*CodeResponse, error) {
	m := new(CodeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *codeWorkflowClient) Decide(ctx context.Context, in *CodeDecision, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/CodeWorkflow/Decide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeWorkflowServer is the server API for CodeWorkflow service.
type CodeWorkflowServer interface {
	New(context.Context, *CodeRequest) (*CodeResponse, error)
	Get(context.Context, *CodeRequest) (*CodeResponse, error)
	List(*empty.Empty, CodeWorkflow_ListServer) error
	Decide(context.Context, *CodeDecision) (*empty.Empty, error)
}

func RegisterCodeWorkflowServer(s *grpc.Server, srv CodeWorkflowServer) {
	s.RegisterService(&_CodeWorkflow_serviceDesc, srv)
}

func _CodeWorkflow_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeWorkflowServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CodeWorkflow/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeWorkflowServer).New(ctx, req.(*CodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeWorkflow_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeWorkflowServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CodeWorkflow/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeWorkflowServer).Get(ctx, req.(*CodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeWorkflow_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CodeWorkflowServer).List(m, &codeWorkflowListServer{stream})
}

type CodeWorkflow_ListServer interface {
	Send(*CodeResponse) error
	grpc.ServerStream
}

type codeWorkflowListServer struct {
	grpc.ServerStream
}

func (x *codeWorkflowListServer) Send(m *CodeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CodeWorkflow_Decide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeDecision)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeWorkflowServer).Decide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CodeWorkflow/Decide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeWorkflowServer).Decide(ctx, req.(*CodeDecision))
	}
	return interceptor(ctx, in, info, handler)
}

var _CodeWorkflow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CodeWorkflow",
	HandlerType: (*CodeWorkflowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _CodeWorkflow_New_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CodeWorkflow_Get_Handler,
		},
		{
			MethodName: "Decide",
			Handler:    _CodeWorkflow_Decide_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _CodeWorkflow_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "code_workflow.proto3",
}

func init() { proto.RegisterFile("code_workflow.proto3", fileDescriptor_code_workflow_9478c0a876ac4e10) }

var fileDescriptor_code_workflow_9478c0a876ac4e10 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xef, 0x6a, 0xda, 0x50,
	0x18, 0xc6, 0x8d, 0x4a, 0xda, 0xbc, 0xd1, 0x2e, 0x1e, 0xb6, 0x22, 0x29, 0x85, 0x36, 0xec, 0x43,
	0xa9, 0x90, 0xae, 0xdd, 0x15, 0x58, 0x0d, 0x1d, 0xd2, 0x26, 0x72, 0x9a, 0x6d, 0xf4, 0x53, 0x68,
	0xcc, 0xab, 0x0d, 0xd5, 0x9c, 0xcc, 0x73, 0x82, 0x78, 0x41, 0xbb, 0x8f, 0x5d, 0xda, 0x38, 0x27,
	0xc9, 0x70, 0x93, 0xd1, 0x6f, 0x79, 0x7f, 0xef, 0x9f, 0x3c, 0xe7, 0xe1, 0x81, 0xf7, 0x33, 0x96,
	0x60, 0xb4, 0x61, 0xeb, 0xd7, 0xf9, 0x92, 0x6d, 0xdc, 0x7c, 0xcd, 0x04, 0xfb, 0x6c, 0x9f, 0x2c,
	0x18, 0x5b, 0x2c, 0xf1, 0x4a, 0x95, 0x71, 0x31, 0xbf, 0xc2, 0x55, 0x2e, 0xb6, 0x65, 0xd7, 0xf9,
	0xa9, 0x81, 0x39, 0x62, 0x09, 0x52, 0xfc, 0x51, 0x20, 0x17, 0xe4, 0x08, 0x9a, 0x69, 0xd2, 0xd7,
	0xce, 0xb4, 0x8b, 0x16, 0x6d, 0xa6, 0x09, 0x39, 0x85, 0xb6, 0xd8, 0xe6, 0xd8, 0x6f, 0x9e, 0x69,
	0x17, 0x47, 0x37, 0x86, 0x2b, 0x67, 0xc3, 0x6d, 0x8e, 0x54, 0x61, 0x7b, 0x01, 0xfa, 0x5d, 0x2a,
	0x5e, 0x8a, 0x98, 0x9c, 0x02, 0x3c, 0x17, 0xe2, 0x25, 0x12, 0xec, 0x15, 0x33, 0x75, 0xc0, 0xa0,
	0x86, 0x24, 0xa1, 0x04, 0xe4, 0x04, 0x8c, 0x79, 0xb1, 0x5c, 0x46, 0x6b, 0xcc, 0x99, 0x3a, 0x66,
	0xd0, 0x43, 0x09, 0x28, 0xe6, 0x8c, 0x9c, 0x43, 0x27, 0xe5, 0xbc, 0xc0, 0x28, 0x2b, 0x56, 0x31,
	0xae, 0xfb, 0x2d, 0xd5, 0x37, 0x15, 0xf3, 0x15, 0x72, 0x66, 0xd0, 0x29, 0x65, 0xf2, 0x9c, 0x65,
	0x1c, 0xf7, 0x74, 0x1e, 0x83, 0xce, 0xc5, 0xb3, 0x28, 0x78, 0xb5, 0x5c, 0x55, 0xf6, 0xe0, 0x8f,
	0xc0, 0x73, 0xe8, 0xe4, 0xa5, 0x02, 0xf5, 0xd2, 0x4a, 0xa2, 0x99, 0x2b, 0x11, 0x0a, 0x39, 0x51,
	0xf9, 0x93, 0x31, 0xce, 0x52, 0x9e, 0xb2, 0x6c, 0xef, 0x27, 0x03, 0x38, 0x48, 0x64, 0x8f, 0x65,
	0x95, 0x1f, 0x3d, 0x77, 0x77, 0x5e, 0xf9, 0x52, 0x4f, 0x10, 0x02, 0xed, 0x98, 0x25, 0xdb, 0x4a,
	0x8f, 0xfa, 0xbe, 0x44, 0x38, 0xac, 0x0d, 0x24, 0xef, 0xc0, 0xa4, 0x5f, 0x6f, 0x9f, 0xa2, 0xfb,
	0xe1, 0xc3, 0xed, 0x78, 0x68, 0x35, 0x48, 0x0f, 0xba, 0x7e, 0x30, 0xf6, 0x26, 0x8f, 0x35, 0xd2,
	0x24, 0x9a, 0x3e, 0x85, 0x5f, 0x02, 0xbf, 0x46, 0x4d, 0x62, 0xc2, 0x01, 0xf5, 0x86, 0xa3, 0x70,
	0xf2, 0x68, 0xb5, 0xc8, 0x07, 0xe8, 0x55, 0x45, 0x34, 0x0a, 0x1e, 0xa6, 0x81, 0xef, 0xf9, 0xa1,
	0xd5, 0xbe, 0x1c, 0x80, 0xf5, 0xaf, 0x2e, 0xb9, 0x37, 0x9c, 0x4e, 0x69, 0xf0, 0xcd, 0xb3, 0x1a,
	0x04, 0x40, 0xa7, 0xde, 0xc4, 0x1b, 0x85, 0x96, 0x76, 0xf3, 0x4b, 0x2b, 0x5f, 0xfd, 0xbd, 0x8a,
	0x0d, 0xf9, 0x08, 0x2d, 0x1f, 0x37, 0xa4, 0xe3, 0xee, 0xe4, 0xc2, 0xee, 0xba, 0xbb, 0xf6, 0x3b,
	0x0d, 0x39, 0x75, 0x87, 0xe2, 0xad, 0xa9, 0x6b, 0x68, 0xdf, 0xa7, 0x5c, 0x90, 0x63, 0xb7, 0x0c,
	0xa1, 0x5b, 0x87, 0xd0, 0xf5, 0x64, 0x08, 0xf7, 0x16, 0x3e, 0x69, 0xe4, 0x1a, 0x74, 0x29, 0x3c,
	0x41, 0xd2, 0xfd, 0xcb, 0x5d, 0xfb, 0x3f, 0x37, 0x9c, 0x46, 0xac, 0x97, 0x49, 0xff, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x8e, 0x9a, 0xc0, 0xeb, 0x00, 0x03, 0x00, 0x00,
}
