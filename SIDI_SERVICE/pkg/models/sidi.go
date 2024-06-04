package models

import (
	"SIDI_SERVICE/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Sidi struct {
	gorm.Model
	Nama_Lengkap  string `json:"nama_lengkap"`
	Tempat_Lahir  string `json:"tempat_lahir"`
	Tanggal_Lahir string `json:"tanggal_lahir"`
	Tenggal_Sidi  string `json:"tenggal_sidi"`
	Lokasi_Sidi   string `json:"lokasi_sidi"`
	Nama_Pendeta  string `json:"nama_pendeta"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Sidi{})
}
func (b *Sidi) CreateSidi() *Sidi {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllSidi() []Sidi {
	var sidi []Sidi
	db.Find(&sidi)
	return sidi
}
func GetSidibyId(reg int64) (*Sidi, *gorm.DB) {
	var getSidi Sidi
	db := db.Where("ID=?", reg).Find(&getSidi)
	return &getSidi, db
}
func DeleteSidi(reg int64) Sidi {
	var sidi Sidi
	db.Where("ID=?", reg).Delete(sidi)
	return sidi
}
