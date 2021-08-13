package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type TacacsServerSequenceService service

type GetTacacsServerSequenceQueryParams struct {
	Page int `url:"page,omitempty"` //Page number
	Size int `url:"size,omitempty"` //Number of objects returned per page
}

type ResponseTacacsServerSequenceGetTacacsServerSequenceByName struct {
	TacacsServerSequence ResponseTacacsServerSequenceGetTacacsServerSequenceByNameTacacsServerSequence `json:"TacacsServerSequence,omitempty"` //
}

type ResponseTacacsServerSequenceGetTacacsServerSequenceByNameTacacsServerSequence struct {
	ID               string                                                                            `json:"id,omitempty"`               //
	Name             string                                                                            `json:"name,omitempty"`             //
	Description      string                                                                            `json:"description,omitempty"`      //
	ServerList       string                                                                            `json:"serverList,omitempty"`       // The names of Tacacs external servers separated by commas. The order of the names in the string is the order of servers that will be used during authentication
	LocalAccounting  bool                                                                              `json:"localAccounting,omitempty"`  //
	RemoteAccounting bool                                                                              `json:"remoteAccounting,omitempty"` //
	PrefixStrip      bool                                                                              `json:"prefixStrip,omitempty"`      // Define if a delimiter will be used for prefix strip
	PrefixDelimiter  string                                                                            `json:"prefixDelimiter,omitempty"`  // The delimiter that will be used for prefix strip
	SuffixStrip      bool                                                                              `json:"suffixStrip,omitempty"`      // Define if a delimiter will be used for suffix strip
	SuffixDelimiter  string                                                                            `json:"suffixDelimiter,omitempty"`  // The delimiter that will be used for suffix strip
	Link             ResponseTacacsServerSequenceGetTacacsServerSequenceByNameTacacsServerSequenceLink `json:"link,omitempty"`             //
}

type ResponseTacacsServerSequenceGetTacacsServerSequenceByNameTacacsServerSequenceLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsServerSequenceGetTacacsServerSequenceByID struct {
	TacacsServerSequence ResponseTacacsServerSequenceGetTacacsServerSequenceByIDTacacsServerSequence `json:"TacacsServerSequence,omitempty"` //
}

type ResponseTacacsServerSequenceGetTacacsServerSequenceByIDTacacsServerSequence struct {
	ID               string                                                                          `json:"id,omitempty"`               //
	Name             string                                                                          `json:"name,omitempty"`             //
	Description      string                                                                          `json:"description,omitempty"`      //
	ServerList       string                                                                          `json:"serverList,omitempty"`       // The names of Tacacs external servers separated by commas. The order of the names in the string is the order of servers that will be used during authentication
	LocalAccounting  bool                                                                            `json:"localAccounting,omitempty"`  //
	RemoteAccounting bool                                                                            `json:"remoteAccounting,omitempty"` //
	PrefixStrip      bool                                                                            `json:"prefixStrip,omitempty"`      // Define if a delimiter will be used for prefix strip
	PrefixDelimiter  string                                                                          `json:"prefixDelimiter,omitempty"`  // The delimiter that will be used for prefix strip
	SuffixStrip      bool                                                                            `json:"suffixStrip,omitempty"`      // Define if a delimiter will be used for suffix strip
	SuffixDelimiter  string                                                                          `json:"suffixDelimiter,omitempty"`  // The delimiter that will be used for suffix strip
	Link             ResponseTacacsServerSequenceGetTacacsServerSequenceByIDTacacsServerSequenceLink `json:"link,omitempty"`             //
}

