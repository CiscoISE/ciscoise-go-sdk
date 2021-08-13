package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type GuestUserService service

type GetGuestUsersQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseGuestUserGetGuestUserByName struct {
	GuestUser ResponseGuestUserGetGuestUserByNameGuestUser `json:"GuestUser,omitempty"` //
}

type ResponseGuestUserGetGuestUserByNameGuestUser struct {
	ID              string                                                      `json:"id,omitempty"`              //
	Name            string                                                      `json:"name,omitempty"`            //
	Description     string                                                      `json:"description,omitempty"`     //
	GuestType       string                                                      `json:"guestType,omitempty"`       //
	Status          string                                                      `json:"status,omitempty"`          //
	StatusReason    string                                                      `json:"statusReason,omitempty"`    //
	ReasonForVisit  string                                                      `json:"reasonForVisit,omitempty"`  //
	SponsorUserID   string                                                      `json:"sponsorUserId,omitempty"`   //
	SponsorUserName string                                                      `json:"sponsorUserName,omitempty"` //
	GuestInfo       ResponseGuestUserGetGuestUserByNameGuestUserGuestInfo       `json:"guestInfo,omitempty"`       //
	GuestAccessInfo ResponseGuestUserGetGuestUserByNameGuestUserGuestAccessInfo `json:"guestAccessInfo,omitempty"` //
	PortalID        string                                                      `json:"portalId,omitempty"`        //
	CustomFields    map[string]interface{}                                      `json:"customFields,omitempty"`    // Key value map
	Link            ResponseGuestUserGetGuestUserByNameGuestUserLink            `json:"link,omitempty"`            //
}

type ResponseGuestUserGetGuestUserByNameGuestUserGuestInfo struct {
	FirstName            string `json:"firstName,omitempty"`            //
	LastName             string `json:"lastName,omitempty"`             //
	Company              string `json:"company,omitempty"`              //
	CreationTime         string `json:"creationTime,omitempty"`         //
	NotificationLanguage string `json:"notificationLanguage,omitempty"` //
	UserName             string `json:"userName,omitempty"`             // If account needs be created with mobile number, please provide mobile number here
	EmailAddress         string `json:"emailAddress,omitempty"`         //
	PhoneNumber          string `json:"phoneNumber,omitempty"`          // Phone number should be E.164 format
	Password             string `json:"password,omitempty"`             //
	Enabled              bool   `json:"enabled,omitempty"`              // This field is only for Get operation not applicable for Create, Update operations
	SmsServiceProvider   string `json:"smsServiceProvider,omitempty"`   //
}

type ResponseGuestUserGetGuestUserByNameGuestUserGuestAccessInfo struct {
	ValidDays int    `json:"validDays,omitempty"` //
	FromDate  string `json:"fromDate,omitempty"`  //
	ToDate    string `json:"toDate,omitempty"`    //
	Location  string `json:"location,omitempty"`  //
	SSID      string `json:"ssid,omitempty"`      //
	GroupTag  string `json:"groupTag,omitempty"`  //
}

