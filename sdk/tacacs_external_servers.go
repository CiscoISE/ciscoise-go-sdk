package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type TacacsExternalServersService service

type GetTacacsExternalServersQueryParams struct {
	Page int `url:"page,omitempty"` //Page number
	Size int `url:"size,omitempty"` //Number of objects returned per page
}

type ResponseTacacsExternalServersGetTacacsExternalServersByName struct {
	TacacsExternalServer ResponseTacacsExternalServersGetTacacsExternalServersByNameTacacsExternalServer `json:"TacacsExternalServer,omitempty"` //
}

type ResponseTacacsExternalServersGetTacacsExternalServersByNameTacacsExternalServer struct {
	ID             string                                                                              `json:"id,omitempty"`             //
	Name           string                                                                              `json:"name,omitempty"`           //
	Description    string                                                                              `json:"description,omitempty"`    //
	HostIP         string                                                                              `json:"hostIP,omitempty"`         // The server IPV4 address
	ConnectionPort int                                                                                 `json:"connectionPort,omitempty"` // The port to connect the server
	SingleConnect  bool                                                                                `json:"singleConnect,omitempty"`  // Define the use of single connection
	SharedSecret   string                                                                              `json:"sharedSecret,omitempty"`   // The server shared secret
	Timeout        int                                                                                 `json:"timeout,omitempty"`        // The server timeout
	Link           ResponseTacacsExternalServersGetTacacsExternalServersByNameTacacsExternalServerLink `json:"link,omitempty"`           //
}

type ResponseTacacsExternalServersGetTacacsExternalServersByNameTacacsExternalServerLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsExternalServersGetTacacsExternalServersByID struct {
	TacacsExternalServer ResponseTacacsExternalServersGetTacacsExternalServersByIDTacacsExternalServer `json:"TacacsExternalServer,omitempty"` //
}

type ResponseTacacsExternalServersGetTacacsExternalServersByIDTacacsExternalServer struct {
	ID             string                                                                            `json:"id,omitempty"`             //
	Name           string                                                                            `json:"name,omitempty"`           //
	Description    string                                                                            `json:"description,omitempty"`    //
	HostIP         string                                                                            `json:"hostIP,omitempty"`         // The server IPV4 address
	ConnectionPort int                                                                               `json:"connectionPort,omitempty"` // The port to connect the server
	SingleConnect  bool                                                                              `json:"singleConnect,omitempty"`  // Define the use of single connection
	SharedSecret   string                                                                            `json:"sharedSecret,omitempty"`   // The server shared secret
	Timeout        int                                                                               `json:"timeout,omitempty"`        // The server timeout
	Link           ResponseTacacsExternalServersGetTacacsExternalServersByIDTacacsExternalServerLink `json:"link,omitempty"`           //
}

