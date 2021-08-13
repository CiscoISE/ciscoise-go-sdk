package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type IDentityGroupsService service

type GetIDentityGroupsQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseIDentityGroupsGetIDentityGroupByName struct {
	IDentityGroup ResponseIDentityGroupsGetIDentityGroupByNameIDentityGroup `json:"IdentityGroup,omitempty"` //
}

type ResponseIDentityGroupsGetIDentityGroupByNameIDentityGroup struct {
	ID          string                                                        `json:"id,omitempty"`          //
	Name        string                                                        `json:"name,omitempty"`        //
	Description string                                                        `json:"description,omitempty"` //
	Parent      string                                                        `json:"parent,omitempty"`      //
	Link        ResponseIDentityGroupsGetIDentityGroupByNameIDentityGroupLink `json:"link,omitempty"`        //
}

type ResponseIDentityGroupsGetIDentityGroupByNameIDentityGroupLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIDentityGroupsGetIDentityGroupByID struct {
	IDentityGroup ResponseIDentityGroupsGetIDentityGroupByIDIDentityGroup `json:"IdentityGroup,omitempty"` //
}

type ResponseIDentityGroupsGetIDentityGroupByIDIDentityGroup struct {
	ID          string                                                      `json:"id,omitempty"`          //
	Name        string                                                      `json:"name,omitempty"`        //
	Description string                                                      `json:"description,omitempty"` //
	Parent      string                                                      `json:"parent,omitempty"`      //
	Link        ResponseIDentityGroupsGetIDentityGroupByIDIDentityGroupLink `json:"link,omitempty"`        //
}

type ResponseIDentityGroupsGetIDentityGroupByIDIDentityGroupLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIDentityGroupsUpdateIDentityGroupByID struct {
	UpdatedFieldsList ResponseIDentityGroupsUpdateIDentityGroupByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseIDentityGroupsUpdateIDentityGroupByIDUpdatedFieldsList struct {
	UpdatedField ResponseIDentityGroupsUpdateIDentityGroupByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                     `json:"field,omitempty"`        //
	OldValue     string                                                                     `json:"oldValue,omitempty"`     //
	NewValue     string                                                                     `json:"newValue,omitempty"`     //
}

type ResponseIDentityGroupsUpdateIDentityGroupByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseIDentityGroupsGetIDentityGroups struct {
	SearchResult ResponseIDentityGroupsGetIDentityGroupsSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseIDentityGroupsGetIDentityGroupsSearchResult struct {
	Total        int                                                               `json:"total,omitempty"`        //
	Resources    []ResponseIDentityGroupsGetIDentityGroupsSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseIDentityGroupsGetIDentityGroupsSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseIDentityGroupsGetIDentityGroupsSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseIDentityGroupsGetIDentityGroupsSearchResultResources struct {
	ID          string                                                           `json:"id,omitempty"`          //
	Name        string                                                           `json:"name,omitempty"`        //
	Description string                                                           `json:"description,omitempty"` //
	Link        ResponseIDentityGroupsGetIDentityGroupsSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseIDentityGroupsGetIDentityGroupsSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIDentityGroupsGetIDentityGroupsSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIDentityGroupsGetIDentityGroupsSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIDentityGroupsGetVersion struct {
	VersionInfo ResponseIDentityGroupsGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseIDentityGroupsGetVersionVersionInfo struct {
	CurrentServerVersion string                                          `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                          `json:"supportedVersions,omitempty"`    //
	Link                 ResponseIDentityGroupsGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseIDentityGroupsGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestIDentityGroupsUpdateIDentityGroupByID struct {
	IDentityGroup *RequestIDentityGroupsUpdateIDentityGroupByIDIDentityGroup `json:"IdentityGroup,omitempty"` //
}

type RequestIDentityGroupsUpdateIDentityGroupByIDIDentityGroup struct {
	ID          string `json:"id,omitempty"`          //
	Name        string `json:"name,omitempty"`        //
	Description string `json:"description,omitempty"` //
	Parent      string `json:"parent,omitempty"`      //
}

type RequestIDentityGroupsCreateIDentityGroup struct {
	IDentityGroup *RequestIDentityGroupsCreateIDentityGroupIDentityGroup `json:"IdentityGroup,omitempty"` //
}

type RequestIDentityGroupsCreateIDentityGroupIDentityGroup struct {
	Name        string `json:"name,omitempty"`        //
	Description string `json:"description,omitempty"` //
	Parent      string `json:"parent,omitempty"`      //
}

//GetIDentityGroupByName Get identity group by name
/* This API allows the client to get an identity group by name.

@param name name path parameter.
*/
func (s *IDentityGroupsService) GetIDentityGroupByName(name string) (*ResponseIDentityGroupsGetIDentityGroupByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/identitygroup/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIDentityGroupsGetIDentityGroupByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIdentityGroupByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIDentityGroupsGetIDentityGroupByName)
	return result, response, err

}

//GetIDentityGroupByID Get identity group by ID
/* This API allows the client to get an identity group by ID.

@param id id path parameter.
*/
func (s *IDentityGroupsService) GetIDentityGroupByID(id string) (*ResponseIDentityGroupsGetIDentityGroupByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/identitygroup/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIDentityGroupsGetIDentityGroupByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIdentityGroupById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIDentityGroupsGetIDentityGroupByID)
	return result, response, err

}

//GetIDentityGroups Get all identity groups
/* This API allows the client to get all the identity groups.

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

@param getIdentityGroupsQueryParams Filtering parameter
*/
func (s *IDentityGroupsService) GetIDentityGroups(getIdentityGroupsQueryParams *GetIDentityGroupsQueryParams) (*ResponseIDentityGroupsGetIDentityGroups, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/identitygroup"

	queryString, _ := query.Values(getIdentityGroupsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseIDentityGroupsGetIDentityGroups{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIdentityGroups")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIDentityGroupsGetIDentityGroups)
	return result, response, err

}

//GetVersion Get identity group version information
/* This API helps to retrieve the version information related to identity group.

 */
func (s *IDentityGroupsService) GetVersion() (*ResponseIDentityGroupsGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/identitygroup/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIDentityGroupsGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIDentityGroupsGetVersion)
	return result, response, err

}

//CreateIDentityGroup Create identity group
/* This API creates an identity group.

 */
func (s *IDentityGroupsService) CreateIDentityGroup(requestIDentityGroupsCreateIdentityGroup *RequestIDentityGroupsCreateIDentityGroup) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/identitygroup"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIDentityGroupsCreateIdentityGroup).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateIdentityGroup")
	}

	return response, err

}

//UpdateIDentityGroupByID Update identity group
/* This API allows the client to update an identity group.

@param id id path parameter.
*/
func (s *IDentityGroupsService) UpdateIDentityGroupByID(id string, requestIDentityGroupsUpdateIdentityGroupById *RequestIDentityGroupsUpdateIDentityGroupByID) (*ResponseIDentityGroupsUpdateIDentityGroupByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/identitygroup/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIDentityGroupsUpdateIdentityGroupById).
		SetResult(&ResponseIDentityGroupsUpdateIDentityGroupByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateIdentityGroupById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIDentityGroupsUpdateIDentityGroupByID)
	return result, response, err

}
