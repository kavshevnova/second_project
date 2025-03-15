package Server

import (
	"google.golang.org/grpc"
	"log"
	"net"
	database "second_project/Databases"
	services "second_project/Services"
	"second_project/gen/go/account"
)

func StartServer() {
	acc := database.NewDatabase_accounts()
	grpcServer := grpc.NewServer()
	// Регистрируем сервис
	accountService := services.NewAccountService(acc)
	account.RegisterAccountServiceServer(grpcServer, accountService)
	// Запускаем сервер
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
