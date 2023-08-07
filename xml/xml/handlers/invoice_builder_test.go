package xml_handlers

import (
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestToMonetary(t *testing.T) {
	// Example usage
	amount := decimal.NewFromFloat(123.456)
	fmt.Println(toMonetary(&amount)) // Output: 123.46
}

func TestAmountString(t *testing.T) {
	// Example usage
	amount := decimal.NewFromFloat(123.456)
	fmt.Println(amountString(&amount)) // Output: 123.46
}
