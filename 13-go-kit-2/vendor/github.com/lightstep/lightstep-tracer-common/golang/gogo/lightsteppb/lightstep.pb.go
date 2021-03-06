// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/lightstep/lightstep-tracer-common/tmpgen/lightstep.proto

/*
	Package lightsteppb is a generated protocol buffer package.

	It is generated from these files:
		github.com/lightstep/lightstep-tracer-common/tmpgen/lightstep.proto

	It has these top-level messages:
		BinaryCarrier
		BasicTracerCarrier
*/
package lightsteppb // import "github.com/lightstep/lightstep-tracer-common/golang/gogo/lightsteppb"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import encoding_binary "encoding/binary"

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

// The standard carrier for binary context propagation into LightStep.
type BinaryCarrier struct {
	// "text_ctx" was deprecated following lightstep-tracer-cpp-0.36
	DeprecatedTextCtx [][]byte `protobuf:"bytes,1,rep,name=deprecated_text_ctx,json=deprecatedTextCtx" json:"deprecated_text_ctx,omitempty"`
	// The Opentracing "basictracer" proto.
	BasicCtx *BasicTracerCarrier `protobuf:"bytes,2,opt,name=basic_ctx,json=basicCtx" json:"basic_ctx,omitempty"`
}

func (m *BinaryCarrier) Reset()                    { *m = BinaryCarrier{} }
func (m *BinaryCarrier) String() string            { return proto.CompactTextString(m) }
func (*BinaryCarrier) ProtoMessage()               {}
func (*BinaryCarrier) Descriptor() ([]byte, []int) { return fileDescriptorLightstep, []int{0} }

func (m *BinaryCarrier) GetDeprecatedTextCtx() [][]byte {
	if m != nil {
		return m.DeprecatedTextCtx
	}
	return nil
}

func (m *BinaryCarrier) GetBasicCtx() *BasicTracerCarrier {
	if m != nil {
		return m.BasicCtx
	}
	return nil
}

