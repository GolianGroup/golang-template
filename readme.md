# Go Project Template

A clean and structured Go project template using Fiber framework and EntGo.

## 📁 Project Structure

```bash
.
├── cmd/ # Command line interface
│ ├── root.go # Root cobra command
│ └── serve.go # Server start command
│
├── config/ # Configuration
│ └── config.yml # Configuration file
│
├── internal/ # Private application code
│ ├── di/ # required injections
│ ├── pkg/ # logic and code implementation of project
│ │ ├── config/ # config model
│ │ ├── errors/ # error models
│ │ ├── handlers/ # HTTP Layer
│ │ │ ├── api/ # implementation in HTTP Layer
│ │ │ └── middlewares/ # middleware in HTTP Layer
│ │ │
│ │ ├── repositories/ # Data access layer
│ │ ├── models/ # EntGo layer of schema
│ │ ├── services/ # Business logic
│ │ ├── utils/ # Utility functions
│ │ └── helpers/ # Helper functions
│ │
│ └── app.go # starter of project
│
├── main.go # Application entry point
└── Makefile.ent # make file for common commands of entgo
```

## 🔧 Technology Stack

- [Fiber](https://gofiber.io/) - Web framework
- [Uber FX](https://uber-go.github.io/fx/) - Dependency injection
- [Viper](https://github.com/spf13/viper) - Configuration management
- [Cobra](https://github.com/spf13/cobra) - CLI commands

## 🏗️ Architecture

This template follows clean architecture principles:

1. **Domain Layer** (`internal/pkg/models/`)
   - Business entities
   - Repository interfaces
   - Service interfaces

2. **Application Layer** (`internal/pkg/services/`)
   - Business logic implementation
   - Use case orchestration
   - Domain service implementation

3. **Infrastructure Layer** (`internal/pkg/repositories/`)
   - Database implementations
   - External service integrations
   - Repository implementations

4. **Interface Layer** (`internal/pkg/handlers/`)
   - HTTP handlers
   - Middleware
   - Route definitions

## 🚀 Getting Started

### installation and run

1. install dependencies

   ```bash
   go mod tidy
   ```

2. configure environment variables in config.yml

   ```bash
   nano config/config.yml.example
   nano config/config.yml
   ```

3. run the server

   ```bash
   go run main.go serve
   ```

4. Also you can set environment variable mappings by writing the following command before executing command `go run main.go serve`:

   ```bash
   APP_SERVER_PORT=8004 APP_SERVER_HOST=0.0.0.0 go run main.go serve
   ```
