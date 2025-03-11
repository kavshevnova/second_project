package Entities

import "time"

type PaymentType string

const (
	PaymentTypeAccountToAccount PaymentType = "account_to_account" // Перевод с аккаунта на аккаунт
	PaymentTypeCardToAccount    PaymentType = "card_to_account"    // Пополнение с карты или биткоинов
)

type Payment struct {
	ID          int         `json:"id"` //id платежа
	Type        PaymentType `json:"type"`
	FromAccount string      `json:"from_account"` // Имя аккаунта или способа оплаты
	ToAccount   string      `json:"to_account"`   // Номер карты или аккаунта
	Amount      float32     `json:"amount"`       // Сумма платежа
	Date        time.Time   `json:"date"`         // Дата платежа
}
