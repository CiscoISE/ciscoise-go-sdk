package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type IDentitySequenceService service

type GetIDentitySequenceQueryParams struct {
	Page int `url: page,omitempty` //Page number
	Size int `url: size,omitempty` //Number of objects returned per page
}

type ResponseIDentitySequenceGetIDentitySequenceByName struct {
	IDStoreSequence ResponseIDentitySequenceGetIDentitySequenceByNameIDStoreSequence `json:"IdStoreSequence,omitempty"` //
}

type ResponseIDentitySequenceGetIDentitySequenceByNameIDStoreSequence struct {
	ID                               string                                                                      `json:"id,omitempty"`                               //
	Name                             string                                                                      `json:"name,omitempty"`                             //
	Description                      string                                                                      `json:"description,omitempty"`                      //
	Parent                           string                                                                      `json:"parent,omitempty"`                           //
	IDSeqItem                        []ResponseIDentitySequenceGetIDentitySequenceByNameIDStoreSequenceIDSeqItem `json:"idSeqItem,omitempty"`                        //
	CertificateAuthenticationProfile string                                                                      `json:"certificateAuthenticationProfile,omitempty"` //
	BreakOnStoreFail                 bool                                                                        `json:"breakOnStoreFail,omitempty"`                 //
	Link                             ResponseIDentitySequenceGetIDentitySequenceByNameIDStoreSequenceLink        `json:"link,omitempty"`                             //
}

type ResponseIDentitySequenceGetIDentitySequenceByNameIDStoreSequenceIDSeqItem struct {
	IDstore string `json:"idstore,omitempty"` //
	Order   int    `json:"order,omitempty"`   //
}

type ResponseIDentitySequenceGetIDentitySequenceByNameIDStoreSequenceLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIDentitySequenceGetIDentitySequenceByID struct {
	IDStoreSequence ResponseIDentitySequenceGetIDentitySequenceByIDIDStoreSequence `json:"IdStoreSequence,omitempty"` //
}

type ResponseIDentitySequenceGetIDentitySequenceByIDIDStoreSequence struct {
	ID                               string                                                                    `json:"id,omitempty"`                               //
	Name                             string                                                                    `json:"name,omitempty"`                             //
	Description                      string                                                                    `json:"description,omitempty"`                      //
	Parent                           string                                                                    `json:"parent,omitempty"`                           //
	IDSeqItem                        []ResponseIDentitySequenceGetIDentitySequenceByIDIDStoreSequenceIDSeqItem `json:"idSeqItem,omitempty"`                        //
	CertificateAuthenticationProfile string                                                                    `json:"certificateAuthenticationProfile,omitempty"` //
	BreakOnStoreFail                 bool                                                                      `json:"breakOnStoreFail,omitempty"`                 //
	Link                             ResponseIDentitySequenceGetIDentitySequenceByIDIDStoreSequenceLink        `json:"link,omitempty"`                             //
}

type ResponseIDentitySequenceGetIDentitySequenceByIDIDStoreSequenceIDSeqItem struct {
	IDstore string `json:"idstore,omitempty"` //
	Order   int    `json:"order,omitempty"`   //
}

