package main

import (
	"Go-Encryptor-Tools/filecrypt"
	"bytes"
	"fmt"
	"os"
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

func printBanner(mode string) {
	const (
		ColorBlue  = "\033[1;34m" // Bold Blue
		ColorRed   = "\033[1;31m" // Bold Red for Decrypting
		ColorGreen = "\033[1;32m" // Bold Green for Encrypting
		ColorReset = "\033[0m"
	)

	var art string
	var color string

	switch mode {
	case "EP": // Encrypting Process
		color = ColorGreen
		art = `
    __               __   _                    
   / /   ____  _____/ /__(_)___  ____ _        
  / /   / __ \/ ___/ //_/ / __ \/ __ ` + "`" + `/        
 / /___/ /_/ / /__/ ,< / / / / / /_/ / _ _ _ _ 
/_____/\____/\___/_/|_/_/_/ /_/\__, (_|_|_|_|_)
                              /____/           `
	case "DP": // Decrypting Process
		color = ColorRed
		art = `
    __  __      __           __   _                    
   / / / /___  / /___  _____/ /__(_)___  ____ _        
  / / / / __ \/ / __ \/ ___/ //_/ / __ \/ __ ` + "`" + `/        
 / /_/ / / / / / /_/ / /__/ ,< / / / / / /_/ / _ _ _ _ 
\____/_/ /_/_/\____/\___/_/|_/_/_/ /_/\__, (_|_|_|_|_)
                                     /____/           `
	default: // Main Menu / Default
		color = ColorBlue
		art = `
   ______ ____                         __ 
  / ____// __ \ ________  __  ______  / /_
 / / __ / / / // ___/ __// / / / __ \/ __/
/ /_/ // /_/ // /__/ /  / /_/ / /_/ / /_  
\____/ \____/ \___/_/   \__, / .___/\__/  
                       /____/_/           
 
          v2.0 - Argon2id Edition`
	}

	fmt.Println(color + art + ColorReset)
	fmt.Println("==================================================================")
}

func printHelp() {
	printBanner("")
	fmt.Println()
	fmt.Println("A simple file encryptor and decryptor written in Go and for personal use only")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  go run . [command] [file_path]")
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
	fmt.Println("  - Argon2id key derivation (Time=1, Mem=64MB, Threads=4)")
	fmt.Println("  - Random salt and nonce generation")
	fmt.Println("  - Supports all file types")
	fmt.Println("===========================================")
}

func encryptHandle() {
	printBanner("EP")
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
	printBanner("DP")
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

func validatePassword(password1 []byte, password2 []byte) bool {
	return bytes.Equal(password1, password2)
}

func fileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}
