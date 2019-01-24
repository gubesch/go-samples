package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gubesch/go-samples/jwtsample/middleware"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}


func CreateTokenEndpoint(w http.ResponseWriter, req *http.Request) {

	//in production you should check in your database if username & password are valid

	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)

	//in the next lines you can define the tokens values
	//in my case username, makita, timestamp

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"maikta": "akkuschrauber",
		"Timestamp": "in 3 jahr",
	})
	//jwt-sample is the JWT secret this should be changed to a secure secret
	tokenString, err := token.SignedString([]byte("JWT-Sample"))
	if err != nil{
		fmt.Println(err)
	}
	_=json.NewEncoder(w).Encode(JwtToken{Token:tokenString})
}


func TestEndpoint(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r,"decoded")
	fmt.Println(decoded)
	//mapstructer is only needed if you want to parse the decode message into an object
	//var user User
	//mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	_=json.NewEncoder(w).Encode(decoded)
}


func main(){
	router := mux.NewRouter()
	fmt.Printf("Starting web server")

	//the router has an authentication route which is not behind a middleware

	router.HandleFunc("/authenticate", CreateTokenEndpoint).Methods("POST")

	//all other routes are behind a middleware
	test:=router.PathPrefix("/test").Subrouter()
	test.HandleFunc("/trick", TestEndpoint)
	test.Use(middleware.ValidateMiddleware)


	//router.Use(ValidateMiddleware)
	log.Fatal(http.ListenAndServe(":9000", router))
}

// create middleware function for routes behind middleware





