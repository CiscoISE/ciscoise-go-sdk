package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type HotspotPortalService service

type GetHotspotPortalQueryParams struct {
	Page       int      `url: page,omitempty`       //Page number
	Size       int      `url: size,omitempty`       //Number of objects returned per page
	Sortasc    string   `url: sortasc,omitempty`    //sort asc
	Sortdsc    string   `url: sortdsc,omitempty`    //sort desc
	Filter     []string `url: filter,omitempty`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url: filterType,omitempty` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseHotspotPortalGetHotspotPortalByID struct {
	HotspotPortal ResponseHotspotPortalGetHotspotPortalByIDHotspotPortal `json:"HotspotPortal,omitempty"` //
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortal struct {
	ID             string                                                               `json:"id,omitempty"`             //
	Name           string                                                               `json:"name,omitempty"`           //
	Description    string                                                               `json:"description,omitempty"`    //
	PortalType     string                                                               `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                               `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a BYOD
	Customizations ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizations `json:"customizations,omitempty"` // Defines all of the Portal Customizations available
	Link           ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalLink           `json:"link,omitempty"`           //
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalSettings struct {
	PortalSettings           ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalSettingsPortalSettings           `json:"portalSettings,omitempty"`           // The port, interface, certificate, and other basic settings of a portal
	AupSettings              ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalSettingsAupSettings              `json:"aupSettings,omitempty"`              // Configuration of the Acceptable Use Policy (AUP) for a portal
	PostAccessBannerSettings ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalSettingsPostAccessBannerSettings `json:"postAccessBannerSettings,omitempty"` //
	AuthSuccessSettings      ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalSettingsAuthSuccessSettings      `json:"authSuccessSettings,omitempty"`      //
	PostLoginBannerSettings  ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalSettingsPostLoginBannerSettings  `json:"postLoginBannerSettings,omitempty"`  //
	SupportInfoSettings      ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalSettingsSupportInfoSettings      `json:"supportInfoSettings,omitempty"`      // Portal Support Information Settings
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalSettingsPortalSettings struct {
	HTTPSPort             int    `json:"httpsPort,omitempty"`             // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces     string `json:"allowedInterfaces,omitempty"`     // Interfaces that the portal will be reachable on. Allowed values: - eth0 - eth1 - eth2 - eth3 - eth4 - eth5 - bond0 - bond1 - bond2
	CertificateGroupTag   string `json:"certificateGroupTag,omitempty"`   // Logical name of the x.509 server certificate that will be used for the portal
	EndpointIDentityGroup string `json:"endpointIdentityGroup,omitempty"` // Unique Id of the endpoint identity group where user's devices will be added. Used only in Hotspot Portal
	CoaType               string `json:"coaType,omitempty"`               // Allowed Values: - COAREAUTHENTICATE, - COATERMINATE
	DisplayLang           string `json:"displayLang,omitempty"`           // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage      string `json:"fallbackLanguage,omitempty"`      // Used when displayLang = USEBROWSERLOCALE
	AlwaysUsedLanguage    string `json:"alwaysUsedLanguage,omitempty"`    // Used when displayLang = ALWAYSUSE
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalSettingsAupSettings struct {
	RequireAccessCode bool   `json:"requireAccessCode,omitempty"` // Require the portal user to enter an access code. Only used in Hotspot portal
	AccessCode        string `json:"accessCode,omitempty"`        // Access code that must be entered by the portal user (only valid if requireAccessCode = true)
	IncludeAup        bool   `json:"includeAup,omitempty"`        // Require the portal user to read and accept an AUP
	RequireScrolling  bool   `json:"requireScrolling,omitempty"`  // Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalSettingsPostAccessBannerSettings struct {
	IncludePostAccessBanner bool `json:"includePostAccessBanner,omitempty"` //
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalSettingsAuthSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` // After an Authentication Success where should device be redirected. Allowed values: - AUTHSUCCESSPAGE, - ORIGINATINGURL, - URL
	RedirectURL     string `json:"redirectUrl,omitempty"`     // Target URL for redirection, used when successRedirect = URL
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalSettingsPostLoginBannerSettings struct {
	IncludePostAccessBanner bool `json:"includePostAccessBanner,omitempty"` // Include a Post-Login Banner page
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  bool   `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          bool   `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        bool   `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent bool   `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     bool   `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      bool   `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizations struct {
	PortalTheme          ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          // Defines the configuration for portal theme
	PortalTweakSettings  ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` //
	PageCustomizations   ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   //
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        // The unique internal identifier of the portal theme
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BackgroundImage  ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerImage      ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BannerTitle      string                                                                                                   `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                                   `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                                   `json:"footerElement,omitempty"`    //
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsPageCustomizations struct {
	Data []ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsPageCustomizationsData `json:"data,omitempty"` //
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

type ResponseHotspotPortalGetHotspotPortalByIDHotspotPortalLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseHotspotPortalUpdateHotspotPortalByID struct {
	UpdatedFieldsList ResponseHotspotPortalUpdateHotspotPortalByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseHotspotPortalUpdateHotspotPortalByIDUpdatedFieldsList struct {
	UpdatedField ResponseHotspotPortalUpdateHotspotPortalByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                    `json:"field,omitempty"`        //
	OldValue     string                                                                    `json:"oldValue,omitempty"`     //
	NewValue     string                                                                    `json:"newValue,omitempty"`     //
}

type ResponseHotspotPortalUpdateHotspotPortalByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseHotspotPortalGetHotspotPortal struct {
	SearchResult ResponseHotspotPortalGetHotspotPortalSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseHotspotPortalGetHotspotPortalSearchResult struct {
	Total        int                                                             `json:"total,omitempty"`        //
	Resources    []ResponseHotspotPortalGetHotspotPortalSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseHotspotPortalGetHotspotPortalSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseHotspotPortalGetHotspotPortalSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseHotspotPortalGetHotspotPortalSearchResultResources struct {
	ID          string                                                         `json:"id,omitempty"`          //
	Name        string                                                         `json:"name,omitempty"`        //
	Description string                                                         `json:"description,omitempty"` //
	Link        ResponseHotspotPortalGetHotspotPortalSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseHotspotPortalGetHotspotPortalSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseHotspotPortalGetHotspotPortalSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseHotspotPortalGetHotspotPortalSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseHotspotPortalGetVersion struct {
	VersionInfo ResponseHotspotPortalGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseHotspotPortalGetVersionVersionInfo struct {
	CurrentServerVersion string                                         `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                         `json:"supportedVersions,omitempty"`    //
	Link                 ResponseHotspotPortalGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseHotspotPortalGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestHotspotPortalUpdateHotspotPortalByID struct {
	HotspotPortal *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortal `json:"HotspotPortal,omitempty"` //
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortal struct {
	ID             string                                                                  `json:"id,omitempty"`             //
	Name           string                                                                  `json:"name,omitempty"`           //
	Description    string                                                                  `json:"description,omitempty"`    //
	PortalType     string                                                                  `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                                  `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a BYOD
	Customizations *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizations `json:"customizations,omitempty"` // Defines all of the Portal Customizations available
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettings struct {
	PortalSettings           *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPortalSettings           `json:"portalSettings,omitempty"`           // The port, interface, certificate, and other basic settings of a portal
	AupSettings              *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsAupSettings              `json:"aupSettings,omitempty"`              // Configuration of the Acceptable Use Policy (AUP) for a portal
	PostAccessBannerSettings *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPostAccessBannerSettings `json:"postAccessBannerSettings,omitempty"` //
	AuthSuccessSettings      *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsAuthSuccessSettings      `json:"authSuccessSettings,omitempty"`      //
	PostLoginBannerSettings  *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPostLoginBannerSettings  `json:"postLoginBannerSettings,omitempty"`  //
	SupportInfoSettings      *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsSupportInfoSettings      `json:"supportInfoSettings,omitempty"`      // Portal Support Information Settings
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPortalSettings struct {
	HTTPSPort             *int   `json:"httpsPort,omitempty"`             // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces     string `json:"allowedInterfaces,omitempty"`     // Interfaces that the portal will be reachable on. Allowed values: - eth0 - eth1 - eth2 - eth3 - eth4 - eth5 - bond0 - bond1 - bond2
	CertificateGroupTag   string `json:"certificateGroupTag,omitempty"`   // Logical name of the x.509 server certificate that will be used for the portal
	EndpointIDentityGroup string `json:"endpointIdentityGroup,omitempty"` // Unique Id of the endpoint identity group where user's devices will be added. Used only in Hotspot Portal
	CoaType               string `json:"coaType,omitempty"`               // Allowed Values: - COAREAUTHENTICATE, - COATERMINATE
	DisplayLang           string `json:"displayLang,omitempty"`           // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage      string `json:"fallbackLanguage,omitempty"`      // Used when displayLang = USEBROWSERLOCALE
	AlwaysUsedLanguage    string `json:"alwaysUsedLanguage,omitempty"`    // Used when displayLang = ALWAYSUSE
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsAupSettings struct {
	RequireAccessCode *bool  `json:"requireAccessCode,omitempty"` // Require the portal user to enter an access code. Only used in Hotspot portal
	AccessCode        string `json:"accessCode,omitempty"`        // Access code that must be entered by the portal user (only valid if requireAccessCode = true)
	IncludeAup        *bool  `json:"includeAup,omitempty"`        // Require the portal user to read and accept an AUP
	RequireScrolling  *bool  `json:"requireScrolling,omitempty"`  // Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPostAccessBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` //
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsAuthSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` // After an Authentication Success where should device be redirected. Allowed values: - AUTHSUCCESSPAGE, - ORIGINATINGURL, - URL
	RedirectURL     string `json:"redirectUrl,omitempty"`     // Target URL for redirection, used when successRedirect = URL
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPostLoginBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` // Include a Post-Login Banner page
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  *bool  `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          *bool  `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        *bool  `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent *bool  `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     *bool  `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      *bool  `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizations struct {
	PortalTheme          *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          // Defines the configuration for portal theme
	PortalTweakSettings  *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` //
	PageCustomizations   *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   //
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        // The unique internal identifier of the portal theme
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BackgroundImage  *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerImage      *RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BannerTitle      string                                                                                                      `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                                      `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                                      `json:"footerElement,omitempty"`    //
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPageCustomizations struct {
	Data *[]RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPageCustomizationsData `json:"data,omitempty"` //
}

type RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

type RequestHotspotPortalCreateHotspotPortal struct {
	HotspotPortal *RequestHotspotPortalCreateHotspotPortalHotspotPortal `json:"HotspotPortal,omitempty"` //
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortal struct {
	Name           string                                                              `json:"name,omitempty"`           //
	Description    string                                                              `json:"description,omitempty"`    //
	PortalType     string                                                              `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                              `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       *RequestHotspotPortalCreateHotspotPortalHotspotPortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a BYOD
	Customizations *RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizations `json:"customizations,omitempty"` // Defines all of the Portal Customizations available
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalSettings struct {
	PortalSettings           *RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPortalSettings           `json:"portalSettings,omitempty"`           // The port, interface, certificate, and other basic settings of a portal
	AupSettings              *RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsAupSettings              `json:"aupSettings,omitempty"`              // Configuration of the Acceptable Use Policy (AUP) for a portal
	PostAccessBannerSettings *RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPostAccessBannerSettings `json:"postAccessBannerSettings,omitempty"` //
	AuthSuccessSettings      *RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsAuthSuccessSettings      `json:"authSuccessSettings,omitempty"`      //
	PostLoginBannerSettings  *RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPostLoginBannerSettings  `json:"postLoginBannerSettings,omitempty"`  //
	SupportInfoSettings      *RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsSupportInfoSettings      `json:"supportInfoSettings,omitempty"`      // Portal Support Information Settings
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPortalSettings struct {
	HTTPSPort             *int   `json:"httpsPort,omitempty"`             // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces     string `json:"allowedInterfaces,omitempty"`     // Interfaces that the portal will be reachable on. Allowed values: - eth0 - eth1 - eth2 - eth3 - eth4 - eth5 - bond0 - bond1 - bond2
	CertificateGroupTag   string `json:"certificateGroupTag,omitempty"`   // Logical name of the x.509 server certificate that will be used for the portal
	EndpointIDentityGroup string `json:"endpointIdentityGroup,omitempty"` // Unique Id of the endpoint identity group where user's devices will be added. Used only in Hotspot Portal
	CoaType               string `json:"coaType,omitempty"`               // Allowed Values: - COAREAUTHENTICATE, - COATERMINATE
	DisplayLang           string `json:"displayLang,omitempty"`           // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage      string `json:"fallbackLanguage,omitempty"`      // Used when displayLang = USEBROWSERLOCALE
	AlwaysUsedLanguage    string `json:"alwaysUsedLanguage,omitempty"`    // Used when displayLang = ALWAYSUSE
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsAupSettings struct {
	RequireAccessCode *bool  `json:"requireAccessCode,omitempty"` // Require the portal user to enter an access code. Only used in Hotspot portal
	AccessCode        string `json:"accessCode,omitempty"`        // Access code that must be entered by the portal user (only valid if requireAccessCode = true)
	IncludeAup        *bool  `json:"includeAup,omitempty"`        // Require the portal user to read and accept an AUP
	RequireScrolling  *bool  `json:"requireScrolling,omitempty"`  // Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPostAccessBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` //
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsAuthSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` // After an Authentication Success where should device be redirected. Allowed values: - AUTHSUCCESSPAGE, - ORIGINATINGURL, - URL
	RedirectURL     string `json:"redirectUrl,omitempty"`     // Target URL for redirection, used when successRedirect = URL
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPostLoginBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` // Include a Post-Login Banner page
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  *bool  `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          *bool  `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        *bool  `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent *bool  `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     *bool  `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      *bool  `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizations struct {
	PortalTheme          *RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          // Defines the configuration for portal theme
	PortalTweakSettings  *RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             *RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations *RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` //
	PageCustomizations   *RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   //
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        // The unique internal identifier of the portal theme
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  *RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage *RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BackgroundImage  *RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerImage      *RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BannerTitle      string                                                                                                  `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                                  `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                                  `json:"footerElement,omitempty"`    //
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPageCustomizations struct {
	Data *[]RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPageCustomizationsData `json:"data,omitempty"` //
}

type RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

//GetHotspotPortalByID Get hotspot portal by ID
/* This API allows the client to get a hotspot portal by ID.

@param id id path parameter.
*/
func (s *HotspotPortalService) GetHotspotPortalByID(id string) (*ResponseHotspotPortalGetHotspotPortalByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/hotspotportal/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseHotspotPortalGetHotspotPortalByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetHotspotPortalById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseHotspotPortalGetHotspotPortalByID)
	return result, response, err

}

//GetHotspotPortal Get all hotspot portals
/* This API allows the client to get all the hotspot portals.

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

@param getHotspotPortalQueryParams Filtering parameter
*/
func (s *HotspotPortalService) GetHotspotPortal(getHotspotPortalQueryParams *GetHotspotPortalQueryParams) (*ResponseHotspotPortalGetHotspotPortal, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/hotspotportal"

	queryString, _ := query.Values(getHotspotPortalQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHotspotPortalGetHotspotPortal{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetHotspotPortal")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseHotspotPortalGetHotspotPortal)
	return result, response, err

}

//GetVersion Get hotspot portal version information
/* This API helps to retrieve the version information related to the hotspot portal.

 */
func (s *HotspotPortalService) GetVersion() (*ResponseHotspotPortalGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/hotspotportal/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseHotspotPortalGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseHotspotPortalGetVersion)
	return result, response, err

}

//CreateHotspotPortal Create hotspot portal
/* This API creates a hotspot portal.

 */
func (s *HotspotPortalService) CreateHotspotPortal(requestHotspotPortalCreateHotspotPortal *RequestHotspotPortalCreateHotspotPortal) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/hotspotportal"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestHotspotPortalCreateHotspotPortal).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateHotspotPortal")
	}

	return response, err

}

//UpdateHotspotPortalByID Update hotspot portal by ID
/* This API allows the client to update a hotspot portal by ID.

@param id id path parameter.
*/
func (s *HotspotPortalService) UpdateHotspotPortalByID(id string, requestHotspotPortalUpdateHotspotPortalById *RequestHotspotPortalUpdateHotspotPortalByID) (*ResponseHotspotPortalUpdateHotspotPortalByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/hotspotportal/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestHotspotPortalUpdateHotspotPortalById).
		SetResult(&ResponseHotspotPortalUpdateHotspotPortalByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateHotspotPortalById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseHotspotPortalUpdateHotspotPortalByID)
	return result, response, err

}

//DeleteHotspotPortalByID Delete hotspot portal by ID
/* This API deletes a hotspot portal by ID.

@param id id path parameter.
*/
func (s *HotspotPortalService) DeleteHotspotPortalByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/hotspotportal/{id}"
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
		return response, fmt.Errorf("error with operation DeleteHotspotPortalById")
	}

	getCSFRToken(response.Header())

	return response, err

}
