package main

import (
	"Go-Calculator/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/calculate", handlers.CalculateHandler)

	addr := ":8080"
	log.Printf("Сервер запущен на %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
