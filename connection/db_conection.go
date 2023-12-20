package connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "postgres"
)

func db() string {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return connStr
}

func createTable(db *sql.DB) error {
	// Create a simple table called "example_table"
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS example_table (
			id SERIAL PRIMARY KEY,
			Book_name VARCHAR(50) NOT NULL,
			Authore_name VARCHAR(50) NOT NULL,
			Price  INT 
		);
	`)
	return err
}



func Connection() {
	db, err := sql.Open("postgres", db())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	err = createTable(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table 'example_table' created successfully")

	

	// fmt.Println("Data inserted into 'example_table' successfully")
}
