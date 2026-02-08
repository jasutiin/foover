## Project Info

This project is for me to learn more about microservices architecture using Golang, gRPC for inter-service communication, Docker to easily run each microservice locally and test them out, and later on Kubernetes for container orchestration.

## Setup

To generate protobuf files from a .proto file, follow these steps:

1. Run `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest` only if you haven't installed protoc-gen-go-grpc on your PC yet.
2. Run `cd <service_name>/proto` to go to the service's proto directory in which you want to generate protobuf files for, assuming the **.proto** file is inside this directory.
3. Create a new directory called **pb/** by doing `mkdir pb`.
4. Run `protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative <file_name>.proto` to generate the protobuf files.
   - The `--go_out` flag specifics where protoc will generate the **.pb.go** file for the service. In this case, it will be in the **pb/** directory we made in step 3.
   - The `--go-grpc_out` flag specifies where protoc will generate the **\_grpc.pb.go** file for the service. In this case, it will be in the **pb/** directory we made in step 3.
   - The `--go_opt` and `--go-grpc_opt` flags make it so that it doesn't generate a bunch of nested directories inside the **pb/** directory based on the module name in the **go.mod** file.

## Protobufs

### <service>.pb.go

This is where all the Go structs are for our protobuf messages. It has basic serialization/deserialization methods.

### <service>\_proto.pb.go

This is where our gRPC service interface is. This is what our server must implement.

Project file structure based off of this repo: https://github.com/rasadov/EcommerceAPI
