package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type GuestLocationService service

type GetGuestLocationQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseGuestLocationGetGuestLocationByID struct {
	LocationIDentification ResponseGuestLocationGetGuestLocationByIDLocationIDentification `json:"LocationIdentification,omitempty"` //
}

type ResponseGuestLocationGetGuestLocationByIDLocationIDentification struct {
	ID   string                                                              `json:"id,omitempty"`   //
	Name string                                                              `json:"name,omitempty"` //
	Link ResponseGuestLocationGetGuestLocationByIDLocationIDentificationLink `json:"link,omitempty"` //
}

type ResponseGuestLocationGetGuestLocationByIDLocationIDentificationLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestLocationGetGuestLocation struct {
	SearchResult ResponseGuestLocationGetGuestLocationSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseGuestLocationGetGuestLocationSearchResult struct {
	Total        int                                                             `json:"total,omitempty"`        //
	Resources    []ResponseGuestLocationGetGuestLocationSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseGuestLocationGetGuestLocationSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseGuestLocationGetGuestLocationSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseGuestLocationGetGuestLocationSearchResultResources struct {
	ID          string                                                         `json:"id,omitempty"`          //
	Name        string                                                         `json:"name,omitempty"`        //
	Description string                                                         `json:"description,omitempty"` //
	Link        ResponseGuestLocationGetGuestLocationSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseGuestLocationGetGuestLocationSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestLocationGetGuestLocationSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestLocationGetGuestLocationSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestLocationGetVersion struct {
	VersionInfo ResponseGuestLocationGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseGuestLocationGetVersionVersionInfo struct {
	CurrentServerVersion string                                         `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                         `json:"supportedVersions,omitempty"`    //
	Link                 ResponseGuestLocationGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseGuestLocationGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

//GetGuestLocationByID Get guest location by ID
/* This API allows the client to get a guest location by ID.

@param id id path parameter.
*/
func (s *GuestLocationService) GetGuestLocationByID(id string) (*ResponseGuestLocationGetGuestLocationByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestlocation/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseGuestLocationGetGuestLocationByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetGuestLocationById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestLocationGetGuestLocationByID)
	return result, response, err

}

//GetGuestLocation Get all guest locations
/* This API allows the client to get all the guest locations.

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

@param getGuestLocationQueryParams Filtering parameter
*/
func (s *GuestLocationService) GetGuestLocation(getGuestLocationQueryParams *GetGuestLocationQueryParams) (*ResponseGuestLocationGetGuestLocation, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestlocation"

	queryString, _ := query.Values(getGuestLocationQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseGuestLocationGetGuestLocation{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetGuestLocation")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestLocationGetGuestLocation)
	return result, response, err

}

//GetVersion Get guest location version information
/* This API helps to retrieve the version information related to the guest location.

 */
func (s *GuestLocationService) GetVersion() (*ResponseGuestLocationGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestlocation/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseGuestLocationGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestLocationGetVersion)
	return result, response, err

}
