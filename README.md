# GOCRYPT

A secure file encryption and decryption tool written in Go. This application provides a simple command-line interface for protecting files using strong encryption standards.

## Features

- **AES-256-GCM encryption**: Industry standards encryption
- **Argon2id key derivation**: Memory-hard key derivation function (Time=1, Memory=64MB, Threads=4) for modern security.
- **Random salt and nonce generation**: Prevents rainbow table attacks
- **Cross-platform support**: Works on Windows, Linux, and macOS
- **Secure password input**: Passwords are hidden during entry
- **Supports all file types**: Encrypt any file format

## Prerequisites

- Go 1.16 or higher (https://go.dev/doc/install)
- Git (for cloning the repository)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Dimm377/Go-Encryptor-Tools.git
   cd Go-Encryptor-Tools
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

The tool provides three main commands:

### Encrypt a file
```bash
go run . encrypt [file_path]
```

### Decrypt a file
```bash
go run . decrypt [file_path]
```

### Show help
```bash
go run . help
```

### Examples

Encrypt a text file:
```bash
go run . encrypt testing.txt
```

Decrypt the same file:
```bash
go run . decrypt testing.txt
```
<img width="1920" height="1080" alt="Image" src="https://github.com/user-attachments/assets/4caecfa4-755f-4902-876a-896ca47fab18" />
## Security Features

- **Encryption Algorithm**: AES-256-GCM (Galois/Counter Mode)
- **Key Derivation**: Argon2id (Memory-hard function) with Time=1, Memory=64MB, Threads=4
- **Backward Compatibility**: Files encrypted with previous versions (PBKDF2) are NOT compatible.
- **Salt Generation**: 16-byte random salt for each encryption
- **Nonce Generation**: 12-byte random nonce for each encryption
- **Password Security**: Passwords are not echoed to the terminal during entry

## How It Works

1. **Encryption Process**:
   - A random 16-byte salt is generated
   - PBKDF2 is used to derive a key from the password using the salt
   - A random 12-byte nonce is generated
   - The file content is encrypted using AES-256-GCM
   - The salt, nonce, and ciphertext are concatenated and written to the file

2. **Decryption Process**:
   - The salt, nonce, and ciphertext are extracted from the file
   - The same PBKDF2 process is used to derive the key from your password
   - AES-256-GCM decryption is performed on the ciphertext
   - The original file content is restored

## Important Notes

- The original file will be overwritten with encrypted/decrypted content
- **Always keep a backup** of important files before encryption
- Losing your password will result in permanent data loss
- Make sure you remember the password you entered
- This tool was created for personal use only

## Support

If you encounter any issues or have questions, please open an issue in the GitHub repository.

