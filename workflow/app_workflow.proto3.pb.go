// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app_workflow.proto3

package workflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

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

type AppRequest struct {
	AppName              string            `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	AppVersion           string            `protobuf:"bytes,2,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	Fields               []*Field          `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	Environment          map[string]string `protobuf:"bytes,4,rep,name=environment,proto3" json:"environment,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Action               string            `protobuf:"bytes,5,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AppRequest) Reset()         { *m = AppRequest{} }
func (m *AppRequest) String() string { return proto.CompactTextString(m) }
func (*AppRequest) ProtoMessage()    {}
func (*AppRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_workflow_5c5a458c08e892f4, []int{0}
}
func (m *AppRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppRequest.Unmarshal(m, b)
}
func (m *AppRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppRequest.Marshal(b, m, deterministic)
}
func (dst *AppRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppRequest.Merge(dst, src)
}
func (m *AppRequest) XXX_Size() int {
	return xxx_messageInfo_AppRequest.Size(m)
}
func (m *AppRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppRequest proto.InternalMessageInfo

func (m *AppRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *AppRequest) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

func (m *AppRequest) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *AppRequest) GetEnvironment() map[string]string {
	if m != nil {
		return m.Environment
	}
	return nil
}

func (m *AppRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type AppResponse struct {
	Fields               []*Field `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	Raw                  *any.Any `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppResponse) Reset()         { *m = AppResponse{} }
func (m *AppResponse) String() string { return proto.CompactTextString(m) }
func (*AppResponse) ProtoMessage()    {}
func (*AppResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_workflow_5c5a458c08e892f4, []int{1}
}
func (m *AppResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppResponse.Unmarshal(m, b)
}
func (m *AppResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppResponse.Marshal(b, m, deterministic)
}
func (dst *AppResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppResponse.Merge(dst, src)
}
func (m *AppResponse) XXX_Size() int {
	return xxx_messageInfo_AppResponse.Size(m)
}
func (m *AppResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AppResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AppResponse proto.InternalMessageInfo

func (m *AppResponse) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *AppResponse) GetRaw() *any.Any {
	if m != nil {
		return m.Raw
	}
	return nil
}

func init() {
	proto.RegisterType((*AppRequest)(nil), "AppRequest")
	proto.RegisterMapType((map[string]string)(nil), "AppRequest.EnvironmentEntry")
	proto.RegisterType((*AppResponse)(nil), "AppResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppWorkflowClient is the client API for AppWorkflow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppWorkflowClient interface {
	Invoke(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppResponse, error)
}

type appWorkflowClient struct {
	cc *grpc.ClientConn
}

func NewAppWorkflowClient(cc *grpc.ClientConn) AppWorkflowClient {
	return &appWorkflowClient{cc}
}

func (c *appWorkflowClient) Invoke(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppResponse, error) {
	out := new(AppResponse)
	err := c.cc.Invoke(ctx, "/AppWorkflow/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppWorkflowServer is the server API for AppWorkflow service.
type AppWorkflowServer interface {
	Invoke(context.Context, *AppRequest) (*AppResponse, error)
}

func RegisterAppWorkflowServer(s *grpc.Server, srv AppWorkflowServer) {
	s.RegisterService(&_AppWorkflow_serviceDesc, srv)
}

func _AppWorkflow_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppWorkflowServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AppWorkflow/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppWorkflowServer).Invoke(ctx, req.(*AppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppWorkflow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AppWorkflow",
	HandlerType: (*AppWorkflowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _AppWorkflow_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app_workflow.proto3",
}

func init() { proto.RegisterFile("app_workflow.proto3", fileDescriptor_app_workflow_5c5a458c08e892f4) }

var fileDescriptor_app_workflow_5c5a458c08e892f4 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x75, 0xbb, 0x76, 0xd5, 0x89, 0x42, 0x89, 0x45, 0xb6, 0x45, 0xb4, 0x14, 0x94, 0x9e, 0x52,
	0xa8, 0x1e, 0xc4, 0x43, 0xa1, 0x87, 0x0a, 0x5e, 0x3c, 0x2c, 0xa8, 0x47, 0x49, 0x75, 0x5a, 0xca,
	0xee, 0x26, 0x31, 0xfb, 0xc5, 0xde, 0xfc, 0xe9, 0x92, 0x64, 0x4b, 0x17, 0xf1, 0x96, 0x97, 0x37,
	0xf3, 0xde, 0xbc, 0x19, 0x38, 0xe7, 0x4a, 0x7d, 0x54, 0x52, 0xc7, 0xeb, 0x44, 0x56, 0x4c, 0x69,
	0x99, 0xcb, 0xbb, 0xe1, 0xd9, 0x7a, 0x8b, 0xc9, 0x57, 0xb6, 0x83, 0x83, 0x8d, 0x94, 0x9b, 0x04,
	0xa7, 0x16, 0xae, 0x8a, 0xf5, 0x94, 0x8b, 0xda, 0x71, 0xe3, 0x9f, 0x0e, 0xc0, 0x42, 0xa9, 0x08,
	0xbf, 0x0b, 0xcc, 0x72, 0x3a, 0x80, 0x63, 0xa3, 0x27, 0x78, 0x8a, 0xa1, 0x37, 0xf2, 0x26, 0x27,
	0xd1, 0x11, 0x57, 0xea, 0x85, 0xa7, 0x48, 0xaf, 0x81, 0x18, 0xaa, 0x44, 0x9d, 0x6d, 0xa5, 0x08,
	0x3b, 0x96, 0x05, 0xae, 0xd4, 0x9b, 0xfb, 0xa1, 0x57, 0x10, 0x38, 0xdb, 0xd0, 0x1f, 0xf9, 0x13,
	0x32, 0x0b, 0xd8, 0x93, 0x81, 0x51, 0xf3, 0x4b, 0xe7, 0x40, 0x50, 0x94, 0x5b, 0x2d, 0x45, 0x8a,
	0x22, 0x0f, 0x0f, 0x6d, 0xd1, 0x25, 0xdb, 0xbb, 0xb3, 0xe5, 0x9e, 0x5e, 0x8a, 0x5c, 0xd7, 0x51,
	0xbb, 0x81, 0x5e, 0x40, 0xc0, 0x3f, 0x73, 0xe3, 0xdd, 0xb5, 0xde, 0x0d, 0x1a, 0xce, 0xa1, 0xf7,
	0xb7, 0x91, 0xf6, 0xc0, 0x8f, 0xb1, 0x6e, 0x22, 0x98, 0x27, 0xed, 0x43, 0xb7, 0xe4, 0x49, 0x81,
	0xcd, 0xe0, 0x0e, 0x3c, 0x76, 0x1e, 0xbc, 0xf1, 0x2b, 0x10, 0x3b, 0x43, 0xa6, 0xa4, 0xc8, 0xb0,
	0x15, 0xc3, 0xfb, 0x37, 0xc6, 0x2d, 0xf8, 0x9a, 0x57, 0x56, 0x86, 0xcc, 0xfa, 0xcc, 0xad, 0x96,
	0xed, 0x56, 0xcb, 0x16, 0xa2, 0x8e, 0x4c, 0xc1, 0xec, 0xde, 0xca, 0xbe, 0x37, 0x97, 0xa1, 0x37,
	0x10, 0x3c, 0x8b, 0x52, 0xc6, 0x48, 0x49, 0x2b, 0xf2, 0xf0, 0x94, 0xb5, 0xbc, 0xc7, 0x07, 0xab,
	0xc0, 0x9d, 0xec, 0x37, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xcc, 0xbb, 0x51, 0xd7, 0x01, 0x00, 0x00,
}
