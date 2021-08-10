package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type PxGridNodeService service

type GetPxGridNodeQueryParams struct {
	Page int `url: page,omitempty` //Page number
	Size int `url: size,omitempty` //Number of objects returned per page
}

type ResponsePxGridNodeGetPxGridNodeByName struct {
	PxgridNode ResponsePxGridNodeGetPxGridNodeByNamePxgridNode `json:"PxgridNode,omitempty"` //
}

type ResponsePxGridNodeGetPxGridNodeByNamePxgridNode struct {
	ID          string                                              `json:"id,omitempty"`          //
	Name        string                                              `json:"name,omitempty"`        //
	Description string                                              `json:"description,omitempty"` //
	Status      string                                              `json:"status,omitempty"`      //
	AuthMethod  string                                              `json:"authMethod,omitempty"`  //
	Groups      string                                              `json:"groups,omitempty"`      //
	Link        ResponsePxGridNodeGetPxGridNodeByNamePxgridNodeLink `json:"link,omitempty"`        //
}

type ResponsePxGridNodeGetPxGridNodeByNamePxgridNodeLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePxGridNodeGetPxGridNodeByID struct {
	PxgridNode ResponsePxGridNodeGetPxGridNodeByIDPxgridNode `json:"PxgridNode,omitempty"` //
}

type ResponsePxGridNodeGetPxGridNodeByIDPxgridNode struct {
	ID          string                                            `json:"id,omitempty"`          //
	Name        string                                            `json:"name,omitempty"`        //
	Description string                                            `json:"description,omitempty"` //
	Status      string                                            `json:"status,omitempty"`      //
	AuthMethod  string                                            `json:"authMethod,omitempty"`  //
	Groups      string                                            `json:"groups,omitempty"`      //
	Link        ResponsePxGridNodeGetPxGridNodeByIDPxgridNodeLink `json:"link,omitempty"`        //
}

type ResponsePxGridNodeGetPxGridNodeByIDPxgridNodeLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePxGridNodeGetPxGridNode struct {
	SearchResult ResponsePxGridNodeGetPxGridNodeSearchResult `json:"SearchResult,omitempty"` //
}

type ResponsePxGridNodeGetPxGridNodeSearchResult struct {
	Total        int                                                       `json:"total,omitempty"`        //
	Resources    []ResponsePxGridNodeGetPxGridNodeSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponsePxGridNodeGetPxGridNodeSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponsePxGridNodeGetPxGridNodeSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponsePxGridNodeGetPxGridNodeSearchResultResources struct {
	ID          string                                                   `json:"id,omitempty"`          //
	Name        string                                                   `json:"name,omitempty"`        //
	Description string                                                   `json:"description,omitempty"` //
	Link        ResponsePxGridNodeGetPxGridNodeSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponsePxGridNodeGetPxGridNodeSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePxGridNodeGetPxGridNodeSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePxGridNodeGetPxGridNodeSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePxGridNodeGetVersion struct {
	VersionInfo ResponsePxGridNodeGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponsePxGridNodeGetVersionVersionInfo struct {
	CurrentServerVersion string                                      `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                      `json:"supportedVersions,omitempty"`    //
	Link                 ResponsePxGridNodeGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponsePxGridNodeGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

//GetPxGridNodeByName Get pxGrid node by name
/* This API allows the client to get a pxGrid node by name.

@param name name path parameter.
*/
func (s *PxGridNodeService) GetPxGridNodeByName(name string) (*ResponsePxGridNodeGetPxGridNodeByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/pxgridnode/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePxGridNodeGetPxGridNodeByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPxGridNodeByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePxGridNodeGetPxGridNodeByName)
	return result, response, err

}

//GetPxGridNodeByID Get pxGrid node by ID
/* This API allows the client to get a pxGrid node by ID.

@param id id path parameter.
*/
func (s *PxGridNodeService) GetPxGridNodeByID(id string) (*ResponsePxGridNodeGetPxGridNodeByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/pxgridnode/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePxGridNodeGetPxGridNodeByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPxGridNodeById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePxGridNodeGetPxGridNodeByID)
	return result, response, err

}

//GetPxGridNode Get all pxGrid nodes
/* This API allows the client to get all the npxGrid nodes.

@param getPxGridNodeQueryParams Filtering parameter
*/
func (s *PxGridNodeService) GetPxGridNode(getPxGridNodeQueryParams *GetPxGridNodeQueryParams) (*ResponsePxGridNodeGetPxGridNode, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/pxgridnode"

	queryString, _ := query.Values(getPxGridNodeQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponsePxGridNodeGetPxGridNode{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPxGridNode")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePxGridNodeGetPxGridNode)
	return result, response, err

}

//GetVersion Get pxGrid node version information
/* This API helps to retrieve the version information related to the pxGrid node.

 */
func (s *PxGridNodeService) GetVersion() (*ResponsePxGridNodeGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/pxgridnode/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePxGridNodeGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePxGridNodeGetVersion)
	return result, response, err

}

//ApprovePxGridNode Approve pxGrid node
/* This API allows the client to approve a pxGrid node.
Only pending pxGrid nodes can be approved

@param name name path parameter.
*/
func (s *PxGridNodeService) ApprovePxGridNode(name string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/pxgridnode/name/{name}/approve"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation ApprovePxGridNode")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeletePxGridNodeByName Delete pxGrid node by name
/* This API deletes a pxGrid node by name.

@param name name path parameter.
*/
func (s *PxGridNodeService) DeletePxGridNodeByName(name string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/pxgridnode/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

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
		return response, fmt.Errorf("error with operation DeletePxGridNodeByName")
	}

	getCSFRToken(response.Header())

	return response, err

}
