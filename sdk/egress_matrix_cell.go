package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type EgressMatrixCellService service

type GetEgressMatrixCellQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseEgressMatrixCellCloneMatrixCell struct {
	OperationResult *ResponseEgressMatrixCellCloneMatrixCellOperationResult `json:"OperationResult,omitempty"` //
}

type ResponseEgressMatrixCellCloneMatrixCellOperationResult struct {
	ResultValue *[]ResponseEgressMatrixCellCloneMatrixCellOperationResultResultValue `json:"resultValue,omitempty"` //
}

type ResponseEgressMatrixCellCloneMatrixCellOperationResultResultValue struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type ResponseEgressMatrixCellGetEgressMatrixCellByID struct {
	EgressMatrixCell *ResponseEgressMatrixCellGetEgressMatrixCellByIDEgressMatrixCell `json:"EgressMatrixCell,omitempty"` //
}

type ResponseEgressMatrixCellGetEgressMatrixCellByIDEgressMatrixCell struct {
	ID               string                                                               `json:"id,omitempty"`               //
	Name             string                                                               `json:"name,omitempty"`             //
	Description      string                                                               `json:"description,omitempty"`      //
	SourceSgtID      string                                                               `json:"sourceSgtId,omitempty"`      //
	DestinationSgtID string                                                               `json:"destinationSgtId,omitempty"` //
	MatrixCellStatus string                                                               `json:"matrixCellStatus,omitempty"` // Allowed values: - DISABLED, - ENABLED, - MONITOR
	DefaultRule      string                                                               `json:"defaultRule,omitempty"`      // Allowed values: - NONE, - DENY_IP, - PERMIT_IP
	Sgacls           []string                                                             `json:"sgacls,omitempty"`           //
	Link             *ResponseEgressMatrixCellGetEgressMatrixCellByIDEgressMatrixCellLink `json:"link,omitempty"`             //
}

