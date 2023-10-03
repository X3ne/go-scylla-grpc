package interceptors

import (
	"context"
	"errors"
	"log"

	api_errors "github.com/X3ne/go-scylla-grpc/internal/errors"
	"github.com/X3ne/go-scylla-grpc/services"

	"connectrpc.com/connect"
)

const tokenHeader = "Authorization"

func NewAuthInterceptor(manager *services.JwtManager) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if err := ctx.Err(); err != nil {
				return nil, err
			}

			token := req.Header().Get(tokenHeader)
			if token == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New(api_errors.ErrUnauthorized))
			}

			token = token[len("Bearer "):]

			log.Println(token)

			claims, err := manager.Verify(token)
			log.Println(claims)
			if err != nil {
				log.Println(err)
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New(api_errors.ErrUnauthorized))
			}

			ctx = context.WithValue(ctx, "username", claims.Username)

      return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
