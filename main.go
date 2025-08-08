package main

import (
	"fmt"

	"github.com/4Noyis/ecommerce-product-api/pkg/database"
)

func main() {
	client, err := database.ConnectMongoDB("ecommerce-api")
	if err != nil {
		fmt.Println("hata")
	}
	defer client.Close()
}
