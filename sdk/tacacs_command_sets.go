package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type TacacsCommandSetsService service

type GetTacacsCommandSetsQueryParams struct {
	Page int `url: page,omitempty` //Page number
	Size int `url: size,omitempty` //Number of objects returned per page
}

type ResponseTacacsCommandSetsGetTacacsCommandSetsByName struct {
	TacacsCommandSets ResponseTacacsCommandSetsGetTacacsCommandSetsByNameTacacsCommandSets `json:"TacacsCommandSets,omitempty"` //
}

type ResponseTacacsCommandSetsGetTacacsCommandSetsByNameTacacsCommandSets struct {
	ID              string                                                                       `json:"id,omitempty"`              //
	Name            string                                                                       `json:"name,omitempty"`            //
	Description     string                                                                       `json:"description,omitempty"`     //
	PermitUnmatched bool                                                                         `json:"permitUnmatched,omitempty"` //
	Commands        ResponseTacacsCommandSetsGetTacacsCommandSetsByNameTacacsCommandSetsCommands `json:"commands,omitempty"`        //
	Link            ResponseTacacsCommandSetsGetTacacsCommandSetsByNameTacacsCommandSetsLink     `json:"link,omitempty"`            //
}

type ResponseTacacsCommandSetsGetTacacsCommandSetsByNameTacacsCommandSetsCommands struct {
	CommandList []ResponseTacacsCommandSetsGetTacacsCommandSetsByNameTacacsCommandSetsCommandsCommandList `json:"commandList,omitempty"` //
}

type ResponseTacacsCommandSetsGetTacacsCommandSetsByNameTacacsCommandSetsCommandsCommandList struct {
	Grant     string `json:"grant,omitempty"`     // Allowed values: PERMIT, DENY, DENY_ALWAYS
	Command   string `json:"command,omitempty"`   //
	Arguments string `json:"arguments,omitempty"` //
}

type ResponseTacacsCommandSetsGetTacacsCommandSetsByNameTacacsCommandSetsLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsCommandSetsGetTacacsCommandSetsByID struct {
	TacacsCommandSets ResponseTacacsCommandSetsGetTacacsCommandSetsByIDTacacsCommandSets `json:"TacacsCommandSets,omitempty"` //
}

type ResponseTacacsCommandSetsGetTacacsCommandSetsByIDTacacsCommandSets struct {
	ID              string                                                                     `json:"id,omitempty"`              //
	Name            string                                                                     `json:"name,omitempty"`            //
	Description     string                                                                     `json:"description,omitempty"`     //
	PermitUnmatched bool                                                                       `json:"permitUnmatched,omitempty"` //
	Commands        ResponseTacacsCommandSetsGetTacacsCommandSetsByIDTacacsCommandSetsCommands `json:"commands,omitempty"`        //
	Link            ResponseTacacsCommandSetsGetTacacsCommandSetsByIDTacacsCommandSetsLink     `json:"link,omitempty"`            //
}

type ResponseTacacsCommandSetsGetTacacsCommandSetsByIDTacacsCommandSetsCommands struct {
	CommandList []ResponseTacacsCommandSetsGetTacacsCommandSetsByIDTacacsCommandSetsCommandsCommandList `json:"commandList,omitempty"` //
}

type ResponseTacacsCommandSetsGetTacacsCommandSetsByIDTacacsCommandSetsCommandsCommandList struct {
	Grant     string `json:"grant,omitempty"`     // Allowed values: PERMIT, DENY, DENY_ALWAYS
	Command   string `json:"command,omitempty"`   //
	Arguments string `json:"arguments,omitempty"` //
}

