package database

import (
	"database/sql"

	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	ChatID   int64
	Language string
}
type Product struct {
	gorm.Model
	Code          string
	NameEN        string
	NameFA        string
	NameAR        string
	DescriptionEN string
	DescriptionFA string
	DescriptionAR string
}
