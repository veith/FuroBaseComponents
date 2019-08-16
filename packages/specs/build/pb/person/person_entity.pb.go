// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: person_entity.proto

package person

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import furo "../furo"

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

// PersonEntity with Person
type PersonEntity struct {
	// contains a person.Person
	Data *Person `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	// Hateoas links
	Links []*furo.Link `protobuf:"bytes,2,rep,name=links" json:"links,omitempty"`
	// Meta for the response
	Meta                 *furo.Meta `protobuf:"bytes,3,opt,name=meta" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PersonEntity) Reset()         { *m = PersonEntity{} }
func (m *PersonEntity) String() string { return proto.CompactTextString(m) }
func (*PersonEntity) ProtoMessage()    {}
func (*PersonEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_person_entity_cd701cb8e149cdc6, []int{0}
}
func (m *PersonEntity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PersonEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PersonEntity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PersonEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonEntity.Merge(dst, src)
}
func (m *PersonEntity) XXX_Size() int {
	return m.Size()
}
func (m *PersonEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonEntity.DiscardUnknown(m)
}

var xxx_messageInfo_PersonEntity proto.InternalMessageInfo

func (m *PersonEntity) GetData() *Person {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PersonEntity) GetLinks() []*furo.Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *PersonEntity) GetMeta() *furo.Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*PersonEntity)(nil), "person.PersonEntity")
}
func (m *PersonEntity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PersonEntity) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPersonEntity(dAtA, i, uint64(m.Data.Size()))
		n1, err := m.Data.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Links) > 0 {
		for _, msg := range m.Links {
			dAtA[i] = 0x12
			i++
			i = encodeVarintPersonEntity(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Meta != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPersonEntity(dAtA, i, uint64(m.Meta.Size()))
		n2, err := m.Meta.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintPersonEntity(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PersonEntity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovPersonEntity(uint64(l))
	}
	if len(m.Links) > 0 {
		for _, e := range m.Links {
			l = e.Size()
			n += 1 + l + sovPersonEntity(uint64(l))
		}
	}
	if m.Meta != nil {
		l = m.Meta.Size()
		n += 1 + l + sovPersonEntity(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPersonEntity(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPersonEntity(x uint64) (n int) {
	return sovPersonEntity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PersonEntity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPersonEntity
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
			return fmt.Errorf("proto: PersonEntity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PersonEntity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPersonEntity
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
				return ErrInvalidLengthPersonEntity
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &Person{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Links", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPersonEntity
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
				return ErrInvalidLengthPersonEntity
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Links = append(m.Links, &furo.Link{})
			if err := m.Links[len(m.Links)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPersonEntity
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
				return ErrInvalidLengthPersonEntity
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Meta == nil {
				m.Meta = &furo.Meta{}
			}
			if err := m.Meta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPersonEntity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPersonEntity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPersonEntity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPersonEntity
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
					return 0, ErrIntOverflowPersonEntity
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
					return 0, ErrIntOverflowPersonEntity
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
				return 0, ErrInvalidLengthPersonEntity
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPersonEntity
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
				next, err := skipPersonEntity(dAtA[start:])
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
	ErrInvalidLengthPersonEntity = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPersonEntity   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("person_entity.proto", fileDescriptor_person_entity_cd701cb8e149cdc6) }

var fileDescriptor_person_entity_cd701cb8e149cdc6 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0x8b, 0x4f, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x83, 0x08, 0x4a, 0xf1, 0x40, 0x68, 0x88, 0xa8, 0x14, 0x7f, 0x5a, 0x69, 0x51, 0xbe, 0x7e,
	0x6e, 0x6a, 0x49, 0x22, 0x8a, 0x40, 0x4e, 0x66, 0x5e, 0x36, 0x44, 0x40, 0xa9, 0x84, 0x8b, 0x27,
	0x00, 0xac, 0xc3, 0x15, 0x6c, 0x9a, 0x90, 0x12, 0x17, 0x4b, 0x4a, 0x62, 0x49, 0xa2, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0xb7, 0x11, 0x9f, 0x1e, 0xd4, 0x38, 0x88, 0x9a, 0x20, 0xb0, 0x9c, 0x90, 0x02,
	0x17, 0x2b, 0xc8, 0x84, 0x62, 0x09, 0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x2e, 0x3d, 0x90, 0xa1,
	0x7a, 0x3e, 0x99, 0x79, 0xd9, 0x41, 0x10, 0x09, 0x21, 0x39, 0x2e, 0x16, 0x90, 0xa5, 0x12, 0xcc,
	0x60, 0x53, 0xa0, 0x0a, 0x7c, 0x53, 0x4b, 0x12, 0x83, 0xc0, 0xe2, 0x4e, 0x3c, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0x63, 0x12, 0x1b, 0xd8, 0x29, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0x98, 0x07, 0x2e, 0xd9, 0x00, 0x00, 0x00,
}
