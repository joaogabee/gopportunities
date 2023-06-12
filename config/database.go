package config

import (
	"github.com/joaogabee/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func InitializerSQL() (*gorm.DB, error) {
	logger := GetLogger("DB")
	dbpath := "./db/main.db"
	_, err := os.Stat(dbpath)
	if os.IsNotExist(err) {
		logger.Info("database path not found, creating...")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbpath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		logger.Err("SQLITE: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Err("SQLITE: failed migrations %v", err)
		return nil, err
	}
	return db, nil
}
