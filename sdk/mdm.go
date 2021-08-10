package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type MdmService service

//GetEndpoints 🚧 getEndpoints
/* 🚧 getEndpoints

 */
func (s *MdmService) GetEndpoints() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/ise/mdm/getEndpoints"

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
		return response, fmt.Errorf("error with operation GetEndpoints")
	}

	return response, err

}

//GetEndpointByMacAddress 🚧 getEndpointByMacAddress
/* 🚧 getEndpointByMacAddress

 */
func (s *MdmService) GetEndpointByMacAddress() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/ise/mdm/getEndpointByMacAddress"

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
		return response, fmt.Errorf("error with operation GetEndpointByMacAddress")
	}

	return response, err

}

//GetEndpointsByType 🚧 getEndpointsByType
/* 🚧 getEndpointsByType

 */
func (s *MdmService) GetEndpointsByType() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/ise/mdm/getEndpointsByType"

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
		return response, fmt.Errorf("error with operation GetEndpointsByType")
	}

	return response, err

}

//GetEndpointsByOsType 🚧 getEndpointsByOsType
/* 🚧 getEndpointsByOsType

 */
func (s *MdmService) GetEndpointsByOsType() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/ise/mdm/getEndpointsByOsType"

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
		return response, fmt.Errorf("error with operation GetEndpointsByOsType")
	}

	return response, err

}
