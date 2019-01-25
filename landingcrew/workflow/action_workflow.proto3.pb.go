// Code generated by protoc-gen-go. DO NOT EDIT.
// source: action_workflow.proto3

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

type ActionRequest struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string            `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Instructions         string            `protobuf:"bytes,3,opt,name=instructions,proto3" json:"instructions,omitempty"`
	InstructionVideo     string            `protobuf:"bytes,4,opt,name=instruction_video,json=instructionVideo,proto3" json:"instruction_video,omitempty"`
	Secrets              map[string]string `protobuf:"bytes,5,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ActionRequest) Reset()         { *m = ActionRequest{} }
func (m *ActionRequest) String() string { return proto.CompactTextString(m) }
func (*ActionRequest) ProtoMessage()    {}
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_action_workflow_f3c4336e6c184f1d, []int{0}
}
func (m *ActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionRequest.Unmarshal(m, b)
}
func (m *ActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionRequest.Marshal(b, m, deterministic)
}
func (dst *ActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionRequest.Merge(dst, src)
}
func (m *ActionRequest) XXX_Size() int {
	return xxx_messageInfo_ActionRequest.Size(m)
}
func (m *ActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActionRequest proto.InternalMessageInfo

func (m *ActionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ActionRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ActionRequest) GetInstructions() string {
	if m != nil {
		return m.Instructions
	}
	return ""
}

func (m *ActionRequest) GetInstructionVideo() string {
	if m != nil {
		return m.InstructionVideo
	}
	return ""
}

func (m *ActionRequest) GetSecrets() map[string]string {
	if m != nil {
		return m.Secrets
	}
	return nil
}

type ActionResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Notes                string   `protobuf:"bytes,3,opt,name=notes,proto3" json:"notes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActionResponse) Reset()         { *m = ActionResponse{} }
func (m *ActionResponse) String() string { return proto.CompactTextString(m) }
func (*ActionResponse) ProtoMessage()    {}
func (*ActionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_action_workflow_f3c4336e6c184f1d, []int{1}
}
func (m *ActionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionResponse.Unmarshal(m, b)
}
func (m *ActionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionResponse.Marshal(b, m, deterministic)
}
func (dst *ActionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionResponse.Merge(dst, src)
}
func (m *ActionResponse) XXX_Size() int {
	return xxx_messageInfo_ActionResponse.Size(m)
}
func (m *ActionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ActionResponse proto.InternalMessageInfo

func (m *ActionResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ActionResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ActionResponse) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

func init() {
	proto.RegisterType((*ActionRequest)(nil), "ActionRequest")
	proto.RegisterMapType((map[string]string)(nil), "ActionRequest.SecretsEntry")
	proto.RegisterType((*ActionResponse)(nil), "ActionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ActionWorkflowClient is the client API for ActionWorkflow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActionWorkflowClient interface {
	New(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	Get(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ActionWorkflow_ListClient, error)
}

type actionWorkflowClient struct {
	cc *grpc.ClientConn
}

func NewActionWorkflowClient(cc *grpc.ClientConn) ActionWorkflowClient {
	return &actionWorkflowClient{cc}
}

func (c *actionWorkflowClient) New(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionResponse, error) {
	out := new(ActionResponse)
	err := c.cc.Invoke(ctx, "/ActionWorkflow/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionWorkflowClient) Get(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionResponse, error) {
	out := new(ActionResponse)
	err := c.cc.Invoke(ctx, "/ActionWorkflow/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionWorkflowClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ActionWorkflow_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ActionWorkflow_serviceDesc.Streams[0], "/ActionWorkflow/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &actionWorkflowListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ActionWorkflow_ListClient interface {
	Recv() (*ActionResponse, error)
	grpc.ClientStream
}

type actionWorkflowListClient struct {
	grpc.ClientStream
}

func (x *actionWorkflowListClient) Recv() (*ActionResponse, error) {
	m := new(ActionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ActionWorkflowServer is the server API for ActionWorkflow service.
type ActionWorkflowServer interface {
	New(context.Context, *ActionRequest) (*ActionResponse, error)
	Get(context.Context, *ActionRequest) (*ActionResponse, error)
	List(*empty.Empty, ActionWorkflow_ListServer) error
}

func RegisterActionWorkflowServer(s *grpc.Server, srv ActionWorkflowServer) {
	s.RegisterService(&_ActionWorkflow_serviceDesc, srv)
}

func _ActionWorkflow_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionWorkflowServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionWorkflow/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionWorkflowServer).New(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionWorkflow_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionWorkflowServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionWorkflow/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionWorkflowServer).Get(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionWorkflow_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActionWorkflowServer).List(m, &actionWorkflowListServer{stream})
}

type ActionWorkflow_ListServer interface {
	Send(*ActionResponse) error
	grpc.ServerStream
}

type actionWorkflowListServer struct {
	grpc.ServerStream
}

func (x *actionWorkflowListServer) Send(m *ActionResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ActionWorkflow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ActionWorkflow",
	HandlerType: (*ActionWorkflowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _ActionWorkflow_New_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ActionWorkflow_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _ActionWorkflow_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "action_workflow.proto3",
}

func init() {
	proto.RegisterFile("action_workflow.proto3", fileDescriptor_action_workflow_f3c4336e6c184f1d)
}

var fileDescriptor_action_workflow_f3c4336e6c184f1d = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4f, 0x6b, 0xfa, 0x40,
	0x10, 0xfd, 0x6d, 0xa2, 0xfe, 0xe8, 0xd4, 0x5a, 0xbb, 0x14, 0x09, 0x7a, 0x91, 0x9c, 0x84, 0xc2,
	0x5a, 0x94, 0x42, 0xf1, 0xd6, 0x83, 0xf4, 0x52, 0x3c, 0x58, 0x68, 0x8f, 0xe2, 0x9f, 0x51, 0x16,
	0xd3, 0xac, 0xcd, 0x4e, 0x94, 0x7c, 0x94, 0x7e, 0xd4, 0xde, 0xca, 0xee, 0x26, 0x25, 0xa1, 0x3d,
	0xf4, 0x96, 0x37, 0xef, 0xbd, 0x9d, 0x37, 0xbc, 0x40, 0x67, 0xb9, 0x26, 0xa9, 0xe2, 0xc5, 0x49,
	0x25, 0xfb, 0x6d, 0xa4, 0x4e, 0xe2, 0x90, 0x28, 0x52, 0xe3, 0x6e, 0x6f, 0xa7, 0xd4, 0x2e, 0xc2,
	0xa1, 0x85, 0xab, 0x74, 0x3b, 0xc4, 0xb7, 0x03, 0x65, 0x8e, 0x0d, 0x3f, 0x19, 0x5c, 0x3c, 0x58,
	0xdf, 0x1c, 0xdf, 0x53, 0xd4, 0xc4, 0x5b, 0xe0, 0xc9, 0x4d, 0xc0, 0xfa, 0x6c, 0x70, 0x36, 0xf7,
	0xe4, 0x86, 0x5f, 0x43, 0x9d, 0x24, 0x45, 0x18, 0x78, 0x76, 0xe4, 0x00, 0x0f, 0xa1, 0x29, 0x63,
	0x4d, 0x49, 0x6a, 0xbd, 0x3a, 0xf0, 0x2d, 0x59, 0x99, 0xf1, 0x1b, 0xb8, 0x2a, 0xe1, 0xc5, 0x51,
	0x6e, 0x50, 0x05, 0x35, 0x2b, 0x6c, 0x97, 0x88, 0x17, 0x33, 0xe7, 0x77, 0xf0, 0x5f, 0xe3, 0x3a,
	0x41, 0xd2, 0x41, 0xbd, 0xef, 0x0f, 0xce, 0x47, 0x3d, 0x51, 0xc9, 0x25, 0x9e, 0x1d, 0x3b, 0x8d,
	0x29, 0xc9, 0xe6, 0x85, 0xb6, 0x3b, 0x81, 0x66, 0x99, 0xe0, 0x6d, 0xf0, 0xf7, 0x98, 0xe5, 0xf1,
	0xcd, 0xa7, 0xc9, 0x7f, 0x5c, 0x46, 0xe9, 0x77, 0x7e, 0x0b, 0x26, 0xde, 0x3d, 0x0b, 0x67, 0xd0,
	0x2a, 0x56, 0xe8, 0x83, 0x8a, 0x35, 0xfe, 0xb8, 0xbd, 0x03, 0x0d, 0x4d, 0x4b, 0x4a, 0x75, 0x6e,
	0xce, 0x91, 0x79, 0x33, 0x56, 0x84, 0xc5, 0xd9, 0x0e, 0x8c, 0x3e, 0x58, 0xf1, 0xe0, 0x6b, 0x5e,
	0x01, 0x1f, 0x80, 0x3f, 0xc3, 0x13, 0x6f, 0x55, 0x6f, 0xe9, 0x5e, 0x8a, 0xea, 0xe2, 0xf0, 0x9f,
	0x51, 0x3e, 0x22, 0xfd, 0x45, 0x39, 0x86, 0xda, 0x93, 0xd4, 0xc4, 0x3b, 0xc2, 0x15, 0x2b, 0x8a,
	0x62, 0xc5, 0xd4, 0x14, 0xfb, 0x8b, 0xe5, 0x96, 0xad, 0x1a, 0xee, 0x67, 0xf8, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0x31, 0x86, 0x54, 0xca, 0x25, 0x02, 0x00, 0x00,
}
