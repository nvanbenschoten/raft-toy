// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transport.proto

/*
	Package transportpb is a generated protocol buffer package.

	It is generated from these files:
		transport.proto

	It has these top-level messages:
		RaftMsg
*/
package transportpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import raftpb "go.etcd.io/etcd/raft/v3/raftpb"
import config "github.com/nvanbenschoten/rafttoy/config"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RaftMsg struct {
	// Epoch indicates the test epoch that this message was sent from. Raft
	// state is reset when an epoch increment is detected and Raft messages
	// from previous epochs are discarded.
	Epoch config.TestEpoch `protobuf:"bytes,1,opt,name=epoch" json:"epoch"`
	// Msgs is a group of wrapped Raft message.
	Msgs []raftpb.Message `protobuf:"bytes,2,rep,name=msgs" json:"msgs"`
}

func (m *RaftMsg) Reset()                    { *m = RaftMsg{} }
func (m *RaftMsg) String() string            { return proto.CompactTextString(m) }
func (*RaftMsg) ProtoMessage()               {}
func (*RaftMsg) Descriptor() ([]byte, []int) { return fileDescriptorTransport, []int{0} }

func (m *RaftMsg) GetEpoch() config.TestEpoch {
	if m != nil {
		return m.Epoch
	}
	return config.TestEpoch{}
}

func (m *RaftMsg) GetMsgs() []raftpb.Message {
	if m != nil {
		return m.Msgs
	}
	return nil
}

