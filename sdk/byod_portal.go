package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ByodPortalService service

type GetByodPortalQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseByodPortalGetByodPortalByID struct {
	ByodPortal *ResponseByodPortalGetByodPortalByIDByodPortal `json:"BYODPortal,omitempty"` //
}

type ResponseByodPortalGetByodPortalByIDByodPortal struct {
	ID             string                                                       `json:"id,omitempty"`             // Resource UUID, mandatory for update
	Name           string                                                       `json:"name,omitempty"`           // Resource Name
	Description    string                                                       `json:"description,omitempty"`    //
	PortalType     string                                                       `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                       `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       *ResponseByodPortalGetByodPortalByIDByodPortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a BYOD
	Customizations *ResponseByodPortalGetByodPortalByIDByodPortalCustomizations `json:"customizations,omitempty"` // Defines all of the Portal Customizations available for a BYOD
	Link           *ResponseByodPortalGetByodPortalByIDByodPortalLink           `json:"link,omitempty"`           //
}

type ResponseByodPortalGetByodPortalByIDByodPortalSettings struct {
	PortalSettings      *ResponseByodPortalGetByodPortalByIDByodPortalSettingsPortalSettings      `json:"portalSettings,omitempty"`      // The port, interface, certificate, and other basic settings of a portal
	ByodSettings        *ResponseByodPortalGetByodPortalByIDByodPortalSettingsByodSettings        `json:"byodSettings,omitempty"`        // Configuration of BYOD Device Welcome, Registration and Success steps
	SupportInfoSettings *ResponseByodPortalGetByodPortalByIDByodPortalSettingsSupportInfoSettings `json:"supportInfoSettings,omitempty"` //
}

type ResponseByodPortalGetByodPortalByIDByodPortalSettingsPortalSettings struct {
	HTTPSPort             *int     `json:"httpsPort,omitempty"`             // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces     []string `json:"allowedInterfaces,omitempty"`     // Interfaces that the portal will be reachable on. Allowed values: - eth0, - eth1, - eth2, - eth3, - eth4, - eth5, - bond0, - bond1, - bond2
	CertificateGroupTag   string   `json:"certificateGroupTag,omitempty"`   // Logical name of the x.509 server certificate that will be used for the portal
	EndpointIDentityGroup string   `json:"endpointIdentityGroup,omitempty"` // Unique Id of the endpoint identity group where user's devices will be added. Used only in Hotspot Portal
	DisplayLang           string   `json:"displayLang,omitempty"`           // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage      string   `json:"fallbackLanguage,omitempty"`      // Used when displayLang = USEBROWSERLOCALE
	AlwaysUsedLanguage    string   `json:"alwaysUsedLanguage,omitempty"`    // Used when displayLang = ALWAYSUSE
}

type ResponseByodPortalGetByodPortalByIDByodPortalSettingsByodSettings struct {
	ByodWelcomeSettings             *ResponseByodPortalGetByodPortalByIDByodPortalSettingsByodSettingsByodWelcomeSettings             `json:"byodWelcomeSettings,omitempty"`             // Configuration of BYOD endpoint welcome step configuration
	ByodRegistrationSettings        *ResponseByodPortalGetByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSettings        `json:"byodRegistrationSettings,omitempty"`        //
	ByodRegistrationSuccessSettings *ResponseByodPortalGetByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSuccessSettings `json:"byodRegistrationSuccessSettings,omitempty"` //
}

type ResponseByodPortalGetByodPortalByIDByodPortalSettingsByodSettingsByodWelcomeSettings struct {
	EnableByod           *bool  `json:"enableBYOD,omitempty"`           //
	EnableGuestAccess    *bool  `json:"enableGuestAccess,omitempty"`    //
	RequireMdm           *bool  `json:"requireMDM,omitempty"`           //
	IncludeAup           *bool  `json:"includeAup,omitempty"`           //
	AupDisplay           string `json:"aupDisplay,omitempty"`           // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: - ONPAGE, - ASLINK
	RequireAupAcceptance *bool  `json:"requireAupAcceptance,omitempty"` //
	RequireScrolling     *bool  `json:"requireScrolling,omitempty"`     // Require BYOD devices to scroll down to the bottom of the AUP, Only valid if includeAup = true
}

type ResponseByodPortalGetByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSettings struct {
	ShowDeviceID            *bool  `json:"showDeviceID,omitempty"`            //
	EndPointIDentityGroupID string `json:"endPointIdentityGroupId,omitempty"` //
}

type ResponseByodPortalGetByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` // After an Authentication Success where should device be redirected. Allowed values:
	RedirectURL     string `json:"redirectUrl,omitempty"`     // Target URL for redirection, used when successRedirect = URL
}

type ResponseByodPortalGetByodPortalByIDByodPortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  *bool  `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          *bool  `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        *bool  `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent *bool  `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     *bool  `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      *bool  `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type ResponseByodPortalGetByodPortalByIDByodPortalCustomizations struct {
	PortalTheme          *ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          // Defines the configuration for portal theme
	PortalTweakSettings  *ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             *ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations *ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` // Represent the portal Global customizations
	PageCustomizations   *ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   // Represent the entire page customization as a giant dictionary
}

type ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        // The unique internal identifier of the portal theme
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  *ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage *ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BannerImage      *ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BackgroundImage  *ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerTitle      string                                                                                           `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                           `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                           `json:"footerElement,omitempty"`    //
}

type ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsPageCustomizations struct {
	Data *[]ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsPageCustomizationsData `json:"data,omitempty"` // The Dictionary will be exposed here as key value pair
}

type ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

