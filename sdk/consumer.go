package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ConsumerService service

type RequestClearThreatsAndVulnerabilitiesCreateAccount struct {
	NodeName string `json:"nodeName,omitempty"` //
}

type RequestClearThreatsAndVulnerabilitiesActivateAccount struct {
	Description string `json:"description,omitempty"` //
}

type RequestClearThreatsAndVulnerabilitiesLookupService struct {
	Name string `json:"name,omitempty"` //
}

type RequestClearThreatsAndVulnerabilitiesAccessSecret struct {
	PeerNodeName string `json:"peerNodeName,omitempty"` //
}


//CreateAccount 🚧 AccountCreate
/* 🚧 AccountCreate

 */
func (s *ConsumerService) CreateAccount(requestClearThreatsAndVulnerabilitiesCreateAccount *RequestClearThreatsAndVulnerabilitiesCreateAccount) (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/control/AccountCreate"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestClearThreatsAndVulnerabilitiesCreateAccount).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateAccount")
	}

	return response, err

}

//ActivateAccount 🚧 AccountActivate
/* 🚧 AccountActivate

 */
func (s *ConsumerService) ActivateAccount(requestClearThreatsAndVulnerabilitiesActivateAccount *RequestClearThreatsAndVulnerabilitiesActivateAccount) (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/control/AccountActivate"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestClearThreatsAndVulnerabilitiesActivateAccount).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation ActivateAccount")
	}

	return response, err

}

//LookupService 🚧 ServiceLookup
/* 🚧 ServiceLookup

 */
func (s *ConsumerService) LookupService(requestClearThreatsAndVulnerabilitiesLookupService *RequestClearThreatsAndVulnerabilitiesLookupService) (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/control/ServiceLookup"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestClearThreatsAndVulnerabilitiesLookupService).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation LookupService")
	}

	return response, err

}

//AccessSecret 🚧 AccessSecret
/* 🚧 AccessSecret

 */
func (s *ConsumerService) AccessSecret(requestClearThreatsAndVulnerabilitiesAccessSecret *RequestClearThreatsAndVulnerabilitiesAccessSecret) (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/control/AccessSecret"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestClearThreatsAndVulnerabilitiesAccessSecret).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation AccessSecret")
	}

	return response, err

}
