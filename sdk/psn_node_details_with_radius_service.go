package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type PsnNodeDetailsWithRadiusServiceService service

type GetSessionServiceNodeQueryParams struct {
	Page int `url:"page,omitempty"` //Page number
	Size int `url:"size,omitempty"` //Number of objects returned per page
}

type ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeByName struct {
	SessionServiceNode *ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeByNameSessionServiceNode `json:"SessionServiceNode,omitempty"` //
}

type ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeByNameSessionServiceNode struct {
	ID          string                                                                                    `json:"id,omitempty"`          //
	Name        string                                                                                    `json:"name,omitempty"`        //
	Description string                                                                                    `json:"description,omitempty"` //
	IPAddress   string                                                                                    `json:"ipAddress,omitempty"`   //
	GateWay     string                                                                                    `json:"gateWay,omitempty"`     //
	Link        *ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeByNameSessionServiceNodeLink `json:"link,omitempty"`        //
}

type ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeByNameSessionServiceNodeLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeByID struct {
	SessionServiceNode *ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeByIDSessionServiceNode `json:"SessionServiceNode,omitempty"` //
}

type ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeByIDSessionServiceNode struct {
	ID          string                                                                                  `json:"id,omitempty"`          //
	Name        string                                                                                  `json:"name,omitempty"`        //
	Description string                                                                                  `json:"description,omitempty"` //
	IPAddress   string                                                                                  `json:"ipAddress,omitempty"`   //
	GateWay     string                                                                                  `json:"gateWay,omitempty"`     //
	Link        *ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeByIDSessionServiceNodeLink `json:"link,omitempty"`        //
}

type ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeByIDSessionServiceNodeLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNode struct {
	SearchResult *ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeSearchResult `json:"SearchResult,omitempty"` //
}

type ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeSearchResult struct {
	Total        *int                                                                                  `json:"total,omitempty"`        //
	Resources    *[]ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeSearchResultResources struct {
	ID          string                                                                                 `json:"id,omitempty"`          //
	Name        string                                                                                 `json:"name,omitempty"`        //
	Description string                                                                                 `json:"description,omitempty"` //
	Link        *ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePsnNodeDetailsWithRadiusServiceGetVersion struct {
	VersionInfo *ResponsePsnNodeDetailsWithRadiusServiceGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponsePsnNodeDetailsWithRadiusServiceGetVersionVersionInfo struct {
	CurrentServerVersion string                                                            `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                                            `json:"supportedVersions,omitempty"`    //
	Link                 *ResponsePsnNodeDetailsWithRadiusServiceGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponsePsnNodeDetailsWithRadiusServiceGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

//GetSessionServiceNodeByName Get PSN node details by name
/* This API allows the client to get a PSN node details by name.

@param name name path parameter.
*/
func (s *PsnNodeDetailsWithRadiusServiceService) GetSessionServiceNodeByName(name string) (*ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sessionservicenode/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSessionServiceNodeByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeByName)
	return result, response, err

}

//GetSessionServiceNodeByID Get PSN node details by ID
/* This API allows the client to get a PSN node details by ID.

@param id id path parameter.
*/
func (s *PsnNodeDetailsWithRadiusServiceService) GetSessionServiceNodeByID(id string) (*ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sessionservicenode/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSessionServiceNodeById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNodeByID)
	return result, response, err

}

//GetSessionServiceNode Get all PSN node details
/* This API allows the client to get all the PSN node details.

@param getSessionServiceNodeQueryParams Filtering parameter
*/
func (s *PsnNodeDetailsWithRadiusServiceService) GetSessionServiceNode(getSessionServiceNodeQueryParams *GetSessionServiceNodeQueryParams) (*ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNode, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sessionservicenode"

	queryString, _ := query.Values(getSessionServiceNodeQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNode{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSessionServiceNode")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePsnNodeDetailsWithRadiusServiceGetSessionServiceNode)
	return result, response, err

}

//GetVersion Get PSN node version information
/* This API helps to retrieve the version information related to the PSN node.

 */
func (s *PsnNodeDetailsWithRadiusServiceService) GetVersion() (*ResponsePsnNodeDetailsWithRadiusServiceGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sessionservicenode/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePsnNodeDetailsWithRadiusServiceGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePsnNodeDetailsWithRadiusServiceGetVersion)
	return result, response, err

}