type ResponseTacacsExternalServersGetTacacsExternalServersByIDTacacsExternalServerLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsExternalServersUpdateTacacsExternalServersByID struct {
	UpdatedFieldsList ResponseTacacsExternalServersUpdateTacacsExternalServersByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseTacacsExternalServersUpdateTacacsExternalServersByIDUpdatedFieldsList struct {
	UpdatedField ResponseTacacsExternalServersUpdateTacacsExternalServersByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                                    `json:"field,omitempty"`        //
	OldValue     string                                                                                    `json:"oldValue,omitempty"`     //
	NewValue     string                                                                                    `json:"newValue,omitempty"`     //
}

type ResponseTacacsExternalServersUpdateTacacsExternalServersByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseTacacsExternalServersGetTacacsExternalServers struct {
	SearchResult ResponseTacacsExternalServersGetTacacsExternalServersSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseTacacsExternalServersGetTacacsExternalServersSearchResult struct {
	Total        int                                                                           `json:"total,omitempty"`        //
	Resources    []ResponseTacacsExternalServersGetTacacsExternalServersSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponseTacacsExternalServersGetTacacsExternalServersSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponseTacacsExternalServersGetTacacsExternalServersSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseTacacsExternalServersGetTacacsExternalServersSearchResultResources struct {
	ID          string                                                                         `json:"id,omitempty"`          //
	Name        string                                                                         `json:"name,omitempty"`        //
	Description string                                                                         `json:"description,omitempty"` //
	Link        ResponseTacacsExternalServersGetTacacsExternalServersSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseTacacsExternalServersGetTacacsExternalServersSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsExternalServersGetTacacsExternalServersSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsExternalServersGetTacacsExternalServersSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsExternalServersGetVersion struct {
	VersionInfo ResponseTacacsExternalServersGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseTacacsExternalServersGetVersionVersionInfo struct {
	CurrentServerVersion string                                                 `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                                 `json:"supportedVersions,omitempty"`    //
	Link                 ResponseTacacsExternalServersGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseTacacsExternalServersGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestTacacsExternalServersUpdateTacacsExternalServersByID struct {
	TacacsExternalServer *RequestTacacsExternalServersUpdateTacacsExternalServersByIDTacacsExternalServer `json:"TacacsExternalServer,omitempty"` //
}

type RequestTacacsExternalServersUpdateTacacsExternalServersByIDTacacsExternalServer struct {
	ID             string `json:"id,omitempty"`             //
	Name           string `json:"name,omitempty"`           //
	Description    string `json:"description,omitempty"`    //
	HostIP         string `json:"hostIP,omitempty"`         // The server IPV4 address
	ConnectionPort *int   `json:"connectionPort,omitempty"` // The port to connect the server
	SingleConnect  *bool  `json:"singleConnect,omitempty"`  // Define the use of single connection
	SharedSecret   string `json:"sharedSecret,omitempty"`   // The server shared secret
	Timeout        *int   `json:"timeout,omitempty"`        // The server timeout
}

type RequestTacacsExternalServersCreateTacacsExternalServers struct {
	TacacsExternalServer *RequestTacacsExternalServersCreateTacacsExternalServersTacacsExternalServer `json:"TacacsExternalServer,omitempty"` //
}

type RequestTacacsExternalServersCreateTacacsExternalServersTacacsExternalServer struct {
	Name           string `json:"name,omitempty"`           //
	Description    string `json:"description,omitempty"`    //
	HostIP         string `json:"hostIP,omitempty"`         // The server IPV4 address
	ConnectionPort *int   `json:"connectionPort,omitempty"` // The port to connect the server
	SingleConnect  *bool  `json:"singleConnect,omitempty"`  // Define the use of single connection
	SharedSecret   string `json:"sharedSecret,omitempty"`   // The server shared secret
	Timeout        *int   `json:"timeout,omitempty"`        // The server timeout
}

//GetTacacsExternalServersByName Get TACACS external servers by name
/* This API allows the client to get TACACS external servers by name.

@param name name path parameter.
*/
func (s *TacacsExternalServersService) GetTacacsExternalServersByName(name string) (*ResponseTacacsExternalServersGetTacacsExternalServersByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsexternalservers/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTacacsExternalServersGetTacacsExternalServersByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTacacsExternalServersByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsExternalServersGetTacacsExternalServersByName)
	return result, response, err

}

//GetTacacsExternalServersByID Get TACACS external servers by ID
/* This API allows the client to get TACACS external servers by ID.

@param id id path parameter.
*/
func (s *TacacsExternalServersService) GetTacacsExternalServersByID(id string) (*ResponseTacacsExternalServersGetTacacsExternalServersByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsexternalservers/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTacacsExternalServersGetTacacsExternalServersByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTacacsExternalServersById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsExternalServersGetTacacsExternalServersByID)
	return result, response, err

}

//GetTacacsExternalServers Get all TACACS external servers
/* This API allows the client to get all the TACACS external servers.

@param getTACACSExternalServersQueryParams Filtering parameter
*/
func (s *TacacsExternalServersService) GetTacacsExternalServers(getTACACSExternalServersQueryParams *GetTacacsExternalServersQueryParams) (*ResponseTacacsExternalServersGetTacacsExternalServers, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsexternalservers"

	queryString, _ := query.Values(getTACACSExternalServersQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTacacsExternalServersGetTacacsExternalServers{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTacacsExternalServers")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsExternalServersGetTacacsExternalServers)
	return result, response, err

}

//GetVersion Get TACACS external servers version information
/* This API helps to retrieve the version information related to the TACACS external servers.

 */
func (s *TacacsExternalServersService) GetVersion() (*ResponseTacacsExternalServersGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsexternalservers/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTacacsExternalServersGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsExternalServersGetVersion)
	return result, response, err

}

//CreateTacacsExternalServers Create TACACS external servers
/* This API creates TACACS external servers.

 */
func (s *TacacsExternalServersService) CreateTacacsExternalServers(requestTacacsExternalServersCreateTACACSExternalServers *RequestTacacsExternalServersCreateTacacsExternalServers) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsexternalservers"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTacacsExternalServersCreateTACACSExternalServers).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateTacacsExternalServers")
	}

	return response, err

}

//UpdateTacacsExternalServersByID Update TACACS external servers
/* This API allows the client to update TACACS external servers.

@param id id path parameter.
*/
func (s *TacacsExternalServersService) UpdateTacacsExternalServersByID(id string, requestTacacsExternalServersUpdateTACACSExternalServersById *RequestTacacsExternalServersUpdateTacacsExternalServersByID) (*ResponseTacacsExternalServersUpdateTacacsExternalServersByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsexternalservers/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTacacsExternalServersUpdateTACACSExternalServersById).
		SetResult(&ResponseTacacsExternalServersUpdateTacacsExternalServersByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateTacacsExternalServersById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsExternalServersUpdateTacacsExternalServersByID)
	return result, response, err

}

//DeleteTacacsExternalServersByID Delete TACACS external servers
/* This API deletes TACACS external servers.

@param id id path parameter.
*/
func (s *TacacsExternalServersService) DeleteTacacsExternalServersByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsexternalservers/{id}"
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
		return response, fmt.Errorf("error with operation DeleteTacacsExternalServersById")
	}

	getCSFRToken(response.Header())

	return response, err

}
