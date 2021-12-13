package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type VirtualNetworkService service

type GetVirtualNetworksQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sort       string   `url:"sort,omitempty"`       //sort type - asc or desc
	SortBy     string   `url:"sortBy,omitempty"`     //sort column by which objects needs to be sorted
	Filter     []string `url:"filter,omitempty"`     //<div> <style type="text/css" scoped> .apiServiceTable td, .apiServiceTable th { padding: 5px 10px !important; text-align: left; } </style> <span> <b>Simple filtering</b> should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the <i>"filterType=or"</i> query string parameter. Each resource Data model description should specify if an attribute is a filtered field. </span> <br /> <table class="apiServiceTable"> <thead> <tr> <th>OPERATOR</th> <th>DESCRIPTION</th> </tr> </thead> <tbody> <tr> <td>EQ</td> <td>Equals</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> </tr> <tr> <td>GT</td> <td>Greater Than</td> </tr> <tr> <td>LT</td> <td>Less Then</td> </tr> <tr> <td>STARTSW</td> <td>Starts With</td> </tr> <tr> <td>NSTARTSW</td> <td>Not Starts With</td> </tr> <tr> <td>ENDSW</td> <td>Ends With</td> </tr> <tr> <td>NENDSW</td> <td>Not Ends With</td> </tr> <tr> <td>CONTAINS</td> <td>Contains</td> </tr> <tr> <td>NCONTAINS</td> <td>Not Contains</td> </tr> </tbody> </table> </div>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseVirtualNetworkGetVirtualNetworks struct {
	Response *[]ResponseVirtualNetworkGetVirtualNetworksResponse `json:"response,omitempty"` // Array of Virtual Network
}

type ResponseVirtualNetworkGetVirtualNetworksResponse struct {
	AdditionalAttributes string `json:"additionalAttributes,omitempty"` // JSON String of additional attributes for the Virtual Network
	ID                   string `json:"id,omitempty"`                   // Identifier of the Virtual Network
	LastUpdate           string `json:"lastUpdate,omitempty"`           // Timestamp for the last update of the Virtual Network
	Name                 string `json:"name,omitempty"`                 // Name of the Virtual Network
}

type ResponseVirtualNetworkCreateVirtualNetwork struct {
	ID string `json:"id,omitempty"` // ID of the newly created object
}

type ResponseVirtualNetworkBulkCreateVirtualNetworks struct {
	ID string `json:"id,omitempty"` //
}

type ResponseVirtualNetworkBulkDeleteVirtualNetworks struct {
	ID string `json:"id,omitempty"` //
}

type ResponseVirtualNetworkBulkUpdateVirtualNetworks struct {
	ID string `json:"id,omitempty"` //
}

type ResponseVirtualNetworkGetVirtualNetworkByID struct {
	Response *[]ResponseVirtualNetworkGetVirtualNetworkByIDResponse `json:"response,omitempty"` // Array of Virtual Network
}

type ResponseVirtualNetworkGetVirtualNetworkByIDResponse struct {
	AdditionalAttributes string `json:"additionalAttributes,omitempty"` // JSON String of additional attributes for the Virtual Network
	ID                   string `json:"id,omitempty"`                   // Identifier of the Virtual Network
	LastUpdate           string `json:"lastUpdate,omitempty"`           // Timestamp for the last update of the Virtual Network
	Name                 string `json:"name,omitempty"`                 // Name of the Virtual Network
}

type ResponseVirtualNetworkUpdateVirtualNetworkByID struct {
	Code    *int   `json:"code,omitempty"`    //
	Message string `json:"message,omitempty"` //
}

type ResponseVirtualNetworkDeleteVirtualNetworkByID struct {
	Code    *int   `json:"code,omitempty"`    //
	Message string `json:"message,omitempty"` //
}

type RequestVirtualNetworkCreateVirtualNetwork struct {
	AdditionalAttributes string `json:"additionalAttributes,omitempty"` // JSON String of additional attributes for the Virtual Network
	ID                   string `json:"id,omitempty"`                   // Identifier of the Virtual Network
	LastUpdate           string `json:"lastUpdate,omitempty"`           // Timestamp for the last update of the Virtual Network
	Name                 string `json:"name,omitempty"`                 // Name of the Virtual Network
}

type RequestVirtualNetworkBulkCreateVirtualNetworks []RequestItemVirtualNetworkBulkCreateVirtualNetworks // Array of RequestVirtualNetworkBulkCreateVirtualNetworks

type RequestItemVirtualNetworkBulkCreateVirtualNetworks struct {
	AdditionalAttributes string `json:"additionalAttributes,omitempty"` // JSON String of additional attributes for the Virtual Network
	ID                   string `json:"id,omitempty"`                   // Identifier of the Virtual Network
	LastUpdate           string `json:"lastUpdate,omitempty"`           // Timestamp for the last update of the Virtual Network
	Name                 string `json:"name,omitempty"`                 // Name of the Virtual Network
}

type RequestVirtualNetworkBulkDeleteVirtualNetworks []string // Array of RequestVirtualNetworkBulkDeleteVirtualNetworks

type RequestVirtualNetworkBulkUpdateVirtualNetworks []RequestItemVirtualNetworkBulkUpdateVirtualNetworks // Array of RequestVirtualNetworkBulkUpdateVirtualNetworks

type RequestItemVirtualNetworkBulkUpdateVirtualNetworks struct {
	AdditionalAttributes string `json:"additionalAttributes,omitempty"` // JSON String of additional attributes for the Virtual Network
	ID                   string `json:"id,omitempty"`                   // Identifier of the Virtual Network
	LastUpdate           string `json:"lastUpdate,omitempty"`           // Timestamp for the last update of the Virtual Network
	Name                 string `json:"name,omitempty"`                 // Name of the Virtual Network
}

