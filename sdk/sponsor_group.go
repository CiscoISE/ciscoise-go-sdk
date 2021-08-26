package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SponsorGroupService service

type GetSponsorGroupQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseSponsorGroupGetSponsorGroupByID struct {
	SponsorGroup *ResponseSponsorGroupGetSponsorGroupByIDSponsorGroup `json:"SponsorGroup,omitempty"` //
}

type ResponseSponsorGroupGetSponsorGroupByIDSponsorGroup struct {
	ID                string                                                                `json:"id,omitempty"`                //
	Name              string                                                                `json:"name,omitempty"`              //
	Description       string                                                                `json:"description,omitempty"`       //
	IsEnabled         *bool                                                                 `json:"isEnabled,omitempty"`         //
	IsDefaultGroup    *bool                                                                 `json:"isDefaultGroup,omitempty"`    //
	MemberGroups      []string                                                              `json:"memberGroups,omitempty"`      //
	GuestTypes        []string                                                              `json:"guestTypes,omitempty"`        //
	Locations         []string                                                              `json:"locations,omitempty"`         //
	AutoNotification  *bool                                                                 `json:"autoNotification,omitempty"`  //
	CreatePermissions *ResponseSponsorGroupGetSponsorGroupByIDSponsorGroupCreatePermissions `json:"createPermissions,omitempty"` //
	ManagePermission  string                                                                `json:"managePermission,omitempty"`  //
	OtherPermissions  *ResponseSponsorGroupGetSponsorGroupByIDSponsorGroupOtherPermissions  `json:"otherPermissions,omitempty"`  //
	Link              *ResponseSponsorGroupGetSponsorGroupByIDSponsorGroupLink              `json:"link,omitempty"`              //
}

type ResponseSponsorGroupGetSponsorGroupByIDSponsorGroupCreatePermissions struct {
	CanImportMultipleAccounts *bool  `json:"canImportMultipleAccounts,omitempty"` //
	ImportBatchSizeLimit      *int   `json:"importBatchSizeLimit,omitempty"`      //
	CanCreateRandomAccounts   *bool  `json:"canCreateRandomAccounts,omitempty"`   //
	RandomBatchSizeLimit      *int   `json:"randomBatchSizeLimit,omitempty"`      //
	DefaultUsernamePrefix     string `json:"defaultUsernamePrefix,omitempty"`     //
	CanSpecifyUsernamePrefix  *bool  `json:"canSpecifyUsernamePrefix,omitempty"`  //
	CanSetFutureStartDate     *bool  `json:"canSetFutureStartDate,omitempty"`     //
	StartDateFutureLimitDays  *int   `json:"startDateFutureLimitDays,omitempty"`  //
}

type ResponseSponsorGroupGetSponsorGroupByIDSponsorGroupOtherPermissions struct {
	CanUpdateGuestContactInfo     *bool `json:"canUpdateGuestContactInfo,omitempty"`     //
	CanViewGuestPasswords         *bool `json:"canViewGuestPasswords,omitempty"`         //
	CanSendSmsNotifications       *bool `json:"canSendSmsNotifications,omitempty"`       //
	CanResetGuestPasswords        *bool `json:"canResetGuestPasswords,omitempty"`        //
	CanExtendGuestAccounts        *bool `json:"canExtendGuestAccounts,omitempty"`        //
	CanDeleteGuestAccounts        *bool `json:"canDeleteGuestAccounts,omitempty"`        //
	CanSuspendGuestAccounts       *bool `json:"canSuspendGuestAccounts,omitempty"`       //
	RequireSuspensionReason       *bool `json:"requireSuspensionReason,omitempty"`       //
	CanReinstateSuspendedAccounts *bool `json:"canReinstateSuspendedAccounts,omitempty"` //
	CanApproveSelfregGuests       *bool `json:"canApproveSelfregGuests,omitempty"`       //
	LimitApprovalToSponsorsGuests *bool `json:"limitApprovalToSponsorsGuests,omitempty"` //
	CanAccessViaRest              *bool `json:"canAccessViaRest,omitempty"`              //
}

