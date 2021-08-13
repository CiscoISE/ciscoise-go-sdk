package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type PortalThemeService service

type GetPortalThemesQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponsePortalThemeGetPortalThemeByID struct {
	PortalTheme ResponsePortalThemeGetPortalThemeByIDPortalTheme `json:"PortalTheme,omitempty"` //
}

type ResponsePortalThemeGetPortalThemeByIDPortalTheme struct {
	ID          string                                               `json:"id,omitempty"`          //
	Name        string                                               `json:"name,omitempty"`        //
	Description string                                               `json:"description,omitempty"` //
	ThemeData   string                                               `json:"themeData,omitempty"`   // Portal Theme for all portals
	Link        ResponsePortalThemeGetPortalThemeByIDPortalThemeLink `json:"link,omitempty"`        //
}

type ResponsePortalThemeGetPortalThemeByIDPortalThemeLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePortalThemeUpdatePortalThemeByID struct {
	UpdatedFieldsList ResponsePortalThemeUpdatePortalThemeByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponsePortalThemeUpdatePortalThemeByIDUpdatedFieldsList struct {
	UpdatedField ResponsePortalThemeUpdatePortalThemeByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                `json:"field,omitempty"`        //
	OldValue     string                                                                `json:"oldValue,omitempty"`     //
	NewValue     string                                                                `json:"newValue,omitempty"`     //
}

type ResponsePortalThemeUpdatePortalThemeByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponsePortalThemeGetPortalThemes struct {
	SearchResult ResponsePortalThemeGetPortalThemesSearchResult `json:"SearchResult,omitempty"` //
}

type ResponsePortalThemeGetPortalThemesSearchResult struct {
	Total        int                                                        `json:"total,omitempty"`        //
	Resources    []ResponsePortalThemeGetPortalThemesSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponsePortalThemeGetPortalThemesSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponsePortalThemeGetPortalThemesSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponsePortalThemeGetPortalThemesSearchResultResources struct {
	ID   string                                                      `json:"id,omitempty"`   //
	Name string                                                      `json:"name,omitempty"` //
	Link ResponsePortalThemeGetPortalThemesSearchResultResourcesLink `json:"link,omitempty"` //
}

type ResponsePortalThemeGetPortalThemesSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePortalThemeGetPortalThemesSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePortalThemeGetPortalThemesSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePortalThemeGetVersion struct {
	VersionInfo ResponsePortalThemeGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponsePortalThemeGetVersionVersionInfo struct {
	CurrentServerVersion string                                       `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                       `json:"supportedVersions,omitempty"`    //
	Link                 ResponsePortalThemeGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponsePortalThemeGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestPortalThemeUpdatePortalThemeByID struct {
	PortalTheme *RequestPortalThemeUpdatePortalThemeByIDPortalTheme `json:"PortalTheme,omitempty"` //
}

type RequestPortalThemeUpdatePortalThemeByIDPortalTheme struct {
	ID          string `json:"id,omitempty"`          //
	Name        string `json:"name,omitempty"`        //
	Description string `json:"description,omitempty"` //
	ThemeData   string `json:"themeData,omitempty"`   // Portal Theme for all portals
}

type RequestPortalThemeCreatePortalTheme struct {
	PortalTheme *RequestPortalThemeCreatePortalThemePortalTheme `json:"PortalTheme,omitempty"` //
}

type RequestPortalThemeCreatePortalThemePortalTheme struct {
	Name        string `json:"name,omitempty"`        //
	Description string `json:"description,omitempty"` //
	ThemeData   string `json:"themeData,omitempty"`   // Portal Theme for all portals
}

//GetPortalThemeByID Get portal theme by ID
/* This API allows the client to get a portal theme by ID.

@param id id path parameter.
*/
func (s *PortalThemeService) GetPortalThemeByID(id string) (*ResponsePortalThemeGetPortalThemeByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/portaltheme/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePortalThemeGetPortalThemeByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPortalThemeById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePortalThemeGetPortalThemeByID)
	return result, response, err

}

//GetPortalThemes Get all portal themes
/* This API allows the client to get all the portal themes.

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

@param getPortalThemesQueryParams Filtering parameter
*/
func (s *PortalThemeService) GetPortalThemes(getPortalThemesQueryParams *GetPortalThemesQueryParams) (*ResponsePortalThemeGetPortalThemes, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/portaltheme"

	queryString, _ := query.Values(getPortalThemesQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponsePortalThemeGetPortalThemes{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPortalThemes")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePortalThemeGetPortalThemes)
	return result, response, err

}

//GetVersion Get portal theme version information
/* This API helps to retrieve the version information related to the portal theme.

 */
func (s *PortalThemeService) GetVersion() (*ResponsePortalThemeGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/portaltheme/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePortalThemeGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePortalThemeGetVersion)
	return result, response, err

}

//CreatePortalTheme Create portal theme
/* This API creates a portal theme.

 */
func (s *PortalThemeService) CreatePortalTheme(requestPortalThemeCreatePortalTheme *RequestPortalThemeCreatePortalTheme) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/portaltheme"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestPortalThemeCreatePortalTheme).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreatePortalTheme")
	}

	return response, err

}

//UpdatePortalThemeByID Update portal theme by ID
/* This API allows the client to update a portal theme by ID.

@param id id path parameter.
*/
func (s *PortalThemeService) UpdatePortalThemeByID(id string, requestPortalThemeUpdatePortalThemeById *RequestPortalThemeUpdatePortalThemeByID) (*ResponsePortalThemeUpdatePortalThemeByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/portaltheme/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestPortalThemeUpdatePortalThemeById).
		SetResult(&ResponsePortalThemeUpdatePortalThemeByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdatePortalThemeById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePortalThemeUpdatePortalThemeByID)
	return result, response, err

}

//DeletePortalThemeByID Delete portal theme by ID
/* This API deletes a portal theme by ID.

@param id id path parameter.
*/
func (s *PortalThemeService) DeletePortalThemeByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/portaltheme/{id}"
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
		return response, fmt.Errorf("error with operation DeletePortalThemeById")
	}

	getCSFRToken(response.Header())

	return response, err

}
