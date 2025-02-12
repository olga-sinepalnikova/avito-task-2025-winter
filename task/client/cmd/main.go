package main

import (
	"context"
	sw "github.com/olgasinepalnikova/avito-task-2025-winter/client/go"
	"log"
	"time"
)

func main() {
	cfg := sw.NewConfiguration()
	// URL, если запускаем в докере, то server, если на машине, то localhost
	cfg.Host = "server:8080"

	apiClient := sw.NewAPIClient(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	infoRequest := apiClient.DefaultAPI.ApiInfoGet(ctx)
	resp, httpResp, err := infoRequest.Execute()
	if err != nil {
		log.Fatalf("Ошибка вызова API: %v\nHTTP ответ: %v", err, httpResp)
	}

	log.Printf("Ответ API: %+v", resp)
}
