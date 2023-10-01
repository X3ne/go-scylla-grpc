package validators

import (
	"fmt"

	adapterv1 "scylla-grpc-adapter/gen/adapter/v1"
)

func ValidateAdapterRequest(req *adapterv1.PostRequest) error {
	if req.Value == "" {
		return fmt.Errorf("value is required")
	}
	return nil
}
