// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kythe/proto/xref_serving.proto

package xref_serving_go_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common_go_proto "kythe.io/kythe/proto/common_go_proto"
import schema_go_proto "kythe.io/kythe/proto/schema_go_proto"
import serving_go_proto "kythe.io/kythe/proto/serving_go_proto"
import storage_go_proto "kythe.io/kythe/proto/storage_go_proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FileDecorations_TargetOverride_Kind int32

const (
	FileDecorations_TargetOverride_OVERRIDES FileDecorations_TargetOverride_Kind = 0
	FileDecorations_TargetOverride_EXTENDS   FileDecorations_TargetOverride_Kind = 1
)

var FileDecorations_TargetOverride_Kind_name = map[int32]string{
	0: "OVERRIDES",
	1: "EXTENDS",
}
var FileDecorations_TargetOverride_Kind_value = map[string]int32{
	"OVERRIDES": 0,
	"EXTENDS":   1,
}

func (x FileDecorations_TargetOverride_Kind) String() string {
	return proto.EnumName(FileDecorations_TargetOverride_Kind_name, int32(x))
}
func (FileDecorations_TargetOverride_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_xref_serving_93298b8ea3ebd096, []int{0, 3, 0}
}

type FileDecorations struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileDecorations) Reset()         { *m = FileDecorations{} }
func (m *FileDecorations) String() string { return proto.CompactTextString(m) }
func (*FileDecorations) ProtoMessage()    {}
func (*FileDecorations) Descriptor() ([]byte, []int) {
	return fileDescriptor_xref_serving_93298b8ea3ebd096, []int{0}
}
func (m *FileDecorations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileDecorations.Unmarshal(m, b)
}
func (m *FileDecorations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileDecorations.Marshal(b, m, deterministic)
}
func (dst *FileDecorations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileDecorations.Merge(dst, src)
}
func (m *FileDecorations) XXX_Size() int {
	return xxx_messageInfo_FileDecorations.Size(m)
}
func (m *FileDecorations) XXX_DiscardUnknown() {
	xxx_messageInfo_FileDecorations.DiscardUnknown(m)
}

var xxx_messageInfo_FileDecorations proto.InternalMessageInfo

type FileDecorations_Index struct {
	TextEncoding         string   `protobuf:"bytes,1,opt,name=text_encoding,json=textEncoding" json:"text_encoding,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileDecorations_Index) Reset()         { *m = FileDecorations_Index{} }
func (m *FileDecorations_Index) String() string { return proto.CompactTextString(m) }
func (*FileDecorations_Index) ProtoMessage()    {}
func (*FileDecorations_Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_xref_serving_93298b8ea3ebd096, []int{0, 0}
}
func (m *FileDecorations_Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileDecorations_Index.Unmarshal(m, b)
}
func (m *FileDecorations_Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileDecorations_Index.Marshal(b, m, deterministic)
}
func (dst *FileDecorations_Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileDecorations_Index.Merge(dst, src)
}
func (m *FileDecorations_Index) XXX_Size() int {
	return xxx_messageInfo_FileDecorations_Index.Size(m)
}
func (m *FileDecorations_Index) XXX_DiscardUnknown() {
	xxx_messageInfo_FileDecorations_Index.DiscardUnknown(m)
}

var xxx_messageInfo_FileDecorations_Index proto.InternalMessageInfo

func (m *FileDecorations_Index) GetTextEncoding() string {
	if m != nil {
		return m.TextEncoding
	}
	return ""
}

type FileDecorations_Text struct {
	StartOffest          int32    `protobuf:"varint,1,opt,name=start_offest,json=startOffest" json:"start_offest,omitempty"`
	EndOffset            int32    `protobuf:"varint,2,opt,name=end_offset,json=endOffset" json:"end_offset,omitempty"`
	Text                 []byte   `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileDecorations_Text) Reset()         { *m = FileDecorations_Text{} }
