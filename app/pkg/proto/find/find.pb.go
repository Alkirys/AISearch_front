// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/find/find.proto

package find

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type FindHandlerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Img string `protobuf:"bytes,1,opt,name=img,proto3" json:"img,omitempty"`
}

func (x *FindHandlerRequest) Reset() {
	*x = FindHandlerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_find_find_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindHandlerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindHandlerRequest) ProtoMessage() {}

func (x *FindHandlerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_find_find_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindHandlerRequest.ProtoReflect.Descriptor instead.
func (*FindHandlerRequest) Descriptor() ([]byte, []int) {
	return file_proto_find_find_proto_rawDescGZIP(), []int{0}
}

func (x *FindHandlerRequest) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

type FindHandlerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imgs []*FindHandlerResponseImage `protobuf:"bytes,1,rep,name=imgs,proto3" json:"imgs,omitempty"`
}

func (x *FindHandlerResponse) Reset() {
	*x = FindHandlerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_find_find_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindHandlerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindHandlerResponse) ProtoMessage() {}

func (x *FindHandlerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_find_find_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindHandlerResponse.ProtoReflect.Descriptor instead.
func (*FindHandlerResponse) Descriptor() ([]byte, []int) {
	return file_proto_find_find_proto_rawDescGZIP(), []int{1}
}

func (x *FindHandlerResponse) GetImgs() []*FindHandlerResponseImage {
	if x != nil {
		return x.Imgs
	}
	return nil
}

type DetectHandlerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Img *DetectHandlerRequestImage `protobuf:"bytes,1,opt,name=img,proto3" json:"img,omitempty"`
}

func (x *DetectHandlerRequest) Reset() {
	*x = DetectHandlerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_find_find_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectHandlerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectHandlerRequest) ProtoMessage() {}

