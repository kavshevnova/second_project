package main

import (
	"net/http"
	"second_project/Controllers"
	database "second_project/Databases"
	"second_project/Services"
)

func main() {
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
