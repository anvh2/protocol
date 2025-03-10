// Code generated by protoc-gen-psrpc v0.6.0, DO NOT EDIT.
// source: rpc/agent.proto

package rpc

import (
	"context"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/client"
	"github.com/livekit/psrpc/pkg/info"
	"github.com/livekit/psrpc/pkg/rand"
	"github.com/livekit/psrpc/pkg/server"
	"github.com/livekit/psrpc/version"
)
import google_protobuf "google.golang.org/protobuf/types/known/emptypb"
import livekit2 "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_6

// ==============================
// AgentInternal Client Interface
// ==============================

type AgentInternalClient interface {
	CheckEnabled(ctx context.Context, req *livekit2.CheckEnabledRequest, opts ...psrpc.RequestOption) (<-chan *psrpc.Response[*livekit2.CheckEnabledResponse], error)

	CheckAvailability(ctx context.Context, req *livekit2.CheckAvailabilityRequest, opts ...psrpc.RequestOption) (<-chan *psrpc.Response[*livekit2.CheckAvailabilityResponse], error)

	JobRequest(ctx context.Context, namespace string, jobType string, req *livekit2.Job, opts ...psrpc.RequestOption) (*JobRequestResponse, error)

	JobTerminate(ctx context.Context, jobId string, req *JobTerminateRequest, opts ...psrpc.RequestOption) (*JobTerminateResponse, error)

	SubscribeWorkerRegistered(ctx context.Context, handlerNamespace string) (psrpc.Subscription[*google_protobuf.Empty], error)

	// Close immediately, without waiting for pending RPCs
	Close()
}

// ==================================
// AgentInternal ServerImpl Interface
// ==================================

type AgentInternalServerImpl interface {
	CheckEnabled(context.Context, *livekit2.CheckEnabledRequest) (*livekit2.CheckEnabledResponse, error)

	CheckAvailability(context.Context, *livekit2.CheckAvailabilityRequest) (*livekit2.CheckAvailabilityResponse, error)

	JobRequest(context.Context, *livekit2.Job) (*JobRequestResponse, error)
	JobRequestAffinity(context.Context, *livekit2.Job) float32

	JobTerminate(context.Context, *JobTerminateRequest) (*JobTerminateResponse, error)
}

// ==============================
// AgentInternal Server Interface
// ==============================

