package database

func RunExtensions() error {
	db := Connect()
	return db.Exec(`CREATE EXTENSION IF NOT EXISTS pg_trgm;`).Error
}
