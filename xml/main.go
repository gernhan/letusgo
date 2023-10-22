package main

import (
	"github.com/gernhan/xml/caches"
	sftputils "github.com/gernhan/xml/sftp"
	xml_handlers "github.com/gernhan/xml/xml/handlers"
	"net/http"
	"time"

	"github.com/gernhan/xml/api"
	"github.com/gernhan/xml/db"
	"github.com/gernhan/xml/pool"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database connection pool.
	connString := "postgresql://aax2tm:aax2tm@tmperf-db.int.compax.at:5433/tmperf"
	err := db.InitDB(connString)
	if err != nil {
		panic(err)
	}
	defer db.CloseDB()

	pool.InitPools()
	caches.InitCaches()

	xml_handlers.InitPrepaidHandler()
	xml_handlers.InitNormalHandler()

	sftputils.Init()

	router := mux.NewRouter()

	// Register apis
	router.HandleFunc("/invoice-interface/actuator/health/liveness", api.HealCheck).Methods("GET")
	router.HandleFunc("/invoice-interface/actuator/health/readiness", api.HealCheck).Methods("GET")

	router.HandleFunc("/invoice-interface/xml/{billRunId}/gen", api.XmlHandler).Methods("POST")
	router.HandleFunc("/invoice-interface/test/xml/update-bill-statuses", api.ResetStatusesHandler).Methods("POST")
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
