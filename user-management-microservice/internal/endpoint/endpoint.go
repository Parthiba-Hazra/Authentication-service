package endpoint

import (
	"context"

	pb "github.com/Parthiba-Hazra/user-management-microservice/api"
)

type Endpoints struct {
	RegisterUserEndpoint func(ctx context.Context, req interface{}) (interface{}, error)
	LoginEndpoint        func(ctx context.Context, req interface{}) (interface{}, error)
}

func MakeEndpoints(svc pb.UserManagementServiceServer) Endpoints {
	return Endpoints{
		RegisterUserEndpoint: MakeRegisterUserEndpoint(svc),
		LoginEndpoint:        MakeLoginEndpoint(svc),
	}
}

func MakeRegisterUserEndpoint(svc pb.UserManagementServiceServer) func(ctx context.Context, req interface{}) (interface{}, error) {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		// Add necessary req/res conversations here.
		return svc.RegisterUser(ctx, req.(*pb.RegisterUserRequest))
	}
}

func MakeLoginEndpoint(svc pb.UserManagementServiceServer) func(ctx context.Context, req interface{}) (interface{}, error) {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		// Add any necessary request/response conversions here.
		return svc.Login(ctx, req.(*pb.LoginRequest))
	}
}
