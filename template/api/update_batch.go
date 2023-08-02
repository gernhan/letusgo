package api

import (
	"encoding/json"
	"fmt"
	"github.com/gernhan/template/db"
	"github.com/gernhan/template/entities"
	"github.com/gernhan/template/tools"
	"net/http"
	"strconv"
)

type UserCountRequest struct {
	UserCount int `json:"userCount"`
	BatchSize int `json:"batchSize"`
}

func getAboutUpdatingUsers(userCount int) []entities.User {
	inputGenerator := tools.NewInputGenerator()
	users := make([]entities.User, userCount)

	for i := 1; i <= userCount; i++ {
		user := entities.User{}
		user.Username = inputGenerator.GenerateString(10)
		user.Password = inputGenerator.GenerateString(20) + "_updated"
		user.ID = json.Number(strconv.Itoa(i))
		users[i-1] = user
	}
	return users
}

func NormalUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody UserCountRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	users := getAboutUpdatingUsers(requestBody.UserCount)

	// Get the batchSize from the request body or use a default value (e.g., 1000).
	batchSize := requestBody.BatchSize
	if batchSize <= 0 {
		batchSize = 1000
	}

	err := db.UpdateAllUsers(users, batchSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message or HTTP status code.
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("Batch update completed successfully.")
}

func SpecialBatchUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody UserCountRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	users := getAboutUpdatingUsers(requestBody.UserCount)

	// Get the batchSize from the request body or use a default value (e.g., 1000).
	batchSize := requestBody.BatchSize
	if batchSize <= 0 {
		batchSize = 1000
	}

	// Call the function to perform the batch update.
	err := db.BatchUpdateStatic(users, batchSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message or HTTP status code.
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Updated %d users in the database", len(users))
}
