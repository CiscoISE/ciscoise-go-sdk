package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type NetworkDeviceGroupService service

type GetNetworkDeviceGroupQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseNetworkDeviceGroupGetNetworkDeviceGroupByName struct {
	NetworkDeviceGroup *ResponseNetworkDeviceGroupGetNetworkDeviceGroupByNameNetworkDeviceGroup `json:"NetworkDeviceGroup,omitempty"` //
}

type ResponseNetworkDeviceGroupGetNetworkDeviceGroupByNameNetworkDeviceGroup struct {
	ID          string                                                                       `json:"id,omitempty"`          //
	Name        string                                                                       `json:"name,omitempty"`        //
	Description string                                                                       `json:"description,omitempty"` //
	Link        *ResponseNetworkDeviceGroupGetNetworkDeviceGroupByNameNetworkDeviceGroupLink `json:"link,omitempty"`        //
	Othername   string                                                                       `json:"othername,omitempty"`   //
}

type ResponseNetworkDeviceGroupGetNetworkDeviceGroupByNameNetworkDeviceGroupLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkDeviceGroupGetNetworkDeviceGroupByID struct {
	NetworkDeviceGroup *ResponseNetworkDeviceGroupGetNetworkDeviceGroupByIDNetworkDeviceGroup `json:"NetworkDeviceGroup,omitempty"` //
}

type ResponseNetworkDeviceGroupGetNetworkDeviceGroupByIDNetworkDeviceGroup struct {
	ID          string                                                                     `json:"id,omitempty"`          //
	Name        string                                                                     `json:"name,omitempty"`        //
	Description string                                                                     `json:"description,omitempty"` //
	Link        *ResponseNetworkDeviceGroupGetNetworkDeviceGroupByIDNetworkDeviceGroupLink `json:"link,omitempty"`        //
	Othername   string                                                                     `json:"othername,omitempty"`   //
}

