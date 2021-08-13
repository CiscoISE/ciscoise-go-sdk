package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type GuestTypeService service

type GetGuestTypeQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseGuestTypeGetGuestTypeByID struct {
	GuestType ResponseGuestTypeGetGuestTypeByIDGuestType `json:"GuestType,omitempty"` //
}

type ResponseGuestTypeGetGuestTypeByIDGuestType struct {
	ID                     string                                                           `json:"id,omitempty"`                     //
	Name                   string                                                           `json:"name,omitempty"`                   //
	Description            string                                                           `json:"description,omitempty"`            //
	IsDefaultType          bool                                                             `json:"isDefaultType,omitempty"`          //
	AccessTime             ResponseGuestTypeGetGuestTypeByIDGuestTypeAccessTime             `json:"accessTime,omitempty"`             //
	LoginOptions           ResponseGuestTypeGetGuestTypeByIDGuestTypeLoginOptions           `json:"loginOptions,omitempty"`           //
	ExpirationNotification ResponseGuestTypeGetGuestTypeByIDGuestTypeExpirationNotification `json:"expirationNotification,omitempty"` // Expiration Notification Settings
	SponsorGroups          []string                                                         `json:"sponsorGroups,omitempty"`          //
	Link                   ResponseGuestTypeGetGuestTypeByIDGuestTypeLink                   `json:"link,omitempty"`                   //
}

type ResponseGuestTypeGetGuestTypeByIDGuestTypeAccessTime struct {
	FromFirstLogin                 bool                                                                `json:"fromFirstLogin,omitempty"`                 // When Account Duration starts from first login or specified date
	MaxAccountDuration             int                                                                 `json:"maxAccountDuration,omitempty"`             // Maximum value of Account Duration
	DurationTimeUnit               string                                                              `json:"durationTimeUnit,omitempty"`               // Allowed values are: - DAYS, - HOURS, - MINUTES
	DefaultDuration                int                                                                 `json:"defaultDuration,omitempty"`                //
	AllowAccessOnSpecificDaysTimes bool                                                                `json:"allowAccessOnSpecificDaysTimes,omitempty"` //
	DayTimeLimits                  []ResponseGuestTypeGetGuestTypeByIDGuestTypeAccessTimeDayTimeLimits `json:"dayTimeLimits,omitempty"`                  // List of Time Ranges for account access
}

type ResponseGuestTypeGetGuestTypeByIDGuestTypeAccessTimeDayTimeLimits struct {
	StartTime string   `json:"startTime,omitempty"` // Start time in HH:mm format
	EndTime   string   `json:"endTime,omitempty"`   // End time in HH:mm format
	Days      []string `json:"days,omitempty"`      // List of Days Values should be one of Week day. Allowed values are: - Sunday, - Monday, - Tuesday, - Wednesday, - Thursday, - Friday, - Saturday
}

type ResponseGuestTypeGetGuestTypeByIDGuestTypeLoginOptions struct {
	LimitSimultaneousLogins bool   `json:"limitSimultaneousLogins,omitempty"` // Enable Simultaneous Logins
	MaxSimultaneousLogins   int    `json:"maxSimultaneousLogins,omitempty"`   // Number of Simultaneous Logins
	FailureAction           string `json:"failureAction,omitempty"`           // When Guest Exceeds limit this action will be invoked. Allowed values are: - Disconnect_Oldest_Connection, - Disconnect_Newest_Connection
	MaxRegisteredDevices    int    `json:"maxRegisteredDevices,omitempty"`    // Maximum devices guests can register
	IDentityGroupID         string `json:"identityGroupId,omitempty"`         //
	AllowGuestPortalBypass  bool   `json:"allowGuestPortalBypass,omitempty"`  //
}

type ResponseGuestTypeGetGuestTypeByIDGuestTypeExpirationNotification struct {
	EnableNotification          bool   `json:"enableNotification,omitempty"`          // Enable Notification settings
	AdvanceNotificationDuration int    `json:"advanceNotificationDuration,omitempty"` // Send Account Expiration Notification Duration before ( Days, Hours, Minutes )
	AdvanceNotificationUnits    string `json:"advanceNotificationUnits,omitempty"`    // Allowed values are: - DAYS, - HOURS, - MINUTES
	SendEmailNotification       bool   `json:"sendEmailNotification,omitempty"`       // Enable Email Notification
	EmailText                   string `json:"emailText,omitempty"`                   //
	SendSmsNotification         bool   `json:"sendSmsNotification,omitempty"`         // Maximum devices guests can register
	SmsText                     string `json:"smsText,omitempty"`                     //
}

