package global

import (
	"gorm.io/gorm"
	"uatbo-backend/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
