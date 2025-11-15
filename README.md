# Go Encryptor Tools

A simple file encryptor and decryptor written in Go for personal use.

## Features

- Encrypt files with password protection
- Decrypt files using the correct password
- Uses AES-256-GCM encryption for security
- PBKDF2 key derivation with SHA-256 and 10,000 iterations
- Secure random salt and nonce generation

## Requirements

- Go 1.21 or higher

## ⚠️ Important Warning

**This tool overwrites your original files during encryption and decryption.** Always keep backups of your important files before using this tool.

## Platform Compatibility

This application is cross-platform and works on Linux, Windows, and macOS. The Go code uses standard library functions that are compatible across platforms.

## Testing

This repository includes a dummy file (`img.jpg`) that you can use to test whether the encryption tool is working properly.

## Installation

1. Clone the repository:

```bash
git clone <your-repo-url>
cd Go-Encryptor-Tools
```

2. Install dependencies:

```bash
go mod tidy
```

## Usage

### Running directly:

```bash
# Show help
go run . help

# Encrypt a file
go run . encrypt /path/to/file

# Decrypt a file
go run . decrypt /path/to/file
```

### Building and running:

```bash
# Build the application
go build .

# Show help
./Go-Encryptor-Tools help

# Encrypt a file
./Go-Encryptor-Tools encrypt /path/to/file

# Decrypt a file
./Go-Encryptor-Tools decrypt /path/to/file
```

### Example with included test file:

```bash
# Encrypt the included test image
go run . encrypt img.jpg

# Decrypt the image (use the same password as used for encryption)
go run . decrypt img.jpg
```

## How it works

1. **Encryption**:

   - A random 16-byte salt is generated
   - Password is used with PBKDF2 to derive a 32-byte key
   - A random 12-byte nonce is generated
   - File content is encrypted using AES-256-GCM
   - Final file format: [16-byte salt + 12-byte nonce + ciphertext]

2. **Decryption**:
   - Salt and nonce are extracted from the beginning of the encrypted file
   - Same PBKDF2 process is used with the provided password and extracted salt
   - Ciphertext is decrypted using AES-256-GCM with the derived key and nonce

## Security

- Uses AES-256-GCM for authenticated encryption
- PBKDF2 with 10,000 iterations for key derivation
- SHA-256 as the hash function
- Random salt and nonce for each encryption
- Secure random number generation
- Support all file type

## License

This project is for learn and personal use, and free for use and contrib
