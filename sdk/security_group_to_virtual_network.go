package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SecurityGroupToVirtualNetworkService service

type GetSecurityGroupsToVnToVLANQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByID struct {
	SgtVnVLANContainer *ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDSgtVnVLANContainer `json:"SgtVNVlanContainer,omitempty"` //
}

type ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDSgtVnVLANContainer struct {
	ID                 string                                                                                                      `json:"id,omitempty"`                 //
	Name               string                                                                                                      `json:"name,omitempty"`               //
	Description        string                                                                                                      `json:"description,omitempty"`        //
	SgtID              string                                                                                                      `json:"sgtId,omitempty"`              //
	Virtualnetworklist *[]ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklist `json:"virtualnetworklist,omitempty"` //
	Link               *ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDSgtVnVLANContainerLink                 `json:"link,omitempty"`               //
}

type ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklist struct {
	ID                    string                                                                                                           `json:"id,omitempty"`                    //
	Name                  string                                                                                                           `json:"name,omitempty"`                  //
	Description           string                                                                                                           `json:"description,omitempty"`           //
	DefaultVirtualNetwork *bool                                                                                                            `json:"defaultVirtualNetwork,omitempty"` //
	VLANs                 *[]ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklistVLANs `json:"vlans,omitempty"`                 //
}

type ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklistVLANs struct {
	ID          string `json:"id,omitempty"`          //
	Name        string `json:"name,omitempty"`        //
	Description string `json:"description,omitempty"` //
	DefaultVLAN *bool  `json:"defaultVlan,omitempty"` //
	MaxValue    *int   `json:"maxValue,omitempty"`    //
	Data        *bool  `json:"data,omitempty"`        //
}

type ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDSgtVnVLANContainerLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByID struct {
	UpdatedFieldsList *ResponseSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDUpdatedFieldsList struct {
	UpdatedField *[]ResponseSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                                                  `json:"field,omitempty"`        //
	OldValue     string                                                                                                  `json:"oldValue,omitempty"`     //
	NewValue     string                                                                                                  `json:"newValue,omitempty"`     //
}

type ResponseSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLAN struct {
	SearchResult *ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResult struct {
	Total        *int                                                                                      `json:"total,omitempty"`        //
	Resources    *[]ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResultResources struct {
	ID          string                                                                                     `json:"id,omitempty"`          //
	Name        string                                                                                     `json:"name,omitempty"`        //
	Description string                                                                                     `json:"description,omitempty"` //
	Link        *ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSecurityGroupToVirtualNetworkGetVersion struct {
	VersionInfo *ResponseSecurityGroupToVirtualNetworkGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseSecurityGroupToVirtualNetworkGetVersionVersionInfo struct {
	CurrentServerVersion string                                                          `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                                          `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseSecurityGroupToVirtualNetworkGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseSecurityGroupToVirtualNetworkGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSecurityGroupToVirtualNetworkMonitorBulkStatusSecurityGroupsToVnToVLAN struct {
	BulkStatus *ResponseSecurityGroupToVirtualNetworkMonitorBulkStatusSecurityGroupsToVnToVLANBulkStatus `json:"BulkStatus,omitempty"` //
}

type ResponseSecurityGroupToVirtualNetworkMonitorBulkStatusSecurityGroupsToVnToVLANBulkStatus struct {
	BulkID          string                                                                                                     `json:"bulkId,omitempty"`          //
	MediaType       string                                                                                                     `json:"mediaType,omitempty"`       //
	ExecutionStatus string                                                                                                     `json:"executionStatus,omitempty"` //
	OperationType   string                                                                                                     `json:"operationType,omitempty"`   //
	StartTime       string                                                                                                     `json:"startTime,omitempty"`       //
	ResourcesCount  *int                                                                                                       `json:"resourcesCount,omitempty"`  //
	SuccessCount    *int                                                                                                       `json:"successCount,omitempty"`    //
	FailCount       *int                                                                                                       `json:"failCount,omitempty"`       //
	ResourcesStatus *[]ResponseSecurityGroupToVirtualNetworkMonitorBulkStatusSecurityGroupsToVnToVLANBulkStatusResourcesStatus `json:"resourcesStatus,omitempty"` //
}

type ResponseSecurityGroupToVirtualNetworkMonitorBulkStatusSecurityGroupsToVnToVLANBulkStatusResourcesStatus struct {
	ID                      string `json:"id,omitempty"`                      //
	Name                    string `json:"name,omitempty"`                    //
	Description             string `json:"description,omitempty"`             //
	ResourceExecutionStatus string `json:"resourceExecutionStatus,omitempty"` //
	Status                  string `json:"status,omitempty"`                  //
}

type RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByID struct {
	SgtVnVLANContainer *RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainer `json:"SgtVNVlanContainer,omitempty"` //
}

type RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainer struct {
	ID                 string                                                                                                        `json:"id,omitempty"`                 //
	Name               string                                                                                                        `json:"name,omitempty"`               //
	Description        string                                                                                                        `json:"description,omitempty"`        //
	SgtID              string                                                                                                        `json:"sgtId,omitempty"`              //
	Virtualnetworklist *[]RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklist `json:"virtualnetworklist,omitempty"` //
}

type RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklist struct {
	ID                    string                                                                                                             `json:"id,omitempty"`                    //
	Name                  string                                                                                                             `json:"name,omitempty"`                  //
	Description           string                                                                                                             `json:"description,omitempty"`           //
	DefaultVirtualNetwork *bool                                                                                                              `json:"defaultVirtualNetwork,omitempty"` //
	VLANs                 *[]RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklistVLANs `json:"vlans,omitempty"`                 //
}

type RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklistVLANs struct {
	ID          string `json:"id,omitempty"`          //
	Name        string `json:"name,omitempty"`        //
	Description string `json:"description,omitempty"` //
	DefaultVLAN *bool  `json:"defaultVlan,omitempty"` //
	MaxValue    *int   `json:"maxValue,omitempty"`    //
	Data        *bool  `json:"data,omitempty"`        //
}

type RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLAN struct {
	SgtVnVLANContainer *RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLANSgtVnVLANContainer `json:"SgtVNVlanContainer,omitempty"` //
}

type RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLANSgtVnVLANContainer struct {
	ID                 string                                                                                                    `json:"id,omitempty"`                 //
	Name               string                                                                                                    `json:"name,omitempty"`               //
	Description        string                                                                                                    `json:"description,omitempty"`        //
	SgtID              string                                                                                                    `json:"sgtId,omitempty"`              //
	Virtualnetworklist *[]RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklist `json:"virtualnetworklist,omitempty"` //
}

type RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklist struct {
	ID                    string                                                                                                         `json:"id,omitempty"`                    //
	Name                  string                                                                                                         `json:"name,omitempty"`                  //
	Description           string                                                                                                         `json:"description,omitempty"`           //
	DefaultVirtualNetwork *bool                                                                                                          `json:"defaultVirtualNetwork,omitempty"` //
	VLANs                 *[]RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklistVLANs `json:"vlans,omitempty"`                 //
}

type RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklistVLANs struct {
	ID          string `json:"id,omitempty"`          //
	Name        string `json:"name,omitempty"`        //
	Description string `json:"description,omitempty"` //
	DefaultVLAN *bool  `json:"defaultVlan,omitempty"` //
	MaxValue    *int   `json:"maxValue,omitempty"`    //
	Data        *bool  `json:"data,omitempty"`        //
}

type RequestSecurityGroupToVirtualNetworkBulkRequestForSecurityGroupsToVnToVLAN struct {
	SgtVnVLANContainerBulkRequest *RequestSecurityGroupToVirtualNetworkBulkRequestForSecurityGroupsToVnToVLANSgtVnVLANContainerBulkRequest `json:"SgtVNVlanContainerBulkRequest,omitempty"` //
}

type RequestSecurityGroupToVirtualNetworkBulkRequestForSecurityGroupsToVnToVLANSgtVnVLANContainerBulkRequest struct {
	OperationType     string `json:"operationType,omitempty"`     //
	ResourceMediaType string `json:"resourceMediaType,omitempty"` //
}

//GetSecurityGroupsToVnToVLANByID Get security group to virtual network by ID
/* This API allows the client to get a security group to virtual network by ID.

@param id id path parameter.
*/
func (s *SecurityGroupToVirtualNetworkService) GetSecurityGroupsToVnToVLANByID(id string) (*ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgtvnvlan/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSecurityGroupsToVnToVlanById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByID)
	return result, response, err

}

//GetSecurityGroupsToVnToVLAN Get all security group to virtual networks
/* This API allows the client to get all the security group ACL to virtual networks.

Filter:

[sgtId]

To search guest users by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


@param getSecurityGroupsToVnToVlanQueryParams Filtering parameter
*/
func (s *SecurityGroupToVirtualNetworkService) GetSecurityGroupsToVnToVLAN(getSecurityGroupsToVnToVlanQueryParams *GetSecurityGroupsToVnToVLANQueryParams) (*ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLAN, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgtvnvlan"

	queryString, _ := query.Values(getSecurityGroupsToVnToVlanQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLAN{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSecurityGroupsToVnToVlan")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLAN)
	return result, response, err

}

//GetVersion Get security group to virtual network version information
/* This API helps to retrieve the version information related to the security group to virtual network.

 */
func (s *SecurityGroupToVirtualNetworkService) GetVersion() (*ResponseSecurityGroupToVirtualNetworkGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgtvnvlan/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityGroupToVirtualNetworkGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSecurityGroupToVirtualNetworkGetVersion)
	return result, response, err

}

//MonitorBulkStatusSecurityGroupsToVnToVLAN Monitor bulk request
/* This API allows the client to monitor the bulk request.

@param bulkid bulkid path parameter.
*/
func (s *SecurityGroupToVirtualNetworkService) MonitorBulkStatusSecurityGroupsToVnToVLAN(bulkid string) (*ResponseSecurityGroupToVirtualNetworkMonitorBulkStatusSecurityGroupsToVnToVLAN, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgtvnvlan/bulk/{bulkid}"
	path = strings.Replace(path, "{bulkid}", fmt.Sprintf("%v", bulkid), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityGroupToVirtualNetworkMonitorBulkStatusSecurityGroupsToVnToVLAN{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MonitorBulkStatusSecurityGroupsToVnToVlan")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSecurityGroupToVirtualNetworkMonitorBulkStatusSecurityGroupsToVnToVLAN)
	return result, response, err

}

//CreateSecurityGroupsToVnToVLAN Create security group to virtual network
/* This API creates a security group to virtual network.

 */
func (s *SecurityGroupToVirtualNetworkService) CreateSecurityGroupsToVnToVLAN(requestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVlan *RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLAN) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgtvnvlan"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVlan).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateSecurityGroupsToVnToVlan")
	}

	return response, err

}

//UpdateSecurityGroupsToVnToVLANByID Update security group to virtual network
/* This API allows the client to update a security group to virtual network.

@param id id path parameter.
*/
func (s *SecurityGroupToVirtualNetworkService) UpdateSecurityGroupsToVnToVLANByID(id string, requestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVlanById *RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByID) (*ResponseSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgtvnvlan/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVlanById).
		SetResult(&ResponseSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSecurityGroupsToVnToVlanById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByID)
	return result, response, err

}

//BulkRequestForSecurityGroupsToVnToVLAN Submit bulk request
/* This API allows the client to submit the bulk request.

 */
func (s *SecurityGroupToVirtualNetworkService) BulkRequestForSecurityGroupsToVnToVLAN(requestSecurityGroupToVirtualNetworkBulkRequestForSecurityGroupsToVnToVlan *RequestSecurityGroupToVirtualNetworkBulkRequestForSecurityGroupsToVnToVLAN) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgtvnvlan/bulk/submit"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSecurityGroupToVirtualNetworkBulkRequestForSecurityGroupsToVnToVlan).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation BulkRequestForSecurityGroupsToVnToVlan")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteSecurityGroupsToVnToVLANByID Delete security group to virtual network
/* This API deletes a security group ACL to virtual network.

@param id id path parameter.
*/
func (s *SecurityGroupToVirtualNetworkService) DeleteSecurityGroupsToVnToVLANByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sgtvnvlan/{id}"
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
		return response, fmt.Errorf("error with operation DeleteSecurityGroupsToVnToVlanById")
	}

	getCSFRToken(response.Header())

	return response, err

}
