// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: article.proto

package dawn

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Article struct {
	Id      string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string            `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string            `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Image   string            `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Video   string            `protobuf:"bytes,5,opt,name=video,proto3" json:"video,omitempty"`
	Owner   string            `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	Hiddens []int32           `protobuf:"varint,7,rep,packed,name=hiddens,proto3" json:"hiddens,omitempty"`
	Created *types.Timestamp  `protobuf:"bytes,8,opt,name=created,proto3" json:"created,omitempty"`
	Labels  map[string]string `protobuf:"bytes,9,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Article) Reset()         { *m = Article{} }
func (m *Article) String() string { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()    {}
func (*Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{0}
}
func (m *Article) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Article) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Article.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Article) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article.Merge(m, src)
}
func (m *Article) XXX_Size() int {
	return m.Size()
}
func (m *Article) XXX_DiscardUnknown() {
	xxx_messageInfo_Article.DiscardUnknown(m)
}

var xxx_messageInfo_Article proto.InternalMessageInfo

func (m *Article) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Article) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Article) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Article) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Article) GetVideo() string {
	if m != nil {
		return m.Video
	}
	return ""
}

func (m *Article) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Article) GetHiddens() []int32 {
	if m != nil {
		return m.Hiddens
	}
	return nil
}

func (m *Article) GetCreated() *types.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Article) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*Article)(nil), "dawn.Article")
	proto.RegisterMapType((map[string]string)(nil), "dawn.Article.LabelsEntry")
}

func init() { proto.RegisterFile("article.proto", fileDescriptor_5c593d380f9840a2) }

var fileDescriptor_5c593d380f9840a2 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x18, 0xc4, 0xe3, 0x64, 0x9b, 0xec, 0xba, 0x5a, 0x84, 0x2c, 0x84, 0x4c, 0x90, 0x42, 0x54, 0x09,
	0x29, 0x27, 0x2f, 0x14, 0x0e, 0xc0, 0x6d, 0x57, 0xac, 0xb8, 0xf4, 0x14, 0xd1, 0x07, 0x70, 0xeb,
	0x8f, 0x62, 0x91, 0x7f, 0x72, 0x5c, 0xaa, 0xbc, 0x05, 0x8f, 0x05, 0xb7, 0x1e, 0x39, 0xa2, 0x56,
	0xbc, 0x07, 0xb2, 0x9d, 0xa0, 0x6d, 0x7b, 0xe8, 0x2d, 0x33, 0xf3, 0xd3, 0x67, 0xcd, 0x28, 0xf8,
	0x9a, 0x2b, 0x2d, 0x97, 0x05, 0xb0, 0x46, 0xd5, 0xba, 0x26, 0x17, 0x82, 0x6f, 0xaa, 0xf8, 0xf9,
	0xaa, 0xae, 0x57, 0x05, 0xdc, 0x58, 0x6f, 0xb1, 0xfe, 0x72, 0x03, 0x65, 0xa3, 0x3b, 0x87, 0xc4,
	0xc9, 0x71, 0xb8, 0x51, 0xbc, 0x69, 0x40, 0xb5, 0x7d, 0xfe, 0xe2, 0x38, 0xd7, 0xb2, 0x84, 0x56,
	0xf3, 0xb2, 0x71, 0xc0, 0xe4, 0x97, 0x8f, 0xa3, 0x5b, 0xf7, 0x2a, 0x79, 0x84, 0x7d, 0x29, 0x28,
	0x4a, 0x51, 0x76, 0x95, 0xfb, 0x52, 0x90, 0x27, 0x78, 0xa4, 0xa5, 0x2e, 0x80, 0xfa, 0xd6, 0x72,
	0x82, 0x50, 0x1c, 0x2d, 0xeb, 0x4a, 0x43, 0xa5, 0x69, 0x60, 0xfd, 0x41, 0x1a, 0x5e, 0x96, 0x7c,
	0x05, 0xf4, 0xc2, 0xf1, 0x56, 0x18, 0xf7, 0xbb, 0x14, 0x50, 0xd3, 0x91, 0x73, 0xad, 0x30, 0x6e,
	0xbd, 0xa9, 0x40, 0xd1, 0xd0, 0xb9, 0x56, 0x98, 0xdb, 0x5f, 0xa5, 0x10, 0x50, 0xb5, 0x34, 0x4a,
	0x83, 0x6c, 0x94, 0x0f, 0x92, 0xbc, 0xc5, 0xd1, 0x52, 0x01, 0xd7, 0x20, 0xe8, 0x65, 0x8a, 0xb2,
	0xf1, 0x34, 0x66, 0xae, 0x1a, 0x1b, 0xaa, 0xb1, 0xcf, 0x43, 0xb5, 0x7c, 0x40, 0xc9, 0x6b, 0x1c,
	0x16, 0x7c, 0x01, 0x45, 0x4b, 0xaf, 0xd2, 0x20, 0x1b, 0x4f, 0x9f, 0x31, 0x33, 0x29, 0xeb, 0x0b,
	0xb3, 0x99, 0xcd, 0xee, 0x2b, 0xad, 0xba, 0xbc, 0x07, 0xe3, 0xf7, 0x78, 0xfc, 0xc0, 0x26, 0x8f,
	0x71, 0xf0, 0x0d, 0xba, 0x7e, 0x14, 0xf3, 0x69, 0xfb, 0xf0, 0x62, 0xfd, 0x7f, 0x15, 0x2b, 0x3e,
	0xf8, 0xef, 0xd0, 0xf4, 0x2f, 0xc2, 0x97, 0xfd, 0xe9, 0x96, 0xbc, 0xc4, 0xc1, 0xad, 0x10, 0xe4,
	0xfa, 0xe0, 0xc5, 0xf8, 0x50, 0x4e, 0x3c, 0x83, 0x7d, 0x02, 0x7d, 0x16, 0xcb, 0x70, 0x38, 0x6f,
	0x04, 0xd7, 0x70, 0x96, 0x64, 0x18, 0xcf, 0x64, 0xab, 0xef, 0xba, 0x79, 0x0b, 0xea, 0x1c, 0xfd,
	0x0a, 0x99, 0x89, 0x3e, 0x42, 0x01, 0xa7, 0x97, 0x9f, 0x9e, 0x0c, 0x7c, 0x6f, 0x7e, 0xbc, 0x89,
	0x77, 0x47, 0x7f, 0xee, 0x12, 0xb4, 0xdd, 0x25, 0xe8, 0xcf, 0x2e, 0x41, 0x3f, 0xf6, 0x89, 0xb7,
	0xdd, 0x27, 0xde, 0xef, 0x7d, 0xe2, 0x2d, 0x42, 0xcb, 0xbe, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff,
	0xb3, 0xd8, 0x00, 0xe1, 0xc9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArticlesClient is the client API for Articles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArticlesClient interface {
	Add(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error)
	Get(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error)
	Update(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error)
	ListByUser(ctx context.Context, in *Article, opts ...grpc.CallOption) (Articles_ListByUserClient, error)
	Delete(ctx context.Context, in *Article, opts ...grpc.CallOption) (*types.Empty, error)
}

type articlesClient struct {
	cc *grpc.ClientConn
}

func NewArticlesClient(cc *grpc.ClientConn) ArticlesClient {
	return &articlesClient{cc}
}

func (c *articlesClient) Add(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/dawn.Articles/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) Get(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/dawn.Articles/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) Update(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/dawn.Articles/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) ListByUser(ctx context.Context, in *Article, opts ...grpc.CallOption) (Articles_ListByUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Articles_serviceDesc.Streams[0], "/dawn.Articles/ListByUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &articlesListByUserClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Articles_ListByUserClient interface {
	Recv() (*Article, error)
	grpc.ClientStream
}

type articlesListByUserClient struct {
	grpc.ClientStream
}

func (x *articlesListByUserClient) Recv() (*Article, error) {
	m := new(Article)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *articlesClient) Delete(ctx context.Context, in *Article, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/dawn.Articles/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticlesServer is the server API for Articles service.
type ArticlesServer interface {
	Add(context.Context, *Article) (*Article, error)
	Get(context.Context, *Article) (*Article, error)
	Update(context.Context, *Article) (*Article, error)
	ListByUser(*Article, Articles_ListByUserServer) error
	Delete(context.Context, *Article) (*types.Empty, error)
}

func RegisterArticlesServer(s *grpc.Server, srv ArticlesServer) {
	s.RegisterService(&_Articles_serviceDesc, srv)
}

func _Articles_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Article)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dawn.Articles/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).Add(ctx, req.(*Article))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Article)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dawn.Articles/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).Get(ctx, req.(*Article))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Article)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dawn.Articles/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).Update(ctx, req.(*Article))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_ListByUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Article)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArticlesServer).ListByUser(m, &articlesListByUserServer{stream})
}

type Articles_ListByUserServer interface {
	Send(*Article) error
	grpc.ServerStream
}

type articlesListByUserServer struct {
	grpc.ServerStream
}

func (x *articlesListByUserServer) Send(m *Article) error {
	return x.ServerStream.SendMsg(m)
}

func _Articles_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Article)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dawn.Articles/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).Delete(ctx, req.(*Article))
	}
	return interceptor(ctx, in, info, handler)
}

var _Articles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dawn.Articles",
	HandlerType: (*ArticlesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Articles_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Articles_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Articles_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Articles_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListByUser",
			Handler:       _Articles_ListByUser_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "article.proto",
}

func (m *Article) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Article) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintArticle(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Title) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintArticle(dAtA, i, uint64(len(m.Title)))
		i += copy(dAtA[i:], m.Title)
	}
	if len(m.Content) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintArticle(dAtA, i, uint64(len(m.Content)))
		i += copy(dAtA[i:], m.Content)
	}
	if len(m.Image) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintArticle(dAtA, i, uint64(len(m.Image)))
		i += copy(dAtA[i:], m.Image)
	}
	if len(m.Video) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintArticle(dAtA, i, uint64(len(m.Video)))
		i += copy(dAtA[i:], m.Video)
	}
	if len(m.Owner) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintArticle(dAtA, i, uint64(len(m.Owner)))
		i += copy(dAtA[i:], m.Owner)
	}
	if len(m.Hiddens) > 0 {
		dAtA2 := make([]byte, len(m.Hiddens)*10)
		var j1 int
		for _, num1 := range m.Hiddens {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x3a
		i++
		i = encodeVarintArticle(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if m.Created != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintArticle(dAtA, i, uint64(m.Created.Size()))
		n3, err := m.Created.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			dAtA[i] = 0x4a
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovArticle(uint64(len(k))) + 1 + len(v) + sovArticle(uint64(len(v)))
			i = encodeVarintArticle(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintArticle(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintArticle(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func encodeVarintArticle(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Article) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovArticle(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovArticle(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovArticle(uint64(l))
	}
	l = len(m.Image)
	if l > 0 {
		n += 1 + l + sovArticle(uint64(l))
	}
	l = len(m.Video)
	if l > 0 {
		n += 1 + l + sovArticle(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovArticle(uint64(l))
	}
	if len(m.Hiddens) > 0 {
		l = 0
		for _, e := range m.Hiddens {
			l += sovArticle(uint64(e))
		}
		n += 1 + sovArticle(uint64(l)) + l
	}
	if m.Created != nil {
		l = m.Created.Size()
		n += 1 + l + sovArticle(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovArticle(uint64(len(k))) + 1 + len(v) + sovArticle(uint64(len(v)))
			n += mapEntrySize + 1 + sovArticle(uint64(mapEntrySize))
		}
	}
	return n
}

func sovArticle(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozArticle(x uint64) (n int) {
	return sovArticle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Article) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowArticle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Article: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Article: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthArticle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArticle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthArticle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArticle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthArticle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArticle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthArticle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArticle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Video", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthArticle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArticle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Video = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthArticle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArticle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowArticle
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Hiddens = append(m.Hiddens, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowArticle
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthArticle
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthArticle
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Hiddens) == 0 {
					m.Hiddens = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowArticle
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Hiddens = append(m.Hiddens, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Hiddens", wireType)
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthArticle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthArticle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Created == nil {
				m.Created = &types.Timestamp{}
			}
			if err := m.Created.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthArticle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthArticle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowArticle
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowArticle
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthArticle
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthArticle
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowArticle
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthArticle
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthArticle
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipArticle(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthArticle
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipArticle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthArticle
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthArticle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipArticle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowArticle
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowArticle
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowArticle
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthArticle
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthArticle
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowArticle
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipArticle(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthArticle
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthArticle = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowArticle   = fmt.Errorf("proto: integer overflow")
)
