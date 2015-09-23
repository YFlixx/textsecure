// Code generated by protoc-gen-go.
// source: TextSecure.proto
// DO NOT EDIT!

package textsecure

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Envelope_Type int32

const (
	Envelope_UNKNOWN       Envelope_Type = 0
	Envelope_CIPHERTEXT    Envelope_Type = 1
	Envelope_KEY_EXCHANGE  Envelope_Type = 2
	Envelope_PREKEY_BUNDLE Envelope_Type = 3
	Envelope_RECEIPT       Envelope_Type = 5
)

var Envelope_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "CIPHERTEXT",
	2: "KEY_EXCHANGE",
	3: "PREKEY_BUNDLE",
	5: "RECEIPT",
}
var Envelope_Type_value = map[string]int32{
	"UNKNOWN":       0,
	"CIPHERTEXT":    1,
	"KEY_EXCHANGE":  2,
	"PREKEY_BUNDLE": 3,
	"RECEIPT":       5,
}

func (x Envelope_Type) Enum() *Envelope_Type {
	p := new(Envelope_Type)
	*p = x
	return p
}
func (x Envelope_Type) String() string {
	return proto.EnumName(Envelope_Type_name, int32(x))
}
func (x Envelope_Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Envelope_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Envelope_Type_value, data, "Envelope_Type")
	if err != nil {
		return err
	}
	*x = Envelope_Type(value)
	return nil
}

type DataMessage_Flags int32

const (
	DataMessage_END_SESSION DataMessage_Flags = 1
)

var DataMessage_Flags_name = map[int32]string{
	1: "END_SESSION",
}
var DataMessage_Flags_value = map[string]int32{
	"END_SESSION": 1,
}

func (x DataMessage_Flags) Enum() *DataMessage_Flags {
	p := new(DataMessage_Flags)
	*p = x
	return p
}
func (x DataMessage_Flags) String() string {
	return proto.EnumName(DataMessage_Flags_name, int32(x))
}
func (x DataMessage_Flags) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *DataMessage_Flags) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DataMessage_Flags_value, data, "DataMessage_Flags")
	if err != nil {
		return err
	}
	*x = DataMessage_Flags(value)
	return nil
}

type SyncMessage_Request_Type int32

const (
	SyncMessage_Request_UNKNOWN  SyncMessage_Request_Type = 0
	SyncMessage_Request_CONTACTS SyncMessage_Request_Type = 1
	SyncMessage_Request_GROUPS   SyncMessage_Request_Type = 2
)

var SyncMessage_Request_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "CONTACTS",
	2: "GROUPS",
}
var SyncMessage_Request_Type_value = map[string]int32{
	"UNKNOWN":  0,
	"CONTACTS": 1,
	"GROUPS":   2,
}

func (x SyncMessage_Request_Type) Enum() *SyncMessage_Request_Type {
	p := new(SyncMessage_Request_Type)
	*p = x
	return p
}
func (x SyncMessage_Request_Type) String() string {
	return proto.EnumName(SyncMessage_Request_Type_name, int32(x))
}
func (x SyncMessage_Request_Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *SyncMessage_Request_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SyncMessage_Request_Type_value, data, "SyncMessage_Request_Type")
	if err != nil {
		return err
	}
	*x = SyncMessage_Request_Type(value)
	return nil
}

type GroupContext_Type int32

const (
	GroupContext_UNKNOWN GroupContext_Type = 0
	GroupContext_UPDATE  GroupContext_Type = 1
	GroupContext_DELIVER GroupContext_Type = 2
	GroupContext_QUIT    GroupContext_Type = 3
)

var GroupContext_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "UPDATE",
	2: "DELIVER",
	3: "QUIT",
}
var GroupContext_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"UPDATE":  1,
	"DELIVER": 2,
	"QUIT":    3,
}

func (x GroupContext_Type) Enum() *GroupContext_Type {
	p := new(GroupContext_Type)
	*p = x
	return p
}
func (x GroupContext_Type) String() string {
	return proto.EnumName(GroupContext_Type_name, int32(x))
}
func (x GroupContext_Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *GroupContext_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GroupContext_Type_value, data, "GroupContext_Type")
	if err != nil {
		return err
	}
	*x = GroupContext_Type(value)
	return nil
}

