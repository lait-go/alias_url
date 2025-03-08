package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"retsAPI/serv/config"
	"retsAPI/serv/logger"
	reading "retsAPI/serv/processing"
	"retsAPI/serv/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.NewConfig()
	log := logger.NewLogger(cfg.Env)

	log.Debug("LOGGER_STARTED")

	err := storage.FileExists(cfg.StoragePath)
	if err != nil{
		panic(err)
	}
	log.Debug("STORAGE_EXIST")
	fmt.Println(cfg)
	storage.StorageWork("asf", "assss")
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(logger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/", reading.ReadRequest())
	
	

	http.ListenAndServe(cfg.HTTPServer.Address, router)
	//run server
}