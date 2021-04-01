package grpc

import (
	"log"
	"net"

	"github.com/1024casts/snake/internal/conf"

	"google.golang.org/grpc"

	"github.com/1024casts/snake/api/grpc/user/v1"
	"github.com/1024casts/snake/internal/service"
)

// Init new grpc server
func Init(c *conf.Config, svc *service.Service) *grpc.Server {
	// todo: get addr from conf
	lis, err := net.Listen("tcp", "127.0.0.1:12349")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	v1.RegisterUserServiceServer(grpcServer, service.UserSvc)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve grpc server: %v", err)
	}
	log.Println("serve grpc server is success, port:1234")

	return grpcServer
}
