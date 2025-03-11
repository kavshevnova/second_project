package Controllers

import (
	"encoding/json"
	"net/http"
	"second_project/Entities"
	"second_project/Services"
)

type PaymentController struct {
	service *Services.PaymentService
}

func NewPaymentController(service *Services.PaymentService) *PaymentController {
	return &PaymentController{service: service}
}

// Jбработчик для выполнения платежа
func (c *PaymentController) ProcessPaymentHandler(w http.ResponseWriter, r *http.Request) {
	var paymentRequest struct {
		Type        Entities.PaymentType `json:"type"`
		FromAccount string               `json:"from_account"`
		ToAccount   string               `json:"to_account"`
		Amount      float32              `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&paymentRequest); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	if err := c.service.ProcessPayment(paymentRequest.Type, paymentRequest.FromAccount, paymentRequest.ToAccount, paymentRequest.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Платеж успешно выполнен"))
}
