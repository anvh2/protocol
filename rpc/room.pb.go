// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.0
// source: rpc/room.proto

package rpc

import (
	livekit "github.com/livekit/protocol/livekit"
	_ "github.com/livekit/psrpc/protoc-gen-psrpc/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_rpc_room_proto protoreflect.FileDescriptor

var file_rpc_room_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6c, 0x69, 0x76, 0x65,
	0x6b, 0x69, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9f,
	0x02, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x5d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x1a, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16,
	0xb2, 0x89, 0x01, 0x12, 0x10, 0x01, 0x1a, 0x0e, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x04,
	0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x12, 0x57, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x18, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0xb2, 0x89, 0x01, 0x12, 0x10, 0x01, 0x1a,
	0x0e, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x12,
	0x5f, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x22, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6c, 0x69, 0x76, 0x65,
	0x6b, 0x69, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x22, 0x16, 0xb2, 0x89, 0x01, 0x12, 0x10, 0x01,
	0x1a, 0x0e, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01,
	0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_rpc_room_proto_goTypes = []any{
	(*livekit.DeleteRoomRequest)(nil),         // 0: livekit.DeleteRoomRequest
	(*livekit.SendDataRequest)(nil),           // 1: livekit.SendDataRequest
	(*livekit.UpdateRoomMetadataRequest)(nil), // 2: livekit.UpdateRoomMetadataRequest
	(*livekit.DeleteRoomResponse)(nil),        // 3: livekit.DeleteRoomResponse
	(*livekit.SendDataResponse)(nil),          // 4: livekit.SendDataResponse
	(*livekit.Room)(nil),                      // 5: livekit.Room
}
var file_rpc_room_proto_depIdxs = []int32{
	0, // 0: rpc.Room.DeleteRoom:input_type -> livekit.DeleteRoomRequest
	1, // 1: rpc.Room.SendData:input_type -> livekit.SendDataRequest
	2, // 2: rpc.Room.UpdateRoomMetadata:input_type -> livekit.UpdateRoomMetadataRequest
	3, // 3: rpc.Room.DeleteRoom:output_type -> livekit.DeleteRoomResponse
	4, // 4: rpc.Room.SendData:output_type -> livekit.SendDataResponse
	5, // 5: rpc.Room.UpdateRoomMetadata:output_type -> livekit.Room
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_room_proto_init() }
func file_rpc_room_proto_init() {
	if File_rpc_room_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_room_proto_goTypes,
		DependencyIndexes: file_rpc_room_proto_depIdxs,
	}.Build()
	File_rpc_room_proto = out.File
	file_rpc_room_proto_rawDesc = nil
	file_rpc_room_proto_goTypes = nil
	file_rpc_room_proto_depIdxs = nil
}
