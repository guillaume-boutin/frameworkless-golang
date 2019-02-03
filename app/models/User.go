package models

import (
	"github.com/guillaume-boutin/frameworkless-golang/app/resources"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"size:255;UNIQUE;NOT NULL"`
	HashPassword string `gorm:"size:255;NOT NULL"`
}

func (u User) Resource() resources.UserResource {
	return resources.UserResource{
		ID:        u.ID,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}
