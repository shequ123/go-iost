// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: consensus/synchronizer/pb/message.proto

package msgpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RequireType int32

const (
	RequireType_GETBLOCKHASHES         RequireType = 0
	RequireType_GETBLOCKHASHESBYNUMBER RequireType = 1
)

var RequireType_name = map[int32]string{
	0: "GETBLOCKHASHES",
	1: "GETBLOCKHASHESBYNUMBER",
}
var RequireType_value = map[string]int32{
	"GETBLOCKHASHES":         0,
	"GETBLOCKHASHESBYNUMBER": 1,
}

func (x RequireType) String() string {
	return proto.EnumName(RequireType_name, int32(x))
}
func (RequireType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_2909aa68f10e2b1c, []int{0}
}

type BlockInfo struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Hash                 []byte   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockInfo) Reset()         { *m = BlockInfo{} }
func (m *BlockInfo) String() string { return proto.CompactTextString(m) }
func (*BlockInfo) ProtoMessage()    {}
func (*BlockInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_2909aa68f10e2b1c, []int{0}
}
func (m *BlockInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BlockInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockInfo.Merge(dst, src)
}
func (m *BlockInfo) XXX_Size() int {
	return m.Size()
}
func (m *BlockInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BlockInfo proto.InternalMessageInfo

func (m *BlockInfo) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *BlockInfo) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type BlockHashQuery struct {
	ReqType              RequireType `protobuf:"varint,1,opt,name=reqType,proto3,enum=msgpb.RequireType" json:"reqType,omitempty"`
	Start                int64       `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End                  int64       `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	Nums                 []int64     `protobuf:"varint,4,rep,packed,name=nums" json:"nums,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BlockHashQuery) Reset()         { *m = BlockHashQuery{} }
func (m *BlockHashQuery) String() string { return proto.CompactTextString(m) }
func (*BlockHashQuery) ProtoMessage()    {}
func (*BlockHashQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_2909aa68f10e2b1c, []int{1}
}
func (m *BlockHashQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHashQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHashQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BlockHashQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHashQuery.Merge(dst, src)
}
func (m *BlockHashQuery) XXX_Size() int {
	return m.Size()
}
func (m *BlockHashQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHashQuery.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHashQuery proto.InternalMessageInfo

func (m *BlockHashQuery) GetReqType() RequireType {
	if m != nil {
		return m.ReqType
	}
	return RequireType_GETBLOCKHASHES
}

func (m *BlockHashQuery) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *BlockHashQuery) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *BlockHashQuery) GetNums() []int64 {
	if m != nil {
		return m.Nums
	}
	return nil
}

type BlockHashResponse struct {
	BlockInfos           []*BlockInfo `protobuf:"bytes,1,rep,name=blockInfos" json:"blockInfos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BlockHashResponse) Reset()         { *m = BlockHashResponse{} }
func (m *BlockHashResponse) String() string { return proto.CompactTextString(m) }
func (*BlockHashResponse) ProtoMessage()    {}
func (*BlockHashResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_2909aa68f10e2b1c, []int{2}
}
func (m *BlockHashResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHashResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHashResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BlockHashResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHashResponse.Merge(dst, src)
}
func (m *BlockHashResponse) XXX_Size() int {
	return m.Size()
}
func (m *BlockHashResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHashResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHashResponse proto.InternalMessageInfo

func (m *BlockHashResponse) GetBlockInfos() []*BlockInfo {
	if m != nil {
		return m.BlockInfos
	}
	return nil
}

type SyncHeight struct {
	Height               int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncHeight) Reset()         { *m = SyncHeight{} }
func (m *SyncHeight) String() string { return proto.CompactTextString(m) }
func (*SyncHeight) ProtoMessage()    {}
func (*SyncHeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_2909aa68f10e2b1c, []int{3}
}
func (m *SyncHeight) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SyncHeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SyncHeight.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SyncHeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncHeight.Merge(dst, src)
}
func (m *SyncHeight) XXX_Size() int {
	return m.Size()
}
func (m *SyncHeight) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncHeight.DiscardUnknown(m)
}

