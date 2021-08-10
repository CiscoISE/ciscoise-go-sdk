package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type DeviceAdministrationProfilesService service

type ResponseDeviceAdministrationProfilesGetDeviceAdminProfiles struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}

//GetDeviceAdminProfiles Device Admin - Returns list of profiles.
/* Device Admin Returns list of profiles.
(Other CRUD APIs available throught ERS)

*/
func (s *DeviceAdministrationProfilesService) GetDeviceAdminProfiles() (*[]ResponseDeviceAdministrationProfilesGetDeviceAdminProfiles, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/profiles"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&[]ResponseDeviceAdministrationProfilesGetDeviceAdminProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminProfiles")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*[]ResponseDeviceAdministrationProfilesGetDeviceAdminProfiles)
	return result, response, err

}
