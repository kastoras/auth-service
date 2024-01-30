package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDB(dbUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
