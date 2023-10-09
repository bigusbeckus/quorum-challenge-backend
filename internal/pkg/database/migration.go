package database

import (
	"github.com/bigusbeckus/quorum-challenge-backend/internal/pkg/database/models"
)

func RunMigrations() error {
	db := Connect()

	if err := db.AutoMigrate(&models.Brand{}); err != nil {
		return err
	}

	return nil
}
