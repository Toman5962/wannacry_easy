package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

// Encrypt
func Encrypt(path string, secretKey []byte)  {
	data, err := ioutil.ReadFile(path)
	if err != nil{
		panic(err)
	}

	block,err := aes.NewCipher(secretKey)
	if err !=nil {
		panic(err)
	}

	cipherText := make([]byte,aes.BlockSize + len(data))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader,iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], data)
	if err := ioutil.WriteFile(path, cipherText, 0644); err != nil {
		panic(err)
	}
}

// Decrypt
func Decrypt(path string, secretKey []byte)  {
	fmt.Println("仿照encrypt")
}