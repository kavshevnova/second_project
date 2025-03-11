package Databases

import (
	"errors"
	"second_project/Entities"
	"sync"
)

type Database struct {
	mu       sync.RWMutex
	payments map[int]Entities.Payment
	accounts map[string]Entities.Account
}

func Newdatabase() *Database {
	return &Database{
		payments: make(map[int]Entities.Payment),
		accounts: make(map[string]Entities.Account),
	}
}

// Сохранение платежа
func (db *Database) SavePayment(payment Entities.Payment) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.payments[payment.ID] = payment
	return nil
}

// Получение аккаунта по логину
func (db *Database) GetAccount(UserName string) (Entities.Account, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	account, exists := db.accounts[UserName]
	if !exists {
		return Entities.Account{}, errors.New("Аккаунт не найден")
	}
	return account, nil
}

// Jбновление баланса аккаунта
func (db *Database) UpdateAccountBalance(UserName string, newBalance float32) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	account, exists := db.accounts[UserName]
	if !exists {
		return errors.New("Аккаунт не найден")
	}

	account.Balance = newBalance
	db.accounts[UserName] = account
	return nil
}
