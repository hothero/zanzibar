// Code generated by zanzibar
// @generated
// Checksum : jH3bjurWR2NjySRaADWlPg==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bar

import (
	json "encoding/json"
	fmt "fmt"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarNormal(in *jlexer.Lexer, out *Bar_Normal_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(BarResponse)
				}
				(*out.Success).UnmarshalEasyJSON(in)
			}
		case "barException":
			if in.IsNull() {
				in.Skip()
				out.BarException = nil
			} else {
				if out.BarException == nil {
					out.BarException = new(BarException)
				}
				(*out.BarException).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarNormal(out *jwriter.Writer, in Bar_Normal_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Success).MarshalEasyJSON(out)
	}
	if in.BarException != nil {
		const prefix string = ",\"barException\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.BarException).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_Normal_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarNormal(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_Normal_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarNormal(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_Normal_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarNormal(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_Normal_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarNormal(l, v)
}
func easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarNormal1(in *jlexer.Lexer, out *Bar_Normal_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var RequestSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "request":
			if in.IsNull() {
				in.Skip()
				out.Request = nil
			} else {
				if out.Request == nil {
					out.Request = new(BarRequest)
				}
				(*out.Request).UnmarshalEasyJSON(in)
			}
			RequestSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !RequestSet {
		in.AddError(fmt.Errorf("key 'request' is required"))
	}
}
func easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarNormal1(out *jwriter.Writer, in Bar_Normal_Args) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"request\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Request == nil {
			out.RawString("null")
		} else {
			(*in.Request).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_Normal_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarNormal1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_Normal_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBea79dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarNormal1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_Normal_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarNormal1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_Normal_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBea79dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarNormal1(l, v)
}