func (m *FileDecorations_Text) String() string { return proto.CompactTextString(m) }
func (*FileDecorations_Text) ProtoMessage()    {}
func (*FileDecorations_Text) Descriptor() ([]byte, []int) {
	return fileDescriptor_xref_serving_93298b8ea3ebd096, []int{0, 1}
}
func (m *FileDecorations_Text) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileDecorations_Text.Unmarshal(m, b)
}
func (m *FileDecorations_Text) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileDecorations_Text.Marshal(b, m, deterministic)
}
func (dst *FileDecorations_Text) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileDecorations_Text.Merge(dst, src)
}
func (m *FileDecorations_Text) XXX_Size() int {
	return xxx_messageInfo_FileDecorations_Text.Size(m)
}
func (m *FileDecorations_Text) XXX_DiscardUnknown() {
	xxx_messageInfo_FileDecorations_Text.DiscardUnknown(m)
}

var xxx_messageInfo_FileDecorations_Text proto.InternalMessageInfo

func (m *FileDecorations_Text) GetStartOffest() int32 {
	if m != nil {
		return m.StartOffest
	}
	return 0
}

func (m *FileDecorations_Text) GetEndOffset() int32 {
	if m != nil {
		return m.EndOffset
	}
	return 0
}

func (m *FileDecorations_Text) GetText() []byte {
	if m != nil {
		return m.Text
	}
	return nil
}

type FileDecorations_Target struct {
	StartOffest int32 `protobuf:"varint,1,opt,name=start_offest,json=startOffest" json:"start_offest,omitempty"`
	EndOffset   int32 `protobuf:"varint,2,opt,name=end_offset,json=endOffset" json:"end_offset,omitempty"`
	// Types that are valid to be assigned to Kind:
	//	*FileDecorations_Target_KytheKind
	//	*FileDecorations_Target_GenericKind
	Kind                 isFileDecorations_Target_Kind `protobuf_oneof:"kind"`
	Target               *storage_go_proto.VName       `protobuf:"bytes,5,opt,name=target" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *FileDecorations_Target) Reset()         { *m = FileDecorations_Target{} }
func (m *FileDecorations_Target) String() string { return proto.CompactTextString(m) }
func (*FileDecorations_Target) ProtoMessage()    {}
func (*FileDecorations_Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_xref_serving_93298b8ea3ebd096, []int{0, 2}
}
func (m *FileDecorations_Target) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileDecorations_Target.Unmarshal(m, b)
}
func (m *FileDecorations_Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileDecorations_Target.Marshal(b, m, deterministic)
}
func (dst *FileDecorations_Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileDecorations_Target.Merge(dst, src)
}
func (m *FileDecorations_Target) XXX_Size() int {
	return xxx_messageInfo_FileDecorations_Target.Size(m)
}
func (m *FileDecorations_Target) XXX_DiscardUnknown() {
	xxx_messageInfo_FileDecorations_Target.DiscardUnknown(m)
}

var xxx_messageInfo_FileDecorations_Target proto.InternalMessageInfo

type isFileDecorations_Target_Kind interface {
	isFileDecorations_Target_Kind()
}

type FileDecorations_Target_KytheKind struct {
	KytheKind schema_go_proto.EdgeKind `protobuf:"varint,3,opt,name=kythe_kind,json=kytheKind,enum=kythe.proto.schema.EdgeKind,oneof"`
}
type FileDecorations_Target_GenericKind struct {
	GenericKind string `protobuf:"bytes,4,opt,name=generic_kind,json=genericKind,oneof"`
}

func (*FileDecorations_Target_KytheKind) isFileDecorations_Target_Kind()   {}
func (*FileDecorations_Target_GenericKind) isFileDecorations_Target_Kind() {}

func (m *FileDecorations_Target) GetKind() isFileDecorations_Target_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *FileDecorations_Target) GetStartOffest() int32 {
	if m != nil {
		return m.StartOffest
	}
	return 0
}

func (m *FileDecorations_Target) GetEndOffset() int32 {
	if m != nil {
		return m.EndOffset
	}
	return 0
}

func (m *FileDecorations_Target) GetKytheKind() schema_go_proto.EdgeKind {
	if x, ok := m.GetKind().(*FileDecorations_Target_KytheKind); ok {
		return x.KytheKind
	}
	return schema_go_proto.EdgeKind_UNKNOWN_EDGE_KIND
}

func (m *FileDecorations_Target) GetGenericKind() string {
	if x, ok := m.GetKind().(*FileDecorations_Target_GenericKind); ok {
		return x.GenericKind
	}
	return ""
}

func (m *FileDecorations_Target) GetTarget() *storage_go_proto.VName {
	if m != nil {
		return m.Target
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FileDecorations_Target) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FileDecorations_Target_OneofMarshaler, _FileDecorations_Target_OneofUnmarshaler, _FileDecorations_Target_OneofSizer, []interface{}{
		(*FileDecorations_Target_KytheKind)(nil),
		(*FileDecorations_Target_GenericKind)(nil),
	}
}

func _FileDecorations_Target_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FileDecorations_Target)
	// kind
	switch x := m.Kind.(type) {
	case *FileDecorations_Target_KytheKind:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.KytheKind))
	case *FileDecorations_Target_GenericKind:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.GenericKind)
	case nil:
	default:
		return fmt.Errorf("FileDecorations_Target.Kind has unexpected type %T", x)
	}
	return nil
}

func _FileDecorations_Target_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FileDecorations_Target)
	switch tag {
	case 3: // kind.kythe_kind
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Kind = &FileDecorations_Target_KytheKind{schema_go_proto.EdgeKind(x)}
		return true, err
	case 4: // kind.generic_kind
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Kind = &FileDecorations_Target_GenericKind{x}
		return true, err
	default:
		return false, nil
	}
}

func _FileDecorations_Target_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FileDecorations_Target)
	// kind
	switch x := m.Kind.(type) {
	case *FileDecorations_Target_KytheKind:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.KytheKind))
	case *FileDecorations_Target_GenericKind:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.GenericKind)))
		n += len(x.GenericKind)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type FileDecorations_TargetOverride struct {
	Overridden           *storage_go_proto.VName             `protobuf:"bytes,1,opt,name=overridden" json:"overridden,omitempty"`
	Kind                 FileDecorations_TargetOverride_Kind `protobuf:"varint,2,opt,name=kind,enum=kythe.proto.serving.xrefs.FileDecorations_TargetOverride_Kind" json:"kind,omitempty"`
	Overriding           *storage_go_proto.VName             `protobuf:"bytes,3,opt,name=overriding" json:"overriding,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *FileDecorations_TargetOverride) Reset()         { *m = FileDecorations_TargetOverride{} }
func (m *FileDecorations_TargetOverride) String() string { return proto.CompactTextString(m) }
func (*FileDecorations_TargetOverride) ProtoMessage()    {}
func (*FileDecorations_TargetOverride) Descriptor() ([]byte, []int) {
	return fileDescriptor_xref_serving_93298b8ea3ebd096, []int{0, 3}
}
func (m *FileDecorations_TargetOverride) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileDecorations_TargetOverride.Unmarshal(m, b)
}
func (m *FileDecorations_TargetOverride) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileDecorations_TargetOverride.Marshal(b, m, deterministic)
}
func (dst *FileDecorations_TargetOverride) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileDecorations_TargetOverride.Merge(dst, src)
}
func (m *FileDecorations_TargetOverride) XXX_Size() int {
	return xxx_messageInfo_FileDecorations_TargetOverride.Size(m)
}
func (m *FileDecorations_TargetOverride) XXX_DiscardUnknown() {
	xxx_messageInfo_FileDecorations_TargetOverride.DiscardUnknown(m)
}

