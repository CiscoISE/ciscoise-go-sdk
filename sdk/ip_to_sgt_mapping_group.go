package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type IPToSgtMappingGroupService service

type GetIPToSgtMappingGroupQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseIPToSgtMappingGroupGetDeployStatusIPToSgtMappingGroup struct {
	OperationResult ResponseIPToSgtMappingGroupGetDeployStatusIPToSgtMappingGroupOperationResult `json:"OperationResult,omitempty"` //
}

type ResponseIPToSgtMappingGroupGetDeployStatusIPToSgtMappingGroupOperationResult struct {
	ResultValue []ResponseIPToSgtMappingGroupGetDeployStatusIPToSgtMappingGroupOperationResultResultValue `json:"resultValue,omitempty"` //
}

type ResponseIPToSgtMappingGroupGetDeployStatusIPToSgtMappingGroupOperationResultResultValue struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupByID struct {
	SgMappingGroup ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupByIDSgMappingGroup `json:"SGMappingGroup,omitempty"` //
}

type ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupByIDSgMappingGroup struct {
	Name       string                                                                  `json:"name,omitempty"`       //
	Sgt        string                                                                  `json:"sgt,omitempty"`        // Mandatory unless mappingGroup is set
	DeployTo   string                                                                  `json:"deployTo,omitempty"`   // Mandatory unless mappingGroup is set or unless deployType=ALL
	DeployType string                                                                  `json:"deployType,omitempty"` // Allowed values: - ALL, - ND, - NDG
	Link       ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupByIDSgMappingGroupLink `json:"link,omitempty"`       //
}

type ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupByIDSgMappingGroupLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIPToSgtMappingGroupUpdateIPToSgtMappingGroupByID struct {
	UpdatedFieldsList ResponseIPToSgtMappingGroupUpdateIPToSgtMappingGroupByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseIPToSgtMappingGroupUpdateIPToSgtMappingGroupByIDUpdatedFieldsList struct {
	UpdatedField ResponseIPToSgtMappingGroupUpdateIPToSgtMappingGroupByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                                `json:"field,omitempty"`        //
	OldValue     string                                                                                `json:"oldValue,omitempty"`     //
	NewValue     string                                                                                `json:"newValue,omitempty"`     //
}

type ResponseIPToSgtMappingGroupUpdateIPToSgtMappingGroupByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseIPToSgtMappingGroupGetIPToSgtMappingGroup struct {
	SearchResult ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupSearchResult struct {
	Total        int                                                                       `json:"total,omitempty"`        //
	Resources    []ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupSearchResultResources struct {
	ID   string                                                                     `json:"id,omitempty"`   //
	Name string                                                                     `json:"name,omitempty"` //
	Link ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupSearchResultResourcesLink `json:"link,omitempty"` //
}

type ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIPToSgtMappingGroupGetVersion struct {
	VersionInfo ResponseIPToSgtMappingGroupGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseIPToSgtMappingGroupGetVersionVersionInfo struct {
	CurrentServerVersion string                                               `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                               `json:"supportedVersions,omitempty"`    //
	Link                 ResponseIPToSgtMappingGroupGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseIPToSgtMappingGroupGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseIPToSgtMappingGroupMonitorBulkStatusIPToSgtMappingGroup struct {
	BulkStatus ResponseIPToSgtMappingGroupMonitorBulkStatusIPToSgtMappingGroupBulkStatus `json:"BulkStatus,omitempty"` //
}

type ResponseIPToSgtMappingGroupMonitorBulkStatusIPToSgtMappingGroupBulkStatus struct {
	BulkID          string                                                                                     `json:"bulkId,omitempty"`          //
	MediaType       string                                                                                     `json:"mediaType,omitempty"`       //
	ExecutionStatus string                                                                                     `json:"executionStatus,omitempty"` //
	OperationType   string                                                                                     `json:"operationType,omitempty"`   //
	StartTime       string                                                                                     `json:"startTime,omitempty"`       //
	ResourcesCount  int                                                                                        `json:"resourcesCount,omitempty"`  //
	SuccessCount    int                                                                                        `json:"successCount,omitempty"`    //
	FailCount       int                                                                                        `json:"failCount,omitempty"`       //
	ResourcesStatus []ResponseIPToSgtMappingGroupMonitorBulkStatusIPToSgtMappingGroupBulkStatusResourcesStatus `json:"resourcesStatus,omitempty"` //
}

type ResponseIPToSgtMappingGroupMonitorBulkStatusIPToSgtMappingGroupBulkStatusResourcesStatus struct {
	ID                      string `json:"id,omitempty"`                      //
	Name                    string `json:"name,omitempty"`                    //
	Description             string `json:"description,omitempty"`             //
	ResourceExecutionStatus string `json:"resourceExecutionStatus,omitempty"` //
	Status                  string `json:"status,omitempty"`                  //
}

type RequestIPToSgtMappingGroupUpdateIPToSgtMappingGroupByID struct {
	SgMappingGroup *RequestIPToSgtMappingGroupUpdateIPToSgtMappingGroupByIDSgMappingGroup `json:"SGMappingGroup,omitempty"` //
}

type RequestIPToSgtMappingGroupUpdateIPToSgtMappingGroupByIDSgMappingGroup struct {
	Name       string `json:"name,omitempty"`       //
	Sgt        string `json:"sgt,omitempty"`        // Mandatory unless mappingGroup is set
	DeployTo   string `json:"deployTo,omitempty"`   // Mandatory unless mappingGroup is set or unless deployType=ALL
	DeployType string `json:"deployType,omitempty"` // Allowed values: - ALL, - ND, - NDG
}

type RequestIPToSgtMappingGroupCreateIPToSgtMappingGroup struct {
	SgMappingGroup *RequestIPToSgtMappingGroupCreateIPToSgtMappingGroupSgMappingGroup `json:"SGMappingGroup,omitempty"` //
}

type RequestIPToSgtMappingGroupCreateIPToSgtMappingGroupSgMappingGroup struct {
	Name       string `json:"name,omitempty"`       //
	Sgt        string `json:"sgt,omitempty"`        // Mandatory unless mappingGroup is set
	DeployTo   string `json:"deployTo,omitempty"`   // Mandatory unless mappingGroup is set or unless deployType=ALL
	DeployType string `json:"deployType,omitempty"` // Allowed values: - ALL, - ND, - NDG
}

type RequestIPToSgtMappingGroupBulkRequestForIPToSgtMappingGroup struct {
	SgMappingGroupBulkRequest *RequestIPToSgtMappingGroupBulkRequestForIPToSgtMappingGroupSgMappingGroupBulkRequest `json:"SGMappingGroupBulkRequest,omitempty"` //
}

type RequestIPToSgtMappingGroupBulkRequestForIPToSgtMappingGroupSgMappingGroupBulkRequest struct {
	OperationType     string `json:"operationType,omitempty"`     //
	ResourceMediaType string `json:"resourceMediaType,omitempty"` //
}

//GetIPToSgtMappingGroupByID Get IP to SGT mapping group by ID
/* This API allows the client to get an IP to SGT mapping group by ID.

@param id id path parameter.
*/
func (s *IPToSgtMappingGroupService) GetIPToSgtMappingGroupByID(id string) (*ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmappinggroup/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIpToSgtMappingGroupById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupByID)
	return result, response, err

}

//GetIPToSgtMappingGroup Get all IP to SGT mapping groups
/* This API allows the client to get all the IP to SGT mapping groups.

Filter:

[name, sgtName]

To search guest users by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:

[name, sgtName]

@param getIPToSGTMappingGroupQueryParams Filtering parameter
*/
func (s *IPToSgtMappingGroupService) GetIPToSgtMappingGroup(getIPToSGTMappingGroupQueryParams *GetIPToSgtMappingGroupQueryParams) (*ResponseIPToSgtMappingGroupGetIPToSgtMappingGroup, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmappinggroup"

	queryString, _ := query.Values(getIPToSGTMappingGroupQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseIPToSgtMappingGroupGetIPToSgtMappingGroup{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIpToSgtMappingGroup")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIPToSgtMappingGroupGetIPToSgtMappingGroup)
	return result, response, err

}

//GetVersion Get IP to SGT mapping group version information
/* This API helps to retrieve the version information related to the IP to SGT mapping group.

 */
func (s *IPToSgtMappingGroupService) GetVersion() (*ResponseIPToSgtMappingGroupGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmappinggroup/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIPToSgtMappingGroupGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIPToSgtMappingGroupGetVersion)
	return result, response, err

}

//MonitorBulkStatusIPToSgtMappingGroup Monitor bulk request
/* This API allows the client to monitor the bulk request.

@param bulkid bulkid path parameter.
*/
func (s *IPToSgtMappingGroupService) MonitorBulkStatusIPToSgtMappingGroup(bulkid string) (*ResponseIPToSgtMappingGroupMonitorBulkStatusIPToSgtMappingGroup, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmappinggroup/bulk/{bulkid}"
	path = strings.Replace(path, "{bulkid}", fmt.Sprintf("%v", bulkid), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIPToSgtMappingGroupMonitorBulkStatusIPToSgtMappingGroup{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MonitorBulkStatusIpToSgtMappingGroup")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIPToSgtMappingGroupMonitorBulkStatusIPToSgtMappingGroup)
	return result, response, err

}

//CreateIPToSgtMappingGroup Create IP to SGT mapping group
/* This API creates an IP to SGT mapping group.

 */
func (s *IPToSgtMappingGroupService) CreateIPToSgtMappingGroup(requestIPToSgtMappingGroupCreateIPToSGTMappingGroup *RequestIPToSgtMappingGroupCreateIPToSgtMappingGroup) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmappinggroup"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIPToSgtMappingGroupCreateIPToSGTMappingGroup).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateIpToSgtMappingGroup")
	}

	return response, err

}

//DeployIPToSgtMappingGroupByID Deploy IP to SGT mapping group by ID
/* This API allows the client to deploy an IP to SGT mapping group by ID.
Only one Deploy process can run at any given time

@param id id path parameter.
*/
func (s *IPToSgtMappingGroupService) DeployIPToSgtMappingGroupByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmappinggroup/{id}/deploy"
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
		return response, fmt.Errorf("error with operation DeployIpToSgtMappingGroupById")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeployAllIPToSgtMappingGroup Deploy all IP to SGT mapping groups
/* This API allows the client to deploy all the IP to SGT mapping groups.
Only one Deploy process can run at any given time

*/
func (s *IPToSgtMappingGroupService) DeployAllIPToSgtMappingGroup() (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmappinggroup/deployall"

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
		return response, fmt.Errorf("error with operation DeployAllIpToSgtMappingGroup")
	}

	getCSFRToken(response.Header())

	return response, err

}

//GetDeployStatusIPToSgtMappingGroup Get IP to SGT mapping group deployment status
/* This API allows the client to get the IP to SGT mapping group deployment status.
Deploy Status will show last Deploy command output. The information will be saved until the next Deploy command

*/
func (s *IPToSgtMappingGroupService) GetDeployStatusIPToSgtMappingGroup() (*ResponseIPToSgtMappingGroupGetDeployStatusIPToSgtMappingGroup, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmappinggroup/deploy/status"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIPToSgtMappingGroupGetDeployStatusIPToSgtMappingGroup{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeployStatusIpToSgtMappingGroup")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIPToSgtMappingGroupGetDeployStatusIPToSgtMappingGroup)
	return result, response, err

}

//UpdateIPToSgtMappingGroupByID Update IP to SGT mapping group by ID
/* This API allows the client to update an IP to SGT mapping group by ID.

@param id id path parameter.
*/
func (s *IPToSgtMappingGroupService) UpdateIPToSgtMappingGroupByID(id string, requestIPToSgtMappingGroupUpdateIPToSGTMappingGroupById *RequestIPToSgtMappingGroupUpdateIPToSgtMappingGroupByID) (*ResponseIPToSgtMappingGroupUpdateIPToSgtMappingGroupByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmappinggroup/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIPToSgtMappingGroupUpdateIPToSGTMappingGroupById).
		SetResult(&ResponseIPToSgtMappingGroupUpdateIPToSgtMappingGroupByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateIpToSgtMappingGroupById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseIPToSgtMappingGroupUpdateIPToSgtMappingGroupByID)
	return result, response, err

}

//BulkRequestForIPToSgtMappingGroup Submit bulk request
/* This API allows the client to submit the bulk request.

 */
func (s *IPToSgtMappingGroupService) BulkRequestForIPToSgtMappingGroup(requestIPToSgtMappingGroupBulkRequestForIPToSGTMappingGroup *RequestIPToSgtMappingGroupBulkRequestForIPToSgtMappingGroup) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmappinggroup/bulk/submit"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIPToSgtMappingGroupBulkRequestForIPToSGTMappingGroup).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation BulkRequestForIpToSgtMappingGroup")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteIPToSgtMappingGroupByID Delete IP to SGT mapping group
/* This API deletes an IP to SGT mapping group.

@param id id path parameter.
*/
func (s *IPToSgtMappingGroupService) DeleteIPToSgtMappingGroupByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgmappinggroup/{id}"
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
		return response, fmt.Errorf("error with operation DeleteIpToSgtMappingGroupById")
	}

	getCSFRToken(response.Header())

	return response, err

}
