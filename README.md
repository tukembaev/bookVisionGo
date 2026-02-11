# BookVisionGo

A Go-based web application for book management with PostgreSQL database integration.

## Overview

BookVisionGo is a RESTful API service built with Go that provides book management capabilities. The project uses a clean architecture pattern with proper separation of concerns.

## Tech Stack

- **Go 1.25.6** - Programming language
- **PostgreSQL** - Primary database
- **Gin Framework** - HTTP web framework
- **SQLC** - SQL query builder
- **Viper** - Configuration management
- **Database Migrations** - Schema versioning

## Project Structure

```
bookVisionGo/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── db/
│   │   ├── migrations/          # Database migration files
│   │   └── queries/             # SQLC generated queries
│   ├── domain/                  # Business logic layer
│   └── handler/                 # HTTP handlers
├── pkg/                         # Reusable packages
├── .env                         # Environment variables
├── .gitignore                   # Git ignore rules
├── go.mod                       # Go module file
├── go.sum                       # Go dependencies checksum
├── sqlc.yaml                    # SQLC configuration
└── README.md                    # Project documentation
```

## Features

- Database migration management
- Configuration management with environment variables
- Clean architecture with separation of concerns
- PostgreSQL integration
- Extensible design for future features

## Getting Started

### Prerequisites

- Go 1.25.6 or higher
- PostgreSQL database
- Migration tool (golang-migrate)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/tukembaev/bookVisionGo.git
cd bookVisionGo
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your database configuration
```

4. Run database migrations:
```bash
go run cmd/api/main.go
```

### Configuration

The application uses environment variables for configuration. Create a `.env` file in the root directory:

```env
# Server Configuration
SERVER_PORT=8080
GIN_MODE=debug

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=bookvisiongo
DB_SSLMODE=disable

# JWT Configuration
JWT_SECRET=your_jwt_secret_key
JWT_EXPIRES_IN=24
```

### Running the Application

```bash
# Run the main application
go run cmd/api/main.go

# Build for production
go build -o bookvisiongo cmd/api/main.go
./bookvisiongo
```

## Database

The project uses PostgreSQL as the primary database. Database migrations are managed through the `internal/db/migrations` directory.

### Migration Commands

```bash
# Run migrations up
migrate -path internal/db/migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up

# Run migrations down
migrate -path internal/db/migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" down

# Create new migration
migrate create -ext sql -dir internal/db/migrations -seq migration_name
```

## Development

### Code Generation

This project uses SQLC for type-safe SQL queries:

```bash
# Generate Go code from SQL queries
sqlc generate
```

### Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## API Documentation

The API endpoints will be documented here as they are implemented.

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Author

- [tukembaev](https://github.com/tukembaev)

## Acknowledgments

- [Go](https://golang.org/) - The programming language
- [Gin](https://gin-gonic.com/) - HTTP web framework
- [SQLC](https://sqlc.dev/) - SQL query builder
- [Viper](https://github.com/spf13/viper) - Configuration management
