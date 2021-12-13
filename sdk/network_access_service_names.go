package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type NetworkAccessServiceNamesService service

type ResponseNetworkAccessServiceNamesGetNetworkAccessServiceNames []ResponseItemNetworkAccessServiceNamesGetNetworkAccessServiceNames // Array of ResponseNetworkAccessServiceNamesGetNetworkAccessServiceNames

type ResponseItemNetworkAccessServiceNamesGetNetworkAccessServiceNames struct {
	ID                   string `json:"id,omitempty"`                   //
	IsLocalAuthorization *bool  `json:"isLocalAuthorization,omitempty"` //
	Name                 string `json:"name,omitempty"`                 //
	ServiceType          string `json:"serviceType,omitempty"`          // Allowed Protocols OR Server Sequence
}

//GetNetworkAccessServiceNames Network Access - Returns list of allowed protocols and server sequences for Policy Set.
/* Returns list of Allowed Protocols and Server Sequences for Network Access Policy Set results.
'isLocalAuthorization' property is available only for Network Access Policy Set results of type Server Sequence.
(Other CRUD APIs available throught ERS)

*/
func (s *NetworkAccessServiceNamesService) GetNetworkAccessServiceNames() (*ResponseNetworkAccessServiceNamesGetNetworkAccessServiceNames, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/service-names"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessServiceNamesGetNetworkAccessServiceNames{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessServiceNames")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessServiceNamesGetNetworkAccessServiceNames)
	return result, response, err

}
