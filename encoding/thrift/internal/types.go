// Code generated by thriftrw v0.4.0
// @generated

// Copyright (c) 2016 Uber Technologies, Inc.
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

package internal

import (
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type ExceptionType int32

const (
	ExceptionTypeUnknown               ExceptionType = 0
	ExceptionTypeUnknownMethod         ExceptionType = 1
	ExceptionTypeInvalidMessageType    ExceptionType = 2
	ExceptionTypeWrongMethodName       ExceptionType = 3
	ExceptionTypeBadSequenceID         ExceptionType = 4
	ExceptionTypeMissingResult         ExceptionType = 5
	ExceptionTypeInternalError         ExceptionType = 6
	ExceptionTypeProtocolError         ExceptionType = 7
	ExceptionTypeInvalidTransform      ExceptionType = 8
	ExceptionTypeInvalidProtocol       ExceptionType = 9
	ExceptionTypeUnsupportedClientType ExceptionType = 10
)

func (v ExceptionType) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *ExceptionType) FromWire(w wire.Value) error {
	*v = (ExceptionType)(w.GetI32())
	return nil
}

func (v ExceptionType) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "UNKNOWN"
	case 1:
		return "UNKNOWN_METHOD"
	case 2:
		return "INVALID_MESSAGE_TYPE"
	case 3:
		return "WRONG_METHOD_NAME"
	case 4:
		return "BAD_SEQUENCE_ID"
	case 5:
		return "MISSING_RESULT"
	case 6:
		return "INTERNAL_ERROR"
	case 7:
		return "PROTOCOL_ERROR"
	case 8:
		return "INVALID_TRANSFORM"
	case 9:
		return "INVALID_PROTOCOL"
	case 10:
		return "UNSUPPORTED_CLIENT_TYPE"
	}
	return fmt.Sprintf("ExceptionType(%d)", w)
}

type TApplicationException struct {
	Message *string        `json:"message,omitempty"`
	Type    *ExceptionType `json:"type,omitempty"`
}

func (v *TApplicationException) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Message != nil {
		w, err = wire.NewValueString(*(v.Message)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.Type != nil {
		w, err = v.Type.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ExceptionType_Read(w wire.Value) (ExceptionType, error) {
	var v ExceptionType
	err := v.FromWire(w)
	return v, err
}

func (v *TApplicationException) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Message = &x
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TI32 {
				var x ExceptionType
				x, err = _ExceptionType_Read(field.Value)
				v.Type = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *TApplicationException) String() string {
	var fields [2]string
	i := 0
	if v.Message != nil {
		fields[i] = fmt.Sprintf("Message: %v", *(v.Message))
		i++
	}
	if v.Type != nil {
		fields[i] = fmt.Sprintf("Type: %v", *(v.Type))
		i++
	}
	return fmt.Sprintf("TApplicationException{%v}", strings.Join(fields[:i], ", "))
}

func (v *TApplicationException) Error() string {
	return v.String()
}