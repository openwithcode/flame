// Copyright 2022 Cisco Systems, Inc. and its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.3
// source: meta.proto

package meta

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

type MetaResponse_Status int32

const (
	MetaResponse_ERROR   MetaResponse_Status = 0 // default
	MetaResponse_SUCCESS MetaResponse_Status = 1
)

// Enum value maps for MetaResponse_Status.
var (
	MetaResponse_Status_name = map[int32]string{
		0: "ERROR",
		1: "SUCCESS",
	}
	MetaResponse_Status_value = map[string]int32{
		"ERROR":   0,
		"SUCCESS": 1,
	}
)

func (x MetaResponse_Status) Enum() *MetaResponse_Status {
	p := new(MetaResponse_Status)
	*p = x
	return p
}

func (x MetaResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetaResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_meta_proto_enumTypes[0].Descriptor()
}

func (MetaResponse_Status) Type() protoreflect.EnumType {
	return &file_meta_proto_enumTypes[0]
}

func (x MetaResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetaResponse_Status.Descriptor instead.
func (MetaResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_meta_proto_rawDescGZIP(), []int{1, 0}
}

type MetaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId    string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	ChName   string `protobuf:"bytes,2,opt,name=ch_name,json=chName,proto3" json:"ch_name,omitempty"`
	Me       string `protobuf:"bytes,3,opt,name=me,proto3" json:"me,omitempty"`
	Other    string `protobuf:"bytes,4,opt,name=other,proto3" json:"other,omitempty"`
	Group    string `protobuf:"bytes,5,opt,name=group,proto3" json:"group,omitempty"`
	Endpoint string `protobuf:"bytes,6,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *MetaInfo) Reset() {
	*x = MetaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaInfo) ProtoMessage() {}

func (x *MetaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaInfo.ProtoReflect.Descriptor instead.
func (*MetaInfo) Descriptor() ([]byte, []int) {
	return file_meta_proto_rawDescGZIP(), []int{0}
}

func (x *MetaInfo) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *MetaInfo) GetChName() string {
	if x != nil {
		return x.ChName
	}
	return ""
}

func (x *MetaInfo) GetMe() string {
	if x != nil {
		return x.Me
	}
	return ""
}

func (x *MetaInfo) GetOther() string {
	if x != nil {
		return x.Other
	}
	return ""
}

func (x *MetaInfo) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *MetaInfo) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type MetaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    MetaResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=grpcMeta.MetaResponse_Status" json:"status,omitempty"`
	Endpoints []string            `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *MetaResponse) Reset() {
	*x = MetaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaResponse) ProtoMessage() {}

func (x *MetaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_meta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaResponse.ProtoReflect.Descriptor instead.
func (*MetaResponse) Descriptor() ([]byte, []int) {
	return file_meta_proto_rawDescGZIP(), []int{1}
}

func (x *MetaResponse) GetStatus() MetaResponse_Status {
	if x != nil {
		return x.Status
	}
	return MetaResponse_ERROR
}

func (x *MetaResponse) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

var File_meta_proto protoreflect.FileDescriptor

var file_meta_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x72,
	0x70, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x92, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x0c,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x01, 0x32, 0x88, 0x01, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x12, 0x40, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74,
	0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x61,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x4d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74,
	0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x73,
	0x63, 0x6f, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x66, 0x6c, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meta_proto_rawDescOnce sync.Once
	file_meta_proto_rawDescData = file_meta_proto_rawDesc
)

func file_meta_proto_rawDescGZIP() []byte {
	file_meta_proto_rawDescOnce.Do(func() {
		file_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_meta_proto_rawDescData)
	})
	return file_meta_proto_rawDescData
}

var file_meta_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_meta_proto_goTypes = []interface{}{
	(MetaResponse_Status)(0), // 0: grpcMeta.MetaResponse.Status
	(*MetaInfo)(nil),         // 1: grpcMeta.MetaInfo
	(*MetaResponse)(nil),     // 2: grpcMeta.MetaResponse
}
var file_meta_proto_depIdxs = []int32{
	0, // 0: grpcMeta.MetaResponse.status:type_name -> grpcMeta.MetaResponse.Status
	1, // 1: grpcMeta.MetaRoute.RegisterMetaInfo:input_type -> grpcMeta.MetaInfo
	1, // 2: grpcMeta.MetaRoute.HeartBeat:input_type -> grpcMeta.MetaInfo
	2, // 3: grpcMeta.MetaRoute.RegisterMetaInfo:output_type -> grpcMeta.MetaResponse
	2, // 4: grpcMeta.MetaRoute.HeartBeat:output_type -> grpcMeta.MetaResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_meta_proto_init() }
func file_meta_proto_init() {
	if File_meta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaInfo); i {
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
		file_meta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaResponse); i {
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
			RawDescriptor: file_meta_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_meta_proto_goTypes,
		DependencyIndexes: file_meta_proto_depIdxs,
		EnumInfos:         file_meta_proto_enumTypes,
		MessageInfos:      file_meta_proto_msgTypes,
	}.Build()
	File_meta_proto = out.File
	file_meta_proto_rawDesc = nil
	file_meta_proto_goTypes = nil
	file_meta_proto_depIdxs = nil
}
