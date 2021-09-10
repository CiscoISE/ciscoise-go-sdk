package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ActiveDirectoryService service

type GetActiveDirectoryQueryParams struct {
	Page int `url:"page,omitempty"` //Page number
	Size int `url:"size,omitempty"` //Number of objects returned per page
}

type ResponseActiveDirectoryGetActiveDirectoryByName struct {
	ERSActiveDirectory *ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectory `json:"ERSActiveDirectory,omitempty"` //
}

type ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectory struct {
	ID                      string                                                                             `json:"id,omitempty"`                      // Resource UUID value
	Name                    string                                                                             `json:"name,omitempty"`                    // Resource Name. Maximum 32 characters allowed. Allowed characters are alphanumeric and .-_/\\ characters
	Description             string                                                                             `json:"description,omitempty"`             // No character restriction
	Domain                  string                                                                             `json:"domain,omitempty"`                  // The AD domain. Alphanumeric, hyphen (-) and dot (.) characters are allowed
	EnableDomainAllowedList *bool                                                                              `json:"enableDomainAllowedList,omitempty"` //
	EnableDomainWhiteList   *bool                                                                              `json:"enableDomainWhiteList,omitempty"`   //
	Adgroups                *ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdgroups         `json:"adgroups,omitempty"`                // Holds list of AD Groups
	AdvancedSettings        *ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdvancedSettings `json:"advancedSettings,omitempty"`        //
	AdAttributes            *ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdAttributes     `json:"adAttributes,omitempty"`            // Holds list of AD Attributes
	AdScopesNames           string                                                                             `json:"adScopesNames,omitempty"`           // String that contains the names of the scopes that the active directory belongs to. Names are separated by comma. Alphanumeric, underscore (_) characters are allowed
	Link                    *ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryLink             `json:"link,omitempty"`                    //
}

type ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdgroups struct {
	Groups *[]ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdgroupsGroups `json:"groups,omitempty"` // List of Groups
}

type ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdgroupsGroups struct {
	Name string `json:"name,omitempty"` // Required for each group in the group list with no duplication between groups. All characters are allowed except %
	Sid  string `json:"sid,omitempty"`  // Cisco ISE uses security identifiers (SIDs) for optimization of group membership evaluation. SIDs are useful for efficiency (speed) when the groups are evaluated. All characters are allowed except %
	Type string `json:"type,omitempty"` // No character restriction
}

type ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdvancedSettings struct {
	EnablePassChange              *bool                                                                                            `json:"enablePassChange,omitempty"`              //
	EnableMachineAuth             *bool                                                                                            `json:"enableMachineAuth,omitempty"`             //
	EnableMachineAccess           *bool                                                                                            `json:"enableMachineAccess,omitempty"`           //
	AgingTime                     *int                                                                                             `json:"agingTime,omitempty"`                     // Range 1-8760 hours
	EnableDialinPermissionCheck   *bool                                                                                            `json:"enableDialinPermissionCheck,omitempty"`   //
	EnableCallbackForDialinClient *bool                                                                                            `json:"enableCallbackForDialinClient,omitempty"` //
	PlaintextAuth                 *bool                                                                                            `json:"plaintextAuth,omitempty"`                 //
	EnableFailedAuthProtection    *bool                                                                                            `json:"enableFailedAuthProtection,omitempty"`    // Enable prevent AD account lockout due to too many bad password attempts
	AuthProtectionType            string                                                                                           `json:"authProtectionType,omitempty"`            // Enable prevent AD account lockout. Allowed values: - WIRELESS, - WIRED, - BOTH
	FailedAuthThreshold           *int                                                                                             `json:"failedAuthThreshold,omitempty"`           // Number of bad password attempts
	IDentityNotInAdBehaviour      string                                                                                           `json:"identityNotInAdBehaviour,omitempty"`      // Allowed values: REJECT, SEARCH_JOINED_FOREST, SEARCH_ALL
	UnreachableDomainsBehaviour   string                                                                                           `json:"unreachableDomainsBehaviour,omitempty"`   // Allowed values: PROCEED, DROP
	EnableRewrites                *bool                                                                                            `json:"enableRewrites,omitempty"`                //
	RewriteRules                  *[]ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdvancedSettingsRewriteRules `json:"rewriteRules,omitempty"`                  // Identity rewrite is an advanced feature that directs Cisco ISE to manipulate the identity before it is passed to the external Active Directory system. You can create rules to change the identity to a desired format that includes or excludes a domain prefix and/or suffix or other additional markup of your choice
	FirstName                     string                                                                                           `json:"firstName,omitempty"`                     // User info attribute. All characters are allowed except %
	Department                    string                                                                                           `json:"department,omitempty"`                    // User info attribute. All characters are allowed except %
	LastName                      string                                                                                           `json:"lastName,omitempty"`                      // User info attribute. All characters are allowed except %
	OrganizationalUnit            string                                                                                           `json:"organizationalUnit,omitempty"`            // User info attribute. All characters are allowed except %
	JobTitle                      string                                                                                           `json:"jobTitle,omitempty"`                      // User info attribute. All characters are allowed except %
	Locality                      string                                                                                           `json:"locality,omitempty"`                      // User info attribute. All characters are allowed except %
	Email                         string                                                                                           `json:"email,omitempty"`                         // User info attribute. All characters are allowed except %
	StateOrProvince               string                                                                                           `json:"stateOrProvince,omitempty"`               // User info attribute. All characters are allowed except %
	Telephone                     string                                                                                           `json:"telephone,omitempty"`                     // User info attribute. All characters are allowed except %
	Country                       string                                                                                           `json:"country,omitempty"`                       // User info attribute. All characters are allowed except %
	StreetAddress                 string                                                                                           `json:"streetAddress,omitempty"`                 // User info attribute. All characters are allowed except %
	Schema                        string                                                                                           `json:"schema,omitempty"`                        // Allowed values: ACTIVE_DIRECTORY, CUSTOM. Choose ACTIVE_DIRECTORY schema when the AD attributes defined in AD can be copied to relevant attributes in Cisco ISE. If customization is needed, choose CUSTOM schema. All User info attributes are always set to default value if schema is ACTIVE_DIRECTORY. Values can be changed only for CUSTOM schema
}

type ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdvancedSettingsRewriteRules struct {
	RowID         *int   `json:"rowId,omitempty"`         // Required for each rule in the list in serial order
	RewriteMatch  string `json:"rewriteMatch,omitempty"`  // Required for each rule in the list with no duplication between rules. All characters are allowed except %"
	RewriteResult string `json:"rewriteResult,omitempty"` // Required for each rule in the list. All characters are allowed except %"
}

type ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdAttributes struct {
	Attributes *[]ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdAttributesAttributes `json:"attributes,omitempty"` // List of Attributes
}

type ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdAttributesAttributes struct {
	Name         string `json:"name,omitempty"`         // Required for each attribute in the attribute list with no duplication between attributes. All characters are allowed except <%"
	Type         string `json:"type,omitempty"`         // Required for each group in the group list. Allowed values: STRING, IP, BOOLEAN, INT, OCTET_STRING
	InternalName string `json:"internalName,omitempty"` // Required for each attribute in the attribute list. All characters are allowed except <%"
	DefaultValue string `json:"defaultValue,omitempty"` // Required for each attribute in the attribute list. Can contain an empty string. All characters are allowed except <%"
}

type ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseActiveDirectoryGetUserGroups struct {
	ERSActiveDirectoryGroups *ResponseActiveDirectoryGetUserGroupsERSActiveDirectoryGroups `json:"ERSActiveDirectoryGroups,omitempty"` //
}

type ResponseActiveDirectoryGetUserGroupsERSActiveDirectoryGroups struct {
	Groups *[]ResponseActiveDirectoryGetUserGroupsERSActiveDirectoryGroupsGroups `json:"groups,omitempty"` //
}

type ResponseActiveDirectoryGetUserGroupsERSActiveDirectoryGroupsGroups struct {
	GroupName string `json:"groupName,omitempty"` //
	Sid       string `json:"sid,omitempty"`       //
	Type      string `json:"type,omitempty"`      //
}

type ResponseActiveDirectoryIsUserMemberOfGroups struct {
	ERSActiveDirectoryGroups *ResponseActiveDirectoryIsUserMemberOfGroupsERSActiveDirectoryGroups `json:"ERSActiveDirectoryGroups,omitempty"` //
}

type ResponseActiveDirectoryIsUserMemberOfGroupsERSActiveDirectoryGroups struct {
	Groups *[]ResponseActiveDirectoryIsUserMemberOfGroupsERSActiveDirectoryGroupsGroups `json:"groups,omitempty"` //
}

type ResponseActiveDirectoryIsUserMemberOfGroupsERSActiveDirectoryGroupsGroups struct {
	GroupName string `json:"groupName,omitempty"` //
	Sid       string `json:"sid,omitempty"`       //
	Type      string `json:"type,omitempty"`      //
}

type ResponseActiveDirectoryGetTrustedDomains struct {
	ERSActiveDirectoryDomains *ResponseActiveDirectoryGetTrustedDomainsERSActiveDirectoryDomains `json:"ERSActiveDirectoryDomains,omitempty"` //
}

type ResponseActiveDirectoryGetTrustedDomainsERSActiveDirectoryDomains struct {
	Domains *[]ResponseActiveDirectoryGetTrustedDomainsERSActiveDirectoryDomainsDomains `json:"domains,omitempty"` //
}

type ResponseActiveDirectoryGetTrustedDomainsERSActiveDirectoryDomainsDomains struct {
	DNSName        string `json:"dnsName,omitempty"`        //
	Forest         string `json:"forest,omitempty"`         //
	UnusableReason string `json:"unusableReason,omitempty"` //
}

type ResponseActiveDirectoryGetGroupsByDomain struct {
	ERSActiveDirectoryGroups *ResponseActiveDirectoryGetGroupsByDomainERSActiveDirectoryGroups `json:"ERSActiveDirectoryGroups,omitempty"` //
}

type ResponseActiveDirectoryGetGroupsByDomainERSActiveDirectoryGroups struct {
	Groups *[]ResponseActiveDirectoryGetGroupsByDomainERSActiveDirectoryGroupsGroups `json:"groups,omitempty"` //
}

type ResponseActiveDirectoryGetGroupsByDomainERSActiveDirectoryGroupsGroups struct {
	GroupName string `json:"groupName,omitempty"` //
	Sid       string `json:"sid,omitempty"`       //
	Type      string `json:"type,omitempty"`      //
}

type ResponseActiveDirectoryGetActiveDirectoryByID struct {
	ERSActiveDirectory *ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectory `json:"ERSActiveDirectory,omitempty"` //
}

type ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectory struct {
	ID                      string                                                                           `json:"id,omitempty"`                      // Resource UUID value
	Name                    string                                                                           `json:"name,omitempty"`                    // Resource Name. Maximum 32 characters allowed. Allowed characters are alphanumeric and .-_/\\ characters
	Description             string                                                                           `json:"description,omitempty"`             // No character restriction
	Domain                  string                                                                           `json:"domain,omitempty"`                  // The AD domain. Alphanumeric, hyphen (-) and dot (.) characters are allowed
	EnableDomainWhiteList   *bool                                                                            `json:"enableDomainWhiteList,omitempty"`   //
	EnableDomainAllowedList *bool                                                                            `json:"enableDomainAllowedList,omitempty"` //
	Adgroups                *ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdgroups         `json:"adgroups,omitempty"`                // Holds list of AD Groups
	AdvancedSettings        *ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdvancedSettings `json:"advancedSettings,omitempty"`        //
	AdAttributes            *ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdAttributes     `json:"adAttributes,omitempty"`            // Holds list of AD Attributes
	AdScopesNames           string                                                                           `json:"adScopesNames,omitempty"`           // String that contains the names of the scopes that the active directory belongs to. Names are separated by comma. Alphanumeric, underscore (_) characters are allowed
	Link                    *ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryLink             `json:"link,omitempty"`                    //
}

type ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdgroups struct {
	Groups *[]ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdgroupsGroups `json:"groups,omitempty"` // List of Groups
}

type ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdgroupsGroups struct {
	Name string `json:"name,omitempty"` // Required for each group in the group list with no duplication between groups. All characters are allowed except %
	Sid  string `json:"sid,omitempty"`  // Cisco ISE uses security identifiers (SIDs) for optimization of group membership evaluation. SIDs are useful for efficiency (speed) when the groups are evaluated. All characters are allowed except %
	Type string `json:"type,omitempty"` // No character restriction
}

type ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdvancedSettings struct {
	EnablePassChange              *bool                                                                                          `json:"enablePassChange,omitempty"`              //
	EnableMachineAuth             *bool                                                                                          `json:"enableMachineAuth,omitempty"`             //
	EnableMachineAccess           *bool                                                                                          `json:"enableMachineAccess,omitempty"`           //
	AgingTime                     *int                                                                                           `json:"agingTime,omitempty"`                     // Range 1-8760 hours
	EnableDialinPermissionCheck   *bool                                                                                          `json:"enableDialinPermissionCheck,omitempty"`   //
	EnableCallbackForDialinClient *bool                                                                                          `json:"enableCallbackForDialinClient,omitempty"` //
	PlaintextAuth                 *bool                                                                                          `json:"plaintextAuth,omitempty"`                 //
	EnableFailedAuthProtection    *bool                                                                                          `json:"enableFailedAuthProtection,omitempty"`    // Enable prevent AD account lockout due to too many bad password attempts
	AuthProtectionType            string                                                                                         `json:"authProtectionType,omitempty"`            // Enable prevent AD account lockout. Allowed values: - WIRELESS, - WIRED, - BOTH
	FailedAuthThreshold           *int                                                                                           `json:"failedAuthThreshold,omitempty"`           // Number of bad password attempts
	IDentityNotInAdBehaviour      string                                                                                         `json:"identityNotInAdBehaviour,omitempty"`      // Allowed values: REJECT, SEARCH_JOINED_FOREST, SEARCH_ALL
	UnreachableDomainsBehaviour   string                                                                                         `json:"unreachableDomainsBehaviour,omitempty"`   // Allowed values: PROCEED, DROP
	EnableRewrites                *bool                                                                                          `json:"enableRewrites,omitempty"`                //
	RewriteRules                  *[]ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdvancedSettingsRewriteRules `json:"rewriteRules,omitempty"`                  // Identity rewrite is an advanced feature that directs Cisco ISE to manipulate the identity before it is passed to the external Active Directory system. You can create rules to change the identity to a desired format that includes or excludes a domain prefix and/or suffix or other additional markup of your choice
	FirstName                     string                                                                                         `json:"firstName,omitempty"`                     // User info attribute. All characters are allowed except %
	Department                    string                                                                                         `json:"department,omitempty"`                    // User info attribute. All characters are allowed except %
	LastName                      string                                                                                         `json:"lastName,omitempty"`                      // User info attribute. All characters are allowed except %
	OrganizationalUnit            string                                                                                         `json:"organizationalUnit,omitempty"`            // User info attribute. All characters are allowed except %
	JobTitle                      string                                                                                         `json:"jobTitle,omitempty"`                      // User info attribute. All characters are allowed except %
	Locality                      string                                                                                         `json:"locality,omitempty"`                      // User info attribute. All characters are allowed except %
	Email                         string                                                                                         `json:"email,omitempty"`                         // User info attribute. All characters are allowed except %
	StateOrProvince               string                                                                                         `json:"stateOrProvince,omitempty"`               // User info attribute. All characters are allowed except %
	Telephone                     string                                                                                         `json:"telephone,omitempty"`                     // User info attribute. All characters are allowed except %
	Country                       string                                                                                         `json:"country,omitempty"`                       // User info attribute. All characters are allowed except %
	StreetAddress                 string                                                                                         `json:"streetAddress,omitempty"`                 // User info attribute. All characters are allowed except %
	Schema                        string                                                                                         `json:"schema,omitempty"`                        // Allowed values: ACTIVE_DIRECTORY, CUSTOM. Choose ACTIVE_DIRECTORY schema when the AD attributes defined in AD can be copied to relevant attributes in Cisco ISE. If customization is needed, choose CUSTOM schema. All User info attributes are always set to default value if schema is ACTIVE_DIRECTORY. Values can be changed only for CUSTOM schema
}

type ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdvancedSettingsRewriteRules struct {
	RowID         *int   `json:"rowId,omitempty"`         // Required for each rule in the list in serial order
	RewriteMatch  string `json:"rewriteMatch,omitempty"`  // Required for each rule in the list with no duplication between rules. All characters are allowed except %"
	RewriteResult string `json:"rewriteResult,omitempty"` // Required for each rule in the list. All characters are allowed except %"
}

type ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdAttributes struct {
	Attributes *[]ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdAttributesAttributes `json:"attributes,omitempty"` // List of Attributes
}

type ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdAttributesAttributes struct {
	Name         string `json:"name,omitempty"`         // Required for each attribute in the attribute list with no duplication between attributes. All characters are allowed except <%"
	Type         string `json:"type,omitempty"`         // Required for each group in the group list. Allowed values: STRING, IP, BOOLEAN, INT, OCTET_STRING
	InternalName string `json:"internalName,omitempty"` // Required for each attribute in the attribute list. All characters are allowed except <%"
	DefaultValue string `json:"defaultValue,omitempty"` // Required for each attribute in the attribute list. Can contain an empty string. All characters are allowed except <%"
}

type ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseActiveDirectoryGetActiveDirectory struct {
	SearchResult *ResponseActiveDirectoryGetActiveDirectorySearchResult `json:"SearchResult,omitempty"` //
}

type ResponseActiveDirectoryGetActiveDirectorySearchResult struct {
	Total        *int                                                               `json:"total,omitempty"`        //
	Resources    *[]ResponseActiveDirectoryGetActiveDirectorySearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponseActiveDirectoryGetActiveDirectorySearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseActiveDirectoryGetActiveDirectorySearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseActiveDirectoryGetActiveDirectorySearchResultResources struct {
	ID          string                                                              `json:"id,omitempty"`          //
	Name        string                                                              `json:"name,omitempty"`        //
	Description string                                                              `json:"description,omitempty"` //
	Link        *ResponseActiveDirectoryGetActiveDirectorySearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseActiveDirectoryGetActiveDirectorySearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseActiveDirectoryGetActiveDirectorySearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseActiveDirectoryGetActiveDirectorySearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseActiveDirectoryGetVersion struct {
	VersionInfo *ResponseActiveDirectoryGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseActiveDirectoryGetVersionVersionInfo struct {
	CurrentServerVersion string                                            `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                            `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseActiveDirectoryGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseActiveDirectoryGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestActiveDirectoryGetUserGroups struct {
	OperationAdditionalData *RequestActiveDirectoryGetUserGroupsOperationAdditionalData `json:"OperationAdditionalData,omitempty"` //
}

type RequestActiveDirectoryGetUserGroupsOperationAdditionalData struct {
	AdditionalData *[]RequestActiveDirectoryGetUserGroupsOperationAdditionalDataAdditionalData `json:"additionalData,omitempty"` //
}

type RequestActiveDirectoryGetUserGroupsOperationAdditionalDataAdditionalData struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type RequestActiveDirectoryLoadGroupsFromDomain struct {
	ERSActiveDirectory *RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectory `json:"ERSActiveDirectory,omitempty"` //
}

type RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectory struct {
	ID                    string                                                                        `json:"id,omitempty"`                    // Resource UUID value
	Name                  string                                                                        `json:"name,omitempty"`                  // Resource Name. Maximum 32 characters allowed. Allowed characters are alphanumeric and .-_/\\ characters
	Description           string                                                                        `json:"description,omitempty"`           // No character restriction
	Domain                string                                                                        `json:"domain,omitempty"`                // The AD domain. Alphanumeric, hyphen (-) and dot (.) characters are allowed
	EnableDomainWhiteList *bool                                                                         `json:"enableDomainWhiteList,omitempty"` //
	Adgroups              *RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdgroups         `json:"adgroups,omitempty"`              // Holds list of AD Groups
	AdvancedSettings      *RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdvancedSettings `json:"advancedSettings,omitempty"`      //
	AdAttributes          *RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdAttributes     `json:"adAttributes,omitempty"`          // Holds list of AD Attributes
	AdScopesNames         string                                                                        `json:"adScopesNames,omitempty"`         // String that contains the names of the scopes that the active directory belongs to. Names are separated by comma. Alphanumeric, underscore (_) characters are allowed
}

type RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdgroups struct {
	Groups *[]RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdgroupsGroups `json:"groups,omitempty"` // List of Groups
}

type RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdgroupsGroups struct {
	Name string `json:"name,omitempty"` // Required for each group in the group list with no duplication between groups. All characters are allowed except %
	Sid  string `json:"sid,omitempty"`  // Cisco ISE uses security identifiers (SIDs) for optimization of group membership evaluation. SIDs are useful for efficiency (speed) when the groups are evaluated. All characters are allowed except %
	Type string `json:"type,omitempty"` // No character restriction
}

type RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdvancedSettings struct {
	EnablePassChange              *bool                                                                                       `json:"enablePassChange,omitempty"`              //
	EnableMachineAuth             *bool                                                                                       `json:"enableMachineAuth,omitempty"`             //
	EnableMachineAccess           *bool                                                                                       `json:"enableMachineAccess,omitempty"`           //
	AgingTime                     *int                                                                                        `json:"agingTime,omitempty"`                     // Range 1-8760 hours
	EnableDialinPermissionCheck   *bool                                                                                       `json:"enableDialinPermissionCheck,omitempty"`   //
	EnableCallbackForDialinClient *bool                                                                                       `json:"enableCallbackForDialinClient,omitempty"` //
	PlaintextAuth                 *bool                                                                                       `json:"plaintextAuth,omitempty"`                 //
	EnableFailedAuthProtection    *bool                                                                                       `json:"enableFailedAuthProtection,omitempty"`    // Enable prevent AD account lockout due to too many bad password attempts
	AuthProtectionType            string                                                                                      `json:"authProtectionType,omitempty"`            // Enable prevent AD account lockout. Allowed values: - WIRELESS, - WIRED, - BOTH
	FailedAuthThreshold           *int                                                                                        `json:"failedAuthThreshold,omitempty"`           // Number of bad password attempts
	IDentityNotInAdBehaviour      string                                                                                      `json:"identityNotInAdBehaviour,omitempty"`      // Allowed values: REJECT, SEARCH_JOINED_FOREST, SEARCH_ALL
	UnreachableDomainsBehaviour   string                                                                                      `json:"unreachableDomainsBehaviour,omitempty"`   // Allowed values: PROCEED, DROP
	EnableRewrites                *bool                                                                                       `json:"enableRewrites,omitempty"`                //
	RewriteRules                  *[]RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdvancedSettingsRewriteRules `json:"rewriteRules,omitempty"`                  // Identity rewrite is an advanced feature that directs Cisco ISE to manipulate the identity before it is passed to the external Active Directory system. You can create rules to change the identity to a desired format that includes or excludes a domain prefix and/or suffix or other additional markup of your choice
	FirstName                     string                                                                                      `json:"firstName,omitempty"`                     // User info attribute. All characters are allowed except %
	Department                    string                                                                                      `json:"department,omitempty"`                    // User info attribute. All characters are allowed except %
	LastName                      string                                                                                      `json:"lastName,omitempty"`                      // User info attribute. All characters are allowed except %
	OrganizationalUnit            string                                                                                      `json:"organizationalUnit,omitempty"`            // User info attribute. All characters are allowed except %
	JobTitle                      string                                                                                      `json:"jobTitle,omitempty"`                      // User info attribute. All characters are allowed except %
	Locality                      string                                                                                      `json:"locality,omitempty"`                      // User info attribute. All characters are allowed except %
	Email                         string                                                                                      `json:"email,omitempty"`                         // User info attribute. All characters are allowed except %
	StateOrProvince               string                                                                                      `json:"stateOrProvince,omitempty"`               // User info attribute. All characters are allowed except %
	Telephone                     string                                                                                      `json:"telephone,omitempty"`                     // User info attribute. All characters are allowed except %
	Country                       string                                                                                      `json:"country,omitempty"`                       // User info attribute. All characters are allowed except %
	StreetAddress                 string                                                                                      `json:"streetAddress,omitempty"`                 // User info attribute. All characters are allowed except %
	Schema                        string                                                                                      `json:"schema,omitempty"`                        // Allowed values: ACTIVE_DIRECTORY, CUSTOM. Choose ACTIVE_DIRECTORY schema when the AD attributes defined in AD can be copied to relevant attributes in Cisco ISE. If customization is needed, choose CUSTOM schema. All User info attributes are always set to default value if schema is ACTIVE_DIRECTORY. Values can be changed only for CUSTOM schema
}

type RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdvancedSettingsRewriteRules struct {
	RowID         *int   `json:"rowId,omitempty"`         // Required for each rule in the list in serial order
	RewriteMatch  string `json:"rewriteMatch,omitempty"`  // Required for each rule in the list with no duplication between rules. All characters are allowed except %"
	RewriteResult string `json:"rewriteResult,omitempty"` // Required for each rule in the list. All characters are allowed except %"
}

type RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdAttributes struct {
	Attributes *[]RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdAttributesAttributes `json:"attributes,omitempty"` // List of Attributes
}

type RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdAttributesAttributes struct {
	Name         string `json:"name,omitempty"`         // Required for each attribute in the attribute list with no duplication between attributes. All characters are allowed except <%"
	Type         string `json:"type,omitempty"`         // Required for each group in the group list. Allowed values: STRING, IP, BOOLEAN, INT, OCTET_STRING
	InternalName string `json:"internalName,omitempty"` // Required for each attribute in the attribute list. All characters are allowed except <%"
	DefaultValue string `json:"defaultValue,omitempty"` // Required for each attribute in the attribute list. Can contain an empty string. All characters are allowed except <%"
}

type RequestActiveDirectoryLeaveDomain struct {
	OperationAdditionalData *RequestActiveDirectoryLeaveDomainOperationAdditionalData `json:"OperationAdditionalData,omitempty"` //
}

type RequestActiveDirectoryLeaveDomainOperationAdditionalData struct {
	AdditionalData *[]RequestActiveDirectoryLeaveDomainOperationAdditionalDataAdditionalData `json:"additionalData,omitempty"` //
}

type RequestActiveDirectoryLeaveDomainOperationAdditionalDataAdditionalData struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type RequestActiveDirectoryIsUserMemberOfGroups struct {
	OperationAdditionalData *RequestActiveDirectoryIsUserMemberOfGroupsOperationAdditionalData `json:"OperationAdditionalData,omitempty"` //
}

type RequestActiveDirectoryIsUserMemberOfGroupsOperationAdditionalData struct {
	AdditionalData *[]RequestActiveDirectoryIsUserMemberOfGroupsOperationAdditionalDataAdditionalData `json:"additionalData,omitempty"` //
}

type RequestActiveDirectoryIsUserMemberOfGroupsOperationAdditionalDataAdditionalData struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type RequestActiveDirectoryJoinDomainWithAllNodes struct {
	OperationAdditionalData *RequestActiveDirectoryJoinDomainWithAllNodesOperationAdditionalData `json:"OperationAdditionalData,omitempty"` //
}

type RequestActiveDirectoryJoinDomainWithAllNodesOperationAdditionalData struct {
	AdditionalData *[]RequestActiveDirectoryJoinDomainWithAllNodesOperationAdditionalDataAdditionalData `json:"additionalData,omitempty"` //
}

type RequestActiveDirectoryJoinDomainWithAllNodesOperationAdditionalDataAdditionalData struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type RequestActiveDirectoryLeaveDomainWithAllNodes struct {
	OperationAdditionalData *RequestActiveDirectoryLeaveDomainWithAllNodesOperationAdditionalData `json:"OperationAdditionalData,omitempty"` //
}

type RequestActiveDirectoryLeaveDomainWithAllNodesOperationAdditionalData struct {
	AdditionalData *[]RequestActiveDirectoryLeaveDomainWithAllNodesOperationAdditionalDataAdditionalData `json:"additionalData,omitempty"` //
}

type RequestActiveDirectoryLeaveDomainWithAllNodesOperationAdditionalDataAdditionalData struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type RequestActiveDirectoryGetGroupsByDomain struct {
	OperationAdditionalData *RequestActiveDirectoryGetGroupsByDomainOperationAdditionalData `json:"OperationAdditionalData,omitempty"` //
}

type RequestActiveDirectoryGetGroupsByDomainOperationAdditionalData struct {
	AdditionalData *[]RequestActiveDirectoryGetGroupsByDomainOperationAdditionalDataAdditionalData `json:"additionalData,omitempty"` //
}

type RequestActiveDirectoryGetGroupsByDomainOperationAdditionalDataAdditionalData struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type RequestActiveDirectoryJoinDomain struct {
	OperationAdditionalData *RequestActiveDirectoryJoinDomainOperationAdditionalData `json:"OperationAdditionalData,omitempty"` //
}

type RequestActiveDirectoryJoinDomainOperationAdditionalData struct {
	AdditionalData *[]RequestActiveDirectoryJoinDomainOperationAdditionalDataAdditionalData `json:"additionalData,omitempty"` //
}

type RequestActiveDirectoryJoinDomainOperationAdditionalDataAdditionalData struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

type RequestActiveDirectoryCreateActiveDirectory struct {
	ERSActiveDirectory *RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectory `json:"ERSActiveDirectory,omitempty"` //
}

type RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectory struct {
	ID                    string                                                                         `json:"id,omitempty"`                    // Resource UUID value
	Name                  string                                                                         `json:"name,omitempty"`                  // Resource Name. Maximum 32 characters allowed. Allowed characters are alphanumeric and .-_/\\ characters
	Description           string                                                                         `json:"description,omitempty"`           // No character restriction
	Domain                string                                                                         `json:"domain,omitempty"`                // The AD domain. Alphanumeric, hyphen (-) and dot (.) characters are allowed
	EnableDomainWhiteList *bool                                                                          `json:"enableDomainWhiteList,omitempty"` //
	Adgroups              *RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroups         `json:"adgroups,omitempty"`              // Holds list of AD Groups
	AdvancedSettings      *RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettings `json:"advancedSettings,omitempty"`      //
	AdAttributes          *RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributes     `json:"adAttributes,omitempty"`          // Holds list of AD Attributes
	AdScopesNames         string                                                                         `json:"adScopesNames,omitempty"`         // String that contains the names of the scopes that the active directory belongs to. Names are separated by comma. Alphanumeric, underscore (_) characters are allowed
}

type RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroups struct {
	Groups *[]RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroups `json:"groups,omitempty"` // List of Groups
}

type RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroups struct {
	Name string `json:"name,omitempty"` // Required for each group in the group list with no duplication between groups. All characters are allowed except %
	Sid  string `json:"sid,omitempty"`  // Cisco ISE uses security identifiers (SIDs) for optimization of group membership evaluation. SIDs are useful for efficiency (speed) when the groups are evaluated. All characters are allowed except %
	Type string `json:"type,omitempty"` // No character restriction
}

type RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettings struct {
	EnablePassChange              *bool                                                                                        `json:"enablePassChange,omitempty"`              //
	EnableMachineAuth             *bool                                                                                        `json:"enableMachineAuth,omitempty"`             //
	EnableMachineAccess           *bool                                                                                        `json:"enableMachineAccess,omitempty"`           //
	AgingTime                     *int                                                                                         `json:"agingTime,omitempty"`                     // Range 1-8760 hours
	EnableDialinPermissionCheck   *bool                                                                                        `json:"enableDialinPermissionCheck,omitempty"`   //
	EnableCallbackForDialinClient *bool                                                                                        `json:"enableCallbackForDialinClient,omitempty"` //
	PlaintextAuth                 *bool                                                                                        `json:"plaintextAuth,omitempty"`                 //
	EnableFailedAuthProtection    *bool                                                                                        `json:"enableFailedAuthProtection,omitempty"`    // Enable prevent AD account lockout due to too many bad password attempts
	AuthProtectionType            string                                                                                       `json:"authProtectionType,omitempty"`            // Enable prevent AD account lockout. Allowed values: - WIRELESS, - WIRED, - BOTH
	FailedAuthThreshold           *int                                                                                         `json:"failedAuthThreshold,omitempty"`           // Number of bad password attempts
	IDentityNotInAdBehaviour      string                                                                                       `json:"identityNotInAdBehaviour,omitempty"`      // Allowed values: REJECT, SEARCH_JOINED_FOREST, SEARCH_ALL
	UnreachableDomainsBehaviour   string                                                                                       `json:"unreachableDomainsBehaviour,omitempty"`   // Allowed values: PROCEED, DROP
	EnableRewrites                *bool                                                                                        `json:"enableRewrites,omitempty"`                //
	RewriteRules                  *[]RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRules `json:"rewriteRules,omitempty"`                  // Identity rewrite is an advanced feature that directs Cisco ISE to manipulate the identity before it is passed to the external Active Directory system. You can create rules to change the identity to a desired format that includes or excludes a domain prefix and/or suffix or other additional markup of your choice
	FirstName                     string                                                                                       `json:"firstName,omitempty"`                     // User info attribute. All characters are allowed except %
	Department                    string                                                                                       `json:"department,omitempty"`                    // User info attribute. All characters are allowed except %
	LastName                      string                                                                                       `json:"lastName,omitempty"`                      // User info attribute. All characters are allowed except %
	OrganizationalUnit            string                                                                                       `json:"organizationalUnit,omitempty"`            // User info attribute. All characters are allowed except %
	JobTitle                      string                                                                                       `json:"jobTitle,omitempty"`                      // User info attribute. All characters are allowed except %
	Locality                      string                                                                                       `json:"locality,omitempty"`                      // User info attribute. All characters are allowed except %
	Email                         string                                                                                       `json:"email,omitempty"`                         // User info attribute. All characters are allowed except %
	StateOrProvince               string                                                                                       `json:"stateOrProvince,omitempty"`               // User info attribute. All characters are allowed except %
	Telephone                     string                                                                                       `json:"telephone,omitempty"`                     // User info attribute. All characters are allowed except %
	Country                       string                                                                                       `json:"country,omitempty"`                       // User info attribute. All characters are allowed except %
	StreetAddress                 string                                                                                       `json:"streetAddress,omitempty"`                 // User info attribute. All characters are allowed except %
	Schema                        string                                                                                       `json:"schema,omitempty"`                        // Allowed values: ACTIVE_DIRECTORY, CUSTOM. Choose ACTIVE_DIRECTORY schema when the AD attributes defined in AD can be copied to relevant attributes in Cisco ISE. If customization is needed, choose CUSTOM schema. All User info attributes are always set to default value if schema is ACTIVE_DIRECTORY. Values can be changed only for CUSTOM schema
}

type RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRules struct {
	RowID         *int   `json:"rowId,omitempty"`         // Required for each rule in the list in serial order
	RewriteMatch  string `json:"rewriteMatch,omitempty"`  // Required for each rule in the list with no duplication between rules. All characters are allowed except %"
	RewriteResult string `json:"rewriteResult,omitempty"` // Required for each rule in the list. All characters are allowed except %"
}

type RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributes struct {
	Attributes *[]RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributes `json:"attributes,omitempty"` // List of Attributes
}

type RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributes struct {
	Name         string `json:"name,omitempty"`         // Required for each attribute in the attribute list with no duplication between attributes. All characters are allowed except <%"
	Type         string `json:"type,omitempty"`         // Required for each group in the group list. Allowed values: STRING, IP, BOOLEAN, INT, OCTET_STRING
	InternalName string `json:"internalName,omitempty"` // Required for each attribute in the attribute list. All characters are allowed except <%"
	DefaultValue string `json:"defaultValue,omitempty"` // Required for each attribute in the attribute list. Can contain an empty string. All characters are allowed except <%"
}

//GetActiveDirectoryByName Get Active Directory by name
/* This API allows the client to get Active Directory by name.

@param name name path parameter.
*/
func (s *ActiveDirectoryService) GetActiveDirectoryByName(name string) (*ResponseActiveDirectoryGetActiveDirectoryByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/activedirectory/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseActiveDirectoryGetActiveDirectoryByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetActiveDirectoryByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseActiveDirectoryGetActiveDirectoryByName)
	return result, response, err

}

//GetActiveDirectoryByID Fetch join point details by ID.
/* This API fetchs the join point details by ID. The ID can be retrieved with the Get All operation.

@param id id path parameter.
*/
func (s *ActiveDirectoryService) GetActiveDirectoryByID(id string) (*ResponseActiveDirectoryGetActiveDirectoryByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/activedirectory/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseActiveDirectoryGetActiveDirectoryByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetActiveDirectoryById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseActiveDirectoryGetActiveDirectoryByID)
	return result, response, err

}

//GetActiveDirectory List all join points for Active Directory domains in Cisco ISE
/* This API lists all the join points for Active Directory domains in Cisco ISE.

@param getActiveDirectoryQueryParams Filtering parameter
*/
func (s *ActiveDirectoryService) GetActiveDirectory(getActiveDirectoryQueryParams *GetActiveDirectoryQueryParams) (*ResponseActiveDirectoryGetActiveDirectory, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/activedirectory"

	queryString, _ := query.Values(getActiveDirectoryQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseActiveDirectoryGetActiveDirectory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetActiveDirectory")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseActiveDirectoryGetActiveDirectory)
	return result, response, err

}

//GetVersion Get Active Directory version information
/* This API helps to retrieve the version information related to the active directory.

 */
func (s *ActiveDirectoryService) GetVersion() (*ResponseActiveDirectoryGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/activedirectory/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseActiveDirectoryGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseActiveDirectoryGetVersion)
	return result, response, err

}

//CreateActiveDirectory Create AD join point in Cisco ISE
/* This API creates an AD join point in Cisco ISE.

 */
func (s *ActiveDirectoryService) CreateActiveDirectory(requestActiveDirectoryCreateActiveDirectory *RequestActiveDirectoryCreateActiveDirectory) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/activedirectory"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestActiveDirectoryCreateActiveDirectory).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateActiveDirectory")
	}

	return response, err

}

