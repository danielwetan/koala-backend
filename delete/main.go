package main

import (
	"fmt"

	"github.com/danielwetan/koala-backend/helpers"
)

func main() {
	res := helpers.GenerateId()
	fmt.Println(res)
}

// func main() {
// 	b := make([]byte, 4) //equals 8 charachters
// 	rand.Read(b)
// 	s := hex.EncodeToString(b)
// 	fmt.Println(s)
// }
