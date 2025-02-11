package main

import (
	"context"
	sw "github.com/olgasinepalnikova/avito-task-2025-winter/client/go"
	"log"
	"time"
)

func main() {
	cfg := sw.NewConfiguration()
	cfg.Host = "localhost:8080" // URL, где запущен ваш сервер

	// Создаем экземпляр API клиента
	apiClient := sw.NewAPIClient(cfg)

	// Создаем контекст с таймаутом для запроса
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Пример вызова эндпоинта /api/info
	// Название метода может отличаться в зависимости от генератора и настроек
	infoRequest := apiClient.DefaultAPI.ApiInfoGet(ctx)
	resp, httpResp, err := infoRequest.Execute()
	if err != nil {
		log.Fatalf("Ошибка вызова API: %v\nHTTP ответ: %v", err, httpResp)
	}

	log.Printf("Ответ API: %+v", resp)
}
