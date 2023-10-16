package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type DeviceAdministrationIDentityStoresService service

type ResponseDeviceAdministrationIDentityStoresGetDeviceAdminIDentityStores []ResponseItemDeviceAdministrationIDentityStoresGetDeviceAdminIDentityStores // Array of ResponseDeviceAdministrationIDentityStoresGetDeviceAdminIdentityStores

type ResponseItemDeviceAdministrationIDentityStoresGetDeviceAdminIDentityStores struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}

//GetDeviceAdminIDentityStores Device Admin - Return list of identity stores for authentication.
/* Device Admin Return list of identity stores for authentication.
(Other CRUD APIs available throught ERS)

*/
func (s *DeviceAdministrationIDentityStoresService) GetDeviceAdminIDentityStores() (*ResponseDeviceAdministrationIDentityStoresGetDeviceAdminIDentityStores, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/identity-stores"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationIDentityStoresGetDeviceAdminIDentityStores{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminIdentityStores")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationIDentityStoresGetDeviceAdminIDentityStores)
	return result, response, err

}
