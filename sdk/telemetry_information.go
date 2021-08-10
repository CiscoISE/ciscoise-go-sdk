package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type TelemetryInformationService service

type GetTelemetryInformationQueryParams struct {
	Page       int      `url: page,omitempty`       //Page number
	Size       int      `url: size,omitempty`       //Number of objects returned per page
	Filter     []string `url: filter,omitempty`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url: filterType,omitempty` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseTelemetryInformationGetTelemetryInfoByID struct {
	TelemetryInfo ResponseTelemetryInformationGetTelemetryInfoByIDTelemetryInfo `json:"TelemetryInfo,omitempty"` //
}

type ResponseTelemetryInformationGetTelemetryInfoByIDTelemetryInfo struct {
	ID           string                                                            `json:"id,omitempty"`           //
	Status       string                                                            `json:"status,omitempty"`       //
	DeploymentID string                                                            `json:"deploymentId,omitempty"` //
	UdiSN        string                                                            `json:"udiSN,omitempty"`        //
	Link         ResponseTelemetryInformationGetTelemetryInfoByIDTelemetryInfoLink `json:"link,omitempty"`         //
}

type ResponseTelemetryInformationGetTelemetryInfoByIDTelemetryInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTelemetryInformationGetTelemetryInformation struct {
	SearchResult ResponseTelemetryInformationGetTelemetryInformationSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseTelemetryInformationGetTelemetryInformationSearchResult struct {
	Total        int                                                                           `json:"total,omitempty"`        //
	Resources    []ResponseTelemetryInformationGetTelemetryInformationSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseTelemetryInformationGetTelemetryInformationSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseTelemetryInformationGetTelemetryInformationSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseTelemetryInformationGetTelemetryInformationSearchResultResources struct {
	ID   string                                                                       `json:"id,omitempty"`   //
	Link ResponseTelemetryInformationGetTelemetryInformationSearchResultResourcesLink `json:"link,omitempty"` //
}

type ResponseTelemetryInformationGetTelemetryInformationSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTelemetryInformationGetTelemetryInformationSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTelemetryInformationGetTelemetryInformationSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTelemetryInformationGetVersion struct {
	VersionInfo ResponseTelemetryInformationGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseTelemetryInformationGetVersionVersionInfo struct {
	CurrentServerVersion string                                                `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                                `json:"supportedVersions,omitempty"`    //
	Link                 ResponseTelemetryInformationGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseTelemetryInformationGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

//GetTelemetryInfoByID Get telemetry information by ID
/* This API allows the client to get telemetry information by ID.

@param id id path parameter.
*/
func (s *TelemetryInformationService) GetTelemetryInfoByID(id string) (*ResponseTelemetryInformationGetTelemetryInfoByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/telemetryinfo/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTelemetryInformationGetTelemetryInfoByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTelemetryInfoById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTelemetryInformationGetTelemetryInfoByID)
	return result, response, err

}

//GetTelemetryInformation Get all telemtry information
/* This API allows the client to get all the telemetry information.

Filter:

[deploymentId]

To search resources by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


@param getTelemetryInformationQueryParams Filtering parameter
*/
func (s *TelemetryInformationService) GetTelemetryInformation(getTelemetryInformationQueryParams *GetTelemetryInformationQueryParams) (*ResponseTelemetryInformationGetTelemetryInformation, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/telemetryinfo"

	queryString, _ := query.Values(getTelemetryInformationQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTelemetryInformationGetTelemetryInformation{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTelemetryInformation")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTelemetryInformationGetTelemetryInformation)
	return result, response, err

}

//GetVersion Get telemtry information version information
/* This API helps to retrieve the version information related to the telemetry information.

 */
func (s *TelemetryInformationService) GetVersion() (*ResponseTelemetryInformationGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/telemetryinfo/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTelemetryInformationGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTelemetryInformationGetVersion)
	return result, response, err

}
