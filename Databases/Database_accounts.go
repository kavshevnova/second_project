package Databases

import (
	"errors"
	"second_project/Entities"
	"sync"
)

type Database_accounts struct {
	mu       sync.RWMutex
	accounts map[string]Entities.Account
}

func NewDatabase_accounts() *Database_accounts {
	return &Database_accounts{
		accounts: make(map[string]Entities.Account)}
}

// сохранение пользователя
func (acc *Database_accounts) SaveAccount(account Entities.Account) error {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	acc.accounts[account.UserName] = account
	return nil
}

// удаление пользователя по имени
func (acc *Database_accounts) DeleteAccount(username string) error {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	if _, exists := acc.accounts[username]; !exists {
		return ErrAccountNotFound
	}
	delete(acc.accounts, username)
	return nil
}

// получение пользователя по имени
func (acc *Database_accounts) GetAccount(username string) (Entities.Account, error) {
	acc.mu.RLock()
	defer acc.mu.RUnlock()

	account, exists := acc.accounts[username]
	if !exists {
		return Entities.Account{}, ErrAccountNotFound
	}

	return account, nil
}

// Ошибка, если анкета не найдена
var ErrAccountNotFound = errors.New("Пользователь не найден")
