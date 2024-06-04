package models

import (
	"REGISTRASI_JEMAAT_SERVICE/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Jemaat struct {
	gorm.Model
	No_Registrasi_Jemaat string `json:"no_registrasi_jemaat"`
	Nama_Lengkap         string `json:"nama_lengkap"`
	Nama_Keluarga        string `json:"nama_keluarga"`
	Tempat_Lahir         string `json:"tempat_lahir"`
	Tanggal_Lahir        string `json:"tanggal_lahir"`
	Alamat               string `json:"alamat"`
	Jenis_Kelamin        string `json:"jenis_kelamin"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Jemaat{})
}
func (b *Jemaat) CreateJemaat() *Jemaat {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllJemaat() []Jemaat {
	var jemaat []Jemaat
	db.Find(&jemaat)
	return jemaat
}
func GetJemaatbyId(reg int64) (*Jemaat, *gorm.DB) {
	var getJemaat Jemaat
	db := db.Where("ID=?", reg).Find(&getJemaat)
	return &getJemaat, db
}
func DeleteJemaat(reg int64) Jemaat {
	var jemaat Jemaat
	db.Where("ID=?", reg).Delete(jemaat)
	return jemaat
}
