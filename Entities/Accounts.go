package Entities

import "time"

type Account struct {
	UserName           string    `json:"user_id"` // Логин
	Balance            float32   `json:"balance"` // Баланс аккаунта
	DateOfRegistration time.Time `json:"date"`    // Дата создания аккаунта
}
