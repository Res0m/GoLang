package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не передан параметр KEY в переменные окружения")
	}
	return &Encrypter{
		Key: key,
	}
}

func (enc *Encrypter) Encrypt(plainStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key)) // создаем блок для шифрования
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block) // передаем для создания шифра
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	return aesGCM.Seal(nonce, nonce, plainStr, nil)
}
func (enc *Encrypter) Decrypt(encryptedStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key)) // создаем блок для шифрования
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block) // передаем для создания шифра
	if err != nil {
		panic(err.Error())
	}
	noneSize := aesGCM.NonceSize()
	nonce, cipherText := encryptedStr[:noneSize], encryptedStr[noneSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return plainText
}
