# my-budgety-gRPC

API endpoint for mobile app of My Budgety to consume. Also serve as a repo to learn gRPC framework in Go.

## Dev Setup

### Install Protocol Buffers Compiler

On Window:

Ensure choco is install and powershell terminal is in adminstrator mode

```bash
    choco install protoc
```

### Use protoc compiler

```bash
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out= --go-grpc_opt=paths=source_relative proto/helloworld.proto
```
