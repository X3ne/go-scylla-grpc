package validators

import (
	"fmt"
	"log"

	adapterv1 "scylla-grpc-adapter/gen/adapter/v1"
)

func ValidateAdapterRequest(req *adapterv1.GetRequest) error {
	log.Println("Validating request: ", req)
	if req.Value == "" {
		return fmt.Errorf("value is required")
	}
	return nil
}
