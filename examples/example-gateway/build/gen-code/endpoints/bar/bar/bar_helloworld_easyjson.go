// Code generated by zanzibar
// @generated
// Checksum : dXUtrHnCuVwvbVJeqmY6fA==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bar

import (
	json "encoding/json"
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

func easyjson687c569aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarHelloWorld(in *jlexer.Lexer, out *Bar_HelloWorld_Result) {
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
					out.Success = new(string)
				}
				*out.Success = string(in.String())
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
func easyjson687c569aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarHelloWorld(out *jwriter.Writer, in Bar_HelloWorld_Result) {
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
		out.String(string(*in.Success))
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
func (v Bar_HelloWorld_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson687c569aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarHelloWorld(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_HelloWorld_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson687c569aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarHelloWorld(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_HelloWorld_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson687c569aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarHelloWorld(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_HelloWorld_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson687c569aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarHelloWorld(l, v)
}
func easyjson687c569aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarHelloWorld1(in *jlexer.Lexer, out *Bar_HelloWorld_Args) {
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
func easyjson687c569aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarHelloWorld1(out *jwriter.Writer, in Bar_HelloWorld_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_HelloWorld_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson687c569aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarHelloWorld1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_HelloWorld_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson687c569aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarHelloWorld1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_HelloWorld_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson687c569aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarHelloWorld1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_HelloWorld_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson687c569aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarHelloWorld1(l, v)
}