func (x *DetectHandlerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_find_find_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectHandlerRequest.ProtoReflect.Descriptor instead.
func (*DetectHandlerRequest) Descriptor() ([]byte, []int) {
	return file_proto_find_find_proto_rawDescGZIP(), []int{2}
}

func (x *DetectHandlerRequest) GetImg() *DetectHandlerRequestImage {
	if x != nil {
		return x.Img
	}
	return nil
}

type DetectHandlerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []string `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *DetectHandlerResponse) Reset() {
	*x = DetectHandlerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_find_find_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectHandlerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectHandlerResponse) ProtoMessage() {}

func (x *DetectHandlerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_find_find_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectHandlerResponse.ProtoReflect.Descriptor instead.
func (*DetectHandlerResponse) Descriptor() ([]byte, []int) {
	return file_proto_find_find_proto_rawDescGZIP(), []int{3}
}

func (x *DetectHandlerResponse) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

type FindHandlerResponseImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Bound []uint64 `protobuf:"varint,2,rep,packed,name=bound,proto3" json:"bound,omitempty"`
}

func (x *FindHandlerResponseImage) Reset() {
	*x = FindHandlerResponseImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_find_find_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindHandlerResponseImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindHandlerResponseImage) ProtoMessage() {}

func (x *FindHandlerResponseImage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_find_find_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindHandlerResponseImage.ProtoReflect.Descriptor instead.
func (*FindHandlerResponseImage) Descriptor() ([]byte, []int) {
	return file_proto_find_find_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FindHandlerResponseImage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FindHandlerResponseImage) GetBound() []uint64 {
	if x != nil {
		return x.Bound
	}
	return nil
}

type DetectHandlerRequestImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Bound []uint64 `protobuf:"varint,2,rep,packed,name=bound,proto3" json:"bound,omitempty"`
}

func (x *DetectHandlerRequestImage) Reset() {
	*x = DetectHandlerRequestImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_find_find_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectHandlerRequestImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectHandlerRequestImage) ProtoMessage() {}

func (x *DetectHandlerRequestImage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_find_find_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectHandlerRequestImage.ProtoReflect.Descriptor instead.
func (*DetectHandlerRequestImage) Descriptor() ([]byte, []int) {
	return file_proto_find_find_proto_rawDescGZIP(), []int{2, 0}
}

func (x *DetectHandlerRequestImage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DetectHandlerRequestImage) GetBound() []uint64 {
	if x != nil {
		return x.Bound
	}
	return nil
}

var File_proto_find_find_proto protoreflect.FileDescriptor

var file_proto_find_find_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x2f, 0x66, 0x69, 0x6e,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x69, 0x6e, 0x64, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x12, 0x46,
	0x69, 0x6e, 0x64, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x69, 0x6d, 0x67, 0x22, 0x7d, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x69, 0x6d,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x69, 0x6e, 0x64, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x04, 0x69, 0x6d, 0x67, 0x73, 0x1a,
	0x31, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x05, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x22, 0x7d, 0x0a, 0x14, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x03, 0x69, 0x6d,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x66, 0x69, 0x6e, 0x64, 0x2e, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x1a, 0x31,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x05, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x22, 0x2d, 0x0a, 0x15, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x32, 0xc8, 0x01, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x5b, 0x0a, 0x0b, 0x46, 0x69, 0x6e,
	0x64, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x66, 0x69, 0x6e, 0x64, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x66, 0x69, 0x6e, 0x64, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66,
	0x69, 0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x63, 0x0a, 0x0d, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x66, 0x69, 0x6e, 0x64, 0x2e, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x69, 0x6e, 0x64, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x30, 0x5a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x6c, 0x64, 0x65, 0x72, 0x6c,
	0x79, 0x2d, 0x41, 0x49, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x3b, 0x66, 0x69, 0x6e, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_find_find_proto_rawDescOnce sync.Once
	file_proto_find_find_proto_rawDescData = file_proto_find_find_proto_rawDesc
)

func file_proto_find_find_proto_rawDescGZIP() []byte {
	file_proto_find_find_proto_rawDescOnce.Do(func() {
		file_proto_find_find_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_find_find_proto_rawDescData)
	})
	return file_proto_find_find_proto_rawDescData
}

var file_proto_find_find_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_find_find_proto_goTypes = []interface{}{
	(*FindHandlerRequest)(nil),        // 0: find.FindHandlerRequest
	(*FindHandlerResponse)(nil),       // 1: find.FindHandlerResponse
	(*DetectHandlerRequest)(nil),      // 2: find.DetectHandlerRequest
	(*DetectHandlerResponse)(nil),     // 3: find.DetectHandlerResponse
	(*FindHandlerResponseImage)(nil),  // 4: find.FindHandlerResponse.image
	(*DetectHandlerRequestImage)(nil), // 5: find.DetectHandlerRequest.image
}
var file_proto_find_find_proto_depIdxs = []int32{
	4, // 0: find.FindHandlerResponse.imgs:type_name -> find.FindHandlerResponse.image
	5, // 1: find.DetectHandlerRequest.img:type_name -> find.DetectHandlerRequest.image
	0, // 2: find.Find.FindHandler:input_type -> find.FindHandlerRequest
	2, // 3: find.Find.DetectHandler:input_type -> find.DetectHandlerRequest
	1, // 4: find.Find.FindHandler:output_type -> find.FindHandlerResponse
	3, // 5: find.Find.DetectHandler:output_type -> find.DetectHandlerResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_find_find_proto_init() }
func file_proto_find_find_proto_init() {
	if File_proto_find_find_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_find_find_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindHandlerRequest); i {
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
		file_proto_find_find_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindHandlerResponse); i {
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
		file_proto_find_find_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectHandlerRequest); i {
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
		file_proto_find_find_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectHandlerResponse); i {
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
		file_proto_find_find_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindHandlerResponseImage); i {
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
		file_proto_find_find_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectHandlerRequestImage); i {
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
			RawDescriptor: file_proto_find_find_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_find_find_proto_goTypes,
		DependencyIndexes: file_proto_find_find_proto_depIdxs,
		MessageInfos:      file_proto_find_find_proto_msgTypes,
	}.Build()
	File_proto_find_find_proto = out.File
	file_proto_find_find_proto_rawDesc = nil
	file_proto_find_find_proto_goTypes = nil
	file_proto_find_find_proto_depIdxs = nil
}
