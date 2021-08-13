package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SxpConnectionsService service

type GetSxpConnectionsQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseSxpConnectionsGetSxpConnectionsByID struct {
	ERSSxpConnection ResponseSxpConnectionsGetSxpConnectionsByIDERSSxpConnection `json:"ERSSxpConnection,omitempty"` //
}

type ResponseSxpConnectionsGetSxpConnectionsByIDERSSxpConnection struct {
	ID          string                                                          `json:"id,omitempty"`          //
	Description string                                                          `json:"description,omitempty"` //
	SxpPeer     string                                                          `json:"sxpPeer,omitempty"`     //
	SxpVpn      string                                                          `json:"sxpVpn,omitempty"`      //
	SxpNode     string                                                          `json:"sxpNode,omitempty"`     //
	IPAddress   string                                                          `json:"ipAddress,omitempty"`   //
	SxpMode     string                                                          `json:"sxpMode,omitempty"`     //
	SxpVersion  string                                                          `json:"sxpVersion,omitempty"`  //
	Enabled     bool                                                            `json:"enabled,omitempty"`     //
	Link        ResponseSxpConnectionsGetSxpConnectionsByIDERSSxpConnectionLink `json:"link,omitempty"`        //
}

type ResponseSxpConnectionsGetSxpConnectionsByIDERSSxpConnectionLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSxpConnectionsUpdateSxpConnectionsByID struct {
	UpdatedFieldsList ResponseSxpConnectionsUpdateSxpConnectionsByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseSxpConnectionsUpdateSxpConnectionsByIDUpdatedFieldsList struct {
	UpdatedField ResponseSxpConnectionsUpdateSxpConnectionsByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                      `json:"field,omitempty"`        //
	OldValue     string                                                                      `json:"oldValue,omitempty"`     //
	NewValue     string                                                                      `json:"newValue,omitempty"`     //
}

type ResponseSxpConnectionsUpdateSxpConnectionsByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseSxpConnectionsGetSxpConnections struct {
	SearchResult ResponseSxpConnectionsGetSxpConnectionsSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseSxpConnectionsGetSxpConnectionsSearchResult struct {
	Total        int                                                               `json:"total,omitempty"`        //
	Resources    []ResponseSxpConnectionsGetSxpConnectionsSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseSxpConnectionsGetSxpConnectionsSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseSxpConnectionsGetSxpConnectionsSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseSxpConnectionsGetSxpConnectionsSearchResultResources struct {
	ID   string                                                           `json:"id,omitempty"`   //
	Link ResponseSxpConnectionsGetSxpConnectionsSearchResultResourcesLink `json:"link,omitempty"` //
}

type ResponseSxpConnectionsGetSxpConnectionsSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSxpConnectionsGetSxpConnectionsSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSxpConnectionsGetSxpConnectionsSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSxpConnectionsGetVersion struct {
	VersionInfo ResponseSxpConnectionsGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseSxpConnectionsGetVersionVersionInfo struct {
	CurrentServerVersion string                                          `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                          `json:"supportedVersions,omitempty"`    //
	Link                 ResponseSxpConnectionsGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseSxpConnectionsGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSxpConnectionsMonitorBulkStatusSxpConnections struct {
	BulkStatus ResponseSxpConnectionsMonitorBulkStatusSxpConnectionsBulkStatus `json:"BulkStatus,omitempty"` //
}

type ResponseSxpConnectionsMonitorBulkStatusSxpConnectionsBulkStatus struct {
	BulkID          string                                                                           `json:"bulkId,omitempty"`          //
	MediaType       string                                                                           `json:"mediaType,omitempty"`       //
	ExecutionStatus string                                                                           `json:"executionStatus,omitempty"` //
	OperationType   string                                                                           `json:"operationType,omitempty"`   //
	StartTime       string                                                                           `json:"startTime,omitempty"`       //
	ResourcesCount  int                                                                              `json:"resourcesCount,omitempty"`  //
	SuccessCount    int                                                                              `json:"successCount,omitempty"`    //
	FailCount       int                                                                              `json:"failCount,omitempty"`       //
	ResourcesStatus []ResponseSxpConnectionsMonitorBulkStatusSxpConnectionsBulkStatusResourcesStatus `json:"resourcesStatus,omitempty"` //
}

type ResponseSxpConnectionsMonitorBulkStatusSxpConnectionsBulkStatusResourcesStatus struct {
	ID                      string `json:"id,omitempty"`                      //
	Name                    string `json:"name,omitempty"`                    //
	Description             string `json:"description,omitempty"`             //
	ResourceExecutionStatus string `json:"resourceExecutionStatus,omitempty"` //
	Status                  string `json:"status,omitempty"`                  //
}

type RequestSxpConnectionsUpdateSxpConnectionsByID struct {
	ERSSxpConnection *RequestSxpConnectionsUpdateSxpConnectionsByIDERSSxpConnection `json:"ERSSxpConnection,omitempty"` //
}

type RequestSxpConnectionsUpdateSxpConnectionsByIDERSSxpConnection struct {
	ID          string `json:"id,omitempty"`          //
	Description string `json:"description,omitempty"` //
	SxpPeer     string `json:"sxpPeer,omitempty"`     //
	SxpVpn      string `json:"sxpVpn,omitempty"`      //
	SxpNode     string `json:"sxpNode,omitempty"`     //
	IPAddress   string `json:"ipAddress,omitempty"`   //
	SxpMode     string `json:"sxpMode,omitempty"`     //
	SxpVersion  string `json:"sxpVersion,omitempty"`  //
	Enabled     *bool  `json:"enabled,omitempty"`     //
}

type RequestSxpConnectionsCreateSxpConnections struct {
	ERSSxpConnection *RequestSxpConnectionsCreateSxpConnectionsERSSxpConnection `json:"ERSSxpConnection,omitempty"` //
}

type RequestSxpConnectionsCreateSxpConnectionsERSSxpConnection struct {
	Description string `json:"description,omitempty"` //
	SxpPeer     string `json:"sxpPeer,omitempty"`     //
	SxpVpn      string `json:"sxpVpn,omitempty"`      //
	SxpNode     string `json:"sxpNode,omitempty"`     //
	IPAddress   string `json:"ipAddress,omitempty"`   //
	SxpMode     string `json:"sxpMode,omitempty"`     //
	SxpVersion  string `json:"sxpVersion,omitempty"`  //
	Enabled     *bool  `json:"enabled,omitempty"`     //
}

type RequestSxpConnectionsBulkRequestForSxpConnections struct {
	ConnectionBulkRequest *RequestSxpConnectionsBulkRequestForSxpConnectionsConnectionBulkRequest `json:"ConnectionBulkRequest,omitempty"` //
}

type RequestSxpConnectionsBulkRequestForSxpConnectionsConnectionBulkRequest struct {
	OperationType     string `json:"operationType,omitempty"`     //
	ResourceMediaType string `json:"resourceMediaType,omitempty"` //
}

//GetSxpConnectionsByID Get SXP connection by ID
/* This API allows the client to get a SXP connection by ID.

@param id id path parameter.
*/
func (s *SxpConnectionsService) GetSxpConnectionsByID(id string) (*ResponseSxpConnectionsGetSxpConnectionsByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxpconnections/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSxpConnectionsGetSxpConnectionsByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSxpConnectionsById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSxpConnectionsGetSxpConnectionsByID)
	return result, response, err

}

//GetSxpConnections Get all SXP connections
/* This API allows the client to get all the SXP connections.

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

@param getSXPConnectionsQueryParams Filtering parameter
*/
func (s *SxpConnectionsService) GetSxpConnections(getSXPConnectionsQueryParams *GetSxpConnectionsQueryParams) (*ResponseSxpConnectionsGetSxpConnections, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxpconnections"

	queryString, _ := query.Values(getSXPConnectionsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSxpConnectionsGetSxpConnections{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSxpConnections")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSxpConnectionsGetSxpConnections)
	return result, response, err

}

//GetVersion Get SXP connections version information
/* This API helps to retrieve the version information related to the SXP connections.

 */
func (s *SxpConnectionsService) GetVersion() (*ResponseSxpConnectionsGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxpconnections/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSxpConnectionsGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSxpConnectionsGetVersion)
	return result, response, err

}

//MonitorBulkStatusSxpConnections Monitor bulk request
/* This API allows the client to monitor the bulk request.

@param bulkid bulkid path parameter.
*/
func (s *SxpConnectionsService) MonitorBulkStatusSxpConnections(bulkid string) (*ResponseSxpConnectionsMonitorBulkStatusSxpConnections, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxpconnections/bulk/{bulkid}"
	path = strings.Replace(path, "{bulkid}", fmt.Sprintf("%v", bulkid), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSxpConnectionsMonitorBulkStatusSxpConnections{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MonitorBulkStatusSxpConnections")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSxpConnectionsMonitorBulkStatusSxpConnections)
	return result, response, err

}

//CreateSxpConnections Create SXP connection
/* This API creates a SXP connection.

 */
func (s *SxpConnectionsService) CreateSxpConnections(requestSxpConnectionsCreateSXPConnections *RequestSxpConnectionsCreateSxpConnections) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxpconnections"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSxpConnectionsCreateSXPConnections).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateSxpConnections")
	}

	return response, err

}

//UpdateSxpConnectionsByID Update SXP connection
/* This API allows the client to update a SXP connection.

@param id id path parameter.
*/
func (s *SxpConnectionsService) UpdateSxpConnectionsByID(id string, requestSxpConnectionsUpdateSXPConnectionsById *RequestSxpConnectionsUpdateSxpConnectionsByID) (*ResponseSxpConnectionsUpdateSxpConnectionsByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxpconnections/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSxpConnectionsUpdateSXPConnectionsById).
		SetResult(&ResponseSxpConnectionsUpdateSxpConnectionsByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSxpConnectionsById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSxpConnectionsUpdateSxpConnectionsByID)
	return result, response, err

}

//BulkRequestForSxpConnections Submit bulk request
/* This API allows the client to submit the bulk request.

 */
func (s *SxpConnectionsService) BulkRequestForSxpConnections(requestSxpConnectionsBulkRequestForSXPConnections *RequestSxpConnectionsBulkRequestForSxpConnections) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxpconnections/bulk/submit"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSxpConnectionsBulkRequestForSXPConnections).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation BulkRequestForSxpConnections")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteSxpConnectionsByID Delete SXP connection
/* This API deletes a SXP connection.

@param id id path parameter.
*/
func (s *SxpConnectionsService) DeleteSxpConnectionsByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxpconnections/{id}"
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
		return response, fmt.Errorf("error with operation DeleteSxpConnectionsById")
	}

	getCSFRToken(response.Header())

	return response, err

}
