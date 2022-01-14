package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SponsoredGuestPortalService service

type GetSponsoredGuestPortalsQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByID struct {
	SponsoredGuestPortal *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortal `json:"SponsoredGuestPortal,omitempty"` //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortal struct {
	ID             string                                                                                     `json:"id,omitempty"`             //
	Name           string                                                                                     `json:"name,omitempty"`           //
	Description    string                                                                                     `json:"description,omitempty"`    //
	PortalType     string                                                                                     `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                                                     `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a portal
	Customizations *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizations `json:"customizations,omitempty"` // Defines all of the Portal Customizations available
	Link           *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalLink           `json:"link,omitempty"`           //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettings struct {
	PortalSettings                  *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPortalSettings                  `json:"portalSettings,omitempty"`                  // The port, interface, certificate, and other basic settings of a portal
	LoginPageSettings               *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettings               `json:"loginPageSettings,omitempty"`               // Portal Login Page settings groups follow
	AupSettings                     *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAupSettings                     `json:"aupSettings,omitempty"`                     //
	GuestChangePasswordSettings     *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestChangePasswordSettings     `json:"guestChangePasswordSettings,omitempty"`     //
	GuestDeviceRegistrationSettings *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestDeviceRegistrationSettings `json:"guestDeviceRegistrationSettings,omitempty"` //
	ByodSettings                    *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettings                    `json:"byodSettings,omitempty"`                    //
	PostAccessBannerSettings        *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostAccessBannerSettings        `json:"postAccessBannerSettings,omitempty"`        //
	AuthSuccessSettings             *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAuthSuccessSettings             `json:"authSuccessSettings,omitempty"`             //
	PostLoginBannerSettings         *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostLoginBannerSettings         `json:"postLoginBannerSettings,omitempty"`         //
	SupportInfoSettings             *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsSupportInfoSettings             `json:"supportInfoSettings,omitempty"`             //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPortalSettings struct {
	HTTPSPort                    *int     `json:"httpsPort,omitempty"`                    // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces            []string `json:"allowedInterfaces,omitempty"`            // Interfaces that the portal will be reachable on. Allowed values: - eth0, - eth1, - eth2, - eth3, - eth4, - eth5, - bond0, - bond1, - bond2
	CertificateGroupTag          string   `json:"certificateGroupTag,omitempty"`          // Logical name of the x.509 server certificate that will be used for the portal
	AuthenticationMethod         string   `json:"authenticationMethod,omitempty"`         // Unique Id of the identity source sequence
	AssignedGuestTypeForEmployee string   `json:"assignedGuestTypeForEmployee,omitempty"` // Unique Id of a guest type. Employees using this portal as a guest inherit login options from the guest type
	DisplayLang                  string   `json:"displayLang,omitempty"`                  // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage             string   `json:"fallbackLanguage,omitempty"`             // Used when displayLang = USEBROWSERLOCALE
	AlwaysUsedLanguage           string   `json:"alwaysUsedLanguage,omitempty"`           //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettings struct {
	RequireAccessCode                *bool                                                                                                                `json:"requireAccessCode,omitempty"`                //
	MaxFailedAttemptsBeforeRateLimit *int                                                                                                                 `json:"maxFailedAttemptsBeforeRateLimit,omitempty"` // Maximum failed login attempts before rate limiting
	TimeBetweenLoginsDuringRateLimit *int                                                                                                                 `json:"timeBetweenLoginsDuringRateLimit,omitempty"` // Time between login attempts when rate limiting
	IncludeAup                       *bool                                                                                                                `json:"includeAup,omitempty"`                       // Include an Acceptable Use Policy (AUP) that should be displayed during login
	AupDisplay                       string                                                                                                               `json:"aupDisplay,omitempty"`                       // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: -  ONPAGE, - ASLINK
	RequireAupAcceptance             *bool                                                                                                                `json:"requireAupAcceptance,omitempty"`             // Require the portal user to accept the AUP. Only valid if includeAup = true
	AccessCode                       string                                                                                                               `json:"accessCode,omitempty"`                       // Access code that must be entered by the portal user (only valid if requireAccessCode = true)
	AllowGuestToCreateAccounts       *bool                                                                                                                `json:"allowGuestToCreateAccounts,omitempty"`       //
	AllowForgotPassword              *bool                                                                                                                `json:"allowForgotPassword,omitempty"`              //
	AllowGuestToChangePassword       *bool                                                                                                                `json:"allowGuestToChangePassword,omitempty"`       // Require the portal user to enter an access code
	AllowAlternateGuestPortal        *bool                                                                                                                `json:"allowAlternateGuestPortal,omitempty"`        //
	SocialConfigs                    *[]ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs `json:"socialConfigs,omitempty"`                    //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs struct {
	SocialMediaType  string `json:"socialMediaType,omitempty"`  //
	SocialMediaValue string `json:"socialMediaValue,omitempty"` //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAupSettings struct {
	IncludeAup                   *bool  `json:"includeAup,omitempty"`                   //
	RequireAupScrolling          *bool  `json:"requireAupScrolling,omitempty"`          // Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true
	UseDiffAupForEmployees       *bool  `json:"useDiffAupForEmployees,omitempty"`       // Only valid if requireAupAcceptance = true
	SkipAupForEmployees          *bool  `json:"skipAupForEmployees,omitempty"`          // Only valid if requireAupAcceptance = true
	DisplayFrequencyIntervalDays *int   `json:"displayFrequencyIntervalDays,omitempty"` // Number of days between AUP confirmations (when displayFrequency = recurring)
	RequireScrolling             *bool  `json:"requireScrolling,omitempty"`             //
	DisplayFrequency             string `json:"displayFrequency,omitempty"`             // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed Values: - FIRSTLOGIN, - EVERYLOGIN, - RECURRING
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestChangePasswordSettings struct {
	AllowChangePasswdAtFirstLogin *bool `json:"allowChangePasswdAtFirstLogin,omitempty"` // Allow guest to change their own passwords
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestDeviceRegistrationSettings struct {
	AutoRegisterGuestDevices     *bool `json:"autoRegisterGuestDevices,omitempty"`     // Automatically register guest devices
	AllowGuestsToRegisterDevices *bool `json:"allowGuestsToRegisterDevices,omitempty"` // Allow guests to register devices
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettings struct {
	ByodWelcomeSettings             *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodWelcomeSettings             `json:"byodWelcomeSettings,omitempty"`             //
	ByodRegistrationSettings        *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSettings        `json:"byodRegistrationSettings,omitempty"`        // Configuration of BYOD endpoint Registration step configuration
	ByodRegistrationSuccessSettings *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSuccessSettings `json:"byodRegistrationSuccessSettings,omitempty"` //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodWelcomeSettings struct {
	EnableByod           *bool  `json:"enableBYOD,omitempty"`           //
	EnableGuestAccess    *bool  `json:"enableGuestAccess,omitempty"`    //
	RequireMdm           *bool  `json:"requireMDM,omitempty"`           //
	IncludeAup           *bool  `json:"includeAup,omitempty"`           //
	AupDisplay           string `json:"aupDisplay,omitempty"`           // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: - ONPAGE, - ASLINK
	RequireAupAcceptance *bool  `json:"requireAupAcceptance,omitempty"` //
	RequireScrolling     *bool  `json:"requireScrolling,omitempty"`     // Require BYOD devices to scroll down to the bottom of the AUP. Only valid if includeAup = true
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSettings struct {
	ShowDeviceID            *bool  `json:"showDeviceID,omitempty"`            // Display Device ID field during registration
	EndPointIDentityGroupID string `json:"endPointIdentityGroupId,omitempty"` // Identity group id for which endpoint belongs
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` // After an Authentication Success where should device be redirected. Allowed values: - AUTHSUCCESSPAGE, - ORIGINATINGURL, - URL
	RedirectURL     string `json:"redirectUrl,omitempty"`     // Target URL for redirection, used when successRedirect = URL
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostAccessBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAuthSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` // After an Authentication Success where should device be redirected. Allowed values: - AUTHSUCCESSPAGE, - ORIGINATINGURL, - URL
	RedirectURL     string `json:"redirectUrl,omitempty"`     // Target URL for redirection, used when successRedirect = URL
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostLoginBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` // Include a Post-Login Banner page
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  *bool  `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          *bool  `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        *bool  `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent *bool  `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     *bool  `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      *bool  `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizations struct {
	PortalTheme          *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          //
	PortalTweakSettings  *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` //
	PageCustomizations   *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   // Represent the entire page customization as a giant dictionary
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BannerImage      *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BackgroundImage  *ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerTitle      string                                                                                                                         `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                                                         `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                                                         `json:"footerElement,omitempty"`    //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizations struct {
	Data *[]ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizationsData `json:"data,omitempty"` // The Dictionary will be exposed here as key value pair
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSponsoredGuestPortalUpdateSponsoredGuestPortalByID struct {
	UpdatedFieldsList *ResponseSponsoredGuestPortalUpdateSponsoredGuestPortalByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseSponsoredGuestPortalUpdateSponsoredGuestPortalByIDUpdatedFieldsList struct {
	UpdatedField *[]ResponseSponsoredGuestPortalUpdateSponsoredGuestPortalByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                                     `json:"field,omitempty"`        //
	OldValue     string                                                                                     `json:"oldValue,omitempty"`     //
	NewValue     string                                                                                     `json:"newValue,omitempty"`     //
}

type ResponseSponsoredGuestPortalUpdateSponsoredGuestPortalByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortals struct {
	SearchResult *ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResult struct {
	Total        *int                                                                          `json:"total,omitempty"`        //
	Resources    *[]ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResultResources struct {
	ID          string                                                                         `json:"id,omitempty"`          //
	Name        string                                                                         `json:"name,omitempty"`        //
	Description string                                                                         `json:"description,omitempty"` //
	Link        *ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSponsoredGuestPortalGetVersion struct {
	VersionInfo *ResponseSponsoredGuestPortalGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseSponsoredGuestPortalGetVersionVersionInfo struct {
	CurrentServerVersion string                                                 `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                                 `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseSponsoredGuestPortalGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseSponsoredGuestPortalGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByID struct {
	SponsoredGuestPortal *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortal `json:"SponsoredGuestPortal,omitempty"` //
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortal struct {
	ID             string                                                                                       `json:"id,omitempty"`             //
	Name           string                                                                                       `json:"name,omitempty"`           //
	Description    string                                                                                       `json:"description,omitempty"`    //
	PortalType     string                                                                                       `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                                                       `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a portal
	Customizations *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizations `json:"customizations,omitempty"` // Defines all of the Portal Customizations available
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettings struct {
	PortalSettings                  *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPortalSettings                  `json:"portalSettings,omitempty"`                  // The port, interface, certificate, and other basic settings of a portal
	LoginPageSettings               *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettings               `json:"loginPageSettings,omitempty"`               // Portal Login Page settings groups follow
	AupSettings                     *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAupSettings                     `json:"aupSettings,omitempty"`                     //
	GuestChangePasswordSettings     *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestChangePasswordSettings     `json:"guestChangePasswordSettings,omitempty"`     //
	GuestDeviceRegistrationSettings *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestDeviceRegistrationSettings `json:"guestDeviceRegistrationSettings,omitempty"` //
	ByodSettings                    *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettings                    `json:"byodSettings,omitempty"`                    //
	PostAccessBannerSettings        *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostAccessBannerSettings        `json:"postAccessBannerSettings,omitempty"`        //
	AuthSuccessSettings             *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAuthSuccessSettings             `json:"authSuccessSettings,omitempty"`             //
	PostLoginBannerSettings         *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostLoginBannerSettings         `json:"postLoginBannerSettings,omitempty"`         //
	SupportInfoSettings             *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsSupportInfoSettings             `json:"supportInfoSettings,omitempty"`             //
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPortalSettings struct {
	HTTPSPort                    *int     `json:"httpsPort,omitempty"`                    // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces            []string `json:"allowedInterfaces,omitempty"`            // Interfaces that the portal will be reachable on. Allowed values: - eth0, - eth1, - eth2, - eth3, - eth4, - eth5, - bond0, - bond1, - bond2
	CertificateGroupTag          string   `json:"certificateGroupTag,omitempty"`          // Logical name of the x.509 server certificate that will be used for the portal
	AuthenticationMethod         string   `json:"authenticationMethod,omitempty"`         // Unique Id of the identity source sequence
	AssignedGuestTypeForEmployee string   `json:"assignedGuestTypeForEmployee,omitempty"` // Unique Id of a guest type. Employees using this portal as a guest inherit login options from the guest type
	DisplayLang                  string   `json:"displayLang,omitempty"`                  // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage             string   `json:"fallbackLanguage,omitempty"`             // Used when displayLang = USEBROWSERLOCALE
	AlwaysUsedLanguage           string   `json:"alwaysUsedLanguage,omitempty"`           //
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettings struct {
	RequireAccessCode                *bool                                                                                                                  `json:"requireAccessCode,omitempty"`                //
	MaxFailedAttemptsBeforeRateLimit *int                                                                                                                   `json:"maxFailedAttemptsBeforeRateLimit,omitempty"` // Maximum failed login attempts before rate limiting
	TimeBetweenLoginsDuringRateLimit *int                                                                                                                   `json:"timeBetweenLoginsDuringRateLimit,omitempty"` // Time between login attempts when rate limiting
	IncludeAup                       *bool                                                                                                                  `json:"includeAup,omitempty"`                       // Include an Acceptable Use Policy (AUP) that should be displayed during login
	AupDisplay                       string                                                                                                                 `json:"aupDisplay,omitempty"`                       // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: -  ONPAGE, - ASLINK
	RequireAupAcceptance             *bool                                                                                                                  `json:"requireAupAcceptance,omitempty"`             // Require the portal user to accept the AUP. Only valid if includeAup = true
	AccessCode                       string                                                                                                                 `json:"accessCode,omitempty"`                       // Access code that must be entered by the portal user (only valid if requireAccessCode = true)
	AllowGuestToCreateAccounts       *bool                                                                                                                  `json:"allowGuestToCreateAccounts,omitempty"`       //
	AllowForgotPassword              *bool                                                                                                                  `json:"allowForgotPassword,omitempty"`              //
	AllowGuestToChangePassword       *bool                                                                                                                  `json:"allowGuestToChangePassword,omitempty"`       // Require the portal user to enter an access code
	AllowAlternateGuestPortal        *bool                                                                                                                  `json:"allowAlternateGuestPortal,omitempty"`        //
	SocialConfigs                    *[]RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs `json:"socialConfigs,omitempty"`                    //
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs struct {
	SocialMediaType  string `json:"socialMediaType,omitempty"`  //
	SocialMediaValue string `json:"socialMediaValue,omitempty"` //
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAupSettings struct {
	IncludeAup                   *bool  `json:"includeAup,omitempty"`                   //
	RequireAupScrolling          *bool  `json:"requireAupScrolling,omitempty"`          // Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true
	UseDiffAupForEmployees       *bool  `json:"useDiffAupForEmployees,omitempty"`       // Only valid if requireAupAcceptance = true
	SkipAupForEmployees          *bool  `json:"skipAupForEmployees,omitempty"`          // Only valid if requireAupAcceptance = true
	DisplayFrequencyIntervalDays *int   `json:"displayFrequencyIntervalDays,omitempty"` // Number of days between AUP confirmations (when displayFrequency = recurring)
	RequireScrolling             *bool  `json:"requireScrolling,omitempty"`             //
	DisplayFrequency             string `json:"displayFrequency,omitempty"`             // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed Values: - FIRSTLOGIN, - EVERYLOGIN, - RECURRING
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestChangePasswordSettings struct {
	AllowChangePasswdAtFirstLogin *bool `json:"allowChangePasswdAtFirstLogin,omitempty"` // Allow guest to change their own passwords
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestDeviceRegistrationSettings struct {
	AutoRegisterGuestDevices     *bool `json:"autoRegisterGuestDevices,omitempty"`     // Automatically register guest devices
	AllowGuestsToRegisterDevices *bool `json:"allowGuestsToRegisterDevices,omitempty"` // Allow guests to register devices
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettings struct {
	ByodWelcomeSettings             *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodWelcomeSettings             `json:"byodWelcomeSettings,omitempty"`             //
	ByodRegistrationSettings        *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSettings        `json:"byodRegistrationSettings,omitempty"`        // Configuration of BYOD endpoint Registration step configuration
	ByodRegistrationSuccessSettings *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSuccessSettings `json:"byodRegistrationSuccessSettings,omitempty"` //
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodWelcomeSettings struct {
	EnableByod           *bool  `json:"enableBYOD,omitempty"`           //
	EnableGuestAccess    *bool  `json:"enableGuestAccess,omitempty"`    //
	RequireMdm           *bool  `json:"requireMDM,omitempty"`           //
	IncludeAup           *bool  `json:"includeAup,omitempty"`           //
	AupDisplay           string `json:"aupDisplay,omitempty"`           // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: - ONPAGE, - ASLINK
	RequireAupAcceptance *bool  `json:"requireAupAcceptance,omitempty"` //
	RequireScrolling     *bool  `json:"requireScrolling,omitempty"`     // Require BYOD devices to scroll down to the bottom of the AUP. Only valid if includeAup = true
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSettings struct {
	ShowDeviceID            *bool  `json:"showDeviceID,omitempty"`            // Display Device ID field during registration
	EndPointIDentityGroupID string `json:"endPointIdentityGroupId,omitempty"` // Identity group id for which endpoint belongs
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` // After an Authentication Success where should device be redirected. Allowed values: - AUTHSUCCESSPAGE, - ORIGINATINGURL, - URL
	RedirectURL     string `json:"redirectUrl,omitempty"`     // Target URL for redirection, used when successRedirect = URL
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostAccessBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` //
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAuthSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` // After an Authentication Success where should device be redirected. Allowed values: - AUTHSUCCESSPAGE, - ORIGINATINGURL, - URL
	RedirectURL     string `json:"redirectUrl,omitempty"`     // Target URL for redirection, used when successRedirect = URL
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostLoginBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` // Include a Post-Login Banner page
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  *bool  `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          *bool  `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        *bool  `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent *bool  `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     *bool  `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      *bool  `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizations struct {
	PortalTheme          *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          //
	PortalTweakSettings  *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` //
	PageCustomizations   *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   // Represent the entire page customization as a giant dictionary
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BannerImage      *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BackgroundImage  *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerTitle      string                                                                                                                           `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                                                           `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                                                           `json:"footerElement,omitempty"`    //
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizations struct {
	Data *[]RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizationsData `json:"data,omitempty"` // The Dictionary will be exposed here as key value pair
}

type RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortal struct {
	SponsoredGuestPortal *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortal `json:"SponsoredGuestPortal,omitempty"` //
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortal struct {
	Name           string                                                                                   `json:"name,omitempty"`           //
	Description    string                                                                                   `json:"description,omitempty"`    //
	PortalType     string                                                                                   `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                                                   `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a portal
	Customizations *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizations `json:"customizations,omitempty"` // Defines all of the Portal Customizations available
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettings struct {
	PortalSettings                  *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPortalSettings                  `json:"portalSettings,omitempty"`                  // The port, interface, certificate, and other basic settings of a portal
	LoginPageSettings               *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsLoginPageSettings               `json:"loginPageSettings,omitempty"`               // Portal Login Page settings groups follow
	AupSettings                     *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsAupSettings                     `json:"aupSettings,omitempty"`                     //
	GuestChangePasswordSettings     *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsGuestChangePasswordSettings     `json:"guestChangePasswordSettings,omitempty"`     //
	GuestDeviceRegistrationSettings *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsGuestDeviceRegistrationSettings `json:"guestDeviceRegistrationSettings,omitempty"` //
	ByodSettings                    *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettings                    `json:"byodSettings,omitempty"`                    //
	PostAccessBannerSettings        *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPostAccessBannerSettings        `json:"postAccessBannerSettings,omitempty"`        //
	AuthSuccessSettings             *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsAuthSuccessSettings             `json:"authSuccessSettings,omitempty"`             //
	PostLoginBannerSettings         *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPostLoginBannerSettings         `json:"postLoginBannerSettings,omitempty"`         //
	SupportInfoSettings             *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsSupportInfoSettings             `json:"supportInfoSettings,omitempty"`             //
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPortalSettings struct {
	HTTPSPort                    *int     `json:"httpsPort,omitempty"`                    // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces            []string `json:"allowedInterfaces,omitempty"`            // Interfaces that the portal will be reachable on. Allowed values: - eth0, - eth1, - eth2, - eth3, - eth4, - eth5, - bond0, - bond1, - bond2
	CertificateGroupTag          string   `json:"certificateGroupTag,omitempty"`          // Logical name of the x.509 server certificate that will be used for the portal
	AuthenticationMethod         string   `json:"authenticationMethod,omitempty"`         // Unique Id of the identity source sequence
	AssignedGuestTypeForEmployee string   `json:"assignedGuestTypeForEmployee,omitempty"` // Unique Id of a guest type. Employees using this portal as a guest inherit login options from the guest type
	DisplayLang                  string   `json:"displayLang,omitempty"`                  // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage             string   `json:"fallbackLanguage,omitempty"`             // Used when displayLang = USEBROWSERLOCALE
	AlwaysUsedLanguage           string   `json:"alwaysUsedLanguage,omitempty"`           //
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsLoginPageSettings struct {
	RequireAccessCode                *bool                                                                                                              `json:"requireAccessCode,omitempty"`                //
	MaxFailedAttemptsBeforeRateLimit *int                                                                                                               `json:"maxFailedAttemptsBeforeRateLimit,omitempty"` // Maximum failed login attempts before rate limiting
	TimeBetweenLoginsDuringRateLimit *int                                                                                                               `json:"timeBetweenLoginsDuringRateLimit,omitempty"` // Time between login attempts when rate limiting
	IncludeAup                       *bool                                                                                                              `json:"includeAup,omitempty"`                       // Include an Acceptable Use Policy (AUP) that should be displayed during login
	AupDisplay                       string                                                                                                             `json:"aupDisplay,omitempty"`                       // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: -  ONPAGE, - ASLINK
	RequireAupAcceptance             *bool                                                                                                              `json:"requireAupAcceptance,omitempty"`             // Require the portal user to accept the AUP. Only valid if includeAup = true
	AccessCode                       string                                                                                                             `json:"accessCode,omitempty"`                       // Access code that must be entered by the portal user (only valid if requireAccessCode = true)
	AllowGuestToCreateAccounts       *bool                                                                                                              `json:"allowGuestToCreateAccounts,omitempty"`       //
	AllowForgotPassword              *bool                                                                                                              `json:"allowForgotPassword,omitempty"`              //
	AllowGuestToChangePassword       *bool                                                                                                              `json:"allowGuestToChangePassword,omitempty"`       // Require the portal user to enter an access code
	AllowAlternateGuestPortal        *bool                                                                                                              `json:"allowAlternateGuestPortal,omitempty"`        //
	SocialConfigs                    *[]RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs `json:"socialConfigs,omitempty"`                    //
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs struct {
	SocialMediaType  string `json:"socialMediaType,omitempty"`  //
	SocialMediaValue string `json:"socialMediaValue,omitempty"` //
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsAupSettings struct {
	IncludeAup                   *bool  `json:"includeAup,omitempty"`                   //
	RequireAupScrolling          *bool  `json:"requireAupScrolling,omitempty"`          // Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true
	UseDiffAupForEmployees       *bool  `json:"useDiffAupForEmployees,omitempty"`       // Only valid if requireAupAcceptance = true
	SkipAupForEmployees          *bool  `json:"skipAupForEmployees,omitempty"`          // Only valid if requireAupAcceptance = true
	DisplayFrequencyIntervalDays *int   `json:"displayFrequencyIntervalDays,omitempty"` // Number of days between AUP confirmations (when displayFrequency = recurring)
	RequireScrolling             *bool  `json:"requireScrolling,omitempty"`             //
	DisplayFrequency             string `json:"displayFrequency,omitempty"`             // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed Values: - FIRSTLOGIN, - EVERYLOGIN, - RECURRING
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsGuestChangePasswordSettings struct {
	AllowChangePasswdAtFirstLogin *bool `json:"allowChangePasswdAtFirstLogin,omitempty"` // Allow guest to change their own passwords
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsGuestDeviceRegistrationSettings struct {
	AutoRegisterGuestDevices     *bool `json:"autoRegisterGuestDevices,omitempty"`     // Automatically register guest devices
	AllowGuestsToRegisterDevices *bool `json:"allowGuestsToRegisterDevices,omitempty"` // Allow guests to register devices
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettings struct {
	ByodWelcomeSettings             *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodWelcomeSettings             `json:"byodWelcomeSettings,omitempty"`             //
	ByodRegistrationSettings        *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodRegistrationSettings        `json:"byodRegistrationSettings,omitempty"`        // Configuration of BYOD endpoint Registration step configuration
	ByodRegistrationSuccessSettings *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodRegistrationSuccessSettings `json:"byodRegistrationSuccessSettings,omitempty"` //
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodWelcomeSettings struct {
	EnableByod           *bool  `json:"enableBYOD,omitempty"`           //
	EnableGuestAccess    *bool  `json:"enableGuestAccess,omitempty"`    //
	RequireMdm           *bool  `json:"requireMDM,omitempty"`           //
	IncludeAup           *bool  `json:"includeAup,omitempty"`           //
	AupDisplay           string `json:"aupDisplay,omitempty"`           // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: - ONPAGE, - ASLINK
	RequireAupAcceptance *bool  `json:"requireAupAcceptance,omitempty"` //
	RequireScrolling     *bool  `json:"requireScrolling,omitempty"`     // Require BYOD devices to scroll down to the bottom of the AUP. Only valid if includeAup = true
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodRegistrationSettings struct {
	ShowDeviceID            *bool  `json:"showDeviceID,omitempty"`            // Display Device ID field during registration
	EndPointIDentityGroupID string `json:"endPointIdentityGroupId,omitempty"` // Identity group id for which endpoint belongs
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodRegistrationSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` // After an Authentication Success where should device be redirected. Allowed values: - AUTHSUCCESSPAGE, - ORIGINATINGURL, - URL
	RedirectURL     string `json:"redirectUrl,omitempty"`     // Target URL for redirection, used when successRedirect = URL
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPostAccessBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` //
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsAuthSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` // After an Authentication Success where should device be redirected. Allowed values: - AUTHSUCCESSPAGE, - ORIGINATINGURL, - URL
	RedirectURL     string `json:"redirectUrl,omitempty"`     // Target URL for redirection, used when successRedirect = URL
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPostLoginBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` // Include a Post-Login Banner page
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  *bool  `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          *bool  `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        *bool  `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent *bool  `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     *bool  `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      *bool  `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizations struct {
	PortalTheme          *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          //
	PortalTweakSettings  *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` //
	PageCustomizations   *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   // Represent the entire page customization as a giant dictionary
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BannerImage      *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BackgroundImage  *RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerTitle      string                                                                                                                       `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                                                       `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                                                       `json:"footerElement,omitempty"`    //
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPageCustomizations struct {
	Data *[]RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPageCustomizationsData `json:"data,omitempty"` // The Dictionary will be exposed here as key value pair
}

type RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

//GetSponsoredGuestPortalByID Get sponsored guest portal by ID
/* This API allows the client to get a sponsored guest portal by ID.

@param id id path parameter.
*/
func (s *SponsoredGuestPortalService) GetSponsoredGuestPortalByID(id string) (*ResponseSponsoredGuestPortalGetSponsoredGuestPortalByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsoredguestportal/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSponsoredGuestPortalGetSponsoredGuestPortalByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSponsoredGuestPortalById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSponsoredGuestPortalGetSponsoredGuestPortalByID)
	return result, response, err

}

//GetSponsoredGuestPortals Get all sponsored guest portals
/* This API allows the client to get all the sponsored guest portals.

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

@param getSponsoredGuestPortalsQueryParams Filtering parameter
*/
func (s *SponsoredGuestPortalService) GetSponsoredGuestPortals(getSponsoredGuestPortalsQueryParams *GetSponsoredGuestPortalsQueryParams) (*ResponseSponsoredGuestPortalGetSponsoredGuestPortals, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsoredguestportal"

	queryString, _ := query.Values(getSponsoredGuestPortalsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSponsoredGuestPortalGetSponsoredGuestPortals{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSponsoredGuestPortals")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSponsoredGuestPortalGetSponsoredGuestPortals)
	return result, response, err

}

//GetVersion Get sponsored guest portal version information
/* This API helps to retrieve the version information related to the sponsored guest portal.

 */
func (s *SponsoredGuestPortalService) GetVersion() (*ResponseSponsoredGuestPortalGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsoredguestportal/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSponsoredGuestPortalGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSponsoredGuestPortalGetVersion)
	return result, response, err

}

//CreateSponsoredGuestPortal Create sponsored guest portal
/* This API creates a sponsored guest portal.

 */
func (s *SponsoredGuestPortalService) CreateSponsoredGuestPortal(requestSponsoredGuestPortalCreateSponsoredGuestPortal *RequestSponsoredGuestPortalCreateSponsoredGuestPortal) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsoredguestportal"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSponsoredGuestPortalCreateSponsoredGuestPortal).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateSponsoredGuestPortal")
	}

	return response, err

}

//UpdateSponsoredGuestPortalByID Update sponsored guest portal by ID
/* This API allows the client to update a sponsored guest portal by ID.

@param id id path parameter.
*/
func (s *SponsoredGuestPortalService) UpdateSponsoredGuestPortalByID(id string, requestSponsoredGuestPortalUpdateSponsoredGuestPortalById *RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByID) (*ResponseSponsoredGuestPortalUpdateSponsoredGuestPortalByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsoredguestportal/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSponsoredGuestPortalUpdateSponsoredGuestPortalById).
		SetResult(&ResponseSponsoredGuestPortalUpdateSponsoredGuestPortalByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseSponsoredGuestPortalUpdateSponsoredGuestPortalByID{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSponsoredGuestPortalById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSponsoredGuestPortalUpdateSponsoredGuestPortalByID)
	return result, response, err

}

//DeleteSponsoredGuestPortalByID Delete sponsored guest portal by ID
/* This API deletes a sponsored guest portal by ID.

@param id id path parameter.
*/
func (s *SponsoredGuestPortalService) DeleteSponsoredGuestPortalByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/sponsoredguestportal/{id}"
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
		return response, fmt.Errorf("error with operation DeleteSponsoredGuestPortalById")
	}

	getCSFRToken(response.Header())

	return response, err

}
