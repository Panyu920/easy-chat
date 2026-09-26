package token

import (
	"context"
	"easy-chat/pkg/xerr"
)

var errUserIDNotFound = xerr.New(xerr.REQUEST_PARAM_ERROR, "user id not found")

func GetUserID(c context.Context) (string, error) {
	if id, ok := c.Value(UserID).(string); ok {
		return id, nil
	}
	return "", errUserIDNotFound
}
