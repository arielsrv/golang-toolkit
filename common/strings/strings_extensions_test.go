package stringsextensions_test

import (
	. "github.com/arielsrv/golang-toolkit/common/strings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	value := ""
	actual := IsEmpty(value)
	assert.True(t, actual)
}

func TestIsNotEmpty(t *testing.T) {
	value := "hello world!"
	actual := IsEmpty(value)
	assert.False(t, actual)
}

func TestGuardIsEmpty(t *testing.T) {
	value := "hello world!"
	err := NotEmpty(value)
	assert.NoError(t, err)
}

func TestGuardIsNotEmpty(t *testing.T) {
	value := ""
	err := NotEmpty(value)
	assert.Error(t, err)
}
