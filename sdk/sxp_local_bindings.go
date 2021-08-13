package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SxpLocalBindingsService service

type GetSxpLocalBindingsQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseSxpLocalBindingsGetSxpLocalBindingsByID struct {
	ERSSxpLocalBindings ResponseSxpLocalBindingsGetSxpLocalBindingsByIDERSSxpLocalBindings `json:"ERSSxpLocalBindings,omitempty"` //
}

type ResponseSxpLocalBindingsGetSxpLocalBindingsByIDERSSxpLocalBindings struct {
	ID              string                                                                 `json:"id,omitempty"`              //
	Description     string                                                                 `json:"description,omitempty"`     //
	BindingName     string                                                                 `json:"bindingName,omitempty"`     // This field is depricated from Cisco ISE 3.0
	IPAddressOrHost string                                                                 `json:"ipAddressOrHost,omitempty"` // IP address for static mapping (hostname is not supported)
	SxpVpn          string                                                                 `json:"sxpVpn,omitempty"`          // List of SXP Domains, separated with comma. At least one of: sxpVpn or vns should be defined
	Sgt             string                                                                 `json:"sgt,omitempty"`             // SGT name or ID
	Vns             string                                                                 `json:"vns,omitempty"`             // List of Virtual Networks, separated with comma. At least one of: sxpVpn or vns should be defined
	Link            ResponseSxpLocalBindingsGetSxpLocalBindingsByIDERSSxpLocalBindingsLink `json:"link,omitempty"`            //
}

