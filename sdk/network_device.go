package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type NetworkDeviceService service

type GetNetworkDeviceQueryParams struct {
	Page       int      `url: page,omitempty`       //Page number
	Size       int      `url: size,omitempty`       //Number of objects returned per page
	Sortasc    string   `url: sortasc,omitempty`    //sort asc
	Sortdsc    string   `url: sortdsc,omitempty`    //sort desc
	Filter     []string `url: filter,omitempty`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url: filterType,omitempty` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseNetworkDeviceGetNetworkDeviceByName struct {
	NetworkDevice ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDevice `json:"NetworkDevice,omitempty"` //
}

type ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDevice struct {
	ID                     string                                                                         `json:"id,omitempty"`                     //
	Name                   string                                                                         `json:"name,omitempty"`                   //
	Description            string                                                                         `json:"description,omitempty"`            //
	AuthenticationSettings ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceAuthenticationSettings `json:"authenticationSettings,omitempty"` //
	SNMPsettings           ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceSNMPsettings           `json:"snmpsettings,omitempty"`           //
	Trustsecsettings       ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceTrustsecsettings       `json:"trustsecsettings,omitempty"`       //
	TacacsSettings         ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceTacacsSettings         `json:"tacacsSettings,omitempty"`         //
	ProfileName            string                                                                         `json:"profileName,omitempty"`            //
	ModelName              string                                                                         `json:"modelName,omitempty"`              //
	SoftwareVersion        string                                                                         `json:"softwareVersion,omitempty"`        //
	CoaPort                int                                                                            `json:"coaPort,omitempty"`                //
	DtlsDNSName            string                                                                         `json:"dtlsDnsName,omitempty"`            // This value is used to verify the client identity contained in the X.509 RADIUS/DTLS client certificate
	NetworkDeviceIPList    []ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceNetworkDeviceIPList  `json:"NetworkDeviceIPList,omitempty"`    // List of IP Subnets for this node
	NetworkDeviceGroupList []string                                                                       `json:"NetworkDeviceGroupList,omitempty"` // List of Network Device Group names for this node
	Link                   ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceLink                   `json:"link,omitempty"`                   //
}

type ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceAuthenticationSettings struct {
	NetworkProtocol             string `json:"networkProtocol,omitempty"`             // Allowed values: - RADIUS, - TACACS_PLUS
	SecondRadiusSharedSecret    string `json:"secondRadiusSharedSecret,omitempty"`    //
	RadiusSharedSecret          string `json:"radiusSharedSecret,omitempty"`          //
	EnableKeyWrap               bool   `json:"enableKeyWrap,omitempty"`               //
	Enabled                     bool   `json:"enabled,omitempty"`                     //
	DtlsRequired                bool   `json:"dtlsRequired,omitempty"`                // This value enforces use of dtls
	EnableMultiSecret           bool   `json:"enableMultiSecret,omitempty"`           //
	KeyEncryptionKey            string `json:"keyEncryptionKey,omitempty"`            //
	MessageAuthenticatorCodeKey string `json:"messageAuthenticatorCodeKey,omitempty"` //
	KeyInputFormat              string `json:"keyInputFormat,omitempty"`              // Allowed values: - ASCII, - HEXADECIMAL
}

type ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceSNMPsettings struct {
	Version                       string `json:"version,omitempty"`                       //
	RoCommunity                   string `json:"roCommunity,omitempty"`                   //
	PollingInterval               int    `json:"pollingInterval,omitempty"`               //
	LinkTrapQuery                 bool   `json:"linkTrapQuery,omitempty"`                 //
	MacTrapQuery                  bool   `json:"macTrapQuery,omitempty"`                  //
	OriginatingPolicyServicesNode string `json:"originatingPolicyServicesNode,omitempty"` //
}

type ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceTrustsecsettings struct {
	DeviceAuthenticationSettings  ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings  `json:"deviceAuthenticationSettings,omitempty"`  //
	SgaNotificationAndUpdates     ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates     `json:"sgaNotificationAndUpdates,omitempty"`     //
	DeviceConfigurationDeployment ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment `json:"deviceConfigurationDeployment,omitempty"` //
	PushIDSupport                 bool                                                                                                  `json:"pushIdSupport,omitempty"`                 //
}

type ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings struct {
	SgaDeviceID       string `json:"sgaDeviceId,omitempty"`       //
	SgaDevicePassword string `json:"sgaDevicePassword,omitempty"` //
}

type ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates struct {
	DownlaodEnvironmentDataEveryXSeconds         int    `json:"downlaodEnvironmentDataEveryXSeconds,omitempty"`         //
	DownlaodPeerAuthorizationPolicyEveryXSeconds int    `json:"downlaodPeerAuthorizationPolicyEveryXSeconds,omitempty"` //
	ReAuthenticationEveryXSeconds                int    `json:"reAuthenticationEveryXSeconds,omitempty"`                //
	DownloadSgACLListsEveryXSeconds              int    `json:"downloadSGACLListsEveryXSeconds,omitempty"`              //
	OtherSgADevicesToTrustThisDevice             bool   `json:"otherSGADevicesToTrustThisDevice,omitempty"`             //
	SendConfigurationToDevice                    bool   `json:"sendConfigurationToDevice,omitempty"`                    //
	SendConfigurationToDeviceUsing               string `json:"sendConfigurationToDeviceUsing,omitempty"`               // Allowed values: - ENABLE_USING_COA, - ENABLE_USING_CLI, - DISABLE_ALL
	CoaSourceHost                                string `json:"coaSourceHost,omitempty"`                                //
}

type ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment struct {
	IncludeWhenDeployingSgtUpdates bool   `json:"includeWhenDeployingSGTUpdates,omitempty"` //
	EnableModePassword             string `json:"enableModePassword,omitempty"`             //
	ExecModePassword               string `json:"execModePassword,omitempty"`               //
	ExecModeUsername               string `json:"execModeUsername,omitempty"`               //
}

type ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceTacacsSettings struct {
	SharedSecret       string `json:"sharedSecret,omitempty"`       //
	ConnectModeOptions string `json:"connectModeOptions,omitempty"` // Allowed values: - OFF, - ON_LEGACY, - ON_DRAFT_COMPLIANT
}

type ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceNetworkDeviceIPList struct {
	IPaddress           string `json:"ipaddress,omitempty"`           //
	Mask                int    `json:"mask,omitempty"`                //
	GetIPaddressExclude string `json:"getIpaddressExclude,omitempty"` // It can be either single IP address or IP range address
}

type ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkDeviceUpdateNetworkDeviceByName struct {
	UpdatedFieldsList ResponseNetworkDeviceUpdateNetworkDeviceByNameUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseNetworkDeviceUpdateNetworkDeviceByNameUpdatedFieldsList struct {
	UpdatedField ResponseNetworkDeviceUpdateNetworkDeviceByNameUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                      `json:"field,omitempty"`        //
	OldValue     string                                                                      `json:"oldValue,omitempty"`     //
	NewValue     string                                                                      `json:"newValue,omitempty"`     //
}

type ResponseNetworkDeviceUpdateNetworkDeviceByNameUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseNetworkDeviceGetNetworkDeviceByID struct {
	NetworkDevice ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDevice `json:"NetworkDevice,omitempty"` //
}

type ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDevice struct {
	ID                     string                                                                       `json:"id,omitempty"`                     //
	Name                   string                                                                       `json:"name,omitempty"`                   //
	Description            string                                                                       `json:"description,omitempty"`            //
	AuthenticationSettings ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceAuthenticationSettings `json:"authenticationSettings,omitempty"` //
	SNMPsettings           ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceSNMPsettings           `json:"snmpsettings,omitempty"`           //
	Trustsecsettings       ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceTrustsecsettings       `json:"trustsecsettings,omitempty"`       //
	TacacsSettings         ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceTacacsSettings         `json:"tacacsSettings,omitempty"`         //
	ProfileName            string                                                                       `json:"profileName,omitempty"`            //
	CoaPort                int                                                                          `json:"coaPort,omitempty"`                //
	DtlsDNSName            string                                                                       `json:"dtlsDnsName,omitempty"`            // This value is used to verify the client identity contained in the X.509 RADIUS/DTLS client certificate
	ModelName              string                                                                       `json:"modelName,omitempty"`              //
	SoftwareVersion        string                                                                       `json:"softwareVersion,omitempty"`        //
	NetworkDeviceIPList    []ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceNetworkDeviceIPList  `json:"NetworkDeviceIPList,omitempty"`    // List of IP Subnets for this node
	NetworkDeviceGroupList []string                                                                     `json:"NetworkDeviceGroupList,omitempty"` // List of Network Device Group names for this node
	Link                   ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceLink                   `json:"link,omitempty"`                   //
}

type ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceAuthenticationSettings struct {
	NetworkProtocol             string `json:"networkProtocol,omitempty"`             // Allowed values: - RADIUS, - TACACS_PLUS
	SecondRadiusSharedSecret    string `json:"secondRadiusSharedSecret,omitempty"`    //
	RadiusSharedSecret          string `json:"radiusSharedSecret,omitempty"`          //
	EnableKeyWrap               bool   `json:"enableKeyWrap,omitempty"`               //
	Enabled                     bool   `json:"enabled,omitempty"`                     //
	DtlsRequired                bool   `json:"dtlsRequired,omitempty"`                // This value enforces use of dtls
	EnableMultiSecret           bool   `json:"enableMultiSecret,omitempty"`           //
	KeyEncryptionKey            string `json:"keyEncryptionKey,omitempty"`            //
	MessageAuthenticatorCodeKey string `json:"messageAuthenticatorCodeKey,omitempty"` //
	KeyInputFormat              string `json:"keyInputFormat,omitempty"`              // Allowed values: - ASCII, - HEXADECIMAL
}

type ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceSNMPsettings struct {
	Version                       string `json:"version,omitempty"`                       //
	RoCommunity                   string `json:"roCommunity,omitempty"`                   //
	PollingInterval               int    `json:"pollingInterval,omitempty"`               //
	LinkTrapQuery                 bool   `json:"linkTrapQuery,omitempty"`                 //
	MacTrapQuery                  bool   `json:"macTrapQuery,omitempty"`                  //
	OriginatingPolicyServicesNode string `json:"originatingPolicyServicesNode,omitempty"` //
}

type ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceTrustsecsettings struct {
	DeviceAuthenticationSettings  ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings  `json:"deviceAuthenticationSettings,omitempty"`  //
	SgaNotificationAndUpdates     ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates     `json:"sgaNotificationAndUpdates,omitempty"`     //
	DeviceConfigurationDeployment ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment `json:"deviceConfigurationDeployment,omitempty"` //
	PushIDSupport                 bool                                                                                                `json:"pushIdSupport,omitempty"`                 //
}

type ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings struct {
	SgaDeviceID       string `json:"sgaDeviceId,omitempty"`       //
	SgaDevicePassword string `json:"sgaDevicePassword,omitempty"` //
}

type ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates struct {
	DownlaodEnvironmentDataEveryXSeconds         int    `json:"downlaodEnvironmentDataEveryXSeconds,omitempty"`         //
	DownlaodPeerAuthorizationPolicyEveryXSeconds int    `json:"downlaodPeerAuthorizationPolicyEveryXSeconds,omitempty"` //
	ReAuthenticationEveryXSeconds                int    `json:"reAuthenticationEveryXSeconds,omitempty"`                //
	DownloadSgACLListsEveryXSeconds              int    `json:"downloadSGACLListsEveryXSeconds,omitempty"`              //
	OtherSgADevicesToTrustThisDevice             bool   `json:"otherSGADevicesToTrustThisDevice,omitempty"`             //
	SendConfigurationToDevice                    bool   `json:"sendConfigurationToDevice,omitempty"`                    //
	SendConfigurationToDeviceUsing               string `json:"sendConfigurationToDeviceUsing,omitempty"`               // Allowed values: - ENABLE_USING_COA, - ENABLE_USING_CLI, - DISABLE_ALL
	CoaSourceHost                                string `json:"coaSourceHost,omitempty"`                                //
}

type ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment struct {
	IncludeWhenDeployingSgtUpdates bool   `json:"includeWhenDeployingSGTUpdates,omitempty"` //
	EnableModePassword             string `json:"enableModePassword,omitempty"`             //
	ExecModePassword               string `json:"execModePassword,omitempty"`               //
	ExecModeUsername               string `json:"execModeUsername,omitempty"`               //
}

type ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceTacacsSettings struct {
	SharedSecret       string `json:"sharedSecret,omitempty"`       //
	ConnectModeOptions string `json:"connectModeOptions,omitempty"` // Allowed values: - OFF, - ON_LEGACY, - ON_DRAFT_COMPLIANT
}

type ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceNetworkDeviceIPList struct {
	IPaddress           string `json:"ipaddress,omitempty"`           //
	Mask                int    `json:"mask,omitempty"`                //
	GetIPaddressExclude string `json:"getIpaddressExclude,omitempty"` // It can be either single IP address or IP range address
}

type ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkDeviceUpdateNetworkDeviceByID struct {
	UpdatedFieldsList ResponseNetworkDeviceUpdateNetworkDeviceByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseNetworkDeviceUpdateNetworkDeviceByIDUpdatedFieldsList struct {
	UpdatedField ResponseNetworkDeviceUpdateNetworkDeviceByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                    `json:"field,omitempty"`        //
	OldValue     string                                                                    `json:"oldValue,omitempty"`     //
	NewValue     string                                                                    `json:"newValue,omitempty"`     //
}

type ResponseNetworkDeviceUpdateNetworkDeviceByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseNetworkDeviceGetNetworkDevice struct {
	SearchResult ResponseNetworkDeviceGetNetworkDeviceSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseNetworkDeviceGetNetworkDeviceSearchResult struct {
	Total        int                                                             `json:"total,omitempty"`        //
	Resources    []ResponseNetworkDeviceGetNetworkDeviceSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseNetworkDeviceGetNetworkDeviceSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseNetworkDeviceGetNetworkDeviceSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseNetworkDeviceGetNetworkDeviceSearchResultResources struct {
	ID          string                                                         `json:"id,omitempty"`          //
	Name        string                                                         `json:"name,omitempty"`        //
	Description string                                                         `json:"description,omitempty"` //
	Link        ResponseNetworkDeviceGetNetworkDeviceSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseNetworkDeviceGetNetworkDeviceSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkDeviceGetNetworkDeviceSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkDeviceGetNetworkDeviceSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkDeviceGetVersion struct {
	VersionInfo ResponseNetworkDeviceGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseNetworkDeviceGetVersionVersionInfo struct {
	CurrentServerVersion string                                         `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                         `json:"supportedVersions,omitempty"`    //
	Link                 ResponseNetworkDeviceGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseNetworkDeviceGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkDeviceMonitorBulkStatusNetworkDevice struct {
	BulkStatus ResponseNetworkDeviceMonitorBulkStatusNetworkDeviceBulkStatus `json:"BulkStatus,omitempty"` //
}

type ResponseNetworkDeviceMonitorBulkStatusNetworkDeviceBulkStatus struct {
	BulkID          string                                                                         `json:"bulkId,omitempty"`          //
	MediaType       string                                                                         `json:"mediaType,omitempty"`       //
	ExecutionStatus string                                                                         `json:"executionStatus,omitempty"` //
	OperationType   string                                                                         `json:"operationType,omitempty"`   //
	StartTime       string                                                                         `json:"startTime,omitempty"`       //
	ResourcesCount  int                                                                            `json:"resourcesCount,omitempty"`  //
	SuccessCount    int                                                                            `json:"successCount,omitempty"`    //
	FailCount       int                                                                            `json:"failCount,omitempty"`       //
	ResourcesStatus []ResponseNetworkDeviceMonitorBulkStatusNetworkDeviceBulkStatusResourcesStatus `json:"resourcesStatus,omitempty"` //
}

type ResponseNetworkDeviceMonitorBulkStatusNetworkDeviceBulkStatusResourcesStatus struct {
	ID                      string `json:"id,omitempty"`                      //
	Name                    string `json:"name,omitempty"`                    //
	Description             string `json:"description,omitempty"`             //
	ResourceExecutionStatus string `json:"resourceExecutionStatus,omitempty"` //
	Status                  string `json:"status,omitempty"`                  //
}

type RequestNetworkDeviceUpdateNetworkDeviceByName struct {
	NetworkDevice *RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDevice `json:"NetworkDevice,omitempty"` //
}

type RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDevice struct {
	ID                     string                                                                            `json:"id,omitempty"`                     //
	Name                   string                                                                            `json:"name,omitempty"`                   //
	Description            string                                                                            `json:"description,omitempty"`            //
	AuthenticationSettings *RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDeviceAuthenticationSettings `json:"authenticationSettings,omitempty"` //
	SNMPsettings           *RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDeviceSNMPsettings           `json:"snmpsettings,omitempty"`           //
	Trustsecsettings       *RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDeviceTrustsecsettings       `json:"trustsecsettings,omitempty"`       //
	TacacsSettings         *RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDeviceTacacsSettings         `json:"tacacsSettings,omitempty"`         //
	ProfileName            string                                                                            `json:"profileName,omitempty"`            //
	CoaPort                *int                                                                              `json:"coaPort,omitempty"`                //
	DtlsDNSName            string                                                                            `json:"dtlsDnsName,omitempty"`            // This value is used to verify the client identity contained in the X.509 RADIUS/DTLS client certificate
	ModelName              string                                                                            `json:"modelName,omitempty"`              //
	SoftwareVersion        string                                                                            `json:"softwareVersion,omitempty"`        //
	NetworkDeviceIPList    *[]RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDeviceNetworkDeviceIPList  `json:"NetworkDeviceIPList,omitempty"`    // List of IP Subnets for this node
	NetworkDeviceGroupList []string                                                                          `json:"NetworkDeviceGroupList,omitempty"` // List of Network Device Group names for this node
}

type RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDeviceAuthenticationSettings struct {
	NetworkProtocol             string `json:"networkProtocol,omitempty"`             // Allowed values: - RADIUS, - TACACS_PLUS
	SecondRadiusSharedSecret    string `json:"secondRadiusSharedSecret,omitempty"`    //
	RadiusSharedSecret          string `json:"radiusSharedSecret,omitempty"`          //
	EnableKeyWrap               *bool  `json:"enableKeyWrap,omitempty"`               //
	Enabled                     *bool  `json:"enabled,omitempty"`                     //
	DtlsRequired                *bool  `json:"dtlsRequired,omitempty"`                // This value enforces use of dtls
	EnableMultiSecret           *bool  `json:"enableMultiSecret,omitempty"`           //
	KeyEncryptionKey            string `json:"keyEncryptionKey,omitempty"`            //
	MessageAuthenticatorCodeKey string `json:"messageAuthenticatorCodeKey,omitempty"` //
	KeyInputFormat              string `json:"keyInputFormat,omitempty"`              // Allowed values: - ASCII, - HEXADECIMAL
}

type RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDeviceSNMPsettings struct {
	Version                       string `json:"version,omitempty"`                       //
	RoCommunity                   string `json:"roCommunity,omitempty"`                   //
	PollingInterval               *int   `json:"pollingInterval,omitempty"`               //
	LinkTrapQuery                 *bool  `json:"linkTrapQuery,omitempty"`                 //
	MacTrapQuery                  *bool  `json:"macTrapQuery,omitempty"`                  //
	OriginatingPolicyServicesNode string `json:"originatingPolicyServicesNode,omitempty"` //
}

type RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDeviceTrustsecsettings struct {
	DeviceAuthenticationSettings  *RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings  `json:"deviceAuthenticationSettings,omitempty"`  //
	SgaNotificationAndUpdates     *RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates     `json:"sgaNotificationAndUpdates,omitempty"`     //
	DeviceConfigurationDeployment *RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment `json:"deviceConfigurationDeployment,omitempty"` //
	PushIDSupport                 *bool                                                                                                    `json:"pushIdSupport,omitempty"`                 //
}

type RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings struct {
	SgaDeviceID       string `json:"sgaDeviceId,omitempty"`       //
	SgaDevicePassword string `json:"sgaDevicePassword,omitempty"` //
}

type RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates struct {
	DownlaodEnvironmentDataEveryXSeconds         *int   `json:"downlaodEnvironmentDataEveryXSeconds,omitempty"`         //
	DownlaodPeerAuthorizationPolicyEveryXSeconds *int   `json:"downlaodPeerAuthorizationPolicyEveryXSeconds,omitempty"` //
	ReAuthenticationEveryXSeconds                *int   `json:"reAuthenticationEveryXSeconds,omitempty"`                //
	DownloadSgACLListsEveryXSeconds              *int   `json:"downloadSGACLListsEveryXSeconds,omitempty"`              //
	OtherSgADevicesToTrustThisDevice             *bool  `json:"otherSGADevicesToTrustThisDevice,omitempty"`             //
	SendConfigurationToDevice                    *bool  `json:"sendConfigurationToDevice,omitempty"`                    //
	SendConfigurationToDeviceUsing               string `json:"sendConfigurationToDeviceUsing,omitempty"`               // Allowed values: - ENABLE_USING_COA, - ENABLE_USING_CLI, - DISABLE_ALL
	CoaSourceHost                                string `json:"coaSourceHost,omitempty"`                                //
}

type RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment struct {
	IncludeWhenDeployingSgtUpdates *bool  `json:"includeWhenDeployingSGTUpdates,omitempty"` //
	EnableModePassword             string `json:"enableModePassword,omitempty"`             //
	ExecModePassword               string `json:"execModePassword,omitempty"`               //
	ExecModeUsername               string `json:"execModeUsername,omitempty"`               //
}

type RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDeviceTacacsSettings struct {
	SharedSecret       string `json:"sharedSecret,omitempty"`       //
	ConnectModeOptions string `json:"connectModeOptions,omitempty"` // Allowed values: - OFF, - ON_LEGACY, - ON_DRAFT_COMPLIANT
}

type RequestNetworkDeviceUpdateNetworkDeviceByNameNetworkDeviceNetworkDeviceIPList struct {
	IPaddress           string `json:"ipaddress,omitempty"`           //
	Mask                *int   `json:"mask,omitempty"`                //
	GetIPaddressExclude string `json:"getIpaddressExclude,omitempty"` // It can be either single IP address or IP range address
}

type RequestNetworkDeviceUpdateNetworkDeviceByID struct {
	NetworkDevice *RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDevice `json:"NetworkDevice,omitempty"` //
}

type RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDevice struct {
	ID                     string                                                                          `json:"id,omitempty"`                     //
	Name                   string                                                                          `json:"name,omitempty"`                   //
	Description            string                                                                          `json:"description,omitempty"`            //
	AuthenticationSettings *RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceAuthenticationSettings `json:"authenticationSettings,omitempty"` //
	SNMPsettings           *RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceSNMPsettings           `json:"snmpsettings,omitempty"`           //
	Trustsecsettings       *RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettings       `json:"trustsecsettings,omitempty"`       //
	TacacsSettings         *RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTacacsSettings         `json:"tacacsSettings,omitempty"`         //
	ProfileName            string                                                                          `json:"profileName,omitempty"`            //
	CoaPort                *int                                                                            `json:"coaPort,omitempty"`                //
	DtlsDNSName            string                                                                          `json:"dtlsDnsName,omitempty"`            // This value is used to verify the client identity contained in the X.509 RADIUS/DTLS client certificate
	ModelName              string                                                                          `json:"modelName,omitempty"`              //
	SoftwareVersion        string                                                                          `json:"softwareVersion,omitempty"`        //
	NetworkDeviceIPList    *[]RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceNetworkDeviceIPList  `json:"NetworkDeviceIPList,omitempty"`    // List of IP Subnets for this node
	NetworkDeviceGroupList []string                                                                        `json:"NetworkDeviceGroupList,omitempty"` // List of Network Device Group names for this node
}

type RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceAuthenticationSettings struct {
	NetworkProtocol             string `json:"networkProtocol,omitempty"`             // Allowed values: - RADIUS, - TACACS_PLUS
	SecondRadiusSharedSecret    string `json:"secondRadiusSharedSecret,omitempty"`    //
	RadiusSharedSecret          string `json:"radiusSharedSecret,omitempty"`          //
	EnableKeyWrap               *bool  `json:"enableKeyWrap,omitempty"`               //
	Enabled                     *bool  `json:"enabled,omitempty"`                     //
	DtlsRequired                *bool  `json:"dtlsRequired,omitempty"`                // This value enforces use of dtls
	EnableMultiSecret           *bool  `json:"enableMultiSecret,omitempty"`           //
	KeyEncryptionKey            string `json:"keyEncryptionKey,omitempty"`            //
	MessageAuthenticatorCodeKey string `json:"messageAuthenticatorCodeKey,omitempty"` //
	KeyInputFormat              string `json:"keyInputFormat,omitempty"`              // Allowed values: - ASCII, - HEXADECIMAL
}

type RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceSNMPsettings struct {
	Version                       string `json:"version,omitempty"`                       //
	RoCommunity                   string `json:"roCommunity,omitempty"`                   //
	PollingInterval               *int   `json:"pollingInterval,omitempty"`               //
	LinkTrapQuery                 *bool  `json:"linkTrapQuery,omitempty"`                 //
	MacTrapQuery                  *bool  `json:"macTrapQuery,omitempty"`                  //
	OriginatingPolicyServicesNode string `json:"originatingPolicyServicesNode,omitempty"` //
}

type RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettings struct {
	DeviceAuthenticationSettings  *RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings  `json:"deviceAuthenticationSettings,omitempty"`  //
	SgaNotificationAndUpdates     *RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates     `json:"sgaNotificationAndUpdates,omitempty"`     //
	DeviceConfigurationDeployment *RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment `json:"deviceConfigurationDeployment,omitempty"` //
	PushIDSupport                 *bool                                                                                                  `json:"pushIdSupport,omitempty"`                 //
}

type RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings struct {
	SgaDeviceID       string `json:"sgaDeviceId,omitempty"`       //
	SgaDevicePassword string `json:"sgaDevicePassword,omitempty"` //
}

type RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates struct {
	DownlaodEnvironmentDataEveryXSeconds         *int   `json:"downlaodEnvironmentDataEveryXSeconds,omitempty"`         //
	DownlaodPeerAuthorizationPolicyEveryXSeconds *int   `json:"downlaodPeerAuthorizationPolicyEveryXSeconds,omitempty"` //
	ReAuthenticationEveryXSeconds                *int   `json:"reAuthenticationEveryXSeconds,omitempty"`                //
	DownloadSgACLListsEveryXSeconds              *int   `json:"downloadSGACLListsEveryXSeconds,omitempty"`              //
	OtherSgADevicesToTrustThisDevice             *bool  `json:"otherSGADevicesToTrustThisDevice,omitempty"`             //
	SendConfigurationToDevice                    *bool  `json:"sendConfigurationToDevice,omitempty"`                    //
	SendConfigurationToDeviceUsing               string `json:"sendConfigurationToDeviceUsing,omitempty"`               // Allowed values: - ENABLE_USING_COA, - ENABLE_USING_CLI, - DISABLE_ALL
	CoaSourceHost                                string `json:"coaSourceHost,omitempty"`                                //
}

type RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment struct {
	IncludeWhenDeployingSgtUpdates *bool  `json:"includeWhenDeployingSGTUpdates,omitempty"` //
	EnableModePassword             string `json:"enableModePassword,omitempty"`             //
	ExecModePassword               string `json:"execModePassword,omitempty"`               //
	ExecModeUsername               string `json:"execModeUsername,omitempty"`               //
}

type RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTacacsSettings struct {
	SharedSecret       string `json:"sharedSecret,omitempty"`       //
	ConnectModeOptions string `json:"connectModeOptions,omitempty"` // Allowed values: - OFF, - ON_LEGACY, - ON_DRAFT_COMPLIANT
}

type RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceNetworkDeviceIPList struct {
	IPaddress           string `json:"ipaddress,omitempty"`           //
	Mask                *int   `json:"mask,omitempty"`                //
	GetIPaddressExclude string `json:"getIpaddressExclude,omitempty"` // It can be either single IP address or IP range address
}

type RequestNetworkDeviceCreateNetworkDevice struct {
	NetworkDevice *RequestNetworkDeviceCreateNetworkDeviceNetworkDevice `json:"NetworkDevice,omitempty"` //
}

type RequestNetworkDeviceCreateNetworkDeviceNetworkDevice struct {
	Name                   string                                                                      `json:"name,omitempty"`                   //
	Description            string                                                                      `json:"description,omitempty"`            //
	AuthenticationSettings *RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceAuthenticationSettings `json:"authenticationSettings,omitempty"` //
	SNMPsettings           *RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceSNMPsettings           `json:"snmpsettings,omitempty"`           //
	Trustsecsettings       *RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettings       `json:"trustsecsettings,omitempty"`       //
	TacacsSettings         *RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTacacsSettings         `json:"tacacsSettings,omitempty"`         //
	ProfileName            string                                                                      `json:"profileName,omitempty"`            //
	CoaPort                *int                                                                        `json:"coaPort,omitempty"`                //
	DtlsDNSName            string                                                                      `json:"dtlsDnsName,omitempty"`            // This value is used to verify the client identity contained in the X.509 RADIUS/DTLS client certificate
	ModelName              string                                                                      `json:"modelName,omitempty"`              //
	SoftwareVersion        string                                                                      `json:"softwareVersion,omitempty"`        //
	NetworkDeviceIPList    *[]RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceNetworkDeviceIPList  `json:"NetworkDeviceIPList,omitempty"`    // List of IP Subnets for this node
	NetworkDeviceGroupList []string                                                                    `json:"NetworkDeviceGroupList,omitempty"` // List of Network Device Group names for this node
}

type RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceAuthenticationSettings struct {
	NetworkProtocol             string `json:"networkProtocol,omitempty"`             // Allowed values: - RADIUS, - TACACS_PLUS
	SecondRadiusSharedSecret    string `json:"secondRadiusSharedSecret,omitempty"`    //
	RadiusSharedSecret          string `json:"radiusSharedSecret,omitempty"`          //
	EnableKeyWrap               *bool  `json:"enableKeyWrap,omitempty"`               //
	Enabled                     *bool  `json:"enabled,omitempty"`                     //
	DtlsRequired                *bool  `json:"dtlsRequired,omitempty"`                // This value enforces use of dtls
	EnableMultiSecret           *bool  `json:"enableMultiSecret,omitempty"`           //
	KeyEncryptionKey            string `json:"keyEncryptionKey,omitempty"`            //
	MessageAuthenticatorCodeKey string `json:"messageAuthenticatorCodeKey,omitempty"` //
	KeyInputFormat              string `json:"keyInputFormat,omitempty"`              // Allowed values: - ASCII, - HEXADECIMAL
}

type RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceSNMPsettings struct {
	Version                       string `json:"version,omitempty"`                       //
	RoCommunity                   string `json:"roCommunity,omitempty"`                   //
	PollingInterval               *int   `json:"pollingInterval,omitempty"`               //
	LinkTrapQuery                 *bool  `json:"linkTrapQuery,omitempty"`                 //
	MacTrapQuery                  *bool  `json:"macTrapQuery,omitempty"`                  //
	OriginatingPolicyServicesNode string `json:"originatingPolicyServicesNode,omitempty"` //
}

type RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettings struct {
	DeviceAuthenticationSettings  *RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings  `json:"deviceAuthenticationSettings,omitempty"`  //
	SgaNotificationAndUpdates     *RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates     `json:"sgaNotificationAndUpdates,omitempty"`     //
	DeviceConfigurationDeployment *RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment `json:"deviceConfigurationDeployment,omitempty"` //
	PushIDSupport                 *bool                                                                                              `json:"pushIdSupport,omitempty"`                 //
}

type RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings struct {
	SgaDeviceID       string `json:"sgaDeviceId,omitempty"`       //
	SgaDevicePassword string `json:"sgaDevicePassword,omitempty"` //
}

type RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates struct {
	DownlaodEnvironmentDataEveryXSeconds         *int   `json:"downlaodEnvironmentDataEveryXSeconds,omitempty"`         //
	DownlaodPeerAuthorizationPolicyEveryXSeconds *int   `json:"downlaodPeerAuthorizationPolicyEveryXSeconds,omitempty"` //
	ReAuthenticationEveryXSeconds                *int   `json:"reAuthenticationEveryXSeconds,omitempty"`                //
	DownloadSgACLListsEveryXSeconds              *int   `json:"downloadSGACLListsEveryXSeconds,omitempty"`              //
	OtherSgADevicesToTrustThisDevice             *bool  `json:"otherSGADevicesToTrustThisDevice,omitempty"`             //
	SendConfigurationToDevice                    *bool  `json:"sendConfigurationToDevice,omitempty"`                    //
	SendConfigurationToDeviceUsing               string `json:"sendConfigurationToDeviceUsing,omitempty"`               // Allowed values: - ENABLE_USING_COA, - ENABLE_USING_CLI, - DISABLE_ALL
	CoaSourceHost                                string `json:"coaSourceHost,omitempty"`                                //
}

type RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment struct {
	IncludeWhenDeployingSgtUpdates *bool  `json:"includeWhenDeployingSGTUpdates,omitempty"` //
	EnableModePassword             string `json:"enableModePassword,omitempty"`             //
	ExecModePassword               string `json:"execModePassword,omitempty"`               //
	ExecModeUsername               string `json:"execModeUsername,omitempty"`               //
}

type RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTacacsSettings struct {
	SharedSecret       string `json:"sharedSecret,omitempty"`       //
	ConnectModeOptions string `json:"connectModeOptions,omitempty"` // Allowed values: - OFF, - ON_LEGACY, - ON_DRAFT_COMPLIANT
}

type RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceNetworkDeviceIPList struct {
	IPaddress           string `json:"ipaddress,omitempty"`           //
	Mask                *int   `json:"mask,omitempty"`                //
	GetIPaddressExclude string `json:"getIpaddressExclude,omitempty"` // It can be either single IP address or IP range address
}

type RequestNetworkDeviceBulkRequestForNetworkDevice struct {
	NetworkDeviceBulkRequest *RequestNetworkDeviceBulkRequestForNetworkDeviceNetworkDeviceBulkRequest `json:"NetworkDeviceBulkRequest,omitempty"` //
}

type RequestNetworkDeviceBulkRequestForNetworkDeviceNetworkDeviceBulkRequest struct {
	OperationType     string `json:"operationType,omitempty"`     //
	ResourceMediaType string `json:"resourceMediaType,omitempty"` //
}

//GetNetworkDeviceByName Get network device by name
/* This API allows the client to get a network device by name.

@param name name path parameter.
*/
func (s *NetworkDeviceService) GetNetworkDeviceByName(name string) (*ResponseNetworkDeviceGetNetworkDeviceByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevice/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkDeviceGetNetworkDeviceByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkDeviceByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkDeviceGetNetworkDeviceByName)
	return result, response, err

}

