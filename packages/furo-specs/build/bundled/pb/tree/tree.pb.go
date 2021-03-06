// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tree/tree.proto

package tree

import (
	furo "../furo"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Navigation tree type with recursive navigation nodes
type Tree struct {
	// description of the tree
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// String representation of the tree
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Id of the tree
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Rootnode of the tree
	Root                 *Navigationnode `protobuf:"bytes,10,opt,name=root,proto3" json:"root,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Tree) Reset()         { *m = Tree{} }
func (m *Tree) String() string { return proto.CompactTextString(m) }
func (*Tree) ProtoMessage()    {}
func (*Tree) Descriptor() ([]byte, []int) {
	return fileDescriptor_87889ecf37a49aa2, []int{0}
}

func (m *Tree) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tree.Unmarshal(m, b)
}
func (m *Tree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tree.Marshal(b, m, deterministic)
}
func (m *Tree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tree.Merge(m, src)
}
func (m *Tree) XXX_Size() int {
	return xxx_messageInfo_Tree.Size(m)
}
func (m *Tree) XXX_DiscardUnknown() {
	xxx_messageInfo_Tree.DiscardUnknown(m)
}

var xxx_messageInfo_Tree proto.InternalMessageInfo

func (m *Tree) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Tree) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Tree) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tree) GetRoot() *Navigationnode {
	if m != nil {
		return m.Root
	}
	return nil
}

// TreeEntity with Tree
type TreeEntity struct {
	// contains a tree.Tree
	Data *Tree `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Hateoas links
	Links []*furo.Link `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
	// Meta for the response
	Meta                 *furo.Meta `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TreeEntity) Reset()         { *m = TreeEntity{} }
func (m *TreeEntity) String() string { return proto.CompactTextString(m) }
func (*TreeEntity) ProtoMessage()    {}
func (*TreeEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_87889ecf37a49aa2, []int{1}
}

func (m *TreeEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TreeEntity.Unmarshal(m, b)
}
func (m *TreeEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TreeEntity.Marshal(b, m, deterministic)
}
func (m *TreeEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TreeEntity.Merge(m, src)
}
func (m *TreeEntity) XXX_Size() int {
	return xxx_messageInfo_TreeEntity.Size(m)
}
func (m *TreeEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_TreeEntity.DiscardUnknown(m)
}

var xxx_messageInfo_TreeEntity proto.InternalMessageInfo

func (m *TreeEntity) GetData() *Tree {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *TreeEntity) GetLinks() []*furo.Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *TreeEntity) GetMeta() *furo.Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

// Item of the navigationtree
type Navigationnode struct {
	// Children of this node
	Children []*Navigationnode `protobuf:"bytes,12,rep,name=children,proto3" json:"children,omitempty"`
	// description of the node
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// String representation of the node
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// if node has error
	HasError bool `protobuf:"varint,8,opt,name=has_error,json=hasError,proto3" json:"has_error,omitempty"`
	// icon of the node
	Icon string `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	// Id of the node
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// This node is a group label
	IsGroupLabel bool `protobuf:"varint,11,opt,name=is_group_label,json=isGroupLabel,proto3" json:"is_group_label,omitempty"`
	// key words of the node
	KeyWords string `protobuf:"bytes,7,opt,name=key_words,json=keyWords,proto3" json:"key_words,omitempty"`
	// Deeplink information of this node
	Link *furo.Link `protobuf:"bytes,10,opt,name=link,proto3" json:"link,omitempty"`
	// node is open or not
	Open bool `protobuf:"varint,9,opt,name=open,proto3" json:"open,omitempty"`
	// Which panel (i.e. view, edit, display) opens the node type (which is defined in property link)
	Panel string `protobuf:"bytes,6,opt,name=panel,proto3" json:"panel,omitempty"`
	// Secondary text of the node
	SecondaryText        string   `protobuf:"bytes,3,opt,name=secondary_text,json=secondaryText,proto3" json:"secondary_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Navigationnode) Reset()         { *m = Navigationnode{} }
func (m *Navigationnode) String() string { return proto.CompactTextString(m) }
func (*Navigationnode) ProtoMessage()    {}
func (*Navigationnode) Descriptor() ([]byte, []int) {
	return fileDescriptor_87889ecf37a49aa2, []int{2}
}

func (m *Navigationnode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Navigationnode.Unmarshal(m, b)
}
func (m *Navigationnode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Navigationnode.Marshal(b, m, deterministic)
}
func (m *Navigationnode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Navigationnode.Merge(m, src)
}
func (m *Navigationnode) XXX_Size() int {
	return xxx_messageInfo_Navigationnode.Size(m)
}
func (m *Navigationnode) XXX_DiscardUnknown() {
	xxx_messageInfo_Navigationnode.DiscardUnknown(m)
}

var xxx_messageInfo_Navigationnode proto.InternalMessageInfo

func (m *Navigationnode) GetChildren() []*Navigationnode {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *Navigationnode) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Navigationnode) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Navigationnode) GetHasError() bool {
	if m != nil {
		return m.HasError
	}
	return false
}

