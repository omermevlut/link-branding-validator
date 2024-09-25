package internal

import (
	"crypto/tls"
	"fmt"
)

// ValidateSSLCertificate returns error if SSL validation fails
func ValidateSSLCertificate(subdomain string) error {
	_, err := tls.Dial("tcp", fmt.Sprintf("%s:443", subdomain), nil)
	if err != nil {
		return fmt.Errorf("failed to validate SSL cerficate for '%s', error %v", subdomain, err)
	}

	return nil
}
