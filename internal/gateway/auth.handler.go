package gateway

import (
	"context"
	"errors"

	authv1 "github.com/X3ne/go-scylla-grpc/gen/auth/v1"
	api_errors "github.com/X3ne/go-scylla-grpc/internal/errors"
	"github.com/X3ne/go-scylla-grpc/internal/validators"
	"github.com/X3ne/go-scylla-grpc/services"

	"connectrpc.com/connect"
)

type AuthServer struct{
	JwtManager *services.JwtManager
}

func (*AuthServer) Register(ctx context.Context, req *connect.Request[authv1.PostRequest]) (*connect.Response[authv1.SuccessResponse], error) {
	if err := ctx.Err(); err != nil {
    return nil, err
  }

	if err := validators.ValidateAuthRequest(req.Msg); err != nil {
    return nil, connect.NewError(connect.CodeInvalidArgument, err)
  }

	err := services.CreateUser(&services.User{
		Username: req.Msg.Username,
		Password: req.Msg.Password,
	})
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&authv1.SuccessResponse{
		Message: "User created successfully",
	})

	res.Header().Set("Auth-Version", "v1")

	return res, nil
}

func (s *AuthServer) Login(ctx context.Context, req *connect.Request[authv1.PostRequest]) (*connect.Response[authv1.LoginResponse], error) {
	if err := ctx.Err(); err != nil {
    return nil, err
  }

	if err := validators.ValidateAuthRequest(req.Msg); err != nil {
    return nil, connect.NewError(connect.CodeInvalidArgument, err)
  }

	user, err := services.GetUserByUsername(req.Msg.Username)
	if err != nil {
		return nil, err
	}

	match, err := services.ComparePasswordAndHash(req.Msg.Password, user.Password)
	if err != nil {
		return nil, err
	}

	if !match {
		return nil, errors.New(api_errors.ErrInvalidPassword)
	}

	token, err := s.JwtManager.Generate(user)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&authv1.LoginResponse{
		Username: user.Username,
		Token: 		token,
	})

	res.Header().Set("Auth-Version", "v1")

	return res, nil
}
