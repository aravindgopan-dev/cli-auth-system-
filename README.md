# CLI Auth System with 2FA

A Go CLI application that supports user registration, authentication, 2FA (TOTP), and session management. Built using dependency injection (DI) .

## Setup & Running

### Using Docker

Start the PostgreSQL database and run the interactive CLI application:

```bash
make run
# or: sudo docker compose run app
```

To stop and remove containers:

```bash
make clean
# or: sudo docker compose down -v
```

## CLI Commands

### Guest State
- `register` - Create a new account
- `login` - Log in with username/password (and 2FA if enabled)
- `help` - Show available commands
- `exit` - Quit application

### Authenticated State
- `whoami` - Show session details and expiration time
- `enable-2fa` - Generate and display a TOTP QR code to scan
- `disable-2fa` - Disable TOTP
- `logout` - End current session
- `help` - Show available commands

## Configuration

The app uses the following environment variables:
- `DATABASE_URL` (default: `postgres://postgres:postgres@localhost:5432/auth_db?sslmode=disable`)
- `SESSION_DURATION` (default: `5m`)
- `LOCKOUT_DURATION` (default: `1m`)

## Architecture & Design Decisions

- **Dependency Injection (DI)**: Components are decoupled. The database repository is injected into the service, and the service is injected into the CLI handler, avoiding global state and keeping code testable.
- **Database Resilience**: Multi-container environments often suffer from startup lag. The application retries the database connection 10 times with a delay on startup to prevent crashing.
- **Self-Contained Migrations**: Migration files are executed automatically in alphabetical order on startup, ensuring the database schema is always set up correctly.
