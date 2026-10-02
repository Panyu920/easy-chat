package handler

import (
	"context"
	"easy-chat/apps/im/ws/internal/svc"
	"net/http"

	mtk "easy-chat/pkg/token"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/token"
)

type JWTAuth struct {
	svc    *svc.ServiceContext
	parser *token.TokenParser
	logx   logx.Logger
}

func NewJWTAuth(svc *svc.ServiceContext) *JWTAuth {
	return &JWTAuth{
		svc:    svc,
		logx:   logx.WithContext(context.Background()),
		parser: token.NewTokenParser(),
	}
}

func (a *JWTAuth) GetUserID(r *http.Request) string {
	userID, err := mtk.GetUserID(r.Context())
	if err != nil {
		return ""
	}
	return userID

}

func (a *JWTAuth) Auth(w http.ResponseWriter, r *http.Request) bool {
	tk, err := a.parser.ParseToken(r, a.svc.Config.JwtAuth.AccessSecret, "")
	if err != nil {
		a.logx.Errorf("ParseToken error: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return false
	}

	if !tk.Valid {
		a.logx.Errorf("Token is not valid")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return false
	}

	claims, ok := tk.Claims.(jwt.MapClaims)
	if !ok {
		a.logx.Errorf("Claims are not of type MapClaims")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return false
	}
	*r = *r.WithContext(context.WithValue(r.Context(), mtk.UserID, claims[mtk.UserID]))
	return true
}
