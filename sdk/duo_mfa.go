package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type DuoMfaService service

type ResponseDuoMfaGetMfa struct {
	Response *[]ResponseDuoMfaGetMfaResponse `json:"response,omitempty"` //
	Version  string                          `json:"version,omitempty"`  //
}

type ResponseDuoMfaGetMfaResponse struct {
	APIHostName  string `json:"apiHostName,omitempty"`  // Duo API HostName
	IDentitySync string `json:"identitySync,omitempty"` // Name of the Identity Sync configuration
	Name         string `json:"name,omitempty"`         // Name of the Duo-MFA configuration
	Provider     string `json:"provider,omitempty"`     // Name of the Mfa provider
	Type         string `json:"type,omitempty"`         // Protocol type for which Mfa can be used
}

// # Review unknown case
// # Review unknown case

type ResponseDuoMfaGetMfaByconnectionName struct {
	Response *ResponseDuoMfaGetMfaByconnectionNameResponse `json:"response,omitempty"` // Duo-MFA configuration information
	Version  string                                        `json:"version,omitempty"`  //
}

type ResponseDuoMfaGetMfaByconnectionNameResponse struct {
	Mfa *ResponseDuoMfaGetMfaByconnectionNameResponseMfa `json:"mfa,omitempty"` // Duo-MFA information
}

type ResponseDuoMfaGetMfaByconnectionNameResponseMfa struct {
	AccountConfigurations *ResponseDuoMfaGetMfaByconnectionNameResponseMfaAccountConfigurations `json:"accountConfigurations,omitempty"` //
	ConnectionName        string                                                                `json:"connectionName,omitempty"`        // Name of the Duo-MFA configuration
	Description           string                                                                `json:"description,omitempty"`           // Description of the Duo-MFA configuration
	IDentitySync          string                                                                `json:"identitySync,omitempty"`          // Name of the Identity Sync configuration
	Type                  string                                                                `json:"type,omitempty"`                  // Protocol type for which this Duo-MFA can be used
}

type ResponseDuoMfaGetMfaByconnectionNameResponseMfaAccountConfigurations struct {
	AdminAPI          *ResponseDuoMfaGetMfaByconnectionNameResponseMfaAccountConfigurationsAdminAPI          `json:"adminApi,omitempty"`          // API type
	APIHostName       string                                                                                 `json:"apiHostName,omitempty"`       // Duo API HostName
	AuthenticationAPI *ResponseDuoMfaGetMfaByconnectionNameResponseMfaAccountConfigurationsAuthenticationAPI `json:"authenticationApi,omitempty"` // API type
}

type ResponseDuoMfaGetMfaByconnectionNameResponseMfaAccountConfigurationsAdminAPI struct {
	Ikey string `json:"ikey,omitempty"` // Integration Key
	SKey string `json:"sKey,omitempty"` // Secret Key
}

type ResponseDuoMfaGetMfaByconnectionNameResponseMfaAccountConfigurationsAuthenticationAPI struct {
	Ikey string `json:"ikey,omitempty"` // Integration Key
	SKey string `json:"sKey,omitempty"` // Secret Key
}

// # Review unknown case
// # Review unknown case

type RequestDuoMfaCreateMfa struct {
	Mfa *RequestDuoMfaCreateMfaMfa `json:"mfa,omitempty"` // Duo-MFA information
}

type RequestDuoMfaCreateMfaMfa struct {
	AccountConfigurations *RequestDuoMfaCreateMfaMfaAccountConfigurations `json:"accountConfigurations,omitempty"` //
	ConnectionName        string                                          `json:"connectionName,omitempty"`        // Name of the Duo-MFA configuration
	Description           string                                          `json:"description,omitempty"`           // Description of the Duo-MFA configuration
	IDentitySync          string                                          `json:"identitySync,omitempty"`          // Name of the Identity Sync configuration
	Type                  string                                          `json:"type,omitempty"`                  // Protocol type for which this Duo-MFA can be used
}

type RequestDuoMfaCreateMfaMfaAccountConfigurations struct {
	AdminAPI          *RequestDuoMfaCreateMfaMfaAccountConfigurationsAdminAPI          `json:"adminApi,omitempty"`          // API type
	APIHostName       string                                                           `json:"apiHostName,omitempty"`       // Duo API HostName
	AuthenticationAPI *RequestDuoMfaCreateMfaMfaAccountConfigurationsAuthenticationAPI `json:"authenticationApi,omitempty"` // API type
}

type RequestDuoMfaCreateMfaMfaAccountConfigurationsAdminAPI struct {
	Ikey string `json:"ikey,omitempty"` // Integration Key
	SKey string `json:"sKey,omitempty"` // Secret Key
}

type RequestDuoMfaCreateMfaMfaAccountConfigurationsAuthenticationAPI struct {
	Ikey string `json:"ikey,omitempty"` // Integration Key
	SKey string `json:"sKey,omitempty"` // Secret Key
}

type RequestDuoMfaTestConnection struct {
	AdminAPI          *RequestDuoMfaTestConnectionAdminAPI          `json:"adminApi,omitempty"`          // API type
	APIHostName       string                                        `json:"apiHostName,omitempty"`       // Duo API HostName
	AuthenticationAPI *RequestDuoMfaTestConnectionAuthenticationAPI `json:"authenticationApi,omitempty"` // API type
}

