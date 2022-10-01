package common_test

import (
	"github.com/arielsrv/golang-toolkit/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	value := ""
	actual := common.IsEmpty(value)
	assert.True(t, actual)
}

func TestIsNotEmpty(t *testing.T) {
	value := "hello world!"
	actual := common.IsEmpty(value)
	assert.False(t, actual)
}

func TestGuardIsEmpty(t *testing.T) {
	value := "hello world!"
	err := common.NotEmpty(value)
	assert.NoError(t, err)
}

func TestGuardIsNotEmpty(t *testing.T) {
	value := ""
	err := common.NotEmpty(value)
	assert.Error(t, err)
}
