package controllers

import (
	"RPP_SERVICE/pkg/models"
	"RPP_SERVICE/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewRpp models.Rpp

func GetRpp(w http.ResponseWriter, r *http.Request) {
	newRpp := models.GetAllRpp()
	res, _ := json.Marshal(newRpp)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetRppbyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rppId := vars["rppId"]
	ID_Rpp, err := strconv.ParseInt(rppId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	rppDetails, _ := models.GetRppbyId(ID_Rpp)
	res, _ := json.Marshal(rppDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateRpp(w http.ResponseWriter, r *http.Request) {
	CreateRpp := &models.Rpp{}
	utils.ParseBody(r, CreateRpp)
	b := CreateRpp.CreateRpp()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteRpp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rppId := vars["rppId"]
	ID_Rpp, err := strconv.ParseInt(rppId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	student := models.DeleteRpp(ID_Rpp)
	res, _ := json.Marshal(student)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateRpp(w http.ResponseWriter, r *http.Request) {
	var updateRpp = &models.Rpp{}
	utils.ParseBody(r, updateRpp)
	vars := mux.Vars(r)
	rpp_Id := vars["rppId"]
	ID_Rpp, err := strconv.ParseInt(rpp_Id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	rppDetails, db := models.GetRppbyId(ID_Rpp)
	if updateRpp.No_Registrasi_Jemaat != "" {
		rppDetails.No_Registrasi_Jemaat = updateRpp.No_Registrasi_Jemaat
	}
	if updateRpp.Pelanggaran != "" {
		rppDetails.Pelanggaran = updateRpp.Pelanggaran
	}
	if updateRpp.Tingkat_Pelanggaran != "" {
		rppDetails.Tingkat_Pelanggaran = updateRpp.Tingkat_Pelanggaran
	}
	if updateRpp.Nama_Lengkap != "" {
		rppDetails.Nama_Lengkap = updateRpp.Nama_Lengkap
	}
	if updateRpp.Jenis_Kelamin != "" {
		rppDetails.Jenis_Kelamin = updateRpp.Jenis_Kelamin
	}
	if updateRpp.Alamat != "" {
		rppDetails.Alamat = updateRpp.Alamat
	}
	if updateRpp.Hukuman != "" {
		rppDetails.Hukuman = updateRpp.Hukuman
	}
	db.Save(&rppDetails)
	res, _ := json.Marshal(rppDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
