package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// initialize sqlite
	db, err = InitializeSQLite()
	if err != nil {
		return err
	}

	return err
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// initialize logger
	logger = NewLogger(p)
	return logger
}