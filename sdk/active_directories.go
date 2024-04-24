package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ActiveDirectoriesService service

type ResponseActiveDirectoriesGetActiveDirectories []ResponseItemActiveDirectoriesGetActiveDirectories // Array of ResponseActiveDirectoriesGetActiveDirectories

type ResponseItemActiveDirectoriesGetActiveDirectories struct {
	DirectoryID string `json:"directoryID,omitempty"` // Active Directory ID
	Domain      string `json:"domain,omitempty"`      // Active Directory domain name
	Name        string `json:"name,omitempty"`        // Name of the Active Directory
}

//GetActiveDirectories Get the list of all configured Active Directories
/* Duo-IdentitySync Get the list of all configured Active Directories

 */
func (s *ActiveDirectoriesService) GetActiveDirectories() (*ResponseActiveDirectoriesGetActiveDirectories, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-identitysync/activedirectories"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseActiveDirectoriesGetActiveDirectories{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetActiveDirectories")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseActiveDirectoriesGetActiveDirectories)
	return result, response, err

}
