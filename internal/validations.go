package internal

import "regexp"

// IsValidSubdomain returns true if domain is valid
func IsValidSubdomain(domain string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9-]+\.([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}(/[^\s]*)?$`)

	return re.MatchString(domain)
}
