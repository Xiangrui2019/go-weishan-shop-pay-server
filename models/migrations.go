package models

func MigrationModels() error {
	DB.AutoMigrate(&Fee{})
	DB.AutoMigrate(&Order{})
	return nil
}
