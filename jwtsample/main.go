package main

import (
	"fmt"
	"github.com/gorilla/mux"
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

func main(){
	router := mux.NewRouter()
	fmt.Printf("Starting web server")
	router.HandleFunc("/authenticate", CreateTokenEndpoint).Methods("POST")

}
