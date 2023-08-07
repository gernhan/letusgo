package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gernhan/xml/db"
	"github.com/gernhan/xml/entities"
	"github.com/gernhan/xml/tools"
)

func getUsers(userCount int) []entities.User {
	inputGenerator := tools.NewInputGenerator()
	users := make([]entities.User, 0, userCount)

	for i := 0; i < userCount; i++ {
		user := entities.User{}
		user.Username = inputGenerator.GenerateString(10)
		user.Password = inputGenerator.GenerateString(20)
		user.CreatedTime = time.Now()
		user.UpdatedTime = time.Now()
		user.UserType = "2"
		user.DateOfBirth = time.Now()

		users = append(users, user)
	}
	return users
}

func InsertBatchHandler(w http.ResponseWriter, r *http.Request) {
	type UserCountRequest struct {
		UserCount int `json:"userCount"`
		BatchSize int `json:"batchSize"`
	}

	var requestBody UserCountRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	users := getUsers(requestBody.UserCount)

	// Define a function to insert users asynchronously using a goroutine.
	insertUsers := func(users []entities.User, wg *sync.WaitGroup) {
		defer wg.Done()
		err := db.InsertBatchUsers(users)
		if err != nil {
			fmt.Printf("Error inserting users: %v\n", err)
		}
	}

	// Get the batchSize from the request body or use a default value (e.g., 1000).
	batchSize := requestBody.BatchSize
	if batchSize <= 0 {
		batchSize = 1000
	}

	// Chunk users into batches of batchSize users each and insert each batch using a goroutine.
	totalUsers := len(users)
	var wg sync.WaitGroup
	for i := 0; i < totalUsers; i += batchSize {
		end := i + batchSize
		if end > totalUsers {
			end = totalUsers
		}
		userBatch := users[i:end]
		wg.Add(1)
		go insertUsers(userBatch, &wg)
	}

	// Wait for all goroutines to finish before responding.
	wg.Wait()

	// Respond with a success message or HTTP status code.
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Finished batch insert for %d users with batch size %d", len(users), batchSize)
}


