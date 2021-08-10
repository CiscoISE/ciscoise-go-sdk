package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ExternalRadiusServerService service

type GetExternalRadiusServerQueryParams struct {
	Page int `url: page,omitempty` //Page number
	Size int `url: size,omitempty` //Number of objects returned per page
}

type ResponseExternalRadiusServerGetExternalRadiusServerByName struct {
	ExternalRadiusServer ResponseExternalRadiusServerGetExternalRadiusServerByNameExternalRadiusServer `json:"ExternalRadiusServer,omitempty"` //
}

type ResponseExternalRadiusServerGetExternalRadiusServerByNameExternalRadiusServer struct {
	ID                 string                                                                            `json:"id,omitempty"`                 //
	Name               string                                                                            `json:"name,omitempty"`               // Resource Name. Allowed charactera are alphanumeric and _ (underscore).
	Description        string                                                                            `json:"description,omitempty"`        //
	HostIP             string                                                                            `json:"hostIP,omitempty"`             // The IP of the host - must be a valid IPV4 address
	SharedSecret       string                                                                            `json:"sharedSecret,omitempty"`       // Shared secret maximum length is 128 characters
	EnableKeyWrap      bool                                                                              `json:"enableKeyWrap,omitempty"`      // KeyWrap may only be enabled if it is supported on the device. When running in FIPS mode this option should be enabled for such devices
	EncryptionKey      string                                                                            `json:"encryptionKey,omitempty"`      // The encryptionKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty. The maximum length is 16 ASCII characters or 32 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')
	AuthenticatorKey   string                                                                            `json:"authenticatorKey,omitempty"`   // The authenticatorKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty. The maximum length is 20 ASCII characters or 40 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')
	KeyInputFormat     string                                                                            `json:"keyInputFormat,omitempty"`     // Specifies the format of the input for fields 'encryptionKey' and 'authenticatorKey'. Allowed Values: - ASCII - HEXADECIMAL
	AuthenticationPort int                                                                               `json:"authenticationPort,omitempty"` // Valid Range 1 to 65535
	AccountingPort     int                                                                               `json:"accountingPort,omitempty"`     // Valid Range 1 to 65535
	Timeout            int                                                                               `json:"timeout,omitempty"`            // Valid Range 1 to 120
	Retries            int                                                                               `json:"retries,omitempty"`            // Valid Range 1 to 9
	ProxyTimeout       int                                                                               `json:"proxyTimeout,omitempty"`       // Valid Range 1 to 600
	Link               ResponseExternalRadiusServerGetExternalRadiusServerByNameExternalRadiusServerLink `json:"link,omitempty"`               //
}

type ResponseExternalRadiusServerGetExternalRadiusServerByNameExternalRadiusServerLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseExternalRadiusServerGetExternalRadiusServerByID struct {
	ExternalRadiusServer ResponseExternalRadiusServerGetExternalRadiusServerByIDExternalRadiusServer `json:"ExternalRadiusServer,omitempty"` //
}

type ResponseExternalRadiusServerGetExternalRadiusServerByIDExternalRadiusServer struct {
	ID                 string                                                                          `json:"id,omitempty"`                 //
	Name               string                                                                          `json:"name,omitempty"`               // Resource Name. Allowed charactera are alphanumeric and _ (underscore).
	Description        string                                                                          `json:"description,omitempty"`        //
	HostIP             string                                                                          `json:"hostIP,omitempty"`             // The IP of the host - must be a valid IPV4 address
	SharedSecret       string                                                                          `json:"sharedSecret,omitempty"`       // Shared secret maximum length is 128 characters
	EnableKeyWrap      bool                                                                            `json:"enableKeyWrap,omitempty"`      // KeyWrap may only be enabled if it is supported on the device. When running in FIPS mode this option should be enabled for such devices
	EncryptionKey      string                                                                          `json:"encryptionKey,omitempty"`      // The encryptionKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty. The maximum length is 16 ASCII characters or 32 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')
	AuthenticatorKey   string                                                                          `json:"authenticatorKey,omitempty"`   // The authenticatorKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty. The maximum length is 20 ASCII characters or 40 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')
	KeyInputFormat     string                                                                          `json:"keyInputFormat,omitempty"`     // Specifies the format of the input for fields 'encryptionKey' and 'authenticatorKey'. Allowed Values: - ASCII - HEXADECIMAL
	AuthenticationPort int                                                                             `json:"authenticationPort,omitempty"` // Valid Range 1 to 65535
	AccountingPort     int                                                                             `json:"accountingPort,omitempty"`     // Valid Range 1 to 65535
	Timeout            int                                                                             `json:"timeout,omitempty"`            // Valid Range 1 to 120
	Retries            int                                                                             `json:"retries,omitempty"`            // Valid Range 1 to 9
	ProxyTimeout       int                                                                             `json:"proxyTimeout,omitempty"`       // Valid Range 1 to 600
	Link               ResponseExternalRadiusServerGetExternalRadiusServerByIDExternalRadiusServerLink `json:"link,omitempty"`               //
}

