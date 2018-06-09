package shortener

import (
	"database/sql"
	"log"
	"os"

	// Add mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//InitDB inits mysql database
func InitDB(datasource string) *sql.DB {
	var db *sql.DB
	var err error
	db, err = sql.Open(datasource, os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp(mysql:"+os.Getenv("DB_PORT")+")/"+os.Getenv("DATABASE_NAME"))
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	return db
}
