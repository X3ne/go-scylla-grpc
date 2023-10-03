package validators

import (
	"fmt"

	usersv1 "scylla-grpc-adapter/gen/users/v1"
)

func ValidateGetUserByIdRequest(req *usersv1.GetByIdRequest) error {
	if req.Id == "" {
		return fmt.Errorf("id is required")
	}
	return nil
}

func ValidateGetUserByUsernameRequest(req *usersv1.GetByUsernameRequest) error {
	if req.Username == "" {
		return fmt.Errorf("username is required")
	}
	return nil
}
