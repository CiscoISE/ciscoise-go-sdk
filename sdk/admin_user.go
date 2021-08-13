package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type AdminUserService service

type GetAdminUsersQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseAdminUserGetAdminUserByID struct {
	AdminUser ResponseAdminUserGetAdminUserByIDAdminUser `json:"AdminUser,omitempty"` //
}

type ResponseAdminUserGetAdminUserByIDAdminUser struct {
	Name                         string                                         `json:"name,omitempty"`                         //
	ID                           string                                         `json:"id,omitempty"`                           //
	Description                  string                                         `json:"description,omitempty"`                  //
	Enabled                      bool                                           `json:"enabled,omitempty"`                      //
	Password                     string                                         `json:"password,omitempty"`                     //
	ChangePassword               bool                                           `json:"changePassword,omitempty"`               //
	IncludeSystemAlarmsInEmail   bool                                           `json:"includeSystemAlarmsInEmail,omitempty"`   //
	ExternalUser                 bool                                           `json:"externalUser,omitempty"`                 //
	InactiveAccountNeverDisabled bool                                           `json:"inactiveAccountNeverDisabled,omitempty"` //
	AdminGroups                  string                                         `json:"adminGroups,omitempty"`                  //
	CustomAttributes             map[string]interface{}                         `json:"customAttributes,omitempty"`             //
	Link                         ResponseAdminUserGetAdminUserByIDAdminUserLink `json:"link,omitempty"`                         //
}

type ResponseAdminUserGetAdminUserByIDAdminUserLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAdminUserGetAdminUsers struct {
	SearchResult ResponseAdminUserGetAdminUsersSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseAdminUserGetAdminUsersSearchResult struct {
	Total        int                                                      `json:"total,omitempty"`        //
	Resources    []ResponseAdminUserGetAdminUsersSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseAdminUserGetAdminUsersSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseAdminUserGetAdminUsersSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseAdminUserGetAdminUsersSearchResultResources struct {
	ID          string                                                  `json:"id,omitempty"`          //
	Name        string                                                  `json:"name,omitempty"`        //
	Description string                                                  `json:"description,omitempty"` //
	Link        ResponseAdminUserGetAdminUsersSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseAdminUserGetAdminUsersSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAdminUserGetAdminUsersSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAdminUserGetAdminUsersSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAdminUserGetVersion struct {
	VersionInfo ResponseAdminUserGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseAdminUserGetVersionVersionInfo struct {
	CurrentServerVersion string                                     `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                     `json:"supportedVersions,omitempty"`    //
	Link                 ResponseAdminUserGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseAdminUserGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

//GetAdminUserByID Get admin user by ID
/* This API allows the client to get an admin user by ID.

@param id id path parameter.
*/
func (s *AdminUserService) GetAdminUserByID(id string) (*ResponseAdminUserGetAdminUserByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/adminuser/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAdminUserGetAdminUserByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAdminUserById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAdminUserGetAdminUserByID)
	return result, response, err

}

//GetAdminUsers Get all admin users
/* This API allows the client to get all the admin users.

Filter:

[firstName, lastName, adminGroups, name, description, inactiveAccountNeverDisabled, includeSystemAlarmsInEmail, email, enabled]

To search resources by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:

[name, description]

@param getAdminUsersQueryParams Filtering parameter
*/
func (s *AdminUserService) GetAdminUsers(getAdminUsersQueryParams *GetAdminUsersQueryParams) (*ResponseAdminUserGetAdminUsers, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/adminuser"

	queryString, _ := query.Values(getAdminUsersQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAdminUserGetAdminUsers{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAdminUsers")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAdminUserGetAdminUsers)
	return result, response, err

}

//GetVersion Get admin user version information
/* This API helps to retrieve the version information related to the admin user.

 */
func (s *AdminUserService) GetVersion() (*ResponseAdminUserGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/adminuser/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAdminUserGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAdminUserGetVersion)
	return result, response, err

}
