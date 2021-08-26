package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ProfilerProfileService service

type GetProfilerProfilesQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseProfilerProfileGetProfilerProfileByID struct {
	ProfilerProfile *ResponseProfilerProfileGetProfilerProfileByIDProfilerProfile `json:"ProfilerProfile,omitempty"` //
}

type ResponseProfilerProfileGetProfilerProfileByIDProfilerProfile struct {
	ID          string                                                            `json:"id,omitempty"`          //
	Name        string                                                            `json:"name,omitempty"`        //
	Description string                                                            `json:"description,omitempty"` //
	ParentID    string                                                            `json:"parentId,omitempty"`    //
	Link        *ResponseProfilerProfileGetProfilerProfileByIDProfilerProfileLink `json:"link,omitempty"`        //
}

type ResponseProfilerProfileGetProfilerProfileByIDProfilerProfileLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseProfilerProfileGetProfilerProfiles struct {
	SearchResult *ResponseProfilerProfileGetProfilerProfilesSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseProfilerProfileGetProfilerProfilesSearchResult struct {
	Total        *int                                                                `json:"total,omitempty"`        //
	Resources    *[]ResponseProfilerProfileGetProfilerProfilesSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponseProfilerProfileGetProfilerProfilesSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseProfilerProfileGetProfilerProfilesSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseProfilerProfileGetProfilerProfilesSearchResultResources struct {
	ID          string                                                               `json:"id,omitempty"`          //
	Name        string                                                               `json:"name,omitempty"`        //
	Description string                                                               `json:"description,omitempty"` //
	Link        *ResponseProfilerProfileGetProfilerProfilesSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseProfilerProfileGetProfilerProfilesSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseProfilerProfileGetProfilerProfilesSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseProfilerProfileGetProfilerProfilesSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseProfilerProfileGetVersion struct {
	VersionInfo *ResponseProfilerProfileGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseProfilerProfileGetVersionVersionInfo struct {
	CurrentServerVersion string                                            `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                            `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseProfilerProfileGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseProfilerProfileGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

//GetProfilerProfileByID Get profiler profile by ID
/* This API allows the client to get a profiler profile by ID.

@param id id path parameter.
*/
func (s *ProfilerProfileService) GetProfilerProfileByID(id string) (*ResponseProfilerProfileGetProfilerProfileByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/profilerprofile/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseProfilerProfileGetProfilerProfileByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetProfilerProfileById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseProfilerProfileGetProfilerProfileByID)
	return result, response, err

}

//GetProfilerProfiles Get all profiler profiles
/* This API allows the client to get all the profiler profiles.

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

@param getProfilerProfilesQueryParams Filtering parameter
*/
func (s *ProfilerProfileService) GetProfilerProfiles(getProfilerProfilesQueryParams *GetProfilerProfilesQueryParams) (*ResponseProfilerProfileGetProfilerProfiles, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/profilerprofile"

	queryString, _ := query.Values(getProfilerProfilesQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseProfilerProfileGetProfilerProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetProfilerProfiles")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseProfilerProfileGetProfilerProfiles)
	return result, response, err

}

//GetVersion Get profiler profile version information
/* This API helps to retrieve the version information related to the profiler profile.

 */
func (s *ProfilerProfileService) GetVersion() (*ResponseProfilerProfileGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/profilerprofile/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseProfilerProfileGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseProfilerProfileGetVersion)
	return result, response, err

}
