// Code generated by thriftrw-plugin-yarpc
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

package keyvalueserver

import (
	"context"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/thrift"
	"go.uber.org/yarpc/internal/examples/thrift-keyvalue/keyvalue/kv"
)

// Interface is the server-side interface for the KeyValue service.
type Interface interface {
	GetValue(
		ctx context.Context,
		Key *string,
	) (string, error)

	SetValue(
		ctx context.Context,
		Key *string,
		Value *string,
	) error
}

// New prepares an implementation of the KeyValue service for
// registration.
//
// 	handler := KeyValueHandler{}
// 	dispatcher.Register(keyvalueserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "KeyValue",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "getValue",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.GetValue),
				},
				Signature:    "GetValue(Key *string) (string)",
				ThriftModule: kv.ThriftModule,
			},

			thrift.Method{
				Name: "setValue",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.SetValue),
				},
				Signature:    "SetValue(Key *string, Value *string)",
				ThriftModule: kv.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 2)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

func (h handler) GetValue(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args kv.KeyValue_GetValue_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.GetValue(ctx, args.Key)

	hadError := err != nil
	result, err := kv.KeyValue_GetValue_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) SetValue(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args kv.KeyValue_SetValue_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.SetValue(ctx, args.Key, args.Value)

	hadError := err != nil
	result, err := kv.KeyValue_SetValue_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}
