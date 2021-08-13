package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type FilterPolicyService service

type GetFilterPolicyQueryParams struct {
	Page int `url:"page,omitempty"` //Page number
	Size int `url:"size,omitempty"` //Number of objects returned per page
}

type ResponseFilterPolicyGetFilterPolicyByID struct {
	ERSFilterPolicy ResponseFilterPolicyGetFilterPolicyByIDERSFilterPolicy `json:"ERSFilterPolicy,omitempty"` //
}

type ResponseFilterPolicyGetFilterPolicyByIDERSFilterPolicy struct {
	Subnet  string `json:"subnet,omitempty"`  // Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
	Domains string `json:"domains,omitempty"` // List of SXP Domains, separated with comma
	Sgt     string `json:"sgt,omitempty"`     // SGT name or ID. At least one of subnet or sgt or vn should be defined
	Vn      string `json:"vn,omitempty"`      // Virtual Network. At least one of subnet or sgt or vn should be defined
}

type ResponseFilterPolicyUpdateFilterPolicyByID struct {
	UpdatedFieldsList ResponseFilterPolicyUpdateFilterPolicyByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseFilterPolicyUpdateFilterPolicyByIDUpdatedFieldsList struct {
	UpdatedField ResponseFilterPolicyUpdateFilterPolicyByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                  `json:"field,omitempty"`        //
	OldValue     string                                                                  `json:"oldValue,omitempty"`     //
	NewValue     string                                                                  `json:"newValue,omitempty"`     //
}

type ResponseFilterPolicyUpdateFilterPolicyByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseFilterPolicyGetFilterPolicy struct {
	SearchResult ResponseFilterPolicyGetFilterPolicySearchResult `json:"SearchResult,omitempty"` //
}

type ResponseFilterPolicyGetFilterPolicySearchResult struct {
	Total        int                                                         `json:"total,omitempty"`        //
	Resources    []ResponseFilterPolicyGetFilterPolicySearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponseFilterPolicyGetFilterPolicySearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponseFilterPolicyGetFilterPolicySearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseFilterPolicyGetFilterPolicySearchResultResources struct {
	ID          string                                                       `json:"id,omitempty"`          //
	Name        string                                                       `json:"name,omitempty"`        //
	Description string                                                       `json:"description,omitempty"` //
	Link        ResponseFilterPolicyGetFilterPolicySearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseFilterPolicyGetFilterPolicySearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseFilterPolicyGetFilterPolicySearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseFilterPolicyGetFilterPolicySearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseFilterPolicyGetVersion struct {
	VersionInfo ResponseFilterPolicyGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseFilterPolicyGetVersionVersionInfo struct {
	CurrentServerVersion string                                        `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                        `json:"supportedVersions,omitempty"`    //
	Link                 ResponseFilterPolicyGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseFilterPolicyGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestFilterPolicyUpdateFilterPolicyByID struct {
	ERSFilterPolicy *RequestFilterPolicyUpdateFilterPolicyByIDERSFilterPolicy `json:"ERSFilterPolicy,omitempty"` //
}

type RequestFilterPolicyUpdateFilterPolicyByIDERSFilterPolicy struct {
	Subnet  string `json:"subnet,omitempty"`  // Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
	Domains string `json:"domains,omitempty"` // List of SXP Domains, separated with comma
	Sgt     string `json:"sgt,omitempty"`     // SGT name or ID. At least one of subnet or sgt or vn should be defined
	Vn      string `json:"vn,omitempty"`      // Virtual Network. At least one of subnet or sgt or vn should be defined
}

type RequestFilterPolicyCreateFilterPolicy struct {
	ERSFilterPolicy *RequestFilterPolicyCreateFilterPolicyERSFilterPolicy `json:"ERSFilterPolicy,omitempty"` //
}

type RequestFilterPolicyCreateFilterPolicyERSFilterPolicy struct {
	Subnet  string `json:"subnet,omitempty"`  // Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
	Domains string `json:"domains,omitempty"` // List of SXP Domains, separated with comma
	Sgt     string `json:"sgt,omitempty"`     // SGT name or ID. At least one of subnet or sgt or vn should be defined
	Vn      string `json:"vn,omitempty"`      // Virtual Network. At least one of subnet or sgt or vn should be defined
}

//GetFilterPolicyByID Get filter policy by ID
/* This API allows the client to get a filter policy by ID.

@param id id path parameter.
*/
func (s *FilterPolicyService) GetFilterPolicyByID(id string) (*ResponseFilterPolicyGetFilterPolicyByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/filterpolicy/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseFilterPolicyGetFilterPolicyByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetFilterPolicyById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseFilterPolicyGetFilterPolicyByID)
	return result, response, err

}

//GetFilterPolicy Get all filter policies
/* This API allows the client to get all the filter policies.

@param getFilterPolicyQueryParams Filtering parameter
*/
func (s *FilterPolicyService) GetFilterPolicy(getFilterPolicyQueryParams *GetFilterPolicyQueryParams) (*ResponseFilterPolicyGetFilterPolicy, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/filterpolicy"

	queryString, _ := query.Values(getFilterPolicyQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseFilterPolicyGetFilterPolicy{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetFilterPolicy")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseFilterPolicyGetFilterPolicy)
	return result, response, err

}

//GetVersion Get filter policy version information
/* This API helps to retrieve the version information related to the filter policy.

 */
func (s *FilterPolicyService) GetVersion() (*ResponseFilterPolicyGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/filterpolicy/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseFilterPolicyGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseFilterPolicyGetVersion)
	return result, response, err

}

//CreateFilterPolicy Create filter policy
/* This API creates a filter policy.

 */
func (s *FilterPolicyService) CreateFilterPolicy(requestFilterPolicyCreateFilterPolicy *RequestFilterPolicyCreateFilterPolicy) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/filterpolicy"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFilterPolicyCreateFilterPolicy).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateFilterPolicy")
	}

	return response, err

}

//UpdateFilterPolicyByID Update filter policy
/* This API allows the client to update a filter policy.

@param id id path parameter.
*/
func (s *FilterPolicyService) UpdateFilterPolicyByID(id string, requestFilterPolicyUpdateFilterPolicyById *RequestFilterPolicyUpdateFilterPolicyByID) (*ResponseFilterPolicyUpdateFilterPolicyByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/filterpolicy/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFilterPolicyUpdateFilterPolicyById).
		SetResult(&ResponseFilterPolicyUpdateFilterPolicyByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateFilterPolicyById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseFilterPolicyUpdateFilterPolicyByID)
	return result, response, err

}

//DeleteFilterPolicyByID Delete filter policy
/* This API deletes a filter policy.

@param id id path parameter.
*/
func (s *FilterPolicyService) DeleteFilterPolicyByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/filterpolicy/{id}"
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
		return response, fmt.Errorf("error with operation DeleteFilterPolicyById")
	}

	getCSFRToken(response.Header())

	return response, err

}