type ResponseGuestUserGetGuestUserByNameGuestUserLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestUserUpdateGuestUserByName struct {
	UpdatedFieldsList ResponseGuestUserUpdateGuestUserByNameUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseGuestUserUpdateGuestUserByNameUpdatedFieldsList struct {
	UpdatedField ResponseGuestUserUpdateGuestUserByNameUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                              `json:"field,omitempty"`        //
	OldValue     string                                                              `json:"oldValue,omitempty"`     //
	NewValue     string                                                              `json:"newValue,omitempty"`     //
}

type ResponseGuestUserUpdateGuestUserByNameUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseGuestUserResetGuestUserPasswordByID struct {
	OperationResult ResponseGuestUserResetGuestUserPasswordByIDOperationResult `json:"OperationResult,omitempty"` //
}

type ResponseGuestUserResetGuestUserPasswordByIDOperationResult struct {
	ResultValue []ResponseGuestUserResetGuestUserPasswordByIDOperationResultResultValue `json:"resultValue,omitempty"` //
}

type ResponseGuestUserResetGuestUserPasswordByIDOperationResultResultValue struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type ResponseGuestUserGetGuestUserByID struct {
	GuestUser ResponseGuestUserGetGuestUserByIDGuestUser `json:"GuestUser,omitempty"` //
}

type ResponseGuestUserGetGuestUserByIDGuestUser struct {
	ID              string                                                    `json:"id,omitempty"`              //
	Name            string                                                    `json:"name,omitempty"`            //
	Description     string                                                    `json:"description,omitempty"`     //
	GuestType       string                                                    `json:"guestType,omitempty"`       //
	Status          string                                                    `json:"status,omitempty"`          //
	StatusReason    string                                                    `json:"statusReason,omitempty"`    //
	ReasonForVisit  string                                                    `json:"reasonForVisit,omitempty"`  //
	SponsorUserID   string                                                    `json:"sponsorUserId,omitempty"`   //
	SponsorUserName string                                                    `json:"sponsorUserName,omitempty"` //
	GuestInfo       ResponseGuestUserGetGuestUserByIDGuestUserGuestInfo       `json:"guestInfo,omitempty"`       //
	GuestAccessInfo ResponseGuestUserGetGuestUserByIDGuestUserGuestAccessInfo `json:"guestAccessInfo,omitempty"` //
	PortalID        string                                                    `json:"portalId,omitempty"`        //
	CustomFields    map[string]interface{}                                    `json:"customFields,omitempty"`    // Key value map
	Link            ResponseGuestUserGetGuestUserByIDGuestUserLink            `json:"link,omitempty"`            //
}

type ResponseGuestUserGetGuestUserByIDGuestUserGuestInfo struct {
	FirstName            string `json:"firstName,omitempty"`            //
	LastName             string `json:"lastName,omitempty"`             //
	Company              string `json:"company,omitempty"`              //
	CreationTime         string `json:"creationTime,omitempty"`         //
	NotificationLanguage string `json:"notificationLanguage,omitempty"` //
	UserName             string `json:"userName,omitempty"`             // If account needs be created with mobile number, please provide mobile number here
	EmailAddress         string `json:"emailAddress,omitempty"`         //
	PhoneNumber          string `json:"phoneNumber,omitempty"`          // Phone number should be E.164 format
	Password             string `json:"password,omitempty"`             //
	Enabled              bool   `json:"enabled,omitempty"`              // This field is only for Get operation not applicable for Create, Update operations
	SmsServiceProvider   string `json:"smsServiceProvider,omitempty"`   //
}

type ResponseGuestUserGetGuestUserByIDGuestUserGuestAccessInfo struct {
	ValidDays int    `json:"validDays,omitempty"` //
	FromDate  string `json:"fromDate,omitempty"`  //
	ToDate    string `json:"toDate,omitempty"`    //
	Location  string `json:"location,omitempty"`  //
	SSID      string `json:"ssid,omitempty"`      //
	GroupTag  string `json:"groupTag,omitempty"`  //
}

type ResponseGuestUserGetGuestUserByIDGuestUserLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestUserUpdateGuestUserByID struct {
	UpdatedFieldsList ResponseGuestUserUpdateGuestUserByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseGuestUserUpdateGuestUserByIDUpdatedFieldsList struct {
	UpdatedField ResponseGuestUserUpdateGuestUserByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                            `json:"field,omitempty"`        //
	OldValue     string                                                            `json:"oldValue,omitempty"`     //
	NewValue     string                                                            `json:"newValue,omitempty"`     //
}

type ResponseGuestUserUpdateGuestUserByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseGuestUserGetGuestUsers struct {
	SearchResult ResponseGuestUserGetGuestUsersSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseGuestUserGetGuestUsersSearchResult struct {
	Total        int                                                    `json:"total,omitempty"`        //
	Resources    []ResponseGuestUserGetGuestUsersSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponseGuestUserGetGuestUsersSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponseGuestUserGetGuestUsersSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseGuestUserGetGuestUsersSearchResultResources struct {
	ID          string                                                  `json:"id,omitempty"`          //
	Name        string                                                  `json:"name,omitempty"`        //
	Description string                                                  `json:"description,omitempty"` //
	Link        ResponseGuestUserGetGuestUsersSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseGuestUserGetGuestUsersSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestUserGetGuestUsersSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestUserGetGuestUsersSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestUserGetVersion struct {
	VersionInfo ResponseGuestUserGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseGuestUserGetVersionVersionInfo struct {
	CurrentServerVersion string                                     `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                     `json:"supportedVersions,omitempty"`    //
	Link                 ResponseGuestUserGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseGuestUserGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestUserMonitorBulkStatusGuestUser struct {
	BulkStatus ResponseGuestUserMonitorBulkStatusGuestUserBulkStatus `json:"BulkStatus,omitempty"` //
}

type ResponseGuestUserMonitorBulkStatusGuestUserBulkStatus struct {
	BulkID          string                                                                 `json:"bulkId,omitempty"`          //
	MediaType       string                                                                 `json:"mediaType,omitempty"`       //
	ExecutionStatus string                                                                 `json:"executionStatus,omitempty"` //
	OperationType   string                                                                 `json:"operationType,omitempty"`   //
	StartTime       string                                                                 `json:"startTime,omitempty"`       //
	ResourcesCount  int                                                                    `json:"resourcesCount,omitempty"`  //
	SuccessCount    int                                                                    `json:"successCount,omitempty"`    //
	FailCount       int                                                                    `json:"failCount,omitempty"`       //
	ResourcesStatus []ResponseGuestUserMonitorBulkStatusGuestUserBulkStatusResourcesStatus `json:"resourcesStatus,omitempty"` //
}

type ResponseGuestUserMonitorBulkStatusGuestUserBulkStatusResourcesStatus struct {
	ID                      string `json:"id,omitempty"`                      //
	Name                    string `json:"name,omitempty"`                    //
	Description             string `json:"description,omitempty"`             //
	ResourceExecutionStatus string `json:"resourceExecutionStatus,omitempty"` //
	Status                  string `json:"status,omitempty"`                  //
}

type RequestGuestUserChangeSponsorPassword struct {
	OperationAdditionalData *RequestGuestUserChangeSponsorPasswordOperationAdditionalData `json:"OperationAdditionalData,omitempty"` //
}

type RequestGuestUserChangeSponsorPasswordOperationAdditionalData struct {
	AdditionalData *[]RequestGuestUserChangeSponsorPasswordOperationAdditionalDataAdditionalData `json:"additionalData,omitempty"` //
}

type RequestGuestUserChangeSponsorPasswordOperationAdditionalDataAdditionalData struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type RequestGuestUserUpdateGuestUserByName struct {
	GuestUser *RequestGuestUserUpdateGuestUserByNameGuestUser `json:"GuestUser,omitempty"` //
}

type RequestGuestUserUpdateGuestUserByNameGuestUser struct {
	ID              string                                                         `json:"id,omitempty"`              //
	Name            string                                                         `json:"name,omitempty"`            //
	Description     string                                                         `json:"description,omitempty"`     //
	GuestType       string                                                         `json:"guestType,omitempty"`       //
	Status          string                                                         `json:"status,omitempty"`          //
	StatusReason    string                                                         `json:"statusReason,omitempty"`    //
	ReasonForVisit  string                                                         `json:"reasonForVisit,omitempty"`  //
	SponsorUserID   string                                                         `json:"sponsorUserId,omitempty"`   //
	SponsorUserName string                                                         `json:"sponsorUserName,omitempty"` //
	GuestInfo       *RequestGuestUserUpdateGuestUserByNameGuestUserGuestInfo       `json:"guestInfo,omitempty"`       //
	GuestAccessInfo *RequestGuestUserUpdateGuestUserByNameGuestUserGuestAccessInfo `json:"guestAccessInfo,omitempty"` //
	PortalID        string                                                         `json:"portalId,omitempty"`        //
	CustomFields    *map[string]interface{}                                        `json:"customFields,omitempty"`    // Key value map
}

type RequestGuestUserUpdateGuestUserByNameGuestUserGuestInfo struct {
	FirstName            string `json:"firstName,omitempty"`            //
	LastName             string `json:"lastName,omitempty"`             //
	Company              string `json:"company,omitempty"`              //
	CreationTime         string `json:"creationTime,omitempty"`         //
	NotificationLanguage string `json:"notificationLanguage,omitempty"` //
	UserName             string `json:"userName,omitempty"`             // If account needs be created with mobile number, please provide mobile number here
	EmailAddress         string `json:"emailAddress,omitempty"`         //
	PhoneNumber          string `json:"phoneNumber,omitempty"`          // Phone number should be E.164 format
	Password             string `json:"password,omitempty"`             //
	Enabled              *bool  `json:"enabled,omitempty"`              // This field is only for Get operation not applicable for Create, Update operations
	SmsServiceProvider   string `json:"smsServiceProvider,omitempty"`   //
}

type RequestGuestUserUpdateGuestUserByNameGuestUserGuestAccessInfo struct {
	ValidDays *int   `json:"validDays,omitempty"` //
	FromDate  string `json:"fromDate,omitempty"`  //
	ToDate    string `json:"toDate,omitempty"`    //
	Location  string `json:"location,omitempty"`  //
	SSID      string `json:"ssid,omitempty"`      //
	GroupTag  string `json:"groupTag,omitempty"`  //
}

type RequestGuestUserUpdateGuestUserEmail struct {
	OperationAdditionalData *RequestGuestUserUpdateGuestUserEmailOperationAdditionalData `json:"OperationAdditionalData,omitempty"` //
}

type RequestGuestUserUpdateGuestUserEmailOperationAdditionalData struct {
	AdditionalData *[]RequestGuestUserUpdateGuestUserEmailOperationAdditionalDataAdditionalData `json:"additionalData,omitempty"` //
}

type RequestGuestUserUpdateGuestUserEmailOperationAdditionalDataAdditionalData struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type RequestGuestUserUpdateGuestUserByID struct {
	GuestUser *RequestGuestUserUpdateGuestUserByIDGuestUser `json:"GuestUser,omitempty"` //
}

type RequestGuestUserUpdateGuestUserByIDGuestUser struct {
	ID              string                                                       `json:"id,omitempty"`              //
	Name            string                                                       `json:"name,omitempty"`            //
	Description     string                                                       `json:"description,omitempty"`     //
	GuestType       string                                                       `json:"guestType,omitempty"`       //
	Status          string                                                       `json:"status,omitempty"`          //
	StatusReason    string                                                       `json:"statusReason,omitempty"`    //
	ReasonForVisit  string                                                       `json:"reasonForVisit,omitempty"`  //
	SponsorUserID   string                                                       `json:"sponsorUserId,omitempty"`   //
	SponsorUserName string                                                       `json:"sponsorUserName,omitempty"` //
	GuestInfo       *RequestGuestUserUpdateGuestUserByIDGuestUserGuestInfo       `json:"guestInfo,omitempty"`       //
	GuestAccessInfo *RequestGuestUserUpdateGuestUserByIDGuestUserGuestAccessInfo `json:"guestAccessInfo,omitempty"` //
	PortalID        string                                                       `json:"portalId,omitempty"`        //
	CustomFields    *map[string]interface{}                                      `json:"customFields,omitempty"`    // Key value map
}

type RequestGuestUserUpdateGuestUserByIDGuestUserGuestInfo struct {
	FirstName            string `json:"firstName,omitempty"`            //
	LastName             string `json:"lastName,omitempty"`             //
	Company              string `json:"company,omitempty"`              //
	CreationTime         string `json:"creationTime,omitempty"`         //
	NotificationLanguage string `json:"notificationLanguage,omitempty"` //
	UserName             string `json:"userName,omitempty"`             // If account needs be created with mobile number, please provide mobile number here
	EmailAddress         string `json:"emailAddress,omitempty"`         //
	PhoneNumber          string `json:"phoneNumber,omitempty"`          // Phone number should be E.164 format
	Password             string `json:"password,omitempty"`             //
	Enabled              *bool  `json:"enabled,omitempty"`              // This field is only for Get operation not applicable for Create, Update operations
	SmsServiceProvider   string `json:"smsServiceProvider,omitempty"`   //
}

type RequestGuestUserUpdateGuestUserByIDGuestUserGuestAccessInfo struct {
	ValidDays *int   `json:"validDays,omitempty"` //
	FromDate  string `json:"fromDate,omitempty"`  //
	ToDate    string `json:"toDate,omitempty"`    //
	Location  string `json:"location,omitempty"`  //
	SSID      string `json:"ssid,omitempty"`      //
	GroupTag  string `json:"groupTag,omitempty"`  //
}

type RequestGuestUserCreateGuestUser struct {
	GuestUser *RequestGuestUserCreateGuestUserGuestUser `json:"GuestUser,omitempty"` //
}

type RequestGuestUserCreateGuestUserGuestUser struct {
	Name            string                                                   `json:"name,omitempty"`            //
	Description     string                                                   `json:"description,omitempty"`     //
	GuestType       string                                                   `json:"guestType,omitempty"`       //
	Status          string                                                   `json:"status,omitempty"`          //
	StatusReason    string                                                   `json:"statusReason,omitempty"`    //
	ReasonForVisit  string                                                   `json:"reasonForVisit,omitempty"`  //
	SponsorUserID   string                                                   `json:"sponsorUserId,omitempty"`   //
	SponsorUserName string                                                   `json:"sponsorUserName,omitempty"` //
	GuestInfo       *RequestGuestUserCreateGuestUserGuestUserGuestInfo       `json:"guestInfo,omitempty"`       //
	GuestAccessInfo *RequestGuestUserCreateGuestUserGuestUserGuestAccessInfo `json:"guestAccessInfo,omitempty"` //
	PortalID        string                                                   `json:"portalId,omitempty"`        //
	CustomFields    *map[string]interface{}                                  `json:"customFields,omitempty"`    // Key value map
}

type RequestGuestUserCreateGuestUserGuestUserGuestInfo struct {
	FirstName            string `json:"firstName,omitempty"`            //
	LastName             string `json:"lastName,omitempty"`             //
	Company              string `json:"company,omitempty"`              //
	CreationTime         string `json:"creationTime,omitempty"`         //
	NotificationLanguage string `json:"notificationLanguage,omitempty"` //
	UserName             string `json:"userName,omitempty"`             // If account needs be created with mobile number, please provide mobile number here
	EmailAddress         string `json:"emailAddress,omitempty"`         //
	PhoneNumber          string `json:"phoneNumber,omitempty"`          // Phone number should be E.164 format
	Password             string `json:"password,omitempty"`             //
	Enabled              *bool  `json:"enabled,omitempty"`              // This field is only for Get operation not applicable for Create, Update operations
	SmsServiceProvider   string `json:"smsServiceProvider,omitempty"`   //
}

type RequestGuestUserCreateGuestUserGuestUserGuestAccessInfo struct {
	ValidDays *int   `json:"validDays,omitempty"` //
	FromDate  string `json:"fromDate,omitempty"`  //
	ToDate    string `json:"toDate,omitempty"`    //
	Location  string `json:"location,omitempty"`  //
	SSID      string `json:"ssid,omitempty"`      //
	GroupTag  string `json:"groupTag,omitempty"`  //
}

type RequestGuestUserSuspendGuestUserByID struct {
	OperationAdditionalData *RequestGuestUserSuspendGuestUserByIDOperationAdditionalData `json:"OperationAdditionalData,omitempty"` //
}

type RequestGuestUserSuspendGuestUserByIDOperationAdditionalData struct {
	AdditionalData *[]RequestGuestUserSuspendGuestUserByIDOperationAdditionalDataAdditionalData `json:"additionalData,omitempty"` //
}

type RequestGuestUserSuspendGuestUserByIDOperationAdditionalDataAdditionalData struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type RequestGuestUserBulkRequestForGuestUser struct {
	GuestUserBulkRequest *RequestGuestUserBulkRequestForGuestUserGuestUserBulkRequest `json:"GuestUserBulkRequest,omitempty"` //
}

type RequestGuestUserBulkRequestForGuestUserGuestUserBulkRequest struct {
	OperationType     string `json:"operationType,omitempty"`     //
	ResourceMediaType string `json:"resourceMediaType,omitempty"` //
}

//GetGuestUserByName Get guest user by name
/* This API allows the client to get a guest user by name.

@param name name path parameter.
*/
func (s *GuestUserService) GetGuestUserByName(name string) (*ResponseGuestUserGetGuestUserByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseGuestUserGetGuestUserByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetGuestUserByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestUserGetGuestUserByName)
	return result, response, err

}

//GetGuestUserByID Get guest user by ID
/* This API allows the client to get a guest user by ID.

@param id id path parameter.
*/
func (s *GuestUserService) GetGuestUserByID(id string) (*ResponseGuestUserGetGuestUserByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseGuestUserGetGuestUserByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetGuestUserById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestUserGetGuestUserByID)
	return result, response, err

}

//GetGuestUsers Get all guest users
/* This API allows the client to get all the guest users.

Filter:

[lastName, sponsor, creationTime, personBeingVisited, toDate, userName, firstName, emailAddress, phoneNumber, groupTag, name, company, guestType, status]

To search resources by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:

[firstName, lastName, emailAddress, name, description]

@param getGuestUsersQueryParams Filtering parameter
*/
func (s *GuestUserService) GetGuestUsers(getGuestUsersQueryParams *GetGuestUsersQueryParams) (*ResponseGuestUserGetGuestUsers, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser"

	queryString, _ := query.Values(getGuestUsersQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseGuestUserGetGuestUsers{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetGuestUsers")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestUserGetGuestUsers)
	return result, response, err

}

//GetVersion Get guest user version information
/* This API helps to retrieve the version information related to the guest user.

 */
func (s *GuestUserService) GetVersion() (*ResponseGuestUserGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseGuestUserGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestUserGetVersion)
	return result, response, err

}

//MonitorBulkStatusGuestUser Monitor bulk request
/* This API allows the client to monitor the bulk request.

@param bulkid bulkid path parameter.
*/
func (s *GuestUserService) MonitorBulkStatusGuestUser(bulkid string) (*ResponseGuestUserMonitorBulkStatusGuestUser, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/bulk/{bulkid}"
	path = strings.Replace(path, "{bulkid}", fmt.Sprintf("%v", bulkid), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseGuestUserMonitorBulkStatusGuestUser{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MonitorBulkStatusGuestUser")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestUserMonitorBulkStatusGuestUser)
	return result, response, err

}

//CreateGuestUser Create guest user
/* This API creates a guest user.

 */
func (s *GuestUserService) CreateGuestUser(requestGuestUserCreateGuestUser *RequestGuestUserCreateGuestUser) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestGuestUserCreateGuestUser).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateGuestUser")
	}

	return response, err

}

//ApproveGuestUserByID Approve guest user by ID
/* This API allows the client to approve a guest user by ID.

@param id id path parameter.
*/
func (s *GuestUserService) ApproveGuestUserByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/approve/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

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
		return response, fmt.Errorf("error with operation ApproveGuestUserById")
	}

	getCSFRToken(response.Header())

	return response, err

}

//ChangeSponsorPassword Change sponsor password
/* This API allows the client to change the sponsor password.

@param portalID portalId path parameter.
*/
func (s *GuestUserService) ChangeSponsorPassword(portalID string, requestGuestUserChangeSponsorPassword *RequestGuestUserChangeSponsorPassword) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/changeSponsorPassword/{portalId}"
	path = strings.Replace(path, "{portalId}", fmt.Sprintf("%v", portalID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestGuestUserChangeSponsorPassword).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation ChangeSponsorPassword")
	}

	getCSFRToken(response.Header())

	return response, err

}

//SuspendGuestUserByName Suspend guest user by name
/* This API allows the client to suspend a guest user by name.

@param name name path parameter.
*/
func (s *GuestUserService) SuspendGuestUserByName(name string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/suspend/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

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
		return response, fmt.Errorf("error with operation SuspendGuestUserByName")
	}

	getCSFRToken(response.Header())

	return response, err

}

//ReinstateGuestUserByName Reinstate guest user by name
/* This API allows the client to reinstate a guest user by name.

@param name name path parameter.
*/
func (s *GuestUserService) ReinstateGuestUserByName(name string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/reinstate/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

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
		return response, fmt.Errorf("error with operation ReinstateGuestUserByName")
	}

	getCSFRToken(response.Header())

	return response, err

}

//UpdateGuestUserByName Update guest user by name
/* This API allows the client to update a guest user by name.

@param name name path parameter.
*/
func (s *GuestUserService) UpdateGuestUserByName(name string, requestGuestUserUpdateGuestUserByName *RequestGuestUserUpdateGuestUserByName) (*ResponseGuestUserUpdateGuestUserByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestGuestUserUpdateGuestUserByName).
		SetResult(&ResponseGuestUserUpdateGuestUserByName{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateGuestUserByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestUserUpdateGuestUserByName)
	return result, response, err

}

//ResetGuestUserPasswordByID Reset guest user password
/* This API allows the client to reset the guest user password.

@param id id path parameter.
*/
func (s *GuestUserService) ResetGuestUserPasswordByID(id string) (*ResponseGuestUserResetGuestUserPasswordByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/resetpassword/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseGuestUserResetGuestUserPasswordByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ResetGuestUserPasswordById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestUserResetGuestUserPasswordByID)
	return result, response, err

}

//ReinstateGuestUserByID Reinstate guest user by ID
/* This API allows the client to reinstate a guest user by ID.

@param id id path parameter.
*/
func (s *GuestUserService) ReinstateGuestUserByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/reinstate/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

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
		return response, fmt.Errorf("error with operation ReinstateGuestUserById")
	}

	getCSFRToken(response.Header())

	return response, err

}

//UpdateGuestUserEmail Update guest user email by ID
/* This API allows the client to update a guest user email by ID.

@param id id path parameter.
@param portalID portalId path parameter.
*/
func (s *GuestUserService) UpdateGuestUserEmail(id string, portalID string, requestGuestUserUpdateGuestUserEmail *RequestGuestUserUpdateGuestUserEmail) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/email/{id}/portalId/{portalId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{portalId}", fmt.Sprintf("%v", portalID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestGuestUserUpdateGuestUserEmail).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateGuestUserEmail")
	}

	getCSFRToken(response.Header())

	return response, err

}

//UpdateGuestUserSms Update guest user sms by ID
/* This API allows the client to update a guest user sms by ID.

@param id id path parameter.
@param portalID portalId path parameter.
*/
func (s *GuestUserService) UpdateGuestUserSms(id string, portalID string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/sms/{id}/portalId/{portalId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{portalId}", fmt.Sprintf("%v", portalID), -1)

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
		return response, fmt.Errorf("error with operation UpdateGuestUserSms")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DenyGuestUserByID Deny guest user by ID
/* This API allows the client to deny a guest user by ID.

@param id id path parameter.
*/
func (s *GuestUserService) DenyGuestUserByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/deny/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

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
		return response, fmt.Errorf("error with operation DenyGuestUserById")
	}

	getCSFRToken(response.Header())

	return response, err

}

//UpdateGuestUserByID Update guest user by ID
/* This API allows the client to update a guest user by ID.

@param id id path parameter.
*/
func (s *GuestUserService) UpdateGuestUserByID(id string, requestGuestUserUpdateGuestUserById *RequestGuestUserUpdateGuestUserByID) (*ResponseGuestUserUpdateGuestUserByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestGuestUserUpdateGuestUserById).
		SetResult(&ResponseGuestUserUpdateGuestUserByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateGuestUserById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestUserUpdateGuestUserByID)
	return result, response, err

}

//SuspendGuestUserByID Suspend guest user by ID
/* This API allows the client to suspend a guest user by ID.

@param id id path parameter.
*/
func (s *GuestUserService) SuspendGuestUserByID(id string, requestGuestUserSuspendGuestUserById *RequestGuestUserSuspendGuestUserByID) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/suspend/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestGuestUserSuspendGuestUserById).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation SuspendGuestUserById")
	}

	getCSFRToken(response.Header())

	return response, err

}

//BulkRequestForGuestUser Submit bulk request
/* This API allows the client to submit the bulk request.

 */
func (s *GuestUserService) BulkRequestForGuestUser(requestGuestUserBulkRequestForGuestUser *RequestGuestUserBulkRequestForGuestUser) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/bulk/submit"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestGuestUserBulkRequestForGuestUser).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation BulkRequestForGuestUser")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteGuestUserByName Delete guest user
/* This API deletes a guest user.

@param name name path parameter.
*/
func (s *GuestUserService) DeleteGuestUserByName(name string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/name/{name}"
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
		return response, fmt.Errorf("error with operation DeleteGuestUserByName")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteGuestUserByID Delete guest user by ID
/* This API deletes a guest user by ID.

@param id id path parameter.
*/
func (s *GuestUserService) DeleteGuestUserByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestuser/{id}"
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
		return response, fmt.Errorf("error with operation DeleteGuestUserById")
	}

	getCSFRToken(response.Header())

	return response, err

}
