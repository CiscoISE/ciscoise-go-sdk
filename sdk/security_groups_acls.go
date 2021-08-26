package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SecurityGroupsACLsService service

type GetSecurityGroupsACLQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseSecurityGroupsACLsGetSecurityGroupsACLByID struct {
	Sgacl *ResponseSecurityGroupsACLsGetSecurityGroupsACLByIDSgacl `json:"Sgacl,omitempty"` //
}

type ResponseSecurityGroupsACLsGetSecurityGroupsACLByIDSgacl struct {
	ID              string                                                       `json:"id,omitempty"`              //
	Name            string                                                       `json:"name,omitempty"`            //
	Description     string                                                       `json:"description,omitempty"`     //
	GenerationID    string                                                       `json:"generationId,omitempty"`    //
	ACLcontent      string                                                       `json:"aclcontent,omitempty"`      //
	IsReadOnly      *bool                                                        `json:"isReadOnly,omitempty"`      //
	ModelledContent *interface{}                                                 `json:"modelledContent,omitempty"` // Modelled content of contract
	IPVersion       string                                                       `json:"ipVersion,omitempty"`       // Allowed values: - IPV4, - IPV6, - IP_AGNOSTIC
	Link            *ResponseSecurityGroupsACLsGetSecurityGroupsACLByIDSgaclLink `json:"link,omitempty"`            //
}

