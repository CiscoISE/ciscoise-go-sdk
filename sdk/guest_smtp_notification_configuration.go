package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type GuestSmtpNotificationConfigurationService service

type GetGuestSmtpNotificationSettingsQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByID struct {
	ERSGuestSmtpNotificationSettings ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByIDERSGuestSmtpNotificationSettings `json:"ERSGuestSmtpNotificationSettings,omitempty"` //
}

type ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByIDERSGuestSmtpNotificationSettings struct {
	ID                        string                                                                                                             `json:"id,omitempty"`                        //
	SmtpServer                string                                                                                                             `json:"smtpServer,omitempty"`                // The SMTP server ip address or fqdn such as outbound.mycompany.com
	NotificationEnabled       bool                                                                                                               `json:"notificationEnabled,omitempty"`       // Indicates if the email notification service is to be enabled
	UseDefaultFromAddress     bool                                                                                                               `json:"useDefaultFromAddress,omitempty"`     // If the default from address should be used rather than using a sponsor user email address
	DefaultFromAddress        string                                                                                                             `json:"defaultFromAddress,omitempty"`        // The default from email address to be used to send emails from
	SmtpPort                  string                                                                                                             `json:"smtpPort,omitempty"`                  // Port at which SMTP Secure Server is listening
	ConnectionTimeout         string                                                                                                             `json:"connectionTimeout,omitempty"`         // Interval in seconds for all the SMTP client connections
	UseTLSorSSLEncryption     bool                                                                                                               `json:"useTLSorSSLEncryption,omitempty"`     // If configured to true, SMTP server authentication will happen using TLS/SSL
	UsePasswordAuthentication bool                                                                                                               `json:"usePasswordAuthentication,omitempty"` // If configured to true, SMTP server authentication will happen using username/password
	UserName                  string                                                                                                             `json:"userName,omitempty"`                  // Username of Secure SMTP server
	Password                  string                                                                                                             `json:"password,omitempty"`                  // Password of Secure SMTP server
	Link                      ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByIDERSGuestSmtpNotificationSettingsLink `json:"link,omitempty"`                      //
}

type ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByIDERSGuestSmtpNotificationSettingsLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsByID struct {
	UpdatedFieldsList ResponseGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsByIDUpdatedFieldsList struct {
	UpdatedField ResponseGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                                                         `json:"field,omitempty"`        //
	OldValue     string                                                                                                         `json:"oldValue,omitempty"`     //
	NewValue     string                                                                                                         `json:"newValue,omitempty"`     //
}

type ResponseGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettings struct {
	SearchResult ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsSearchResult struct {
	Total        int                                                                                                `json:"total,omitempty"`        //
	Resources    []ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsSearchResultResources struct {
	ID   string                                                                                              `json:"id,omitempty"`   //
	Link ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsSearchResultResourcesLink `json:"link,omitempty"` //
}

type ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseGuestSmtpNotificationConfigurationGetVersion struct {
	VersionInfo ResponseGuestSmtpNotificationConfigurationGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseGuestSmtpNotificationConfigurationGetVersionVersionInfo struct {
	CurrentServerVersion string                                                              `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                                              `json:"supportedVersions,omitempty"`    //
	Link                 ResponseGuestSmtpNotificationConfigurationGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseGuestSmtpNotificationConfigurationGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsByID struct {
	ERSGuestSmtpNotificationSettings *RequestGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsByIDERSGuestSmtpNotificationSettings `json:"ERSGuestSmtpNotificationSettings,omitempty"` //
}

type RequestGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsByIDERSGuestSmtpNotificationSettings struct {
	ID                        string `json:"id,omitempty"`                        //
	SmtpServer                string `json:"smtpServer,omitempty"`                // The SMTP server ip address or fqdn such as outbound.mycompany.com
	NotificationEnabled       *bool  `json:"notificationEnabled,omitempty"`       // Indicates if the email notification service is to be enabled
	UseDefaultFromAddress     *bool  `json:"useDefaultFromAddress,omitempty"`     // If the default from address should be used rather than using a sponsor user email address
	DefaultFromAddress        string `json:"defaultFromAddress,omitempty"`        // The default from email address to be used to send emails from
	SmtpPort                  string `json:"smtpPort,omitempty"`                  // Port at which SMTP Secure Server is listening
	ConnectionTimeout         string `json:"connectionTimeout,omitempty"`         // Interval in seconds for all the SMTP client connections
	UseTLSorSSLEncryption     *bool  `json:"useTLSorSSLEncryption,omitempty"`     // If configured to true, SMTP server authentication will happen using TLS/SSL
	UsePasswordAuthentication *bool  `json:"usePasswordAuthentication,omitempty"` // If configured to true, SMTP server authentication will happen using username/password
	UserName                  string `json:"userName,omitempty"`                  // Username of Secure SMTP server
	Password                  string `json:"password,omitempty"`                  // Password of Secure SMTP server
}

type RequestGuestSmtpNotificationConfigurationCreateGuestSmtpNotificationSettings struct {
	ERSGuestSmtpNotificationSettings *RequestGuestSmtpNotificationConfigurationCreateGuestSmtpNotificationSettingsERSGuestSmtpNotificationSettings `json:"ERSGuestSmtpNotificationSettings,omitempty"` //
}

type RequestGuestSmtpNotificationConfigurationCreateGuestSmtpNotificationSettingsERSGuestSmtpNotificationSettings struct {
	SmtpServer                string `json:"smtpServer,omitempty"`                // The SMTP server ip address or fqdn such as outbound.mycompany.com
	NotificationEnabled       *bool  `json:"notificationEnabled,omitempty"`       // Indicates if the email notification service is to be enabled
	UseDefaultFromAddress     *bool  `json:"useDefaultFromAddress,omitempty"`     // If the default from address should be used rather than using a sponsor user email address
	DefaultFromAddress        string `json:"defaultFromAddress,omitempty"`        // The default from email address to be used to send emails from
	SmtpPort                  string `json:"smtpPort,omitempty"`                  // Port at which SMTP Secure Server is listening
	ConnectionTimeout         string `json:"connectionTimeout,omitempty"`         // Interval in seconds for all the SMTP client connections
	UseTLSorSSLEncryption     *bool  `json:"useTLSorSSLEncryption,omitempty"`     // If configured to true, SMTP server authentication will happen using TLS/SSL
	UsePasswordAuthentication *bool  `json:"usePasswordAuthentication,omitempty"` // If configured to true, SMTP server authentication will happen using username/password
	UserName                  string `json:"userName,omitempty"`                  // Username of Secure SMTP server
	Password                  string `json:"password,omitempty"`                  // Password of Secure SMTP server
}

//GetGuestSmtpNotificationSettingsByID Get guest SMTP notification configuration by ID
/* This API allows the client to get a guest SMTP notification configuration by ID.

@param id id path parameter.
*/
func (s *GuestSmtpNotificationConfigurationService) GetGuestSmtpNotificationSettingsByID(id string) (*ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestsmtpnotificationsettings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetGuestSmtpNotificationSettingsById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByID)
	return result, response, err

}

//GetGuestSmtpNotificationSettings Get all guest SMTP notification configurations
/* This API allows the client to get all the guest SMTP notification configurations.

Filter:

[name]

To search guest users by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:

[name, description]

@param getGuestSmtpNotificationSettingsQueryParams Filtering parameter
*/
func (s *GuestSmtpNotificationConfigurationService) GetGuestSmtpNotificationSettings(getGuestSmtpNotificationSettingsQueryParams *GetGuestSmtpNotificationSettingsQueryParams) (*ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettings, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestsmtpnotificationsettings"

	queryString, _ := query.Values(getGuestSmtpNotificationSettingsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetGuestSmtpNotificationSettings")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettings)
	return result, response, err

}

//GetVersion Get guest smtp notification configuration version information
/* This API helps to retrieve the version information related to the guest smtp notification configuration.

 */
func (s *GuestSmtpNotificationConfigurationService) GetVersion() (*ResponseGuestSmtpNotificationConfigurationGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestsmtpnotificationsettings/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseGuestSmtpNotificationConfigurationGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestSmtpNotificationConfigurationGetVersion)
	return result, response, err

}

//CreateGuestSmtpNotificationSettings Create guest SMTP notification configuration
/* This API creates a guest SMTP notification configuration.

 */
func (s *GuestSmtpNotificationConfigurationService) CreateGuestSmtpNotificationSettings(requestGuestSmtpNotificationConfigurationCreateGuestSmtpNotificationSettings *RequestGuestSmtpNotificationConfigurationCreateGuestSmtpNotificationSettings) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestsmtpnotificationsettings"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestGuestSmtpNotificationConfigurationCreateGuestSmtpNotificationSettings).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateGuestSmtpNotificationSettings")
	}

	return response, err

}

//UpdateGuestSmtpNotificationSettingsByID Update SMTP notification configuration
/* This API allows the client to update a SMTP configuration setting.

@param id id path parameter.
*/
func (s *GuestSmtpNotificationConfigurationService) UpdateGuestSmtpNotificationSettingsByID(id string, requestGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsById *RequestGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsByID) (*ResponseGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/guestsmtpnotificationsettings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsById).
		SetResult(&ResponseGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateGuestSmtpNotificationSettingsById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsByID)
	return result, response, err

}
