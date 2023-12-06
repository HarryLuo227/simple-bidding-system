package global

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/jinzhu/gorm"
)

var (
	DBEngine  *gorm.DB
	RedisSync *redsync.Redsync
)
