package api

import (
	"entryleveltask/model"
	"entryleveltask/service"
	"net/http"

	"database/sql"
	// "github.com/gorilla/mux"
	// "github.com/gorilla/sessions"

	"log"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"

)

// product represents data about a collection .
type createProductRequest struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Artist      string  `json:"artist"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Comments    string  `json:"comments"`
}

// products slice to seed product data.
var products = []createProductRequest{
	{ID: "1", Title: "Pants", Artist: "John", Description: "Pants", Category: "Pants", Price: 20000},
	{ID: "2", Title: "Shirt", Artist: "Paul", Description: "Shirt", Category: "Shirt", Price: 50000},
	{ID: "3", Title: "Underwear", Artist: "Jiro", Description: "Underwear", Category: "Underwear", Price: 10000},
}

//create factory that have some API, which is for this example products, get products, post products
//create struct name it produ
type Productapi struct {
	// this struct will be called by our main function
	//object that have api for products
	svc  service.ProductService
	usvc service.UserService
}



func InitApi(router *gin.Engine, db *sql.DB) {
	newProduct := Productapi{
		svc:  service.InitProductService(db),
		usvc: service.UserService{db},
	}

	router.GET("/products", newProduct.getProducts)
	router.POST("/products", newProduct.postProducts)
	router.GET("/products/title/:title", newProduct.getProductsbyTitle2)
	router.GET("/products/category/:category", newProduct.getProductsbyCategory)
	router.POST("/products/comment", newProduct.addComment)
	router.GET("/products/view/:productID", newProduct.viewProduct)
	router.POST("/users/register", newProduct.register)
	router.POST("/users/signin", newProduct.login)
}



// func main(){
// 	r := mux.NewRouter()
//     r.HandleFunc("/login", loginHandler).Methods("POST")
// }

// func (pa Productapi) loginHandler(c *gin.Context){
// 	username := c.Param("username")
// 	password := c.Param("password")

// 	//todo ask Pak Paulus how to check if the username and password
// 	if exists {
//         // It returns a new session if the sessions doesn't exist
//         session, _ := store.Get(r, "session.id")
//         if storedPassword == password {
//             session.Values["authenticated"] = true
//             // Saves all sessions used during the current request
//             session.Save(r, w)
//         } else {
//             http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
//         }
//         w.Write([]byte("Login successfully!"))
//     }
// }

// getAlbums responds with the list of all albums as JSON.
func (pa Productapi) getProducts(c *gin.Context) {
	answer, err := pa.svc.GetProductsWithService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, answer)
	// c.IndentedJSON(http.StatusOK, products)
}

func (pa Productapi) viewProduct(c *gin.Context) {
	productID := c.Param("productID")
	answer, err := pa.svc.ViewProduct(productID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, answer)
	// c.IndentedJSON(http.StatusOK, products)
}
func (pa Productapi) addComment(c *gin.Context) {
	//check token
	var jwtKey = []byte("my_secret_key")

	tokenHeader := c.Request.Header["Authorization"][0]
	if !strings.Contains(tokenHeader, "Bearer") {
		log.Println("1")
		c.IndentedJSON(http.StatusInternalServerError, "err")
		return 
	}
	tokenStringArray := strings.Split(tokenHeader, " ")
	tokenString := tokenStringArray[1]

	claims := &model.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		log.Print("parsewclaim")
		return jwtKey, nil
	})
	log.Print("parsewclaimpassedone")
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println("1")
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		log.Println(err)
		return
	}
	if !tkn.Valid {
		log.Println("2")
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	//check expiration and if the token is valid 
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 1*time.Second {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	var newComment model.CommentRequest
	if err := c.BindJSON(&newComment); err != nil {
		log.Println("4")
		c.IndentedJSON(http.StatusInternalServerError, err)
		return 
	}
	answer, err := pa.usvc.AddComment(newComment)

	if err != nil {
		log.Println("5")
		c.IndentedJSON(http.StatusInternalServerError, err)
		return 
	}
	c.IndentedJSON(http.StatusCreated, "Succesfully Added Comment: " + answer.Comment)
	return
}

func (pa Productapi) register(c *gin.Context) {
	var register model.Register
	if err := c.BindJSON(&register); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	answer, err := pa.usvc.RegisterUser(register.Username, register.Password, register.Email)
	log.Printf(register.Username +"hello")

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusCreated, "Successfully created user: "+ answer)
}
func (pa Productapi) login(c *gin.Context) {
	var login model.Loginrequest
	if err := c.BindJSON(&login); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	answer, err := pa.usvc.Signin(login)
	log.Printf(answer)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusCreated, "Login Successful")
}

func (pa Productapi) getProductsbyTitle2(c *gin.Context) {
	title := c.Param("title")
	answer, err := pa.svc.GetProductbytitleWithService(title)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, answer)
	// c.IndentedJSON(http.StatusOK, products)
}
func (pa Productapi) getProductsbyCategory(c *gin.Context) {
	category := c.Param("category")
	answer, err := pa.svc.GetProductbycategoryWithService(category)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, answer)
	// c.IndentedJSON(http.StatusOK, products)
}

// postAlbums adds an album from JSON received in the request body.
func (pa Productapi) postProducts(c *gin.Context) {
	var newProduct createProductRequest

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	// Add the new album to the slice.
	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func (pa Productapi) getProductByTitle(c *gin.Context) {
	title := c.Param("title")

	// Loop over the list of products, looking for
	// an album whose ID value matches the parameter.
	for _, a := range products {
		if a.Title == title {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}
 