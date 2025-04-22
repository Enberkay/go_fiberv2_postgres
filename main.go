package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	sdb, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	db = sdb

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	//Connection Database Successful.
	fmt.Println("Connection Database Successful.")

	err = createProduct(&Product{Name: "Go product 2", Price: 444})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Create Successful !")
}

func createProduct(product *Product) error {

	_, err := db.Exec(
		"INSERT INTO public.products(name, price)VALUES ($1, $2);",
		product.Name,
		product.Price,
	)

	return err
}
