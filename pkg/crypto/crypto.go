package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
	"os"

	"github.com/mergermarket/go-pkcs7"
)

func Decrypt(encrypted string) (string, error) {
	key := []byte(os.Getenv("CYPHER_KEY"))

	cipherText, _ := hex.DecodeString(encrypted)

	block, err := aes.NewCipher(key)

	if err != nil {
		log.Fatalln(err)
	}

	if len(cipherText) < aes.BlockSize {
		log.Fatalln("cipherText too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	if len(cipherText)%aes.BlockSize != 0 {
		log.Fatalln("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, _ = pkcs7.Unpad(cipherText, aes.BlockSize)
	return string(cipherText), nil
}
