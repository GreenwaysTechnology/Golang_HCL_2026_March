package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// connect to db
func connectdb() *pgxpool.Pool {
	//connect string
	connStr := "postgres://admin:secret@localhost:5432/mydb?sslmode=disable"
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse config: %v\n", err)
		os.Exit(1)
	}
	// Set pool limits
	config.MaxConns = 10
	config.MaxConnIdleTime = 5 * time.Minute
	// Create the pool
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return pool

}
func createUser(pool *pgxpool.Pool, name string, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `INSERT INTO users (name, email) VALUES ($1, $2)`
	_, err := pool.Exec(ctx, query, name, email)
	return err
}
func getAllUsers(pool *pgxpool.Pool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, name, email FROM users`
	rows, err := pool.Query(ctx, query)
	if err != nil {
		return err
	}
	// Crucial: Close rows to release the connection back to the pool
	defer rows.Close()

	fmt.Println("ID | Name | Email")
	fmt.Println("-------------------")

	for rows.Next() {
		var id int
		var name, email string

		// Scan must match the order of columns in your SELECT statement
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			return err
		}
		fmt.Printf("%d | %s | %s\n", id, name, email)
	}

	// Check if the loop stopped because of an error
	return rows.Err()
}

func main() {
	dbPool := connectdb()
	//close method is usefull for mememory data
	defer dbPool.Close()
	err := createUser(dbPool, "Subramanian", "subu@example.com")
	if err != nil {
		fmt.Printf("Create failed: %v\n", err)
	}
	err1 := getAllUsers(dbPool)
	if err1 != nil {
		fmt.Printf("GetAllUsers failed: %v\n", err1)
	}

	fmt.Println("Database operations completed successfully!")
}
