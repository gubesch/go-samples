package main

import (
	"fmt"
	"github.com/gorilla/mux"
	. "github.com/gubesch/go-samples/apisample/models"
	"net/http"
	"time"
)




func main(){
	fmt.Println("test")
	Appointments = append(Appointments, &Appointment{ID:1, Brand:"Makita", Model:"DVR450Z", DateJoined:time.Now(),Owner:&Person{ID:1, Firstname:"Anton", Lastname:"Horvath"}})
	Appointments = append(Appointments, &Appointment{ID:2, Brand:"HP", Model:"Laserjet Pro M15W", DateJoined:time.Now(),Owner:&Person{ID:2, Firstname:"Dominik", Lastname:"Aschbacher"}})
	fmt.Println(Appointments)
	router := mux.NewRouter()
	router.HandleFunc("/service", GetAllAppointments).Methods("GET")
	router.HandleFunc("/service/{id}",GetSingleAppointment).Methods("GET")
	router.HandleFunc("/service", CreateNewAppointment).Methods("POST")
	router.HandleFunc("/service/{id}", DeleteAppointment).Methods("DELETE")
	router.HandleFunc("/service/{id}", UpdateAppointment).Methods("PUT")
	http.ListenAndServe(":8000", router)

}