package jsonchan

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

type xcoder struct {
	typ reflect2.Type
}

func (x xcoder) IsEmpty(ptr unsafe.Pointer) bool {
	return x.typ.UnsafeIsNil(ptr)
}

func (x xcoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	c := x.typ.UnsafeIndirect(ptr).(iChan)
	c.encode(stream)
}

func (x xcoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	c := x.typ.UnsafeIndirect(ptr).(iChan)
	c.decode(iter)
}
