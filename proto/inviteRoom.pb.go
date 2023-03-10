// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/inviteRoom.proto

package proto

import (
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

type GetInviteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetInviteRequest) Reset() {
	*x = GetInviteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inviteRoom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInviteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInviteRequest) ProtoMessage() {}

func (x *GetInviteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inviteRoom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInviteRequest.ProtoReflect.Descriptor instead.
func (*GetInviteRequest) Descriptor() ([]byte, []int) {
	return file_proto_inviteRoom_proto_rawDescGZIP(), []int{0}
}

func (x *GetInviteRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetInviteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCreatorID string `protobuf:"bytes,1,opt,name=userCreatorID,proto3" json:"userCreatorID,omitempty"`
	RoomID        string `protobuf:"bytes,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	Place         string `protobuf:"bytes,3,opt,name=place,proto3" json:"place,omitempty"`
	Date          string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetInviteResponse) Reset() {
	*x = GetInviteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inviteRoom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInviteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInviteResponse) ProtoMessage() {}

func (x *GetInviteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inviteRoom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInviteResponse.ProtoReflect.Descriptor instead.
func (*GetInviteResponse) Descriptor() ([]byte, []int) {
	return file_proto_inviteRoom_proto_rawDescGZIP(), []int{1}
}

func (x *GetInviteResponse) GetUserCreatorID() string {
	if x != nil {
		return x.UserCreatorID
	}
	return ""
}

func (x *GetInviteResponse) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

func (x *GetInviteResponse) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *GetInviteResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type StorageInviteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCreatorID string `protobuf:"bytes,1,opt,name=userCreatorID,proto3" json:"userCreatorID,omitempty"`
	RoomID        string `protobuf:"bytes,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	Place         string `protobuf:"bytes,3,opt,name=place,proto3" json:"place,omitempty"`
	Date          string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *StorageInviteRequest) Reset() {
	*x = StorageInviteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inviteRoom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageInviteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageInviteRequest) ProtoMessage() {}

func (x *StorageInviteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inviteRoom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageInviteRequest.ProtoReflect.Descriptor instead.
func (*StorageInviteRequest) Descriptor() ([]byte, []int) {
	return file_proto_inviteRoom_proto_rawDescGZIP(), []int{2}
}

func (x *StorageInviteRequest) GetUserCreatorID() string {
	if x != nil {
		return x.UserCreatorID
	}
	return ""
}

func (x *StorageInviteRequest) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

func (x *StorageInviteRequest) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *StorageInviteRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type StorageInviteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StorageInviteResponse) Reset() {
	*x = StorageInviteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inviteRoom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageInviteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageInviteResponse) ProtoMessage() {}

func (x *StorageInviteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inviteRoom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageInviteResponse.ProtoReflect.Descriptor instead.
func (*StorageInviteResponse) Descriptor() ([]byte, []int) {
	return file_proto_inviteRoom_proto_rawDescGZIP(), []int{3}
}

var File_proto_inviteRoom_proto protoreflect.FileDescriptor

var file_proto_inviteRoom_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x6f,
	0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x52, 0x6f, 0x6f, 0x6d, 0x22, 0x2a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x22, 0x7b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x73,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x7e, 0x0a,
	0x14, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x73,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x17, 0x0a,
	0x15, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb0, 0x01, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x4a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x12, 0x1c, 0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x56, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x12, 0x20, 0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x6f, 0x6f,
	0x6d, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x76, 0x61, 0x6e, 0x56, 0x6f, 0x6a, 0x6e,
	0x69, 0x63, 0x2f, 0x62, 0x61, 0x6e, 0x64, 0x45, 0x46, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_inviteRoom_proto_rawDescOnce sync.Once
	file_proto_inviteRoom_proto_rawDescData = file_proto_inviteRoom_proto_rawDesc
)

func file_proto_inviteRoom_proto_rawDescGZIP() []byte {
	file_proto_inviteRoom_proto_rawDescOnce.Do(func() {
		file_proto_inviteRoom_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_inviteRoom_proto_rawDescData)
	})
	return file_proto_inviteRoom_proto_rawDescData
}

var file_proto_inviteRoom_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_inviteRoom_proto_goTypes = []interface{}{
	(*GetInviteRequest)(nil),      // 0: inviteRoom.GetInviteRequest
	(*GetInviteResponse)(nil),     // 1: inviteRoom.GetInviteResponse
	(*StorageInviteRequest)(nil),  // 2: inviteRoom.StorageInviteRequest
	(*StorageInviteResponse)(nil), // 3: inviteRoom.StorageInviteResponse
}
var file_proto_inviteRoom_proto_depIdxs = []int32{
	0, // 0: inviteRoom.inviteRoom.GetInvite:input_type -> inviteRoom.GetInviteRequest
	2, // 1: inviteRoom.inviteRoom.StorageInvite:input_type -> inviteRoom.StorageInviteRequest
	1, // 2: inviteRoom.inviteRoom.GetInvite:output_type -> inviteRoom.GetInviteResponse
	3, // 3: inviteRoom.inviteRoom.StorageInvite:output_type -> inviteRoom.StorageInviteResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_inviteRoom_proto_init() }
func file_proto_inviteRoom_proto_init() {
	if File_proto_inviteRoom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_inviteRoom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInviteRequest); i {
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
		file_proto_inviteRoom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInviteResponse); i {
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
		file_proto_inviteRoom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageInviteRequest); i {
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
		file_proto_inviteRoom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageInviteResponse); i {
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
			RawDescriptor: file_proto_inviteRoom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_inviteRoom_proto_goTypes,
		DependencyIndexes: file_proto_inviteRoom_proto_depIdxs,
		MessageInfos:      file_proto_inviteRoom_proto_msgTypes,
	}.Build()
	File_proto_inviteRoom_proto = out.File
	file_proto_inviteRoom_proto_rawDesc = nil
	file_proto_inviteRoom_proto_goTypes = nil
	file_proto_inviteRoom_proto_depIdxs = nil
}