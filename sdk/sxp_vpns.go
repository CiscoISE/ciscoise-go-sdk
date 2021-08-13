package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SxpVpnsService service

type GetSxpVpnsQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
}

type ResponseSxpVpnsGetSxpVpnByID struct {
	ERSSxpVpn ResponseSxpVpnsGetSxpVpnByIDERSSxpVpn `json:"ERSSxpVpn,omitempty"` //
}

type ResponseSxpVpnsGetSxpVpnByIDERSSxpVpn struct {
	ID         string                                    `json:"id,omitempty"`         //
	SxpVpnName string                                    `json:"sxpVpnName,omitempty"` //
	Link       ResponseSxpVpnsGetSxpVpnByIDERSSxpVpnLink `json:"link,omitempty"`       //
}

type ResponseSxpVpnsGetSxpVpnByIDERSSxpVpnLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSxpVpnsGetSxpVpns struct {
	SearchResult ResponseSxpVpnsGetSxpVpnsSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseSxpVpnsGetSxpVpnsSearchResult struct {
	Total        int                                               `json:"total,omitempty"`        //
	Resources    []ResponseSxpVpnsGetSxpVpnsSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponseSxpVpnsGetSxpVpnsSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponseSxpVpnsGetSxpVpnsSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseSxpVpnsGetSxpVpnsSearchResultResources struct {
	ID   string                                             `json:"id,omitempty"`   //
	Link ResponseSxpVpnsGetSxpVpnsSearchResultResourcesLink `json:"link,omitempty"` //
}

type ResponseSxpVpnsGetSxpVpnsSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSxpVpnsGetSxpVpnsSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSxpVpnsGetSxpVpnsSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSxpVpnsGetVersion struct {
	VersionInfo ResponseSxpVpnsGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseSxpVpnsGetVersionVersionInfo struct {
	CurrentServerVersion string                                   `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                   `json:"supportedVersions,omitempty"`    //
	Link                 ResponseSxpVpnsGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseSxpVpnsGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSxpVpnsMonitorBulkStatusSxpVpns struct {
	BulkStatus ResponseSxpVpnsMonitorBulkStatusSxpVpnsBulkStatus `json:"BulkStatus,omitempty"` //
}

type ResponseSxpVpnsMonitorBulkStatusSxpVpnsBulkStatus struct {
	BulkID          string                                                             `json:"bulkId,omitempty"`          //
	MediaType       string                                                             `json:"mediaType,omitempty"`       //
	ExecutionStatus string                                                             `json:"executionStatus,omitempty"` //
	OperationType   string                                                             `json:"operationType,omitempty"`   //
	StartTime       string                                                             `json:"startTime,omitempty"`       //
	ResourcesCount  int                                                                `json:"resourcesCount,omitempty"`  //
	SuccessCount    int                                                                `json:"successCount,omitempty"`    //
	FailCount       int                                                                `json:"failCount,omitempty"`       //
	ResourcesStatus []ResponseSxpVpnsMonitorBulkStatusSxpVpnsBulkStatusResourcesStatus `json:"resourcesStatus,omitempty"` //
}

type ResponseSxpVpnsMonitorBulkStatusSxpVpnsBulkStatusResourcesStatus struct {
	ID                      string `json:"id,omitempty"`                      //
	Name                    string `json:"name,omitempty"`                    //
	Description             string `json:"description,omitempty"`             //
	ResourceExecutionStatus string `json:"resourceExecutionStatus,omitempty"` //
	Status                  string `json:"status,omitempty"`                  //
}

type RequestSxpVpnsCreateSxpVpn struct {
	ERSSxpVpn *RequestSxpVpnsCreateSxpVpnERSSxpVpn `json:"ERSSxpVpn,omitempty"` //
}

type RequestSxpVpnsCreateSxpVpnERSSxpVpn struct {
	SxpVpnName string `json:"sxpVpnName,omitempty"` //
}

type RequestSxpVpnsBulkRequestForSxpVpns struct {
	VpnBulkRequest *RequestSxpVpnsBulkRequestForSxpVpnsVpnBulkRequest `json:"VpnBulkRequest,omitempty"` //
}

type RequestSxpVpnsBulkRequestForSxpVpnsVpnBulkRequest struct {
	OperationType     string `json:"operationType,omitempty"`     //
	ResourceMediaType string `json:"resourceMediaType,omitempty"` //
}

//GetSxpVpnByID Get SXP VPN by ID
/* This API allows the client to get a SXP VPN by ID.

@param id id path parameter.
*/
func (s *SxpVpnsService) GetSxpVpnByID(id string) (*ResponseSxpVpnsGetSxpVpnByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxpvpns/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSxpVpnsGetSxpVpnByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSxpVpnById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSxpVpnsGetSxpVpnByID)
	return result, response, err

}

//GetSxpVpns Get all SXP VPNs
/* This API allows the client to get all the SXP VPNs.

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

@param getSXPVPNsQueryParams Filtering parameter
*/
func (s *SxpVpnsService) GetSxpVpns(getSXPVPNsQueryParams *GetSxpVpnsQueryParams) (*ResponseSxpVpnsGetSxpVpns, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxpvpns"

	queryString, _ := query.Values(getSXPVPNsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSxpVpnsGetSxpVpns{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSxpVpns")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSxpVpnsGetSxpVpns)
	return result, response, err

}

//GetVersion Get SXP VPNs version information
/* This API helps to retrieve the version information related to the SXP VPNs.

 */
func (s *SxpVpnsService) GetVersion() (*ResponseSxpVpnsGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxpvpns/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSxpVpnsGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSxpVpnsGetVersion)
	return result, response, err

}

//MonitorBulkStatusSxpVpns Monitor bulk request
/* This API allows the client to monitor the bulk request.

@param bulkid bulkid path parameter.
*/
func (s *SxpVpnsService) MonitorBulkStatusSxpVpns(bulkid string) (*ResponseSxpVpnsMonitorBulkStatusSxpVpns, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxpvpns/bulk/{bulkid}"
	path = strings.Replace(path, "{bulkid}", fmt.Sprintf("%v", bulkid), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSxpVpnsMonitorBulkStatusSxpVpns{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MonitorBulkStatusSxpVpns")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSxpVpnsMonitorBulkStatusSxpVpns)
	return result, response, err

}

//CreateSxpVpn Create SXP VPN
/* This API creates a SXP VPN.

 */
func (s *SxpVpnsService) CreateSxpVpn(requestSxpVpnsCreateSXPVPN *RequestSxpVpnsCreateSxpVpn) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxpvpns"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSxpVpnsCreateSXPVPN).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateSxpVpn")
	}

	return response, err

}

//BulkRequestForSxpVpns Submit bulk request
/* This API allows the client to submit the bulk request.

 */
func (s *SxpVpnsService) BulkRequestForSxpVpns(requestSxpVpnsBulkRequestForSXPVPNs *RequestSxpVpnsBulkRequestForSxpVpns) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxpvpns/bulk/submit"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSxpVpnsBulkRequestForSXPVPNs).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation BulkRequestForSxpVpns")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteSxpVpnByID Delete SXP VPN
/* This API deletes a SXP VPN.

@param id id path parameter.
*/
func (s *SxpVpnsService) DeleteSxpVpnByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxpvpns/{id}"
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
		return response, fmt.Errorf("error with operation DeleteSxpVpnById")
	}

	getCSFRToken(response.Header())

	return response, err

}
