// Code generated by protoc-gen-psrpc v0.6.0, DO NOT EDIT.
// source: rpc/io.proto

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
import livekit4 "github.com/livekit/protocol/livekit"
import livekit5 "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_6

// =======================
// IOInfo Client Interface
// =======================

type IOInfoClient interface {
	// egress
	CreateEgress(ctx context.Context, req *livekit4.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	UpdateEgress(ctx context.Context, req *livekit4.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	GetEgress(ctx context.Context, req *GetEgressRequest, opts ...psrpc.RequestOption) (*livekit4.EgressInfo, error)

	ListEgress(ctx context.Context, req *livekit4.ListEgressRequest, opts ...psrpc.RequestOption) (*livekit4.ListEgressResponse, error)

	UpdateMetrics(ctx context.Context, req *UpdateMetricsRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	// ingress
	CreateIngress(ctx context.Context, req *livekit5.IngressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	GetIngressInfo(ctx context.Context, req *GetIngressInfoRequest, opts ...psrpc.RequestOption) (*GetIngressInfoResponse, error)

	UpdateIngressState(ctx context.Context, req *UpdateIngressStateRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	// sip
	GetSIPTrunkAuthentication(ctx context.Context, req *GetSIPTrunkAuthenticationRequest, opts ...psrpc.RequestOption) (*GetSIPTrunkAuthenticationResponse, error)

	EvaluateSIPDispatchRules(ctx context.Context, req *EvaluateSIPDispatchRulesRequest, opts ...psrpc.RequestOption) (*EvaluateSIPDispatchRulesResponse, error)

	UpdateSIPCallState(ctx context.Context, req *UpdateSIPCallStateRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	// Close immediately, without waiting for pending RPCs
	Close()
}

// ===========================
// IOInfo ServerImpl Interface
// ===========================

type IOInfoServerImpl interface {
	// egress
	CreateEgress(context.Context, *livekit4.EgressInfo) (*google_protobuf.Empty, error)

	UpdateEgress(context.Context, *livekit4.EgressInfo) (*google_protobuf.Empty, error)

	GetEgress(context.Context, *GetEgressRequest) (*livekit4.EgressInfo, error)

	ListEgress(context.Context, *livekit4.ListEgressRequest) (*livekit4.ListEgressResponse, error)

	UpdateMetrics(context.Context, *UpdateMetricsRequest) (*google_protobuf.Empty, error)

	// ingress
	CreateIngress(context.Context, *livekit5.IngressInfo) (*google_protobuf.Empty, error)

	GetIngressInfo(context.Context, *GetIngressInfoRequest) (*GetIngressInfoResponse, error)

	UpdateIngressState(context.Context, *UpdateIngressStateRequest) (*google_protobuf.Empty, error)

	// sip
	GetSIPTrunkAuthentication(context.Context, *GetSIPTrunkAuthenticationRequest) (*GetSIPTrunkAuthenticationResponse, error)

	EvaluateSIPDispatchRules(context.Context, *EvaluateSIPDispatchRulesRequest) (*EvaluateSIPDispatchRulesResponse, error)

	UpdateSIPCallState(context.Context, *UpdateSIPCallStateRequest) (*google_protobuf.Empty, error)
}

// =======================
// IOInfo Server Interface
// =======================

type IOInfoServer interface {

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// =============
// IOInfo Client
// =============

type iOInfoClient struct {
	client *client.RPCClient
}

// NewIOInfoClient creates a psrpc client that implements the IOInfoClient interface.
func NewIOInfoClient(bus psrpc.MessageBus, opts ...psrpc.ClientOption) (IOInfoClient, error) {
	sd := &info.ServiceDefinition{
		Name: "IOInfo",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("CreateEgress", false, false, true, true)
	sd.RegisterMethod("UpdateEgress", false, false, true, true)
	sd.RegisterMethod("GetEgress", false, false, true, true)
	sd.RegisterMethod("ListEgress", false, false, true, true)
	sd.RegisterMethod("UpdateMetrics", false, false, true, true)
	sd.RegisterMethod("CreateIngress", false, false, true, true)
	sd.RegisterMethod("GetIngressInfo", false, false, true, true)
	sd.RegisterMethod("UpdateIngressState", false, false, true, true)
	sd.RegisterMethod("GetSIPTrunkAuthentication", false, false, true, true)
	sd.RegisterMethod("EvaluateSIPDispatchRules", false, false, true, true)
	sd.RegisterMethod("UpdateSIPCallState", false, false, true, true)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &iOInfoClient{
		client: rpcClient,
	}, nil
}

func (c *iOInfoClient) CreateEgress(ctx context.Context, req *livekit4.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "CreateEgress", nil, req, opts...)
}

func (c *iOInfoClient) UpdateEgress(ctx context.Context, req *livekit4.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "UpdateEgress", nil, req, opts...)
}

func (c *iOInfoClient) GetEgress(ctx context.Context, req *GetEgressRequest, opts ...psrpc.RequestOption) (*livekit4.EgressInfo, error) {
	return client.RequestSingle[*livekit4.EgressInfo](ctx, c.client, "GetEgress", nil, req, opts...)
}

func (c *iOInfoClient) ListEgress(ctx context.Context, req *livekit4.ListEgressRequest, opts ...psrpc.RequestOption) (*livekit4.ListEgressResponse, error) {
	return client.RequestSingle[*livekit4.ListEgressResponse](ctx, c.client, "ListEgress", nil, req, opts...)
}

func (c *iOInfoClient) UpdateMetrics(ctx context.Context, req *UpdateMetricsRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "UpdateMetrics", nil, req, opts...)
}

func (c *iOInfoClient) CreateIngress(ctx context.Context, req *livekit5.IngressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "CreateIngress", nil, req, opts...)
}

func (c *iOInfoClient) GetIngressInfo(ctx context.Context, req *GetIngressInfoRequest, opts ...psrpc.RequestOption) (*GetIngressInfoResponse, error) {
	return client.RequestSingle[*GetIngressInfoResponse](ctx, c.client, "GetIngressInfo", nil, req, opts...)
}

func (c *iOInfoClient) UpdateIngressState(ctx context.Context, req *UpdateIngressStateRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "UpdateIngressState", nil, req, opts...)
}

func (c *iOInfoClient) GetSIPTrunkAuthentication(ctx context.Context, req *GetSIPTrunkAuthenticationRequest, opts ...psrpc.RequestOption) (*GetSIPTrunkAuthenticationResponse, error) {
	return client.RequestSingle[*GetSIPTrunkAuthenticationResponse](ctx, c.client, "GetSIPTrunkAuthentication", nil, req, opts...)
}

func (c *iOInfoClient) EvaluateSIPDispatchRules(ctx context.Context, req *EvaluateSIPDispatchRulesRequest, opts ...psrpc.RequestOption) (*EvaluateSIPDispatchRulesResponse, error) {
	return client.RequestSingle[*EvaluateSIPDispatchRulesResponse](ctx, c.client, "EvaluateSIPDispatchRules", nil, req, opts...)
}

func (c *iOInfoClient) UpdateSIPCallState(ctx context.Context, req *UpdateSIPCallStateRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "UpdateSIPCallState", nil, req, opts...)
}

func (s *iOInfoClient) Close() {
	s.client.Close()
}

// =============
// IOInfo Server
// =============

type iOInfoServer struct {
	svc IOInfoServerImpl
	rpc *server.RPCServer
}

// NewIOInfoServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewIOInfoServer(svc IOInfoServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (IOInfoServer, error) {
	sd := &info.ServiceDefinition{
		Name: "IOInfo",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("CreateEgress", false, false, true, true)
	var err error
	err = server.RegisterHandler(s, "CreateEgress", nil, svc.CreateEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("UpdateEgress", false, false, true, true)
	err = server.RegisterHandler(s, "UpdateEgress", nil, svc.UpdateEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("GetEgress", false, false, true, true)
	err = server.RegisterHandler(s, "GetEgress", nil, svc.GetEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("ListEgress", false, false, true, true)
	err = server.RegisterHandler(s, "ListEgress", nil, svc.ListEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("UpdateMetrics", false, false, true, true)
	err = server.RegisterHandler(s, "UpdateMetrics", nil, svc.UpdateMetrics, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("CreateIngress", false, false, true, true)
	err = server.RegisterHandler(s, "CreateIngress", nil, svc.CreateIngress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("GetIngressInfo", false, false, true, true)
	err = server.RegisterHandler(s, "GetIngressInfo", nil, svc.GetIngressInfo, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("UpdateIngressState", false, false, true, true)
	err = server.RegisterHandler(s, "UpdateIngressState", nil, svc.UpdateIngressState, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("GetSIPTrunkAuthentication", false, false, true, true)
	err = server.RegisterHandler(s, "GetSIPTrunkAuthentication", nil, svc.GetSIPTrunkAuthentication, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("EvaluateSIPDispatchRules", false, false, true, true)
	err = server.RegisterHandler(s, "EvaluateSIPDispatchRules", nil, svc.EvaluateSIPDispatchRules, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("UpdateSIPCallState", false, false, true, true)
	err = server.RegisterHandler(s, "UpdateSIPCallState", nil, svc.UpdateSIPCallState, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	return &iOInfoServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *iOInfoServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *iOInfoServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor4 = []byte{
	// 1458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x57, 0xdd, 0x56, 0xdb, 0xc6,
	0x16, 0x3e, 0xb6, 0xc1, 0xc1, 0xdb, 0x60, 0xcc, 0x60, 0x38, 0xc6, 0xac, 0x93, 0x10, 0xe7, 0xa4,
	0xa5, 0xed, 0x5a, 0xa6, 0xa1, 0x17, 0x6d, 0xd3, 0xd5, 0xac, 0x00, 0x71, 0x88, 0x5a, 0x02, 0xae,
	0x80, 0x8b, 0xf6, 0x46, 0x1d, 0xa4, 0xc1, 0x4c, 0x90, 0x35, 0xea, 0xcc, 0x08, 0xc2, 0x13, 0xb4,
	0x8f, 0xd0, 0xcb, 0x3e, 0x42, 0x5f, 0xa1, 0x2f, 0xd5, 0xeb, 0xae, 0xf9, 0x91, 0x2d, 0x63, 0x3b,
	0xe0, 0xde, 0x69, 0xf6, 0xdf, 0x7c, 0xfb, 0x57, 0x7b, 0x60, 0x9e, 0xc7, 0xfe, 0x16, 0x65, 0xad,
	0x98, 0x33, 0xc9, 0x50, 0x81, 0xc7, 0x7e, 0xa3, 0x16, 0xd2, 0x2b, 0x72, 0x49, 0xa5, 0x47, 0xba,
	0x9c, 0x08, 0x61, 0x58, 0x8d, 0x95, 0x94, 0x4a, 0xa3, 0x2c, 0x79, 0x29, 0x25, 0x0b, 0x1a, 0x5b,
	0xd2, 0x7a, 0x97, 0xb1, 0x6e, 0x48, 0xb6, 0xf4, 0xe9, 0x2c, 0x39, 0xdf, 0x22, 0xbd, 0x58, 0xde,
	0x58, 0xe6, 0xc3, 0xdb, 0xcc, 0x20, 0xe1, 0x58, 0x52, 0x16, 0x19, 0x7e, 0x73, 0x0b, 0xaa, 0xfb,
	0x44, 0xb6, 0xf5, 0x15, 0x2e, 0xf9, 0x25, 0x21, 0x42, 0xa2, 0x75, 0x28, 0x19, 0x28, 0x1e, 0x0d,
	0xea, 0xb9, 0x8d, 0xdc, 0x66, 0xc9, 0x9d, 0x33, 0x04, 0x27, 0x68, 0xfe, 0x9a, 0x83, 0xda, 0x69,
	0x1c, 0x60, 0x49, 0xde, 0x12, 0xc9, 0xa9, 0xdf, 0xd7, 0xfa, 0x18, 0x66, 0x68, 0x74, 0xce, 0xb4,
	0x42, 0x79, 0x7b, 0xb9, 0x65, 0x81, 0xb6, 0x8c, 0x6d, 0x27, 0x3a, 0x67, 0xae, 0x16, 0x40, 0x4d,
	0x58, 0xc0, 0x57, 0x5d, 0xcf, 0x8f, 0x13, 0x2f, 0x11, 0xb8, 0x4b, 0xea, 0x85, 0x8d, 0xdc, 0x66,
	0xde, 0x2d, 0xe3, 0xab, 0xee, 0x5e, 0x9c, 0x9c, 0x2a, 0x92, 0x92, 0xe9, 0xe1, 0xf7, 0x19, 0x99,
	0x19, 0x23, 0xd3, 0xc3, 0xef, 0x53, 0x99, 0xe6, 0x29, 0xac, 0xec, 0x13, 0xe9, 0x44, 0x03, 0xfb,
	0x16, 0xc9, 0xff, 0x00, 0x6c, 0xd0, 0x06, 0x0e, 0x94, 0x2c, 0xc5, 0x09, 0x14, 0x5b, 0x48, 0x4e,
	0x70, 0xcf, 0xbb, 0x24, 0x37, 0xf5, 0xbc, 0x61, 0x1b, 0xca, 0xf7, 0xe4, 0xa6, 0xf9, 0x5b, 0x1e,
	0x56, 0x6f, 0xdb, 0x15, 0x31, 0x8b, 0x04, 0x41, 0x9b, 0x43, 0x2e, 0xd6, 0xfa, 0x2e, 0x66, 0x65,
	0x8d, 0x8f, 0x35, 0x98, 0x95, 0xec, 0x92, 0x44, 0xd6, 0xbc, 0x39, 0xa0, 0x15, 0x28, 0x5e, 0x0b,
	0x2f, 0xe1, 0xa1, 0x76, 0xb9, 0xe4, 0xce, 0x5e, 0x8b, 0x53, 0x1e, 0xa2, 0x53, 0xa8, 0x84, 0xac,
	0xdb, 0xa5, 0x51, 0xd7, 0x3b, 0xa7, 0x24, 0x0c, 0x44, 0x7d, 0x66, 0xa3, 0xb0, 0x59, 0xde, 0x6e,
	0xb5, 0x78, 0xec, 0xb7, 0xc6, 0x63, 0x69, 0x1d, 0x18, 0x8d, 0xd7, 0x5a, 0xa1, 0x1d, 0x49, 0x7e,
	0xe3, 0x2e, 0x84, 0x59, 0x5a, 0xe3, 0x25, 0xa0, 0x51, 0x21, 0x54, 0x85, 0x82, 0x72, 0xdb, 0x44,
	0x45, 0x7d, 0x2a, 0xac, 0x57, 0x38, 0x4c, 0x48, 0x8a, 0x55, 0x1f, 0x9e, 0xe7, 0xbf, 0xca, 0x35,
	0xbb, 0xb0, 0x66, 0x52, 0x6d, 0x01, 0x1c, 0x4b, 0x2c, 0xc9, 0x3d, 0xa3, 0xfc, 0x19, 0xcc, 0x0a,
	0x25, 0xae, 0xad, 0x96, 0xb7, 0x57, 0x6e, 0x07, 0xcb, 0xd8, 0x32, 0x32, 0xcd, 0x3f, 0x72, 0xb0,
	0xb1, 0x4f, 0xe4, 0xb1, 0xd3, 0x39, 0xe1, 0x49, 0x74, 0xb9, 0x93, 0xc8, 0x0b, 0x12, 0x49, 0xea,
	0xeb, 0x4a, 0x4d, 0x2f, 0x7c, 0x08, 0x65, 0x41, 0x63, 0xcf, 0xc7, 0x61, 0xa8, 0x6e, 0x2c, 0xda,
	0xc4, 0xd1, 0x78, 0x0f, 0x87, 0xa1, 0x13, 0x20, 0x04, 0x33, 0xe7, 0x9c, 0xf5, 0xac, 0x1b, 0xfa,
	0x1b, 0x55, 0x20, 0x2f, 0x99, 0x8d, 0x76, 0x5e, 0x32, 0xf4, 0x08, 0xca, 0x82, 0xfb, 0x1e, 0x0e,
	0x02, 0x85, 0x41, 0x57, 0x55, 0xc9, 0x05, 0xc1, 0xfd, 0x1d, 0x43, 0x41, 0xff, 0x85, 0x07, 0x92,
	0x79, 0x17, 0x4c, 0xc8, 0xfa, 0xac, 0x66, 0x16, 0x25, 0x7b, 0xc3, 0x84, 0x6c, 0xfe, 0x99, 0x83,
	0xc7, 0x1f, 0x80, 0x68, 0x2b, 0xa4, 0x01, 0x73, 0x89, 0x20, 0x3c, 0xc2, 0x3d, 0x92, 0x76, 0x4e,
	0x7a, 0x56, 0xbc, 0x18, 0x0b, 0x71, 0xcd, 0x78, 0x60, 0x31, 0xf6, 0xcf, 0x0a, 0x7b, 0xc0, 0x59,
	0xac, 0x91, 0xce, 0xb9, 0xfa, 0x1b, 0x6d, 0xc0, 0xbc, 0xf2, 0x57, 0xaa, 0xeb, 0x94, 0xc3, 0x29,
	0x58, 0x1a, 0x6b, 0x04, 0xa6, 0x92, 0x63, 0xce, 0xde, 0x11, 0x5f, 0x2a, 0xbe, 0xc1, 0x5b, 0xb2,
	0x14, 0x27, 0x68, 0xfe, 0x5d, 0x80, 0x47, 0x6d, 0x95, 0x4d, 0x2c, 0xc9, 0xb1, 0xd3, 0x79, 0x45,
	0x45, 0x8c, 0xa5, 0x7f, 0xe1, 0x26, 0x21, 0x11, 0x13, 0x82, 0x3a, 0x77, 0x3b, 0xa8, 0x9f, 0x03,
	0x52, 0xfc, 0x18, 0x73, 0x49, 0x7d, 0x1a, 0xe3, 0x48, 0xf6, 0xb3, 0xbd, 0x9b, 0xaf, 0xe7, 0xdc,
	0xaa, 0xa0, 0x71, 0x67, 0xc0, 0x74, 0x82, 0x11, 0xd8, 0x30, 0x02, 0xfb, 0x29, 0x54, 0xd4, 0x7d,
	0xaa, 0xde, 0xa3, 0xa4, 0x77, 0x46, 0xb8, 0x0d, 0xc7, 0x82, 0xa5, 0x1e, 0x6a, 0x22, 0x7a, 0x02,
	0x9a, 0x40, 0x82, 0x54, 0xca, 0xa4, 0x71, 0xde, 0x10, 0xad, 0xd0, 0x9d, 0x09, 0xad, 0x42, 0x21,
	0xa6, 0x91, 0x0d, 0x8e, 0xfa, 0x54, 0x5d, 0x18, 0x31, 0x4f, 0x11, 0x8b, 0x3a, 0xda, 0xb3, 0x11,
	0xeb, 0xd0, 0x48, 0x59, 0xb2, 0xd7, 0xe9, 0xec, 0x3f, 0x30, 0x96, 0x0c, 0x49, 0x55, 0x00, 0x0a,
	0xa0, 0x4a, 0xde, 0x4b, 0x8e, 0x3d, 0x2c, 0x25, 0xa7, 0x67, 0x89, 0x24, 0xa2, 0x5e, 0xd2, 0x8d,
	0xfa, 0xb5, 0x6e, 0xd4, 0x3b, 0x42, 0xdd, 0x6a, 0x2b, 0xe5, 0x9d, 0xbe, 0xae, 0xe9, 0xd9, 0x45,
	0x32, 0x4c, 0x6d, 0xec, 0x42, 0x6d, 0x9c, 0xe0, 0x54, 0x7d, 0xfb, 0x7b, 0x19, 0x36, 0x26, 0xa3,
	0xb1, 0xa5, 0xba, 0x0e, 0x25, 0xce, 0x58, 0xcf, 0xcb, 0xd6, 0xaa, 0x22, 0x1c, 0xaa, 0x5a, 0x7d,
	0x06, 0xb5, 0xe1, 0x94, 0xab, 0x5a, 0x97, 0xe9, 0xb4, 0x5c, 0x8e, 0xb3, 0x19, 0x37, 0x2c, 0xf4,
	0x09, 0x54, 0xb3, 0x2a, 0xda, 0xac, 0x09, 0xe2, 0x62, 0x86, 0x3e, 0xce, 0x7a, 0x8f, 0x48, 0x1c,
	0x60, 0x89, 0x6d, 0xf5, 0x65, 0xad, 0xbf, 0xb5, 0x2c, 0x74, 0x0d, 0xab, 0x59, 0x95, 0x4c, 0x0a,
	0xca, 0x3a, 0x05, 0x2f, 0xef, 0x48, 0x81, 0x9d, 0x9a, 0x99, 0x52, 0xbd, 0x9d, 0x89, 0x95, 0x78,
	0x1c, 0x0f, 0x3d, 0x81, 0x32, 0x37, 0x09, 0xd4, 0x25, 0xa3, 0x1b, 0x54, 0x57, 0x3e, 0x58, 0xb2,
	0xaa, 0x9d, 0xfe, 0xb8, 0x9f, 0x19, 0x3f, 0xee, 0x67, 0xb3, 0xe3, 0xbe, 0x05, 0x45, 0x4e, 0x44,
	0x12, 0x4a, 0x5d, 0x7f, 0x95, 0xed, 0x55, 0x0d, 0x3d, 0x0b, 0x59, 0x73, 0x5d, 0x2b, 0x35, 0xd2,
	0x50, 0xa5, 0x91, 0x86, 0xda, 0x82, 0x9a, 0x92, 0x08, 0xac, 0xbe, 0xc7, 0x93, 0x90, 0x0c, 0x5a,
	0x6f, 0x49, 0xd0, 0x38, 0x1b, 0x8d, 0x91, 0xc1, 0x31, 0x7f, 0x6b, 0x70, 0xa0, 0x03, 0x78, 0x70,
	0x41, 0x70, 0x40, 0xb8, 0xa8, 0x2f, 0xe8, 0xe8, 0x6e, 0xdf, 0x2f, 0xba, 0x6f, 0x8c, 0x92, 0x89,
	0x67, 0x6a, 0x02, 0x71, 0x58, 0xb1, 0x9f, 0x9e, 0x64, 0xd9, 0xcc, 0x55, 0xb4, 0xed, 0x17, 0x53,
	0xd9, 0x3e, 0x61, 0xb7, 0xf3, 0xb6, 0x7c, 0x31, 0xca, 0x51, 0x77, 0x0e, 0x2e, 0x52, 0xd7, 0xa6,
	0xfe, 0xa0, 0x69, 0xee, 0x1c, 0x18, 0x3c, 0x61, 0x43, 0xbe, 0x2d, 0xe3, 0x51, 0x0e, 0xda, 0x85,
	0x45, 0x1a, 0xf9, 0x61, 0x12, 0x90, 0xfe, 0x6d, 0xcb, 0x3a, 0xc1, 0x6b, 0xfd, 0x7f, 0xdf, 0xb1,
	0xd3, 0x31, 0xd2, 0x47, 0xb1, 0xfa, 0x6d, 0x08, 0xb7, 0x62, 0x35, 0x52, 0x1b, 0x2f, 0xa0, 0x4a,
	0x22, 0x7c, 0xa6, 0xa6, 0xd0, 0x39, 0xc1, 0x32, 0xe1, 0x44, 0xd4, 0x17, 0x37, 0x0a, 0x9b, 0x95,
	0xcc, 0x42, 0x75, 0xec, 0x74, 0x5e, 0x1b, 0x9e, 0xbb, 0x68, 0x85, 0xed, 0x59, 0x63, 0xe0, 0x34,
	0xd2, 0xab, 0x84, 0xa4, 0x3d, 0xc2, 0x12, 0x59, 0xaf, 0xea, 0xff, 0xef, 0x5a, 0xcb, 0x2c, 0x82,
	0xad, 0x74, 0x11, 0x6c, 0xbd, 0xb2, 0x8b, 0xa0, 0x5b, 0xb1, 0x1a, 0x27, 0x46, 0x01, 0xb5, 0x61,
	0x49, 0xef, 0x5e, 0xea, 0x97, 0x90, 0x6e, 0x8b, 0xf5, 0xa5, 0xbb, 0xac, 0x2c, 0xaa, 0xd5, 0x0c,
	0x87, 0x61, 0x4a, 0x68, 0xbc, 0x81, 0xc6, 0xe4, 0x6e, 0x9b, 0x66, 0x9c, 0x35, 0x9e, 0xc3, 0x7c,
	0x36, 0xfa, 0x53, 0xe9, 0xbe, 0x86, 0xfa, 0xa4, 0xca, 0x99, 0xd6, 0xce, 0xa4, 0x6a, 0x98, 0x6a,
	0x34, 0x1f, 0xa6, 0x2b, 0xd5, 0xb1, 0xd3, 0x51, 0xe1, 0x1a, 0x5a, 0xa9, 0x9e, 0x41, 0xc9, 0xfc,
	0x88, 0xc7, 0x2d, 0x99, 0x56, 0x41, 0x2f, 0x81, 0x73, 0xbe, 0xfd, 0xfa, 0xf4, 0x67, 0x58, 0x1a,
	0x99, 0x1c, 0xa8, 0x0e, 0xb5, 0x83, 0xf6, 0xfe, 0xce, 0xde, 0x8f, 0xde, 0xce, 0xde, 0x5e, 0xbb,
	0x73, 0xe2, 0x1d, 0xb9, 0x5e, 0xc7, 0x39, 0xac, 0xfe, 0x07, 0x01, 0x14, 0x0d, 0xa9, 0x9a, 0x43,
	0x8b, 0x50, 0x76, 0xdb, 0x3f, 0x9c, 0xb6, 0x8f, 0x4f, 0x34, 0x33, 0xaf, 0x98, 0x6e, 0xfb, 0xbb,
	0xf6, 0xde, 0x49, 0xb5, 0x80, 0xe6, 0x60, 0xe6, 0x95, 0x7b, 0xd4, 0xa9, 0xce, 0x6c, 0xff, 0x55,
	0x84, 0xa2, 0x73, 0xa4, 0x2e, 0x43, 0xdf, 0xc0, 0xfc, 0x1e, 0x27, 0x58, 0x12, 0xb3, 0xd3, 0xa3,
	0x71, 0x4b, 0x7e, 0x63, 0x75, 0xa4, 0x46, 0xda, 0xea, 0x3d, 0xa2, 0x94, 0x8d, 0xe7, 0xff, 0x46,
	0xf9, 0x4b, 0x28, 0xf5, 0x9f, 0x29, 0x68, 0x25, 0xdd, 0x8b, 0x87, 0x9e, 0x2d, 0x8d, 0x71, 0x06,
	0x51, 0x1b, 0xe0, 0x80, 0x8a, 0x54, 0xb3, 0xd1, 0x17, 0x19, 0x10, 0x53, 0xf5, 0xf5, 0xb1, 0x3c,
	0xfb, 0xb3, 0xdc, 0x85, 0x85, 0xa1, 0x47, 0x0f, 0x5a, 0xd3, 0x18, 0xc6, 0x3d, 0x84, 0x26, 0xfa,
	0xf0, 0x2d, 0x2c, 0x98, 0xe8, 0xd9, 0x0d, 0x18, 0x8d, 0x7d, 0x40, 0x4c, 0x54, 0x77, 0xa0, 0x32,
	0xfc, 0x14, 0x40, 0x8d, 0xb1, 0xef, 0x83, 0xd4, 0x9b, 0xc9, 0x6f, 0x07, 0x74, 0x00, 0x68, 0x74,
	0xaf, 0x47, 0x0f, 0x33, 0x2e, 0x8d, 0x59, 0xf8, 0x27, 0x02, 0x7b, 0x07, 0x6b, 0x13, 0x17, 0x63,
	0xf4, 0x34, 0xc5, 0xf1, 0xc1, 0xdd, 0xbe, 0xf1, 0xd1, 0x5d, 0x62, 0x16, 0x79, 0x17, 0xea, 0x93,
	0xa6, 0x36, 0xfa, 0xff, 0x7d, 0xb6, 0xb0, 0xc6, 0xd3, 0x7b, 0x8d, 0xfe, 0x41, 0x88, 0xb2, 0x7d,
	0x3a, 0x14, 0xa2, 0x31, 0x0d, 0x3c, 0x29, 0x44, 0xbb, 0x8f, 0x7f, 0x7a, 0xd4, 0xa5, 0xf2, 0x22,
	0x39, 0x6b, 0xf9, 0xac, 0xb7, 0x65, 0xb3, 0x6e, 0xde, 0xe4, 0x3e, 0x0b, 0xb7, 0x78, 0xec, 0x9f,
	0x15, 0xf5, 0xe9, 0x8b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x43, 0x29, 0xa8, 0x35, 0x21, 0x10,
	0x00, 0x00,
}
