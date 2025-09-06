package migrations

import (
	"api/config"
	"api/entities"
	"log"
)

func Migrate() {
	err := config.DB.AutoMigrate(
		&entities.User{},
		&entities.Machine{},
		&entities.Reservation{},
	)

	if err != nil {
		log.Fatal("error migrating: " + err.Error())
	}
	log.Println("migrations executed successfully")
}
