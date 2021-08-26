package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type RadiusServerSequenceService service

type GetRadiusServerSequenceQueryParams struct {
	Page int `url:"page,omitempty"` //Page number
	Size int `url:"size,omitempty"` //Number of objects returned per page
}

type ResponseRadiusServerSequenceGetRadiusServerSequenceByID struct {
	RadiusServerSequence *ResponseRadiusServerSequenceGetRadiusServerSequenceByIDRadiusServerSequence `json:"RadiusServerSequence,omitempty"` //
}

type ResponseRadiusServerSequenceGetRadiusServerSequenceByIDRadiusServerSequence struct {
	ID                               string                                                                                                         `json:"id,omitempty"`                               //
	Name                             string                                                                                                         `json:"name,omitempty"`                             //
	Description                      string                                                                                                         `json:"description,omitempty"`                      //
	StripPrefix                      *bool                                                                                                          `json:"stripPrefix,omitempty"`                      //
	StripSuffix                      *bool                                                                                                          `json:"stripSuffix,omitempty"`                      //
	PrefixSeparator                  string                                                                                                         `json:"prefixSeparator,omitempty"`                  // The prefixSeparator is required only if stripPrefix is true. The maximum length is 1 character
	SuffixSeparator                  string                                                                                                         `json:"suffixSeparator,omitempty"`                  // The suffixSeparator is required only if stripSuffix is true. The maximum length is 1 character
	RemoteAccounting                 *bool                                                                                                          `json:"remoteAccounting,omitempty"`                 //
	LocalAccounting                  *bool                                                                                                          `json:"localAccounting,omitempty"`                  //
	UseAttrSetOnRequest              *bool                                                                                                          `json:"useAttrSetOnRequest,omitempty"`              //
	UseAttrSetBeforeAcc              *bool                                                                                                          `json:"useAttrSetBeforeAcc,omitempty"`              //
	ContinueAuthorzPolicy            *bool                                                                                                          `json:"continueAuthorzPolicy,omitempty"`            //
	RadiusServerList                 []string                                                                                                       `json:"RadiusServerList,omitempty"`                 //
	OnRequestAttrManipulatorList     *[]ResponseRadiusServerSequenceGetRadiusServerSequenceByIDRadiusServerSequenceOnRequestAttrManipulatorList     `json:"OnRequestAttrManipulatorList,omitempty"`     // The onRequestAttrManipulators is required only if useAttrSetOnRequest is true
	BeforeAcceptAttrManipulatorsList *[]ResponseRadiusServerSequenceGetRadiusServerSequenceByIDRadiusServerSequenceBeforeAcceptAttrManipulatorsList `json:"BeforeAcceptAttrManipulatorsList,omitempty"` // The beforeAcceptAttrManipulators is required only if useAttrSetBeforeAcc is true
	Link                             *ResponseRadiusServerSequenceGetRadiusServerSequenceByIDRadiusServerSequenceLink                               `json:"link,omitempty"`                             //
}

type ResponseRadiusServerSequenceGetRadiusServerSequenceByIDRadiusServerSequenceOnRequestAttrManipulatorList struct {
	Action         string `json:"action,omitempty"`         // Allowed Values: - ADD, - UPDATE, - REMOVE, - REMOVEANY
	DictionaryName string `json:"dictionaryName,omitempty"` //
	AttributeName  string `json:"attributeName,omitempty"`  //
	Value          string `json:"value,omitempty"`          //
	ChangedVal     string `json:"changedVal,omitempty"`     // The changedVal is required only if the action equals to 'UPDATE'
}

type ResponseRadiusServerSequenceGetRadiusServerSequenceByIDRadiusServerSequenceBeforeAcceptAttrManipulatorsList struct {
	Action         string `json:"action,omitempty"`         // Allowed Values: - ADD, - UPDATE, - REMOVE, - REMOVEANY
	DictionaryName string `json:"dictionaryName,omitempty"` //
	AttributeName  string `json:"attributeName,omitempty"`  //
	Value          string `json:"value,omitempty"`          //
	ChangedVal     string `json:"changedVal,omitempty"`     // The changedVal is required only if the action equals to 'UPDATE'
}