func (m *Navigationnode) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Navigationnode) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Navigationnode) GetIsGroupLabel() bool {
	if m != nil {
		return m.IsGroupLabel
	}
	return false
}

func (m *Navigationnode) GetKeyWords() string {
	if m != nil {
		return m.KeyWords
	}
	return ""
}

func (m *Navigationnode) GetLink() *furo.Link {
	if m != nil {
		return m.Link
	}
	return nil
}

func (m *Navigationnode) GetOpen() bool {
	if m != nil {
		return m.Open
	}
	return false
}

func (m *Navigationnode) GetPanel() string {
	if m != nil {
		return m.Panel
	}
	return ""
}

func (m *Navigationnode) GetSecondaryText() string {
	if m != nil {
		return m.SecondaryText
	}
	return ""
}

// TreeCollection with repeated TreeEntity
type TreeCollection struct {
	// Contains a tree.TreeEntity repeated
	Entities []*TreeEntity `protobuf:"bytes,4,rep,name=entities,proto3" json:"entities,omitempty"`
	// Hateoas links
	Links []*furo.Link `protobuf:"bytes,3,rep,name=links,proto3" json:"links,omitempty"`
	// Meta for the response
	Meta                 *furo.Meta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TreeCollection) Reset()         { *m = TreeCollection{} }
func (m *TreeCollection) String() string { return proto.CompactTextString(m) }
func (*TreeCollection) ProtoMessage()    {}
func (*TreeCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_87889ecf37a49aa2, []int{3}
}

func (m *TreeCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TreeCollection.Unmarshal(m, b)
}
func (m *TreeCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TreeCollection.Marshal(b, m, deterministic)
}
func (m *TreeCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TreeCollection.Merge(m, src)
}
func (m *TreeCollection) XXX_Size() int {
	return xxx_messageInfo_TreeCollection.Size(m)
}
func (m *TreeCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_TreeCollection.DiscardUnknown(m)
}

var xxx_messageInfo_TreeCollection proto.InternalMessageInfo

func (m *TreeCollection) GetEntities() []*TreeEntity {
	if m != nil {
		return m.Entities
	}
	return nil
}

func (m *TreeCollection) GetLinks() []*furo.Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *TreeCollection) GetMeta() *furo.Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*Tree)(nil), "tree.Tree")
	proto.RegisterType((*TreeEntity)(nil), "tree.TreeEntity")
	proto.RegisterType((*Navigationnode)(nil), "tree.Navigationnode")
	proto.RegisterType((*TreeCollection)(nil), "tree.TreeCollection")
}

func init() { proto.RegisterFile("tree/tree.proto", fileDescriptor_87889ecf37a49aa2) }

