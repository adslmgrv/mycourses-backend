// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: llm-service/proto/llm_service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Role int32

const (
	Role_SYSTEM    Role = 0
	Role_USER      Role = 1
	Role_ASSISTANT Role = 2
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "SYSTEM",
		1: "USER",
		2: "ASSISTANT",
	}
	Role_value = map[string]int32{
		"SYSTEM":    0,
		"USER":      1,
		"ASSISTANT": 2,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_llm_service_proto_llm_service_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_llm_service_proto_llm_service_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_llm_service_proto_llm_service_proto_rawDescGZIP(), []int{0}
}

type Type struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Type) Reset() {
	*x = Type{}
	mi := &file_llm_service_proto_llm_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Type) ProtoMessage() {}

func (x *Type) ProtoReflect() protoreflect.Message {
	mi := &file_llm_service_proto_llm_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Type.ProtoReflect.Descriptor instead.
func (*Type) Descriptor() ([]byte, []int) {
	return file_llm_service_proto_llm_service_proto_rawDescGZIP(), []int{0}
}

func (x *Type) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ToolCall struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Parameters    map[string]*anypb.Any  `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ToolCall) Reset() {
	*x = ToolCall{}
	mi := &file_llm_service_proto_llm_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToolCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolCall) ProtoMessage() {}

func (x *ToolCall) ProtoReflect() protoreflect.Message {
	mi := &file_llm_service_proto_llm_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolCall.ProtoReflect.Descriptor instead.
func (*ToolCall) Descriptor() ([]byte, []int) {
	return file_llm_service_proto_llm_service_proto_rawDescGZIP(), []int{1}
}

func (x *ToolCall) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ToolCall) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ToolCall) GetParameters() map[string]*anypb.Any {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Role          Role                   `protobuf:"varint,1,opt,name=role,proto3,enum=llm_service.v1.Role" json:"role,omitempty"`
	Text          string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	ToolCalls     []*ToolCall            `protobuf:"bytes,3,rep,name=tool_calls,json=toolCalls,proto3" json:"tool_calls,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_llm_service_proto_llm_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_llm_service_proto_llm_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_llm_service_proto_llm_service_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_SYSTEM
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Message) GetToolCalls() []*ToolCall {
	if x != nil {
		return x.ToolCalls
	}
	return nil
}

type ToolParameter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          *Type                  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	IsRequired    bool                   `protobuf:"varint,4,opt,name=is_required,json=isRequired,proto3" json:"is_required,omitempty"`
	Enum          []string               `protobuf:"bytes,5,rep,name=enum,proto3" json:"enum,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ToolParameter) Reset() {
	*x = ToolParameter{}
	mi := &file_llm_service_proto_llm_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToolParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolParameter) ProtoMessage() {}

func (x *ToolParameter) ProtoReflect() protoreflect.Message {
	mi := &file_llm_service_proto_llm_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolParameter.ProtoReflect.Descriptor instead.
func (*ToolParameter) Descriptor() ([]byte, []int) {
	return file_llm_service_proto_llm_service_proto_rawDescGZIP(), []int{3}
}

func (x *ToolParameter) GetType() *Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *ToolParameter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ToolParameter) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ToolParameter) GetIsRequired() bool {
	if x != nil {
		return x.IsRequired
	}
	return false
}

func (x *ToolParameter) GetEnum() []string {
	if x != nil {
		return x.Enum
	}
	return nil
}

type ToolDefinition struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Name             string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description      string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ParametersSchema *Schema                `protobuf:"bytes,3,opt,name=parameters_schema,json=parametersSchema,proto3" json:"parameters_schema,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ToolDefinition) Reset() {
	*x = ToolDefinition{}
	mi := &file_llm_service_proto_llm_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToolDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolDefinition) ProtoMessage() {}

func (x *ToolDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_llm_service_proto_llm_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolDefinition.ProtoReflect.Descriptor instead.
func (*ToolDefinition) Descriptor() ([]byte, []int) {
	return file_llm_service_proto_llm_service_proto_rawDescGZIP(), []int{4}
}

func (x *ToolDefinition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ToolDefinition) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ToolDefinition) GetParametersSchema() *Schema {
	if x != nil {
		return x.ParametersSchema
	}
	return nil
}

