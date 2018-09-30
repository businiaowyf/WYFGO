package main

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"log"

	"github.com/businiaowyf/wyfgo/test/aes/ecb"
)

func DecryptOpenId(openId string, key []byte) (uint8, uint32, uint64, error) {
	id, err := hex.DecodeString(openId)
	if err != nil {
		return 0, 0, 0, err
	}
	if len(id) != 16 {
		return 0, 0, 0, errors.New("openid's length illegal")
	}
	plaintext, err := ecb.EcbDecryptNoPadding(id, key)
	if err != nil {
		log.Fatal("decrypt error=%v", err)
		return 0, 0, 0, err
	}
	log.Printf("plaintext=%v", plaintext)

	platform := uint8(plaintext[0])
	appId := binary.BigEndian.Uint32(plaintext[4:8])
	uid := binary.BigEndian.Uint64(plaintext[8:16])
	return platform, appId, uid, nil
}

func EncryptOpenId(platform uint8, appId uint32, uid uint64, key []byte) (string, error) {
	buf := make([]byte, 16)
	var p uint32 = uint32(platform) << 24
	binary.BigEndian.PutUint32(buf[:4], p)
	binary.BigEndian.PutUint32(buf[4:8], appId)
	binary.BigEndian.PutUint64(buf[8:16], uid)
	ciphertext, err := ecb.EcbEncryptNoPadding(buf, key)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(ciphertext), nil
}

func main() {
	key, err := hex.DecodeString("b9a5ffbb7cdd8727898f9cb607331c94")
	if err != nil {
		log.Printf("Decode key error=%v", err)
		return
	}
	openId := "15086c53ea86f3d00b013dfd265cb9a0"
	log.Printf("before openid=%v", openId)

	platform, appId, uid, err := DecryptOpenId(openId, key)
	if err != nil {
		log.Println("DecryptOpenId error=", err)
		return
	}
	log.Printf("platform=%v, appid=%v, uid=%v", platform, appId, uid)
	id, err := EncryptOpenId(platform, appId, uid, key)
	if err != nil {
		log.Println("EncryptOpenId error=", err)
		return
	}
	log.Printf("after openid=%v", id)
}
