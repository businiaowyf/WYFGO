package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/BLHT/HT_GOGO/gotcp/examples/tcpfw"
)

const (
	packetHeadLen = 49
)

func ConstructPacket(payLoad []byte) (buff []byte) {
	buff = make([]byte, packetHeadLen+len(payLoad)+1)
	buff[0] = 0x0a
	buff[1] = 0xF0
	buff[2] = 0x04
	buff[3] = 0x00
	buff[4] = 0x00
	binary.BigEndian.PutUint16(buff[5:7], 0x1003)
	binary.BigEndian.PutUint16(buff[7:9], 0x0001)
	binary.BigEndian.PutUint32(buff[9:13], 1946612)
	binary.BigEndian.PutUint32(buff[13:17], 2325928)
	var packetLenth uint32 = packetHeadLen + uint32(len(payLoad)) + 1
	binary.BigEndian.PutUint32(buff[17:21], packetLenth)
	binary.BigEndian.PutUint16(buff[21:23], 0)
	binary.BigEndian.PutUint16(buff[23:25], 0)
	copy(buff[49:len(buff)-1], payLoad)
	buff[len(buff)-1] = 0x0b
	return buff
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8989")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	headV3Protocol := &tcpfw.HeadV3Protocol{}

	// ping <--> pong
	for i := 0; i < 3; i++ {
		buff := ConstructPacket([]byte("hello"))
		fmt.Println("packet=", buff)
		// write
		conn.Write(tcpfw.NewHeadV3Packet(buff).Serialize())

		// read
		p, err := headV3Protocol.ReadPacket(conn)
		if err == nil {
			headV3Packet := p.(*tcpfw.HeadV3Packet)
			fmt.Printf("Server reply:[%v] [%v]\n", headV3Packet.GetLength(), string(headV3Packet.GetBody()))
		}

		time.Sleep(2 * time.Second)
	}

	conn.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
