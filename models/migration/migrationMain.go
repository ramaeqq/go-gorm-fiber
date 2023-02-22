package migration

import (
	"fmt"
	"log"
	"ramaeqq/go-fiber-gorm/models/database"
	"ramaeqq/go-fiber-gorm/models/entities"
)

func DbMigration() {

	err := database.DB.AutoMigrate(&entities.Person{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("DB migrated success")
}
