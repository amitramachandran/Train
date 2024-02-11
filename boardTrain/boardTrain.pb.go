// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.1
// source: boardTrain.proto

package boardTrain

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname    string `protobuf:"bytes,1,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname     string `protobuf:"bytes,2,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
	EmailAddress string `protobuf:"bytes,3,opt,name=EmailAddress,proto3" json:"EmailAddress,omitempty"`
	Id           int32  `protobuf:"varint,4,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boardTrain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_boardTrain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_boardTrain_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *User) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *User) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Transit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string `protobuf:"bytes,1,opt,name=From,proto3" json:"From,omitempty"`
	To   string `protobuf:"bytes,2,opt,name=To,proto3" json:"To,omitempty"`
}

func (x *Transit) Reset() {
	*x = Transit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boardTrain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transit) ProtoMessage() {}

func (x *Transit) ProtoReflect() protoreflect.Message {
	mi := &file_boardTrain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transit.ProtoReflect.Descriptor instead.
func (*Transit) Descriptor() ([]byte, []int) {
	return file_boardTrain_proto_rawDescGZIP(), []int{1}
}

func (x *Transit) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Transit) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount   float64 `protobuf:"fixed64,1,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Currency string  `protobuf:"bytes,2,opt,name=Currency,proto3" json:"Currency,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boardTrain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_boardTrain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_boardTrain_proto_rawDescGZIP(), []int{2}
}

func (x *Price) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Price) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type TicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Route     *Transit `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	Passenger *User    `protobuf:"bytes,2,opt,name=passenger,proto3" json:"passenger,omitempty"`
}

func (x *TicketRequest) Reset() {
	*x = TicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boardTrain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketRequest) ProtoMessage() {}

func (x *TicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boardTrain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketRequest.ProtoReflect.Descriptor instead.
func (*TicketRequest) Descriptor() ([]byte, []int) {
	return file_boardTrain_proto_rawDescGZIP(), []int{3}
}

func (x *TicketRequest) GetRoute() *Transit {
	if x != nil {
		return x.Route
	}
	return nil
}

func (x *TicketRequest) GetPassenger() *User {
	if x != nil {
		return x.Passenger
	}
	return nil
}

type TicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From         string `protobuf:"bytes,1,opt,name=From,proto3" json:"From,omitempty"`
	To           string `protobuf:"bytes,2,opt,name=To,proto3" json:"To,omitempty"`
	UserDetails  *User  `protobuf:"bytes,3,opt,name=UserDetails,proto3" json:"UserDetails,omitempty"`
	PriceDetails *Price `protobuf:"bytes,4,opt,name=PriceDetails,proto3" json:"PriceDetails,omitempty"`
}

func (x *TicketResponse) Reset() {
	*x = TicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boardTrain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketResponse) ProtoMessage() {}

func (x *TicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boardTrain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketResponse.ProtoReflect.Descriptor instead.
func (*TicketResponse) Descriptor() ([]byte, []int) {
	return file_boardTrain_proto_rawDescGZIP(), []int{4}
}

func (x *TicketResponse) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TicketResponse) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TicketResponse) GetUserDetails() *User {
	if x != nil {
		return x.UserDetails
	}
	return nil
}

func (x *TicketResponse) GetPriceDetails() *Price {
	if x != nil {
		return x.PriceDetails
	}
	return nil
}

type Seat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section string `protobuf:"bytes,1,opt,name=Section,proto3" json:"Section,omitempty"`
	Number  int32  `protobuf:"varint,2,opt,name=Number,proto3" json:"Number,omitempty"`
}

func (x *Seat) Reset() {
	*x = Seat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boardTrain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Seat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Seat) ProtoMessage() {}

func (x *Seat) ProtoReflect() protoreflect.Message {
	mi := &file_boardTrain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Seat.ProtoReflect.Descriptor instead.
func (*Seat) Descriptor() ([]byte, []int) {
	return file_boardTrain_proto_rawDescGZIP(), []int{5}
}

func (x *Seat) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

func (x *Seat) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailAddress string `protobuf:"bytes,1,opt,name=EmailAddress,proto3" json:"EmailAddress,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boardTrain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boardTrain_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_boardTrain_proto_rawDescGZIP(), []int{6}
}

func (x *UserRequest) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

type SectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section string `protobuf:"bytes,1,opt,name=Section,proto3" json:"Section,omitempty"`
}

func (x *SectionRequest) Reset() {
	*x = SectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boardTrain_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SectionRequest) ProtoMessage() {}

func (x *SectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boardTrain_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SectionRequest.ProtoReflect.Descriptor instead.
func (*SectionRequest) Descriptor() ([]byte, []int) {
	return file_boardTrain_proto_rawDescGZIP(), []int{7}
}

func (x *SectionRequest) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

type SeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
}

func (x *SeatResponse) Reset() {
	*x = SeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boardTrain_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeatResponse) ProtoMessage() {}

