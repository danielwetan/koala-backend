package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danielwetan/koala-backend/helpers"
	jwt "github.com/dgrijalva/jwt-go"
)

func Authorization(w http.ResponseWriter, r *http.Request) {
	tokenString := w.Header().Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		body := "Authorization failed!"
		res := helpers.ResponseMsg(false, body)
		json.NewEncoder(w).Encode(res)
	}

}
