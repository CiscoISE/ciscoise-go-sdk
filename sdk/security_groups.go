package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SecurityGroupsService service

type GetSecurityGroupsQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseSecurityGroupsGetSecurityGroupByID struct {
	Sgt *ResponseSecurityGroupsGetSecurityGroupByIDSgt `json:"Sgt,omitempty"` //
}

type ResponseSecurityGroupsGetSecurityGroupByIDSgt struct {
	ID              string                                             `json:"id,omitempty"`              //
	Name            string                                             `json:"name,omitempty"`            //
	Description     string                                             `json:"description,omitempty"`     //
	Value           *int                                               `json:"value,omitempty"`           // Value range: 2 ot 65519 or -1 to auto-generate
	GenerationID    *int                                               `json:"generationId,omitempty"`    //
	IsReadOnly      *bool                                              `json:"isReadOnly,omitempty"`      //
	PropogateToAPIc *bool                                              `json:"propogateToApic,omitempty"` //
	DefaultSgACLs   *[]interface{}                                     `json:"defaultSGACLs,omitempty"`   //
	Link            *ResponseSecurityGroupsGetSecurityGroupByIDSgtLink `json:"link,omitempty"`            //
}

type ResponseSecurityGroupsGetSecurityGroupByIDSgtLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSecurityGroupsUpdateSecurityGroupByID struct {
	UpdatedFieldsList *ResponseSecurityGroupsUpdateSecurityGroupByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseSecurityGroupsUpdateSecurityGroupByIDUpdatedFieldsList struct {
	UpdatedField *[]ResponseSecurityGroupsUpdateSecurityGroupByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                        `json:"field,omitempty"`        //
	OldValue     string                                                                        `json:"oldValue,omitempty"`     //
	NewValue     string                                                                        `json:"newValue,omitempty"`     //
}

type ResponseSecurityGroupsUpdateSecurityGroupByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseSecurityGroupsGetSecurityGroups struct {
	SearchResult *ResponseSecurityGroupsGetSecurityGroupsSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseSecurityGroupsGetSecurityGroupsSearchResult struct {
	Total        *int                                                             `json:"total,omitempty"`        //
	Resources    *[]ResponseSecurityGroupsGetSecurityGroupsSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponseSecurityGroupsGetSecurityGroupsSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseSecurityGroupsGetSecurityGroupsSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseSecurityGroupsGetSecurityGroupsSearchResultResources struct {
	ID          string                                                            `json:"id,omitempty"`          //
	Name        string                                                            `json:"name,omitempty"`        //
	Description string                                                            `json:"description,omitempty"` //
	Link        *ResponseSecurityGroupsGetSecurityGroupsSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseSecurityGroupsGetSecurityGroupsSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSecurityGroupsGetSecurityGroupsSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSecurityGroupsGetSecurityGroupsSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSecurityGroupsGetVersion struct {
	VersionInfo *ResponseSecurityGroupsGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseSecurityGroupsGetVersionVersionInfo struct {
	CurrentServerVersion string                                           `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                           `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseSecurityGroupsGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseSecurityGroupsGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSecurityGroupsMonitorBulkStatusSecurityGroup struct {
	BulkStatus *ResponseSecurityGroupsMonitorBulkStatusSecurityGroupBulkStatus `json:"BulkStatus,omitempty"` //
}

type ResponseSecurityGroupsMonitorBulkStatusSecurityGroupBulkStatus struct {
	BulkID          string                                                                           `json:"bulkId,omitempty"`          //
	MediaType       string                                                                           `json:"mediaType,omitempty"`       //
	ExecutionStatus string                                                                           `json:"executionStatus,omitempty"` //
	OperationType   string                                                                           `json:"operationType,omitempty"`   //
	StartTime       string                                                                           `json:"startTime,omitempty"`       //
	ResourcesCount  *int                                                                             `json:"resourcesCount,omitempty"`  //
	SuccessCount    *int                                                                             `json:"successCount,omitempty"`    //
	FailCount       *int                                                                             `json:"failCount,omitempty"`       //
	ResourcesStatus *[]ResponseSecurityGroupsMonitorBulkStatusSecurityGroupBulkStatusResourcesStatus `json:"resourcesStatus,omitempty"` //
}

type ResponseSecurityGroupsMonitorBulkStatusSecurityGroupBulkStatusResourcesStatus struct {
	ID                      string `json:"id,omitempty"`                      //
	Name                    string `json:"name,omitempty"`                    //
	Description             string `json:"description,omitempty"`             //
	ResourceExecutionStatus string `json:"resourceExecutionStatus,omitempty"` //
	Status                  string `json:"status,omitempty"`                  //
}

type RequestSecurityGroupsUpdateSecurityGroupByID struct {
	Sgt *RequestSecurityGroupsUpdateSecurityGroupByIDSgt `json:"Sgt,omitempty"` //
}

type RequestSecurityGroupsUpdateSecurityGroupByIDSgt struct {
	ID              string         `json:"id,omitempty"`              //
	Name            string         `json:"name,omitempty"`            //
	Description     string         `json:"description,omitempty"`     //
	Value           *int           `json:"value,omitempty"`           // Value range: 2 ot 65519 or -1 to auto-generate
	GenerationID    *int           `json:"generationId,omitempty"`    //
	IsReadOnly      *bool          `json:"isReadOnly,omitempty"`      //
	PropogateToAPIc *bool          `json:"propogateToApic,omitempty"` //
	DefaultSgACLs   *[]interface{} `json:"defaultSGACLs,omitempty"`   //
}

type RequestSecurityGroupsCreateSecurityGroup struct {
	Sgt *RequestSecurityGroupsCreateSecurityGroupSgt `json:"Sgt,omitempty"` //
}

type RequestSecurityGroupsCreateSecurityGroupSgt struct {
	Name            string         `json:"name,omitempty"`            //
	Description     string         `json:"description,omitempty"`     //
	Value           *int           `json:"value,omitempty"`           // Value range: 2 ot 65519 or -1 to auto-generate
	GenerationID    *int           `json:"generationId,omitempty"`    //
	IsReadOnly      *bool          `json:"isReadOnly,omitempty"`      //
	PropogateToAPIc *bool          `json:"propogateToApic,omitempty"` //
	DefaultSgACLs   *[]interface{} `json:"defaultSGACLs,omitempty"`   //
}

type RequestSecurityGroupsBulkRequestForSecurityGroup struct {
	SgtBulkRequest *RequestSecurityGroupsBulkRequestForSecurityGroupSgtBulkRequest `json:"SgtBulkRequest,omitempty"` //
}

type RequestSecurityGroupsBulkRequestForSecurityGroupSgtBulkRequest struct {
	OperationType     string `json:"operationType,omitempty"`     //
	ResourceMediaType string `json:"resourceMediaType,omitempty"` //
}

//GetSecurityGroupByID Get security group by ID
/* This API allows the client to get a security group by ID.

@param id id path parameter.
*/
func (s *SecurityGroupsService) GetSecurityGroupByID(id string) (*ResponseSecurityGroupsGetSecurityGroupByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgt/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityGroupsGetSecurityGroupByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSecurityGroupById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSecurityGroupsGetSecurityGroupByID)
	return result, response, err

}

//GetSecurityGroups Get all security groups
/* This API allows the client to get all the security groups.

Filter:

[propogateToApic, name, description, value]

To search resources by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:

[name, description, value]

@param getSecurityGroupsQueryParams Filtering parameter
*/
func (s *SecurityGroupsService) GetSecurityGroups(getSecurityGroupsQueryParams *GetSecurityGroupsQueryParams) (*ResponseSecurityGroupsGetSecurityGroups, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgt"

	queryString, _ := query.Values(getSecurityGroupsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSecurityGroupsGetSecurityGroups{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSecurityGroups")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSecurityGroupsGetSecurityGroups)
	return result, response, err

}

//GetVersion Get security groups version information
/* This API helps to retrieve the version information related to the security groups.

 */
func (s *SecurityGroupsService) GetVersion() (*ResponseSecurityGroupsGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgt/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityGroupsGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSecurityGroupsGetVersion)
	return result, response, err

}

//MonitorBulkStatusSecurityGroup Monitor bulk request
/* This API allows the client to monitor the bulk request.

@param bulkid bulkid path parameter.
*/
func (s *SecurityGroupsService) MonitorBulkStatusSecurityGroup(bulkid string) (*ResponseSecurityGroupsMonitorBulkStatusSecurityGroup, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgt/bulk/{bulkid}"
	path = strings.Replace(path, "{bulkid}", fmt.Sprintf("%v", bulkid), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityGroupsMonitorBulkStatusSecurityGroup{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MonitorBulkStatusSecurityGroup")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSecurityGroupsMonitorBulkStatusSecurityGroup)
	return result, response, err

}

//CreateSecurityGroup Create security group
/* This API creates a security group.

 */
func (s *SecurityGroupsService) CreateSecurityGroup(requestSecurityGroupsCreateSecurityGroup *RequestSecurityGroupsCreateSecurityGroup) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgt"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSecurityGroupsCreateSecurityGroup).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateSecurityGroup")
	}

	return response, err

}

//UpdateSecurityGroupByID Update security group
/* This API allows the client to update a security group.

@param id id path parameter.
*/
func (s *SecurityGroupsService) UpdateSecurityGroupByID(id string, requestSecurityGroupsUpdateSecurityGroupById *RequestSecurityGroupsUpdateSecurityGroupByID) (*ResponseSecurityGroupsUpdateSecurityGroupByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgt/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSecurityGroupsUpdateSecurityGroupById).
		SetResult(&ResponseSecurityGroupsUpdateSecurityGroupByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSecurityGroupById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSecurityGroupsUpdateSecurityGroupByID)
	return result, response, err

}

//BulkRequestForSecurityGroup Submit bulk request
/* This API allows the client to submit the bulk request.

 */
func (s *SecurityGroupsService) BulkRequestForSecurityGroup(requestSecurityGroupsBulkRequestForSecurityGroup *RequestSecurityGroupsBulkRequestForSecurityGroup) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgt/bulk/submit"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSecurityGroupsBulkRequestForSecurityGroup).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation BulkRequestForSecurityGroup")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteSecurityGroupByID Delete security group
/* This API deletes a security group.

@param id id path parameter.
*/
func (s *SecurityGroupsService) DeleteSecurityGroupByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgt/{id}"
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
		return response, fmt.Errorf("error with operation DeleteSecurityGroupById")
	}

	getCSFRToken(response.Header())

	return response, err

}
