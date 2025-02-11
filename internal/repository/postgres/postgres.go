package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(strConn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(strConn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
