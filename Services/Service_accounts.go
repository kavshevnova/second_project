package Services

import (
	"second_project/Databases"
	"second_project/Entities"
)

type AccountService struct {
	acc *Databases.Database_accounts
}

func NewAccountService(acc *Databases.Database_accounts) *AccountService {
	return &AccountService{acc: acc}
}

// создание
func (s *AccountService) CreateAccount(account Entities.Account) error {
	return s.acc.SaveAccount(account)
}

// удаление
func (s *AccountService) DeleteAccount(username string) error {
	return s.acc.DeleteAccount(username)
}

// получение
func (s *AccountService) GetAccount(username string) (Entities.Account, error) {
	return s.acc.GetAccount(username)
}