type ResponseEgressMatrixCellGetEgressMatrixCellByIDEgressMatrixCellLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseEgressMatrixCellUpdateEgressMatrixCellByID struct {
	UpdatedFieldsList *ResponseEgressMatrixCellUpdateEgressMatrixCellByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseEgressMatrixCellUpdateEgressMatrixCellByIDUpdatedFieldsList struct {
	UpdatedField *[]ResponseEgressMatrixCellUpdateEgressMatrixCellByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                             `json:"field,omitempty"`        //
	OldValue     string                                                                             `json:"oldValue,omitempty"`     //
	NewValue     string                                                                             `json:"newValue,omitempty"`     //
}

type ResponseEgressMatrixCellUpdateEgressMatrixCellByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseEgressMatrixCellGetEgressMatrixCell struct {
	SearchResult *ResponseEgressMatrixCellGetEgressMatrixCellSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseEgressMatrixCellGetEgressMatrixCellSearchResult struct {
	Total        *int                                                                 `json:"total,omitempty"`        //
	Resources    *[]ResponseEgressMatrixCellGetEgressMatrixCellSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponseEgressMatrixCellGetEgressMatrixCellSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseEgressMatrixCellGetEgressMatrixCellSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseEgressMatrixCellGetEgressMatrixCellSearchResultResources struct {
	ID          string                                                                `json:"id,omitempty"`          //
	Name        string                                                                `json:"name,omitempty"`        //
	Description string                                                                `json:"description,omitempty"` //
	Link        *ResponseEgressMatrixCellGetEgressMatrixCellSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseEgressMatrixCellGetEgressMatrixCellSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseEgressMatrixCellGetEgressMatrixCellSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseEgressMatrixCellGetEgressMatrixCellSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseEgressMatrixCellGetVersion struct {
	VersionInfo *ResponseEgressMatrixCellGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseEgressMatrixCellGetVersionVersionInfo struct {
	CurrentServerVersion string                                             `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                             `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseEgressMatrixCellGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseEgressMatrixCellGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseEgressMatrixCellMonitorBulkStatusEgressMatrixCell struct {
	BulkStatus *ResponseEgressMatrixCellMonitorBulkStatusEgressMatrixCellBulkStatus `json:"BulkStatus,omitempty"` //
}

type ResponseEgressMatrixCellMonitorBulkStatusEgressMatrixCellBulkStatus struct {
	BulkID          string                                                                                `json:"bulkId,omitempty"`          //
	MediaType       string                                                                                `json:"mediaType,omitempty"`       //
	ExecutionStatus string                                                                                `json:"executionStatus,omitempty"` //
	OperationType   string                                                                                `json:"operationType,omitempty"`   //
	StartTime       string                                                                                `json:"startTime,omitempty"`       //
	ResourcesCount  *int                                                                                  `json:"resourcesCount,omitempty"`  //
	SuccessCount    *int                                                                                  `json:"successCount,omitempty"`    //
	FailCount       *int                                                                                  `json:"failCount,omitempty"`       //
	ResourcesStatus *[]ResponseEgressMatrixCellMonitorBulkStatusEgressMatrixCellBulkStatusResourcesStatus `json:"resourcesStatus,omitempty"` //
}

type ResponseEgressMatrixCellMonitorBulkStatusEgressMatrixCellBulkStatusResourcesStatus struct {
	ID                      string `json:"id,omitempty"`                      //
	Name                    string `json:"name,omitempty"`                    //
	Description             string `json:"description,omitempty"`             //
	ResourceExecutionStatus string `json:"resourceExecutionStatus,omitempty"` //
	Status                  string `json:"status,omitempty"`                  //
}

type RequestEgressMatrixCellUpdateEgressMatrixCellByID struct {
	EgressMatrixCell *RequestEgressMatrixCellUpdateEgressMatrixCellByIDEgressMatrixCell `json:"EgressMatrixCell,omitempty"` //
}

type RequestEgressMatrixCellUpdateEgressMatrixCellByIDEgressMatrixCell struct {
	ID               string   `json:"id,omitempty"`               //
	Name             string   `json:"name,omitempty"`             //
	Description      string   `json:"description,omitempty"`      //
	SourceSgtID      string   `json:"sourceSgtId,omitempty"`      //
	DestinationSgtID string   `json:"destinationSgtId,omitempty"` //
	MatrixCellStatus string   `json:"matrixCellStatus,omitempty"` // Allowed values: - DISABLED, - ENABLED, - MONITOR
	DefaultRule      string   `json:"defaultRule,omitempty"`      // Allowed values: - NONE, - DENY_IP, - PERMIT_IP
	Sgacls           []string `json:"sgacls,omitempty"`           //
}

type RequestEgressMatrixCellCreateEgressMatrixCell struct {
	EgressMatrixCell *RequestEgressMatrixCellCreateEgressMatrixCellEgressMatrixCell `json:"EgressMatrixCell,omitempty"` //
}

type RequestEgressMatrixCellCreateEgressMatrixCellEgressMatrixCell struct {
	Name             string   `json:"name,omitempty"`             //
	Description      string   `json:"description,omitempty"`      //
	SourceSgtID      string   `json:"sourceSgtId,omitempty"`      //
	DestinationSgtID string   `json:"destinationSgtId,omitempty"` //
	MatrixCellStatus string   `json:"matrixCellStatus,omitempty"` // Allowed values: - DISABLED, - ENABLED, - MONITOR
	DefaultRule      string   `json:"defaultRule,omitempty"`      // Allowed values: - NONE, - DENY_IP, - PERMIT_IP
	Sgacls           []string `json:"sgacls,omitempty"`           //
}

type RequestEgressMatrixCellBulkRequestForEgressMatrixCell struct {
	EgressMatrixCellBulkRequest *RequestEgressMatrixCellBulkRequestForEgressMatrixCellEgressMatrixCellBulkRequest `json:"EgressMatrixCellBulkRequest,omitempty"` //
}

type RequestEgressMatrixCellBulkRequestForEgressMatrixCellEgressMatrixCellBulkRequest struct {
	OperationType     string `json:"operationType,omitempty"`     //
	ResourceMediaType string `json:"resourceMediaType,omitempty"` //
}

//GetEgressMatrixCellByID Get egress matrix cell by ID
/* This API allows the client to get an egress matrix cell by ID.

@param id id path parameter.
*/
func (s *EgressMatrixCellService) GetEgressMatrixCellByID(id string) (*ResponseEgressMatrixCellGetEgressMatrixCellByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/egressmatrixcell/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEgressMatrixCellGetEgressMatrixCellByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetEgressMatrixCellById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEgressMatrixCellGetEgressMatrixCellByID)
	return result, response, err

}

//GetEgressMatrixCell Get all egress matrix cell
/* This API allows the client to get all the egress matrix cell.

Filter:

[sgtSrcValue, matrixStatus, description, sgtSrcName, sgtDstName, sgtDstValue]

To search resources by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:

[description, sgtSrcName, sgtDstName, sgtDstValue]

@param getEgressMatrixCellQueryParams Filtering parameter
*/
func (s *EgressMatrixCellService) GetEgressMatrixCell(getEgressMatrixCellQueryParams *GetEgressMatrixCellQueryParams) (*ResponseEgressMatrixCellGetEgressMatrixCell, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/egressmatrixcell"

	queryString, _ := query.Values(getEgressMatrixCellQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEgressMatrixCellGetEgressMatrixCell{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetEgressMatrixCell")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEgressMatrixCellGetEgressMatrixCell)
	return result, response, err

}

//GetVersion Get egress matrix cell version information
/* This API helps to retrieve the version information related to the egress matrix cell.

 */
func (s *EgressMatrixCellService) GetVersion() (*ResponseEgressMatrixCellGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/egressmatrixcell/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEgressMatrixCellGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEgressMatrixCellGetVersion)
	return result, response, err

}

//MonitorBulkStatusEgressMatrixCell Monitor bulk request
/* This API allows the client to monitor the bulk request.

@param bulkid bulkid path parameter.
*/
func (s *EgressMatrixCellService) MonitorBulkStatusEgressMatrixCell(bulkid string) (*ResponseEgressMatrixCellMonitorBulkStatusEgressMatrixCell, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/egressmatrixcell/bulk/{bulkid}"
	path = strings.Replace(path, "{bulkid}", fmt.Sprintf("%v", bulkid), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEgressMatrixCellMonitorBulkStatusEgressMatrixCell{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MonitorBulkStatusEgressMatrixCell")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEgressMatrixCellMonitorBulkStatusEgressMatrixCell)
	return result, response, err

}

//CreateEgressMatrixCell Create egress matrix cell
/* This API creates an egress matrix cell.

 */
func (s *EgressMatrixCellService) CreateEgressMatrixCell(requestEgressMatrixCellCreateEgressMatrixCell *RequestEgressMatrixCellCreateEgressMatrixCell) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/egressmatrixcell"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEgressMatrixCellCreateEgressMatrixCell).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateEgressMatrixCell")
	}

	return response, err

}

//ClearAllMatrixCells Clear all egress matrix cells
/* This API allows the client to clear all the egress matrix cells.

 */
func (s *EgressMatrixCellService) ClearAllMatrixCells() (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/egressmatrixcell/clearallmatrixcells"

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
		return response, fmt.Errorf("error with operation ClearAllMatrixCells")
	}

	getCSFRToken(response.Header())

	return response, err

}

//SetAllCellsStatus Set status of all egress matrix cells
/* his API allows the client to set status of all the egress matrix cells.

@param status status path parameter.
*/
func (s *EgressMatrixCellService) SetAllCellsStatus(status string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/egressmatrixcell/status/{status}"
	path = strings.Replace(path, "{status}", fmt.Sprintf("%v", status), -1)

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
		return response, fmt.Errorf("error with operation SetAllCellsStatus")
	}

	getCSFRToken(response.Header())

	return response, err

}

//CloneMatrixCell Clone egress matrix cell
/* This API allows the client to clone an egress matrix cell.

@param id id path parameter.
@param srcSgtID srcSgtId path parameter.
@param dstSgtID dstSgtId path parameter.
*/
func (s *EgressMatrixCellService) CloneMatrixCell(id string, srcSgtID string, dstSgtID string) (*ResponseEgressMatrixCellCloneMatrixCell, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/egressmatrixcell/clonecell/{id}/srcSgt/{srcSgtId}/dstSgt/{dstSgtId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{srcSgtId}", fmt.Sprintf("%v", srcSgtID), -1)
	path = strings.Replace(path, "{dstSgtId}", fmt.Sprintf("%v", dstSgtID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEgressMatrixCellCloneMatrixCell{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseEgressMatrixCellCloneMatrixCell{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CloneMatrixCell")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEgressMatrixCellCloneMatrixCell)
	return result, response, err

}

//UpdateEgressMatrixCellByID Update egress matrix cell
/* This API allows the client to update an egress matrix cell.

@param id id path parameter.
*/
func (s *EgressMatrixCellService) UpdateEgressMatrixCellByID(id string, requestEgressMatrixCellUpdateEgressMatrixCellById *RequestEgressMatrixCellUpdateEgressMatrixCellByID) (*ResponseEgressMatrixCellUpdateEgressMatrixCellByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/egressmatrixcell/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEgressMatrixCellUpdateEgressMatrixCellById).
		SetResult(&ResponseEgressMatrixCellUpdateEgressMatrixCellByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseEgressMatrixCellUpdateEgressMatrixCellByID{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateEgressMatrixCellById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEgressMatrixCellUpdateEgressMatrixCellByID)
	return result, response, err

}

//BulkRequestForEgressMatrixCell Submit bulk request
/* This API allows the client to submit the bulk request.

 */
func (s *EgressMatrixCellService) BulkRequestForEgressMatrixCell(requestEgressMatrixCellBulkRequestForEgressMatrixCell *RequestEgressMatrixCellBulkRequestForEgressMatrixCell) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/egressmatrixcell/bulk/submit"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEgressMatrixCellBulkRequestForEgressMatrixCell).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation BulkRequestForEgressMatrixCell")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteEgressMatrixCellByID Delete egress matrix cell
/* This API deletes an egress matrix cell.

@param id id path parameter.
*/
func (s *EgressMatrixCellService) DeleteEgressMatrixCellByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/egressmatrixcell/{id}"
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
		return response, fmt.Errorf("error with operation DeleteEgressMatrixCellById")
	}

	getCSFRToken(response.Header())

	return response, err

}
