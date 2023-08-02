package main

import (
	"net/http"
	"time"

	"github.com/gernhan/template/api"
	"github.com/gernhan/template/db"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database connection pool.
	connString := "postgres://postgres:1@localhost/postgres"
	err := db.InitDB(connString)
	if err != nil {
		panic(err)
	}
	defer db.CloseDB()

	router := mux.NewRouter()

	// Register the "/api/insert-batch" route to the InsertBatchHandler function.
	router.HandleFunc("/api/insert-batch", api.InsertBatchHandler).Methods("POST")
	// Register the "/api/update-batch" route to the SpecialBatchUpdateHandler function.
	router.HandleFunc("/api/update-batch", api.SpecialBatchUpdateHandler).Methods("PUT")
	// Register the "/api/normal-update-batch" route to the SpecialBatchUpdateHandler function.
	router.HandleFunc("/api/normal-update-batch", api.NormalUpdateHandler).Methods("PUT")

	server := &http.Server{
		Addr:         ":9872",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
