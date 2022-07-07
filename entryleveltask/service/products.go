package service

import (
	"database/sql"
	"entryleveltask/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//contains functions that will be called by the api

type ProductService struct {
	DB *sql.DB
}

func (ps ProductService) GetProductsWithService() ([]model.Product, error) {
	rows, err := ps.DB.Query("SELECT * FROM products LIMIT 20;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// create array of products
	//add querry results to array
	// return the array
	var products []model.Product

	for rows.Next() {

		var pro model.Product
		if err := rows.Scan(&pro.ID, &pro.Title, &pro.Artist, &pro.Description, &pro.Category, &pro.Price); err != nil {
			log.Println(err.Error())
			return nil, err
		}
		products = append(products, pro)
	}

	return products, nil
}

func (ps ProductService) GetProductbytitleWithService(titleinput string) ([]model.Product, error) {
	rows, err := ps.DB.Query("SELECT * FROM products WHERE title = ? LIMIT 20", titleinput)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var pro model.Product
		if err := rows.Scan(&pro.ID, &pro.Title, &pro.Artist, &pro.Description, &pro.Category, &pro.Price); err != nil {
			return nil, err
		}
		products = append(products, pro)
	}

	return products, nil

}

func (ps ProductService) GetProductbycategoryWithService(categoryinput string) ([]model.Product, error) {
	rows, err := ps.DB.Query("SELECT * FROM products WHERE category = ? LIMIT 20", categoryinput)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var pro model.Product
		if err := rows.Scan(&pro.ID, &pro.Title, &pro.Artist, &pro.Description, &pro.Category, &pro.Price); err != nil {
			return nil, err
		}
		products = append(products, pro)
	}

	return products, nil

}
func (ps ProductService) ViewProduct(productID string) (model.ProductView, error) {
	log.Printf("start function")
	rows, err := ps.DB.Query("SELECT * FROM comments WHERE product_id = ? LIMIT 20", productID)
	log.Printf("afterdb")
	if err != nil {
		log.Printf("error")
		return model.ProductView{}, err
	}
	defer rows.Close()

	var commentsarray []model.Comment
	log.Printf("array initialized")
	for rows.Next() {
		var com model.Comment
		if err := rows.Scan(&com.ID, &com.Comment, &com.ProductID, &com.UserID, &com.ParentID); err != nil {
			return model.ProductView{}, err
		}
		commentsarray = append(commentsarray, com)
		log.Printf("error3")
		log.Printf("%v", com)
	}

	productrows := ps.DB.QueryRow("SELECT * FROM products WHERE id = ? LIMIT 20", productID)
	prod := &model.Product{}
	err = productrows.Scan(&prod.ID,
		&prod.Title,
		&prod.Artist,
		&prod.Description,
		&prod.Category,
		&prod.Price)
	log.Printf("%v", prod)

	if err != nil && err != sql.ErrNoRows {
		// log the error
		log.Print(err)
		return model.ProductView{}, err
	}

	viewanswer := model.ProductView{
		ID:          prod.ID,
		Title:       prod.Title,
		Description: prod.Description,
		Category:    prod.Category,
		Photos:      "",
		Comments:    commentsarray,
	}

	return viewanswer, nil
}

// func (ps ProductService) AddComment(comm model.CommentRequest) (model.Comment, error) {
// 	//todo: allow user to truly add a new comment
// 	//1 cases: one if parent comment doesnt exist (id not in table), error aka this is a reply... needs to be a valid reply
// 	//2 cases: one if parentid of comment is null add a new comment
// 	parentIDofComment := comm.ParentID
// 	commentrows := ps.DB.QueryRow("SELECT id FROM comments WHERE id = ? ", parentIDofComment)
// 	var idek int
// 	err := commentrows.Scan(&idek)

// if err != nil && err != sql.ErrNoRows {
//     return model.Comment{}, fmt.Errorf("No such parent comment exists")
// }

// 	result, err := ps.DB.Exec("INSERT INTO comments (comment, product_id, user_id, parent_id) VALUES (?, ?, ?, ?)", comm.Comment, comm.ProductID, comm.UserID, comm.ParentID)
// 	if err != nil {
// 		return model.Comment{}, err
// 	}
// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	log.Print(id)
// 	comm.ID = id
// 	return comm, nil
// }
// func (ps ProductService) RegisterUser(username string, password string, email string) (string, error) {
// 	//check
// 	//1. if username is in database, return error saying "username taken"
// 	//2. if email is in database, return email has already been signed up"

// 	passwordHash, errr := getPasswordHash(password)
// 	if errr != nil {
// 		return "error", errr
// 	}

// 	result, err := ps.DB.Exec("INSERT INTO users (username, email, passwordHash) VALUES (?, ?, ?)", username, email, passwordHash)
// 	if err != nil {
// 		return "error", err
// 	}
// 	id, err := result.LastInsertId()

// 	newUser := model.User{
// 		ID:           id,
// 		Username:     username,
// 		Email:        email,
// 		PasswordHash: passwordHash,
// 	}
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	return newUser.Username, nil
// }

// func (ps ProductService) AuthenticateUser(username string, password string) (bool, error) {

// 	userrows := ps.DB.QueryRow("SELECT * FROM users WHERE username = ? ", username)
// 	us := &model.User{}
// 	err := userrows.Scan(&us.ID,
// 		us.Username,
// 		us.Email,
// 		us.PasswordHash,
// 	)
// 	if err != nil {
// 		return false, err
// 	}

// 	check := bcrypt.CompareHashAndPassword(
// 		[]byte(us.PasswordHash),
// 		[]byte(password),
// 	)
// 	return check == nil, nil
// }

// func getPasswordHash(password string) (string, error) {
// 	hash, error := bcrypt.GenerateFromPassword([]byte(password), 0)
// 	return string(hash), error
// }

func InitProductService(db *sql.DB) ProductService {
	newProductService := ProductService{
		DB: db,
	}
	return newProductService
}

//create the querry
//use the database
//pass the database created in main to service
//we need to fill the db variables with the values from main
//the route goes from main to api to service
