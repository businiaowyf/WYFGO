package gcm

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
)

func TestGCM() {
	key := "0123456789abcdef0123456789abcdef"
	plaintext := "12345678901234567"
	ciphertext, err := GcmEncrypt([]byte(plaintext), []byte(key))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s => %x %v\n", plaintext, ciphertext, len(ciphertext))

	result, err := GcmDecrypt([]byte(ciphertext), []byte(key))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x => %s %v\n", ciphertext, result, len(result))
}

func GcmEncrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	out := gcm.Seal(nonce, nonce, plaintext, nil)
	return out, nil
}

func GcmDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