var fileDescriptor_87889ecf37a49aa2 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x1d, 0xa7, 0x38, 0x93, 0xe0, 0xa2, 0x55, 0x0f, 0x2b, 0x90, 0x90, 0x89, 0x40, 0xca,
	0x01, 0xa5, 0xa8, 0x7c, 0x02, 0xaa, 0xb8, 0x94, 0x1e, 0xac, 0x4a, 0x1c, 0xad, 0xad, 0x77, 0x68,
	0x56, 0xd9, 0xec, 0x5a, 0xbb, 0x5b, 0xa8, 0x6f, 0x5c, 0xf8, 0x6b, 0x0e, 0x68, 0xc6, 0x51, 0x5a,
	0xa0, 0x12, 0x12, 0x97, 0x68, 0xe7, 0xcd, 0x9b, 0x97, 0x99, 0x37, 0x63, 0x38, 0x4e, 0x01, 0xf1,
	0x94, 0x7e, 0xd6, 0x7d, 0xf0, 0xc9, 0x8b, 0x82, 0xde, 0xcf, 0x8f, 0xbf, 0xdc, 0x06, 0x7f, 0xba,
	0xc3, 0xa4, 0x46, 0x78, 0x0f, 0x58, 0xe3, 0xb6, 0x23, 0xb0, 0xfc, 0x91, 0x41, 0x71, 0x15, 0x10,
	0x45, 0x0d, 0x73, 0x8d, 0xb1, 0x0b, 0xa6, 0x4f, 0xc6, 0x3b, 0x39, 0xa9, 0xb3, 0xd5, 0xac, 0x79,
	0x08, 0x89, 0x57, 0xb0, 0xd0, 0x26, 0xf6, 0x56, 0x0d, 0xad, 0x53, 0x3b, 0x94, 0xf9, 0x9e, 0x32,
	0x62, 0x97, 0x6a, 0x87, 0xa2, 0x82, 0xdc, 0x68, 0x99, 0x71, 0x22, 0x37, 0x5a, 0xac, 0xa0, 0x08,
	0xde, 0x27, 0x09, 0x75, 0xb6, 0x9a, 0x9f, 0x9d, 0xac, 0xb9, 0xc1, 0x4b, 0xf5, 0xd5, 0xdc, 0x28,
	0x92, 0x74, 0x5e, 0x63, 0xc3, 0x8c, 0xa5, 0x03, 0xa0, 0x36, 0xce, 0x5d, 0x32, 0x69, 0x10, 0x2f,
	0xa1, 0xd0, 0x2a, 0x29, 0x56, 0x9a, 0x9f, 0xc1, 0x58, 0x47, 0xf9, 0x86, 0x71, 0x51, 0xc3, 0x94,
	0x66, 0x88, 0x32, 0xaf, 0x27, 0x4c, 0xa0, 0xb1, 0xd6, 0x17, 0xc6, 0x6d, 0x9b, 0x31, 0x41, 0x0a,
	0x34, 0x36, 0xcf, 0x71, 0x20, 0x7c, 0xc2, 0xa4, 0x1a, 0xc6, 0x97, 0x3f, 0x73, 0xa8, 0x7e, 0x6f,
	0x44, 0xbc, 0x83, 0xb2, 0xdb, 0x18, 0xab, 0x03, 0x3a, 0xb9, 0x60, 0xdd, 0xc7, 0x1b, 0x3e, 0xb0,
	0xfe, 0xf4, 0xac, 0xf8, 0x2f, 0xcf, 0x5e, 0xc0, 0x6c, 0xa3, 0x62, 0x8b, 0x21, 0xf8, 0x20, 0xcb,
	0x3a, 0x5b, 0x95, 0x4d, 0xb9, 0x51, 0xf1, 0x9c, 0x62, 0x21, 0xa0, 0x30, 0x9d, 0x77, 0x72, 0xca,
	0x75, 0xfc, 0xfe, 0xcb, 0xe4, 0xd7, 0x50, 0x99, 0xd8, 0xde, 0x04, 0x7f, 0xdb, 0xb7, 0x56, 0x5d,
	0xa3, 0x95, 0x73, 0x56, 0x59, 0x98, 0xf8, 0x91, 0xc0, 0x0b, 0xc2, 0xe8, 0x6f, 0xb6, 0x38, 0xb4,
	0xdf, 0x7c, 0xd0, 0x51, 0x3e, 0xe1, 0xe2, 0x72, 0x8b, 0xc3, 0x67, 0x8a, 0xc9, 0x2d, 0xb2, 0x6d,
	0xbf, 0xa7, 0x87, 0x76, 0x32, 0x4e, 0x6d, 0xf8, 0x1e, 0x9d, 0x9c, 0xb1, 0x30, 0xbf, 0xc5, 0x09,
	0x4c, 0x7b, 0xe5, 0xd0, 0xca, 0x23, 0x16, 0x1b, 0x03, 0xf1, 0x06, 0xaa, 0x88, 0x9d, 0x77, 0x5a,
	0x85, 0xa1, 0x4d, 0x78, 0x97, 0xf6, 0x97, 0xf4, 0xf4, 0x80, 0x5e, 0xe1, 0x5d, 0x5a, 0x7e, 0xcf,
	0xa0, 0xa2, 0x7d, 0x7e, 0xf0, 0xd6, 0x62, 0xc7, 0x56, 0xbd, 0x85, 0x12, 0x69, 0xfb, 0x06, 0xa3,
	0x2c, 0xd8, 0xfe, 0x67, 0xf7, 0x7b, 0x1f, 0xef, 0xa2, 0x39, 0x30, 0xee, 0x2f, 0x60, 0xf2, 0xaf,
	0x0b, 0xc8, 0x1f, 0xbf, 0x80, 0xeb, 0x23, 0xfe, 0x00, 0xde, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x2b, 0x03, 0x22, 0xc8, 0x3b, 0x03, 0x00, 0x00,
}
