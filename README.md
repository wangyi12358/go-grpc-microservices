# Go gRPC micro service
This is a simple gRPC micro service written in Go. It is a simple service that takes a string and returns the string in uppercase.

## Environment Installation

* Go 1.22
* protoc
* PostgreSQL 14

## Running the service

You can run the following command to run the service:

```bash
make start SERVICE=user
```

This will start the gRPC server on port 4001.