package main

import (
	"fmt"
	"gormginapp/handler"
	"gormginapp/middlewares"
	"gormginapp/model"
	"gormginapp/repository"
	"gormginapp/service"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, error) {
	dsn := "host=localhost user=admin password=secret dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil

}

func main() {
	// 1. Attempt to connect
	gormDB, err := ConnectDb()
	if err != nil {
		log.Fatalf("Failed to initialize GORM: %v", err)
	}

	// 2. Get the underlying sql.DB instance to Ping
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying DB instance: %v", err)
	}

	// 3. Ping the database
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Database is unreachable: %v", err)
	}

	fmt.Println("✅ Connection successful! Database is alive and kicking.")
	// Auto migration
	gormDB.AutoMigrate(&model.Employee{})
	r := gin.New()
	r.Use(middlewares.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	// Wiring
	repo := repository.NewEmployeeRepository(gormDB)
	service := service.NewEmployeeService(repo)
	handler := handler.NewEmployeeHandler(service)

	handler.RegisterRoutes(api.Group("/employees"))

	log.Println("🚀 Server running on :8080")
	r.Run(":8080")

}
