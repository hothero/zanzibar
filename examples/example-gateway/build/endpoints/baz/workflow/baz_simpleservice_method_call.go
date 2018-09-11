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

	zanzibar "github.com/uber/zanzibar/runtime"

	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
	endpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/baz/baz"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/module"
	"go.uber.org/zap"
)

// SimpleServiceCallWorkflow defines the interface for SimpleServiceCall workflow
type SimpleServiceCallWorkflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
		r *endpointsBazBaz.SimpleService_Call_Args,
	) (zanzibar.Header, error)
}

// NewSimpleServiceCallWorkflow creates a workflow
func NewSimpleServiceCallWorkflow(deps *module.Dependencies) SimpleServiceCallWorkflow {
	return &simpleServiceCallWorkflow{
		Clients: deps.Client,
		Logger:  deps.Default.Logger,
	}
}

// simpleServiceCallWorkflow calls thrift client Baz.Call
type simpleServiceCallWorkflow struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
}

// Handle calls thrift client.
func (w simpleServiceCallWorkflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBazBaz.SimpleService_Call_Args,
) (zanzibar.Header, error) {
	clientRequest := convertToCallClientRequest(r)

	clientHeaders := map[string]string{}

	var ok bool
	var h string
	h, ok = reqHeaders.Get("X-Deputy-Forwarded")
	if ok {
		clientHeaders["X-Deputy-Forwarded"] = h
	}
	h, ok = reqHeaders.Get("X-Shadow-Start-Trace-Id")
	if ok {
		clientHeaders["X-Shadow-Start-Trace-Id"] = h
	}
	h, ok = reqHeaders.Get("X-Token")
	if ok {
		clientHeaders["X-Token"] = h
	}
	h, ok = reqHeaders.Get("X-Uuid")
	if ok {
		clientHeaders["X-Uuid"] = h
	}
	h, ok = reqHeaders.Get("X-Zanzibar-Use-Staging")
	if ok {
		clientHeaders["X-Zanzibar-Use-Staging"] = h
	}

	cliRespHeaders, err := w.Clients.Baz.Call(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBazBaz.AuthErr:
			serverErr := convertCallAuthErr(
				errValue,
			)

			return nil, serverErr

		default:
			w.Logger.Warn("Could not make client request",
				zap.Error(errValue),
				zap.String("client", "Baz"),
			)

			return nil, err

		}
	}

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}
	resHeaders.Set("Some-Res-Header", cliRespHeaders["Some-Res-Header"])

	return resHeaders, nil
}

func convertToCallClientRequest(in *endpointsBazBaz.SimpleService_Call_Args) *clientsBazBaz.SimpleService_Call_Args {
	out := &clientsBazBaz.SimpleService_Call_Args{}

	if in.Arg != nil {
		out.Arg = &clientsBazBaz.BazRequest{}
		out.Arg.B1 = bool(in.Arg.B1)
		out.Arg.S2 = string(in.Arg.S2)
		out.Arg.I3 = int32(in.Arg.I3)
	} else {
		out.Arg = nil
	}
	out.I64Optional = (*int64)(in.I64Optional)
	out.TestUUID = (*clientsBazBaz.UUID)(in.TestUUID)

	return out
}

func convertCallAuthErr(
	clientError *clientsBazBaz.AuthErr,
) *endpointsBazBaz.AuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.AuthErr{}
	return serverError
}
