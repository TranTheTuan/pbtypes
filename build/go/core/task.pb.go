// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: proto/core/task.proto

package core

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

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_core_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_proto_core_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_proto_core_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type BatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberOfTasks           uint32   `protobuf:"varint,1,opt,name=number_of_tasks,json=numberOfTasks,proto3" json:"number_of_tasks,omitempty"`
	NumberOfSuccessfulTasks uint32   `protobuf:"varint,2,opt,name=number_of_successful_tasks,json=numberOfSuccessfulTasks,proto3" json:"number_of_successful_tasks,omitempty"`
	FailedTaskIds           []uint32 `protobuf:"varint,3,rep,packed,name=failed_task_ids,json=failedTaskIds,proto3" json:"failed_task_ids,omitempty"`
}

func (x *BatchResponse) Reset() {
	*x = BatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_core_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchResponse) ProtoMessage() {}

func (x *BatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_core_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchResponse.ProtoReflect.Descriptor instead.
func (*BatchResponse) Descriptor() ([]byte, []int) {
	return file_proto_core_task_proto_rawDescGZIP(), []int{1}
}

func (x *BatchResponse) GetNumberOfTasks() uint32 {
	if x != nil {
		return x.NumberOfTasks
	}
	return 0
}

func (x *BatchResponse) GetNumberOfSuccessfulTasks() uint32 {
	if x != nil {
		return x.NumberOfSuccessfulTasks
	}
	return 0
}

func (x *BatchResponse) GetFailedTaskIds() []uint32 {
	if x != nil {
		return x.FailedTaskIds
	}
	return nil
}

var File_proto_core_task_proto protoreflect.FileDescriptor

var file_proto_core_task_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x22, 0x4c, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9c,
	0x01, 0x0a, 0x0d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x4f, 0x66, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x3b, 0x0a, 0x1a, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d,
	0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x73, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x72, 0x61, 0x6e,
	0x54, 0x68, 0x65, 0x54, 0x75, 0x61, 0x6e, 0x2f, 0x70, 0x62, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_core_task_proto_rawDescOnce sync.Once
	file_proto_core_task_proto_rawDescData = file_proto_core_task_proto_rawDesc
)

func file_proto_core_task_proto_rawDescGZIP() []byte {
	file_proto_core_task_proto_rawDescOnce.Do(func() {
		file_proto_core_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_core_task_proto_rawDescData)
	})
	return file_proto_core_task_proto_rawDescData
}

var file_proto_core_task_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_core_task_proto_goTypes = []interface{}{
	(*Task)(nil),          // 0: core.v1.Task
	(*BatchResponse)(nil), // 1: core.v1.BatchResponse
}
var file_proto_core_task_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_core_task_proto_init() }
func file_proto_core_task_proto_init() {
	if File_proto_core_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_core_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_proto_core_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchResponse); i {
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
			RawDescriptor: file_proto_core_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_core_task_proto_goTypes,
		DependencyIndexes: file_proto_core_task_proto_depIdxs,
		MessageInfos:      file_proto_core_task_proto_msgTypes,
	}.Build()
	File_proto_core_task_proto = out.File
	file_proto_core_task_proto_rawDesc = nil
	file_proto_core_task_proto_goTypes = nil
	file_proto_core_task_proto_depIdxs = nil
}
