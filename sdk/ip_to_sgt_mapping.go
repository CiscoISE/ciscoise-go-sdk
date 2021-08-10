package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type IPToSgtMappingService service

type GetIPToSgtMappingQueryParams struct {
	Page       int      `url: page,omitempty`       //Page number
	Size       int      `url: size,omitempty`       //Number of objects returned per page
	Sortasc    string   `url: sortasc,omitempty`    //sort asc
	Sortdsc    string   `url: sortdsc,omitempty`    //sort desc
	Filter     []string `url: filter,omitempty`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url: filterType,omitempty` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseIPToSgtMappingGetDeployStatusIPToSgtMapping struct {
	OperationResult ResponseIPToSgtMappingGetDeployStatusIPToSgtMappingOperationResult `json:"OperationResult,omitempty"` //
}

type ResponseIPToSgtMappingGetDeployStatusIPToSgtMappingOperationResult struct {
	ResultValue []ResponseIPToSgtMappingGetDeployStatusIPToSgtMappingOperationResultResultValue `json:"resultValue,omitempty"` //
}

type ResponseIPToSgtMappingGetDeployStatusIPToSgtMappingOperationResultResultValue struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type ResponseIPToSgtMappingGetIPToSgtMappingByID struct {
	SgMapping ResponseIPToSgtMappingGetIPToSgtMappingByIDSgMapping `json:"SGMapping,omitempty"` //
}

type ResponseIPToSgtMappingGetIPToSgtMappingByIDSgMapping struct {
	ID           string                                                   `json:"id,omitempty"`           //
	Name         string                                                   `json:"name,omitempty"`         //
	Sgt          string                                                   `json:"sgt,omitempty"`          // Mandatory unless mappingGroup is set
	DeployTo     string                                                   `json:"deployTo,omitempty"`     // Mandatory unless mappingGroup is set or unless deployType=ALL
	DeployType   string                                                   `json:"deployType,omitempty"`   // Allowed values: - ALL, - ND, - NDG
	HostName     string                                                   `json:"hostName,omitempty"`     // Mandatory if hostIp is empty
	HostIP       string                                                   `json:"hostIp,omitempty"`       // Mandatory if hostName is empty -- valid IP
	MappingGroup string                                                   `json:"mappingGroup,omitempty"` // Mapping Group Id. Mandatory unless sgt and deployTo and deployType are set
	Link         ResponseIPToSgtMappingGetIPToSgtMappingByIDSgMappingLink `json:"link,omitempty"`         //
}

type ResponseIPToSgtMappingGetIPToSgtMappingByIDSgMappingLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIPToSgtMappingUpdateIPToSgtMappingByID struct {
	UpdatedFieldsList ResponseIPToSgtMappingUpdateIPToSgtMappingByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseIPToSgtMappingUpdateIPToSgtMappingByIDUpdatedFieldsList struct {
	UpdatedField ResponseIPToSgtMappingUpdateIPToSgtMappingByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                      `json:"field,omitempty"`        //
	OldValue     string                                                                      `json:"oldValue,omitempty"`     //
	NewValue     string                                                                      `json:"newValue,omitempty"`     //
}

type ResponseIPToSgtMappingUpdateIPToSgtMappingByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseIPToSgtMappingGetIPToSgtMapping struct {
	SearchResult ResponseIPToSgtMappingGetIPToSgtMappingSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseIPToSgtMappingGetIPToSgtMappingSearchResult struct {
	Total        int                                                               `json:"total,omitempty"`        //
	Resources    []ResponseIPToSgtMappingGetIPToSgtMappingSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseIPToSgtMappingGetIPToSgtMappingSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseIPToSgtMappingGetIPToSgtMappingSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseIPToSgtMappingGetIPToSgtMappingSearchResultResources struct {
	ID   string                                                           `json:"id,omitempty"`   //
	Name string                                                           `json:"name,omitempty"` //
	Link ResponseIPToSgtMappingGetIPToSgtMappingSearchResultResourcesLink `json:"link,omitempty"` //
}

type ResponseIPToSgtMappingGetIPToSgtMappingSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIPToSgtMappingGetIPToSgtMappingSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIPToSgtMappingGetIPToSgtMappingSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIPToSgtMappingGetVersion struct {
	VersionInfo ResponseIPToSgtMappingGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseIPToSgtMappingGetVersionVersionInfo struct {
	CurrentServerVersion string                                          `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                          `json:"supportedVersions,omitempty"`    //
	Link                 ResponseIPToSgtMappingGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseIPToSgtMappingGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIPToSgtMappingMonitorBulkStatusIPToSgtMapping struct {
	BulkStatus ResponseIPToSgtMappingMonitorBulkStatusIPToSgtMappingBulkStatus `json:"BulkStatus,omitempty"` //
}

type ResponseIPToSgtMappingMonitorBulkStatusIPToSgtMappingBulkStatus struct {
	BulkID          string                                                                           `json:"bulkId,omitempty"`          //
	MediaType       string                                                                           `json:"mediaType,omitempty"`       //
	ExecutionStatus string                                                                           `json:"executionStatus,omitempty"` //
	OperationType   string                                                                           `json:"operationType,omitempty"`   //
	StartTime       string                                                                           `json:"startTime,omitempty"`       //
	ResourcesCount  int                                                                              `json:"resourcesCount,omitempty"`  //
	SuccessCount    int                                                                              `json:"successCount,omitempty"`    //
	FailCount       int                                                                              `json:"failCount,omitempty"`       //
	ResourcesStatus []ResponseIPToSgtMappingMonitorBulkStatusIPToSgtMappingBulkStatusResourcesStatus `json:"resourcesStatus,omitempty"` //
}

type ResponseIPToSgtMappingMonitorBulkStatusIPToSgtMappingBulkStatusResourcesStatus struct {
	ID                      string `json:"id,omitempty"`                      //
	Name                    string `json:"name,omitempty"`                    //
	Description             string `json:"description,omitempty"`             //
	ResourceExecutionStatus string `json:"resourceExecutionStatus,omitempty"` //
	Status                  string `json:"status,omitempty"`                  //
}

type RequestIPToSgtMappingUpdateIPToSgtMappingByID struct {
	SgMapping *RequestIPToSgtMappingUpdateIPToSgtMappingByIDSgMapping `json:"SGMapping,omitempty"` //
}

type RequestIPToSgtMappingUpdateIPToSgtMappingByIDSgMapping struct {
	ID           string `json:"id,omitempty"`           //
	Name         string `json:"name,omitempty"`         //
	Sgt          string `json:"sgt,omitempty"`          // Mandatory unless mappingGroup is set
	DeployTo     string `json:"deployTo,omitempty"`     // Mandatory unless mappingGroup is set or unless deployType=ALL
	DeployType   string `json:"deployType,omitempty"`   // Allowed values: - ALL, - ND, - NDG
	HostName     string `json:"hostName,omitempty"`     // Mandatory if hostIp is empty
	HostIP       string `json:"hostIp,omitempty"`       // Mandatory if hostName is empty -- valid IP
	MappingGroup string `json:"mappingGroup,omitempty"` // Mapping Group Id. Mandatory unless sgt and deployTo and deployType are set
}

type RequestIPToSgtMappingCreateIPToSgtMapping struct {
	SgMapping *RequestIPToSgtMappingCreateIPToSgtMappingSgMapping `json:"SGMapping,omitempty"` //
}

type RequestIPToSgtMappingCreateIPToSgtMappingSgMapping struct {
	Name         string `json:"name,omitempty"`         //
	Sgt          string `json:"sgt,omitempty"`          // Mandatory unless mappingGroup is set
	DeployTo     string `json:"deployTo,omitempty"`     // Mandatory unless mappingGroup is set or unless deployType=ALL
	DeployType   string `json:"deployType,omitempty"`   // Allowed values: - ALL, - ND, - NDG
	HostName     string `json:"hostName,omitempty"`     // Mandatory if hostIp is empty
	HostIP       string `json:"hostIp,omitempty"`       // Mandatory if hostName is empty -- valid IP
	MappingGroup string `json:"mappingGroup,omitempty"` // Mapping Group Id. Mandatory unless sgt and deployTo and deployType are set
}

type RequestIPToSgtMappingBulkRequestForIPToSgtMapping struct {
	SgMappingBulkRequest *RequestIPToSgtMappingBulkRequestForIPToSgtMappingSgMappingBulkRequest `json:"SGMappingBulkRequest,omitempty"` //
}

type RequestIPToSgtMappingBulkRequestForIPToSgtMappingSgMappingBulkRequest struct {
	OperationType     string `json:"operationType,omitempty"`     //
	ResourceMediaType string `json:"resourceMediaType,omitempty"` //
}

//GetIPToSgtMappingByID Get IP to SGT mapping by ID
/* This API allows the client to get an IP to SGT mapping by ID.

@param id id path parameter.
*/
func (s *IPToSgtMappingService) GetIPToSgtMappingByID(id string) (*ResponseIPToSgtMappingGetIPToSgtMappingByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmapping/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIPToSgtMappingGetIPToSgtMappingByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIpToSgtMappingById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIPToSgtMappingGetIPToSgtMappingByID)
	return result, response, err

}

