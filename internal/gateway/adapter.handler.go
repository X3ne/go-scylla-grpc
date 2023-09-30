package gateway

import (
	"context"
	"fmt"
	"log"

	"connectrpc.com/connect"

	adapterv1 "scylla-grpc-adapter/gen/adapter/v1"
	"scylla-grpc-adapter/internal/validators"
)

type AdaperServer struct{}

func (*AdaperServer) Get(ctx context.Context, req *connect.Request[adapterv1.GetRequest]) (*connect.Response[adapterv1.GetResponse], error) {
	if err := ctx.Err(); err != nil {
    return nil, err
  }

	if err := validators.ValidateAdapterRequest(req.Msg); err != nil {
    return nil, connect.NewError(connect.CodeInvalidArgument, err)
  }

	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&adapterv1.GetResponse{
		Message: fmt.Sprintf("Hello, %s!", req.Msg.Value),
	})
	res.Header().Set("Adaper-Version", "v1")
	return res, nil
}
