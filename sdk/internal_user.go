package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type InternalUserService service

type GetInternalUserQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseInternalUserGetInternalUserByName struct {
	InternalUser ResponseInternalUserGetInternalUserByNameInternalUser `json:"InternalUser,omitempty"` //
}

type ResponseInternalUserGetInternalUserByNameInternalUser struct {
	ID                string                                                    `json:"id,omitempty"`                //
	Name              string                                                    `json:"name,omitempty"`              //
	Description       string                                                    `json:"description,omitempty"`       //
	Enabled           bool                                                      `json:"enabled,omitempty"`           // Whether the user is enabled/disabled. To use it as filter, the values should be 'Enabled' or 'Disabled'. The values are case sensitive. For example, '[ERSObjectURL]?filter=enabled.EQ.Enabled'
	Email             string                                                    `json:"email,omitempty"`             //
	Password          string                                                    `json:"password,omitempty"`          //
	FirstName         string                                                    `json:"firstName,omitempty"`         //
	LastName          string                                                    `json:"lastName,omitempty"`          //
	ChangePassword    bool                                                      `json:"changePassword,omitempty"`    //
	IDentityGroups    string                                                    `json:"identityGroups,omitempty"`    // CSV of identity group IDs
	ExpiryDateEnabled bool                                                      `json:"expiryDateEnabled,omitempty"` //
	ExpiryDate        string                                                    `json:"expiryDate,omitempty"`        // To store the internal user's expiry date information. It's format is = 'YYYY-MM-DD'
	EnablePassword    string                                                    `json:"enablePassword,omitempty"`    //
	CustomAttributes  map[string]interface{}                                    `json:"customAttributes,omitempty"`  // Key value map
	PasswordIDStore   string                                                    `json:"passwordIDStore,omitempty"`   // The id store where the internal user's password is kept
	Link              ResponseInternalUserGetInternalUserByNameInternalUserLink `json:"link,omitempty"`              //
}

