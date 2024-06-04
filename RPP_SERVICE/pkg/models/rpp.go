package models

import (
	"RPP_SERVICE/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Rpp struct {
	gorm.Model
	No_Registrasi_Jemaat string `json:"no_registrasi_jemaat"`
	Pelanggaran          string `json:"pelanggaran"`
	Tingkat_Pelanggaran  string `json:"tingkat_pelanggaran"`
	Nama_Lengkap         string `json:"nama_lengkap"`
	Jenis_Kelamin        string `json:"jenis_kelamin"`
	Alamat               string `json:"alamat"`
	Hukuman              string `json:"hukuman"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Rpp{})
}
func (b *Rpp) CreateRpp() *Rpp {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllRpp() []Rpp {
	var rpp []Rpp
	db.Find(&rpp)
	return rpp
}
func GetRppbyId(reg int64) (*Rpp, *gorm.DB) {
	var getRpp Rpp
	db := db.Where("ID=?", reg).Find(&getRpp)
	return &getRpp, db
}
func DeleteRpp(reg int64) Rpp {
	var rpp Rpp
	db.Where("ID=?", reg).Delete(rpp)
	return rpp
}
