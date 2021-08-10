package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type SystemHealthService service

//GetHealths 🚧 getHealths
/* 🚧 getHealths

 */
func (s *SystemHealthService) GetHealths() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/ise/system/getHealths"

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
		return response, fmt.Errorf("error with operation GetHealths")
	}

	return response, err

}

//GetPerformances 🚧 getPerformances
/* 🚧 getPerformances

 */
func (s *SystemHealthService) GetPerformances() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/ise/system/getPerformances"

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
		return response, fmt.Errorf("error with operation GetPerformances")
	}

	return response, err

}
