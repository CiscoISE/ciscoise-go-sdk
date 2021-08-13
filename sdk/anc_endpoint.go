package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type AncEndpointService service

type GetAncEndpointQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseAncEndpointGetAncEndpointByID struct {
	ErsAncEndpoint ResponseAncEndpointGetAncEndpointByIDErsAncEndpoint `json:"ErsAncEndpoint,omitempty"` //
}

type ResponseAncEndpointGetAncEndpointByIDErsAncEndpoint struct {
	ID         string                                                  `json:"id,omitempty"`         //
	MacAddress string                                                  `json:"macAddress,omitempty"` //
	PolicyName string                                                  `json:"policyName,omitempty"` //
	Link       ResponseAncEndpointGetAncEndpointByIDErsAncEndpointLink `json:"link,omitempty"`       //
}

type ResponseAncEndpointGetAncEndpointByIDErsAncEndpointLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAncEndpointGetAncEndpoint struct {
	SearchResult ResponseAncEndpointGetAncEndpointSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseAncEndpointGetAncEndpointSearchResult struct {
	Total        int                                                       `json:"total,omitempty"`        //
	Resources    []ResponseAncEndpointGetAncEndpointSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponseAncEndpointGetAncEndpointSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponseAncEndpointGetAncEndpointSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseAncEndpointGetAncEndpointSearchResultResources struct {
	ID   string                                                     `json:"id,omitempty"`   //
	Link ResponseAncEndpointGetAncEndpointSearchResultResourcesLink `json:"link,omitempty"` //
}

type ResponseAncEndpointGetAncEndpointSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAncEndpointGetAncEndpointSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAncEndpointGetAncEndpointSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAncEndpointGetVersion struct {
	VersionInfo ResponseAncEndpointGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseAncEndpointGetVersionVersionInfo struct {
	CurrentServerVersion string                                       `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                       `json:"supportedVersions,omitempty"`    //
	Link                 ResponseAncEndpointGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseAncEndpointGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAncEndpointMonitorBulkStatusAncEndpoint struct {
	BulkStatus ResponseAncEndpointMonitorBulkStatusAncEndpointBulkStatus `json:"BulkStatus,omitempty"` //
}

type ResponseAncEndpointMonitorBulkStatusAncEndpointBulkStatus struct {
	BulkID          string                                                                     `json:"bulkID,omitempty"`          //
	ExecutionStatus string                                                                     `json:"executionStatus,omitempty"` //
	OperationType   string                                                                     `json:"operationType,omitempty"`   //
	StartTime       string                                                                     `json:"startTime,omitempty"`       //
	ResourcesCount  int                                                                        `json:"resourcesCount,omitempty"`  //
	SuccessCount    int                                                                        `json:"successCount,omitempty"`    //
	FailCount       int                                                                        `json:"failCount,omitempty"`       //
	ResourcesStatus []ResponseAncEndpointMonitorBulkStatusAncEndpointBulkStatusResourcesStatus `json:"resourcesStatus,omitempty"` //
}

type ResponseAncEndpointMonitorBulkStatusAncEndpointBulkStatusResourcesStatus struct {
	ID                      string `json:"id,omitempty"`                      //
	Name                    string `json:"name,omitempty"`                    //
	Description             string `json:"description,omitempty"`             //
	ResourceExecutionStatus string `json:"resourceExecutionStatus,omitempty"` //
	Status                  string `json:"status,omitempty"`                  //
}

type RequestAncEndpointClearAncEndpoint struct {
	OperationAdditionalData *RequestAncEndpointClearAncEndpointOperationAdditionalData `json:"OperationAdditionalData,omitempty"` //
}

type RequestAncEndpointClearAncEndpointOperationAdditionalData struct {
	AdditionalData *[]RequestAncEndpointClearAncEndpointOperationAdditionalDataAdditionalData `json:"additionalData,omitempty"` //
}

type RequestAncEndpointClearAncEndpointOperationAdditionalDataAdditionalData struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type RequestAncEndpointApplyAncEndpoint struct {
	OperationAdditionalData *RequestAncEndpointApplyAncEndpointOperationAdditionalData `json:"OperationAdditionalData,omitempty"` //
}

type RequestAncEndpointApplyAncEndpointOperationAdditionalData struct {
	AdditionalData *[]RequestAncEndpointApplyAncEndpointOperationAdditionalDataAdditionalData `json:"additionalData,omitempty"` //
}

type RequestAncEndpointApplyAncEndpointOperationAdditionalDataAdditionalData struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type RequestAncEndpointBulkRequestForAncEndpoint struct {
	ErsAncEndpointBulkRequest *RequestAncEndpointBulkRequestForAncEndpointErsAncEndpointBulkRequest `json:"ErsAncEndpointBulkRequest,omitempty"` //
}

type RequestAncEndpointBulkRequestForAncEndpointErsAncEndpointBulkRequest struct {
	OperationType     string `json:"operationType,omitempty"`     //
	ResourceMediaType string `json:"resourceMediaType,omitempty"` //
}

//GetAncEndpointByID Get ANC endpoint by ID
/* This API allows the client to get an ANC endpoint by ID.

@param id id path parameter.
*/
func (s *AncEndpointService) GetAncEndpointByID(id string) (*ResponseAncEndpointGetAncEndpointByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ancendpoint/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAncEndpointGetAncEndpointByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAncEndpointById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAncEndpointGetAncEndpointByID)
	return result, response, err

}

//GetAncEndpoint Get all ANC endpoints
/* This API allows the client to get all the ANC endpoints.

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
[name]

@param getAncEndpointQueryParams Filtering parameter
*/
func (s *AncEndpointService) GetAncEndpoint(getAncEndpointQueryParams *GetAncEndpointQueryParams) (*ResponseAncEndpointGetAncEndpoint, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ancendpoint"

	queryString, _ := query.Values(getAncEndpointQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAncEndpointGetAncEndpoint{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAncEndpoint")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAncEndpointGetAncEndpoint)
	return result, response, err

}

//GetVersion Get ANC endpoint version information
/* This API helps to retrieve the version information related to the ANC Endpoint.

 */
func (s *AncEndpointService) GetVersion() (*ResponseAncEndpointGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ancendpoint/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAncEndpointGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAncEndpointGetVersion)
	return result, response, err

}

//MonitorBulkStatusAncEndpoint Monitor bulk request
/* This API allows the client to monitor the bulk request.

@param bulkid bulkid path parameter.
*/
func (s *AncEndpointService) MonitorBulkStatusAncEndpoint(bulkid string) (*ResponseAncEndpointMonitorBulkStatusAncEndpoint, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ancendpoint/bulk/{bulkid}"
	path = strings.Replace(path, "{bulkid}", fmt.Sprintf("%v", bulkid), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAncEndpointMonitorBulkStatusAncEndpoint{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MonitorBulkStatusAncEndpoint")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAncEndpointMonitorBulkStatusAncEndpoint)
	return result, response, err

}

//ClearAncEndpoint Clear required configuration
/* This API allows the client to clear the required configuration.

 */
func (s *AncEndpointService) ClearAncEndpoint(requestAncEndpointClearAncEndpoint *RequestAncEndpointClearAncEndpoint) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ancendpoint/clear"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAncEndpointClearAncEndpoint).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation ClearAncEndpoint")
	}

	getCSFRToken(response.Header())

	return response, err

}

//ApplyAncEndpoint Apply required configuration
/* This API allows the client to apply the required configuration.

 */
func (s *AncEndpointService) ApplyAncEndpoint(requestAncEndpointApplyAncEndpoint *RequestAncEndpointApplyAncEndpoint) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ancendpoint/apply"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAncEndpointApplyAncEndpoint).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation ApplyAncEndpoint")
	}

	getCSFRToken(response.Header())

	return response, err

}

//BulkRequestForAncEndpoint Submit bulk request
/* This API allows the client to submit the bulk request.

 */
func (s *AncEndpointService) BulkRequestForAncEndpoint(requestAncEndpointBulkRequestForAncEndpoint *RequestAncEndpointBulkRequestForAncEndpoint) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ancendpoint/bulk/submit"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAncEndpointBulkRequestForAncEndpoint).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation BulkRequestForAncEndpoint")
	}

	getCSFRToken(response.Header())

	return response, err

}