var xxx_messageInfo_SyncHeight proto.InternalMessageInfo

func (m *SyncHeight) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *SyncHeight) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterType((*BlockInfo)(nil), "msgpb.BlockInfo")
	proto.RegisterType((*BlockHashQuery)(nil), "msgpb.BlockHashQuery")
	proto.RegisterType((*BlockHashResponse)(nil), "msgpb.BlockHashResponse")
	proto.RegisterType((*SyncHeight)(nil), "msgpb.SyncHeight")
	proto.RegisterEnum("msgpb.RequireType", RequireType_name, RequireType_value)
}
func (m *BlockInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Number != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Number))
	}
	if len(m.Hash) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Hash)))
		i += copy(dAtA[i:], m.Hash)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BlockHashQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHashQuery) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ReqType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.ReqType))
	}
	if m.Start != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Start))
	}
	if m.End != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.End))
	}
	if len(m.Nums) > 0 {
		dAtA2 := make([]byte, len(m.Nums)*10)
		var j1 int
		for _, num1 := range m.Nums {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x22
		i++
		i = encodeVarintMessage(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BlockHashResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHashResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.BlockInfos) > 0 {
		for _, msg := range m.BlockInfos {
			dAtA[i] = 0xa
			i++
			i = encodeVarintMessage(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SyncHeight) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncHeight) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Height))
	}
	if m.Time != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Time))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BlockInfo) Size() (n int) {
	var l int
	_ = l
	if m.Number != 0 {
		n += 1 + sovMessage(uint64(m.Number))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BlockHashQuery) Size() (n int) {
	var l int
	_ = l
	if m.ReqType != 0 {
		n += 1 + sovMessage(uint64(m.ReqType))
	}
	if m.Start != 0 {
		n += 1 + sovMessage(uint64(m.Start))
	}
	if m.End != 0 {
		n += 1 + sovMessage(uint64(m.End))
	}
	if len(m.Nums) > 0 {
		l = 0
		for _, e := range m.Nums {
			l += sovMessage(uint64(e))
		}
		n += 1 + sovMessage(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BlockHashResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.BlockInfos) > 0 {
		for _, e := range m.BlockInfos {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SyncHeight) Size() (n int) {
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovMessage(uint64(m.Height))
	}
	if m.Time != 0 {
		n += 1 + sovMessage(uint64(m.Time))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMessage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: BlockInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *BlockHashQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: BlockHashQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHashQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReqType", wireType)
			}
			m.ReqType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReqType |= (RequireType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			m.End = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.End |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMessage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Nums = append(m.Nums, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMessage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthMessage
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMessage
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Nums = append(m.Nums, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Nums", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *BlockHashResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: BlockHashResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHashResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockInfos = append(m.BlockInfos, &BlockInfo{})
			if err := m.BlockInfos[len(m.BlockInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *SyncHeight) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: SyncHeight: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncHeight: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessage
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
				next, err := skipMessage(dAtA[start:])
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
	ErrInvalidLengthMessage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("consensus/synchronizer/pb/message.proto", fileDescriptor_message_2909aa68f10e2b1c)
}

var fileDescriptor_message_2909aa68f10e2b1c = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcd, 0x4e, 0x83, 0x40,
	0x14, 0x85, 0x3b, 0xd2, 0xd6, 0x78, 0x6b, 0x1a, 0x9c, 0x98, 0x86, 0xb8, 0x20, 0x84, 0x8d, 0xc4,
	0x98, 0xd6, 0xd4, 0x85, 0x6e, 0x5c, 0x88, 0x21, 0x62, 0xfc, 0x8b, 0xd3, 0xba, 0x70, 0x09, 0x78,
	0x2d, 0x44, 0x19, 0xe8, 0x0c, 0x2c, 0xe8, 0x93, 0xf8, 0x48, 0x2e, 0x7d, 0x04, 0x53, 0x5f, 0xc4,
	0x30, 0xfd, 0x49, 0xdd, 0x9d, 0x73, 0xb9, 0x97, 0x73, 0xbe, 0x0c, 0x1c, 0x46, 0x19, 0x97, 0xc8,
	0x65, 0x29, 0x07, 0xb2, 0xe2, 0x51, 0x2c, 0x32, 0x9e, 0xcc, 0x50, 0x0c, 0xf2, 0x70, 0x90, 0xa2,
	0x94, 0xc1, 0x04, 0xfb, 0xb9, 0xc8, 0x8a, 0x8c, 0xb6, 0x52, 0x39, 0xc9, 0x43, 0xfb, 0x0c, 0x76,
	0xdc, 0x8f, 0x2c, 0x7a, 0xbf, 0xe1, 0x6f, 0x19, 0xed, 0x41, 0x9b, 0x97, 0x69, 0x88, 0xc2, 0x20,
	0x16, 0x71, 0x34, 0xb6, 0x74, 0x94, 0x42, 0x33, 0x0e, 0x64, 0x6c, 0x6c, 0x59, 0xc4, 0xd9, 0x65,
	0x4a, 0xdb, 0x33, 0xe8, 0xaa, 0x43, 0x3f, 0x90, 0xf1, 0x53, 0x89, 0xa2, 0xa2, 0xc7, 0xb0, 0x2d,
	0x70, 0x3a, 0xae, 0x72, 0x54, 0xe7, 0xdd, 0x21, 0xed, 0xab, 0x8c, 0x3e, 0xc3, 0x69, 0x99, 0x08,
	0xac, 0xbf, 0xb0, 0xd5, 0x0a, 0xdd, 0x87, 0x96, 0x2c, 0x02, 0x51, 0xa8, 0x9f, 0x6a, 0x6c, 0x61,
	0xa8, 0x0e, 0x1a, 0xf2, 0x57, 0x43, 0x53, 0xb3, 0x5a, 0xd6, 0xd9, 0xbc, 0x4c, 0xa5, 0xd1, 0xb4,
	0x34, 0x47, 0x63, 0x4a, 0xdb, 0x1e, 0xec, 0xad, 0xb3, 0x19, 0xca, 0xbc, 0x46, 0xa6, 0x27, 0x00,
	0xe1, 0x8a, 0x44, 0x1a, 0xc4, 0xd2, 0x9c, 0xce, 0x50, 0x5f, 0x36, 0x58, 0x23, 0xb2, 0x8d, 0x1d,
	0xfb, 0x1c, 0x60, 0x54, 0xf1, 0xc8, 0xc7, 0x64, 0x12, 0x17, 0x35, 0x7c, 0xac, 0xd4, 0x0a, 0x7e,
	0xe1, 0xea, 0x02, 0x45, 0x92, 0xe2, 0xb2, 0xa7, 0xd2, 0x47, 0x17, 0xd0, 0xd9, 0x80, 0xa2, 0x14,
	0xba, 0xd7, 0xde, 0xd8, 0xbd, 0x7b, 0xbc, 0xba, 0xf5, 0x2f, 0x47, 0xbe, 0x37, 0xd2, 0x1b, 0xf4,
	0x00, 0x7a, 0xff, 0x67, 0xee, 0xcb, 0xc3, 0xf3, 0xbd, 0xeb, 0x31, 0x9d, 0xb8, 0xfa, 0xd7, 0xdc,
	0x24, 0xdf, 0x73, 0x93, 0xfc, 0xcc, 0x4d, 0xf2, 0xf9, 0x6b, 0x36, 0xc2, 0xb6, 0x7a, 0x94, 0xd3,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x06, 0x49, 0xa6, 0xc4, 0xbf, 0x01, 0x00, 0x00,
}
