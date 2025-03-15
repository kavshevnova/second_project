package Entities

import "time"

type Account struct {
	UserName           string    `yaml:"username"` // Логин
	Password           string    `yaml:"password"` //Пароль
	Balance            float32   `yaml:"balance"`  // Баланс аккаунта
	DateOfRegistration time.Time `yaml:"date"`     // Дата создания аккаунта
}
