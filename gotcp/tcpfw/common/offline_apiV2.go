package common

import (
	"time"

	"github.com/businiaowyf/wyfgo/gotcp"
	"github.com/businiaowyf/wyfgo/gotcp/tcpfw/include/ht_offline"
	"github.com/golang/protobuf/proto"
)

// Conn exposes a set of callbacks for the various events that occur on a connection
type OfflineApiV2 struct {
	pool *Pool
}

func newPool(ip, port string, readTimeout, writeTimeout time.Duration, maxConn int, proto gotcp.Protocol) *Pool {
	return &Pool{
		MaxIdle:     3,
		MaxActive:   maxConn,
		IdleTimeout: 240 * time.Second,
		Dial: func() (Conn, error) {
			c, err := Dial(ip, port, proto, DialReadTimeout(readTimeout), DialWriteTimeout(writeTimeout))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

func NewOfflineApiV2(ip, port string, readTimeout, writeTimeout time.Duration, protocol gotcp.Protocol, maxConn int) *OfflineApiV2 {
	pool := newPool(ip, port, readTimeout, writeTimeout, maxConn, protocol)
	return &OfflineApiV2{
		pool: pool,
	}
}

func (c *OfflineApiV2) SendPacketWithHeadV3(headV3 *HeadV3, payLoad []uint8, push []byte) (ret uint16, err error) {
	// 存离线消息的包（HTHeadV3的包）
	newHead := new(HeadV3)
	*newHead = *headV3
	newHead.Len = uint32(PacketV3HeadLen) + uint32(len(payLoad)) + 1
	buf := make([]byte, newHead.Len)
	buf[0] = HTV3MagicBegin
	err = SerialHeadV3ToSlice(newHead, buf[1:])
	if err != nil {
		ret = uint16(CRetMarshallFailed)
		return ret, err
	}
	copy(buf[PacketV3HeadLen:], payLoad)
	buf[newHead.Len-1] = HTV3MagicEnd

	// 构造HTHeadV2 报文
	headV2 := &HeadV2{Cmd: uint32(ht_offline.CMD_TYPE_CMD_SAVE_OFFLINE_MSG),
		Uid:     headV3.To,
		SysType: uint16(ht_offline.SYS_TYPE_SYS_VOIP_SERVER),
	}

	reqBody := new(ht_offline.ReqBody)
	reqBody.SaveOfflineMsgReqbody = new(ht_offline.SaveOfflineMsgReqBody)
	subReqBody := reqBody.GetSaveOfflineMsgReqbody()
	subReqBody.FromUid = new(uint32)
	*(subReqBody.FromUid) = headV3.From

	subReqBody.ToUid = new(uint32)
	*(subReqBody.ToUid) = headV3.To

	subReqBody.Cmd = new(uint32)
	*(subReqBody.Cmd) = uint32(headV3.Cmd)

	subReqBody.Format = new(uint32)
	*(subReqBody.Format) = 10 // HTHeadV3的消息格式用10

	subReqBody.Content = make([]byte, newHead.Len)
	copy(subReqBody.Content, buf)

	subReqBody.MsgTime = new(uint32)
	*(subReqBody.MsgTime) = uint32(time.Now().Unix()) // 添加时间戳

	subReqBody.PushType = new(uint32)
	if len(push) > 0 {
		*(subReqBody.PushType) = uint32(ht_offline.SaveOfflineMsgReqBody_EM_PUSH_CONFIRM)
		subReqBody.PushPkg = make([]byte, len(push))
		copy(subReqBody.PushPkg, push)
	} else {
		*(subReqBody.PushType) = uint32(ht_offline.SaveOfflineMsgReqBody_EM_PUSH_DISCARD)
	}

	s, err := proto.Marshal(reqBody)
	if err != nil {
		ret = uint16(CRetMarshallFailed)
		return ret, err
	}

	headV2.Len = uint32(PacketV2HeadLen) + uint32(len(s)) + 1
	headV2Buf := make([]byte, headV2.Len)
	headV2Buf[0] = HTV2MagicBegin
	err = SerialHeadV2ToSlice(headV2, headV2Buf[1:])
	if err != nil {
		ret = uint16(CRetMarshallFailed)
		return ret, err
	}
	copy(headV2Buf[PacketV2HeadLen:], s)
	headV2Buf[headV2.Len-1] = HTV2MagicEnd

	packet := NewHeadV2Packet(headV2Buf)

	conn := c.pool.Get()
	defer conn.Close()

	rsp, err := conn.Do(packet)
	rspPacket := rsp.(*HeadV2Packet)
	rspHead, err := rspPacket.GetHead()
	if err != nil {
		ret := uint16(CRetUnMarshallFailed)
		return ret, err
	}
	ret = rspHead.Ret
	return ret, nil
}
