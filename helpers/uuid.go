package helpers

import (
	"github.com/lithammer/shortuuid"
)

func GenerateId() string {
	id := shortuuid.New()
	return id[:6] // return only 6 character
}
