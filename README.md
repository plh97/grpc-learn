# grpc-study

Simple Go net/rpc example with separate server and client.

## Prerequisites

- Go 1.20+ (any recent Go version should work)

## Run

### Server

```bash
cd server
go run .
```

### Client

```bash
cd client
go run .
```

The client calls `Server.Add` on the server and prints the sum.
