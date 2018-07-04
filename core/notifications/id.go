package notifications

import (
	"strings"

	"github.com/pborman/uuid"
)

// ID uniquely identifies a particular entity.
type ID string

// NextID generates a new ID.
func NextID() ID {
	return ID(strings.Split(strings.ToUpper(uuid.New()), "-")[0])
}