var xxx_messageInfo_FileDecorations_TargetOverride proto.InternalMessageInfo

func (m *FileDecorations_TargetOverride) GetOverridden() *storage_go_proto.VName {
	if m != nil {
		return m.Overridden
	}
	return nil
}

func (m *FileDecorations_TargetOverride) GetKind() FileDecorations_TargetOverride_Kind {
	if m != nil {
		return m.Kind
	}
	return FileDecorations_TargetOverride_OVERRIDES
}

func (m *FileDecorations_TargetOverride) GetOverriding() *storage_go_proto.VName {
	if m != nil {
		return m.Overriding
	}
	return nil
}

type FileDecorations_TargetNode struct {
	Node                 *schema_go_proto.Node `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FileDecorations_TargetNode) Reset()         { *m = FileDecorations_TargetNode{} }
func (m *FileDecorations_TargetNode) String() string { return proto.CompactTextString(m) }
func (*FileDecorations_TargetNode) ProtoMessage()    {}
func (*FileDecorations_TargetNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_xref_serving_93298b8ea3ebd096, []int{0, 4}
}
func (m *FileDecorations_TargetNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileDecorations_TargetNode.Unmarshal(m, b)
}
func (m *FileDecorations_TargetNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileDecorations_TargetNode.Marshal(b, m, deterministic)
}
func (dst *FileDecorations_TargetNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileDecorations_TargetNode.Merge(dst, src)
}
func (m *FileDecorations_TargetNode) XXX_Size() int {
	return xxx_messageInfo_FileDecorations_TargetNode.Size(m)
}
func (m *FileDecorations_TargetNode) XXX_DiscardUnknown() {
	xxx_messageInfo_FileDecorations_TargetNode.DiscardUnknown(m)
}

var xxx_messageInfo_FileDecorations_TargetNode proto.InternalMessageInfo

func (m *FileDecorations_TargetNode) GetNode() *schema_go_proto.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type FileDecorations_TargetDefinition struct {
	Definition           *storage_go_proto.VName `protobuf:"bytes,1,opt,name=definition" json:"definition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FileDecorations_TargetDefinition) Reset()         { *m = FileDecorations_TargetDefinition{} }
func (m *FileDecorations_TargetDefinition) String() string { return proto.CompactTextString(m) }
func (*FileDecorations_TargetDefinition) ProtoMessage()    {}
func (*FileDecorations_TargetDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_xref_serving_93298b8ea3ebd096, []int{0, 5}
}
func (m *FileDecorations_TargetDefinition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileDecorations_TargetDefinition.Unmarshal(m, b)
}
func (m *FileDecorations_TargetDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileDecorations_TargetDefinition.Marshal(b, m, deterministic)
}
func (dst *FileDecorations_TargetDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileDecorations_TargetDefinition.Merge(dst, src)
}
func (m *FileDecorations_TargetDefinition) XXX_Size() int {
	return xxx_messageInfo_FileDecorations_TargetDefinition.Size(m)
}
func (m *FileDecorations_TargetDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_FileDecorations_TargetDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_FileDecorations_TargetDefinition proto.InternalMessageInfo

func (m *FileDecorations_TargetDefinition) GetDefinition() *storage_go_proto.VName {
	if m != nil {
		return m.Definition
	}
	return nil
}

type FileDecorations_DefinitionLocation struct {
	Location             *serving_go_proto.ExpandedAnchor `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *FileDecorations_DefinitionLocation) Reset()         { *m = FileDecorations_DefinitionLocation{} }
func (m *FileDecorations_DefinitionLocation) String() string { return proto.CompactTextString(m) }
func (*FileDecorations_DefinitionLocation) ProtoMessage()    {}
func (*FileDecorations_DefinitionLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_xref_serving_93298b8ea3ebd096, []int{0, 6}
}
func (m *FileDecorations_DefinitionLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileDecorations_DefinitionLocation.Unmarshal(m, b)
}
func (m *FileDecorations_DefinitionLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileDecorations_DefinitionLocation.Marshal(b, m, deterministic)
}
func (dst *FileDecorations_DefinitionLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileDecorations_DefinitionLocation.Merge(dst, src)
}
func (m *FileDecorations_DefinitionLocation) XXX_Size() int {
	return xxx_messageInfo_FileDecorations_DefinitionLocation.Size(m)
}
func (m *FileDecorations_DefinitionLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_FileDecorations_DefinitionLocation.DiscardUnknown(m)
}

var xxx_messageInfo_FileDecorations_DefinitionLocation proto.InternalMessageInfo

func (m *FileDecorations_DefinitionLocation) GetLocation() *serving_go_proto.ExpandedAnchor {
	if m != nil {
		return m.Location
	}
	return nil
}

type FileDecorations_Override struct {
	MarkedSource         *common_go_proto.MarkedSource `protobuf:"bytes,1,opt,name=marked_source,json=markedSource" json:"marked_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *FileDecorations_Override) Reset()         { *m = FileDecorations_Override{} }
func (m *FileDecorations_Override) String() string { return proto.CompactTextString(m) }
func (*FileDecorations_Override) ProtoMessage()    {}
func (*FileDecorations_Override) Descriptor() ([]byte, []int) {
	return fileDescriptor_xref_serving_93298b8ea3ebd096, []int{0, 7}
}
func (m *FileDecorations_Override) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileDecorations_Override.Unmarshal(m, b)
}
func (m *FileDecorations_Override) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileDecorations_Override.Marshal(b, m, deterministic)
}
func (dst *FileDecorations_Override) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileDecorations_Override.Merge(dst, src)
}
func (m *FileDecorations_Override) XXX_Size() int {
	return xxx_messageInfo_FileDecorations_Override.Size(m)
}
func (m *FileDecorations_Override) XXX_DiscardUnknown() {
	xxx_messageInfo_FileDecorations_Override.DiscardUnknown(m)
}

