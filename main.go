package main

import (
	"net/http"
	"second_project/Controllers"
	database "second_project/Databases"
	"second_project/Server"
	"second_project/Services"
)

func main() {
	Server.StartServer()
	setapPaymentHandler()
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
