package routes

import (
	"net/http"
)

func Auth() {
	http.HandleFunc("/auth/register", controllers.Register)
}
