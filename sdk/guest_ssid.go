package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type GuestSSIDService service

type GetGuestSSIDQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseGuestSSIDGetGuestSSIDByID struct {
	GuestSSID *ResponseGuestSSIDGetGuestSSIDByIDGuestSSID `json:"GuestSSID,omitempty"` //
}

type ResponseGuestSSIDGetGuestSSIDByIDGuestSSID struct {
	ID   string                                          `json:"id,omitempty"`   //
	Name string                                          `json:"name,omitempty"` // Resource Name. Name may contain alphanumeric or any of the following characters [_.-]
	Link *ResponseGuestSSIDGetGuestSSIDByIDGuestSSIDLink `json:"link,omitempty"` //
}

type ResponseGuestSSIDGetGuestSSIDByIDGuestSSIDLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestSSIDUpdateGuestSSIDByID struct {
	UpdatedFieldsList *ResponseGuestSSIDUpdateGuestSSIDByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseGuestSSIDUpdateGuestSSIDByIDUpdatedFieldsList struct {
	UpdatedField *[]ResponseGuestSSIDUpdateGuestSSIDByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                               `json:"field,omitempty"`        //
	OldValue     string                                                               `json:"oldValue,omitempty"`     //
	NewValue     string                                                               `json:"newValue,omitempty"`     //
}

type ResponseGuestSSIDUpdateGuestSSIDByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseGuestSSIDGetGuestSSID struct {
	SearchResult *ResponseGuestSSIDGetGuestSSIDSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseGuestSSIDGetGuestSSIDSearchResult struct {
	Total        *int                                                   `json:"total,omitempty"`        //
	Resources    *[]ResponseGuestSSIDGetGuestSSIDSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponseGuestSSIDGetGuestSSIDSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseGuestSSIDGetGuestSSIDSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseGuestSSIDGetGuestSSIDSearchResultResources struct {
	ID   string                                                  `json:"id,omitempty"`   //
	Name string                                                  `json:"name,omitempty"` //
	Link *ResponseGuestSSIDGetGuestSSIDSearchResultResourcesLink `json:"link,omitempty"` //
}

type ResponseGuestSSIDGetGuestSSIDSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestSSIDGetGuestSSIDSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestSSIDGetGuestSSIDSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestSSIDGetVersion struct {
	VersionInfo *ResponseGuestSSIDGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseGuestSSIDGetVersionVersionInfo struct {
	CurrentServerVersion string                                      `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                      `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseGuestSSIDGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseGuestSSIDGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestGuestSSIDUpdateGuestSSIDByID struct {
	GuestSSID *RequestGuestSSIDUpdateGuestSSIDByIDGuestSSID `json:"GuestSSID,omitempty"` //
}

type RequestGuestSSIDUpdateGuestSSIDByIDGuestSSID struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` // Resource Name. Name may contain alphanumeric or any of the following characters [_.-]
}

type RequestGuestSSIDCreateGuestSSID struct {
	GuestSSID *RequestGuestSSIDCreateGuestSSIDGuestSSID `json:"GuestSSID,omitempty"` //
}

type RequestGuestSSIDCreateGuestSSIDGuestSSID struct {
	Name string `json:"name,omitempty"` // Resource Name. Name may contain alphanumeric or any of the following characters [_.-]
}

//GetGuestSSIDByID Get guest SSID by ID
/* This API allows the client to get a guest SSID by ID.

@param id id path parameter.
*/
func (s *GuestSSIDService) GetGuestSSIDByID(id string) (*ResponseGuestSSIDGetGuestSSIDByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestssid/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseGuestSSIDGetGuestSSIDByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetGuestSsidById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestSSIDGetGuestSSIDByID)
	return result, response, err

}

//GetGuestSSID Get all guest SSIDs
/* This API allows the client to get all the guest SSIDs.

Filter:

[name]

To search resources by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:

[name, description]

@param getGuestSsidQueryParams Filtering parameter
*/
func (s *GuestSSIDService) GetGuestSSID(getGuestSsidQueryParams *GetGuestSSIDQueryParams) (*ResponseGuestSSIDGetGuestSSID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestssid"

	queryString, _ := query.Values(getGuestSsidQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseGuestSSIDGetGuestSSID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetGuestSsid")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestSSIDGetGuestSSID)
	return result, response, err

}

//GetVersion Get guest SSID version information
/* This API helps to retrieve the version information related to the guest SSID.

 */
func (s *GuestSSIDService) GetVersion() (*ResponseGuestSSIDGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestssid/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseGuestSSIDGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestSSIDGetVersion)
	return result, response, err

}

//CreateGuestSSID Create guest SSID
/* This API creates a guest SSID.

 */
func (s *GuestSSIDService) CreateGuestSSID(requestGuestSSIDCreateGuestSsid *RequestGuestSSIDCreateGuestSSID) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestssid"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestGuestSSIDCreateGuestSsid).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateGuestSsid")
	}

	return response, err

}

//UpdateGuestSSIDByID Update guest SSID by ID
/* This API allows the client to update a guest SSID by ID.

@param id id path parameter.
*/
func (s *GuestSSIDService) UpdateGuestSSIDByID(id string, requestGuestSSIDUpdateGuestSsidById *RequestGuestSSIDUpdateGuestSSIDByID) (*ResponseGuestSSIDUpdateGuestSSIDByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestssid/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestGuestSSIDUpdateGuestSsidById).
		SetResult(&ResponseGuestSSIDUpdateGuestSSIDByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateGuestSsidById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestSSIDUpdateGuestSSIDByID)
	return result, response, err

}

//DeleteGuestSSIDByID Delete guest SSID by ID
/* This API deletes a guest SSID by ID.

@param id id path parameter.
*/
func (s *GuestSSIDService) DeleteGuestSSIDByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestssid/{id}"
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
		return response, fmt.Errorf("error with operation DeleteGuestSsidById")
	}

	getCSFRToken(response.Header())

	return response, err

}
