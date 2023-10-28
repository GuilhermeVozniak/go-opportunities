package config

import (
	"os"

	"github.com/GuilhermeVozniak/go-opportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	var (
		db     *gorm.DB
		err    error
		logger *Logger
		dbPath = "./db/opportunity.db"
	)

	// chek if db is already exists
	_, err = os.Stat(dbPath)
	if os.IsNotExist(err) {
		// create db file
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("sqlite db folder creation error %v", err)
			return db, err
		}

		// create db file
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("sqlite db creation error %v", err)
			return db, err
		}

		file.Close()
	}

	// initialize logger
	logger = NewLogger("sqlite")

	// create sqlite database and connect
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite connection error %v", err)
		return db, err
	}

	// migrate the schema
	err = db.AutoMigrate(&schemas.Opportunity{})
	if err != nil {
		logger.Errorf("sqlite migration error %v", err)
		return db, err
	}

	return db, err
}
