# Contributing Guide

Thank you for helping build the Raspberry Pi Backend! 

## Code Style & Architecture
- **Keep `cmd/` thin:** The main entry point should only handle initialization and startup.
- **Isolate Hardware:** All direct GPIO or hardware-level code belongs strictly inside `internal/device/`.
- **No Semicolons:** Follow standard idiomatic Go formatting. Run `go fmt ./...` before committing.
- **Security First:** Never hardcode secrets, Wi-Fi credentials, or tokens. Use configuration files or environment variables.

## Branching & Pull Requests
1. Fork the repo and create your branch (`feature/amazing-feature`).
2. Ensure your code passes all integration tests (`go test ./test/integration/...`).
3. Open a Pull Request with a clear description of what changed.
