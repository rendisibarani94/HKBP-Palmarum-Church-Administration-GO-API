package models

import (
	"EVENT_SERVICE/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Event struct {
	gorm.Model
	Nama_Event    string `json:"nama_event"`
	Tempat_Event  string `json:"tempat_event"`
	Tanggal_Event string `json:"tanggal_event"`
	Pelayan_Event string `json:"pelayan_event"`
	Jumlah_Hadir  string `json:"jumlah_hadir"`
	Keterangan    string `json:"keterangan"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Event{})
}
func (b *Event) CreateEvent() *Event {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllEvent() []Event {
	var event []Event
	db.Find(&event)
	return event
}
func GetEventbyId(reg int64) (*Event, *gorm.DB) {
	var getEvent Event
	db := db.Where("ID=?", reg).Find(&getEvent)
	return &getEvent, db
}
func DeleteEvent(reg int64) Event {
	var event Event
	db.Where("ID=?", reg).Delete(event)
	return event
}