type Envelope struct {
	Type             *Envelope_Type `protobuf:"varint,1,opt,name=type,enum=textsecure.Envelope_Type" json:"type,omitempty"`
	Source           *string        `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	SourceDevice     *uint32        `protobuf:"varint,7,opt,name=sourceDevice" json:"sourceDevice,omitempty"`
	Relay            *string        `protobuf:"bytes,3,opt,name=relay" json:"relay,omitempty"`
	Timestamp        *uint64        `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
	LegacyMessage    []byte         `protobuf:"bytes,6,opt,name=legacyMessage" json:"legacyMessage,omitempty"`
	Content          []byte         `protobuf:"bytes,8,opt,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}

func (m *Envelope) GetType() Envelope_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Envelope) GetSource() string {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return ""
}

func (m *Envelope) GetSourceDevice() uint32 {
	if m != nil && m.SourceDevice != nil {
		return *m.SourceDevice
	}
	return 0
}

func (m *Envelope) GetRelay() string {
	if m != nil && m.Relay != nil {
		return *m.Relay
	}
	return ""
}

func (m *Envelope) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Envelope) GetLegacyMessage() []byte {
	if m != nil {
		return m.LegacyMessage
	}
	return nil
}

func (m *Envelope) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type Content struct {
	DataMessage      *DataMessage `protobuf:"bytes,1,opt,name=dataMessage" json:"dataMessage,omitempty"`
	SyncMessage      *SyncMessage `protobuf:"bytes,2,opt,name=syncMessage" json:"syncMessage,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}

func (m *Content) GetDataMessage() *DataMessage {
	if m != nil {
		return m.DataMessage
	}
	return nil
}

func (m *Content) GetSyncMessage() *SyncMessage {
	if m != nil {
		return m.SyncMessage
	}
	return nil
}

type DataMessage struct {
	Body             *string              `protobuf:"bytes,1,opt,name=body" json:"body,omitempty"`
	Attachments      []*AttachmentPointer `protobuf:"bytes,2,rep,name=attachments" json:"attachments,omitempty"`
	Group            *GroupContext        `protobuf:"bytes,3,opt,name=group" json:"group,omitempty"`
	Flags            *uint32              `protobuf:"varint,4,opt,name=flags" json:"flags,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DataMessage) Reset()         { *m = DataMessage{} }
func (m *DataMessage) String() string { return proto.CompactTextString(m) }
func (*DataMessage) ProtoMessage()    {}

func (m *DataMessage) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func (m *DataMessage) GetAttachments() []*AttachmentPointer {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *DataMessage) GetGroup() *GroupContext {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *DataMessage) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

type SyncMessage struct {
	Sent             *SyncMessage_Sent     `protobuf:"bytes,1,opt,name=sent" json:"sent,omitempty"`
	Contacts         *SyncMessage_Contacts `protobuf:"bytes,2,opt,name=contacts" json:"contacts,omitempty"`
	Groups           *SyncMessage_Groups   `protobuf:"bytes,3,opt,name=groups" json:"groups,omitempty"`
	Request          *SyncMessage_Request  `protobuf:"bytes,4,opt,name=request" json:"request,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *SyncMessage) Reset()         { *m = SyncMessage{} }
func (m *SyncMessage) String() string { return proto.CompactTextString(m) }
func (*SyncMessage) ProtoMessage()    {}

func (m *SyncMessage) GetSent() *SyncMessage_Sent {
	if m != nil {
		return m.Sent
	}
	return nil
}

func (m *SyncMessage) GetContacts() *SyncMessage_Contacts {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func (m *SyncMessage) GetGroups() *SyncMessage_Groups {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *SyncMessage) GetRequest() *SyncMessage_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

type SyncMessage_Sent struct {
	Destination      *string      `protobuf:"bytes,1,opt,name=destination" json:"destination,omitempty"`
	Timestamp        *uint64      `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Message          *DataMessage `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SyncMessage_Sent) Reset()         { *m = SyncMessage_Sent{} }
func (m *SyncMessage_Sent) String() string { return proto.CompactTextString(m) }
func (*SyncMessage_Sent) ProtoMessage()    {}

func (m *SyncMessage_Sent) GetDestination() string {
	if m != nil && m.Destination != nil {
		return *m.Destination
	}
	return ""
}

func (m *SyncMessage_Sent) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *SyncMessage_Sent) GetMessage() *DataMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

type SyncMessage_Contacts struct {
	Blob             *AttachmentPointer `protobuf:"bytes,1,opt,name=blob" json:"blob,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *SyncMessage_Contacts) Reset()         { *m = SyncMessage_Contacts{} }
func (m *SyncMessage_Contacts) String() string { return proto.CompactTextString(m) }
func (*SyncMessage_Contacts) ProtoMessage()    {}

func (m *SyncMessage_Contacts) GetBlob() *AttachmentPointer {
	if m != nil {
		return m.Blob
	}
	return nil
}

type SyncMessage_Groups struct {
	Blob             *AttachmentPointer `protobuf:"bytes,1,opt,name=blob" json:"blob,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *SyncMessage_Groups) Reset()         { *m = SyncMessage_Groups{} }
func (m *SyncMessage_Groups) String() string { return proto.CompactTextString(m) }
func (*SyncMessage_Groups) ProtoMessage()    {}

func (m *SyncMessage_Groups) GetBlob() *AttachmentPointer {
	if m != nil {
		return m.Blob
	}
	return nil
}

type SyncMessage_Request struct {
	Type             *SyncMessage_Request_Type `protobuf:"varint,1,opt,name=type,enum=textsecure.SyncMessage_Request_Type" json:"type,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *SyncMessage_Request) Reset()         { *m = SyncMessage_Request{} }
func (m *SyncMessage_Request) String() string { return proto.CompactTextString(m) }
func (*SyncMessage_Request) ProtoMessage()    {}

func (m *SyncMessage_Request) GetType() SyncMessage_Request_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

type AttachmentPointer struct {
	Id               *uint64 `protobuf:"fixed64,1,opt,name=id" json:"id,omitempty"`
	ContentType      *string `protobuf:"bytes,2,opt,name=contentType" json:"contentType,omitempty"`
	Key              []byte  `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Size             *uint32 `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
	Thumbnail        []byte  `protobuf:"bytes,5,opt,name=thumbnail" json:"thumbnail,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AttachmentPointer) Reset()         { *m = AttachmentPointer{} }
func (m *AttachmentPointer) String() string { return proto.CompactTextString(m) }
func (*AttachmentPointer) ProtoMessage()    {}

func (m *AttachmentPointer) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *AttachmentPointer) GetContentType() string {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ""
}

func (m *AttachmentPointer) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *AttachmentPointer) GetSize() uint32 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *AttachmentPointer) GetThumbnail() []byte {
	if m != nil {
		return m.Thumbnail
	}
	return nil
}

type GroupContext struct {
	Id               []byte             `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type             *GroupContext_Type `protobuf:"varint,2,opt,name=type,enum=textsecure.GroupContext_Type" json:"type,omitempty"`
	Name             *string            `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Members          []string           `protobuf:"bytes,4,rep,name=members" json:"members,omitempty"`
	Avatar           *AttachmentPointer `protobuf:"bytes,5,opt,name=avatar" json:"avatar,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *GroupContext) Reset()         { *m = GroupContext{} }
func (m *GroupContext) String() string { return proto.CompactTextString(m) }
func (*GroupContext) ProtoMessage()    {}

func (m *GroupContext) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GroupContext) GetType() GroupContext_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *GroupContext) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *GroupContext) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *GroupContext) GetAvatar() *AttachmentPointer {
	if m != nil {
		return m.Avatar
	}
	return nil
}

type ContactDetails struct {
	Number           *string                `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Name             *string                `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Avatar           *ContactDetails_Avatar `protobuf:"bytes,3,opt,name=avatar" json:"avatar,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ContactDetails) Reset()         { *m = ContactDetails{} }
func (m *ContactDetails) String() string { return proto.CompactTextString(m) }
func (*ContactDetails) ProtoMessage()    {}

func (m *ContactDetails) GetNumber() string {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return ""
}

func (m *ContactDetails) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ContactDetails) GetAvatar() *ContactDetails_Avatar {
	if m != nil {
		return m.Avatar
	}
	return nil
}

type ContactDetails_Avatar struct {
	ContentType      *string `protobuf:"bytes,1,opt,name=contentType" json:"contentType,omitempty"`
	Length           *uint32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ContactDetails_Avatar) Reset()         { *m = ContactDetails_Avatar{} }
func (m *ContactDetails_Avatar) String() string { return proto.CompactTextString(m) }
func (*ContactDetails_Avatar) ProtoMessage()    {}

func (m *ContactDetails_Avatar) GetContentType() string {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ""
}

func (m *ContactDetails_Avatar) GetLength() uint32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

type GroupDetails struct {
	Id               []byte               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name             *string              `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Members          []string             `protobuf:"bytes,3,rep,name=members" json:"members,omitempty"`
	Avatar           *GroupDetails_Avatar `protobuf:"bytes,4,opt,name=avatar" json:"avatar,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *GroupDetails) Reset()         { *m = GroupDetails{} }
func (m *GroupDetails) String() string { return proto.CompactTextString(m) }
func (*GroupDetails) ProtoMessage()    {}

func (m *GroupDetails) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GroupDetails) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *GroupDetails) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *GroupDetails) GetAvatar() *GroupDetails_Avatar {
	if m != nil {
		return m.Avatar
	}
	return nil
}

type GroupDetails_Avatar struct {
	ContentType      *string `protobuf:"bytes,1,opt,name=contentType" json:"contentType,omitempty"`
	Length           *uint32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GroupDetails_Avatar) Reset()         { *m = GroupDetails_Avatar{} }
func (m *GroupDetails_Avatar) String() string { return proto.CompactTextString(m) }
func (*GroupDetails_Avatar) ProtoMessage()    {}

func (m *GroupDetails_Avatar) GetContentType() string {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ""
}

func (m *GroupDetails_Avatar) GetLength() uint32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func init() {
	proto.RegisterEnum("textsecure.Envelope_Type", Envelope_Type_name, Envelope_Type_value)
	proto.RegisterEnum("textsecure.DataMessage_Flags", DataMessage_Flags_name, DataMessage_Flags_value)
	proto.RegisterEnum("textsecure.SyncMessage_Request_Type", SyncMessage_Request_Type_name, SyncMessage_Request_Type_value)
	proto.RegisterEnum("textsecure.GroupContext_Type", GroupContext_Type_name, GroupContext_Type_value)
}
