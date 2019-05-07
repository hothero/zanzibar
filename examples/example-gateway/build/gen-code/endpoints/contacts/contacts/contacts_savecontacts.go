// Code generated by thriftrw v1.19.0. DO NOT EDIT.
// @generated

package contacts

import (
	errors "errors"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

// Contacts_SaveContacts_Args represents the arguments for the Contacts.saveContacts function.
//
// The arguments for saveContacts are sent and received over the wire as this struct.
type Contacts_SaveContacts_Args struct {
	SaveContactsRequest *SaveContactsRequest `json:"saveContactsRequest,required"`
}

// ToWire translates a Contacts_SaveContacts_Args struct into a Thrift-level intermediate
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
func (v *Contacts_SaveContacts_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.SaveContactsRequest == nil {
		return w, errors.New("field SaveContactsRequest of Contacts_SaveContacts_Args is required")
	}
	w, err = v.SaveContactsRequest.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _SaveContactsRequest_Read(w wire.Value) (*SaveContactsRequest, error) {
	var v SaveContactsRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Contacts_SaveContacts_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Contacts_SaveContacts_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Contacts_SaveContacts_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Contacts_SaveContacts_Args) FromWire(w wire.Value) error {
	var err error

	saveContactsRequestIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.SaveContactsRequest, err = _SaveContactsRequest_Read(field.Value)
				if err != nil {
					return err
				}
				saveContactsRequestIsSet = true
			}
		}
	}

	if !saveContactsRequestIsSet {
		return errors.New("field SaveContactsRequest of Contacts_SaveContacts_Args is required")
	}

	return nil
}

// String returns a readable string representation of a Contacts_SaveContacts_Args
// struct.
func (v *Contacts_SaveContacts_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("SaveContactsRequest: %v", v.SaveContactsRequest)
	i++

	return fmt.Sprintf("Contacts_SaveContacts_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Contacts_SaveContacts_Args match the
// provided Contacts_SaveContacts_Args.
//
// This function performs a deep comparison.
func (v *Contacts_SaveContacts_Args) Equals(rhs *Contacts_SaveContacts_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !v.SaveContactsRequest.Equals(rhs.SaveContactsRequest) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Contacts_SaveContacts_Args.
func (v *Contacts_SaveContacts_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	err = multierr.Append(err, enc.AddObject("saveContactsRequest", v.SaveContactsRequest))
	return err
}

// GetSaveContactsRequest returns the value of SaveContactsRequest if it is set or its
// zero value if it is unset.
func (v *Contacts_SaveContacts_Args) GetSaveContactsRequest() (o *SaveContactsRequest) {
	if v != nil {
		o = v.SaveContactsRequest
	}
	return
}

// IsSetSaveContactsRequest returns true if SaveContactsRequest is not nil.
func (v *Contacts_SaveContacts_Args) IsSetSaveContactsRequest() bool {
	return v != nil && v.SaveContactsRequest != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "saveContacts" for this struct.
func (v *Contacts_SaveContacts_Args) MethodName() string {
	return "saveContacts"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Contacts_SaveContacts_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Contacts_SaveContacts_Helper provides functions that aid in handling the
// parameters and return values of the Contacts.saveContacts
// function.
var Contacts_SaveContacts_Helper = struct {
	// Args accepts the parameters of saveContacts in-order and returns
	// the arguments struct for the function.
	Args func(
		saveContactsRequest *SaveContactsRequest,
	) *Contacts_SaveContacts_Args

	// IsException returns true if the given error can be thrown
	// by saveContacts.
	//
	// An error can be thrown by saveContacts only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for saveContacts
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// saveContacts into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by saveContacts
	//
	//   value, err := saveContacts(args)
	//   result, err := Contacts_SaveContacts_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from saveContacts: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*SaveContactsResponse, error) (*Contacts_SaveContacts_Result, error)

	// UnwrapResponse takes the result struct for saveContacts
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if saveContacts threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Contacts_SaveContacts_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Contacts_SaveContacts_Result) (*SaveContactsResponse, error)
}{}

func init() {
	Contacts_SaveContacts_Helper.Args = func(
		saveContactsRequest *SaveContactsRequest,
	) *Contacts_SaveContacts_Args {
		return &Contacts_SaveContacts_Args{
			SaveContactsRequest: saveContactsRequest,
		}
	}

	Contacts_SaveContacts_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Contacts_SaveContacts_Helper.WrapResponse = func(success *SaveContactsResponse, err error) (*Contacts_SaveContacts_Result, error) {
		if err == nil {
			return &Contacts_SaveContacts_Result{Success: success}, nil
		}

		return nil, err
	}
	Contacts_SaveContacts_Helper.UnwrapResponse = func(result *Contacts_SaveContacts_Result) (success *SaveContactsResponse, err error) {

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Contacts_SaveContacts_Result represents the result of a Contacts.saveContacts function call.
//
// The result of a saveContacts execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Contacts_SaveContacts_Result struct {
	// Value returned by saveContacts after a successful execution.
	Success *SaveContactsResponse `json:"success,omitempty"`
}

// ToWire translates a Contacts_SaveContacts_Result struct into a Thrift-level intermediate
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
func (v *Contacts_SaveContacts_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Contacts_SaveContacts_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _SaveContactsResponse_Read(w wire.Value) (*SaveContactsResponse, error) {
	var v SaveContactsResponse
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Contacts_SaveContacts_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Contacts_SaveContacts_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Contacts_SaveContacts_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Contacts_SaveContacts_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _SaveContactsResponse_Read(field.Value)
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
	if count != 1 {
		return fmt.Errorf("Contacts_SaveContacts_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Contacts_SaveContacts_Result
// struct.
func (v *Contacts_SaveContacts_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}

	return fmt.Sprintf("Contacts_SaveContacts_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Contacts_SaveContacts_Result match the
// provided Contacts_SaveContacts_Result.
//
// This function performs a deep comparison.
func (v *Contacts_SaveContacts_Result) Equals(rhs *Contacts_SaveContacts_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Contacts_SaveContacts_Result.
func (v *Contacts_SaveContacts_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Contacts_SaveContacts_Result) GetSuccess() (o *SaveContactsResponse) {
	if v != nil && v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *Contacts_SaveContacts_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "saveContacts" for this struct.
func (v *Contacts_SaveContacts_Result) MethodName() string {
	return "saveContacts"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Contacts_SaveContacts_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
