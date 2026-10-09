# Changelog

## Unreleased

### Added

- Add initial config and app startup scaffolding for the backend bootstrap flow.
- Add a service shell and app wiring so the server can be started from the app layer instead of directly from main.
- Add an HTTP server with a `GET /healthz` endpoint that returns the service health status.