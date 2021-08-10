package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type AncPolicyService service

type GetAncPolicyQueryParams struct {
	Page       int      `url: page,omitempty`       //Page number
	Size       int      `url: size,omitempty`       //Number of objects returned per page
	Sortasc    string   `url: sortasc,omitempty`    //sort asc
	Sortdsc    string   `url: sortdsc,omitempty`    //sort desc
	Filter     []string `url: filter,omitempty`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url: filterType,omitempty` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseAncPolicyGetAncPolicyByName struct {
	ErsAncPolicy ResponseAncPolicyGetAncPolicyByNameErsAncPolicy `json:"ErsAncPolicy,omitempty"` //
}

type ResponseAncPolicyGetAncPolicyByNameErsAncPolicy struct {
	ID      string                                              `json:"id,omitempty"`      //
	Name    string                                              `json:"name,omitempty"`    //
	Actions []string                                            `json:"actions,omitempty"` // - QUARANTINE: Allows you to use Exception policies (authorization policies) to limit or deny an endpoint access to the network. - PORTBOUNCE: Resets the port on the network device to which the endpoint is connected. - SHUTDOWN : Shuts down the port on the network device to which the endpoint is connected. - RE_AUTHENTICATE: Re-authenticates the session from the endpoint.
	Link    ResponseAncPolicyGetAncPolicyByNameErsAncPolicyLink `json:"link,omitempty"`    //
}

type ResponseAncPolicyGetAncPolicyByNameErsAncPolicyLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAncPolicyGetAncPolicyByID struct {
	ErsAncPolicy ResponseAncPolicyGetAncPolicyByIDErsAncPolicy `json:"ErsAncPolicy,omitempty"` //
}

type ResponseAncPolicyGetAncPolicyByIDErsAncPolicy struct {
	ID      string                                            `json:"id,omitempty"`      //
	Name    string                                            `json:"name,omitempty"`    //
	Actions []string                                          `json:"actions,omitempty"` // - QUARANTINE: Allows you to use Exception policies (authorization policies) to limit or deny an endpoint access to the network. - PORTBOUNCE: Resets the port on the network device to which the endpoint is connected. - SHUTDOWN : Shuts down the port on the network device to which the endpoint is connected. - RE_AUTHENTICATE: Re-authenticates the session from the endpoint.
	Link    ResponseAncPolicyGetAncPolicyByIDErsAncPolicyLink `json:"link,omitempty"`    //
}

type ResponseAncPolicyGetAncPolicyByIDErsAncPolicyLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAncPolicyUpdateAncPolicyByID struct {
	UpdatedFieldsList ResponseAncPolicyUpdateAncPolicyByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseAncPolicyUpdateAncPolicyByIDUpdatedFieldsList struct {
	UpdatedField []ResponseAncPolicyUpdateAncPolicyByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
}

type ResponseAncPolicyUpdateAncPolicyByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseAncPolicyGetAncPolicy struct {
	SearchResult ResponseAncPolicyGetAncPolicySearchResult `json:"SearchResult,omitempty"` //
}

type ResponseAncPolicyGetAncPolicySearchResult struct {
	Total        int                                                     `json:"total,omitempty"`        //
	Resources    []ResponseAncPolicyGetAncPolicySearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseAncPolicyGetAncPolicySearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseAncPolicyGetAncPolicySearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseAncPolicyGetAncPolicySearchResultResources struct {
	ID   string                                                 `json:"id,omitempty"`   //
	Name string                                                 `json:"name,omitempty"` //
	Link ResponseAncPolicyGetAncPolicySearchResultResourcesLink `json:"link,omitempty"` //
}

type ResponseAncPolicyGetAncPolicySearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAncPolicyGetAncPolicySearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAncPolicyGetAncPolicySearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAncPolicyGetVersion struct {
	VersionInfo ResponseAncPolicyGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseAncPolicyGetVersionVersionInfo struct {
	CurrentServerVersion string                                     `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                     `json:"supportedVersions,omitempty"`    //
	Link                 ResponseAncPolicyGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseAncPolicyGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAncPolicyMonitorBulkStatusAncPolicy struct {
	BulkStatus ResponseAncPolicyMonitorBulkStatusAncPolicyBulkStatus `json:"BulkStatus,omitempty"` //
}

type ResponseAncPolicyMonitorBulkStatusAncPolicyBulkStatus struct {
	BulkID          string                                                                 `json:"bulkID,omitempty"`          //
	ExecutionStatus string                                                                 `json:"executionStatus,omitempty"` //
	OperationType   string                                                                 `json:"operationType,omitempty"`   //
	StartTime       string                                                                 `json:"startTime,omitempty"`       //
	ResourcesCount  int                                                                    `json:"resourcesCount,omitempty"`  //
	SuccessCount    int                                                                    `json:"successCount,omitempty"`    //
	FailCount       int                                                                    `json:"failCount,omitempty"`       //
	ResourcesStatus []ResponseAncPolicyMonitorBulkStatusAncPolicyBulkStatusResourcesStatus `json:"resourcesStatus,omitempty"` //
}

type ResponseAncPolicyMonitorBulkStatusAncPolicyBulkStatusResourcesStatus struct {
	ID                      string `json:"id,omitempty"`                      //
	Name                    string `json:"name,omitempty"`                    //
	Description             string `json:"description,omitempty"`             //
	ResourceExecutionStatus string `json:"resourceExecutionStatus,omitempty"` //
	Status                  string `json:"status,omitempty"`                  //
}

type RequestAncPolicyUpdateAncPolicyByID struct {
	ErsAncPolicy *RequestAncPolicyUpdateAncPolicyByIDErsAncPolicy `json:"ErsAncPolicy,omitempty"` //
}

type RequestAncPolicyUpdateAncPolicyByIDErsAncPolicy struct {
	ID      string   `json:"id,omitempty"`      //
	Name    string   `json:"name,omitempty"`    //
	Actions []string `json:"actions,omitempty"` // - QUARANTINE: Allows you to use Exception policies (authorization policies) to limit or deny an endpoint access to the network. - PORTBOUNCE: Resets the port on the network device to which the endpoint is connected. - SHUTDOWN : Shuts down the port on the network device to which the endpoint is connected. - RE_AUTHENTICATE: Re-authenticates the session from the endpoint.
}

type RequestAncPolicyCreateAncPolicy struct {
	ErsAncPolicy *RequestAncPolicyCreateAncPolicyErsAncPolicy `json:"ErsAncPolicy,omitempty"` //
}

type RequestAncPolicyCreateAncPolicyErsAncPolicy struct {
	Name    string   `json:"name,omitempty"`    //
	Actions []string `json:"actions,omitempty"` // - QUARANTINE: Allows you to use Exception policies (authorization policies) to limit or deny an endpoint access to the network. - PORTBOUNCE: Resets the port on the network device to which the endpoint is connected. - SHUTDOWN : Shuts down the port on the network device to which the endpoint is connected. - RE_AUTHENTICATE: Re-authenticates the session from the endpoint.
}

type RequestAncPolicyBulkRequestForAncPolicy struct {
	ErsAncPolicyBulkRequest *RequestAncPolicyBulkRequestForAncPolicyErsAncPolicyBulkRequest `json:"ErsAncPolicyBulkRequest,omitempty"` //
}

type RequestAncPolicyBulkRequestForAncPolicyErsAncPolicyBulkRequest struct {
	OperationType     string `json:"operationType,omitempty"`     //
	ResourceMediaType string `json:"resourceMediaType,omitempty"` //
}

//GetAncPolicyByName Get ANC policy by name
/* This API allows the client to get an ANC policy by name.

@param name name path parameter.
*/
func (s *AncPolicyService) GetAncPolicyByName(name string) (*ResponseAncPolicyGetAncPolicyByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ancpolicy/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAncPolicyGetAncPolicyByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAncPolicyByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAncPolicyGetAncPolicyByName)
	return result, response, err

}