type ResponseSxpLocalBindingsGetSxpLocalBindingsByIDERSSxpLocalBindingsLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSxpLocalBindingsUpdateSxpLocalBindingsByID struct {
	UpdatedFieldsList ResponseSxpLocalBindingsUpdateSxpLocalBindingsByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseSxpLocalBindingsUpdateSxpLocalBindingsByIDUpdatedFieldsList struct {
	UpdatedField ResponseSxpLocalBindingsUpdateSxpLocalBindingsByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                          `json:"field,omitempty"`        //
	OldValue     string                                                                          `json:"oldValue,omitempty"`     //
	NewValue     string                                                                          `json:"newValue,omitempty"`     //
}

type ResponseSxpLocalBindingsUpdateSxpLocalBindingsByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseSxpLocalBindingsGetSxpLocalBindings struct {
	SearchResult ResponseSxpLocalBindingsGetSxpLocalBindingsSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseSxpLocalBindingsGetSxpLocalBindingsSearchResult struct {
	Total        int                                                                 `json:"total,omitempty"`        //
	Resources    []ResponseSxpLocalBindingsGetSxpLocalBindingsSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponseSxpLocalBindingsGetSxpLocalBindingsSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponseSxpLocalBindingsGetSxpLocalBindingsSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseSxpLocalBindingsGetSxpLocalBindingsSearchResultResources struct {
	ID   string                                                               `json:"id,omitempty"`   //
	Link ResponseSxpLocalBindingsGetSxpLocalBindingsSearchResultResourcesLink `json:"link,omitempty"` //
}

type ResponseSxpLocalBindingsGetSxpLocalBindingsSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSxpLocalBindingsGetSxpLocalBindingsSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSxpLocalBindingsGetSxpLocalBindingsSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSxpLocalBindingsGetVersion struct {
	VersionInfo ResponseSxpLocalBindingsGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseSxpLocalBindingsGetVersionVersionInfo struct {
	CurrentServerVersion string                                            `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                            `json:"supportedVersions,omitempty"`    //
	Link                 ResponseSxpLocalBindingsGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseSxpLocalBindingsGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSxpLocalBindingsMonitorBulkStatusSxpLocalBindings struct {
	BulkStatus ResponseSxpLocalBindingsMonitorBulkStatusSxpLocalBindingsBulkStatus `json:"BulkStatus,omitempty"` //
}

type ResponseSxpLocalBindingsMonitorBulkStatusSxpLocalBindingsBulkStatus struct {
	BulkID          string                                                                               `json:"bulkId,omitempty"`          //
	MediaType       string                                                                               `json:"mediaType,omitempty"`       //
	ExecutionStatus string                                                                               `json:"executionStatus,omitempty"` //
	OperationType   string                                                                               `json:"operationType,omitempty"`   //
	StartTime       string                                                                               `json:"startTime,omitempty"`       //
	ResourcesCount  int                                                                                  `json:"resourcesCount,omitempty"`  //
	SuccessCount    int                                                                                  `json:"successCount,omitempty"`    //
	FailCount       int                                                                                  `json:"failCount,omitempty"`       //
	ResourcesStatus []ResponseSxpLocalBindingsMonitorBulkStatusSxpLocalBindingsBulkStatusResourcesStatus `json:"resourcesStatus,omitempty"` //
}

type ResponseSxpLocalBindingsMonitorBulkStatusSxpLocalBindingsBulkStatusResourcesStatus struct {
	ID                      string `json:"id,omitempty"`                      //
	Name                    string `json:"name,omitempty"`                    //
	Description             string `json:"description,omitempty"`             //
	ResourceExecutionStatus string `json:"resourceExecutionStatus,omitempty"` //
	Status                  string `json:"status,omitempty"`                  //
}

type RequestSxpLocalBindingsUpdateSxpLocalBindingsByID struct {
	ERSSxpLocalBindings *RequestSxpLocalBindingsUpdateSxpLocalBindingsByIDERSSxpLocalBindings `json:"ERSSxpLocalBindings,omitempty"` //
}

type RequestSxpLocalBindingsUpdateSxpLocalBindingsByIDERSSxpLocalBindings struct {
	ID              string `json:"id,omitempty"`              //
	Description     string `json:"description,omitempty"`     //
	BindingName     string `json:"bindingName,omitempty"`     // This field is depricated from Cisco ISE 3.0
	IPAddressOrHost string `json:"ipAddressOrHost,omitempty"` // IP address for static mapping (hostname is not supported)
	SxpVpn          string `json:"sxpVpn,omitempty"`          // List of SXP Domains, separated with comma. At least one of: sxpVpn or vns should be defined
	Sgt             string `json:"sgt,omitempty"`             // SGT name or ID
	Vns             string `json:"vns,omitempty"`             // List of Virtual Networks, separated with comma. At least one of: sxpVpn or vns should be defined
}

type RequestSxpLocalBindingsCreateSxpLocalBindings struct {
	ERSSxpLocalBindings *RequestSxpLocalBindingsCreateSxpLocalBindingsERSSxpLocalBindings `json:"ERSSxpLocalBindings,omitempty"` //
}

type RequestSxpLocalBindingsCreateSxpLocalBindingsERSSxpLocalBindings struct {
	ID              string `json:"id,omitempty"`              //
	Description     string `json:"description,omitempty"`     //
	BindingName     string `json:"bindingName,omitempty"`     // This field is depricated from Cisco ISE 3.0
	IPAddressOrHost string `json:"ipAddressOrHost,omitempty"` // IP address for static mapping (hostname is not supported)
	SxpVpn          string `json:"sxpVpn,omitempty"`          // List of SXP Domains, separated with comma. At least one of: sxpVpn or vns should be defined
	Sgt             string `json:"sgt,omitempty"`             // SGT name or ID
	Vns             string `json:"vns,omitempty"`             // List of Virtual Networks, separated with comma. At least one of: sxpVpn or vns should be defined
}

type RequestSxpLocalBindingsBulkRequestForSxpLocalBindings struct {
	LocalBindingBulkRequest *RequestSxpLocalBindingsBulkRequestForSxpLocalBindingsLocalBindingBulkRequest `json:"LocalBindingBulkRequest,omitempty"` //
}

type RequestSxpLocalBindingsBulkRequestForSxpLocalBindingsLocalBindingBulkRequest struct {
	OperationType     string `json:"operationType,omitempty"`     //
	ResourceMediaType string `json:"resourceMediaType,omitempty"` //
}

//GetSxpLocalBindingsByID Get SXP local binding by ID
/* This API allows the client to get a SXP local binding by ID.

@param id id path parameter.
*/
func (s *SxpLocalBindingsService) GetSxpLocalBindingsByID(id string) (*ResponseSxpLocalBindingsGetSxpLocalBindingsByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxplocalbindings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSxpLocalBindingsGetSxpLocalBindingsByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSxpLocalBindingsById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSxpLocalBindingsGetSxpLocalBindingsByID)
	return result, response, err

}

//GetSxpLocalBindings Get all SXP local bindings
/* This API allows the client to get all the SXP local bindings.

Filter:

[name, description]

To search resources by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:

[name, description]

@param getSXPLocalBindingsQueryParams Filtering parameter
*/
func (s *SxpLocalBindingsService) GetSxpLocalBindings(getSXPLocalBindingsQueryParams *GetSxpLocalBindingsQueryParams) (*ResponseSxpLocalBindingsGetSxpLocalBindings, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxplocalbindings"

	queryString, _ := query.Values(getSXPLocalBindingsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSxpLocalBindingsGetSxpLocalBindings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSxpLocalBindings")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSxpLocalBindingsGetSxpLocalBindings)
	return result, response, err

}

//GetVersion Get SXP local bindings version information
/* This API helps to retrieve the version information related to the SXP local bindings.

 */
func (s *SxpLocalBindingsService) GetVersion() (*ResponseSxpLocalBindingsGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxplocalbindings/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSxpLocalBindingsGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSxpLocalBindingsGetVersion)
	return result, response, err

}

//MonitorBulkStatusSxpLocalBindings Monitor bulk request
/* This API allows the client to monitor the bulk request.

@param bulkid bulkid path parameter.
*/
func (s *SxpLocalBindingsService) MonitorBulkStatusSxpLocalBindings(bulkid string) (*ResponseSxpLocalBindingsMonitorBulkStatusSxpLocalBindings, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxplocalbindings/bulk/{bulkid}"
	path = strings.Replace(path, "{bulkid}", fmt.Sprintf("%v", bulkid), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSxpLocalBindingsMonitorBulkStatusSxpLocalBindings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MonitorBulkStatusSxpLocalBindings")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSxpLocalBindingsMonitorBulkStatusSxpLocalBindings)
	return result, response, err

}

//CreateSxpLocalBindings Create SXP local binding
/* This API creates a SXP local binding.

 */
func (s *SxpLocalBindingsService) CreateSxpLocalBindings(requestSxpLocalBindingsCreateSXPLocalBindings *RequestSxpLocalBindingsCreateSxpLocalBindings) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxplocalbindings"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSxpLocalBindingsCreateSXPLocalBindings).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateSxpLocalBindings")
	}

	return response, err

}

//UpdateSxpLocalBindingsByID Update SXP local binding
/* This API allows the client to update a SXP local binding.

@param id id path parameter.
*/
func (s *SxpLocalBindingsService) UpdateSxpLocalBindingsByID(id string, requestSxpLocalBindingsUpdateSXPLocalBindingsById *RequestSxpLocalBindingsUpdateSxpLocalBindingsByID) (*ResponseSxpLocalBindingsUpdateSxpLocalBindingsByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxplocalbindings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSxpLocalBindingsUpdateSXPLocalBindingsById).
		SetResult(&ResponseSxpLocalBindingsUpdateSxpLocalBindingsByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSxpLocalBindingsById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSxpLocalBindingsUpdateSxpLocalBindingsByID)
	return result, response, err

}

//BulkRequestForSxpLocalBindings Submit bulk request
/* This API allows the client to submit the bulk request.

 */
func (s *SxpLocalBindingsService) BulkRequestForSxpLocalBindings(requestSxpLocalBindingsBulkRequestForSXPLocalBindings *RequestSxpLocalBindingsBulkRequestForSxpLocalBindings) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxplocalbindings/bulk/submit"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSxpLocalBindingsBulkRequestForSXPLocalBindings).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation BulkRequestForSxpLocalBindings")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteSxpLocalBindingsByID Delete SXP local binding
/* This API deletes a SXP local binding.

@param id id path parameter.
*/
func (s *SxpLocalBindingsService) DeleteSxpLocalBindingsByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sxplocalbindings/{id}"
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
		return response, fmt.Errorf("error with operation DeleteSxpLocalBindingsById")
	}

	getCSFRToken(response.Header())

	return response, err

}
