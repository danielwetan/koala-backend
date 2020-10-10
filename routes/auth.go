package routes

import (
	"net/http"

	"github.com/danielwetan/koala-backend/controllers"
)

func Auth() {
	http.HandleFunc("/api/auth/register", controllers.Register)
}
