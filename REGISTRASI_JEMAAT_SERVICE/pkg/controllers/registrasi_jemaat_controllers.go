package controllers

import (
	"REGISTRASI_JEMAAT_SERVICE/pkg/models"
	"REGISTRASI_JEMAAT_SERVICE/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewJemaat models.Jemaat

func GetJemaat(w http.ResponseWriter, r *http.Request) {
	newJemaat := models.GetAllJemaat()
	res, _ := json.Marshal(newJemaat)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetJemaatbyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jemaatId := vars["jemaatId"]
	ID_Jemaat, err := strconv.ParseInt(jemaatId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	jemaatDetails, _ := models.GetJemaatbyId(ID_Jemaat)
	res, _ := json.Marshal(jemaatDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateJemaat(w http.ResponseWriter, r *http.Request) {
	CreateJemaat := &models.Jemaat{}
	utils.ParseBody(r, CreateJemaat)
	b := CreateJemaat.CreateJemaat()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteJemaat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jemaatId := vars["jemaatId"]
	ID_Jemaat, err := strconv.ParseInt(jemaatId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	jemaat := models.DeleteJemaat(ID_Jemaat)
	res, _ := json.Marshal(jemaat)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateJemaat(w http.ResponseWriter, r *http.Request) {
	var updateJemaat = &models.Jemaat{}
	utils.ParseBody(r, updateJemaat)
	vars := mux.Vars(r)
	jemaat_Id := vars["jemaatId"]
	ID_Jemaat, err := strconv.ParseInt(jemaat_Id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	jemaatDetails, db := models.GetJemaatbyId(ID_Jemaat)
	if updateJemaat.No_Registrasi_Jemaat != "" {
		jemaatDetails.No_Registrasi_Jemaat = updateJemaat.No_Registrasi_Jemaat
	}
	if updateJemaat.Nama_Lengkap != "" {
		jemaatDetails.Nama_Lengkap = updateJemaat.Nama_Lengkap
	}
	if updateJemaat.Nama_Keluarga != "" {
		jemaatDetails.Nama_Keluarga = updateJemaat.Nama_Keluarga
	}
	if updateJemaat.Tempat_Lahir != "" {
		jemaatDetails.Tempat_Lahir = updateJemaat.Tempat_Lahir
	}
	if updateJemaat.Tanggal_Lahir != "" {
		jemaatDetails.Tanggal_Lahir = updateJemaat.Tanggal_Lahir
	}
	if updateJemaat.Alamat != "" {
		jemaatDetails.Alamat = updateJemaat.Alamat
	}
	if updateJemaat.Jenis_Kelamin != "" {
		jemaatDetails.Jenis_Kelamin = updateJemaat.Jenis_Kelamin
	}
	db.Save(&jemaatDetails)
	res, _ := json.Marshal(jemaatDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
