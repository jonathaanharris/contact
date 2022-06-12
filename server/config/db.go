package config

import (
	"api-contact/models"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// var (
// 	contactSeed = []models.Contact{
// 		{Name: "jon", PhoneNumber: "0811223344",Email:"jon@mail.com"},
// 		{Name: "Mil", PhoneNumber: "0813121110",Email:"Mil@mail.com"},
// 	}
// )

func ConnectDatabase() *gorm.DB {
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	db, err := gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Succesfully connected to Database!")
	}

	// defer db.Close()

	db.AutoMigrate(&models.Contact{})

	// for idx := range contactSeed {
	// 	db.Create(&contactSeed[idx])
	// }

	return db
}
