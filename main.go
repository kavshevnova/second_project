package main

import (
	"fmt"
	"log"
	"net/http"
	"second_project/Controllers"
	database "second_project/Databases"
	"second_project/Services"
	"time"
)

func main() {
	acc := &account.Account{
		UserName:           "john_doe",
		Balance:            100.50,
		DateOfRegistration: timestamppb.New(time.Now()),
	}

	data, err := proto.Marshal(acc)
	if err != nil {
		log.Fatalf("Serialization error: %v", err)
	}

	newAcc := &account.Account{}
	err = proto.Unmarshal(data, newAcc)
	if err != nil {
		log.Fatalf("Deserialization error: %v", err)
	}
	fmt.Printf("Original: %+v\n", acc)
	fmt.Printf("New: %+v\n", newAcc)

	setapAccountHandler()
	setapPaymentHandler()
}
func setapAccountHandler() {
	acc := database.NewDatabase_accounts()
	AccountService := Services.NewAccountService(acc)
	AccountController := Controllers.NewAccountController(AccountService)
	http.HandleFunc("/account/create", AccountController.CreateAccountHandler)
	http.HandleFunc("/account/delete", AccountController.DeleteClientHandler)
	http.HandleFunc("/account/get", AccountController.GetClientHandler)
}
func setapPaymentHandler() {
	db := database.Newdatabase()
	PaymentService := Services.NewPaymentService(db)
	PaymentController := Controllers.NewPaymentController(PaymentService)
	// Регистрируем обработчики
	http.HandleFunc("/payment/process", PaymentController.ProcessPaymentHandler)
	// Запускаем сервер
	http.ListenAndServe(":8081", nil)
}
