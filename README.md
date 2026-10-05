# Koala - Raspberry Pi Backend Server

Koala is a Go backend server designed to run on a Raspberry Pi. A separate frontend can send HTTP/JSON requests to the backend, which can communicate with Raspberry Pi hardware and return device data or operation results.

For example, a frontend could request an action such as turning an LED or motor on or off, and Koala would handle the corresponding GPIO operation.

I might add a frontend to this project in the future.

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
