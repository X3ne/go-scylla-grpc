package gateway

import (
	"context"
	"fmt"
	"log"

	"connectrpc.com/connect"
	"github.com/bwmarrin/snowflake"
	"github.com/scylladb/gocqlx/v2/table"

	adapterv1 "scylla-grpc-adapter/gen/adapter/v1"
	"scylla-grpc-adapter/internal/app"
	"scylla-grpc-adapter/internal/validators"
)

type AdaperServer struct{
	App *app.App
}

var usersMetadata = table.Metadata{
	Name:    "users",
	Columns: []string{"id", "value"},
	PartKey: []string{"id"},
	SortKey: []string{"value"},
}

type User struct {
	ID    string `db:"id"`
	Value string `db:"value"`
}

func (s *AdaperServer) Post(ctx context.Context, req *connect.Request[adapterv1.PostRequest]) (*connect.Response[adapterv1.PostResponse], error) {
	log.Printf("Received: %v", req.Msg)
	if err := ctx.Err(); err != nil {
    return nil, err
  }

	if err := validators.ValidateAdapterRequest(req.Msg); err != nil {
    return nil, connect.NewError(connect.CodeInvalidArgument, err)
  }

	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	p := User{
		ID:    node.Generate().String(),
		Value: req.Msg.Value,
	}

	personTable := table.New(usersMetadata)

	// insert user into db if not exists
	if err := s.App.Db.Session.Query(personTable.Insert()).BindStruct(p).ExecRelease(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// get user by id
	var user User
	if err := s.App.Db.Session.Query(personTable.Get()).BindStruct(p).GetRelease(&user); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Printf("User: %v", user)

	res := connect.NewResponse(&adapterv1.PostResponse{
		Message: fmt.Sprintf("Hello %s", user),
	})

	res.Header().Set("Adaper-Version", "v1")

	return res, nil
}
