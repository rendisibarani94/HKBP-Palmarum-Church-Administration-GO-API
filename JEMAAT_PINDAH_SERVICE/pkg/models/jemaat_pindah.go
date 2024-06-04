package models

import (
	"JEMAAT_PINDAH_SERVICE/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type JemaatPindah struct {
	gorm.Model
	Nama_Lengkap  string `json:"nama_lengkap"`
	Nama_Keluarga string `json:"nama_keluarga"`
	Alamat        string `json:"alamat"`
	Gereja_Asal   string `json:"gereja_asal"`
	Gereja_Tujuan string `json:"gereja_tujuan"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&JemaatPindah{})
}
func (b *JemaatPindah) CreateJemaatPindah() *JemaatPindah {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllJemaatPindah() []JemaatPindah {
	var jemaatPindah []JemaatPindah
	db.Find(&jemaatPindah)
	return jemaatPindah
}
func GetJemaatPindahbyId(reg int64) (*JemaatPindah, *gorm.DB) {
	var getJemaatPindah JemaatPindah
	db := db.Where("ID=?", reg).Find(&getJemaatPindah)
	return &getJemaatPindah, db
}
func DeleteJemaatPindah(reg int64) JemaatPindah {
	var jemaatPindah JemaatPindah
	db.Where("ID=?", reg).Delete(jemaatPindah)
	return jemaatPindah
}
