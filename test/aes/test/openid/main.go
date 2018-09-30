package main

import (
	"encoding/hex"
	"log"

	"github.com/businiaowyf/wyfgo/test/aes/ecb"
)

func main() {
	key, err := hex.DecodeString("a9a5ffbb7cdd8727898f9cb607331c93")
	if err != nil {
		log.Printf("Decode key error=%v", err)
		return
	}
	openid := "15086c53ea86f3d00b013dfd265cb9a0"
	log.Printf("openid len=%v", len(openid))

	plaintext, err := ecb.EcbDecrypt([]byte(openid), key)
	if err != nil {
		log.Fatal("decrypt error=%v", err)
		return
	}
	log.Printf("plaintext=%v", plaintext)
}