type AgentInternalServer interface {
	RegisterJobRequestTopic(namespace string, jobType string) error
	DeregisterJobRequestTopic(namespace string, jobType string)
	RegisterJobTerminateTopic(jobId string) error
	DeregisterJobTerminateTopic(jobId string)
	PublishWorkerRegistered(ctx context.Context, handlerNamespace string, msg *google_protobuf.Empty) error

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ====================
// AgentInternal Client
// ====================

type agentInternalClient struct {
	client *client.RPCClient
}

// NewAgentInternalClient creates a psrpc client that implements the AgentInternalClient interface.
func NewAgentInternalClient(bus psrpc.MessageBus, opts ...psrpc.ClientOption) (AgentInternalClient, error) {
	sd := &info.ServiceDefinition{
		Name: "AgentInternal",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("CheckEnabled", false, true, false, false)
	sd.RegisterMethod("CheckAvailability", false, true, false, false)
	sd.RegisterMethod("JobRequest", true, false, true, false)
	sd.RegisterMethod("JobTerminate", false, false, true, true)
	sd.RegisterMethod("WorkerRegistered", false, true, false, false)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &agentInternalClient{
		client: rpcClient,
	}, nil
}

func (c *agentInternalClient) CheckEnabled(ctx context.Context, req *livekit2.CheckEnabledRequest, opts ...psrpc.RequestOption) (<-chan *psrpc.Response[*livekit2.CheckEnabledResponse], error) {
	return client.RequestMulti[*livekit2.CheckEnabledResponse](ctx, c.client, "CheckEnabled", nil, req, opts...)
}

func (c *agentInternalClient) CheckAvailability(ctx context.Context, req *livekit2.CheckAvailabilityRequest, opts ...psrpc.RequestOption) (<-chan *psrpc.Response[*livekit2.CheckAvailabilityResponse], error) {
	return client.RequestMulti[*livekit2.CheckAvailabilityResponse](ctx, c.client, "CheckAvailability", nil, req, opts...)
}

func (c *agentInternalClient) JobRequest(ctx context.Context, namespace string, jobType string, req *livekit2.Job, opts ...psrpc.RequestOption) (*JobRequestResponse, error) {
	return client.RequestSingle[*JobRequestResponse](ctx, c.client, "JobRequest", []string{namespace, jobType}, req, opts...)
}

func (c *agentInternalClient) JobTerminate(ctx context.Context, jobId string, req *JobTerminateRequest, opts ...psrpc.RequestOption) (*JobTerminateResponse, error) {
	return client.RequestSingle[*JobTerminateResponse](ctx, c.client, "JobTerminate", []string{jobId}, req, opts...)
}

func (c *agentInternalClient) SubscribeWorkerRegistered(ctx context.Context, handlerNamespace string) (psrpc.Subscription[*google_protobuf.Empty], error) {
	return client.Join[*google_protobuf.Empty](ctx, c.client, "WorkerRegistered", []string{handlerNamespace})
}

func (s *agentInternalClient) Close() {
	s.client.Close()
}

// ====================
// AgentInternal Server
// ====================

type agentInternalServer struct {
	svc AgentInternalServerImpl
	rpc *server.RPCServer
}

// NewAgentInternalServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewAgentInternalServer(svc AgentInternalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (AgentInternalServer, error) {
	sd := &info.ServiceDefinition{
		Name: "AgentInternal",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("CheckEnabled", false, true, false, false)
	var err error
	err = server.RegisterHandler(s, "CheckEnabled", nil, svc.CheckEnabled, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("CheckAvailability", false, true, false, false)
	err = server.RegisterHandler(s, "CheckAvailability", nil, svc.CheckAvailability, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("JobRequest", true, false, true, false)
	sd.RegisterMethod("JobTerminate", false, false, true, true)
	sd.RegisterMethod("WorkerRegistered", false, true, false, false)
	return &agentInternalServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *agentInternalServer) RegisterJobRequestTopic(namespace string, jobType string) error {
	return server.RegisterHandler(s.rpc, "JobRequest", []string{namespace, jobType}, s.svc.JobRequest, s.svc.JobRequestAffinity)
}

func (s *agentInternalServer) DeregisterJobRequestTopic(namespace string, jobType string) {
	s.rpc.DeregisterHandler("JobRequest", []string{namespace, jobType})
}

func (s *agentInternalServer) RegisterJobTerminateTopic(jobId string) error {
	return server.RegisterHandler(s.rpc, "JobTerminate", []string{jobId}, s.svc.JobTerminate, nil)
}

func (s *agentInternalServer) DeregisterJobTerminateTopic(jobId string) {
	s.rpc.DeregisterHandler("JobTerminate", []string{jobId})
}

func (s *agentInternalServer) PublishWorkerRegistered(ctx context.Context, handlerNamespace string, msg *google_protobuf.Empty) error {
	return s.rpc.Publish(ctx, "WorkerRegistered", []string{handlerNamespace}, msg)
}

func (s *agentInternalServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *agentInternalServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor0 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xe1, 0x6e, 0xd3, 0x30,
	0x10, 0xc6, 0x1b, 0x9b, 0x3a, 0xd3, 0xb1, 0xd4, 0xa5, 0xd0, 0x05, 0x26, 0xd6, 0xfc, 0xa1, 0xe2,
	0x47, 0x82, 0xca, 0x6f, 0xc4, 0x0a, 0x0b, 0xa8, 0x15, 0x6b, 0x45, 0x96, 0x81, 0x84, 0x84, 0xa2,
	0x38, 0x39, 0x5a, 0xaf, 0xa9, 0x6d, 0x1c, 0x77, 0x52, 0x1f, 0xa1, 0xaf, 0xd3, 0x47, 0xe2, 0x49,
	0x50, 0x9a, 0xac, 0xa4, 0x63, 0x43, 0xe2, 0xa7, 0xef, 0xbb, 0xfb, 0xee, 0xee, 0xbb, 0xcf, 0xf8,
	0x40, 0xc9, 0xc8, 0x09, 0x47, 0xc0, 0xb5, 0x2d, 0x95, 0xd0, 0x82, 0x6c, 0x2b, 0x19, 0x99, 0x4f,
	0x47, 0x42, 0x8c, 0x12, 0x70, 0x56, 0x21, 0x3a, 0xfb, 0xe1, 0xc0, 0x54, 0xea, 0x79, 0x9e, 0x61,
	0xee, 0x0b, 0xa9, 0x99, 0xe0, 0x69, 0xf1, 0xac, 0x27, 0xec, 0x0a, 0x26, 0x4c, 0x07, 0x25, 0x16,
	0xeb, 0x0d, 0x26, 0x7d, 0x41, 0x3d, 0xf8, 0x39, 0x83, 0x54, 0x7b, 0x90, 0x4a, 0xc1, 0x53, 0x20,
	0x2f, 0xf0, 0x4e, 0xaa, 0x43, 0x0d, 0x4d, 0x74, 0x8c, 0xda, 0x0f, 0x3a, 0x35, 0xbb, 0x28, 0xb5,
	0xfb, 0x82, 0x9e, 0x67, 0x80, 0x97, 0xe3, 0xd6, 0x77, 0x5c, 0xef, 0x0b, 0xea, 0x83, 0x9a, 0x32,
	0x9e, 0x85, 0x73, 0x1e, 0xd2, 0xc0, 0xbb, 0x97, 0x82, 0x06, 0x2c, 0x5e, 0x11, 0xec, 0x79, 0x3b,
	0x97, 0x82, 0xf6, 0x62, 0xe2, 0xe0, 0x5d, 0x05, 0x61, 0x2a, 0x78, 0x73, 0xeb, 0x18, 0xb5, 0x1f,
	0x76, 0x9e, 0xd8, 0x4a, 0x46, 0xf6, 0x26, 0x41, 0x06, 0x7b, 0x45, 0x9a, 0xf5, 0x16, 0x3f, 0xda,
	0x44, 0xff, 0x73, 0xbe, 0x97, 0xa7, 0xab, 0xf5, 0x6e, 0xd0, 0x93, 0x43, 0xdc, 0xf0, 0x5d, 0xef,
	0xac, 0x37, 0xe8, 0xfa, 0xbd, 0xe1, 0x20, 0xf0, 0xdc, 0xcf, 0x17, 0xee, 0xb9, 0xef, 0x9e, 0x1a,
	0xf7, 0x48, 0x1d, 0x1f, 0x74, 0x3f, 0xba, 0x03, 0x3f, 0xf8, 0xe4, 0x7e, 0xf0, 0x03, 0x6f, 0x38,
	0x3c, 0x33, 0x50, 0xe7, 0xd7, 0x36, 0xde, 0xef, 0x66, 0xa2, 0xf5, 0xb8, 0x06, 0xc5, 0xc3, 0x84,
	0x5c, 0xe0, 0xea, 0xfb, 0x31, 0x44, 0x13, 0x97, 0x87, 0x34, 0x81, 0x98, 0x3c, 0x5b, 0x4f, 0x50,
	0x0e, 0x17, 0x72, 0x98, 0x47, 0x77, 0xa0, 0xf9, 0x36, 0x56, 0x65, 0xb9, 0x40, 0xf7, 0x4f, 0xb6,
	0xda, 0x88, 0xc4, 0xb8, 0xb6, 0xca, 0xe8, 0x5e, 0x85, 0x2c, 0x09, 0x29, 0x4b, 0x98, 0x9e, 0x93,
	0xd6, 0x66, 0x75, 0x19, 0xbb, 0x6e, 0x60, 0xfd, 0x2b, 0xe5, 0xaf, 0x2e, 0x5f, 0x30, 0xfe, 0x73,
	0x73, 0x52, 0x2d, 0x8b, 0x67, 0xae, 0x4f, 0x72, 0xc3, 0x12, 0x56, 0x6b, 0xb9, 0x40, 0x47, 0x06,
	0x32, 0x1b, 0x64, 0x8f, 0x87, 0x53, 0x48, 0x65, 0x18, 0x01, 0xa9, 0x64, 0x57, 0xd6, 0x73, 0x09,
	0x27, 0xe8, 0x15, 0xca, 0x44, 0x29, 0x8b, 0x4d, 0x9a, 0xb7, 0x9c, 0x37, 0x9f, 0xf7, 0xf0, 0x16,
	0xa4, 0xe8, 0x63, 0x2c, 0x17, 0xa8, 0x6a, 0x20, 0xb3, 0x42, 0x0a, 0x0b, 0x11, 0xc0, 0xc6, 0x57,
	0xa1, 0x26, 0xa0, 0x3c, 0x18, 0xb1, 0x54, 0x83, 0x82, 0x98, 0x3c, 0xb6, 0x73, 0xe3, 0xdb, 0xd7,
	0xc6, 0xb7, 0xdd, 0xcc, 0xf8, 0xe6, 0x1d, 0xf1, 0x7c, 0xfa, 0x0a, 0x32, 0x90, 0x59, 0x27, 0xb5,
	0x71, 0xc8, 0xe3, 0x04, 0x54, 0xb0, 0xde, 0x23, 0x53, 0xe5, 0x5d, 0xeb, 0xdb, 0xf3, 0x11, 0xd3,
	0xe3, 0x19, 0xb5, 0x23, 0x31, 0x75, 0x0a, 0x4d, 0xf2, 0x8f, 0x15, 0x89, 0xc4, 0x51, 0x32, 0xa2,
	0xbb, 0xab, 0xd7, 0xeb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xb9, 0x37, 0x76, 0x8c, 0x03,
	0x00, 0x00,
}
