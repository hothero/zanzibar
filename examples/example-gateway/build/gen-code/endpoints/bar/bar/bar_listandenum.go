// Code generated by thriftrw v1.16.1. DO NOT EDIT.
// @generated

package bar

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// Bar_ListAndEnum_Args represents the arguments for the Bar.listAndEnum function.
//
// The arguments for listAndEnum are sent and received over the wire as this struct.
type Bar_ListAndEnum_Args struct {
	DemoIds  []string  `json:"demoIds,required"`
	DemoType *DemoType `json:"demoType,omitempty"`
}

// ToWire translates a Bar_ListAndEnum_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Bar_ListAndEnum_Args) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.DemoIds == nil {
		return w, errors.New("field DemoIds of Bar_ListAndEnum_Args is required")
	}
	w, err = wire.NewValueList(_List_String_ValueList(v.DemoIds)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.DemoType != nil {
		w, err = v.DemoType.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _DemoType_Read(w wire.Value) (DemoType, error) {
	var v DemoType
	err := v.FromWire(w)
	return v, err
}

// FromWire deserializes a Bar_ListAndEnum_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_ListAndEnum_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_ListAndEnum_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_ListAndEnum_Args) FromWire(w wire.Value) error {
	var err error

	demoIdsIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TList {
				v.DemoIds, err = _List_String_Read(field.Value.GetList())
				if err != nil {
					return err
				}
				demoIdsIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TI32 {
				var x DemoType
				x, err = _DemoType_Read(field.Value)
				v.DemoType = &x
				if err != nil {
					return err
				}

			}
		}
	}

	if !demoIdsIsSet {
		return errors.New("field DemoIds of Bar_ListAndEnum_Args is required")
	}

	return nil
}

// String returns a readable string representation of a Bar_ListAndEnum_Args
// struct.
func (v *Bar_ListAndEnum_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("DemoIds: %v", v.DemoIds)
	i++
	if v.DemoType != nil {
		fields[i] = fmt.Sprintf("DemoType: %v", *(v.DemoType))
		i++
	}

	return fmt.Sprintf("Bar_ListAndEnum_Args{%v}", strings.Join(fields[:i], ", "))
}

