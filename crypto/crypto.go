package crypto //import "go.iondynamics.net/iDhelper/crypto"

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"go.iondynamics.net/iDlogger"
	"io"
)

func Encrypt(keystr, text string) string {
	sum := sha256.Sum256([]byte(keystr))
	key := sum[:]
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		iDlogger.Panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		iDlogger.Panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext)
}

func Decrypt(keystr, cryptoText string) string {
	sum := sha256.Sum256([]byte(keystr))
	key := sum[:]
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		iDlogger.Panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		iDlogger.Panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}