//GetNetworkDeviceByID Get network device by ID
/* This API allows the client to get a network device by ID.

@param id id path parameter.
*/
func (s *NetworkDeviceService) GetNetworkDeviceByID(id string) (*ResponseNetworkDeviceGetNetworkDeviceByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevice/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkDeviceGetNetworkDeviceByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkDeviceById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkDeviceGetNetworkDeviceByID)
	return result, response, err

}

//GetNetworkDevice Get all network devices
/* This API allows the client to get all the network devices.

Filter:

[ipaddress, name, description, location, type]

To search resources by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:

[name, description]

@param getNetworkDeviceQueryParams Filtering parameter
*/
func (s *NetworkDeviceService) GetNetworkDevice(getNetworkDeviceQueryParams *GetNetworkDeviceQueryParams) (*ResponseNetworkDeviceGetNetworkDevice, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevice"

	queryString, _ := query.Values(getNetworkDeviceQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkDeviceGetNetworkDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkDevice")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkDeviceGetNetworkDevice)
	return result, response, err

}

//GetVersion Get network device version information
/* This API helps to retrieve the version information related to the network device.

 */
func (s *NetworkDeviceService) GetVersion() (*ResponseNetworkDeviceGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevice/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkDeviceGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkDeviceGetVersion)
	return result, response, err

}

