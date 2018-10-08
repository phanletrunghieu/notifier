package pg

import (
	"github.com/jinzhu/gorm"
	"github.com/phanletrunghieu/notifier/domain"
)

// New create new postgres database connection
// it return postgres connection and cleanup postgres connection
func New(ds string) (*gorm.DB, func()) {
	db, err := gorm.Open("postgres", ds)
	if err != nil {
		panic(err)
	}

	return db, func() {
		db.Close()
	}
}

// MigrateTables .
func MigrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
		domain.User{},
		domain.Device{},
	).Error
}
