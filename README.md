# my-budgety-gRPC

API endpoint for mobile app of My Budgety to consume. Also serve as a repo to learn gRPC framework in Go.

## Dev Setup

### Install Protocol Buffers Compiler

On Window:

Ensure choco is install and powershell terminal is in adminstrator mode

```bash
    choco install protoc
```

### Use protoc compiler to complie the Go files for the proto files

```bash
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out= --go-grpc_opt=paths=source_relative expanse/protos/expanse.proto
```

### Login in using google auth to create a ADC file locally

```bash
    gcloud auth application-default login
```

### Create a .env file in the project root folder

In the .env file have the following variables:

- PROJECT_ID
- PROJECT_NAME
- PORT

### Run server

```bash
    go run .
```

## How to run test locally

```bash
    go test ./tests/
```

## If there are problem with debugging tests

Open the launch.json file by going to Run > Add Configuration...
Edit the json file to include the following:

```
  "configurations": [
    {
      "name": "Launch Package",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceRoot}",
      "env": {},
      "args": [],
      "showLog": true
    },
    {
      "name": "Test Current File",
      "type": "go",
      "request": "launch",
      "mode": "test",
      "port": 2345,
      "host": "127.0.0.1",
      "program": "${file}",
      "env": {},
      "args": [],
      "showLog": true
    }
  ]
}
```