type ResponseByodPortalGetByodPortalByIDByodPortalLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseByodPortalUpdateByodPortalByID struct {
	UpdatedFieldsList *ResponseByodPortalUpdateByodPortalByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseByodPortalUpdateByodPortalByIDUpdatedFieldsList struct {
	UpdatedField *[]ResponseByodPortalUpdateByodPortalByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                 `json:"field,omitempty"`        //
	OldValue     string                                                                 `json:"oldValue,omitempty"`     //
	NewValue     string                                                                 `json:"newValue,omitempty"`     //
}

type ResponseByodPortalUpdateByodPortalByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseByodPortalGetByodPortal struct {
	SearchResult *ResponseByodPortalGetByodPortalSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseByodPortalGetByodPortalSearchResult struct {
	Total        *int                                                     `json:"total,omitempty"`        //
	Resources    *[]ResponseByodPortalGetByodPortalSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponseByodPortalGetByodPortalSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseByodPortalGetByodPortalSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseByodPortalGetByodPortalSearchResultResources struct {
	ID          string                                                    `json:"id,omitempty"`          //
	Name        string                                                    `json:"name,omitempty"`        //
	Description string                                                    `json:"description,omitempty"` //
	Link        *ResponseByodPortalGetByodPortalSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseByodPortalGetByodPortalSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseByodPortalGetByodPortalSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseByodPortalGetByodPortalSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseByodPortalGetVersion struct {
	VersionInfo *ResponseByodPortalGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseByodPortalGetVersionVersionInfo struct {
	CurrentServerVersion string                                       `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                       `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseByodPortalGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseByodPortalGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestByodPortalUpdateByodPortalByID struct {
	ByodPortal *RequestByodPortalUpdateByodPortalByIDByodPortal `json:"BYODPortal,omitempty"` //
}

type RequestByodPortalUpdateByodPortalByIDByodPortal struct {
	ID             string                                                         `json:"id,omitempty"`             // Resource UUID, mandatory for update
	Name           string                                                         `json:"name,omitempty"`           // Resource Name
	Description    string                                                         `json:"description,omitempty"`    //
	PortalType     string                                                         `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                         `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       *RequestByodPortalUpdateByodPortalByIDByodPortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a BYOD
	Customizations *RequestByodPortalUpdateByodPortalByIDByodPortalCustomizations `json:"customizations,omitempty"` // Defines all of the Portal Customizations available for a BYOD
}

type RequestByodPortalUpdateByodPortalByIDByodPortalSettings struct {
	PortalSettings      *RequestByodPortalUpdateByodPortalByIDByodPortalSettingsPortalSettings      `json:"portalSettings,omitempty"`      // The port, interface, certificate, and other basic settings of a portal
	ByodSettings        *RequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettings        `json:"byodSettings,omitempty"`        // Configuration of BYOD Device Welcome, Registration and Success steps
	SupportInfoSettings *RequestByodPortalUpdateByodPortalByIDByodPortalSettingsSupportInfoSettings `json:"supportInfoSettings,omitempty"` //
}

type RequestByodPortalUpdateByodPortalByIDByodPortalSettingsPortalSettings struct {
	HTTPSPort             *int     `json:"httpsPort,omitempty"`             // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces     []string `json:"allowedInterfaces,omitempty"`     // Interfaces that the portal will be reachable on. Allowed values: - eth0, - eth1, - eth2, - eth3, - eth4, - eth5, - bond0, - bond1, - bond2
	CertificateGroupTag   string   `json:"certificateGroupTag,omitempty"`   // Logical name of the x.509 server certificate that will be used for the portal
	EndpointIDentityGroup string   `json:"endpointIdentityGroup,omitempty"` // Unique Id of the endpoint identity group where user's devices will be added. Used only in Hotspot Portal
	DisplayLang           string   `json:"displayLang,omitempty"`           // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage      string   `json:"fallbackLanguage,omitempty"`      // Used when displayLang = USEBROWSERLOCALE
	AlwaysUsedLanguage    string   `json:"alwaysUsedLanguage,omitempty"`    // Used when displayLang = ALWAYSUSE
}

type RequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettings struct {
	ByodWelcomeSettings             *RequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodWelcomeSettings             `json:"byodWelcomeSettings,omitempty"`             // Configuration of BYOD endpoint welcome step configuration
	ByodRegistrationSettings        *RequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSettings        `json:"byodRegistrationSettings,omitempty"`        //
	ByodRegistrationSuccessSettings *RequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSuccessSettings `json:"byodRegistrationSuccessSettings,omitempty"` //
}

type RequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodWelcomeSettings struct {
	EnableByod           *bool  `json:"enableBYOD,omitempty"`           //
	EnableGuestAccess    *bool  `json:"enableGuestAccess,omitempty"`    //
	RequireMdm           *bool  `json:"requireMDM,omitempty"`           //
	IncludeAup           *bool  `json:"includeAup,omitempty"`           //
	AupDisplay           string `json:"aupDisplay,omitempty"`           // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: - ONPAGE, - ASLINK
	RequireAupAcceptance *bool  `json:"requireAupAcceptance,omitempty"` //
	RequireScrolling     *bool  `json:"requireScrolling,omitempty"`     // Require BYOD devices to scroll down to the bottom of the AUP, Only valid if includeAup = true
}

type RequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSettings struct {
	ShowDeviceID            *bool  `json:"showDeviceID,omitempty"`            //
	EndPointIDentityGroupID string `json:"endPointIdentityGroupId,omitempty"` //
}

type RequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` // After an Authentication Success where should device be redirected. Allowed values:
	RedirectURL     string `json:"redirectUrl,omitempty"`     // Target URL for redirection, used when successRedirect = URL
}

type RequestByodPortalUpdateByodPortalByIDByodPortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  *bool  `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          *bool  `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        *bool  `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent *bool  `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     *bool  `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      *bool  `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type RequestByodPortalUpdateByodPortalByIDByodPortalCustomizations struct {
	PortalTheme          *RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          // Defines the configuration for portal theme
	PortalTweakSettings  *RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             *RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations *RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` // Represent the portal Global customizations
	PageCustomizations   *RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   // Represent the entire page customization as a giant dictionary
}

type RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        // The unique internal identifier of the portal theme
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  *RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage *RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BannerImage      *RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BackgroundImage  *RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerTitle      string                                                                                             `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                             `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                             `json:"footerElement,omitempty"`    //
}

type RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPageCustomizations struct {
	Data *[]RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPageCustomizationsData `json:"data,omitempty"` // The Dictionary will be exposed here as key value pair
}

type RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

type RequestByodPortalCreateByodPortal struct {
	ByodPortal *RequestByodPortalCreateByodPortalByodPortal `json:"BYODPortal,omitempty"` //
}

type RequestByodPortalCreateByodPortalByodPortal struct {
	ID             string                                                     `json:"id,omitempty"`             // Resource UUID, mandatory for update
	Name           string                                                     `json:"name,omitempty"`           // Resource Name
	Description    string                                                     `json:"description,omitempty"`    //
	PortalType     string                                                     `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                     `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       *RequestByodPortalCreateByodPortalByodPortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a BYOD
	Customizations *RequestByodPortalCreateByodPortalByodPortalCustomizations `json:"customizations,omitempty"` // Defines all of the Portal Customizations available for a BYOD
}

type RequestByodPortalCreateByodPortalByodPortalSettings struct {
	PortalSettings      *RequestByodPortalCreateByodPortalByodPortalSettingsPortalSettings      `json:"portalSettings,omitempty"`      // The port, interface, certificate, and other basic settings of a portal
	ByodSettings        *RequestByodPortalCreateByodPortalByodPortalSettingsByodSettings        `json:"byodSettings,omitempty"`        // Configuration of BYOD Device Welcome, Registration and Success steps
	SupportInfoSettings *RequestByodPortalCreateByodPortalByodPortalSettingsSupportInfoSettings `json:"supportInfoSettings,omitempty"` //
}

type RequestByodPortalCreateByodPortalByodPortalSettingsPortalSettings struct {
	HTTPSPort             *int     `json:"httpsPort,omitempty"`             // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces     []string `json:"allowedInterfaces,omitempty"`     // Interfaces that the portal will be reachable on. Allowed values: - eth0, - eth1, - eth2, - eth3, - eth4, - eth5, - bond0, - bond1, - bond2
	CertificateGroupTag   string   `json:"certificateGroupTag,omitempty"`   // Logical name of the x.509 server certificate that will be used for the portal
	EndpointIDentityGroup string   `json:"endpointIdentityGroup,omitempty"` // Unique Id of the endpoint identity group where user's devices will be added. Used only in Hotspot Portal
	DisplayLang           string   `json:"displayLang,omitempty"`           // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage      string   `json:"fallbackLanguage,omitempty"`      // Used when displayLang = USEBROWSERLOCALE
	AlwaysUsedLanguage    string   `json:"alwaysUsedLanguage,omitempty"`    // Used when displayLang = ALWAYSUSE
}

type RequestByodPortalCreateByodPortalByodPortalSettingsByodSettings struct {
	ByodWelcomeSettings             *RequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodWelcomeSettings             `json:"byodWelcomeSettings,omitempty"`             // Configuration of BYOD endpoint welcome step configuration
	ByodRegistrationSettings        *RequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodRegistrationSettings        `json:"byodRegistrationSettings,omitempty"`        //
	ByodRegistrationSuccessSettings *RequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodRegistrationSuccessSettings `json:"byodRegistrationSuccessSettings,omitempty"` //
}

type RequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodWelcomeSettings struct {
	EnableByod           *bool  `json:"enableBYOD,omitempty"`           //
	EnableGuestAccess    *bool  `json:"enableGuestAccess,omitempty"`    //
	RequireMdm           *bool  `json:"requireMDM,omitempty"`           //
	IncludeAup           *bool  `json:"includeAup,omitempty"`           //
	AupDisplay           string `json:"aupDisplay,omitempty"`           // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: - ONPAGE, - ASLINK
	RequireAupAcceptance *bool  `json:"requireAupAcceptance,omitempty"` //
	RequireScrolling     *bool  `json:"requireScrolling,omitempty"`     // Require BYOD devices to scroll down to the bottom of the AUP, Only valid if includeAup = true
}

type RequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodRegistrationSettings struct {
	ShowDeviceID            *bool  `json:"showDeviceID,omitempty"`            //
	EndPointIDentityGroupID string `json:"endPointIdentityGroupId,omitempty"` //
}

type RequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodRegistrationSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` // After an Authentication Success where should device be redirected. Allowed values:
	RedirectURL     string `json:"redirectUrl,omitempty"`     // Target URL for redirection, used when successRedirect = URL
}

type RequestByodPortalCreateByodPortalByodPortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  *bool  `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          *bool  `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        *bool  `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent *bool  `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     *bool  `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      *bool  `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type RequestByodPortalCreateByodPortalByodPortalCustomizations struct {
	PortalTheme          *RequestByodPortalCreateByodPortalByodPortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          // Defines the configuration for portal theme
	PortalTweakSettings  *RequestByodPortalCreateByodPortalByodPortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             *RequestByodPortalCreateByodPortalByodPortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations *RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` // Represent the portal Global customizations
	PageCustomizations   *RequestByodPortalCreateByodPortalByodPortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   // Represent the entire page customization as a giant dictionary
}

type RequestByodPortalCreateByodPortalByodPortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        // The unique internal identifier of the portal theme
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type RequestByodPortalCreateByodPortalByodPortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type RequestByodPortalCreateByodPortalByodPortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  *RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage *RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BannerImage      *RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BackgroundImage  *RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerTitle      string                                                                                         `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                         `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                         `json:"footerElement,omitempty"`    //
}

type RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestByodPortalCreateByodPortalByodPortalCustomizationsPageCustomizations struct {
	Data *[]RequestByodPortalCreateByodPortalByodPortalCustomizationsPageCustomizationsData `json:"data,omitempty"` // The Dictionary will be exposed here as key value pair
}

type RequestByodPortalCreateByodPortalByodPortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

//GetByodPortalByID Get BYOD portal by ID
/* This API allows the client to get a BYOD portal by ID.

@param id id path parameter. Portal id
*/
func (s *ByodPortalService) GetByodPortalByID(id string) (*ResponseByodPortalGetByodPortalByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/byodportal/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseByodPortalGetByodPortalByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetByodPortalById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseByodPortalGetByodPortalByID)
	return result, response, err

}

//GetByodPortal Get all BYOD portals
/* This API allows the client to get all the BYOD portals.

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

@param getByodPortalQueryParams Filtering parameter
*/
func (s *ByodPortalService) GetByodPortal(getByodPortalQueryParams *GetByodPortalQueryParams) (*ResponseByodPortalGetByodPortal, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/byodportal"

	queryString, _ := query.Values(getByodPortalQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseByodPortalGetByodPortal{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetByodPortal")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseByodPortalGetByodPortal)
	return result, response, err

}

//GetVersion Get BYOD portal version information
/* This API helps to retrieve the version information related to the BYOD portal.

 */
func (s *ByodPortalService) GetVersion() (*ResponseByodPortalGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/byodportal/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseByodPortalGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseByodPortalGetVersion)
	return result, response, err

}

//CreateByodPortal Create BYOD portal
/* This API creates a BYOD portal.

 */
func (s *ByodPortalService) CreateByodPortal(requestByodPortalCreateByodPortal *RequestByodPortalCreateByodPortal) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/byodportal"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestByodPortalCreateByodPortal).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateByodPortal")
	}

	return response, err

}

//UpdateByodPortalByID Update BYOD portal by ID
/* This API allows the client to update a BYOD portal by ID.

@param id id path parameter. Portal id
*/
func (s *ByodPortalService) UpdateByodPortalByID(id string, requestByodPortalUpdateByodPortalById *RequestByodPortalUpdateByodPortalByID) (*ResponseByodPortalUpdateByodPortalByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/byodportal/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestByodPortalUpdateByodPortalById).
		SetResult(&ResponseByodPortalUpdateByodPortalByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateByodPortalById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseByodPortalUpdateByodPortalByID)
	return result, response, err

}

//DeleteByodPortalByID Delete BYOD portal by ID
/* This API deletes a BYOD portal by ID.

@param id id path parameter. Portal id
*/
func (s *ByodPortalService) DeleteByodPortalByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/byodportal/{id}"
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
		return response, fmt.Errorf("error with operation DeleteByodPortalById")
	}

	getCSFRToken(response.Header())

	return response, err

}
