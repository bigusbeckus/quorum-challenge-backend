package database

import (
	"fmt"

	"github.com/bigusbeckus/quorum-challenge-backend/internal/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Initialize() error {
	config := config.AppConfig.Database
	dsn := GenerateDsn(
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
		config.EnableSSL,
	)

	var err error
	c := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
		// Logger: logger.Default.LogMode(logger.Silent),
	}

	db, err = gorm.Open(postgres.Open(dsn), c)
	if err != nil || db == nil {
		return err
	}

	return nil
}

func GenerateDsn(
	host string,
	user string,
	password string,
	dbName string,
	port uint16,
	enableSSL bool,
) string {
	ssl := "require"
	if !enableSSL {
		ssl = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		host,
		user,
		password,
		dbName,
		port,
		ssl,
	)
	return dsn
}

func Connect() *gorm.DB {
	return db
}
