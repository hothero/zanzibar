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

	clientsWithexceptionsWithexceptions "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/withexceptions/withexceptions"
	endpointsWithexceptionsWithexceptions "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/withexceptions/withexceptions"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/withexceptions/module"
	"go.uber.org/zap"
)

// WithExceptionsFunc1Workflow defines the interface for WithExceptionsFunc1 workflow
type WithExceptionsFunc1Workflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
	) (*endpointsWithexceptionsWithexceptions.Response, zanzibar.Header, error)
}

// NewWithExceptionsFunc1Workflow creates a workflow
func NewWithExceptionsFunc1Workflow(deps *module.Dependencies) WithExceptionsFunc1Workflow {
	return &withExceptionsFunc1Workflow{
		Clients: deps.Client,
		Logger:  deps.Default.Logger,
	}
}

// withExceptionsFunc1Workflow calls thrift client Withexceptions.Func1
type withExceptionsFunc1Workflow struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
}

// Handle calls thrift client.
func (w withExceptionsFunc1Workflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
) (*endpointsWithexceptionsWithexceptions.Response, zanzibar.Header, error) {

	clientHeaders := map[string]string{}

	var ok bool
	var h string
	h, ok = reqHeaders.Get("X-Deputy-Forwarded")
	if ok {
		clientHeaders["X-Deputy-Forwarded"] = h
	}

	clientRespBody, _, err := w.Clients.Withexceptions.Func1(
		ctx, clientHeaders,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsWithexceptionsWithexceptions.ExceptionType1:
			serverErr := convertFunc1E1(
				errValue,
			)

			return nil, nil, serverErr

		case *clientsWithexceptionsWithexceptions.ExceptionType2:
			serverErr := convertFunc1E2(
				errValue,
			)

			return nil, nil, serverErr

		default:
			w.Logger.Warn("Client failure: could not make client request",
				zap.Error(errValue),
				zap.String("client", "Withexceptions"),
			)

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertWithExceptionsFunc1ClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertFunc1E1(
	clientError *clientsWithexceptionsWithexceptions.ExceptionType1,
) *endpointsWithexceptionsWithexceptions.EndpointExceptionType1 {
	// TODO: Add error fields mapping here.
	serverError := &endpointsWithexceptionsWithexceptions.EndpointExceptionType1{}
	return serverError
}
func convertFunc1E2(
	clientError *clientsWithexceptionsWithexceptions.ExceptionType2,
) *endpointsWithexceptionsWithexceptions.EndpointExceptionType2 {
	// TODO: Add error fields mapping here.
	serverError := &endpointsWithexceptionsWithexceptions.EndpointExceptionType2{}
	return serverError
}

func convertWithExceptionsFunc1ClientResponse(in *clientsWithexceptionsWithexceptions.Response) *endpointsWithexceptionsWithexceptions.Response {
	out := &endpointsWithexceptionsWithexceptions.Response{}

	return out
}