type ResponseNetworkDeviceGroupGetNetworkDeviceGroupByIDNetworkDeviceGroupLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkDeviceGroupUpdateNetworkDeviceGroupByID struct {
	UpdatedFieldsList *ResponseNetworkDeviceGroupUpdateNetworkDeviceGroupByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseNetworkDeviceGroupUpdateNetworkDeviceGroupByIDUpdatedFieldsList struct {
	UpdatedField *[]ResponseNetworkDeviceGroupUpdateNetworkDeviceGroupByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                                 `json:"field,omitempty"`        //
	OldValue     string                                                                                 `json:"oldValue,omitempty"`     //
	NewValue     string                                                                                 `json:"newValue,omitempty"`     //
}

type ResponseNetworkDeviceGroupUpdateNetworkDeviceGroupByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseNetworkDeviceGroupGetNetworkDeviceGroup struct {
	SearchResult *ResponseNetworkDeviceGroupGetNetworkDeviceGroupSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseNetworkDeviceGroupGetNetworkDeviceGroupSearchResult struct {
	Total        *int                                                                     `json:"total,omitempty"`        //
	Resources    *[]ResponseNetworkDeviceGroupGetNetworkDeviceGroupSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponseNetworkDeviceGroupGetNetworkDeviceGroupSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseNetworkDeviceGroupGetNetworkDeviceGroupSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseNetworkDeviceGroupGetNetworkDeviceGroupSearchResultResources struct {
	ID          string                                                                    `json:"id,omitempty"`          //
	Name        string                                                                    `json:"name,omitempty"`        //
	Description string                                                                    `json:"description,omitempty"` //
	Link        *ResponseNetworkDeviceGroupGetNetworkDeviceGroupSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseNetworkDeviceGroupGetNetworkDeviceGroupSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkDeviceGroupGetNetworkDeviceGroupSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkDeviceGroupGetNetworkDeviceGroupSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkDeviceGroupGetVersion struct {
	VersionInfo *ResponseNetworkDeviceGroupGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseNetworkDeviceGroupGetVersionVersionInfo struct {
	CurrentServerVersion string                                               `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                               `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseNetworkDeviceGroupGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseNetworkDeviceGroupGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkDeviceGroupUpdateNetworkDeviceGroupByID struct {
	NetworkDeviceGroup *RequestNetworkDeviceGroupUpdateNetworkDeviceGroupByIDNetworkDeviceGroup `json:"NetworkDeviceGroup,omitempty"` //
}

type RequestNetworkDeviceGroupUpdateNetworkDeviceGroupByIDNetworkDeviceGroup struct {
	ID          string `json:"id,omitempty"`          //
	Name        string `json:"name,omitempty"`        //
	Description string `json:"description,omitempty"` //
	Othername   string `json:"othername,omitempty"`   //
}

type RequestNetworkDeviceGroupCreateNetworkDeviceGroup struct {
	NetworkDeviceGroup *RequestNetworkDeviceGroupCreateNetworkDeviceGroupNetworkDeviceGroup `json:"NetworkDeviceGroup,omitempty"` //
}

type RequestNetworkDeviceGroupCreateNetworkDeviceGroupNetworkDeviceGroup struct {
	Name        string `json:"name,omitempty"`        //
	Description string `json:"description,omitempty"` //
	Othername   string `json:"othername,omitempty"`   //
}

//GetNetworkDeviceGroupByName Get network device group by name
/* This API allows the client to get a network device group by name.

@param name name path parameter. colon is used as a separator instead of (#) in the NDG name.
For example, if the name is a#b#c it should appear in the request URL as a:b:c
*/
func (s *NetworkDeviceGroupService) GetNetworkDeviceGroupByName(name string) (*ResponseNetworkDeviceGroupGetNetworkDeviceGroupByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevicegroup/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkDeviceGroupGetNetworkDeviceGroupByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkDeviceGroupByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkDeviceGroupGetNetworkDeviceGroupByName)
	return result, response, err

}

//GetNetworkDeviceGroupByID Get network device group by ID
/* This API allows the client to get a network device group by ID.

@param id id path parameter.
*/
func (s *NetworkDeviceGroupService) GetNetworkDeviceGroupByID(id string) (*ResponseNetworkDeviceGroupGetNetworkDeviceGroupByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevicegroup/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkDeviceGroupGetNetworkDeviceGroupByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkDeviceGroupById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkDeviceGroupGetNetworkDeviceGroupByID)
	return result, response, err

}

//GetNetworkDeviceGroup Get all network device groups
/* This API allows the client to get all the network device groups.

Filter:

[name, description, type]

To search guest users by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:

[name, description]

@param getNetworkDeviceGroupQueryParams Filtering parameter
*/
func (s *NetworkDeviceGroupService) GetNetworkDeviceGroup(getNetworkDeviceGroupQueryParams *GetNetworkDeviceGroupQueryParams) (*ResponseNetworkDeviceGroupGetNetworkDeviceGroup, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevicegroup"

	queryString, _ := query.Values(getNetworkDeviceGroupQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkDeviceGroupGetNetworkDeviceGroup{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkDeviceGroup")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkDeviceGroupGetNetworkDeviceGroup)
	return result, response, err

}

//GetVersion Get network device group version information
/* This API helps to retrieve the version information related to the network device group.

 */
func (s *NetworkDeviceGroupService) GetVersion() (*ResponseNetworkDeviceGroupGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevicegroup/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkDeviceGroupGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkDeviceGroupGetVersion)
	return result, response, err

}

//CreateNetworkDeviceGroup Create network device group
/* This API creates a network device group.

 */
func (s *NetworkDeviceGroupService) CreateNetworkDeviceGroup(requestNetworkDeviceGroupCreateNetworkDeviceGroup *RequestNetworkDeviceGroupCreateNetworkDeviceGroup) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevicegroup"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkDeviceGroupCreateNetworkDeviceGroup).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkDeviceGroup")
	}

	return response, err

}

//UpdateNetworkDeviceGroupByID Update network device group
/* This API allows the client to update a network device group.

@param id id path parameter.
*/
func (s *NetworkDeviceGroupService) UpdateNetworkDeviceGroupByID(id string, requestNetworkDeviceGroupUpdateNetworkDeviceGroupById *RequestNetworkDeviceGroupUpdateNetworkDeviceGroupByID) (*ResponseNetworkDeviceGroupUpdateNetworkDeviceGroupByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevicegroup/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkDeviceGroupUpdateNetworkDeviceGroupById).
		SetResult(&ResponseNetworkDeviceGroupUpdateNetworkDeviceGroupByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkDeviceGroupById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkDeviceGroupUpdateNetworkDeviceGroupByID)
	return result, response, err

}

//DeleteNetworkDeviceGroupByID Delete network device group
/* This API deletes a network device group.

@param id id path parameter.
*/
func (s *NetworkDeviceGroupService) DeleteNetworkDeviceGroupByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevicegroup/{id}"
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
		return response, fmt.Errorf("error with operation DeleteNetworkDeviceGroupById")
	}

	getCSFRToken(response.Header())

	return response, err

}
