package controllers

import (
	"SIDI_SERVICE/pkg/models"
	"SIDI_SERVICE/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewSidi models.Sidi

func GetSidi(w http.ResponseWriter, r *http.Request) {
	newSidi := models.GetAllSidi()
	res, _ := json.Marshal(newSidi)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetSidibyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sidiId := vars["sidiId"]
	ID_Sidi, err := strconv.ParseInt(sidiId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	sidiDetails, _ := models.GetSidibyId(ID_Sidi)
	res, _ := json.Marshal(sidiDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateSidi(w http.ResponseWriter, r *http.Request) {
	CreateSidi := &models.Sidi{}
	utils.ParseBody(r, CreateSidi)
	b := CreateSidi.CreateSidi()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteSidi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sidiId := vars["sidiId"]
	ID_Sidi, err := strconv.ParseInt(sidiId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	sidi := models.DeleteSidi(ID_Sidi)
	res, _ := json.Marshal(sidi)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateSidi(w http.ResponseWriter, r *http.Request) {
	var updateSidi = &models.Sidi{}
	utils.ParseBody(r, updateSidi)
	vars := mux.Vars(r)
	sidiId := vars["sidiId"]
	ID_Sidi, err := strconv.ParseInt(sidiId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	sidiDetails, db := models.GetSidibyId(ID_Sidi)
	if updateSidi.Nama_Lengkap != "" {
		sidiDetails.Nama_Lengkap = updateSidi.Nama_Lengkap
	}
	if updateSidi.Tempat_Lahir != "" {
		sidiDetails.Tempat_Lahir = updateSidi.Tempat_Lahir
	}
	if updateSidi.Tanggal_Lahir != "" {
		sidiDetails.Tanggal_Lahir = updateSidi.Tanggal_Lahir
	}
	if updateSidi.Tenggal_Sidi != "" {
		sidiDetails.Tenggal_Sidi = updateSidi.Tenggal_Sidi
	}
	if updateSidi.Lokasi_Sidi != "" {
		sidiDetails.Lokasi_Sidi = updateSidi.Lokasi_Sidi
	}
	if updateSidi.Nama_Pendeta != "" {
		sidiDetails.Nama_Pendeta = updateSidi.Nama_Pendeta
	}
	db.Save(&sidiDetails)
	res, _ := json.Marshal(sidiDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
