# API Template

The `api-template` project is a Golang-based template for building versioned RESTful APIs using the Gin framework. This template supports multiple API versions, organised for modularity and maintainability.

## Project Structure

Please check for each directory as I provide how they're used in separated MD files. So I won't explain too much details in this one.

```bash
api-template/
├── cmd/
│   └── main.go
├── internal/
│   ├── app/
│   │   ├── v1/
│   │   │   ├── handlers/
│   │   │   │   ├── user/
│   │   │   │   │   └── user.go
│   │   │   │   ├── product/
│   │   │   │   │   └── product.go
│   │   │   ├── middleware/
│   │   │   ├── models/
│   │   │   ├── repositories/
│   │   │   ├── services/
│   │   │   └── router.go
│   │   ├── v2/
│   │   │   ├── handlers/
│   │   │   │   ├── user/
│   │   │   │   │   └── user.go
│   │   │   │   ├── product/
│   │   │   │   │   └── product.go
│   │   │   ├── middleware/
│   │   │   ├── models/
│   │   │   ├── repositories/
│   │   │   ├── services/
│   │   │   └── router.go
│   ├── logging/
│   │   └── logger.go
│   ├── errors/
│   │   └── errors.go
├── config/
├── platform/
├── pkg/
├── test/
│   ├── integration/
│   ├── e2e/
│   ├── data/
│   └── utils/
├── go.mod
└── go.sum
```

## Directories and Their Purposes

### cmd/
Contains the main application entry point.

### internal/
Contains packages intended for internal use within the project.

- **app/**: Core application logic, organised by version (e.g., v1, v2).
  - **v1/**: Handlers, middleware, models, repositories, services.
  - **v2/**: Handlers, middleware, models, repositories, services.
- **logging/**: Custom logging setup and configuration.
- **errors/**: Custom error definitions and handling.

### config/
Contains configuration files and settings.

### platform/
Contains code and configurations for external dependencies or services.

### pkg/
Contains reusable utility and helper packages shared across the application.

### test/
Organises and stores test files and resources for the application.

- **integration/**: Integration tests.
  - **v1/**: Integration tests for API v1.
  - **v2/**: Integration tests for API v2.
- **e2e/**: End-to-end tests.
  - **v1/**: E2E tests for API v1.
  - **v2/**: E2E tests for API v2.
- **data/**: Mock data and fixtures used for testing.
- **utils/**: Helper functions and utilities to facilitate testing.

## Getting Started

### Prerequisites
- Go (1.16 or later)
- Gin framework

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/eugeneChuBNE/ginTemplate
   cd api-template
   ```

OR:
    ```bash
   git clone https://github.com/vietnix-org/api-template.git
   cd api-template
   ```

2. Install dependencies
    ```bash
    go mod tidy
    ```

### Running the Application

Start the application servers for different API versions:

```bash
go run cmd/main.go
```

### Running tests

Run the tests to ensure everything is working:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss any changes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Authors

- Eugene Chu - [eugeneChuBNE](https://github.com/eugeneChuBNE)
- As a contribution to Vietnix JSC

Ho Chi Minh city, Vietnam | 2024