type Schema struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          *Type                  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Properties    map[string]*Schema     `protobuf:"bytes,3,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Items         *Schema                `protobuf:"bytes,4,opt,name=items,proto3" json:"items,omitempty"`
	MaxItems      int64                  `protobuf:"varint,5,opt,name=max_items,json=maxItems,proto3" json:"max_items,omitempty"`
	MinItems      int64                  `protobuf:"varint,6,opt,name=min_items,json=minItems,proto3" json:"min_items,omitempty"`
	Required      []string               `protobuf:"bytes,7,rep,name=required,proto3" json:"required,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Schema) Reset() {
	*x = Schema{}
	mi := &file_llm_service_proto_llm_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Schema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema) ProtoMessage() {}

func (x *Schema) ProtoReflect() protoreflect.Message {
	mi := &file_llm_service_proto_llm_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema.ProtoReflect.Descriptor instead.
func (*Schema) Descriptor() ([]byte, []int) {
	return file_llm_service_proto_llm_service_proto_rawDescGZIP(), []int{5}
}

func (x *Schema) GetType() *Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Schema) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Schema) GetProperties() map[string]*Schema {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Schema) GetItems() *Schema {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Schema) GetMaxItems() int64 {
	if x != nil {
		return x.MaxItems
	}
	return 0
}

func (x *Schema) GetMinItems() int64 {
	if x != nil {
		return x.MinItems
	}
	return 0
}

func (x *Schema) GetRequired() []string {
	if x != nil {
		return x.Required
	}
	return nil
}

type GenerateResponseRequest struct {
	state                  protoimpl.MessageState `protogen:"open.v1"`
	ChatHistory            []*Message             `protobuf:"bytes,1,rep,name=chat_history,json=chatHistory,proto3" json:"chat_history,omitempty"`
	Tools                  []*ToolDefinition      `protobuf:"bytes,2,rep,name=tools,proto3" json:"tools,omitempty"`
	StructuredOutputSchema *Schema                `protobuf:"bytes,3,opt,name=structured_output_schema,json=structuredOutputSchema,proto3" json:"structured_output_schema,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *GenerateResponseRequest) Reset() {
	*x = GenerateResponseRequest{}
	mi := &file_llm_service_proto_llm_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateResponseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateResponseRequest) ProtoMessage() {}

func (x *GenerateResponseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_llm_service_proto_llm_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateResponseRequest.ProtoReflect.Descriptor instead.
func (*GenerateResponseRequest) Descriptor() ([]byte, []int) {
	return file_llm_service_proto_llm_service_proto_rawDescGZIP(), []int{6}
}

func (x *GenerateResponseRequest) GetChatHistory() []*Message {
	if x != nil {
		return x.ChatHistory
	}
	return nil
}

func (x *GenerateResponseRequest) GetTools() []*ToolDefinition {
	if x != nil {
		return x.Tools
	}
	return nil
}

func (x *GenerateResponseRequest) GetStructuredOutputSchema() *Schema {
	if x != nil {
		return x.StructuredOutputSchema
	}
	return nil
}

type GenerateResponseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       *Message               `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateResponseResponse) Reset() {
	*x = GenerateResponseResponse{}
	mi := &file_llm_service_proto_llm_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateResponseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateResponseResponse) ProtoMessage() {}

