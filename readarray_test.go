package jsonchan

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadArray(t *testing.T) {
	r := strings.NewReader(`{"Obj":{"Ints":[1,2,3]}}`)
	c := make(chan int, 3)

	err := ReadArray(r, "Obj.Ints", c)
	assert.NoError(t, err)
	assert.Equal(t, 1, <-c)
	assert.Equal(t, 2, <-c)
	assert.Equal(t, 3, <-c)
}
