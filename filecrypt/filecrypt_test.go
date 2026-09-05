package filecrypt

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestEncryptDecryptRoundTrip(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "test.txt")
	originalContent := []byte("hello world this is a test string")

	err := os.WriteFile(filePath, originalContent, 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	password := []byte("supersecret")

	// Encrypt
	err = Encrypt(filePath, password)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	encryptedContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read encrypted file: %v", err)
	}

	if bytes.Equal(encryptedContent, originalContent) {
		t.Fatalf("Encrypted file content is identical to original")
	}

	// Decrypt
	err = Decrypt(filePath, password)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	decryptedContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read decrypted file: %v", err)
	}

	if !bytes.Equal(decryptedContent, originalContent) {
		t.Fatalf("Decrypted content does not match original. Got %q, want %q", decryptedContent, originalContent)
	}
}

func TestDecryptWrongPassword(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "test.txt")
	originalContent := []byte("hello world")

	os.WriteFile(filePath, originalContent, 0644)

	err := Encrypt(filePath, []byte("correctpassword"))
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	err = Decrypt(filePath, []byte("wrongpassword"))
	if err == nil {
		t.Errorf("Decrypt with wrong password should have failed")
	}
}

func TestDecryptCorruptedCiphertext(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "test.txt")
	os.WriteFile(filePath, []byte("hello world"), 0644)
	
	err := Encrypt(filePath, []byte("password"))
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	content, _ := os.ReadFile(filePath)
	// Corrupt a byte in the ciphertext part (salt=16, nonce=12)
	if len(content) > 30 {
		content[30] ^= 0xff
		os.WriteFile(filePath, content, 0644)
	}

	err = Decrypt(filePath, []byte("password"))
	if err == nil {
		t.Errorf("Decrypt with corrupted ciphertext should have failed")
	}
}

func TestEncryptEmptyFile(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "empty.txt")
	os.WriteFile(filePath, []byte(""), 0644)

	password := []byte("password")
	err := Encrypt(filePath, password)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}
	
	err = Decrypt(filePath, password)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	decryptedContent, _ := os.ReadFile(filePath)
	if len(decryptedContent) != 0 {
		t.Fatalf("Expected empty file, got length %d", len(decryptedContent))
	}
}

func TestDecryptMalformedInput(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "short.txt")
	// Too short to contain salt and nonce (28 bytes)
	os.WriteFile(filePath, []byte("short"), 0644)

	err := Decrypt(filePath, []byte("password"))
	if err == nil {
		t.Errorf("Decrypt with too short input should have failed")
	}
}
