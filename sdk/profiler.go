package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ProfilerService service

//GetProfiles 🚧 getProfiles
/* 🚧 getProfiles

 */
func (s *ProfilerService) GetProfiles() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/ise/config/profiler/getProfiles"

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
		return response, fmt.Errorf("error with operation GetProfiles")
	}

	return response, err

}