type ResponseGuestTypeGetGuestTypeByIDGuestTypeLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestTypeUpdateGuestTypeByID struct {
	UpdatedFieldsList ResponseGuestTypeUpdateGuestTypeByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseGuestTypeUpdateGuestTypeByIDUpdatedFieldsList struct {
	UpdatedField ResponseGuestTypeUpdateGuestTypeByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                            `json:"field,omitempty"`        //
	OldValue     string                                                            `json:"oldValue,omitempty"`     //
	NewValue     string                                                            `json:"newValue,omitempty"`     //
}

type ResponseGuestTypeUpdateGuestTypeByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseGuestTypeGetGuestType struct {
	SearchResult ResponseGuestTypeGetGuestTypeSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseGuestTypeGetGuestTypeSearchResult struct {
	Total        int                                                   `json:"total,omitempty"`        //
	Resources    []ResponseGuestTypeGetGuestTypeSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponseGuestTypeGetGuestTypeSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponseGuestTypeGetGuestTypeSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseGuestTypeGetGuestTypeSearchResultResources struct {
	ID          string                                                 `json:"id,omitempty"`          //
	Name        string                                                 `json:"name,omitempty"`        //
	Description string                                                 `json:"description,omitempty"` //
	Link        ResponseGuestTypeGetGuestTypeSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseGuestTypeGetGuestTypeSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestTypeGetGuestTypeSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestTypeGetGuestTypeSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestTypeGetVersion struct {
	VersionInfo ResponseGuestTypeGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseGuestTypeGetVersionVersionInfo struct {
	CurrentServerVersion string                                     `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                     `json:"supportedVersions,omitempty"`    //
	Link                 ResponseGuestTypeGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseGuestTypeGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestGuestTypeUpdateGuestTypeEmail struct {
	OperationAdditionalData *RequestGuestTypeUpdateGuestTypeEmailOperationAdditionalData `json:"OperationAdditionalData,omitempty"` //
}

type RequestGuestTypeUpdateGuestTypeEmailOperationAdditionalData struct {
	AdditionalData *[]RequestGuestTypeUpdateGuestTypeEmailOperationAdditionalDataAdditionalData `json:"additionalData,omitempty"` //
}

type RequestGuestTypeUpdateGuestTypeEmailOperationAdditionalDataAdditionalData struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type RequestGuestTypeUpdateGuestTypeSms struct {
	OperationAdditionalData *RequestGuestTypeUpdateGuestTypeSmsOperationAdditionalData `json:"OperationAdditionalData,omitempty"` //
}

type RequestGuestTypeUpdateGuestTypeSmsOperationAdditionalData struct {
	AdditionalData *[]RequestGuestTypeUpdateGuestTypeSmsOperationAdditionalDataAdditionalData `json:"additionalData,omitempty"` //
}

type RequestGuestTypeUpdateGuestTypeSmsOperationAdditionalDataAdditionalData struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type RequestGuestTypeUpdateGuestTypeByID struct {
	GuestType *RequestGuestTypeUpdateGuestTypeByIDGuestType `json:"GuestType,omitempty"` //
}

type RequestGuestTypeUpdateGuestTypeByIDGuestType struct {
	ID                     string                                                              `json:"id,omitempty"`                     //
	Name                   string                                                              `json:"name,omitempty"`                   //
	Description            string                                                              `json:"description,omitempty"`            //
	IsDefaultType          *bool                                                               `json:"isDefaultType,omitempty"`          //
	AccessTime             *RequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTime             `json:"accessTime,omitempty"`             //
	LoginOptions           *RequestGuestTypeUpdateGuestTypeByIDGuestTypeLoginOptions           `json:"loginOptions,omitempty"`           //
	ExpirationNotification *RequestGuestTypeUpdateGuestTypeByIDGuestTypeExpirationNotification `json:"expirationNotification,omitempty"` // Expiration Notification Settings
	SponsorGroups          []string                                                            `json:"sponsorGroups,omitempty"`          //
}

type RequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTime struct {
	FromFirstLogin                 *bool                                                                  `json:"fromFirstLogin,omitempty"`                 // When Account Duration starts from first login or specified date
	MaxAccountDuration             *int                                                                   `json:"maxAccountDuration,omitempty"`             // Maximum value of Account Duration
	DurationTimeUnit               string                                                                 `json:"durationTimeUnit,omitempty"`               // Allowed values are: - DAYS, - HOURS, - MINUTES
	DefaultDuration                *int                                                                   `json:"defaultDuration,omitempty"`                //
	AllowAccessOnSpecificDaysTimes *bool                                                                  `json:"allowAccessOnSpecificDaysTimes,omitempty"` //
	DayTimeLimits                  *[]RequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimits `json:"dayTimeLimits,omitempty"`                  // List of Time Ranges for account access
}

type RequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimits struct {
	StartTime string   `json:"startTime,omitempty"` // Start time in HH:mm format
	EndTime   string   `json:"endTime,omitempty"`   // End time in HH:mm format
	Days      []string `json:"days,omitempty"`      // List of Days Values should be one of Week day. Allowed values are: - Sunday, - Monday, - Tuesday, - Wednesday, - Thursday, - Friday, - Saturday
}

type RequestGuestTypeUpdateGuestTypeByIDGuestTypeLoginOptions struct {
	LimitSimultaneousLogins *bool  `json:"limitSimultaneousLogins,omitempty"` // Enable Simultaneous Logins
	MaxSimultaneousLogins   *int   `json:"maxSimultaneousLogins,omitempty"`   // Number of Simultaneous Logins
	FailureAction           string `json:"failureAction,omitempty"`           // When Guest Exceeds limit this action will be invoked. Allowed values are: - Disconnect_Oldest_Connection, - Disconnect_Newest_Connection
	MaxRegisteredDevices    *int   `json:"maxRegisteredDevices,omitempty"`    // Maximum devices guests can register
	IDentityGroupID         string `json:"identityGroupId,omitempty"`         //
	AllowGuestPortalBypass  *bool  `json:"allowGuestPortalBypass,omitempty"`  //
}

type RequestGuestTypeUpdateGuestTypeByIDGuestTypeExpirationNotification struct {
	EnableNotification          *bool  `json:"enableNotification,omitempty"`          // Enable Notification settings
	AdvanceNotificationDuration *int   `json:"advanceNotificationDuration,omitempty"` // Send Account Expiration Notification Duration before ( Days, Hours, Minutes )
	AdvanceNotificationUnits    string `json:"advanceNotificationUnits,omitempty"`    // Allowed values are: - DAYS, - HOURS, - MINUTES
	SendEmailNotification       *bool  `json:"sendEmailNotification,omitempty"`       // Enable Email Notification
	EmailText                   string `json:"emailText,omitempty"`                   //
	SendSmsNotification         *bool  `json:"sendSmsNotification,omitempty"`         // Maximum devices guests can register
	SmsText                     string `json:"smsText,omitempty"`                     //
}

type RequestGuestTypeCreateGuestType struct {
	GuestType *RequestGuestTypeCreateGuestTypeGuestType `json:"GuestType,omitempty"` //
}

type RequestGuestTypeCreateGuestTypeGuestType struct {
	Name                   string                                                          `json:"name,omitempty"`                   //
	Description            string                                                          `json:"description,omitempty"`            //
	IsDefaultType          *bool                                                           `json:"isDefaultType,omitempty"`          //
	AccessTime             *RequestGuestTypeCreateGuestTypeGuestTypeAccessTime             `json:"accessTime,omitempty"`             //
	LoginOptions           *RequestGuestTypeCreateGuestTypeGuestTypeLoginOptions           `json:"loginOptions,omitempty"`           //
	ExpirationNotification *RequestGuestTypeCreateGuestTypeGuestTypeExpirationNotification `json:"expirationNotification,omitempty"` // Expiration Notification Settings
	SponsorGroups          []string                                                        `json:"sponsorGroups,omitempty"`          //
}

type RequestGuestTypeCreateGuestTypeGuestTypeAccessTime struct {
	FromFirstLogin                 *bool                                                              `json:"fromFirstLogin,omitempty"`                 // When Account Duration starts from first login or specified date
	MaxAccountDuration             *int                                                               `json:"maxAccountDuration,omitempty"`             // Maximum value of Account Duration
	DurationTimeUnit               string                                                             `json:"durationTimeUnit,omitempty"`               // Allowed values are: - DAYS, - HOURS, - MINUTES
	DefaultDuration                *int                                                               `json:"defaultDuration,omitempty"`                //
	AllowAccessOnSpecificDaysTimes *bool                                                              `json:"allowAccessOnSpecificDaysTimes,omitempty"` //
	DayTimeLimits                  *[]RequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimits `json:"dayTimeLimits,omitempty"`                  // List of Time Ranges for account access
}

type RequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimits struct {
	StartTime string   `json:"startTime,omitempty"` // Start time in HH:mm format
	EndTime   string   `json:"endTime,omitempty"`   // End time in HH:mm format
	Days      []string `json:"days,omitempty"`      // List of Days Values should be one of Week day. Allowed values are: - Sunday, - Monday, - Tuesday, - Wednesday, - Thursday, - Friday, - Saturday
}

type RequestGuestTypeCreateGuestTypeGuestTypeLoginOptions struct {
	LimitSimultaneousLogins *bool  `json:"limitSimultaneousLogins,omitempty"` // Enable Simultaneous Logins
	MaxSimultaneousLogins   *int   `json:"maxSimultaneousLogins,omitempty"`   // Number of Simultaneous Logins
	FailureAction           string `json:"failureAction,omitempty"`           // When Guest Exceeds limit this action will be invoked. Allowed values are: - Disconnect_Oldest_Connection, - Disconnect_Newest_Connection
	MaxRegisteredDevices    *int   `json:"maxRegisteredDevices,omitempty"`    // Maximum devices guests can register
	IDentityGroupID         string `json:"identityGroupId,omitempty"`         //
	AllowGuestPortalBypass  *bool  `json:"allowGuestPortalBypass,omitempty"`  //
}

type RequestGuestTypeCreateGuestTypeGuestTypeExpirationNotification struct {
	EnableNotification          *bool  `json:"enableNotification,omitempty"`          // Enable Notification settings
	AdvanceNotificationDuration *int   `json:"advanceNotificationDuration,omitempty"` // Send Account Expiration Notification Duration before ( Days, Hours, Minutes )
	AdvanceNotificationUnits    string `json:"advanceNotificationUnits,omitempty"`    // Allowed values are: - DAYS, - HOURS, - MINUTES
	SendEmailNotification       *bool  `json:"sendEmailNotification,omitempty"`       // Enable Email Notification
	EmailText                   string `json:"emailText,omitempty"`                   //
	SendSmsNotification         *bool  `json:"sendSmsNotification,omitempty"`         // Maximum devices guests can register
	SmsText                     string `json:"smsText,omitempty"`                     //
}

//GetGuestTypeByID Get guest type by ID
/* This API allows the client to get a guest type by ID.

@param id id path parameter.
*/
func (s *GuestTypeService) GetGuestTypeByID(id string) (*ResponseGuestTypeGetGuestTypeByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guesttype/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseGuestTypeGetGuestTypeByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetGuestTypeById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestTypeGetGuestTypeByID)
	return result, response, err

}

//GetGuestType Get all guest types
/* This API allows the client to get all the guest types.

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

@param getGuestTypeQueryParams Filtering parameter
*/
func (s *GuestTypeService) GetGuestType(getGuestTypeQueryParams *GetGuestTypeQueryParams) (*ResponseGuestTypeGetGuestType, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guesttype"

	queryString, _ := query.Values(getGuestTypeQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseGuestTypeGetGuestType{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetGuestType")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestTypeGetGuestType)
	return result, response, err

}

//GetVersion Get guest type version information
/* This API helps to retrieve the version information related to the guest type.

 */
func (s *GuestTypeService) GetVersion() (*ResponseGuestTypeGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guesttype/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseGuestTypeGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestTypeGetVersion)
	return result, response, err

}

//CreateGuestType Create guest type
/* This API creates a guest type.

 */
func (s *GuestTypeService) CreateGuestType(requestGuestTypeCreateGuestType *RequestGuestTypeCreateGuestType) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guesttype"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestGuestTypeCreateGuestType).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateGuestType")
	}

	return response, err

}

//UpdateGuestTypeEmail Update guest type email by ID
/* This API allows the client to update a guest type email by ID.

@param id id path parameter.
*/
func (s *GuestTypeService) UpdateGuestTypeEmail(id string, requestGuestTypeUpdateGuestTypeEmail *RequestGuestTypeUpdateGuestTypeEmail) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guesttype/email/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestGuestTypeUpdateGuestTypeEmail).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateGuestTypeEmail")
	}

	getCSFRToken(response.Header())

	return response, err

}

//UpdateGuestTypeSms Update guest type sms by ID
/* This API allows the client to update a guest type sms by ID.

@param id id path parameter.
*/
func (s *GuestTypeService) UpdateGuestTypeSms(id string, requestGuestTypeUpdateGuestTypeSms *RequestGuestTypeUpdateGuestTypeSms) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guesttype/sms/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestGuestTypeUpdateGuestTypeSms).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateGuestTypeSms")
	}

	getCSFRToken(response.Header())

	return response, err

}

//UpdateGuestTypeByID Update guest type
/* This API allows the client to update a guest type.

@param id id path parameter.
*/
func (s *GuestTypeService) UpdateGuestTypeByID(id string, requestGuestTypeUpdateGuestTypeById *RequestGuestTypeUpdateGuestTypeByID) (*ResponseGuestTypeUpdateGuestTypeByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guesttype/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestGuestTypeUpdateGuestTypeById).
		SetResult(&ResponseGuestTypeUpdateGuestTypeByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateGuestTypeById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestTypeUpdateGuestTypeByID)
	return result, response, err

}

//DeleteGuestTypeByID Delete guest type
/* This API deletes a guest type.

@param id id path parameter.
*/
func (s *GuestTypeService) DeleteGuestTypeByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guesttype/{id}"
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
		return response, fmt.Errorf("error with operation DeleteGuestTypeById")
	}

	getCSFRToken(response.Header())

	return response, err

}
