package tcpfw

import (
	"encoding/binary"
	"errors"
	"github.com/BLHT/HT_GOGO/gotcp"
	"io"
	"net"
)

const (
	HeadV3Len     = 48
	PacketHeadLen = 49
	EmptyPacktLen = 50
	PacketLimit   = 64 * 1024
)

// Error type
var (
	ErrShortLen = errors.New("use byte[] lne is not enough")
)

type HeadV3 struct {
	Flag         uint8     // 0xF0客户端请求，0xF1 服务器应答, 0xF2  服务器主动发包, 0xF3  客户端应答, 0xF4 服务器之间的包
	Version      uint8     // 版本号  VER_MMEDIA = 4
	CryptoKey    uint8     // 加密类型  E_NONE_KEY = 0, E_SESSION_KEY = 1, E_RAND_KEY = 2, E_SERV_KEY= 3
	TerminalType uint8     // 终端类型
	Cmd          uint16    // 命令字
	Seq          uint16    // 序列号
	FromUid      uint32    // uint32_t uiFrom
	ToUid        uint32    // 目的UID TO_SERVER = 0
	PacketLen    uint32    // 包总长度
	Ret          uint16    // 返回码
	SysType      uint16    // 包来源
	Echo         [8]uint8  // 回带字段
	Reserved     [16]uint8 // 保留字段
}

func NewHeadV3(buf []byte) (head *HeadV3, err error) {
	if len(buf) < HeadV3Len {
		return nil, ErrShortLen
	}

	head = new(HeadV3)
	head.Flag = buf[0]
	head.Version = buf[1]
	head.CryptoKey = buf[2]
	head.TerminalType = buf[3]
	head.Cmd = binary.BigEndian.Uint16(buf[4:6])
	head.Seq = binary.BigEndian.Uint16(buf[6:8])
	head.FromUid = binary.BigEndian.Uint32(buf[8:12])
	head.ToUid = binary.BigEndian.Uint32(buf[12:16])
	head.PacketLen = binary.BigEndian.Uint32(buf[16:20])
	head.Ret = binary.BigEndian.Uint16(buf[20:22])
	head.SysType = binary.BigEndian.Uint16(buf[22:24])
	head.Echo[0] = buf[24]
	head.Echo[1] = buf[25]
	head.Echo[2] = buf[26]
	head.Echo[3] = buf[27]
	head.Echo[4] = buf[28]
	head.Echo[5] = buf[29]
	head.Echo[6] = buf[30]
	head.Echo[7] = buf[31]
	head.Reserved[0] = buf[32]
	head.Reserved[1] = buf[33]
	head.Reserved[2] = buf[34]
	head.Reserved[3] = buf[35]
	head.Reserved[4] = buf[36]
	head.Reserved[5] = buf[37]
	head.Reserved[6] = buf[38]
	head.Reserved[7] = buf[39]
	head.Reserved[8] = buf[40]
	head.Reserved[9] = buf[41]
	head.Reserved[10] = buf[42]
	head.Reserved[11] = buf[43]
	head.Reserved[12] = buf[44]
	head.Reserved[13] = buf[45]
	head.Reserved[14] = buf[46]
	head.Reserved[15] = buf[47]

	return head, nil
}

func SerialHeadV3ToSlice(head *HeadV3, buf []byte) (err error) {
	if len(buf) < HeadV3Len {
		return ErrShortLen
	}
	buf[0] = head.Flag
	buf[1] = head.Version
	buf[2] = head.CryptoKey
	buf[3] = head.TerminalType
	binary.BigEndian.PutUint16(buf[4:6], head.Cmd)
	binary.BigEndian.PutUint16(buf[6:8], head.Seq)
	binary.BigEndian.PutUint32(buf[8:12], head.FromUid)
	binary.BigEndian.PutUint32(buf[12:16], head.ToUid)
	binary.BigEndian.PutUint32(buf[16:20], head.PacketLen)
	binary.BigEndian.PutUint16(buf[20:22], head.Ret)
	binary.BigEndian.PutUint16(buf[22:24], head.SysType)
	echoSlice := head.Echo[:]
	copy(buf[24:32], echoSlice)
	reservedSlice := head.Reserved[:]
	copy(buf[32:48], reservedSlice)

	return nil
}

