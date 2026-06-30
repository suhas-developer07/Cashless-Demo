package main

import (
	"fmt"
	"net/http"

	"github.com/suhas-developer07/cashless-demo/internals/db"
	"github.com/suhas-developer07/cashless-demo/internals/routes"
)

// features
// 1. card need to initialize
// 2. card recharge using payment gateway
// 3. POS Machine installation
// 4. POS Machine Maintanance
// 5. Card Payments Using POS Machine
// 6. POS Machine Revenue Count for 24h  {POS acts as a vendor}
// 7. Card Balance Minimum 50rs if not cannot able to make payment
// 8. Card Transaction
// 9. POS Machine Transaction

func main(){
	fmt.Print("Hello World")
	db,err := db.Connect()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()
	
	fmt.Println("Connected to the database successfully!")
	
	server := &http.Server{
		Addr:    ":8080",
		Handler: routes.SetupRoutes(),
	}

	fmt.Println("Server is running on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}