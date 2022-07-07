package testing

import (
	"database/sql"
	"entryleveltask/service"
	"fmt"
	"log"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)
	
func Setup() (*sql.DB){
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

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db
}

func TestGetProducts(t *testing.T){
	db := Setup()
	productServiceTest := service.InitProductService(db)
	products, err := productServiceTest.GetProductsWithService()

	require.NoError(t, err)
	require.NotEmpty(t, products)

	require.Equal(t, products[0].ID, int64(1))
	require.Equal(t, products[0].Title, "White Pants")
	require.Equal(t, products[0].Description, "Nice Pants")
	require.Equal(t, products[0].Category, "Pants")
	require.Equal(t, products[0].Price, 34.98)

	require.Equal(t, products[1].ID, int64(2))
	require.Equal(t, products[1].Title, "White Pants")
	require.Equal(t, products[1].Description, "Nice Pants")
	require.Equal(t, products[1].Category, "Pants")
	require.Equal(t, products[1].Price, 34.98)

}

func TestGetProductsByTitle(t *testing.T){
	db := Setup()
	productServiceTest := service.InitProductService(db)
	title := "White Pants"
	products, err := productServiceTest.GetProductbytitleWithService(title)

	require.NoError(t, err)
	require.NotEmpty(t, products)

	require.Equal(t, products[0].ID, int64(1))
	require.Equal(t, products[0].Title, "White Pants")
	require.Equal(t, products[0].Description, "Nice Pants")
	require.Equal(t, products[0].Category, "Pants")
	require.Equal(t, products[0].Price, 34.98)

	require.Equal(t, products[1].ID, int64(2))
	require.Equal(t, products[1].Title, "White Pants")
	require.Equal(t, products[1].Description, "Nice Pants")
	require.Equal(t, products[1].Category, "Pants")
	require.Equal(t, products[1].Price, 34.98)
}

func TestGetProductsByCategory(t *testing.T){
	db := Setup()
	productServiceTest := service.InitProductService(db)
	category := "Pants"
	products, err := productServiceTest.GetProductbycategoryWithService(category)

	require.NoError(t, err)
	require.NotEmpty(t, products)

	require.Equal(t, products[0].ID, int64(1))
	require.Equal(t, products[0].Title, "White Pants")
	require.Equal(t, products[0].Description, "Nice Pants")
	require.Equal(t, products[0].Category, "Pants")
	require.Equal(t, products[0].Price, 34.98)

	require.Equal(t, products[1].ID, int64(2))
	require.Equal(t, products[1].Title, "White Pants")
	require.Equal(t, products[1].Description, "Nice Pants")
	require.Equal(t, products[1].Category, "Pants")
	require.Equal(t, products[1].Price, 34.98)
}

func TestViewProduct(t *testing.T){
	db := Setup()
	productServiceTest := service.InitProductService(db)
	productview, err := productServiceTest.ViewProduct("1")

	require.NoError(t, err)
	require.NotEmpty(t, productview)

	require.Equal(t, productview.ID, int64(1))
	require.Equal(t, productview.Title, "White Pants")
	require.Equal(t, productview.Category, "Pants")
	require.Equal(t, productview.Photos, "")



	actual := productview.Comments

	require.Equal(t, actual[0].ID, int64(1))
	require.Equal(t, actual[0].Comment, "Worth it and Good price")
	require.Equal(t, actual[0].ProductID, int64(1))
	require.Equal(t, actual[0].UserID, int64(1))
	require.Equal(t, actual[0].ParentID, int64(0))

	require.Equal(t, actual[1].ID, int64(3))
	require.Equal(t, actual[1].Comment, "Stop lying its bad")
	require.Equal(t, actual[1].ProductID, int64(1))
	require.Equal(t, actual[1].UserID, int64(1))
	require.Equal(t, actual[1].ParentID, int64(1))


}