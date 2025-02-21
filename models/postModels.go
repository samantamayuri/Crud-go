package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string
	Body     string
	Comments []Comment `gorm:"foreignKey:PostId;constraint:OnDelete:CASCADE;"`
}
