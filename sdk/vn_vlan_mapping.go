package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type VnVLANMappingService service

type GetVnVLANMappingsQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sort       string   `url:"sort,omitempty"`       //sort type - asc or desc
	SortBy     string   `url:"sortBy,omitempty"`     //sort column by which objects needs to be sorted
	Filter     []string `url:"filter,omitempty"`     //<div> <style type="text/css" scoped> .apiServiceTable td, .apiServiceTable th { padding: 5px 10px !important; text-align: left; } </style> <span> <b>Simple filtering</b> should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the <i>"filterType=or"</i> query string parameter. Each resource Data model description should specify if an attribute is a filtered field. </span> <br /> <table class="apiServiceTable"> <thead> <tr> <th>OPERATOR</th> <th>DESCRIPTION</th> </tr> </thead> <tbody> <tr> <td>EQ</td> <td>Equals</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> </tr> <tr> <td>GT</td> <td>Greater Than</td> </tr> <tr> <td>LT</td> <td>Less Then</td> </tr> <tr> <td>STARTSW</td> <td>Starts With</td> </tr> <tr> <td>NSTARTSW</td> <td>Not Starts With</td> </tr> <tr> <td>ENDSW</td> <td>Ends With</td> </tr> <tr> <td>NENDSW</td> <td>Not Ends With</td> </tr> <tr> <td>CONTAINS</td> <td>Contains</td> </tr> <tr> <td>NCONTAINS</td> <td>Not Contains</td> </tr> </tbody> </table> </div>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseVnVLANMappingGetVnVLANMappings struct {
	Response *[]ResponseVnVLANMappingGetVnVLANMappingsResponse `json:"response,omitempty"` // Array of VN-Vlan Mappings
}

type ResponseVnVLANMappingGetVnVLANMappingsResponse struct {
	ID            string `json:"id,omitempty"`            // Identifier of the VN-Vlan Mapping
	IsData        *bool  `json:"isData,omitempty"`        // Flag which indicates whether the Vlan is data or voice type
	IsDefaultVLAN *bool  `json:"isDefaultVlan,omitempty"` // Flag which indicates if the Vlan is default
	LastUpdate    string `json:"lastUpdate,omitempty"`    // Timestamp for the last update of the VN-Vlan Mapping
	MaxValue      *int   `json:"maxValue,omitempty"`      // Max value
	Name          string `json:"name,omitempty"`          // Name of the Vlan
	VnID          string `json:"vnId,omitempty"`          // Identifier for the associated Virtual Network which is required unless its name is provided
	VnName        string `json:"vnName,omitempty"`        // Name of the associated Virtual Network to be used for identity if id is not provided
}

type ResponseVnVLANMappingCreateVnVLANMapping struct {
	ID string `json:"id,omitempty"` // ID of the newly created object
}

type ResponseVnVLANMappingBulkCreateVnVLANMappings struct {
	ID string `json:"id,omitempty"` //
}

type ResponseVnVLANMappingBulkDeleteVnVLANMappings struct {
	ID string `json:"id,omitempty"` //
}

type ResponseVnVLANMappingBulkUpdateVnVLANMappings struct {
	ID string `json:"id,omitempty"` //
}

type ResponseVnVLANMappingGetVnVLANMappingByID struct {
	Response *[]ResponseVnVLANMappingGetVnVLANMappingByIDResponse `json:"response,omitempty"` // Array of VN-Vlan Mappings
}

type ResponseVnVLANMappingGetVnVLANMappingByIDResponse struct {
	ID            string `json:"id,omitempty"`            // Identifier of the VN-Vlan Mapping
	IsData        *bool  `json:"isData,omitempty"`        // Flag which indicates whether the Vlan is data or voice type
	IsDefaultVLAN *bool  `json:"isDefaultVlan,omitempty"` // Flag which indicates if the Vlan is default
	LastUpdate    string `json:"lastUpdate,omitempty"`    // Timestamp for the last update of the VN-Vlan Mapping
	MaxValue      *int   `json:"maxValue,omitempty"`      // Max value
	Name          string `json:"name,omitempty"`          // Name of the Vlan
	VnID          string `json:"vnId,omitempty"`          // Identifier for the associated Virtual Network which is required unless its name is provided
	VnName        string `json:"vnName,omitempty"`        // Name of the associated Virtual Network to be used for identity if id is not provided
}

type ResponseVnVLANMappingUpdateVnVLANMappingByID struct {
	Code    *int   `json:"code,omitempty"`    //
	Message string `json:"message,omitempty"` //
}

type ResponseVnVLANMappingDeleteVnVLANMappingByID struct {
	Code    *int   `json:"code,omitempty"`    //
	Message string `json:"message,omitempty"` //
}

