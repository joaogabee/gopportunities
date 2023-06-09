package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitializerSQL()
	if err != nil {
		return fmt.Errorf("failed to initialize sql: %v", err)
	}
	return nil
}
func GetDB() *gorm.DB {
	return db
}
func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
