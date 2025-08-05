package main

import (
	"exception-reporter-agent/server"
	"log"
)

func main() {
	addr := ":9000" // может вынести а вдруг он занят другим приложением

	log.Printf("Starting exception reporter daemon on %s...", addr)

	if err := server.ListenAndServe(addr); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
