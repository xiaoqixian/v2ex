// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: protobuf/post.proto

package postpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type PublishPostRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint64                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Node          string                 `protobuf:"bytes,4,opt,name=node,proto3" json:"node,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublishPostRequest) Reset() {
	*x = PublishPostRequest{}
	mi := &file_protobuf_post_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishPostRequest) ProtoMessage() {}

func (x *PublishPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_post_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishPostRequest.ProtoReflect.Descriptor instead.
func (*PublishPostRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_post_proto_rawDescGZIP(), []int{0}
}

func (x *PublishPostRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PublishPostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PublishPostRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PublishPostRequest) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

type PublishPostResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PostId        uint64                 `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublishPostResponse) Reset() {
	*x = PublishPostResponse{}
	mi := &file_protobuf_post_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishPostResponse) ProtoMessage() {}

func (x *PublishPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_post_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishPostResponse.ProtoReflect.Descriptor instead.
func (*PublishPostResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_post_proto_rawDescGZIP(), []int{1}
}

func (x *PublishPostResponse) GetPostId() uint64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *PublishPostResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PostEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PostId        uint64                 `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	AuthorId      uint64                 `protobuf:"varint,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Node          string                 `protobuf:"bytes,4,opt,name=node,proto3" json:"node,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostEntry) Reset() {
	*x = PostEntry{}
	mi := &file_protobuf_post_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostEntry) ProtoMessage() {}

func (x *PostEntry) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_post_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostEntry.ProtoReflect.Descriptor instead.
func (*PostEntry) Descriptor() ([]byte, []int) {
	return file_protobuf_post_proto_rawDescGZIP(), []int{2}
}

func (x *PostEntry) GetPostId() uint64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *PostEntry) GetAuthorId() uint64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *PostEntry) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostEntry) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *PostEntry) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type GetPostRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PostId        uint64                 `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	mi := &file_protobuf_post_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_post_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_post_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostRequest) GetPostId() uint64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

type GetPostResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Found         bool                   `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	PostId        uint64                 `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	AuthorId      uint64                 `protobuf:"varint,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title         string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Node          string                 `protobuf:"bytes,5,opt,name=node,proto3" json:"node,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Content       string                 `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPostResponse) Reset() {
	*x = GetPostResponse{}
	mi := &file_protobuf_post_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostResponse) ProtoMessage() {}

func (x *GetPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_post_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostResponse.ProtoReflect.Descriptor instead.
func (*GetPostResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_post_proto_rawDescGZIP(), []int{4}
}

func (x *GetPostResponse) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

func (x *GetPostResponse) GetPostId() uint64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *GetPostResponse) GetAuthorId() uint64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *GetPostResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetPostResponse) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *GetPostResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GetPostResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type GetPostsForUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint64                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPostsForUserRequest) Reset() {
	*x = GetPostsForUserRequest{}
	mi := &file_protobuf_post_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostsForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsForUserRequest) ProtoMessage() {}

func (x *GetPostsForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_post_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsForUserRequest.ProtoReflect.Descriptor instead.
func (*GetPostsForUserRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_post_proto_rawDescGZIP(), []int{5}
}

func (x *GetPostsForUserRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetPostsForUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Posts         []*PostEntry           `protobuf:"bytes,2,rep,name=posts,proto3" json:"posts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPostsForUserResponse) Reset() {
	*x = GetPostsForUserResponse{}
	mi := &file_protobuf_post_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostsForUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsForUserResponse) ProtoMessage() {}

func (x *GetPostsForUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_post_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsForUserResponse.ProtoReflect.Descriptor instead.
func (*GetPostsForUserResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_post_proto_rawDescGZIP(), []int{6}
}

func (x *GetPostsForUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetPostsForUserResponse) GetPosts() []*PostEntry {
	if x != nil {
		return x.Posts
	}
	return nil
}

type AddPostViewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddPostViewRequest) Reset() {
	*x = AddPostViewRequest{}
	mi := &file_protobuf_post_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddPostViewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPostViewRequest) ProtoMessage() {}

func (x *AddPostViewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_post_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPostViewRequest.ProtoReflect.Descriptor instead.
func (*AddPostViewRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_post_proto_rawDescGZIP(), []int{7}
}

type AddPostViewResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddPostViewResponse) Reset() {
	*x = AddPostViewResponse{}
	mi := &file_protobuf_post_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddPostViewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPostViewResponse) ProtoMessage() {}

