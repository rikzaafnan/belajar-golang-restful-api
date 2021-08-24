package middleware

import "net/http"

type AuthMidleware struct {
	Handler http.Handler
}

func NewAuthMidleware(handler http.Handler) *AuthMidleware {
	return &AuthMidleware{Handler: handler}
}

func (midleware *AuthMidleware) ServerHTP(writer http.ResponseWriter, request *http.Request) {

}
