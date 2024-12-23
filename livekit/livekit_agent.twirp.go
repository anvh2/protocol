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

	CheckAvailibility(context.Context, *CheckAvailibilityRequest) (*CheckAvailibilityResponse, error)
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
		serviceURL + "CheckAvailibility",
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

func (c *agentServiceProtobufClient) CheckAvailibility(ctx context.Context, in *CheckAvailibilityRequest) (*CheckAvailibilityResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "livekit")
	ctx = ctxsetters.WithServiceName(ctx, "AgentService")
	ctx = ctxsetters.WithMethodName(ctx, "CheckAvailibility")
	caller := c.callCheckAvailibility
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CheckAvailibilityRequest) (*CheckAvailibilityResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CheckAvailibilityRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CheckAvailibilityRequest) when calling interceptor")
					}
					return c.callCheckAvailibility(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CheckAvailibilityResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CheckAvailibilityResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *agentServiceProtobufClient) callCheckAvailibility(ctx context.Context, in *CheckAvailibilityRequest) (*CheckAvailibilityResponse, error) {
	out := new(CheckAvailibilityResponse)
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
		serviceURL + "CheckAvailibility",
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

func (c *agentServiceJSONClient) CheckAvailibility(ctx context.Context, in *CheckAvailibilityRequest) (*CheckAvailibilityResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "livekit")
	ctx = ctxsetters.WithServiceName(ctx, "AgentService")
	ctx = ctxsetters.WithMethodName(ctx, "CheckAvailibility")
	caller := c.callCheckAvailibility
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CheckAvailibilityRequest) (*CheckAvailibilityResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CheckAvailibilityRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CheckAvailibilityRequest) when calling interceptor")
					}
					return c.callCheckAvailibility(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CheckAvailibilityResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CheckAvailibilityResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *agentServiceJSONClient) callCheckAvailibility(ctx context.Context, in *CheckAvailibilityRequest) (*CheckAvailibilityResponse, error) {
	out := new(CheckAvailibilityResponse)
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
	case "CheckAvailibility":
		s.serveCheckAvailibility(ctx, resp, req)
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

func (s *agentServiceServer) serveCheckAvailibility(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveCheckAvailibilityJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveCheckAvailibilityProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *agentServiceServer) serveCheckAvailibilityJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CheckAvailibility")
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
	reqContent := new(CheckAvailibilityRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.AgentService.CheckAvailibility
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CheckAvailibilityRequest) (*CheckAvailibilityResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CheckAvailibilityRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CheckAvailibilityRequest) when calling interceptor")
					}
					return s.AgentService.CheckAvailibility(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CheckAvailibilityResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CheckAvailibilityResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CheckAvailibilityResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CheckAvailibilityResponse and nil error while calling CheckAvailibility. nil responses are not supported"))
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

func (s *agentServiceServer) serveCheckAvailibilityProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CheckAvailibility")
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
	reqContent := new(CheckAvailibilityRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.AgentService.CheckAvailibility
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CheckAvailibilityRequest) (*CheckAvailibilityResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CheckAvailibilityRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CheckAvailibilityRequest) when calling interceptor")
					}
					return s.AgentService.CheckAvailibility(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CheckAvailibilityResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CheckAvailibilityResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CheckAvailibilityResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CheckAvailibilityResponse and nil error while calling CheckAvailibility. nil responses are not supported"))
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
	// 1594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x57, 0xcd, 0x72, 0xe3, 0xc6,
	0x11, 0x26, 0x00, 0x52, 0x24, 0x9b, 0x3f, 0xcb, 0x1d, 0x49, 0x6b, 0x5a, 0xf6, 0xda, 0x12, 0x9c,
	0xd4, 0xca, 0x4a, 0x56, 0xae, 0x28, 0x39, 0x6c, 0xe4, 0xb8, 0x12, 0x52, 0xd6, 0x86, 0xa4, 0x25,
	0xad, 0x3c, 0x94, 0xb2, 0x55, 0xbe, 0xa0, 0x40, 0x62, 0x4c, 0x41, 0x02, 0x30, 0x08, 0x66, 0xa8,
	0x2d, 0xdd, 0x73, 0xc8, 0x2b, 0xe4, 0x90, 0x4b, 0x8e, 0xc9, 0x03, 0xe4, 0x98, 0x1c, 0xf3, 0x04,
	0xa9, 0x3c, 0x43, 0x9e, 0x22, 0x35, 0x83, 0xc1, 0x1f, 0x7f, 0x36, 0x2a, 0xdf, 0x38, 0xfd, 0x87,
	0x99, 0xfe, 0xba, 0xbf, 0x6e, 0xc2, 0xa6, 0xe7, 0xde, 0x93, 0x3b, 0x97, 0x5b, 0xf6, 0x8c, 0x04,
	0xfc, 0x30, 0x8c, 0x28, 0xa7, 0xa8, 0xaa, 0x84, 0x3b, 0x5b, 0x89, 0xd6, 0xa7, 0x0e, 0xf1, 0x58,
	0xac, 0x36, 0xff, 0xad, 0x83, 0x31, 0xa2, 0x13, 0xd4, 0x06, 0xdd, 0x75, 0xba, 0xda, 0xae, 0xb6,
	0x5f, 0xc7, 0xba, 0xeb, 0xa0, 0x4f, 0xa1, 0xe1, 0xb8, 0x2c, 0xb4, 0xf9, 0xf4, 0xc6, 0x72, 0x9d,
	0x6e, 0x5d, 0x2a, 0x20, 0x11, 0x0d, 0x1d, 0xf4, 0x23, 0x28, 0xf3, 0x87, 0x90, 0x74, 0xf5, 0x5d,
	0x6d, 0xbf, 0x7d, 0xd4, 0x39, 0x54, 0xd1, 0x0f, 0x47, 0x74, 0x72, 0xf5, 0x10, 0x12, 0x2c, 0xb5,
	0x68, 0x0f, 0xca, 0x11, 0xa5, 0x7e, 0xd7, 0xd8, 0xd5, 0xf6, 0x1b, 0x47, 0xad, 0xd4, 0x0a, 0x53,
	0xea, 0x63, 0xa9, 0x42, 0xbf, 0x86, 0x46, 0x68, 0x47, 0xdc, 0x9d, 0xba, 0xa1, 0x1d, 0xf0, 0x6e,
	0x59, 0x5a, 0x76, 0x53, 0xcb, 0xcb, 0x4c, 0x37, 0x0c, 0xbe, 0xa7, 0x83, 0x12, 0xce, 0x9b, 0xff,
	0x51, 0xd3, 0xd0, 0x2e, 0xd4, 0x03, 0xdb, 0x27, 0x2c, 0xb4, 0xa7, 0xa4, 0x5b, 0x11, 0x17, 0xed,
	0xeb, 0x5d, 0x0d, 0x67, 0x42, 0xb4, 0x03, 0x35, 0x9f, 0x70, 0xdb, 0xb1, 0xb9, 0xdd, 0xdd, 0x90,
	0x2f, 0x49, 0xcf, 0xe8, 0x39, 0x80, 0x4c, 0x97, 0x25, 0xcc, 0xbb, 0x55, 0xa9, 0xad, 0x4b, 0xc9,
	0x85, 0xed, 0x13, 0xf4, 0x02, 0x2a, 0x8c, 0xdb, 0x9c, 0x74, 0x6b, 0xf2, 0x5e, 0x4f, 0xf3, 0xef,
	0x1c, 0x0b, 0x05, 0x8e, 0xf5, 0xfd, 0x36, 0x34, 0xad, 0xdc, 0xc5, 0xcc, 0xff, 0x68, 0x50, 0x4b,
	0x6c, 0xd0, 0x01, 0x6c, 0x08, 0xab, 0x39, 0x93, 0x19, 0x6e, 0x1f, 0xa1, 0xc5, 0x30, 0x73, 0x86,
	0x95, 0x05, 0xda, 0x82, 0x0a, 0x89, 0x22, 0x1a, 0xc9, 0xcc, 0xd6, 0x71, 0x7c, 0x10, 0xd7, 0x64,
	0xdc, 0x8e, 0x38, 0x71, 0x2c, 0x9b, 0xcb, 0x74, 0x1a, 0xb8, 0xae, 0x24, 0x3d, 0x8e, 0x3e, 0x84,
	0x1a, 0x09, 0x9c, 0x58, 0x59, 0x96, 0xca, 0xaa, 0x3c, 0xf7, 0xb8, 0xf0, 0x9c, 0x87, 0x8e, 0xad,
	0x3c, 0x2b, 0xb1, 0xa7, 0x92, 0xf4, 0x38, 0xfa, 0x19, 0x6c, 0xe5, 0xae, 0x6d, 0xb9, 0x0e, 0x09,
	0xb8, 0xcb, 0x1f, 0x54, 0x9e, 0x36, 0x73, 0xba, 0xa1, 0x52, 0x99, 0x7f, 0xd1, 0x01, 0xde, 0xd2,
	0xe8, 0x8e, 0x44, 0x02, 0x91, 0xa5, 0xd2, 0x29, 0x66, 0x54, 0x5f, 0xcc, 0x68, 0x17, 0xaa, 0xf7,
	0x24, 0x62, 0x2e, 0x0d, 0xe4, 0x33, 0xea, 0x38, 0x39, 0xa2, 0x97, 0x69, 0x96, 0xca, 0x32, 0x4b,
	0xdb, 0x69, 0x96, 0xe2, 0xaf, 0x2d, 0x24, 0x2a, 0x8f, 0x6a, 0x65, 0x01, 0x55, 0x04, 0x65, 0x8f,
	0xda, 0x8e, 0x7c, 0x85, 0x8e, 0xe5, 0x6f, 0xf4, 0x11, 0xd4, 0x6f, 0xe9, 0xc4, 0x9a, 0xd2, 0x79,
	0xc0, 0x25, 0xd0, 0x2d, 0x5c, 0xbb, 0xa5, 0x93, 0x13, 0x71, 0x46, 0x6f, 0x60, 0xd3, 0xf6, 0x3c,
	0xfa, 0x8e, 0x38, 0x56, 0x48, 0x22, 0xdf, 0x65, 0xe2, 0x46, 0x4c, 0xa1, 0xfe, 0xc9, 0xaa, 0x6a,
	0xbc, 0x4c, 0xcd, 0x30, 0x52, 0xae, 0x99, 0x88, 0x99, 0xff, 0x32, 0xa0, 0x15, 0x5f, 0xfb, 0x9c,
	0x30, 0x66, 0xcf, 0x08, 0xfa, 0x15, 0xd4, 0x22, 0x32, 0x73, 0x19, 0x27, 0x91, 0xcc, 0x56, 0x3e,
	0x2e, 0x56, 0x8a, 0xd8, 0x03, 0x93, 0xdf, 0xcf, 0x09, 0xe3, 0x83, 0x12, 0x4e, 0x3d, 0xd0, 0x09,
	0x34, 0xed, 0x7b, 0xdb, 0xf5, 0xec, 0x89, 0xeb, 0x09, 0x7c, 0x74, 0x19, 0xe1, 0x79, 0x1a, 0xa1,
	0x97, 0x53, 0x62, 0xc2, 0x42, 0x1a, 0x30, 0x32, 0x28, 0xe1, 0x82, 0x13, 0xea, 0x43, 0x2b, 0x46,
	0xde, 0x7a, 0x27, 0x3f, 0xa4, 0xfa, 0xf2, 0xa3, 0x34, 0xca, 0xb5, 0xd4, 0xe6, 0xd3, 0x2d, 0x62,
	0xcc, 0x73, 0x52, 0xf4, 0xcb, 0xa4, 0x9e, 0xac, 0x5b, 0x3a, 0x59, 0x6a, 0xd7, 0x38, 0x40, 0x5a,
	0xd5, 0x83, 0x52, 0x52, 0x6b, 0x82, 0x64, 0x3e, 0x87, 0x72, 0xe8, 0x06, 0x33, 0x89, 0x56, 0xe3,
	0x68, 0x73, 0x01, 0xde, 0x4b, 0x37, 0x98, 0x0d, 0x4a, 0x58, 0x9a, 0xa0, 0xdf, 0x40, 0x93, 0xb9,
	0xfe, 0xdc, 0x4b, 0xbe, 0xb3, 0xb1, 0x70, 0xd1, 0xb1, 0x52, 0x8e, 0xe8, 0x24, 0xcb, 0x56, 0x83,
	0x65, 0x52, 0xf4, 0x15, 0x34, 0x7c, 0x77, 0x16, 0x25, 0x01, 0xaa, 0x32, 0xc0, 0x4e, 0x1a, 0xe0,
	0x3c, 0xd6, 0x15, 0xfc, 0xc1, 0x4f, 0x85, 0xfd, 0x3a, 0x54, 0xfd, 0x18, 0x38, 0xf3, 0x1f, 0x3a,
	0xb4, 0xc6, 0x24, 0xba, 0xcf, 0xa0, 0xfc, 0x6a, 0x09, 0xca, 0x4f, 0xd7, 0x42, 0x99, 0x42, 0x91,
	0x61, 0xd9, 0x5f, 0x89, 0xe5, 0xc7, 0x6b, 0xb0, 0x4c, 0x6e, 0x57, 0x84, 0xf2, 0x15, 0x80, 0xcd,
	0x98, 0x3b, 0x0b, 0x7c, 0x12, 0x70, 0x85, 0xe3, 0xb3, 0x3c, 0xad, 0xf4, 0x52, 0xad, 0x78, 0x59,
	0x66, 0x8b, 0xbe, 0x84, 0x06, 0x17, 0x85, 0x1a, 0xd8, 0x5c, 0x34, 0x61, 0x0c, 0xc6, 0x07, 0x05,
	0x02, 0xcf, 0xd4, 0x22, 0xab, 0x39, 0x6b, 0x09, 0x21, 0x0d, 0x66, 0x0a, 0xf7, 0x25, 0x08, 0xa9,
	0x82, 0x90, 0x06, 0xb3, 0x7c, 0x06, 0xff, 0xa4, 0x01, 0x5a, 0x46, 0x2c, 0x9d, 0x21, 0xda, 0xa3,
	0x66, 0x88, 0xbe, 0x7e, 0x86, 0x1c, 0x17, 0x67, 0x88, 0xf1, 0xfe, 0x19, 0x52, 0x98, 0x20, 0xe6,
	0x41, 0x42, 0x66, 0xa2, 0xfe, 0xd0, 0xc7, 0x50, 0xe7, 0xae, 0x4f, 0x18, 0xb7, 0xfd, 0x50, 0xde,
	0xcb, 0xc0, 0x99, 0xc0, 0xfc, 0x36, 0xb5, 0xa5, 0xc1, 0x0c, 0xfd, 0x18, 0xda, 0x9e, 0xcd, 0xb8,
	0xb5, 0xe8, 0xd0, 0x12, 0xd2, 0xab, 0x44, 0x58, 0x0c, 0xa9, 0x2f, 0x86, 0xfc, 0xb3, 0x0e, 0xdb,
	0x2b, 0xbb, 0xff, 0x91, 0xd9, 0x29, 0xb2, 0x6d, 0xed, 0xf1, 0x6c, 0xfb, 0x19, 0xb4, 0x44, 0xa7,
	0x59, 0x6e, 0xc0, 0x49, 0x74, 0x6f, 0x7b, 0xb2, 0x10, 0x5a, 0xb8, 0x29, 0x84, 0x43, 0x25, 0x43,
	0x7b, 0xf9, 0xd9, 0x2a, 0x47, 0x82, 0xe8, 0xe8, 0x54, 0x24, 0xc6, 0xef, 0x1a, 0xe6, 0xac, 0xfe,
	0x50, 0xe6, 0xec, 0x37, 0x01, 0xac, 0xf4, 0x0b, 0xe6, 0x1d, 0x3c, 0x5b, 0xdd, 0x51, 0x82, 0xcf,
	0x63, 0x16, 0xb3, 0xd2, 0xf1, 0x53, 0x8b, 0x05, 0x43, 0x07, 0xfd, 0x02, 0x1a, 0x4c, 0xb6, 0xac,
	0xe5, 0x06, 0xdf, 0x53, 0x55, 0x11, 0x59, 0xb9, 0xc6, 0xed, 0x2c, 0x8b, 0x01, 0x58, 0xfa, 0xdb,
	0xfc, 0x29, 0x3c, 0x5d, 0xe2, 0x05, 0xf4, 0x01, 0x54, 0xc5, 0xdc, 0x70, 0x1d, 0xd6, 0xd5, 0x77,
	0x8d, 0xfd, 0x3a, 0xde, 0xb8, 0xa5, 0x93, 0xa1, 0xc3, 0xcc, 0x6f, 0x61, 0x73, 0x45, 0xa7, 0xa2,
	0x4f, 0xc0, 0x10, 0x84, 0x13, 0xf3, 0x42, 0x33, 0x0f, 0x1b, 0x16, 0x0a, 0x31, 0xb7, 0x22, 0xc2,
	0xe6, 0xbe, 0x60, 0x42, 0x51, 0x0e, 0x35, 0x9c, 0x9e, 0xcd, 0x7f, 0x1a, 0xb0, 0xb5, 0x8a, 0xc9,
	0xd1, 0x36, 0x6c, 0xc4, 0x97, 0x50, 0x2f, 0xad, 0xc8, 0x3b, 0x88, 0xda, 0x52, 0xac, 0xe0, 0x11,
	0x15, 0x2c, 0x13, 0xa0, 0x17, 0xf0, 0x84, 0xcd, 0xc3, 0x90, 0x46, 0x9c, 0x59, 0xf2, 0x13, 0x44,
	0x26, 0xa2, 0x86, 0xdb, 0x89, 0x18, 0x4b, 0x29, 0xfa, 0x1c, 0x3a, 0xf9, 0x25, 0x40, 0x96, 0x52,
	0x59, 0x7e, 0xe7, 0x49, 0x4e, 0x2e, 0x0b, 0x6a, 0xdd, 0xbe, 0x50, 0x59, 0xbb, 0x2f, 0x2c, 0xba,
	0x2c, 0xac, 0x62, 0x79, 0x97, 0xf3, 0x64, 0x7e, 0x53, 0x78, 0x96, 0x77, 0xb1, 0x39, 0x8f, 0xdc,
	0xc9, 0x9c, 0x13, 0x51, 0x57, 0xc6, 0x7e, 0xe3, 0xe8, 0xd5, 0x7b, 0xe7, 0x5e, 0xbe, 0xd8, 0x7a,
	0xa9, 0xeb, 0x69, 0xc0, 0xa3, 0x07, 0xbc, 0x1d, 0xae, 0xd2, 0xed, 0x0c, 0x60, 0x67, 0xbd, 0x13,
	0xea, 0x80, 0x71, 0x47, 0x1e, 0x54, 0xea, 0xc5, 0x4f, 0xb1, 0xa5, 0xdd, 0xdb, 0xde, 0x3c, 0xd9,
	0x6f, 0xe2, 0xc3, 0xb1, 0xfe, 0x4a, 0x33, 0x6f, 0xe1, 0xc9, 0xc2, 0x10, 0x5c, 0x07, 0x5e, 0xb6,
	0x15, 0xea, 0x8f, 0xdf, 0x0a, 0x8d, 0xdc, 0x56, 0x68, 0xfe, 0x41, 0x03, 0xb4, 0x3c, 0xb2, 0xd1,
	0xd1, 0xc2, 0xba, 0xb9, 0x7a, 0x91, 0x1a, 0x94, 0x92, 0xe8, 0xa2, 0x8d, 0x93, 0x8d, 0xc9, 0x58,
	0xb7, 0x31, 0x95, 0x8b, 0x1b, 0x93, 0xa0, 0x77, 0x2b, 0x76, 0x37, 0x1d, 0x68, 0x15, 0x06, 0xce,
	0xff, 0x6d, 0x81, 0x6d, 0x30, 0xe6, 0x91, 0x17, 0xe7, 0x6e, 0x50, 0xc2, 0xe2, 0x20, 0xee, 0xb0,
	0x05, 0x15, 0x4e, 0xef, 0x48, 0x42, 0x55, 0xf1, 0xa1, 0xbf, 0x01, 0x65, 0x6b, 0x1e, 0x79, 0xe6,
	0x0b, 0x68, 0x17, 0x67, 0xd3, 0x9a, 0xbc, 0x9a, 0xdb, 0xb0, 0x79, 0x72, 0x43, 0xa6, 0x77, 0xa7,
	0x81, 0xe8, 0x02, 0x47, 0xf5, 0xa5, 0xf9, 0x37, 0x0d, 0xb6, 0x8a, 0x72, 0xd5, 0x5b, 0x7b, 0xd0,
	0x14, 0x53, 0xc4, 0x22, 0xb1, 0x5c, 0x06, 0xab, 0xe1, 0x86, 0x90, 0x29, 0x53, 0xf4, 0x13, 0x78,
	0x1a, 0xce, 0x27, 0x9e, 0xcb, 0x6e, 0x48, 0x94, 0xda, 0xc5, 0xfd, 0xd6, 0x49, 0x15, 0x89, 0xb1,
	0x09, 0x90, 0xf2, 0x17, 0xeb, 0x1a, 0x82, 0x33, 0xe4, 0x3f, 0x92, 0x9c, 0x54, 0xfc, 0xbf, 0xca,
	0x68, 0x5b, 0x2c, 0xbc, 0x82, 0x58, 0x20, 0xe5, 0x6d, 0x66, 0x7e, 0x09, 0x5d, 0x79, 0x59, 0x59,
	0xdf, 0x6e, 0x91, 0x61, 0x16, 0x9c, 0xb5, 0x25, 0xe7, 0x11, 0x7c, 0xb8, 0xc2, 0x59, 0x3d, 0xf7,
	0x25, 0x54, 0x63, 0x9a, 0x8c, 0x3d, 0x97, 0xa7, 0xb8, 0xa4, 0xc5, 0xc4, 0xe6, 0x60, 0x1f, 0xaa,
	0x6a, 0xe2, 0xa0, 0x06, 0x54, 0x47, 0x57, 0x16, 0x7e, 0xf3, 0xe6, 0xbc, 0x53, 0x42, 0x1d, 0x68,
	0x8e, 0xae, 0xac, 0xcb, 0xeb, 0xfe, 0xd9, 0x70, 0x3c, 0x38, 0xc5, 0x1d, 0xed, 0xe0, 0x25, 0x34,
	0x0b, 0x65, 0xd8, 0x81, 0xe6, 0xdb, 0xb1, 0xd5, 0xfb, 0x5d, 0x6f, 0x78, 0xd6, 0xeb, 0x9f, 0x9d,
	0x76, 0x4a, 0x22, 0xc0, 0xdb, 0xb1, 0xf5, 0xfa, 0xfa, 0xec, 0xac, 0xa3, 0x1d, 0x8c, 0xa0, 0x9e,
	0xb5, 0x48, 0x1b, 0x60, 0x34, 0xb6, 0x2e, 0x4f, 0x2f, 0xbe, 0x1e, 0x5e, 0xfc, 0xb6, 0x53, 0x52,
	0x67, 0x7c, 0x7d, 0x71, 0x21, 0xce, 0x9a, 0x3a, 0x8f, 0xaf, 0x4f, 0x4e, 0x4e, 0xc7, 0xe3, 0x8e,
	0x8e, 0x5a, 0x50, 0x1f, 0x8d, 0xad, 0xd7, 0xbd, 0xe1, 0xd9, 0xe9, 0xd7, 0x1d, 0xe3, 0xe8, 0xef,
	0x1a, 0x34, 0x7b, 0xe2, 0xfd, 0x82, 0xd8, 0xdd, 0x29, 0x41, 0xdf, 0x40, 0x33, 0x8f, 0x35, 0xca,
	0x96, 0xab, 0x15, 0xa5, 0xb1, 0xf3, 0x7c, 0x8d, 0x56, 0x65, 0xec, 0x3b, 0x78, 0xba, 0x94, 0x4e,
	0xb4, 0x57, 0xf4, 0x59, 0x81, 0xd3, 0x8e, 0xf9, 0x3e, 0x93, 0x38, 0x76, 0xff, 0xf5, 0x77, 0x9f,
	0xcd, 0x5c, 0x7e, 0x33, 0x9f, 0x1c, 0x4e, 0xa9, 0xff, 0x85, 0xb2, 0xff, 0x42, 0xfe, 0x39, 0x9f,
	0x52, 0x2f, 0x11, 0xfc, 0x55, 0x6f, 0x9d, 0xb9, 0xf7, 0xe4, 0x1b, 0x31, 0x48, 0x85, 0xea, 0xbf,
	0x7a, 0x5b, 0x9d, 0x8f, 0x8f, 0xa5, 0x60, 0xb2, 0x21, 0x5d, 0x7e, 0xfe, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf3, 0x6b, 0x7a, 0x55, 0x05, 0x10, 0x00, 0x00,
}
