package models

import (
	"BAPTIS_SERVICE/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Baptis struct {
	gorm.Model
	Nama_Lengkap   string `json:"nama_lengkap"`
	Tempat_Lahir   string `json:"tempat_lahir"`
	Tanggal_Lahir  string `json:"tanggal_lahir"`
	Nama_Ayah      string `json:"nama_ayah"`
	Nama_Ibu       string `json:"nama_ibu"`
	Tanggal_Baptis string `json:"tanggal_baptis"`
	Pendeta        string `json:"pendeta"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Baptis{})
}
func (b *Baptis) CreateBaptis() *Baptis {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllBaptis() []Baptis {
	var baptis []Baptis
	db.Find(&baptis)
	return baptis
}
func GetBaptisbyId(reg int64) (*Baptis, *gorm.DB) {
	var getBaptis Baptis
	db := db.Where("ID=?", reg).Find(&getBaptis)
	return &getBaptis, db
}
func DeleteBaptis(reg int64) Baptis {
	var baptis Baptis
	db.Where("ID=?", reg).Delete(baptis)
	return baptis
}
