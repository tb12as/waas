package database

import (
	"fmt"
	"time"

	"github.com/tb12as/why-as-a-service/internal/config"
	"github.com/tb12as/why-as-a-service/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Load(cfg *config.Config) (*gorm.DB, error) {
	username := cfg.DbUsername
	password := cfg.DbPassword
	dbName := cfg.DbName
	host := cfg.DbHost
	port := cfg.DbPort

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbName)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  false, // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    false, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   false, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get sql db: %w", err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&models.Reason{})

	return db, nil
}
