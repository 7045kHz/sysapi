package main

import (
	"github.com/7045kHz/sysapi/controllers"
	"github.com/7045kHz/sysapi/driver"
	"github.com/7045kHz/sysapi/utils"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()

	controller := controllers.Controller{}

	router.HandleFunc("/systems", utils.TokenVerifyMiddleWare(controller.GetSystems(db))).Methods("GET")
	router.HandleFunc("/sysupdate", utils.TokenVerifyMiddleWare(controller.UpdateSystems(db))).Methods("POST")
	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
