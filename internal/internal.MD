# Internal Directory

The `internal` directory contains packages that are intended for use only within your project. This helps enforce encapsulation by preventing other projects from importing these packages.

## Common Uses

1. **App:**
   - Core application logic, organised by version (e.g., v1, v2).
   - Handlers, middleware, models, repositories, and services.

2. **Logging:**
   - Custom logging setup and configuration.

3. **Errors:**
   - Custom error definitions and handling.

## Example Structure

- **app/**: Versioned application logic (e.g., v1, v2)
  - **v1/**: Handlers, middleware, models, repositories, services
  - **v2/**: Handlers, middleware, models, repositories, services
- **logging/**: Logging setup (`logger.go`)
- **errors/**: Custom error handling (`errors.go`)

The `internal` directory organizes the project's core logic and utilities, enhancing maintainability and preventing unintended external usage.
