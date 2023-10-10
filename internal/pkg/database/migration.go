package database

import (
	"github.com/bigusbeckus/quorum-challenge-backend/internal/pkg/database/models"
)

func RunMigrations() error {
	db := Connect()

	db.AutoMigrate(&models.Brand{})

	return nil
}
