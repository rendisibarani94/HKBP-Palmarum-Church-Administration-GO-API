package controllers

import (
	"JADWAL_IBADAH_SERVICE/pkg/models"
	"JADWAL_IBADAH_SERVICE/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var newIbadah models.Ibadah

func GetIbadah(w http.ResponseWriter, r *http.Request) {
	newIbadah := models.GetAllIbadah()
	res, _ := json.Marshal(newIbadah)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetIbadahbyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ibadahId := vars["ibadahId"]
	ID_Ibadah, err := strconv.ParseInt(ibadahId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	ibadahDetails, _ := models.GetIbadahbyId(ID_Ibadah)
	res, _ := json.Marshal(ibadahDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateIbadah(w http.ResponseWriter, r *http.Request) {
	CreateIbadah := &models.Ibadah{}
	utils.ParseBody(r, CreateIbadah)
	b := CreateIbadah.CreateIbadah()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteIbadah(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ibadahId := vars["ibadahId"]
	ID_Ibadah, err := strconv.ParseInt(ibadahId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	ibadah := models.DeleteIbadah(ID_Ibadah)
	res, _ := json.Marshal(ibadah)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateIbadah(w http.ResponseWriter, r *http.Request) {
	var updateIbadah = &models.Ibadah{}
	utils.ParseBody(r, updateIbadah)
	vars := mux.Vars(r)
	ibadahId := vars["ibadahId"]
	ID_Ibadah, err := strconv.ParseInt(ibadahId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	ibadahDetails, db := models.GetIbadahbyId(ID_Ibadah)
	if updateIbadah.Nama_Minggu != "" {
		ibadahDetails.Nama_Minggu = updateIbadah.Nama_Minggu
	}
	if updateIbadah.Jenis_Ibadah != "" {
		ibadahDetails.Jenis_Ibadah = updateIbadah.Jenis_Ibadah
	}
	if updateIbadah.Tanggal_Ibadah != "" {
		ibadahDetails.Tanggal_Ibadah = updateIbadah.Tanggal_Ibadah
	}
	if updateIbadah.Pendeta_Melayani != "" {
		ibadahDetails.Pendeta_Melayani = updateIbadah.Pendeta_Melayani
	}
	db.Save(&ibadahDetails)
	res, _ := json.Marshal(ibadahDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