type RequestVnVLANMappingCreateVnVLANMapping struct {
	ID            string `json:"id,omitempty"`            // Identifier of the VN-Vlan Mapping
	IsData        *bool  `json:"isData,omitempty"`        // Flag which indicates whether the Vlan is data or voice type
	IsDefaultVLAN *bool  `json:"isDefaultVlan,omitempty"` // Flag which indicates if the Vlan is default
	LastUpdate    string `json:"lastUpdate,omitempty"`    // Timestamp for the last update of the VN-Vlan Mapping
	MaxValue      *int   `json:"maxValue,omitempty"`      // Max value
	Name          string `json:"name,omitempty"`          // Name of the Vlan
	VnID          string `json:"vnId,omitempty"`          // Identifier for the associated Virtual Network which is required unless its name is provided
	VnName        string `json:"vnName,omitempty"`        // Name of the associated Virtual Network to be used for identity if id is not provided
}

type RequestVnVLANMappingBulkCreateVnVLANMappings []RequestItemVnVLANMappingBulkCreateVnVLANMappings // Array of RequestVnVLANMappingBulkCreateVnVlanMappings

type RequestItemVnVLANMappingBulkCreateVnVLANMappings struct {
	ID            string `json:"id,omitempty"`            // Identifier of the VN-Vlan Mapping
	IsData        *bool  `json:"isData,omitempty"`        // Flag which indicates whether the Vlan is data or voice type
	IsDefaultVLAN *bool  `json:"isDefaultVlan,omitempty"` // Flag which indicates if the Vlan is default
	LastUpdate    string `json:"lastUpdate,omitempty"`    // Timestamp for the last update of the VN-Vlan Mapping
	MaxValue      *int   `json:"maxValue,omitempty"`      // Max value
	Name          string `json:"name,omitempty"`          // Name of the Vlan
	VnID          string `json:"vnId,omitempty"`          // Identifier for the associated Virtual Network which is required unless its name is provided
	VnName        string `json:"vnName,omitempty"`        // Name of the associated Virtual Network to be used for identity if id is not provided
}

type RequestVnVLANMappingBulkDeleteVnVLANMappings []string // Array of RequestVnVLANMappingBulkDeleteVnVlanMappings

type RequestVnVLANMappingBulkUpdateVnVLANMappings []RequestItemVnVLANMappingBulkUpdateVnVLANMappings // Array of RequestVnVLANMappingBulkUpdateVnVlanMappings

type RequestItemVnVLANMappingBulkUpdateVnVLANMappings struct {
	ID            string `json:"id,omitempty"`            // Identifier of the VN-Vlan Mapping
	IsData        *bool  `json:"isData,omitempty"`        // Flag which indicates whether the Vlan is data or voice type
	IsDefaultVLAN *bool  `json:"isDefaultVlan,omitempty"` // Flag which indicates if the Vlan is default
	LastUpdate    string `json:"lastUpdate,omitempty"`    // Timestamp for the last update of the VN-Vlan Mapping
	MaxValue      *int   `json:"maxValue,omitempty"`      // Max value
	Name          string `json:"name,omitempty"`          // Name of the Vlan
	VnID          string `json:"vnId,omitempty"`          // Identifier for the associated Virtual Network which is required unless its name is provided
	VnName        string `json:"vnName,omitempty"`        // Name of the associated Virtual Network to be used for identity if id is not provided
}

type RequestVnVLANMappingUpdateVnVLANMappingByID struct {
	ID            string `json:"id,omitempty"`            // Identifier of the VN-Vlan Mapping
	IsData        *bool  `json:"isData,omitempty"`        // Flag which indicates whether the Vlan is data or voice type
	IsDefaultVLAN *bool  `json:"isDefaultVlan,omitempty"` // Flag which indicates if the Vlan is default
	LastUpdate    string `json:"lastUpdate,omitempty"`    // Timestamp for the last update of the VN-Vlan Mapping
	MaxValue      *int   `json:"maxValue,omitempty"`      // Max value
	Name          string `json:"name,omitempty"`          // Name of the Vlan
	VnID          string `json:"vnId,omitempty"`          // Identifier for the associated Virtual Network which is required unless its name is provided
	VnName        string `json:"vnName,omitempty"`        // Name of the associated Virtual Network to be used for identity if id is not provided
}

//GetVnVLANMappings Get all VN-Vlan Mappings
/* Get all VN-Vlan Mappings

@param getVnVlanMappingsQueryParams Filtering parameter
*/
func (s *VnVLANMappingService) GetVnVLANMappings(getVnVlanMappingsQueryParams *GetVnVLANMappingsQueryParams) (*ResponseVnVLANMappingGetVnVLANMappings, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/vnvlanmapping"

	queryString, _ := query.Values(getVnVlanMappingsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseVnVLANMappingGetVnVLANMappings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVnVlanMappings")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseVnVLANMappingGetVnVLANMappings)
	return result, response, err

}

//GetVnVLANMappingByID Get VN-Vlan Mapping by id
/* Get VN-Vlan Mapping by id

@param id id path parameter.
*/
func (s *VnVLANMappingService) GetVnVLANMappingByID(id string) (*ResponseVnVLANMappingGetVnVLANMappingByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/vnvlanmapping/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseVnVLANMappingGetVnVLANMappingByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVnVlanMappingById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseVnVLANMappingGetVnVLANMappingByID)
	return result, response, err

}

