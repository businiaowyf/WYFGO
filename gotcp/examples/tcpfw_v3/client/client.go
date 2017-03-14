package main

import (
	"github.com/BLHT/HT_GOGO/gotcp/examples/tcpfw"
	"log"
	"net"
	"os"
	"time"
)

// func ConstructPacket(payLoad []byte) (buff []byte) {
// 	buff = make([]byte, packetHeadLen+len(payLoad)+1)
// 	buff[0] = 0x0a
// 	buff[1] = 0xF0
// 	buff[2] = 0x04
// 	buff[3] = 0x00
// 	buff[4] = 0x00
// 	binary.BigEndian.PutUint16(buff[5:7], 0x1003)
// 	binary.BigEndian.PutUint16(buff[7:9], 0x0001)
// 	binary.BigEndian.PutUint32(buff[9:13], 1946612)
// 	binary.BigEndian.PutUint32(buff[13:17], 2325928)
// 	var packetLenth uint32 = packetHeadLen + uint32(len(payLoad)) + 1
// 	binary.BigEndian.PutUint32(buff[17:21], packetLenth)
// 	binary.BigEndian.PutUint16(buff[21:23], 0)
// 	binary.BigEndian.PutUint16(buff[23:25], 0)
// 	copy(buff[49:len(buff)-1], payLoad)
// 	buff[len(buff)-1] = 0x0b
// 	return buff
// }

func main() {
	// 定义一个文件
	fileName := "test.log"
	logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	// 创建一个日志对象
	debugLog := log.New(logFile, "[Info]", log.LstdFlags)
	// 配置log的Flag参数
	debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8989")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	headV3Protocol := &tcpfw.HeadV3Protocol{}

	// ping <--> pong
	for i := 0; i < 3; i++ {
		head := new(tcpfw.HeadV3)
		head.Flag = 0xF0
		head.Version = 0x04
		head.CryptoKey = 0x00
		head.TerminalType = 0x00
		head.Cmd = 0x1003
		head.Seq = 0x001
		head.FromUid = 1946612
		head.ToUid = 2325928
		head.PacketLen = tcpfw.EmptyPacktLen
		head.Ret = 0
		head.SysType = 0

		payLoad := []byte("hello")
		head.PacketLen = uint32(tcpfw.PacketHeadLen) + uint32(len(payLoad)) + 1
		buf := make([]byte, head.PacketLen)
		buf[0] = 0x0a
		err := tcpfw.SerialHeadV3ToSlice(head, buf[1:])
		if err != nil {
			debugLog.Println("SerialHeadV3ToSlice failed")
			continue
		}
		copy(buf[tcpfw.PacketHeadLen:], payLoad)
		buf[head.PacketLen-1] = 0x0b

		debugLog.Println("payLoad=", tcpfw.NewHeadV3Packet(buf).Serialize())
		// write
		conn.Write(tcpfw.NewHeadV3Packet(buf).Serialize())

		// read
		p, err := headV3Protocol.ReadPacket(conn)
		if err == nil {
			headV3Packet := p.(*tcpfw.HeadV3Packet)
			debugLog.Printf("Server reply:[%v] [%v]\n", headV3Packet.GetLength(), string(headV3Packet.GetBody()))

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
