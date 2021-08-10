package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type NetworkAccessSecurityGroupsService service

type ResponseNetworkAccessSecurityGroupsGetNetworkAccessSecurityGroups struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}

//GetNetworkAccessSecurityGroups Network Access - Return list of available security groups for authorization policy definition.
/* Network Access Return list of available security groups for authorization policy definition.
(Other CRUD APIs available throught ERS)

*/
func (s *NetworkAccessSecurityGroupsService) GetNetworkAccessSecurityGroups() (*[]ResponseNetworkAccessSecurityGroupsGetNetworkAccessSecurityGroups, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/security-groups"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&[]ResponseNetworkAccessSecurityGroupsGetNetworkAccessSecurityGroups{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessSecurityGroups")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*[]ResponseNetworkAccessSecurityGroupsGetNetworkAccessSecurityGroups)
	return result, response, err

}