//GetUserGroups List groups of which given user is member
/* This API allows the client to get groups of which a given user is a member.

@param id id path parameter.
*/
func (s *ActiveDirectoryService) GetUserGroups(id string, requestActiveDirectoryGetUserGroups *RequestActiveDirectoryGetUserGroups) (*ResponseActiveDirectoryGetUserGroups, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/activedirectory/{id}/getUserGroups"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestActiveDirectoryGetUserGroups).
		SetResult(&ResponseActiveDirectoryGetUserGroups{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetUserGroups")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseActiveDirectoryGetUserGroups)
	return result, response, err

}

//LoadGroupsFromDomain Reload domain groups configuration from Active Directory into Cisco ISE
/* This API loads domain groups configuration from Active Directory into Cisco ISE.

@param id id path parameter.
*/
func (s *ActiveDirectoryService) LoadGroupsFromDomain(id string, requestActiveDirectoryLoadGroupsFromDomain *RequestActiveDirectoryLoadGroupsFromDomain) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/activedirectory/{id}/addGroups"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestActiveDirectoryLoadGroupsFromDomain).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation LoadGroupsFromDomain")
	}

	getCSFRToken(response.Header())

	return response, err

}

//LeaveDomain Make Cisco ISE node leave Active Directory domain
/* This API makes a Cisco ISE node to leave an Active Directory domain.

@param id id path parameter.
*/
func (s *ActiveDirectoryService) LeaveDomain(id string, requestActiveDirectoryLeaveDomain *RequestActiveDirectoryLeaveDomain) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/activedirectory/{id}/leave"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestActiveDirectoryLeaveDomain).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation LeaveDomain")
	}

	getCSFRToken(response.Header())

	return response, err

}

