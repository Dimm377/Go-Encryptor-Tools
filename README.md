# Go File Encryptor

A simple file encryptor and decryptor written in Go for personal use.

## Features

- Encrypt files with password protection
- Decrypt files using the correct password
- Uses AES-256-GCM encryption for security
- PBKDF2 key derivation with SHA-256 and 10,000 iterations
- Secure random salt and nonce generation

## Requirements

- Go 1.21 or higher

## Installation

1. Clone the repository:

```bash
git clone <your-repo-url>
cd Go-file-encryptor
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
./Go-file-encryptor help

# Encrypt a file
./Go-file-encryptor encrypt /path/to/file

# Decrypt a file
./Go-file-encryptor decrypt /path/to/file
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
