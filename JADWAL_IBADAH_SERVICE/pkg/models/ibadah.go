package models

import (
	"JADWAL_IBADAH_SERVICE/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Ibadah struct {
	gorm.Model
	Nama_Minggu      string `json:"nama_minggu"`
	Jenis_Ibadah     string `json:"jenis_ibadah"`
	Tanggal_Ibadah   string `json:"tanggal_ibadah"`
	Pendeta_Melayani string `json:"pendeta_melayani"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Ibadah{})
}
func (b *Ibadah) CreateIbadah() *Ibadah {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllIbadah() []Ibadah {
	var ibadah []Ibadah
	db.Find(&ibadah)
	return ibadah
}
func GetIbadahbyId(reg int64) (*Ibadah, *gorm.DB) {
	var getibadah Ibadah
	db := db.Where("ID=?", reg).Find(&getibadah)
	return &getibadah, db
}
func DeleteIbadah(reg int64) Ibadah {
	var ibadah Ibadah
	db.Where("ID=?", reg).Delete(ibadah)
	return ibadah
}
