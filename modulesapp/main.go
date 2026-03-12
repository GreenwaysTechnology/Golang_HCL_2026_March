package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println("Hello World")
	uuid := uuid.New()
	fmt.Println(uuid)
	color.Cyan("Hello World")
	color.Blue(uuid.String())
	company := os.Getenv("COMPANY_NAME")
	fmt.Println(company)
}
