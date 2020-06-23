package model

import "time"

type Contact struct {
	ID string `gorm:"primary_key" json:"id"`
	Name string `gorm:"type:varchar(255);NOT NULL" json:"name" binding:"required"`
	Email string `gorm:"type:varchar(255);NOT NULL" json:"email" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Contacts []Contact