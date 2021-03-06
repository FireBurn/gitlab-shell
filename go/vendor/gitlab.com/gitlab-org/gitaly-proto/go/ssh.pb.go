// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssh.proto

package gitaly

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SSHUploadPackRequest struct {
	// 'repository' must be present in the first message.
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// A chunk of raw data to be copied to 'git upload-pack' standard input
	Stdin []byte `protobuf:"bytes,2,opt,name=stdin,proto3" json:"stdin,omitempty"`
}

func (m *SSHUploadPackRequest) Reset()                    { *m = SSHUploadPackRequest{} }
func (m *SSHUploadPackRequest) String() string            { return proto.CompactTextString(m) }
func (*SSHUploadPackRequest) ProtoMessage()               {}
func (*SSHUploadPackRequest) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *SSHUploadPackRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *SSHUploadPackRequest) GetStdin() []byte {
	if m != nil {
		return m.Stdin
	}
	return nil
}

type SSHUploadPackResponse struct {
	// A chunk of raw data from 'git upload-pack' standard output
	Stdout []byte `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	// A chunk of raw data from 'git upload-pack' standard error
	Stderr []byte `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	// This field may be nil. This is intentional: only when the remote
	// command has finished can we return its exit status.
	ExitStatus *ExitStatus `protobuf:"bytes,3,opt,name=exit_status,json=exitStatus" json:"exit_status,omitempty"`
}

func (m *SSHUploadPackResponse) Reset()                    { *m = SSHUploadPackResponse{} }
func (m *SSHUploadPackResponse) String() string            { return proto.CompactTextString(m) }
func (*SSHUploadPackResponse) ProtoMessage()               {}
func (*SSHUploadPackResponse) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *SSHUploadPackResponse) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *SSHUploadPackResponse) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func (m *SSHUploadPackResponse) GetExitStatus() *ExitStatus {
	if m != nil {
		return m.ExitStatus
	}
	return nil
}

type SSHReceivePackRequest struct {
	// 'repository' must be present in the first message.
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// A chunk of raw data to be copied to 'git upload-pack' standard input
	Stdin []byte `protobuf:"bytes,2,opt,name=stdin,proto3" json:"stdin,omitempty"`
	// Contents of GL_ID and GL_REPOSITORY environment variables for
	// 'git receive-pack'
	GlId         string `protobuf:"bytes,3,opt,name=gl_id,json=glId" json:"gl_id,omitempty"`
	GlRepository string `protobuf:"bytes,4,opt,name=gl_repository,json=glRepository" json:"gl_repository,omitempty"`
}

func (m *SSHReceivePackRequest) Reset()                    { *m = SSHReceivePackRequest{} }
func (m *SSHReceivePackRequest) String() string            { return proto.CompactTextString(m) }
func (*SSHReceivePackRequest) ProtoMessage()               {}
func (*SSHReceivePackRequest) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *SSHReceivePackRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *SSHReceivePackRequest) GetStdin() []byte {
	if m != nil {
		return m.Stdin
	}
	return nil
}

func (m *SSHReceivePackRequest) GetGlId() string {
	if m != nil {
		return m.GlId
	}
	return ""
}

func (m *SSHReceivePackRequest) GetGlRepository() string {
	if m != nil {
		return m.GlRepository
	}
	return ""
}

type SSHReceivePackResponse struct {
	// A chunk of raw data from 'git receive-pack' standard output
	Stdout []byte `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	// A chunk of raw data from 'git receive-pack' standard error
	Stderr []byte `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	// This field may be nil. This is intentional: only when the remote
	// command has finished can we return its exit status.
	ExitStatus *ExitStatus `protobuf:"bytes,3,opt,name=exit_status,json=exitStatus" json:"exit_status,omitempty"`
}

