package seeds

import (
	"github.com/bxcodec/faker/v3"
)

func (s Seed) CustomerSeed() {

	for i := 0; i < 100; i++ {
		//prepare the statement
			artist := "Jiro"
			description := "Nice Pants"
			category := "Pants"
			price := 34.98

		stmt, _ := s.db.Prepare(`INSERT INTO products (title, artist, descriptions, category, price) VALUES (?,?,?,?,?)`)
		// execute query
		_, err := stmt.Exec("John",faker.Name(), artist, description, category,price )
		if err != nil {
			panic(err)
		}
	}
}