type ResponseIDentitySequenceGetIDentitySequenceByIDIDStoreSequenceLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIDentitySequenceUpdateIDentitySequenceByID struct {
	UpdatedFieldsList ResponseIDentitySequenceUpdateIDentitySequenceByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseIDentitySequenceUpdateIDentitySequenceByIDUpdatedFieldsList struct {
	UpdatedField ResponseIDentitySequenceUpdateIDentitySequenceByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                          `json:"field,omitempty"`        //
	OldValue     string                                                                          `json:"oldValue,omitempty"`     //
	NewValue     string                                                                          `json:"newValue,omitempty"`     //
}

type ResponseIDentitySequenceUpdateIDentitySequenceByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseIDentitySequenceGetIDentitySequence struct {
	SearchResult ResponseIDentitySequenceGetIDentitySequenceSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseIDentitySequenceGetIDentitySequenceSearchResult struct {
	Total        int                                                                   `json:"total,omitempty"`        //
	Resources    []ResponseIDentitySequenceGetIDentitySequenceSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseIDentitySequenceGetIDentitySequenceSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseIDentitySequenceGetIDentitySequenceSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseIDentitySequenceGetIDentitySequenceSearchResultResources struct {
	ID          string                                                               `json:"id,omitempty"`          //
	Name        string                                                               `json:"name,omitempty"`        //
	Description string                                                               `json:"description,omitempty"` //
	Link        ResponseIDentitySequenceGetIDentitySequenceSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseIDentitySequenceGetIDentitySequenceSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIDentitySequenceGetIDentitySequenceSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIDentitySequenceGetIDentitySequenceSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIDentitySequenceGetVersion struct {
	VersionInfo ResponseIDentitySequenceGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseIDentitySequenceGetVersionVersionInfo struct {
	CurrentServerVersion string                                            `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                            `json:"supportedVersions,omitempty"`    //
	Link                 ResponseIDentitySequenceGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseIDentitySequenceGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestIDentitySequenceUpdateIDentitySequenceByID struct {
	IDStoreSequence *RequestIDentitySequenceUpdateIDentitySequenceByIDIDStoreSequence `json:"IdStoreSequence,omitempty"` //
}

type RequestIDentitySequenceUpdateIDentitySequenceByIDIDStoreSequence struct {
	ID                               string                                                                       `json:"id,omitempty"`                               //
	Name                             string                                                                       `json:"name,omitempty"`                             //
	Description                      string                                                                       `json:"description,omitempty"`                      //
	Parent                           string                                                                       `json:"parent,omitempty"`                           //
	IDSeqItem                        *[]RequestIDentitySequenceUpdateIDentitySequenceByIDIDStoreSequenceIDSeqItem `json:"idSeqItem,omitempty"`                        //
	CertificateAuthenticationProfile string                                                                       `json:"certificateAuthenticationProfile,omitempty"` //
	BreakOnStoreFail                 *bool                                                                        `json:"breakOnStoreFail,omitempty"`                 //
}

type RequestIDentitySequenceUpdateIDentitySequenceByIDIDStoreSequenceIDSeqItem struct {
	IDstore string `json:"idstore,omitempty"` //
	Order   *int   `json:"order,omitempty"`   //
}

type RequestIDentitySequenceCreateIDentitySequence struct {
	IDStoreSequence *RequestIDentitySequenceCreateIDentitySequenceIDStoreSequence `json:"IdStoreSequence,omitempty"` //
}

type RequestIDentitySequenceCreateIDentitySequenceIDStoreSequence struct {
	Name                             string                                                                   `json:"name,omitempty"`                             //
	Description                      string                                                                   `json:"description,omitempty"`                      //
	Parent                           string                                                                   `json:"parent,omitempty"`                           //
	IDSeqItem                        *[]RequestIDentitySequenceCreateIDentitySequenceIDStoreSequenceIDSeqItem `json:"idSeqItem,omitempty"`                        //
	CertificateAuthenticationProfile string                                                                   `json:"certificateAuthenticationProfile,omitempty"` //
	BreakOnStoreFail                 *bool                                                                    `json:"breakOnStoreFail,omitempty"`                 //
}

type RequestIDentitySequenceCreateIDentitySequenceIDStoreSequenceIDSeqItem struct {
	IDstore string `json:"idstore,omitempty"` //
	Order   *int   `json:"order,omitempty"`   //
}

//GetIDentitySequenceByName Get identity sequence by name
/* This API allows the client to get an identity sequence by name.

@param name name path parameter.
*/
func (s *IDentitySequenceService) GetIDentitySequenceByName(name string) (*ResponseIDentitySequenceGetIDentitySequenceByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/idstoresequence/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIDentitySequenceGetIDentitySequenceByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIdentitySequenceByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIDentitySequenceGetIDentitySequenceByName)
	return result, response, err

}

//GetIDentitySequenceByID Get identity sequence by ID
/* This API allows the client to get an identity sequence by ID.

@param id id path parameter.
*/
func (s *IDentitySequenceService) GetIDentitySequenceByID(id string) (*ResponseIDentitySequenceGetIDentitySequenceByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/idstoresequence/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIDentitySequenceGetIDentitySequenceByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIdentitySequenceById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIDentitySequenceGetIDentitySequenceByID)
	return result, response, err

}

//GetIDentitySequence Get all identity sequences
/* This API allows the client to get all the identity sequences.

@param getIdentitySequenceQueryParams Filtering parameter
*/
func (s *IDentitySequenceService) GetIDentitySequence(getIdentitySequenceQueryParams *GetIDentitySequenceQueryParams) (*ResponseIDentitySequenceGetIDentitySequence, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/idstoresequence"

	queryString, _ := query.Values(getIdentitySequenceQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseIDentitySequenceGetIDentitySequence{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIdentitySequence")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIDentitySequenceGetIDentitySequence)
	return result, response, err

}

//GetVersion Get identity sequence version information
/* This API helps to retrieve the version information related to the identity sequence.

 */
func (s *IDentitySequenceService) GetVersion() (*ResponseIDentitySequenceGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/idstoresequence/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIDentitySequenceGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIDentitySequenceGetVersion)
	return result, response, err

}

//CreateIDentitySequence Create identity sequence
/* This API creates an identity sequence.

 */
func (s *IDentitySequenceService) CreateIDentitySequence(requestIDentitySequenceCreateIdentitySequence *RequestIDentitySequenceCreateIDentitySequence) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/idstoresequence"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIDentitySequenceCreateIdentitySequence).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateIdentitySequence")
	}

	return response, err

}

//UpdateIDentitySequenceByID Update identity sequence
/* This API allows the client to update an identity sequence.
Partial update is not supported

@param id id path parameter.
*/
func (s *IDentitySequenceService) UpdateIDentitySequenceByID(id string, requestIDentitySequenceUpdateIdentitySequenceById *RequestIDentitySequenceUpdateIDentitySequenceByID) (*ResponseIDentitySequenceUpdateIDentitySequenceByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/idstoresequence/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIDentitySequenceUpdateIdentitySequenceById).
		SetResult(&ResponseIDentitySequenceUpdateIDentitySequenceByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateIdentitySequenceById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIDentitySequenceUpdateIDentitySequenceByID)
	return result, response, err

}

//DeleteIDentitySequenceByID Delete identity sequence
/* This API deletes an identity sequence.

@param id id path parameter.
*/
func (s *IDentitySequenceService) DeleteIDentitySequenceByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/idstoresequence/{id}"
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
		return response, fmt.Errorf("error with operation DeleteIdentitySequenceById")
	}

	getCSFRToken(response.Header())

	return response, err

}
