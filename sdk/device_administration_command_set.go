package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type DeviceAdministrationCommandSetService service

type ResponseDeviceAdministrationCommandSetGetDeviceAdminCommandSets struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` // Command used in Device Admin authorization policies
}

//GetDeviceAdminCommandSets Device Admin - Return list of command sets.
/* Device Admin Return list of command sets.
(Other CRUD APIs available through ERS)

*/
func (s *DeviceAdministrationCommandSetService) GetDeviceAdminCommandSets() (*[]ResponseDeviceAdministrationCommandSetGetDeviceAdminCommandSets, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/command-sets"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&[]ResponseDeviceAdministrationCommandSetGetDeviceAdminCommandSets{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminCommandSets")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*[]ResponseDeviceAdministrationCommandSetGetDeviceAdminCommandSets)
	return result, response, err

}
