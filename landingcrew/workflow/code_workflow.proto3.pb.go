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
	return fileDescriptor_code_workflow_132c04d5f01be5c8, []int{0}
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
	return fileDescriptor_code_workflow_132c04d5f01be5c8, []int{1}
}

type CodeRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 CodeType `protobuf:"varint,2,opt,name=type,proto3,enum=CodeType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeRequest) Reset()         { *m = CodeRequest{} }
func (m *CodeRequest) String() string { return proto.CompactTextString(m) }
func (*CodeRequest) ProtoMessage()    {}
func (*CodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_workflow_132c04d5f01be5c8, []int{0}
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

func (m *CodeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CodeRequest) GetType() CodeType {
	if m != nil {
		return m.Type
	}
	return CodeType_RUBY_LAMBDA
}

type CodeResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Identifier           string   `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeResponse) Reset()         { *m = CodeResponse{} }
func (m *CodeResponse) String() string { return proto.CompactTextString(m) }
func (*CodeResponse) ProtoMessage()    {}
func (*CodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_workflow_132c04d5f01be5c8, []int{1}
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

func (m *CodeResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CodeResponse) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *CodeResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type CodeResponse_Github struct {
	Repo                 string   `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeResponse_Github) Reset()         { *m = CodeResponse_Github{} }
func (m *CodeResponse_Github) String() string { return proto.CompactTextString(m) }
func (*CodeResponse_Github) ProtoMessage()    {}
func (*CodeResponse_Github) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_workflow_132c04d5f01be5c8, []int{1, 0}
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

func (m *CodeResponse_Github) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

type CodeDecision struct {
	Id                   string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
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
	return fileDescriptor_code_workflow_132c04d5f01be5c8, []int{2}
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

func (m *CodeDecision) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
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

func init() { proto.RegisterFile("code_workflow.proto3", fileDescriptor_code_workflow_132c04d5f01be5c8) }

var fileDescriptor_code_workflow_132c04d5f01be5c8 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xb5, 0x6c, 0xa3, 0xd4, 0xe3, 0x38, 0x95, 0x97, 0x36, 0x18, 0xf7, 0x83, 0x20, 0x7a, 0x08,
	0x09, 0x6c, 0x9a, 0xf4, 0xda, 0x8b, 0x22, 0x8b, 0x14, 0x93, 0x48, 0x62, 0xa3, 0xb6, 0xe4, 0x24,
	0xaa, 0xec, 0x38, 0x5d, 0x9a, 0x6a, 0x55, 0x69, 0x85, 0xd1, 0xbf, 0xeb, 0x4f, 0x2b, 0xab, 0x0f,
	0x50, 0x6b, 0x4a, 0x6e, 0x3b, 0x6f, 0xde, 0xcc, 0xbc, 0xd9, 0x79, 0xf0, 0xe2, 0x5e, 0x72, 0x8c,
	0xb7, 0x32, 0xff, 0xb1, 0x79, 0x94, 0x5b, 0x9a, 0xe5, 0x52, 0xc9, 0x0f, 0xcb, 0x57, 0x0f, 0x52,
	0x3e, 0x3c, 0xe2, 0x59, 0x1d, 0x26, 0xe5, 0xe6, 0x0c, 0x7f, 0x66, 0xaa, 0x6a, 0xb2, 0xf6, 0x47,
	0x98, 0xba, 0x92, 0x23, 0xc3, 0x5f, 0x25, 0x16, 0x8a, 0x1c, 0xc0, 0x50, 0xf0, 0x85, 0x71, 0x64,
	0x1c, 0x4f, 0xd8, 0x50, 0x70, 0xf2, 0x06, 0xc6, 0xaa, 0xca, 0x70, 0x31, 0x3c, 0x32, 0x8e, 0x0f,
	0x2e, 0x26, 0x54, 0x73, 0xa3, 0x2a, 0x43, 0x56, 0xc3, 0xb6, 0x82, 0xfd, 0xa6, 0xba, 0xc8, 0x64,
	0x5a, 0xe0, 0x4e, 0xf9, 0x5b, 0x00, 0xc1, 0x31, 0x55, 0x62, 0x23, 0x30, 0xaf, 0x9b, 0x4c, 0x58,
	0x0f, 0x21, 0x87, 0x60, 0x16, 0xea, 0x9b, 0x2a, 0x8b, 0xc5, 0xa8, 0xce, 0xb5, 0xd1, 0xf2, 0x35,
	0x98, 0x57, 0x42, 0x7d, 0x2f, 0x13, 0x42, 0x60, 0x9c, 0x63, 0x26, 0xdb, 0x9e, 0xf5, 0xdb, 0x8e,
	0x9b, 0xa9, 0x2b, 0xbc, 0x17, 0x85, 0x90, 0xe9, 0xce, 0xd4, 0x53, 0xd8, 0xe3, 0x3a, 0x27, 0xd3,
	0x56, 0xf7, 0x9c, 0xf6, 0xf9, 0xb5, 0xfe, 0x8e, 0xa1, 0x07, 0x24, 0x92, 0x57, 0xad, 0x80, 0xfa,
	0x7d, 0x82, 0xf0, 0xac, 0x5b, 0x94, 0x3c, 0x87, 0x29, 0xfb, 0x7c, 0x79, 0x17, 0x5f, 0x3b, 0x37,
	0x97, 0x2b, 0xc7, 0x1a, 0x90, 0x39, 0xcc, 0xfc, 0x60, 0xe5, 0xad, 0x6f, 0x3b, 0xc8, 0xd0, 0x50,
	0x78, 0x17, 0x7d, 0x0a, 0xfc, 0x0e, 0x1a, 0x92, 0x29, 0xec, 0x31, 0xcf, 0x71, 0xa3, 0xf5, 0xad,
	0x35, 0x22, 0x2f, 0x61, 0xde, 0x06, 0xb1, 0x1b, 0xdc, 0x84, 0x81, 0xef, 0xf9, 0x91, 0x35, 0x3e,
	0x39, 0x05, 0xeb, 0x5f, 0x5d, 0xba, 0xce, 0x09, 0x43, 0x16, 0x7c, 0xf1, 0xac, 0x01, 0x01, 0x30,
	0x99, 0xb7, 0xf6, 0xdc, 0xc8, 0x32, 0x2e, 0x7e, 0x1b, 0xcd, 0xd6, 0x5f, 0xdb, 0xeb, 0x92, 0x77,
	0x30, 0xf2, 0x71, 0x4b, 0xf6, 0x69, 0xef, 0x7e, 0xcb, 0x19, 0xed, 0xdf, 0xc3, 0x1e, 0x68, 0xd6,
	0x15, 0xaa, 0xa7, 0x58, 0xe7, 0x30, 0xbe, 0x16, 0x85, 0x22, 0x87, 0xb4, 0xf1, 0x0a, 0xed, 0xbc,
	0x42, 0x3d, 0xed, 0x95, 0x9d, 0x82, 0xf7, 0x06, 0x39, 0x07, 0x53, 0x0b, 0xe7, 0x48, 0x66, 0x7f,
	0xfd, 0xee, 0xf2, 0x3f, 0x3d, 0xec, 0x41, 0x62, 0x36, 0x86, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff,
	0xe0, 0x4d, 0xff, 0x2f, 0xa7, 0x02, 0x00, 0x00,
}
