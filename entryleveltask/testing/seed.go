package testing

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	"github.com/rs/xid"
)

var logdetail bool

func SeedCommand(db *sql.DB, tblName string) error {
	var err error
	defer db.Close()
	start := time.Now()
	for x := 0; x < 994; x++ {
		err = seed(db, tblName, x)
		if err != nil {
			break
		}
	}
	duration := time.Since(start)
	log.Println("done seeding, elapsed: ", duration)

	return err
}

func seed(db *sql.DB, tblName string, x int) error {
	sql := "INSERT INTO " + "products" + "(title, artist, descriptions, category, price) VALUES "
	vals := []interface{}{}

	start := time.Now()

	for i := 0; i < 1000; i++ {
		var title string
		var artist string
		var description string
		var category string
		var price float64

		guid := xid.New()

		s1 := rand.NewSource(time.Now().UnixNano() + int64(guid.Counter()))
		r1 := rand.New(s1)
		i1 := r1.Intn(100)

	
		if i1 <= 40 {
			title = "Red Pants"
			artist = "Jiro"
			description= "Nice Pants"
			category = "Pants"
			price= 34.98
		} else {
			title = "White Shirt"
			artist = "Paul"
			description= "Nice Shirt"
			category = "Shirt"
			price= 34.98
		}
		

		sql += "(?, ?, ?, ?, ?, ?),"
		vals = append(vals, title, artist, description, category, price, )
	}

	sql = sql[0 : len(sql)-1]
	stmt, _ := db.Prepare(sql)
	res, err := stmt.Exec(vals...)

	if err != nil {
		panic(err.Error())
	}

	_, err = res.LastInsertId()

	if err != nil {
		log.Fatal(err)
		return err
	}

	// fmt.Printf("The last inserted row id: %d\n", lastId)
	if logdetail {
		duration := time.Since(start)
		log.Printf("%d. seeded for 1000, elapsed: %s\n", x, duration)
	}

	return nil
}