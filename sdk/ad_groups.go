package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type ADGroupsService service

type ResponseADGroupsGetAdgroups []ResponseItemADGroupsGetAdgroups // Array of ResponseADGroupsGetAdgroups

type ResponseItemADGroupsGetAdgroups struct {
	Name   string `json:"name,omitempty"`   // Active Directory Group ID
	Source string `json:"source,omitempty"` // Source of the Active Directory Group
}

//GetAdgroups Get the list of all AD groups for the specified Active Directory
/* Duo-IdentitySync  Get the list of all AD groups for the specified Active Directory

@param activeDirectory activeDirectory path parameter. List of AD groups for the specified Active Directory
*/
func (s *ADGroupsService) GetAdgroups(activeDirectory string) (*ResponseADGroupsGetAdgroups, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-identitysync/adgroups/{activeDirectory}"
	path = strings.Replace(path, "{activeDirectory}", fmt.Sprintf("%v", activeDirectory), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseADGroupsGetAdgroups{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAdgroups")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseADGroupsGetAdgroups)
	return result, response, err

}
