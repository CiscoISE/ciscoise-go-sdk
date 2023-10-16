package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type NetworkAccessProfilesService service

type ResponseNetworkAccessProfilesGetNetworkAccessProfiles []ResponseItemNetworkAccessProfilesGetNetworkAccessProfiles // Array of ResponseNetworkAccessProfilesGetNetworkAccessProfiles

type ResponseItemNetworkAccessProfilesGetNetworkAccessProfiles struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}

//GetNetworkAccessProfiles Network Access - Returns list of profiles.
/* Network Access Returns list of profiles.
(Other CRUD APIs available throught ERS)

*/
func (s *NetworkAccessProfilesService) GetNetworkAccessProfiles() (*ResponseNetworkAccessProfilesGetNetworkAccessProfiles, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/profiles"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessProfilesGetNetworkAccessProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessProfiles")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessProfilesGetNetworkAccessProfiles)
	return result, response, err

}
