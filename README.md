# Product Service

This is a Go project that uses Echo framework and Swagger for API documentation.

## Swagger Update

We use [echo-swagger](https://github.com/swaggo/echo-swagger) for API documentation.

To install the Swagger tool, run the following command:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

To generate the Swagger documentation, run the following command:

```bash
swag init --parseDependency --parseInternal --parseDepth 5 -g main.go
```

After starting the server, you can view the Swagger UI at:

```
http://localhost:1323/swagger/index.html
```

## Running the Server

To start the server, run the following command:

```bash
go run main.go
```

This will start the server on port 1323. You can access the API endpoints as per the Swagger documentation.

## Frontend Setup and Start

To start the frontend take a  look under the folder frontend

[README.md Frontend](./frontend/README.md)

