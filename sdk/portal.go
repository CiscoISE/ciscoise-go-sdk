package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type PortalService service

type GetPortalsQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponsePortalGetPortalByID struct {
	ERSPortal ResponsePortalGetPortalByIDERSPortal `json:"ERSPortal,omitempty"` //
}

type ResponsePortalGetPortalByIDERSPortal struct {
	ID          string                                   `json:"id,omitempty"`          //
	Name        string                                   `json:"name,omitempty"`        //
	Description string                                   `json:"description,omitempty"` //
	PortalType  string                                   `json:"portalType,omitempty"`  //
	Link        ResponsePortalGetPortalByIDERSPortalLink `json:"link,omitempty"`        //
}

type ResponsePortalGetPortalByIDERSPortalLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePortalGetPortals struct {
	SearchResult ResponsePortalGetPortalsSearchResult `json:"SearchResult,omitempty"` //
}

type ResponsePortalGetPortalsSearchResult struct {
	Total        int                                              `json:"total,omitempty"`        //
	Resources    []ResponsePortalGetPortalsSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponsePortalGetPortalsSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponsePortalGetPortalsSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponsePortalGetPortalsSearchResultResources struct {
	ID          string                                            `json:"id,omitempty"`          //
	Name        string                                            `json:"name,omitempty"`        //
	Description string                                            `json:"description,omitempty"` //
	Link        ResponsePortalGetPortalsSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponsePortalGetPortalsSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePortalGetPortalsSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePortalGetPortalsSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePortalGetVersion struct {
	VersionInfo ResponsePortalGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponsePortalGetVersionVersionInfo struct {
	CurrentServerVersion string                                  `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                  `json:"supportedVersions,omitempty"`    //
	Link                 ResponsePortalGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponsePortalGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

//GetPortalByID Get portal by ID
/* This API allows the client to get a portal by ID.

@param id id path parameter.
*/
func (s *PortalService) GetPortalByID(id string) (*ResponsePortalGetPortalByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/portal/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePortalGetPortalByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPortalById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePortalGetPortalByID)
	return result, response, err

}

//GetPortals Get all portals
/* This API allows the client to get all the portals.

Filter:

[name, description]

To search resources by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:

[name, description]

@param getPortalsQueryParams Filtering parameter
*/
func (s *PortalService) GetPortals(getPortalsQueryParams *GetPortalsQueryParams) (*ResponsePortalGetPortals, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/portal"

	queryString, _ := query.Values(getPortalsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponsePortalGetPortals{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPortals")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePortalGetPortals)
	return result, response, err

}

//GetVersion Get portal version information
/* This API helps to retrieve the version information realted to the portal.

 */
func (s *PortalService) GetVersion() (*ResponsePortalGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/portal/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePortalGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePortalGetVersion)
	return result, response, err

}
