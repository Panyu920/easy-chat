package resultx

import (
	"context"
	"easy-chat/pkg/xerr"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	zerr "github.com/zeromicro/x/errors"
	"google.golang.org/grpc/status"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func success(data any) Response {
	return Response{
		Code: 200,
		Msg:  "",
		Data: data,
	}
}

func fail(code int, msg string) Response {
	return Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func OkHandler(_ context.Context, data any) any {
	return success(data)
}

func ErrorHandler(name string) func(ctx context.Context, err error) (int, any) {
	return func(ctx context.Context, err error) (int, any) {
		errcode := xerr.SERVER_COMMON_ERROR
		errmsg := xerr.ErrMsg(errcode)

		causeErr := errors.Cause(err)

		if e, ok := causeErr.(*zerr.CodeMsg); ok {
			errcode = e.Code
			errmsg = e.Msg
		} else {
			if gstatus, ok := status.FromError(causeErr); ok {
				errcode = int(gstatus.Code())
				errmsg = gstatus.Message()
			}
		}
		logx.WithContext(ctx).Errorf("[%s] err %v", name, err)

		return http.StatusBadRequest, fail(errcode, errmsg)
	}
}