type ResponseTacacsCommandSetsGetTacacsCommandSetsByIDTacacsCommandSetsLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsCommandSetsUpdateTacacsCommandSetsByID struct {
	UpdatedFieldsList ResponseTacacsCommandSetsUpdateTacacsCommandSetsByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseTacacsCommandSetsUpdateTacacsCommandSetsByIDUpdatedFieldsList struct {
	UpdatedField ResponseTacacsCommandSetsUpdateTacacsCommandSetsByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                            `json:"field,omitempty"`        //
	OldValue     string                                                                            `json:"oldValue,omitempty"`     //
	NewValue     string                                                                            `json:"newValue,omitempty"`     //
}

type ResponseTacacsCommandSetsUpdateTacacsCommandSetsByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseTacacsCommandSetsGetTacacsCommandSets struct {
	SearchResult ResponseTacacsCommandSetsGetTacacsCommandSetsSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseTacacsCommandSetsGetTacacsCommandSetsSearchResult struct {
	Total        int                                                                     `json:"total,omitempty"`        //
	Resources    []ResponseTacacsCommandSetsGetTacacsCommandSetsSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseTacacsCommandSetsGetTacacsCommandSetsSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseTacacsCommandSetsGetTacacsCommandSetsSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseTacacsCommandSetsGetTacacsCommandSetsSearchResultResources struct {
	ID          string                                                                 `json:"id,omitempty"`          //
	Name        string                                                                 `json:"name,omitempty"`        //
	Description string                                                                 `json:"description,omitempty"` //
	Link        ResponseTacacsCommandSetsGetTacacsCommandSetsSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseTacacsCommandSetsGetTacacsCommandSetsSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsCommandSetsGetTacacsCommandSetsSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsCommandSetsGetTacacsCommandSetsSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsCommandSetsGetVersion struct {
	VersionInfo ResponseTacacsCommandSetsGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseTacacsCommandSetsGetVersionVersionInfo struct {
	CurrentServerVersion string                                             `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                             `json:"supportedVersions,omitempty"`    //
	Link                 ResponseTacacsCommandSetsGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseTacacsCommandSetsGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestTacacsCommandSetsUpdateTacacsCommandSetsByID struct {
	TacacsCommandSets *RequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSets `json:"TacacsCommandSets,omitempty"` //
}

type RequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSets struct {
	ID              string                                                                        `json:"id,omitempty"`              //
	Name            string                                                                        `json:"name,omitempty"`            //
	Description     string                                                                        `json:"description,omitempty"`     //
	PermitUnmatched *bool                                                                         `json:"permitUnmatched,omitempty"` //
	Commands        *RequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSetsCommands `json:"commands,omitempty"`        //
}

type RequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSetsCommands struct {
	CommandList *[]RequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSetsCommandsCommandList `json:"commandList,omitempty"` //
}

type RequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSetsCommandsCommandList struct {
	Grant     string `json:"grant,omitempty"`     // Allowed values: PERMIT, DENY, DENY_ALWAYS
	Command   string `json:"command,omitempty"`   //
	Arguments string `json:"arguments,omitempty"` //
}

type RequestTacacsCommandSetsCreateTacacsCommandSets struct {
	TacacsCommandSets *RequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSets `json:"TacacsCommandSets,omitempty"` //
}

type RequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSets struct {
	Name            string                                                                    `json:"name,omitempty"`            //
	Description     string                                                                    `json:"description,omitempty"`     //
	PermitUnmatched *bool                                                                     `json:"permitUnmatched,omitempty"` //
	Commands        *RequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSetsCommands `json:"commands,omitempty"`        //
}

type RequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSetsCommands struct {
	CommandList *[]RequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSetsCommandsCommandList `json:"commandList,omitempty"` //
}

type RequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSetsCommandsCommandList struct {
	Grant     string `json:"grant,omitempty"`     // Allowed values: PERMIT, DENY, DENY_ALWAYS
	Command   string `json:"command,omitempty"`   //
	Arguments string `json:"arguments,omitempty"` //
}

//GetTacacsCommandSetsByName Get TACACS command sets by name
/* This API allows the client to get TACACS command sets by name.

@param name name path parameter.
*/
func (s *TacacsCommandSetsService) GetTacacsCommandSetsByName(name string) (*ResponseTacacsCommandSetsGetTacacsCommandSetsByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacscommandsets/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTacacsCommandSetsGetTacacsCommandSetsByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTacacsCommandSetsByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsCommandSetsGetTacacsCommandSetsByName)
	return result, response, err

}

//GetTacacsCommandSetsByID Get TACACS command sets by ID
/* This API allows the client to get TACACS command sets by ID.

@param id id path parameter.
*/
func (s *TacacsCommandSetsService) GetTacacsCommandSetsByID(id string) (*ResponseTacacsCommandSetsGetTacacsCommandSetsByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacscommandsets/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTacacsCommandSetsGetTacacsCommandSetsByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTacacsCommandSetsById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsCommandSetsGetTacacsCommandSetsByID)
	return result, response, err

}

//GetTacacsCommandSets Get all TACACS command sets
/* This API allows the client to get all the TACACS command sets.

@param getTACACSCommandSetsQueryParams Filtering parameter
*/
func (s *TacacsCommandSetsService) GetTacacsCommandSets(getTACACSCommandSetsQueryParams *GetTacacsCommandSetsQueryParams) (*ResponseTacacsCommandSetsGetTacacsCommandSets, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacscommandsets"

	queryString, _ := query.Values(getTACACSCommandSetsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTacacsCommandSetsGetTacacsCommandSets{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTacacsCommandSets")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsCommandSetsGetTacacsCommandSets)
	return result, response, err

}

//GetVersion Get TACACS command sets version information
/* This API helps to retrieve the version information related to the TACACS command sets.

 */
func (s *TacacsCommandSetsService) GetVersion() (*ResponseTacacsCommandSetsGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacscommandsets/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTacacsCommandSetsGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsCommandSetsGetVersion)
	return result, response, err

}

//CreateTacacsCommandSets Create TACACS command sets
/* This API creates TACACS command sets.

 */
func (s *TacacsCommandSetsService) CreateTacacsCommandSets(requestTacacsCommandSetsCreateTACACSCommandSets *RequestTacacsCommandSetsCreateTacacsCommandSets) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacscommandsets"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTacacsCommandSetsCreateTACACSCommandSets).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateTacacsCommandSets")
	}

	return response, err

}

//UpdateTacacsCommandSetsByID Update TACACS command sets
/* This API allows the client to update TACACS command sets.

@param id id path parameter.
*/
func (s *TacacsCommandSetsService) UpdateTacacsCommandSetsByID(id string, requestTacacsCommandSetsUpdateTACACSCommandSetsById *RequestTacacsCommandSetsUpdateTacacsCommandSetsByID) (*ResponseTacacsCommandSetsUpdateTacacsCommandSetsByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacscommandsets/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTacacsCommandSetsUpdateTACACSCommandSetsById).
		SetResult(&ResponseTacacsCommandSetsUpdateTacacsCommandSetsByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateTacacsCommandSetsById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsCommandSetsUpdateTacacsCommandSetsByID)
	return result, response, err

}

//DeleteTacacsCommandSetsByID Delete TACACS command sets
/* This API deletes TACACS command sets.

@param id id path parameter.
*/
func (s *TacacsCommandSetsService) DeleteTacacsCommandSetsByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacscommandsets/{id}"
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
		return response, fmt.Errorf("error with operation DeleteTacacsCommandSetsById")
	}

	getCSFRToken(response.Header())

	return response, err

}
