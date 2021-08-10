package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type RadiusFailureService service

//GetFailures 🚧 getFailures
/* 🚧 getFailures

 */
func (s *RadiusFailureService) GetFailures() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/getFailures"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetFailures")
	}

	getCSFRToken(response.Header())

	return response, err

}
