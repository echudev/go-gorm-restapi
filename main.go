package main

import (
	"net/http"

	"github.com/echudev/go-gorm-restapi/db"
	"github.com/echudev/go-gorm-restapi/models"
	"github.com/echudev/go-gorm-restapi/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter() // declaro un nuevo router utilizando mux

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