func (x *AddPostViewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_post_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPostViewResponse.ProtoReflect.Descriptor instead.
func (*AddPostViewResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_post_proto_rawDescGZIP(), []int{8}
}

var File_protobuf_post_proto protoreflect.FileDescriptor

const file_protobuf_post_proto_rawDesc = "" +
	"\n" +
	"\x13protobuf/post.proto\x12\x04post\x1a\x1fgoogle/protobuf/timestamp.proto\"q\n" +
	"\x12PublishPostRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x04R\x06userId\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x03 \x01(\tR\acontent\x12\x12\n" +
	"\x04node\x18\x04 \x01(\tR\x04node\"H\n" +
	"\x13PublishPostResponse\x12\x17\n" +
	"\apost_id\x18\x01 \x01(\x04R\x06postId\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"\xa6\x01\n" +
	"\tPostEntry\x12\x17\n" +
	"\apost_id\x18\x01 \x01(\x04R\x06postId\x12\x1b\n" +
	"\tauthor_id\x18\x02 \x01(\x04R\bauthorId\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x12\x12\n" +
	"\x04node\x18\x04 \x01(\tR\x04node\x129\n" +
	"\n" +
	"created_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\")\n" +
	"\x0eGetPostRequest\x12\x17\n" +
	"\apost_id\x18\x01 \x01(\x04R\x06postId\"\xdc\x01\n" +
	"\x0fGetPostResponse\x12\x14\n" +
	"\x05found\x18\x01 \x01(\bR\x05found\x12\x17\n" +
	"\apost_id\x18\x02 \x01(\x04R\x06postId\x12\x1b\n" +
	"\tauthor_id\x18\x03 \x01(\x04R\bauthorId\x12\x14\n" +
	"\x05title\x18\x04 \x01(\tR\x05title\x12\x12\n" +
	"\x04node\x18\x05 \x01(\tR\x04node\x129\n" +
	"\n" +
	"created_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x18\n" +
	"\acontent\x18\a \x01(\tR\acontent\"1\n" +
	"\x16GetPostsForUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x04R\x06userId\"Z\n" +
	"\x17GetPostsForUserResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12%\n" +
	"\x05posts\x18\x02 \x03(\v2\x0f.post.PostEntryR\x05posts\"\x14\n" +
	"\x12AddPostViewRequest\"\x15\n" +
	"\x13AddPostViewResponse2\x9d\x02\n" +
	"\vPostService\x12B\n" +
	"\vPublishPost\x12\x18.post.PublishPostRequest\x1a\x19.post.PublishPostResponse\x12N\n" +
	"\x0fGetPostsForUser\x12\x1c.post.GetPostsForUserRequest\x1a\x1d.post.GetPostsForUserResponse\x126\n" +
	"\aGetPost\x12\x14.post.GetPostRequest\x1a\x15.post.GetPostResponse\x12B\n" +
	"\vAddPostView\x12\x18.post.AddPostViewRequest\x1a\x19.post.AddPostViewResponseB\x17Z\x15rpc_gen/postpb;postpbb\x06proto3"

var (
	file_protobuf_post_proto_rawDescOnce sync.Once
	file_protobuf_post_proto_rawDescData []byte
)

func file_protobuf_post_proto_rawDescGZIP() []byte {
	file_protobuf_post_proto_rawDescOnce.Do(func() {
		file_protobuf_post_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protobuf_post_proto_rawDesc), len(file_protobuf_post_proto_rawDesc)))
	})
	return file_protobuf_post_proto_rawDescData
}

var file_protobuf_post_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protobuf_post_proto_goTypes = []any{
	(*PublishPostRequest)(nil),      // 0: post.PublishPostRequest
	(*PublishPostResponse)(nil),     // 1: post.PublishPostResponse
	(*PostEntry)(nil),               // 2: post.PostEntry
	(*GetPostRequest)(nil),          // 3: post.GetPostRequest
	(*GetPostResponse)(nil),         // 4: post.GetPostResponse
	(*GetPostsForUserRequest)(nil),  // 5: post.GetPostsForUserRequest
	(*GetPostsForUserResponse)(nil), // 6: post.GetPostsForUserResponse
	(*AddPostViewRequest)(nil),      // 7: post.AddPostViewRequest
	(*AddPostViewResponse)(nil),     // 8: post.AddPostViewResponse
	(*timestamppb.Timestamp)(nil),   // 9: google.protobuf.Timestamp
}
var file_protobuf_post_proto_depIdxs = []int32{
	9, // 0: post.PostEntry.created_at:type_name -> google.protobuf.Timestamp
	9, // 1: post.GetPostResponse.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: post.GetPostsForUserResponse.posts:type_name -> post.PostEntry
	0, // 3: post.PostService.PublishPost:input_type -> post.PublishPostRequest
	5, // 4: post.PostService.GetPostsForUser:input_type -> post.GetPostsForUserRequest
	3, // 5: post.PostService.GetPost:input_type -> post.GetPostRequest
	7, // 6: post.PostService.AddPostView:input_type -> post.AddPostViewRequest
	1, // 7: post.PostService.PublishPost:output_type -> post.PublishPostResponse
	6, // 8: post.PostService.GetPostsForUser:output_type -> post.GetPostsForUserResponse
	4, // 9: post.PostService.GetPost:output_type -> post.GetPostResponse
	8, // 10: post.PostService.AddPostView:output_type -> post.AddPostViewResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protobuf_post_proto_init() }
func file_protobuf_post_proto_init() {
	if File_protobuf_post_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protobuf_post_proto_rawDesc), len(file_protobuf_post_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_post_proto_goTypes,
		DependencyIndexes: file_protobuf_post_proto_depIdxs,
		MessageInfos:      file_protobuf_post_proto_msgTypes,
	}.Build()
	File_protobuf_post_proto = out.File
	file_protobuf_post_proto_goTypes = nil
	file_protobuf_post_proto_depIdxs = nil
}
