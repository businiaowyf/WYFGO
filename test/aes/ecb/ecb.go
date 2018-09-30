package ecb

import (
	"bytes"
	"crypto/aes"
	"fmt"
)

func TestECB() {
	key := "0123456789abcdef0123456789abcdef"
	plaintext := "12345678901234567"
	ciphertext, err := EcbEncrypt([]byte(plaintext), []byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v -> %v, len=%v\n", plaintext, ciphertext, len(ciphertext))
	result, err := EcbDecrypt([]byte(ciphertext), []byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v -> %v\n", ciphertext, string(result))
}

func EcbEncrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	plaintext = PKCS5Padding(plaintext, block.BlockSize())
	ciphertext := plaintext
	for len(plaintext) > 0 {
		block.Encrypt(plaintext, plaintext[:block.BlockSize()])
		plaintext = plaintext[block.BlockSize():]
	}
	return ciphertext, nil
}

func EcbDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	plaintext := ciphertext
	for len(ciphertext) > 0 {
		block.Decrypt(ciphertext, ciphertext[:block.BlockSize()])
		ciphertext = ciphertext[block.BlockSize():]
	}

	plaintext = PKCS5UnPadding(plaintext)

	return plaintext, nil
}

func PKCS5Padding(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, padtext...)
}

func PKCS5UnPadding(text []byte) []byte {
	length := len(text)
	unpadding := int(text[length-1])
	return text[:(length - unpadding)]
}
