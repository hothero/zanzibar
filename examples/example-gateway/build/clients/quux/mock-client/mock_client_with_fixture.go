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

package clientmock

import (
	"github.com/golang/mock/gomock"
)

// MockClientWithFixture is a mock of Client interface with preset fixture
type MockClientWithFixture struct {
	*MockClient
	fixture *ClientFixture
}

// New creates a new mock instance
func New(ctrl *gomock.Controller, fixture *ClientFixture) *MockClientWithFixture {
	return &MockClientWithFixture{
		MockClient: NewMockClient(ctrl),
		fixture:    fixture,
	}
}

// EXPECT shadows the EXPECT method on the underlying mock client.
// It should not be called directly.
func (m *MockClientWithFixture) EXPECT() {
	panic("should not call EXPECT directly.")
}

// EchoStringMock mocks the EchoString method
type EchoStringMock struct {
	scenarios  *EchoStringScenarios
	mockClient *MockClient
}

// ExpectEchoString returns an object that allows the caller to choose expected scenario for EchoString
func (m *MockClientWithFixture) ExpectEchoString() *EchoStringMock {
	return &EchoStringMock{
		scenarios:  m.fixture.EchoString,
		mockClient: m.MockClient,
	}
}

// Success sets the expected scenario as defined in the concrete fixture package
// github.com/uber/zanzibar/examples/example-gateway/clients/quux/fixture
func (s *EchoStringMock) Success() {
	f := s.scenarios.Success

	var arg0 interface{}
	arg0 = f.Arg0
	if f.Arg0Any {
		arg0 = gomock.Any()
	}

	ret0 := f.Ret0

	s.mockClient.EXPECT().EchoString(arg0).Return(ret0)
}
