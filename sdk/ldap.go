package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type LdapService service

type GetLdapQueryParams struct {
	Page int `url:"page,omitempty"` //Page Number (0...N)
	Size int `url:"size,omitempty"` //Items by Page
}

type ResponseLdapGetLdapRootcacertificates struct {
	ERSLdap *ResponseLdapGetLdapRootcacertificatesERSLdap `json:"ERSLdap,omitempty"` //
}

type ResponseLdapGetLdapRootcacertificatesERSLdap struct {
	GeneralSettings          *ResponseLdapGetLdapRootcacertificatesERSLdapGeneralSettings       `json:"generalSettings,omitempty"`          //
	ConnectionSettings       *ResponseLdapGetLdapRootcacertificatesERSLdapConnectionSettings    `json:"connectionSettings,omitempty"`       //
	DirectoryOrganization    *ResponseLdapGetLdapRootcacertificatesERSLdapDirectoryOrganization `json:"directoryOrganization,omitempty"`    //
	Groups                   *ResponseLdapGetLdapRootcacertificatesERSLdapGroups                `json:"groups,omitempty"`                   //
	Attributes               *ResponseLdapGetLdapRootcacertificatesERSLdapAttributes            `json:"attributes,omitempty"`               //
	EnablePasswordChangeLDAP *bool                                                              `json:"enablePasswordChangeLDAP,omitempty"` // enablePasswordChangeLDAP
	Name                     string                                                             `json:"name,omitempty"`                     // name
	ID                       string                                                             `json:"id,omitempty"`                       // Id
	Description              string                                                             `json:"description,omitempty"`              // Description
	Link                     *ResponseLdapGetLdapRootcacertificatesERSLdapLink                  `json:"link,omitempty"`                     //
}

type ResponseLdapGetLdapRootcacertificatesERSLdapGeneralSettings struct {
	UserObjectClass       string                                                                           `json:"userObjectClass,omitempty"`       // userObjectClass.
	UserNameAttribute     string                                                                           `json:"userNameAttribute,omitempty"`     // userNameAttribute
	GroupNameAttribute    string                                                                           `json:"groupNameAttribute,omitempty"`    // groupNameAttribute
	GroupObjectClass      string                                                                           `json:"groupObjectClass,omitempty"`      // groupObjectClass
	GroupMapAttributeName string                                                                           `json:"groupMapAttributeName,omitempty"` // groupMapAttributeName
	Certificate           string                                                                           `json:"certificate,omitempty"`           // certificate
	Schema                *ResponseLdapGetLdapRootcacertificatesERSLdapGeneralSettingsSchema               `json:"schema,omitempty"`                // schema
	GroupMemberReference  *ResponseLdapGetLdapRootcacertificatesERSLdapGeneralSettingsGroupMemberReference `json:"groupMemberReference,omitempty"`  // groupMemberReference
	UserInfoAttributes    *ResponseLdapGetLdapRootcacertificatesERSLdapGeneralSettingsUserInfoAttributes   `json:"userInfoAttributes,omitempty"`    //
}

type ResponseLdapGetLdapRootcacertificatesERSLdapGeneralSettingsSchema interface{}

type ResponseLdapGetLdapRootcacertificatesERSLdapGeneralSettingsGroupMemberReference interface{}

type ResponseLdapGetLdapRootcacertificatesERSLdapGeneralSettingsUserInfoAttributes struct {
	FirstName           string `json:"firstName,omitempty"`           // firstName
	Department          string `json:"department,omitempty"`          // department
	LastName            string `json:"lastName,omitempty"`            // lastName
	OrganizationalUnit  string `json:"organizationalUnit,omitempty"`  // organizationalUnit
	JobTitle            string `json:"jobTitle,omitempty"`            // jobTitle
	Locality            string `json:"locality,omitempty"`            // locality
	Email               string `json:"email,omitempty"`               // email
	StateOrProvince     string `json:"stateOrProvince,omitempty"`     // stateOrProvince
	Telephone           string `json:"telephone,omitempty"`           // telephone
	Country             string `json:"country,omitempty"`             // country
	StreetAddress       string `json:"streetAddress,omitempty"`       // streetAddress
	AdditionalAttribute string `json:"additionalAttribute,omitempty"` // additionalAttribute
}

type ResponseLdapGetLdapRootcacertificatesERSLdapConnectionSettings struct {
	PrimaryServer               *ResponseLdapGetLdapRootcacertificatesERSLdapConnectionSettingsPrimaryServer   `json:"primaryServer,omitempty"`               //
	SecondaryServer             *ResponseLdapGetLdapRootcacertificatesERSLdapConnectionSettingsSecondaryServer `json:"secondaryServer,omitempty"`             //
	LdapNodeData                *[]ResponseLdapGetLdapRootcacertificatesERSLdapConnectionSettingsLdapNodeData  `json:"ldapNodeData,omitempty"`                // ldapNodeData
	FailoverToSecondary         *bool                                                                          `json:"failoverToSecondary,omitempty"`         // failoverToSecondary
	FailbackRetryDelay          *float64                                                                       `json:"failbackRetryDelay,omitempty"`          // failbackRetryDelay
	SpecifyServerForEachIseNode *bool                                                                          `json:"specifyServerForEachISENode,omitempty"` // specifyServerForEachISENode
	AlwaysAccessPrimaryFirst    *bool                                                                          `json:"alwaysAccessPrimaryFirst,omitempty"`    // alwaysAccessPrimaryFirst
}

type ResponseLdapGetLdapRootcacertificatesERSLdapConnectionSettingsPrimaryServer struct {
	HostName                  string   `json:"hostName,omitempty"`                  // hostName
	Port                      *float64 `json:"port,omitempty"`                      // port
	MaxConnections            *float64 `json:"maxConnections,omitempty"`            // maxConnections
	ServerTimeout             *float64 `json:"serverTimeout,omitempty"`             // serverTimeout
	UseAdminAccess            *bool    `json:"useAdminAccess,omitempty"`            // useAdminAccess
	AdminDN                   string   `json:"adminDN,omitempty"`                   // adminDN
	AdminPassword             string   `json:"adminPassword,omitempty"`             // adminPassword
	EnableSecureConnection    *bool    `json:"enableSecureConnection,omitempty"`    // enableSecureConnection
	EnableServerIDentityCheck *bool    `json:"enableServerIdentityCheck,omitempty"` // enableServerIdentityCheck
	TrustCertificate          string   `json:"trustCertificate,omitempty"`          // trustCertificate
	IssuerCaCertificate       string   `json:"issuerCACertificate,omitempty"`       // issuerCACertificate
	EnableForceReconnect      *bool    `json:"enableForceReconnect,omitempty"`      // enableForceReconnect
	ForceReconnectTime        *float64 `json:"forceReconnectTime,omitempty"`        // forceReconnectTime
}

type ResponseLdapGetLdapRootcacertificatesERSLdapConnectionSettingsSecondaryServer struct {
	HostName                  string   `json:"hostName,omitempty"`                  // hostName
	Port                      *float64 `json:"port,omitempty"`                      // port
	MaxConnections            *float64 `json:"maxConnections,omitempty"`            // maxConnections
	ServerTimeout             *float64 `json:"serverTimeout,omitempty"`             // serverTimeout
	UseAdminAccess            *bool    `json:"useAdminAccess,omitempty"`            // useAdminAccess
	AdminDN                   string   `json:"adminDN,omitempty"`                   // adminDN
	AdminPassword             string   `json:"adminPassword,omitempty"`             // adminPassword
	EnableSecureConnection    *bool    `json:"enableSecureConnection,omitempty"`    // enableSecureConnection
	EnableServerIDentityCheck *bool    `json:"enableServerIdentityCheck,omitempty"` // enableServerIdentityCheck
	TrustCertificate          string   `json:"trustCertificate,omitempty"`          // trustCertificate
	IssuerCaCertificate       string   `json:"issuerCACertificate,omitempty"`       // issuerCACertificate
	EnableForceReconnect      *bool    `json:"enableForceReconnect,omitempty"`      // enableForceReconnect
	ForceReconnectTime        *float64 `json:"forceReconnectTime,omitempty"`        // forceReconnectTime
}

type ResponseLdapGetLdapRootcacertificatesERSLdapConnectionSettingsLdapNodeData struct {
	Name              string   `json:"name,omitempty"`              // name
	PrimaryHostname   string   `json:"primaryHostname,omitempty"`   // ipaddress
	SecondaryHostname string   `json:"secondaryHostname,omitempty"` // ipaddress
	PrimaryPort       *float64 `json:"primaryPort,omitempty"`       // primaryPort
	SecondaryPort     *float64 `json:"secondaryPort,omitempty"`     // secondaryPort
}

type ResponseLdapGetLdapRootcacertificatesERSLdapDirectoryOrganization struct {
	UserDirectorySubtree  string                                                                      `json:"userDirectorySubtree,omitempty"`  // userDirectorySubtree
	GroupDirectorySubtree string                                                                      `json:"groupDirectorySubtree,omitempty"` // groupDirectorySubtree
	MacFormat             *ResponseLdapGetLdapRootcacertificatesERSLdapDirectoryOrganizationMacFormat `json:"macFormat,omitempty"`             // macFormat
	StripPrefix           *bool                                                                       `json:"stripPrefix,omitempty"`           // stripPrefix
	StripSuffix           string                                                                      `json:"stripSuffix,omitempty"`           // stripSuffix
	PrefixSeparator       string                                                                      `json:"prefixSeparator,omitempty"`       // prefixSeparator
	SuffixSeparator       string                                                                      `json:"suffixSeparator,omitempty"`       // suffixSeparator
}

