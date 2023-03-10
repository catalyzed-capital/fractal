// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fractal/balancecheck/usdcbalance.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Usdcbalance struct {
	Creator         string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Chain           string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	Unit            string `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	Chainaddress    string `protobuf:"bytes,4,opt,name=chainaddress,proto3" json:"chainaddress,omitempty"`
	Requiredamount  string `protobuf:"bytes,5,opt,name=requiredamount,proto3" json:"requiredamount,omitempty"`
	Availableamount string `protobuf:"bytes,6,opt,name=availableamount,proto3" json:"availableamount,omitempty"`
	Unsettledamount string `protobuf:"bytes,7,opt,name=unsettledamount,proto3" json:"unsettledamount,omitempty"`
	Id              uint64 `protobuf:"varint,8,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *Usdcbalance) Reset()         { *m = Usdcbalance{} }
func (m *Usdcbalance) String() string { return proto.CompactTextString(m) }
func (*Usdcbalance) ProtoMessage()    {}
func (*Usdcbalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2d08459d7a4ce65, []int{0}
}
func (m *Usdcbalance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Usdcbalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Usdcbalance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Usdcbalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Usdcbalance.Merge(m, src)
}
func (m *Usdcbalance) XXX_Size() int {
	return m.Size()
}
func (m *Usdcbalance) XXX_DiscardUnknown() {
	xxx_messageInfo_Usdcbalance.DiscardUnknown(m)
}

var xxx_messageInfo_Usdcbalance proto.InternalMessageInfo

func (m *Usdcbalance) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Usdcbalance) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *Usdcbalance) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *Usdcbalance) GetChainaddress() string {
	if m != nil {
		return m.Chainaddress
	}
	return ""
}

func (m *Usdcbalance) GetRequiredamount() string {
	if m != nil {
		return m.Requiredamount
	}
	return ""
}

func (m *Usdcbalance) GetAvailableamount() string {
	if m != nil {
		return m.Availableamount
	}
	return ""
}

func (m *Usdcbalance) GetUnsettledamount() string {
	if m != nil {
		return m.Unsettledamount
	}
	return ""
}

func (m *Usdcbalance) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Usdcbalance)(nil), "fractal.balancecheck.Usdcbalance")
}

func init() {
	proto.RegisterFile("fractal/balancecheck/usdcbalance.proto", fileDescriptor_a2d08459d7a4ce65)
}

var fileDescriptor_a2d08459d7a4ce65 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0xe3, 0x90, 0xb6, 0x60, 0x50, 0x91, 0xac, 0x0e, 0x1e, 0x90, 0x55, 0x75, 0xa8, 0x32,
	0xb5, 0x03, 0x12, 0x0f, 0xc0, 0x23, 0x54, 0x62, 0x61, 0xbb, 0xb1, 0x2f, 0xaa, 0x85, 0x71, 0x8a,
	0x7f, 0x10, 0x8c, 0xbc, 0x01, 0x8f, 0xc5, 0xd8, 0x91, 0x11, 0x25, 0x2f, 0x82, 0x70, 0x53, 0x68,
	0xb3, 0xdd, 0xf3, 0x9d, 0x4f, 0x77, 0x38, 0x74, 0xfe, 0xe0, 0x40, 0x06, 0x30, 0xcb, 0x0a, 0x0c,
	0x58, 0x89, 0x72, 0x8d, 0xf2, 0x71, 0x19, 0xbd, 0x92, 0x1d, 0x58, 0x6c, 0x5c, 0x1d, 0x6a, 0x36,
	0xe9, 0xbc, 0xc5, 0xa1, 0x37, 0x7b, 0xcf, 0xe9, 0xf9, 0xdd, 0xbf, 0xcb, 0x38, 0x1d, 0x49, 0x87,
	0x10, 0x6a, 0xc7, 0xc9, 0x94, 0x94, 0x67, 0xab, 0x7d, 0x64, 0x13, 0x3a, 0x90, 0x6b, 0xd0, 0x96,
	0xe7, 0x89, 0xef, 0x02, 0x63, 0xb4, 0x88, 0x56, 0x07, 0x7e, 0x92, 0x60, 0xba, 0xd9, 0x8c, 0x5e,
	0xa4, 0x12, 0x94, 0x72, 0xe8, 0x3d, 0x2f, 0x52, 0x77, 0xc4, 0xd8, 0x9c, 0x8e, 0x1d, 0x3e, 0x47,
	0xed, 0x50, 0xc1, 0x53, 0x1d, 0x6d, 0xe0, 0x83, 0x64, 0xf5, 0x28, 0x2b, 0xe9, 0x25, 0xbc, 0x80,
	0x36, 0x50, 0x19, 0xec, 0xc4, 0x61, 0x12, 0xfb, 0xf8, 0xd7, 0x8c, 0xd6, 0x63, 0x08, 0xe6, 0xef,
	0xe5, 0x68, 0x67, 0xf6, 0x30, 0x1b, 0xd3, 0x5c, 0x2b, 0x7e, 0x3a, 0x25, 0x65, 0xb1, 0xca, 0xb5,
	0xba, 0xbd, 0xf9, 0x6c, 0x04, 0xd9, 0x36, 0x82, 0x7c, 0x37, 0x82, 0x7c, 0xb4, 0x22, 0xdb, 0xb6,
	0x22, 0xfb, 0x6a, 0x45, 0x76, 0x7f, 0xb5, 0xdf, 0xf6, 0xf5, 0x78, 0xdd, 0xf0, 0xb6, 0x41, 0x5f,
	0x0d, 0xd3, 0xb0, 0xd7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x28, 0x0b, 0x8f, 0xdc, 0x82, 0x01,
	0x00, 0x00,
}

