package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// DNSResponseAnswer DNS response answer section
type DNSResponseAnswer struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (ans *DNSResponseAnswer) dataContains(value string) bool {
	return strings.Contains(ans.Data, value)
}

// DNSQueryResponse response acquired from Google DNS resolve
type DNSQueryResponse struct {
	Status int                 `json:"Status"`
	Answer []DNSResponseAnswer `json:"Answer"`
}

func (dqr *DNSQueryResponse) dataContains(value string) bool {
	for _, ans := range dqr.Answer {
		if ans.dataContains(value) {
			return true
		}
	}

	return false
}

// ValidateDNSRecords returns error if valdation fails
func ValidateDNSRecords(subdomain string) error {
	res, err := http.Get(fmt.Sprintf("https://dns.google/resolve?name=%s", subdomain))
	if err != nil {
		return fmt.Errorf("failed to fetch DNS information for %s, error %v", subdomain, err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response from Google DNS Resolver for %s, error %v", subdomain, err)
	}

	var dnsResponse DNSQueryResponse
	if err := json.Unmarshal(body, &dnsResponse); err != nil {
		return fmt.Errorf("failed to parse DNS response data from Google, error %v", err)
	}

	if dnsResponse.dataContains("sendgrid.net") {
		return nil
	}

	return fmt.Errorf("required value 'sendgrid.net' was not found in DNS Answer section")
}
