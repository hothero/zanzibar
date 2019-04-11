// Code generated by zanzibar
// @generated
// Checksum : 5Fk021vPF9RJRBO2ZFtXCw==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package baz

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

func easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceSillyNoop(in *jlexer.Lexer, out *SimpleService_SillyNoop_Result) {
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
		case "authErr":
			if in.IsNull() {
				in.Skip()
				out.AuthErr = nil
			} else {
				if out.AuthErr == nil {
					out.AuthErr = new(AuthErr)
				}
				(*out.AuthErr).UnmarshalEasyJSON(in)
			}
		case "serverErr":
			if in.IsNull() {
				in.Skip()
				out.ServerErr = nil
			} else {
				if out.ServerErr == nil {
					out.ServerErr = new(ServerErr)
				}
				(*out.ServerErr).UnmarshalEasyJSON(in)
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
func easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceSillyNoop(out *jwriter.Writer, in SimpleService_SillyNoop_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AuthErr != nil {
		const prefix string = ",\"authErr\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.AuthErr).MarshalEasyJSON(out)
	}
	if in.ServerErr != nil {
		const prefix string = ",\"serverErr\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.ServerErr).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SimpleService_SillyNoop_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceSillyNoop(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SimpleService_SillyNoop_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceSillyNoop(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SimpleService_SillyNoop_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceSillyNoop(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SimpleService_SillyNoop_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceSillyNoop(l, v)
}
func easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceSillyNoop1(in *jlexer.Lexer, out *SimpleService_SillyNoop_Args) {
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
func easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceSillyNoop1(out *jwriter.Writer, in SimpleService_SillyNoop_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SimpleService_SillyNoop_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceSillyNoop1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SimpleService_SillyNoop_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceSillyNoop1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SimpleService_SillyNoop_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceSillyNoop1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SimpleService_SillyNoop_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceSillyNoop1(l, v)
}
