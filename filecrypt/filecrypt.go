package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha256"
	"io"
	"os"
)

func Encrypt(source string, password []byte) {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}

	srcFile, err := os.Open(source)
	if err != nil {
		panic(err.Error())
	}
	defer srcFile.Close()
	plaintext, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err.Error())
	}
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		panic(err.Error())
	}
	passwordStr := string(password)
	derivedKey, err := pbkdf2.Key(sha256.New, passwordStr, salt, 10000, 32)
	if err != nil {
		panic(err.Error())
	}
	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	result := append(salt, nonce...)
	result = append(result, ciphertext...)

	destFile, err := os.Create(source)
	if err != nil {
		panic(err.Error())
	}
	defer destFile.Close()

	_, err = destFile.Write(result)
	if err != nil {
		panic(err.Error())
	}
}

func Decrypt(source string, password []byte) {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}

	srcFile, err := os.Open(source)
	if err != nil {
		panic(err.Error())
	}
	defer srcFile.Close()

	ciphertext, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err.Error())
	}

	salt := ciphertext[:16]
	nonceSize := 12
	nonce := ciphertext[16 : 16+nonceSize]
	actualCiphertext := ciphertext[16+nonceSize:]

	passwordStr := string(password)

	derivedKey, err := pbkdf2.Key(sha256.New, passwordStr, salt, 10000, 32)
	if err != nil {
		panic(err.Error())
	}

	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, actualCiphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	destFile, err := os.Create(source)
	if err != nil {
		panic(err.Error())
	}
	defer destFile.Close()

	_, err = destFile.Write(plaintext)
	if err != nil {
		panic(err.Error())
	}
}
