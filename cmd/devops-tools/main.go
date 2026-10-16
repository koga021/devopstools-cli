package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
	log.Printf("DevOpsToolsTarget: %s", os.Getenv("DEVOPS_TOOLS_TARGET"))

	FirstMenu()
}
