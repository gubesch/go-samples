package models

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)



func GetAllAppointments(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Appointments)
}
func GetSingleAppointment(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	for _, item := range Appointments  {
		id,err := strconv.ParseInt(parameters["id"],0,64)
		if err == nil {
			if int64(item.ID) == id {
				json.NewEncoder(w).Encode(item)
			}
		}
	}
}
func CreateNewAppointment(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	newId := len(Appointments) + 1
	var appointment Appointment
	json.NewDecoder(r.Body).Decode(&appointment)
	appointment.ID = newId
	Appointments = append(Appointments,&appointment)
	json.NewEncoder(w).Encode(Appointments)
}
func DeleteAppointment(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	id,err := strconv.ParseInt(parameters["id"],0,64)
	if(err == nil){
		for index, item := range Appointments {
			if int64(item.ID) == id {
				Appointments = append(Appointments[:index], Appointments[index+1])
				break
			}
		}
	}
	json.NewEncoder(w).Encode(Appointments)
}
func UpdateAppointment(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	id,err := strconv.ParseInt(parameters["id"],0,64)
	var appointment Appointment
	json.NewDecoder(r.Body).Decode(&appointment)
	if(err == nil){
		for _, item := range Appointments {
			if int64(item.ID) == id {
				item.Model = appointment.Model
				item.Owner = appointment.Owner;
				item.DateJoined = appointment.DateJoined
				item.Brand = appointment.Brand
			}
		}
	}
	json.NewEncoder(w).Encode(Appointments)
}