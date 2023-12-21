package jsonchan

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func init() {
	jsoniter.RegisterExtension(&extension{})
}

var chanType = reflect2.TypeOf((*iChan)(nil)).(reflect2.PtrType).Elem()

type extension struct {
	jsoniter.DummyExtension
}

func (extension) CreateEncoder(typ reflect2.Type) jsoniter.ValEncoder {
	if typ.Implements(chanType) {
		return &xcoder{typ: typ}
	}
	return nil
}

func (extension) CreateDecoder(typ reflect2.Type) jsoniter.ValDecoder {
	if typ.Implements(chanType) {
		return &xcoder{typ: typ}
	}
	return nil
}
