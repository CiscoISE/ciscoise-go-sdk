package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type NodeDetailsService service

type GetNodeDetailsQueryParams struct {
	Page       int      `url: page,omitempty`       //Page number
	Size       int      `url: size,omitempty`       //Number of objects returned per page
	Filter     []string `url: filter,omitempty`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url: filterType,omitempty` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseNodeDetailsGetNodeDetailByName struct {
	Node ResponseNodeDetailsGetNodeDetailByNameNode `json:"Node,omitempty"` //
}

type ResponseNodeDetailsGetNodeDetailByNameNode struct {
	ID               string                                         `json:"id,omitempty"`               //
	Name             string                                         `json:"name,omitempty"`             //
	GateWay          string                                         `json:"gateWay,omitempty"`          //
	UserName         string                                         `json:"userName,omitempty"`         //
	PassWord         string                                         `json:"passWord,omitempty"`         //
	DisplayName      string                                         `json:"displayName,omitempty"`      //
	InDeployment     bool                                           `json:"inDeployment,omitempty"`     //
	OtherPapFqdn     string                                         `json:"otherPapFqdn,omitempty"`     //
	IPAddresses      []string                                       `json:"ipAddresses,omitempty"`      //
	IPAddress        string                                         `json:"ipAddress,omitempty"`        //
	SxpIPAddress     string                                         `json:"sxpIpAddress,omitempty"`     //
	NodeServiceTypes string                                         `json:"nodeServiceTypes,omitempty"` //
	Fqdn             string                                         `json:"fqdn,omitempty"`             //
	PapNode          bool                                           `json:"papNode,omitempty"`          //
	PrimaryPapNode   bool                                           `json:"primaryPapNode,omitempty"`   //
	PxGridNode       bool                                           `json:"pxGridNode,omitempty"`       //
	Link             ResponseNodeDetailsGetNodeDetailByNameNodeLink `json:"link,omitempty"`             //
}

type ResponseNodeDetailsGetNodeDetailByNameNodeLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNodeDetailsGetNodeDetailByID struct {
	Node ResponseNodeDetailsGetNodeDetailByIDNode `json:"Node,omitempty"` //
}

type ResponseNodeDetailsGetNodeDetailByIDNode struct {
	ID               string                                       `json:"id,omitempty"`               //
	Name             string                                       `json:"name,omitempty"`             //
	GateWay          string                                       `json:"gateWay,omitempty"`          //
	UserName         string                                       `json:"userName,omitempty"`         //
	PassWord         string                                       `json:"passWord,omitempty"`         //
	DisplayName      string                                       `json:"displayName,omitempty"`      //
	InDeployment     bool                                         `json:"inDeployment,omitempty"`     //
	OtherPapFqdn     string                                       `json:"otherPapFqdn,omitempty"`     //
	IPAddresses      []string                                     `json:"ipAddresses,omitempty"`      //
	IPAddress        string                                       `json:"ipAddress,omitempty"`        //
	SxpIPAddress     string                                       `json:"sxpIpAddress,omitempty"`     //
	NodeServiceTypes string                                       `json:"nodeServiceTypes,omitempty"` //
	Fqdn             string                                       `json:"fqdn,omitempty"`             //
	PapNode          bool                                         `json:"papNode,omitempty"`          //
	PrimaryPapNode   bool                                         `json:"primaryPapNode,omitempty"`   //
	PxGridNode       bool                                         `json:"pxGridNode,omitempty"`       //
	Link             ResponseNodeDetailsGetNodeDetailByIDNodeLink `json:"link,omitempty"`             //
}

type ResponseNodeDetailsGetNodeDetailByIDNodeLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNodeDetailsGetNodeDetails struct {
	SearchResult ResponseNodeDetailsGetNodeDetailsSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseNodeDetailsGetNodeDetailsSearchResult struct {
	Total        int                                                         `json:"total,omitempty"`        //
	Resources    []ResponseNodeDetailsGetNodeDetailsSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseNodeDetailsGetNodeDetailsSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseNodeDetailsGetNodeDetailsSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseNodeDetailsGetNodeDetailsSearchResultResources struct {
	ID          string                                                     `json:"id,omitempty"`          //
	Name        string                                                     `json:"name,omitempty"`        //
	Description string                                                     `json:"description,omitempty"` //
	Link        ResponseNodeDetailsGetNodeDetailsSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseNodeDetailsGetNodeDetailsSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNodeDetailsGetNodeDetailsSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNodeDetailsGetNodeDetailsSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNodeDetailsGetVersion struct {
	VersionInfo ResponseNodeDetailsGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseNodeDetailsGetVersionVersionInfo struct {
	CurrentServerVersion string                                       `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                       `json:"supportedVersions,omitempty"`    //
	Link                 ResponseNodeDetailsGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseNodeDetailsGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

//GetNodeDetailByName Get node details by name
/* This API allows the client to get node details by name.

@param name name path parameter.
*/
func (s *NodeDetailsService) GetNodeDetailByName(name string) (*ResponseNodeDetailsGetNodeDetailByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/node/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeDetailsGetNodeDetailByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNodeDetailByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeDetailsGetNodeDetailByName)
	return result, response, err

}

//GetNodeDetailByID Get node details by ID
/* This API allows the client to get node details by ID.

@param id id path parameter.
*/
func (s *NodeDetailsService) GetNodeDetailByID(id string) (*ResponseNodeDetailsGetNodeDetailByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/node/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeDetailsGetNodeDetailByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNodeDetailById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeDetailsGetNodeDetailByID)
	return result, response, err

}

//GetNodeDetails Get all node details
/* This API allows the client to get all the node details.

Filter:

[nodeservicetypes]

To search resources by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


@param getNodeDetailsQueryParams Filtering parameter
*/
func (s *NodeDetailsService) GetNodeDetails(getNodeDetailsQueryParams *GetNodeDetailsQueryParams) (*ResponseNodeDetailsGetNodeDetails, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/node"

	queryString, _ := query.Values(getNodeDetailsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNodeDetailsGetNodeDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNodeDetails")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeDetailsGetNodeDetails)
	return result, response, err

}

//GetVersion Get node details version information
/* This API helps to retrieve the version information realted to the node details.

 */
func (s *NodeDetailsService) GetVersion() (*ResponseNodeDetailsGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/node/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeDetailsGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeDetailsGetVersion)
	return result, response, err

}
