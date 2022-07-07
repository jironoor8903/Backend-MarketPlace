package seeds

import (
	"database/sql"
	"flag"

	"log"
	"os"
	"github.com/go-sql-driver/mysql"
	"github.com/redhajuanda/goseeder/seeds"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	handleArgs()
}

func handleArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			cfg := mysql.Config{
				User:   "root",
				Passwd: "Jiron256$",
				Net:    "tcp",
				Addr:   "127.0.0.1:3306",
				DBName: "test",
			}
			// Get a database handle.
			var err error
			db, err := sql.Open("mysql", cfg.FormatDSN())
			if err != nil {
				log.Fatal(err)
			}
			seeds.Execute(db, args[1:]...)
			os.Exit(0)
		}
	}
}