//CreateVnVLANMapping Create VN-Vlan Mapping
/* Create VN-Vlan Mapping

 */
func (s *VnVLANMappingService) CreateVnVLANMapping(requestVnVLANMappingCreateVnVlanMapping *RequestVnVLANMappingCreateVnVLANMapping) (*ResponseVnVLANMappingCreateVnVLANMapping, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/vnvlanmapping"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestVnVLANMappingCreateVnVlanMapping).
		SetResult(&ResponseVnVLANMappingCreateVnVLANMapping{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseVnVLANMappingCreateVnVLANMapping{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateVnVlanMapping")
	}

	result := response.Result().(*ResponseVnVLANMappingCreateVnVLANMapping)
	return result, response, err

}

//BulkCreateVnVLANMappings Create VN-Vlan Mappings in bulk
/* Create VN-Vlan Mappings in bulk

 */
func (s *VnVLANMappingService) BulkCreateVnVLANMappings(requestVnVLANMappingBulkCreateVnVlanMappings *RequestVnVLANMappingBulkCreateVnVLANMappings) (*ResponseVnVLANMappingBulkCreateVnVLANMappings, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/vnvlanmapping/bulk/create"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestVnVLANMappingBulkCreateVnVlanMappings).
		SetResult(&ResponseVnVLANMappingBulkCreateVnVLANMappings{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseVnVLANMappingBulkCreateVnVLANMappings{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BulkCreateVnVlanMappings")
	}

	result := response.Result().(*ResponseVnVLANMappingBulkCreateVnVLANMappings)
	return result, response, err

}

//BulkDeleteVnVLANMappings Delete VN-Vlan Mappings in bulk
/* Delete VN-Vlan Mappings in bulk

 */
func (s *VnVLANMappingService) BulkDeleteVnVLANMappings(requestVnVLANMappingBulkDeleteVnVlanMappings *RequestVnVLANMappingBulkDeleteVnVLANMappings) (*ResponseVnVLANMappingBulkDeleteVnVLANMappings, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/vnvlanmapping/bulk/delete"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestVnVLANMappingBulkDeleteVnVlanMappings).
		SetResult(&ResponseVnVLANMappingBulkDeleteVnVLANMappings{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseVnVLANMappingBulkDeleteVnVLANMappings{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BulkDeleteVnVlanMappings")
	}

	result := response.Result().(*ResponseVnVLANMappingBulkDeleteVnVLANMappings)
	return result, response, err

}

//BulkUpdateVnVLANMappings Update VN-Vlan Mappings in bulk
/* Update VN-Vlan Mappings in bulk

 */
func (s *VnVLANMappingService) BulkUpdateVnVLANMappings(requestVnVLANMappingBulkUpdateVnVlanMappings *RequestVnVLANMappingBulkUpdateVnVLANMappings) (*ResponseVnVLANMappingBulkUpdateVnVLANMappings, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/vnvlanmapping/bulk/update"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestVnVLANMappingBulkUpdateVnVlanMappings).
		SetResult(&ResponseVnVLANMappingBulkUpdateVnVLANMappings{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseVnVLANMappingBulkUpdateVnVLANMappings{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BulkUpdateVnVlanMappings")
	}

	result := response.Result().(*ResponseVnVLANMappingBulkUpdateVnVLANMappings)
	return result, response, err

}

//UpdateVnVLANMappingByID Update VN-Vlan Mapping
/* Update VN-Vlan Mapping

@param id id path parameter.
*/
func (s *VnVLANMappingService) UpdateVnVLANMappingByID(id string, requestVnVLANMappingUpdateVnVlanMappingById *RequestVnVLANMappingUpdateVnVLANMappingByID) (*ResponseVnVLANMappingUpdateVnVLANMappingByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/vnvlanmapping/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestVnVLANMappingUpdateVnVlanMappingById).
		SetResult(&ResponseVnVLANMappingUpdateVnVLANMappingByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseVnVLANMappingUpdateVnVLANMappingByID{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateVnVlanMappingById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseVnVLANMappingUpdateVnVLANMappingByID)
	return result, response, err

}

//DeleteVnVLANMappingByID Delete VN-Vlan Mapping
/* Delete VN-Vlan Mapping

@param id id path parameter.
*/
func (s *VnVLANMappingService) DeleteVnVLANMappingByID(id string) (*ResponseVnVLANMappingDeleteVnVLANMappingByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/vnvlanmapping/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseVnVLANMappingDeleteVnVLANMappingByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteVnVlanMappingById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseVnVLANMappingDeleteVnVLANMappingByID)
	return result, response, err

}