// Copy of https://github.com/opentracing/basictracer-go/blob/master/wire/wire.proto
type BasicTracerCarrier struct {
	TraceId      uint64            `protobuf:"fixed64,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SpanId       uint64            `protobuf:"fixed64,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	Sampled      bool              `protobuf:"varint,3,opt,name=sampled,proto3" json:"sampled,omitempty"`
	BaggageItems map[string]string `protobuf:"bytes,4,rep,name=baggage_items,json=baggageItems" json:"baggage_items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *BasicTracerCarrier) Reset()                    { *m = BasicTracerCarrier{} }
func (m *BasicTracerCarrier) String() string            { return proto.CompactTextString(m) }
func (*BasicTracerCarrier) ProtoMessage()               {}
func (*BasicTracerCarrier) Descriptor() ([]byte, []int) { return fileDescriptorLightstep, []int{1} }

func (m *BasicTracerCarrier) GetTraceId() uint64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

func (m *BasicTracerCarrier) GetSpanId() uint64 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

func (m *BasicTracerCarrier) GetSampled() bool {
	if m != nil {
		return m.Sampled
	}
	return false
}

func (m *BasicTracerCarrier) GetBaggageItems() map[string]string {
	if m != nil {
		return m.BaggageItems
	}
	return nil
}

func init() {
	proto.RegisterType((*BinaryCarrier)(nil), "lightstep.BinaryCarrier")
	proto.RegisterType((*BasicTracerCarrier)(nil), "lightstep.BasicTracerCarrier")
}
func (m *BinaryCarrier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BinaryCarrier) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.DeprecatedTextCtx) > 0 {
		for _, b := range m.DeprecatedTextCtx {
			dAtA[i] = 0xa
			i++
			i = encodeVarintLightstep(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if m.BasicCtx != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLightstep(dAtA, i, uint64(m.BasicCtx.Size()))
		n1, err := m.BasicCtx.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *BasicTracerCarrier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasicTracerCarrier) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TraceId != 0 {
		dAtA[i] = 0x9
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.TraceId))
		i += 8
	}
	if m.SpanId != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.SpanId))
		i += 8
	}
	if m.Sampled {
		dAtA[i] = 0x18
		i++
		if m.Sampled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.BaggageItems) > 0 {
		for k, _ := range m.BaggageItems {
			dAtA[i] = 0x22
			i++
			v := m.BaggageItems[k]
			mapSize := 1 + len(k) + sovLightstep(uint64(len(k))) + 1 + len(v) + sovLightstep(uint64(len(v)))
			i = encodeVarintLightstep(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintLightstep(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintLightstep(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func encodeVarintLightstep(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BinaryCarrier) Size() (n int) {
	var l int
	_ = l
	if len(m.DeprecatedTextCtx) > 0 {
		for _, b := range m.DeprecatedTextCtx {
			l = len(b)
			n += 1 + l + sovLightstep(uint64(l))
		}
	}
	if m.BasicCtx != nil {
		l = m.BasicCtx.Size()
		n += 1 + l + sovLightstep(uint64(l))
	}
	return n
}

func (m *BasicTracerCarrier) Size() (n int) {
	var l int
	_ = l
	if m.TraceId != 0 {
		n += 9
	}
	if m.SpanId != 0 {
		n += 9
	}
	if m.Sampled {
		n += 2
	}
	if len(m.BaggageItems) > 0 {
		for k, v := range m.BaggageItems {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovLightstep(uint64(len(k))) + 1 + len(v) + sovLightstep(uint64(len(v)))
			n += mapEntrySize + 1 + sovLightstep(uint64(mapEntrySize))
		}
	}
	return n
}

func sovLightstep(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLightstep(x uint64) (n int) {
	return sovLightstep(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BinaryCarrier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLightstep
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
			return fmt.Errorf("proto: BinaryCarrier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BinaryCarrier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeprecatedTextCtx", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLightstep
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLightstep
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeprecatedTextCtx = append(m.DeprecatedTextCtx, make([]byte, postIndex-iNdEx))
			copy(m.DeprecatedTextCtx[len(m.DeprecatedTextCtx)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasicCtx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLightstep
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
				return ErrInvalidLengthLightstep
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BasicCtx == nil {
				m.BasicCtx = &BasicTracerCarrier{}
			}
			if err := m.BasicCtx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLightstep(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLightstep
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
func (m *BasicTracerCarrier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLightstep
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
			return fmt.Errorf("proto: BasicTracerCarrier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasicTracerCarrier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceId", wireType)
			}
			m.TraceId = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.TraceId = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpanId", wireType)
			}
			m.SpanId = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.SpanId = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sampled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLightstep
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Sampled = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaggageItems", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLightstep
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
				return ErrInvalidLengthLightstep
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaggageItems == nil {
				m.BaggageItems = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLightstep
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLightstep
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthLightstep
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLightstep
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthLightstep
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipLightstep(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthLightstep
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.BaggageItems[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLightstep(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLightstep
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
func skipLightstep(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLightstep
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
					return 0, ErrIntOverflowLightstep
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
					return 0, ErrIntOverflowLightstep
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
				return 0, ErrInvalidLengthLightstep
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLightstep
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
				next, err := skipLightstep(dAtA[start:])
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
	ErrInvalidLengthLightstep = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLightstep   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/lightstep/lightstep-tracer-common/tmpgen/lightstep.proto", fileDescriptorLightstep)
}

var fileDescriptorLightstep = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xdd, 0x46, 0xdb, 0x66, 0xdb, 0x82, 0x5d, 0x05, 0xa3, 0xd0, 0x10, 0x7a, 0xca, 0xa5,
	0x29, 0xd4, 0x8b, 0xf4, 0x22, 0x24, 0x78, 0xe8, 0x35, 0xf4, 0xe4, 0x25, 0x6c, 0x92, 0x21, 0x0d,
	0x36, 0xc9, 0xb2, 0x99, 0x4a, 0x8a, 0x2f, 0xe1, 0x63, 0x79, 0xf4, 0x11, 0xa4, 0x3e, 0x85, 0x37,
	0xc9, 0x56, 0x1b, 0xa1, 0xe0, 0x6d, 0xfe, 0xf9, 0xff, 0xe1, 0xff, 0x96, 0xa5, 0x5e, 0x92, 0xe2,
	0x6a, 0x13, 0x3a, 0x51, 0x91, 0x4d, 0xd7, 0x69, 0xb2, 0xc2, 0x12, 0x41, 0x34, 0xd3, 0x04, 0x25,
	0x8f, 0x40, 0x4e, 0xa2, 0x22, 0xcb, 0x8a, 0x7c, 0x8a, 0x99, 0x48, 0x20, 0x6f, 0x6c, 0x47, 0xc8,
	0x02, 0x0b, 0xa6, 0x1f, 0x16, 0xe3, 0x17, 0x3a, 0x70, 0xd3, 0x9c, 0xcb, 0xad, 0xc7, 0xa5, 0x4c,
	0x41, 0x32, 0x87, 0x5e, 0xc4, 0x20, 0x24, 0x44, 0x1c, 0x21, 0x0e, 0x10, 0x2a, 0x0c, 0x22, 0xac,
	0x0c, 0x62, 0x69, 0x76, 0xdf, 0x1f, 0x36, 0xd6, 0x12, 0x2a, 0xf4, 0xb0, 0x62, 0x73, 0xaa, 0x87,
	0xbc, 0x4c, 0x23, 0x95, 0x6a, 0x59, 0xc4, 0xee, 0xcd, 0x46, 0x4e, 0x53, 0xe8, 0xd6, 0xde, 0x52,
	0x41, 0xfd, 0x34, 0xf8, 0x5d, 0x95, 0xf7, 0xb0, 0x1a, 0x7f, 0x11, 0xca, 0x8e, 0x03, 0xec, 0x9a,
	0x76, 0xd5, 0x33, 0x82, 0x34, 0x36, 0x88, 0x45, 0xec, 0xb6, 0xdf, 0x51, 0x7a, 0x11, 0xb3, 0x2b,
	0xda, 0x29, 0x05, 0xcf, 0x6b, 0xa7, 0xa5, 0x9c, 0x76, 0x2d, 0x17, 0x31, 0x33, 0x68, 0xa7, 0xe4,
	0x99, 0x58, 0x43, 0x6c, 0x68, 0x16, 0xb1, 0xbb, 0xfe, 0xaf, 0x64, 0x4b, 0x3a, 0x08, 0x79, 0x92,
	0xf0, 0x04, 0x82, 0x14, 0x21, 0x2b, 0x8d, 0x53, 0x4b, 0xb3, 0x7b, 0xb3, 0xe9, 0xbf, 0x90, 0x8e,
	0xbb, 0x3f, 0x59, 0xd4, 0x17, 0x0f, 0x39, 0xca, 0xad, 0xdf, 0x0f, 0xff, 0xac, 0x6e, 0xee, 0xe9,
	0xf0, 0x28, 0xc2, 0xce, 0xa9, 0xf6, 0x04, 0x5b, 0xc5, 0xac, 0xfb, 0xf5, 0xc8, 0x2e, 0xe9, 0xd9,
	0x33, 0x5f, 0x6f, 0x40, 0xd1, 0xea, 0xfe, 0x5e, 0xcc, 0x5b, 0x77, 0xc4, 0x1d, 0xbd, 0xed, 0x4c,
	0xf2, 0xbe, 0x33, 0xc9, 0xc7, 0xce, 0x24, 0xaf, 0x9f, 0xe6, 0xc9, 0x63, 0xef, 0x00, 0x24, 0xc2,
	0xb0, 0xad, 0x7e, 0xea, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x95, 0x72, 0x6a, 0xa9, 0xf0, 0x01,
	0x00, 0x00,
}
