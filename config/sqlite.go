package config

import (
	"os"

	"github.com/joaogabsoaresf/go-stock-watcher/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// check DB exist
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("db file not found, creating now...")
		// create DB file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create((dbPath))
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// create DB
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite start error: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Lead{})
	if err != nil {
		logger.Errorf("sqlite automigrate error: %v", err)
		return nil, err
	}
	// return DB
	return db, nil
}
