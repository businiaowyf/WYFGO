package main

import (
	"github.com/businiaowyf/wyfgo/test/aes/cbc"
	"github.com/businiaowyf/wyfgo/test/aes/ecb"
	"github.com/businiaowyf/wyfgo/test/aes/gcm"
)

func main() {
	ecb.TestECB()
	gcm.TestGCM()
	cbc.TestCBC()
}
