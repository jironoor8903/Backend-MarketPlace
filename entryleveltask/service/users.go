package service

import (
	"database/sql"
	"entryleveltask/model"
	"log"

	"errors"

	"golang.org/x/crypto/bcrypt"

	"time"

	jwt "github.com/dgrijalva/jwt-go"

	_ "github.com/go-sql-driver/mysql"
)

//contains functions that will be called by the api
type UserService struct {
	DB *sql.DB
}

func (ps UserService) AddComment(comm model.CommentRequest) (model.Comment, error) {
	//todo: allow user to truly add a new comment
	//1 cases: one if parent comment doesnt exist (id not in table), error aka this is a reply... needs to be a valid reply
	//2 cases: one if parentid of comment is null add a new comment
	parentIDofComment := comm.ParentID
	if parentIDofComment != 0 {
		commentrows := ps.DB.QueryRow("SELECT id FROM comments WHERE id = ? ", parentIDofComment)
		var idek int
		err := commentrows.Scan(&idek)
		if err != nil {
			err1 := errors.New("no such parent comment exists")
			//todo: create new error
			return model.Comment{}, err1
		}
	}
	//check if the id of product id exists
	var count_product int64
	rows, err := ps.DB.Query("select count(*) as count_product from products")
	if err != nil {
		return model.Comment{}, err
	} else {
		for rows.Next() {
			rows.Scan(&count_product)
		}
	}

	if (count_product < comm.ProductID){
		log.Println("it reaches here ?",count_product)
		return model.Comment{}, errors.New("product does not exist")
	}
	

	result, err := ps.DB.Exec("INSERT INTO comments (comment, product_id, user_id, parent_id) VALUES (?, ?, ?, ?)", comm.Comment, comm.ProductID, 1, comm.ParentID)
	if err != nil {
		return model.Comment{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	answer := model.Comment{
		ID:        id,
		Comment:   comm.Comment,
		ProductID: comm.ProductID,
		UserID:    1,
		ParentID:  comm.ParentID,
	}

	return answer, nil
}


func (ps UserService) RegisterUser(username string, password string, email string) (string, error) {
	//check
	//1. if username is in database, return error saying "username taken"
	//2. if email is in database, return email has already been signed up"

	passwordHash, errr := getPasswordHash(password)
	if errr != nil {
		return "error", errr
	}

	result, err := ps.DB.Exec("INSERT INTO users (username, email, passwordHash) VALUES (?, ?, ?)", username, email, passwordHash)
	if err != nil {
		return "error", err
	}
	id, err := result.LastInsertId()

	newUser := model.User{
		ID:           id,
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
	}
	if err != nil {
		panic(err.Error())
	}

	return newUser.Username, nil
}
func (ps UserService) AuthenticateUser(username string, password string) (bool, error) {

	userrows := ps.DB.QueryRow("SELECT * FROM users WHERE username = ? ", username)
	us := &model.User{}
	err := userrows.Scan(&us.ID,
		us.Username,
		us.Email,
		us.PasswordHash,
	)
	if err != nil {
		return false, err
	}

	check := bcrypt.CompareHashAndPassword(
		[]byte(us.PasswordHash),
		[]byte(password),
	)
	return check == nil, nil
}

func getPasswordHash(password string) (string, error) {
	hash, error := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(hash), error
}

func InitUserService(db *sql.DB) UserService {
	newUserService := UserService{
		DB: db,
	}
	return newUserService
}

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

func (ps UserService) Signin(cred model.Loginrequest) (string, error) {
	var creds model.Loginrequest
	// Get the expected password from our in memory map
	userrow := ps.DB.QueryRow("SELECT * FROM users WHERE username = ? ", cred.Username)

	user := &model.User{}
	err := userrow.Scan(&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
	)

	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	expectedPassword := user.PasswordHash

	check := bcrypt.CompareHashAndPassword(
		[]byte(expectedPassword),
		[]byte(cred.Password),
	)

	if check != nil {
		return "", check
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(168 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &model.Claims{
		Username: creds.Username,
		ID:       user.ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", err
	}

	return tokenString, err

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself

}
