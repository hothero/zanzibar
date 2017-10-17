// Code generated by zanzibar
// @generated
// Checksum : hSu9t3tnXqTAcJEBG9REXg==
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

func easyjsonD6bc7640DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStructMap(in *jlexer.Lexer, out *Echo_EchoStructMap_Result) {
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
				in.Delim('[')
				if out.Success == nil {
					if !in.IsDelim(']') {
						out.Success = make([]struct {
							Key   *BarResponse
							Value string
						}, 0, 2)
					} else {
						out.Success = []struct {
							Key   *BarResponse
							Value string
						}{}
					}
				} else {
					out.Success = (out.Success)[:0]
				}
				for !in.IsDelim(']') {
					var v1 struct {
						Key   *BarResponse
						Value string
					}
					easyjsonD6bc7640Decode(in, &v1)
					out.Success = append(out.Success, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonD6bc7640EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStructMap(out *jwriter.Writer, in Echo_EchoStructMap_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Success) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"success\":")
		if in.Success == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Success {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjsonD6bc7640Encode(out, v3)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Echo_EchoStructMap_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD6bc7640EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStructMap(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Echo_EchoStructMap_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD6bc7640EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStructMap(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Echo_EchoStructMap_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD6bc7640DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStructMap(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Echo_EchoStructMap_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD6bc7640DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStructMap(l, v)
}
func easyjsonD6bc7640Decode(in *jlexer.Lexer, out *struct {
	Key   *BarResponse
	Value string
}) {
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
		case "Key":
			if in.IsNull() {
				in.Skip()
				out.Key = nil
			} else {
				if out.Key == nil {
					out.Key = new(BarResponse)
				}
				easyjsonD6bc7640DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(in, &*out.Key)
			}
		case "Value":
			out.Value = string(in.String())
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
func easyjsonD6bc7640Encode(out *jwriter.Writer, in struct {
	Key   *BarResponse
	Value string
}) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Key\":")
	if in.Key == nil {
		out.RawString("null")
	} else {
		easyjsonD6bc7640EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(out, *in.Key)
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Value\":")
	out.String(string(in.Value))
	out.RawByte('}')
}
func easyjsonD6bc7640DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(in *jlexer.Lexer, out *BarResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var StringFieldSet bool
	var IntWithRangeSet bool
	var IntWithoutRangeSet bool
	var MapIntWithRangeSet bool
	var MapIntWithoutRangeSet bool
	var BinaryFieldSet bool
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
		case "stringField":
			out.StringField = string(in.String())
			StringFieldSet = true
		case "intWithRange":
			out.IntWithRange = int32(in.Int32())
			IntWithRangeSet = true
		case "intWithoutRange":
			out.IntWithoutRange = int32(in.Int32())
			IntWithoutRangeSet = true
		case "mapIntWithRange":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MapIntWithRange = make(map[UUID]int32)
				} else {
					out.MapIntWithRange = nil
				}
				for !in.IsDelim('}') {
					key := UUID(in.String())
					in.WantColon()
					var v4 int32
					v4 = int32(in.Int32())
					(out.MapIntWithRange)[key] = v4
					in.WantComma()
				}
				in.Delim('}')
			}
			MapIntWithRangeSet = true
		case "mapIntWithoutRange":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MapIntWithoutRange = make(map[string]int32)
				} else {
					out.MapIntWithoutRange = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v5 int32
					v5 = int32(in.Int32())
					(out.MapIntWithoutRange)[key] = v5
					in.WantComma()
				}
				in.Delim('}')
			}
			MapIntWithoutRangeSet = true
		case "binaryField":
			if in.IsNull() {
				in.Skip()
				out.BinaryField = nil
			} else {
				out.BinaryField = in.Bytes()
			}
			BinaryFieldSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !StringFieldSet {
		in.AddError(fmt.Errorf("key 'stringField' is required"))
	}
	if !IntWithRangeSet {
		in.AddError(fmt.Errorf("key 'intWithRange' is required"))
	}
	if !IntWithoutRangeSet {
		in.AddError(fmt.Errorf("key 'intWithoutRange' is required"))
	}
	if !MapIntWithRangeSet {
		in.AddError(fmt.Errorf("key 'mapIntWithRange' is required"))
	}
	if !MapIntWithoutRangeSet {
		in.AddError(fmt.Errorf("key 'mapIntWithoutRange' is required"))
	}
	if !BinaryFieldSet {
		in.AddError(fmt.Errorf("key 'binaryField' is required"))
	}
}
func easyjsonD6bc7640EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(out *jwriter.Writer, in BarResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stringField\":")
	out.String(string(in.StringField))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"intWithRange\":")
	out.Int32(int32(in.IntWithRange))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"intWithoutRange\":")
	out.Int32(int32(in.IntWithoutRange))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mapIntWithRange\":")
	if in.MapIntWithRange == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v7First := true
		for v7Name, v7Value := range in.MapIntWithRange {
			if !v7First {
				out.RawByte(',')
			}
			v7First = false
			out.String(string(v7Name))
			out.RawByte(':')
			out.Int32(int32(v7Value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mapIntWithoutRange\":")
	if in.MapIntWithoutRange == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v8First := true
		for v8Name, v8Value := range in.MapIntWithoutRange {
			if !v8First {
				out.RawByte(',')
			}
			v8First = false
			out.String(string(v8Name))
			out.RawByte(':')
			out.Int32(int32(v8Value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"binaryField\":")
	out.Base64Bytes(in.BinaryField)
	out.RawByte('}')
}
func easyjsonD6bc7640DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStructMap1(in *jlexer.Lexer, out *Echo_EchoStructMap_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var ArgSet bool
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
		case "arg":
			if in.IsNull() {
				in.Skip()
				out.Arg = nil
			} else {
				in.Delim('[')
				if out.Arg == nil {
					if !in.IsDelim(']') {
						out.Arg = make([]struct {
							Key   *BarResponse
							Value string
						}, 0, 2)
					} else {
						out.Arg = []struct {
							Key   *BarResponse
							Value string
						}{}
					}
				} else {
					out.Arg = (out.Arg)[:0]
				}
				for !in.IsDelim(']') {
					var v11 struct {
						Key   *BarResponse
						Value string
					}
					easyjsonD6bc7640Decode(in, &v11)
					out.Arg = append(out.Arg, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
			ArgSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !ArgSet {
		in.AddError(fmt.Errorf("key 'arg' is required"))
	}
}
func easyjsonD6bc7640EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStructMap1(out *jwriter.Writer, in Echo_EchoStructMap_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"arg\":")
	if in.Arg == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v12, v13 := range in.Arg {
			if v12 > 0 {
				out.RawByte(',')
			}
			easyjsonD6bc7640Encode(out, v13)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Echo_EchoStructMap_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD6bc7640EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStructMap1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Echo_EchoStructMap_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD6bc7640EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStructMap1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Echo_EchoStructMap_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD6bc7640DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStructMap1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Echo_EchoStructMap_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD6bc7640DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStructMap1(l, v)
}
