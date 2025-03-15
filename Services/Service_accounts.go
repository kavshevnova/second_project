package services

import (
	"context"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"second_project/Databases"
	"second_project/entities"
	"second_project/gen/go/account"
)

type AccountService struct {
	account.UnimplementedAccountServiceServer
	db *Databases.Database_accounts
}

func NewAccountService(db *Databases.Database_accounts) *AccountService {
	return &AccountService{db: db}
}

// hashPassword хэширует пароль с использованием bcrypt
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *AccountService) CreateAccount(ctx context.Context, req *account.CreateAccountRequest) (*account.CreateAccountResponse, error) {
	// Проверяем, что имя пользователя и пароль переданы
	if req.GetUsername() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Username is required")
	}
	if req.GetPassword() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Password is required")
	}

	// Хэшируем пароль
	hashedPassword, err := hashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to hash password: %v", err)
	}

	// Создаем сущность Account
	entity := Entities.Account{
		UserName:           req.GetUsername(), // Имя пользователя из запроса
		Password:           hashedPassword,    // Хэшированный пароль
		Balance:            0,                 // Баланс по умолчанию
		DateOfRegistration: time.Now(),        // Текущее время
	}

	// Сохраняем аккаунт в базу данных
	if err := s.db.SaveAccount(entity); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to save account: %v", err)
	}

	// Преобразуем сущность обратно в gRPC-сообщение для ответа
	acc := &account.Account{
		Username:           entity.UserName,
		Password:           entity.Password,                            // Хэшированный пароль
		Balance:            entity.Balance,                             // Баланс будет 0
		DateOfRegistration: timestamppb.New(entity.DateOfRegistration), // Текущее время
	}

	// Возвращаем ответ
	return &account.CreateAccountResponse{
		Account: acc,
	}, nil
}

func (s *AccountService) GetAccount(ctx context.Context, req *account.GetAccountRequest) (*account.GetAccountResponse, error) {
	username := req.GetUsername()
	if username == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Username is required")
	}

	// Получаем аккаунт из базы данных
	entity, err := s.db.GetAccount(username)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Account not found: %v", err)
	}

	// Преобразуем сущность в gRPC-сообщение
	acc := &account.Account{
		Username:           entity.UserName,
		Password:           entity.Password, // Хэшированный пароль
		Balance:            entity.Balance,
		DateOfRegistration: timestamppb.New(entity.DateOfRegistration),
	}

	// Возвращаем ответ
	return &account.GetAccountResponse{
		Account: acc,
	}, nil
}

func (s *AccountService) DeleteAccount(ctx context.Context, req *account.DeleteAccountRequest) (*account.DeleteAccountResponse, error) {
	username := req.GetUsername()
	if username == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Username is required")
	}

	// Удаляем аккаунт из базы данных
	if err := s.db.DeleteAccount(username); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete account: %v", err)
	}

	// Возвращаем ответ
	return &account.DeleteAccountResponse{
		Success: true,
	}, nil
}
