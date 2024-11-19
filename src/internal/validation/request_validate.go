package validation

import "regexp"

// IsUuid - Check if the string is a valid UUID.
func IsUuid(value string) bool {
	regex, err := regexp.Compile("^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$")
	if err != nil {
		return false
	}

	// Check if the string is a valid UUID
	return regex.MatchString(value)
}
