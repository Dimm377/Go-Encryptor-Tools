package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"

	"golang.org/x/crypto/argon2"
)

const (
	// Argon2id parameters (RFC 9106 recommended for interactive sessions)
	time    = 1
	memory  = 64 * 1024 // 64 MB
	threads = 4
	keyLen  = 32
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

	// argon2.IDKey signature: []byte password, []byte salt, uint32 time, uint32 memory, uint8 threads, uint32 keyLen
	derivedKey := argon2.IDKey(password, salt, time, memory, threads, keyLen)

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

	// Safe write: write to temp file first
	tmpFile := source + ".tmp"
	destFile, err := os.Create(tmpFile)
	if err != nil {
		panic(err.Error())
	}
	
	if _, err := destFile.Write(result); err != nil {
		destFile.Close()
		os.Remove(tmpFile) // Clean up
		panic(err.Error())
	}
	destFile.Close()

	// Atomic rename
	if err := os.Rename(tmpFile, source); err != nil {
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

	fileContent, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err.Error())
	}

	if len(fileContent) < 16+12 {
		panic("invalid file format: too short")
	}

	salt := fileContent[:16]
	nonceSize := 12
	nonce := fileContent[16 : 16+nonceSize]
	actualCiphertext := fileContent[16+nonceSize:]

	// argon2.IDKey signature: []byte password, []byte salt, uint32 time, uint32 memory, uint8 threads, uint32 keyLen
	derivedKey := argon2.IDKey(password, salt, time, memory, threads, keyLen)

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
		panic("decryption failed: " + err.Error())
	}

	// Safe write: write to temp file first
	tmpFile := source + ".tmp"
	destFile, err := os.Create(tmpFile)
	if err != nil {
		panic(err.Error())
	}

	if _, err := destFile.Write(plaintext); err != nil {
		destFile.Close()
		os.Remove(tmpFile)
		panic(err.Error())
	}
	destFile.Close()

	// Atomic rename
	if err := os.Rename(tmpFile, source); err != nil {
		panic(err.Error())
	}
}
