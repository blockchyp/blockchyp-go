package blockchyp

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"math/rand"
	"time"
)

/*
Encrypt performs AES 256/CBC/PKCS7 encryption on the given plain text with the given key.
*/
func Encrypt(key []byte, plainText string) string {

	plainBytes := []byte(plainText)

	key = key[:16]

	plainBytes = pkcs7Pad(plainBytes, aes.BlockSize)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	randometer := rand.New(rand.NewSource(time.Now().UnixNano()))

	cipherText := make([]byte, aes.BlockSize+len(plainBytes))
	iv := cipherText[:aes.BlockSize]
	randometer.Read(iv)

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainBytes)

	return hex.EncodeToString(cipherText)

}

/*
Decrypt performs AES 256/CBC/PKCS7 decryption on the given cipherText with the given key.
*/
func Decrypt(key []byte, cipherText string) string {

	key = key[:16]

	cipherBytes, _ := hex.DecodeString(cipherText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	iv := cipherBytes[:aes.BlockSize]
	cipherBytes = cipherBytes[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(cipherBytes, cipherBytes)

	return string(pkcs7Unpad(cipherBytes, aes.BlockSize))
}

func pkcs7Pad(b []byte, blocksize int) []byte {
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
	return pb
}

func pkcs7Unpad(b []byte, blocksize int) []byte {
	c := b[len(b)-1]
	n := int(c)
	if n == 0 || n > len(b) {
		return nil
	}
	for i := 0; i < n; i++ {
		if b[len(b)-n+i] != c {
			return nil
		}
	}
	return b[:len(b)-n]
}
