package routes

import (
	"net/http"

	"github.com/danielwetan/koala-backend/controllers"
)

func Orders() {
	http.HandleFunc("/api/orders", controllers.Orders)
}
