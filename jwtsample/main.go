package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

func CreateTokenEndpoint(w http.ResponseWriter, req *http.Request) {

	//in production you should check in your database if username & password are valid

	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"maikta": "akkuschrauber",
		"Timestamp": "in 3 jahr",
	})
	tokenString, err := token.SignedString([]byte("JWT-Sample"))
	if err != nil{
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(JwtToken{Token:tokenString})
}


func TestEndpoint(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r,"decoded")
	fmt.Println(decoded)
	//var user User
	//mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	json.NewEncoder(w).Encode(decoded)
}


func main(){
	router := mux.NewRouter()
	fmt.Printf("Starting web server")

	router.HandleFunc("/authenticate", CreateTokenEndpoint).Methods("POST")
	test:=router.PathPrefix("/test").Subrouter()
	test.HandleFunc("/trick", TestEndpoint);
	test.Use(ValidateMiddleware)


	//router.Use(ValidateMiddleware)
	log.Fatal(http.ListenAndServe(":9000", router))
}
type MiddlewareFunc func(http.Handler) http.Handler
func ValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("JWT-Sample"), nil
				})
				if err != nil{
					json.NewEncoder(w).Encode(Exception{Message:err.Error()})
					return
				}
				if token.Valid {
					context.Set(r,"decoded", token.Claims)
					next.ServeHTTP(w,r)
				} else {
					json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				}
			} else {
				json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
			}
		} else {
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}
	})
}



