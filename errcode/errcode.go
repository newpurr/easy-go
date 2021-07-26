package errcode

import (
	"fmt"
)

type Error struct {
	// 错误码
	code int
	// 错误消息
	msg     string
	details []map[string]string
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg, details: make([]map[string]string, 0)}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息:：%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) WithDetail(key, val string) *Error {
	e.details = append(e.details, map[string]string{
		"key": key,
		"val": val,
	})
	return e
}

func (e *Error) Details() []map[string]string {
	return e.details
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}
