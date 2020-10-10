package helpers

var Query = map[string]string{
	"login":    "SELECT * FROM Customers WHERE email = ?",
	"register": "INSERT INTO Customers (customer_id, customer_name, email, phone_number, dob, sex, salt, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
	"orders":   "INSERT INTO Orders (order_id, customer_id, order_number, payment_method_id) VALUE (?, ?, ?, ?)",
}
