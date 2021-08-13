package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type EndpointService service

type GetEndpointsQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseEndpointGetRejectedEndpoints struct {
	OperationResult ResponseEndpointGetRejectedEndpointsOperationResult `json:"OperationResult,omitempty"` //
}

type ResponseEndpointGetRejectedEndpointsOperationResult struct {
	ResultValue []ResponseEndpointGetRejectedEndpointsOperationResultResultValue `json:"resultValue,omitempty"` //
}

type ResponseEndpointGetRejectedEndpointsOperationResultResultValue struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type ResponseEndpointGetEndpointByName struct {
	ERSEndPoint ResponseEndpointGetEndpointByNameERSEndPoint `json:"ERSEndPoint,omitempty"` //
}

type ResponseEndpointGetEndpointByNameERSEndPoint struct {
	ID                      string                                                    `json:"id,omitempty"`                      //
	Name                    string                                                    `json:"name,omitempty"`                    //
	Description             string                                                    `json:"description,omitempty"`             //
	Mac                     string                                                    `json:"mac,omitempty"`                     //
	ProfileID               string                                                    `json:"profileId,omitempty"`               //
	StaticProfileAssignment bool                                                      `json:"staticProfileAssignment,omitempty"` //
	GroupID                 string                                                    `json:"groupId,omitempty"`                 //
	StaticGroupAssignment   bool                                                      `json:"staticGroupAssignment,omitempty"`   //
	PortalUser              string                                                    `json:"portalUser,omitempty"`              //
	IDentityStore           string                                                    `json:"identityStore,omitempty"`           //
	IDentityStoreID         string                                                    `json:"identityStoreId,omitempty"`         //
	MdmAttributes           ResponseEndpointGetEndpointByNameERSEndPointMdmAttributes `json:"mdmAttributes,omitempty"`           //
	CustomAttributes        map[string]interface{}                                    `json:"customAttributes,omitempty"`        //
	Link                    ResponseEndpointGetEndpointByNameERSEndPointLink          `json:"link,omitempty"`                    //
}

type ResponseEndpointGetEndpointByNameERSEndPointMdmAttributes struct {
	MdmServerName       string `json:"mdmServerName,omitempty"`       //
	MdmReachable        bool   `json:"mdmReachable,omitempty"`        //
	MdmEnrolled         bool   `json:"mdmEnrolled,omitempty"`         //
	MdmComplianceStatus bool   `json:"mdmComplianceStatus,omitempty"` //
	MdmOS               string `json:"mdmOS,omitempty"`               //
	MdmManufacturer     string `json:"mdmManufacturer,omitempty"`     //
	MdmModel            string `json:"mdmModel,omitempty"`            //
	MdmSerial           string `json:"mdmSerial,omitempty"`           //
	MdmEncrypted        bool   `json:"mdmEncrypted,omitempty"`        //
	MdmPinlock          bool   `json:"mdmPinlock,omitempty"`          //
	MdmJailBroken       bool   `json:"mdmJailBroken,omitempty"`       //
	MdmIMEI             string `json:"mdmIMEI,omitempty"`             //
	MdmPhoneNumber      string `json:"mdmPhoneNumber,omitempty"`      //
}

type ResponseEndpointGetEndpointByNameERSEndPointLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseEndpointGetEndpointByID struct {
	ERSEndPoint ResponseEndpointGetEndpointByIDERSEndPoint `json:"ERSEndPoint,omitempty"` //
}

type ResponseEndpointGetEndpointByIDERSEndPoint struct {
	ID                      string                                                  `json:"id,omitempty"`                      //
	Name                    string                                                  `json:"name,omitempty"`                    //
	Description             string                                                  `json:"description,omitempty"`             //
	Mac                     string                                                  `json:"mac,omitempty"`                     //
	ProfileID               string                                                  `json:"profileId,omitempty"`               //
	StaticProfileAssignment bool                                                    `json:"staticProfileAssignment,omitempty"` //
	GroupID                 string                                                  `json:"groupId,omitempty"`                 //
	StaticGroupAssignment   bool                                                    `json:"staticGroupAssignment,omitempty"`   //
	PortalUser              string                                                  `json:"portalUser,omitempty"`              //
	IDentityStore           string                                                  `json:"identityStore,omitempty"`           //
	IDentityStoreID         string                                                  `json:"identityStoreId,omitempty"`         //
	MdmAttributes           ResponseEndpointGetEndpointByIDERSEndPointMdmAttributes `json:"mdmAttributes,omitempty"`           //
	CustomAttributes        map[string]interface{}                                  `json:"customAttributes,omitempty"`        //
	Link                    ResponseEndpointGetEndpointByIDERSEndPointLink          `json:"link,omitempty"`                    //
}

