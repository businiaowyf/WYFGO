package tcpfw

import (
	"encoding/binary"
	"errors"
	"github.com/BLHT/HT_GOGO/gotcp"
	"io"
	"net"
)

// HTHeadV3 {
// uint8_t ucFlag			// 0xF0客户端请求，0xF1 服务器应答, 0xF2  服务器主动发包, 0xF3  客户端应答, 0xF4 服务器之间的包
// uint8_t ucVersion		// 版本号  VER_MMEDIA = 4
// uint8_t ucKey			//加密类型  E_NONE_KEY = 0, E_SESSION_KEY = 1, E_RAND_KEY = 2, E_SERV_KEY= 3
// uint8_t ucReserved		//保留字节, 这个是为了兼容XT_head
// uint16_t usCmd			//命令字
// uint16_t usSeq			//序列号
// uint32_t uiFrom		//源UID  FROM_SERVER = 0
// uint32_t uiTo			//目的UID TO_SERVER = 0
// uint32_t uiLen			//包总长度
// uint16_t usRet			//返回码
// uint16_t usSysType		//包来源
// char acEcho[8]			//回带字段
// char acReserved[16]	//保留字段
//}
//
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
	const (
		packetHeadLen = 49
		packetLimit   = 64 * 1024
	)

	var (
		headBytes []byte = make([]byte, packetHeadLen)
		length    uint32
	)

	// read length

	if _, err := io.ReadFull(conn, headBytes); err != nil {
		return nil, err
	}

	if length = binary.BigEndian.Uint32(headBytes[17:21]); length > packetLimit {
		return nil, errors.New("the size of packet is larger than the limit")
	}

	buff := make([]byte, length)
	copy(buff[0:packetHeadLen], headBytes)

	// read body ( buff = lengthBytes + body )
	if _, err := io.ReadFull(conn, buff[packetHeadLen:]); err != nil {
		return nil, err
	}

	return NewHeadV3Packet(buff), nil
}