func (x *SeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boardTrain_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeatResponse.ProtoReflect.Descriptor instead.
func (*SeatResponse) Descriptor() ([]byte, []int) {
	return file_boardTrain_proto_rawDescGZIP(), []int{8}
}

func (x *SeatResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boardTrain_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boardTrain_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_boardTrain_proto_rawDescGZIP(), []int{9}
}

func (x *StatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ModifySeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailAddress string `protobuf:"bytes,1,opt,name=EmailAddress,proto3" json:"EmailAddress,omitempty"`
	NewSeat      string `protobuf:"bytes,2,opt,name=NewSeat,proto3" json:"NewSeat,omitempty"`
}

func (x *ModifySeatRequest) Reset() {
	*x = ModifySeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boardTrain_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySeatRequest) ProtoMessage() {}

func (x *ModifySeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boardTrain_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySeatRequest.ProtoReflect.Descriptor instead.
func (*ModifySeatRequest) Descriptor() ([]byte, []int) {
	return file_boardTrain_proto_rawDescGZIP(), []int{10}
}

func (x *ModifySeatRequest) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *ModifySeatRequest) GetNewSeat() string {
	if x != nil {
		return x.NewSeat
	}
	return ""
}

var File_boardTrain_proto protoreflect.FileDescriptor

var file_boardTrain_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x74, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69,
	0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46,
	0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x54, 0x6f, 0x22, 0x3b, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x22, 0x54, 0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x52, 0x05,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x09, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x09, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x22, 0x89, 0x01, 0x0a, 0x0e, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x72, 0x6f,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x54,
	0x6f, 0x12, 0x27, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2a, 0x0a, 0x0c, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x06, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x0c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x38, 0x0a, 0x04, 0x53, 0x65, 0x61, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x31, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x2b, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1b, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x2a, 0x0a, 0x0e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x51, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x61, 0x74, 0x32, 0xf7, 0x01, 0x0a, 0x08,
	0x42, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x31, 0x0a, 0x0e, 0x70, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x0a, 0x67,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x0c, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53,
	0x65, 0x61, 0x74, 0x12, 0x12, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x76, 0x69, 0x65, 0x77,
	0x53, 0x65, 0x61, 0x74, 0x73, 0x12, 0x0f, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6d, 0x69, 0x74, 0x72, 0x61, 0x6d, 0x61, 0x63, 0x68, 0x61, 0x6e,
	0x64, 0x72, 0x61, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_boardTrain_proto_rawDescOnce sync.Once
	file_boardTrain_proto_rawDescData = file_boardTrain_proto_rawDesc
)

func file_boardTrain_proto_rawDescGZIP() []byte {
	file_boardTrain_proto_rawDescOnce.Do(func() {
		file_boardTrain_proto_rawDescData = protoimpl.X.CompressGZIP(file_boardTrain_proto_rawDescData)
	})
	return file_boardTrain_proto_rawDescData
}

var file_boardTrain_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_boardTrain_proto_goTypes = []interface{}{
	(*User)(nil),              // 0: User
	(*Transit)(nil),           // 1: Transit
	(*Price)(nil),             // 2: Price
	(*TicketRequest)(nil),     // 3: TicketRequest
	(*TicketResponse)(nil),    // 4: TicketResponse
	(*Seat)(nil),              // 5: Seat
	(*UserRequest)(nil),       // 6: UserRequest
	(*SectionRequest)(nil),    // 7: SectionRequest
	(*SeatResponse)(nil),      // 8: SeatResponse
	(*StatusResponse)(nil),    // 9: StatusResponse
	(*ModifySeatRequest)(nil), // 10: ModifySeatRequest
}
var file_boardTrain_proto_depIdxs = []int32{
	1,  // 0: TicketRequest.route:type_name -> Transit
	0,  // 1: TicketRequest.passenger:type_name -> User
	0,  // 2: TicketResponse.UserDetails:type_name -> User
	2,  // 3: TicketResponse.PriceDetails:type_name -> Price
	0,  // 4: SeatResponse.Users:type_name -> User
	3,  // 5: Boarding.purchaseTicket:input_type -> TicketRequest
	6,  // 6: Boarding.getReceipt:input_type -> UserRequest
	6,  // 7: Boarding.removeUser:input_type -> UserRequest
	10, // 8: Boarding.modifySeat:input_type -> ModifySeatRequest
	7,  // 9: Boarding.viewSeats:input_type -> SectionRequest
	4,  // 10: Boarding.purchaseTicket:output_type -> TicketResponse
	4,  // 11: Boarding.getReceipt:output_type -> TicketResponse
	9,  // 12: Boarding.removeUser:output_type -> StatusResponse
	9,  // 13: Boarding.modifySeat:output_type -> StatusResponse
	8,  // 14: Boarding.viewSeats:output_type -> SeatResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_boardTrain_proto_init() }
func file_boardTrain_proto_init() {
	if File_boardTrain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_boardTrain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_boardTrain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transit); i {
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
		file_boardTrain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_boardTrain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketRequest); i {
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
		file_boardTrain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketResponse); i {
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
		file_boardTrain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Seat); i {
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
		file_boardTrain_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_boardTrain_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SectionRequest); i {
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
		file_boardTrain_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeatResponse); i {
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
		file_boardTrain_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_boardTrain_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifySeatRequest); i {
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
			RawDescriptor: file_boardTrain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_boardTrain_proto_goTypes,
		DependencyIndexes: file_boardTrain_proto_depIdxs,
		MessageInfos:      file_boardTrain_proto_msgTypes,
	}.Build()
	File_boardTrain_proto = out.File
	file_boardTrain_proto_rawDesc = nil
	file_boardTrain_proto_goTypes = nil
	file_boardTrain_proto_depIdxs = nil
}
