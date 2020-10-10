package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danielwetan/koala-backend/helpers"
)

func Register(w http.ResponseWriter, r *http.Request) {
	helpers.Headers(&w)

	if r.Method == "POST" {
		r.ParseForm()

		db, err := helpers.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		id := helpers.GenerateId()
		salt := helpers.GenerateId() // salt generate from random string
		email, phoneNumber, customerName := r.FormValue("email"), r.FormValue("phone_number"), r.FormValue("customer_name")
		dob, sex, password := r.FormValue("dob"), r.FormValue("sex"), r.FormValue("password")
		hash, _ := helpers.HashPassword(password)

		_, err = db.Exec(helpers.Query["register"], id, customerName, email, phoneNumber, dob, sex, salt, hash)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		body := "Register success"
		res := helpers.ResponseMsg(true, body)
		json.NewEncoder(w).Encode(res)
	} else {
		body := "Invalid HTTP method"
		res := helpers.ResponseMsg(false, body)
		json.NewEncoder(w).Encode(res)
	}
}
