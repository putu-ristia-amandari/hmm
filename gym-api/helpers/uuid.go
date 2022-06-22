package helpers

import (
	"strings"

	"github.com/google/uuid"
)

func UUID() (string, error) {
	uuidGenerator := uuid.New()
	result := strings.Replace(uuidGenerator.String(), "-", "", -1)

	return result, nil
}
