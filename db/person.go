package db

import (
	"fmt"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name  string
	Score float64
}

func (db *DB) GetPeople() ([]Person, error) {
	var people []Person
	result := db.db.Find(&people)

	if result.Error != nil {
		return nil, result.Error
	}

	return people, nil
}

func (db *DB) AddPerson(person Person) error {
	return db.db.Create(&person).Error
}

func (db *DB) UpdatePerson(person Person) error {
	return db.db.Model(&person).Updates(person).Error
}

func (db *DB) DeletePerson(id string) error {
	fmt.Println(id)
	return db.db.Delete(&Person{}, id).Error
}
