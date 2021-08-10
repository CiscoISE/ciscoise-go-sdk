package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type NetworkAccessIDentityStoresService service

type ResponseNetworkAccessIDentityStoresGetNetworkAccessIDentityStores struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}

//GetNetworkAccessIDentityStores Network Access - Return list of identity stores for authentication policy definition.
/* Network Access Return list of identity stores for authentication policy definition.
(Other CRUD APIs available throught ERS)

*/
func (s *NetworkAccessIDentityStoresService) GetNetworkAccessIDentityStores() (*[]ResponseNetworkAccessIDentityStoresGetNetworkAccessIDentityStores, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/identity-stores"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&[]ResponseNetworkAccessIDentityStoresGetNetworkAccessIDentityStores{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessIdentityStores")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*[]ResponseNetworkAccessIDentityStoresGetNetworkAccessIDentityStores)
	return result, response, err

}