//MonitorBulkStatusNetworkDevice Monitor bulk request
/* This API allows the client to monitor the bulk request.

@param bulkid bulkid path parameter.
*/
func (s *NetworkDeviceService) MonitorBulkStatusNetworkDevice(bulkid string) (*ResponseNetworkDeviceMonitorBulkStatusNetworkDevice, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevice/bulk/{bulkid}"
	path = strings.Replace(path, "{bulkid}", fmt.Sprintf("%v", bulkid), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkDeviceMonitorBulkStatusNetworkDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MonitorBulkStatusNetworkDevice")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkDeviceMonitorBulkStatusNetworkDevice)
	return result, response, err

}

//CreateNetworkDevice Create network device
/* This API creates a network device.

 */
func (s *NetworkDeviceService) CreateNetworkDevice(requestNetworkDeviceCreateNetworkDevice *RequestNetworkDeviceCreateNetworkDevice) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevice"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkDeviceCreateNetworkDevice).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkDevice")
	}

	return response, err

}

//UpdateNetworkDeviceByName Update network device by name
/* This API allows the client to update a network device by name.

@param name name path parameter.
*/
func (s *NetworkDeviceService) UpdateNetworkDeviceByName(name string, requestNetworkDeviceUpdateNetworkDeviceByName *RequestNetworkDeviceUpdateNetworkDeviceByName) (*ResponseNetworkDeviceUpdateNetworkDeviceByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevice/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkDeviceUpdateNetworkDeviceByName).
		SetResult(&ResponseNetworkDeviceUpdateNetworkDeviceByName{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkDeviceByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkDeviceUpdateNetworkDeviceByName)
	return result, response, err

}

