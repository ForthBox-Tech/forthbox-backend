# Forthbox Backend

This repository contains the Go backend for the Forthbox account and verification API. It provides user registration, login, password management, email/mobile verification code delivery, JWT-based authentication, and database migration utilities.

## Overview

The service is built with:

- Go `1.17`
- Gin for HTTP routing
- GORM v1 for MySQL access
- `go-ini` for configuration loading
- Mailgun for email delivery
- JWT for stateless authentication

The HTTP server listens on `0.0.0.0:<HttpPort>` and uses `Asia/Shanghai` as its runtime timezone.

## Repository Layout

```text
.
|-- app
|   |-- http
|   |   |-- controller     # Request handlers and response wrappers
|   |   `-- middleware     # HTTP middleware
|   |-- model              # GORM models and database access
|   `-- service            # Business logic for auth, users, mail, and tokens
|-- cmd
|   |-- migrate            # Database migration entrypoint
|   `-- server             # API server entrypoint
|-- conf                   # Environment configuration files
|-- docs                   # Supplemental project notes
|-- pkg
|   |-- setting            # INI configuration bootstrap
|   `-- util               # Shared helpers
`-- tools                  # Supporting scripts and helper assets
```

## Runtime Flow

1. Configuration is loaded from `conf/app.ini` during process startup.
2. The database connection is initialized automatically from the configured MySQL DSN.
3. Gin is started with the configured run mode and HTTP timeouts.
4. CORS is relaxed in development and restricted to known Forthbox domains in non-development environments.
5. Authentication APIs issue JWTs after successful login or signup.
6. The server shuts down gracefully on `SIGINT` and `SIGTERM`.

## Configuration

The application expects `conf/app.ini` to exist. The repository includes environment-specific samples:

- `conf/app_dev.ini`
- `conf/app_prd.ini`

Create the runtime config before starting the service:

```powershell
Copy-Item conf\app_dev.ini conf\app.ini
```

### Config Sections

#### `[app]`

- `IsDev`: enables development behavior such as returning verification codes in API responses
- `DebugLevel`: enables verbose GORM logging when greater than `0`
- `LogSavePath`, `LogSaveName`, `TimeFormat`: reserved logging metadata

#### `[server]`

- `RunMode`: Gin mode such as `debug` or `release`
- `HttpPort`: HTTP listening port
- `ReadTimeout`: request read timeout in seconds
- `WriteTimeout`: response write timeout in seconds

#### `[database]`

- `UserDsn`: MySQL DSN used by GORM

#### `[mail]`

- `Domain`: Mailgun sending domain
- `MGkey`: Mailgun API key

## Local Development

### Prerequisites

- Go `1.17`
- A reachable MySQL instance
- Valid Mailgun credentials if email delivery is needed

### Start the API

```powershell
Copy-Item conf\app_dev.ini conf\app.ini
go run .\cmd\server\main.go
```

By default, the development config runs the API on `http://localhost:8080`.

### Run Database Migrations

The migration command creates or updates the following tables:

- `users`
- `verify_tokens`
- `pics`

Run:

```powershell
Copy-Item conf\app_dev.ini conf\app.ini
go run .\cmd\migrate\db.go
```

## Docker

The repository includes a multi-stage Docker build:

```powershell
docker build -t forthbox-be .
docker run --rm -p 8080:8080 forthbox-be
```

The image:

- builds the server binary from `cmd/server/main.go`
- copies the `conf` directory into the runtime image
- defaults `conf/app.ini` from `conf/app_dev.ini`
- exposes port `8080`
- sets the container timezone to `Asia/Shanghai`

## API Endpoints

### Health / Root

- `GET /`

### User APIs

