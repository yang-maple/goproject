// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.5
// source: userinfo.proto

package protoc

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

// 枚举类型
type Sex int32

const (
	Sex_Man   Sex = 0
	Sex_Woman Sex = 1
)

// Enum value maps for Sex.
var (
	Sex_name = map[int32]string{
		0: "Man",
		1: "Woman",
	}
	Sex_value = map[string]int32{
		"Man":   0,
		"Woman": 1,
	}
)

func (x Sex) Enum() *Sex {
	p := new(Sex)
	*p = x
	return p
}

func (x Sex) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sex) Descriptor() protoreflect.EnumDescriptor {
	return file_userinfo_proto_enumTypes[0].Descriptor()
}

func (Sex) Type() protoreflect.EnumType {
	return &file_userinfo_proto_enumTypes[0]
}

func (x Sex) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sex.Descriptor instead.
func (Sex) EnumDescriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{0}
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 字符串类型
	UserName string `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	// 整型类型
	Age int32 `protobuf:"varint,2,opt,name=Age,proto3" json:"Age,omitempty"`
	// 数组类型
	Hobby []string `protobuf:"bytes,3,rep,name=hobby,proto3" json:"hobby,omitempty"`
	// 枚举类型
	Sex Sex `protobuf:"varint,4,opt,name=Sex,proto3,enum=Sex" json:"Sex,omitempty"`
	// 内嵌类型
	Food *Food `protobuf:"bytes,5,opt,name=Food,proto3" json:"Food,omitempty"`
	// map类型
	Address map[string]*UserInfo_Address `protobuf:"bytes,6,rep,name=address,proto3" json:"address,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Score   map[string]int32             `protobuf:"bytes,7,rep,name=score,proto3" json:"score,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserInfo) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserInfo) GetHobby() []string {
	if x != nil {
		return x.Hobby
	}
	return nil
}

func (x *UserInfo) GetSex() Sex {
	if x != nil {
		return x.Sex
	}
	return Sex_Man
}

func (x *UserInfo) GetFood() *Food {
	if x != nil {
		return x.Food
	}
	return nil
}

func (x *UserInfo) GetAddress() map[string]*UserInfo_Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *UserInfo) GetScore() map[string]int32 {
	if x != nil {
		return x.Score
	}
	return nil
}

// 嵌套
type Food struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *Food) Reset() {
	*x = Food{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Food) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Food) ProtoMessage() {}

func (x *Food) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Food.ProtoReflect.Descriptor instead.
func (*Food) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{1}
}

func (x *Food) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 内嵌套类型
type UserInfo_Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Province string `protobuf:"bytes,1,opt,name=Province,proto3" json:"Province,omitempty"`
	City     string `protobuf:"bytes,2,opt,name=City,proto3" json:"City,omitempty"`
}

func (x *UserInfo_Address) Reset() {
	*x = UserInfo_Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo_Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo_Address) ProtoMessage() {}

func (x *UserInfo_Address) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo_Address.ProtoReflect.Descriptor instead.
func (*UserInfo_Address) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{0, 0}
}

func (x *UserInfo_Address) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *UserInfo_Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

var File_userinfo_proto protoreflect.FileDescriptor

var file_userinfo_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa3, 0x03, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x41, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x68,
	0x6f, 0x62, 0x62, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x62, 0x62,
	0x79, 0x12, 0x16, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x04,
	0x2e, 0x53, 0x65, 0x78, 0x52, 0x03, 0x53, 0x65, 0x78, 0x12, 0x19, 0x0a, 0x04, 0x46, 0x6f, 0x6f,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x04,
	0x46, 0x6f, 0x6f, 0x64, 0x12, 0x30, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x1a, 0x39, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x1a, 0x4d, 0x0a,
	0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x38, 0x0a, 0x0a,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1a, 0x0a, 0x04, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x2a, 0x19, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x61, 0x6e,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x6f, 0x6d, 0x61, 0x6e, 0x10, 0x01, 0x42, 0x0b, 0x5a,
	0x09, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_userinfo_proto_rawDescOnce sync.Once
	file_userinfo_proto_rawDescData = file_userinfo_proto_rawDesc
)

func file_userinfo_proto_rawDescGZIP() []byte {
	file_userinfo_proto_rawDescOnce.Do(func() {
		file_userinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_userinfo_proto_rawDescData)
	})
	return file_userinfo_proto_rawDescData
}

var file_userinfo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_userinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_userinfo_proto_goTypes = []interface{}{
	(Sex)(0),                 // 0: Sex
	(*UserInfo)(nil),         // 1: UserInfo
	(*Food)(nil),             // 2: Food
	(*UserInfo_Address)(nil), // 3: UserInfo.Address
	nil,                      // 4: UserInfo.AddressEntry
	nil,                      // 5: UserInfo.ScoreEntry
}
var file_userinfo_proto_depIdxs = []int32{
	0, // 0: UserInfo.Sex:type_name -> Sex
	2, // 1: UserInfo.Food:type_name -> Food
	4, // 2: UserInfo.address:type_name -> UserInfo.AddressEntry
	5, // 3: UserInfo.score:type_name -> UserInfo.ScoreEntry
	3, // 4: UserInfo.AddressEntry.value:type_name -> UserInfo.Address
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_userinfo_proto_init() }
func file_userinfo_proto_init() {
	if File_userinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_userinfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Food); i {
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
		file_userinfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo_Address); i {
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
			RawDescriptor: file_userinfo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_userinfo_proto_goTypes,
		DependencyIndexes: file_userinfo_proto_depIdxs,
		EnumInfos:         file_userinfo_proto_enumTypes,
		MessageInfos:      file_userinfo_proto_msgTypes,
	}.Build()
	File_userinfo_proto = out.File
	file_userinfo_proto_rawDesc = nil
	file_userinfo_proto_goTypes = nil
	file_userinfo_proto_depIdxs = nil
}
