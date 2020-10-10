package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/danielwetan/koala-backend/helpers"
	"github.com/danielwetan/koala-backend/models"
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

func Login(w http.ResponseWriter, r *http.Request) {
	helpers.Headers(&w)

	if r.Method == "POST" {
		r.ParseForm()

		db, err := helpers.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		result := models.Customers{}
		email, password := r.FormValue("email"), r.FormValue("password")
		err = db.
			QueryRow(helpers.Query["login"], email).
			Scan(&result.CustomerID, &result.CustomerName, &result.Email, &result.PhoneNumber, &result.Dob, &result.Sex, &result.Salt, &result.Password, &result.CreatedAt)

		match := helpers.CheckPasswordHash(password, result.Password)
		if match {

			const SECRET = "quatre vingt dix neuf heures"
			sign := jwt.New(jwt.GetSigningMethod("HS256"))
			token, err := sign.SignedString([]byte(SECRET))
			if err != nil {
				body := "Login failed!"
				res := helpers.ResponseMsg(false, body)
				json.NewEncoder(w).Encode(res)
				return
			}

			const REFRESH = "bonjour comment ca va"
			refreshSign := jwt.New(jwt.GetSigningMethod("HS256"))
			refreshToken, err := refreshSign.SignedString([]byte(REFRESH))
			if err != nil {
				body := "Login failed!"
				res := helpers.ResponseMsg(false, body)
				json.NewEncoder(w).Encode(res)
				return
			}

			body := map[string]string{
				"accessToken":  token,
				"refreshToken": refreshToken,
			}

			// return JWT token
			res := helpers.ResponseMsg(true, body)
			json.NewEncoder(w).Encode(res)
		} else {
			body := "Username or password is wrong!"
			res := helpers.ResponseMsg(false, body)
			json.NewEncoder(w).Encode(res)
		}
	} else {
		body := "Invalid HTTP method"
		res := helpers.ResponseMsg(false, body)
		json.NewEncoder(w).Encode(res)
	}
}
