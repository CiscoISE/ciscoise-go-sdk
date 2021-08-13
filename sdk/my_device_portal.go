package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type MyDevicePortalService service

type GetMyDevicePortalQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseMyDevicePortalGetMyDevicePortalByID struct {
	MyDevicePortal ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortal `json:"MyDevicePortal,omitempty"` //
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortal struct {
	ID             string                                                                  `json:"id,omitempty"`             //
	Name           string                                                                  `json:"name,omitempty"`           //
	Description    string                                                                  `json:"description,omitempty"`    //
	PortalType     string                                                                  `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                                  `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a Mydevice portal
	Customizations ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizations `json:"customizations,omitempty"` // Defines all of the Portal Customizations available
	Link           ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalLink           `json:"link,omitempty"`           //
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettings struct {
	PortalSettings                 ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsPortalSettings                 `json:"portalSettings,omitempty"`                 // The port, interface, certificate, and other basic settings of a portal
	LoginPageSettings              ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettings              `json:"loginPageSettings,omitempty"`              //
	AupSettings                    ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsAupSettings                    `json:"aupSettings,omitempty"`                    // Configuration of the Acceptable Use Policy (AUP) for a portal
	EmployeeChangePasswordSettings ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsEmployeeChangePasswordSettings `json:"employeeChangePasswordSettings,omitempty"` //
	PostLoginBannerSettings        ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsPostLoginBannerSettings        `json:"postLoginBannerSettings,omitempty"`        //
	PostAccessBannerSettings       ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsPostAccessBannerSettings       `json:"postAccessBannerSettings,omitempty"`       //
	SupportInfoSettings            ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsSupportInfoSettings            `json:"supportInfoSettings,omitempty"`            //
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsPortalSettings struct {
	HTTPSPort             int    `json:"httpsPort,omitempty"`             // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces     string `json:"allowedInterfaces,omitempty"`     // Interfaces that the portal will be reachable on. Allowed values: - eth0, - eth1, - eth2, - eth3, - eth4, - eth5, - bond0, - bond1, - bond2
	CertificateGroupTag   string `json:"certificateGroupTag,omitempty"`   // Logical name of the x.509 server certificate that will be used for the portal
	EndpointIDentityGroup string `json:"endpointIdentityGroup,omitempty"` // Unique Id of the endpoint identity group where user's devices will be added. Used only in Hotspot Portal
	DisplayLang           string `json:"displayLang,omitempty"`           // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage      string `json:"fallbackLanguage,omitempty"`      // Used when displayLang = USEBROWSERLOCALE
	AlwaysUsedLanguage    string `json:"alwaysUsedLanguage,omitempty"`    //
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettings struct {
	MaxFailedAttemptsBeforeRateLimit int           `json:"maxFailedAttemptsBeforeRateLimit,omitempty"` // Maximum failed login attempts before rate limiting
	TimeBetweenLoginsDuringRateLimit int           `json:"timeBetweenLoginsDuringRateLimit,omitempty"` // Time between login attempts when rate limiting
	IncludeAup                       bool          `json:"includeAup,omitempty"`                       // Include an Acceptable Use Policy (AUP) that should be displayed during login
	AupDisplay                       string        `json:"aupDisplay,omitempty"`                       // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: -  ONPAGE, - ASLINK
	RequireAupAcceptance             bool          `json:"requireAupAcceptance,omitempty"`             // Require the portal user to accept the AUP. Only valid if includeAup = true
	RequireScrolling                 bool          `json:"requireScrolling,omitempty"`                 // Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true
	SocialConfigs                    []interface{} `json:"socialConfigs,omitempty"`                    //
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsAupSettings struct {
	DisplayFrequencyIntervalDays int    `json:"displayFrequencyIntervalDays,omitempty"` // Number of days between AUP confirmations (when displayFrequency = recurring)
	DisplayFrequency             string `json:"displayFrequency,omitempty"`             // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed Values: - FIRSTLOGIN, - EVERYLOGIN, - RECURRING
	IncludeAup                   bool   `json:"includeAup,omitempty"`                   // Require the portal user to read and accept an AUP
	RequireScrolling             bool   `json:"requireScrolling,omitempty"`             // Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsEmployeeChangePasswordSettings struct {
	AllowEmployeeToChangePwd bool `json:"allowEmployeeToChangePwd,omitempty"` //
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsPostLoginBannerSettings struct {
	IncludePostAccessBanner bool `json:"includePostAccessBanner,omitempty"` // Include a Post-Login Banner page
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsPostAccessBannerSettings struct {
	IncludePostAccessBanner bool `json:"includePostAccessBanner,omitempty"` //
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  bool   `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          bool   `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        bool   `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent bool   `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     bool   `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      bool   `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizations struct {
	PortalTheme          ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          //
	PortalTweakSettings  ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` //
	PageCustomizations   ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   // Represent the entire page customization as a giant dictionary
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BannerImage      ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BackgroundImage  ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerTitle      string                                                                                                      `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                                      `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                                      `json:"footerElement,omitempty"`    //
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizations struct {
	Data []ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsData `json:"data,omitempty"` // The Dictionary will be exposed here as key value pair
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

type ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseMyDevicePortalUpdateMyDevicePortalByID struct {
	UpdatedFieldsList ResponseMyDevicePortalUpdateMyDevicePortalByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseMyDevicePortalUpdateMyDevicePortalByIDUpdatedFieldsList struct {
	UpdatedField ResponseMyDevicePortalUpdateMyDevicePortalByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                      `json:"field,omitempty"`        //
	OldValue     string                                                                      `json:"oldValue,omitempty"`     //
	NewValue     string                                                                      `json:"newValue,omitempty"`     //
}

type ResponseMyDevicePortalUpdateMyDevicePortalByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseMyDevicePortalGetMyDevicePortal struct {
	SearchResult ResponseMyDevicePortalGetMyDevicePortalSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseMyDevicePortalGetMyDevicePortalSearchResult struct {
	Total        int                                                               `json:"total,omitempty"`        //
	Resources    []ResponseMyDevicePortalGetMyDevicePortalSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseMyDevicePortalGetMyDevicePortalSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseMyDevicePortalGetMyDevicePortalSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseMyDevicePortalGetMyDevicePortalSearchResultResources struct {
	ID          string                                                           `json:"id,omitempty"`          //
	Name        string                                                           `json:"name,omitempty"`        //
	Description string                                                           `json:"description,omitempty"` //
	Link        ResponseMyDevicePortalGetMyDevicePortalSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseMyDevicePortalGetMyDevicePortalSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseMyDevicePortalGetMyDevicePortalSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseMyDevicePortalGetMyDevicePortalSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseMyDevicePortalGetVersion struct {
	VersionInfo ResponseMyDevicePortalGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseMyDevicePortalGetVersionVersionInfo struct {
	CurrentServerVersion string                                          `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                          `json:"supportedVersions,omitempty"`    //
	Link                 ResponseMyDevicePortalGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseMyDevicePortalGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestMyDevicePortalUpdateMyDevicePortalByID struct {
	MyDevicePortal *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortal `json:"MyDevicePortal,omitempty"` //
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortal struct {
	ID             string                                                                     `json:"id,omitempty"`             //
	Name           string                                                                     `json:"name,omitempty"`           //
	Description    string                                                                     `json:"description,omitempty"`    //
	PortalType     string                                                                     `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                                     `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a Mydevice portal
	Customizations *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizations `json:"customizations,omitempty"` // Defines all of the Portal Customizations available
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettings struct {
	PortalSettings                 *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPortalSettings                 `json:"portalSettings,omitempty"`                 // The port, interface, certificate, and other basic settings of a portal
	LoginPageSettings              *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettings              `json:"loginPageSettings,omitempty"`              //
	AupSettings                    *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsAupSettings                    `json:"aupSettings,omitempty"`                    // Configuration of the Acceptable Use Policy (AUP) for a portal
	EmployeeChangePasswordSettings *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsEmployeeChangePasswordSettings `json:"employeeChangePasswordSettings,omitempty"` //
	PostLoginBannerSettings        *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostLoginBannerSettings        `json:"postLoginBannerSettings,omitempty"`        //
	PostAccessBannerSettings       *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostAccessBannerSettings       `json:"postAccessBannerSettings,omitempty"`       //
	SupportInfoSettings            *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsSupportInfoSettings            `json:"supportInfoSettings,omitempty"`            //
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPortalSettings struct {
	HTTPSPort             *int   `json:"httpsPort,omitempty"`             // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces     string `json:"allowedInterfaces,omitempty"`     // Interfaces that the portal will be reachable on. Allowed values: - eth0, - eth1, - eth2, - eth3, - eth4, - eth5, - bond0, - bond1, - bond2
	CertificateGroupTag   string `json:"certificateGroupTag,omitempty"`   // Logical name of the x.509 server certificate that will be used for the portal
	EndpointIDentityGroup string `json:"endpointIdentityGroup,omitempty"` // Unique Id of the endpoint identity group where user's devices will be added. Used only in Hotspot Portal
	DisplayLang           string `json:"displayLang,omitempty"`           // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage      string `json:"fallbackLanguage,omitempty"`      // Used when displayLang = USEBROWSERLOCALE
	AlwaysUsedLanguage    string `json:"alwaysUsedLanguage,omitempty"`    //
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettings struct {
	MaxFailedAttemptsBeforeRateLimit *int           `json:"maxFailedAttemptsBeforeRateLimit,omitempty"` // Maximum failed login attempts before rate limiting
	TimeBetweenLoginsDuringRateLimit *int           `json:"timeBetweenLoginsDuringRateLimit,omitempty"` // Time between login attempts when rate limiting
	IncludeAup                       *bool          `json:"includeAup,omitempty"`                       // Include an Acceptable Use Policy (AUP) that should be displayed during login
	AupDisplay                       string         `json:"aupDisplay,omitempty"`                       // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: -  ONPAGE, - ASLINK
	RequireAupAcceptance             *bool          `json:"requireAupAcceptance,omitempty"`             // Require the portal user to accept the AUP. Only valid if includeAup = true
	RequireScrolling                 *bool          `json:"requireScrolling,omitempty"`                 // Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true
	SocialConfigs                    *[]interface{} `json:"socialConfigs,omitempty"`                    //
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsAupSettings struct {
	DisplayFrequencyIntervalDays *int   `json:"displayFrequencyIntervalDays,omitempty"` // Number of days between AUP confirmations (when displayFrequency = recurring)
	DisplayFrequency             string `json:"displayFrequency,omitempty"`             // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed Values: - FIRSTLOGIN, - EVERYLOGIN, - RECURRING
	IncludeAup                   *bool  `json:"includeAup,omitempty"`                   // Require the portal user to read and accept an AUP
	RequireScrolling             *bool  `json:"requireScrolling,omitempty"`             // Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsEmployeeChangePasswordSettings struct {
	AllowEmployeeToChangePwd *bool `json:"allowEmployeeToChangePwd,omitempty"` //
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostLoginBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` // Include a Post-Login Banner page
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostAccessBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` //
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  *bool  `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          *bool  `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        *bool  `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent *bool  `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     *bool  `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      *bool  `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizations struct {
	PortalTheme          *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          //
	PortalTweakSettings  *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` //
	PageCustomizations   *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   // Represent the entire page customization as a giant dictionary
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BannerImage      *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BackgroundImage  *RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerTitle      string                                                                                                         `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                                         `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                                         `json:"footerElement,omitempty"`    //
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizations struct {
	Data *[]RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsData `json:"data,omitempty"` // The Dictionary will be exposed here as key value pair
}

type RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

type RequestMyDevicePortalCreateMyDevicePortal struct {
	MyDevicePortal *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortal `json:"MyDevicePortal,omitempty"` //
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortal struct {
	Name           string                                                                 `json:"name,omitempty"`           //
	Description    string                                                                 `json:"description,omitempty"`    //
	PortalType     string                                                                 `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                                 `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a Mydevice portal
	Customizations *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizations `json:"customizations,omitempty"` // Defines all of the Portal Customizations available
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettings struct {
	PortalSettings                 *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPortalSettings                 `json:"portalSettings,omitempty"`                 // The port, interface, certificate, and other basic settings of a portal
	LoginPageSettings              *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettings              `json:"loginPageSettings,omitempty"`              //
	AupSettings                    *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsAupSettings                    `json:"aupSettings,omitempty"`                    // Configuration of the Acceptable Use Policy (AUP) for a portal
	EmployeeChangePasswordSettings *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsEmployeeChangePasswordSettings `json:"employeeChangePasswordSettings,omitempty"` //
	PostLoginBannerSettings        *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostLoginBannerSettings        `json:"postLoginBannerSettings,omitempty"`        //
	PostAccessBannerSettings       *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostAccessBannerSettings       `json:"postAccessBannerSettings,omitempty"`       //
	SupportInfoSettings            *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsSupportInfoSettings            `json:"supportInfoSettings,omitempty"`            //
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPortalSettings struct {
	HTTPSPort             *int   `json:"httpsPort,omitempty"`             // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces     string `json:"allowedInterfaces,omitempty"`     // Interfaces that the portal will be reachable on. Allowed values: - eth0, - eth1, - eth2, - eth3, - eth4, - eth5, - bond0, - bond1, - bond2
	CertificateGroupTag   string `json:"certificateGroupTag,omitempty"`   // Logical name of the x.509 server certificate that will be used for the portal
	EndpointIDentityGroup string `json:"endpointIdentityGroup,omitempty"` // Unique Id of the endpoint identity group where user's devices will be added. Used only in Hotspot Portal
	DisplayLang           string `json:"displayLang,omitempty"`           // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage      string `json:"fallbackLanguage,omitempty"`      // Used when displayLang = USEBROWSERLOCALE
	AlwaysUsedLanguage    string `json:"alwaysUsedLanguage,omitempty"`    //
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettings struct {
	MaxFailedAttemptsBeforeRateLimit *int           `json:"maxFailedAttemptsBeforeRateLimit,omitempty"` // Maximum failed login attempts before rate limiting
	TimeBetweenLoginsDuringRateLimit *int           `json:"timeBetweenLoginsDuringRateLimit,omitempty"` // Time between login attempts when rate limiting
	IncludeAup                       *bool          `json:"includeAup,omitempty"`                       // Include an Acceptable Use Policy (AUP) that should be displayed during login
	AupDisplay                       string         `json:"aupDisplay,omitempty"`                       // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: -  ONPAGE, - ASLINK
	RequireAupAcceptance             *bool          `json:"requireAupAcceptance,omitempty"`             // Require the portal user to accept the AUP. Only valid if includeAup = true
	RequireScrolling                 *bool          `json:"requireScrolling,omitempty"`                 // Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true
	SocialConfigs                    *[]interface{} `json:"socialConfigs,omitempty"`                    //
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsAupSettings struct {
	DisplayFrequencyIntervalDays *int   `json:"displayFrequencyIntervalDays,omitempty"` // Number of days between AUP confirmations (when displayFrequency = recurring)
	DisplayFrequency             string `json:"displayFrequency,omitempty"`             // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed Values: - FIRSTLOGIN, - EVERYLOGIN, - RECURRING
	IncludeAup                   *bool  `json:"includeAup,omitempty"`                   // Require the portal user to read and accept an AUP
	RequireScrolling             *bool  `json:"requireScrolling,omitempty"`             // Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsEmployeeChangePasswordSettings struct {
	AllowEmployeeToChangePwd *bool `json:"allowEmployeeToChangePwd,omitempty"` //
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostLoginBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` // Include a Post-Login Banner page
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostAccessBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` //
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  *bool  `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          *bool  `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        *bool  `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent *bool  `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     *bool  `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      *bool  `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizations struct {
	PortalTheme          *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          //
	PortalTweakSettings  *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` //
	PageCustomizations   *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   // Represent the entire page customization as a giant dictionary
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BannerImage      *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BackgroundImage  *RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerTitle      string                                                                                                     `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                                     `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                                     `json:"footerElement,omitempty"`    //
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizations struct {
	Data *[]RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsData `json:"data,omitempty"` // The Dictionary will be exposed here as key value pair
}

type RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

//GetMyDevicePortalByID Get my device portal by ID
/* This API allows the client to get a my device portal by ID.

@param id id path parameter.
*/
func (s *MyDevicePortalService) GetMyDevicePortalByID(id string) (*ResponseMyDevicePortalGetMyDevicePortalByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/mydeviceportal/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseMyDevicePortalGetMyDevicePortalByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetMyDevicePortalById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseMyDevicePortalGetMyDevicePortalByID)
	return result, response, err

}

//GetMyDevicePortal Get all my device portals
/* This API allows the client to get all the my device portals.

Filter:

[name, description]

To search resources by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:

[name, description]

@param getMyDevicePortalQueryParams Filtering parameter
*/
func (s *MyDevicePortalService) GetMyDevicePortal(getMyDevicePortalQueryParams *GetMyDevicePortalQueryParams) (*ResponseMyDevicePortalGetMyDevicePortal, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/mydeviceportal"

	queryString, _ := query.Values(getMyDevicePortalQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseMyDevicePortalGetMyDevicePortal{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetMyDevicePortal")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseMyDevicePortalGetMyDevicePortal)
	return result, response, err

}

//GetVersion Get my device portal version information
/* This API helps to retrieve the version information related to the my device portal.

 */
func (s *MyDevicePortalService) GetVersion() (*ResponseMyDevicePortalGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/mydeviceportal/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseMyDevicePortalGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseMyDevicePortalGetVersion)
	return result, response, err

}

//CreateMyDevicePortal Create my device portal
/* This API creates a my device portal.

 */
func (s *MyDevicePortalService) CreateMyDevicePortal(requestMyDevicePortalCreateMyDevicePortal *RequestMyDevicePortalCreateMyDevicePortal) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/mydeviceportal"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestMyDevicePortalCreateMyDevicePortal).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateMyDevicePortal")
	}

	return response, err

}

//UpdateMyDevicePortalByID Update my device portal by ID
/* This API allows the client to update a my device portal by ID.

@param id id path parameter.
*/
func (s *MyDevicePortalService) UpdateMyDevicePortalByID(id string, requestMyDevicePortalUpdateMyDevicePortalById *RequestMyDevicePortalUpdateMyDevicePortalByID) (*ResponseMyDevicePortalUpdateMyDevicePortalByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/mydeviceportal/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestMyDevicePortalUpdateMyDevicePortalById).
		SetResult(&ResponseMyDevicePortalUpdateMyDevicePortalByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateMyDevicePortalById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseMyDevicePortalUpdateMyDevicePortalByID)
	return result, response, err

}

//DeleteMyDevicePortalByID Delete my device portal by ID
/* This API deletes a my device portal by ID.

@param id id path parameter.
*/
func (s *MyDevicePortalService) DeleteMyDevicePortalByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/mydeviceportal/{id}"
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
		return response, fmt.Errorf("error with operation DeleteMyDevicePortalById")
	}

	getCSFRToken(response.Header())

	return response, err

}