type ResponseTacacsServerSequenceGetTacacsServerSequenceByIDTacacsServerSequenceLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsServerSequenceUpdateTacacsServerSequenceByID struct {
	UpdatedFieldsList ResponseTacacsServerSequenceUpdateTacacsServerSequenceByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseTacacsServerSequenceUpdateTacacsServerSequenceByIDUpdatedFieldsList struct {
	UpdatedField ResponseTacacsServerSequenceUpdateTacacsServerSequenceByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                                  `json:"field,omitempty"`        //
	OldValue     string                                                                                  `json:"oldValue,omitempty"`     //
	NewValue     string                                                                                  `json:"newValue,omitempty"`     //
}

type ResponseTacacsServerSequenceUpdateTacacsServerSequenceByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseTacacsServerSequenceGetTacacsServerSequence struct {
	SearchResult ResponseTacacsServerSequenceGetTacacsServerSequenceSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseTacacsServerSequenceGetTacacsServerSequenceSearchResult struct {
	Total        int                                                                           `json:"total,omitempty"`        //
	Resources    []ResponseTacacsServerSequenceGetTacacsServerSequenceSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseTacacsServerSequenceGetTacacsServerSequenceSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseTacacsServerSequenceGetTacacsServerSequenceSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseTacacsServerSequenceGetTacacsServerSequenceSearchResultResources struct {
	ID          string                                                                       `json:"id,omitempty"`          //
	Name        string                                                                       `json:"name,omitempty"`        //
	Description string                                                                       `json:"description,omitempty"` //
	Link        ResponseTacacsServerSequenceGetTacacsServerSequenceSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseTacacsServerSequenceGetTacacsServerSequenceSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsServerSequenceGetTacacsServerSequenceSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsServerSequenceGetTacacsServerSequenceSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsServerSequenceGetVersion struct {
	VersionInfo ResponseTacacsServerSequenceGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseTacacsServerSequenceGetVersionVersionInfo struct {
	CurrentServerVersion string                                                `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                                `json:"supportedVersions,omitempty"`    //
	Link                 ResponseTacacsServerSequenceGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseTacacsServerSequenceGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestTacacsServerSequenceUpdateTacacsServerSequenceByID struct {
	TacacsServerSequence *RequestTacacsServerSequenceUpdateTacacsServerSequenceByIDTacacsServerSequence `json:"TacacsServerSequence,omitempty"` //
}

type RequestTacacsServerSequenceUpdateTacacsServerSequenceByIDTacacsServerSequence struct {
	ID               string `json:"id,omitempty"`               //
	Name             string `json:"name,omitempty"`             //
	Description      string `json:"description,omitempty"`      //
	ServerList       string `json:"serverList,omitempty"`       // The names of Tacacs external servers separated by commas. The order of the names in the string is the order of servers that will be used during authentication
	LocalAccounting  *bool  `json:"localAccounting,omitempty"`  //
	RemoteAccounting *bool  `json:"remoteAccounting,omitempty"` //
	PrefixStrip      *bool  `json:"prefixStrip,omitempty"`      // Define if a delimiter will be used for prefix strip
	PrefixDelimiter  string `json:"prefixDelimiter,omitempty"`  // The delimiter that will be used for prefix strip
	SuffixStrip      *bool  `json:"suffixStrip,omitempty"`      // Define if a delimiter will be used for suffix strip
	SuffixDelimiter  string `json:"suffixDelimiter,omitempty"`  // The delimiter that will be used for suffix strip
}

type RequestTacacsServerSequenceCreateTacacsServerSequence struct {
	TacacsServerSequence *RequestTacacsServerSequenceCreateTacacsServerSequenceTacacsServerSequence `json:"TacacsServerSequence,omitempty"` //
}

type RequestTacacsServerSequenceCreateTacacsServerSequenceTacacsServerSequence struct {
	Name             string `json:"name,omitempty"`             //
	Description      string `json:"description,omitempty"`      //
	ServerList       string `json:"serverList,omitempty"`       // The names of Tacacs external servers separated by commas. The order of the names in the string is the order of servers that will be used during authentication
	LocalAccounting  *bool  `json:"localAccounting,omitempty"`  //
	RemoteAccounting *bool  `json:"remoteAccounting,omitempty"` //
	PrefixStrip      *bool  `json:"prefixStrip,omitempty"`      // Define if a delimiter will be used for prefix strip
	PrefixDelimiter  string `json:"prefixDelimiter,omitempty"`  // The delimiter that will be used for prefix strip
	SuffixStrip      *bool  `json:"suffixStrip,omitempty"`      // Define if a delimiter will be used for suffix strip
	SuffixDelimiter  string `json:"suffixDelimiter,omitempty"`  // The delimiter that will be used for suffix strip
}

//GetTacacsServerSequenceByName Get TACACS server sequence by name
/* This API allows the client to get a TACACS server sequence by name.

@param name name path parameter.
*/
func (s *TacacsServerSequenceService) GetTacacsServerSequenceByName(name string) (*ResponseTacacsServerSequenceGetTacacsServerSequenceByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsserversequence/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTacacsServerSequenceGetTacacsServerSequenceByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTacacsServerSequenceByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsServerSequenceGetTacacsServerSequenceByName)
	return result, response, err

}

//GetTacacsServerSequenceByID Get TACACS server sequence by ID
/* This API allows the client to get a TACACS server sequence by ID.

@param id id path parameter.
*/
func (s *TacacsServerSequenceService) GetTacacsServerSequenceByID(id string) (*ResponseTacacsServerSequenceGetTacacsServerSequenceByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsserversequence/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTacacsServerSequenceGetTacacsServerSequenceByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTacacsServerSequenceById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsServerSequenceGetTacacsServerSequenceByID)
	return result, response, err

}

//GetTacacsServerSequence Get all TACACS server sequences
/* This API allows the client to get all the TACACS server sequences.

@param getTACACSServerSequenceQueryParams Filtering parameter
*/
func (s *TacacsServerSequenceService) GetTacacsServerSequence(getTACACSServerSequenceQueryParams *GetTacacsServerSequenceQueryParams) (*ResponseTacacsServerSequenceGetTacacsServerSequence, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsserversequence"

	queryString, _ := query.Values(getTACACSServerSequenceQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTacacsServerSequenceGetTacacsServerSequence{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTacacsServerSequence")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsServerSequenceGetTacacsServerSequence)
	return result, response, err

}

//GetVersion Get TACACS server sequence version information
/* This API helps to retrieve the version information related to the TACACS server sequence.

 */
func (s *TacacsServerSequenceService) GetVersion() (*ResponseTacacsServerSequenceGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsserversequence/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTacacsServerSequenceGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsServerSequenceGetVersion)
	return result, response, err

}

//CreateTacacsServerSequence Create TACACS server sequence
/* This API creates a TACACS server sequence.

 */
func (s *TacacsServerSequenceService) CreateTacacsServerSequence(requestTacacsServerSequenceCreateTACACSServerSequence *RequestTacacsServerSequenceCreateTacacsServerSequence) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsserversequence"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTacacsServerSequenceCreateTACACSServerSequence).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateTacacsServerSequence")
	}

	return response, err

}

//UpdateTacacsServerSequenceByID Update TACACS server sequence
/* This API allows the client to update a TACACS server sequence.

@param id id path parameter.
*/
func (s *TacacsServerSequenceService) UpdateTacacsServerSequenceByID(id string, requestTacacsServerSequenceUpdateTACACSServerSequenceById *RequestTacacsServerSequenceUpdateTacacsServerSequenceByID) (*ResponseTacacsServerSequenceUpdateTacacsServerSequenceByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsserversequence/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTacacsServerSequenceUpdateTACACSServerSequenceById).
		SetResult(&ResponseTacacsServerSequenceUpdateTacacsServerSequenceByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateTacacsServerSequenceById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsServerSequenceUpdateTacacsServerSequenceByID)
	return result, response, err

}

//DeleteTacacsServerSequenceByID Delete TACACS server sequence
/* This API deletes a TACACS server sequence.

@param id id path parameter.
*/
func (s *TacacsServerSequenceService) DeleteTacacsServerSequenceByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsserversequence/{id}"
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
		return response, fmt.Errorf("error with operation DeleteTacacsServerSequenceById")
	}

	getCSFRToken(response.Header())

	return response, err

}
