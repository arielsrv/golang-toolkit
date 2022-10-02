package stringsextensions

import (
	"errors"
	"github.com/tjarratt/babble"
	"strings"
)

func IsEmpty(value string) bool {
	return value == ""
}

func NotEmpty(value string) error {
	if !IsEmpty(value) {
		return nil
	}

	return errors.New("string value cannot be empty")
}

func RandomString() string {
	babbler := babble.NewBabbler()
	babbler.Separator = "_"
	return strings.ToLower(babbler.Babble())
}
