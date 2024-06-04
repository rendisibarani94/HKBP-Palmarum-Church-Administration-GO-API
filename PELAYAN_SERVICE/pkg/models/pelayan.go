package models

import (
	"PELAYAN_SERVICE/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Pelayan struct {
	gorm.Model
	Nama_Lengkap          string `json:"nama_lengkap"`
	Jabatan               string `json:"jabatan"`
	Alamat                string `json:"alamat"`
	Tanggal_Menjabat      string `json:"tanggal_menjabat"`
	Tanggal_Akhir_Jabatan string `json:"tanggal_akhir_jabatan"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Pelayan{})
}
func (b *Pelayan) CreatePelayan() *Pelayan {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllPelayan() []Pelayan {
	var pelayan []Pelayan
	db.Find(&pelayan)
	return pelayan
}
func GetPelayanbyId(reg int64) (*Pelayan, *gorm.DB) {
	var getPelayan Pelayan
	db := db.Where("ID=?", reg).Find(&getPelayan)
	return &getPelayan, db
}
func DeletePelayan(reg int64) Pelayan {
	var pelayan Pelayan
	db.Where("ID=?", reg).Delete(pelayan)
	return pelayan
}
