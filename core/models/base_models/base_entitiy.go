package models

import (
	"gorm.io/gorm"
)

type BaseEntitiy struct {
	gorm.Model
	CreatedBy uint
	UpdatedBy uint
}
