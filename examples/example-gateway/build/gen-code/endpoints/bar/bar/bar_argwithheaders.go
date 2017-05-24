// Code generated by thriftrw v1.3.0
// @generated

package bar

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Bar_ArgWithHeaders_Args struct {
	Name     string  `json:"name,required"`
	UserUUID *string `json:"-"`
}

func (v *Bar_ArgWithHeaders_Args) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.Name), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.UserUUID != nil {
		w, err = wire.NewValueString(*(v.UserUUID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Bar_ArgWithHeaders_Args) FromWire(w wire.Value) error {
	var err error
	nameIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Name, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				nameIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.UserUUID = &x
				if err != nil {
					return err
				}
			}
		}
	}
	if !nameIsSet {
		return errors.New("field Name of Bar_ArgWithHeaders_Args is required")
	}
	return nil
}

func (v *Bar_ArgWithHeaders_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Name: %v", v.Name)
	i++
	if v.UserUUID != nil {
		fields[i] = fmt.Sprintf("UserUUID: %v", *(v.UserUUID))
		i++
	}
	return fmt.Sprintf("Bar_ArgWithHeaders_Args{%v}", strings.Join(fields[:i], ", "))
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {
		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func (v *Bar_ArgWithHeaders_Args) Equals(rhs *Bar_ArgWithHeaders_Args) bool {
	if !(v.Name == rhs.Name) {
		return false
	}
	if !_String_EqualsPtr(v.UserUUID, rhs.UserUUID) {
		return false
	}
	return true
}

func (v *Bar_ArgWithHeaders_Args) MethodName() string {
	return "argWithHeaders"
}

func (v *Bar_ArgWithHeaders_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var Bar_ArgWithHeaders_Helper = struct {
	Args           func(name string, userUUID *string) *Bar_ArgWithHeaders_Args
	IsException    func(error) bool
	WrapResponse   func(*BarResponse, error) (*Bar_ArgWithHeaders_Result, error)
	UnwrapResponse func(*Bar_ArgWithHeaders_Result) (*BarResponse, error)
}{}

func init() {
	Bar_ArgWithHeaders_Helper.Args = func(name string, userUUID *string) *Bar_ArgWithHeaders_Args {
		return &Bar_ArgWithHeaders_Args{Name: name, UserUUID: userUUID}
	}
	Bar_ArgWithHeaders_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	Bar_ArgWithHeaders_Helper.WrapResponse = func(success *BarResponse, err error) (*Bar_ArgWithHeaders_Result, error) {
		if err == nil {
			return &Bar_ArgWithHeaders_Result{Success: success}, nil
		}
		return nil, err
	}
	Bar_ArgWithHeaders_Helper.UnwrapResponse = func(result *Bar_ArgWithHeaders_Result) (success *BarResponse, err error) {
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type Bar_ArgWithHeaders_Result struct {
	Success *BarResponse `json:"success,omitempty"`
}

func (v *Bar_ArgWithHeaders_Result) ToWire() (wire.Value, error) {
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
		return wire.Value{}, fmt.Errorf("Bar_ArgWithHeaders_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _BarResponse_Read(w wire.Value) (*BarResponse, error) {
	var v BarResponse
	err := v.FromWire(w)
	return &v, err
}

func (v *Bar_ArgWithHeaders_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _BarResponse_Read(field.Value)
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
		return fmt.Errorf("Bar_ArgWithHeaders_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *Bar_ArgWithHeaders_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	return fmt.Sprintf("Bar_ArgWithHeaders_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *Bar_ArgWithHeaders_Result) Equals(rhs *Bar_ArgWithHeaders_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	return true
}

func (v *Bar_ArgWithHeaders_Result) MethodName() string {
	return "argWithHeaders"
}

func (v *Bar_ArgWithHeaders_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
