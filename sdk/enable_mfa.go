package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type EnableMFAService service

// # Review unknown case

type RequestEnableMFAEnableMFA struct {
	Mfa *RequestEnableMFAEnableMFAMfa `json:"mfa,omitempty"` // Duo-MFA feature config
}

type RequestEnableMFAEnableMFAMfa struct {
	Status *bool `json:"status,omitempty"` // Enable or Disable MFA
}

//EnableMFA Enable MFA feature
/* Enable or Disable MFA feature

 */
func (s *EnableMFAService) EnableMFA(requestEnableMFAEnableMFA *RequestEnableMFAEnableMFA) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-mfa/enable"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEnableMFAEnableMFA).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation EnableMFA")
	}

	getCSFRToken(response.Header())

	return response, err

}
