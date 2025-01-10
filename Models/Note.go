package Models

import (
	"first-api/Config"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func GetAllNote(notes *[]Note) (err error) {
	if err = Config.DB.Where("isDelete = ?", false).Find(notes).Error; err != nil {
		return err
	}
	return nil
}

func CreateNote(note *Note) (err error) {
	if err = Config.DB.Create(note).Error; err != nil {
		return err
	}
	return nil
}

func GetNoteById(note *Note, id int) (err error) {
	if err = Config.DB.Where("id = ?", id).First(note).Error; err != nil {
		return err
	}
	return nil
}

func UpdateNote(note *Note, id int) (err error) {
	fmt.Println(note)
	Config.DB.Save(note)
	return nil
}

func DeleteNote(note *Note, id int) (err error) {
	Config.DB.Where("id = ?", id).First(note)
	note.IsDelete = true
	Config.DB.Save(note)
	return nil
}