type RequestVirtualNetworkUpdateVirtualNetworkByID struct {
	AdditionalAttributes string `json:"additionalAttributes,omitempty"` // JSON String of additional attributes for the Virtual Network
	ID                   string `json:"id,omitempty"`                   // Identifier of the Virtual Network
	LastUpdate           string `json:"lastUpdate,omitempty"`           // Timestamp for the last update of the Virtual Network
	Name                 string `json:"name,omitempty"`                 // Name of the Virtual Network
}

//GetVirtualNetworks Get all Virtual Networks
/* Get all Virtual Networks

@param getVirtualNetworksQueryParams Filtering parameter
*/
func (s *VirtualNetworkService) GetVirtualNetworks(getVirtualNetworksQueryParams *GetVirtualNetworksQueryParams) (*ResponseVirtualNetworkGetVirtualNetworks, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/virtualnetwork"

	queryString, _ := query.Values(getVirtualNetworksQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseVirtualNetworkGetVirtualNetworks{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVirtualNetworks")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseVirtualNetworkGetVirtualNetworks)
	return result, response, err

}

//GetVirtualNetworkByID Get Virtual Network by id
/* Get Virtual Network by id

@param id id path parameter.
*/
func (s *VirtualNetworkService) GetVirtualNetworkByID(id string) (*ResponseVirtualNetworkGetVirtualNetworkByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/virtualnetwork/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseVirtualNetworkGetVirtualNetworkByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVirtualNetworkById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseVirtualNetworkGetVirtualNetworkByID)
	return result, response, err

}

//CreateVirtualNetwork Create Virtual Network
/* Create Virtual Network

 */
func (s *VirtualNetworkService) CreateVirtualNetwork(requestVirtualNetworkCreateVirtualNetwork *RequestVirtualNetworkCreateVirtualNetwork) (*ResponseVirtualNetworkCreateVirtualNetwork, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/virtualnetwork"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestVirtualNetworkCreateVirtualNetwork).
		SetResult(&ResponseVirtualNetworkCreateVirtualNetwork{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateVirtualNetwork")
	}

	result := response.Result().(*ResponseVirtualNetworkCreateVirtualNetwork)
	return result, response, err

}

//BulkCreateVirtualNetworks Create Virtual Network in bulk
/* Create Virtual Network in bulk

 */
func (s *VirtualNetworkService) BulkCreateVirtualNetworks(requestVirtualNetworkBulkCreateVirtualNetworks *RequestVirtualNetworkBulkCreateVirtualNetworks) (*ResponseVirtualNetworkBulkCreateVirtualNetworks, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/virtualnetwork/bulk/create"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestVirtualNetworkBulkCreateVirtualNetworks).
		SetResult(&ResponseVirtualNetworkBulkCreateVirtualNetworks{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BulkCreateVirtualNetworks")
	}

	result := response.Result().(*ResponseVirtualNetworkBulkCreateVirtualNetworks)
	return result, response, err

}

//BulkDeleteVirtualNetworks Delete Virtual Network in bulk
/* Delete Virtual Network in bulk

 */
func (s *VirtualNetworkService) BulkDeleteVirtualNetworks(requestVirtualNetworkBulkDeleteVirtualNetworks *RequestVirtualNetworkBulkDeleteVirtualNetworks) (*ResponseVirtualNetworkBulkDeleteVirtualNetworks, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/virtualnetwork/bulk/delete"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestVirtualNetworkBulkDeleteVirtualNetworks).
		SetResult(&ResponseVirtualNetworkBulkDeleteVirtualNetworks{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BulkDeleteVirtualNetworks")
	}

	result := response.Result().(*ResponseVirtualNetworkBulkDeleteVirtualNetworks)
	return result, response, err

}

//BulkUpdateVirtualNetworks Update Virtual Network in bulk
/* Update Virtual Network in bulk

 */
func (s *VirtualNetworkService) BulkUpdateVirtualNetworks(requestVirtualNetworkBulkUpdateVirtualNetworks *RequestVirtualNetworkBulkUpdateVirtualNetworks) (*ResponseVirtualNetworkBulkUpdateVirtualNetworks, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/virtualnetwork/bulk/update"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestVirtualNetworkBulkUpdateVirtualNetworks).
		SetResult(&ResponseVirtualNetworkBulkUpdateVirtualNetworks{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BulkUpdateVirtualNetworks")
	}

	result := response.Result().(*ResponseVirtualNetworkBulkUpdateVirtualNetworks)
	return result, response, err

}

//UpdateVirtualNetworkByID Update Virtual Network
/* Update Virtual Network

@param id id path parameter.
*/
func (s *VirtualNetworkService) UpdateVirtualNetworkByID(id string, requestVirtualNetworkUpdateVirtualNetworkById *RequestVirtualNetworkUpdateVirtualNetworkByID) (*ResponseVirtualNetworkUpdateVirtualNetworkByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/virtualnetwork/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestVirtualNetworkUpdateVirtualNetworkById).
		SetResult(&ResponseVirtualNetworkUpdateVirtualNetworkByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateVirtualNetworkById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseVirtualNetworkUpdateVirtualNetworkByID)
	return result, response, err

}

//DeleteVirtualNetworkByID Delete Virtual Network
/* Delete Virtual Network

@param id id path parameter.
*/
func (s *VirtualNetworkService) DeleteVirtualNetworkByID(id string) (*ResponseVirtualNetworkDeleteVirtualNetworkByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/virtualnetwork/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseVirtualNetworkDeleteVirtualNetworkByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteVirtualNetworkById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseVirtualNetworkDeleteVirtualNetworkByID)
	return result, response, err

}