type ResponseLdapGetLdapRootcacertificatesERSLdapDirectoryOrganizationMacFormat interface{}

type ResponseLdapGetLdapRootcacertificatesERSLdapGroups struct {
	GroupsNames []string `json:"groupsNames,omitempty"` // List of groups
}

type ResponseLdapGetLdapRootcacertificatesERSLdapAttributes struct {
	Attributes []string `json:"attributes,omitempty"` // List of Attributes
}

type ResponseLdapGetLdapRootcacertificatesERSLdapLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseLdapGetLdap struct {
	SearchResult *ResponseLdapGetLdapSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseLdapGetLdapSearchResult struct {
	Total        *int                                         `json:"total,omitempty"`         //
	Resources    *[]ResponseLdapGetLdapSearchResultResources  `json:"resources,omitempty"`     //
	NextPage     *ResponseLdapGetLdapSearchResultNextPage     `json:"nextHyperLink,omitempty"` //
	PreviousPage *ResponseLdapGetLdapSearchResultPreviousPage `json:"prevHyperLink,omitempty"` //
}

type ResponseLdapGetLdapSearchResultResources struct {
	ID          string                                        `json:"id,omitempty"`          // Id description
	Name        string                                        `json:"name,omitempty"`        // name description
	Description string                                        `json:"description,omitempty"` // description
	Link        *ResponseLdapGetLdapSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseLdapGetLdapSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseLdapGetLdapSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseLdapGetLdapSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseLdapPostLdap struct {
	Operation string                          `json:"operation,omitempty"` //
	Messages  *[]ResponseLdapPostLdapMessages `json:"messages,omitempty"`  //
}

type ResponseLdapPostLdapMessages struct {
	Title string `json:"title,omitempty"` //
	Type  string `json:"type,omitempty"`  //
	Code  string `json:"code,omitempty"`  //
}

type ResponseLdapGetLdapHosts struct {
	ERSLdap *ResponseLdapGetLdapHostsERSLdap `json:"ERSLdap,omitempty"` //
}

type ResponseLdapGetLdapHostsERSLdap struct {
	GeneralSettings          *ResponseLdapGetLdapHostsERSLdapGeneralSettings       `json:"generalSettings,omitempty"`          //
	ConnectionSettings       *ResponseLdapGetLdapHostsERSLdapConnectionSettings    `json:"connectionSettings,omitempty"`       //
	DirectoryOrganization    *ResponseLdapGetLdapHostsERSLdapDirectoryOrganization `json:"directoryOrganization,omitempty"`    //
	Groups                   *ResponseLdapGetLdapHostsERSLdapGroups                `json:"groups,omitempty"`                   //
	Attributes               *ResponseLdapGetLdapHostsERSLdapAttributes            `json:"attributes,omitempty"`               //
	EnablePasswordChangeLDAP *bool                                                 `json:"enablePasswordChangeLDAP,omitempty"` // enablePasswordChangeLDAP
	Name                     string                                                `json:"name,omitempty"`                     // name
	ID                       string                                                `json:"id,omitempty"`                       // Id
	Description              string                                                `json:"description,omitempty"`              // Description
	Link                     *ResponseLdapGetLdapHostsERSLdapLink                  `json:"link,omitempty"`                     //
}

type ResponseLdapGetLdapHostsERSLdapGeneralSettings struct {
	UserObjectClass       string                                                              `json:"userObjectClass,omitempty"`       // userObjectClass.
	UserNameAttribute     string                                                              `json:"userNameAttribute,omitempty"`     // userNameAttribute
	GroupNameAttribute    string                                                              `json:"groupNameAttribute,omitempty"`    // groupNameAttribute
	GroupObjectClass      string                                                              `json:"groupObjectClass,omitempty"`      // groupObjectClass
	GroupMapAttributeName string                                                              `json:"groupMapAttributeName,omitempty"` // groupMapAttributeName
	Certificate           string                                                              `json:"certificate,omitempty"`           // certificate
	Schema                *ResponseLdapGetLdapHostsERSLdapGeneralSettingsSchema               `json:"schema,omitempty"`                // schema
	GroupMemberReference  *ResponseLdapGetLdapHostsERSLdapGeneralSettingsGroupMemberReference `json:"groupMemberReference,omitempty"`  // groupMemberReference
	UserInfoAttributes    *ResponseLdapGetLdapHostsERSLdapGeneralSettingsUserInfoAttributes   `json:"userInfoAttributes,omitempty"`    //
}

type ResponseLdapGetLdapHostsERSLdapGeneralSettingsSchema interface{}

type ResponseLdapGetLdapHostsERSLdapGeneralSettingsGroupMemberReference interface{}

type ResponseLdapGetLdapHostsERSLdapGeneralSettingsUserInfoAttributes struct {
	FirstName           string `json:"firstName,omitempty"`           // firstName
	Department          string `json:"department,omitempty"`          // department
	LastName            string `json:"lastName,omitempty"`            // lastName
	OrganizationalUnit  string `json:"organizationalUnit,omitempty"`  // organizationalUnit
	JobTitle            string `json:"jobTitle,omitempty"`            // jobTitle
	Locality            string `json:"locality,omitempty"`            // locality
	Email               string `json:"email,omitempty"`               // email
	StateOrProvince     string `json:"stateOrProvince,omitempty"`     // stateOrProvince
	Telephone           string `json:"telephone,omitempty"`           // telephone
	Country             string `json:"country,omitempty"`             // country
	StreetAddress       string `json:"streetAddress,omitempty"`       // streetAddress
	AdditionalAttribute string `json:"additionalAttribute,omitempty"` // additionalAttribute
}

type ResponseLdapGetLdapHostsERSLdapConnectionSettings struct {
	PrimaryServer               *ResponseLdapGetLdapHostsERSLdapConnectionSettingsPrimaryServer   `json:"primaryServer,omitempty"`               //
	SecondaryServer             *ResponseLdapGetLdapHostsERSLdapConnectionSettingsSecondaryServer `json:"secondaryServer,omitempty"`             //
	LdapNodeData                *[]ResponseLdapGetLdapHostsERSLdapConnectionSettingsLdapNodeData  `json:"ldapNodeData,omitempty"`                // ldapNodeData
	FailoverToSecondary         *bool                                                             `json:"failoverToSecondary,omitempty"`         // failoverToSecondary
	FailbackRetryDelay          *float64                                                          `json:"failbackRetryDelay,omitempty"`          // failbackRetryDelay
	SpecifyServerForEachIseNode *bool                                                             `json:"specifyServerForEachISENode,omitempty"` // specifyServerForEachISENode
	AlwaysAccessPrimaryFirst    *bool                                                             `json:"alwaysAccessPrimaryFirst,omitempty"`    // alwaysAccessPrimaryFirst
}

type ResponseLdapGetLdapHostsERSLdapConnectionSettingsPrimaryServer struct {
	HostName                  string   `json:"hostName,omitempty"`                  // hostName
	Port                      *float64 `json:"port,omitempty"`                      // port
	MaxConnections            *float64 `json:"maxConnections,omitempty"`            // maxConnections
	ServerTimeout             *float64 `json:"serverTimeout,omitempty"`             // serverTimeout
	UseAdminAccess            *bool    `json:"useAdminAccess,omitempty"`            // useAdminAccess
	AdminDN                   string   `json:"adminDN,omitempty"`                   // adminDN
	AdminPassword             string   `json:"adminPassword,omitempty"`             // adminPassword
	EnableSecureConnection    *bool    `json:"enableSecureConnection,omitempty"`    // enableSecureConnection
	EnableServerIDentityCheck *bool    `json:"enableServerIdentityCheck,omitempty"` // enableServerIdentityCheck
	TrustCertificate          string   `json:"trustCertificate,omitempty"`          // trustCertificate
	IssuerCaCertificate       string   `json:"issuerCACertificate,omitempty"`       // issuerCACertificate
	EnableForceReconnect      *bool    `json:"enableForceReconnect,omitempty"`      // enableForceReconnect
	ForceReconnectTime        *float64 `json:"forceReconnectTime,omitempty"`        // forceReconnectTime
}

type ResponseLdapGetLdapHostsERSLdapConnectionSettingsSecondaryServer struct {
	HostName                  string   `json:"hostName,omitempty"`                  // hostName
	Port                      *float64 `json:"port,omitempty"`                      // port
	MaxConnections            *float64 `json:"maxConnections,omitempty"`            // maxConnections
	ServerTimeout             *float64 `json:"serverTimeout,omitempty"`             // serverTimeout
	UseAdminAccess            *bool    `json:"useAdminAccess,omitempty"`            // useAdminAccess
	AdminDN                   string   `json:"adminDN,omitempty"`                   // adminDN
	AdminPassword             string   `json:"adminPassword,omitempty"`             // adminPassword
	EnableSecureConnection    *bool    `json:"enableSecureConnection,omitempty"`    // enableSecureConnection
	EnableServerIDentityCheck *bool    `json:"enableServerIdentityCheck,omitempty"` // enableServerIdentityCheck
	TrustCertificate          string   `json:"trustCertificate,omitempty"`          // trustCertificate
	IssuerCaCertificate       string   `json:"issuerCACertificate,omitempty"`       // issuerCACertificate
	EnableForceReconnect      *bool    `json:"enableForceReconnect,omitempty"`      // enableForceReconnect
	ForceReconnectTime        *float64 `json:"forceReconnectTime,omitempty"`        // forceReconnectTime
}

type ResponseLdapGetLdapHostsERSLdapConnectionSettingsLdapNodeData struct {
	Name              string   `json:"name,omitempty"`              // name
	PrimaryHostname   string   `json:"primaryHostname,omitempty"`   // ipaddress
	SecondaryHostname string   `json:"secondaryHostname,omitempty"` // ipaddress
	PrimaryPort       *float64 `json:"primaryPort,omitempty"`       // primaryPort
	SecondaryPort     *float64 `json:"secondaryPort,omitempty"`     // secondaryPort
}

type ResponseLdapGetLdapHostsERSLdapDirectoryOrganization struct {
	UserDirectorySubtree  string                                                         `json:"userDirectorySubtree,omitempty"`  // userDirectorySubtree
	GroupDirectorySubtree string                                                         `json:"groupDirectorySubtree,omitempty"` // groupDirectorySubtree
	MacFormat             *ResponseLdapGetLdapHostsERSLdapDirectoryOrganizationMacFormat `json:"macFormat,omitempty"`             // macFormat
	StripPrefix           *bool                                                          `json:"stripPrefix,omitempty"`           // stripPrefix
	StripSuffix           string                                                         `json:"stripSuffix,omitempty"`           // stripSuffix
	PrefixSeparator       string                                                         `json:"prefixSeparator,omitempty"`       // prefixSeparator
	SuffixSeparator       string                                                         `json:"suffixSeparator,omitempty"`       // suffixSeparator
}

type ResponseLdapGetLdapHostsERSLdapDirectoryOrganizationMacFormat interface{}

type ResponseLdapGetLdapHostsERSLdapGroups struct {
	GroupsNames []string `json:"groupsNames,omitempty"` // List of groups
}

type ResponseLdapGetLdapHostsERSLdapAttributes struct {
	Attributes []string `json:"attributes,omitempty"` // List of Attributes
}

type ResponseLdapGetLdapHostsERSLdapLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseLdapGetLdapissuercacertificates struct {
	ERSLdap *ResponseLdapGetLdapissuercacertificatesERSLdap `json:"ERSLdap,omitempty"` //
}

type ResponseLdapGetLdapissuercacertificatesERSLdap struct {
	GeneralSettings          *ResponseLdapGetLdapissuercacertificatesERSLdapGeneralSettings       `json:"generalSettings,omitempty"`          //
	ConnectionSettings       *ResponseLdapGetLdapissuercacertificatesERSLdapConnectionSettings    `json:"connectionSettings,omitempty"`       //
	DirectoryOrganization    *ResponseLdapGetLdapissuercacertificatesERSLdapDirectoryOrganization `json:"directoryOrganization,omitempty"`    //
	Groups                   *ResponseLdapGetLdapissuercacertificatesERSLdapGroups                `json:"groups,omitempty"`                   //
	Attributes               *ResponseLdapGetLdapissuercacertificatesERSLdapAttributes            `json:"attributes,omitempty"`               //
	EnablePasswordChangeLDAP *bool                                                                `json:"enablePasswordChangeLDAP,omitempty"` // enablePasswordChangeLDAP
	Name                     string                                                               `json:"name,omitempty"`                     // name
	ID                       string                                                               `json:"id,omitempty"`                       // Id
	Description              string                                                               `json:"description,omitempty"`              // Description
	Link                     *ResponseLdapGetLdapissuercacertificatesERSLdapLink                  `json:"link,omitempty"`                     //
}

type ResponseLdapGetLdapissuercacertificatesERSLdapGeneralSettings struct {
	UserObjectClass       string                                                                             `json:"userObjectClass,omitempty"`       // userObjectClass.
	UserNameAttribute     string                                                                             `json:"userNameAttribute,omitempty"`     // userNameAttribute
	GroupNameAttribute    string                                                                             `json:"groupNameAttribute,omitempty"`    // groupNameAttribute
	GroupObjectClass      string                                                                             `json:"groupObjectClass,omitempty"`      // groupObjectClass
	GroupMapAttributeName string                                                                             `json:"groupMapAttributeName,omitempty"` // groupMapAttributeName
	Certificate           string                                                                             `json:"certificate,omitempty"`           // certificate
	Schema                *ResponseLdapGetLdapissuercacertificatesERSLdapGeneralSettingsSchema               `json:"schema,omitempty"`                // schema
	GroupMemberReference  *ResponseLdapGetLdapissuercacertificatesERSLdapGeneralSettingsGroupMemberReference `json:"groupMemberReference,omitempty"`  // groupMemberReference
	UserInfoAttributes    *ResponseLdapGetLdapissuercacertificatesERSLdapGeneralSettingsUserInfoAttributes   `json:"userInfoAttributes,omitempty"`    //
}

type ResponseLdapGetLdapissuercacertificatesERSLdapGeneralSettingsSchema interface{}

type ResponseLdapGetLdapissuercacertificatesERSLdapGeneralSettingsGroupMemberReference interface{}

type ResponseLdapGetLdapissuercacertificatesERSLdapGeneralSettingsUserInfoAttributes struct {
	FirstName           string `json:"firstName,omitempty"`           // firstName
	Department          string `json:"department,omitempty"`          // department
	LastName            string `json:"lastName,omitempty"`            // lastName
	OrganizationalUnit  string `json:"organizationalUnit,omitempty"`  // organizationalUnit
	JobTitle            string `json:"jobTitle,omitempty"`            // jobTitle
	Locality            string `json:"locality,omitempty"`            // locality
	Email               string `json:"email,omitempty"`               // email
	StateOrProvince     string `json:"stateOrProvince,omitempty"`     // stateOrProvince
	Telephone           string `json:"telephone,omitempty"`           // telephone
	Country             string `json:"country,omitempty"`             // country
	StreetAddress       string `json:"streetAddress,omitempty"`       // streetAddress
	AdditionalAttribute string `json:"additionalAttribute,omitempty"` // additionalAttribute
}

type ResponseLdapGetLdapissuercacertificatesERSLdapConnectionSettings struct {
	PrimaryServer               *ResponseLdapGetLdapissuercacertificatesERSLdapConnectionSettingsPrimaryServer   `json:"primaryServer,omitempty"`               //
	SecondaryServer             *ResponseLdapGetLdapissuercacertificatesERSLdapConnectionSettingsSecondaryServer `json:"secondaryServer,omitempty"`             //
	LdapNodeData                *[]ResponseLdapGetLdapissuercacertificatesERSLdapConnectionSettingsLdapNodeData  `json:"ldapNodeData,omitempty"`                // ldapNodeData
	FailoverToSecondary         *bool                                                                            `json:"failoverToSecondary,omitempty"`         // failoverToSecondary
	FailbackRetryDelay          *float64                                                                         `json:"failbackRetryDelay,omitempty"`          // failbackRetryDelay
	SpecifyServerForEachIseNode *bool                                                                            `json:"specifyServerForEachISENode,omitempty"` // specifyServerForEachISENode
	AlwaysAccessPrimaryFirst    *bool                                                                            `json:"alwaysAccessPrimaryFirst,omitempty"`    // alwaysAccessPrimaryFirst
}

type ResponseLdapGetLdapissuercacertificatesERSLdapConnectionSettingsPrimaryServer struct {
	HostName                  string   `json:"hostName,omitempty"`                  // hostName
	Port                      *float64 `json:"port,omitempty"`                      // port
	MaxConnections            *float64 `json:"maxConnections,omitempty"`            // maxConnections
	ServerTimeout             *float64 `json:"serverTimeout,omitempty"`             // serverTimeout
	UseAdminAccess            *bool    `json:"useAdminAccess,omitempty"`            // useAdminAccess
	AdminDN                   string   `json:"adminDN,omitempty"`                   // adminDN
	AdminPassword             string   `json:"adminPassword,omitempty"`             // adminPassword
	EnableSecureConnection    *bool    `json:"enableSecureConnection,omitempty"`    // enableSecureConnection
	EnableServerIDentityCheck *bool    `json:"enableServerIdentityCheck,omitempty"` // enableServerIdentityCheck
	TrustCertificate          string   `json:"trustCertificate,omitempty"`          // trustCertificate
	IssuerCaCertificate       string   `json:"issuerCACertificate,omitempty"`       // issuerCACertificate
	EnableForceReconnect      *bool    `json:"enableForceReconnect,omitempty"`      // enableForceReconnect
	ForceReconnectTime        *float64 `json:"forceReconnectTime,omitempty"`        // forceReconnectTime
}

type ResponseLdapGetLdapissuercacertificatesERSLdapConnectionSettingsSecondaryServer struct {
	HostName                  string   `json:"hostName,omitempty"`                  // hostName
	Port                      *float64 `json:"port,omitempty"`                      // port
	MaxConnections            *float64 `json:"maxConnections,omitempty"`            // maxConnections
	ServerTimeout             *float64 `json:"serverTimeout,omitempty"`             // serverTimeout
	UseAdminAccess            *bool    `json:"useAdminAccess,omitempty"`            // useAdminAccess
	AdminDN                   string   `json:"adminDN,omitempty"`                   // adminDN
	AdminPassword             string   `json:"adminPassword,omitempty"`             // adminPassword
	EnableSecureConnection    *bool    `json:"enableSecureConnection,omitempty"`    // enableSecureConnection
	EnableServerIDentityCheck *bool    `json:"enableServerIdentityCheck,omitempty"` // enableServerIdentityCheck
	TrustCertificate          string   `json:"trustCertificate,omitempty"`          // trustCertificate
	IssuerCaCertificate       string   `json:"issuerCACertificate,omitempty"`       // issuerCACertificate
	EnableForceReconnect      *bool    `json:"enableForceReconnect,omitempty"`      // enableForceReconnect
	ForceReconnectTime        *float64 `json:"forceReconnectTime,omitempty"`        // forceReconnectTime
}

type ResponseLdapGetLdapissuercacertificatesERSLdapConnectionSettingsLdapNodeData struct {
	Name              string   `json:"name,omitempty"`              // name
	PrimaryHostname   string   `json:"primaryHostname,omitempty"`   // ipaddress
	SecondaryHostname string   `json:"secondaryHostname,omitempty"` // ipaddress
	PrimaryPort       *float64 `json:"primaryPort,omitempty"`       // primaryPort
	SecondaryPort     *float64 `json:"secondaryPort,omitempty"`     // secondaryPort
}

type ResponseLdapGetLdapissuercacertificatesERSLdapDirectoryOrganization struct {
	UserDirectorySubtree  string                                                                        `json:"userDirectorySubtree,omitempty"`  // userDirectorySubtree
	GroupDirectorySubtree string                                                                        `json:"groupDirectorySubtree,omitempty"` // groupDirectorySubtree
	MacFormat             *ResponseLdapGetLdapissuercacertificatesERSLdapDirectoryOrganizationMacFormat `json:"macFormat,omitempty"`             // macFormat
	StripPrefix           *bool                                                                         `json:"stripPrefix,omitempty"`           // stripPrefix
	StripSuffix           string                                                                        `json:"stripSuffix,omitempty"`           // stripSuffix
	PrefixSeparator       string                                                                        `json:"prefixSeparator,omitempty"`       // prefixSeparator
	SuffixSeparator       string                                                                        `json:"suffixSeparator,omitempty"`       // suffixSeparator
}

type ResponseLdapGetLdapissuercacertificatesERSLdapDirectoryOrganizationMacFormat interface{}

type ResponseLdapGetLdapissuercacertificatesERSLdapGroups struct {
	GroupsNames []string `json:"groupsNames,omitempty"` // List of groups
}

type ResponseLdapGetLdapissuercacertificatesERSLdapAttributes struct {
	Attributes []string `json:"attributes,omitempty"` // List of Attributes
}

type ResponseLdapGetLdapissuercacertificatesERSLdapLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseLdapPutLdapidTestbindprimary struct {
	Operation string                                          `json:"operation,omitempty"` //
	Messages  *[]ResponseLdapPutLdapidTestbindprimaryMessages `json:"messages,omitempty"`  //
}

type ResponseLdapPutLdapidTestbindprimaryMessages struct {
	Title string `json:"title,omitempty"` //
	Type  string `json:"type,omitempty"`  //
	Code  string `json:"code,omitempty"`  //
}

type ResponseLdapGetLdapNameName struct {
	ERSLdap *ResponseLdapGetLdapNameNameERSLdap `json:"ERSLdap,omitempty"` //
}

type ResponseLdapGetLdapNameNameERSLdap struct {
	GeneralSettings          *ResponseLdapGetLdapNameNameERSLdapGeneralSettings       `json:"generalSettings,omitempty"`          //
	ConnectionSettings       *ResponseLdapGetLdapNameNameERSLdapConnectionSettings    `json:"connectionSettings,omitempty"`       //
	DirectoryOrganization    *ResponseLdapGetLdapNameNameERSLdapDirectoryOrganization `json:"directoryOrganization,omitempty"`    //
	Groups                   *ResponseLdapGetLdapNameNameERSLdapGroups                `json:"groups,omitempty"`                   //
	Attributes               *ResponseLdapGetLdapNameNameERSLdapAttributes            `json:"attributes,omitempty"`               //
	EnablePasswordChangeLDAP *bool                                                    `json:"enablePasswordChangeLDAP,omitempty"` // enablePasswordChangeLDAP
	Name                     string                                                   `json:"name,omitempty"`                     // name
	ID                       string                                                   `json:"id,omitempty"`                       // Id
	Description              string                                                   `json:"description,omitempty"`              // Description
	Link                     *ResponseLdapGetLdapNameNameERSLdapLink                  `json:"link,omitempty"`                     //
}

type ResponseLdapGetLdapNameNameERSLdapGeneralSettings struct {
	UserObjectClass       string                                                                 `json:"userObjectClass,omitempty"`       // userObjectClass.
	UserNameAttribute     string                                                                 `json:"userNameAttribute,omitempty"`     // userNameAttribute
	GroupNameAttribute    string                                                                 `json:"groupNameAttribute,omitempty"`    // groupNameAttribute
	GroupObjectClass      string                                                                 `json:"groupObjectClass,omitempty"`      // groupObjectClass
	GroupMapAttributeName string                                                                 `json:"groupMapAttributeName,omitempty"` // groupMapAttributeName
	Certificate           string                                                                 `json:"certificate,omitempty"`           // certificate
	Schema                *ResponseLdapGetLdapNameNameERSLdapGeneralSettingsSchema               `json:"schema,omitempty"`                // schema
	GroupMemberReference  *ResponseLdapGetLdapNameNameERSLdapGeneralSettingsGroupMemberReference `json:"groupMemberReference,omitempty"`  // groupMemberReference
	UserInfoAttributes    *ResponseLdapGetLdapNameNameERSLdapGeneralSettingsUserInfoAttributes   `json:"userInfoAttributes,omitempty"`    //
}

type ResponseLdapGetLdapNameNameERSLdapGeneralSettingsSchema interface{}

type ResponseLdapGetLdapNameNameERSLdapGeneralSettingsGroupMemberReference interface{}

type ResponseLdapGetLdapNameNameERSLdapGeneralSettingsUserInfoAttributes struct {
	FirstName           string `json:"firstName,omitempty"`           // firstName
	Department          string `json:"department,omitempty"`          // department
	LastName            string `json:"lastName,omitempty"`            // lastName
	OrganizationalUnit  string `json:"organizationalUnit,omitempty"`  // organizationalUnit
	JobTitle            string `json:"jobTitle,omitempty"`            // jobTitle
	Locality            string `json:"locality,omitempty"`            // locality
	Email               string `json:"email,omitempty"`               // email
	StateOrProvince     string `json:"stateOrProvince,omitempty"`     // stateOrProvince
	Telephone           string `json:"telephone,omitempty"`           // telephone
	Country             string `json:"country,omitempty"`             // country
	StreetAddress       string `json:"streetAddress,omitempty"`       // streetAddress
	AdditionalAttribute string `json:"additionalAttribute,omitempty"` // additionalAttribute
}

type ResponseLdapGetLdapNameNameERSLdapConnectionSettings struct {
	PrimaryServer               *ResponseLdapGetLdapNameNameERSLdapConnectionSettingsPrimaryServer   `json:"primaryServer,omitempty"`               //
	SecondaryServer             *ResponseLdapGetLdapNameNameERSLdapConnectionSettingsSecondaryServer `json:"secondaryServer,omitempty"`             //
	LdapNodeData                *[]ResponseLdapGetLdapNameNameERSLdapConnectionSettingsLdapNodeData  `json:"ldapNodeData,omitempty"`                // ldapNodeData
	FailoverToSecondary         *bool                                                                `json:"failoverToSecondary,omitempty"`         // failoverToSecondary
	FailbackRetryDelay          *float64                                                             `json:"failbackRetryDelay,omitempty"`          // failbackRetryDelay
	SpecifyServerForEachIseNode *bool                                                                `json:"specifyServerForEachISENode,omitempty"` // specifyServerForEachISENode
	AlwaysAccessPrimaryFirst    *bool                                                                `json:"alwaysAccessPrimaryFirst,omitempty"`    // alwaysAccessPrimaryFirst
}

type ResponseLdapGetLdapNameNameERSLdapConnectionSettingsPrimaryServer struct {
	HostName                  string   `json:"hostName,omitempty"`                  // hostName
	Port                      *float64 `json:"port,omitempty"`                      // port
	MaxConnections            *float64 `json:"maxConnections,omitempty"`            // maxConnections
	ServerTimeout             *float64 `json:"serverTimeout,omitempty"`             // serverTimeout
	UseAdminAccess            *bool    `json:"useAdminAccess,omitempty"`            // useAdminAccess
	AdminDN                   string   `json:"adminDN,omitempty"`                   // adminDN
	AdminPassword             string   `json:"adminPassword,omitempty"`             // adminPassword
	EnableSecureConnection    *bool    `json:"enableSecureConnection,omitempty"`    // enableSecureConnection
	EnableServerIDentityCheck *bool    `json:"enableServerIdentityCheck,omitempty"` // enableServerIdentityCheck
	TrustCertificate          string   `json:"trustCertificate,omitempty"`          // trustCertificate
	IssuerCaCertificate       string   `json:"issuerCACertificate,omitempty"`       // issuerCACertificate
	EnableForceReconnect      *bool    `json:"enableForceReconnect,omitempty"`      // enableForceReconnect
	ForceReconnectTime        *float64 `json:"forceReconnectTime,omitempty"`        // forceReconnectTime
}

type ResponseLdapGetLdapNameNameERSLdapConnectionSettingsSecondaryServer struct {
	HostName                  string   `json:"hostName,omitempty"`                  // hostName
	Port                      *float64 `json:"port,omitempty"`                      // port
	MaxConnections            *float64 `json:"maxConnections,omitempty"`            // maxConnections
	ServerTimeout             *float64 `json:"serverTimeout,omitempty"`             // serverTimeout
	UseAdminAccess            *bool    `json:"useAdminAccess,omitempty"`            // useAdminAccess
	AdminDN                   string   `json:"adminDN,omitempty"`                   // adminDN
	AdminPassword             string   `json:"adminPassword,omitempty"`             // adminPassword
	EnableSecureConnection    *bool    `json:"enableSecureConnection,omitempty"`    // enableSecureConnection
	EnableServerIDentityCheck *bool    `json:"enableServerIdentityCheck,omitempty"` // enableServerIdentityCheck
	TrustCertificate          string   `json:"trustCertificate,omitempty"`          // trustCertificate
	IssuerCaCertificate       string   `json:"issuerCACertificate,omitempty"`       // issuerCACertificate
	EnableForceReconnect      *bool    `json:"enableForceReconnect,omitempty"`      // enableForceReconnect
	ForceReconnectTime        *float64 `json:"forceReconnectTime,omitempty"`        // forceReconnectTime
}

type ResponseLdapGetLdapNameNameERSLdapConnectionSettingsLdapNodeData struct {
	Name              string   `json:"name,omitempty"`              // name
	PrimaryHostname   string   `json:"primaryHostname,omitempty"`   // ipaddress
	SecondaryHostname string   `json:"secondaryHostname,omitempty"` // ipaddress
	PrimaryPort       *float64 `json:"primaryPort,omitempty"`       // primaryPort
	SecondaryPort     *float64 `json:"secondaryPort,omitempty"`     // secondaryPort
}

type ResponseLdapGetLdapNameNameERSLdapDirectoryOrganization struct {
	UserDirectorySubtree  string                                                            `json:"userDirectorySubtree,omitempty"`  // userDirectorySubtree
	GroupDirectorySubtree string                                                            `json:"groupDirectorySubtree,omitempty"` // groupDirectorySubtree
	MacFormat             *ResponseLdapGetLdapNameNameERSLdapDirectoryOrganizationMacFormat `json:"macFormat,omitempty"`             // macFormat
	StripPrefix           *bool                                                             `json:"stripPrefix,omitempty"`           // stripPrefix
	StripSuffix           string                                                            `json:"stripSuffix,omitempty"`           // stripSuffix
	PrefixSeparator       string                                                            `json:"prefixSeparator,omitempty"`       // prefixSeparator
	SuffixSeparator       string                                                            `json:"suffixSeparator,omitempty"`       // suffixSeparator
}

type ResponseLdapGetLdapNameNameERSLdapDirectoryOrganizationMacFormat interface{}

type ResponseLdapGetLdapNameNameERSLdapGroups struct {
	GroupsNames []string `json:"groupsNames,omitempty"` // List of groups
}

type ResponseLdapGetLdapNameNameERSLdapAttributes struct {
	Attributes []string `json:"attributes,omitempty"` // List of Attributes
}

type ResponseLdapGetLdapNameNameERSLdapLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseLdapPutLdapidTestbindsecondary struct {
	Operation string                                            `json:"operation,omitempty"` //
	Messages  *[]ResponseLdapPutLdapidTestbindsecondaryMessages `json:"messages,omitempty"`  //
}

type ResponseLdapPutLdapidTestbindsecondaryMessages struct {
	Title string `json:"title,omitempty"` //
	Type  string `json:"type,omitempty"`  //
	Code  string `json:"code,omitempty"`  //
}

type ResponseLdapGetLdapid struct {
	ERSLdap *ResponseLdapGetLdapidERSLdap `json:"ERSLdap,omitempty"` //
}

type ResponseLdapGetLdapidERSLdap struct {
	GeneralSettings          *ResponseLdapGetLdapidERSLdapGeneralSettings       `json:"generalSettings,omitempty"`          //
	ConnectionSettings       *ResponseLdapGetLdapidERSLdapConnectionSettings    `json:"connectionSettings,omitempty"`       //
	DirectoryOrganization    *ResponseLdapGetLdapidERSLdapDirectoryOrganization `json:"directoryOrganization,omitempty"`    //
	Groups                   *ResponseLdapGetLdapidERSLdapGroups                `json:"groups,omitempty"`                   //
	Attributes               *ResponseLdapGetLdapidERSLdapAttributes            `json:"attributes,omitempty"`               //
	EnablePasswordChangeLDAP *bool                                              `json:"enablePasswordChangeLDAP,omitempty"` // enablePasswordChangeLDAP
	Name                     string                                             `json:"name,omitempty"`                     // name
	ID                       string                                             `json:"id,omitempty"`                       // Id
	Description              string                                             `json:"description,omitempty"`              // Description
	Link                     *ResponseLdapGetLdapidERSLdapLink                  `json:"link,omitempty"`                     //
}

type ResponseLdapGetLdapidERSLdapGeneralSettings struct {
	UserObjectClass       string                                                           `json:"userObjectClass,omitempty"`       // userObjectClass.
	UserNameAttribute     string                                                           `json:"userNameAttribute,omitempty"`     // userNameAttribute
	GroupNameAttribute    string                                                           `json:"groupNameAttribute,omitempty"`    // groupNameAttribute
	GroupObjectClass      string                                                           `json:"groupObjectClass,omitempty"`      // groupObjectClass
	GroupMapAttributeName string                                                           `json:"groupMapAttributeName,omitempty"` // groupMapAttributeName
	Certificate           string                                                           `json:"certificate,omitempty"`           // certificate
	Schema                *ResponseLdapGetLdapidERSLdapGeneralSettingsSchema               `json:"schema,omitempty"`                // schema
	GroupMemberReference  *ResponseLdapGetLdapidERSLdapGeneralSettingsGroupMemberReference `json:"groupMemberReference,omitempty"`  // groupMemberReference
	UserInfoAttributes    *ResponseLdapGetLdapidERSLdapGeneralSettingsUserInfoAttributes   `json:"userInfoAttributes,omitempty"`    //
}

type ResponseLdapGetLdapidERSLdapGeneralSettingsSchema interface{}

type ResponseLdapGetLdapidERSLdapGeneralSettingsGroupMemberReference interface{}

type ResponseLdapGetLdapidERSLdapGeneralSettingsUserInfoAttributes struct {
	FirstName           string `json:"firstName,omitempty"`           // firstName
	Department          string `json:"department,omitempty"`          // department
	LastName            string `json:"lastName,omitempty"`            // lastName
	OrganizationalUnit  string `json:"organizationalUnit,omitempty"`  // organizationalUnit
	JobTitle            string `json:"jobTitle,omitempty"`            // jobTitle
	Locality            string `json:"locality,omitempty"`            // locality
	Email               string `json:"email,omitempty"`               // email
	StateOrProvince     string `json:"stateOrProvince,omitempty"`     // stateOrProvince
	Telephone           string `json:"telephone,omitempty"`           // telephone
	Country             string `json:"country,omitempty"`             // country
	StreetAddress       string `json:"streetAddress,omitempty"`       // streetAddress
	AdditionalAttribute string `json:"additionalAttribute,omitempty"` // additionalAttribute
}

type ResponseLdapGetLdapidERSLdapConnectionSettings struct {
	PrimaryServer               *ResponseLdapGetLdapidERSLdapConnectionSettingsPrimaryServer   `json:"primaryServer,omitempty"`               //
	SecondaryServer             *ResponseLdapGetLdapidERSLdapConnectionSettingsSecondaryServer `json:"secondaryServer,omitempty"`             //
	LdapNodeData                *[]ResponseLdapGetLdapidERSLdapConnectionSettingsLdapNodeData  `json:"ldapNodeData,omitempty"`                // ldapNodeData
	FailoverToSecondary         *bool                                                          `json:"failoverToSecondary,omitempty"`         // failoverToSecondary
	FailbackRetryDelay          *float64                                                       `json:"failbackRetryDelay,omitempty"`          // failbackRetryDelay
	SpecifyServerForEachIseNode *bool                                                          `json:"specifyServerForEachISENode,omitempty"` // specifyServerForEachISENode
	AlwaysAccessPrimaryFirst    *bool                                                          `json:"alwaysAccessPrimaryFirst,omitempty"`    // alwaysAccessPrimaryFirst
}

type ResponseLdapGetLdapidERSLdapConnectionSettingsPrimaryServer struct {
	HostName                  string   `json:"hostName,omitempty"`                  // hostName
	Port                      *float64 `json:"port,omitempty"`                      // port
	MaxConnections            *float64 `json:"maxConnections,omitempty"`            // maxConnections
	ServerTimeout             *float64 `json:"serverTimeout,omitempty"`             // serverTimeout
	UseAdminAccess            *bool    `json:"useAdminAccess,omitempty"`            // useAdminAccess
	AdminDN                   string   `json:"adminDN,omitempty"`                   // adminDN
	AdminPassword             string   `json:"adminPassword,omitempty"`             // adminPassword
	EnableSecureConnection    *bool    `json:"enableSecureConnection,omitempty"`    // enableSecureConnection
	EnableServerIDentityCheck *bool    `json:"enableServerIdentityCheck,omitempty"` // enableServerIdentityCheck
	TrustCertificate          string   `json:"trustCertificate,omitempty"`          // trustCertificate
	IssuerCaCertificate       string   `json:"issuerCACertificate,omitempty"`       // issuerCACertificate
	EnableForceReconnect      *bool    `json:"enableForceReconnect,omitempty"`      // enableForceReconnect
	ForceReconnectTime        *float64 `json:"forceReconnectTime,omitempty"`        // forceReconnectTime
}

type ResponseLdapGetLdapidERSLdapConnectionSettingsSecondaryServer struct {
	HostName                  string   `json:"hostName,omitempty"`                  // hostName
	Port                      *float64 `json:"port,omitempty"`                      // port
	MaxConnections            *float64 `json:"maxConnections,omitempty"`            // maxConnections
	ServerTimeout             *float64 `json:"serverTimeout,omitempty"`             // serverTimeout
	UseAdminAccess            *bool    `json:"useAdminAccess,omitempty"`            // useAdminAccess
	AdminDN                   string   `json:"adminDN,omitempty"`                   // adminDN
	AdminPassword             string   `json:"adminPassword,omitempty"`             // adminPassword
	EnableSecureConnection    *bool    `json:"enableSecureConnection,omitempty"`    // enableSecureConnection
	EnableServerIDentityCheck *bool    `json:"enableServerIdentityCheck,omitempty"` // enableServerIdentityCheck
	TrustCertificate          string   `json:"trustCertificate,omitempty"`          // trustCertificate
	IssuerCaCertificate       string   `json:"issuerCACertificate,omitempty"`       // issuerCACertificate
	EnableForceReconnect      *bool    `json:"enableForceReconnect,omitempty"`      // enableForceReconnect
	ForceReconnectTime        *float64 `json:"forceReconnectTime,omitempty"`        // forceReconnectTime
}

type ResponseLdapGetLdapidERSLdapConnectionSettingsLdapNodeData struct {
	Name              string   `json:"name,omitempty"`              // name
	PrimaryHostname   string   `json:"primaryHostname,omitempty"`   // ipaddress
	SecondaryHostname string   `json:"secondaryHostname,omitempty"` // ipaddress
	PrimaryPort       *float64 `json:"primaryPort,omitempty"`       // primaryPort
	SecondaryPort     *float64 `json:"secondaryPort,omitempty"`     // secondaryPort
}

type ResponseLdapGetLdapidERSLdapDirectoryOrganization struct {
	UserDirectorySubtree  string                                                      `json:"userDirectorySubtree,omitempty"`  // userDirectorySubtree
	GroupDirectorySubtree string                                                      `json:"groupDirectorySubtree,omitempty"` // groupDirectorySubtree
	MacFormat             *ResponseLdapGetLdapidERSLdapDirectoryOrganizationMacFormat `json:"macFormat,omitempty"`             // macFormat
	StripPrefix           *bool                                                       `json:"stripPrefix,omitempty"`           // stripPrefix
	StripSuffix           string                                                      `json:"stripSuffix,omitempty"`           // stripSuffix
	PrefixSeparator       string                                                      `json:"prefixSeparator,omitempty"`       // prefixSeparator
	SuffixSeparator       string                                                      `json:"suffixSeparator,omitempty"`       // suffixSeparator
}

type ResponseLdapGetLdapidERSLdapDirectoryOrganizationMacFormat interface{}

type ResponseLdapGetLdapidERSLdapGroups struct {
	GroupsNames []string `json:"groupsNames,omitempty"` // List of groups
}

type ResponseLdapGetLdapidERSLdapAttributes struct {
	Attributes []string `json:"attributes,omitempty"` // List of Attributes
}

type ResponseLdapGetLdapidERSLdapLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseLdapPutLdapid struct {
	UpdatedFieldsList *ResponseLdapPutLdapidUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseLdapPutLdapidUpdatedFieldsList struct {
	UpdatedField *[]ResponseLdapPutLdapidUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
}

type ResponseLdapPutLdapidUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseLdapDeleteLdapid struct {
	Operation string                              `json:"operation,omitempty"` //
	Messages  *[]ResponseLdapDeleteLdapidMessages `json:"messages,omitempty"`  //
}

type ResponseLdapDeleteLdapidMessages struct {
	Title string `json:"title,omitempty"` //
	Type  string `json:"type,omitempty"`  //
	Code  string `json:"code,omitempty"`  //
}

type RequestLdapPostLdap struct {
	ERSLdap *RequestLdapPostLdapERSLdap `json:"ERSLdap,omitempty"` //
}

type RequestLdapPostLdapERSLdap struct {
	GeneralSettings          *RequestLdapPostLdapERSLdapGeneralSettings       `json:"generalSettings,omitempty"`          //
	ConnectionSettings       *RequestLdapPostLdapERSLdapConnectionSettings    `json:"connectionSettings,omitempty"`       //
	DirectoryOrganization    *RequestLdapPostLdapERSLdapDirectoryOrganization `json:"directoryOrganization,omitempty"`    //
	Groups                   *RequestLdapPostLdapERSLdapGroups                `json:"groups,omitempty"`                   //
	Attributes               *RequestLdapPostLdapERSLdapAttributes            `json:"attributes,omitempty"`               //
	EnablePasswordChangeLDAP *bool                                            `json:"enablePasswordChangeLDAP,omitempty"` // enablePasswordChangeLDAP
	Name                     string                                           `json:"name,omitempty"`                     // name
	ID                       string                                           `json:"id,omitempty"`                       // Id
	Description              string                                           `json:"description,omitempty"`              // Description
}

type RequestLdapPostLdapERSLdapGeneralSettings struct {
	UserObjectClass       string                                                         `json:"userObjectClass,omitempty"`       // userObjectClass.
	UserNameAttribute     string                                                         `json:"userNameAttribute,omitempty"`     // userNameAttribute
	GroupNameAttribute    string                                                         `json:"groupNameAttribute,omitempty"`    // groupNameAttribute
	GroupObjectClass      string                                                         `json:"groupObjectClass,omitempty"`      // groupObjectClass
	GroupMapAttributeName string                                                         `json:"groupMapAttributeName,omitempty"` // groupMapAttributeName
	Certificate           string                                                         `json:"certificate,omitempty"`           // certificate
	Schema                *RequestLdapPostLdapERSLdapGeneralSettingsSchema               `json:"schema,omitempty"`                // schema
	GroupMemberReference  *RequestLdapPostLdapERSLdapGeneralSettingsGroupMemberReference `json:"groupMemberReference,omitempty"`  // groupMemberReference
	UserInfoAttributes    *RequestLdapPostLdapERSLdapGeneralSettingsUserInfoAttributes   `json:"userInfoAttributes,omitempty"`    //
}

type RequestLdapPostLdapERSLdapGeneralSettingsSchema interface{}

type RequestLdapPostLdapERSLdapGeneralSettingsGroupMemberReference interface{}

type RequestLdapPostLdapERSLdapGeneralSettingsUserInfoAttributes struct {
	FirstName           string `json:"firstName,omitempty"`           // firstName
	Department          string `json:"department,omitempty"`          // department
	LastName            string `json:"lastName,omitempty"`            // lastName
	OrganizationalUnit  string `json:"organizationalUnit,omitempty"`  // organizationalUnit
	JobTitle            string `json:"jobTitle,omitempty"`            // jobTitle
	Locality            string `json:"locality,omitempty"`            // locality
	Email               string `json:"email,omitempty"`               // email
	StateOrProvince     string `json:"stateOrProvince,omitempty"`     // stateOrProvince
	Telephone           string `json:"telephone,omitempty"`           // telephone
	Country             string `json:"country,omitempty"`             // country
	StreetAddress       string `json:"streetAddress,omitempty"`       // streetAddress
	AdditionalAttribute string `json:"additionalAttribute,omitempty"` // additionalAttribute
}

type RequestLdapPostLdapERSLdapConnectionSettings struct {
	PrimaryServer               *RequestLdapPostLdapERSLdapConnectionSettingsPrimaryServer   `json:"primaryServer,omitempty"`               //
	SecondaryServer             *RequestLdapPostLdapERSLdapConnectionSettingsSecondaryServer `json:"secondaryServer,omitempty"`             //
	LdapNodeData                *[]RequestLdapPostLdapERSLdapConnectionSettingsLdapNodeData  `json:"ldapNodeData,omitempty"`                // ldapNodeData
	FailoverToSecondary         *bool                                                        `json:"failoverToSecondary,omitempty"`         // failoverToSecondary
	FailbackRetryDelay          *float64                                                     `json:"failbackRetryDelay,omitempty"`          // failbackRetryDelay
	SpecifyServerForEachIseNode *bool                                                        `json:"specifyServerForEachISENode,omitempty"` // specifyServerForEachISENode
	AlwaysAccessPrimaryFirst    *bool                                                        `json:"alwaysAccessPrimaryFirst,omitempty"`    // alwaysAccessPrimaryFirst
}

type RequestLdapPostLdapERSLdapConnectionSettingsPrimaryServer struct {
	HostName                  string   `json:"hostName,omitempty"`                  // hostName
	Port                      *float64 `json:"port,omitempty"`                      // port
	MaxConnections            *float64 `json:"maxConnections,omitempty"`            // maxConnections
	ServerTimeout             *float64 `json:"serverTimeout,omitempty"`             // serverTimeout
	UseAdminAccess            *bool    `json:"useAdminAccess,omitempty"`            // useAdminAccess
	AdminDN                   string   `json:"adminDN,omitempty"`                   // adminDN
	AdminPassword             string   `json:"adminPassword,omitempty"`             // adminPassword
	EnableSecureConnection    *bool    `json:"enableSecureConnection,omitempty"`    // enableSecureConnection
	EnableServerIDentityCheck *bool    `json:"enableServerIdentityCheck,omitempty"` // enableServerIdentityCheck
	TrustCertificate          string   `json:"trustCertificate,omitempty"`          // trustCertificate
	IssuerCaCertificate       string   `json:"issuerCACertificate,omitempty"`       // issuerCACertificate
	EnableForceReconnect      *bool    `json:"enableForceReconnect,omitempty"`      // enableForceReconnect
	ForceReconnectTime        *float64 `json:"forceReconnectTime,omitempty"`        // forceReconnectTime
}

type RequestLdapPostLdapERSLdapConnectionSettingsSecondaryServer struct {
	HostName                  string   `json:"hostName,omitempty"`                  // hostName
	Port                      *float64 `json:"port,omitempty"`                      // port
	MaxConnections            *float64 `json:"maxConnections,omitempty"`            // maxConnections
	ServerTimeout             *float64 `json:"serverTimeout,omitempty"`             // serverTimeout
	UseAdminAccess            *bool    `json:"useAdminAccess,omitempty"`            // useAdminAccess
	AdminDN                   string   `json:"adminDN,omitempty"`                   // adminDN
	AdminPassword             string   `json:"adminPassword,omitempty"`             // adminPassword
	EnableSecureConnection    *bool    `json:"enableSecureConnection,omitempty"`    // enableSecureConnection
	EnableServerIDentityCheck *bool    `json:"enableServerIdentityCheck,omitempty"` // enableServerIdentityCheck
	TrustCertificate          string   `json:"trustCertificate,omitempty"`          // trustCertificate
	IssuerCaCertificate       string   `json:"issuerCACertificate,omitempty"`       // issuerCACertificate
	EnableForceReconnect      *bool    `json:"enableForceReconnect,omitempty"`      // enableForceReconnect
	ForceReconnectTime        *float64 `json:"forceReconnectTime,omitempty"`        // forceReconnectTime
}

type RequestLdapPostLdapERSLdapConnectionSettingsLdapNodeData struct {
	Name              string   `json:"name,omitempty"`              // name
	PrimaryHostname   string   `json:"primaryHostname,omitempty"`   // ipaddress
	SecondaryHostname string   `json:"secondaryHostname,omitempty"` // ipaddress
	PrimaryPort       *float64 `json:"primaryPort,omitempty"`       // primaryPort
	SecondaryPort     *float64 `json:"secondaryPort,omitempty"`     // secondaryPort
}

type RequestLdapPostLdapERSLdapDirectoryOrganization struct {
	UserDirectorySubtree  string                                                    `json:"userDirectorySubtree,omitempty"`  // userDirectorySubtree
	GroupDirectorySubtree string                                                    `json:"groupDirectorySubtree,omitempty"` // groupDirectorySubtree
	MacFormat             *RequestLdapPostLdapERSLdapDirectoryOrganizationMacFormat `json:"macFormat,omitempty"`             // macFormat
	StripPrefix           *bool                                                     `json:"stripPrefix,omitempty"`           // stripPrefix
	StripSuffix           string                                                    `json:"stripSuffix,omitempty"`           // stripSuffix
	PrefixSeparator       string                                                    `json:"prefixSeparator,omitempty"`       // prefixSeparator
	SuffixSeparator       string                                                    `json:"suffixSeparator,omitempty"`       // suffixSeparator
}

type RequestLdapPostLdapERSLdapDirectoryOrganizationMacFormat interface{}

type RequestLdapPostLdapERSLdapGroups struct {
	GroupsNames []string `json:"groupsNames,omitempty"` // List of groups
}

type RequestLdapPostLdapERSLdapAttributes struct {
	Attributes []string `json:"attributes,omitempty"` // List of Attributes
}

type RequestLdapPutLdapid struct {
	ERSLdap *RequestLdapPutLdapidERSLdap `json:"ERSLdap,omitempty"` //
}

type RequestLdapPutLdapidERSLdap struct {
	GeneralSettings          *RequestLdapPutLdapidERSLdapGeneralSettings       `json:"generalSettings,omitempty"`          //
	ConnectionSettings       *RequestLdapPutLdapidERSLdapConnectionSettings    `json:"connectionSettings,omitempty"`       //
	DirectoryOrganization    *RequestLdapPutLdapidERSLdapDirectoryOrganization `json:"directoryOrganization,omitempty"`    //
	Groups                   *RequestLdapPutLdapidERSLdapGroups                `json:"groups,omitempty"`                   //
	Attributes               *RequestLdapPutLdapidERSLdapAttributes            `json:"attributes,omitempty"`               //
	EnablePasswordChangeLDAP *bool                                             `json:"enablePasswordChangeLDAP,omitempty"` // enablePasswordChangeLDAP
	Name                     string                                            `json:"name,omitempty"`                     // name
	ID                       string                                            `json:"id,omitempty"`                       // Id
	Description              string                                            `json:"description,omitempty"`              // Description
}

type RequestLdapPutLdapidERSLdapGeneralSettings struct {
	UserObjectClass       string                                                          `json:"userObjectClass,omitempty"`       // userObjectClass.
	UserNameAttribute     string                                                          `json:"userNameAttribute,omitempty"`     // userNameAttribute
	GroupNameAttribute    string                                                          `json:"groupNameAttribute,omitempty"`    // groupNameAttribute
	GroupObjectClass      string                                                          `json:"groupObjectClass,omitempty"`      // groupObjectClass
	GroupMapAttributeName string                                                          `json:"groupMapAttributeName,omitempty"` // groupMapAttributeName
	Certificate           string                                                          `json:"certificate,omitempty"`           // certificate
	Schema                *RequestLdapPutLdapidERSLdapGeneralSettingsSchema               `json:"schema,omitempty"`                // schema
	GroupMemberReference  *RequestLdapPutLdapidERSLdapGeneralSettingsGroupMemberReference `json:"groupMemberReference,omitempty"`  // groupMemberReference
	UserInfoAttributes    *RequestLdapPutLdapidERSLdapGeneralSettingsUserInfoAttributes   `json:"userInfoAttributes,omitempty"`    //
}

type RequestLdapPutLdapidERSLdapGeneralSettingsSchema interface{}

type RequestLdapPutLdapidERSLdapGeneralSettingsGroupMemberReference interface{}

type RequestLdapPutLdapidERSLdapGeneralSettingsUserInfoAttributes struct {
	FirstName           string `json:"firstName,omitempty"`           // firstName
	Department          string `json:"department,omitempty"`          // department
	LastName            string `json:"lastName,omitempty"`            // lastName
	OrganizationalUnit  string `json:"organizationalUnit,omitempty"`  // organizationalUnit
	JobTitle            string `json:"jobTitle,omitempty"`            // jobTitle
	Locality            string `json:"locality,omitempty"`            // locality
	Email               string `json:"email,omitempty"`               // email
	StateOrProvince     string `json:"stateOrProvince,omitempty"`     // stateOrProvince
	Telephone           string `json:"telephone,omitempty"`           // telephone
	Country             string `json:"country,omitempty"`             // country
	StreetAddress       string `json:"streetAddress,omitempty"`       // streetAddress
	AdditionalAttribute string `json:"additionalAttribute,omitempty"` // additionalAttribute
}

type RequestLdapPutLdapidERSLdapConnectionSettings struct {
	PrimaryServer               *RequestLdapPutLdapidERSLdapConnectionSettingsPrimaryServer   `json:"primaryServer,omitempty"`               //
	SecondaryServer             *RequestLdapPutLdapidERSLdapConnectionSettingsSecondaryServer `json:"secondaryServer,omitempty"`             //
	LdapNodeData                *[]RequestLdapPutLdapidERSLdapConnectionSettingsLdapNodeData  `json:"ldapNodeData,omitempty"`                // ldapNodeData
	FailoverToSecondary         *bool                                                         `json:"failoverToSecondary,omitempty"`         // failoverToSecondary
	FailbackRetryDelay          *float64                                                      `json:"failbackRetryDelay,omitempty"`          // failbackRetryDelay
	SpecifyServerForEachIseNode *bool                                                         `json:"specifyServerForEachISENode,omitempty"` // specifyServerForEachISENode
	AlwaysAccessPrimaryFirst    *bool                                                         `json:"alwaysAccessPrimaryFirst,omitempty"`    // alwaysAccessPrimaryFirst
}

type RequestLdapPutLdapidERSLdapConnectionSettingsPrimaryServer struct {
	HostName                  string   `json:"hostName,omitempty"`                  // hostName
	Port                      *float64 `json:"port,omitempty"`                      // port
	MaxConnections            *float64 `json:"maxConnections,omitempty"`            // maxConnections
	ServerTimeout             *float64 `json:"serverTimeout,omitempty"`             // serverTimeout
	UseAdminAccess            *bool    `json:"useAdminAccess,omitempty"`            // useAdminAccess
	AdminDN                   string   `json:"adminDN,omitempty"`                   // adminDN
	AdminPassword             string   `json:"adminPassword,omitempty"`             // adminPassword
	EnableSecureConnection    *bool    `json:"enableSecureConnection,omitempty"`    // enableSecureConnection
	EnableServerIDentityCheck *bool    `json:"enableServerIdentityCheck,omitempty"` // enableServerIdentityCheck
	TrustCertificate          string   `json:"trustCertificate,omitempty"`          // trustCertificate
	IssuerCaCertificate       string   `json:"issuerCACertificate,omitempty"`       // issuerCACertificate
	EnableForceReconnect      *bool    `json:"enableForceReconnect,omitempty"`      // enableForceReconnect
	ForceReconnectTime        *float64 `json:"forceReconnectTime,omitempty"`        // forceReconnectTime
}

type RequestLdapPutLdapidERSLdapConnectionSettingsSecondaryServer struct {
	HostName                  string   `json:"hostName,omitempty"`                  // hostName
	Port                      *float64 `json:"port,omitempty"`                      // port
	MaxConnections            *float64 `json:"maxConnections,omitempty"`            // maxConnections
	ServerTimeout             *float64 `json:"serverTimeout,omitempty"`             // serverTimeout
	UseAdminAccess            *bool    `json:"useAdminAccess,omitempty"`            // useAdminAccess
	AdminDN                   string   `json:"adminDN,omitempty"`                   // adminDN
	AdminPassword             string   `json:"adminPassword,omitempty"`             // adminPassword
	EnableSecureConnection    *bool    `json:"enableSecureConnection,omitempty"`    // enableSecureConnection
	EnableServerIDentityCheck *bool    `json:"enableServerIdentityCheck,omitempty"` // enableServerIdentityCheck
	TrustCertificate          string   `json:"trustCertificate,omitempty"`          // trustCertificate
	IssuerCaCertificate       string   `json:"issuerCACertificate,omitempty"`       // issuerCACertificate
	EnableForceReconnect      *bool    `json:"enableForceReconnect,omitempty"`      // enableForceReconnect
	ForceReconnectTime        *float64 `json:"forceReconnectTime,omitempty"`        // forceReconnectTime
}

type RequestLdapPutLdapidERSLdapConnectionSettingsLdapNodeData struct {
	Name              string   `json:"name,omitempty"`              // name
	PrimaryHostname   string   `json:"primaryHostname,omitempty"`   // ipaddress
	SecondaryHostname string   `json:"secondaryHostname,omitempty"` // ipaddress
	PrimaryPort       *float64 `json:"primaryPort,omitempty"`       // primaryPort
	SecondaryPort     *float64 `json:"secondaryPort,omitempty"`     // secondaryPort
}

type RequestLdapPutLdapidERSLdapDirectoryOrganization struct {
	UserDirectorySubtree  string                                                     `json:"userDirectorySubtree,omitempty"`  // userDirectorySubtree
	GroupDirectorySubtree string                                                     `json:"groupDirectorySubtree,omitempty"` // groupDirectorySubtree
	MacFormat             *RequestLdapPutLdapidERSLdapDirectoryOrganizationMacFormat `json:"macFormat,omitempty"`             // macFormat
	StripPrefix           *bool                                                      `json:"stripPrefix,omitempty"`           // stripPrefix
	StripSuffix           string                                                     `json:"stripSuffix,omitempty"`           // stripSuffix
	PrefixSeparator       string                                                     `json:"prefixSeparator,omitempty"`       // prefixSeparator
	SuffixSeparator       string                                                     `json:"suffixSeparator,omitempty"`       // suffixSeparator
}

type RequestLdapPutLdapidERSLdapDirectoryOrganizationMacFormat interface{}

type RequestLdapPutLdapidERSLdapGroups struct {
	GroupsNames []string `json:"groupsNames,omitempty"` // List of groups
}

type RequestLdapPutLdapidERSLdapAttributes struct {
	Attributes []string `json:"attributes,omitempty"` // List of Attributes
}

//GetLdapRootcacertificates get-rootca-certificates
/* get-rootca-certificates

 */
func (s *LdapService) GetLdapRootcacertificates() (*ResponseLdapGetLdapRootcacertificates, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ldap/rootcacertificates"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLdapGetLdapRootcacertificates{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetLdapRootcacertificates")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseLdapGetLdapRootcacertificates)
	return result, response, err

}

//GetLdap Get-All
/* Get-All

@param getLdapQueryParams Filtering parameter
*/
func (s *LdapService) GetLdap(getLdapQueryParams *GetLdapQueryParams) (*ResponseLdapGetLdap, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ldap"

	queryString, _ := query.Values(getLdapQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLdapGetLdap{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetLdap")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseLdapGetLdap)
	return result, response, err

}

//GetLdapHosts get-hosts
/* get-hosts

 */
func (s *LdapService) GetLdapHosts() (*ResponseLdapGetLdapHosts, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ldap/hosts"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLdapGetLdapHosts{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetLdapHosts")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseLdapGetLdapHosts)
	return result, response, err

}

//GetLdapissuercacertificates get-issuerca-certificates
/* get-issuerca-certificates

 */
func (s *LdapService) GetLdapissuercacertificates() (*ResponseLdapGetLdapissuercacertificates, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ldap/issuercacertificates"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLdapGetLdapissuercacertificates{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetLdapissuercacertificates")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseLdapGetLdapissuercacertificates)
	return result, response, err

}

