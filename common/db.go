// db/db.go

package db

import (
    "database/sql"
    "fmt"
    "github.com/labstack/echo/v4"
    _ "github.com/go-sql-driver/mysql"
	"os"
    "github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() error {
	var err error
	log := echo.New()

	if err := godotenv.Load(); err != nil {
		msg := "Erro ao carregar arquivo .env: %v"
        log.Logger.Fatal(msg, err)
        return fmt.Errorf(msg, err)
    }


	dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")

    dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbName
    DB, err = sql.Open("mysql", dsn)
    
	
	if err != nil {
		msg := "error database connection: %v"
		log.Logger.Fatal(msg, err)
        return fmt.Errorf(msg, err)
    }

    fmt.Println("Conexão com o banco de dados estabelecida com sucesso")
    return nil
}