//GetIPToSgtMapping Get all IP to SGT mappings
/* This API allows the client to get all the IP to SGT mappings.

Filter:

[hostName, groupName, ip, sgtName]

To search guest users by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:

[hostName, groupName, ip, sgtName]

@param getIPToSGTMappingQueryParams Filtering parameter
*/
func (s *IPToSgtMappingService) GetIPToSgtMapping(getIPToSGTMappingQueryParams *GetIPToSgtMappingQueryParams) (*ResponseIPToSgtMappingGetIPToSgtMapping, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmapping"

	queryString, _ := query.Values(getIPToSGTMappingQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseIPToSgtMappingGetIPToSgtMapping{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIpToSgtMapping")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIPToSgtMappingGetIPToSgtMapping)
	return result, response, err

}

//GetVersion Get IP to SGT mapping version information
/* This API helps to retrieve the version information related to IP to SGT mapping.

 */
func (s *IPToSgtMappingService) GetVersion() (*ResponseIPToSgtMappingGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmapping/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIPToSgtMappingGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIPToSgtMappingGetVersion)
	return result, response, err

}

//MonitorBulkStatusIPToSgtMapping Monitor bulk request
/* This API allows the client to monitor the bulk request.

@param bulkid bulkid path parameter.
*/
func (s *IPToSgtMappingService) MonitorBulkStatusIPToSgtMapping(bulkid string) (*ResponseIPToSgtMappingMonitorBulkStatusIPToSgtMapping, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmapping/bulk/{bulkid}"
	path = strings.Replace(path, "{bulkid}", fmt.Sprintf("%v", bulkid), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIPToSgtMappingMonitorBulkStatusIPToSgtMapping{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MonitorBulkStatusIpToSgtMapping")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIPToSgtMappingMonitorBulkStatusIPToSgtMapping)
	return result, response, err

}

//CreateIPToSgtMapping Create IP to SGT mapping
/* This API creates an IP to SGT mapping.

 */
func (s *IPToSgtMappingService) CreateIPToSgtMapping(requestIPToSgtMappingCreateIPToSGTMapping *RequestIPToSgtMappingCreateIPToSgtMapping) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmapping"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIPToSgtMappingCreateIPToSGTMapping).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateIpToSgtMapping")
	}

	return response, err

}