//GetAncPolicyByID Get ANC policy by ID
/* This API allows the client to get an ANC policy by ID.

@param id id path parameter.
*/
func (s *AncPolicyService) GetAncPolicyByID(id string) (*ResponseAncPolicyGetAncPolicyByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ancpolicy/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAncPolicyGetAncPolicyByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAncPolicyById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAncPolicyGetAncPolicyByID)
	return result, response, err

}

//GetAncPolicy Get all ANC policies
/* This API allows the client to get all the ANC policies.

Filter:
[name]

To search resources by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:
[name]

@param getAncPolicyQueryParams Filtering parameter
*/
func (s *AncPolicyService) GetAncPolicy(getAncPolicyQueryParams *GetAncPolicyQueryParams) (*ResponseAncPolicyGetAncPolicy, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ancpolicy"

	queryString, _ := query.Values(getAncPolicyQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAncPolicyGetAncPolicy{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAncPolicy")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAncPolicyGetAncPolicy)
	return result, response, err

}

//GetVersion Get ANC policy version information
/* This API helps to retrieve the version information related to the ANC policy

 */
func (s *AncPolicyService) GetVersion() (*ResponseAncPolicyGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ancpolicy/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAncPolicyGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAncPolicyGetVersion)
	return result, response, err

}

//MonitorBulkStatusAncPolicy Monitor bulk request
/* This API allows the client to monitor the bulk request.

@param bulkid bulkid path parameter.
*/
func (s *AncPolicyService) MonitorBulkStatusAncPolicy(bulkid string) (*ResponseAncPolicyMonitorBulkStatusAncPolicy, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ancpolicy/bulk/{bulkid}"
	path = strings.Replace(path, "{bulkid}", fmt.Sprintf("%v", bulkid), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAncPolicyMonitorBulkStatusAncPolicy{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MonitorBulkStatusAncPolicy")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAncPolicyMonitorBulkStatusAncPolicy)
	return result, response, err

}

//CreateAncPolicy Create ANC policy
/* This API allows the client to create an ANC policy.

 */
func (s *AncPolicyService) CreateAncPolicy(requestAncPolicyCreateAncPolicy *RequestAncPolicyCreateAncPolicy) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ancpolicy"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAncPolicyCreateAncPolicy).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateAncPolicy")
	}

	return response, err

}

