package database

import "github.com/bigusbeckus/quorum-challenge-backend/internal/pkg/database/models"

func Seed() error {
	db := Connect()
	tx := db.Begin()

	if err := models.SeedBrands(tx); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
