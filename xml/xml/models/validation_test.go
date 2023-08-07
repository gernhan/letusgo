package xml

import "testing"

func TestDateTypeValidation(t *testing.T) {
	validDates := []string{
		"31.12.2022",
		"01.01.2023",
	}

	invalidDates := []string{
		"2022-12-31",
		"31/12/2022",
		"12.31.2022",
		"32.12.2022", // Invalid day (32)
		"31.13.2022", // Invalid month (13)
		"31.12.22",   // Invalid year (short year)
	}

	for _, date := range validDates {
		if !ValidateDate(date) {
			t.Errorf("Validation failed for valid date: %s", date)
		}
	}

	for _, date := range invalidDates {
		if ValidateDate(date) {
			t.Errorf("Validation passed for invalid date: %s", date)
		}
	}
}

func TestTimeTypeValidation(t *testing.T) {
	validTimes := []string{
		"12:30:45",
		"23:59:59",
	}

	invalidTimes := []string{
		"12-30-45",
		"12:30",
		"12:30:61",
	}

	for _, time := range validTimes {
		if !ValidateTime(time) {
			t.Errorf("Validation failed for valid time: %s", time)
		}
	}

	for _, time := range invalidTimes {
		if ValidateTime(time) {
			t.Errorf("Validation passed for invalid time: %s", time)
		}
	}
}

func TestTaxRateTypeValidation(t *testing.T) {
	validTaxRates := []string{
		"19%",
		"7,7%",
	}

	invalidTaxRates := []string{
		"19",
		"7,7",
		"100%",
		"1,000%",
	}

	for _, taxRate := range validTaxRates {
		if !ValidateTaxRate(taxRate) {
			t.Errorf("Validation failed for valid tax rate: %s", taxRate)
		}
	}

	for _, taxRate := range invalidTaxRates {
		if ValidateTaxRate(taxRate) {
			t.Errorf("Validation passed for invalid tax rate: %s", taxRate)
		}
	}
}
