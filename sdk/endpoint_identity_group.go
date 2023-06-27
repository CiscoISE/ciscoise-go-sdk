package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type EndpointIDentityGroupService service

type GetEndpointGroupsQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseEndpointIDentityGroupGetEndpointGroupByName struct {
	EndPointGroup *ResponseEndpointIDentityGroupGetEndpointGroupByNameEndPointGroup `json:"EndPointGroup,omitempty"` //
}

type ResponseEndpointIDentityGroupGetEndpointGroupByNameEndPointGroup struct {
	ID            string                                                                `json:"id,omitempty"`            //
	Name          string                                                                `json:"name,omitempty"`          //
	Description   string                                                                `json:"description,omitempty"`   //
	SystemDefined *bool                                                                 `json:"systemDefined,omitempty"` //
	Link          *ResponseEndpointIDentityGroupGetEndpointGroupByNameEndPointGroupLink `json:"link,omitempty"`          //
	ParentID      string                                                                `json:"parentId,omitempty"`      //
}

type ResponseEndpointIDentityGroupGetEndpointGroupByNameEndPointGroupLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseEndpointIDentityGroupGetEndpointGroupByID struct {
	EndPointGroup *ResponseEndpointIDentityGroupGetEndpointGroupByIDEndPointGroup `json:"EndPointGroup,omitempty"` //
}

type ResponseEndpointIDentityGroupGetEndpointGroupByIDEndPointGroup struct {
	ID            string                                                              `json:"id,omitempty"`            //
	Name          string                                                              `json:"name,omitempty"`          //
	Description   string                                                              `json:"description,omitempty"`   //
	SystemDefined *bool                                                               `json:"systemDefined,omitempty"` //
	Link          *ResponseEndpointIDentityGroupGetEndpointGroupByIDEndPointGroupLink `json:"link,omitempty"`          //
	ParentID      string                                                              `json:"parentId,omitempty"`      //
}

