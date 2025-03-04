// Code generated by protoc-gen-twirp v8.1.3, DO NOT EDIT.
// source: livekit_agent.proto

package livekit

import context "context"
import fmt "fmt"
import http "net/http"
import io "io"
import json "encoding/json"
import strconv "strconv"
import strings "strings"

import protojson "google.golang.org/protobuf/encoding/protojson"
import proto "google.golang.org/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// Version compatibility assertion.
// If the constant is not defined in the package, that likely means
// the package needs to be updated to work with this generated code.
// See https://twitchtv.github.io/twirp/docs/version_matrix.html
const _ = twirp.TwirpPackageMinVersion_8_1_0

// ======================
// AgentService Interface
// ======================

type AgentService interface {
	CheckEnabled(context.Context, *CheckEnabledRequest) (*CheckEnabledResponse, error)

	CheckAvailability(context.Context, *CheckAvailabilityRequest) (*CheckAvailabilityResponse, error)
}

// ============================
// AgentService Protobuf Client
// ============================

type agentServiceProtobufClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewAgentServiceProtobufClient creates a Protobuf client that implements the AgentService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewAgentServiceProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) AgentService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwards compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "livekit", "AgentService")
	urls := [2]string{
		serviceURL + "CheckEnabled",
		serviceURL + "CheckAvailability",
	}

	return &agentServiceProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *agentServiceProtobufClient) CheckEnabled(ctx context.Context, in *CheckEnabledRequest) (*CheckEnabledResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "livekit")
	ctx = ctxsetters.WithServiceName(ctx, "AgentService")
	ctx = ctxsetters.WithMethodName(ctx, "CheckEnabled")
	caller := c.callCheckEnabled
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CheckEnabledRequest) (*CheckEnabledResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CheckEnabledRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CheckEnabledRequest) when calling interceptor")
					}
					return c.callCheckEnabled(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CheckEnabledResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CheckEnabledResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *agentServiceProtobufClient) callCheckEnabled(ctx context.Context, in *CheckEnabledRequest) (*CheckEnabledResponse, error) {
	out := new(CheckEnabledResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *agentServiceProtobufClient) CheckAvailability(ctx context.Context, in *CheckAvailabilityRequest) (*CheckAvailabilityResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "livekit")
	ctx = ctxsetters.WithServiceName(ctx, "AgentService")
	ctx = ctxsetters.WithMethodName(ctx, "CheckAvailability")
	caller := c.callCheckAvailability
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CheckAvailabilityRequest) (*CheckAvailabilityResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CheckAvailabilityRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CheckAvailabilityRequest) when calling interceptor")
					}
					return c.callCheckAvailability(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CheckAvailabilityResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CheckAvailabilityResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *agentServiceProtobufClient) callCheckAvailability(ctx context.Context, in *CheckAvailabilityRequest) (*CheckAvailabilityResponse, error) {
	out := new(CheckAvailabilityResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ========================
// AgentService JSON Client
// ========================

type agentServiceJSONClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewAgentServiceJSONClient creates a JSON client that implements the AgentService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewAgentServiceJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) AgentService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwards compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "livekit", "AgentService")
	urls := [2]string{
		serviceURL + "CheckEnabled",
		serviceURL + "CheckAvailability",
	}

	return &agentServiceJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *agentServiceJSONClient) CheckEnabled(ctx context.Context, in *CheckEnabledRequest) (*CheckEnabledResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "livekit")
	ctx = ctxsetters.WithServiceName(ctx, "AgentService")
	ctx = ctxsetters.WithMethodName(ctx, "CheckEnabled")
	caller := c.callCheckEnabled
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CheckEnabledRequest) (*CheckEnabledResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CheckEnabledRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CheckEnabledRequest) when calling interceptor")
					}
					return c.callCheckEnabled(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CheckEnabledResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CheckEnabledResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *agentServiceJSONClient) callCheckEnabled(ctx context.Context, in *CheckEnabledRequest) (*CheckEnabledResponse, error) {
	out := new(CheckEnabledResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *agentServiceJSONClient) CheckAvailability(ctx context.Context, in *CheckAvailabilityRequest) (*CheckAvailabilityResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "livekit")
	ctx = ctxsetters.WithServiceName(ctx, "AgentService")
	ctx = ctxsetters.WithMethodName(ctx, "CheckAvailability")
	caller := c.callCheckAvailability
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CheckAvailabilityRequest) (*CheckAvailabilityResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CheckAvailabilityRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CheckAvailabilityRequest) when calling interceptor")
					}
					return c.callCheckAvailability(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CheckAvailabilityResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CheckAvailabilityResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *agentServiceJSONClient) callCheckAvailability(ctx context.Context, in *CheckAvailabilityRequest) (*CheckAvailabilityResponse, error) {
	out := new(CheckAvailabilityResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ===========================
// AgentService Server Handler
// ===========================

type agentServiceServer struct {
	AgentService
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
	jsonCamelCase    bool   // JSON fields are serialized as lowerCamelCase rather than keeping the original proto names
}

// NewAgentServiceServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewAgentServiceServer(svc AgentService, opts ...interface{}) TwirpServer {
	serverOpts := newServerOpts(opts)

	// Using ReadOpt allows backwards and forwards compatibility with new options in the future
	jsonSkipDefaults := false
	_ = serverOpts.ReadOpt("jsonSkipDefaults", &jsonSkipDefaults)
	jsonCamelCase := false
	_ = serverOpts.ReadOpt("jsonCamelCase", &jsonCamelCase)
	var pathPrefix string
	if ok := serverOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	return &agentServiceServer{
		AgentService:     svc,
		hooks:            serverOpts.Hooks,
		interceptor:      twirp.ChainInterceptors(serverOpts.Interceptors...),
		pathPrefix:       pathPrefix,
		jsonSkipDefaults: jsonSkipDefaults,
		jsonCamelCase:    jsonCamelCase,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *agentServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *agentServiceServer) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
	if context.Canceled == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.Canceled, "failed to read request: context canceled"))
		return
	}
	if context.DeadlineExceeded == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.DeadlineExceeded, "failed to read request: deadline exceeded"))
		return
	}
	s.writeError(ctx, resp, twirp.WrapError(malformedRequestError(msg), err))
}

