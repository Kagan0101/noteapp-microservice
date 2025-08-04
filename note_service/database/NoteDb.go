package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"NoteApp-microservice/note_service/models"
)

var (
	db *gorm.DB
	err error
	N models.Note
)

func Migrate() {
	db, err = gorm.Open(mysql.Open(dns),&gorm.Config{})
	if err != nil{
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&N)
}



func AddNote(note *models.Note) error {
    return db.Create(note).Error
}


func DeleteNote() {
	db.Delete(&N,N.ID)
}

func GetNote(where ...interface{}) models.Note {
	db.First(&N,where...)
	return N
}

func GetAllNotes(where ...interface{}) []models.Note {
	var notes []models.Note
	db.Find(&notes,where...)
	return notes
}

func UptadeNote(column string,value interface{}) {
	db.Model(&N).Update(column,value)
}

func UpdateNotes(data models.Note) {
	db.Model(&N).Updates(data)
}