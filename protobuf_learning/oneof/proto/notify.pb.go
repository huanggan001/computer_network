// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.20.3
// source: notify.proto

//声明protobuf中的包名

package notify

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Notify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// Types that are assignable to NoticeWay:
	//
	//	*Notify_Phone
	//	*Notify_Email
	NoticeWay isNotify_NoticeWay `protobuf_oneof:"notice_way"`
}

func (x *Notify) Reset() {
	*x = Notify{}
	mi := &file_notify_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notify) ProtoMessage() {}

func (x *Notify) ProtoReflect() protoreflect.Message {
	mi := &file_notify_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notify.ProtoReflect.Descriptor instead.
func (*Notify) Descriptor() ([]byte, []int) {
	return file_notify_proto_rawDescGZIP(), []int{0}
}

func (x *Notify) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (m *Notify) GetNoticeWay() isNotify_NoticeWay {
	if m != nil {
		return m.NoticeWay
	}
	return nil
}

func (x *Notify) GetPhone() string {
	if x, ok := x.GetNoticeWay().(*Notify_Phone); ok {
		return x.Phone
	}
	return ""
}

func (x *Notify) GetEmail() string {
	if x, ok := x.GetNoticeWay().(*Notify_Email); ok {
		return x.Email
	}
	return ""
}

type isNotify_NoticeWay interface {
	isNotify_NoticeWay()
}

type Notify_Phone struct {
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3,oneof"`
}

type Notify_Email struct {
	Email string `protobuf:"bytes,3,opt,name=email,proto3,oneof"`
}

func (*Notify_Phone) isNotify_NoticeWay() {}

func (*Notify_Email) isNotify_NoticeWay() {}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// optional作用是该值可赋值可不赋值，比如price传了0，服务端那边不知道是默认值，还是它就是0
	Price *int64     `protobuf:"varint,2,opt,name=price,proto3,oneof" json:"price,omitempty"`
	Info  *Book_Info `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	mi := &file_notify_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_notify_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_notify_proto_rawDescGZIP(), []int{1}
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetPrice() int64 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *Book) GetInfo() *Book_Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op   string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`     //更新的人
	Book *Book  `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"` //书的信息
	// 要更新的字段
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateBookRequest) Reset() {
	*x = UpdateBookRequest{}
	mi := &file_notify_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookRequest) ProtoMessage() {}

func (x *UpdateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notify_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return file_notify_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateBookRequest) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *UpdateBookRequest) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

func (x *UpdateBookRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type Book_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B string `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *Book_Info) Reset() {
	*x = Book_Info{}
	mi := &file_notify_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book_Info) ProtoMessage() {}

func (x *Book_Info) ProtoReflect() protoreflect.Message {
	mi := &file_notify_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book_Info.ProtoReflect.Descriptor instead.
func (*Book_Info) Descriptor() ([]byte, []int) {
	return file_notify_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Book_Info) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *Book_Info) GetB() string {
	if x != nil {
		return x.B
	}
	return ""
}

var File_notify_proto protoreflect.FileDescriptor

var file_notify_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x0c, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65,
	0x5f, 0x77, 0x61, 0x79, 0x22, 0x8c, 0x01, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x1a, 0x22, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x0a,
	0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x62, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x20, 0x0a, 0x04, 0x62, 0x6f, 0x6f,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x37, 0x5a, 0x35, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notify_proto_rawDescOnce sync.Once
	file_notify_proto_rawDescData = file_notify_proto_rawDesc
)

func file_notify_proto_rawDescGZIP() []byte {
	file_notify_proto_rawDescOnce.Do(func() {
		file_notify_proto_rawDescData = protoimpl.X.CompressGZIP(file_notify_proto_rawDescData)
	})
	return file_notify_proto_rawDescData
}

var file_notify_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_notify_proto_goTypes = []any{
	(*Notify)(nil),                // 0: notify.Notify
	(*Book)(nil),                  // 1: notify.Book
	(*UpdateBookRequest)(nil),     // 2: notify.updateBookRequest
	(*Book_Info)(nil),             // 3: notify.Book.Info
	(*fieldmaskpb.FieldMask)(nil), // 4: google.protobuf.FieldMask
}
var file_notify_proto_depIdxs = []int32{
	3, // 0: notify.Book.info:type_name -> notify.Book.Info
	1, // 1: notify.updateBookRequest.book:type_name -> notify.Book
	4, // 2: notify.updateBookRequest.update_mask:type_name -> google.protobuf.FieldMask
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_notify_proto_init() }
func file_notify_proto_init() {
	if File_notify_proto != nil {
		return
	}
	file_notify_proto_msgTypes[0].OneofWrappers = []any{
		(*Notify_Phone)(nil),
		(*Notify_Email)(nil),
	}
	file_notify_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notify_proto_goTypes,
		DependencyIndexes: file_notify_proto_depIdxs,
		MessageInfos:      file_notify_proto_msgTypes,
	}.Build()
	File_notify_proto = out.File
	file_notify_proto_rawDesc = nil
	file_notify_proto_goTypes = nil
	file_notify_proto_depIdxs = nil
}
