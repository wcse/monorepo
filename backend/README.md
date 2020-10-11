# WCSE - monorepo - backend

## Install Protobuf
### Step 1
To build and install the C++ Protocol Buffer runtime and the Protocol Buffer compiler (protoc) execute the following
Go to this (link)[https://github.com/protocolbuffers/protobuf/releases/tag/v3.12.4]
And download the pre-build binary for your OS

Unzip the file (read protoc-xxx/README.md)
```
1. copy protoc-xxx/bin/protoc to your bin directory
2. copy protoc-xxx/includes and put into the same level with bin directory
```

### Step 2
Install Go support for Google's protocol buffers
Go to this (link)[https://github.com/protocolbuffers/protobuf-go/releases/tag/v1.25.0]
And download the pre-build binary for your OS & copy to your bin directory

### Step 3
Install The Go language implementation of gRPC. HTTP/2 based RPC
Go to this (link)[https://github.com/grpc/grpc-go/releases/tag/v1.31.0]
And download the pre-build binary for your OS & copy to your bin directory

### Step 4
Install gRPC for Web Clients
Go to this (link)[https://github.com/grpc/grpc-web/releases/tag/1.2.1]
And download the pre-build binary for your OS & copy to your bin directory

### Step 5
Install 
We're using version (v0.4.1)[https://github.com/envoyproxy/protoc-gen-validate/releases/tag/v0.4.1]
You can checkout exactly this version and run `go install` in the root on this project
```
go install github.com/envoyproxy/protoc-gen-validate
```

### Step 6
Install ent for generate code for storage
Go to this (link)[https://github.com/facebook/ent/releases/tag/v0.4.0]
You can checkout the repository to your GOPATH
```
cd cmd/entc
go install
```

### Summary
```bash
protoc --version
# libprotoc 3.12.4

protoc-gen-go --version
# protoc-gen-go v1.25.0
```