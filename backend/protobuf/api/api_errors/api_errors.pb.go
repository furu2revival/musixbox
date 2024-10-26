// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: api/api_errors/api_errors.proto

package api_errors

import (
	code "google.golang.org/genproto/googleapis/rpc/code"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorHandlingType int32

const (
	ErrorHandlingType_ERROR_HANDLING_TYPE_UNSPECIFIED ErrorHandlingType = 0
	// サバクラ間の前提条件が満たされていない場合に発生する致命的なエラー。
	// 緊急メンテナンスで修正しなければならない。
	ErrorHandlingType_ERROR_HANDLING_TYPE_FATAL ErrorHandlingType = 1
	// サーバ側の一時的な不具合により発生したエラー。
	// リトライすることにより解決する可能性がある。
	ErrorHandlingType_ERROR_HANDLING_TYPE_TEMPORARY ErrorHandlingType = 2
	// クライアント側の不具合により発生したエラー。
	// 1つ前の画面に遷移したりタイトルに戻して前提条件を変えることにより、解決する可能性がある。
	ErrorHandlingType_ERROR_HANDLING_TYPE_RECOVERABLE ErrorHandlingType = 3
)

// Enum value maps for ErrorHandlingType.
var (
	ErrorHandlingType_name = map[int32]string{
		0: "ERROR_HANDLING_TYPE_UNSPECIFIED",
		1: "ERROR_HANDLING_TYPE_FATAL",
		2: "ERROR_HANDLING_TYPE_TEMPORARY",
		3: "ERROR_HANDLING_TYPE_RECOVERABLE",
	}
	ErrorHandlingType_value = map[string]int32{
		"ERROR_HANDLING_TYPE_UNSPECIFIED": 0,
		"ERROR_HANDLING_TYPE_FATAL":       1,
		"ERROR_HANDLING_TYPE_TEMPORARY":   2,
		"ERROR_HANDLING_TYPE_RECOVERABLE": 3,
	}
)

func (x ErrorHandlingType) Enum() *ErrorHandlingType {
	p := new(ErrorHandlingType)
	*p = x
	return p
}

func (x ErrorHandlingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorHandlingType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_api_errors_api_errors_proto_enumTypes[0].Descriptor()
}

func (ErrorHandlingType) Type() protoreflect.EnumType {
	return &file_api_api_errors_api_errors_proto_enumTypes[0]
}

func (x ErrorHandlingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorHandlingType.Descriptor instead.
func (ErrorHandlingType) EnumDescriptor() ([]byte, []int) {
	return file_api_api_errors_api_errors_proto_rawDescGZIP(), []int{0}
}

type ErrorSeverity int32

const (
	ErrorSeverity_ERROR_SEVERITY_UNSPECIFIED ErrorSeverity = 0
	ErrorSeverity_ERROR_SEVERITY_DEBUG       ErrorSeverity = 1
	ErrorSeverity_ERROR_SEVERITY_INFO        ErrorSeverity = 2
	ErrorSeverity_ERROR_SEVERITY_NOTICE      ErrorSeverity = 3
	ErrorSeverity_ERROR_SEVERITY_WARNING     ErrorSeverity = 4
	ErrorSeverity_ERROR_SEVERITY_ERROR       ErrorSeverity = 5
	ErrorSeverity_ERROR_SEVERITY_CRITICAL    ErrorSeverity = 6
	ErrorSeverity_ERROR_SEVERITY_ALERT       ErrorSeverity = 7
	ErrorSeverity_ERROR_SEVERITY_EMERGENCY   ErrorSeverity = 8
)

// Enum value maps for ErrorSeverity.
var (
	ErrorSeverity_name = map[int32]string{
		0: "ERROR_SEVERITY_UNSPECIFIED",
		1: "ERROR_SEVERITY_DEBUG",
		2: "ERROR_SEVERITY_INFO",
		3: "ERROR_SEVERITY_NOTICE",
		4: "ERROR_SEVERITY_WARNING",
		5: "ERROR_SEVERITY_ERROR",
		6: "ERROR_SEVERITY_CRITICAL",
		7: "ERROR_SEVERITY_ALERT",
		8: "ERROR_SEVERITY_EMERGENCY",
	}
	ErrorSeverity_value = map[string]int32{
		"ERROR_SEVERITY_UNSPECIFIED": 0,
		"ERROR_SEVERITY_DEBUG":       1,
		"ERROR_SEVERITY_INFO":        2,
		"ERROR_SEVERITY_NOTICE":      3,
		"ERROR_SEVERITY_WARNING":     4,
		"ERROR_SEVERITY_ERROR":       5,
		"ERROR_SEVERITY_CRITICAL":    6,
		"ERROR_SEVERITY_ALERT":       7,
		"ERROR_SEVERITY_EMERGENCY":   8,
	}
)

func (x ErrorSeverity) Enum() *ErrorSeverity {
	p := new(ErrorSeverity)
	*p = x
	return p
}

func (x ErrorSeverity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorSeverity) Descriptor() protoreflect.EnumDescriptor {
	return file_api_api_errors_api_errors_proto_enumTypes[1].Descriptor()
}

func (ErrorSeverity) Type() protoreflect.EnumType {
	return &file_api_api_errors_api_errors_proto_enumTypes[1]
}

func (x ErrorSeverity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorSeverity.Descriptor instead.
func (ErrorSeverity) EnumDescriptor() ([]byte, []int) {
	return file_api_api_errors_api_errors_proto_rawDescGZIP(), []int{1}
}

type ErrorCode_Common int32

const (
	ErrorCode_COMMON_UNSPECIFIED               ErrorCode_Common = 0
	ErrorCode_COMMON_UNKNOWN                   ErrorCode_Common = 1000
	ErrorCode_COMMON_INVALID_SESSION           ErrorCode_Common = 1001
	ErrorCode_COMMON_INVALID_USER_AVAILABILITY ErrorCode_Common = 1002
)

// Enum value maps for ErrorCode_Common.
var (
	ErrorCode_Common_name = map[int32]string{
		0:    "COMMON_UNSPECIFIED",
		1000: "COMMON_UNKNOWN",
		1001: "COMMON_INVALID_SESSION",
		1002: "COMMON_INVALID_USER_AVAILABILITY",
	}
	ErrorCode_Common_value = map[string]int32{
		"COMMON_UNSPECIFIED":               0,
		"COMMON_UNKNOWN":                   1000,
		"COMMON_INVALID_SESSION":           1001,
		"COMMON_INVALID_USER_AVAILABILITY": 1002,
	}
)

func (x ErrorCode_Common) Enum() *ErrorCode_Common {
	p := new(ErrorCode_Common)
	*p = x
	return p
}

func (x ErrorCode_Common) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode_Common) Descriptor() protoreflect.EnumDescriptor {
	return file_api_api_errors_api_errors_proto_enumTypes[2].Descriptor()
}

func (ErrorCode_Common) Type() protoreflect.EnumType {
	return &file_api_api_errors_api_errors_proto_enumTypes[2]
}

func (x ErrorCode_Common) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode_Common.Descriptor instead.
func (ErrorCode_Common) EnumDescriptor() ([]byte, []int) {
	return file_api_api_errors_api_errors_proto_rawDescGZIP(), []int{0, 0}
}

type ErrorCode_Method int32

const (
	ErrorCode_METHOD_UNSPECIFIED ErrorCode_Method = 0
	// 入力値のバリデーションエラー。
	// エラーを表示した後、アプリを終了させる。
	ErrorCode_METHOD_ILLEGAL_ARGUMENT ErrorCode_Method = 2000
	// 要求の操作を行うための前提条件を満たしていない。
	// 前提となる状態に遷移し、ユーザに再操作を促す必要がある。
	ErrorCode_METHOD_RESOURCE_NOT_FOUND ErrorCode_Method = 2001
	// 要求の操作を行うための前提条件を満たしていない。
	// 前提となる状態に遷移し、ユーザに再操作を促す必要がある。
	ErrorCode_METHOD_RESOURCE_CONFLICT ErrorCode_Method = 2002
	// 要求の操作を行うための前提条件を満たしていない。
	// 前提となる状態に遷移し、ユーザに再操作を促す必要がある。
	ErrorCode_METHOD_RESOURCE_INSUFFICIENT ErrorCode_Method = 2003
)

// Enum value maps for ErrorCode_Method.
var (
	ErrorCode_Method_name = map[int32]string{
		0:    "METHOD_UNSPECIFIED",
		2000: "METHOD_ILLEGAL_ARGUMENT",
		2001: "METHOD_RESOURCE_NOT_FOUND",
		2002: "METHOD_RESOURCE_CONFLICT",
		2003: "METHOD_RESOURCE_INSUFFICIENT",
	}
	ErrorCode_Method_value = map[string]int32{
		"METHOD_UNSPECIFIED":           0,
		"METHOD_ILLEGAL_ARGUMENT":      2000,
		"METHOD_RESOURCE_NOT_FOUND":    2001,
		"METHOD_RESOURCE_CONFLICT":     2002,
		"METHOD_RESOURCE_INSUFFICIENT": 2003,
	}
)

func (x ErrorCode_Method) Enum() *ErrorCode_Method {
	p := new(ErrorCode_Method)
	*p = x
	return p
}

func (x ErrorCode_Method) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode_Method) Descriptor() protoreflect.EnumDescriptor {
	return file_api_api_errors_api_errors_proto_enumTypes[3].Descriptor()
}

func (ErrorCode_Method) Type() protoreflect.EnumType {
	return &file_api_api_errors_api_errors_proto_enumTypes[3]
}

func (x ErrorCode_Method) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode_Method.Descriptor instead.
func (ErrorCode_Method) EnumDescriptor() ([]byte, []int) {
	return file_api_api_errors_api_errors_proto_rawDescGZIP(), []int{0, 1}
}

type ErrorCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ErrorCode) Reset() {
	*x = ErrorCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_errors_api_errors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorCode) ProtoMessage() {}

