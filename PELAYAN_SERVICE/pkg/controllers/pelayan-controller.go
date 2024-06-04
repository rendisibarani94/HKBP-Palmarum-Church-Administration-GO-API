package controllers

import (
	"PELAYAN_SERVICE/pkg/models"
	"PELAYAN_SERVICE/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var newPelayan models.Pelayan

func GetPelayan(w http.ResponseWriter, r *http.Request) {
	newPelayan := models.GetAllPelayan()
	res, _ := json.Marshal(newPelayan)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetPelayanbyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pelayanId := vars["pelayanId"]
	ID_Pelayan, err := strconv.ParseInt(pelayanId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	pelayanDetails, _ := models.GetPelayanbyId(ID_Pelayan)
	res, _ := json.Marshal(pelayanDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreatePelayan(w http.ResponseWriter, r *http.Request) {
	createPelayan := &models.Pelayan{}
	utils.ParseBody(r, createPelayan)
	b := createPelayan.CreatePelayan()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeletePelayan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pelayanId := vars["pelayanId"]
	ID_Pelayan, err := strconv.ParseInt(pelayanId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	pelayan := models.DeletePelayan(ID_Pelayan)
	res, _ := json.Marshal(pelayan)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdatePelayan(w http.ResponseWriter, r *http.Request) {
	var updatePelayan = &models.Pelayan{}
	utils.ParseBody(r, updatePelayan)
	vars := mux.Vars(r)
	pelayanId := vars["pelayanId"]
	ID_Pelayan, err := strconv.ParseInt(pelayanId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	pelayanDetails, db := models.GetPelayanbyId(ID_Pelayan)
	if updatePelayan.Nama_Lengkap != "" {
		pelayanDetails.Nama_Lengkap = updatePelayan.Nama_Lengkap
	}
	if updatePelayan.Jabatan != "" {
		pelayanDetails.Jabatan = updatePelayan.Jabatan
	}
	if updatePelayan.Alamat != "" {
		pelayanDetails.Alamat = updatePelayan.Alamat
	}
	if updatePelayan.Tanggal_Menjabat != "" {
		pelayanDetails.Tanggal_Menjabat = updatePelayan.Tanggal_Menjabat
	}
	if updatePelayan.Tanggal_Akhir_Jabatan != "" {
		pelayanDetails.Tanggal_Akhir_Jabatan = updatePelayan.Tanggal_Akhir_Jabatan
	}
	db.Save(&pelayanDetails)
	res, _ := json.Marshal(pelayanDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
