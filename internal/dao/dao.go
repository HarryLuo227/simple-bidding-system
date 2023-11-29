package dao

import (
	"github.com/jinzhu/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	engine.Begin()
	return &Dao{engine: engine}
}