func (x *GenerateResponseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_llm_service_proto_llm_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateResponseResponse.ProtoReflect.Descriptor instead.
func (*GenerateResponseResponse) Descriptor() ([]byte, []int) {
	return file_llm_service_proto_llm_service_proto_rawDescGZIP(), []int{7}
}

func (x *GenerateResponseResponse) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_llm_service_proto_llm_service_proto protoreflect.FileDescriptor

const file_llm_service_proto_llm_service_proto_rawDesc = "" +
	"\n" +
	"#llm-service/proto/llm_service.proto\x12\x0ellm_service.v1\x1a\x19google/protobuf/any.proto\"\x1c\n" +
	"\x04Type\x12\x14\n" +
	"\x05value\x18\x01 \x01(\tR\x05value\"\xcd\x01\n" +
	"\bToolCall\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12H\n" +
	"\n" +
	"parameters\x18\x03 \x03(\v2(.llm_service.v1.ToolCall.ParametersEntryR\n" +
	"parameters\x1aS\n" +
	"\x0fParametersEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12*\n" +
	"\x05value\x18\x02 \x01(\v2\x14.google.protobuf.AnyR\x05value:\x028\x01\"\x80\x01\n" +
	"\aMessage\x12(\n" +
	"\x04role\x18\x01 \x01(\x0e2\x14.llm_service.v1.RoleR\x04role\x12\x12\n" +
	"\x04text\x18\x02 \x01(\tR\x04text\x127\n" +
	"\n" +
	"tool_calls\x18\x03 \x03(\v2\x18.llm_service.v1.ToolCallR\ttoolCalls\"\xa4\x01\n" +
	"\rToolParameter\x12(\n" +
	"\x04type\x18\x01 \x01(\v2\x14.llm_service.v1.TypeR\x04type\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1f\n" +
	"\vis_required\x18\x04 \x01(\bR\n" +
	"isRequired\x12\x12\n" +
	"\x04enum\x18\x05 \x03(\tR\x04enum\"\x8b\x01\n" +
	"\x0eToolDefinition\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12C\n" +
	"\x11parameters_schema\x18\x03 \x01(\v2\x16.llm_service.v1.SchemaR\x10parametersSchema\"\xf7\x02\n" +
	"\x06Schema\x12(\n" +
	"\x04type\x18\x01 \x01(\v2\x14.llm_service.v1.TypeR\x04type\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12F\n" +
	"\n" +
	"properties\x18\x03 \x03(\v2&.llm_service.v1.Schema.PropertiesEntryR\n" +
	"properties\x12,\n" +
	"\x05items\x18\x04 \x01(\v2\x16.llm_service.v1.SchemaR\x05items\x12\x1b\n" +
	"\tmax_items\x18\x05 \x01(\x03R\bmaxItems\x12\x1b\n" +
	"\tmin_items\x18\x06 \x01(\x03R\bminItems\x12\x1a\n" +
	"\brequired\x18\a \x03(\tR\brequired\x1aU\n" +
	"\x0fPropertiesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12,\n" +
	"\x05value\x18\x02 \x01(\v2\x16.llm_service.v1.SchemaR\x05value:\x028\x01\"\xdd\x01\n" +
	"\x17GenerateResponseRequest\x12:\n" +
	"\fchat_history\x18\x01 \x03(\v2\x17.llm_service.v1.MessageR\vchatHistory\x124\n" +
	"\x05tools\x18\x02 \x03(\v2\x1e.llm_service.v1.ToolDefinitionR\x05tools\x12P\n" +
	"\x18structured_output_schema\x18\x03 \x01(\v2\x16.llm_service.v1.SchemaR\x16structuredOutputSchema\"M\n" +
	"\x18GenerateResponseResponse\x121\n" +
	"\amessage\x18\x01 \x01(\v2\x17.llm_service.v1.MessageR\amessage*+\n" +
	"\x04Role\x12\n" +
	"\n" +
	"\x06SYSTEM\x10\x00\x12\b\n" +
	"\x04USER\x10\x01\x12\r\n" +
	"\tASSISTANT\x10\x022s\n" +
	"\n" +
	"LLMService\x12e\n" +
	"\x10GenerateResponse\x12'.llm_service.v1.GenerateResponseRequest\x1a(.llm_service.v1.GenerateResponseResponseB\x13Z\x11internal/proto/v1b\x06proto3"

var (
	file_llm_service_proto_llm_service_proto_rawDescOnce sync.Once
	file_llm_service_proto_llm_service_proto_rawDescData []byte
)

func file_llm_service_proto_llm_service_proto_rawDescGZIP() []byte {
	file_llm_service_proto_llm_service_proto_rawDescOnce.Do(func() {
		file_llm_service_proto_llm_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_llm_service_proto_llm_service_proto_rawDesc), len(file_llm_service_proto_llm_service_proto_rawDesc)))
	})
	return file_llm_service_proto_llm_service_proto_rawDescData
}