//GetLdapNameName Get-By-Name
/* Get-By-Name

@param name name path parameter.
*/
func (s *LdapService) GetLdapNameName(name string) (*ResponseLdapGetLdapNameName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ldap/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLdapGetLdapNameName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetLdapNameName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseLdapGetLdapNameName)
	return result, response, err

}

//GetLdapid Get-By-Id
/* Get-By-Id

@param id id path parameter.
*/
func (s *LdapService) GetLdapid(id string) (*ResponseLdapGetLdapid, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ldap/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLdapGetLdapid{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetLdapid")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseLdapGetLdapid)
	return result, response, err

}

//PostLdap Create
/* Create

 */
func (s *LdapService) PostLdap(requestLdapPostLdap *RequestLdapPostLdap) (*ResponseLdapPostLdap, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ldap"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLdapPostLdap).
		SetResult(&ResponseLdapPostLdap{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseLdapPostLdap{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation PostLdap")
	}

	result := response.Result().(*ResponseLdapPostLdap)
	return result, response, err

}

//PutLdapidTestbindprimary test-bind-primary
/* test-bind-primary

@param id id path parameter.
*/
func (s *LdapService) PutLdapidTestbindprimary(id string) (*ResponseLdapPutLdapidTestbindprimary, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ldap/{id}/testbindprimary"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLdapPutLdapidTestbindprimary{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseLdapPutLdapidTestbindprimary{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation PutLdapidTestbindprimary")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseLdapPutLdapidTestbindprimary)
	return result, response, err

}

//PutLdapidTestbindsecondary test-bind-secondary
/* test-bind-secondary

@param id id path parameter.
*/
func (s *LdapService) PutLdapidTestbindsecondary(id string) (*ResponseLdapPutLdapidTestbindsecondary, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ldap/{id}/testbindsecondary"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLdapPutLdapidTestbindsecondary{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseLdapPutLdapidTestbindsecondary{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation PutLdapidTestbindsecondary")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseLdapPutLdapidTestbindsecondary)
	return result, response, err

}

//PutLdapid Update
/* Update

@param id id path parameter.
*/
func (s *LdapService) PutLdapid(id string, requestLdapPutLdapId *RequestLdapPutLdapid) (*ResponseLdapPutLdapid, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ldap/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLdapPutLdapId).
		SetResult(&ResponseLdapPutLdapid{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseLdapPutLdapid{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation PutLdapid")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseLdapPutLdapid)
	return result, response, err

}

//DeleteLdapid Delete
/* Delete

@param id id path parameter.
*/
func (s *LdapService) DeleteLdapid(id string) (*ResponseLdapDeleteLdapid, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/ldap/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLdapDeleteLdapid{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteLdapid")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseLdapDeleteLdapid)
	return result, response, err

}