// AgentServicePathPrefix is a convenience constant that may identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// with the default "/twirp" prefix and default CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const AgentServicePathPrefix = "/twirp/livekit.AgentService/"

func (s *agentServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "livekit")
	ctx = ctxsetters.WithServiceName(ctx, "AgentService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	// Verify path format: [<prefix>]/<package>.<Service>/<Method>
	prefix, pkgService, method := parseTwirpPath(req.URL.Path)
	if pkgService != "livekit.AgentService" {
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
	if prefix != s.pathPrefix {
		msg := fmt.Sprintf("invalid path prefix %q, expected %q, on path %q", prefix, s.pathPrefix, req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	switch method {
	case "CheckEnabled":
		s.serveCheckEnabled(ctx, resp, req)
		return
	case "CheckAvailability":
		s.serveCheckAvailability(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *agentServiceServer) serveCheckEnabled(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveCheckEnabledJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveCheckEnabledProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *agentServiceServer) serveCheckEnabledJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CheckEnabled")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(CheckEnabledRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.AgentService.CheckEnabled
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CheckEnabledRequest) (*CheckEnabledResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CheckEnabledRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CheckEnabledRequest) when calling interceptor")
					}
					return s.AgentService.CheckEnabled(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CheckEnabledResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CheckEnabledResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CheckEnabledResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CheckEnabledResponse and nil error while calling CheckEnabled. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *agentServiceServer) serveCheckEnabledProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CheckEnabled")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := io.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(CheckEnabledRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.AgentService.CheckEnabled
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CheckEnabledRequest) (*CheckEnabledResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CheckEnabledRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CheckEnabledRequest) when calling interceptor")
					}
					return s.AgentService.CheckEnabled(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CheckEnabledResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CheckEnabledResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CheckEnabledResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CheckEnabledResponse and nil error while calling CheckEnabled. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *agentServiceServer) serveCheckAvailability(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveCheckAvailabilityJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveCheckAvailabilityProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *agentServiceServer) serveCheckAvailabilityJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CheckAvailability")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(CheckAvailabilityRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.AgentService.CheckAvailability
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CheckAvailabilityRequest) (*CheckAvailabilityResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CheckAvailabilityRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CheckAvailabilityRequest) when calling interceptor")
					}
					return s.AgentService.CheckAvailability(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CheckAvailabilityResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CheckAvailabilityResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CheckAvailabilityResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CheckAvailabilityResponse and nil error while calling CheckAvailability. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *agentServiceServer) serveCheckAvailabilityProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CheckAvailability")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := io.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(CheckAvailabilityRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.AgentService.CheckAvailability
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CheckAvailabilityRequest) (*CheckAvailabilityResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CheckAvailabilityRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CheckAvailabilityRequest) when calling interceptor")
					}
					return s.AgentService.CheckAvailability(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CheckAvailabilityResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CheckAvailabilityResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CheckAvailabilityResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CheckAvailabilityResponse and nil error while calling CheckAvailability. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *agentServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor5, 0
}

