package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ProviderService service

type RequestProfilerProfileRegisterService struct {
	Name string `json:"name,omitempty"` //
}

//RegisterService 🚧 ServiceRegister
/* 🚧 ServiceRegister

 */
func (s *ProviderService) RegisterService() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/control/ServiceRegister"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation RegisterService")
	}

	return response, err

}

//UnregisterService 🚧 ServiceUnregister
/* 🚧 ServiceUnregister

 */
func (s *ProviderService) UnregisterService() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/control/ServiceUnregister"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation UnregisterService")
	}

	return response, err

}

//ReregisterService 🚧 ServiceReregister
/* 🚧 ServiceReregister

 */
func (s *ProviderService) ReregisterService() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/control/ServiceReregister"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation ReregisterService")
	}

	return response, err

}

//Authorization 🚧 Authorization
/* 🚧 Authorization

 */
func (s *ProviderService) Authorization() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/control/Authorization"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation Authorization")
	}

	return response, err

}
