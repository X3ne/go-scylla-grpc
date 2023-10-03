package validators

import (
	"fmt"
	authv1 "scylla-grpc-adapter/gen/auth/v1"
)

func ValidateAuthRequest(req *authv1.PostRequest) error {
	if req.Username == "" {
		return fmt.Errorf("username is required")
	}
	if req.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}
