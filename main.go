package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var globalDB *gorm.DB
var globalGameRepository *SqliteGameRepository

func main() {
	globalDB, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
	}

	err = globalDB.AutoMigrate(
		&GameRegistry{},
		&Authentication{},
	)

	if err != nil {
		panic("Failed to migrate database")
	}

	globalGameRepository = NewSqliteGameRegistryRepository(globalDB)

	err = globalDB.AutoMigrate(&GameRegistry{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate: %v", err))
	}

	NewSqliteGameRegistryRepository(globalDB)
}
