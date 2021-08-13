package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type DownloadableACLService service

type GetDownloadableACLQueryParams struct {
	Page int `url:"page,omitempty"` //Page number
	Size int `url:"size,omitempty"` //Number of objects returned per page
}

type ResponseDownloadableACLGetDownloadableACLByID struct {
	DownloadableACL ResponseDownloadableACLGetDownloadableACLByIDDownloadableACL `json:"DownloadableAcl,omitempty"` //
}

type ResponseDownloadableACLGetDownloadableACLByIDDownloadableACL struct {
	ID          string                                                            `json:"id,omitempty"`          //
	Name        string                                                            `json:"name,omitempty"`        // Resource Name. Name may contain alphanumeric or any of the following characters [_.-]
	Description string                                                            `json:"description,omitempty"` // Use the string \\n for a newline
	Dacl        string                                                            `json:"dacl,omitempty"`        // The DACL Content. Use the string \\n for a newline
	DaclType    string                                                            `json:"daclType,omitempty"`    // Allowed values: - IPV4, - IPV6, - IP_AGNOSTIC
	Link1       ResponseDownloadableACLGetDownloadableACLByIDDownloadableACLLink1 `json:"link1,omitempty"`       //
}

type ResponseDownloadableACLGetDownloadableACLByIDDownloadableACLLink1 struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseDownloadableACLUpdateDownloadableACLByID struct {
	UpdatedFieldsList ResponseDownloadableACLUpdateDownloadableACLByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseDownloadableACLUpdateDownloadableACLByIDUpdatedFieldsList struct {
	UpdatedField ResponseDownloadableACLUpdateDownloadableACLByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                        `json:"field,omitempty"`        //
	OldValue     string                                                                        `json:"oldValue,omitempty"`     //
	NewValue     string                                                                        `json:"newValue,omitempty"`     //
}

type ResponseDownloadableACLUpdateDownloadableACLByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseDownloadableACLGetDownloadableACL struct {
	SearchResult ResponseDownloadableACLGetDownloadableACLSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseDownloadableACLGetDownloadableACLSearchResult struct {
	Total        int                                                                 `json:"total,omitempty"`        //
	Resources    []ResponseDownloadableACLGetDownloadableACLSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseDownloadableACLGetDownloadableACLSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseDownloadableACLGetDownloadableACLSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseDownloadableACLGetDownloadableACLSearchResultResources struct {
	ID          string                                                             `json:"id,omitempty"`          //
	Name        string                                                             `json:"name,omitempty"`        //
	Description string                                                             `json:"description,omitempty"` //
	Link        ResponseDownloadableACLGetDownloadableACLSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseDownloadableACLGetDownloadableACLSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseDownloadableACLGetDownloadableACLSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseDownloadableACLGetDownloadableACLSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseDownloadableACLGetVersion struct {
	VersionInfo ResponseDownloadableACLGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseDownloadableACLGetVersionVersionInfo struct {
	CurrentServerVersion string                                           `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                           `json:"supportedVersions,omitempty"`    //
	Link                 ResponseDownloadableACLGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseDownloadableACLGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestDownloadableACLUpdateDownloadableACLByID struct {
	DownloadableACL *RequestDownloadableACLUpdateDownloadableACLByIDDownloadableACL `json:"DownloadableAcl,omitempty"` //
}

type RequestDownloadableACLUpdateDownloadableACLByIDDownloadableACL struct {
	ID          string `json:"id,omitempty"`          //
	Name        string `json:"name,omitempty"`        // Resource Name. Name may contain alphanumeric or any of the following characters [_.-]
	Description string `json:"description,omitempty"` // Use the string \\n for a newline
	Dacl        string `json:"dacl,omitempty"`        // The DACL Content. Use the string \\n for a newline
	DaclType    string `json:"daclType,omitempty"`    // Allowed values: - IPV4, - IPV6, - IP_AGNOSTIC
}

type RequestDownloadableACLCreateDownloadableACL struct {
	DownloadableACL *RequestDownloadableACLCreateDownloadableACLDownloadableACL `json:"DownloadableAcl,omitempty"` //
}

type RequestDownloadableACLCreateDownloadableACLDownloadableACL struct {
	Name        string `json:"name,omitempty"`        // Resource Name. Name may contain alphanumeric or any of the following characters [_.-]
	Description string `json:"description,omitempty"` // Use the string \\n for a newline
	Dacl        string `json:"dacl,omitempty"`        // The DACL Content. Use the string \\n for a newline
	DaclType    string `json:"daclType,omitempty"`    // Allowed values: - IPV4, - IPV6, - IP_AGNOSTIC
}

//GetDownloadableACLByID Get downloadable ACL by ID
/* This API allows the client to get a downloadable ACL by ID.

@param id id path parameter.
*/
func (s *DownloadableACLService) GetDownloadableACLByID(id string) (*ResponseDownloadableACLGetDownloadableACLByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/downloadableacl/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDownloadableACLGetDownloadableACLByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDownloadableAclById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDownloadableACLGetDownloadableACLByID)
	return result, response, err

}

//GetDownloadableACL Get all downloadable ACLs
/* This API allows the client to get all downloadable ACLs.

@param getDownloadableAclQueryParams Filtering parameter
*/
func (s *DownloadableACLService) GetDownloadableACL(getDownloadableAclQueryParams *GetDownloadableACLQueryParams) (*ResponseDownloadableACLGetDownloadableACL, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/downloadableacl"

	queryString, _ := query.Values(getDownloadableAclQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDownloadableACLGetDownloadableACL{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDownloadableAcl")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDownloadableACLGetDownloadableACL)
	return result, response, err

}

//GetVersion Get downloadable ACL version information
/* This API helps to retrieve the version information related to the downloadable ACL.

 */
func (s *DownloadableACLService) GetVersion() (*ResponseDownloadableACLGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/downloadableacl/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDownloadableACLGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDownloadableACLGetVersion)
	return result, response, err

}

//CreateDownloadableACL Create downloadable ACL
/* This API creates a downloadable ACL.

 */
func (s *DownloadableACLService) CreateDownloadableACL(requestDownloadableACLCreateDownloadableAcl *RequestDownloadableACLCreateDownloadableACL) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/downloadableacl"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDownloadableACLCreateDownloadableAcl).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateDownloadableAcl")
	}

	return response, err

}

//UpdateDownloadableACLByID Update downloadable ACL
/* This API allows the client to update a downloadable ACL.

@param id id path parameter.
*/
func (s *DownloadableACLService) UpdateDownloadableACLByID(id string, requestDownloadableACLUpdateDownloadableAclById *RequestDownloadableACLUpdateDownloadableACLByID) (*ResponseDownloadableACLUpdateDownloadableACLByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/downloadableacl/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDownloadableACLUpdateDownloadableAclById).
		SetResult(&ResponseDownloadableACLUpdateDownloadableACLByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDownloadableAclById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDownloadableACLUpdateDownloadableACLByID)
	return result, response, err

}

//DeleteDownloadableACLByID Delete downloadable ACL
/* This API deletes a downloadable ACL.

@param id id path parameter.
*/
func (s *DownloadableACLService) DeleteDownloadableACLByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/downloadableacl/{id}"
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
		return response, fmt.Errorf("error with operation DeleteDownloadableAclById")
	}

	getCSFRToken(response.Header())

	return response, err

}