func (m *SSHReceivePackResponse) Reset()                    { *m = SSHReceivePackResponse{} }
func (m *SSHReceivePackResponse) String() string            { return proto.CompactTextString(m) }
func (*SSHReceivePackResponse) ProtoMessage()               {}
func (*SSHReceivePackResponse) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *SSHReceivePackResponse) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *SSHReceivePackResponse) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func (m *SSHReceivePackResponse) GetExitStatus() *ExitStatus {
	if m != nil {
		return m.ExitStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*SSHUploadPackRequest)(nil), "gitaly.SSHUploadPackRequest")
	proto.RegisterType((*SSHUploadPackResponse)(nil), "gitaly.SSHUploadPackResponse")
	proto.RegisterType((*SSHReceivePackRequest)(nil), "gitaly.SSHReceivePackRequest")
	proto.RegisterType((*SSHReceivePackResponse)(nil), "gitaly.SSHReceivePackResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SSHService service

type SSHServiceClient interface {
	// To forward 'git upload-pack' to Gitaly for SSH sessions
	SSHUploadPack(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHUploadPackClient, error)
	// To forward 'git receive-pack' to Gitaly for SSH sessions
	SSHReceivePack(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHReceivePackClient, error)
}

type sSHServiceClient struct {
	cc *grpc.ClientConn
}

func NewSSHServiceClient(cc *grpc.ClientConn) SSHServiceClient {
	return &sSHServiceClient{cc}
}

func (c *sSHServiceClient) SSHUploadPack(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHUploadPackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SSHService_serviceDesc.Streams[0], c.cc, "/gitaly.SSHService/SSHUploadPack", opts...)
	if err != nil {
		return nil, err
	}
	x := &sSHServiceSSHUploadPackClient{stream}
	return x, nil
}

type SSHService_SSHUploadPackClient interface {
	Send(*SSHUploadPackRequest) error
	Recv() (*SSHUploadPackResponse, error)
	grpc.ClientStream
}

type sSHServiceSSHUploadPackClient struct {
	grpc.ClientStream
}

func (x *sSHServiceSSHUploadPackClient) Send(m *SSHUploadPackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sSHServiceSSHUploadPackClient) Recv() (*SSHUploadPackResponse, error) {
	m := new(SSHUploadPackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sSHServiceClient) SSHReceivePack(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHReceivePackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SSHService_serviceDesc.Streams[1], c.cc, "/gitaly.SSHService/SSHReceivePack", opts...)
	if err != nil {
		return nil, err
	}
	x := &sSHServiceSSHReceivePackClient{stream}
	return x, nil
}

type SSHService_SSHReceivePackClient interface {
	Send(*SSHReceivePackRequest) error
	Recv() (*SSHReceivePackResponse, error)
	grpc.ClientStream
}

type sSHServiceSSHReceivePackClient struct {
	grpc.ClientStream
}

func (x *sSHServiceSSHReceivePackClient) Send(m *SSHReceivePackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sSHServiceSSHReceivePackClient) Recv() (*SSHReceivePackResponse, error) {
	m := new(SSHReceivePackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SSHService service

type SSHServiceServer interface {
	// To forward 'git upload-pack' to Gitaly for SSH sessions
	SSHUploadPack(SSHService_SSHUploadPackServer) error
	// To forward 'git receive-pack' to Gitaly for SSH sessions
	SSHReceivePack(SSHService_SSHReceivePackServer) error
}

func RegisterSSHServiceServer(s *grpc.Server, srv SSHServiceServer) {
	s.RegisterService(&_SSHService_serviceDesc, srv)
}

func _SSHService_SSHUploadPack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SSHServiceServer).SSHUploadPack(&sSHServiceSSHUploadPackServer{stream})
}

type SSHService_SSHUploadPackServer interface {
	Send(*SSHUploadPackResponse) error
	Recv() (*SSHUploadPackRequest, error)
	grpc.ServerStream
}

type sSHServiceSSHUploadPackServer struct {
	grpc.ServerStream
}

func (x *sSHServiceSSHUploadPackServer) Send(m *SSHUploadPackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sSHServiceSSHUploadPackServer) Recv() (*SSHUploadPackRequest, error) {
	m := new(SSHUploadPackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SSHService_SSHReceivePack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SSHServiceServer).SSHReceivePack(&sSHServiceSSHReceivePackServer{stream})
}

type SSHService_SSHReceivePackServer interface {
	Send(*SSHReceivePackResponse) error
	Recv() (*SSHReceivePackRequest, error)
	grpc.ServerStream
}

type sSHServiceSSHReceivePackServer struct {
	grpc.ServerStream
}

