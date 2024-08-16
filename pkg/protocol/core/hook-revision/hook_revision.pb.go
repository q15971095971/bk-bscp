// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: hook_revision.proto

package pbhr

import (
	base "github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/protocol/core/base"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// HookRevision source resource reference: pkg/dal/table/hook_revision.go
type HookRevision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec       *HookRevisionSpec       `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Attachment *HookRevisionAttachment `protobuf:"bytes,3,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Revision   *base.Revision          `protobuf:"bytes,4,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *HookRevision) Reset() {
	*x = HookRevision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hook_revision_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HookRevision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HookRevision) ProtoMessage() {}

func (x *HookRevision) ProtoReflect() protoreflect.Message {
	mi := &file_hook_revision_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HookRevision.ProtoReflect.Descriptor instead.
func (*HookRevision) Descriptor() ([]byte, []int) {
	return file_hook_revision_proto_rawDescGZIP(), []int{0}
}

func (x *HookRevision) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HookRevision) GetSpec() *HookRevisionSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *HookRevision) GetAttachment() *HookRevisionAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *HookRevision) GetRevision() *base.Revision {
	if x != nil {
		return x.Revision
	}
	return nil
}

// HookRevisionAttachment source resource reference: pkg/dal/table/hook_revision.go
type HookRevisionSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	State   string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Memo    string `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *HookRevisionSpec) Reset() {
	*x = HookRevisionSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hook_revision_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HookRevisionSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HookRevisionSpec) ProtoMessage() {}

func (x *HookRevisionSpec) ProtoReflect() protoreflect.Message {
	mi := &file_hook_revision_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HookRevisionSpec.ProtoReflect.Descriptor instead.
func (*HookRevisionSpec) Descriptor() ([]byte, []int) {
	return file_hook_revision_proto_rawDescGZIP(), []int{1}
}

func (x *HookRevisionSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HookRevisionSpec) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *HookRevisionSpec) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *HookRevisionSpec) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

// HookRevisionAttachment source resource reference: pkg/dal/table/hook_revision.go
type HookRevisionAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId  uint32 `protobuf:"varint,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	HookId uint32 `protobuf:"varint,2,opt,name=hook_id,json=hookId,proto3" json:"hook_id,omitempty"`
}

func (x *HookRevisionAttachment) Reset() {
	*x = HookRevisionAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hook_revision_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HookRevisionAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HookRevisionAttachment) ProtoMessage() {}

func (x *HookRevisionAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_hook_revision_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HookRevisionAttachment.ProtoReflect.Descriptor instead.
func (*HookRevisionAttachment) Descriptor() ([]byte, []int) {
	return file_hook_revision_proto_rawDescGZIP(), []int{2}
}

func (x *HookRevisionAttachment) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *HookRevisionAttachment) GetHookId() uint32 {
	if x != nil {
		return x.HookId
	}
	return 0
}

var File_hook_revision_proto protoreflect.FileDescriptor

var file_hook_revision_proto_rawDesc = []byte{
	0x0a, 0x13, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x62, 0x68, 0x72, 0x1a, 0x21, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6,
	0x01, 0x0a, 0x0c, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2a, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x62, 0x68, 0x72, 0x2e, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3c, 0x0a, 0x0a, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x70, 0x62, 0x68, 0x72, 0x2e, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x72, 0x65, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x6a, 0x0a, 0x10, 0x48, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x22, 0x48, 0x0a, 0x16, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a,
	0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x62,
	0x69, 0x7a, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x42, 0x5e, 0x5a,
	0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x6e, 0x63,
	0x65, 0x6e, 0x74, 0x42, 0x6c, 0x75, 0x65, 0x4b, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x6b, 0x2d, 0x62,
	0x63, 0x73, 0x2f, 0x62, 0x63, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x62, 0x63, 0x73, 0x2d, 0x62, 0x73, 0x63, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x6f, 0x6f, 0x6b, 0x2d,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x3b, 0x70, 0x62, 0x68, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hook_revision_proto_rawDescOnce sync.Once
	file_hook_revision_proto_rawDescData = file_hook_revision_proto_rawDesc
)

func file_hook_revision_proto_rawDescGZIP() []byte {
	file_hook_revision_proto_rawDescOnce.Do(func() {
		file_hook_revision_proto_rawDescData = protoimpl.X.CompressGZIP(file_hook_revision_proto_rawDescData)
	})
	return file_hook_revision_proto_rawDescData
}

var file_hook_revision_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_hook_revision_proto_goTypes = []interface{}{
	(*HookRevision)(nil),           // 0: pbhr.HookRevision
	(*HookRevisionSpec)(nil),       // 1: pbhr.HookRevisionSpec
	(*HookRevisionAttachment)(nil), // 2: pbhr.HookRevisionAttachment
	(*base.Revision)(nil),          // 3: pbbase.Revision
}
var file_hook_revision_proto_depIdxs = []int32{
	1, // 0: pbhr.HookRevision.spec:type_name -> pbhr.HookRevisionSpec
	2, // 1: pbhr.HookRevision.attachment:type_name -> pbhr.HookRevisionAttachment
	3, // 2: pbhr.HookRevision.revision:type_name -> pbbase.Revision
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_hook_revision_proto_init() }
func file_hook_revision_proto_init() {
	if File_hook_revision_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hook_revision_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HookRevision); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hook_revision_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HookRevisionSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hook_revision_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HookRevisionAttachment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hook_revision_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hook_revision_proto_goTypes,
		DependencyIndexes: file_hook_revision_proto_depIdxs,
		MessageInfos:      file_hook_revision_proto_msgTypes,
	}.Build()
	File_hook_revision_proto = out.File
	file_hook_revision_proto_rawDesc = nil
	file_hook_revision_proto_goTypes = nil
	file_hook_revision_proto_depIdxs = nil
}