package gateway

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	adapterv1 "scylla-grpc-adapter/gen/adapter/v1"
	"scylla-grpc-adapter/internal/validators"
)

type AdaperServer struct{}

func (*AdaperServer) Post(ctx context.Context, req *connect.Request[adapterv1.PostRequest]) (*connect.Response[adapterv1.PostResponse], error) {
	if err := ctx.Err(); err != nil {
    return nil, err
  }

	if err := validators.ValidateAdapterRequest(req.Msg); err != nil {
    return nil, connect.NewError(connect.CodeInvalidArgument, err)
  }

	res := connect.NewResponse(&adapterv1.PostResponse{
		Message: fmt.Sprintf("Hello %s", req.Msg.Value),
	})

	res.Header().Set("Adaper-Version", "v1")

	return res, nil
}
