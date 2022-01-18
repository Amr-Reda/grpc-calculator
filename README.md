# Calculator GRPC

Simple calculator using protocol buffers.

## Functionallity

* Sum API (gRPC Unary).

## Getting Started

### Dependencies

* Golang.
* Protoc (Find the correct protocol buffers version based on your Linux Distripution: https://github.com/google/protobuf/releases).
  ```
    # Make sure you grab the latest version
    curl -OL https://github.com/google/protobuf/releases/download/v3.19.3/protoc-3.19.3-linux-x86_64.zip
    # Unzip
    unzip protoc-3.19.3-linux-x86_64.zip -d protoc3
    # Move protoc to /usr/local/bin/
    sudo mv protoc3/bin/* /usr/local/bin/
    # Move protoc3/include to /usr/local/include/
    sudo mv protoc3/include/* /usr/local/include/
    # Optional: change owner
    sudo chown [user] /usr/local/bin/protoc
    sudo chown -R [user] /usr/local/include/google
  ```

### Installing

```
go get -u google.golang.org/grpc
go get -u google.golang.org/protobuf
chmod u+x generate.sh
./generate.sh
```

### Executing program

```
code blocks for commands
```