type ResponseInternalUserGetInternalUserByNameInternalUserLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseInternalUserUpdateInternalUserByName struct {
	UpdatedFieldsList ResponseInternalUserUpdateInternalUserByNameUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseInternalUserUpdateInternalUserByNameUpdatedFieldsList struct {
	UpdatedField ResponseInternalUserUpdateInternalUserByNameUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                    `json:"field,omitempty"`        //
	OldValue     string                                                                    `json:"oldValue,omitempty"`     //
	NewValue     string                                                                    `json:"newValue,omitempty"`     //
}

type ResponseInternalUserUpdateInternalUserByNameUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseInternalUserGetInternalUserByID struct {
	InternalUser ResponseInternalUserGetInternalUserByIDInternalUser `json:"InternalUser,omitempty"` //
}

type ResponseInternalUserGetInternalUserByIDInternalUser struct {
	ID                string                                                  `json:"id,omitempty"`                //
	Name              string                                                  `json:"name,omitempty"`              //
	Description       string                                                  `json:"description,omitempty"`       //
	Enabled           bool                                                    `json:"enabled,omitempty"`           // Whether the user is enabled/disabled. To use it as filter, the values should be 'Enabled' or 'Disabled'. The values are case sensitive. For example, '[ERSObjectURL]?filter=enabled.EQ.Enabled'
	Email             string                                                  `json:"email,omitempty"`             //
	Password          string                                                  `json:"password,omitempty"`          //
	FirstName         string                                                  `json:"firstName,omitempty"`         //
	LastName          string                                                  `json:"lastName,omitempty"`          //
	ChangePassword    bool                                                    `json:"changePassword,omitempty"`    //
	IDentityGroups    string                                                  `json:"identityGroups,omitempty"`    // CSV of identity group IDs
	ExpiryDateEnabled bool                                                    `json:"expiryDateEnabled,omitempty"` //
	ExpiryDate        string                                                  `json:"expiryDate,omitempty"`        // To store the internal user's expiry date information. It's format is = 'YYYY-MM-DD'
	EnablePassword    string                                                  `json:"enablePassword,omitempty"`    //
	CustomAttributes  map[string]interface{}                                  `json:"customAttributes,omitempty"`  // Key value map
	PasswordIDStore   string                                                  `json:"passwordIDStore,omitempty"`   // The id store where the internal user's password is kept
	Link              ResponseInternalUserGetInternalUserByIDInternalUserLink `json:"link,omitempty"`              //
}

type ResponseInternalUserGetInternalUserByIDInternalUserLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseInternalUserUpdateInternalUserByID struct {
	UpdatedFieldsList ResponseInternalUserUpdateInternalUserByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseInternalUserUpdateInternalUserByIDUpdatedFieldsList struct {
	UpdatedField ResponseInternalUserUpdateInternalUserByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                  `json:"field,omitempty"`        //
	OldValue     string                                                                  `json:"oldValue,omitempty"`     //
	NewValue     string                                                                  `json:"newValue,omitempty"`     //
}

type ResponseInternalUserUpdateInternalUserByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseInternalUserGetInternalUser struct {
	SearchResult ResponseInternalUserGetInternalUserSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseInternalUserGetInternalUserSearchResult struct {
	Total        int                                                         `json:"total,omitempty"`        //
	Resources    []ResponseInternalUserGetInternalUserSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponseInternalUserGetInternalUserSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponseInternalUserGetInternalUserSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseInternalUserGetInternalUserSearchResultResources struct {
	ID          string                                                       `json:"id,omitempty"`          //
	Name        string                                                       `json:"name,omitempty"`        //
	Description string                                                       `json:"description,omitempty"` //
	Link        ResponseInternalUserGetInternalUserSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseInternalUserGetInternalUserSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseInternalUserGetInternalUserSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseInternalUserGetInternalUserSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseInternalUserGetVersion struct {
	VersionInfo ResponseInternalUserGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseInternalUserGetVersionVersionInfo struct {
	CurrentServerVersion string                                        `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                        `json:"supportedVersions,omitempty"`    //
	Link                 ResponseInternalUserGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseInternalUserGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestInternalUserUpdateInternalUserByName struct {
	InternalUser *RequestInternalUserUpdateInternalUserByNameInternalUser `json:"InternalUser,omitempty"` //
}

type RequestInternalUserUpdateInternalUserByNameInternalUser struct {
	ID                string                  `json:"id,omitempty"`                //
	Name              string                  `json:"name,omitempty"`              //
	Description       string                  `json:"description,omitempty"`       //
	Enabled           *bool                   `json:"enabled,omitempty"`           // Whether the user is enabled/disabled. To use it as filter, the values should be 'Enabled' or 'Disabled'. The values are case sensitive. For example, '[ERSObjectURL]?filter=enabled.EQ.Enabled'
	Email             string                  `json:"email,omitempty"`             //
	Password          string                  `json:"password,omitempty"`          //
	FirstName         string                  `json:"firstName,omitempty"`         //
	LastName          string                  `json:"lastName,omitempty"`          //
	ChangePassword    *bool                   `json:"changePassword,omitempty"`    //
	IDentityGroups    string                  `json:"identityGroups,omitempty"`    // CSV of identity group IDs
	ExpiryDateEnabled *bool                   `json:"expiryDateEnabled,omitempty"` //
	ExpiryDate        string                  `json:"expiryDate,omitempty"`        // To store the internal user's expiry date information. It's format is = 'YYYY-MM-DD'
	EnablePassword    string                  `json:"enablePassword,omitempty"`    //
	CustomAttributes  *map[string]interface{} `json:"customAttributes,omitempty"`  // Key value map
	PasswordIDStore   string                  `json:"passwordIDStore,omitempty"`   // The id store where the internal user's password is kept
}

type RequestInternalUserUpdateInternalUserByID struct {
	InternalUser *RequestInternalUserUpdateInternalUserByIDInternalUser `json:"InternalUser,omitempty"` //
}

type RequestInternalUserUpdateInternalUserByIDInternalUser struct {
	ID                string                  `json:"id,omitempty"`                //
	Name              string                  `json:"name,omitempty"`              //
	Description       string                  `json:"description,omitempty"`       //
	Enabled           *bool                   `json:"enabled,omitempty"`           // Whether the user is enabled/disabled. To use it as filter, the values should be 'Enabled' or 'Disabled'. The values are case sensitive. For example, '[ERSObjectURL]?filter=enabled.EQ.Enabled'
	Email             string                  `json:"email,omitempty"`             //
	Password          string                  `json:"password,omitempty"`          //
	FirstName         string                  `json:"firstName,omitempty"`         //
	LastName          string                  `json:"lastName,omitempty"`          //
	ChangePassword    *bool                   `json:"changePassword,omitempty"`    //
	IDentityGroups    string                  `json:"identityGroups,omitempty"`    // CSV of identity group IDs
	ExpiryDateEnabled *bool                   `json:"expiryDateEnabled,omitempty"` //
	ExpiryDate        string                  `json:"expiryDate,omitempty"`        // To store the internal user's expiry date information. It's format is = 'YYYY-MM-DD'
	EnablePassword    string                  `json:"enablePassword,omitempty"`    //
	CustomAttributes  *map[string]interface{} `json:"customAttributes,omitempty"`  // Key value map
	PasswordIDStore   string                  `json:"passwordIDStore,omitempty"`   // The id store where the internal user's password is kept
}

type RequestInternalUserCreateInternalUser struct {
	InternalUser *RequestInternalUserCreateInternalUserInternalUser `json:"InternalUser,omitempty"` //
}

type RequestInternalUserCreateInternalUserInternalUser struct {
	Name              string                  `json:"name,omitempty"`              //
	Description       string                  `json:"description,omitempty"`       //
	Enabled           *bool                   `json:"enabled,omitempty"`           // Whether the user is enabled/disabled. To use it as filter, the values should be 'Enabled' or 'Disabled'. The values are case sensitive. For example, '[ERSObjectURL]?filter=enabled.EQ.Enabled'
	Email             string                  `json:"email,omitempty"`             //
	Password          string                  `json:"password,omitempty"`          //
	FirstName         string                  `json:"firstName,omitempty"`         //
	LastName          string                  `json:"lastName,omitempty"`          //
	ChangePassword    *bool                   `json:"changePassword,omitempty"`    //
	IDentityGroups    string                  `json:"identityGroups,omitempty"`    // CSV of identity group IDs
	ExpiryDateEnabled *bool                   `json:"expiryDateEnabled,omitempty"` //
	ExpiryDate        string                  `json:"expiryDate,omitempty"`        // To store the internal user's expiry date information. It's format is = 'YYYY-MM-DD'
	EnablePassword    string                  `json:"enablePassword,omitempty"`    //
	CustomAttributes  *map[string]interface{} `json:"customAttributes,omitempty"`  // Key value map
	PasswordIDStore   string                  `json:"passwordIDStore,omitempty"`   // The id store where the internal user's password is kept
}

//GetInternalUserByName Get internal user by name
/* This API allows the client to get an internal user by name.

@param name name path parameter.
*/
func (s *InternalUserService) GetInternalUserByName(name string) (*ResponseInternalUserGetInternalUserByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/internaluser/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseInternalUserGetInternalUserByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetInternalUserByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseInternalUserGetInternalUserByName)
	return result, response, err

}

//GetInternalUserByID Get internal user by ID
/* This API allows the client to get an internal user by ID.

@param id id path parameter.
*/
func (s *InternalUserService) GetInternalUserByID(id string) (*ResponseInternalUserGetInternalUserByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/internaluser/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseInternalUserGetInternalUserByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetInternalUserById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseInternalUserGetInternalUserByID)
	return result, response, err

}

//GetInternalUser Get all internal users
/* This API allows the client to get all the internal users.

Filter:

[firstName, lastName, identityGroup, name, description, email, enabled]

To search resources by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:

[name, description]

@param getInternalUserQueryParams Filtering parameter
*/
func (s *InternalUserService) GetInternalUser(getInternalUserQueryParams *GetInternalUserQueryParams) (*ResponseInternalUserGetInternalUser, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/internaluser"

	queryString, _ := query.Values(getInternalUserQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseInternalUserGetInternalUser{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetInternalUser")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseInternalUserGetInternalUser)
	return result, response, err

}

//GetVersion Get internal user version information
/* This API helps to retrieve the version information related to the internal user.

 */
func (s *InternalUserService) GetVersion() (*ResponseInternalUserGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/internaluser/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseInternalUserGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseInternalUserGetVersion)
	return result, response, err

}

//CreateInternalUser Create internal user
/* This API creates an internal user.

 */
func (s *InternalUserService) CreateInternalUser(requestInternalUserCreateInternalUser *RequestInternalUserCreateInternalUser) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/internaluser"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestInternalUserCreateInternalUser).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateInternalUser")
	}

	return response, err

}

//UpdateInternalUserByName Update internal user by name
/* This API allows the client to update an internal user by name.

@param name name path parameter.
*/
func (s *InternalUserService) UpdateInternalUserByName(name string, requestInternalUserUpdateInternalUserByName *RequestInternalUserUpdateInternalUserByName) (*ResponseInternalUserUpdateInternalUserByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/internaluser/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestInternalUserUpdateInternalUserByName).
		SetResult(&ResponseInternalUserUpdateInternalUserByName{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateInternalUserByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseInternalUserUpdateInternalUserByName)
	return result, response, err

}

//UpdateInternalUserByID Update internal user by ID
/* This API allows the client to update an internal user by ID.

@param id id path parameter.
*/
func (s *InternalUserService) UpdateInternalUserByID(id string, requestInternalUserUpdateInternalUserById *RequestInternalUserUpdateInternalUserByID) (*ResponseInternalUserUpdateInternalUserByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/internaluser/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestInternalUserUpdateInternalUserById).
		SetResult(&ResponseInternalUserUpdateInternalUserByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateInternalUserById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseInternalUserUpdateInternalUserByID)
	return result, response, err

}

//DeleteInternalUserByName Delete internal user by name
/* This API deletes an internal user by name.

@param name name path parameter.
*/
func (s *InternalUserService) DeleteInternalUserByName(name string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/internaluser/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

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
		return response, fmt.Errorf("error with operation DeleteInternalUserByName")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteInternalUserByID Delete internal user by ID
/* This API deletes an internal user by ID.

@param id id path parameter.
*/
func (s *InternalUserService) DeleteInternalUserByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/internaluser/{id}"
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
		return response, fmt.Errorf("error with operation DeleteInternalUserById")
	}

	getCSFRToken(response.Header())

	return response, err

}