//DeployIPToSgtMappingByID Deploy IP to SGT mapping by ID
/* This API allows the client to deploy an IP to SGT mapping by ID.
Only one Deploy process can run at any given time

@param id id path parameter.
*/
func (s *IPToSgtMappingService) DeployIPToSgtMappingByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmapping/{id}/deploy"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

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
		return response, fmt.Errorf("error with operation DeployIpToSgtMappingById")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeployAllIPToSgtMapping Deploy all IP to SGT mappings
/* This API allows the client to deploy all the IP to SGT mappings.
Only one Deploy process can run at any given time

*/
func (s *IPToSgtMappingService) DeployAllIPToSgtMapping() (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmapping/deployall"

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
		return response, fmt.Errorf("error with operation DeployAllIpToSgtMapping")
	}

	getCSFRToken(response.Header())

	return response, err

}

//GetDeployStatusIPToSgtMapping Get IP to SGT mapping deployment status
/* This API allows the client to get the IP to SGT mapping deployment status.
Deploy Status will show last Deploy command output. The information will be saved until the next Deploy command

*/
func (s *IPToSgtMappingService) GetDeployStatusIPToSgtMapping() (*ResponseIPToSgtMappingGetDeployStatusIPToSgtMapping, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmapping/deploy/status"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIPToSgtMappingGetDeployStatusIPToSgtMapping{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeployStatusIpToSgtMapping")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIPToSgtMappingGetDeployStatusIPToSgtMapping)
	return result, response, err

}

//UpdateIPToSgtMappingByID Update IP to SGT mapping by ID
/* This API allows the client to update an IP to SGT mapping by ID.

@param id id path parameter.
*/
func (s *IPToSgtMappingService) UpdateIPToSgtMappingByID(id string, requestIPToSgtMappingUpdateIPToSGTMappingById *RequestIPToSgtMappingUpdateIPToSgtMappingByID) (*ResponseIPToSgtMappingUpdateIPToSgtMappingByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmapping/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIPToSgtMappingUpdateIPToSGTMappingById).
		SetResult(&ResponseIPToSgtMappingUpdateIPToSgtMappingByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateIpToSgtMappingById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIPToSgtMappingUpdateIPToSgtMappingByID)
	return result, response, err

}

//BulkRequestForIPToSgtMapping Submit bulk request
/* This API allows the client to submit the bulk request.

 */
func (s *IPToSgtMappingService) BulkRequestForIPToSgtMapping(requestIPToSgtMappingBulkRequestForIPToSGTMapping *RequestIPToSgtMappingBulkRequestForIPToSgtMapping) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmapping/bulk/submit"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIPToSgtMappingBulkRequestForIPToSGTMapping).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation BulkRequestForIpToSgtMapping")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteIPToSgtMappingByID Delete IP to SGT mapping
/* This API deletes an IP to SGT mapping.

@param id id path parameter.
*/
func (s *IPToSgtMappingService) DeleteIPToSgtMappingByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmapping/{id}"
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
		return response, fmt.Errorf("error with operation DeleteIpToSgtMappingById")
	}

	getCSFRToken(response.Header())

	return response, err

}
