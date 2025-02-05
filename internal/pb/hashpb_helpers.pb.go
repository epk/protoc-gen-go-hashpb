// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb (devel)

package pb

import (
	protowire "google.golang.org/protobuf/encoding/protowire"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	hash "hash"
	math "math"
	sort "sort"
)

func cerbos_hashpb_test_NestedTestAllTypes_hashpb_sum(m *NestedTestAllTypes, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.hashpb.test.NestedTestAllTypes.child"]; !ok {
		if m.Child != nil {
			cerbos_hashpb_test_NestedTestAllTypes_hashpb_sum(m.Child, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.hashpb.test.NestedTestAllTypes.payload"]; !ok {
		if m.Payload != nil {
			cerbos_hashpb_test_TestAllTypes_hashpb_sum(m.Payload, hasher, ignore)
		}

	}
}

func cerbos_hashpb_test_NoFields_hashpb_sum(m *NoFields, hasher hash.Hash, ignore map[string]struct{}) {
}

func cerbos_hashpb_test_TestAllTypes_NestedMessage_hashpb_sum(m *TestAllTypes_NestedMessage, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.NestedMessage.bb"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Bb)))

	}
}

func cerbos_hashpb_test_TestAllTypes_hashpb_sum(m *TestAllTypes, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_int32"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.SingleInt32)))

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_int64"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.SingleInt64)))

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_uint32"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.SingleUint32)))

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_uint64"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, m.SingleUint64))

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_sint32"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeZigZag(int64(m.SingleSint32))))

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_sint64"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeZigZag(m.SingleSint64)))

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_fixed32"]; !ok {
		_, _ = hasher.Write(protowire.AppendFixed32(nil, uint32(m.SingleFixed32)))

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_fixed64"]; !ok {
		_, _ = hasher.Write(protowire.AppendFixed64(nil, m.SingleFixed64))

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_sfixed32"]; !ok {
		_, _ = hasher.Write(protowire.AppendFixed32(nil, uint32(m.SingleSfixed32)))

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_sfixed64"]; !ok {
		_, _ = hasher.Write(protowire.AppendFixed64(nil, uint64(m.SingleSfixed64)))

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_float"]; !ok {
		_, _ = hasher.Write(protowire.AppendFixed32(nil, math.Float32bits(m.SingleFloat)))

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_double"]; !ok {
		_, _ = hasher.Write(protowire.AppendFixed64(nil, math.Float64bits(m.SingleDouble)))

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_bool"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.SingleBool)))

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_string"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.SingleString))

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_bytes"]; !ok {
		_, _ = hasher.Write(protowire.AppendBytes(nil, m.SingleBytes))

	}
	if m.NestedType != nil {
		if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.nested_type"]; !ok {
			switch t := m.NestedType.(type) {
			case *TestAllTypes_SingleNestedMessage:
				if t.SingleNestedMessage != nil {
					cerbos_hashpb_test_TestAllTypes_NestedMessage_hashpb_sum(t.SingleNestedMessage, hasher, ignore)
				}

			case *TestAllTypes_SingleNestedEnum:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.SingleNestedEnum)))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.standalone_enum"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.StandaloneEnum)))

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_int32"]; !ok {
		if len(m.RepeatedInt32) > 0 {
			for _, v := range m.RepeatedInt32 {
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(v)))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_int64"]; !ok {
		if len(m.RepeatedInt64) > 0 {
			for _, v := range m.RepeatedInt64 {
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(v)))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_uint32"]; !ok {
		if len(m.RepeatedUint32) > 0 {
			for _, v := range m.RepeatedUint32 {
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(v)))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_uint64"]; !ok {
		if len(m.RepeatedUint64) > 0 {
			for _, v := range m.RepeatedUint64 {
				_, _ = hasher.Write(protowire.AppendVarint(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_sint32"]; !ok {
		if len(m.RepeatedSint32) > 0 {
			for _, v := range m.RepeatedSint32 {
				_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeZigZag(int64(v))))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_sint64"]; !ok {
		if len(m.RepeatedSint64) > 0 {
			for _, v := range m.RepeatedSint64 {
				_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeZigZag(v)))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_fixed32"]; !ok {
		if len(m.RepeatedFixed32) > 0 {
			for _, v := range m.RepeatedFixed32 {
				_, _ = hasher.Write(protowire.AppendFixed32(nil, uint32(v)))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_fixed64"]; !ok {
		if len(m.RepeatedFixed64) > 0 {
			for _, v := range m.RepeatedFixed64 {
				_, _ = hasher.Write(protowire.AppendFixed64(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_sfixed32"]; !ok {
		if len(m.RepeatedSfixed32) > 0 {
			for _, v := range m.RepeatedSfixed32 {
				_, _ = hasher.Write(protowire.AppendFixed32(nil, uint32(v)))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_sfixed64"]; !ok {
		if len(m.RepeatedSfixed64) > 0 {
			for _, v := range m.RepeatedSfixed64 {
				_, _ = hasher.Write(protowire.AppendFixed64(nil, uint64(v)))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_float"]; !ok {
		if len(m.RepeatedFloat) > 0 {
			for _, v := range m.RepeatedFloat {
				_, _ = hasher.Write(protowire.AppendFixed32(nil, math.Float32bits(v)))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_double"]; !ok {
		if len(m.RepeatedDouble) > 0 {
			for _, v := range m.RepeatedDouble {
				_, _ = hasher.Write(protowire.AppendFixed64(nil, math.Float64bits(v)))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_bool"]; !ok {
		if len(m.RepeatedBool) > 0 {
			for _, v := range m.RepeatedBool {
				_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(v)))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_string"]; !ok {
		if len(m.RepeatedString) > 0 {
			for _, v := range m.RepeatedString {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_bytes"]; !ok {
		if len(m.RepeatedBytes) > 0 {
			for _, v := range m.RepeatedBytes {
				_, _ = hasher.Write(protowire.AppendBytes(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_nested_message"]; !ok {
		if len(m.RepeatedNestedMessage) > 0 {
			for _, v := range m.RepeatedNestedMessage {
				if v != nil {
					cerbos_hashpb_test_TestAllTypes_NestedMessage_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_nested_enum"]; !ok {
		if len(m.RepeatedNestedEnum) > 0 {
			for _, v := range m.RepeatedNestedEnum {
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(v)))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_string_piece"]; !ok {
		if len(m.RepeatedStringPiece) > 0 {
			for _, v := range m.RepeatedStringPiece {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_cord"]; !ok {
		if len(m.RepeatedCord) > 0 {
			for _, v := range m.RepeatedCord {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.repeated_lazy_message"]; !ok {
		if len(m.RepeatedLazyMessage) > 0 {
			for _, v := range m.RepeatedLazyMessage {
				if v != nil {
					cerbos_hashpb_test_TestAllTypes_NestedMessage_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.map_string_string"]; !ok {
		if len(m.MapStringString) > 0 {
			keys := make([]string, len(m.MapStringString))
			i := 0
			for k := range m.MapStringString {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendString(nil, k))

				_, _ = hasher.Write(protowire.AppendString(nil, m.MapStringString[k]))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.map_uint64_string"]; !ok {
		if len(m.MapUint64String) > 0 {
			keys := make([]uint64, len(m.MapUint64String))
			i := 0
			for k := range m.MapUint64String {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendVarint(nil, k))

				_, _ = hasher.Write(protowire.AppendString(nil, m.MapUint64String[k]))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.map_int32_string"]; !ok {
		if len(m.MapInt32String) > 0 {
			keys := make([]int32, len(m.MapInt32String))
			i := 0
			for k := range m.MapInt32String {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(k)))

				_, _ = hasher.Write(protowire.AppendString(nil, m.MapInt32String[k]))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.map_bool_string"]; !ok {
		if len(m.MapBoolString) > 0 {
			keys := make([]bool, len(m.MapBoolString))
			i := 0
			for k := range m.MapBoolString {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return !keys[i] && keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(k)))

				_, _ = hasher.Write(protowire.AppendString(nil, m.MapBoolString[k]))

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.map_int64_nested_type"]; !ok {
		if len(m.MapInt64NestedType) > 0 {
			keys := make([]int64, len(m.MapInt64NestedType))
			i := 0
			for k := range m.MapInt64NestedType {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(k)))

				if m.MapInt64NestedType[k] != nil {
					cerbos_hashpb_test_TestAllTypes_NestedMessage_hashpb_sum(m.MapInt64NestedType[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_any"]; !ok {
		if m.SingleAny != nil {
			google_protobuf_Any_hashpb_sum(m.SingleAny, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_duration"]; !ok {
		if m.SingleDuration != nil {
			google_protobuf_Duration_hashpb_sum(m.SingleDuration, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_timestamp"]; !ok {
		if m.SingleTimestamp != nil {
			google_protobuf_Timestamp_hashpb_sum(m.SingleTimestamp, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_struct"]; !ok {
		if m.SingleStruct != nil {
			google_protobuf_Struct_hashpb_sum(m.SingleStruct, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_value"]; !ok {
		if m.SingleValue != nil {
			google_protobuf_Value_hashpb_sum(m.SingleValue, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_int64_wrapper"]; !ok {
		if m.SingleInt64Wrapper != nil {
			google_protobuf_Int64Value_hashpb_sum(m.SingleInt64Wrapper, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_int32_wrapper"]; !ok {
		if m.SingleInt32Wrapper != nil {
			google_protobuf_Int32Value_hashpb_sum(m.SingleInt32Wrapper, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_double_wrapper"]; !ok {
		if m.SingleDoubleWrapper != nil {
			google_protobuf_DoubleValue_hashpb_sum(m.SingleDoubleWrapper, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_float_wrapper"]; !ok {
		if m.SingleFloatWrapper != nil {
			google_protobuf_FloatValue_hashpb_sum(m.SingleFloatWrapper, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_uint64_wrapper"]; !ok {
		if m.SingleUint64Wrapper != nil {
			google_protobuf_UInt64Value_hashpb_sum(m.SingleUint64Wrapper, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_uint32_wrapper"]; !ok {
		if m.SingleUint32Wrapper != nil {
			google_protobuf_UInt32Value_hashpb_sum(m.SingleUint32Wrapper, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_string_wrapper"]; !ok {
		if m.SingleStringWrapper != nil {
			google_protobuf_StringValue_hashpb_sum(m.SingleStringWrapper, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_bool_wrapper"]; !ok {
		if m.SingleBoolWrapper != nil {
			google_protobuf_BoolValue_hashpb_sum(m.SingleBoolWrapper, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.hashpb.test.TestAllTypes.single_bytes_wrapper"]; !ok {
		if m.SingleBytesWrapper != nil {
			google_protobuf_BytesValue_hashpb_sum(m.SingleBytesWrapper, hasher, ignore)
		}

	}
}

func google_protobuf_Any_hashpb_sum(m *anypb.Any, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Any.type_url"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.TypeUrl))

	}
	if _, ok := ignore["google.protobuf.Any.value"]; !ok {
		_, _ = hasher.Write(protowire.AppendBytes(nil, m.Value))

	}
}

func google_protobuf_BoolValue_hashpb_sum(m *wrapperspb.BoolValue, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.BoolValue.value"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.Value)))

	}
}

func google_protobuf_BytesValue_hashpb_sum(m *wrapperspb.BytesValue, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.BytesValue.value"]; !ok {
		_, _ = hasher.Write(protowire.AppendBytes(nil, m.Value))

	}
}

func google_protobuf_DoubleValue_hashpb_sum(m *wrapperspb.DoubleValue, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.DoubleValue.value"]; !ok {
		_, _ = hasher.Write(protowire.AppendFixed64(nil, math.Float64bits(m.Value)))

	}
}

func google_protobuf_Duration_hashpb_sum(m *durationpb.Duration, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Duration.seconds"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Seconds)))

	}
	if _, ok := ignore["google.protobuf.Duration.nanos"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Nanos)))

	}
}

func google_protobuf_FloatValue_hashpb_sum(m *wrapperspb.FloatValue, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.FloatValue.value"]; !ok {
		_, _ = hasher.Write(protowire.AppendFixed32(nil, math.Float32bits(m.Value)))

	}
}

func google_protobuf_Int32Value_hashpb_sum(m *wrapperspb.Int32Value, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Int32Value.value"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Value)))

	}
}

func google_protobuf_Int64Value_hashpb_sum(m *wrapperspb.Int64Value, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Int64Value.value"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Value)))

	}
}

func google_protobuf_ListValue_hashpb_sum(m *structpb.ListValue, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.ListValue.values"]; !ok {
		if len(m.Values) > 0 {
			for _, v := range m.Values {
				if v != nil {
					google_protobuf_Value_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func google_protobuf_StringValue_hashpb_sum(m *wrapperspb.StringValue, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.StringValue.value"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Value))

	}
}

func google_protobuf_Struct_hashpb_sum(m *structpb.Struct, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Struct.fields"]; !ok {
		if len(m.Fields) > 0 {
			keys := make([]string, len(m.Fields))
			i := 0
			for k := range m.Fields {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendString(nil, k))

				if m.Fields[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Fields[k], hasher, ignore)
				}

			}
		}
	}
}

func google_protobuf_Timestamp_hashpb_sum(m *timestamppb.Timestamp, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Timestamp.seconds"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Seconds)))

	}
	if _, ok := ignore["google.protobuf.Timestamp.nanos"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Nanos)))

	}
}

func google_protobuf_UInt32Value_hashpb_sum(m *wrapperspb.UInt32Value, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.UInt32Value.value"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Value)))

	}
}

func google_protobuf_UInt64Value_hashpb_sum(m *wrapperspb.UInt64Value, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.UInt64Value.value"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, m.Value))

	}
}

func google_protobuf_Value_hashpb_sum(m *structpb.Value, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Kind != nil {
		if _, ok := ignore["google.protobuf.Value.kind"]; !ok {
			switch t := m.Kind.(type) {
			case *structpb.Value_NullValue:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.NullValue)))

			case *structpb.Value_NumberValue:
				_, _ = hasher.Write(protowire.AppendFixed64(nil, math.Float64bits(t.NumberValue)))

			case *structpb.Value_StringValue:
				_, _ = hasher.Write(protowire.AppendString(nil, t.StringValue))

			case *structpb.Value_BoolValue:
				_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(t.BoolValue)))

			case *structpb.Value_StructValue:
				if t.StructValue != nil {
					google_protobuf_Struct_hashpb_sum(t.StructValue, hasher, ignore)
				}

			case *structpb.Value_ListValue:
				if t.ListValue != nil {
					google_protobuf_ListValue_hashpb_sum(t.ListValue, hasher, ignore)
				}

			}
		}
	}
}
