package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type IsMFAEnabledService service

type ResponseIsMFAEnabledIsMfaEnabled struct {
	Status *bool `json:"status,omitempty"` //
}

//IsMfaEnabled MFA feature enabled status
/* Is MFA feature enabled?

 */
func (s *IsMFAEnabledService) IsMfaEnabled() (*ResponseIsMFAEnabledIsMfaEnabled, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-mfa/status"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIsMFAEnabledIsMfaEnabled{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation IsMfaEnabled")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIsMFAEnabledIsMfaEnabled)
	return result, response, err

}
