# gRPC CRUD Example 


This is an example project demonstrating basic CRUD (Create, Read, Update, Delete) operations using gRPC with an in-memory database.

## Prerequisites

- Go programming language (version go1.19 or above)
- Protobuf compiler (`protoc`) (version 3.21.5 or above)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/akshay-singla/gRPC-example.git
   cd gRPC-example
   ``` 

## Usage

1. Start the server:

    ```bash
   go run server.go
   ``` 

2. Start the Client:

    ```bash
   go run client/main.go
   ``` 


## Testing
To run the unit tests, use the following command:

    go test ./...
