// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: cube.proto

package pb

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

type CellData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoreData []byte `protobuf:"bytes,1,opt,name=coreData,proto3" json:"coreData,omitempty"`
}

func (x *CellData) Reset() {
	*x = CellData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cube_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CellData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CellData) ProtoMessage() {}

func (x *CellData) ProtoReflect() protoreflect.Message {
	mi := &file_cube_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CellData.ProtoReflect.Descriptor instead.
func (*CellData) Descriptor() ([]byte, []int) {
	return file_cube_proto_rawDescGZIP(), []int{0}
}

func (x *CellData) GetCoreData() []byte {
	if x != nil {
		return x.CoreData
	}
	return nil
}

type RowData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CellMap map[string]*CellData `protobuf:"bytes,1,rep,name=cellMap,proto3" json:"cellMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RowData) Reset() {
	*x = RowData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cube_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RowData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RowData) ProtoMessage() {}

func (x *RowData) ProtoReflect() protoreflect.Message {
	mi := &file_cube_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RowData.ProtoReflect.Descriptor instead.
func (*RowData) Descriptor() ([]byte, []int) {
	return file_cube_proto_rawDescGZIP(), []int{1}
}

func (x *RowData) GetCellMap() map[string]*CellData {
	if x != nil {
		return x.CellMap
	}
	return nil
}

type BoardData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowMap map[string]*RowData `protobuf:"bytes,1,rep,name=rowMap,proto3" json:"rowMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BoardData) Reset() {
	*x = BoardData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cube_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardData) ProtoMessage() {}

func (x *BoardData) ProtoReflect() protoreflect.Message {
	mi := &file_cube_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardData.ProtoReflect.Descriptor instead.
func (*BoardData) Descriptor() ([]byte, []int) {
	return file_cube_proto_rawDescGZIP(), []int{2}
}

func (x *BoardData) GetRowMap() map[string]*RowData {
	if x != nil {
		return x.RowMap
	}
	return nil
}

type CubeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardMap map[string]*BoardData `protobuf:"bytes,1,rep,name=boardMap,proto3" json:"boardMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CubeData) Reset() {
	*x = CubeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cube_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CubeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CubeData) ProtoMessage() {}

func (x *CubeData) ProtoReflect() protoreflect.Message {
	mi := &file_cube_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CubeData.ProtoReflect.Descriptor instead.
func (*CubeData) Descriptor() ([]byte, []int) {
	return file_cube_proto_rawDescGZIP(), []int{3}
}

func (x *CubeData) GetBoardMap() map[string]*BoardData {
	if x != nil {
		return x.BoardMap
	}
	return nil
}

var File_cube_proto protoreflect.FileDescriptor

var file_cube_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x08,
	0x43, 0x65, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x72, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6f, 0x72, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x81, 0x01, 0x0a, 0x07, 0x52, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x2f, 0x0a, 0x07, 0x63, 0x65, 0x6c, 0x6c, 0x4d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x52, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x65, 0x6c, 0x6c,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x65, 0x6c, 0x6c, 0x4d, 0x61,
	0x70, 0x1a, 0x45, 0x0a, 0x0c, 0x43, 0x65, 0x6c, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x80, 0x01, 0x0a, 0x09, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x06, 0x72, 0x6f, 0x77, 0x4d, 0x61, 0x70,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x52, 0x6f, 0x77, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x72, 0x6f, 0x77, 0x4d, 0x61, 0x70, 0x1a, 0x43, 0x0a, 0x0b, 0x52, 0x6f, 0x77, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x88, 0x01, 0x0a, 0x08,
	0x43, 0x75, 0x62, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x08, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x4d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x43, 0x75, 0x62,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4d, 0x61, 0x70, 0x1a, 0x47, 0x0a,
	0x0d, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cube_proto_rawDescOnce sync.Once
	file_cube_proto_rawDescData = file_cube_proto_rawDesc
)

func file_cube_proto_rawDescGZIP() []byte {
	file_cube_proto_rawDescOnce.Do(func() {
		file_cube_proto_rawDescData = protoimpl.X.CompressGZIP(file_cube_proto_rawDescData)
	})
	return file_cube_proto_rawDescData
}

var file_cube_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cube_proto_goTypes = []interface{}{
	(*CellData)(nil),  // 0: CellData
	(*RowData)(nil),   // 1: RowData
	(*BoardData)(nil), // 2: BoardData
	(*CubeData)(nil),  // 3: CubeData
	nil,               // 4: RowData.CellMapEntry
	nil,               // 5: BoardData.RowMapEntry
	nil,               // 6: CubeData.BoardMapEntry
}
var file_cube_proto_depIdxs = []int32{
	4, // 0: RowData.cellMap:type_name -> RowData.CellMapEntry
	5, // 1: BoardData.rowMap:type_name -> BoardData.RowMapEntry
	6, // 2: CubeData.boardMap:type_name -> CubeData.BoardMapEntry
	0, // 3: RowData.CellMapEntry.value:type_name -> CellData
	1, // 4: BoardData.RowMapEntry.value:type_name -> RowData
	2, // 5: CubeData.BoardMapEntry.value:type_name -> BoardData
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cube_proto_init() }
func file_cube_proto_init() {
	if File_cube_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cube_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CellData); i {
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
		file_cube_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RowData); i {
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
		file_cube_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardData); i {
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
		file_cube_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CubeData); i {
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
			RawDescriptor: file_cube_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cube_proto_goTypes,
		DependencyIndexes: file_cube_proto_depIdxs,
		MessageInfos:      file_cube_proto_msgTypes,
	}.Build()
	File_cube_proto = out.File
	file_cube_proto_rawDesc = nil
	file_cube_proto_goTypes = nil
	file_cube_proto_depIdxs = nil
}