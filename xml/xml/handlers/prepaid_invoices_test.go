package xml_handlers

import (
	"fmt"
	"os"
	"testing"

	"github.com/gernhan/xml/caches"
	"github.com/gernhan/xml/db"
	"github.com/gernhan/xml/entities/views"
	"github.com/gernhan/xml/pool"
	"github.com/gernhan/xml/utils"
)

func TestPrepareXmlData(t *testing.T) {
	// Initialize the database connection pool.
	connString := "postgresql://aax2tm:aax2tm@tmperf-db.int.compax.at:5433/tmperf"
	err := db.InitDB(connString)
	if err != nil {
		panic(err)
	}
	defer db.CloseDB()

	pool.InitPools()
	caches.InitCaches()

	content, err := getFileContent()
	if err != nil {
		t.Errorf("Caught error %v", err)
	}

	var bills []views.VExportingBillsV3
	err = utils.FromJSON(content, &bills) // Pass the pointer to the slice
	if err != nil {
		t.Errorf("Caught error %v", err)
	}

	handler := PrepaidInvoiceHandler{}
	data, err := handler.PrepareXmlData(bills)
	if err != nil {
		t.Errorf("Caught error %v", err)
	} else {
		fmt.Println(data)
	}
}

func getFileContent() (string, error) {
	// Replace "file.txt" with the path to your file
	filePath := "prepaid_bills.json"

	// Read the entire file into a byte slice
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}

	// Convert the byte slice to a string and return
	return string(content), nil
}
