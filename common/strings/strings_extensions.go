package strings

import "errors"

func IsEmpty(value string) bool {
	return value == ""
}

func NotEmpty(value string) error {
	if !IsEmpty(value) {
		return nil
	}

	return errors.New("string value cannot be empty")
}
