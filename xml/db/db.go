package db

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gernhan/xml/entities"
	"github.com/gernhan/xml/tools"
	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

func GetPool() *pgxpool.Pool {
	return pool
}

// InitDB initializes the database connection pool.
func InitDB(connString string) error {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return err
	}

	// Customize connection pool configuration.
	config.MaxConns = 20                       // Set maximum number of connections in the pool.
	config.MaxConnIdleTime = 5 * time.Minute   // Set the maximum idle time for a connection.
	config.HealthCheckPeriod = 1 * time.Minute // Set the interval between health checks.
	// Add more specialized configurations as needed.

	pool, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return err
	}

	return nil
}

// CloseDB closes the database connection pool.
func CloseDB() {
	pool.Close()
}

// InsertBatchUsers performs a batch insert of users into the "d_user" table.
func InsertBatchUsers(users []entities.User) error {
	// Prepare the SQL query string for the bulk insert.
	query := "INSERT INTO d_user (username, password, created_time, updated_time, user_type, date_of_birth) VALUES "
	valueStrings := make([]string, 0, len(users))
	valueArgs := make([]interface{}, 0, len(users)*6)

	for _, user := range users {
		valuePlaceHolders := fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)",
			len(valueArgs)+1, len(valueArgs)+2, len(valueArgs)+3, len(valueArgs)+4, len(valueArgs)+5, len(valueArgs)+6)
		valueStrings = append(valueStrings, valuePlaceHolders)
		valueArgs = append(valueArgs, user.Username, user.Password, user.CreatedTime, user.UpdatedTime, user.UserType, user.DateOfBirth)
	}

	query += strings.Join(valueStrings, ", ")

	// Execute the batch insert query using the connection pool.
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), query, valueArgs...)
	if err != nil {
		return err
	}

	return nil
}

func UpdateAllUsers(users []entities.User, batchSize int) error {
	startTime := time.Now()
	batches := 0
	updateCount := int64(0)

	var e error
	var wg sync.WaitGroup
	for i := 0; i < len(users); i += batchSize {
		batch := users[i:tools.Min(i+batchSize, len(users))]
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Call UpdateAllUsers for each batch
			err := UpdateBatch(batch)
			if err != nil {
				e = err
				log.Printf("Caught error when updating batch: %v", err)
			} else {
				atomic.AddInt64(&updateCount, int64(len(batch)))
			}
		}()
		batches++
	}

	wg.Wait()
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	log.Printf("normal batch update for %d users use %d batches", len(users), batches)
	log.Printf("normal batch update for %d batches of %d-user -> Total update count: %d in %v", batches, batchSize, updateCount, executionTime)
	return e
}

func UpdateBatch(users []entities.User) error {

	sql := "UPDATE d_user SET password=$1, username=$2, updated_time=$3 WHERE ID=$4"

	ctx := context.Background()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, u := range users {
		updatedAt := time.Now()
		_, err := tx.Exec(ctx, sql, u.Password, strings.ToUpper(u.Username), updatedAt, u.ID)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

// BatchUpdateStatic performs batch update on the list of users.
func BatchUpdateStatic(users []entities.User, batchUpdateSize int) error {
	startTime := time.Now()
	var wg sync.WaitGroup
	var updateCount int64

	ch := make(chan string, len(users))

	var e error
	for i := 0; i < len(users); i += batchUpdateSize {
		wg.Add(1)
		finalI := i
		go func() {
			err := batchUpdate(users[finalI:tools.Min(finalI+batchUpdateSize, len(users))], ch, &updateCount, &wg)
			if err != nil {
				log.Printf("Caught error: %v", err)
				e = err
			}
		}()
	}

	wg.Wait()

	if e != nil {
		return e
	}

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	batches := (len(users) + batchUpdateSize - 1) / batchUpdateSize
	log.Printf("batchUpdateStatic for %d users use %d batches", len(users), batches)
	log.Printf("batchUpdateStatic for %d batches of %d-user -> Total update count: %d in %v", batches, batchUpdateSize, updateCount, executionTime)
	return nil
}

func batchUpdate(users []entities.User, ch chan<- string, updateCount *int64, wg *sync.WaitGroup) error {
	defer wg.Done()

	query := buildUpdateQuery(users)
	ch <- query

	// Execute the batch insert query using the connection pool.
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	// Update the updateCount with the number of records updated by this batch
	atomic.AddInt64(updateCount, int64(len(users)))
	return nil
}

func buildUpdateQuery(users []entities.User) string {
	var query strings.Builder
	var setUserName strings.Builder
	var setPassword strings.Builder
	var setUpdatedTime strings.Builder

	query.WriteString("UPDATE d_user SET username = CASE ")

	for _, u := range users {
		setUserName.WriteString(fmt.Sprintf("WHEN ID='%s' THEN '%s' ", u.ID, strings.ToLower(u.Username)))
		setPassword.WriteString(fmt.Sprintf("WHEN ID='%s' THEN '%s' ", u.ID, strings.ToLower(u.Password)))
		setUpdatedTime.WriteString(fmt.Sprintf("WHEN ID='%s' THEN NOW() ", u.ID))
	}

	query.WriteString(setUserName.String())
	query.WriteString("END, PASSWORD = CASE ")
	query.WriteString(setPassword.String())
	query.WriteString("END, updated_time = CASE ")
	query.WriteString(setUpdatedTime.String())

	ids := make([]string, len(users))
	for i, u := range users {
		ids[i] = u.ID.String()
	}
	query.WriteString(fmt.Sprintf("END WHERE ID IN (%s)", strings.Join(ids, ",")))

	return query.String()
}
