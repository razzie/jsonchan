package jsonchan

import (
	"io"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

func ReadArray[T any](r io.Reader, path string, c chan<- T) error {
	iter := jsoniter.Parse(jsoniter.ConfigCompatibleWithStandardLibrary, r, 4096)
	if err := progressIterator(iter, path); err != nil {
		return err
	}
	iter.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
		var v T
		iter.ReadVal(&v)
		c <- v
		return true
	})
	return iter.Error
}

func progressIterator(iter *jsoniter.Iterator, path string) error {
	pathSegments := strings.Split(path, ".")
	for _, segment := range pathSegments {
		iter.ReadObjectCB(func(iter *jsoniter.Iterator, s string) bool {
			return s != segment
		})
	}
	return iter.Error
}