//IsUserMemberOfGroups Verify if user is member of given groups
/* This API verifies if the user is a member of the given groups.

@param id id path parameter.
*/
func (s *ActiveDirectoryService) IsUserMemberOfGroups(id string, requestActiveDirectoryIsUserMemberOfGroups *RequestActiveDirectoryIsUserMemberOfGroups) (*ResponseActiveDirectoryIsUserMemberOfGroups, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/activedirectory/{id}/isUserMemberOf"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestActiveDirectoryIsUserMemberOfGroups).
		SetResult(&ResponseActiveDirectoryIsUserMemberOfGroups{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation IsUserMemberOfGroups")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseActiveDirectoryIsUserMemberOfGroups)
	return result, response, err

}

//GetTrustedDomains Get list of domains that are accessible through given join point via trust relationships
/* This API gets the list of domains that are accessible through the given join point via trust relationships.

@param id id path parameter.
*/
func (s *ActiveDirectoryService) GetTrustedDomains(id string) (*ResponseActiveDirectoryGetTrustedDomains, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/activedirectory/{id}/getTrustedDomains"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseActiveDirectoryGetTrustedDomains{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTrustedDomains")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseActiveDirectoryGetTrustedDomains)
	return result, response, err

}

//JoinDomainWithAllNodes Join all Cisco ISE Nodes to Active Directory domain
/* This API joins all Cisco ISE Nodes to an Active Directory domain.

@param id id path parameter.
*/
func (s *ActiveDirectoryService) JoinDomainWithAllNodes(id string, requestActiveDirectoryJoinDomainWithAllNodes *RequestActiveDirectoryJoinDomainWithAllNodes) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/activedirectory/{id}/joinAllNodes"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestActiveDirectoryJoinDomainWithAllNodes).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation JoinDomainWithAllNodes")
	}

	getCSFRToken(response.Header())

	return response, err

}

