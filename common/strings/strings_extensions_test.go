package strings_test

import (
	"github.com/arielsrv/golang-toolkit/common/strings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	value := ""
	actual := strings.IsEmpty(value)
	assert.True(t, actual)
}

func TestIsNotEmpty(t *testing.T) {
	value := "hello world!"
	actual := strings.IsEmpty(value)
	assert.False(t, actual)
}

func TestGuardIsEmpty(t *testing.T) {
	value := "hello world!"
	err := strings.NotEmpty(value)
	assert.NoError(t, err)
}

func TestGuardIsNotEmpty(t *testing.T) {
	value := ""
	err := strings.NotEmpty(value)
	assert.Error(t, err)
}
