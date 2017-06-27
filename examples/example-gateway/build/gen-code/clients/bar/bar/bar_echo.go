// Code generated by thriftrw v1.3.0
// @generated

package bar

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Bar_Echo_Args struct {
	Name string `json:"name,required"`
}

func (v *Bar_Echo_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
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
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Bar_Echo_Args) FromWire(w wire.Value) error {
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
		}
	}
	if !nameIsSet {
		return errors.New("field Name of Bar_Echo_Args is required")
	}
	return nil
}

func (v *Bar_Echo_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Name: %v", v.Name)
	i++
	return fmt.Sprintf("Bar_Echo_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *Bar_Echo_Args) Equals(rhs *Bar_Echo_Args) bool {
	if !(v.Name == rhs.Name) {
		return false
	}
	return true
}

func (v *Bar_Echo_Args) MethodName() string {
	return "Echo"
}

func (v *Bar_Echo_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var Bar_Echo_Helper = struct {
	Args           func(name string) *Bar_Echo_Args
	IsException    func(error) bool
	WrapResponse   func(string, error) (*Bar_Echo_Result, error)
	UnwrapResponse func(*Bar_Echo_Result) (string, error)
}{}

func init() {
	Bar_Echo_Helper.Args = func(name string) *Bar_Echo_Args {
		return &Bar_Echo_Args{Name: name}
	}
	Bar_Echo_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	Bar_Echo_Helper.WrapResponse = func(success string, err error) (*Bar_Echo_Result, error) {
		if err == nil {
			return &Bar_Echo_Result{Success: &success}, nil
		}
		return nil, err
	}
	Bar_Echo_Helper.UnwrapResponse = func(result *Bar_Echo_Result) (success string, err error) {
		if result.Success != nil {
			success = *result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type Bar_Echo_Result struct {
	Success *string `json:"success,omitempty"`
}

func (v *Bar_Echo_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
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
	if i != 1 {
		return wire.Value{}, fmt.Errorf("Bar_Echo_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Bar_Echo_Result) FromWire(w wire.Value) error {
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
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Bar_Echo_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *Bar_Echo_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}
	return fmt.Sprintf("Bar_Echo_Result{%v}", strings.Join(fields[:i], ", "))
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {
		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func (v *Bar_Echo_Result) Equals(rhs *Bar_Echo_Result) bool {
	if !_String_EqualsPtr(v.Success, rhs.Success) {
		return false
	}
	return true
}

func (v *Bar_Echo_Result) MethodName() string {
	return "Echo"
}

func (v *Bar_Echo_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
