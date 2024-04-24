package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SgVnMappingService service

type GetSgVnMappingListQueryParams struct {
	Page       int    `url:"page,omitempty"`       //Page number
	Size       int    `url:"size,omitempty"`       //Number of objects returned per page
	Sort       string `url:"sort,omitempty"`       //sort type - asc or desc
	SortBy     string `url:"sortBy,omitempty"`     //sort column by which objects needs to be sorted
	Filter     string `url:"filter,omitempty"`     //<div> <style type="text/css" scoped> .apiServiceTable td, .apiServiceTable th { padding: 5px 10px !important; text-align: left; } </style> <span> <b>Simple filtering</b> should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the <i>"filterType=or"</i> query string parameter. Each resource Data model description should specify if an attribute is a filtered field. </span> <br /> <table class="apiServiceTable"> <thead> <tr> <th>OPERATOR</th> <th>DESCRIPTION</th> </tr> </thead> <tbody> <tr> <td>EQ</td> <td>Equals</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> </tr> <tr> <td>GT</td> <td>Greater Than</td> </tr> <tr> <td>LT</td> <td>Less Then</td> </tr> <tr> <td>STARTSW</td> <td>Starts With</td> </tr> <tr> <td>NSTARTSW</td> <td>Not Starts With</td> </tr> <tr> <td>ENDSW</td> <td>Ends With</td> </tr> <tr> <td>NENDSW</td> <td>Not Ends With</td> </tr> <tr> <td>CONTAINS</td> <td>Contains</td> </tr> <tr> <td>NCONTAINS</td> <td>Not Contains</td> </tr> </tbody> </table> </div>
	FilterType string `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseSgVnMappingGetSgVnMappingList struct {
	Response *[]ResponseSgVnMappingGetSgVnMappingListResponse `json:"response,omitempty"` // Array of SG-VN Mapping
}

type ResponseSgVnMappingGetSgVnMappingListResponse struct {
	ID         string `json:"id,omitempty"`         // Identifier of the SG-VN mapping
	LastUpdate string `json:"lastUpdate,omitempty"` // Timestamp for the last update of the SG-VN mapping
	SgName     string `json:"sgName,omitempty"`     // Name of the associated Security Group to be used for identity if id is not provided
	SgtID      string `json:"sgtId,omitempty"`      // Identifier of the associated Security Group which is required unless its name is provided
	VnID       string `json:"vnId,omitempty"`       // Identifier for the associated Virtual Network which is required unless its name is provided
	VnName     string `json:"vnName,omitempty"`     // Name of the associated Virtual Network to be used for identity if id is not provided
}

type ResponseSgVnMappingCreateSgVnMapping struct {
	ID string `json:"id,omitempty"` // ID of the newly created object
}

type ResponseSgVnMappingBulkCreateSgVnMappings struct {
	ID string `json:"id,omitempty"` //
}

type ResponseSgVnMappingBulkDeleteSgVnMappings struct {
	ID string `json:"id,omitempty"` //
}

type ResponseSgVnMappingBulkUpdateSgVnMappings struct {
	ID string `json:"id,omitempty"` //
}

type ResponseSgVnMappingGetSgVnMappingByID struct {
	Response *[]ResponseSgVnMappingGetSgVnMappingByIDResponse `json:"response,omitempty"` // Array of SG-VN Mapping
}

type ResponseSgVnMappingGetSgVnMappingByIDResponse struct {
	ID         string `json:"id,omitempty"`         // Identifier of the SG-VN mapping
	LastUpdate string `json:"lastUpdate,omitempty"` // Timestamp for the last update of the SG-VN mapping
	SgName     string `json:"sgName,omitempty"`     // Name of the associated Security Group to be used for identity if id is not provided
	SgtID      string `json:"sgtId,omitempty"`      // Identifier of the associated Security Group which is required unless its name is provided
	VnID       string `json:"vnId,omitempty"`       // Identifier for the associated Virtual Network which is required unless its name is provided
	VnName     string `json:"vnName,omitempty"`     // Name of the associated Virtual Network to be used for identity if id is not provided
}

type ResponseSgVnMappingUpdateSgVnMapping struct {
	ID         string `json:"id,omitempty"`         // Identifier of the SG-VN mapping
	LastUpdate string `json:"lastUpdate,omitempty"` // Timestamp for the last update of the SG-VN mapping
	SgName     string `json:"sgName,omitempty"`     // Name of the associated Security Group to be used for identity if id is not provided
	SgtID      string `json:"sgtId,omitempty"`      // Identifier of the associated Security Group which is required unless its name is provided
	VnID       string `json:"vnId,omitempty"`       // Identifier for the associated Virtual Network which is required unless its name is provided
	VnName     string `json:"vnName,omitempty"`     // Name of the associated Virtual Network to be used for identity if id is not provided
}

type ResponseSgVnMappingDeleteSgVnMapping struct {
	Message string `json:"message,omitempty"` //
}

type RequestSgVnMappingCreateSgVnMapping struct {
	ID         string `json:"id,omitempty"`         // Identifier of the SG-VN mapping
	LastUpdate string `json:"lastUpdate,omitempty"` // Timestamp for the last update of the SG-VN mapping
	SgName     string `json:"sgName,omitempty"`     // Name of the associated Security Group to be used for identity if id is not provided
	SgtID      string `json:"sgtId,omitempty"`      // Identifier of the associated Security Group which is required unless its name is provided
	VnID       string `json:"vnId,omitempty"`       // Identifier for the associated Virtual Network which is required unless its name is provided
	VnName     string `json:"vnName,omitempty"`     // Name of the associated Virtual Network to be used for identity if id is not provided
}

type RequestSgVnMappingBulkCreateSgVnMappings []RequestItemSgVnMappingBulkCreateSgVnMappings // Array of RequestSgVnMappingBulkCreateSgVnMappings

type RequestItemSgVnMappingBulkCreateSgVnMappings struct {
	ID         string `json:"id,omitempty"`         // Identifier of the SG-VN mapping
	LastUpdate string `json:"lastUpdate,omitempty"` // Timestamp for the last update of the SG-VN mapping
	SgName     string `json:"sgName,omitempty"`     // Name of the associated Security Group to be used for identity if id is not provided
	SgtID      string `json:"sgtId,omitempty"`      // Identifier of the associated Security Group which is required unless its name is provided
	VnID       string `json:"vnId,omitempty"`       // Identifier for the associated Virtual Network which is required unless its name is provided
	VnName     string `json:"vnName,omitempty"`     // Name of the associated Virtual Network to be used for identity if id is not provided
}

type RequestSgVnMappingBulkDeleteSgVnMappings []string // Array of RequestSgVnMappingBulkDeleteSgVnMappings

type RequestSgVnMappingBulkUpdateSgVnMappings []RequestItemSgVnMappingBulkUpdateSgVnMappings // Array of RequestSgVnMappingBulkUpdateSgVnMappings

type RequestItemSgVnMappingBulkUpdateSgVnMappings struct {
	ID         string `json:"id,omitempty"`         // Identifier of the SG-VN mapping
	LastUpdate string `json:"lastUpdate,omitempty"` // Timestamp for the last update of the SG-VN mapping
	SgName     string `json:"sgName,omitempty"`     // Name of the associated Security Group to be used for identity if id is not provided
	SgtID      string `json:"sgtId,omitempty"`      // Identifier of the associated Security Group which is required unless its name is provided
	VnID       string `json:"vnId,omitempty"`       // Identifier for the associated Virtual Network which is required unless its name is provided
	VnName     string `json:"vnName,omitempty"`     // Name of the associated Virtual Network to be used for identity if id is not provided
}

type RequestSgVnMappingUpdateSgVnMapping struct {
	ID         string `json:"id,omitempty"`         // Identifier of the SG-VN mapping
	LastUpdate string `json:"lastUpdate,omitempty"` // Timestamp for the last update of the SG-VN mapping
	SgName     string `json:"sgName,omitempty"`     // Name of the associated Security Group to be used for identity if id is not provided
	SgtID      string `json:"sgtId,omitempty"`      // Identifier of the associated Security Group which is required unless its name is provided
	VnID       string `json:"vnId,omitempty"`       // Identifier for the associated Virtual Network which is required unless its name is provided
	VnName     string `json:"vnName,omitempty"`     // Name of the associated Virtual Network to be used for identity if id is not provided
}

//GetSgVnMappingList Get all SG-VN Mappings
/* Get all Security Group and Virtual Network mappings

@param getSgVnMappingListQueryParams Filtering parameter
*/
func (s *SgVnMappingService) GetSgVnMappingList(getSgVnMappingListQueryParams *GetSgVnMappingListQueryParams) (*ResponseSgVnMappingGetSgVnMappingList, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/sgvnmapping"

	queryString, _ := query.Values(getSgVnMappingListQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSgVnMappingGetSgVnMappingList{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSgVnMappingList")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSgVnMappingGetSgVnMappingList)
	return result, response, err

}

//GetSgVnMappingByID Get SG-VN Mapping by id
/* Get Security Group and Virtual Network mapping by id

@param id id path parameter.
*/
func (s *SgVnMappingService) GetSgVnMappingByID(id string) (*ResponseSgVnMappingGetSgVnMappingByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/sgvnmapping/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSgVnMappingGetSgVnMappingByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSgVnMappingById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSgVnMappingGetSgVnMappingByID)
	return result, response, err

}

//CreateSgVnMapping Create SG-VN Mapping
/* Create Security Group and Virtual Network mapping

 */
func (s *SgVnMappingService) CreateSgVnMapping(requestSgVnMappingCreateSgVnMapping *RequestSgVnMappingCreateSgVnMapping) (*ResponseSgVnMappingCreateSgVnMapping, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/sgvnmapping"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSgVnMappingCreateSgVnMapping).
		SetResult(&ResponseSgVnMappingCreateSgVnMapping{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseSgVnMappingCreateSgVnMapping{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateSgVnMapping")
	}

	result := response.Result().(*ResponseSgVnMappingCreateSgVnMapping)
	return result, response, err

}

//BulkCreateSgVnMappings Create SG-VN Mappings in bulk
/* Create SG-VN Mappings in bulk

 */
func (s *SgVnMappingService) BulkCreateSgVnMappings(requestSgVnMappingBulkCreateSgVnMappings *RequestSgVnMappingBulkCreateSgVnMappings) (*ResponseSgVnMappingBulkCreateSgVnMappings, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/sgvnmapping/bulk/create"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSgVnMappingBulkCreateSgVnMappings).
		SetResult(&ResponseSgVnMappingBulkCreateSgVnMappings{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseSgVnMappingBulkCreateSgVnMappings{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BulkCreateSgVnMappings")
	}

	result := response.Result().(*ResponseSgVnMappingBulkCreateSgVnMappings)
	return result, response, err

}

//BulkDeleteSgVnMappings Delete SG-VN Mappings in bulk
/* Delete SG-VN Mappings in bulk

 */
func (s *SgVnMappingService) BulkDeleteSgVnMappings(requestSgVnMappingBulkDeleteSgVnMappings *RequestSgVnMappingBulkDeleteSgVnMappings) (*ResponseSgVnMappingBulkDeleteSgVnMappings, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/sgvnmapping/bulk/delete"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSgVnMappingBulkDeleteSgVnMappings).
		SetResult(&ResponseSgVnMappingBulkDeleteSgVnMappings{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseSgVnMappingBulkDeleteSgVnMappings{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BulkDeleteSgVnMappings")
	}

	result := response.Result().(*ResponseSgVnMappingBulkDeleteSgVnMappings)
	return result, response, err

}

//BulkUpdateSgVnMappings Update SG-VN Mappings in bulk
/* Update SG-VN Mappings in bulk

 */
func (s *SgVnMappingService) BulkUpdateSgVnMappings(requestSgVnMappingBulkUpdateSgVnMappings *RequestSgVnMappingBulkUpdateSgVnMappings) (*ResponseSgVnMappingBulkUpdateSgVnMappings, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/sgvnmapping/bulk/update"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSgVnMappingBulkUpdateSgVnMappings).
		SetResult(&ResponseSgVnMappingBulkUpdateSgVnMappings{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseSgVnMappingBulkUpdateSgVnMappings{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BulkUpdateSgVnMappings")
	}

	result := response.Result().(*ResponseSgVnMappingBulkUpdateSgVnMappings)
	return result, response, err

}

//UpdateSgVnMapping Update SG-VN mapping
/* Update Security Group and Virtual Network mapping

@param id id path parameter.
*/
func (s *SgVnMappingService) UpdateSgVnMapping(id string, requestSgVnMappingUpdateSgVnMapping *RequestSgVnMappingUpdateSgVnMapping) (*ResponseSgVnMappingUpdateSgVnMapping, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/sgvnmapping/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSgVnMappingUpdateSgVnMapping).
		SetResult(&ResponseSgVnMappingUpdateSgVnMapping{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseSgVnMappingUpdateSgVnMapping{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSgVnMapping")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSgVnMappingUpdateSgVnMapping)
	return result, response, err

}

//DeleteSgVnMapping Delete SG-VN mapping
/* Delete Security Group and Virtual Network mapping

@param id id path parameter.
*/
func (s *SgVnMappingService) DeleteSgVnMapping(id string) (*ResponseSgVnMappingDeleteSgVnMapping, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/sgvnmapping/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSgVnMappingDeleteSgVnMapping{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteSgVnMapping")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSgVnMappingDeleteSgVnMapping)
	return result, response, err

}
