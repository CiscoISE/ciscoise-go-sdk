package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ConsumerService service

type RequestConfigurationCreateAccount struct {
	NodeName string `json:"nodeName,omitempty"` //
}

type RequestConfigurationActivateAccount struct {
	Description string `json:"description,omitempty"` //
}

type RequestConfigurationLookupService struct {
	Name string `json:"name,omitempty"` //
}

type RequestConfigurationAccessSecret struct {
	PeerNodeName string `json:"peerNodeName,omitempty"` //
}

//CreateAccount 🚧 AccountCreate
/* 🚧 AccountCreate

 */
func (s *ConsumerService) CreateAccount(requestConfigurationCreateAccount *RequestConfigurationCreateAccount) (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/control/AccountCreate"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationCreateAccount).
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
func (s *ConsumerService) ActivateAccount(requestConfigurationActivateAccount *RequestConfigurationActivateAccount) (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/control/AccountActivate"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationActivateAccount).
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
func (s *ConsumerService) LookupService(requestConfigurationLookupService *RequestConfigurationLookupService) (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/control/ServiceLookup"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationLookupService).
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
func (s *ConsumerService) AccessSecret(requestConfigurationAccessSecret *RequestConfigurationAccessSecret) (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/control/AccessSecret"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationAccessSecret).
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