type ResponseRadiusServerSequenceGetRadiusServerSequenceByIDRadiusServerSequenceLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseRadiusServerSequenceUpdateRadiusServerSequenceByID struct {
	UpdatedFieldsList *ResponseRadiusServerSequenceUpdateRadiusServerSequenceByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseRadiusServerSequenceUpdateRadiusServerSequenceByIDUpdatedFieldsList struct {
	UpdatedField *[]ResponseRadiusServerSequenceUpdateRadiusServerSequenceByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                                     `json:"field,omitempty"`        //
	OldValue     string                                                                                     `json:"oldValue,omitempty"`     //
	NewValue     string                                                                                     `json:"newValue,omitempty"`     //
}

type ResponseRadiusServerSequenceUpdateRadiusServerSequenceByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseRadiusServerSequenceGetRadiusServerSequence struct {
	SearchResult *ResponseRadiusServerSequenceGetRadiusServerSequenceSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseRadiusServerSequenceGetRadiusServerSequenceSearchResult struct {
	Total        *int                                                                         `json:"total,omitempty"`        //
	Resources    *[]ResponseRadiusServerSequenceGetRadiusServerSequenceSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponseRadiusServerSequenceGetRadiusServerSequenceSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseRadiusServerSequenceGetRadiusServerSequenceSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseRadiusServerSequenceGetRadiusServerSequenceSearchResultResources struct {
	ID          string                                                                        `json:"id,omitempty"`          //
	Name        string                                                                        `json:"name,omitempty"`        //
	Description string                                                                        `json:"description,omitempty"` //
	Link        *ResponseRadiusServerSequenceGetRadiusServerSequenceSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseRadiusServerSequenceGetRadiusServerSequenceSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseRadiusServerSequenceGetRadiusServerSequenceSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseRadiusServerSequenceGetRadiusServerSequenceSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseRadiusServerSequenceGetVersion struct {
	VersionInfo *ResponseRadiusServerSequenceGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseRadiusServerSequenceGetVersionVersionInfo struct {
	CurrentServerVersion string                                                 `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                                 `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseRadiusServerSequenceGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseRadiusServerSequenceGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestRadiusServerSequenceUpdateRadiusServerSequenceByID struct {
	RadiusServerSequence *RequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequence `json:"RadiusServerSequence,omitempty"` //
}

type RequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequence struct {
	ID                               string                                                                                                           `json:"id,omitempty"`                               //
	Name                             string                                                                                                           `json:"name,omitempty"`                             //
	Description                      string                                                                                                           `json:"description,omitempty"`                      //
	StripPrefix                      *bool                                                                                                            `json:"stripPrefix,omitempty"`                      //
	StripSuffix                      *bool                                                                                                            `json:"stripSuffix,omitempty"`                      //
	PrefixSeparator                  string                                                                                                           `json:"prefixSeparator,omitempty"`                  // The prefixSeparator is required only if stripPrefix is true. The maximum length is 1 character
	SuffixSeparator                  string                                                                                                           `json:"suffixSeparator,omitempty"`                  // The suffixSeparator is required only if stripSuffix is true. The maximum length is 1 character
	RemoteAccounting                 *bool                                                                                                            `json:"remoteAccounting,omitempty"`                 //
	LocalAccounting                  *bool                                                                                                            `json:"localAccounting,omitempty"`                  //
	UseAttrSetOnRequest              *bool                                                                                                            `json:"useAttrSetOnRequest,omitempty"`              //
	UseAttrSetBeforeAcc              *bool                                                                                                            `json:"useAttrSetBeforeAcc,omitempty"`              //
	ContinueAuthorzPolicy            *bool                                                                                                            `json:"continueAuthorzPolicy,omitempty"`            //
	RadiusServerList                 []string                                                                                                         `json:"RadiusServerList,omitempty"`                 //
	OnRequestAttrManipulatorList     *[]RequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceOnRequestAttrManipulatorList     `json:"OnRequestAttrManipulatorList,omitempty"`     // The onRequestAttrManipulators is required only if useAttrSetOnRequest is true
	BeforeAcceptAttrManipulatorsList *[]RequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceBeforeAcceptAttrManipulatorsList `json:"BeforeAcceptAttrManipulatorsList,omitempty"` // The beforeAcceptAttrManipulators is required only if useAttrSetBeforeAcc is true
}

type RequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceOnRequestAttrManipulatorList struct {
	Action         string `json:"action,omitempty"`         // Allowed Values: - ADD, - UPDATE, - REMOVE, - REMOVEANY
	DictionaryName string `json:"dictionaryName,omitempty"` //
	AttributeName  string `json:"attributeName,omitempty"`  //
	Value          string `json:"value,omitempty"`          //
	ChangedVal     string `json:"changedVal,omitempty"`     // The changedVal is required only if the action equals to 'UPDATE'
}

type RequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceBeforeAcceptAttrManipulatorsList struct {
	Action         string `json:"action,omitempty"`         // Allowed Values: - ADD, - UPDATE, - REMOVE, - REMOVEANY
	DictionaryName string `json:"dictionaryName,omitempty"` //
	AttributeName  string `json:"attributeName,omitempty"`  //
	Value          string `json:"value,omitempty"`          //
	ChangedVal     string `json:"changedVal,omitempty"`     // The changedVal is required only if the action equals to 'UPDATE'
}

type RequestRadiusServerSequenceCreateRadiusServerSequence struct {
	RadiusServerSequence *RequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequence `json:"RadiusServerSequence,omitempty"` //
}

type RequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequence struct {
	Name                             string                                                                                                       `json:"name,omitempty"`                             //
	Description                      string                                                                                                       `json:"description,omitempty"`                      //
	StripPrefix                      *bool                                                                                                        `json:"stripPrefix,omitempty"`                      //
	StripSuffix                      *bool                                                                                                        `json:"stripSuffix,omitempty"`                      //
	PrefixSeparator                  string                                                                                                       `json:"prefixSeparator,omitempty"`                  // The prefixSeparator is required only if stripPrefix is true. The maximum length is 1 character
	SuffixSeparator                  string                                                                                                       `json:"suffixSeparator,omitempty"`                  // The suffixSeparator is required only if stripSuffix is true. The maximum length is 1 character
	RemoteAccounting                 *bool                                                                                                        `json:"remoteAccounting,omitempty"`                 //
	LocalAccounting                  *bool                                                                                                        `json:"localAccounting,omitempty"`                  //
	UseAttrSetOnRequest              *bool                                                                                                        `json:"useAttrSetOnRequest,omitempty"`              //
	UseAttrSetBeforeAcc              *bool                                                                                                        `json:"useAttrSetBeforeAcc,omitempty"`              //
	ContinueAuthorzPolicy            *bool                                                                                                        `json:"continueAuthorzPolicy,omitempty"`            //
	RadiusServerList                 []string                                                                                                     `json:"RadiusServerList,omitempty"`                 //
	OnRequestAttrManipulatorList     *[]RequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceOnRequestAttrManipulatorList     `json:"OnRequestAttrManipulatorList,omitempty"`     // The onRequestAttrManipulators is required only if useAttrSetOnRequest is true
	BeforeAcceptAttrManipulatorsList *[]RequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceBeforeAcceptAttrManipulatorsList `json:"BeforeAcceptAttrManipulatorsList,omitempty"` // The beforeAcceptAttrManipulators is required only if useAttrSetBeforeAcc is true
}

type RequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceOnRequestAttrManipulatorList struct {
	Action         string `json:"action,omitempty"`         // Allowed Values: - ADD, - UPDATE, - REMOVE, - REMOVEANY
	DictionaryName string `json:"dictionaryName,omitempty"` //
	AttributeName  string `json:"attributeName,omitempty"`  //
	Value          string `json:"value,omitempty"`          //
	ChangedVal     string `json:"changedVal,omitempty"`     // The changedVal is required only if the action equals to 'UPDATE'
}

type RequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceBeforeAcceptAttrManipulatorsList struct {
	Action         string `json:"action,omitempty"`         // Allowed Values: - ADD, - UPDATE, - REMOVE, - REMOVEANY
	DictionaryName string `json:"dictionaryName,omitempty"` //
	AttributeName  string `json:"attributeName,omitempty"`  //
	Value          string `json:"value,omitempty"`          //
	ChangedVal     string `json:"changedVal,omitempty"`     // The changedVal is required only if the action equals to 'UPDATE'
}

//GetRadiusServerSequenceByID Get RADIUS server sequence by ID
/* This API allows the client to get a RADIUS server sequence by ID.

@param id id path parameter.
*/
func (s *RadiusServerSequenceService) GetRadiusServerSequenceByID(id string) (*ResponseRadiusServerSequenceGetRadiusServerSequenceByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/radiusserversequence/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseRadiusServerSequenceGetRadiusServerSequenceByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetRadiusServerSequenceById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseRadiusServerSequenceGetRadiusServerSequenceByID)
	return result, response, err

}

//GetRadiusServerSequence Get all RADIUS server sequences
/* This API allows the client to get all the RADIUS server sequences.

@param getRadiusServerSequenceQueryParams Filtering parameter
*/
func (s *RadiusServerSequenceService) GetRadiusServerSequence(getRadiusServerSequenceQueryParams *GetRadiusServerSequenceQueryParams) (*ResponseRadiusServerSequenceGetRadiusServerSequence, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/radiusserversequence"

	queryString, _ := query.Values(getRadiusServerSequenceQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseRadiusServerSequenceGetRadiusServerSequence{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetRadiusServerSequence")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseRadiusServerSequenceGetRadiusServerSequence)
	return result, response, err

}

//GetVersion Get RADIUS server sequence version information
/* This API helps to retrieve the version information related to the RADIUS server sequence.

 */
func (s *RadiusServerSequenceService) GetVersion() (*ResponseRadiusServerSequenceGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/radiusserversequence/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseRadiusServerSequenceGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseRadiusServerSequenceGetVersion)
	return result, response, err

}

//CreateRadiusServerSequence Create RADIUS server sequence
/* This API creates a RADIUS server sequence.

 */
func (s *RadiusServerSequenceService) CreateRadiusServerSequence(requestRadiusServerSequenceCreateRadiusServerSequence *RequestRadiusServerSequenceCreateRadiusServerSequence) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/radiusserversequence"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestRadiusServerSequenceCreateRadiusServerSequence).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateRadiusServerSequence")
	}

	return response, err

}

//UpdateRadiusServerSequenceByID Update RADIUS server sequence
/* This API allows the client to update a RADIUS server sequence.

@param id id path parameter.
*/
func (s *RadiusServerSequenceService) UpdateRadiusServerSequenceByID(id string, requestRadiusServerSequenceUpdateRadiusServerSequenceById *RequestRadiusServerSequenceUpdateRadiusServerSequenceByID) (*ResponseRadiusServerSequenceUpdateRadiusServerSequenceByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/radiusserversequence/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestRadiusServerSequenceUpdateRadiusServerSequenceById).
		SetResult(&ResponseRadiusServerSequenceUpdateRadiusServerSequenceByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateRadiusServerSequenceById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseRadiusServerSequenceUpdateRadiusServerSequenceByID)
	return result, response, err

}

//DeleteRadiusServerSequenceByID Delete RADIUS server sequence
/* This API deletes a RADIUS server sequence.

@param id id path parameter.
*/
func (s *RadiusServerSequenceService) DeleteRadiusServerSequenceByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/radiusserversequence/{id}"
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
		return response, fmt.Errorf("error with operation DeleteRadiusServerSequenceById")
	}

	getCSFRToken(response.Header())

	return response, err

}
