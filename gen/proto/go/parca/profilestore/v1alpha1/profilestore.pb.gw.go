// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: parca/profilestore/v1alpha1/profilestore.proto

/*
Package profilestorev1alpha1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package profilestorev1alpha1

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var (
	_ codes.Code
	_ io.Reader
	_ status.Status
	_ = errors.New
	_ = runtime.String
	_ = utilities.NewDoubleArray
	_ = metadata.Join
)

func request_ProfileStoreService_WriteRaw_0(ctx context.Context, marshaler runtime.Marshaler, client ProfileStoreServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq WriteRawRequest
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.WriteRaw(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_ProfileStoreService_WriteRaw_0(ctx context.Context, marshaler runtime.Marshaler, server ProfileStoreServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq WriteRawRequest
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := server.WriteRaw(ctx, &protoReq)
	return msg, metadata, err
}

func request_ProfileStoreService_Write_0(ctx context.Context, marshaler runtime.Marshaler, client ProfileStoreServiceClient, req *http.Request, pathParams map[string]string) (ProfileStoreService_WriteClient, runtime.ServerMetadata, chan error, error) {
	var metadata runtime.ServerMetadata
	errChan := make(chan error, 1)
	stream, err := client.Write(ctx)
	if err != nil {
		grpclog.Errorf("Failed to start streaming: %v", err)
		close(errChan)
		return nil, metadata, errChan, err
	}
	dec := marshaler.NewDecoder(req.Body)
	handleSend := func() error {
		var protoReq WriteRequest
		err := dec.Decode(&protoReq)
		if errors.Is(err, io.EOF) {
			return err
		}
		if err != nil {
			grpclog.Errorf("Failed to decode request: %v", err)
			return status.Errorf(codes.InvalidArgument, "Failed to decode request: %v", err)
		}
		if err := stream.Send(&protoReq); err != nil {
			grpclog.Errorf("Failed to send request: %v", err)
			return err
		}
		return nil
	}
	go func() {
		defer close(errChan)
		for {
			if err := handleSend(); err != nil {
				errChan <- err
				break
			}
		}
		if err := stream.CloseSend(); err != nil {
			grpclog.Errorf("Failed to terminate client stream: %v", err)
		}
	}()
	header, err := stream.Header()
	if err != nil {
		grpclog.Errorf("Failed to get header from client: %v", err)
		return nil, metadata, errChan, err
	}
	metadata.HeaderMD = header
	return stream, metadata, errChan, nil
}

func request_AgentsService_Agents_0(ctx context.Context, marshaler runtime.Marshaler, client AgentsServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq AgentsRequest
		metadata runtime.ServerMetadata
	)
	msg, err := client.Agents(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_AgentsService_Agents_0(ctx context.Context, marshaler runtime.Marshaler, server AgentsServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq AgentsRequest
		metadata runtime.ServerMetadata
	)
	msg, err := server.Agents(ctx, &protoReq)
	return msg, metadata, err
}

// RegisterProfileStoreServiceHandlerServer registers the http handlers for service ProfileStoreService to "mux".
// UnaryRPC     :call ProfileStoreServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterProfileStoreServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterProfileStoreServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ProfileStoreServiceServer) error {
	mux.Handle(http.MethodPost, pattern_ProfileStoreService_WriteRaw_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/parca.profilestore.v1alpha1.ProfileStoreService/WriteRaw", runtime.WithHTTPPathPattern("/profiles/writeraw"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ProfileStoreService_WriteRaw_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_ProfileStoreService_WriteRaw_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	mux.Handle(http.MethodPost, pattern_ProfileStoreService_Write_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterAgentsServiceHandlerServer registers the http handlers for service AgentsService to "mux".
// UnaryRPC     :call AgentsServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterAgentsServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterAgentsServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server AgentsServiceServer) error {
	mux.Handle(http.MethodGet, pattern_AgentsService_Agents_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/parca.profilestore.v1alpha1.AgentsService/Agents", runtime.WithHTTPPathPattern("/agents"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_AgentsService_Agents_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_AgentsService_Agents_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	return nil
}

// RegisterProfileStoreServiceHandlerFromEndpoint is same as RegisterProfileStoreServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterProfileStoreServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterProfileStoreServiceHandler(ctx, mux, conn)
}

// RegisterProfileStoreServiceHandler registers the http handlers for service ProfileStoreService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterProfileStoreServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterProfileStoreServiceHandlerClient(ctx, mux, NewProfileStoreServiceClient(conn))
}

// RegisterProfileStoreServiceHandlerClient registers the http handlers for service ProfileStoreService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ProfileStoreServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ProfileStoreServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ProfileStoreServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterProfileStoreServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ProfileStoreServiceClient) error {
	mux.Handle(http.MethodPost, pattern_ProfileStoreService_WriteRaw_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/parca.profilestore.v1alpha1.ProfileStoreService/WriteRaw", runtime.WithHTTPPathPattern("/profiles/writeraw"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ProfileStoreService_WriteRaw_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_ProfileStoreService_WriteRaw_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_ProfileStoreService_Write_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/parca.profilestore.v1alpha1.ProfileStoreService/Write", runtime.WithHTTPPathPattern("/profiles/write"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		resp, md, reqErrChan, err := request_ProfileStoreService_Write_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		go func() {
			for err := range reqErrChan {
				if err != nil && !errors.Is(err, io.EOF) {
					runtime.HTTPStreamError(annotatedContext, mux, outboundMarshaler, w, req, err)
				}
			}
		}()
		forward_ProfileStoreService_Write_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_ProfileStoreService_WriteRaw_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"profiles", "writeraw"}, ""))
	pattern_ProfileStoreService_Write_0    = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"profiles", "write"}, ""))
)

var (
	forward_ProfileStoreService_WriteRaw_0 = runtime.ForwardResponseMessage
	forward_ProfileStoreService_Write_0    = runtime.ForwardResponseStream
)

// RegisterAgentsServiceHandlerFromEndpoint is same as RegisterAgentsServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterAgentsServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterAgentsServiceHandler(ctx, mux, conn)
}

// RegisterAgentsServiceHandler registers the http handlers for service AgentsService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterAgentsServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterAgentsServiceHandlerClient(ctx, mux, NewAgentsServiceClient(conn))
}

// RegisterAgentsServiceHandlerClient registers the http handlers for service AgentsService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "AgentsServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "AgentsServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "AgentsServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterAgentsServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client AgentsServiceClient) error {
	mux.Handle(http.MethodGet, pattern_AgentsService_Agents_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/parca.profilestore.v1alpha1.AgentsService/Agents", runtime.WithHTTPPathPattern("/agents"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_AgentsService_Agents_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_AgentsService_Agents_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_AgentsService_Agents_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"agents"}, ""))
)

var (
	forward_AgentsService_Agents_0 = runtime.ForwardResponseMessage
)
