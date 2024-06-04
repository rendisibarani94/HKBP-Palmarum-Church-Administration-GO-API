package controllers

import (
	"BAPTIS_SERVICE/pkg/models"
	"BAPTIS_SERVICE/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBaptis models.Baptis

func GetBaptis(w http.ResponseWriter, r *http.Request) {
	newBaptis := models.GetAllBaptis()
	res, _ := json.Marshal(newBaptis)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBaptisbyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	baptisId := vars["baptisId"]
	ID_Baptis, err := strconv.ParseInt(baptisId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	baptisDetails, _ := models.GetBaptisbyId(ID_Baptis)
	res, _ := json.Marshal(baptisDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBaptis(w http.ResponseWriter, r *http.Request) {
	CreateBaptis := &models.Baptis{}
	utils.ParseBody(r, CreateBaptis)
	b := CreateBaptis.CreateBaptis()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBaptis(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	baptisId := vars["baptisId"]
	ID_Baptis, err := strconv.ParseInt(baptisId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	baptis := models.DeleteBaptis(ID_Baptis)
	res, _ := json.Marshal(baptis)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateBaptis(w http.ResponseWriter, r *http.Request) {
	var updateBaptis = &models.Baptis{}
	utils.ParseBody(r, updateBaptis)
	vars := mux.Vars(r)
	baptisId := vars["baptisId"]
	ID_Baptis, err := strconv.ParseInt(baptisId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	baptisDetails, db := models.GetBaptisbyId(ID_Baptis)
	if updateBaptis.Nama_Lengkap != "" {
		baptisDetails.Nama_Lengkap = updateBaptis.Nama_Lengkap
	}
	if updateBaptis.Tempat_Lahir != "" {
		baptisDetails.Tempat_Lahir = updateBaptis.Tempat_Lahir
	}
	if updateBaptis.Tanggal_Lahir != "" {
		baptisDetails.Tanggal_Lahir = updateBaptis.Tanggal_Lahir
	}
	if updateBaptis.Nama_Ayah != "" {
		baptisDetails.Nama_Ayah = updateBaptis.Nama_Ayah
	}
	if updateBaptis.Nama_Ibu != "" {
		baptisDetails.Nama_Ibu = updateBaptis.Nama_Ibu
	}
	if updateBaptis.Tanggal_Baptis != "" {
		baptisDetails.Tanggal_Baptis = updateBaptis.Tanggal_Baptis
	}
	if updateBaptis.Pendeta != "" {
		baptisDetails.Pendeta = updateBaptis.Pendeta
	}
	db.Save(&baptisDetails)
	res, _ := json.Marshal(baptisDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
