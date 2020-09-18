package api

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/yuanyu90221/golang_jwt_api_server/api/controllers"
	"github.com/yuanyu90221/golang_jwt_api_server/api/seed"
)

var server = controllers.Server{}

//Run main logic
func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	var dbInitialer controllers.Initialer
	switch os.Getenv("DB_DRIVER") {
	case "postgres":
		dbInitialer = &controllers.PostgresInitialer{}
	default:
		dbInitialer = &controllers.MysqlInitialer{}
	}
	server.Initialize(os.Getenv("DB_DRIVER"), dbInitialer)
	seed.Load(server.DB)

	server.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
