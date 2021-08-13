package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type AuthorizationProfileService service

type GetAuthorizationProfilesQueryParams struct {
	Page int `url:"page,omitempty"` //Page number
	Size int `url:"size,omitempty"` //Number of objects returned per page
}

type ResponseAuthorizationProfileGetAuthorizationProfileByName struct {
	AuthorizationProfile ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfile `json:"AuthorizationProfile,omitempty"` //
}

type ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfile struct {
	ID                        string                                                                                            `json:"id,omitempty"`                        // Resource UUID value
	Name                      string                                                                                            `json:"name,omitempty"`                      // Resource Name
	Description               string                                                                                            `json:"description,omitempty"`               //
	AdvancedAttributes        []ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileAdvancedAttributes `json:"advancedAttributes,omitempty"`        //
	AccessType                string                                                                                            `json:"accessType,omitempty"`                // Allowed Values: - ACCESS_ACCEPT, - ACCESS_REJECT
	AuthzProfileType          string                                                                                            `json:"authzProfileType,omitempty"`          // Allowed Values: - SWITCH, - TRUSTSEC, - TACACS SWITCH is used for Standard Authorization Profiles
	VLAN                      ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileVLAN                 `json:"vlan,omitempty"`                      //
	Reauth                    ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileReauth               `json:"reauth,omitempty"`                    //
	AirespaceACL              string                                                                                            `json:"airespaceACL,omitempty"`              //
	AirespaceIPv6ACL          string                                                                                            `json:"airespaceIPv6ACL,omitempty"`          //
	WebRedirection            ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileWebRedirection       `json:"webRedirection,omitempty"`            //
	ACL                       string                                                                                            `json:"acl,omitempty"`                       //
	TrackMovement             bool                                                                                              `json:"trackMovement,omitempty"`             //
	AgentlessPosture          bool                                                                                              `json:"agentlessPosture,omitempty"`          //
	ServiceTemplate           bool                                                                                              `json:"serviceTemplate,omitempty"`           //
	EasywiredSessionCandidate bool                                                                                              `json:"easywiredSessionCandidate,omitempty"` //
	DaclName                  string                                                                                            `json:"daclName,omitempty"`                  //
	VoiceDomainPermission     bool                                                                                              `json:"voiceDomainPermission,omitempty"`     //
	Neat                      bool                                                                                              `json:"neat,omitempty"`                      //
	WebAuth                   bool                                                                                              `json:"webAuth,omitempty"`                   //
	AutoSmartPort             string                                                                                            `json:"autoSmartPort,omitempty"`             //
	InterfaceTemplate         string                                                                                            `json:"interfaceTemplate,omitempty"`         //
	IPv6ACLFilter             string                                                                                            `json:"ipv6ACLFilter,omitempty"`             //
	AvcProfile                string                                                                                            `json:"avcProfile,omitempty"`                //
	MacSecPolicy              string                                                                                            `json:"macSecPolicy,omitempty"`              // Allowed Values: - MUST_SECURE, - MUST_NOT_SECURE, - SHOULD_SECURE
	AsaVpn                    string                                                                                            `json:"asaVpn,omitempty"`                    //
	ProfileName               string                                                                                            `json:"profileName,omitempty"`               //
	IPv6DaclName              string                                                                                            `json:"ipv6DaclName,omitempty"`              //
	Link                      ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileLink                 `json:"link,omitempty"`                      //
}

type ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileAdvancedAttributes struct {
	LeftHandSideDictionaryAttribue ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue `json:"leftHandSideDictionaryAttribue,omitempty"` //
	RightHandSideAttribueValue     ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue     `json:"rightHandSideAttribueValue,omitempty"`     // Attribute value can be of type AttributeValue or AdvancedDictionaryAttribute. For AttributeValue the value is String, For AdvancedDictionaryAttribute the value is dictionaryName and attributeName properties
}

type ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue struct {
	AdvancedAttributeValueType string `json:"AdvancedAttributeValueType,omitempty"` //
	DictionaryName             string `json:"dictionaryName,omitempty"`             //
	AttributeName              string `json:"attributeName,omitempty"`              //
	Value                      string `json:"value,omitempty"`                      //
}

type ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue struct {
	AdvancedAttributeValueType string `json:"AdvancedAttributeValueType,omitempty"` //
	DictionaryName             string `json:"dictionaryName,omitempty"`             //
	AttributeName              string `json:"attributeName,omitempty"`              //
	Value                      string `json:"value,omitempty"`                      //
}

type ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileVLAN struct {
	NameID string `json:"nameID,omitempty"` //
	TagID  int    `json:"tagID,omitempty"`  // Valid range is 0-31
}

type ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileReauth struct {
	Timer        int    `json:"timer,omitempty"`        // Valid range is 1-65535
	Connectivity string `json:"connectivity,omitempty"` // Allowed Values: - DEFAULT, - RADIUS_REQUEST
}

type ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileWebRedirection struct {
	WebRedirectionType                 string `json:"WebRedirectionType,omitempty"`                 // Value MUST be one of the following:CentralizedWebAuth, HotSpot, NativeSupplicanProvisioning, ClientProvisioning.  The WebRedirectionType must fit the portalName
	ACL                                string `json:"acl,omitempty"`                                //
	PortalName                         string `json:"portalName,omitempty"`                         // A portal that exist in the DB and fits the WebRedirectionType
	StaticIPHostNameFQDN               string `json:"staticIPHostNameFQDN,omitempty"`               //
	DisplayCertificatesRenewalMessages bool   `json:"displayCertificatesRenewalMessages,omitempty"` // The displayCertificatesRenewalMessages is mandatory when 'WebRedirectionType' value is 'CentralizedWebAuth'. For all other 'WebRedirectionType' values the field must be ignored
}

type ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAuthorizationProfileGetAuthorizationProfileByID struct {
	AuthorizationProfile ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfile `json:"AuthorizationProfile,omitempty"` //
}

type ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfile struct {
	ID                        string                                                                                          `json:"id,omitempty"`                        // Resource UUID value
	Name                      string                                                                                          `json:"name,omitempty"`                      // Resource Name
	Description               string                                                                                          `json:"description,omitempty"`               //
	AdvancedAttributes        []ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileAdvancedAttributes `json:"advancedAttributes,omitempty"`        //
	AccessType                string                                                                                          `json:"accessType,omitempty"`                // Allowed Values: - ACCESS_ACCEPT, - ACCESS_REJECT
	AuthzProfileType          string                                                                                          `json:"authzProfileType,omitempty"`          // Allowed Values: - SWITCH, - TRUSTSEC, - TACACS SWITCH is used for Standard Authorization Profiles
	VLAN                      ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileVLAN                 `json:"vlan,omitempty"`                      //
	Reauth                    ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileReauth               `json:"reauth,omitempty"`                    //
	AirespaceACL              string                                                                                          `json:"airespaceACL,omitempty"`              //
	AirespaceIPv6ACL          string                                                                                          `json:"airespaceIPv6ACL,omitempty"`          //
	WebRedirection            ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileWebRedirection       `json:"webRedirection,omitempty"`            //
	ACL                       string                                                                                          `json:"acl,omitempty"`                       //
	TrackMovement             bool                                                                                            `json:"trackMovement,omitempty"`             //
	AgentlessPosture          bool                                                                                            `json:"agentlessPosture,omitempty"`          //
	ServiceTemplate           bool                                                                                            `json:"serviceTemplate,omitempty"`           //
	EasywiredSessionCandidate bool                                                                                            `json:"easywiredSessionCandidate,omitempty"` //
	DaclName                  string                                                                                          `json:"daclName,omitempty"`                  //
	VoiceDomainPermission     bool                                                                                            `json:"voiceDomainPermission,omitempty"`     //
	Neat                      bool                                                                                            `json:"neat,omitempty"`                      //
	WebAuth                   bool                                                                                            `json:"webAuth,omitempty"`                   //
	AutoSmartPort             string                                                                                          `json:"autoSmartPort,omitempty"`             //
	InterfaceTemplate         string                                                                                          `json:"interfaceTemplate,omitempty"`         //
	IPv6ACLFilter             string                                                                                          `json:"ipv6ACLFilter,omitempty"`             //
	AvcProfile                string                                                                                          `json:"avcProfile,omitempty"`                //
	MacSecPolicy              string                                                                                          `json:"macSecPolicy,omitempty"`              // Allowed Values: - MUST_SECURE, - MUST_NOT_SECURE, - SHOULD_SECURE
	AsaVpn                    string                                                                                          `json:"asaVpn,omitempty"`                    //
	ProfileName               string                                                                                          `json:"profileName,omitempty"`               //
	IPv6DaclName              string                                                                                          `json:"ipv6DaclName,omitempty"`              //
	Link                      ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileLink                 `json:"link,omitempty"`                      //
}

type ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileAdvancedAttributes struct {
	LeftHandSideDictionaryAttribue ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue `json:"leftHandSideDictionaryAttribue,omitempty"` //
	RightHandSideAttribueValue     ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue     `json:"rightHandSideAttribueValue,omitempty"`     // Attribute value can be of type AttributeValue or AdvancedDictionaryAttribute. For AttributeValue the value is String, For AdvancedDictionaryAttribute the value is dictionaryName and attributeName properties
}

type ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue struct {
	AdvancedAttributeValueType string `json:"AdvancedAttributeValueType,omitempty"` //
	DictionaryName             string `json:"dictionaryName,omitempty"`             //
	AttributeName              string `json:"attributeName,omitempty"`              //
	Value                      string `json:"value,omitempty"`                      //
}

type ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue struct {
	AdvancedAttributeValueType string `json:"AdvancedAttributeValueType,omitempty"` //
	DictionaryName             string `json:"dictionaryName,omitempty"`             //
	AttributeName              string `json:"attributeName,omitempty"`              //
	Value                      string `json:"value,omitempty"`                      //
}

type ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileVLAN struct {
	NameID string `json:"nameID,omitempty"` //
	TagID  int    `json:"tagID,omitempty"`  // Valid range is 0-31
}

type ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileReauth struct {
	Timer        int    `json:"timer,omitempty"`        // Valid range is 1-65535
	Connectivity string `json:"connectivity,omitempty"` // Allowed Values: - DEFAULT, - RADIUS_REQUEST
}

type ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileWebRedirection struct {
	WebRedirectionType                 string `json:"WebRedirectionType,omitempty"`                 // Value MUST be one of the following:CentralizedWebAuth, HotSpot, NativeSupplicanProvisioning, ClientProvisioning.  The WebRedirectionType must fit the portalName
	ACL                                string `json:"acl,omitempty"`                                //
	PortalName                         string `json:"portalName,omitempty"`                         // A portal that exist in the DB and fits the WebRedirectionType
	StaticIPHostNameFQDN               string `json:"staticIPHostNameFQDN,omitempty"`               //
	DisplayCertificatesRenewalMessages bool   `json:"displayCertificatesRenewalMessages,omitempty"` // The displayCertificatesRenewalMessages is mandatory when 'WebRedirectionType' value is 'CentralizedWebAuth'. For all other 'WebRedirectionType' values the field must be ignored
}

type ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAuthorizationProfileUpdateAuthorizationProfileByID struct {
	UpdatedFieldsList ResponseAuthorizationProfileUpdateAuthorizationProfileByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseAuthorizationProfileUpdateAuthorizationProfileByIDUpdatedFieldsList struct {
	UpdatedField ResponseAuthorizationProfileUpdateAuthorizationProfileByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                                  `json:"field,omitempty"`        //
	OldValue     string                                                                                  `json:"oldValue,omitempty"`     //
	NewValue     string                                                                                  `json:"newValue,omitempty"`     //
}

type ResponseAuthorizationProfileUpdateAuthorizationProfileByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseAuthorizationProfileGetAuthorizationProfiles struct {
	SearchResult ResponseAuthorizationProfileGetAuthorizationProfilesSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseAuthorizationProfileGetAuthorizationProfilesSearchResult struct {
	Total        int                                                                          `json:"total,omitempty"`        //
	Resources    []ResponseAuthorizationProfileGetAuthorizationProfilesSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponseAuthorizationProfileGetAuthorizationProfilesSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponseAuthorizationProfileGetAuthorizationProfilesSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseAuthorizationProfileGetAuthorizationProfilesSearchResultResources struct {
	ID          string                                                                        `json:"id,omitempty"`          //
	Name        string                                                                        `json:"name,omitempty"`        //
	Description string                                                                        `json:"description,omitempty"` //
	Link        ResponseAuthorizationProfileGetAuthorizationProfilesSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseAuthorizationProfileGetAuthorizationProfilesSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAuthorizationProfileGetAuthorizationProfilesSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAuthorizationProfileGetAuthorizationProfilesSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAuthorizationProfileGetVersion struct {
	VersionInfo ResponseAuthorizationProfileGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseAuthorizationProfileGetVersionVersionInfo struct {
	CurrentServerVersion string                                                `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                                `json:"supportedVersions,omitempty"`    //
	Link                 ResponseAuthorizationProfileGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseAuthorizationProfileGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestAuthorizationProfileUpdateAuthorizationProfileByID struct {
	AuthorizationProfile *RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfile `json:"AuthorizationProfile,omitempty"` //
}

type RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfile struct {
	ID                        string                                                                                             `json:"id,omitempty"`                        // Resource UUID value
	Name                      string                                                                                             `json:"name,omitempty"`                      // Resource Name
	Description               string                                                                                             `json:"description,omitempty"`               //
	AdvancedAttributes        *[]RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributes `json:"advancedAttributes,omitempty"`        //
	AccessType                string                                                                                             `json:"accessType,omitempty"`                // Allowed Values: - ACCESS_ACCEPT, - ACCESS_REJECT
	AuthzProfileType          string                                                                                             `json:"authzProfileType,omitempty"`          // Allowed Values: - SWITCH, - TRUSTSEC, - TACACS SWITCH is used for Standard Authorization Profiles
	VLAN                      *RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileVLAN                 `json:"vlan,omitempty"`                      //
	Reauth                    *RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileReauth               `json:"reauth,omitempty"`                    //
	AirespaceACL              string                                                                                             `json:"airespaceACL,omitempty"`              //
	AirespaceIPv6ACL          string                                                                                             `json:"airespaceIPv6ACL,omitempty"`          //
	WebRedirection            *RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileWebRedirection       `json:"webRedirection,omitempty"`            //
	ACL                       string                                                                                             `json:"acl,omitempty"`                       //
	TrackMovement             *bool                                                                                              `json:"trackMovement,omitempty"`             //
	AgentlessPosture          *bool                                                                                              `json:"agentlessPosture,omitempty"`          //
	ServiceTemplate           *bool                                                                                              `json:"serviceTemplate,omitempty"`           //
	EasywiredSessionCandidate *bool                                                                                              `json:"easywiredSessionCandidate,omitempty"` //
	DaclName                  string                                                                                             `json:"daclName,omitempty"`                  //
	VoiceDomainPermission     *bool                                                                                              `json:"voiceDomainPermission,omitempty"`     //
	Neat                      *bool                                                                                              `json:"neat,omitempty"`                      //
	WebAuth                   *bool                                                                                              `json:"webAuth,omitempty"`                   //
	AutoSmartPort             string                                                                                             `json:"autoSmartPort,omitempty"`             //
	InterfaceTemplate         string                                                                                             `json:"interfaceTemplate,omitempty"`         //
	IPv6ACLFilter             string                                                                                             `json:"ipv6ACLFilter,omitempty"`             //
	AvcProfile                string                                                                                             `json:"avcProfile,omitempty"`                //
	MacSecPolicy              string                                                                                             `json:"macSecPolicy,omitempty"`              // Allowed Values: - MUST_SECURE, - MUST_NOT_SECURE, - SHOULD_SECURE
	AsaVpn                    string                                                                                             `json:"asaVpn,omitempty"`                    //
	ProfileName               string                                                                                             `json:"profileName,omitempty"`               //
	IPv6DaclName              string                                                                                             `json:"ipv6DaclName,omitempty"`              //
}

type RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributes struct {
	LeftHandSideDictionaryAttribue *RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue `json:"leftHandSideDictionaryAttribue,omitempty"` //
	RightHandSideAttribueValue     *RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue     `json:"rightHandSideAttribueValue,omitempty"`     // Attribute value can be of type AttributeValue or AdvancedDictionaryAttribute. For AttributeValue the value is String, For AdvancedDictionaryAttribute the value is dictionaryName and attributeName properties
}

type RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue struct {
	AdvancedAttributeValueType string `json:"AdvancedAttributeValueType,omitempty"` //
	DictionaryName             string `json:"dictionaryName,omitempty"`             //
	AttributeName              string `json:"attributeName,omitempty"`              //
	Value                      string `json:"value,omitempty"`                      //
}

type RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue struct {
	AdvancedAttributeValueType string `json:"AdvancedAttributeValueType,omitempty"` //
	DictionaryName             string `json:"dictionaryName,omitempty"`             //
	AttributeName              string `json:"attributeName,omitempty"`              //
	Value                      string `json:"value,omitempty"`                      //
}

type RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileVLAN struct {
	NameID string `json:"nameID,omitempty"` //
	TagID  *int   `json:"tagID,omitempty"`  // Valid range is 0-31
}

type RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileReauth struct {
	Timer        *int   `json:"timer,omitempty"`        // Valid range is 1-65535
	Connectivity string `json:"connectivity,omitempty"` // Allowed Values: - DEFAULT, - RADIUS_REQUEST
}

type RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileWebRedirection struct {
	WebRedirectionType                 string `json:"WebRedirectionType,omitempty"`                 // Value MUST be one of the following:CentralizedWebAuth, HotSpot, NativeSupplicanProvisioning, ClientProvisioning.  The WebRedirectionType must fit the portalName
	ACL                                string `json:"acl,omitempty"`                                //
	PortalName                         string `json:"portalName,omitempty"`                         // A portal that exist in the DB and fits the WebRedirectionType
	StaticIPHostNameFQDN               string `json:"staticIPHostNameFQDN,omitempty"`               //
	DisplayCertificatesRenewalMessages *bool  `json:"displayCertificatesRenewalMessages,omitempty"` // The displayCertificatesRenewalMessages is mandatory when 'WebRedirectionType' value is 'CentralizedWebAuth'. For all other 'WebRedirectionType' values the field must be ignored
}

type RequestAuthorizationProfileCreateAuthorizationProfile struct {
	AuthorizationProfile *RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfile `json:"AuthorizationProfile,omitempty"` //
}

type RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfile struct {
	ID                        string                                                                                         `json:"id,omitempty"`                        // Resource UUID value
	Name                      string                                                                                         `json:"name,omitempty"`                      // Resource Name
	Description               string                                                                                         `json:"description,omitempty"`               //
	AdvancedAttributes        *[]RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributes `json:"advancedAttributes,omitempty"`        //
	AccessType                string                                                                                         `json:"accessType,omitempty"`                // Allowed Values: - ACCESS_ACCEPT, - ACCESS_REJECT
	AuthzProfileType          string                                                                                         `json:"authzProfileType,omitempty"`          // Allowed Values: - SWITCH, - TRUSTSEC, - TACACS SWITCH is used for Standard Authorization Profiles
	VLAN                      *RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileVLAN                 `json:"vlan,omitempty"`                      //
	Reauth                    *RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileReauth               `json:"reauth,omitempty"`                    //
	AirespaceACL              string                                                                                         `json:"airespaceACL,omitempty"`              //
	AirespaceIPv6ACL          string                                                                                         `json:"airespaceIPv6ACL,omitempty"`          //
	WebRedirection            *RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileWebRedirection       `json:"webRedirection,omitempty"`            //
	ACL                       string                                                                                         `json:"acl,omitempty"`                       //
	TrackMovement             *bool                                                                                          `json:"trackMovement,omitempty"`             //
	AgentlessPosture          *bool                                                                                          `json:"agentlessPosture,omitempty"`          //
	ServiceTemplate           *bool                                                                                          `json:"serviceTemplate,omitempty"`           //
	EasywiredSessionCandidate *bool                                                                                          `json:"easywiredSessionCandidate,omitempty"` //
	DaclName                  string                                                                                         `json:"daclName,omitempty"`                  //
	VoiceDomainPermission     *bool                                                                                          `json:"voiceDomainPermission,omitempty"`     //
	Neat                      *bool                                                                                          `json:"neat,omitempty"`                      //
	WebAuth                   *bool                                                                                          `json:"webAuth,omitempty"`                   //
	AutoSmartPort             string                                                                                         `json:"autoSmartPort,omitempty"`             //
	InterfaceTemplate         string                                                                                         `json:"interfaceTemplate,omitempty"`         //
	IPv6ACLFilter             string                                                                                         `json:"ipv6ACLFilter,omitempty"`             //
	AvcProfile                string                                                                                         `json:"avcProfile,omitempty"`                //
	MacSecPolicy              string                                                                                         `json:"macSecPolicy,omitempty"`              // Allowed Values: - MUST_SECURE, - MUST_NOT_SECURE, - SHOULD_SECURE
	AsaVpn                    string                                                                                         `json:"asaVpn,omitempty"`                    //
	ProfileName               string                                                                                         `json:"profileName,omitempty"`               //
	IPv6DaclName              string                                                                                         `json:"ipv6DaclName,omitempty"`              //
}

type RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributes struct {
	LeftHandSideDictionaryAttribue *RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue `json:"leftHandSideDictionaryAttribue,omitempty"` //
	RightHandSideAttribueValue     *RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue     `json:"rightHandSideAttribueValue,omitempty"`     // Attribute value can be of type AttributeValue or AdvancedDictionaryAttribute. For AttributeValue the value is String, For AdvancedDictionaryAttribute the value is dictionaryName and attributeName properties
}

type RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue struct {
	AdvancedAttributeValueType string `json:"AdvancedAttributeValueType,omitempty"` //
	DictionaryName             string `json:"dictionaryName,omitempty"`             //
	AttributeName              string `json:"attributeName,omitempty"`              //
	Value                      string `json:"value,omitempty"`                      //
}

type RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue struct {
	AdvancedAttributeValueType string `json:"AdvancedAttributeValueType,omitempty"` //
	DictionaryName             string `json:"dictionaryName,omitempty"`             //
	AttributeName              string `json:"attributeName,omitempty"`              //
	Value                      string `json:"value,omitempty"`                      //
}

type RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileVLAN struct {
	NameID string `json:"nameID,omitempty"` //
	TagID  *int   `json:"tagID,omitempty"`  // Valid range is 0-31
}

type RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileReauth struct {
	Timer        *int   `json:"timer,omitempty"`        // Valid range is 1-65535
	Connectivity string `json:"connectivity,omitempty"` // Allowed Values: - DEFAULT, - RADIUS_REQUEST
}

type RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileWebRedirection struct {
	WebRedirectionType                 string `json:"WebRedirectionType,omitempty"`                 // Value MUST be one of the following:CentralizedWebAuth, HotSpot, NativeSupplicanProvisioning, ClientProvisioning.  The WebRedirectionType must fit the portalName
	ACL                                string `json:"acl,omitempty"`                                //
	PortalName                         string `json:"portalName,omitempty"`                         // A portal that exist in the DB and fits the WebRedirectionType
	StaticIPHostNameFQDN               string `json:"staticIPHostNameFQDN,omitempty"`               //
	DisplayCertificatesRenewalMessages *bool  `json:"displayCertificatesRenewalMessages,omitempty"` // The displayCertificatesRenewalMessages is mandatory when 'WebRedirectionType' value is 'CentralizedWebAuth'. For all other 'WebRedirectionType' values the field must be ignored
}

//GetAuthorizationProfileByName Get authorization profile by name
/* This API allows the client to get an authorization profile by name.

@param name name path parameter.
*/
func (s *AuthorizationProfileService) GetAuthorizationProfileByName(name string) (*ResponseAuthorizationProfileGetAuthorizationProfileByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/authorizationprofile/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAuthorizationProfileGetAuthorizationProfileByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAuthorizationProfileByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAuthorizationProfileGetAuthorizationProfileByName)
	return result, response, err

}

//GetAuthorizationProfileByID Get authorization profile by ID
/* This API allows the client to get an authorization profile by ID.

@param id id path parameter.
*/
func (s *AuthorizationProfileService) GetAuthorizationProfileByID(id string) (*ResponseAuthorizationProfileGetAuthorizationProfileByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/authorizationprofile/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAuthorizationProfileGetAuthorizationProfileByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAuthorizationProfileById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAuthorizationProfileGetAuthorizationProfileByID)
	return result, response, err

}

//GetAuthorizationProfiles Get all authorization profiles
/* This API allows the client to get all authorization profiles.

@param getAuthorizationProfilesQueryParams Filtering parameter
*/
func (s *AuthorizationProfileService) GetAuthorizationProfiles(getAuthorizationProfilesQueryParams *GetAuthorizationProfilesQueryParams) (*ResponseAuthorizationProfileGetAuthorizationProfiles, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/authorizationprofile"

	queryString, _ := query.Values(getAuthorizationProfilesQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAuthorizationProfileGetAuthorizationProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAuthorizationProfiles")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAuthorizationProfileGetAuthorizationProfiles)
	return result, response, err

}

//GetVersion Get authorization profile version information
/* This API helps to retrieve the version information related to the authorization profile.

 */
func (s *AuthorizationProfileService) GetVersion() (*ResponseAuthorizationProfileGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/authorizationprofile/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAuthorizationProfileGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAuthorizationProfileGetVersion)
	return result, response, err

}

//CreateAuthorizationProfile Create authorization profiles
/* This API creates an authorization profile.

 */
func (s *AuthorizationProfileService) CreateAuthorizationProfile(requestAuthorizationProfileCreateAuthorizationProfile *RequestAuthorizationProfileCreateAuthorizationProfile) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/authorizationprofile"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAuthorizationProfileCreateAuthorizationProfile).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateAuthorizationProfile")
	}

	return response, err

}

//UpdateAuthorizationProfileByID Update authorization profile
/* This API allows the client to update an authorization profile.

@param id id path parameter.
*/
func (s *AuthorizationProfileService) UpdateAuthorizationProfileByID(id string, requestAuthorizationProfileUpdateAuthorizationProfileById *RequestAuthorizationProfileUpdateAuthorizationProfileByID) (*ResponseAuthorizationProfileUpdateAuthorizationProfileByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/authorizationprofile/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAuthorizationProfileUpdateAuthorizationProfileById).
		SetResult(&ResponseAuthorizationProfileUpdateAuthorizationProfileByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateAuthorizationProfileById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAuthorizationProfileUpdateAuthorizationProfileByID)
	return result, response, err

}

//DeleteAuthorizationProfileByID Delete authorization profile
/* This API deletes an authorization profile.

@param id id path parameter.
*/
func (s *AuthorizationProfileService) DeleteAuthorizationProfileByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/authorizationprofile/{id}"
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
		return response, fmt.Errorf("error with operation DeleteAuthorizationProfileById")
	}

	getCSFRToken(response.Header())

	return response, err

}
