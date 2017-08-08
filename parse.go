package libtmx

import (
	"github.com/voidshard/libtmx/common"
	"github.com/voidshard/libtmx/codecs/v1"
)

type TmxCodec interface {
	Unmarshal([]byte) (*common.Map, error)
	Marshal(*common.Map) ([]byte, error)
}

func CodecV1() TmxCodec {
	return &v1.CodecV1{}
}