type ResponseSecurityGroupsACLsGetSecurityGroupsACLByIDSgaclLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSecurityGroupsACLsUpdateSecurityGroupsACLByID struct {
	UpdatedFieldsList *ResponseSecurityGroupsACLsUpdateSecurityGroupsACLByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseSecurityGroupsACLsUpdateSecurityGroupsACLByIDUpdatedFieldsList struct {
	UpdatedField *[]ResponseSecurityGroupsACLsUpdateSecurityGroupsACLByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                                `json:"field,omitempty"`        //
	OldValue     string                                                                                `json:"oldValue,omitempty"`     //
	NewValue     string                                                                                `json:"newValue,omitempty"`     //
}

type ResponseSecurityGroupsACLsUpdateSecurityGroupsACLByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseSecurityGroupsACLsGetSecurityGroupsACL struct {
	SearchResult *ResponseSecurityGroupsACLsGetSecurityGroupsACLSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseSecurityGroupsACLsGetSecurityGroupsACLSearchResult struct {
	Total        *int                                                                    `json:"total,omitempty"`        //
	Resources    *[]ResponseSecurityGroupsACLsGetSecurityGroupsACLSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponseSecurityGroupsACLsGetSecurityGroupsACLSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseSecurityGroupsACLsGetSecurityGroupsACLSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseSecurityGroupsACLsGetSecurityGroupsACLSearchResultResources struct {
	ID          string                                                                   `json:"id,omitempty"`          //
	Name        string                                                                   `json:"name,omitempty"`        //
	Description string                                                                   `json:"description,omitempty"` //
	Link        *ResponseSecurityGroupsACLsGetSecurityGroupsACLSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseSecurityGroupsACLsGetSecurityGroupsACLSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSecurityGroupsACLsGetSecurityGroupsACLSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSecurityGroupsACLsGetSecurityGroupsACLSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSecurityGroupsACLsGetVersion struct {
	VersionInfo *ResponseSecurityGroupsACLsGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseSecurityGroupsACLsGetVersionVersionInfo struct {
	CurrentServerVersion string                                               `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                               `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseSecurityGroupsACLsGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseSecurityGroupsACLsGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSecurityGroupsACLsMonitorBulkStatusSecurityGroupsACL struct {
	BulkStatus *ResponseSecurityGroupsACLsMonitorBulkStatusSecurityGroupsACLBulkStatus `json:"BulkStatus,omitempty"` //
}

type ResponseSecurityGroupsACLsMonitorBulkStatusSecurityGroupsACLBulkStatus struct {
	BulkID          string                                                                                   `json:"bulkId,omitempty"`          //
	MediaType       string                                                                                   `json:"mediaType,omitempty"`       //
	ExecutionStatus string                                                                                   `json:"executionStatus,omitempty"` //
	OperationType   string                                                                                   `json:"operationType,omitempty"`   //
	StartTime       string                                                                                   `json:"startTime,omitempty"`       //
	ResourcesCount  *int                                                                                     `json:"resourcesCount,omitempty"`  //
	SuccessCount    *int                                                                                     `json:"successCount,omitempty"`    //
	FailCount       *int                                                                                     `json:"failCount,omitempty"`       //
	ResourcesStatus *[]ResponseSecurityGroupsACLsMonitorBulkStatusSecurityGroupsACLBulkStatusResourcesStatus `json:"resourcesStatus,omitempty"` //
}

type ResponseSecurityGroupsACLsMonitorBulkStatusSecurityGroupsACLBulkStatusResourcesStatus struct {
	ID                      string `json:"id,omitempty"`                      //
	Name                    string `json:"name,omitempty"`                    //
	Description             string `json:"description,omitempty"`             //
	ResourceExecutionStatus string `json:"resourceExecutionStatus,omitempty"` //
	Status                  string `json:"status,omitempty"`                  //
}

type RequestSecurityGroupsACLsUpdateSecurityGroupsACLByID struct {
	Sgacl *RequestSecurityGroupsACLsUpdateSecurityGroupsACLByIDSgacl `json:"Sgacl,omitempty"` //
}

type RequestSecurityGroupsACLsUpdateSecurityGroupsACLByIDSgacl struct {
	ID              string       `json:"id,omitempty"`              //
	Name            string       `json:"name,omitempty"`            //
	Description     string       `json:"description,omitempty"`     //
	GenerationID    string       `json:"generationId,omitempty"`    //
	ACLcontent      string       `json:"aclcontent,omitempty"`      //
	IsReadOnly      *bool        `json:"isReadOnly,omitempty"`      //
	ModelledContent *interface{} `json:"modelledContent,omitempty"` // Modelled content of contract
	IPVersion       string       `json:"ipVersion,omitempty"`       // Allowed values: - IPV4, - IPV6, - IP_AGNOSTIC
}

type RequestSecurityGroupsACLsCreateSecurityGroupsACL struct {
	Sgacl *RequestSecurityGroupsACLsCreateSecurityGroupsACLSgacl `json:"Sgacl,omitempty"` //
}

type RequestSecurityGroupsACLsCreateSecurityGroupsACLSgacl struct {
	Name            string       `json:"name,omitempty"`            //
	Description     string       `json:"description,omitempty"`     //
	GenerationID    string       `json:"generationId,omitempty"`    //
	ACLcontent      string       `json:"aclcontent,omitempty"`      //
	IsReadOnly      *bool        `json:"isReadOnly,omitempty"`      //
	ModelledContent *interface{} `json:"modelledContent,omitempty"` // Modelled content of contract
	IPVersion       string       `json:"ipVersion,omitempty"`       // Allowed values: - IPV4, - IPV6, - IP_AGNOSTIC
}

type RequestSecurityGroupsACLsBulkRequestForSecurityGroupsACL struct {
	SgaclBulkRequest *RequestSecurityGroupsACLsBulkRequestForSecurityGroupsACLSgaclBulkRequest `json:"SgaclBulkRequest,omitempty"` //
}

type RequestSecurityGroupsACLsBulkRequestForSecurityGroupsACLSgaclBulkRequest struct {
	OperationType     string `json:"operationType,omitempty"`     //
	ResourceMediaType string `json:"resourceMediaType,omitempty"` //
}

//GetSecurityGroupsACLByID Get security group ACL by ID
/* This API allows the client to get a security group ACL by ID.

@param id id path parameter.
*/
func (s *SecurityGroupsACLsService) GetSecurityGroupsACLByID(id string) (*ResponseSecurityGroupsACLsGetSecurityGroupsACLByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgacl/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityGroupsACLsGetSecurityGroupsACLByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSecurityGroupsAclById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSecurityGroupsACLsGetSecurityGroupsACLByID)
	return result, response, err

}

//GetSecurityGroupsACL Get all security group ACLs
/* This API allows the client to get all the security group ACLs.

Filter:

[ipVersion, name, description]

To search guest users by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:

[ipVersion, name, description]

@param getSecurityGroupsAclQueryParams Filtering parameter
*/
func (s *SecurityGroupsACLsService) GetSecurityGroupsACL(getSecurityGroupsAclQueryParams *GetSecurityGroupsACLQueryParams) (*ResponseSecurityGroupsACLsGetSecurityGroupsACL, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgacl"

	queryString, _ := query.Values(getSecurityGroupsAclQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSecurityGroupsACLsGetSecurityGroupsACL{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSecurityGroupsAcl")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSecurityGroupsACLsGetSecurityGroupsACL)
	return result, response, err

}

//GetVersion Get security group ACLs version information
/* This API helps to retrieve the version information related to the security group ACLs.

 */
func (s *SecurityGroupsACLsService) GetVersion() (*ResponseSecurityGroupsACLsGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgacl/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityGroupsACLsGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSecurityGroupsACLsGetVersion)
	return result, response, err

}

//MonitorBulkStatusSecurityGroupsACL Monitor bulk request
/* This API allows the client to monitor the bulk request.

@param bulkid bulkid path parameter.
*/
func (s *SecurityGroupsACLsService) MonitorBulkStatusSecurityGroupsACL(bulkid string) (*ResponseSecurityGroupsACLsMonitorBulkStatusSecurityGroupsACL, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgacl/bulk/{bulkid}"
	path = strings.Replace(path, "{bulkid}", fmt.Sprintf("%v", bulkid), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityGroupsACLsMonitorBulkStatusSecurityGroupsACL{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MonitorBulkStatusSecurityGroupsAcl")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSecurityGroupsACLsMonitorBulkStatusSecurityGroupsACL)
	return result, response, err

}

//CreateSecurityGroupsACL Create security group ACL
/* This API creates a security group ACL.

 */
func (s *SecurityGroupsACLsService) CreateSecurityGroupsACL(requestSecurityGroupsACLsCreateSecurityGroupsAcl *RequestSecurityGroupsACLsCreateSecurityGroupsACL) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgacl"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSecurityGroupsACLsCreateSecurityGroupsAcl).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateSecurityGroupsAcl")
	}

	return response, err

}

//UpdateSecurityGroupsACLByID Update security group ACL
/* This API allows the client to update a security group ACL.

@param id id path parameter.
*/
func (s *SecurityGroupsACLsService) UpdateSecurityGroupsACLByID(id string, requestSecurityGroupsACLsUpdateSecurityGroupsAclById *RequestSecurityGroupsACLsUpdateSecurityGroupsACLByID) (*ResponseSecurityGroupsACLsUpdateSecurityGroupsACLByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgacl/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSecurityGroupsACLsUpdateSecurityGroupsAclById).
		SetResult(&ResponseSecurityGroupsACLsUpdateSecurityGroupsACLByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSecurityGroupsAclById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSecurityGroupsACLsUpdateSecurityGroupsACLByID)
	return result, response, err

}

//BulkRequestForSecurityGroupsACL Submit bulk request
/* This API allows the client to submit the bulk request.

 */
func (s *SecurityGroupsACLsService) BulkRequestForSecurityGroupsACL(requestSecurityGroupsACLsBulkRequestForSecurityGroupsAcl *RequestSecurityGroupsACLsBulkRequestForSecurityGroupsACL) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgacl/bulk/submit"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSecurityGroupsACLsBulkRequestForSecurityGroupsAcl).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation BulkRequestForSecurityGroupsAcl")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteSecurityGroupsACLByID Delete security group ACL
/* This API deletes a security group ACL.

@param id id path parameter.
*/
func (s *SecurityGroupsACLsService) DeleteSecurityGroupsACLByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgacl/{id}"
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
		return response, fmt.Errorf("error with operation DeleteSecurityGroupsAclById")
	}

	getCSFRToken(response.Header())

	return response, err

}
