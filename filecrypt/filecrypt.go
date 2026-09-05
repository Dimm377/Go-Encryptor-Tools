package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
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

func Encrypt(source string, password []byte) error {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return err
	}

	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	plaintext, err := io.ReadAll(srcFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return fmt.Errorf("failed to generate salt: %w", err)
	}

	// argon2.IDKey signature: []byte password, []byte salt, uint32 time, uint32 memory, uint8 threads, uint32 keyLen
	derivedKey := argon2.IDKey(password, salt, time, memory, threads, keyLen)

	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return fmt.Errorf("failed to create cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("failed to create GCM: %w", err)
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return fmt.Errorf("failed to generate nonce: %w", err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	result := append(salt, nonce...)
	result = append(result, ciphertext...)

	// Safe write: write to temp file first
	tmpFile := source + ".tmp"
	destFile, err := os.Create(tmpFile)
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	
	if _, err := destFile.Write(result); err != nil {
		destFile.Close()
		os.Remove(tmpFile) // Clean up
		return fmt.Errorf("failed to write encrypted data: %w", err)
	}
	destFile.Close()

	// Atomic rename
	if err := os.Rename(tmpFile, source); err != nil {
		return fmt.Errorf("failed to rename temp file: %w", err)
	}
	
	return nil
}

func Decrypt(source string, password []byte) error {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return err
	}

	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	fileContent, err := io.ReadAll(srcFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	if len(fileContent) < 16+12 {
		return errors.New("invalid file format: too short")
	}

	salt := fileContent[:16]
	nonceSize := 12
	nonce := fileContent[16 : 16+nonceSize]
	actualCiphertext := fileContent[16+nonceSize:]

	// argon2.IDKey signature: []byte password, []byte salt, uint32 time, uint32 memory, uint8 threads, uint32 keyLen
	derivedKey := argon2.IDKey(password, salt, time, memory, threads, keyLen)

	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return fmt.Errorf("failed to create cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("failed to create GCM: %w", err)
	}

	plaintext, err := aesgcm.Open(nil, nonce, actualCiphertext, nil)
	if err != nil {
		return fmt.Errorf("decryption failed: %w", err)
	}

	// Safe write: write to temp file first
	tmpFile := source + ".tmp"
	destFile, err := os.Create(tmpFile)
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}

	if _, err := destFile.Write(plaintext); err != nil {
		destFile.Close()
		os.Remove(tmpFile)
		return fmt.Errorf("failed to write decrypted data: %w", err)
	}
	destFile.Close()

	// Atomic rename
	if err := os.Rename(tmpFile, source); err != nil {
		return fmt.Errorf("failed to rename temp file: %w", err)
	}
	
	return nil
}
