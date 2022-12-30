package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/gobuffalo/packr/v2"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

//by using flag pass the flag in the terminal in order to get that flag use --flag=-up || --flag=down
//if flag ibns down aplliy down migration if up apply up migration!

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:08091@localhost:5432/test?sslmode=disable")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}

	migrations := &migrate.PackrMigrationSource{
		Box: packr.NewBox("/migration"),
	}

	flagValue := flag.String("migration", "", "The name of the migration file")
	flag.Parse()

	switch *flagValue {
	case "up":
		n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("Applied %d migrations!\n", n)
	case "down":
		n, err := migrate.Exec(db, "postgres", migrations, migrate.Down)
		if err != nil {

			log.Fatal(err)
			return
		}
		fmt.Printf("Applied %d migrations!\n", n)
	}

}
