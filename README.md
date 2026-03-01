# yolo

**yolo** - A fast and simple Golang project scaffolding tool

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Features

- 🚀 Quick project initialization with a single command
- 📁 Standard Golang project structure with best practices
- 🎯 Built-in templates for common web service patterns
- ⚙️ Automatic Go module setup
- 🔧 Configurable port and project name
- 📝 Verbose mode for debugging
- 🧪 Health check endpoint included

## Installation

```bash
go install github.com/jotiao/yolo@latest
```

By default, `yolo` is installed in the `$GOPATH/bin` or `$GOBIN` directory. Make sure this directory is in your `PATH`.

## Usage

### Create a New Project

```bash
yolo init -name <project_name> -port <project_port>
```

**Options:**
- `-name` : Project name (required)
- `-port` : Project port (default: 8080)
- `-v` : Enable verbose output

**Example:**
```bash
yolo init -name myapi -port 8080
```

### Check Version

```bash
yolo version
```

## Project Structure

The generated project follows Golang best practices:

```
myapi/
├── cmd/
│   └── server.go          # Application entry point
├── pkg/
│   ├── config/            # Configuration management
│   │   ├── config.go
│   │   ├── define.go
│   │   ├── dev.yaml
│   │   └── log.go
│   ├── controller/        # HTTP handlers
│   │   ├── controller.go
│   │   └── typedef.go
│   ├── model/             # Data models
│   │   └── base.go
│   ├── router/            # HTTP routing
│   │   ├── middleware.go
│   │   └── url.go
│   ├── service/           # Business logic
│   │   └── service.go
│   └── util/
│       └── cm/            # Common utilities
│           └── cm.go
├── etc/                   # Configuration files
│   └── dev.yaml
├── script/                # Build and deployment scripts
├── go.mod
└── go.sum
```

## Running Your Project

```bash
cd myapi
go run cmd/server.go
```

## Testing

The generated project includes a liveness health check endpoint:

```bash
curl http://127.0.0.1:8080/v1/liveness
```

Response:
```json
{
  "status": "ok"
}
```

## Configuration

Configuration is managed through YAML files in the `etc/` directory. The default configuration is in `etc/dev.yaml`.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Changelog

### 2.0.0 (Latest)
- Project scaffolding with standard structure
- Built-in templates for web services
- Automatic Go module initialization
- Health check endpoint
- Verbose mode for debugging
- Port configuration support