var xxx_messageInfo_FileDecorations_Override proto.InternalMessageInfo

func (m *FileDecorations_Override) GetMarkedSource() *common_go_proto.MarkedSource {
	if m != nil {
		return m.MarkedSource
	}
	return nil
}

type FileDecorations_Diagnostic struct {
	Diagnostic           *common_go_proto.Diagnostic `protobuf:"bytes,1,opt,name=diagnostic" json:"diagnostic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *FileDecorations_Diagnostic) Reset()         { *m = FileDecorations_Diagnostic{} }
func (m *FileDecorations_Diagnostic) String() string { return proto.CompactTextString(m) }
func (*FileDecorations_Diagnostic) ProtoMessage()    {}
func (*FileDecorations_Diagnostic) Descriptor() ([]byte, []int) {
	return fileDescriptor_xref_serving_93298b8ea3ebd096, []int{0, 8}
}
func (m *FileDecorations_Diagnostic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileDecorations_Diagnostic.Unmarshal(m, b)
}
func (m *FileDecorations_Diagnostic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileDecorations_Diagnostic.Marshal(b, m, deterministic)
}
func (dst *FileDecorations_Diagnostic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileDecorations_Diagnostic.Merge(dst, src)
}
func (m *FileDecorations_Diagnostic) XXX_Size() int {
	return xxx_messageInfo_FileDecorations_Diagnostic.Size(m)
}
func (m *FileDecorations_Diagnostic) XXX_DiscardUnknown() {
	xxx_messageInfo_FileDecorations_Diagnostic.DiscardUnknown(m)
}

var xxx_messageInfo_FileDecorations_Diagnostic proto.InternalMessageInfo

func (m *FileDecorations_Diagnostic) GetDiagnostic() *common_go_proto.Diagnostic {
	if m != nil {
		return m.Diagnostic
	}
	return nil
}

type CrossReferences struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CrossReferences) Reset()         { *m = CrossReferences{} }
func (m *CrossReferences) String() string { return proto.CompactTextString(m) }
func (*CrossReferences) ProtoMessage()    {}
func (*CrossReferences) Descriptor() ([]byte, []int) {
	return fileDescriptor_xref_serving_93298b8ea3ebd096, []int{1}
}
func (m *CrossReferences) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrossReferences.Unmarshal(m, b)
}
func (m *CrossReferences) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrossReferences.Marshal(b, m, deterministic)
}
func (dst *CrossReferences) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossReferences.Merge(dst, src)
}
func (m *CrossReferences) XXX_Size() int {
	return xxx_messageInfo_CrossReferences.Size(m)
}
func (m *CrossReferences) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossReferences.DiscardUnknown(m)
}

var xxx_messageInfo_CrossReferences proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FileDecorations)(nil), "kythe.proto.serving.xrefs.FileDecorations")
	proto.RegisterType((*FileDecorations_Index)(nil), "kythe.proto.serving.xrefs.FileDecorations.Index")
	proto.RegisterType((*FileDecorations_Text)(nil), "kythe.proto.serving.xrefs.FileDecorations.Text")
	proto.RegisterType((*FileDecorations_Target)(nil), "kythe.proto.serving.xrefs.FileDecorations.Target")
	proto.RegisterType((*FileDecorations_TargetOverride)(nil), "kythe.proto.serving.xrefs.FileDecorations.TargetOverride")
	proto.RegisterType((*FileDecorations_TargetNode)(nil), "kythe.proto.serving.xrefs.FileDecorations.TargetNode")
	proto.RegisterType((*FileDecorations_TargetDefinition)(nil), "kythe.proto.serving.xrefs.FileDecorations.TargetDefinition")
	proto.RegisterType((*FileDecorations_DefinitionLocation)(nil), "kythe.proto.serving.xrefs.FileDecorations.DefinitionLocation")
	proto.RegisterType((*FileDecorations_Override)(nil), "kythe.proto.serving.xrefs.FileDecorations.Override")
	proto.RegisterType((*FileDecorations_Diagnostic)(nil), "kythe.proto.serving.xrefs.FileDecorations.Diagnostic")
	proto.RegisterType((*CrossReferences)(nil), "kythe.proto.serving.xrefs.CrossReferences")
	proto.RegisterEnum("kythe.proto.serving.xrefs.FileDecorations_TargetOverride_Kind", FileDecorations_TargetOverride_Kind_name, FileDecorations_TargetOverride_Kind_value)
}

func init() {
	proto.RegisterFile("kythe/proto/xref_serving.proto", fileDescriptor_xref_serving_93298b8ea3ebd096)
}

var fileDescriptor_xref_serving_93298b8ea3ebd096 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xad, 0x7f, 0x75, 0xd3, 0x66, 0x92, 0xfe, 0xf9, 0xed, 0xc9, 0xb5, 0xa0, 0xa4, 0xe9, 0xa5,
	0x42, 0x95, 0x2b, 0x85, 0x1b, 0x12, 0x45, 0x94, 0xb8, 0x6a, 0x45, 0x69, 0xc4, 0xb6, 0x54, 0x1c,
	0x90, 0x22, 0xe3, 0x9d, 0xb8, 0xab, 0xc4, 0xbb, 0xd5, 0xee, 0x12, 0x85, 0xaf, 0xcb, 0x27, 0x80,
	0x6f, 0x80, 0x76, 0xed, 0x26, 0x76, 0x15, 0xe0, 0xc2, 0x29, 0x3b, 0x33, 0x6f, 0xde, 0xcc, 0x7b,
	0x13, 0xc3, 0xde, 0xf8, 0x9b, 0xb9, 0xc3, 0xe3, 0x7b, 0x25, 0x8d, 0x3c, 0x9e, 0x29, 0x1c, 0x0d,
	0x35, 0xaa, 0x29, 0x17, 0x59, 0xe4, 0x52, 0x64, 0xd7, 0xd5, 0x8b, 0x20, 0x7a, 0x28, 0x59, 0x9c,
	0x0e, 0x83, 0x6a, 0x6b, 0x2a, 0xf3, 0x5c, 0x8a, 0x02, 0x57, 0xaf, 0xe8, 0xf4, 0x0e, 0xf3, 0xa4,
	0xac, 0xec, 0xd6, 0x2a, 0xd5, 0x49, 0x8f, 0x4a, 0x46, 0xaa, 0x24, 0x2b, 0xe7, 0x76, 0x7f, 0xae,
	0xc3, 0xf6, 0x19, 0x9f, 0x60, 0x1f, 0x53, 0xa9, 0x12, 0xc3, 0xa5, 0xd0, 0xe1, 0x11, 0xac, 0x5d,
	0x08, 0x86, 0x33, 0x72, 0x00, 0x9b, 0x06, 0x67, 0x66, 0x88, 0x22, 0x95, 0x8c, 0x8b, 0x2c, 0xf0,
	0x3a, 0xde, 0x61, 0x93, 0xb6, 0x6d, 0x32, 0x2e, 0x73, 0xe1, 0x67, 0xf0, 0x6f, 0x70, 0x66, 0xc8,
	0x3e, 0xb4, 0xb5, 0x49, 0x94, 0x19, 0xca, 0xd1, 0x08, 0xb5, 0x71, 0xd8, 0x35, 0xda, 0x72, 0xb9,
	0x81, 0x4b, 0x91, 0xa7, 0x00, 0x28, 0x98, 0x05, 0x68, 0x34, 0xc1, 0x7f, 0x0e, 0xd0, 0x44, 0xc1,
	0x06, 0x2e, 0x41, 0x08, 0xf8, 0x96, 0x39, 0x58, 0xed, 0x78, 0x87, 0x6d, 0xea, 0xde, 0xe1, 0x77,
	0x0f, 0x1a, 0x37, 0x89, 0xca, 0xf0, 0x5f, 0x0c, 0x78, 0x05, 0xe0, 0x9c, 0x18, 0x8e, 0xb9, 0x60,
	0x6e, 0xcc, 0x56, 0xef, 0x49, 0x54, 0x3b, 0x43, 0xe1, 0x68, 0xcc, 0x32, 0x7c, 0xc7, 0x05, 0x3b,
	0x5f, 0xa1, 0x4d, 0x57, 0xb6, 0x01, 0x39, 0x80, 0x76, 0x86, 0x02, 0x15, 0x4f, 0x0b, 0x02, 0xdf,
	0xba, 0x71, 0xbe, 0x42, 0x5b, 0x65, 0xd6, 0x81, 0x9e, 0x43, 0xc3, 0xb8, 0x7d, 0x83, 0xb5, 0x8e,
	0x77, 0xd8, 0xea, 0x91, 0x1a, 0xff, 0xed, 0x55, 0x92, 0x23, 0x2d, 0x11, 0xa7, 0x0d, 0xf0, 0x2d,
	0x51, 0xf8, 0xc3, 0x83, 0xad, 0x42, 0xe4, 0x60, 0x8a, 0x4a, 0x71, 0x86, 0xa4, 0x07, 0x20, 0x8b,
	0x37, 0x43, 0xe1, 0xa4, 0x2e, 0xa7, 0xaa, 0xa0, 0x08, 0x2d, 0xe8, 0x9c, 0xee, 0xad, 0xde, 0x49,
	0xf4, 0xdb, 0xff, 0x57, 0xf4, 0xe8, 0xe2, 0x51, 0x7d, 0x78, 0x64, 0x85, 0x50, 0xc7, 0x55, 0xd9,
	0xc3, 0xde, 0x7f, 0xf5, 0xaf, 0x7b, 0x70, 0x91, 0x75, 0xbb, 0xe0, 0x3b, 0x2b, 0x36, 0xa1, 0x39,
	0xb8, 0x8d, 0x29, 0xbd, 0xe8, 0xc7, 0xd7, 0x3b, 0x2b, 0xa4, 0x05, 0xeb, 0xf1, 0xa7, 0x9b, 0xf8,
	0xaa, 0x7f, 0xbd, 0xe3, 0x85, 0x2f, 0x01, 0x8a, 0xa1, 0x57, 0x92, 0x21, 0x39, 0x02, 0x5f, 0x48,
	0x86, 0xa5, 0xce, 0x60, 0xd9, 0x49, 0x2c, 0x8e, 0x3a, 0x54, 0x78, 0x06, 0x3b, 0x45, 0x6f, 0x1f,
	0x47, 0x5c, 0x70, 0x2b, 0xc1, 0xee, 0xc9, 0xe6, 0xd1, 0x9f, 0xfc, 0x5a, 0xa0, 0xc2, 0x8f, 0x40,
	0x16, 0x0c, 0x97, 0x32, 0x75, 0x66, 0x90, 0xd7, 0xb0, 0x31, 0x29, 0xdf, 0x25, 0xcf, 0xc1, 0x52,
	0x27, 0xe3, 0xd9, 0x7d, 0x22, 0x18, 0xb2, 0x37, 0x22, 0xbd, 0x93, 0x8a, 0xce, 0x9b, 0xc2, 0x0f,
	0xb0, 0x31, 0x3f, 0x63, 0x0c, 0x9b, 0x79, 0xa2, 0xc6, 0xc8, 0x86, 0x5a, 0x7e, 0x55, 0xe9, 0x83,
	0xc2, 0x4e, 0x8d, 0xb1, 0xfc, 0xc0, 0xdf, 0x3b, 0xe0, 0xb5, 0xc3, 0xd1, 0x76, 0x5e, 0x89, 0xc2,
	0x4b, 0x80, 0x3e, 0x4f, 0x32, 0x21, 0xb5, 0xe1, 0x29, 0x39, 0x01, 0x60, 0xf3, 0xa8, 0x64, 0xdc,
	0x5b, 0xc6, 0xb8, 0xe8, 0xa1, 0x95, 0x8e, 0xee, 0xff, 0xb0, 0xfd, 0x56, 0x49, 0xad, 0x29, 0x8e,
	0x50, 0xa1, 0x48, 0x51, 0x9f, 0xee, 0xc3, 0xb3, 0x54, 0xe6, 0x51, 0x26, 0x65, 0x36, 0xc1, 0x88,
	0xe1, 0xd4, 0x48, 0x39, 0xd1, 0x55, 0xce, 0x2f, 0x0d, 0xf7, 0xf3, 0xe2, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf6, 0x72, 0x40, 0x50, 0xd7, 0x04, 0x00, 0x00,
}
