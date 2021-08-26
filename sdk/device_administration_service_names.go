package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type DeviceAdministrationServiceNamesService service

type ResponseDeviceAdministrationServiceNamesGetDeviceAdminServiceNames struct {
	ID                   string `json:"id,omitempty"`                   //
	IsLocalAuthorization *bool  `json:"isLocalAuthorization,omitempty"` //
	Name                 string `json:"name,omitempty"`                 //
	ServiceType          string `json:"serviceType,omitempty"`          // Allowed Protocols OR Server Sequence
}

//GetDeviceAdminServiceNames Device Admin - Returns list of allowed protocols and server sequences.
/* Returns list of Allowed Protocols and Server Sequences for Device Admin Policy Set results.
'isLocalAuthorization' property is available only for Network Access Policy Set results of type Server Sequence.
(Other CRUD APIs available throught ERS)

*/
func (s *DeviceAdministrationServiceNamesService) GetDeviceAdminServiceNames() (*[]ResponseDeviceAdministrationServiceNamesGetDeviceAdminServiceNames, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/service-names"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&[]ResponseDeviceAdministrationServiceNamesGetDeviceAdminServiceNames{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminServiceNames")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*[]ResponseDeviceAdministrationServiceNamesGetDeviceAdminServiceNames)
	return result, response, err

}
