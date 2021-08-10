package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SponsorPortalService service

type GetSponsorPortalQueryParams struct {
	Page       int      `url: page,omitempty`       //Page number
	Size       int      `url: size,omitempty`       //Number of objects returned per page
	Sortasc    string   `url: sortasc,omitempty`    //sort asc
	Sortdsc    string   `url: sortdsc,omitempty`    //sort desc
	Filter     []string `url: filter,omitempty`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url: filterType,omitempty` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseSponsorPortalGetSponsorPortalByID struct {
	SponsorPortal ResponseSponsorPortalGetSponsorPortalByIDSponsorPortal `json:"SponsorPortal,omitempty"` //
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortal struct {
	ID             string                                                               `json:"id,omitempty"`             //
	Name           string                                                               `json:"name,omitempty"`           //
	Description    string                                                               `json:"description,omitempty"`    //
	PortalType     string                                                               `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                               `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a portal
	Customizations ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizations `json:"customizations,omitempty"` //
	Link           ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalLink           `json:"link,omitempty"`           //
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettings struct {
	PortalSettings                ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsPortalSettings                `json:"portalSettings,omitempty"`                // The port, interface, certificate, and other basic settings of a portal
	LoginPageSettings             ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsLoginPageSettings             `json:"loginPageSettings,omitempty"`             // Portal Login Page settings groups follow
	AupSettings                   ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsAupSettings                   `json:"aupSettings,omitempty"`                   //
	SponsorChangePasswordSettings ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsSponsorChangePasswordSettings `json:"sponsorChangePasswordSettings,omitempty"` //
	PostLoginBannerSettings       ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsPostLoginBannerSettings       `json:"postLoginBannerSettings,omitempty"`       //
	PostAccessBannerSettings      ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsPostAccessBannerSettings      `json:"postAccessBannerSettings,omitempty"`      //
	SupportInfoSettings           ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsSupportInfoSettings           `json:"supportInfoSettings,omitempty"`           //
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsPortalSettings struct {
	HTTPSPort            int    `json:"httpsPort,omitempty"`            // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces    string `json:"allowedInterfaces,omitempty"`    // Interfaces that the portal will be reachable on. Allowed values: - eth0, - eth1, - eth2, - eth3, - eth4, - eth5, - bond0, - bond1, - bond2
	CertificateGroupTag  string `json:"certificateGroupTag,omitempty"`  // Logical name of the x.509 server certificate that will be used for the portal
	Fqdn                 string `json:"fqdn,omitempty"`                 // The fully-qualified domain name (FQDN) that end-users will use to access this portal. Used only in Sponsor portal
	AuthenticationMethod string `json:"authenticationMethod,omitempty"` // Unique Id of the identity source sequence
	IDleTimeout          int    `json:"idleTimeout,omitempty"`          //
	DisplayLang          string `json:"displayLang,omitempty"`          // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage     string `json:"fallbackLanguage,omitempty"`     // Used when displayLang = USEBROWSERLOCALE
	AvailableSSIDs       string `json:"availableSsids,omitempty"`       // Names of the SSIDs available for assignment to guest users by sponsors
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsLoginPageSettings struct {
	MaxFailedAttemptsBeforeRateLimit int           `json:"maxFailedAttemptsBeforeRateLimit,omitempty"` // Maximum failed login attempts before rate limiting
	TimeBetweenLoginsDuringRateLimit int           `json:"timeBetweenLoginsDuringRateLimit,omitempty"` // Time between login attempts when rate limiting
	IncludeAup                       bool          `json:"includeAup,omitempty"`                       // Include an Acceptable Use Policy (AUP) that should be displayed during login
	AupDisplay                       string        `json:"aupDisplay,omitempty"`                       // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: -  ONPAGE, - ASLINK
	RequireAupAcceptance             bool          `json:"requireAupAcceptance,omitempty"`             // Require the portal user to accept the AUP. Only valid if includeAup = true
	RequireAupScrolling              bool          `json:"requireAupScrolling,omitempty"`              //
	SocialConfigs                    []interface{} `json:"socialConfigs,omitempty"`                    //
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsAupSettings struct {
	IncludeAup                   bool   `json:"includeAup,omitempty"`                   //
	RequireScrolling             bool   `json:"requireScrolling,omitempty"`             //
	DisplayFrequency             string `json:"displayFrequency,omitempty"`             // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed Values: - FIRSTLOGIN, - EVERYLOGIN, - RECURRING
	DisplayFrequencyIntervalDays int    `json:"displayFrequencyIntervalDays,omitempty"` // Number of days between AUP confirmations (when displayFrequency = recurring)
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsSponsorChangePasswordSettings struct {
	AllowSponsorToChangePwd bool `json:"allowSponsorToChangePwd,omitempty"` // Allow sponsors to change their own passwords
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsPostLoginBannerSettings struct {
	IncludePostAccessBanner bool `json:"includePostAccessBanner,omitempty"` // Include a Post-Login Banner page
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsPostAccessBannerSettings struct {
	IncludePostAccessBanner bool `json:"includePostAccessBanner,omitempty"` //
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  bool   `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          bool   `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        bool   `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent bool   `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     bool   `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      bool   `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizations struct {
	PortalTheme          ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          //
	PortalTweakSettings  ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` //
	PageCustomizations   ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   //
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BannerImage      ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BackgroundImage  ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerTitle      string                                                                                                   `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                                   `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                                   `json:"footerElement,omitempty"`    //
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsPageCustomizations struct {
	Data []ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsPageCustomizationsData `json:"data,omitempty"` // The Dictionary will be exposed here as key value pair
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

type ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSponsorPortalUpdateSponsorPortalByID struct {
	UpdatedFieldsList ResponseSponsorPortalUpdateSponsorPortalByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseSponsorPortalUpdateSponsorPortalByIDUpdatedFieldsList struct {
	UpdatedField ResponseSponsorPortalUpdateSponsorPortalByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                    `json:"field,omitempty"`        //
	OldValue     string                                                                    `json:"oldValue,omitempty"`     //
	NewValue     string                                                                    `json:"newValue,omitempty"`     //
}

type ResponseSponsorPortalUpdateSponsorPortalByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseSponsorPortalGetSponsorPortal struct {
	SearchResult ResponseSponsorPortalGetSponsorPortalSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseSponsorPortalGetSponsorPortalSearchResult struct {
	Total        int                                                             `json:"total,omitempty"`        //
	Resources    []ResponseSponsorPortalGetSponsorPortalSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseSponsorPortalGetSponsorPortalSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseSponsorPortalGetSponsorPortalSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseSponsorPortalGetSponsorPortalSearchResultResources struct {
	ID          string                                                         `json:"id,omitempty"`          //
	Name        string                                                         `json:"name,omitempty"`        //
	Description string                                                         `json:"description,omitempty"` //
	Link        ResponseSponsorPortalGetSponsorPortalSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseSponsorPortalGetSponsorPortalSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSponsorPortalGetSponsorPortalSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSponsorPortalGetSponsorPortalSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSponsorPortalGetVersion struct {
	VersionInfo ResponseSponsorPortalGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseSponsorPortalGetVersionVersionInfo struct {
	CurrentServerVersion string                                         `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                         `json:"supportedVersions,omitempty"`    //
	Link                 ResponseSponsorPortalGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseSponsorPortalGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestSponsorPortalUpdateSponsorPortalByID struct {
	SponsorPortal *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortal `json:"SponsorPortal,omitempty"` //
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortal struct {
	ID             string                                                                  `json:"id,omitempty"`             //
	Name           string                                                                  `json:"name,omitempty"`           //
	Description    string                                                                  `json:"description,omitempty"`    //
	PortalType     string                                                                  `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                                  `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a portal
	Customizations *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizations `json:"customizations,omitempty"` //
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettings struct {
	PortalSettings                *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPortalSettings                `json:"portalSettings,omitempty"`                // The port, interface, certificate, and other basic settings of a portal
	LoginPageSettings             *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsLoginPageSettings             `json:"loginPageSettings,omitempty"`             // Portal Login Page settings groups follow
	AupSettings                   *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsAupSettings                   `json:"aupSettings,omitempty"`                   //
	SponsorChangePasswordSettings *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsSponsorChangePasswordSettings `json:"sponsorChangePasswordSettings,omitempty"` //
	PostLoginBannerSettings       *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPostLoginBannerSettings       `json:"postLoginBannerSettings,omitempty"`       //
	PostAccessBannerSettings      *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPostAccessBannerSettings      `json:"postAccessBannerSettings,omitempty"`      //
	SupportInfoSettings           *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsSupportInfoSettings           `json:"supportInfoSettings,omitempty"`           //
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPortalSettings struct {
	HTTPSPort            *int   `json:"httpsPort,omitempty"`            // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces    string `json:"allowedInterfaces,omitempty"`    // Interfaces that the portal will be reachable on. Allowed values: - eth0, - eth1, - eth2, - eth3, - eth4, - eth5, - bond0, - bond1, - bond2
	CertificateGroupTag  string `json:"certificateGroupTag,omitempty"`  // Logical name of the x.509 server certificate that will be used for the portal
	Fqdn                 string `json:"fqdn,omitempty"`                 // The fully-qualified domain name (FQDN) that end-users will use to access this portal. Used only in Sponsor portal
	AuthenticationMethod string `json:"authenticationMethod,omitempty"` // Unique Id of the identity source sequence
	IDleTimeout          *int   `json:"idleTimeout,omitempty"`          //
	DisplayLang          string `json:"displayLang,omitempty"`          // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage     string `json:"fallbackLanguage,omitempty"`     // Used when displayLang = USEBROWSERLOCALE
	AvailableSSIDs       string `json:"availableSsids,omitempty"`       // Names of the SSIDs available for assignment to guest users by sponsors
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsLoginPageSettings struct {
	MaxFailedAttemptsBeforeRateLimit *int           `json:"maxFailedAttemptsBeforeRateLimit,omitempty"` // Maximum failed login attempts before rate limiting
	TimeBetweenLoginsDuringRateLimit *int           `json:"timeBetweenLoginsDuringRateLimit,omitempty"` // Time between login attempts when rate limiting
	IncludeAup                       *bool          `json:"includeAup,omitempty"`                       // Include an Acceptable Use Policy (AUP) that should be displayed during login
	AupDisplay                       string         `json:"aupDisplay,omitempty"`                       // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: -  ONPAGE, - ASLINK
	RequireAupAcceptance             *bool          `json:"requireAupAcceptance,omitempty"`             // Require the portal user to accept the AUP. Only valid if includeAup = true
	RequireAupScrolling              *bool          `json:"requireAupScrolling,omitempty"`              //
	SocialConfigs                    *[]interface{} `json:"socialConfigs,omitempty"`                    //
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsAupSettings struct {
	IncludeAup                   *bool  `json:"includeAup,omitempty"`                   //
	RequireScrolling             *bool  `json:"requireScrolling,omitempty"`             //
	DisplayFrequency             string `json:"displayFrequency,omitempty"`             // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed Values: - FIRSTLOGIN, - EVERYLOGIN, - RECURRING
	DisplayFrequencyIntervalDays *int   `json:"displayFrequencyIntervalDays,omitempty"` // Number of days between AUP confirmations (when displayFrequency = recurring)
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsSponsorChangePasswordSettings struct {
	AllowSponsorToChangePwd *bool `json:"allowSponsorToChangePwd,omitempty"` // Allow sponsors to change their own passwords
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPostLoginBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` // Include a Post-Login Banner page
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPostAccessBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` //
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  *bool  `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          *bool  `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        *bool  `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent *bool  `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     *bool  `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      *bool  `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizations struct {
	PortalTheme          *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          //
	PortalTweakSettings  *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` //
	PageCustomizations   *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   //
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BannerImage      *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BackgroundImage  *RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerTitle      string                                                                                                      `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                                      `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                                      `json:"footerElement,omitempty"`    //
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPageCustomizations struct {
	Data *[]RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPageCustomizationsData `json:"data,omitempty"` // The Dictionary will be exposed here as key value pair
}

type RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

type RequestSponsorPortalCreateSponsorPortal struct {
	SponsorPortal *RequestSponsorPortalCreateSponsorPortalSponsorPortal `json:"SponsorPortal,omitempty"` //
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortal struct {
	Name           string                                                              `json:"name,omitempty"`           //
	Description    string                                                              `json:"description,omitempty"`    //
	PortalType     string                                                              `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                              `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       *RequestSponsorPortalCreateSponsorPortalSponsorPortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a portal
	Customizations *RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizations `json:"customizations,omitempty"` //
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalSettings struct {
	PortalSettings                *RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPortalSettings                `json:"portalSettings,omitempty"`                // The port, interface, certificate, and other basic settings of a portal
	LoginPageSettings             *RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsLoginPageSettings             `json:"loginPageSettings,omitempty"`             // Portal Login Page settings groups follow
	AupSettings                   *RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsAupSettings                   `json:"aupSettings,omitempty"`                   //
	SponsorChangePasswordSettings *RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsSponsorChangePasswordSettings `json:"sponsorChangePasswordSettings,omitempty"` //
	PostLoginBannerSettings       *RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPostLoginBannerSettings       `json:"postLoginBannerSettings,omitempty"`       //
	PostAccessBannerSettings      *RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPostAccessBannerSettings      `json:"postAccessBannerSettings,omitempty"`      //
	SupportInfoSettings           *RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsSupportInfoSettings           `json:"supportInfoSettings,omitempty"`           //
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPortalSettings struct {
	HTTPSPort            *int   `json:"httpsPort,omitempty"`            // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces    string `json:"allowedInterfaces,omitempty"`    // Interfaces that the portal will be reachable on. Allowed values: - eth0, - eth1, - eth2, - eth3, - eth4, - eth5, - bond0, - bond1, - bond2
	CertificateGroupTag  string `json:"certificateGroupTag,omitempty"`  // Logical name of the x.509 server certificate that will be used for the portal
	Fqdn                 string `json:"fqdn,omitempty"`                 // The fully-qualified domain name (FQDN) that end-users will use to access this portal. Used only in Sponsor portal
	AuthenticationMethod string `json:"authenticationMethod,omitempty"` // Unique Id of the identity source sequence
	IDleTimeout          *int   `json:"idleTimeout,omitempty"`          //
	DisplayLang          string `json:"displayLang,omitempty"`          // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage     string `json:"fallbackLanguage,omitempty"`     // Used when displayLang = USEBROWSERLOCALE
	AvailableSSIDs       string `json:"availableSsids,omitempty"`       // Names of the SSIDs available for assignment to guest users by sponsors
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsLoginPageSettings struct {
	MaxFailedAttemptsBeforeRateLimit *int           `json:"maxFailedAttemptsBeforeRateLimit,omitempty"` // Maximum failed login attempts before rate limiting
	TimeBetweenLoginsDuringRateLimit *int           `json:"timeBetweenLoginsDuringRateLimit,omitempty"` // Time between login attempts when rate limiting
	IncludeAup                       *bool          `json:"includeAup,omitempty"`                       // Include an Acceptable Use Policy (AUP) that should be displayed during login
	AupDisplay                       string         `json:"aupDisplay,omitempty"`                       // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: -  ONPAGE, - ASLINK
	RequireAupAcceptance             *bool          `json:"requireAupAcceptance,omitempty"`             // Require the portal user to accept the AUP. Only valid if includeAup = true
	RequireAupScrolling              *bool          `json:"requireAupScrolling,omitempty"`              //
	SocialConfigs                    *[]interface{} `json:"socialConfigs,omitempty"`                    //
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsAupSettings struct {
	IncludeAup                   *bool  `json:"includeAup,omitempty"`                   //
	RequireScrolling             *bool  `json:"requireScrolling,omitempty"`             //
	DisplayFrequency             string `json:"displayFrequency,omitempty"`             // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed Values: - FIRSTLOGIN, - EVERYLOGIN, - RECURRING
	DisplayFrequencyIntervalDays *int   `json:"displayFrequencyIntervalDays,omitempty"` // Number of days between AUP confirmations (when displayFrequency = recurring)
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsSponsorChangePasswordSettings struct {
	AllowSponsorToChangePwd *bool `json:"allowSponsorToChangePwd,omitempty"` // Allow sponsors to change their own passwords
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPostLoginBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` // Include a Post-Login Banner page
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPostAccessBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` //
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  *bool  `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          *bool  `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        *bool  `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent *bool  `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     *bool  `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      *bool  `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizations struct {
	PortalTheme          *RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          //
	PortalTweakSettings  *RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             *RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations *RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` //
	PageCustomizations   *RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   //
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  *RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage *RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BannerImage      *RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BackgroundImage  *RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerTitle      string                                                                                                  `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                                  `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                                  `json:"footerElement,omitempty"`    //
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPageCustomizations struct {
	Data *[]RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPageCustomizationsData `json:"data,omitempty"` // The Dictionary will be exposed here as key value pair
}

type RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

//GetSponsorPortalByID Get sponsor portal by ID
/* This API allows the client to get a sponsor portal by ID.

@param id id path parameter.
*/
func (s *SponsorPortalService) GetSponsorPortalByID(id string) (*ResponseSponsorPortalGetSponsorPortalByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsorportal/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSponsorPortalGetSponsorPortalByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSponsorPortalById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSponsorPortalGetSponsorPortalByID)
	return result, response, err

}

//GetSponsorPortal Get all sponsor portals
/* This API allows the client to get all the sponsor portals.

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

@param getSponsorPortalQueryParams Filtering parameter
*/
func (s *SponsorPortalService) GetSponsorPortal(getSponsorPortalQueryParams *GetSponsorPortalQueryParams) (*ResponseSponsorPortalGetSponsorPortal, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsorportal"

	queryString, _ := query.Values(getSponsorPortalQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSponsorPortalGetSponsorPortal{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSponsorPortal")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSponsorPortalGetSponsorPortal)
	return result, response, err

}

//GetVersion Get sponsor portal version information
/* This API helps to retrieve the version information related to the sponsor portal.

 */
func (s *SponsorPortalService) GetVersion() (*ResponseSponsorPortalGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsorportal/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSponsorPortalGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSponsorPortalGetVersion)
	return result, response, err

}

//CreateSponsorPortal Create sponsor portal
/* This API creates a sponsor portal.

 */
func (s *SponsorPortalService) CreateSponsorPortal(requestSponsorPortalCreateSponsorPortal *RequestSponsorPortalCreateSponsorPortal) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsorportal"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSponsorPortalCreateSponsorPortal).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateSponsorPortal")
	}

	return response, err

}

//UpdateSponsorPortalByID Update sponsor portal by ID
/* This API allows the client to update a sponsor portal by ID.

@param id id path parameter.
*/
func (s *SponsorPortalService) UpdateSponsorPortalByID(id string, requestSponsorPortalUpdateSponsorPortalById *RequestSponsorPortalUpdateSponsorPortalByID) (*ResponseSponsorPortalUpdateSponsorPortalByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsorportal/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSponsorPortalUpdateSponsorPortalById).
		SetResult(&ResponseSponsorPortalUpdateSponsorPortalByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSponsorPortalById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSponsorPortalUpdateSponsorPortalByID)
	return result, response, err

}

//DeleteSponsorPortalByID Delete sponsor portal by ID
/* This API deletes a sponsor portal by ID.

@param id id path parameter.
*/
func (s *SponsorPortalService) DeleteSponsorPortalByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsorportal/{id}"
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
		return response, fmt.Errorf("error with operation DeleteSponsorPortalById")
	}

	getCSFRToken(response.Header())

	return response, err

}
