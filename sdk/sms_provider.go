package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SmsProviderService service

type GetSmsProviderQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseSmsProviderGetSmsProvider struct {
	SearchResult ResponseSmsProviderGetSmsProviderSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseSmsProviderGetSmsProviderSearchResult struct {
	Total        int                                                       `json:"total,omitempty"`        //
	Resources    []ResponseSmsProviderGetSmsProviderSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponseSmsProviderGetSmsProviderSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponseSmsProviderGetSmsProviderSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseSmsProviderGetSmsProviderSearchResultResources struct {
	ID   string                                                     `json:"id,omitempty"`   //
	Name string                                                     `json:"name,omitempty"` //
	Link ResponseSmsProviderGetSmsProviderSearchResultResourcesLink `json:"link,omitempty"` //
}

type ResponseSmsProviderGetSmsProviderSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSmsProviderGetSmsProviderSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSmsProviderGetSmsProviderSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSmsProviderGetVersion struct {
	VersionInfo ResponseSmsProviderGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseSmsProviderGetVersionVersionInfo struct {
	CurrentServerVersion string                                       `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                       `json:"supportedVersions,omitempty"`    //
	Link                 ResponseSmsProviderGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseSmsProviderGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

//GetSmsProvider Get all SMS providers
/* This API allows the client to get all the SMS providers.

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

[name, description]

@param getSMSProviderQueryParams Filtering parameter
*/
func (s *SmsProviderService) GetSmsProvider(getSMSProviderQueryParams *GetSmsProviderQueryParams) (*ResponseSmsProviderGetSmsProvider, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/smsprovider"

	queryString, _ := query.Values(getSMSProviderQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSmsProviderGetSmsProvider{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSmsProvider")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSmsProviderGetSmsProvider)
	return result, response, err

}

//GetVersion Get SMS provider version information
/* This API helps to retrieve the version information related to the SMS provider.

 */
func (s *SmsProviderService) GetVersion() (*ResponseSmsProviderGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/smsprovider/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSmsProviderGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSmsProviderGetVersion)
	return result, response, err

}
