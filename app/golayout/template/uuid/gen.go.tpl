package uuid

import uuid "github.com/satori/go.uuid"

// Gen generate new uuid
func Gen() string {
	return uuid.NewV4().String()
}