type ResponseSponsorGroupGetSponsorGroupByIDSponsorGroupLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSponsorGroupUpdateSponsorGroupByID struct {
	UpdatedFieldsList *ResponseSponsorGroupUpdateSponsorGroupByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseSponsorGroupUpdateSponsorGroupByIDUpdatedFieldsList struct {
	UpdatedField *[]ResponseSponsorGroupUpdateSponsorGroupByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                     `json:"field,omitempty"`        //
	OldValue     string                                                                     `json:"oldValue,omitempty"`     //
	NewValue     string                                                                     `json:"newValue,omitempty"`     //
}

type ResponseSponsorGroupUpdateSponsorGroupByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseSponsorGroupGetSponsorGroup struct {
	SearchResult *ResponseSponsorGroupGetSponsorGroupSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseSponsorGroupGetSponsorGroupSearchResult struct {
	Total        *int                                                         `json:"total,omitempty"`        //
	Resources    *[]ResponseSponsorGroupGetSponsorGroupSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponseSponsorGroupGetSponsorGroupSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseSponsorGroupGetSponsorGroupSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseSponsorGroupGetSponsorGroupSearchResultResources struct {
	ID          string                                                        `json:"id,omitempty"`          //
	Name        string                                                        `json:"name,omitempty"`        //
	Description string                                                        `json:"description,omitempty"` //
	Link        *ResponseSponsorGroupGetSponsorGroupSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseSponsorGroupGetSponsorGroupSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSponsorGroupGetSponsorGroupSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSponsorGroupGetSponsorGroupSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSponsorGroupGetVersion struct {
	VersionInfo *ResponseSponsorGroupGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseSponsorGroupGetVersionVersionInfo struct {
	CurrentServerVersion string                                         `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                         `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseSponsorGroupGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseSponsorGroupGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestSponsorGroupUpdateSponsorGroupByID struct {
	SponsorGroup *RequestSponsorGroupUpdateSponsorGroupByIDSponsorGroup `json:"SponsorGroup,omitempty"` //
}

type RequestSponsorGroupUpdateSponsorGroupByIDSponsorGroup struct {
	ID                string                                                                  `json:"id,omitempty"`                //
	Name              string                                                                  `json:"name,omitempty"`              //
	Description       string                                                                  `json:"description,omitempty"`       //
	IsEnabled         *bool                                                                   `json:"isEnabled,omitempty"`         //
	IsDefaultGroup    *bool                                                                   `json:"isDefaultGroup,omitempty"`    //
	MemberGroups      []string                                                                `json:"memberGroups,omitempty"`      //
	GuestTypes        []string                                                                `json:"guestTypes,omitempty"`        //
	Locations         []string                                                                `json:"locations,omitempty"`         //
	AutoNotification  *bool                                                                   `json:"autoNotification,omitempty"`  //
	CreatePermissions *RequestSponsorGroupUpdateSponsorGroupByIDSponsorGroupCreatePermissions `json:"createPermissions,omitempty"` //
	ManagePermission  string                                                                  `json:"managePermission,omitempty"`  //
	OtherPermissions  *RequestSponsorGroupUpdateSponsorGroupByIDSponsorGroupOtherPermissions  `json:"otherPermissions,omitempty"`  //
}

type RequestSponsorGroupUpdateSponsorGroupByIDSponsorGroupCreatePermissions struct {
	CanImportMultipleAccounts *bool  `json:"canImportMultipleAccounts,omitempty"` //
	ImportBatchSizeLimit      *int   `json:"importBatchSizeLimit,omitempty"`      //
	CanCreateRandomAccounts   *bool  `json:"canCreateRandomAccounts,omitempty"`   //
	RandomBatchSizeLimit      *int   `json:"randomBatchSizeLimit,omitempty"`      //
	DefaultUsernamePrefix     string `json:"defaultUsernamePrefix,omitempty"`     //
	CanSpecifyUsernamePrefix  *bool  `json:"canSpecifyUsernamePrefix,omitempty"`  //
	CanSetFutureStartDate     *bool  `json:"canSetFutureStartDate,omitempty"`     //
	StartDateFutureLimitDays  *int   `json:"startDateFutureLimitDays,omitempty"`  //
}

type RequestSponsorGroupUpdateSponsorGroupByIDSponsorGroupOtherPermissions struct {
	CanUpdateGuestContactInfo     *bool `json:"canUpdateGuestContactInfo,omitempty"`     //
	CanViewGuestPasswords         *bool `json:"canViewGuestPasswords,omitempty"`         //
	CanSendSmsNotifications       *bool `json:"canSendSmsNotifications,omitempty"`       //
	CanResetGuestPasswords        *bool `json:"canResetGuestPasswords,omitempty"`        //
	CanExtendGuestAccounts        *bool `json:"canExtendGuestAccounts,omitempty"`        //
	CanDeleteGuestAccounts        *bool `json:"canDeleteGuestAccounts,omitempty"`        //
	CanSuspendGuestAccounts       *bool `json:"canSuspendGuestAccounts,omitempty"`       //
	RequireSuspensionReason       *bool `json:"requireSuspensionReason,omitempty"`       //
	CanReinstateSuspendedAccounts *bool `json:"canReinstateSuspendedAccounts,omitempty"` //
	CanApproveSelfregGuests       *bool `json:"canApproveSelfregGuests,omitempty"`       //
	LimitApprovalToSponsorsGuests *bool `json:"limitApprovalToSponsorsGuests,omitempty"` //
	CanAccessViaRest              *bool `json:"canAccessViaRest,omitempty"`              //
}

type RequestSponsorGroupCreateSponsorGroup struct {
	SponsorGroup *RequestSponsorGroupCreateSponsorGroupSponsorGroup `json:"SponsorGroup,omitempty"` //
}

type RequestSponsorGroupCreateSponsorGroupSponsorGroup struct {
	Name              string                                                              `json:"name,omitempty"`              //
	Description       string                                                              `json:"description,omitempty"`       //
	IsEnabled         *bool                                                               `json:"isEnabled,omitempty"`         //
	IsDefaultGroup    *bool                                                               `json:"isDefaultGroup,omitempty"`    //
	MemberGroups      []string                                                            `json:"memberGroups,omitempty"`      //
	GuestTypes        []string                                                            `json:"guestTypes,omitempty"`        //
	Locations         []string                                                            `json:"locations,omitempty"`         //
	AutoNotification  *bool                                                               `json:"autoNotification,omitempty"`  //
	CreatePermissions *RequestSponsorGroupCreateSponsorGroupSponsorGroupCreatePermissions `json:"createPermissions,omitempty"` //
	ManagePermission  string                                                              `json:"managePermission,omitempty"`  //
	OtherPermissions  *RequestSponsorGroupCreateSponsorGroupSponsorGroupOtherPermissions  `json:"otherPermissions,omitempty"`  //
}

type RequestSponsorGroupCreateSponsorGroupSponsorGroupCreatePermissions struct {
	CanImportMultipleAccounts *bool  `json:"canImportMultipleAccounts,omitempty"` //
	ImportBatchSizeLimit      *int   `json:"importBatchSizeLimit,omitempty"`      //
	CanCreateRandomAccounts   *bool  `json:"canCreateRandomAccounts,omitempty"`   //
	RandomBatchSizeLimit      *int   `json:"randomBatchSizeLimit,omitempty"`      //
	DefaultUsernamePrefix     string `json:"defaultUsernamePrefix,omitempty"`     //
	CanSpecifyUsernamePrefix  *bool  `json:"canSpecifyUsernamePrefix,omitempty"`  //
	CanSetFutureStartDate     *bool  `json:"canSetFutureStartDate,omitempty"`     //
	StartDateFutureLimitDays  *int   `json:"startDateFutureLimitDays,omitempty"`  //
}

type RequestSponsorGroupCreateSponsorGroupSponsorGroupOtherPermissions struct {
	CanUpdateGuestContactInfo     *bool `json:"canUpdateGuestContactInfo,omitempty"`     //
	CanViewGuestPasswords         *bool `json:"canViewGuestPasswords,omitempty"`         //
	CanSendSmsNotifications       *bool `json:"canSendSmsNotifications,omitempty"`       //
	CanResetGuestPasswords        *bool `json:"canResetGuestPasswords,omitempty"`        //
	CanExtendGuestAccounts        *bool `json:"canExtendGuestAccounts,omitempty"`        //
	CanDeleteGuestAccounts        *bool `json:"canDeleteGuestAccounts,omitempty"`        //
	CanSuspendGuestAccounts       *bool `json:"canSuspendGuestAccounts,omitempty"`       //
	RequireSuspensionReason       *bool `json:"requireSuspensionReason,omitempty"`       //
	CanReinstateSuspendedAccounts *bool `json:"canReinstateSuspendedAccounts,omitempty"` //
	CanApproveSelfregGuests       *bool `json:"canApproveSelfregGuests,omitempty"`       //
	LimitApprovalToSponsorsGuests *bool `json:"limitApprovalToSponsorsGuests,omitempty"` //
	CanAccessViaRest              *bool `json:"canAccessViaRest,omitempty"`              //
}

//GetSponsorGroupByID Get sponsor group by ID
/* This API allows the client to get a sponsor group by ID.

@param id id path parameter.
*/
func (s *SponsorGroupService) GetSponsorGroupByID(id string) (*ResponseSponsorGroupGetSponsorGroupByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsorgroup/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSponsorGroupGetSponsorGroupByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSponsorGroupById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSponsorGroupGetSponsorGroupByID)
	return result, response, err

}

//GetSponsorGroup Get all sponsor groups
/* This API allows the client to get all the sponsor groups.

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

@param getSponsorGroupQueryParams Filtering parameter
*/
func (s *SponsorGroupService) GetSponsorGroup(getSponsorGroupQueryParams *GetSponsorGroupQueryParams) (*ResponseSponsorGroupGetSponsorGroup, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsorgroup"

	queryString, _ := query.Values(getSponsorGroupQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSponsorGroupGetSponsorGroup{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSponsorGroup")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSponsorGroupGetSponsorGroup)
	return result, response, err

}

//GetVersion Get sponsor group version information
/* This API helps to retrieve the version information related to the sponsor group.

 */
func (s *SponsorGroupService) GetVersion() (*ResponseSponsorGroupGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsorgroup/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSponsorGroupGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSponsorGroupGetVersion)
	return result, response, err

}

//CreateSponsorGroup Create sponsor group
/* This API creates a sponsor group.

 */
func (s *SponsorGroupService) CreateSponsorGroup(requestSponsorGroupCreateSponsorGroup *RequestSponsorGroupCreateSponsorGroup) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsorgroup"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSponsorGroupCreateSponsorGroup).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateSponsorGroup")
	}

	return response, err

}

//UpdateSponsorGroupByID Update sponsor group by ID
/* This API allows the client to update a sponsor group by ID.

@param id id path parameter.
*/
func (s *SponsorGroupService) UpdateSponsorGroupByID(id string, requestSponsorGroupUpdateSponsorGroupById *RequestSponsorGroupUpdateSponsorGroupByID) (*ResponseSponsorGroupUpdateSponsorGroupByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsorgroup/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSponsorGroupUpdateSponsorGroupById).
		SetResult(&ResponseSponsorGroupUpdateSponsorGroupByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSponsorGroupById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSponsorGroupUpdateSponsorGroupByID)
	return result, response, err

}

//DeleteSponsorGroupByID Delete sponsor group by ID
/* This API deletes a sponsor group by ID.

@param id id path parameter.
*/
func (s *SponsorGroupService) DeleteSponsorGroupByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsorgroup/{id}"
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
		return response, fmt.Errorf("error with operation DeleteSponsorGroupById")
	}

	getCSFRToken(response.Header())

	return response, err

}
