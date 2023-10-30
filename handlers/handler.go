package handlers

import (
	"github.com/GuilhermeVozniak/go-opportunities/config"
	"gorm.io/gorm"
)

var (
	logger            *config.Logger
	db                *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handlers")
	db = config.GetSQLite()
}
