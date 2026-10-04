# Raspberry Pi Backend Server

This is a Go backend for a separate frontend. The frontend sends it requests, the backend talks to Raspberry Pi hardware, and then sends results or device data back over HTTP/JSON.

This could also be used as a backend server template.

## Project structure

- `cmd/server/` - starts the backend
- `internal/app/` - connects the parts of the backend at startup
- `internal/config/` - settings
- `internal/transport/http/` - HTTP routes and JSON requests and responses
- `internal/service/` - device actions and other backend logic
- `internal/device/` - Raspberry Pi hardware support
- `internal/auth/` - login and permissions, if the frontend needs accounts
- `internal/repository/` - saved data, if the backend needs storage
- `internal/domain/` - shared types and rules, if they become useful
- `deploy/systemd/` - files for running the backend as a service
- `test/integration/` - tests for parts working together

The project is just a starting layout right now. There is no Go server implementation yet. 

See [docs/architecture.md](docs/architecture.md) for how the backend components fit together.

## Recommended Hardware and Software

### Hardware
- Raspberry Pi 3, 4, 5, CM3, CM4, or CM5
- 512 MB RAM for a headless setup; 2 GB recommended for a desktop environment

### Software
- Raspberry Pi OS (64-bit Lite recommended)
- Go 1.25 or higher on the build machine
- GPIO access for the account running the backend, such as access to `/dev/gpiomem`