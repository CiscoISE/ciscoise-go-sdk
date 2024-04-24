package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type EndpointsService service

type List1QueryParams struct {
	Page       int    `url:"page,omitempty"`       //Page number
	Size       int    `url:"size,omitempty"`       //Number of objects returned per page
	Sort       string `url:"sort,omitempty"`       //sort type - asc or desc
	SortBy     string `url:"sortBy,omitempty"`     //sort column by which objects needs to be sorted
	Filter     string `url:"filter,omitempty"`     //<div> <style type='text/css' scoped> .apiServiceTable td, .apiServiceTable th { padding: 5px 10px !important; text-align: left; } </style> <span> <b>Simple filtering</b> should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the <i>'filterType=or'</i> query string parameter. Each resource Data model description should specify if an attribute is a filtered field. </span> <br /> <table class='apiServiceTable'> <thead> <tr> <th>OPERATOR</th> <th>DESCRIPTION</th> </tr> </thead> <tbody> <tr> <td>EQ</td> <td>Equals</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> </tr> <tr> <td>GT</td> <td>Greater Than</td> </tr> <tr> <td>LT</td> <td>Less Then</td> </tr> <tr> <td>STARTSW</td> <td>Starts With</td> </tr> <tr> <td>NSTARTSW</td> <td>Not Starts With</td> </tr> <tr> <td>ENDSW</td> <td>Ends With</td> </tr> <tr> <td>NENDSW</td> <td>Not Ends With</td> </tr> <tr> <td>CONTAINS</td> <td>Contains</td> </tr> <tr> <td>NCONTAINS</td> <td>Not Contains</td> </tr> </tbody> </table> </div>
	FilterType string `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseEndpointsList1 []ResponseItemEndpointsList1 // Array of ResponseEndpointsList_1

type ResponseItemEndpointsList1 struct {
	ConnectedLinks          *ResponseItemEndpointsList1ConnectedLinks   `json:"connectedLinks,omitempty"`          //
	CustomAttributes        *ResponseItemEndpointsList1CustomAttributes `json:"customAttributes,omitempty"`        //
	Description             string                                      `json:"description,omitempty"`             //
	DeviceType              string                                      `json:"deviceType,omitempty"`              //
	GroupID                 string                                      `json:"groupId,omitempty"`                 //
	HardwareRevision        string                                      `json:"hardwareRevision,omitempty"`        //
	ID                      string                                      `json:"id,omitempty"`                      //
	IDentityStore           string                                      `json:"identityStore,omitempty"`           //
	IDentityStoreID         string                                      `json:"identityStoreId,omitempty"`         //
	IPAddress               string                                      `json:"ipAddress,omitempty"`               //
	Mac                     string                                      `json:"mac,omitempty"`                     //
	MdmAttributes           *ResponseItemEndpointsList1MdmAttributes    `json:"mdmAttributes,omitempty"`           //
	Name                    string                                      `json:"name,omitempty"`                    //
	PortalUser              string                                      `json:"portalUser,omitempty"`              //
	ProductID               string                                      `json:"productId,omitempty"`               //
	ProfileID               string                                      `json:"profileId,omitempty"`               //
	Protocol                string                                      `json:"protocol,omitempty"`                //
	SerialNumber            string                                      `json:"serialNumber,omitempty"`            //
	SoftwareRevision        string                                      `json:"softwareRevision,omitempty"`        //
	StaticGroupAssignment   *bool                                       `json:"staticGroupAssignment,omitempty"`   //
	StaticProfileAssignment *bool                                       `json:"staticProfileAssignment,omitempty"` //
	Vendor                  string                                      `json:"vendor,omitempty"`                  //
}

type ResponseItemEndpointsList1ConnectedLinks interface{}

type ResponseItemEndpointsList1CustomAttributes map[string]interface{}

type ResponseItemEndpointsList1MdmAttributes interface{}

// # Review unknown case

type ResponseEndpointsUpdateBulkEndPoints struct {
	ID string `json:"id,omitempty"` //
}

type ResponseEndpointsCreateBulkEndPoints struct {
	ID string `json:"id,omitempty"` //
}

type ResponseEndpointsDeleteBulkEndPoints struct {
	ID string `json:"id,omitempty"` //
}

type ResponseEndpointsGetDeviceTypeSummary []ResponseItemEndpointsGetDeviceTypeSummary // Array of ResponseEndpointsGetDeviceTypeSummary

type ResponseItemEndpointsGetDeviceTypeSummary struct {
	DeviceType string `json:"deviceType,omitempty"` //
	Total      string `json:"total,omitempty"`      //
}

type ResponseEndpointsGet1 struct {
	ConnectedLinks          *ResponseEndpointsGet1ConnectedLinks   `json:"connectedLinks,omitempty"`          //
	CustomAttributes        *ResponseEndpointsGet1CustomAttributes `json:"customAttributes,omitempty"`        //
	Description             string                                 `json:"description,omitempty"`             //
	DeviceType              string                                 `json:"deviceType,omitempty"`              //
	GroupID                 string                                 `json:"groupId,omitempty"`                 //
	HardwareRevision        string                                 `json:"hardwareRevision,omitempty"`        //
	ID                      string                                 `json:"id,omitempty"`                      //
	IDentityStore           string                                 `json:"identityStore,omitempty"`           //
	IDentityStoreID         string                                 `json:"identityStoreId,omitempty"`         //
	IPAddress               string                                 `json:"ipAddress,omitempty"`               //
	Mac                     string                                 `json:"mac,omitempty"`                     //
	MdmAttributes           *ResponseEndpointsGet1MdmAttributes    `json:"mdmAttributes,omitempty"`           //
	Name                    string                                 `json:"name,omitempty"`                    //
	PortalUser              string                                 `json:"portalUser,omitempty"`              //
	ProductID               string                                 `json:"productId,omitempty"`               //
	ProfileID               string                                 `json:"profileId,omitempty"`               //
	Protocol                string                                 `json:"protocol,omitempty"`                //
	SerialNumber            string                                 `json:"serialNumber,omitempty"`            //
	SoftwareRevision        string                                 `json:"softwareRevision,omitempty"`        //
	StaticGroupAssignment   *bool                                  `json:"staticGroupAssignment,omitempty"`   //
	StaticProfileAssignment *bool                                  `json:"staticProfileAssignment,omitempty"` //
	Vendor                  string                                 `json:"vendor,omitempty"`                  //
}

type ResponseEndpointsGet1ConnectedLinks interface{}

type ResponseEndpointsGet1CustomAttributes map[string]interface{}

type ResponseEndpointsGet1MdmAttributes interface{}

type ResponseEndpointsUpdateEndpoint struct {
	ConnectedLinks          *ResponseEndpointsUpdateEndpointConnectedLinks   `json:"connectedLinks,omitempty"`          //
	CustomAttributes        *ResponseEndpointsUpdateEndpointCustomAttributes `json:"customAttributes,omitempty"`        //
	Description             string                                           `json:"description,omitempty"`             //
	DeviceType              string                                           `json:"deviceType,omitempty"`              //
	GroupID                 string                                           `json:"groupId,omitempty"`                 //
	HardwareRevision        string                                           `json:"hardwareRevision,omitempty"`        //
	ID                      string                                           `json:"id,omitempty"`                      //
	IDentityStore           string                                           `json:"identityStore,omitempty"`           //
	IDentityStoreID         string                                           `json:"identityStoreId,omitempty"`         //
	IPAddress               string                                           `json:"ipAddress,omitempty"`               //
	Mac                     string                                           `json:"mac,omitempty"`                     //
	MdmAttributes           *ResponseEndpointsUpdateEndpointMdmAttributes    `json:"mdmAttributes,omitempty"`           //
	Name                    string                                           `json:"name,omitempty"`                    //
	PortalUser              string                                           `json:"portalUser,omitempty"`              //
	ProductID               string                                           `json:"productId,omitempty"`               //
	ProfileID               string                                           `json:"profileId,omitempty"`               //
	Protocol                string                                           `json:"protocol,omitempty"`                //
	SerialNumber            string                                           `json:"serialNumber,omitempty"`            //
	SoftwareRevision        string                                           `json:"softwareRevision,omitempty"`        //
	StaticGroupAssignment   *bool                                            `json:"staticGroupAssignment,omitempty"`   //
	StaticProfileAssignment *bool                                            `json:"staticProfileAssignment,omitempty"` //
	Vendor                  string                                           `json:"vendor,omitempty"`                  //
}

type ResponseEndpointsUpdateEndpointConnectedLinks interface{}

type ResponseEndpointsUpdateEndpointCustomAttributes map[string]interface{}

type ResponseEndpointsUpdateEndpointMdmAttributes interface{}

// # Review unknown case

type ResponseEndpointsCreateEndPointTask struct {
	ID string `json:"id,omitempty"` //
}

type RequestEndpointsCreateEndPoint struct {
	ConnectedLinks          *RequestEndpointsCreateEndPointConnectedLinks   `json:"connectedLinks,omitempty"`          //
	CustomAttributes        *RequestEndpointsCreateEndPointCustomAttributes `json:"customAttributes,omitempty"`        //
	Description             string                                          `json:"description,omitempty"`             //
	DeviceType              string                                          `json:"deviceType,omitempty"`              //
	GroupID                 string                                          `json:"groupId,omitempty"`                 //
	HardwareRevision        string                                          `json:"hardwareRevision,omitempty"`        //
	ID                      string                                          `json:"id,omitempty"`                      //
	IDentityStore           string                                          `json:"identityStore,omitempty"`           //
	IDentityStoreID         string                                          `json:"identityStoreId,omitempty"`         //
	IPAddress               string                                          `json:"ipAddress,omitempty"`               //
	Mac                     string                                          `json:"mac,omitempty"`                     //
	MdmAttributes           *RequestEndpointsCreateEndPointMdmAttributes    `json:"mdmAttributes,omitempty"`           //
	Name                    string                                          `json:"name,omitempty"`                    //
	PortalUser              string                                          `json:"portalUser,omitempty"`              //
	ProductID               string                                          `json:"productId,omitempty"`               //
	ProfileID               string                                          `json:"profileId,omitempty"`               //
	Protocol                string                                          `json:"protocol,omitempty"`                //
	SerialNumber            string                                          `json:"serialNumber,omitempty"`            //
	SoftwareRevision        string                                          `json:"softwareRevision,omitempty"`        //
	StaticGroupAssignment   *bool                                           `json:"staticGroupAssignment,omitempty"`   //
	StaticProfileAssignment *bool                                           `json:"staticProfileAssignment,omitempty"` //
	Vendor                  string                                          `json:"vendor,omitempty"`                  //
}

type RequestEndpointsCreateEndPointConnectedLinks interface{}

type RequestEndpointsCreateEndPointCustomAttributes map[string]interface{}

type RequestEndpointsCreateEndPointMdmAttributes interface{}

type RequestEndpointsUpdateBulkEndPoints []RequestItemEndpointsUpdateBulkEndPoints // Array of RequestEndpointsUpdateBulkEndPoints

type RequestItemEndpointsUpdateBulkEndPoints struct {
	ConnectedLinks          *RequestItemEndpointsUpdateBulkEndPointsConnectedLinks   `json:"connectedLinks,omitempty"`          //
	CustomAttributes        *RequestItemEndpointsUpdateBulkEndPointsCustomAttributes `json:"customAttributes,omitempty"`        //
	Description             string                                                   `json:"description,omitempty"`             //
	DeviceType              string                                                   `json:"deviceType,omitempty"`              //
	GroupID                 string                                                   `json:"groupId,omitempty"`                 //
	HardwareRevision        string                                                   `json:"hardwareRevision,omitempty"`        //
	ID                      string                                                   `json:"id,omitempty"`                      //
	IDentityStore           string                                                   `json:"identityStore,omitempty"`           //
	IDentityStoreID         string                                                   `json:"identityStoreId,omitempty"`         //
	IPAddress               string                                                   `json:"ipAddress,omitempty"`               //
	Mac                     string                                                   `json:"mac,omitempty"`                     //
	MdmAttributes           *RequestItemEndpointsUpdateBulkEndPointsMdmAttributes    `json:"mdmAttributes,omitempty"`           //
	Name                    string                                                   `json:"name,omitempty"`                    //
	PortalUser              string                                                   `json:"portalUser,omitempty"`              //
	ProductID               string                                                   `json:"productId,omitempty"`               //
	ProfileID               string                                                   `json:"profileId,omitempty"`               //
	Protocol                string                                                   `json:"protocol,omitempty"`                //
	SerialNumber            string                                                   `json:"serialNumber,omitempty"`            //
	SoftwareRevision        string                                                   `json:"softwareRevision,omitempty"`        //
	StaticGroupAssignment   *bool                                                    `json:"staticGroupAssignment,omitempty"`   //
	StaticProfileAssignment *bool                                                    `json:"staticProfileAssignment,omitempty"` //
	Vendor                  string                                                   `json:"vendor,omitempty"`                  //
}

type RequestItemEndpointsUpdateBulkEndPointsConnectedLinks interface{}

type RequestItemEndpointsUpdateBulkEndPointsCustomAttributes map[string]interface{}

type RequestItemEndpointsUpdateBulkEndPointsMdmAttributes interface{}

type RequestEndpointsCreateBulkEndPoints []RequestItemEndpointsCreateBulkEndPoints // Array of RequestEndpointsCreateBulkEndPoints

type RequestItemEndpointsCreateBulkEndPoints struct {
	ConnectedLinks          *RequestItemEndpointsCreateBulkEndPointsConnectedLinks   `json:"connectedLinks,omitempty"`          //
	CustomAttributes        *RequestItemEndpointsCreateBulkEndPointsCustomAttributes `json:"customAttributes,omitempty"`        //
	Description             string                                                   `json:"description,omitempty"`             //
	DeviceType              string                                                   `json:"deviceType,omitempty"`              //
	GroupID                 string                                                   `json:"groupId,omitempty"`                 //
	HardwareRevision        string                                                   `json:"hardwareRevision,omitempty"`        //
	ID                      string                                                   `json:"id,omitempty"`                      //
	IDentityStore           string                                                   `json:"identityStore,omitempty"`           //
	IDentityStoreID         string                                                   `json:"identityStoreId,omitempty"`         //
	IPAddress               string                                                   `json:"ipAddress,omitempty"`               //
	Mac                     string                                                   `json:"mac,omitempty"`                     //
	MdmAttributes           *RequestItemEndpointsCreateBulkEndPointsMdmAttributes    `json:"mdmAttributes,omitempty"`           //
	Name                    string                                                   `json:"name,omitempty"`                    //
	PortalUser              string                                                   `json:"portalUser,omitempty"`              //
	ProductID               string                                                   `json:"productId,omitempty"`               //
	ProfileID               string                                                   `json:"profileId,omitempty"`               //
	Protocol                string                                                   `json:"protocol,omitempty"`                //
	SerialNumber            string                                                   `json:"serialNumber,omitempty"`            //
	SoftwareRevision        string                                                   `json:"softwareRevision,omitempty"`        //
	StaticGroupAssignment   *bool                                                    `json:"staticGroupAssignment,omitempty"`   //
	StaticProfileAssignment *bool                                                    `json:"staticProfileAssignment,omitempty"` //
	Vendor                  string                                                   `json:"vendor,omitempty"`                  //
}

type RequestItemEndpointsCreateBulkEndPointsConnectedLinks interface{}

type RequestItemEndpointsCreateBulkEndPointsCustomAttributes map[string]interface{}

type RequestItemEndpointsCreateBulkEndPointsMdmAttributes interface{}

type RequestEndpointsUpdateEndpoint struct {
	ConnectedLinks          *RequestEndpointsUpdateEndpointConnectedLinks   `json:"connectedLinks,omitempty"`          //
	CustomAttributes        *RequestEndpointsUpdateEndpointCustomAttributes `json:"customAttributes,omitempty"`        //
	Description             string                                          `json:"description,omitempty"`             //
	DeviceType              string                                          `json:"deviceType,omitempty"`              //
	GroupID                 string                                          `json:"groupId,omitempty"`                 //
	HardwareRevision        string                                          `json:"hardwareRevision,omitempty"`        //
	ID                      string                                          `json:"id,omitempty"`                      //
	IDentityStore           string                                          `json:"identityStore,omitempty"`           //
	IDentityStoreID         string                                          `json:"identityStoreId,omitempty"`         //
	IPAddress               string                                          `json:"ipAddress,omitempty"`               //
	Mac                     string                                          `json:"mac,omitempty"`                     //
	MdmAttributes           *RequestEndpointsUpdateEndpointMdmAttributes    `json:"mdmAttributes,omitempty"`           //
	Name                    string                                          `json:"name,omitempty"`                    //
	PortalUser              string                                          `json:"portalUser,omitempty"`              //
	ProductID               string                                          `json:"productId,omitempty"`               //
	ProfileID               string                                          `json:"profileId,omitempty"`               //
	Protocol                string                                          `json:"protocol,omitempty"`                //
	SerialNumber            string                                          `json:"serialNumber,omitempty"`            //
	SoftwareRevision        string                                          `json:"softwareRevision,omitempty"`        //
	StaticGroupAssignment   *bool                                           `json:"staticGroupAssignment,omitempty"`   //
	StaticProfileAssignment *bool                                           `json:"staticProfileAssignment,omitempty"` //
	Vendor                  string                                          `json:"vendor,omitempty"`                  //
}

type RequestEndpointsUpdateEndpointConnectedLinks interface{}

type RequestEndpointsUpdateEndpointCustomAttributes map[string]interface{}

type RequestEndpointsUpdateEndpointMdmAttributes interface{}

type RequestEndpointsCreateEndPointTask struct {
	ConnectedLinks          *RequestEndpointsCreateEndPointTaskConnectedLinks   `json:"connectedLinks,omitempty"`          //
	CustomAttributes        *RequestEndpointsCreateEndPointTaskCustomAttributes `json:"customAttributes,omitempty"`        //
	Description             string                                              `json:"description,omitempty"`             //
	DeviceType              string                                              `json:"deviceType,omitempty"`              //
	GroupID                 string                                              `json:"groupId,omitempty"`                 //
	HardwareRevision        string                                              `json:"hardwareRevision,omitempty"`        //
	ID                      string                                              `json:"id,omitempty"`                      //
	IDentityStore           string                                              `json:"identityStore,omitempty"`           //
	IDentityStoreID         string                                              `json:"identityStoreId,omitempty"`         //
	IPAddress               string                                              `json:"ipAddress,omitempty"`               //
	Mac                     string                                              `json:"mac,omitempty"`                     //
	MdmAttributes           *RequestEndpointsCreateEndPointTaskMdmAttributes    `json:"mdmAttributes,omitempty"`           //
	Name                    string                                              `json:"name,omitempty"`                    //
	PortalUser              string                                              `json:"portalUser,omitempty"`              //
	ProductID               string                                              `json:"productId,omitempty"`               //
	ProfileID               string                                              `json:"profileId,omitempty"`               //
	Protocol                string                                              `json:"protocol,omitempty"`                //
	SerialNumber            string                                              `json:"serialNumber,omitempty"`            //
	SoftwareRevision        string                                              `json:"softwareRevision,omitempty"`        //
	StaticGroupAssignment   *bool                                               `json:"staticGroupAssignment,omitempty"`   //
	StaticProfileAssignment *bool                                               `json:"staticProfileAssignment,omitempty"` //
	Vendor                  string                                              `json:"vendor,omitempty"`                  //
}

type RequestEndpointsCreateEndPointTaskConnectedLinks interface{}

type RequestEndpointsCreateEndPointTaskCustomAttributes map[string]interface{}

type RequestEndpointsCreateEndPointTaskMdmAttributes interface{}

//List1 Get all endpoints
/* Get all endpoints

@param list_1QueryParams Filtering parameter
*/
func (s *EndpointsService) List1(list_1QueryParams *List1QueryParams) (*ResponseEndpointsList1, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/endpoint"

	queryString, _ := query.Values(list_1QueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEndpointsList1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation List1")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointsList1)
	return result, response, err

}

//GetDeviceTypeSummary Get aggregate of device types
/* Get aggregate of device types

 */
func (s *EndpointsService) GetDeviceTypeSummary() (*ResponseEndpointsGetDeviceTypeSummary, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/endpoint/deviceType/summary"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEndpointsGetDeviceTypeSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceTypeSummary")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointsGetDeviceTypeSummary)
	return result, response, err

}

//Get1 Get endpoint by id or MAC
/* Get endpoint by id or MAC

@param value value path parameter. The id or MAC of the endpoint
*/
func (s *EndpointsService) Get1(value string) (*ResponseEndpointsGet1, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/endpoint/{value}"
	path = strings.Replace(path, "{value}", fmt.Sprintf("%v", value), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEndpointsGet1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation Get1")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointsGet1)
	return result, response, err

}

//CreateEndPoint Create Endpoint
/* Create Endpoint

 */
func (s *EndpointsService) CreateEndPoint(requestEndpointsCreateEndPoint *RequestEndpointsCreateEndPoint) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/endpoint"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEndpointsCreateEndPoint).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateEndPoint")
	}

	return response, err

}

//CreateBulkEndPoints Create Endpoint in bulk
/* Create Endpoint in bulk

 */
func (s *EndpointsService) CreateBulkEndPoints(requestEndpointsCreateBulkEndPoints *RequestEndpointsCreateBulkEndPoints) (*ResponseEndpointsCreateBulkEndPoints, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/endpoint/bulk"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEndpointsCreateBulkEndPoints).
		SetResult(&ResponseEndpointsCreateBulkEndPoints{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseEndpointsCreateBulkEndPoints{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateBulkEndPoints")
	}

	result := response.Result().(*ResponseEndpointsCreateBulkEndPoints)
	return result, response, err

}

//CreateEndPointTask Create Endpoint task
/* Create Endpoint task

 */
func (s *EndpointsService) CreateEndPointTask(requestEndpointsCreateEndPointTask *RequestEndpointsCreateEndPointTask) (*ResponseEndpointsCreateEndPointTask, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/endpointTask"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEndpointsCreateEndPointTask).
		SetResult(&ResponseEndpointsCreateEndPointTask{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseEndpointsCreateEndPointTask{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateEndPointTask")
	}

	result := response.Result().(*ResponseEndpointsCreateEndPointTask)
	return result, response, err

}

//UpdateBulkEndPoints Update Endpoint in bulk
/* Update Endpoint in bulk

 */
func (s *EndpointsService) UpdateBulkEndPoints(requestEndpointsUpdateBulkEndPoints *RequestEndpointsUpdateBulkEndPoints) (*ResponseEndpointsUpdateBulkEndPoints, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/endpoint/bulk"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEndpointsUpdateBulkEndPoints).
		SetResult(&ResponseEndpointsUpdateBulkEndPoints{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseEndpointsUpdateBulkEndPoints{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateBulkEndPoints")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointsUpdateBulkEndPoints)
	return result, response, err

}

//UpdateEndpoint Update Endpoint by id or mac
/* Update Endpoint by id or mac

@param value value path parameter. The id or MAC of the endpoint
*/
func (s *EndpointsService) UpdateEndpoint(value string, requestEndpointsUpdateEndpoint *RequestEndpointsUpdateEndpoint) (*ResponseEndpointsUpdateEndpoint, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/endpoint/{value}"
	path = strings.Replace(path, "{value}", fmt.Sprintf("%v", value), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEndpointsUpdateEndpoint).
		SetResult(&ResponseEndpointsUpdateEndpoint{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseEndpointsUpdateEndpoint{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateEndpoint")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointsUpdateEndpoint)
	return result, response, err

}

//DeleteBulkEndPoints Delete Endpoint in bulk
/* Delete Endpoint in bulk

 */
func (s *EndpointsService) DeleteBulkEndPoints() (*ResponseEndpointsDeleteBulkEndPoints, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/endpoint/bulk"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEndpointsDeleteBulkEndPoints{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteBulkEndPoints")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointsDeleteBulkEndPoints)
	return result, response, err

}

//DeleteEndpoint Delete endpoint by id or mac
/* Delete endpoint by id or mac

@param value value path parameter. The id or MAC of the endpoint
*/
func (s *EndpointsService) DeleteEndpoint(value string) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/endpoint/{value}"
	path = strings.Replace(path, "{value}", fmt.Sprintf("%v", value), -1)

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
		return response, fmt.Errorf("error with operation DeleteEndpoint")
	}

	getCSFRToken(response.Header())

	return response, err

}