type ResponseEndpointGetEndpointByIDERSEndPointMdmAttributes struct {
	MdmServerName       string `json:"mdmServerName,omitempty"`       //
	MdmReachable        bool   `json:"mdmReachable,omitempty"`        //
	MdmEnrolled         bool   `json:"mdmEnrolled,omitempty"`         //
	MdmComplianceStatus bool   `json:"mdmComplianceStatus,omitempty"` //
	MdmOS               string `json:"mdmOS,omitempty"`               //
	MdmManufacturer     string `json:"mdmManufacturer,omitempty"`     //
	MdmModel            string `json:"mdmModel,omitempty"`            //
	MdmSerial           string `json:"mdmSerial,omitempty"`           //
	MdmEncrypted        bool   `json:"mdmEncrypted,omitempty"`        //
	MdmPinlock          bool   `json:"mdmPinlock,omitempty"`          //
	MdmJailBroken       bool   `json:"mdmJailBroken,omitempty"`       //
	MdmIMEI             string `json:"mdmIMEI,omitempty"`             //
	MdmPhoneNumber      string `json:"mdmPhoneNumber,omitempty"`      //
}

type ResponseEndpointGetEndpointByIDERSEndPointLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseEndpointUpdateEndpointByID struct {
	UpdatedFieldsList ResponseEndpointUpdateEndpointByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseEndpointUpdateEndpointByIDUpdatedFieldsList struct {
	UpdatedField ResponseEndpointUpdateEndpointByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                          `json:"field,omitempty"`        //
	OldValue     string                                                          `json:"oldValue,omitempty"`     //
	NewValue     string                                                          `json:"newValue,omitempty"`     //
}

type ResponseEndpointUpdateEndpointByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseEndpointGetEndpoints struct {
	SearchResult ResponseEndpointGetEndpointsSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseEndpointGetEndpointsSearchResult struct {
	Total        int                                                  `json:"total,omitempty"`        //
	Resources    []ResponseEndpointGetEndpointsSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponseEndpointGetEndpointsSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponseEndpointGetEndpointsSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseEndpointGetEndpointsSearchResultResources struct {
	ID          string                                                `json:"id,omitempty"`          //
	Name        string                                                `json:"name,omitempty"`        //
	Description string                                                `json:"description,omitempty"` //
	Link        ResponseEndpointGetEndpointsSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseEndpointGetEndpointsSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseEndpointGetEndpointsSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseEndpointGetEndpointsSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseEndpointGetVersion struct {
	VersionInfo ResponseEndpointGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseEndpointGetVersionVersionInfo struct {
	CurrentServerVersion string                                    `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                    `json:"supportedVersions,omitempty"`    //
	Link                 ResponseEndpointGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseEndpointGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseEndpointMonitorBulkStatusEndpoint struct {
	BulkStatus ResponseEndpointMonitorBulkStatusEndpointBulkStatus `json:"BulkStatus,omitempty"` //
}

type ResponseEndpointMonitorBulkStatusEndpointBulkStatus struct {
	BulkID          string                                                               `json:"bulkId,omitempty"`          //
	MediaType       string                                                               `json:"mediaType,omitempty"`       //
	ExecutionStatus string                                                               `json:"executionStatus,omitempty"` //
	OperationType   string                                                               `json:"operationType,omitempty"`   //
	StartTime       string                                                               `json:"startTime,omitempty"`       //
	ResourcesCount  int                                                                  `json:"resourcesCount,omitempty"`  //
	SuccessCount    int                                                                  `json:"successCount,omitempty"`    //
	FailCount       int                                                                  `json:"failCount,omitempty"`       //
	ResourcesStatus []ResponseEndpointMonitorBulkStatusEndpointBulkStatusResourcesStatus `json:"resourcesStatus,omitempty"` //
}

type ResponseEndpointMonitorBulkStatusEndpointBulkStatusResourcesStatus struct {
	ID                      string `json:"id,omitempty"`                      //
	Name                    string `json:"name,omitempty"`                    //
	Description             string `json:"description,omitempty"`             //
	ResourceExecutionStatus string `json:"resourceExecutionStatus,omitempty"` //
	Status                  string `json:"status,omitempty"`                  //
}

type RequestEndpointUpdateEndpointByID struct {
	ERSEndPoint *RequestEndpointUpdateEndpointByIDERSEndPoint `json:"ERSEndPoint,omitempty"` //
}

type RequestEndpointUpdateEndpointByIDERSEndPoint struct {
	ID                      string                                                     `json:"id,omitempty"`                      //
	Name                    string                                                     `json:"name,omitempty"`                    //
	Description             string                                                     `json:"description,omitempty"`             //
	Mac                     string                                                     `json:"mac,omitempty"`                     //
	ProfileID               string                                                     `json:"profileId,omitempty"`               //
	StaticProfileAssignment *bool                                                      `json:"staticProfileAssignment,omitempty"` //
	GroupID                 string                                                     `json:"groupId,omitempty"`                 //
	StaticGroupAssignment   *bool                                                      `json:"staticGroupAssignment,omitempty"`   //
	PortalUser              string                                                     `json:"portalUser,omitempty"`              //
	IDentityStore           string                                                     `json:"identityStore,omitempty"`           //
	IDentityStoreID         string                                                     `json:"identityStoreId,omitempty"`         //
	MdmAttributes           *RequestEndpointUpdateEndpointByIDERSEndPointMdmAttributes `json:"mdmAttributes,omitempty"`           //
	CustomAttributes        *map[string]interface{}                                    `json:"customAttributes,omitempty"`        //
}

type RequestEndpointUpdateEndpointByIDERSEndPointMdmAttributes struct {
	MdmServerName       string `json:"mdmServerName,omitempty"`       //
	MdmReachable        *bool  `json:"mdmReachable,omitempty"`        //
	MdmEnrolled         *bool  `json:"mdmEnrolled,omitempty"`         //
	MdmComplianceStatus *bool  `json:"mdmComplianceStatus,omitempty"` //
	MdmOS               string `json:"mdmOS,omitempty"`               //
	MdmManufacturer     string `json:"mdmManufacturer,omitempty"`     //
	MdmModel            string `json:"mdmModel,omitempty"`            //
	MdmSerial           string `json:"mdmSerial,omitempty"`           //
	MdmEncrypted        *bool  `json:"mdmEncrypted,omitempty"`        //
	MdmPinlock          *bool  `json:"mdmPinlock,omitempty"`          //
	MdmJailBroken       *bool  `json:"mdmJailBroken,omitempty"`       //
	MdmIMEI             string `json:"mdmIMEI,omitempty"`             //
	MdmPhoneNumber      string `json:"mdmPhoneNumber,omitempty"`      //
}

type RequestEndpointRegisterEndpoint struct {
	ERSEndPoint *RequestEndpointRegisterEndpointERSEndPoint `json:"ERSEndPoint,omitempty"` //
}

type RequestEndpointRegisterEndpointERSEndPoint struct {
	ID                      string                                                   `json:"id,omitempty"`                      //
	Name                    string                                                   `json:"name,omitempty"`                    //
	Description             string                                                   `json:"description,omitempty"`             //
	Mac                     string                                                   `json:"mac,omitempty"`                     //
	ProfileID               string                                                   `json:"profileId,omitempty"`               //
	StaticProfileAssignment *bool                                                    `json:"staticProfileAssignment,omitempty"` //
	GroupID                 string                                                   `json:"groupId,omitempty"`                 //
	StaticGroupAssignment   *bool                                                    `json:"staticGroupAssignment,omitempty"`   //
	PortalUser              string                                                   `json:"portalUser,omitempty"`              //
	IDentityStore           string                                                   `json:"identityStore,omitempty"`           //
	IDentityStoreID         string                                                   `json:"identityStoreId,omitempty"`         //
	MdmAttributes           *RequestEndpointRegisterEndpointERSEndPointMdmAttributes `json:"mdmAttributes,omitempty"`           //
	CustomAttributes        *map[string]interface{}                                  `json:"customAttributes,omitempty"`        //
}

type RequestEndpointRegisterEndpointERSEndPointMdmAttributes struct {
	MdmServerName       string `json:"mdmServerName,omitempty"`       //
	MdmReachable        *bool  `json:"mdmReachable,omitempty"`        //
	MdmEnrolled         *bool  `json:"mdmEnrolled,omitempty"`         //
	MdmComplianceStatus *bool  `json:"mdmComplianceStatus,omitempty"` //
	MdmOS               string `json:"mdmOS,omitempty"`               //
	MdmManufacturer     string `json:"mdmManufacturer,omitempty"`     //
	MdmModel            string `json:"mdmModel,omitempty"`            //
	MdmSerial           string `json:"mdmSerial,omitempty"`           //
	MdmEncrypted        *bool  `json:"mdmEncrypted,omitempty"`        //
	MdmPinlock          *bool  `json:"mdmPinlock,omitempty"`          //
	MdmJailBroken       *bool  `json:"mdmJailBroken,omitempty"`       //
	MdmIMEI             string `json:"mdmIMEI,omitempty"`             //
	MdmPhoneNumber      string `json:"mdmPhoneNumber,omitempty"`      //
}

type RequestEndpointCreateEndpoint struct {
	ERSEndPoint *RequestEndpointCreateEndpointERSEndPoint `json:"ERSEndPoint,omitempty"` //
}

type RequestEndpointCreateEndpointERSEndPoint struct {
	Name                    string                                                 `json:"name,omitempty"`                    //
	Description             string                                                 `json:"description,omitempty"`             //
	Mac                     string                                                 `json:"mac,omitempty"`                     //
	ProfileID               string                                                 `json:"profileId,omitempty"`               //
	StaticProfileAssignment *bool                                                  `json:"staticProfileAssignment,omitempty"` //
	GroupID                 string                                                 `json:"groupId,omitempty"`                 //
	StaticGroupAssignment   *bool                                                  `json:"staticGroupAssignment,omitempty"`   //
	PortalUser              string                                                 `json:"portalUser,omitempty"`              //
	IDentityStore           string                                                 `json:"identityStore,omitempty"`           //
	IDentityStoreID         string                                                 `json:"identityStoreId,omitempty"`         //
	MdmAttributes           *RequestEndpointCreateEndpointERSEndPointMdmAttributes `json:"mdmAttributes,omitempty"`           //
	CustomAttributes        *map[string]interface{}                                `json:"customAttributes,omitempty"`        //
}

type RequestEndpointCreateEndpointERSEndPointMdmAttributes struct {
	MdmServerName       string `json:"mdmServerName,omitempty"`       //
	MdmReachable        *bool  `json:"mdmReachable,omitempty"`        //
	MdmEnrolled         *bool  `json:"mdmEnrolled,omitempty"`         //
	MdmComplianceStatus *bool  `json:"mdmComplianceStatus,omitempty"` //
	MdmOS               string `json:"mdmOS,omitempty"`               //
	MdmManufacturer     string `json:"mdmManufacturer,omitempty"`     //
	MdmModel            string `json:"mdmModel,omitempty"`            //
	MdmSerial           string `json:"mdmSerial,omitempty"`           //
	MdmEncrypted        *bool  `json:"mdmEncrypted,omitempty"`        //
	MdmPinlock          *bool  `json:"mdmPinlock,omitempty"`          //
	MdmJailBroken       *bool  `json:"mdmJailBroken,omitempty"`       //
	MdmIMEI             string `json:"mdmIMEI,omitempty"`             //
	MdmPhoneNumber      string `json:"mdmPhoneNumber,omitempty"`      //
}

type RequestEndpointBulkRequestForEndpoint struct {
	EndpointBulkRequest *RequestEndpointBulkRequestForEndpointEndpointBulkRequest `json:"EndpointBulkRequest,omitempty"` //
}

type RequestEndpointBulkRequestForEndpointEndpointBulkRequest struct {
	OperationType     string `json:"operationType,omitempty"`     //
	ResourceMediaType string `json:"resourceMediaType,omitempty"` //
}

//GetRejectedEndpoints Get rejected enpoints
/* This API allows the client to get the rejected endpoints.

 */
func (s *EndpointService) GetRejectedEndpoints() (*ResponseEndpointGetRejectedEndpoints, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpoint/getrejectedendpoints"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEndpointGetRejectedEndpoints{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetRejectedEndpoints")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointGetRejectedEndpoints)
	return result, response, err

}

//GetEndpointByName Get endpoint by name
/* This API allows the client to get an endpoint by name.

@param name name path parameter.
*/
func (s *EndpointService) GetEndpointByName(name string) (*ResponseEndpointGetEndpointByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpoint/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEndpointGetEndpointByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetEndpointByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointGetEndpointByName)
	return result, response, err

}

//GetEndpointByID Get endpoint by ID
/* This API allows the client to get an endpoint by ID.

@param id id path parameter.
*/
func (s *EndpointService) GetEndpointByID(id string) (*ResponseEndpointGetEndpointByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpoint/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEndpointGetEndpointByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetEndpointById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointGetEndpointByID)
	return result, response, err

}

//GetEndpoints Get all endpoints
/* This API allows the client to get all the endpoints.

Filter:
Filters can be used to filter out Endpoints based on a set of attributes. This API currently provides the following filters:
[logicalProfileName, portalUser, staticProfileAssignment, profileId, profile, groupId, staticGroupAssignment, mac]

Example 1:

The
logicalProfileName
 filter can be used to get Enpoints that belong  to a specific Logical Profile. The supported operator for logicalProfileNamefilter is EQ (equal to). The syntax to invoke the API with this filter:

/ers/config/endpoint?filter={filter name}.{operator}.{logical profile name}

Example:

https://{ise-ip}:9060/ers/config/endpoint?filter=logicalProfileName.EQ.LP_Apple

Example 2:

To search resources by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:
[name, description]

@param getEndpointsQueryParams Filtering parameter
*/
func (s *EndpointService) GetEndpoints(getEndpointsQueryParams *GetEndpointsQueryParams) (*ResponseEndpointGetEndpoints, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpoint"

	queryString, _ := query.Values(getEndpointsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEndpointGetEndpoints{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetEndpoints")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointGetEndpoints)
	return result, response, err

}

//GetVersion Get endpoint version information
/* This API helps to retrieve the version information related to the endpoint.

 */
func (s *EndpointService) GetVersion() (*ResponseEndpointGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpoint/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEndpointGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointGetVersion)
	return result, response, err

}

//MonitorBulkStatusEndpoint Monitor bulk request
/* This API allows the client to monitor the bulk request.

@param bulkid bulkid path parameter.
*/
func (s *EndpointService) MonitorBulkStatusEndpoint(bulkid string) (*ResponseEndpointMonitorBulkStatusEndpoint, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpoint/bulk/{bulkid}"
	path = strings.Replace(path, "{bulkid}", fmt.Sprintf("%v", bulkid), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEndpointMonitorBulkStatusEndpoint{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MonitorBulkStatusEndpoint")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointMonitorBulkStatusEndpoint)
	return result, response, err

}

//CreateEndpoint Create endpoint
/* This API creates an endpoint.

 */
func (s *EndpointService) CreateEndpoint(requestEndpointCreateEndpoint *RequestEndpointCreateEndpoint) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpoint"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEndpointCreateEndpoint).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateEndpoint")
	}

	return response, err

}

//ReleaseRejectedEndpoint Release rejected endpoint
/* This API allows the client to release a rejected endpoint.

@param id id path parameter.
*/
func (s *EndpointService) ReleaseRejectedEndpoint(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpoint/{id}/releaserejectedendpoint"
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
		return response, fmt.Errorf("error with operation ReleaseRejectedEndpoint")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeregisterEndpoint De-register endpoint
/* This API allows the client to de-register an endpoint.

@param id id path parameter.
*/
func (s *EndpointService) DeregisterEndpoint(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpoint/{id}/deregister"
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
		return response, fmt.Errorf("error with operation DeregisterEndpoint")
	}

	getCSFRToken(response.Header())

	return response, err

}

//UpdateEndpointByID Update endpoint
/* This API allows the client to update an endpoint.

@param id id path parameter.
*/
func (s *EndpointService) UpdateEndpointByID(id string, requestEndpointUpdateEndpointById *RequestEndpointUpdateEndpointByID) (*ResponseEndpointUpdateEndpointByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpoint/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEndpointUpdateEndpointById).
		SetResult(&ResponseEndpointUpdateEndpointByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateEndpointById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointUpdateEndpointByID)
	return result, response, err

}

//RegisterEndpoint Register endpoint
/* This API allows the client to register an endpoint.

 */
func (s *EndpointService) RegisterEndpoint(requestEndpointRegisterEndpoint *RequestEndpointRegisterEndpoint) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpoint/register"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEndpointRegisterEndpoint).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation RegisterEndpoint")
	}

	getCSFRToken(response.Header())

	return response, err

}

//BulkRequestForEndpoint Submit bulk request
/* This API allows the client to submit the bulk request.

 */
func (s *EndpointService) BulkRequestForEndpoint(requestEndpointBulkRequestForEndpoint *RequestEndpointBulkRequestForEndpoint) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpoint/bulk/submit"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEndpointBulkRequestForEndpoint).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation BulkRequestForEndpoint")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteEndpointByID Delete endpoint
/* This API deletes an endpoint.

@param id id path parameter.
*/
func (s *EndpointService) DeleteEndpointByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpoint/{id}"
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
		return response, fmt.Errorf("error with operation DeleteEndpointById")
	}

	getCSFRToken(response.Header())

	return response, err

}
