package models

import (
	_ "github.com/jinzhu/gorm"
	"time"
)

type Todo struct {
	ID          uint   `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Description string `json:"description"`
	Done        bool   `json:"done" gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `json:"-"`
}

type TodoModel struct{}

func (TodoModel) All() []Todo {
	t := []Todo{}
	db.Find(&t)

	return t
}

func (TodoModel) Find(id uint) (Todo, error) {
	t := Todo{}
	db.First(&t, id)

	if db.Error != nil {
		return t, db.Error
	}

	return t, nil
}

func (TodoModel) New(description string, done bool) (uint, error) {
	t := Todo{Description: description, Done: done}
	result := db.Create(&t)

	if result.Error != nil {
		return 0, result.Error
	}

	return t.ID, nil
}

func (TodoModel) Update(id uint, description string, done bool) (int64, error) {
	t := Todo{}
	result := db.First(&t, id)

	if result.Error != nil {
		return 0, result.Error
	}

	t.Description = description
	t.Done = done
	result = db.Save(&t)

	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}

func (TodoModel) Destroy(id uint) int64 {
	t := Todo{ID: id}
	result := db.Delete(&t)

	return result.RowsAffected
}
