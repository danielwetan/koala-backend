package helpers

var Query = map[string]string{
	"login":    "SELECT * FROM users WHERE email = ?",
	"register": "INSERT INTO Customers (customer_id, customer_name, email, phone_number, dob, sex, salt, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
}
