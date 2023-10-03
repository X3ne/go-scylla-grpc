package gateway

import (
	"context"

	usersv1 "github.com/X3ne/go-scylla-grpc/gen/users/v1"
	"github.com/X3ne/go-scylla-grpc/internal/validators"
	"github.com/X3ne/go-scylla-grpc/services"

	"connectrpc.com/connect"
)

type UsersServer struct{}

func (*UsersServer) GetById(ctx context.Context, req *connect.Request[usersv1.GetByIdRequest]) (*connect.Response[usersv1.GetByIdResponse], error) {
	if err := ctx.Err(); err != nil {
    return nil, err
  }

	if err := validators.ValidateGetUserByIdRequest(req.Msg); err != nil {
    return nil, connect.NewError(connect.CodeInvalidArgument, err)
  }

	user, err := services.GetUserById(req.Msg.Id)
	if err != nil {
		if err.Error() == "not found" {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, err
	}

	res := connect.NewResponse(&usersv1.GetByIdResponse{
		User: &usersv1.GetResponse{
			Id:       user.Id,
			Username: user.Username,
		},
	})

	res.Header().Set("Users-Version", "v1")

	return res, nil
}

func (*UsersServer) GetByUsername(ctx context.Context, req *connect.Request[usersv1.GetByUsernameRequest]) (*connect.Response[usersv1.GetByUsernameResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validators.ValidateGetUserByUsernameRequest(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	user, err := services.GetUserByUsername(req.Msg.Username)
	if err != nil {
		if err.Error() == "not found" {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, err
	}

	res := connect.NewResponse(&usersv1.GetByUsernameResponse{
		User: &usersv1.GetResponse{
			Id:       user.Id,
			Username: user.Username,
		},
	})

	res.Header().Set("Users-Version", "v1")

	return res, nil
}

func (*UsersServer) GetAll(ctx context.Context, req *connect.Request[usersv1.GetAllRequest]) (*connect.Response[usersv1.GetAllResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	users, err := services.GetUsers()
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&usersv1.GetAllResponse{
		Users: make([]*usersv1.GetResponse, len(users)),
	})

	for i, user := range users {
		res.Msg.Users[i] = &usersv1.GetResponse{
			Id:       user.Id,
			Username: user.Username,
		}
	}

	res.Header().Set("Users-Version", "v1")

	return res, nil
}
