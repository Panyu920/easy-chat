package xerr

import (
	"github.com/zeromicro/x/errors"
)

func NewError(msg string) error {
	return errors.New(SERVER_COMMON_ERROR, msg)
}

func New(code int, msg string) error {
	return errors.New(code, msg)
}

func NewRequestParamError() error {
	return errors.New(REQUEST_PARAM_ERROR, ErrMsg(REQUEST_PARAM_ERROR))
}

func NewDBError() error {
	return errors.New(DB_ERROR, ErrMsg(DB_ERROR))
}

func NewInternalError() error {
	return errors.New(SERVER_COMMON_ERROR, ErrMsg(SERVER_COMMON_ERROR))
}