func _DemoType_EqualsPtr(lhs, rhs *DemoType) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return x.Equals(y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this Bar_ListAndEnum_Args match the
// provided Bar_ListAndEnum_Args.
//
// This function performs a deep comparison.
func (v *Bar_ListAndEnum_Args) Equals(rhs *Bar_ListAndEnum_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_List_String_Equals(v.DemoIds, rhs.DemoIds) {
		return false
	}
	if !_DemoType_EqualsPtr(v.DemoType, rhs.DemoType) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Bar_ListAndEnum_Args.
func (v *Bar_ListAndEnum_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	err = multierr.Append(err, enc.AddArray("demoIds", (_List_String_Zapper)(v.DemoIds)))
	if v.DemoType != nil {
		err = multierr.Append(err, enc.AddObject("demoType", *v.DemoType))
	}
	return err
}

// GetDemoIds returns the value of DemoIds if it is set or its
// zero value if it is unset.
func (v *Bar_ListAndEnum_Args) GetDemoIds() (o []string) {
	if v != nil {
		o = v.DemoIds
	}
	return
}

// IsSetDemoIds returns true if DemoIds is not nil.
func (v *Bar_ListAndEnum_Args) IsSetDemoIds() bool {
	return v != nil && v.DemoIds != nil
}

// GetDemoType returns the value of DemoType if it is set or its
// zero value if it is unset.
func (v *Bar_ListAndEnum_Args) GetDemoType() (o DemoType) {
	if v != nil && v.DemoType != nil {
		return *v.DemoType
	}

	return
}

// IsSetDemoType returns true if DemoType is not nil.
func (v *Bar_ListAndEnum_Args) IsSetDemoType() bool {
	return v != nil && v.DemoType != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "listAndEnum" for this struct.
func (v *Bar_ListAndEnum_Args) MethodName() string {
	return "listAndEnum"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Bar_ListAndEnum_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Bar_ListAndEnum_Helper provides functions that aid in handling the
// parameters and return values of the Bar.listAndEnum
// function.
var Bar_ListAndEnum_Helper = struct {
	// Args accepts the parameters of listAndEnum in-order and returns
	// the arguments struct for the function.
	Args func(
		demoIds []string,
		demoType *DemoType,
	) *Bar_ListAndEnum_Args

	// IsException returns true if the given error can be thrown
	// by listAndEnum.
	//
	// An error can be thrown by listAndEnum only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for listAndEnum
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// listAndEnum into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by listAndEnum
	//
	//   value, err := listAndEnum(args)
	//   result, err := Bar_ListAndEnum_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from listAndEnum: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(string, error) (*Bar_ListAndEnum_Result, error)

	// UnwrapResponse takes the result struct for listAndEnum
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if listAndEnum threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Bar_ListAndEnum_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Bar_ListAndEnum_Result) (string, error)
}{}

func init() {
	Bar_ListAndEnum_Helper.Args = func(
		demoIds []string,
		demoType *DemoType,
	) *Bar_ListAndEnum_Args {
		return &Bar_ListAndEnum_Args{
			DemoIds:  demoIds,
			DemoType: demoType,
		}
	}

	Bar_ListAndEnum_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BarException:
			return true
		default:
			return false
		}
	}

	Bar_ListAndEnum_Helper.WrapResponse = func(success string, err error) (*Bar_ListAndEnum_Result, error) {
		if err == nil {
			return &Bar_ListAndEnum_Result{Success: &success}, nil
		}

		switch e := err.(type) {
		case *BarException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Bar_ListAndEnum_Result.BarException")
			}
			return &Bar_ListAndEnum_Result{BarException: e}, nil
		}

		return nil, err
	}
	Bar_ListAndEnum_Helper.UnwrapResponse = func(result *Bar_ListAndEnum_Result) (success string, err error) {
		if result.BarException != nil {
			err = result.BarException
			return
		}

		if result.Success != nil {
			success = *result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Bar_ListAndEnum_Result represents the result of a Bar.listAndEnum function call.
//
// The result of a listAndEnum execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Bar_ListAndEnum_Result struct {
	// Value returned by listAndEnum after a successful execution.
	Success      *string       `json:"success,omitempty"`
	BarException *BarException `json:"barException,omitempty"`
}

// ToWire translates a Bar_ListAndEnum_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Bar_ListAndEnum_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueString(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.BarException != nil {
		w, err = v.BarException.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Bar_ListAndEnum_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Bar_ListAndEnum_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_ListAndEnum_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_ListAndEnum_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_ListAndEnum_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Success = &x
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BarException, err = _BarException_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.BarException != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Bar_ListAndEnum_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Bar_ListAndEnum_Result
// struct.
func (v *Bar_ListAndEnum_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}
	if v.BarException != nil {
		fields[i] = fmt.Sprintf("BarException: %v", v.BarException)
		i++
	}

	return fmt.Sprintf("Bar_ListAndEnum_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Bar_ListAndEnum_Result match the
// provided Bar_ListAndEnum_Result.
//
// This function performs a deep comparison.
func (v *Bar_ListAndEnum_Result) Equals(rhs *Bar_ListAndEnum_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Success, rhs.Success) {
		return false
	}
	if !((v.BarException == nil && rhs.BarException == nil) || (v.BarException != nil && rhs.BarException != nil && v.BarException.Equals(rhs.BarException))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Bar_ListAndEnum_Result.
func (v *Bar_ListAndEnum_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		enc.AddString("success", *v.Success)
	}
	if v.BarException != nil {
		err = multierr.Append(err, enc.AddObject("barException", v.BarException))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Bar_ListAndEnum_Result) GetSuccess() (o string) {
	if v != nil && v.Success != nil {
		return *v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *Bar_ListAndEnum_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// GetBarException returns the value of BarException if it is set or its
// zero value if it is unset.
func (v *Bar_ListAndEnum_Result) GetBarException() (o *BarException) {
	if v != nil && v.BarException != nil {
		return v.BarException
	}

	return
}

// IsSetBarException returns true if BarException is not nil.
func (v *Bar_ListAndEnum_Result) IsSetBarException() bool {
	return v != nil && v.BarException != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "listAndEnum" for this struct.
func (v *Bar_ListAndEnum_Result) MethodName() string {
	return "listAndEnum"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Bar_ListAndEnum_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