//UpdateAncPolicyByID Update ANC policy
/* This API allows the client to update an ANC policy.

@param id id path parameter.
*/
func (s *AncPolicyService) UpdateAncPolicyByID(id string, requestAncPolicyUpdateAncPolicyById *RequestAncPolicyUpdateAncPolicyByID) (*ResponseAncPolicyUpdateAncPolicyByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ancpolicy/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAncPolicyUpdateAncPolicyById).
		SetResult(&ResponseAncPolicyUpdateAncPolicyByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateAncPolicyById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAncPolicyUpdateAncPolicyByID)
	return result, response, err

}

//BulkRequestForAncPolicy Submit bulk request
/* This API allows the client to submit the bulk request.

 */
func (s *AncPolicyService) BulkRequestForAncPolicy(requestAncPolicyBulkRequestForAncPolicy *RequestAncPolicyBulkRequestForAncPolicy) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ancpolicy/bulk/submit"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAncPolicyBulkRequestForAncPolicy).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation BulkRequestForAncPolicy")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteAncPolicyByID Delete ANC policy
/* This API allows the client to delete an ANC policy.

@param id id path parameter.
*/
func (s *AncPolicyService) DeleteAncPolicyByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ancpolicy/{id}"
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
		return response, fmt.Errorf("error with operation DeleteAncPolicyById")
	}

	getCSFRToken(response.Header())

	return response, err

}