//UpdateNetworkDeviceByID Update network device by ID
/* This API allows the client to update a network device by ID.

@param id id path parameter.
*/
func (s *NetworkDeviceService) UpdateNetworkDeviceByID(id string, requestNetworkDeviceUpdateNetworkDeviceById *RequestNetworkDeviceUpdateNetworkDeviceByID) (*ResponseNetworkDeviceUpdateNetworkDeviceByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevice/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkDeviceUpdateNetworkDeviceById).
		SetResult(&ResponseNetworkDeviceUpdateNetworkDeviceByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkDeviceById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkDeviceUpdateNetworkDeviceByID)
	return result, response, err

}

//BulkRequestForNetworkDevice Submit bulk request
/* This API allows the client to submit the bulk request.

 */
func (s *NetworkDeviceService) BulkRequestForNetworkDevice(requestNetworkDeviceBulkRequestForNetworkDevice *RequestNetworkDeviceBulkRequestForNetworkDevice) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevice/bulk/submit"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkDeviceBulkRequestForNetworkDevice).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation BulkRequestForNetworkDevice")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteNetworkDeviceByName Delete network device by name
/* This API deletes a network device by name.

@param name name path parameter.
*/
func (s *NetworkDeviceService) DeleteNetworkDeviceByName(name string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevice/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

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
		return response, fmt.Errorf("error with operation DeleteNetworkDeviceByName")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteNetworkDeviceByID Delete network device by ID
/* This API deletes a network device by ID.

@param id id path parameter.
*/
func (s *NetworkDeviceService) DeleteNetworkDeviceByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/networkdevice/{id}"
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
		return response, fmt.Errorf("error with operation DeleteNetworkDeviceById")
	}

	getCSFRToken(response.Header())

	return response, err

}
