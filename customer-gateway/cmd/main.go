package main

import (
	"context"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/pkg/customer/servicespb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка при подключении: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// Создание клиентского объекта
	client := servicespb.NewCustomerServiceClient(conn)

	// Напиши код, чтобы он слал запросы раз в 5 минут
	// Вызов метода сервера
	for i := 0; i < 100; i++ {
		request := &servicespb.CreateCustomerRequest{
			Email: "Привет, сервер!",
		}

		response, err := client.CreateCustomer(context.Background(), request)
		if err != nil {
			log.Fatalf("Ошибка при вызове метода: %v", err)
		}

		// Обработка ответа от сервера
		fmt.Printf("Ответ сервера: %s\n", response.Id)
		time.Sleep(5 * time.Second)
	}
}