func (x *sSHServiceSSHReceivePackServer) Send(m *SSHReceivePackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sSHServiceSSHReceivePackServer) Recv() (*SSHReceivePackRequest, error) {
	m := new(SSHReceivePackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SSHService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.SSHService",
	HandlerType: (*SSHServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SSHUploadPack",
			Handler:       _SSHService_SSHUploadPack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SSHReceivePack",
			Handler:       _SSHService_SSHReceivePack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ssh.proto",
}

func init() { proto.RegisterFile("ssh.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0xcf, 0x4e, 0x32, 0x31,
	0x10, 0xff, 0xfa, 0x09, 0x24, 0x0c, 0x8b, 0x87, 0x11, 0x09, 0x21, 0x6a, 0xc8, 0x7a, 0xe1, 0x44,
	0x0c, 0x3c, 0x83, 0x09, 0xde, 0x4c, 0x1b, 0xce, 0xb8, 0xd2, 0xc9, 0xd2, 0xd8, 0xd0, 0xb5, 0x2d,
	0x04, 0x12, 0x7d, 0x12, 0x1f, 0xc4, 0xd7, 0x33, 0xe9, 0xae, 0xb8, 0x8b, 0x72, 0xd4, 0xdb, 0xce,
	0xfc, 0x76, 0x7e, 0x7f, 0x3a, 0x03, 0x4d, 0xe7, 0x96, 0xa3, 0xcc, 0x1a, 0x6f, 0xb0, 0x91, 0x2a,
	0x9f, 0xe8, 0x5d, 0x3f, 0x72, 0xcb, 0xc4, 0x92, 0xcc, 0xbb, 0xf1, 0x03, 0x74, 0x84, 0x98, 0xce,
	0x32, 0x6d, 0x12, 0x79, 0x9f, 0x2c, 0x9e, 0x38, 0x3d, 0xaf, 0xc9, 0x79, 0x1c, 0x03, 0x58, 0xca,
	0x8c, 0x53, 0xde, 0xd8, 0x5d, 0x8f, 0x0d, 0xd8, 0xb0, 0x35, 0xc6, 0x51, 0x4e, 0x31, 0xe2, 0x7b,
	0x84, 0x97, 0xfe, 0xc2, 0x0e, 0xd4, 0x9d, 0x97, 0x6a, 0xd5, 0xfb, 0x3f, 0x60, 0xc3, 0x88, 0xe7,
	0x45, 0xfc, 0x02, 0xe7, 0x07, 0x0a, 0x2e, 0x33, 0x2b, 0x47, 0xd8, 0x85, 0x86, 0xf3, 0xd2, 0xac,
	0x7d, 0xa0, 0x8f, 0x78, 0x51, 0x15, 0x7d, 0xb2, 0xb6, 0xe0, 0x29, 0x2a, 0x9c, 0x40, 0x8b, 0xb6,
	0xca, 0xcf, 0x9d, 0x4f, 0xfc, 0xda, 0xf5, 0x4e, 0xaa, 0x9e, 0x6e, 0xb7, 0xca, 0x8b, 0x80, 0x70,
	0xa0, 0xfd, 0x77, 0xfc, 0xc6, 0x82, 0x3c, 0xa7, 0x05, 0xa9, 0x0d, 0xfd, 0x4a, 0x42, 0x3c, 0x83,
	0x7a, 0xaa, 0xe7, 0x4a, 0x06, 0x4b, 0x4d, 0x5e, 0x4b, 0xf5, 0x9d, 0xc4, 0x6b, 0x68, 0xa7, 0x7a,
	0x5e, 0x52, 0xa8, 0x05, 0x30, 0x4a, 0xf5, 0x17, 0x77, 0xfc, 0x0a, 0xdd, 0x43, 0x73, 0x7f, 0xf8,
	0x38, 0xe3, 0x77, 0x06, 0x20, 0xc4, 0x54, 0x90, 0xdd, 0xa8, 0x05, 0x21, 0x87, 0x76, 0x65, 0x53,
	0x78, 0xf1, 0x39, 0xff, 0xd3, 0x89, 0xf4, 0x2f, 0x8f, 0xa0, 0x79, 0x82, 0xf8, 0xdf, 0x90, 0xdd,
	0x30, 0x9c, 0xc1, 0x69, 0x35, 0x21, 0x96, 0xc7, 0xbe, 0xaf, 0xa5, 0x7f, 0x75, 0x0c, 0x2e, 0xd3,
	0x3e, 0x36, 0xc2, 0xf5, 0x4e, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x79, 0x5b, 0x32, 0x2b, 0xe0,
	0x02, 0x00, 0x00,
}
