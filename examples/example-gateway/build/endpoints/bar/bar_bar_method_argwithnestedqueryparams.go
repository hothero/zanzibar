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

package barEndpoint

import (
	"context"

	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/thriftrw/ptr"
	"go.uber.org/zap"

	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
)

// BarArgWithNestedQueryParamsHandler is the handler for "/bar/argWithNestedQueryParams"
type BarArgWithNestedQueryParamsHandler struct {
	Clients  *module.ClientDependencies
	endpoint *zanzibar.RouterEndpoint
}

// NewBarArgWithNestedQueryParamsHandler creates a handler
func NewBarArgWithNestedQueryParamsHandler(deps *module.Dependencies) *BarArgWithNestedQueryParamsHandler {
	handler := &BarArgWithNestedQueryParamsHandler{
		Clients: deps.Client,
	}
	handler.endpoint = zanzibar.NewRouterEndpoint(
		deps.Default.Logger, deps.Default.Scope,
		"bar", "argWithNestedQueryParams",
		handler.HandleRequest,
	)
	return handler
}

// Register adds the http handler to the gateway's http router
func (h *BarArgWithNestedQueryParamsHandler) Register(g *zanzibar.Gateway) error {
	g.HTTPRouter.Register(
		"GET", "/bar/argWithNestedQueryParams",
		h.endpoint,
	)
	// TODO: register should return errors on route conflicts
	return nil
}

// HandleRequest handles "/bar/argWithNestedQueryParams".
func (h *BarArgWithNestedQueryParamsHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	var requestBody endpointsBarBar.Bar_ArgWithNestedQueryParams_Args

	if requestBody.Request == nil {
		requestBody.Request = &endpointsBarBar.QueryParamsStruct{}
	}
	xUUIDValue, xUUIDValueExists := req.Header.Get("x-uuid")
	if xUUIDValueExists {
		requestBody.Request.AuthUUID = ptr.String(xUUIDValue)
	}
	xUUID2Value, xUUID2ValueExists := req.Header.Get("x-uuid2")
	if xUUID2ValueExists {
		requestBody.Request.AuthUUID2 = ptr.String(xUUID2Value)
	}

	if requestBody.Request == nil {
		requestBody.Request = &endpointsBarBar.QueryParamsStruct{}
	}
	requestNameOk := req.CheckQueryValue("request.name")
	if !requestNameOk {
		return
	}
	requestNameQuery, ok := req.GetQueryValue("request.name")
	if !ok {
		return
	}
	requestBody.Request.Name = requestNameQuery

	requestUserUUIDOk := req.HasQueryValue("request.userUUID")
	if requestUserUUIDOk {
		requestUserUUIDQuery, ok := req.GetQueryValue("request.userUUID")
		if !ok {
			return
		}
		requestBody.Request.UserUUID = ptr.String(requestUserUUIDQuery)
	}

	if req.HasQueryPrefix("opt") || requestBody.Opt != nil {
		if requestBody.Opt == nil {
			requestBody.Opt = &endpointsBarBar.QueryParamsOptsStruct{}
		}
		optNameOk := req.CheckQueryValue("opt.name")
		if !optNameOk {
			return
		}
		optNameQuery, ok := req.GetQueryValue("opt.name")
		if !ok {
			return
		}
		requestBody.Opt.Name = optNameQuery

		optUserUUIDOk := req.HasQueryValue("opt.userUUID")
		if optUserUUIDOk {
			optUserUUIDQuery, ok := req.GetQueryValue("opt.userUUID")
			if !ok {
				return
			}
			requestBody.Opt.UserUUID = ptr.String(optUserUUIDQuery)
		}

		optAuthUUIDOk := req.HasQueryValue("opt.authUUID")
		if optAuthUUIDOk {
			optAuthUUIDQuery, ok := req.GetQueryValue("opt.authUUID")
			if !ok {
				return
			}
			requestBody.Opt.AuthUUID = ptr.String(optAuthUUIDQuery)
		}

		optAuthUUID2Ok := req.HasQueryValue("opt.authUUID2")
		if optAuthUUID2Ok {
			optAuthUUID2Query, ok := req.GetQueryValue("opt.authUUID2")
			if !ok {
				return
			}
			requestBody.Opt.AuthUUID2 = ptr.String(optAuthUUID2Query)
		}

	}

	workflow := BarArgWithNestedQueryParamsEndpoint{
		Clients: h.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, cliRespHeaders, err := workflow.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		switch errValue := err.(type) {

		default:
			req.Logger.Warn("Workflow for endpoint returned error", zap.Error(errValue))
			res.SendErrorString(500, "Unexpected server error")
			return
		}
	}
	// TODO(jakev): implement writing fields into response headers

	res.WriteJSON(200, cliRespHeaders, response)
}

// BarArgWithNestedQueryParamsEndpoint calls thrift client Bar.ArgWithNestedQueryParams
type BarArgWithNestedQueryParamsEndpoint struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w BarArgWithNestedQueryParamsEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBarBar.Bar_ArgWithNestedQueryParams_Args,
) (*endpointsBarBar.BarResponse, zanzibar.Header, error) {
	clientRequest := convertToArgWithNestedQueryParamsClientRequest(r)

	clientHeaders := map[string]string{}

	clientRespBody, _, err := w.Clients.Bar.ArgWithNestedQueryParams(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		default:
			w.Logger.Warn("Could not make client request", zap.Error(errValue))
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertBarArgWithNestedQueryParamsClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToArgWithNestedQueryParamsClientRequest(in *endpointsBarBar.Bar_ArgWithNestedQueryParams_Args) *clientsBarBar.Bar_ArgWithNestedQueryParams_Args {
	out := &clientsBarBar.Bar_ArgWithNestedQueryParams_Args{}

	if in.Request != nil {
		out.Request = &clientsBarBar.QueryParamsStruct{}
		out.Request.Name = string(in.Request.Name)
		out.Request.UserUUID = (*string)(in.Request.UserUUID)
		out.Request.AuthUUID = (*string)(in.Request.AuthUUID)
		out.Request.AuthUUID2 = (*string)(in.Request.AuthUUID2)
	} else {
		out.Request = nil
	}
	if in.Opt != nil {
		out.Opt = &clientsBarBar.QueryParamsOptsStruct{}
		out.Opt.Name = string(in.Opt.Name)
		out.Opt.UserUUID = (*string)(in.Opt.UserUUID)
		out.Opt.AuthUUID = (*string)(in.Opt.AuthUUID)
		out.Opt.AuthUUID2 = (*string)(in.Opt.AuthUUID2)
	} else {
		out.Opt = nil
	}

	return out
}

func convertBarArgWithNestedQueryParamsClientResponse(in *clientsBarBar.BarResponse) *endpointsBarBar.BarResponse {
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

	return out
}
