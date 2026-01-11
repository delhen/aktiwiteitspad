package models

import "github.com/google/uuid"

type ActivityTracker struct {
	ID   		 			uuid.UUID 	`json:"id" gorm:"type:uuid;primaryKey"`
	Title 				string 			`json:"nickname" validate:"required,max=32,min=8`
	Description		string 			`json:"-" validate:"required,email,max=255"`
	Image					string 			`json:"-" validate:"required,min=8`
	UserId				string
	User   				User 				`gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
