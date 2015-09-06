package models

// Migrate the database
func Migrate() {
	DB.AutoMigrate(&User{})
}
