package internal

const (
	SubDomainInputMessage      = "Please enter the domain you want to validate:"
	ValidationTypeInputMessage = `
Please select validation type:
D: for DNS and SSL
U: for Universal Link and SSL Validation

Validation Type:`
	OSSelectionInputMessage = `
Please enter (1, 2 or 3) for what OS you want to validate:
1: For Apple iOS
2: For Android
3: For both
Option:`

	FailureMessageInvalidSubdomainInput = "FAILURE: domain '%s' is not a valid subdomain!"
	FailureMessageLinkBranding          = "FAILURE: One or more of required steps are not valid, link branding will not work, unless you meant to run it for Universal Links!"
	FailureMessageUniversalLink         = "FAILURE: One or more of required steps are not valid, link branding will not work!"

	ErrorMessage = "ERROR: %v"

	WarningMessageInvalidValidationType = "WARNING: Unknown option '%s' was selected, please type D: for DNS U: for Universal Link"
	WarningMessageInvalidOsType         = "WARNING: Uknown option '%s' was selected of OS type. Please enter 1: for iOS, 2: for Android, 3: for both"

	SuccessMessageDNSValidation = "INFO: successfully validated DNS records"
	SuccessMessageSSLValidation = "INFO: successfully validated SSL Certificate for '%s'"
	SuccessMessageLinkBranding  = "SUCCESS: Link branding should function as expected!"
	SuccessMessageAssetLink     = "INFO: successfully validated Asset files for '%s'"
	SuccessMessageUniversal     = "SUCCESS: Universal Links should function as expected!"
)