type RequestDuoMfaTestConnectionAdminAPI struct {
	Ikey string `json:"ikey,omitempty"` // Integration Key
	SKey string `json:"sKey,omitempty"` // Secret Key
}

type RequestDuoMfaTestConnectionAuthenticationAPI struct {
	Ikey string `json:"ikey,omitempty"` // Integration Key
	SKey string `json:"sKey,omitempty"` // Secret Key
}

type RequestDuoMfaUpdateMFaByConnectionName struct {
	Mfa *RequestDuoMfaUpdateMFaByConnectionNameMfa `json:"mfa,omitempty"` // Duo-MFA information
}

type RequestDuoMfaUpdateMFaByConnectionNameMfa struct {
	AccountConfigurations *RequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurations `json:"accountConfigurations,omitempty"` //
	ConnectionName        string                                                          `json:"connectionName,omitempty"`        // Name of the Duo-MFA configuration
	Description           string                                                          `json:"description,omitempty"`           // Description of the Duo-MFA configuration
	IDentitySync          string                                                          `json:"identitySync,omitempty"`          // Name of the Identity Sync configuration
	Type                  string                                                          `json:"type,omitempty"`                  // Protocol type for which this Duo-MFA can be used
}

type RequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurations struct {
	AdminAPI          *RequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurationsAdminAPI          `json:"adminApi,omitempty"`          // API type
	APIHostName       string                                                                           `json:"apiHostName,omitempty"`       // Duo API HostName
	AuthenticationAPI *RequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurationsAuthenticationAPI `json:"authenticationApi,omitempty"` // API type
}

type RequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurationsAdminAPI struct {
	Ikey string `json:"ikey,omitempty"` // Integration Key
	SKey string `json:"sKey,omitempty"` // Secret Key
}

type RequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurationsAuthenticationAPI struct {
	Ikey string `json:"ikey,omitempty"` // Integration Key
	SKey string `json:"sKey,omitempty"` // Secret Key
}

//GetMfa Get the list of all Duo-MFA configurations
/* Duo-MFA List of Duo-MFA configurations

 */
func (s *DuoMfaService) GetMfa() (*ResponseDuoMfaGetMfa, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-mfa/mfa"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDuoMfaGetMfa{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetMfa")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDuoMfaGetMfa)
	return result, response, err

}

//GetMfaByconnectionName Get the specified Duo-MFA configuration
/* Duo-MFA Get the Duo-MFA configuration specified in the connectionName.

@param connectionName connectionName path parameter. This name is used to update, delete or retrieve the specific Duo-MFA configuration.
*/
func (s *DuoMfaService) GetMfaByconnectionName(connectionName string) (*ResponseDuoMfaGetMfaByconnectionName, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-mfa/mfa/{connectionName}"
	path = strings.Replace(path, "{connectionName}", fmt.Sprintf("%v", connectionName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDuoMfaGetMfaByconnectionName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetMfaByconnectionName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDuoMfaGetMfaByconnectionName)
	return result, response, err

}

//CreateMfa Create a new Duo-MFA configuration
/* Duo-MFA Create a new Duo-MFA configuration

 */
func (s *DuoMfaService) CreateMfa(requestDuoMfaCreateMfa *RequestDuoMfaCreateMfa) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-mfa/mfa"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDuoMfaCreateMfa).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateMfa")
	}

	return response, err

}

//TestConnection Verify the Auth and Admin API keys of the Duo Host
/* Duo-MFA Verify the Auth and Admin API keys of the Duo Host.

@param connectionName connectionName path parameter. This name is used to retrieve secret keys for testing connection of the specified Duo-MFA configuration in case none are specified.
*/
func (s *DuoMfaService) TestConnection(connectionName string, requestDuoMfaTestConnection *RequestDuoMfaTestConnection) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-mfa/mfa/testconnection/{connectionName}"
	path = strings.Replace(path, "{connectionName}", fmt.Sprintf("%v", connectionName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDuoMfaTestConnection).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation TestConnection")
	}

	return response, err

}

//UpdateMFaByConnectionName Update the specified Duo-MFA configuration
/* Duo-MFA Update the Duo-MFA configuration specified in the connectionName.

@param connectionName connectionName path parameter. This name is used to update, delete or retrieve the specific Duo-MFA configuration.
*/
func (s *DuoMfaService) UpdateMFaByConnectionName(connectionName string, requestDuoMfaUpdateMFaByConnectionName *RequestDuoMfaUpdateMFaByConnectionName) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-mfa/mfa/{connectionName}"
	path = strings.Replace(path, "{connectionName}", fmt.Sprintf("%v", connectionName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDuoMfaUpdateMFaByConnectionName).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateMFaByConnectionName")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteMfaByConnectionName Delete the specified Duo-MFA configuration
/* Duo-MFA Delete the Duo-MFA configuration specified in the connectionName.

@param connectionName connectionName path parameter. This name is used to update, delete or retrieve the specific Duo-MFA configuration.
*/
func (s *DuoMfaService) DeleteMfaByConnectionName(connectionName string) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-mfa/mfa/{connectionName}"
	path = strings.Replace(path, "{connectionName}", fmt.Sprintf("%v", connectionName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteMfaByConnectionName")
	}

	getCSFRToken(response.Header())

	return response, err

}
