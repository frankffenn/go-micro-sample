package errors

import "golang.org/x/xerrors"

type ErrorCode uint64

const (
	UnknownErr ErrorCode = 1000 + iota
)

var Errors = map[ErrorCode]error{
	UnknownErr: xerrors.New("unknown error"),
}

func (e *ErrorCode) Message() string {
	if msg, ok := Errors[*e]; ok {
		return msg.Error()
	}
	return Errors[UnknownErr].Error()
}
