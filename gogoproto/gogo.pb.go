// Code generated by protoc-gen-gogo.
// source: gogo.proto
// DO NOT EDIT!

/*
Package gogoproto is a generated protocol buffer package.

It is generated from these files:
	gogo.proto

It has these top-level messages:
*/
package gogoproto

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// renamed import google/protobuf/descriptor to code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor
import google_protobuf "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

var E_GoprotoEnumPrefix = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62001,
	Name:          "gogoproto.goproto_enum_prefix",
	Tag:           "varint,62001,opt,name=goproto_enum_prefix",
}

var E_GoprotoEnumStringer = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62021,
	Name:          "gogoproto.goproto_enum_stringer",
	Tag:           "varint,62021,opt,name=goproto_enum_stringer",
}

var E_EnumStringer = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62022,
	Name:          "gogoproto.enum_stringer",
	Tag:           "varint,62022,opt,name=enum_stringer",
}

var E_GoprotoGettersAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63001,
	Name:          "gogoproto.goproto_getters_all",
	Tag:           "varint,63001,opt,name=goproto_getters_all",
}

var E_GoprotoEnumPrefixAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63002,
	Name:          "gogoproto.goproto_enum_prefix_all",
	Tag:           "varint,63002,opt,name=goproto_enum_prefix_all",
}

var E_GoprotoStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63003,
	Name:          "gogoproto.goproto_stringer_all",
	Tag:           "varint,63003,opt,name=goproto_stringer_all",
}

var E_VerboseEqualAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63004,
	Name:          "gogoproto.verbose_equal_all",
	Tag:           "varint,63004,opt,name=verbose_equal_all",
}

var E_FaceAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63005,
	Name:          "gogoproto.face_all",
	Tag:           "varint,63005,opt,name=face_all",
}

var E_GostringAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63006,
	Name:          "gogoproto.gostring_all",
	Tag:           "varint,63006,opt,name=gostring_all",
}

var E_PopulateAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63007,
	Name:          "gogoproto.populate_all",
	Tag:           "varint,63007,opt,name=populate_all",
}

var E_StringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63008,
	Name:          "gogoproto.stringer_all",
	Tag:           "varint,63008,opt,name=stringer_all",
}

var E_OnlyoneAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63009,
	Name:          "gogoproto.onlyone_all",
	Tag:           "varint,63009,opt,name=onlyone_all",
}

var E_EqualAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63013,
	Name:          "gogoproto.equal_all",
	Tag:           "varint,63013,opt,name=equal_all",
}

var E_DescriptionAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63014,
	Name:          "gogoproto.description_all",
	Tag:           "varint,63014,opt,name=description_all",
}

var E_TestgenAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63015,
	Name:          "gogoproto.testgen_all",
	Tag:           "varint,63015,opt,name=testgen_all",
}

var E_BenchgenAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63016,
	Name:          "gogoproto.benchgen_all",
	Tag:           "varint,63016,opt,name=benchgen_all",
}

var E_MarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63017,
	Name:          "gogoproto.marshaler_all",
	Tag:           "varint,63017,opt,name=marshaler_all",
}

var E_UnmarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63018,
	Name:          "gogoproto.unmarshaler_all",
	Tag:           "varint,63018,opt,name=unmarshaler_all",
}

var E_BuffertoAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63019,
	Name:          "gogoproto.bufferto_all",
	Tag:           "varint,63019,opt,name=bufferto_all",
}

var E_SizerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63020,
	Name:          "gogoproto.sizer_all",
	Tag:           "varint,63020,opt,name=sizer_all",
}

var E_GoprotoEnumStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63021,
	Name:          "gogoproto.goproto_enum_stringer_all",
	Tag:           "varint,63021,opt,name=goproto_enum_stringer_all",
}

var E_EnumStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63022,
	Name:          "gogoproto.enum_stringer_all",
	Tag:           "varint,63022,opt,name=enum_stringer_all",
}

var E_UnsafeMarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63023,
	Name:          "gogoproto.unsafe_marshaler_all",
	Tag:           "varint,63023,opt,name=unsafe_marshaler_all",
}

var E_UnsafeUnmarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63024,
	Name:          "gogoproto.unsafe_unmarshaler_all",
	Tag:           "varint,63024,opt,name=unsafe_unmarshaler_all",
}

var E_GoprotoGetters = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64001,
	Name:          "gogoproto.goproto_getters",
	Tag:           "varint,64001,opt,name=goproto_getters",
}

var E_GoprotoStringer = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64003,
	Name:          "gogoproto.goproto_stringer",
	Tag:           "varint,64003,opt,name=goproto_stringer",
}

var E_VerboseEqual = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64004,
	Name:          "gogoproto.verbose_equal",
	Tag:           "varint,64004,opt,name=verbose_equal",
}

var E_Face = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64005,
	Name:          "gogoproto.face",
	Tag:           "varint,64005,opt,name=face",
}

var E_Gostring = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64006,
	Name:          "gogoproto.gostring",
	Tag:           "varint,64006,opt,name=gostring",
}

