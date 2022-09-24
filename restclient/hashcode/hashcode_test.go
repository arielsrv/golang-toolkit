package hashcode_test

import (
	"github.com/arielsrv/golang-toolkit/restclient/hashcode"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	value := hashcode.String("hello world!")
	assert.Equal(t, 62177901, value)
}
