package internal

import (
	"fmt"
	"net"

	"github.com/jasutiin/foover/user-service/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	pb.UnimplementedUserServiceServer
}

func ListenGRPC () error {
	port := 8080
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...) // take all the items in opts and apply them as args
	fmt.Println("created new server")

	pb.RegisterUserServiceServer(server, &grpcServer{})
	reflection.Register(server) // to support testing w/ grpcurl
	fmt.Println("registered server with grpcServer struct")

	return server.Serve(lis)
}