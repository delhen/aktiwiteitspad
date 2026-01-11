package models

import "github.com/google/uuid"

type User struct {
	ID   		 		uuid.UUID 					`json:"id" gorm:"type:uuid;primaryKey"`
	Nickname 		string 							`json:"nickname" validate:"required,max=32,min=8`
	Email 			string 							`json:"-" validate:"required,email,max=255"`
	Password		string 							`json:"-" validate:"required,min=8`
	Activities	[]ActivityTracker		`gorm:"foreignKey:UserId"`
}