var E_Populate = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64007,
	Name:          "gogoproto.populate",
	Tag:           "varint,64007,opt,name=populate",
}

var E_Stringer = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         67008,
	Name:          "gogoproto.stringer",
	Tag:           "varint,67008,opt,name=stringer",
}

var E_Onlyone = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64009,
	Name:          "gogoproto.onlyone",
	Tag:           "varint,64009,opt,name=onlyone",
}

var E_Equal = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64013,
	Name:          "gogoproto.equal",
	Tag:           "varint,64013,opt,name=equal",
}

var E_Description = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64014,
	Name:          "gogoproto.description",
	Tag:           "varint,64014,opt,name=description",
}

var E_Testgen = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64015,
	Name:          "gogoproto.testgen",
	Tag:           "varint,64015,opt,name=testgen",
}

var E_Benchgen = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64016,
	Name:          "gogoproto.benchgen",
	Tag:           "varint,64016,opt,name=benchgen",
}

var E_Marshaler = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64017,
	Name:          "gogoproto.marshaler",
	Tag:           "varint,64017,opt,name=marshaler",
}

var E_Unmarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64018,
	Name:          "gogoproto.unmarshaler",
	Tag:           "varint,64018,opt,name=unmarshaler",
}

var E_Bufferto = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64019,
	Name:          "gogoproto.bufferto",
	Tag:           "varint,64019,opt,name=bufferto",
}

var E_Sizer = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64020,
	Name:          "gogoproto.sizer",
	Tag:           "varint,64020,opt,name=sizer",
}

var E_UnsafeMarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64023,
	Name:          "gogoproto.unsafe_marshaler",
	Tag:           "varint,64023,opt,name=unsafe_marshaler",
}

var E_UnsafeUnmarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64024,
	Name:          "gogoproto.unsafe_unmarshaler",
	Tag:           "varint,64024,opt,name=unsafe_unmarshaler",
}

var E_Nullable = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65001,
	Name:          "gogoproto.nullable",
	Tag:           "varint,65001,opt,name=nullable",
}

var E_Embed = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65002,
	Name:          "gogoproto.embed",
	Tag:           "varint,65002,opt,name=embed",
}

var E_Customtype = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65003,
	Name:          "gogoproto.customtype",
	Tag:           "bytes,65003,opt,name=customtype",
}

var E_Customname = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65004,
	Name:          "gogoproto.customname",
	Tag:           "bytes,65004,opt,name=customname",
}

func init() {
	proto.RegisterExtension(E_GoprotoEnumPrefix)
	proto.RegisterExtension(E_GoprotoEnumStringer)
	proto.RegisterExtension(E_EnumStringer)
	proto.RegisterExtension(E_GoprotoGettersAll)
	proto.RegisterExtension(E_GoprotoEnumPrefixAll)
	proto.RegisterExtension(E_GoprotoStringerAll)
	proto.RegisterExtension(E_VerboseEqualAll)
	proto.RegisterExtension(E_FaceAll)
	proto.RegisterExtension(E_GostringAll)
	proto.RegisterExtension(E_PopulateAll)
	proto.RegisterExtension(E_StringerAll)
	proto.RegisterExtension(E_OnlyoneAll)
	proto.RegisterExtension(E_EqualAll)
	proto.RegisterExtension(E_DescriptionAll)
	proto.RegisterExtension(E_TestgenAll)
	proto.RegisterExtension(E_BenchgenAll)
	proto.RegisterExtension(E_MarshalerAll)
	proto.RegisterExtension(E_UnmarshalerAll)
	proto.RegisterExtension(E_BuffertoAll)
	proto.RegisterExtension(E_SizerAll)
	proto.RegisterExtension(E_GoprotoEnumStringerAll)
	proto.RegisterExtension(E_EnumStringerAll)
	proto.RegisterExtension(E_UnsafeMarshalerAll)
	proto.RegisterExtension(E_UnsafeUnmarshalerAll)
	proto.RegisterExtension(E_GoprotoGetters)
	proto.RegisterExtension(E_GoprotoStringer)
	proto.RegisterExtension(E_VerboseEqual)
	proto.RegisterExtension(E_Face)
	proto.RegisterExtension(E_Gostring)
	proto.RegisterExtension(E_Populate)
	proto.RegisterExtension(E_Stringer)
	proto.RegisterExtension(E_Onlyone)
	proto.RegisterExtension(E_Equal)
	proto.RegisterExtension(E_Description)
	proto.RegisterExtension(E_Testgen)
	proto.RegisterExtension(E_Benchgen)
	proto.RegisterExtension(E_Marshaler)
	proto.RegisterExtension(E_Unmarshaler)
	proto.RegisterExtension(E_Bufferto)
	proto.RegisterExtension(E_Sizer)
	proto.RegisterExtension(E_UnsafeMarshaler)
	proto.RegisterExtension(E_UnsafeUnmarshaler)
	proto.RegisterExtension(E_Nullable)
	proto.RegisterExtension(E_Embed)
	proto.RegisterExtension(E_Customtype)
	proto.RegisterExtension(E_Customname)
}
