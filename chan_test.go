package jsonchan

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestChanEncode(t *testing.T) {
	o := struct {
		Ints Chan[int]
	}{
		Ints: make(Chan[int], 3),
	}
	o.Ints <- 1
	o.Ints <- 2
	o.Ints <- 3
	close(o.Ints)

	json, err := jsoniter.MarshalToString(o)
	assert.NoError(t, err)
	assert.Equal(t, `{"Ints":[1,2,3]}`, json)
}

func TestChanDecode(t *testing.T) {
	json := `{"Ints":[1,2,3]}`
	o := struct {
		Ints Chan[int]
	}{
		Ints: make(Chan[int], 3),
	}

	err := jsoniter.UnmarshalFromString(json, &o)
	assert.NoError(t, err)
	assert.Equal(t, 1, <-o.Ints)
	assert.Equal(t, 2, <-o.Ints)
	assert.Equal(t, 3, <-o.Ints)
}
