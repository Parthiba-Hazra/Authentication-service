package transport

import (
	"context"
	"log"
	"net"

	pb "github.com/Parthiba-Hazra/user-management-microservice/api"
	"github.com/Parthiba-Hazra/user-management-microservice/internal/endpoint"
	"google.golang.org/grpc"
)

type grpcServer struct {
	pb.UnimplementedUserManagementServiceServer
	registerUserHandler grpc.MethodDesc
	loginHandler        grpc.MethodDesc
}

func NewGRPCServer(endpoints endpoint.Endpoints) pb.UserManagementServiceServer {
	return &grpcServer{
		registerUserHandler: grpc.MethodDesc{
			MethodName: "RegisterUser",
			Handler:    makeRegisterUserHandler(endpoints),
		},
		loginHandler: grpc.MethodDesc{
			MethodName: "Login",
			Handler:    makeLoginHandler(endpoints),
		},
	}
}

func (s *grpcServer) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	return s.registerUserHandler.Handler(s, ctx, req)
}

func (s *grpcServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return s.loginHandler.Handler(s, ctx, req)
}

func makeRegisterUserHandler(endpoints endpoint.Endpoints) grpc.UnaryHandler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return endpoints.RegisterUserEndpoint(ctx, req.(*pb.RegisterUserRequest))
	}
}

func makeLoginHandler(endpoints endpoint.Endpoints) grpc.UnaryHandler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return endpoints.LoginEndpoint(ctx, req.(*pb.LoginRequest))
	}
}

func StartGRPCServer(svc pb.UserManagementServiceServer, port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterUserManagementServiceServer(server, svc)

	log.Printf("gRPC server listening on port %s", port)
	return server.Serve(lis)
}
