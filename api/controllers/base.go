package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/yuanyu90221/golang_jwt_api_server/api/models"
)

//Server server
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

//Initialize basic init
func (server *Server) Initialize(Dbdriver string, intialer Initialer) {
	var err error
	DBURL := intialer.initialize(os.Getenv("DB_USER"), os.Getenv("DB_PASSWD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}
	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{})
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

//Run basic server
func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port: %s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
