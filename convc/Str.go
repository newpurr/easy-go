package convc

import (
	"github.com/spf13/cast"
	"strings"
)

type Str string

func (s Str) String() string {
	return string(s)
}

func (s Str) Int() (int, error) {
	return cast.ToIntE(s.String())
}

func (s Str) MustInt() int {
	return cast.ToInt(s.String())
}

func (s Str) UInt32() (uint32, error) {
	return cast.ToUint32E(s.String())
}

func (s Str) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}

func (s Str) MustSplitSlice(sep string) []string {
	return strings.Split(s.String(), sep)
}

func (s Str) SplitIntSlice(sep string) ([]int, error) {
	return cast.ToIntSliceE(s.MustSplitSlice(sep))
}
