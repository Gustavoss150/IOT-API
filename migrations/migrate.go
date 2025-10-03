package migrations

import (
	"api/config"
	"api/entities"
	"log"
)

func Migrate() {
	err := config.DB.AutoMigrate(
		&entities.User{},
		&entities.Equipment{},
		&entities.Reservation{},
		&entities.AccessKey{},
		&entities.EventLog{},
		&entities.BotConfig{},
	)

	if err != nil {
		log.Fatal("error migrating: " + err.Error())
	}
	log.Println("migrations executed successfully")
}
