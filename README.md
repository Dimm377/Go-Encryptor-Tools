# Go File Encryptor

A secure file encryption and decryption tool written in Go. This application provides a simple command-line interface for protecting files using strong encryption standards.

## Features

- **AES-256-GCM encryption**: Military-grade encryption standard
- **PBKDF2 key derivation**: 10,000 iterations for strong password hashing
- **Random salt and nonce generation**: Prevents rainbow table attacks
- **Cross-platform support**: Works on Windows, Linux, and macOS
- **Secure password input**: Passwords are hidden during entry
- **Supports all file types**: Encrypt any file format

## Prerequisites

- Go 1.16 or higher
- Git (for cloning the repository)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/Go-Encryptor-Tools.git
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
go run . encrypt document.txt
```

Decrypt the same file:
```bash
go run . decrypt document.txt
```

## Security Features

- **Encryption Algorithm**: AES-256-GCM (Galois/Counter Mode)
- **Key Derivation**: PBKDF2 with SHA-256 hash function and 10,000 iterations
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
- This tool was created for personal use only

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

If you encounter any issues or have questions, please open an issue in the GitHub repository.
