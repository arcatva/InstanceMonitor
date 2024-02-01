package repository

import (
	"backend/model"
	"github.com/influxdata/influxdb-client-go/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var PostgresDB *gorm.DB
var InfluxDB influxdb2.Client

func Connect() {
	connectMysqlDB()
	connectInfluxDB()
}
func connectMysqlDB() {
	dsn := os.Getenv("POSTGRES_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	PostgresDB = db
	err = PostgresDB.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Database loaded")
	user := &model.User{
		Email:     "testadmin@admin.lab",
		FirstName: "testadmin",
		LastName:  "testadmin",
	}
	if err := user.SetPassword("pA1zsWO10FDsBd"); err != nil {
		log.Println("Failed to set password")
	}
	PostgresDB.Create(user)
	log.Println("Initialized admin user")
}

func connectInfluxDB() {
	url := os.Getenv("INFLUXDB_URL")
	token := os.Getenv("INFLUXDB_TOKEN")
	InfluxDB = influxdb2.NewClient(url, token)

}