type ResponseExternalRadiusServerGetExternalRadiusServerByIDExternalRadiusServerLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseExternalRadiusServerUpdateExternalRadiusServerByID struct {
	UpdatedFieldsList ResponseExternalRadiusServerUpdateExternalRadiusServerByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseExternalRadiusServerUpdateExternalRadiusServerByIDUpdatedFieldsList struct {
	UpdatedField ResponseExternalRadiusServerUpdateExternalRadiusServerByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                                  `json:"field,omitempty"`        //
	OldValue     string                                                                                  `json:"oldValue,omitempty"`     //
	NewValue     string                                                                                  `json:"newValue,omitempty"`     //
}

type ResponseExternalRadiusServerUpdateExternalRadiusServerByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseExternalRadiusServerGetExternalRadiusServer struct {
	SearchResult ResponseExternalRadiusServerGetExternalRadiusServerSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseExternalRadiusServerGetExternalRadiusServerSearchResult struct {
	Total        int                                                                           `json:"total,omitempty"`        //
	Resources    []ResponseExternalRadiusServerGetExternalRadiusServerSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseExternalRadiusServerGetExternalRadiusServerSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseExternalRadiusServerGetExternalRadiusServerSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseExternalRadiusServerGetExternalRadiusServerSearchResultResources struct {
	ID          string                                                                       `json:"id,omitempty"`          //
	Name        string                                                                       `json:"name,omitempty"`        //
	Description string                                                                       `json:"description,omitempty"` //
	Link        ResponseExternalRadiusServerGetExternalRadiusServerSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseExternalRadiusServerGetExternalRadiusServerSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseExternalRadiusServerGetExternalRadiusServerSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseExternalRadiusServerGetExternalRadiusServerSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseExternalRadiusServerGetVersion struct {
	VersionInfo ResponseExternalRadiusServerGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseExternalRadiusServerGetVersionVersionInfo struct {
	CurrentServerVersion string                                                `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                                `json:"supportedVersions,omitempty"`    //
	Link                 ResponseExternalRadiusServerGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseExternalRadiusServerGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestExternalRadiusServerUpdateExternalRadiusServerByID struct {
	ExternalRadiusServer *RequestExternalRadiusServerUpdateExternalRadiusServerByIDExternalRadiusServer `json:"ExternalRadiusServer,omitempty"` //
}

type RequestExternalRadiusServerUpdateExternalRadiusServerByIDExternalRadiusServer struct {
	ID                 string `json:"id,omitempty"`                 //
	Name               string `json:"name,omitempty"`               // Resource Name. Allowed charactera are alphanumeric and _ (underscore).
	Description        string `json:"description,omitempty"`        //
	HostIP             string `json:"hostIP,omitempty"`             // The IP of the host - must be a valid IPV4 address
	SharedSecret       string `json:"sharedSecret,omitempty"`       // Shared secret maximum length is 128 characters
	EnableKeyWrap      *bool  `json:"enableKeyWrap,omitempty"`      // KeyWrap may only be enabled if it is supported on the device. When running in FIPS mode this option should be enabled for such devices
	EncryptionKey      string `json:"encryptionKey,omitempty"`      // The encryptionKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty. The maximum length is 16 ASCII characters or 32 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')
	AuthenticatorKey   string `json:"authenticatorKey,omitempty"`   // The authenticatorKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty. The maximum length is 20 ASCII characters or 40 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')
	KeyInputFormat     string `json:"keyInputFormat,omitempty"`     // Specifies the format of the input for fields 'encryptionKey' and 'authenticatorKey'. Allowed Values: - ASCII - HEXADECIMAL
	AuthenticationPort *int   `json:"authenticationPort,omitempty"` // Valid Range 1 to 65535
	AccountingPort     *int   `json:"accountingPort,omitempty"`     // Valid Range 1 to 65535
	Timeout            *int   `json:"timeout,omitempty"`            // Valid Range 1 to 120
	Retries            *int   `json:"retries,omitempty"`            // Valid Range 1 to 9
	ProxyTimeout       *int   `json:"proxyTimeout,omitempty"`       // Valid Range 1 to 600
}

type RequestExternalRadiusServerCreateExternalRadiusServer struct {
	ExternalRadiusServer *RequestExternalRadiusServerCreateExternalRadiusServerExternalRadiusServer `json:"ExternalRadiusServer,omitempty"` //
}

type RequestExternalRadiusServerCreateExternalRadiusServerExternalRadiusServer struct {
	Name               string `json:"name,omitempty"`               // Resource Name. Allowed charactera are alphanumeric and _ (underscore).
	Description        string `json:"description,omitempty"`        //
	HostIP             string `json:"hostIP,omitempty"`             // The IP of the host - must be a valid IPV4 address
	SharedSecret       string `json:"sharedSecret,omitempty"`       // Shared secret maximum length is 128 characters
	EnableKeyWrap      *bool  `json:"enableKeyWrap,omitempty"`      // KeyWrap may only be enabled if it is supported on the device. When running in FIPS mode this option should be enabled for such devices
	EncryptionKey      string `json:"encryptionKey,omitempty"`      // The encryptionKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty. The maximum length is 16 ASCII characters or 32 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')
	AuthenticatorKey   string `json:"authenticatorKey,omitempty"`   // The authenticatorKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty. The maximum length is 20 ASCII characters or 40 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')
	KeyInputFormat     string `json:"keyInputFormat,omitempty"`     // Specifies the format of the input for fields 'encryptionKey' and 'authenticatorKey'. Allowed Values: - ASCII - HEXADECIMAL
	AuthenticationPort *int   `json:"authenticationPort,omitempty"` // Valid Range 1 to 65535
	AccountingPort     *int   `json:"accountingPort,omitempty"`     // Valid Range 1 to 65535
	Timeout            *int   `json:"timeout,omitempty"`            // Valid Range 1 to 120
	Retries            *int   `json:"retries,omitempty"`            // Valid Range 1 to 9
	ProxyTimeout       *int   `json:"proxyTimeout,omitempty"`       // Valid Range 1 to 600
}

//GetExternalRadiusServerByName Get external RADIUS server by name
/* This API allows the client to get an external RADIUS server by name.

@param name name path parameter.
*/
func (s *ExternalRadiusServerService) GetExternalRadiusServerByName(name string) (*ResponseExternalRadiusServerGetExternalRadiusServerByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/externalradiusserver/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseExternalRadiusServerGetExternalRadiusServerByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetExternalRadiusServerByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseExternalRadiusServerGetExternalRadiusServerByName)
	return result, response, err

}

//GetExternalRadiusServerByID Get external RADIUS server by ID
/* This API allows the client to get an external RADIUS server by ID.

@param id id path parameter.
*/
func (s *ExternalRadiusServerService) GetExternalRadiusServerByID(id string) (*ResponseExternalRadiusServerGetExternalRadiusServerByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/externalradiusserver/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseExternalRadiusServerGetExternalRadiusServerByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetExternalRadiusServerById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseExternalRadiusServerGetExternalRadiusServerByID)
	return result, response, err

}

//GetExternalRadiusServer Get all external RADIUS servers
/* This API allows the client to get all the external RADIUS servers.

@param getExternalRadiusServerQueryParams Filtering parameter
*/
func (s *ExternalRadiusServerService) GetExternalRadiusServer(getExternalRadiusServerQueryParams *GetExternalRadiusServerQueryParams) (*ResponseExternalRadiusServerGetExternalRadiusServer, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/externalradiusserver"

	queryString, _ := query.Values(getExternalRadiusServerQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseExternalRadiusServerGetExternalRadiusServer{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetExternalRadiusServer")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseExternalRadiusServerGetExternalRadiusServer)
	return result, response, err

}

//GetVersion Get external RADIUS server version information
/* This API helps to retrieve the version information related to the external RADIUS server.

 */
func (s *ExternalRadiusServerService) GetVersion() (*ResponseExternalRadiusServerGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/externalradiusserver/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseExternalRadiusServerGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseExternalRadiusServerGetVersion)
	return result, response, err

}

//CreateExternalRadiusServer Create external RADIUS server
/* This API creates an external RADIUS server.

 */
func (s *ExternalRadiusServerService) CreateExternalRadiusServer(requestExternalRadiusServerCreateExternalRadiusServer *RequestExternalRadiusServerCreateExternalRadiusServer) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/externalradiusserver"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestExternalRadiusServerCreateExternalRadiusServer).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateExternalRadiusServer")
	}

	return response, err

}

//UpdateExternalRadiusServerByID Update external RADIUS server
/* This API allows the client to update an external RADIUS server.

@param id id path parameter.
*/
func (s *ExternalRadiusServerService) UpdateExternalRadiusServerByID(id string, requestExternalRadiusServerUpdateExternalRadiusServerById *RequestExternalRadiusServerUpdateExternalRadiusServerByID) (*ResponseExternalRadiusServerUpdateExternalRadiusServerByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/externalradiusserver/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestExternalRadiusServerUpdateExternalRadiusServerById).
		SetResult(&ResponseExternalRadiusServerUpdateExternalRadiusServerByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateExternalRadiusServerById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseExternalRadiusServerUpdateExternalRadiusServerByID)
	return result, response, err

}

//DeleteExternalRadiusServerByID Delete external RADIUS server
/* This API deletes an external RADIUS server.

@param id id path parameter.
*/
func (s *ExternalRadiusServerService) DeleteExternalRadiusServerByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/externalradiusserver/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

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
		return response, fmt.Errorf("error with operation DeleteExternalRadiusServerById")
	}

	getCSFRToken(response.Header())

	return response, err

}
