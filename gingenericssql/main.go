package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gingenericssql/handler"
	"github.com/gingenericssql/repository"
	"github.com/gingenericssql/service"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := "postgres://admin:secret@localhost:5432/mydb?sslmode=disable"
	return sql.Open("postgres", connStr)
}

func main() {
	con, error := Connect()
	if error != nil {
		fmt.Println(error)
	}
	//defer con.Close()
	fmt.Println("Connected to database")
	r := gin.Default()
	// Middleware
	r.Use()
	api := r.Group("/api")

	// Employee wiring
	empRepo := repository.NewEmployeeRepo(con)
	empService := service.NewService(empRepo)
	empHandler := handler.NewHandler(empService)
	empHandler.RegisterRoutes(api.Group("/employees"))

	log.Println("Server running on :8080")
	r.Run(":8080")

}
