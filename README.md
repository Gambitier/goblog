# Blog API

## Description

This is a simple blog API built with Go and Fiber. It allows you to create, read, update, and delete blogs.

## Prerequisites

Before running the application, make sure you have the following tools installed:

- Go 1.x or higher
- PostgreSQL
- [golang-migrate](https://github.com/golang-migrate/migrate) for database migrations: 
  ```bash
  go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  ```
- [sqlc](https://github.com/sqlc-dev/sqlc) for type-safe SQL: 
  ```bash
  go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
  ```
- [jq](https://stedolan.github.io/jq/) for JSON processing:
  - Ubuntu/Debian: `sudo apt-get install jq`
  - Mac: `brew install jq`
  - Windows: `choco install jq`

## Usage

### Create a config.json file

Create a config.json file by copying the sample configuration:

```bash
cp sample.config.json config.json
```

### Setup Database

```bash
# Run database migrations
make migrate-up

# Generate type-safe database code
make sqlc
```

### Run the server

```bash
# Install dependencies
go mod tidy

# Start the server
make server
```

### Available Make Commands

- `make server` - Run the application
- `make migrate-up` - Apply database migrations
- `make migrate-down` - Rollback database migrations
- `make sqlc` - Generate database code
