# My Budgety gRPC Server

API endpoint for mobile app of My Budgety to consume. Also serve as a repo to learn gRPC framework in Go.

## Dev Setup

### Login in using google auth to create a ADC file locally

```bash
    gcloud auth application-default login
```

### Create a .env file in the project root folder

In the .env file have the following variables:

- PROJECT_ID
- PROJECT_NAME
- PORT

### Setup the Pre-Commit Hook

In the .git folder, remove the .sample extension from pre-commit file and copy the code in pre-commit.sample into it.
If the .git folder isn't visible, then ensure hidden folder can be view (Google it if you don't know how on your machine XD).

### Run server

```bash
    go run .
```

## How to run test locally

```bash
    go test ./tests/
```

## Steps if there been changes to the proto files and they need to be re-generated

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

## If there are problem with debugging tests in VS Code IDE

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