// HeadV3Packet 格式如下
// 0x0a + HTHeadV3 + payload + 0x0b

type HeadV3Packet struct {
	buff []byte
}

func (this *HeadV3Packet) Serialize() []byte {
	return this.buff
}

// index:0 SOH filed
func (this *HeadV3Packet) GetSoh() uint8 {
	return this.buff[0]
}

// index:1 Falg field
func (this *HeadV3Packet) GetFlag() uint8 {
	return this.buff[1]
}

// index:2 Version field
func (this *HeadV3Packet) GetVersion() uint8 {
	return this.buff[2]
}

// index:3 Crypto field
func (this *HeadV3Packet) GetKey() uint8 {
	return this.buff[3]
}

// index:4 Termianl type field
func (this *HeadV3Packet) GetTerminalType() uint8 {
	return this.buff[4]
}

// index:5 Command field
func (this *HeadV3Packet) GetCommand() uint16 {
	return binary.BigEndian.Uint16(this.buff[5:7])
}

// index:6 Sequence field
func (this *HeadV3Packet) GetSeq() uint16 {
	return binary.BigEndian.Uint16(this.buff[7:9])
}

// index:7 From Uid field
func (this *HeadV3Packet) GetFromUid() uint32 {
	return binary.BigEndian.Uint32(this.buff[9:13])
}

// index:8 To Uid field
func (this *HeadV3Packet) GetToUid() uint32 {
	return binary.BigEndian.Uint32(this.buff[13:17])
}

// index:9 Packet Length field length(SOH+HeadV3+PayLoad+EOT)
func (this *HeadV3Packet) GetLength() uint32 {
	return binary.BigEndian.Uint32(this.buff[17:21])
}

// index:10 Return Result field
func (this *HeadV3Packet) GetRet() uint16 {
	return binary.BigEndian.Uint16(this.buff[21:23])
}

// index:11 Packet come from field
func (this *HeadV3Packet) GetSystem() uint16 {
	return binary.BigEndian.Uint16(this.buff[23:25])
}

// index:12 Echo field
func (this *HeadV3Packet) GetEcho() []byte {
	return this.buff[25:33]
}

// index:13 Reserved filed
func (this *HeadV3Packet) GetReserve() []byte {
	return this.buff[33:49]
}

// index:14 Body filed
func (this *HeadV3Packet) GetBody() []byte {
	return this.buff[49 : len(this.buff)-1]
}

// index:15 EOT filed
func (this *HeadV3Packet) GetEot() uint8 {
	return this.buff[len(this.buff)-1]
}

func NewHeadV3Packet(buff []byte) *HeadV3Packet {
	p := &HeadV3Packet{}
	p.buff = buff
	return p
}

type HeadV3Protocol struct {
}

func (this *HeadV3Protocol) ReadPacket(conn *net.TCPConn) (gotcp.Packet, error) {

	var (
		headBytes []byte = make([]byte, PacketHeadLen)
		length    uint32
	)

	// read length
	if _, err := io.ReadFull(conn, headBytes); err != nil {
		return nil, err
	}

	head, err := NewHeadV3(headBytes[1:])
	if err != nil {
		return nil, err
	}

	if length = head.PacketLen; length > PacketLimit {
		return nil, errors.New("the size of packet is larger than the limit")
	}

	buff := make([]byte, length)
	copy(buff[0:PacketHeadLen], headBytes)

	// read body ( buff = lengthBytes + body )
	if _, err := io.ReadFull(conn, buff[PacketHeadLen:]); err != nil {
		return nil, err
	}

	return NewHeadV3Packet(buff), nil
}
