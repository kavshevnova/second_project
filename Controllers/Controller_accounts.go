package Controllers

import (
	"encoding/json"
	"net/http"
	"second_project/Entities"
	"second_project/Services"
)

type AccountController struct {
	service *Services.AccountService
}

func NewAccountController(service *Services.AccountService) *AccountController {
	return &AccountController{service: service}
}

// обработчик для создания
func (c *AccountController) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var account Entities.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.service.CreateAccount(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Пользователь создан"))
}

// обработчик для удаления
func (c *AccountController) DeleteClientHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Пользователь не найден", http.StatusBadRequest)
		return
	}

	if err := c.service.DeleteAccount(username); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Пользователь удален"))
}

// обработчик для получения
func (c *AccountController) GetClientHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Пользователь не найден", http.StatusBadRequest)
		return
	}

	account, err := c.service.GetAccount(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}
