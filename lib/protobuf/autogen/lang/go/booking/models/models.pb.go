// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: src/booking/models/models.proto

package models

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Booking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VenueId    string `protobuf:"bytes,2,opt,name=venueId,proto3" json:"venueId,omitempty"`
	Email      string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	People     uint32 `protobuf:"varint,4,opt,name=people,proto3" json:"people,omitempty"`
	StartsAt   string `protobuf:"bytes,5,opt,name=startsAt,proto3" json:"startsAt,omitempty"`
	EndsAt     string `protobuf:"bytes,6,opt,name=endsAt,proto3" json:"endsAt,omitempty"`
	Duration   uint32 `protobuf:"varint,7,opt,name=duration,proto3" json:"duration,omitempty"`
	TableId    string `protobuf:"bytes,8,opt,name=tableId,proto3" json:"tableId,omitempty"`
	FamilyName string `protobuf:"bytes,9,opt,name=familyName,proto3" json:"familyName,omitempty"`
	GivenName  string `protobuf:"bytes,10,opt,name=givenName,proto3" json:"givenName,omitempty"`
}

func (x *Booking) Reset() {
	*x = Booking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_booking_models_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

func (x *Booking) ProtoReflect() protoreflect.Message {
	mi := &file_src_booking_models_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booking.ProtoReflect.Descriptor instead.
func (*Booking) Descriptor() ([]byte, []int) {
	return file_src_booking_models_models_proto_rawDescGZIP(), []int{0}
}

func (x *Booking) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Booking) GetVenueId() string {
	if x != nil {
		return x.VenueId
	}
	return ""
}

func (x *Booking) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Booking) GetPeople() uint32 {
	if x != nil {
		return x.People
	}
	return 0
}

func (x *Booking) GetStartsAt() string {
	if x != nil {
		return x.StartsAt
	}
	return ""
}

func (x *Booking) GetEndsAt() string {
	if x != nil {
		return x.EndsAt
	}
	return ""
}

func (x *Booking) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Booking) GetTableId() string {
	if x != nil {
		return x.TableId
	}
	return ""
}

func (x *Booking) GetFamilyName() string {
	if x != nil {
		return x.FamilyName
	}
	return ""
}

func (x *Booking) GetGivenName() string {
	if x != nil {
		return x.GivenName
	}
	return ""
}

type Slot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VenueId  string `protobuf:"bytes,1,opt,name=venueId,proto3" json:"venueId,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	People   uint32 `protobuf:"varint,3,opt,name=people,proto3" json:"people,omitempty"`
	StartsAt string `protobuf:"bytes,4,opt,name=startsAt,proto3" json:"startsAt,omitempty"`
	EndsAt   string `protobuf:"bytes,5,opt,name=endsAt,proto3" json:"endsAt,omitempty"`
	Duration uint32 `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *Slot) Reset() {
	*x = Slot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_booking_models_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slot) ProtoMessage() {}

func (x *Slot) ProtoReflect() protoreflect.Message {
	mi := &file_src_booking_models_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slot.ProtoReflect.Descriptor instead.
func (*Slot) Descriptor() ([]byte, []int) {
	return file_src_booking_models_models_proto_rawDescGZIP(), []int{1}
}

func (x *Slot) GetVenueId() string {
	if x != nil {
		return x.VenueId
	}
	return ""
}

func (x *Slot) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Slot) GetPeople() uint32 {
	if x != nil {
		return x.People
	}
	return 0
}

func (x *Slot) GetStartsAt() string {
	if x != nil {
		return x.StartsAt
	}
	return ""
}

func (x *Slot) GetEndsAt() string {
	if x != nil {
		return x.EndsAt
	}
	return ""
}

func (x *Slot) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

var File_src_booking_models_models_proto protoreflect.FileDescriptor

var file_src_booking_models_models_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x72, 0x63, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x22, 0x89, 0x02, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x6e, 0x75, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70,
	0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x73, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x6e, 0x64, 0x73, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x9e, 0x01,
	0x0a, 0x04, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e,
	0x64, 0x73, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x64, 0x73,
	0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x52,
	0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x62,
	0x62, 0x69, 0x6e, 0x6d, 0x61, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x6c, 0x61, 0x6e, 0x67,
	0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_booking_models_models_proto_rawDescOnce sync.Once
	file_src_booking_models_models_proto_rawDescData = file_src_booking_models_models_proto_rawDesc
)

func file_src_booking_models_models_proto_rawDescGZIP() []byte {
	file_src_booking_models_models_proto_rawDescOnce.Do(func() {
		file_src_booking_models_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_booking_models_models_proto_rawDescData)
	})
	return file_src_booking_models_models_proto_rawDescData
}

var file_src_booking_models_models_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_booking_models_models_proto_goTypes = []interface{}{
	(*Booking)(nil), // 0: booking.models.Booking
	(*Slot)(nil),    // 1: booking.models.Slot
}
var file_src_booking_models_models_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_booking_models_models_proto_init() }
func file_src_booking_models_models_proto_init() {
	if File_src_booking_models_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_booking_models_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Booking); i {
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
		file_src_booking_models_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slot); i {
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
			RawDescriptor: file_src_booking_models_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_booking_models_models_proto_goTypes,
		DependencyIndexes: file_src_booking_models_models_proto_depIdxs,
		MessageInfos:      file_src_booking_models_models_proto_msgTypes,
	}.Build()
	File_src_booking_models_models_proto = out.File
	file_src_booking_models_models_proto_rawDesc = nil
	file_src_booking_models_models_proto_goTypes = nil
	file_src_booking_models_models_proto_depIdxs = nil
}