func (m *Usdcbalance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Usdcbalance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Usdcbalance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintUsdcbalance(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Unsettledamount) > 0 {
		i -= len(m.Unsettledamount)
		copy(dAtA[i:], m.Unsettledamount)
		i = encodeVarintUsdcbalance(dAtA, i, uint64(len(m.Unsettledamount)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Availableamount) > 0 {
		i -= len(m.Availableamount)
		copy(dAtA[i:], m.Availableamount)
		i = encodeVarintUsdcbalance(dAtA, i, uint64(len(m.Availableamount)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Requiredamount) > 0 {
		i -= len(m.Requiredamount)
		copy(dAtA[i:], m.Requiredamount)
		i = encodeVarintUsdcbalance(dAtA, i, uint64(len(m.Requiredamount)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Chainaddress) > 0 {
		i -= len(m.Chainaddress)
		copy(dAtA[i:], m.Chainaddress)
		i = encodeVarintUsdcbalance(dAtA, i, uint64(len(m.Chainaddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Unit) > 0 {
		i -= len(m.Unit)
		copy(dAtA[i:], m.Unit)
		i = encodeVarintUsdcbalance(dAtA, i, uint64(len(m.Unit)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintUsdcbalance(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintUsdcbalance(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUsdcbalance(dAtA []byte, offset int, v uint64) int {
	offset -= sovUsdcbalance(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Usdcbalance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovUsdcbalance(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovUsdcbalance(uint64(l))
	}
	l = len(m.Unit)
	if l > 0 {
		n += 1 + l + sovUsdcbalance(uint64(l))
	}
	l = len(m.Chainaddress)
	if l > 0 {
		n += 1 + l + sovUsdcbalance(uint64(l))
	}
	l = len(m.Requiredamount)
	if l > 0 {
		n += 1 + l + sovUsdcbalance(uint64(l))
	}
	l = len(m.Availableamount)
	if l > 0 {
		n += 1 + l + sovUsdcbalance(uint64(l))
	}
	l = len(m.Unsettledamount)
	if l > 0 {
		n += 1 + l + sovUsdcbalance(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovUsdcbalance(uint64(m.Id))
	}
	return n
}

func sovUsdcbalance(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUsdcbalance(x uint64) (n int) {
	return sovUsdcbalance(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Usdcbalance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUsdcbalance
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
			return fmt.Errorf("proto: Usdcbalance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Usdcbalance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsdcbalance
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
				return ErrInvalidLengthUsdcbalance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUsdcbalance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsdcbalance
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
				return ErrInvalidLengthUsdcbalance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUsdcbalance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsdcbalance
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
				return ErrInvalidLengthUsdcbalance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUsdcbalance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Unit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chainaddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsdcbalance
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
				return ErrInvalidLengthUsdcbalance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUsdcbalance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chainaddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requiredamount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsdcbalance
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
				return ErrInvalidLengthUsdcbalance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUsdcbalance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requiredamount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Availableamount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsdcbalance
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
				return ErrInvalidLengthUsdcbalance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUsdcbalance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Availableamount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unsettledamount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsdcbalance
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
				return ErrInvalidLengthUsdcbalance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUsdcbalance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Unsettledamount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsdcbalance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUsdcbalance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUsdcbalance
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
func skipUsdcbalance(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUsdcbalance
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
					return 0, ErrIntOverflowUsdcbalance
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUsdcbalance
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
				return 0, ErrInvalidLengthUsdcbalance
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUsdcbalance
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUsdcbalance
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUsdcbalance        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUsdcbalance          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUsdcbalance = fmt.Errorf("proto: unexpected end of group")
)