//LeaveDomainWithAllNodes Make all Cisco ISE nodes leave Active Directory domain
/* This API joins makes all Cisco ISE nodes leave an Active Directory domain.

@param id id path parameter.
*/
func (s *ActiveDirectoryService) LeaveDomainWithAllNodes(id string, requestActiveDirectoryLeaveDomainWithAllNodes *RequestActiveDirectoryLeaveDomainWithAllNodes) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/activedirectory/{id}/leaveAllNodes"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestActiveDirectoryLeaveDomainWithAllNodes).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation LeaveDomainWithAllNodes")
	}

	getCSFRToken(response.Header())

	return response, err

}

//GetGroupsByDomain List the groups of given domain
/* This API lists the groups of the given domain.

@param id id path parameter.
*/
func (s *ActiveDirectoryService) GetGroupsByDomain(id string, requestActiveDirectoryGetGroupsByDomain *RequestActiveDirectoryGetGroupsByDomain) (*ResponseActiveDirectoryGetGroupsByDomain, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/activedirectory/{id}/getGroupsByDomain"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestActiveDirectoryGetGroupsByDomain).
		SetResult(&ResponseActiveDirectoryGetGroupsByDomain{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetGroupsByDomain")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseActiveDirectoryGetGroupsByDomain)
	return result, response, err

}

//JoinDomain Join Cisco ISE node to Active Directory domain
/* This API joins a Cisco ISE node to an Active Directory domain.

@param id id path parameter.
*/
func (s *ActiveDirectoryService) JoinDomain(id string, requestActiveDirectoryJoinDomain *RequestActiveDirectoryJoinDomain) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/activedirectory/{id}/join"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestActiveDirectoryJoinDomain).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation JoinDomain")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteActiveDirectoryByID Delete AD join point from Cisco ISE
/* This API deletes an AD join point from Cisco ISE.

@param id id path parameter.
*/
func (s *ActiveDirectoryService) DeleteActiveDirectoryByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/activedirectory/{id}"
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
		return response, fmt.Errorf("error with operation DeleteActiveDirectoryById")
	}

	getCSFRToken(response.Header())

	return response, err

}