var file_llm_service_proto_llm_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_llm_service_proto_llm_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_llm_service_proto_llm_service_proto_goTypes = []any{
	(Role)(0),                        // 0: llm_service.v1.Role
	(*Type)(nil),                     // 1: llm_service.v1.Type
	(*ToolCall)(nil),                 // 2: llm_service.v1.ToolCall
	(*Message)(nil),                  // 3: llm_service.v1.Message
	(*ToolParameter)(nil),            // 4: llm_service.v1.ToolParameter
	(*ToolDefinition)(nil),           // 5: llm_service.v1.ToolDefinition
	(*Schema)(nil),                   // 6: llm_service.v1.Schema
	(*GenerateResponseRequest)(nil),  // 7: llm_service.v1.GenerateResponseRequest
	(*GenerateResponseResponse)(nil), // 8: llm_service.v1.GenerateResponseResponse
	nil,                              // 9: llm_service.v1.ToolCall.ParametersEntry
	nil,                              // 10: llm_service.v1.Schema.PropertiesEntry
	(*anypb.Any)(nil),                // 11: google.protobuf.Any
}
var file_llm_service_proto_llm_service_proto_depIdxs = []int32{
	9,  // 0: llm_service.v1.ToolCall.parameters:type_name -> llm_service.v1.ToolCall.ParametersEntry
	0,  // 1: llm_service.v1.Message.role:type_name -> llm_service.v1.Role
	2,  // 2: llm_service.v1.Message.tool_calls:type_name -> llm_service.v1.ToolCall
	1,  // 3: llm_service.v1.ToolParameter.type:type_name -> llm_service.v1.Type
	6,  // 4: llm_service.v1.ToolDefinition.parameters_schema:type_name -> llm_service.v1.Schema
	1,  // 5: llm_service.v1.Schema.type:type_name -> llm_service.v1.Type
	10, // 6: llm_service.v1.Schema.properties:type_name -> llm_service.v1.Schema.PropertiesEntry
	6,  // 7: llm_service.v1.Schema.items:type_name -> llm_service.v1.Schema
	3,  // 8: llm_service.v1.GenerateResponseRequest.chat_history:type_name -> llm_service.v1.Message
	5,  // 9: llm_service.v1.GenerateResponseRequest.tools:type_name -> llm_service.v1.ToolDefinition
	6,  // 10: llm_service.v1.GenerateResponseRequest.structured_output_schema:type_name -> llm_service.v1.Schema
	3,  // 11: llm_service.v1.GenerateResponseResponse.message:type_name -> llm_service.v1.Message
	11, // 12: llm_service.v1.ToolCall.ParametersEntry.value:type_name -> google.protobuf.Any
	6,  // 13: llm_service.v1.Schema.PropertiesEntry.value:type_name -> llm_service.v1.Schema
	7,  // 14: llm_service.v1.LLMService.GenerateResponse:input_type -> llm_service.v1.GenerateResponseRequest
	8,  // 15: llm_service.v1.LLMService.GenerateResponse:output_type -> llm_service.v1.GenerateResponseResponse
	15, // [15:16] is the sub-list for method output_type
	14, // [14:15] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_llm_service_proto_llm_service_proto_init() }
func file_llm_service_proto_llm_service_proto_init() {
	if File_llm_service_proto_llm_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_llm_service_proto_llm_service_proto_rawDesc), len(file_llm_service_proto_llm_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_llm_service_proto_llm_service_proto_goTypes,
		DependencyIndexes: file_llm_service_proto_llm_service_proto_depIdxs,
		EnumInfos:         file_llm_service_proto_llm_service_proto_enumTypes,
		MessageInfos:      file_llm_service_proto_llm_service_proto_msgTypes,
	}.Build()
	File_llm_service_proto_llm_service_proto = out.File
	file_llm_service_proto_llm_service_proto_goTypes = nil
	file_llm_service_proto_llm_service_proto_depIdxs = nil
}
