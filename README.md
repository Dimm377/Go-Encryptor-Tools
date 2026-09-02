# Go-Encryptor

A small command-line tool for encrypting and decrypting files with Go. It is a personal-use project designed to make the encryption flow easy to inspect and understand.

## Features

- AES-256-GCM authenticated encryption
- Argon2id password-based key derivation
- Random 16-byte salt and 12-byte nonce for every encryption
- Hidden password input on supported terminals
- Cross-platform source layout for Linux, macOS, and Windows
- Atomic replacement through a temporary file

## Requirements

- Go 1.24 or newer
- Git

## Installation

```bash
git clone https://github.com/Dimm377/Go-Encryptor.git
cd Go-Encryptor
go build -o go-encryptor .
```

## Usage

Encrypt a file:

```bash
./go-encryptor encrypt path/to/file
```

Decrypt a file:

```bash
./go-encryptor decrypt path/to/file
```

Show help:

```bash
./go-encryptor help
```

During encryption, the tool asks for a password and confirmation. During decryption, enter the same password.

## How it works

1. Generate a random 16-byte salt.
2. Derive a 32-byte key with Argon2id using time `1`, memory `64 MiB`, and `4` threads.
3. Generate a random 12-byte AES-GCM nonce.
4. Encrypt the file contents with AES-256-GCM.
5. Write `salt || nonce || ciphertext` through a temporary file, then atomically replace the source file.

Decryption derives the same key, authenticates the ciphertext, and restores the plaintext through the same temporary-file approach.

## Important notes

- The source file is replaced in place. Keep a separate backup before testing.
- Losing the password means the encrypted contents cannot be recovered by this tool.
- The project has not received a formal security audit and is intended for personal use and learning.
- Do not treat this tool as a replacement for a mature, independently reviewed encryption utility.

## Development

```bash
go test ./...
go vet ./...
```

## License

No license is currently declared. All rights remain with the repository owner unless a license is added.
