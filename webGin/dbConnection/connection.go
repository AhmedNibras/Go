package dbConnection

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/gobuffalo/packr/v2"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "08091"
	dbname   = "album"
)

func Connection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func Migrate() {
	db := Connection()

	migrations := &migrate.PackrMigrationSource{
		Box: packr.NewBox("./db/migration"),
	}

	flagValue := flag.String("migration", "", "The name of the migration file")
	flag.Parse()

	switch *flagValue {
	case "up":
		n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Applied %d migrations!\n", n)
	case "down":
		n, err := migrate.Exec(db, "postgres", migrations, migrate.Down)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Applied %d migrations!\n", n)
	default:
		fmt.Println("Please specify a migration flag")
	}
}
