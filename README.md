# monorepo README.md

# Monorepo for Xwing and Deathstar Services

This monorepo contains two services: **Xwing** and **Deathstar**. Each service has a single endpoint `/status` that returns specific status messages.

## Services

### Xwing
- **Endpoint**: `/status`
- **Response**: "systems normal"

### Deathstar
- **Endpoint**: `/status`
- **Response**: "this battle station is fully operational"

## Project Structure

```
monorepo
├── services
│   ├── xwing
│   ├── deathstar
├── go.mod
├── Makefile
└── README.md
```

## Getting Started

1. Clone the repository.
2. Navigate to the `monorepo` directory.
3. Use the `Makefile` to build and run the services.

## Building the Services

To build the services, run:

```
make build
```

## Running the Services

To run the services, use:

```
make run
```

## License

This project is licensed under the MIT License.