func (x *ErrorCode) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_errors_api_errors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorCode.ProtoReflect.Descriptor instead.
func (*ErrorCode) Descriptor() ([]byte, []int) {
	return file_api_api_errors_api_errors_proto_rawDescGZIP(), []int{0}
}

type ErrorDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode         int64             `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorHandlingType ErrorHandlingType `protobuf:"varint,2,opt,name=error_handling_type,json=errorHandlingType,proto3,enum=api.api_errors.ErrorHandlingType" json:"error_handling_type,omitempty"`
}

func (x *ErrorDetail) Reset() {
	*x = ErrorDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_errors_api_errors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetail) ProtoMessage() {}

func (x *ErrorDetail) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_errors_api_errors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetail.ProtoReflect.Descriptor instead.
func (*ErrorDetail) Descriptor() ([]byte, []int) {
	return file_api_api_errors_api_errors_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorDetail) GetErrorCode() int64 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *ErrorDetail) GetErrorHandlingType() ErrorHandlingType {
	if x != nil {
		return x.ErrorHandlingType
	}
	return ErrorHandlingType_ERROR_HANDLING_TYPE_UNSPECIFIED
}

var file_api_api_errors_api_errors_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*code.Code)(nil),
		Field:         50100,
		Name:          "api.api_errors.grpc_code",
		Tag:           "varint,50100,opt,name=grpc_code,enum=google.rpc.Code",
		Filename:      "api/api_errors/api_errors.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*ErrorHandlingType)(nil),
		Field:         50101,
		Name:          "api.api_errors.error_handling_type",
		Tag:           "varint,50101,opt,name=error_handling_type,enum=api.api_errors.ErrorHandlingType",
		Filename:      "api/api_errors/api_errors.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional google.rpc.Code grpc_code = 50100;
	E_GrpcCode = &file_api_api_errors_api_errors_proto_extTypes[0]
	// optional api.api_errors.ErrorHandlingType error_handling_type = 50101;
	E_ErrorHandlingType = &file_api_api_errors_api_errors_proto_extTypes[1]
)

var File_api_api_errors_api_errors_proto protoreflect.FileDescriptor

var file_api_api_errors_api_errors_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x03, 0x0a, 0x09, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x0e, 0x43,
	0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0xe8, 0x07,
	0x1a, 0x08, 0xa0, 0xbb, 0x18, 0x02, 0xa8, 0xbb, 0x18, 0x02, 0x12, 0x25, 0x0a, 0x16, 0x43, 0x4f,
	0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x45, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0xe9, 0x07, 0x1a, 0x08, 0xa0, 0xbb, 0x18, 0x10, 0xa8, 0xbb, 0x18,
	0x03, 0x12, 0x2f, 0x0a, 0x20, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0xea, 0x07, 0x1a, 0x08, 0xa0, 0xbb, 0x18, 0x07, 0xa8, 0xbb,
	0x18, 0x03, 0x22, 0x05, 0x08, 0x01, 0x10, 0xe7, 0x07, 0x22, 0x09, 0x08, 0xd0, 0x0f, 0x10, 0xff,
	0xff, 0xff, 0xff, 0x07, 0x22, 0xda, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x16, 0x0a, 0x12, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x17, 0x4d, 0x45, 0x54, 0x48, 0x4f,
	0x44, 0x5f, 0x49, 0x4c, 0x4c, 0x45, 0x47, 0x41, 0x4c, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45,
	0x4e, 0x54, 0x10, 0xd0, 0x0f, 0x1a, 0x08, 0xa0, 0xbb, 0x18, 0x03, 0xa8, 0xbb, 0x18, 0x01, 0x12,
	0x28, 0x0a, 0x19, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xd1, 0x0f, 0x1a,
	0x08, 0xa0, 0xbb, 0x18, 0x05, 0xa8, 0xbb, 0x18, 0x03, 0x12, 0x27, 0x0a, 0x18, 0x4d, 0x45, 0x54,
	0x48, 0x4f, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x43, 0x4f, 0x4e,
	0x46, 0x4c, 0x49, 0x43, 0x54, 0x10, 0xd2, 0x0f, 0x1a, 0x08, 0xa0, 0xbb, 0x18, 0x06, 0xa8, 0xbb,
	0x18, 0x03, 0x12, 0x2b, 0x0a, 0x1c, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x52, 0x45, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x4e, 0x53, 0x55, 0x46, 0x46, 0x49, 0x43, 0x49, 0x45,
	0x4e, 0x54, 0x10, 0xd3, 0x0f, 0x1a, 0x08, 0xa0, 0xbb, 0x18, 0x08, 0xa8, 0xbb, 0x18, 0x03, 0x22,
	0x05, 0x08, 0x01, 0x10, 0xcf, 0x0f, 0x22, 0x09, 0x08, 0xb8, 0x17, 0x10, 0xff, 0xff, 0xff, 0xff,
	0x07, 0x22, 0x7f, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x51, 0x0a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e,
	0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x11, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x2a, 0x9f, 0x01, 0x0a, 0x11, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a,
	0x19, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x49, 0x4e, 0x47, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x4f, 0x52, 0x41, 0x52, 0x59, 0x10, 0x02, 0x12,
	0x23, 0x0a, 0x1f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x49, 0x4e,
	0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x03, 0x2a, 0x88, 0x02, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x01,
	0x12, 0x17, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x4f, 0x54, 0x49,
	0x43, 0x45, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x53, 0x45,
	0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x04,
	0x12, 0x18, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x52, 0x49,
	0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x06, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x10,
	0x07, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52,
	0x49, 0x54, 0x59, 0x5f, 0x45, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x08, 0x3a,
	0x55, 0x0a, 0x09, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xb4, 0x87, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x67, 0x72, 0x70, 0x63, 0x43,
	0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x3a, 0x79, 0x0a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xb5, 0x87, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x70, 0x69, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x11, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01,
	0x01, 0x42, 0xb3, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x0e, 0x41, 0x70, 0x69, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x75, 0x72, 0x75, 0x32, 0x72, 0x65, 0x76, 0x69,
	0x76, 0x61, 0x6c, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x62, 0x6f, 0x78, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x41, 0x58, 0xaa, 0x02, 0x0d, 0x41, 0x70, 0x69,
	0x2e, 0x41, 0x70, 0x69, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x0d, 0x41, 0x70, 0x69,
	0x5c, 0x41, 0x70, 0x69, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xe2, 0x02, 0x19, 0x41, 0x70, 0x69,
	0x5c, 0x41, 0x70, 0x69, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_api_errors_api_errors_proto_rawDescOnce sync.Once
	file_api_api_errors_api_errors_proto_rawDescData = file_api_api_errors_api_errors_proto_rawDesc
)

func file_api_api_errors_api_errors_proto_rawDescGZIP() []byte {
	file_api_api_errors_api_errors_proto_rawDescOnce.Do(func() {
		file_api_api_errors_api_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_api_errors_api_errors_proto_rawDescData)
	})
	return file_api_api_errors_api_errors_proto_rawDescData
}

var file_api_api_errors_api_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_api_api_errors_api_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_api_errors_api_errors_proto_goTypes = []any{
	(ErrorHandlingType)(0),                // 0: api.api_errors.ErrorHandlingType
	(ErrorSeverity)(0),                    // 1: api.api_errors.ErrorSeverity
	(ErrorCode_Common)(0),                 // 2: api.api_errors.ErrorCode.Common
	(ErrorCode_Method)(0),                 // 3: api.api_errors.ErrorCode.Method
	(*ErrorCode)(nil),                     // 4: api.api_errors.ErrorCode
	(*ErrorDetail)(nil),                   // 5: api.api_errors.ErrorDetail
	(*descriptorpb.EnumValueOptions)(nil), // 6: google.protobuf.EnumValueOptions
	(code.Code)(0),                        // 7: google.rpc.Code
}
var file_api_api_errors_api_errors_proto_depIdxs = []int32{
	0, // 0: api.api_errors.ErrorDetail.error_handling_type:type_name -> api.api_errors.ErrorHandlingType
	6, // 1: api.api_errors.grpc_code:extendee -> google.protobuf.EnumValueOptions
	6, // 2: api.api_errors.error_handling_type:extendee -> google.protobuf.EnumValueOptions
	7, // 3: api.api_errors.grpc_code:type_name -> google.rpc.Code
	0, // 4: api.api_errors.error_handling_type:type_name -> api.api_errors.ErrorHandlingType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	3, // [3:5] is the sub-list for extension type_name
	1, // [1:3] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_api_errors_api_errors_proto_init() }
func file_api_api_errors_api_errors_proto_init() {
	if File_api_api_errors_api_errors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_api_errors_api_errors_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ErrorCode); i {
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
		file_api_api_errors_api_errors_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ErrorDetail); i {
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
			RawDescriptor: file_api_api_errors_api_errors_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_api_api_errors_api_errors_proto_goTypes,
		DependencyIndexes: file_api_api_errors_api_errors_proto_depIdxs,
		EnumInfos:         file_api_api_errors_api_errors_proto_enumTypes,
		MessageInfos:      file_api_api_errors_api_errors_proto_msgTypes,
		ExtensionInfos:    file_api_api_errors_api_errors_proto_extTypes,
	}.Build()
	File_api_api_errors_api_errors_proto = out.File
	file_api_api_errors_api_errors_proto_rawDesc = nil
	file_api_api_errors_api_errors_proto_goTypes = nil
	file_api_api_errors_api_errors_proto_depIdxs = nil
}
