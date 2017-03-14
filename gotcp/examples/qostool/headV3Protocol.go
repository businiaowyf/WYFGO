package echo

import (
	"encoding/binary"
	"encoding/gob"
	"errors"
	"io"
	"net"
	"unsafe"
	"github.com/BLHT/HT_GOGO/gotcp"
)

type HeadV2 struct {
	version  uint8
	cmd      uint32
	seq      uint32
	ret      uint16
	reserved uint8
	len      uint32
	uid      uint32
	sysType  uint16
	echo     [8]uint8
}

// The HeadV3 type has unexported fields, which the package cannot access.
// We therefore write a BinaryMarshal/BinaryUnmarshal method pair to allow us
// to send and receive the type with the gob package. These interfaces are
// defined in the "encoding" package.
// We could equivalently use the locally defined GobEncode/GobDecoder
// interfaces.

func (head HeadV2) MarshalBinary() ([]byte, error) {
	// A simple encoding: plain text.
	var b bytes.Buffer
	fmt.Fprintln(&b,
		head.version,
		head.cmd,
		head.seq,
		head.ret,
		head.reserved,
		head.len,
		head.uid,
		head.sysType,
		head.echo[0],
		head.echo[1],
		head.echo[2],
		head.echo[3],
		head.echo[4],
		head.echo[5],
		head.echo[6],
		head.echo[7])
	return b.Bytes(), nil
}

// UnmarshalBinary modifies the receiver so it must take a pointer receiver.
func (head *HeadV2) UnmarshalBinary(data []byte) error {
	// A simple encoding: plain text.
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b,
		&head.version,
		&head.cmd,
		&head.seq,
		&head.ret,
		&head.reserved,
		&head.len,
		&head.uid,
		&head.sysType,
		&head.echo[0],
		&head.echo[1],
		&head.echo[2],
		&head.echo[3],
		&head.echo[4],
		&head.echo[5],
		&head.echo[6],
		&head.echo[7])

	return err
}

func Htonl(i uint32) uint32 {
	return uint32(Htons(uint16(i)))<<16 | uint32(Htons(uint16(i>>16)))
}

func Ntohl(i uint32) uint32 {
	return Htonl(i)
}

func Htons(i uint16) uint16 {
	return uint16(i>>8) | uint16(i<<8)
}

func Ntohs(i uint16) uint16 {
	return Htons(i)
}

func (head *HeadV2) Hton() {
	head.cmd = Htonl(head.cmd)
	head.seq = Htonl(head.seq)
	head.ret = Htons(head.ret)
	head.len = Htonl(head.len)
	head.uid = Htonl(head.uid)
	head.sysType = Htons(head.sysType)
}

func (head *HeadV2) Ntoh() {
	head.cmd = Ntohl(head.cmd)
	head.seq = Ntohl(head.seq)
	head.ret = Ntohs(head.ret)
	head.len = Ntohl(head.len)
	head.uid = Ntohl(head.uid)
	head.sysType = Htons(head.sysType)
}

func (head *HeadV2) Len() uint32 {
	return head.len()
}

type HeadV2Packet struct {
	buff []byte
}

func (this *HeadV2Packet) Serialize() []byte {
	return this.buff
}

func (this *HeadV2Packet) GetLength() uint32 {
	var head HeadV2
	bufHead := bytes.NewBuffer(this.buff[])
	dec := gob.NewDecoder(bufHead)
	err = dec.Decode(&head)
	if err != nil {
		log.Fatal("decode:", err)
	}
	head.Ntoh()
	return head.Len() // whole packet lenght include packet head
}

func (this *HeadV2Packet) GetBody() []byte {
	return this.buff[:]
}

func NewHeadV3Packet(buff []byte, hasLengthField bool) *HeadV3Packet {
	p := &HeadV3Packet{}

	if hasLengthField {
		p.buff = buff

	} else {
		p.buff = make([]byte, 4+len(buff))
		binary.BigEndian.PutUint32(p.buff[0:4], uint32(len(buff)))
		copy(p.buff[4:], buff)
	}

	return p
}

type EchoProtocol struct {
}

func (this *EchoProtocol) ReadPacket(conn *net.TCPConn) (gotcp.Packet, error) {
	var (
		lengthBytes []byte = make([]byte, 4)
		length      uint32
	)

	// read length
	if _, err := io.ReadFull(conn, lengthBytes); err != nil {
		return nil, err
	}
	if length = binary.BigEndian.Uint32(lengthBytes); length > 1024 {
		return nil, errors.New("the size of packet is larger than the limit")
	}

	buff := make([]byte, 4+length)
	copy(buff[0:4], lengthBytes)

	// read body ( buff = lengthBytes + body )
	if _, err := io.ReadFull(conn, buff[4:]); err != nil {
		return nil, err
	}

	return NewHeadV3Packet(buff, true), nil
}
