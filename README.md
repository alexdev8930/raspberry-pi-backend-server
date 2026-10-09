# Raspberry Pi Backend Server

This is a general purpose Go backend server designed to run on a Raspberry Pi. It provides an HTTP/JSON API for communication between a frontend and the backend, with support for interacting with Raspberry Pi hardware.

For example, a frontend can send requests to control an LED or motor, and the backend handles the corresponding GPIO operations.

The project can serve as a reusable Go backend template for building other applications.

A frontend may be added to this project in the future.

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

Still in active dev so somthings might not be done or implemented yet.

See [docs/architecture.md](docs/architecture.md) for how the backend components fit together.  
see [CHANGELOG.md](CHANGELOG.md) for everything I've done so far.

## Recommended Hardware and Software

### Hardware
- Raspberry Pi 3, 4, 5, CM3, CM4, or CM5
- 512 MB RAM for a headless setup; 2 GB recommended for a desktop environment

### Software
- Raspberry Pi OS (64-bit Lite recommended)
- Go 1.25 or higher on the build machine
- Optional: GPIO access for the account running the backend, such as access to `/dev/gpiomem`
