package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/akhilsharma90/go-file-encryption/filecrypt"
	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}
	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run the encrypt to encrypt a file, or decrypt to decrypt a file")
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("===========================================")
	fmt.Println("          GO FILE ENCRYPTOR")
	fmt.Println("    A Secure File Encryption Tool")
	fmt.Println("===========================================")
	fmt.Println()
	fmt.Println("A simple file encryptor and decryptor written in Go and for personal use only")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  go run . [command] [file_path]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  encrypt [file]    Encrypt a file with password protection")
	fmt.Println("  decrypt [file]    Decrypt a file using the correct password")
	fmt.Println("  help              Display this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run . encrypt document.txt")
	fmt.Println("  go run . decrypt document.txt")
	fmt.Println("  go run . help")
	fmt.Println()
	fmt.Println("Security Features:")
	fmt.Println("  - AES-256-GCM encryption")
	fmt.Println("  - PBKDF2 key derivation (10,000 iterations)")
	fmt.Println("  - Random salt and nonce generation")
	fmt.Println("  - Supports all file types")
	fmt.Println("===========================================")
}

func encryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide the path to the file to encrypt. for more info run the help command -> go run . help")
		os.Exit(0)
	}
	filePath := os.Args[2]
	if !fileExists(filePath) {
		panic("The file you are trying to encrypt does not exist")
	}
	password := getPassword()
	fmt.Println("\nEncrypting your file.....", filePath)
	filecrypt.Encrypt(filePath, password)
	fmt.Println("Your file is fully protected!")
}

func decryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide the path to the file to decrypt. for more info run the help command -> go run . help")
		os.Exit(0)
	}
	filePath := os.Args[2]
	if !fileExists(filePath) {
		panic("The file you are trying to decrypt does not exist")
	}
	password := getPassword()
	fmt.Println("\nDecrypting your file.....", filePath)
	filecrypt.Decrypt(filePath, password)
	fmt.Println("Your file decrypted successfully!")
}

func getPassword() []byte {
	fmt.Print("Enter password: ")
	password, _ := term.ReadPassword(0)
	fmt.Println("\nconfirm password: ")
	password2, _ := term.ReadPassword(0)
	if !validatePassword(password, password2) {
		fmt.Println("\nPasswords do not match, please try again")
		return getPassword()
	}
	return password
}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}
	return true
}

func fileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}
