package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SelfRegisteredPortalService service

type GetSelfRegisteredPortalsQueryParams struct {
	Page       int      `url: page,omitempty`       //Page number
	Size       int      `url: size,omitempty`       //Number of objects returned per page
	Sortasc    string   `url: sortasc,omitempty`    //sort asc
	Sortdsc    string   `url: sortdsc,omitempty`    //sort desc
	Filter     []string `url: filter,omitempty`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url: filterType,omitempty` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByID struct {
	SelfRegPortal ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortal `json:"SelfRegPortal,omitempty"` //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortal struct {
	ID             string                                                                             `json:"id,omitempty"`             //
	Name           string                                                                             `json:"name,omitempty"`           //
	Description    string                                                                             `json:"description,omitempty"`    //
	PortalType     string                                                                             `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                                             `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a portal
	Customizations ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizations `json:"customizations,omitempty"` // Defines all of the Portal Customizations available
	Link           ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalLink           `json:"link,omitempty"`           //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettings struct {
	PortalSettings                  ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsPortalSettings                  `json:"portalSettings,omitempty"`                  // The port, interface, certificate, and other basic settings of a portal
	LoginPageSettings               ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettings               `json:"loginPageSettings,omitempty"`               // Portal Login Page settings groups follow
	SelfRegPageSettings             ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettings             `json:"selfRegPageSettings,omitempty"`             //
	SelfRegSuccessSettings          ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegSuccessSettings          `json:"selfRegSuccessSettings,omitempty"`          //
	AupSettings                     ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsAupSettings                     `json:"aupSettings,omitempty"`                     //
	GuestChangePasswordSettings     ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsGuestChangePasswordSettings     `json:"guestChangePasswordSettings,omitempty"`     //
	GuestDeviceRegistrationSettings ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsGuestDeviceRegistrationSettings `json:"guestDeviceRegistrationSettings,omitempty"` //
	ByodSettings                    ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettings                    `json:"byodSettings,omitempty"`                    // Configuration of BYOD Device Welcome, Registration and Success steps
	PostLoginBannerSettings         ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsPostLoginBannerSettings         `json:"postLoginBannerSettings,omitempty"`         //
	PostAccessBannerSettings        ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsPostAccessBannerSettings        `json:"postAccessBannerSettings,omitempty"`        //
	AuthSuccessSettings             ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsAuthSuccessSettings             `json:"authSuccessSettings,omitempty"`             //
	SupportInfoSettings             ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSupportInfoSettings             `json:"supportInfoSettings,omitempty"`             //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsPortalSettings struct {
	HTTPSPort                    int    `json:"httpsPort,omitempty"`                    // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces            string `json:"allowedInterfaces,omitempty"`            // Interfaces that the portal will be reachable on. Allowed values: - eth0, - eth1, - eth2, - eth3, - eth4, - eth5, - bond0, - bond1, - bond2
	CertificateGroupTag          string `json:"certificateGroupTag,omitempty"`          // Logical name of the x.509 server certificate that will be used for the portal
	AuthenticationMethod         string `json:"authenticationMethod,omitempty"`         // Unique Id of the identity source sequence
	AssignedGuestTypeForEmployee string `json:"assignedGuestTypeForEmployee,omitempty"` // Unique Id of a guest type. Employees using this portal as a guest inherit login options from the guest type
	DisplayLang                  string `json:"displayLang,omitempty"`                  // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage             string `json:"fallbackLanguage,omitempty"`             // Used when displayLang = USEBROWSERLOCALE
	AlwaysUsedLanguage           string `json:"alwaysUsedLanguage,omitempty"`           //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettings struct {
	RequireAccessCode                bool                                                                                                         `json:"requireAccessCode,omitempty"`                // Require the portal user to enter an access code
	MaxFailedAttemptsBeforeRateLimit int                                                                                                          `json:"maxFailedAttemptsBeforeRateLimit,omitempty"` // Maximum failed login attempts before rate limiting
	TimeBetweenLoginsDuringRateLimit int                                                                                                          `json:"timeBetweenLoginsDuringRateLimit,omitempty"` // Time between login attempts when rate limiting
	IncludeAup                       bool                                                                                                         `json:"includeAup,omitempty"`                       // Include an Acceptable Use Policy (AUP) that should be displayed during login
	AupDisplay                       string                                                                                                       `json:"aupDisplay,omitempty"`                       // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: -  ONPAGE, - ASLINK
	RequireAupAcceptance             bool                                                                                                         `json:"requireAupAcceptance,omitempty"`             // Require the portal user to accept the AUP. Only valid if includeAup = true
	AccessCode                       string                                                                                                       `json:"accessCode,omitempty"`                       // Access code that must be entered by the portal user (only valid if requireAccessCode = true)
	AllowGuestToCreateAccounts       bool                                                                                                         `json:"allowGuestToCreateAccounts,omitempty"`       //
	AllowForgotPassword              bool                                                                                                         `json:"allowForgotPassword,omitempty"`              //
	AllowGuestToChangePassword       bool                                                                                                         `json:"allowGuestToChangePassword,omitempty"`       // Require the portal user to enter an access code
	AllowAlternateGuestPortal        bool                                                                                                         `json:"allowAlternateGuestPortal,omitempty"`        //
	AlternateGuestPortal             string                                                                                                       `json:"alternateGuestPortal,omitempty"`             //
	AllowGuestToUseSocialAccounts    bool                                                                                                         `json:"allowGuestToUseSocialAccounts,omitempty"`    //
	AllowShowGuestForm               bool                                                                                                         `json:"allowShowGuestForm,omitempty"`               //
	SocialConfigs                    []ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettingsSocialConfigs `json:"socialConfigs,omitempty"`                    //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettingsSocialConfigs struct {
	SocialMediaType  string `json:"socialMediaType,omitempty"`  //
	SocialMediaValue string `json:"socialMediaValue,omitempty"` //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettings struct {
	AssignGuestsToGuestType             string                                                                                                                 `json:"assignGuestsToGuestType,omitempty"`             // Guests are assigned to this guest type
	AccountValidityDuration             int                                                                                                                    `json:"accountValidityDuration,omitempty"`             // Self-registered guest account is valid for this many account_validity_time_units
	AccountValidityTimeUnits            string                                                                                                                 `json:"accountValidityTimeUnits,omitempty"`            // Time units for account_validity_duration. Allowed Values: - DAYS, - HOURS, - MINUTES
	RequireRegistrationCode             bool                                                                                                                   `json:"requireRegistrationCode,omitempty"`             // Self-registered guests are required to enter a registration code
	RegistrationCode                    string                                                                                                                 `json:"registrationCode,omitempty"`                    // The registration code that the guest user must enter
	FieldUserName                       ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldUserName           `json:"fieldUserName,omitempty"`                       //
	FieldFirstName                      ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldFirstName          `json:"fieldFirstName,omitempty"`                      //
	FieldLastName                       ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLastName           `json:"fieldLastName,omitempty"`                       //
	FieldEmailAddr                      ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldEmailAddr          `json:"fieldEmailAddr,omitempty"`                      //
	FieldPhoneNo                        ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPhoneNo            `json:"fieldPhoneNo,omitempty"`                        //
	FieldCompany                        ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldCompany            `json:"fieldCompany,omitempty"`                        //
	FieldLocation                       ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLocation           `json:"fieldLocation,omitempty"`                       //
	SelectableLocations                 []string                                                                                                               `json:"selectableLocations,omitempty"`                 // Guests can choose from these locations to set their time zone
	FieldSmsProvider                    ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldSmsProvider        `json:"fieldSmsProvider,omitempty"`                    //
	SelectableSmsProviders              []string                                                                                                               `json:"selectableSmsProviders,omitempty"`              // This attribute is an array of SMS provider names
	FieldPersonBeingVisited             ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPersonBeingVisited `json:"fieldPersonBeingVisited,omitempty"`             //
	FieldReasonForVisit                 ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldReasonForVisit     `json:"fieldReasonForVisit,omitempty"`                 //
	IncludeAup                          bool                                                                                                                   `json:"includeAup,omitempty"`                          // Include an Acceptable Use Policy (AUP) that should be displayed during login
	AupDisplay                          string                                                                                                                 `json:"aupDisplay,omitempty"`                          // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: - ONPAGE, - ASLINK
	RequireAupAcceptance                bool                                                                                                                   `json:"requireAupAcceptance,omitempty"`                // Require the portal user to accept the AUP. Only valid if includeAup = true
	EnableGuestEmailWhitelist           bool                                                                                                                   `json:"enableGuestEmailWhitelist,omitempty"`           // Allow guests with an e-mail address from selected domains
	GuestEmailWhitelistDomains          string                                                                                                                 `json:"guestEmailWhitelistDomains,omitempty"`          // Self-registered guests whose e-mail address is in one of these domains will be allowed. Only valid if enableGuestEmailWhitelist = true
	EnableGuestEmailBlacklist           bool                                                                                                                   `json:"enableGuestEmailBlacklist,omitempty"`           // Disallow guests with an e-mail address from selected domains
	GuestEmailBlacklistDomains          string                                                                                                                 `json:"guestEmailBlacklistDomains,omitempty"`          // Disallow guests with an e-mail address from selected domains
	RequireGuestApproval                bool                                                                                                                   `json:"requireGuestApproval,omitempty"`                // Require self-registered guests to be approved if true
	AutoLoginSelfWait                   bool                                                                                                                   `json:"autoLoginSelfWait,omitempty"`                   // Allow guests to login automatically from self-registration after sponsor's approval. No need to provide the credentials by guest to login
	AutoLoginTimePeriod                 int                                                                                                                    `json:"autoLoginTimePeriod,omitempty"`                 // Waiting period for auto login until sponsor's approval. If time exceeds, guest has to login manually by providing the credentials. Default value is 5 minutes
	AllowGraceAccess                    bool                                                                                                                   `json:"allowGraceAccess,omitempty"`                    //
	GraceAccessExpireInterval           int                                                                                                                    `json:"graceAccessExpireInterval,omitempty"`           //
	GraceAccessSendAccountExpiration    bool                                                                                                                   `json:"graceAccessSendAccountExpiration,omitempty"`    //
	SendApprovalRequestTo               string                                                                                                                 `json:"sendApprovalRequestTo,omitempty"`               // Specifies where approval requests are sent. Only valid if requireGuestApproval = true. Allowed Values: - SELECTEDEMAILADDRESSES, - PERSONBEINGVISITED
	ApprovalEmailAddresses              string                                                                                                                 `json:"approvalEmailAddresses,omitempty"`              // Only valid if requireGuestApproval = true and sendApprovalRequestTo = SELECTEDEMAILADDRESSES
	PostRegistrationRedirect            string                                                                                                                 `json:"postRegistrationRedirect,omitempty"`            // After the registration submission direct the guest user to one of the following pages. Only valid if requireGuestApproval = true. Allowed Values: - SELFREGISTRATIONSUCCESS, - LOGINPAGEWITHINSTRUCTIONS - URL
	PostRegistrationRedirectURL         string                                                                                                                 `json:"postRegistrationRedirectUrl,omitempty"`         // URL where guest user is redirected after registration. Only valid if requireGuestApproval = true and postRegistrationRedirect = URL
	CredentialNotificationUsingEmail    bool                                                                                                                   `json:"credentialNotificationUsingEmail,omitempty"`    // If true, send credential notification upon approval using email. Only valid if requireGuestApproval = true
	CredentialNotificationUsingSms      bool                                                                                                                   `json:"credentialNotificationUsingSms,omitempty"`      // If true, send credential notification upon approval using SMS. Only valid if requireGuestApproval = true
	ApproveDenyLinksValidFor            int                                                                                                                    `json:"approveDenyLinksValidFor,omitempty"`            // This attribute, along with approveDenyLinksTimeUnits, specifies how long the link can be used. Only valid if requireGuestApproval = true
	ApproveDenyLinksTimeUnits           string                                                                                                                 `json:"approveDenyLinksTimeUnits,omitempty"`           // This attribute, along with approveDenyLinksValidFor, specifies how long the link can be used. Only valid if requireGuestApproval = true. Allowed Values: - DAYS, - HOURS, - MINUTES
	RequireApproverToAuthenticate       bool                                                                                                                   `json:"requireApproverToAuthenticate,omitempty"`       // When self-registered guests require approval, an approval request is e-mailed to one or more sponsor users. If the Cisco ISE Administrator chooses to include an approval link in the e-mail, a sponsor user who clicks the link will be required to enter their username and password if this attribute is true. Only valid if requireGuestApproval = true
	AuthenticateSponsorsUsingPortalList string                                                                                                                 `json:"authenticateSponsorsUsingPortalList,omitempty"` //
	SponsorPortalList                   []interface{}                                                                                                          `json:"sponsorPortalList,omitempty"`                   //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldUserName struct {
	Include bool `json:"include,omitempty"` //
	Require bool `json:"require,omitempty"` // Only applicable if include = true
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldFirstName struct {
	Include bool `json:"include,omitempty"` //
	Require bool `json:"require,omitempty"` // Only applicable if include = true
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLastName struct {
	Include bool `json:"include,omitempty"` //
	Require bool `json:"require,omitempty"` // Only applicable if include = true
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldEmailAddr struct {
	Include bool `json:"include,omitempty"` //
	Require bool `json:"require,omitempty"` // Only applicable if include = true
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPhoneNo struct {
	Include bool `json:"include,omitempty"` //
	Require bool `json:"require,omitempty"` // Only applicable if include = true
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldCompany struct {
	Include bool `json:"include,omitempty"` //
	Require bool `json:"require,omitempty"` // Only applicable if include = true
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLocation struct {
	Include bool `json:"include,omitempty"` //
	Require bool `json:"require,omitempty"` // Only applicable if include = true
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldSmsProvider struct {
	Include bool `json:"include,omitempty"` //
	Require bool `json:"require,omitempty"` // Only applicable if include = true
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPersonBeingVisited struct {
	Include bool `json:"include,omitempty"` //
	Require bool `json:"require,omitempty"` // Only applicable if include = true
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldReasonForVisit struct {
	Include bool `json:"include,omitempty"` //
	Require bool `json:"require,omitempty"` // Only applicable if include = true
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegSuccessSettings struct {
	IncludeUserName                       bool `json:"includeUserName,omitempty"`                       //
	IncludePassword                       bool `json:"includePassword,omitempty"`                       //
	IncludeFirstName                      bool `json:"includeFirstName,omitempty"`                      //
	IncludeLastName                       bool `json:"includeLastName,omitempty"`                       //
	IncludeEmailAddr                      bool `json:"includeEmailAddr,omitempty"`                      //
	IncludePhoneNo                        bool `json:"includePhoneNo,omitempty"`                        //
	IncludeCompany                        bool `json:"includeCompany,omitempty"`                        //
	IncludeLocation                       bool `json:"includeLocation,omitempty"`                       //
	IncludeSmsProvider                    bool `json:"includeSmsProvider,omitempty"`                    //
	IncludePersonBeingVisited             bool `json:"includePersonBeingVisited,omitempty"`             //
	IncludeReasonForVisit                 bool `json:"includeReasonForVisit,omitempty"`                 //
	AllowGuestSendSelfUsingPrint          bool `json:"allowGuestSendSelfUsingPrint,omitempty"`          //
	AllowGuestSendSelfUsingEmail          bool `json:"allowGuestSendSelfUsingEmail,omitempty"`          //
	AllowGuestSendSelfUsingSms            bool `json:"allowGuestSendSelfUsingSms,omitempty"`            //
	IncludeAup                            bool `json:"includeAup,omitempty"`                            //
	AupOnPage                             bool `json:"aupOnPage,omitempty"`                             //
	RequireAupAcceptance                  bool `json:"requireAupAcceptance,omitempty"`                  //
	RequireAupScrolling                   bool `json:"requireAupScrolling,omitempty"`                   //
	AllowGuestLoginFromSelfregSuccessPage bool `json:"allowGuestLoginFromSelfregSuccessPage,omitempty"` //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsAupSettings struct {
	IncludeAup                   bool   `json:"includeAup,omitempty"`                   // Require the portal user to read and accept an AUP
	UseDiffAupForEmployees       bool   `json:"useDiffAupForEmployees,omitempty"`       // Only valid if requireAupAcceptance = trueG
	SkipAupForEmployees          bool   `json:"skipAupForEmployees,omitempty"`          // Only valid if requireAupAcceptance = trueG
	RequireScrolling             bool   `json:"requireScrolling,omitempty"`             //
	RequireAupScrolling          bool   `json:"requireAupScrolling,omitempty"`          // Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true
	DisplayFrequency             string `json:"displayFrequency,omitempty"`             // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed Values: - FIRSTLOGIN, - EVERYLOGIN, - RECURRING
	DisplayFrequencyIntervalDays int    `json:"displayFrequencyIntervalDays,omitempty"` // Number of days between AUP confirmations (when displayFrequency = recurring)
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsGuestChangePasswordSettings struct {
	AllowChangePasswdAtFirstLogin bool `json:"allowChangePasswdAtFirstLogin,omitempty"` // Allow guest to change their own passwords
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsGuestDeviceRegistrationSettings struct {
	AutoRegisterGuestDevices     bool `json:"autoRegisterGuestDevices,omitempty"`     // Automatically register guest devices
	AllowGuestsToRegisterDevices bool `json:"allowGuestsToRegisterDevices,omitempty"` // Allow guests to register devices
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettings struct {
	ByodWelcomeSettings             ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodWelcomeSettings             `json:"byodWelcomeSettings,omitempty"`             // Configuration of BYOD endpoint welcome step configuration
	ByodRegistrationSettings        ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSettings        `json:"byodRegistrationSettings,omitempty"`        // Configuration of BYOD endpoint Registration step configuration
	ByodRegistrationSuccessSettings ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSuccessSettings `json:"byodRegistrationSuccessSettings,omitempty"` // Configuration of BYOD endpoint Registration Success step configuration
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodWelcomeSettings struct {
	EnableByod           bool   `json:"enableBYOD,omitempty"`           //
	EnableGuestAccess    bool   `json:"enableGuestAccess,omitempty"`    //
	RequireMdm           bool   `json:"requireMDM,omitempty"`           //
	IncludeAup           bool   `json:"includeAup,omitempty"`           //
	AupDisplay           string `json:"aupDisplay,omitempty"`           // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: - ONPAGE, - ASLINK
	RequireAupAcceptance bool   `json:"requireAupAcceptance,omitempty"` //
	RequireScrolling     bool   `json:"requireScrolling,omitempty"`     // Require BYOD devices to scroll down to the bottom of the AUP,  Only valid if includeAup = true
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSettings struct {
	ShowDeviceID            bool   `json:"showDeviceID,omitempty"`            // Display Device ID field during registration
	EndPointIDentityGroupID string `json:"endPointIdentityGroupId,omitempty"` // Identity group id for which endpoint belongs
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` // After an Authentication Success where should device be redirected. Allowed values: - AUTHSUCCESSPAGE, - ORIGINATINGURL, - URL
	RedirectURL     string `json:"redirectUrl,omitempty"`     // Target URL for redirection, used when successRedirect = URL
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsPostLoginBannerSettings struct {
	IncludePostAccessBanner bool `json:"includePostAccessBanner,omitempty"` // Include a Post-Login Banner page
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsPostAccessBannerSettings struct {
	IncludePostAccessBanner bool `json:"includePostAccessBanner,omitempty"` //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsAuthSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` //
	RedirectURL     string `json:"redirectUrl,omitempty"`     //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  bool   `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          bool   `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        bool   `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent bool   `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     bool   `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      bool   `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizations struct {
	PortalTheme          ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          //
	PortalTweakSettings  ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` //
	PageCustomizations   ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   // Represent the entire page customization as a giant dictionary
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BannerImage      ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BackgroundImage  ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerTitle      string                                                                                                                 `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                                                 `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                                                 `json:"footerElement,omitempty"`    //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizations struct {
	Data []ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizationsData `json:"data,omitempty"` // The Dictionary will be exposed here as key value pair
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSelfRegisteredPortalUpdateSelfRegisteredPortalByID struct {
	UpdatedFieldsList ResponseSelfRegisteredPortalUpdateSelfRegisteredPortalByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseSelfRegisteredPortalUpdateSelfRegisteredPortalByIDUpdatedFieldsList struct {
	UpdatedField ResponseSelfRegisteredPortalUpdateSelfRegisteredPortalByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                                  `json:"field,omitempty"`        //
	OldValue     string                                                                                  `json:"oldValue,omitempty"`     //
	NewValue     string                                                                                  `json:"newValue,omitempty"`     //
}

type ResponseSelfRegisteredPortalUpdateSelfRegisteredPortalByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortals struct {
	SearchResult ResponseSelfRegisteredPortalGetSelfRegisteredPortalsSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalsSearchResult struct {
	Total        int                                                                            `json:"total,omitempty"`        //
	Resources    []ResponseSelfRegisteredPortalGetSelfRegisteredPortalsSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseSelfRegisteredPortalGetSelfRegisteredPortalsSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseSelfRegisteredPortalGetSelfRegisteredPortalsSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalsSearchResultResources struct {
	ID          string                                                                        `json:"id,omitempty"`          //
	Name        string                                                                        `json:"name,omitempty"`        //
	Description string                                                                        `json:"description,omitempty"` //
	Link        ResponseSelfRegisteredPortalGetSelfRegisteredPortalsSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalsSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalsSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSelfRegisteredPortalGetSelfRegisteredPortalsSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSelfRegisteredPortalGetVersion struct {
	VersionInfo ResponseSelfRegisteredPortalGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseSelfRegisteredPortalGetVersionVersionInfo struct {
	CurrentServerVersion string                                                `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                                `json:"supportedVersions,omitempty"`    //
	Link                 ResponseSelfRegisteredPortalGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseSelfRegisteredPortalGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByID struct {
	SelfRegPortal *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortal `json:"SelfRegPortal,omitempty"` //
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortal struct {
	ID             string                                                                                `json:"id,omitempty"`             //
	Name           string                                                                                `json:"name,omitempty"`           //
	Description    string                                                                                `json:"description,omitempty"`    //
	PortalType     string                                                                                `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                                                `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a portal
	Customizations *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizations `json:"customizations,omitempty"` // Defines all of the Portal Customizations available
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettings struct {
	PortalSettings                  *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPortalSettings                  `json:"portalSettings,omitempty"`                  // The port, interface, certificate, and other basic settings of a portal
	LoginPageSettings               *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettings               `json:"loginPageSettings,omitempty"`               // Portal Login Page settings groups follow
	SelfRegPageSettings             *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettings             `json:"selfRegPageSettings,omitempty"`             //
	SelfRegSuccessSettings          *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegSuccessSettings          `json:"selfRegSuccessSettings,omitempty"`          //
	AupSettings                     *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsAupSettings                     `json:"aupSettings,omitempty"`                     //
	GuestChangePasswordSettings     *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsGuestChangePasswordSettings     `json:"guestChangePasswordSettings,omitempty"`     //
	GuestDeviceRegistrationSettings *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsGuestDeviceRegistrationSettings `json:"guestDeviceRegistrationSettings,omitempty"` //
	ByodSettings                    *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettings                    `json:"byodSettings,omitempty"`                    // Configuration of BYOD Device Welcome, Registration and Success steps
	PostLoginBannerSettings         *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPostLoginBannerSettings         `json:"postLoginBannerSettings,omitempty"`         //
	PostAccessBannerSettings        *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPostAccessBannerSettings        `json:"postAccessBannerSettings,omitempty"`        //
	AuthSuccessSettings             *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsAuthSuccessSettings             `json:"authSuccessSettings,omitempty"`             //
	SupportInfoSettings             *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSupportInfoSettings             `json:"supportInfoSettings,omitempty"`             //
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPortalSettings struct {
	HTTPSPort                    *int   `json:"httpsPort,omitempty"`                    // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces            string `json:"allowedInterfaces,omitempty"`            // Interfaces that the portal will be reachable on. Allowed values: - eth0, - eth1, - eth2, - eth3, - eth4, - eth5, - bond0, - bond1, - bond2
	CertificateGroupTag          string `json:"certificateGroupTag,omitempty"`          // Logical name of the x.509 server certificate that will be used for the portal
	AuthenticationMethod         string `json:"authenticationMethod,omitempty"`         // Unique Id of the identity source sequence
	AssignedGuestTypeForEmployee string `json:"assignedGuestTypeForEmployee,omitempty"` // Unique Id of a guest type. Employees using this portal as a guest inherit login options from the guest type
	DisplayLang                  string `json:"displayLang,omitempty"`                  // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage             string `json:"fallbackLanguage,omitempty"`             // Used when displayLang = USEBROWSERLOCALE
	AlwaysUsedLanguage           string `json:"alwaysUsedLanguage,omitempty"`           //
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettings struct {
	RequireAccessCode                *bool                                                                                                           `json:"requireAccessCode,omitempty"`                // Require the portal user to enter an access code
	MaxFailedAttemptsBeforeRateLimit *int                                                                                                            `json:"maxFailedAttemptsBeforeRateLimit,omitempty"` // Maximum failed login attempts before rate limiting
	TimeBetweenLoginsDuringRateLimit *int                                                                                                            `json:"timeBetweenLoginsDuringRateLimit,omitempty"` // Time between login attempts when rate limiting
	IncludeAup                       *bool                                                                                                           `json:"includeAup,omitempty"`                       // Include an Acceptable Use Policy (AUP) that should be displayed during login
	AupDisplay                       string                                                                                                          `json:"aupDisplay,omitempty"`                       // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: -  ONPAGE, - ASLINK
	RequireAupAcceptance             *bool                                                                                                           `json:"requireAupAcceptance,omitempty"`             // Require the portal user to accept the AUP. Only valid if includeAup = true
	AccessCode                       string                                                                                                          `json:"accessCode,omitempty"`                       // Access code that must be entered by the portal user (only valid if requireAccessCode = true)
	AllowGuestToCreateAccounts       *bool                                                                                                           `json:"allowGuestToCreateAccounts,omitempty"`       //
	AllowForgotPassword              *bool                                                                                                           `json:"allowForgotPassword,omitempty"`              //
	AllowGuestToChangePassword       *bool                                                                                                           `json:"allowGuestToChangePassword,omitempty"`       // Require the portal user to enter an access code
	AllowAlternateGuestPortal        *bool                                                                                                           `json:"allowAlternateGuestPortal,omitempty"`        //
	AlternateGuestPortal             string                                                                                                          `json:"alternateGuestPortal,omitempty"`             //
	AllowGuestToUseSocialAccounts    *bool                                                                                                           `json:"allowGuestToUseSocialAccounts,omitempty"`    //
	AllowShowGuestForm               *bool                                                                                                           `json:"allowShowGuestForm,omitempty"`               //
	SocialConfigs                    *[]RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettingsSocialConfigs `json:"socialConfigs,omitempty"`                    //
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettingsSocialConfigs struct {
	SocialMediaType  string `json:"socialMediaType,omitempty"`  //
	SocialMediaValue string `json:"socialMediaValue,omitempty"` //
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettings struct {
	AssignGuestsToGuestType             string                                                                                                                    `json:"assignGuestsToGuestType,omitempty"`             // Guests are assigned to this guest type
	AccountValidityDuration             *int                                                                                                                      `json:"accountValidityDuration,omitempty"`             // Self-registered guest account is valid for this many account_validity_time_units
	AccountValidityTimeUnits            string                                                                                                                    `json:"accountValidityTimeUnits,omitempty"`            // Time units for account_validity_duration. Allowed Values: - DAYS, - HOURS, - MINUTES
	RequireRegistrationCode             *bool                                                                                                                     `json:"requireRegistrationCode,omitempty"`             // Self-registered guests are required to enter a registration code
	RegistrationCode                    string                                                                                                                    `json:"registrationCode,omitempty"`                    // The registration code that the guest user must enter
	FieldUserName                       *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldUserName           `json:"fieldUserName,omitempty"`                       //
	FieldFirstName                      *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldFirstName          `json:"fieldFirstName,omitempty"`                      //
	FieldLastName                       *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLastName           `json:"fieldLastName,omitempty"`                       //
	FieldEmailAddr                      *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldEmailAddr          `json:"fieldEmailAddr,omitempty"`                      //
	FieldPhoneNo                        *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPhoneNo            `json:"fieldPhoneNo,omitempty"`                        //
	FieldCompany                        *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldCompany            `json:"fieldCompany,omitempty"`                        //
	FieldLocation                       *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLocation           `json:"fieldLocation,omitempty"`                       //
	SelectableLocations                 []string                                                                                                                  `json:"selectableLocations,omitempty"`                 // Guests can choose from these locations to set their time zone
	FieldSmsProvider                    *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldSmsProvider        `json:"fieldSmsProvider,omitempty"`                    //
	SelectableSmsProviders              []string                                                                                                                  `json:"selectableSmsProviders,omitempty"`              // This attribute is an array of SMS provider names
	FieldPersonBeingVisited             *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPersonBeingVisited `json:"fieldPersonBeingVisited,omitempty"`             //
	FieldReasonForVisit                 *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldReasonForVisit     `json:"fieldReasonForVisit,omitempty"`                 //
	IncludeAup                          *bool                                                                                                                     `json:"includeAup,omitempty"`                          // Include an Acceptable Use Policy (AUP) that should be displayed during login
	AupDisplay                          string                                                                                                                    `json:"aupDisplay,omitempty"`                          // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: - ONPAGE, - ASLINK
	RequireAupAcceptance                *bool                                                                                                                     `json:"requireAupAcceptance,omitempty"`                // Require the portal user to accept the AUP. Only valid if includeAup = true
	EnableGuestEmailWhitelist           *bool                                                                                                                     `json:"enableGuestEmailWhitelist,omitempty"`           // Allow guests with an e-mail address from selected domains
	GuestEmailWhitelistDomains          string                                                                                                                    `json:"guestEmailWhitelistDomains,omitempty"`          // Self-registered guests whose e-mail address is in one of these domains will be allowed. Only valid if enableGuestEmailWhitelist = true
	EnableGuestEmailBlacklist           *bool                                                                                                                     `json:"enableGuestEmailBlacklist,omitempty"`           // Disallow guests with an e-mail address from selected domains
	GuestEmailBlacklistDomains          string                                                                                                                    `json:"guestEmailBlacklistDomains,omitempty"`          // Disallow guests with an e-mail address from selected domains
	RequireGuestApproval                *bool                                                                                                                     `json:"requireGuestApproval,omitempty"`                // Require self-registered guests to be approved if true
	AutoLoginSelfWait                   *bool                                                                                                                     `json:"autoLoginSelfWait,omitempty"`                   // Allow guests to login automatically from self-registration after sponsor's approval. No need to provide the credentials by guest to login
	AutoLoginTimePeriod                 *int                                                                                                                      `json:"autoLoginTimePeriod,omitempty"`                 // Waiting period for auto login until sponsor's approval. If time exceeds, guest has to login manually by providing the credentials. Default value is 5 minutes
	AllowGraceAccess                    *bool                                                                                                                     `json:"allowGraceAccess,omitempty"`                    //
	GraceAccessExpireInterval           *int                                                                                                                      `json:"graceAccessExpireInterval,omitempty"`           //
	GraceAccessSendAccountExpiration    *bool                                                                                                                     `json:"graceAccessSendAccountExpiration,omitempty"`    //
	SendApprovalRequestTo               string                                                                                                                    `json:"sendApprovalRequestTo,omitempty"`               // Specifies where approval requests are sent. Only valid if requireGuestApproval = true. Allowed Values: - SELECTEDEMAILADDRESSES, - PERSONBEINGVISITED
	ApprovalEmailAddresses              string                                                                                                                    `json:"approvalEmailAddresses,omitempty"`              // Only valid if requireGuestApproval = true and sendApprovalRequestTo = SELECTEDEMAILADDRESSES
	PostRegistrationRedirect            string                                                                                                                    `json:"postRegistrationRedirect,omitempty"`            // After the registration submission direct the guest user to one of the following pages. Only valid if requireGuestApproval = true. Allowed Values: - SELFREGISTRATIONSUCCESS, - LOGINPAGEWITHINSTRUCTIONS - URL
	PostRegistrationRedirectURL         string                                                                                                                    `json:"postRegistrationRedirectUrl,omitempty"`         // URL where guest user is redirected after registration. Only valid if requireGuestApproval = true and postRegistrationRedirect = URL
	CredentialNotificationUsingEmail    *bool                                                                                                                     `json:"credentialNotificationUsingEmail,omitempty"`    // If true, send credential notification upon approval using email. Only valid if requireGuestApproval = true
	CredentialNotificationUsingSms      *bool                                                                                                                     `json:"credentialNotificationUsingSms,omitempty"`      // If true, send credential notification upon approval using SMS. Only valid if requireGuestApproval = true
	ApproveDenyLinksValidFor            *int                                                                                                                      `json:"approveDenyLinksValidFor,omitempty"`            // This attribute, along with approveDenyLinksTimeUnits, specifies how long the link can be used. Only valid if requireGuestApproval = true
	ApproveDenyLinksTimeUnits           string                                                                                                                    `json:"approveDenyLinksTimeUnits,omitempty"`           // This attribute, along with approveDenyLinksValidFor, specifies how long the link can be used. Only valid if requireGuestApproval = true. Allowed Values: - DAYS, - HOURS, - MINUTES
	RequireApproverToAuthenticate       *bool                                                                                                                     `json:"requireApproverToAuthenticate,omitempty"`       // When self-registered guests require approval, an approval request is e-mailed to one or more sponsor users. If the Cisco ISE Administrator chooses to include an approval link in the e-mail, a sponsor user who clicks the link will be required to enter their username and password if this attribute is true. Only valid if requireGuestApproval = true
	AuthenticateSponsorsUsingPortalList string                                                                                                                    `json:"authenticateSponsorsUsingPortalList,omitempty"` //
	SponsorPortalList                   *[]interface{}                                                                                                            `json:"sponsorPortalList,omitempty"`                   //
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldUserName struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldFirstName struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLastName struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldEmailAddr struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPhoneNo struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldCompany struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLocation struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldSmsProvider struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPersonBeingVisited struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldReasonForVisit struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegSuccessSettings struct {
	IncludeUserName                       *bool `json:"includeUserName,omitempty"`                       //
	IncludePassword                       *bool `json:"includePassword,omitempty"`                       //
	IncludeFirstName                      *bool `json:"includeFirstName,omitempty"`                      //
	IncludeLastName                       *bool `json:"includeLastName,omitempty"`                       //
	IncludeEmailAddr                      *bool `json:"includeEmailAddr,omitempty"`                      //
	IncludePhoneNo                        *bool `json:"includePhoneNo,omitempty"`                        //
	IncludeCompany                        *bool `json:"includeCompany,omitempty"`                        //
	IncludeLocation                       *bool `json:"includeLocation,omitempty"`                       //
	IncludeSmsProvider                    *bool `json:"includeSmsProvider,omitempty"`                    //
	IncludePersonBeingVisited             *bool `json:"includePersonBeingVisited,omitempty"`             //
	IncludeReasonForVisit                 *bool `json:"includeReasonForVisit,omitempty"`                 //
	AllowGuestSendSelfUsingPrint          *bool `json:"allowGuestSendSelfUsingPrint,omitempty"`          //
	AllowGuestSendSelfUsingEmail          *bool `json:"allowGuestSendSelfUsingEmail,omitempty"`          //
	AllowGuestSendSelfUsingSms            *bool `json:"allowGuestSendSelfUsingSms,omitempty"`            //
	IncludeAup                            *bool `json:"includeAup,omitempty"`                            //
	AupOnPage                             *bool `json:"aupOnPage,omitempty"`                             //
	RequireAupAcceptance                  *bool `json:"requireAupAcceptance,omitempty"`                  //
	RequireAupScrolling                   *bool `json:"requireAupScrolling,omitempty"`                   //
	AllowGuestLoginFromSelfregSuccessPage *bool `json:"allowGuestLoginFromSelfregSuccessPage,omitempty"` //
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsAupSettings struct {
	IncludeAup                   *bool  `json:"includeAup,omitempty"`                   // Require the portal user to read and accept an AUP
	UseDiffAupForEmployees       *bool  `json:"useDiffAupForEmployees,omitempty"`       // Only valid if requireAupAcceptance = trueG
	SkipAupForEmployees          *bool  `json:"skipAupForEmployees,omitempty"`          // Only valid if requireAupAcceptance = trueG
	RequireScrolling             *bool  `json:"requireScrolling,omitempty"`             //
	RequireAupScrolling          *bool  `json:"requireAupScrolling,omitempty"`          // Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true
	DisplayFrequency             string `json:"displayFrequency,omitempty"`             // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed Values: - FIRSTLOGIN, - EVERYLOGIN, - RECURRING
	DisplayFrequencyIntervalDays *int   `json:"displayFrequencyIntervalDays,omitempty"` // Number of days between AUP confirmations (when displayFrequency = recurring)
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsGuestChangePasswordSettings struct {
	AllowChangePasswdAtFirstLogin *bool `json:"allowChangePasswdAtFirstLogin,omitempty"` // Allow guest to change their own passwords
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsGuestDeviceRegistrationSettings struct {
	AutoRegisterGuestDevices     *bool `json:"autoRegisterGuestDevices,omitempty"`     // Automatically register guest devices
	AllowGuestsToRegisterDevices *bool `json:"allowGuestsToRegisterDevices,omitempty"` // Allow guests to register devices
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettings struct {
	ByodWelcomeSettings             *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodWelcomeSettings             `json:"byodWelcomeSettings,omitempty"`             // Configuration of BYOD endpoint welcome step configuration
	ByodRegistrationSettings        *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSettings        `json:"byodRegistrationSettings,omitempty"`        // Configuration of BYOD endpoint Registration step configuration
	ByodRegistrationSuccessSettings *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSuccessSettings `json:"byodRegistrationSuccessSettings,omitempty"` // Configuration of BYOD endpoint Registration Success step configuration
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodWelcomeSettings struct {
	EnableByod           *bool  `json:"enableBYOD,omitempty"`           //
	EnableGuestAccess    *bool  `json:"enableGuestAccess,omitempty"`    //
	RequireMdm           *bool  `json:"requireMDM,omitempty"`           //
	IncludeAup           *bool  `json:"includeAup,omitempty"`           //
	AupDisplay           string `json:"aupDisplay,omitempty"`           // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: - ONPAGE, - ASLINK
	RequireAupAcceptance *bool  `json:"requireAupAcceptance,omitempty"` //
	RequireScrolling     *bool  `json:"requireScrolling,omitempty"`     // Require BYOD devices to scroll down to the bottom of the AUP,  Only valid if includeAup = true
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSettings struct {
	ShowDeviceID            *bool  `json:"showDeviceID,omitempty"`            // Display Device ID field during registration
	EndPointIDentityGroupID string `json:"endPointIdentityGroupId,omitempty"` // Identity group id for which endpoint belongs
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` // After an Authentication Success where should device be redirected. Allowed values: - AUTHSUCCESSPAGE, - ORIGINATINGURL, - URL
	RedirectURL     string `json:"redirectUrl,omitempty"`     // Target URL for redirection, used when successRedirect = URL
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPostLoginBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` // Include a Post-Login Banner page
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPostAccessBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` //
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsAuthSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` //
	RedirectURL     string `json:"redirectUrl,omitempty"`     //
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  *bool  `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          *bool  `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        *bool  `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent *bool  `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     *bool  `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      *bool  `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizations struct {
	PortalTheme          *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          //
	PortalTweakSettings  *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` //
	PageCustomizations   *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   // Represent the entire page customization as a giant dictionary
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BannerImage      *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BackgroundImage  *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerTitle      string                                                                                                                    `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                                                    `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                                                    `json:"footerElement,omitempty"`    //
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizations struct {
	Data *[]RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizationsData `json:"data,omitempty"` // The Dictionary will be exposed here as key value pair
}

type RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortal struct {
	SelfRegPortal *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortal `json:"SelfRegPortal,omitempty"` //
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortal struct {
	Name           string                                                                            `json:"name,omitempty"`           //
	Description    string                                                                            `json:"description,omitempty"`    //
	PortalType     string                                                                            `json:"portalType,omitempty"`     // Allowed values: - BYOD, - HOTSPOTGUEST, - MYDEVICE, - SELFREGGUEST, - SPONSOR, - SPONSOREDGUEST
	PortalTestURL  string                                                                            `json:"portalTestUrl,omitempty"`  // URL to bring up a test page for this portal
	Settings       *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettings       `json:"settings,omitempty"`       // Defines all of the settings groups available for a portal
	Customizations *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizations `json:"customizations,omitempty"` // Defines all of the Portal Customizations available
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettings struct {
	PortalSettings                  *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPortalSettings                  `json:"portalSettings,omitempty"`                  // The port, interface, certificate, and other basic settings of a portal
	LoginPageSettings               *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsLoginPageSettings               `json:"loginPageSettings,omitempty"`               // Portal Login Page settings groups follow
	SelfRegPageSettings             *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettings             `json:"selfRegPageSettings,omitempty"`             //
	SelfRegSuccessSettings          *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegSuccessSettings          `json:"selfRegSuccessSettings,omitempty"`          //
	AupSettings                     *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsAupSettings                     `json:"aupSettings,omitempty"`                     //
	GuestChangePasswordSettings     *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsGuestChangePasswordSettings     `json:"guestChangePasswordSettings,omitempty"`     //
	GuestDeviceRegistrationSettings *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsGuestDeviceRegistrationSettings `json:"guestDeviceRegistrationSettings,omitempty"` //
	ByodSettings                    *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettings                    `json:"byodSettings,omitempty"`                    // Configuration of BYOD Device Welcome, Registration and Success steps
	PostLoginBannerSettings         *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPostLoginBannerSettings         `json:"postLoginBannerSettings,omitempty"`         //
	PostAccessBannerSettings        *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPostAccessBannerSettings        `json:"postAccessBannerSettings,omitempty"`        //
	AuthSuccessSettings             *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsAuthSuccessSettings             `json:"authSuccessSettings,omitempty"`             //
	SupportInfoSettings             *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSupportInfoSettings             `json:"supportInfoSettings,omitempty"`             //
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPortalSettings struct {
	HTTPSPort                    *int   `json:"httpsPort,omitempty"`                    // The port number that the allowed interfaces will listen on. Range from 8000 to 8999
	AllowedInterfaces            string `json:"allowedInterfaces,omitempty"`            // Interfaces that the portal will be reachable on. Allowed values: - eth0, - eth1, - eth2, - eth3, - eth4, - eth5, - bond0, - bond1, - bond2
	CertificateGroupTag          string `json:"certificateGroupTag,omitempty"`          // Logical name of the x.509 server certificate that will be used for the portal
	AuthenticationMethod         string `json:"authenticationMethod,omitempty"`         // Unique Id of the identity source sequence
	AssignedGuestTypeForEmployee string `json:"assignedGuestTypeForEmployee,omitempty"` // Unique Id of a guest type. Employees using this portal as a guest inherit login options from the guest type
	DisplayLang                  string `json:"displayLang,omitempty"`                  // Allowed values: - USEBROWSERLOCALE, - ALWAYSUSE
	FallbackLanguage             string `json:"fallbackLanguage,omitempty"`             // Used when displayLang = USEBROWSERLOCALE
	AlwaysUsedLanguage           string `json:"alwaysUsedLanguage,omitempty"`           //
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsLoginPageSettings struct {
	RequireAccessCode                *bool                                                                                                       `json:"requireAccessCode,omitempty"`                // Require the portal user to enter an access code
	MaxFailedAttemptsBeforeRateLimit *int                                                                                                        `json:"maxFailedAttemptsBeforeRateLimit,omitempty"` // Maximum failed login attempts before rate limiting
	TimeBetweenLoginsDuringRateLimit *int                                                                                                        `json:"timeBetweenLoginsDuringRateLimit,omitempty"` // Time between login attempts when rate limiting
	IncludeAup                       *bool                                                                                                       `json:"includeAup,omitempty"`                       // Include an Acceptable Use Policy (AUP) that should be displayed during login
	AupDisplay                       string                                                                                                      `json:"aupDisplay,omitempty"`                       // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: -  ONPAGE, - ASLINK
	RequireAupAcceptance             *bool                                                                                                       `json:"requireAupAcceptance,omitempty"`             // Require the portal user to accept the AUP. Only valid if includeAup = true
	AccessCode                       string                                                                                                      `json:"accessCode,omitempty"`                       // Access code that must be entered by the portal user (only valid if requireAccessCode = true)
	AllowGuestToCreateAccounts       *bool                                                                                                       `json:"allowGuestToCreateAccounts,omitempty"`       //
	AllowForgotPassword              *bool                                                                                                       `json:"allowForgotPassword,omitempty"`              //
	AllowGuestToChangePassword       *bool                                                                                                       `json:"allowGuestToChangePassword,omitempty"`       // Require the portal user to enter an access code
	AllowAlternateGuestPortal        *bool                                                                                                       `json:"allowAlternateGuestPortal,omitempty"`        //
	AlternateGuestPortal             string                                                                                                      `json:"alternateGuestPortal,omitempty"`             //
	AllowGuestToUseSocialAccounts    *bool                                                                                                       `json:"allowGuestToUseSocialAccounts,omitempty"`    //
	AllowShowGuestForm               *bool                                                                                                       `json:"allowShowGuestForm,omitempty"`               //
	SocialConfigs                    *[]RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsLoginPageSettingsSocialConfigs `json:"socialConfigs,omitempty"`                    //
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsLoginPageSettingsSocialConfigs struct {
	SocialMediaType  string `json:"socialMediaType,omitempty"`  //
	SocialMediaValue string `json:"socialMediaValue,omitempty"` //
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettings struct {
	AssignGuestsToGuestType             string                                                                                                                `json:"assignGuestsToGuestType,omitempty"`             // Guests are assigned to this guest type
	AccountValidityDuration             *int                                                                                                                  `json:"accountValidityDuration,omitempty"`             // Self-registered guest account is valid for this many account_validity_time_units
	AccountValidityTimeUnits            string                                                                                                                `json:"accountValidityTimeUnits,omitempty"`            // Time units for account_validity_duration. Allowed Values: - DAYS, - HOURS, - MINUTES
	RequireRegistrationCode             *bool                                                                                                                 `json:"requireRegistrationCode,omitempty"`             // Self-registered guests are required to enter a registration code
	RegistrationCode                    string                                                                                                                `json:"registrationCode,omitempty"`                    // The registration code that the guest user must enter
	FieldUserName                       *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldUserName           `json:"fieldUserName,omitempty"`                       //
	FieldFirstName                      *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldFirstName          `json:"fieldFirstName,omitempty"`                      //
	FieldLastName                       *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldLastName           `json:"fieldLastName,omitempty"`                       //
	FieldEmailAddr                      *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldEmailAddr          `json:"fieldEmailAddr,omitempty"`                      //
	FieldPhoneNo                        *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldPhoneNo            `json:"fieldPhoneNo,omitempty"`                        //
	FieldCompany                        *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldCompany            `json:"fieldCompany,omitempty"`                        //
	FieldLocation                       *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldLocation           `json:"fieldLocation,omitempty"`                       //
	SelectableLocations                 []string                                                                                                              `json:"selectableLocations,omitempty"`                 // Guests can choose from these locations to set their time zone
	FieldSmsProvider                    *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldSmsProvider        `json:"fieldSmsProvider,omitempty"`                    //
	SelectableSmsProviders              []string                                                                                                              `json:"selectableSmsProviders,omitempty"`              // This attribute is an array of SMS provider names
	FieldPersonBeingVisited             *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldPersonBeingVisited `json:"fieldPersonBeingVisited,omitempty"`             //
	FieldReasonForVisit                 *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldReasonForVisit     `json:"fieldReasonForVisit,omitempty"`                 //
	IncludeAup                          *bool                                                                                                                 `json:"includeAup,omitempty"`                          // Include an Acceptable Use Policy (AUP) that should be displayed during login
	AupDisplay                          string                                                                                                                `json:"aupDisplay,omitempty"`                          // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: - ONPAGE, - ASLINK
	RequireAupAcceptance                *bool                                                                                                                 `json:"requireAupAcceptance,omitempty"`                // Require the portal user to accept the AUP. Only valid if includeAup = true
	EnableGuestEmailWhitelist           *bool                                                                                                                 `json:"enableGuestEmailWhitelist,omitempty"`           // Allow guests with an e-mail address from selected domains
	GuestEmailWhitelistDomains          string                                                                                                                `json:"guestEmailWhitelistDomains,omitempty"`          // Self-registered guests whose e-mail address is in one of these domains will be allowed. Only valid if enableGuestEmailWhitelist = true
	EnableGuestEmailBlacklist           *bool                                                                                                                 `json:"enableGuestEmailBlacklist,omitempty"`           // Disallow guests with an e-mail address from selected domains
	GuestEmailBlacklistDomains          string                                                                                                                `json:"guestEmailBlacklistDomains,omitempty"`          // Disallow guests with an e-mail address from selected domains
	RequireGuestApproval                *bool                                                                                                                 `json:"requireGuestApproval,omitempty"`                // Require self-registered guests to be approved if true
	AutoLoginSelfWait                   *bool                                                                                                                 `json:"autoLoginSelfWait,omitempty"`                   // Allow guests to login automatically from self-registration after sponsor's approval. No need to provide the credentials by guest to login
	AutoLoginTimePeriod                 *int                                                                                                                  `json:"autoLoginTimePeriod,omitempty"`                 // Waiting period for auto login until sponsor's approval. If time exceeds, guest has to login manually by providing the credentials. Default value is 5 minutes
	AllowGraceAccess                    *bool                                                                                                                 `json:"allowGraceAccess,omitempty"`                    //
	GraceAccessExpireInterval           *int                                                                                                                  `json:"graceAccessExpireInterval,omitempty"`           //
	GraceAccessSendAccountExpiration    *bool                                                                                                                 `json:"graceAccessSendAccountExpiration,omitempty"`    //
	SendApprovalRequestTo               string                                                                                                                `json:"sendApprovalRequestTo,omitempty"`               // Specifies where approval requests are sent. Only valid if requireGuestApproval = true. Allowed Values: - SELECTEDEMAILADDRESSES, - PERSONBEINGVISITED
	ApprovalEmailAddresses              string                                                                                                                `json:"approvalEmailAddresses,omitempty"`              // Only valid if requireGuestApproval = true and sendApprovalRequestTo = SELECTEDEMAILADDRESSES
	PostRegistrationRedirect            string                                                                                                                `json:"postRegistrationRedirect,omitempty"`            // After the registration submission direct the guest user to one of the following pages. Only valid if requireGuestApproval = true. Allowed Values: - SELFREGISTRATIONSUCCESS, - LOGINPAGEWITHINSTRUCTIONS - URL
	PostRegistrationRedirectURL         string                                                                                                                `json:"postRegistrationRedirectUrl,omitempty"`         // URL where guest user is redirected after registration. Only valid if requireGuestApproval = true and postRegistrationRedirect = URL
	CredentialNotificationUsingEmail    *bool                                                                                                                 `json:"credentialNotificationUsingEmail,omitempty"`    // If true, send credential notification upon approval using email. Only valid if requireGuestApproval = true
	CredentialNotificationUsingSms      *bool                                                                                                                 `json:"credentialNotificationUsingSms,omitempty"`      // If true, send credential notification upon approval using SMS. Only valid if requireGuestApproval = true
	ApproveDenyLinksValidFor            *int                                                                                                                  `json:"approveDenyLinksValidFor,omitempty"`            // This attribute, along with approveDenyLinksTimeUnits, specifies how long the link can be used. Only valid if requireGuestApproval = true
	ApproveDenyLinksTimeUnits           string                                                                                                                `json:"approveDenyLinksTimeUnits,omitempty"`           // This attribute, along with approveDenyLinksValidFor, specifies how long the link can be used. Only valid if requireGuestApproval = true. Allowed Values: - DAYS, - HOURS, - MINUTES
	RequireApproverToAuthenticate       *bool                                                                                                                 `json:"requireApproverToAuthenticate,omitempty"`       // When self-registered guests require approval, an approval request is e-mailed to one or more sponsor users. If the Cisco ISE Administrator chooses to include an approval link in the e-mail, a sponsor user who clicks the link will be required to enter their username and password if this attribute is true. Only valid if requireGuestApproval = true
	AuthenticateSponsorsUsingPortalList string                                                                                                                `json:"authenticateSponsorsUsingPortalList,omitempty"` //
	SponsorPortalList                   *[]interface{}                                                                                                        `json:"sponsorPortalList,omitempty"`                   //
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldUserName struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldFirstName struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldLastName struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldEmailAddr struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldPhoneNo struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldCompany struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldLocation struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldSmsProvider struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldPersonBeingVisited struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldReasonForVisit struct {
	Include *bool `json:"include,omitempty"` //
	Require *bool `json:"require,omitempty"` // Only applicable if include = true
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegSuccessSettings struct {
	IncludeUserName                       *bool `json:"includeUserName,omitempty"`                       //
	IncludePassword                       *bool `json:"includePassword,omitempty"`                       //
	IncludeFirstName                      *bool `json:"includeFirstName,omitempty"`                      //
	IncludeLastName                       *bool `json:"includeLastName,omitempty"`                       //
	IncludeEmailAddr                      *bool `json:"includeEmailAddr,omitempty"`                      //
	IncludePhoneNo                        *bool `json:"includePhoneNo,omitempty"`                        //
	IncludeCompany                        *bool `json:"includeCompany,omitempty"`                        //
	IncludeLocation                       *bool `json:"includeLocation,omitempty"`                       //
	IncludeSmsProvider                    *bool `json:"includeSmsProvider,omitempty"`                    //
	IncludePersonBeingVisited             *bool `json:"includePersonBeingVisited,omitempty"`             //
	IncludeReasonForVisit                 *bool `json:"includeReasonForVisit,omitempty"`                 //
	AllowGuestSendSelfUsingPrint          *bool `json:"allowGuestSendSelfUsingPrint,omitempty"`          //
	AllowGuestSendSelfUsingEmail          *bool `json:"allowGuestSendSelfUsingEmail,omitempty"`          //
	AllowGuestSendSelfUsingSms            *bool `json:"allowGuestSendSelfUsingSms,omitempty"`            //
	IncludeAup                            *bool `json:"includeAup,omitempty"`                            //
	AupOnPage                             *bool `json:"aupOnPage,omitempty"`                             //
	RequireAupAcceptance                  *bool `json:"requireAupAcceptance,omitempty"`                  //
	RequireAupScrolling                   *bool `json:"requireAupScrolling,omitempty"`                   //
	AllowGuestLoginFromSelfregSuccessPage *bool `json:"allowGuestLoginFromSelfregSuccessPage,omitempty"` //
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsAupSettings struct {
	IncludeAup                   *bool  `json:"includeAup,omitempty"`                   // Require the portal user to read and accept an AUP
	UseDiffAupForEmployees       *bool  `json:"useDiffAupForEmployees,omitempty"`       // Only valid if requireAupAcceptance = trueG
	SkipAupForEmployees          *bool  `json:"skipAupForEmployees,omitempty"`          // Only valid if requireAupAcceptance = trueG
	RequireScrolling             *bool  `json:"requireScrolling,omitempty"`             //
	RequireAupScrolling          *bool  `json:"requireAupScrolling,omitempty"`          // Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true
	DisplayFrequency             string `json:"displayFrequency,omitempty"`             // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed Values: - FIRSTLOGIN, - EVERYLOGIN, - RECURRING
	DisplayFrequencyIntervalDays *int   `json:"displayFrequencyIntervalDays,omitempty"` // Number of days between AUP confirmations (when displayFrequency = recurring)
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsGuestChangePasswordSettings struct {
	AllowChangePasswdAtFirstLogin *bool `json:"allowChangePasswdAtFirstLogin,omitempty"` // Allow guest to change their own passwords
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsGuestDeviceRegistrationSettings struct {
	AutoRegisterGuestDevices     *bool `json:"autoRegisterGuestDevices,omitempty"`     // Automatically register guest devices
	AllowGuestsToRegisterDevices *bool `json:"allowGuestsToRegisterDevices,omitempty"` // Allow guests to register devices
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettings struct {
	ByodWelcomeSettings             *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodWelcomeSettings             `json:"byodWelcomeSettings,omitempty"`             // Configuration of BYOD endpoint welcome step configuration
	ByodRegistrationSettings        *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodRegistrationSettings        `json:"byodRegistrationSettings,omitempty"`        // Configuration of BYOD endpoint Registration step configuration
	ByodRegistrationSuccessSettings *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodRegistrationSuccessSettings `json:"byodRegistrationSuccessSettings,omitempty"` // Configuration of BYOD endpoint Registration Success step configuration
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodWelcomeSettings struct {
	EnableByod           *bool  `json:"enableBYOD,omitempty"`           //
	EnableGuestAccess    *bool  `json:"enableGuestAccess,omitempty"`    //
	RequireMdm           *bool  `json:"requireMDM,omitempty"`           //
	IncludeAup           *bool  `json:"includeAup,omitempty"`           //
	AupDisplay           string `json:"aupDisplay,omitempty"`           // How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed values: - ONPAGE, - ASLINK
	RequireAupAcceptance *bool  `json:"requireAupAcceptance,omitempty"` //
	RequireScrolling     *bool  `json:"requireScrolling,omitempty"`     // Require BYOD devices to scroll down to the bottom of the AUP,  Only valid if includeAup = true
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodRegistrationSettings struct {
	ShowDeviceID            *bool  `json:"showDeviceID,omitempty"`            // Display Device ID field during registration
	EndPointIDentityGroupID string `json:"endPointIdentityGroupId,omitempty"` // Identity group id for which endpoint belongs
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodRegistrationSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` // After an Authentication Success where should device be redirected. Allowed values: - AUTHSUCCESSPAGE, - ORIGINATINGURL, - URL
	RedirectURL     string `json:"redirectUrl,omitempty"`     // Target URL for redirection, used when successRedirect = URL
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPostLoginBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` // Include a Post-Login Banner page
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPostAccessBannerSettings struct {
	IncludePostAccessBanner *bool `json:"includePostAccessBanner,omitempty"` //
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsAuthSuccessSettings struct {
	SuccessRedirect string `json:"successRedirect,omitempty"` //
	RedirectURL     string `json:"redirectUrl,omitempty"`     //
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSupportInfoSettings struct {
	IncludeSupportInfoPage  *bool  `json:"includeSupportInfoPage,omitempty"`  //
	IncludeMacAddr          *bool  `json:"includeMacAddr,omitempty"`          //
	IncludeIPAddress        *bool  `json:"includeIpAddress,omitempty"`        //
	IncludeBrowserUserAgent *bool  `json:"includeBrowserUserAgent,omitempty"` //
	IncludePolicyServer     *bool  `json:"includePolicyServer,omitempty"`     //
	IncludeFailureCode      *bool  `json:"includeFailureCode,omitempty"`      //
	EmptyFieldDisplay       string `json:"emptyFieldDisplay,omitempty"`       // Specifies how empty fields are handled on the Support Information Page. Allowed values: - HIDE, - DISPLAYWITHNOVALUE, - DISPLAYWITHDEFAULTVALUE
	DefaultEmptyFieldValue  string `json:"defaultEmptyFieldValue,omitempty"`  // The default value displayed for an empty field. Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizations struct {
	PortalTheme          *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPortalTheme          `json:"portalTheme,omitempty"`          //
	PortalTweakSettings  *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPortalTweakSettings  `json:"portalTweakSettings,omitempty"`  // The Tweak Settings are a customization of the Portal Theme that has been selected for the portal. When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme. The Tweak Settings can subsequently be changed by the user
	Language             *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsLanguage             `json:"language,omitempty"`             // This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported
	GlobalCustomizations *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizations `json:"globalCustomizations,omitempty"` //
	PageCustomizations   *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPageCustomizations   `json:"pageCustomizations,omitempty"`   // Represent the entire page customization as a giant dictionary
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPortalTheme struct {
	ID        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      // The system- or user-assigned name of the portal theme
	ThemeData string `json:"themeData,omitempty"` // A CSS file, represented as a Base64-encoded byte array
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPortalTweakSettings struct {
	BannerColor           string `json:"bannerColor,omitempty"`           // Hex value of color
	BannerTextColor       string `json:"bannerTextColor,omitempty"`       //
	PageBackgroundColor   string `json:"pageBackgroundColor,omitempty"`   //
	PageLabelAndTextColor string `json:"pageLabelAndTextColor,omitempty"` //
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsLanguage struct {
	ViewLanguage string `json:"viewLanguage,omitempty"` //
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizations struct {
	MobileLogoImage  *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsMobileLogoImage  `json:"mobileLogoImage,omitempty"`  //
	DesktopLogoImage *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsDesktopLogoImage `json:"desktopLogoImage,omitempty"` //
	BannerImage      *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsBannerImage      `json:"bannerImage,omitempty"`      //
	BackgroundImage  *RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsBackgroundImage  `json:"backgroundImage,omitempty"`  //
	BannerTitle      string                                                                                                                `json:"bannerTitle,omitempty"`      //
	ContactText      string                                                                                                                `json:"contactText,omitempty"`      //
	FooterElement    string                                                                                                                `json:"footerElement,omitempty"`    //
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsMobileLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsDesktopLogoImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsBannerImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsBackgroundImage struct {
	Data string `json:"data,omitempty"` // Represented as base 64 encoded string of the image byte array
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPageCustomizations struct {
	Data *[]RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPageCustomizationsData `json:"data,omitempty"` // The Dictionary will be exposed here as key value pair
}

type RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPageCustomizationsData struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

//GetSelfRegisteredPortalByID Get self registered portal by ID
/* This API allows the client to get a self registered portal by ID.

@param id id path parameter.
*/
func (s *SelfRegisteredPortalService) GetSelfRegisteredPortalByID(id string) (*ResponseSelfRegisteredPortalGetSelfRegisteredPortalByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/selfregportal/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSelfRegisteredPortalGetSelfRegisteredPortalByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSelfRegisteredPortalById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSelfRegisteredPortalGetSelfRegisteredPortalByID)
	return result, response, err

}

//GetSelfRegisteredPortals Get all self registered portals
/* This API allows the client to get all the self registered portals.

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

@param getSelfRegisteredPortalsQueryParams Filtering parameter
*/
func (s *SelfRegisteredPortalService) GetSelfRegisteredPortals(getSelfRegisteredPortalsQueryParams *GetSelfRegisteredPortalsQueryParams) (*ResponseSelfRegisteredPortalGetSelfRegisteredPortals, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/selfregportal"

	queryString, _ := query.Values(getSelfRegisteredPortalsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSelfRegisteredPortalGetSelfRegisteredPortals{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSelfRegisteredPortals")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSelfRegisteredPortalGetSelfRegisteredPortals)
	return result, response, err

}

//GetVersion Get self registered portal version information
/* This API helps to retrieve the version information related to the self registered portal.

 */
func (s *SelfRegisteredPortalService) GetVersion() (*ResponseSelfRegisteredPortalGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/selfregportal/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSelfRegisteredPortalGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSelfRegisteredPortalGetVersion)
	return result, response, err

}

//CreateSelfRegisteredPortal Create self registered portal
/* This API creates a self registered portal.

 */
func (s *SelfRegisteredPortalService) CreateSelfRegisteredPortal(requestSelfRegisteredPortalCreateSelfRegisteredPortal *RequestSelfRegisteredPortalCreateSelfRegisteredPortal) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/selfregportal"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSelfRegisteredPortalCreateSelfRegisteredPortal).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateSelfRegisteredPortal")
	}

	return response, err

}

//UpdateSelfRegisteredPortalByID Update self registered portal by ID
/* This API allows the client to update a self registered portal by ID.

@param id id path parameter.
*/
func (s *SelfRegisteredPortalService) UpdateSelfRegisteredPortalByID(id string, requestSelfRegisteredPortalUpdateSelfRegisteredPortalById *RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByID) (*ResponseSelfRegisteredPortalUpdateSelfRegisteredPortalByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/selfregportal/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSelfRegisteredPortalUpdateSelfRegisteredPortalById).
		SetResult(&ResponseSelfRegisteredPortalUpdateSelfRegisteredPortalByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSelfRegisteredPortalById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSelfRegisteredPortalUpdateSelfRegisteredPortalByID)
	return result, response, err

}

//DeleteSelfRegisteredPortalByID Delete self registered portal by ID
/* This API deletes a self registered portal by ID.

@param id id path parameter.
*/
func (s *SelfRegisteredPortalService) DeleteSelfRegisteredPortalByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/selfregportal/{id}"
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
		return response, fmt.Errorf("error with operation DeleteSelfRegisteredPortalById")
	}

	getCSFRToken(response.Header())

	return response, err

}
