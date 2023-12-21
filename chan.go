package jsonchan

import (
	jsoniter "github.com/json-iterator/go"
)

type iChan interface {
	encode(stream *jsoniter.Stream)
	decode(iter *jsoniter.Iterator)
}

type Chan[T any] chan T

func (c Chan[T]) encode(stream *jsoniter.Stream) {
	stream.WriteArrayStart()
	count := 0
	for v := range c {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteVal(v)
		count++
	}
	stream.WriteArrayEnd()
}

func (c Chan[T]) decode(iter *jsoniter.Iterator) {
	iter.ReadArrayCB(func(i *jsoniter.Iterator) bool {
		var v T
		iter.ReadVal(&v)
		c <- v
		return true
	})
}
