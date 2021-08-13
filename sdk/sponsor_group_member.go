package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SponsorGroupMemberService service

type GetSponsorGroupMemberQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseSponsorGroupMemberGetSponsorGroupMember struct {
	SearchResult ResponseSponsorGroupMemberGetSponsorGroupMemberSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseSponsorGroupMemberGetSponsorGroupMemberSearchResult struct {
	Total        int                                                                     `json:"total,omitempty"`        //
	Resources    []ResponseSponsorGroupMemberGetSponsorGroupMemberSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponseSponsorGroupMemberGetSponsorGroupMemberSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponseSponsorGroupMemberGetSponsorGroupMemberSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseSponsorGroupMemberGetSponsorGroupMemberSearchResultResources struct {
	ID   string                                                                   `json:"id,omitempty"`   //
	Name string                                                                   `json:"name,omitempty"` //
	Link ResponseSponsorGroupMemberGetSponsorGroupMemberSearchResultResourcesLink `json:"link,omitempty"` //
}

type ResponseSponsorGroupMemberGetSponsorGroupMemberSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSponsorGroupMemberGetSponsorGroupMemberSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSponsorGroupMemberGetSponsorGroupMemberSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSponsorGroupMemberGetVersion struct {
	VersionInfo ResponseSponsorGroupMemberGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseSponsorGroupMemberGetVersionVersionInfo struct {
	CurrentServerVersion string                                              `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                              `json:"supportedVersions,omitempty"`    //
	Link                 ResponseSponsorGroupMemberGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseSponsorGroupMemberGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

//GetSponsorGroupMember Get all sponsor group members
/* This API allows the client to get all the sponsor group members.

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

@param getSponsorGroupMemberQueryParams Filtering parameter
*/
func (s *SponsorGroupMemberService) GetSponsorGroupMember(getSponsorGroupMemberQueryParams *GetSponsorGroupMemberQueryParams) (*ResponseSponsorGroupMemberGetSponsorGroupMember, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsorgroupmember"

	queryString, _ := query.Values(getSponsorGroupMemberQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSponsorGroupMemberGetSponsorGroupMember{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSponsorGroupMember")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSponsorGroupMemberGetSponsorGroupMember)
	return result, response, err

}

//GetVersion Get sponsor group member version information
/* This API helps to retrieve the version information related to the sponsor group member.

 */
func (s *SponsorGroupMemberService) GetVersion() (*ResponseSponsorGroupMemberGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsorgroupmember/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSponsorGroupMemberGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSponsorGroupMemberGetVersion)
	return result, response, err

}