func init() {
	proto.RegisterType((*RaftMsg)(nil), "transportpb.RaftMsg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RaftService service

type RaftServiceClient interface {
	RaftMessage(ctx context.Context, opts ...grpc.CallOption) (RaftService_RaftMessageClient, error)
}

type raftServiceClient struct {
	cc *grpc.ClientConn
}

func NewRaftServiceClient(cc *grpc.ClientConn) RaftServiceClient {
	return &raftServiceClient{cc}
}

func (c *raftServiceClient) RaftMessage(ctx context.Context, opts ...grpc.CallOption) (RaftService_RaftMessageClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RaftService_serviceDesc.Streams[0], c.cc, "/transportpb.RaftService/RaftMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &raftServiceRaftMessageClient{stream}
	return x, nil
}

type RaftService_RaftMessageClient interface {
	Send(*RaftMsg) error
	Recv() (*RaftMsg, error)
	grpc.ClientStream
}

type raftServiceRaftMessageClient struct {
	grpc.ClientStream
}

func (x *raftServiceRaftMessageClient) Send(m *RaftMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *raftServiceRaftMessageClient) Recv() (*RaftMsg, error) {
	m := new(RaftMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RaftService service

type RaftServiceServer interface {
	RaftMessage(RaftService_RaftMessageServer) error
}

func RegisterRaftServiceServer(s *grpc.Server, srv RaftServiceServer) {
	s.RegisterService(&_RaftService_serviceDesc, srv)
}

func _RaftService_RaftMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RaftServiceServer).RaftMessage(&raftServiceRaftMessageServer{stream})
}

type RaftService_RaftMessageServer interface {
	Send(*RaftMsg) error
	Recv() (*RaftMsg, error)
	grpc.ServerStream
}

type raftServiceRaftMessageServer struct {
	grpc.ServerStream
}

func (x *raftServiceRaftMessageServer) Send(m *RaftMsg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *raftServiceRaftMessageServer) Recv() (*RaftMsg, error) {
	m := new(RaftMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RaftService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transportpb.RaftService",
	HandlerType: (*RaftServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RaftMessage",
			Handler:       _RaftService_RaftMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "transport.proto",
}

func (m *RaftMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaftMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintTransport(dAtA, i, uint64(m.Epoch.Size()))
	n1, err := m.Epoch.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Msgs) > 0 {
		for _, msg := range m.Msgs {
			dAtA[i] = 0x12
			i++
			i = encodeVarintTransport(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintTransport(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RaftMsg) Size() (n int) {
	var l int
	_ = l
	l = m.Epoch.Size()
	n += 1 + l + sovTransport(uint64(l))
	if len(m.Msgs) > 0 {
		for _, e := range m.Msgs {
			l = e.Size()
			n += 1 + l + sovTransport(uint64(l))
		}
	}
	return n
}

func sovTransport(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTransport(x uint64) (n int) {
	return sovTransport(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RaftMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RaftMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTransport
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Epoch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTransport
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msgs = append(m.Msgs, raftpb.Message{})
			if err := m.Msgs[len(m.Msgs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransport
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
func skipTransport(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransport
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
					return 0, ErrIntOverflowTransport
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
					return 0, ErrIntOverflowTransport
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTransport
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTransport
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
				next, err := skipTransport(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthTransport = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransport   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("transport.proto", fileDescriptorTransport) }

var fileDescriptorTransport = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0x28, 0x20, 0x25, 0x43, 0x21, 0xea, 0x50, 0x65, 0x08, 0x55, 0x07, 0x14, 0x06,
	0x1c, 0x14, 0xc4, 0x8c, 0x54, 0x89, 0xb1, 0x0c, 0x85, 0x17, 0xb0, 0x8d, 0xe3, 0x64, 0xa8, 0xcf,
	0xb2, 0x8f, 0x4a, 0xbc, 0x05, 0x8f, 0xd5, 0x91, 0x27, 0x40, 0x28, 0xbc, 0x08, 0x8a, 0x6d, 0x21,
	0x86, 0x2e, 0xfe, 0xed, 0xdf, 0xdf, 0xe9, 0xfe, 0xbb, 0x74, 0x8a, 0x96, 0x69, 0x67, 0xc0, 0x22,
	0x35, 0x16, 0x10, 0xf2, 0xec, 0xcf, 0x30, 0xbc, 0x98, 0x29, 0x50, 0xe0, 0xfd, 0x7a, 0xbc, 0x05,
	0xa4, 0xb8, 0x52, 0x40, 0x25, 0x8a, 0x57, 0xda, 0x43, 0x3d, 0x6a, 0x6d, 0x59, 0x8b, 0xfe, 0x30,
	0xdc, 0x4b, 0xe4, 0xee, 0x55, 0x8f, 0xdd, 0x1b, 0xa7, 0x02, 0xb6, 0xb5, 0xde, 0x31, 0xcd, 0xa5,
	0x76, 0xa2, 0x03, 0x94, 0xda, 0x53, 0x08, 0xef, 0xb5, 0x00, 0xdd, 0xf6, 0x2a, 0x4a, 0x28, 0x5b,
	0x8a, 0xf4, 0x6c, 0xc3, 0x5a, 0x5c, 0x3b, 0x95, 0xdf, 0xa4, 0x27, 0xd2, 0x80, 0xe8, 0xe6, 0x64,
	0x41, 0xaa, 0xac, 0xb9, 0xa0, 0x11, 0x7c, 0x91, 0x0e, 0x1f, 0xc7, 0x8f, 0xd5, 0x64, 0xff, 0x75,
	0x99, 0x6c, 0x02, 0x95, 0x5f, 0xa7, 0x93, 0xad, 0x53, 0x6e, 0x7e, 0xb4, 0x38, 0xae, 0xb2, 0x66,
	0x4a, 0x43, 0x24, 0xba, 0x96, 0xce, 0x31, 0x25, 0x23, 0xeb, 0x91, 0xe6, 0x29, 0xcd, 0xc6, 0x26,
	0xcf, 0xd2, 0xee, 0x7a, 0x21, 0xf3, 0x87, 0xf0, 0x8c, 0x64, 0x3e, 0xa3, 0xff, 0xb6, 0x40, 0x63,
	0x9a, 0xe2, 0xa0, 0xbb, 0x4c, 0x2a, 0x72, 0x4b, 0x56, 0xe7, 0xfb, 0xa1, 0x24, 0x9f, 0x43, 0x49,
	0xbe, 0x87, 0x92, 0x7c, 0xfc, 0x94, 0x09, 0x3f, 0xf5, 0xd3, 0xdc, 0xfd, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x00, 0x6c, 0x3a, 0xcd, 0x62, 0x01, 0x00, 0x00,
}
