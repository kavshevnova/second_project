package Services

import (
	"errors"
	"second_project/Databases"
	"second_project/Entities"
	"time"
)

type PaymentService struct {
	db *Databases.Database
}

func NewPaymentService(db *Databases.Database) *PaymentService {
	return &PaymentService{db: db}
}

// Обработка платежа
func (s *PaymentService) ProcessPayment(paymentType Entities.PaymentType, fromAccount, toAccount string, amount float32) error {
	// Проверяем, что сумма платежа положительная
	if amount <= 0 {
		return errors.New("сумма платежа должна быть больше нуля")
	}

	// Для пополнения с карты или биткоинов минимальная сумма - 5000 рублей
	if paymentType == Entities.PaymentTypeCardToAccount && amount < 5000 {
		return errors.New("Минимальная сумма пополнения 5000 рублей")
	}

	// Получаем аккаунт получателя
	toAcc, err := s.db.GetAccount(toAccount)
	if err != nil {
		return errors.New("Аккаунт получателя не найден")
	}

	// Если это перевод с аккаунта на аккаунт
	if paymentType == Entities.PaymentTypeAccountToAccount {
		// Получаем аккаунт отправителя
		fromAcc, err := s.db.GetAccount(fromAccount)
		if err != nil {
			return errors.New("Аккаунт отправителя не найден")
		}

		// Проверяем баланс отправителя
		if fromAcc.Balance < amount {
			return errors.New("Недостаточно средств на счете")
		}

		// Обновляем балансы
		fromAcc.Balance -= amount
		toAcc.Balance += amount

		// Сохраняем изменения
		if err := s.db.UpdateAccountBalance(fromAccount, fromAcc.Balance); err != nil {
			return err
		}
	} else {
		// Для пополнения с карты или биткоинов просто увеличиваем баланс получателя
		toAcc.Balance += amount
	}

	// Сохраняем изменения
	if err := s.db.UpdateAccountBalance(toAccount, toAcc.Balance); err != nil {
		return err
	}

	// Сохраняем платеж
	payment := Entities.Payment{
		ID:          len(s.db.Payments) + 1, // Генерация ID
		Type:        paymentType,
		FromAccount: fromAccount,
		ToAccount:   toAccount,
		Amount:      amount,
		Date:        time.Now(),
	}

	return s.db.SavePayment(payment)
}
