package controllers

import (
	"JEMAAT_PINDAH_SERVICE/pkg/models"
	"JEMAAT_PINDAH_SERVICE/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var newJemaatPindah models.JemaatPindah

func GetJemaatPindah(w http.ResponseWriter, r *http.Request) {
	newJemaatPindah := models.GetAllJemaatPindah()
	res, _ := json.Marshal(newJemaatPindah)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetJemaatPindahbyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jemaatPindahId := vars["pindahId"]
	ID_Jemaat_Pindah, err := strconv.ParseInt(jemaatPindahId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	JemaatPindahDetails, _ := models.GetJemaatPindahbyId(ID_Jemaat_Pindah)
	res, _ := json.Marshal(JemaatPindahDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateJemaatPindah(w http.ResponseWriter, r *http.Request) {
	createJemaatPindah := &models.JemaatPindah{}
	utils.ParseBody(r, createJemaatPindah)
	b := createJemaatPindah.CreateJemaatPindah()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteJemaatPindah(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jemaatPindahId := vars["pindahId"]
	ID_Jemaat_Pindah, err := strconv.ParseInt(jemaatPindahId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	jemaatPindah := models.DeleteJemaatPindah(ID_Jemaat_Pindah)
	res, _ := json.Marshal(jemaatPindah)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateJemaatPindah(w http.ResponseWriter, r *http.Request) {
	var updateJemaatPindah = &models.JemaatPindah{}
	utils.ParseBody(r, updateJemaatPindah)
	vars := mux.Vars(r)
	jemaatPindahId := vars["pindahId"]
	ID_Jemaat_Pindah, err := strconv.ParseInt(jemaatPindahId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	jemaatPindahDetails, db := models.GetJemaatPindahbyId(ID_Jemaat_Pindah)
	if updateJemaatPindah.Nama_Lengkap != "" {
		jemaatPindahDetails.Nama_Lengkap = updateJemaatPindah.Nama_Lengkap
	}
	if updateJemaatPindah.Nama_Keluarga != "" {
		jemaatPindahDetails.Nama_Keluarga = updateJemaatPindah.Nama_Keluarga
	}
	if updateJemaatPindah.Alamat != "" {
		jemaatPindahDetails.Alamat = updateJemaatPindah.Alamat
	}
	if updateJemaatPindah.Gereja_Asal != "" {
		jemaatPindahDetails.Gereja_Asal = updateJemaatPindah.Gereja_Asal
	}
	if updateJemaatPindah.Gereja_Tujuan != "" {
		jemaatPindahDetails.Gereja_Tujuan = updateJemaatPindah.Gereja_Tujuan
	}
	db.Save(&jemaatPindahDetails)
	res, _ := json.Marshal(jemaatPindahDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