type ResponseEndpointIDentityGroupGetEndpointGroupByIDEndPointGroupLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseEndpointIDentityGroupUpdateEndpointGroupByID struct {
	UpdatedFieldsList *ResponseEndpointIDentityGroupUpdateEndpointGroupByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseEndpointIDentityGroupUpdateEndpointGroupByIDUpdatedFieldsList struct {
	UpdatedField *[]ResponseEndpointIDentityGroupUpdateEndpointGroupByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                               `json:"field,omitempty"`        //
	OldValue     string                                                                               `json:"oldValue,omitempty"`     //
	NewValue     string                                                                               `json:"newValue,omitempty"`     //
}

type ResponseEndpointIDentityGroupUpdateEndpointGroupByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseEndpointIDentityGroupGetEndpointGroups struct {
	SearchResult *ResponseEndpointIDentityGroupGetEndpointGroupsSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseEndpointIDentityGroupGetEndpointGroupsSearchResult struct {
	Total        *int                                                                    `json:"total,omitempty"`        //
	Resources    *[]ResponseEndpointIDentityGroupGetEndpointGroupsSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponseEndpointIDentityGroupGetEndpointGroupsSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseEndpointIDentityGroupGetEndpointGroupsSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseEndpointIDentityGroupGetEndpointGroupsSearchResultResources struct {
	ID          string                                                                   `json:"id,omitempty"`          //
	Name        string                                                                   `json:"name,omitempty"`        //
	Description string                                                                   `json:"description,omitempty"` //
	Link        *ResponseEndpointIDentityGroupGetEndpointGroupsSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseEndpointIDentityGroupGetEndpointGroupsSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseEndpointIDentityGroupGetEndpointGroupsSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseEndpointIDentityGroupGetEndpointGroupsSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseEndpointIDentityGroupGetVersion struct {
	VersionInfo *ResponseEndpointIDentityGroupGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseEndpointIDentityGroupGetVersionVersionInfo struct {
	CurrentServerVersion string                                                  `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                                  `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseEndpointIDentityGroupGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseEndpointIDentityGroupGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestEndpointIDentityGroupUpdateEndpointGroupByID struct {
	EndPointGroup *RequestEndpointIDentityGroupUpdateEndpointGroupByIDEndPointGroup `json:"EndPointGroup,omitempty"` //
}

type RequestEndpointIDentityGroupUpdateEndpointGroupByIDEndPointGroup struct {
	ID            string `json:"id,omitempty"`            //
	Name          string `json:"name,omitempty"`          //
	Description   string `json:"description,omitempty"`   //
	SystemDefined *bool  `json:"systemDefined,omitempty"` //
}

type RequestEndpointIDentityGroupCreateEndpointGroup struct {
	EndPointGroup *RequestEndpointIDentityGroupCreateEndpointGroupEndPointGroup `json:"EndPointGroup,omitempty"` //
}

type RequestEndpointIDentityGroupCreateEndpointGroupEndPointGroup struct {
	Name          string `json:"name,omitempty"`          //
	Description   string `json:"description,omitempty"`   //
	SystemDefined *bool  `json:"systemDefined,omitempty"` //
	ParentID      string `json:"parentId,omitempty"`      //
}

//GetEndpointGroupByName Get endpoint identity group by name
/* This API allows the client to get an endpoint identity group by name.

@param name name path parameter.
*/
func (s *EndpointIDentityGroupService) GetEndpointGroupByName(name string) (*ResponseEndpointIDentityGroupGetEndpointGroupByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpointgroup/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEndpointIDentityGroupGetEndpointGroupByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetEndpointGroupByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointIDentityGroupGetEndpointGroupByName)
	return result, response, err

}

//GetEndpointGroupByID Get endpoint identity group by ID
/* This API allows the client to get an endpoint identity group by ID.

@param id id path parameter.
*/
func (s *EndpointIDentityGroupService) GetEndpointGroupByID(id string) (*ResponseEndpointIDentityGroupGetEndpointGroupByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpointgroup/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEndpointIDentityGroupGetEndpointGroupByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetEndpointGroupById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointIDentityGroupGetEndpointGroupByID)
	return result, response, err

}

//GetEndpointGroups Get all endpoint identity groups
/* This API allows the client to get all the endpoint identity groups.

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

@param getEndpointGroupsQueryParams Filtering parameter
*/
func (s *EndpointIDentityGroupService) GetEndpointGroups(getEndpointGroupsQueryParams *GetEndpointGroupsQueryParams) (*ResponseEndpointIDentityGroupGetEndpointGroups, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpointgroup"

	queryString, _ := query.Values(getEndpointGroupsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEndpointIDentityGroupGetEndpointGroups{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetEndpointGroups")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointIDentityGroupGetEndpointGroups)
	return result, response, err

}

//GetVersion Get endpoint identity group version information
/* This API helps to retrieve the version information related to the endpoint identity group.

 */
func (s *EndpointIDentityGroupService) GetVersion() (*ResponseEndpointIDentityGroupGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpointgroup/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEndpointIDentityGroupGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointIDentityGroupGetVersion)
	return result, response, err

}

//CreateEndpointGroup Create endpoint identity group
/* This API creates an endpoint identity group.

 */
func (s *EndpointIDentityGroupService) CreateEndpointGroup(requestEndpointIDentityGroupCreateEndpointGroup *RequestEndpointIDentityGroupCreateEndpointGroup) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpointgroup"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEndpointIDentityGroupCreateEndpointGroup).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateEndpointGroup")
	}

	return response, err

}

//UpdateEndpointGroupByID Update endpoint identity group
/* This API allows the client to update an endpoint identity group.

@param id id path parameter.
*/
func (s *EndpointIDentityGroupService) UpdateEndpointGroupByID(id string, requestEndpointIDentityGroupUpdateEndpointGroupById *RequestEndpointIDentityGroupUpdateEndpointGroupByID) (*ResponseEndpointIDentityGroupUpdateEndpointGroupByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpointgroup/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEndpointIDentityGroupUpdateEndpointGroupById).
		SetResult(&ResponseEndpointIDentityGroupUpdateEndpointGroupByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseEndpointIDentityGroupUpdateEndpointGroupByID{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateEndpointGroupById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointIDentityGroupUpdateEndpointGroupByID)
	return result, response, err

}

//DeleteEndpointGroupByID Delete endpoint identity group
/* This API deletes an endpoint identity group.

@param id id path parameter.
*/
func (s *EndpointIDentityGroupService) DeleteEndpointGroupByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpointgroup/{id}"
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
		return response, fmt.Errorf("error with operation DeleteEndpointGroupById")
	}

	getCSFRToken(response.Header())

	return response, err

}
