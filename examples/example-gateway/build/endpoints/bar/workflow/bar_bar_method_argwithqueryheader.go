// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package workflow

import (
	"context"
	"net/textproto"

	"github.com/uber/zanzibar/config"

	zanzibar "github.com/uber/zanzibar/runtime"

	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
	"go.uber.org/zap"
)

// BarArgWithQueryHeaderWorkflow defines the interface for BarArgWithQueryHeader workflow
type BarArgWithQueryHeaderWorkflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
		r *endpointsBarBar.Bar_ArgWithQueryHeader_Args,
	) (*endpointsBarBar.BarResponse, zanzibar.Header, error)
}

// NewBarArgWithQueryHeaderWorkflow creates a workflow
func NewBarArgWithQueryHeaderWorkflow(deps *module.Dependencies) BarArgWithQueryHeaderWorkflow {
	var whitelistedDynamicHeaders []string
	if deps.Default.Config.ContainsKey("clients.bar.alternates") {
		var alternateServiceDetail config.AlternateServiceDetail
		deps.Default.Config.MustGetStruct("clients.bar.alternates", &alternateServiceDetail)
		for _, routingConfig := range alternateServiceDetail.RoutingConfigs {
			whitelistedDynamicHeaders = append(whitelistedDynamicHeaders, routingConfig.HeaderName)
		}
	}

	return &barArgWithQueryHeaderWorkflow{
		Clients:                   deps.Client,
		Logger:                    deps.Default.Logger,
		whitelistedDynamicHeaders: whitelistedDynamicHeaders,
	}
}

// barArgWithQueryHeaderWorkflow calls thrift client Bar.ArgWithQueryHeader
type barArgWithQueryHeaderWorkflow struct {
	Clients                   *module.ClientDependencies
	Logger                    *zap.Logger
	whitelistedDynamicHeaders []string
}

// Handle calls thrift client.
func (w barArgWithQueryHeaderWorkflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBarBar.Bar_ArgWithQueryHeader_Args,
) (*endpointsBarBar.BarResponse, zanzibar.Header, error) {
	clientRequest := convertToArgWithQueryHeaderClientRequest(r)

	clientHeaders := map[string]string{}

	var ok bool
	var h string
	h, ok = reqHeaders.Get("X-Deputy-Forwarded")
	if ok {
		clientHeaders["X-Deputy-Forwarded"] = h
	}
	for _, whitelistedHeader := range w.whitelistedDynamicHeaders {
		transformedHeaderName := textproto.CanonicalMIMEHeaderKey(whitelistedHeader)
		headerVal, ok := reqHeaders.Get(transformedHeaderName)
		if ok {
			clientHeaders[transformedHeaderName] = headerVal
		}
	}

	clientRespBody, _, err := w.Clients.Bar.ArgWithQueryHeader(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		default:
			w.Logger.Warn("Client failure: could not make client request",
				zap.Error(errValue),
				zap.String("client", "Bar"),
			)

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertBarArgWithQueryHeaderClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToArgWithQueryHeaderClientRequest(in *endpointsBarBar.Bar_ArgWithQueryHeader_Args) *clientsBarBar.Bar_ArgWithQueryHeader_Args {
	out := &clientsBarBar.Bar_ArgWithQueryHeader_Args{}

	out.UserUUID = (*string)(in.UserUUID)

	return out
}

func convertBarArgWithQueryHeaderClientResponse(in *clientsBarBar.BarResponse) *endpointsBarBar.BarResponse {
	out := &endpointsBarBar.BarResponse{}

	out.StringField = string(in.StringField)
	out.IntWithRange = int32(in.IntWithRange)
	out.IntWithoutRange = int32(in.IntWithoutRange)
	out.MapIntWithRange = make(map[endpointsBarBar.UUID]int32, len(in.MapIntWithRange))
	for key1, value2 := range in.MapIntWithRange {
		out.MapIntWithRange[endpointsBarBar.UUID(key1)] = int32(value2)
	}
	out.MapIntWithoutRange = make(map[string]int32, len(in.MapIntWithoutRange))
	for key3, value4 := range in.MapIntWithoutRange {
		out.MapIntWithoutRange[key3] = int32(value4)
	}
	out.BinaryField = []byte(in.BinaryField)
	var convertBarResponseHelper5 func(in *clientsBarBar.BarResponse) (out *endpointsBarBar.BarResponse)
	convertBarResponseHelper5 = func(in *clientsBarBar.BarResponse) (out *endpointsBarBar.BarResponse) {
		if in != nil {
			out = &endpointsBarBar.BarResponse{}
			out.StringField = string(in.StringField)
			out.IntWithRange = int32(in.IntWithRange)
			out.IntWithoutRange = int32(in.IntWithoutRange)
			out.MapIntWithRange = make(map[endpointsBarBar.UUID]int32, len(in.MapIntWithRange))
			for key6, value7 := range in.MapIntWithRange {
				out.MapIntWithRange[endpointsBarBar.UUID(key6)] = int32(value7)
			}
			out.MapIntWithoutRange = make(map[string]int32, len(in.MapIntWithoutRange))
			for key8, value9 := range in.MapIntWithoutRange {
				out.MapIntWithoutRange[key8] = int32(value9)
			}
			out.BinaryField = []byte(in.BinaryField)
			out.NextResponse = convertBarResponseHelper5(in.NextResponse)
		} else {
			out = nil
		}
		return
	}
	out.NextResponse = convertBarResponseHelper5(in.NextResponse)

	return out
}
