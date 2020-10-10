package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danielwetan/koala-backend/helpers"
)

func Orders(w http.ResponseWriter, r *http.Request) {
	helpers.Headers(&w)

	if r.Method == "POST" {
		r.ParseForm()

		db, err := helpers.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		orderID := helpers.GenerateId()
		orderNumber := 1
		customerID, paymentMethodID := r.FormValue("customer_id"), r.FormValue("payment_method_id")

		_, err = db.Exec(helpers.Query["orders"], orderID, customerID, orderNumber, paymentMethodID)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		body := map[string]string{
			"message": "Order success!",
			"orderID": orderID,
		}
		res := helpers.ResponseMsg(true, body)
		json.NewEncoder(w).Encode(res)
	} else {
		body := "Invalid HTTP method"
		res := helpers.ResponseMsg(false, body)
		json.NewEncoder(w).Encode(res)
	}
}
