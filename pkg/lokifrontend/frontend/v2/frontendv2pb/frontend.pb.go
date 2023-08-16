// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/lokifrontend/frontend/v2/frontendv2pb/frontend.proto

package frontendv2pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	httpgrpc "github.com/grafana/dskit/httpgrpc"
	stats "github.com/grafana/loki/pkg/querier/stats"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryResultRequest struct {
	QueryID      uint64                 `protobuf:"varint,1,opt,name=queryID,proto3" json:"queryID,omitempty"`
	HttpResponse *httpgrpc.HTTPResponse `protobuf:"bytes,2,opt,name=httpResponse,proto3" json:"httpResponse,omitempty"`
	Stats        *stats.Stats           `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (m *QueryResultRequest) Reset()      { *m = QueryResultRequest{} }
func (*QueryResultRequest) ProtoMessage() {}
func (*QueryResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_85a7e5cdf8261f06, []int{0}
}
func (m *QueryResultRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryResultRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResultRequest.Merge(m, src)
}
func (m *QueryResultRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResultRequest proto.InternalMessageInfo

func (m *QueryResultRequest) GetQueryID() uint64 {
	if m != nil {
		return m.QueryID
	}
	return 0
}

func (m *QueryResultRequest) GetHttpResponse() *httpgrpc.HTTPResponse {
	if m != nil {
		return m.HttpResponse
	}
	return nil
}

func (m *QueryResultRequest) GetStats() *stats.Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type QueryResultResponse struct {
}

func (m *QueryResultResponse) Reset()      { *m = QueryResultResponse{} }
func (*QueryResultResponse) ProtoMessage() {}
func (*QueryResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85a7e5cdf8261f06, []int{1}
}
func (m *QueryResultResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryResultResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResultResponse.Merge(m, src)
}
func (m *QueryResultResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResultResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryResultRequest)(nil), "frontendv2pb.QueryResultRequest")
	proto.RegisterType((*QueryResultResponse)(nil), "frontendv2pb.QueryResultResponse")
}

func init() {
	proto.RegisterFile("pkg/lokifrontend/frontend/v2/frontendv2pb/frontend.proto", fileDescriptor_85a7e5cdf8261f06)
}

var fileDescriptor_85a7e5cdf8261f06 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xbd, 0x6e, 0xf2, 0x30,
	0x14, 0xb5, 0xbf, 0x9f, 0x56, 0x32, 0x4c, 0xee, 0x8f, 0x22, 0xa4, 0x5e, 0xd1, 0x4c, 0x4c, 0x71,
	0x95, 0x2e, 0x55, 0x47, 0x54, 0xa1, 0x76, 0x2b, 0x29, 0x53, 0xb7, 0x00, 0x26, 0xa4, 0xd0, 0x38,
	0xd8, 0x0e, 0x12, 0x5b, 0x9f, 0xa0, 0xea, 0x63, 0xf4, 0x51, 0x3a, 0x32, 0x32, 0x16, 0xb3, 0x74,
	0xe4, 0x11, 0xaa, 0xc4, 0x10, 0x81, 0x2a, 0x75, 0xb9, 0x3a, 0xd7, 0xe7, 0x1e, 0xdd, 0x73, 0x8f,
	0xc9, 0x55, 0x3a, 0x8a, 0xd8, 0x58, 0x8c, 0xe2, 0x81, 0x14, 0x89, 0xe6, 0x49, 0x9f, 0x95, 0x60,
	0xea, 0x97, 0x78, 0xea, 0xa7, 0xdd, 0xb2, 0xf1, 0x52, 0x29, 0xb4, 0xa0, 0xd5, 0x5d, 0xb2, 0x76,
	0x11, 0xc5, 0x7a, 0x98, 0x75, 0xbd, 0x9e, 0x78, 0x66, 0x91, 0x0c, 0x07, 0x61, 0x12, 0xb2, 0xbe,
	0x1a, 0xc5, 0x9a, 0x0d, 0xb5, 0x4e, 0x23, 0x99, 0xf6, 0x4a, 0x60, 0xf5, 0xb5, 0xe3, 0x48, 0x44,
	0xa2, 0x80, 0x2c, 0x47, 0x9b, 0xd7, 0xb3, 0xdc, 0xcf, 0x24, 0xe3, 0x32, 0xe6, 0x92, 0x29, 0x1d,
	0x6a, 0x65, 0xab, 0xa5, 0xdd, 0x57, 0x4c, 0x68, 0x3b, 0xe3, 0x72, 0x16, 0x70, 0x95, 0x8d, 0x75,
	0xc0, 0x27, 0x19, 0x57, 0x9a, 0x3a, 0xe4, 0x30, 0xd7, 0xcc, 0xee, 0x6e, 0x1c, 0x5c, 0xc7, 0x8d,
	0x7f, 0xc1, 0xb6, 0xa5, 0xd7, 0xa4, 0x9a, 0xef, 0x0d, 0xb8, 0x4a, 0x45, 0xa2, 0xb8, 0xf3, 0xa7,
	0x8e, 0x1b, 0x15, 0xff, 0xd4, 0x2b, 0xcd, 0xdc, 0x76, 0x3a, 0xf7, 0x5b, 0x36, 0xd8, 0x9b, 0xa5,
	0x2e, 0xf9, 0x5f, 0xec, 0x76, 0xfe, 0x16, 0xa2, 0xaa, 0x67, 0x9d, 0x3c, 0xe4, 0x35, 0xb0, 0x94,
	0x7b, 0x42, 0x8e, 0xf6, 0xfc, 0x58, 0xa9, 0xff, 0x44, 0x68, 0x6b, 0x13, 0x4f, 0x4b, 0xc8, 0xb6,
	0xbd, 0x87, 0x76, 0x48, 0x65, 0x67, 0x98, 0xd6, 0xbd, 0xdd, 0x08, 0xbd, 0x9f, 0x77, 0xd5, 0xce,
	0x7f, 0x99, 0xb0, 0x9b, 0x5c, 0xd4, 0x6c, 0xce, 0x97, 0x80, 0x16, 0x4b, 0x40, 0xeb, 0x25, 0xe0,
	0x17, 0x03, 0xf8, 0xdd, 0x00, 0xfe, 0x30, 0x80, 0xe7, 0x06, 0xf0, 0xa7, 0x01, 0xfc, 0x65, 0x00,
	0xad, 0x0d, 0xe0, 0xb7, 0x15, 0xa0, 0xf9, 0x0a, 0xd0, 0x62, 0x05, 0xe8, 0x71, 0xef, 0xfb, 0xba,
	0x07, 0x45, 0xbc, 0x97, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x9c, 0x3b, 0xb1, 0x0f, 0x02,
	0x00, 0x00,
}

func (this *QueryResultRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryResultRequest)
	if !ok {
		that2, ok := that.(QueryResultRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.QueryID != that1.QueryID {
		return false
	}
	if !this.HttpResponse.Equal(that1.HttpResponse) {
		return false
	}
	if !this.Stats.Equal(that1.Stats) {
		return false
	}
	return true
}
func (this *QueryResultResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryResultResponse)
	if !ok {
		that2, ok := that.(QueryResultResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *QueryResultRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&frontendv2pb.QueryResultRequest{")
	s = append(s, "QueryID: "+fmt.Sprintf("%#v", this.QueryID)+",\n")
	if this.HttpResponse != nil {
		s = append(s, "HttpResponse: "+fmt.Sprintf("%#v", this.HttpResponse)+",\n")
	}
	if this.Stats != nil {
		s = append(s, "Stats: "+fmt.Sprintf("%#v", this.Stats)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QueryResultResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&frontendv2pb.QueryResultResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringFrontend(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FrontendForQuerierClient is the client API for FrontendForQuerier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FrontendForQuerierClient interface {
	QueryResult(ctx context.Context, in *QueryResultRequest, opts ...grpc.CallOption) (*QueryResultResponse, error)
}

type frontendForQuerierClient struct {
	cc *grpc.ClientConn
}

func NewFrontendForQuerierClient(cc *grpc.ClientConn) FrontendForQuerierClient {
	return &frontendForQuerierClient{cc}
}

func (c *frontendForQuerierClient) QueryResult(ctx context.Context, in *QueryResultRequest, opts ...grpc.CallOption) (*QueryResultResponse, error) {
	out := new(QueryResultResponse)
	err := c.cc.Invoke(ctx, "/frontendv2pb.FrontendForQuerier/QueryResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FrontendForQuerierServer is the server API for FrontendForQuerier service.
type FrontendForQuerierServer interface {
	QueryResult(context.Context, *QueryResultRequest) (*QueryResultResponse, error)
}

// UnimplementedFrontendForQuerierServer can be embedded to have forward compatible implementations.
type UnimplementedFrontendForQuerierServer struct {
}

func (*UnimplementedFrontendForQuerierServer) QueryResult(ctx context.Context, req *QueryResultRequest) (*QueryResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryResult not implemented")
}

func RegisterFrontendForQuerierServer(s *grpc.Server, srv FrontendForQuerierServer) {
	s.RegisterService(&_FrontendForQuerier_serviceDesc, srv)
}

func _FrontendForQuerier_QueryResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendForQuerierServer).QueryResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/frontendv2pb.FrontendForQuerier/QueryResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendForQuerierServer).QueryResult(ctx, req.(*QueryResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FrontendForQuerier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "frontendv2pb.FrontendForQuerier",
	HandlerType: (*FrontendForQuerierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryResult",
			Handler:    _FrontendForQuerier_QueryResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/lokifrontend/frontend/v2/frontendv2pb/frontend.proto",
}

func (m *QueryResultRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryResultRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryResultRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Stats != nil {
		{
			size, err := m.Stats.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFrontend(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.HttpResponse != nil {
		{
			size, err := m.HttpResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFrontend(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.QueryID != 0 {
		i = encodeVarintFrontend(dAtA, i, uint64(m.QueryID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryResultResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryResultResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryResultResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintFrontend(dAtA []byte, offset int, v uint64) int {
	offset -= sovFrontend(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryResultRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.QueryID != 0 {
		n += 1 + sovFrontend(uint64(m.QueryID))
	}
	if m.HttpResponse != nil {
		l = m.HttpResponse.Size()
		n += 1 + l + sovFrontend(uint64(l))
	}
	if m.Stats != nil {
		l = m.Stats.Size()
		n += 1 + l + sovFrontend(uint64(l))
	}
	return n
}

func (m *QueryResultResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovFrontend(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFrontend(x uint64) (n int) {
	return sovFrontend(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *QueryResultRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueryResultRequest{`,
		`QueryID:` + fmt.Sprintf("%v", this.QueryID) + `,`,
		`HttpResponse:` + strings.Replace(fmt.Sprintf("%v", this.HttpResponse), "HTTPResponse", "httpgrpc.HTTPResponse", 1) + `,`,
		`Stats:` + strings.Replace(fmt.Sprintf("%v", this.Stats), "Stats", "stats.Stats", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QueryResultResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueryResultResponse{`,
		`}`,
	}, "")
	return s
}
func valueToStringFrontend(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *QueryResultRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFrontend
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
			return fmt.Errorf("proto: QueryResultRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryResultRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryID", wireType)
			}
			m.QueryID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QueryID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
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
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HttpResponse == nil {
				m.HttpResponse = &httpgrpc.HTTPResponse{}
			}
			if err := m.HttpResponse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
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
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Stats == nil {
				m.Stats = &stats.Stats{}
			}
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFrontend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFrontend
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
func (m *QueryResultResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFrontend
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
			return fmt.Errorf("proto: QueryResultResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryResultResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFrontend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFrontend
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
func skipFrontend(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFrontend
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
					return 0, ErrIntOverflowFrontend
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
					return 0, ErrIntOverflowFrontend
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
				return 0, ErrInvalidLengthFrontend
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthFrontend
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFrontend
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
				next, err := skipFrontend(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthFrontend
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
	ErrInvalidLengthFrontend = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFrontend   = fmt.Errorf("proto: integer overflow")
)
