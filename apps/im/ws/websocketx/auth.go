package websocketx

import "net/http"

type IAuth interface {
	Auth(w http.ResponseWriter, r *http.Request) bool
	GetUser(r *http.Request) string
}

type DefaultAuth struct {
}

func (a *DefaultAuth) Auth(w http.ResponseWriter, r *http.Request) bool {
	return true
}
func (a *DefaultAuth) GetUser(r *http.Request) string {

	return "abc"
}
