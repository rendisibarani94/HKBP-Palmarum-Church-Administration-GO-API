package controllers

import (
	"EVENT_SERVICE/pkg/models"
	"EVENT_SERVICE/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBaptis models.Event

func GetEvent(w http.ResponseWriter, r *http.Request) {
	newEvent := models.GetAllEvent()
	res, _ := json.Marshal(newEvent)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetEventbyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	ID_Event, err := strconv.ParseInt(eventId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	eventDetails, _ := models.GetEventbyId(ID_Event)
	res, _ := json.Marshal(eventDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	CreateEvent := &models.Event{}
	utils.ParseBody(r, CreateEvent)
	b := CreateEvent.CreateEvent()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	ID_Event, err := strconv.ParseInt(eventId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	event := models.DeleteEvent(ID_Event)
	res, _ := json.Marshal(event)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var updateEvent = &models.Event{}
	utils.ParseBody(r, updateEvent)
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	ID_Event, err := strconv.ParseInt(eventId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	eventDetails, db := models.GetEventbyId(ID_Event)
	if updateEvent.Nama_Event != "" {
		eventDetails.Nama_Event = updateEvent.Nama_Event
	}
	if updateEvent.Tempat_Event != "" {
		eventDetails.Tempat_Event = updateEvent.Tempat_Event
	}
	if updateEvent.Tanggal_Event != "" {
		eventDetails.Tanggal_Event = updateEvent.Tanggal_Event
	}
	if updateEvent.Pelayan_Event != "" {
		eventDetails.Pelayan_Event = updateEvent.Pelayan_Event
	}
	if updateEvent.Jumlah_Hadir != "" {
		eventDetails.Jumlah_Hadir = updateEvent.Jumlah_Hadir
	}
	if updateEvent.Keterangan != "" {
		eventDetails.Keterangan = updateEvent.Keterangan
	}
	db.Save(&eventDetails)
	res, _ := json.Marshal(eventDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
