package common

import (
	"errors"
	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
)

var (
	ErrPacketShortLen = errors.New("use byte[] length is not enough")
)

const (
	COnlineStatLen  = 70 //70 = 67+2个字节长度 + 1个字节的'\0'
	CMCOnlinePrefix = "state#"
)

const (
	ST_OFFLINE   = 0
	ST_ONLINE    = 1
	ST_BACKGROUD = 2
	ST_LOGOUT    = 3
	ST_INIT      = 100 // 初始化
)

type UserState struct {
	Uid        uint32 //用户id
	ClientType uint8  //客户端类型 TERMIAL_TYPE
	OnlineStat uint8  //在线状态
	SvrIp      uint32 //所在服务器IP
	UpdateTs   uint64 //更新时间戳
	Session    []byte //Session
	Version    uint32 //版本
	Port       uint32 //端口
	Wid        uint64 //wid
	UserType   uint8  //用户类型
}

func (this *UserState) unMarshall(buf []byte) (err error) {
	if len(buf) < COnlineStatLen {
		err = ErrPacketShortLen
		return
	}
	this.Uid = UnMarshalUint32(&buf)
	this.ClientType = UnMarshalUint8(&buf)
	this.OnlineStat = UnMarshalUint8(&buf)
	this.SvrIp = UnMarshalUint32(&buf)
	this.UpdateTs = UnMarshalUint64(&buf)
	this.Session = UnMarshalSlice(&buf)
	this.Version = UnMarshalUint32(&buf)
	this.Port = UnMarshalUint32(&buf)
	this.Wid = UnMarshalUint64(&buf)
	this.UserType = UnMarshalUint8(&buf)
	return nil
}

func (this *UserState) marshall(buf []byte) (err error) {
	if len(buf) < COnlineStatLen {
		err = ErrPacketShortLen
		return
	}
	MarshalUint32(this.Uid, &buf)
	MarshalUint8(this.ClientType, &buf)
	MarshalUint8(this.OnlineStat, &buf)
	MarshalUint32(this.SvrIp, &buf)
	MarshalUint64(this.UpdateTs, &buf)
	MarshalSlice(this.Session, &buf)
	MarshalUint32(this.Version, &buf)
	MarshalUint32(this.Port, &buf)
	MarshalUint64(this.Wid, &buf)
	MarshalUint8(this.UserType, &buf)
	return nil

}

// Conn exposes a set of callbacks for the various events that occur on a connection
type MemcacheApi struct {
	client *memcache.Client
}

func (c *MemcacheApi) Init(server ...string) {
	c.client = memcache.New(server...)
}

func (c *MemcacheApi) GetUserOnlineStat(uid uint32) (stat *UserState, err error) {
	key := CMCOnlinePrefix + strconv.Itoa(int(uid))
	it, err := c.client.Get(key)
	if err != nil {
		return nil, err
	}
	stat = new(UserState)
	err = stat.unMarshall(it.Value)
	return stat, err
}
