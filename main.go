package main

import (
	"domain-validator/internal"
	"fmt"

	"github.com/fatih/color"
)

func main() {
subdomainInput:
	color.New(color.FgBlue).Print(internal.SubDomainInputMessage)
	var subdomain string
	fmt.Scanln(&subdomain)

	// validate subdomain
	if !internal.IsValidSubdomain(subdomain) {
		color.Red(internal.FailureMessageInvalidSubdomainInput, subdomain)
		goto subdomainInput
	}

typeSelection:
	color.New(color.FgBlue).Print(internal.ValidationTypeInputMessage)
	var validationType string
	fmt.Scanln(&validationType)

	switch validationType {
	case "D":
		validateDNS(subdomain)
	case "U":
		validateUniversalLink(subdomain)
	default:
		color.Yellow(internal.WarningMessageInvalidValidationType, validationType)
		goto typeSelection
	}
}

func validateDNS(subdomain string) {
	var hasErrors bool
	if err := internal.ValidateDNSRecords(subdomain); err != nil {
		color.Red(internal.ErrorMessage, err)
		hasErrors = true
	} else {
		color.Green(internal.SuccessMessageDNSValidation)
	}

	if err := internal.ValidateSSLCertificate(subdomain); err != nil {
		color.Red(internal.ErrorMessage, err)
		hasErrors = true
	} else {
		color.Green(internal.SuccessMessageSSLValidation, subdomain)
	}

	if !hasErrors {
		color.Green(internal.SuccessMessageLinkBranding)
	} else {
		color.Red(internal.FailureMessageLinkBranding)
	}
}

func validateUniversalLink(subdomain string) {
osSelection:
	color.New(color.FgBlue).Print(internal.OSSelectionInputMessage)
	var osType string
	fmt.Scanln(&osType)

	switch osType {
	case "1":
	case "2":
	case "3":
	default:
		color.Yellow(internal.WarningMessageInvalidOsType, osType)
		goto osSelection
	}

	var hasErrors bool
	if err := internal.ValidateSSLCertificate(subdomain); err != nil {
		color.Red(internal.ErrorMessage, err)
		hasErrors = true
	} else {
		color.Green(internal.SuccessMessageSSLValidation, subdomain)
	}

	if message, err := internal.ValidateUniversalLink(subdomain, osType); err != nil {
		color.Red(internal.ErrorMessage, err)
	} else {
		color.Green(internal.SuccessMessageAssetLink, message)
	}

	if !hasErrors {
		color.Green(internal.SuccessMessageUniversal)
	} else {
		color.Red(internal.FailureMessageUniversalLink)
	}
}
