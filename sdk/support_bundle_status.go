package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SupportBundleStatusService service

type GetSupportBundleStatusQueryParams struct {
	Page int `url: page,omitempty` //Page number
	Size int `url: size,omitempty` //Number of objects returned per page
}

type ResponseSupportBundleStatusGetSupportBundleStatusByID struct {
	SBStatus ResponseSupportBundleStatusGetSupportBundleStatusByIDSBStatus `json:"SBStatus,omitempty"` //
}

type ResponseSupportBundleStatusGetSupportBundleStatusByIDSBStatus struct {
	ID          string                                                            `json:"id,omitempty"`          //
	Name        string                                                            `json:"name,omitempty"`        //
	Description string                                                            `json:"description,omitempty"` //
	FileName    string                                                            `json:"fileName,omitempty"`    //
	FileSize    string                                                            `json:"fileSize,omitempty"`    //
	HostName    string                                                            `json:"hostName,omitempty"`    //
	Message     string                                                            `json:"message,omitempty"`     //
	StartTime   string                                                            `json:"startTime,omitempty"`   //
	Status      string                                                            `json:"status,omitempty"`      //
	Link        ResponseSupportBundleStatusGetSupportBundleStatusByIDSBStatusLink `json:"link,omitempty"`        //
}

type ResponseSupportBundleStatusGetSupportBundleStatusByIDSBStatusLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSupportBundleStatusGetSupportBundleStatus struct {
	SearchResult ResponseSupportBundleStatusGetSupportBundleStatusSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseSupportBundleStatusGetSupportBundleStatusSearchResult struct {
	Total        int                                                                         `json:"total,omitempty"`        //
	Resources    []ResponseSupportBundleStatusGetSupportBundleStatusSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseSupportBundleStatusGetSupportBundleStatusSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseSupportBundleStatusGetSupportBundleStatusSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseSupportBundleStatusGetSupportBundleStatusSearchResultResources struct {
	ID          string                                                                     `json:"id,omitempty"`          //
	Name        string                                                                     `json:"name,omitempty"`        //
	Description string                                                                     `json:"description,omitempty"` //
	Link        ResponseSupportBundleStatusGetSupportBundleStatusSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseSupportBundleStatusGetSupportBundleStatusSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSupportBundleStatusGetSupportBundleStatusSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSupportBundleStatusGetSupportBundleStatusSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSupportBundleStatusGetVersion struct {
	VersionInfo ResponseSupportBundleStatusGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseSupportBundleStatusGetVersionVersionInfo struct {
	CurrentServerVersion string                                               `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                               `json:"supportedVersions,omitempty"`    //
	Link                 ResponseSupportBundleStatusGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseSupportBundleStatusGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

//GetSupportBundleStatusByID Get support bundle status by ID
/* This API allows the client to get a support bundle status by ID.

@param id id path parameter.
*/
func (s *SupportBundleStatusService) GetSupportBundleStatusByID(id string) (*ResponseSupportBundleStatusGetSupportBundleStatusByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/supportbundlestatus/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSupportBundleStatusGetSupportBundleStatusByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSupportBundleStatusById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSupportBundleStatusGetSupportBundleStatusByID)
	return result, response, err

}

//GetSupportBundleStatus Get all support bundle status
/* This API allows the client to get all the support bundle status.

@param getSupportBundleStatusQueryParams Filtering parameter
*/
func (s *SupportBundleStatusService) GetSupportBundleStatus(getSupportBundleStatusQueryParams *GetSupportBundleStatusQueryParams) (*ResponseSupportBundleStatusGetSupportBundleStatus, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/supportbundlestatus"

	queryString, _ := query.Values(getSupportBundleStatusQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSupportBundleStatusGetSupportBundleStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSupportBundleStatus")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSupportBundleStatusGetSupportBundleStatus)
	return result, response, err

}

//GetVersion Get support bundle status version information
/* This API helps to retrieve the version information related to the support bundle status.

 */
func (s *SupportBundleStatusService) GetVersion() (*ResponseSupportBundleStatusGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/supportbundlestatus/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSupportBundleStatusGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSupportBundleStatusGetVersion)
	return result, response, err

}
