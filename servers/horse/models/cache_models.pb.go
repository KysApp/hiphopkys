// Code generated by protoc-gen-gogo.
// source: cache_models.proto
// DO NOT EDIT!

/*
	Package models is a generated protocol buffer package.

	It is generated from these files:
		cache_models.proto

	It has these top-level messages:
		AppointmentPlayerCacheModel
		AppointmentRoomCacheModel
*/
package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AppointmentPlayerCacheModel struct {
	AppointmentId        string `protobuf:"bytes,1,opt,name=appointment_id,proto3" json:"appointment_id,omitempty"`
	UserId               string `protobuf:"bytes,2,opt,name=user_id,proto3" json:"user_id,omitempty"`
	AppointmentTimestamp int64  `protobuf:"varint,3,opt,name=appointment_timestamp,proto3" json:"appointment_timestamp,omitempty"`
	LimiteCanPlayCount   int32  `protobuf:"varint,4,opt,name=limite_can_play_count,proto3" json:"limite_can_play_count,omitempty"`
}

func (m *AppointmentPlayerCacheModel) Reset()         { *m = AppointmentPlayerCacheModel{} }
func (m *AppointmentPlayerCacheModel) String() string { return proto.CompactTextString(m) }
func (*AppointmentPlayerCacheModel) ProtoMessage()    {}

type AppointmentRoomCacheModel struct {
	AppointmentId string                         `protobuf:"bytes,1,opt,name=appointment_id,proto3" json:"appointment_id,omitempty"`
	PlayCount     int32                          `protobuf:"varint,2,opt,name=play_count,proto3" json:"play_count,omitempty"`
	ModelIdArray  []*AppointmentPlayerCacheModel `protobuf:"bytes,3,rep,name=model_id_array" json:"model_id_array,omitempty"`
}

func (m *AppointmentRoomCacheModel) Reset()         { *m = AppointmentRoomCacheModel{} }
func (m *AppointmentRoomCacheModel) String() string { return proto.CompactTextString(m) }
func (*AppointmentRoomCacheModel) ProtoMessage()    {}

func (m *AppointmentRoomCacheModel) GetModelIdArray() []*AppointmentPlayerCacheModel {
	if m != nil {
		return m.ModelIdArray
	}
	return nil
}

func (m *AppointmentPlayerCacheModel) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *AppointmentPlayerCacheModel) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AppointmentId) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintCacheModels(data, i, uint64(len(m.AppointmentId)))
		i += copy(data[i:], m.AppointmentId)
	}
	if len(m.UserId) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintCacheModels(data, i, uint64(len(m.UserId)))
		i += copy(data[i:], m.UserId)
	}
	if m.AppointmentTimestamp != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintCacheModels(data, i, uint64(m.AppointmentTimestamp))
	}
	if m.LimiteCanPlayCount != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintCacheModels(data, i, uint64(m.LimiteCanPlayCount))
	}
	return i, nil
}

func (m *AppointmentRoomCacheModel) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *AppointmentRoomCacheModel) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AppointmentId) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintCacheModels(data, i, uint64(len(m.AppointmentId)))
		i += copy(data[i:], m.AppointmentId)
	}
	if m.PlayCount != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintCacheModels(data, i, uint64(m.PlayCount))
	}
	if len(m.ModelIdArray) > 0 {
		for _, msg := range m.ModelIdArray {
			data[i] = 0x1a
			i++
			i = encodeVarintCacheModels(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64CacheModels(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32CacheModels(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCacheModels(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *AppointmentPlayerCacheModel) Size() (n int) {
	var l int
	_ = l
	l = len(m.AppointmentId)
	if l > 0 {
		n += 1 + l + sovCacheModels(uint64(l))
	}
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovCacheModels(uint64(l))
	}
	if m.AppointmentTimestamp != 0 {
		n += 1 + sovCacheModels(uint64(m.AppointmentTimestamp))
	}
	if m.LimiteCanPlayCount != 0 {
		n += 1 + sovCacheModels(uint64(m.LimiteCanPlayCount))
	}
	return n
}

func (m *AppointmentRoomCacheModel) Size() (n int) {
	var l int
	_ = l
	l = len(m.AppointmentId)
	if l > 0 {
		n += 1 + l + sovCacheModels(uint64(l))
	}
	if m.PlayCount != 0 {
		n += 1 + sovCacheModels(uint64(m.PlayCount))
	}
	if len(m.ModelIdArray) > 0 {
		for _, e := range m.ModelIdArray {
			l = e.Size()
			n += 1 + l + sovCacheModels(uint64(l))
		}
	}
	return n
}

func sovCacheModels(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCacheModels(x uint64) (n int) {
	return sovCacheModels(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AppointmentPlayerCacheModel) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCacheModels
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AppointmentPlayerCacheModel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppointmentPlayerCacheModel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppointmentId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCacheModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCacheModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppointmentId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCacheModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCacheModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppointmentTimestamp", wireType)
			}
			m.AppointmentTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCacheModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.AppointmentTimestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimiteCanPlayCount", wireType)
			}
			m.LimiteCanPlayCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCacheModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.LimiteCanPlayCount |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCacheModels(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCacheModels
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
func (m *AppointmentRoomCacheModel) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCacheModels
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AppointmentRoomCacheModel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppointmentRoomCacheModel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppointmentId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCacheModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCacheModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppointmentId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayCount", wireType)
			}
			m.PlayCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCacheModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.PlayCount |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModelIdArray", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCacheModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCacheModels
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModelIdArray = append(m.ModelIdArray, &AppointmentPlayerCacheModel{})
			if err := m.ModelIdArray[len(m.ModelIdArray)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCacheModels(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCacheModels
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
func skipCacheModels(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCacheModels
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowCacheModels
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowCacheModels
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCacheModels
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCacheModels
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipCacheModels(data[start:])
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
	ErrInvalidLengthCacheModels = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCacheModels   = fmt.Errorf("proto: integer overflow")
)
