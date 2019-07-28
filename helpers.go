package gohelpers

import (
	"github.com/brightappsllc/gohelpers/strings"
	"github.com/google/uuid"
)

// NewGUID -
func NewGUID() string {
	guidAsString := strings.RandomString(50)
	id, err := uuid.NewUUID()
	if err == nil {
		guidAsString = id.String()
	}

	return guidAsString
}