func (s *agentServiceServer) ProtocGenTwirpVersion() string {
	return "v8.1.3"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *agentServiceServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "livekit", "AgentService")
}

var twirpFileDescriptor5 = []byte{
	// 1634 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x58, 0xcd, 0x72, 0xdb, 0xc8,
	0x11, 0x26, 0x00, 0x52, 0x24, 0x9b, 0x3f, 0xa6, 0x46, 0x92, 0x97, 0xab, 0x5d, 0xef, 0x4a, 0xd8,
	0xa4, 0xac, 0x55, 0x62, 0xb9, 0xa2, 0xe4, 0xe0, 0x68, 0x77, 0x2b, 0x21, 0xb5, 0x72, 0x48, 0x5a,
	0x92, 0xe5, 0xa1, 0x14, 0x57, 0xf9, 0x82, 0x02, 0x89, 0x31, 0x05, 0x09, 0xc0, 0x20, 0x98, 0xa1,
	0x5c, 0xba, 0xe5, 0x90, 0x43, 0x5e, 0x21, 0xd7, 0x1c, 0xf3, 0x02, 0xa9, 0x9c, 0x92, 0x63, 0x9e,
	0x20, 0x95, 0x43, 0x9e, 0x20, 0x4f, 0x91, 0x9a, 0xc1, 0xe0, 0x87, 0x7f, 0x8e, 0x6a, 0x6f, 0x9c,
	0xfe, 0xc3, 0x4c, 0x7f, 0xdd, 0x5f, 0xb7, 0x04, 0x1b, 0x9e, 0x7b, 0x47, 0x6e, 0x5d, 0x6e, 0xd9,
	0x13, 0x12, 0xf0, 0x83, 0x30, 0xa2, 0x9c, 0xa2, 0xb2, 0x12, 0x6e, 0x6f, 0x26, 0x5a, 0x9f, 0x3a,
	0xc4, 0x63, 0xb1, 0xda, 0xfc, 0x97, 0x0e, 0xc6, 0x80, 0x8e, 0x50, 0x13, 0x74, 0xd7, 0x69, 0x6b,
	0x3b, 0xda, 0x5e, 0x15, 0xeb, 0xae, 0x83, 0xbe, 0x84, 0x9a, 0xe3, 0xb2, 0xd0, 0xe6, 0xe3, 0x6b,
	0xcb, 0x75, 0xda, 0x55, 0xa9, 0x80, 0x44, 0xd4, 0x77, 0xd0, 0x8f, 0xa0, 0xc8, 0xef, 0x43, 0xd2,
	0xd6, 0x77, 0xb4, 0xbd, 0xe6, 0x61, 0xeb, 0x40, 0x45, 0x3f, 0x18, 0xd0, 0xd1, 0xe5, 0x7d, 0x48,
	0xb0, 0xd4, 0xa2, 0x5d, 0x28, 0x46, 0x94, 0xfa, 0x6d, 0x63, 0x47, 0xdb, 0xab, 0x1d, 0x36, 0x52,
	0x2b, 0x4c, 0xa9, 0x8f, 0xa5, 0x0a, 0xfd, 0x0a, 0x6a, 0xa1, 0x1d, 0x71, 0x77, 0xec, 0x86, 0x76,
	0xc0, 0xdb, 0x45, 0x69, 0xd9, 0x4e, 0x2d, 0x2f, 0x32, 0x5d, 0x3f, 0x78, 0x4f, 0x7b, 0x05, 0x9c,
	0x37, 0xff, 0xa3, 0xa6, 0xa1, 0x1d, 0xa8, 0x06, 0xb6, 0x4f, 0x58, 0x68, 0x8f, 0x49, 0xbb, 0x24,
	0x2e, 0xda, 0xd5, 0xdb, 0x1a, 0xce, 0x84, 0x68, 0x1b, 0x2a, 0x3e, 0xe1, 0xb6, 0x63, 0x73, 0xbb,
	0xbd, 0x26, 0x5f, 0x92, 0x9e, 0xd1, 0x13, 0x00, 0x99, 0x2e, 0x4b, 0x98, 0xb7, 0xcb, 0x52, 0x5b,
	0x95, 0x92, 0x73, 0xdb, 0x27, 0xe8, 0x29, 0x94, 0x18, 0xb7, 0x39, 0x69, 0x57, 0xe4, 0xbd, 0xd6,
	0xf3, 0xef, 0x1c, 0x0a, 0x05, 0x8e, 0xf5, 0xdd, 0x26, 0xd4, 0xad, 0xdc, 0xc5, 0xcc, 0x7f, 0x6b,
	0x50, 0x49, 0x6c, 0xd0, 0x3e, 0xac, 0x09, 0xab, 0x29, 0x93, 0x19, 0x6e, 0x1e, 0xa2, 0xf9, 0x30,
	0x53, 0x86, 0x95, 0x05, 0xda, 0x84, 0x12, 0x89, 0x22, 0x1a, 0xc9, 0xcc, 0x56, 0x71, 0x7c, 0x10,
	0xd7, 0x64, 0xdc, 0x8e, 0x38, 0x71, 0x2c, 0x9b, 0xcb, 0x74, 0x1a, 0xb8, 0xaa, 0x24, 0x1d, 0x8e,
	0x3e, 0x85, 0x0a, 0x09, 0x9c, 0x58, 0x59, 0x94, 0xca, 0xb2, 0x3c, 0x77, 0xb8, 0xf0, 0x9c, 0x86,
	0x8e, 0xad, 0x3c, 0x4b, 0xb1, 0xa7, 0x92, 0x74, 0x38, 0xfa, 0x19, 0x6c, 0xe6, 0xae, 0x6d, 0xb9,
	0x0e, 0x09, 0xb8, 0xcb, 0xef, 0x55, 0x9e, 0x36, 0x72, 0xba, 0xbe, 0x52, 0x99, 0x7f, 0xd6, 0x01,
	0xde, 0xd2, 0xe8, 0x96, 0x44, 0x02, 0x91, 0x85, 0xd2, 0x99, 0xcd, 0xa8, 0x3e, 0x9f, 0xd1, 0x36,
	0x94, 0xef, 0x48, 0xc4, 0x5c, 0x1a, 0xc8, 0x67, 0x54, 0x71, 0x72, 0x44, 0xcf, 0xd2, 0x2c, 0x15,
	0x65, 0x96, 0xb6, 0xd2, 0x2c, 0xc5, 0x5f, 0x9b, 0x4b, 0x54, 0x1e, 0xd5, 0xd2, 0x1c, 0xaa, 0x08,
	0x8a, 0x1e, 0xb5, 0x1d, 0xf9, 0x0a, 0x1d, 0xcb, 0xdf, 0xe8, 0x33, 0xa8, 0xde, 0xd0, 0x91, 0x35,
	0xa6, 0xd3, 0x80, 0x4b, 0xa0, 0x1b, 0xb8, 0x72, 0x43, 0x47, 0xc7, 0xe2, 0x8c, 0x5e, 0xc3, 0x86,
	0xed, 0x79, 0xf4, 0x03, 0x71, 0xac, 0x90, 0x44, 0xbe, 0xcb, 0xc4, 0x8d, 0x98, 0x42, 0xfd, 0x8b,
	0x65, 0xd5, 0x78, 0x91, 0x9a, 0x61, 0xa4, 0x5c, 0x33, 0x11, 0x33, 0xff, 0x69, 0x40, 0x23, 0xbe,
	0xf6, 0x19, 0x61, 0xcc, 0x9e, 0x10, 0xf4, 0x2d, 0x54, 0x22, 0x32, 0x71, 0x19, 0x27, 0x91, 0xcc,
	0x56, 0x3e, 0x2e, 0x56, 0x8a, 0xd8, 0x03, 0x93, 0xdf, 0x4d, 0x09, 0xe3, 0xbd, 0x02, 0x4e, 0x3d,
	0xd0, 0x31, 0xd4, 0xed, 0x3b, 0xdb, 0xf5, 0xec, 0x91, 0xeb, 0x09, 0x7c, 0x74, 0x19, 0xe1, 0x49,
	0x1a, 0xa1, 0x93, 0x53, 0x62, 0xc2, 0x42, 0x1a, 0x30, 0xd2, 0x2b, 0xe0, 0x19, 0x27, 0xd4, 0x85,
	0x46, 0x8c, 0xbc, 0xf5, 0x41, 0x7e, 0x48, 0xf5, 0xe5, 0x67, 0x69, 0x94, 0x2b, 0xa9, 0xcd, 0xa7,
	0x5b, 0xc4, 0x98, 0xe6, 0xa4, 0xe8, 0x97, 0x49, 0x3d, 0x59, 0x37, 0x74, 0xb4, 0xd0, 0xae, 0x71,
	0x80, 0xb4, 0xaa, 0x7b, 0x85, 0xa4, 0xd6, 0x04, 0xc9, 0x7c, 0x0d, 0xc5, 0xd0, 0x0d, 0x26, 0x12,
	0xad, 0xda, 0xe1, 0xc6, 0x1c, 0xbc, 0x17, 0x6e, 0x30, 0xe9, 0x15, 0xb0, 0x34, 0x41, 0xbf, 0x86,
	0x3a, 0x73, 0xfd, 0xa9, 0x97, 0x7c, 0x67, 0x6d, 0xee, 0xa2, 0x43, 0xa5, 0x1c, 0xd0, 0x51, 0x96,
	0xad, 0x1a, 0xcb, 0xa4, 0xe8, 0x3b, 0xa8, 0xf9, 0xee, 0x24, 0x4a, 0x02, 0x94, 0x65, 0x80, 0xed,
	0x34, 0xc0, 0x59, 0xac, 0x9b, 0xf1, 0x07, 0x3f, 0x15, 0x76, 0xab, 0x50, 0xf6, 0x63, 0xe0, 0xcc,
	0xbf, 0xeb, 0xd0, 0x18, 0x92, 0xe8, 0x2e, 0x83, 0xf2, 0xbb, 0x05, 0x28, 0xbf, 0x5c, 0x09, 0x65,
	0x0a, 0x45, 0x86, 0x65, 0x77, 0x29, 0x96, 0x9f, 0xaf, 0xc0, 0x32, 0xb9, 0xdd, 0x2c, 0x94, 0x2f,
	0x00, 0x6c, 0xc6, 0xdc, 0x49, 0xe0, 0x93, 0x80, 0x2b, 0x1c, 0x1f, 0xe7, 0x69, 0xa5, 0x93, 0x6a,
	0xc5, 0xcb, 0x32, 0x5b, 0xf4, 0x0d, 0xd4, 0xb8, 0x28, 0xd4, 0xc0, 0xe6, 0xa2, 0x09, 0x63, 0x30,
	0x3e, 0x99, 0x21, 0xf0, 0x4c, 0x2d, 0xb2, 0x9a, 0xb3, 0x96, 0x10, 0xd2, 0x60, 0xa2, 0x70, 0x5f,
	0x80, 0x90, 0x2a, 0x08, 0x69, 0x30, 0xc9, 0x67, 0xf0, 0x4f, 0x1a, 0xa0, 0x45, 0xc4, 0xd2, 0x19,
	0xa2, 0x3d, 0x68, 0x86, 0xe8, 0xab, 0x67, 0xc8, 0xd1, 0xec, 0x0c, 0x31, 0x3e, 0x3e, 0x43, 0x66,
	0x26, 0x88, 0xb9, 0x9f, 0x90, 0x99, 0xa8, 0x3f, 0xf4, 0x39, 0x54, 0xb9, 0xeb, 0x13, 0xc6, 0x6d,
	0x3f, 0x94, 0xf7, 0x32, 0x70, 0x26, 0x30, 0xdf, 0xa4, 0xb6, 0x34, 0x98, 0xa0, 0x1f, 0x43, 0xd3,
	0xb3, 0x19, 0xb7, 0xe6, 0x1d, 0x1a, 0x42, 0x7a, 0x99, 0x08, 0x67, 0x43, 0xea, 0xf3, 0x21, 0xff,
	0xa6, 0xc3, 0xd6, 0xd2, 0xee, 0x7f, 0x60, 0x76, 0x66, 0xd9, 0xb6, 0xf2, 0x70, 0xb6, 0xcd, 0xd3,
	0x67, 0x71, 0x8e, 0x3e, 0xbf, 0x82, 0x86, 0xe8, 0x42, 0xcb, 0x0d, 0x38, 0x89, 0xee, 0x6c, 0x4f,
	0x16, 0x49, 0x03, 0xd7, 0x85, 0xb0, 0xaf, 0x64, 0x68, 0x37, 0x3f, 0x77, 0xe5, 0xb8, 0x10, 0xdd,
	0x9e, 0x8a, 0xc4, 0x68, 0x5e, 0xc1, 0xaa, 0xe5, 0x1f, 0xca, 0xaa, 0xdd, 0x3a, 0x80, 0x95, 0x7e,
	0xc1, 0xbc, 0x85, 0xc7, 0xcb, 0xbb, 0x4d, 0x70, 0x7d, 0xcc, 0x70, 0x56, 0x3a, 0x9a, 0x2a, 0xb1,
	0xa0, 0xef, 0xa0, 0x5f, 0x40, 0x8d, 0xc9, 0x76, 0xb6, 0xdc, 0xe0, 0x3d, 0x55, 0xd5, 0x92, 0x95,
	0x72, 0xdc, 0xea, 0xb2, 0x50, 0x80, 0xa5, 0xbf, 0xcd, 0x9f, 0xc2, 0xfa, 0x02, 0x67, 0xa0, 0x4f,
	0xa0, 0x2c, 0x66, 0x8a, 0xeb, 0xb0, 0xb6, 0xbe, 0x63, 0xec, 0x55, 0xf1, 0xda, 0x0d, 0x1d, 0xf5,
	0x1d, 0x66, 0xbe, 0x81, 0x8d, 0x25, 0x5d, 0x8c, 0xbe, 0x00, 0x43, 0x90, 0x51, 0xcc, 0x19, 0xf5,
	0x3c, 0xa4, 0x58, 0x28, 0x04, 0x28, 0x11, 0x61, 0x53, 0x5f, 0xb0, 0xa4, 0x28, 0x95, 0x0a, 0x4e,
	0xcf, 0xe6, 0x3f, 0x0c, 0xd8, 0x5c, 0xc6, 0xf2, 0x68, 0x0b, 0xd6, 0xe2, 0x4b, 0xa8, 0x97, 0x96,
	0xe4, 0x1d, 0x44, 0xdd, 0x29, 0xc6, 0xf0, 0x88, 0x0a, 0x96, 0x09, 0xd0, 0x53, 0x78, 0xc4, 0xa6,
	0x61, 0x48, 0x23, 0xce, 0x2c, 0xf9, 0x09, 0x22, 0x13, 0x51, 0xc1, 0xcd, 0x44, 0x8c, 0xa5, 0x14,
	0x7d, 0x0d, 0xad, 0xfc, 0x82, 0x20, 0xcb, 0x2c, 0xae, 0x97, 0x47, 0x39, 0xb9, 0x2c, 0xb6, 0x55,
	0xbb, 0x44, 0x69, 0xe5, 0x2e, 0x31, 0xef, 0x32, 0xb7, 0xa6, 0xe5, 0x5d, 0xce, 0x92, 0xe2, 0xa4,
	0xf0, 0x38, 0xef, 0x62, 0x73, 0x1e, 0xb9, 0xa3, 0x29, 0x27, 0xa2, 0xae, 0x8c, 0xbd, 0xda, 0xe1,
	0x8b, 0x8f, 0xce, 0xc4, 0x7c, 0xb1, 0x75, 0x52, 0xd7, 0x93, 0x80, 0x47, 0xf7, 0x78, 0x2b, 0x5c,
	0xa6, 0xdb, 0xee, 0xc1, 0xf6, 0x6a, 0x27, 0xd4, 0x02, 0xe3, 0x96, 0xdc, 0xab, 0xd4, 0x8b, 0x9f,
	0x62, 0x83, 0xbb, 0xb3, 0xbd, 0x69, 0xb2, 0xfb, 0xc4, 0x87, 0x23, 0xfd, 0x85, 0x66, 0xde, 0xc0,
	0xa3, 0xb9, 0x01, 0xb9, 0x0a, 0xbc, 0x6c, 0x63, 0xd4, 0x1f, 0xbe, 0x31, 0x1a, 0xb9, 0x8d, 0xd1,
	0xfc, 0x83, 0x06, 0x68, 0x71, 0x9c, 0xa3, 0xc3, 0xb9, 0x55, 0x74, 0xf9, 0x92, 0xd5, 0x2b, 0x24,
	0xd1, 0x45, 0x1b, 0x27, 0xdb, 0x94, 0xb1, 0x6a, 0x9b, 0x2a, 0xce, 0x6e, 0x53, 0x82, 0xfa, 0xad,
	0xd8, 0xdd, 0x74, 0xa0, 0x31, 0x33, 0x8c, 0xfe, 0x6f, 0x0b, 0x6c, 0x81, 0x31, 0x8d, 0xbc, 0x38,
	0x77, 0xbd, 0x02, 0x16, 0x07, 0x71, 0x87, 0x4d, 0x28, 0x71, 0x7a, 0x4b, 0x12, 0x1a, 0x8b, 0x0f,
	0xdd, 0x35, 0x28, 0x5a, 0xd3, 0xc8, 0x33, 0x9f, 0x42, 0x73, 0x76, 0x6e, 0xad, 0xc8, 0xab, 0xb9,
	0x05, 0x1b, 0xc7, 0xd7, 0x64, 0x7c, 0x7b, 0x12, 0x88, 0x2e, 0x70, 0x54, 0x5f, 0x9a, 0xff, 0xd1,
	0x60, 0x73, 0x56, 0xae, 0x7a, 0x6b, 0x17, 0xea, 0x62, 0xc2, 0x58, 0x24, 0x96, 0xcb, 0x60, 0x15,
	0x5c, 0x13, 0x32, 0x65, 0x8a, 0x7e, 0x02, 0xeb, 0xe1, 0x74, 0xe4, 0xb9, 0xec, 0x9a, 0x44, 0xa9,
	0x5d, 0xdc, 0x6f, 0xad, 0x54, 0x91, 0x18, 0x3f, 0x87, 0x7c, 0x4d, 0xa7, 0xe6, 0x25, 0x69, 0x8e,
	0x72, 0xaa, 0xc4, 0xc1, 0x04, 0x48, 0x09, 0x8f, 0xb5, 0x0d, 0x41, 0x32, 0xf2, 0xcf, 0x9b, 0x9c,
	0x54, 0xfc, 0xb1, 0x96, 0xcd, 0x00, 0xb1, 0x3d, 0x0b, 0x26, 0x82, 0x74, 0x08, 0x30, 0xf3, 0x1b,
	0x68, 0xcb, 0xd7, 0x2d, 0xa3, 0xa4, 0x39, 0x67, 0x6d, 0xc1, 0xf9, 0xf7, 0x1a, 0x7c, 0xba, 0xc4,
	0x5b, 0x25, 0x68, 0x8e, 0x4c, 0xb5, 0x07, 0x91, 0x29, 0x7a, 0x06, 0xe5, 0x98, 0x8e, 0x63, 0xde,
	0x5c, 0xdc, 0x24, 0xa4, 0x47, 0x62, 0xb3, 0xff, 0x2d, 0x94, 0xd5, 0xd4, 0x43, 0x35, 0x28, 0x0f,
	0x2e, 0x2d, 0xfc, 0xfa, 0xf5, 0x59, 0xab, 0x80, 0x5a, 0x50, 0x1f, 0x5c, 0x5a, 0x17, 0x57, 0xdd,
	0xd3, 0xfe, 0xb0, 0x77, 0x82, 0x5b, 0xa2, 0x54, 0x9b, 0x42, 0xd2, 0xc1, 0x97, 0xfd, 0xe3, 0xfe,
	0x45, 0xe7, 0xfc, 0xb2, 0xa5, 0xef, 0x3f, 0x83, 0xfa, 0x4c, 0x0b, 0xb4, 0xa0, 0xfe, 0x76, 0x68,
	0x75, 0x7e, 0xdb, 0xe9, 0x9f, 0x76, 0xba, 0xa7, 0x27, 0xad, 0x82, 0x08, 0xfa, 0x76, 0x68, 0xbd,
	0xbc, 0x3a, 0x3d, 0x6d, 0x69, 0xfb, 0x03, 0xa8, 0x66, 0xed, 0xd9, 0x04, 0x18, 0x0c, 0xad, 0x8b,
	0x93, 0xf3, 0xef, 0xfb, 0xe7, 0xbf, 0x69, 0x15, 0xd4, 0x19, 0x5f, 0x9d, 0x9f, 0x8b, 0xb3, 0xa6,
	0xce, 0xc3, 0xab, 0xe3, 0xe3, 0x93, 0xe1, 0xb0, 0xa5, 0xa3, 0x06, 0x54, 0x07, 0x43, 0xeb, 0x65,
	0xa7, 0x7f, 0x7a, 0xf2, 0x7d, 0xcb, 0x38, 0xfc, 0xab, 0x06, 0xf5, 0x8e, 0x48, 0xa5, 0xc8, 0x83,
	0x3b, 0x26, 0xe8, 0x15, 0xd4, 0xf3, 0x75, 0x86, 0xb2, 0xa5, 0x6f, 0x49, 0x59, 0x6e, 0x3f, 0x59,
	0xa1, 0x55, 0xb9, 0x7f, 0x07, 0xeb, 0x0b, 0xc0, 0xa0, 0xdd, 0x59, 0x9f, 0x25, 0x90, 0x6f, 0x9b,
	0x1f, 0x33, 0x89, 0x63, 0x77, 0x5f, 0xbe, 0xfb, 0x6a, 0xe2, 0xf2, 0xeb, 0xe9, 0xe8, 0x60, 0x4c,
	0xfd, 0xe7, 0xca, 0xfe, 0xb9, 0xfc, 0xa7, 0xc1, 0x98, 0x7a, 0x89, 0xe0, 0x2f, 0x7a, 0xe3, 0xd4,
	0xbd, 0x23, 0xaf, 0xc4, 0x10, 0x17, 0xaa, 0xff, 0xea, 0x4d, 0x75, 0x3e, 0x3a, 0x92, 0x82, 0xd1,
	0x9a, 0x74, 0xf9, 0xf9, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x4b, 0xc0, 0xb7, 0x9d, 0x10,
	0x00, 0x00,
}
