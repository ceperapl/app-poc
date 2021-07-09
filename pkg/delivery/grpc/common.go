package grpc

import "github.com/gofrs/uuid"

func isValidUUID(u string) bool {
	_, err := uuid.FromString(u)
	return err == nil
}
