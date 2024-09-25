package internal

import (
	"fmt"
	"net/http"
)

// ValidateUniversalLink returns error if universal link validation fails
func ValidateUniversalLink(subdomain string, osType string) (string, error) {
	if osType == "1" {
		if err := validateIOS(subdomain); err != nil {
			return "", err
		} else {
			return "Apple iOS", err
		}
	}

	if osType == "2" {
		if err := validateAndroid(subdomain); err != nil {
			return "", err
		} else {
			return "Android", nil
		}
	}

	if err := validateIOS(subdomain); err != nil {
		return "", err
	}

	return "Apple iOS and Android", validateAndroid(subdomain)
}

func validateIOS(subdomain string) error {
	res, err := http.Get(fmt.Sprintf("https://%s/.well-known/apple-app-site-association", subdomain))
	if err != nil {
		return fmt.Errorf("failed to get Apple Site Association (iOS), for %s, error %v", subdomain, err)
	}

	if res.StatusCode < http.StatusOK || res.StatusCode > http.StatusPermanentRedirect {
		return fmt.Errorf("failed to validate Applie Site Association (iOS), server responded with code: %d", res.StatusCode)
	}

	if res.Header.Get("content-type") != "application/json" {
		return fmt.Errorf("failed to validate response header for Apple Site Association (iOS), expected content-type: application/json, got: %s", res.Header.Get("content-type"))
	}

	return nil
}

func validateAndroid(subdomain string) error {
	res, err := http.Get(fmt.Sprintf("https://%s/.well-known/assetlinks.json", subdomain))
	if err != nil {
		return fmt.Errorf("failed to get Android Assetlink, for %s, error %v", subdomain, err)
	}

	if res.StatusCode < http.StatusOK || res.StatusCode > http.StatusPermanentRedirect {
		return fmt.Errorf("failed to validate Android Assetlink, server responded with code: %d", res.StatusCode)
	}

	if res.Header.Get("content-type") != "application/json" {
		return fmt.Errorf("failed to validate response header for Android Assetlink, expected content-type: application/json, got: %s", res.Header.Get("content-type"))
	}

	return nil
}
