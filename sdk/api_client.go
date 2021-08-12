package isegosdk

import (
	"crypto/tls"
	"fmt"
	"os"

	"io/ioutil"
	"path/filepath"

	"github.com/go-resty/resty/v2"
)

// RestyClient is the REST Client
var host = ""
var UseAPIGateway = false
var useCSRFToken = false
var CSRFToken string

const ISE_BASE_URL = "ISE_BASE_URL"
const ISE_USERNAME = "ISE_USERNAME"
const ISE_PASSWORD = "ISE_PASSWORD"
const ISE_DEBUG = "ISE_DEBUG"
const ISE_SSL_VERIFY = "ISE_SSL_VERIFY"
const ISE_USE_API_GATEWAY = "ISE_USE_API_GATEWAY"
const ISE_USE_CSRF_TOKEN = "ISE_USE_CSRF_TOKEN"

type FileDownload struct {
	FileName string
	FileData []byte
}

func (f *FileDownload) SaveDownload(path string) error {
	fpath := filepath.Join(path, f.FileName)
	return ioutil.WriteFile(fpath, f.FileData, 0664)
}

type Client struct {
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	AciBindings                                           *AciBindingsService
	AciSettings                                           *AciSettingsService
	AncEndpoint                                           *AncEndpointService
	ActiveDirectory                                       *ActiveDirectoryService
	AdminUser                                             *AdminUserService
	AllowedProtocols                                      *AllowedProtocolsService
	AncPolicy                                             *AncPolicyService
	AuthorizationProfile                                  *AuthorizationProfileService
	ByodPortal                                            *ByodPortalService
	BackupAndRestore                                      *BackupAndRestoreService
	CertificateProfile                                    *CertificateProfileService
	CertificateTemplate                                   *CertificateTemplateService
	Certificates                                          *CertificatesService
	ClearThreatsAndVulnerabilities                        *ClearThreatsAndVulnerabilitiesService
	Consumer                                              *ConsumerService
	DeviceAdministrationAuthenticationRules               *DeviceAdministrationAuthenticationRulesService
	DeviceAdministrationAuthorizationExceptionRules       *DeviceAdministrationAuthorizationExceptionRulesService
	DeviceAdministrationAuthorizationGlobalExceptionRules *DeviceAdministrationAuthorizationGlobalExceptionRulesService
	DeviceAdministrationAuthorizationRules                *DeviceAdministrationAuthorizationRulesService
	DeviceAdministrationCommandSet                        *DeviceAdministrationCommandSetService
	DeviceAdministrationConditions                        *DeviceAdministrationConditionsService
	DeviceAdministrationDictionaryAttributesList          *DeviceAdministrationDictionaryAttributesListService
	DeviceAdministrationIDentityStores                    *DeviceAdministrationIDentityStoresService
	DeviceAdministrationNetworkConditions                 *DeviceAdministrationNetworkConditionsService
	DeviceAdministrationPolicySet                         *DeviceAdministrationPolicySetService
	DeviceAdministrationProfiles                          *DeviceAdministrationProfilesService
	DeviceAdministrationServiceNames                      *DeviceAdministrationServiceNamesService
	DeviceAdministrationTimeDateConditions                *DeviceAdministrationTimeDateConditionsService
	DownloadableACL                                       *DownloadableACLService
	EgressMatrixCell                                      *EgressMatrixCellService
	EndpointCertificate                                   *EndpointCertificateService
	EndpointIDentityGroup                                 *EndpointIDentityGroupService
	ExternalRadiusServer                                  *ExternalRadiusServerService
	FilterPolicy                                          *FilterPolicyService
	GuestLocation                                         *GuestLocationService
	GuestSmtpNotificationConfiguration                    *GuestSmtpNotificationConfigurationService
	GuestSSID                                             *GuestSSIDService
	GuestType                                             *GuestTypeService
	GuestUser                                             *GuestUserService
	HotspotPortal                                         *HotspotPortalService
	IPToSgtMapping                                        *IPToSgtMappingService
	IPToSgtMappingGroup                                   *IPToSgtMappingGroupService
	IDentityGroups                                        *IDentityGroupsService
	IDentitySequence                                      *IDentitySequenceService
	InternalUser                                          *InternalUserService
	Mdm                                                   *MdmService
	Misc                                                  *MiscService
	MyDevicePortal                                        *MyDevicePortalService
	NativeSupplicantProfile                               *NativeSupplicantProfileService
	NetworkAccessAuthenticationRules                      *NetworkAccessAuthenticationRulesService
	NetworkAccessAuthorizationExceptionRules              *NetworkAccessAuthorizationExceptionRulesService
	NetworkAccessAuthorizationGlobalExceptionRules        *NetworkAccessAuthorizationGlobalExceptionRulesService
	NetworkAccessAuthorizationRules                       *NetworkAccessAuthorizationRulesService
	NetworkAccessConditions                               *NetworkAccessConditionsService
	NetworkAccessDictionary                               *NetworkAccessDictionaryService
	NetworkAccessDictionaryAttribute                      *NetworkAccessDictionaryAttributeService
	NetworkAccessDictionaryAttributesList                 *NetworkAccessDictionaryAttributesListService
	NetworkAccessIDentityStores                           *NetworkAccessIDentityStoresService
	NetworkAccessNetworkConditions                        *NetworkAccessNetworkConditionsService
	NetworkAccessPolicySet                                *NetworkAccessPolicySetService
	NetworkAccessProfiles                                 *NetworkAccessProfilesService
	NetworkAccessSecurityGroups                           *NetworkAccessSecurityGroupsService
	NetworkAccessServiceNames                             *NetworkAccessServiceNamesService
	NetworkAccessTimeDateConditions                       *NetworkAccessTimeDateConditionsService
	NetworkDevice                                         *NetworkDeviceService
	NetworkDeviceGroup                                    *NetworkDeviceGroupService
	NodeDeployment                                        *NodeDeploymentService
	NodeGroup                                             *NodeGroupService
	NodeDetails                                           *NodeDetailsService
	PanHa                                                 *PanHaService
	PortalGlobalSetting                                   *PortalGlobalSettingService
	PortalTheme                                           *PortalThemeService
	Profiler                                              *ProfilerService
	ProfilerProfile                                       *ProfilerProfileService
	Provider                                              *ProviderService
	PsnNodeDetailsWithRadiusService                       *PsnNodeDetailsWithRadiusServiceService
	PullDeploymentInfo                                    *PullDeploymentInfoService
	PxGridSettings                                        *PxGridSettingsService
	RadiusFailure                                         *RadiusFailureService
	RadiusServerSequence                                  *RadiusServerSequenceService
	RestidStore                                           *RestidStoreService
	ReplicationStatus                                     *ReplicationStatusService
	Repository                                            *RepositoryService
	SmsProvider                                           *SmsProviderService
	SxpConnections                                        *SxpConnectionsService
	SxpLocalBindings                                      *SxpLocalBindingsService
	SxpVpns                                               *SxpVpnsService
	SecurityGroupToVirtualNetwork                         *SecurityGroupToVirtualNetworkService
	SecurityGroups                                        *SecurityGroupsService
	SecurityGroupsACLs                                    *SecurityGroupsACLsService
	SelfRegisteredPortal                                  *SelfRegisteredPortalService
	SessionDirectory                                      *SessionDirectoryService
	SponsorGroup                                          *SponsorGroupService
	SponsorGroupMember                                    *SponsorGroupMemberService
	SponsorPortal                                         *SponsorPortalService
	SponsoredGuestPortal                                  *SponsoredGuestPortalService
	SupportBundleDownload                                 *SupportBundleDownloadService
	SupportBundleStatus                                   *SupportBundleStatusService
	SupportBundleTriggerConfiguration                     *SupportBundleTriggerConfigurationService
	SyncIseNode                                           *SyncIseNodeService
	SystemHealth                                          *SystemHealthService
	SystemCertificate                                     *SystemCertificateService
	TacacsCommandSets                                     *TacacsCommandSetsService
	TacacsExternalServers                                 *TacacsExternalServersService
	TacacsProfile                                         *TacacsProfileService
	TacacsServerSequence                                  *TacacsServerSequenceService
	TelemetryInformation                                  *TelemetryInformationService
	TrustSecConfiguration                                 *TrustSecConfigurationService
	TrustSecSxp                                           *TrustSecSxpService
	VersionAndPatch                                       *VersionAndPatchService
	VersionInfo                                           *VersionInfoService
	Endpoint                                              *EndpointService
	Portal                                                *PortalService
	PxGridNode                                            *PxGridNodeService
	Tasks                                                 *TasksService
}

type service struct {
	client *resty.Client
}

func setHost(client *resty.Client, module string) {
	port := getPort(module)
	client.SetHostURL(host + port)
}

func getCSFRToken(headers map[string][]string) {
	if useCSRFToken {
		if val, ok := headers["X-Csrf-Token"]; ok {
			CSRFToken = val[0]
		}
	}
}

func setCSRFToken(client *resty.Client) {
	if useCSRFToken {
		tval := "fetch"
		if CSRFToken != "" {
			tval = CSRFToken
		}
		client.SetHeader("X-CSRF-Token", tval)
	}
}

func getPort(group string) string {
	if UseAPIGateway {
		return ""
	}
	if group == "_ui" {
		return ":443"
	}
	if group == "_ers" {
		return ":9060"
	}
	if group == "_mnt" {
		return ":443"
	}
	if group == "_px_grid" {
		return ":8910"
	}
	return ""
}

// Error indicates an error from the invocation of a Cisco ISE.
var Error map[string]interface{}

// NewClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewClient() (*Client, error) {
	var err error
	client := resty.New()
	c := &Client{}
	c.common.client = client
	username := ""
	password := ""

	if os.Getenv(ISE_DEBUG) == "true" {
		client.SetDebug(true)
	}

	if os.Getenv(ISE_SSL_VERIFY) == "false" {
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	if os.Getenv(ISE_BASE_URL) != "" {
		host = os.Getenv(ISE_BASE_URL)
	} else {
		err = fmt.Errorf("enviroment variable %s was not defined", ISE_BASE_URL)
	}

	if os.Getenv(ISE_USERNAME) != "" {
		username = os.Getenv(ISE_USERNAME)
	} else {
		err = fmt.Errorf("enviroment variable %s was not defined", ISE_USERNAME)
	}

	if os.Getenv(ISE_PASSWORD) != "" {
		password = os.Getenv(ISE_PASSWORD)
	} else {
		err = fmt.Errorf("enviroment variable %s was not defined", ISE_PASSWORD)
	}

	if os.Getenv(ISE_USE_API_GATEWAY) == "true" {
		UseAPIGateway = true
	}

	if os.Getenv(ISE_USE_CSRF_TOKEN) == "true" {
		useCSRFToken = true
	}

	if err != nil {
		return nil, err
	}
	c.AciBindings = (*AciBindingsService)(&c.common)
	c.AciSettings = (*AciSettingsService)(&c.common)
	c.AncEndpoint = (*AncEndpointService)(&c.common)
	c.ActiveDirectory = (*ActiveDirectoryService)(&c.common)
	c.AdminUser = (*AdminUserService)(&c.common)
	c.AllowedProtocols = (*AllowedProtocolsService)(&c.common)
	c.AncPolicy = (*AncPolicyService)(&c.common)
	c.AuthorizationProfile = (*AuthorizationProfileService)(&c.common)
	c.ByodPortal = (*ByodPortalService)(&c.common)
	c.BackupAndRestore = (*BackupAndRestoreService)(&c.common)
	c.CertificateProfile = (*CertificateProfileService)(&c.common)
	c.CertificateTemplate = (*CertificateTemplateService)(&c.common)
	c.Certificates = (*CertificatesService)(&c.common)
	c.ClearThreatsAndVulnerabilities = (*ClearThreatsAndVulnerabilitiesService)(&c.common)
	c.Consumer = (*ConsumerService)(&c.common)
	c.DeviceAdministrationAuthenticationRules = (*DeviceAdministrationAuthenticationRulesService)(&c.common)
	c.DeviceAdministrationAuthorizationExceptionRules = (*DeviceAdministrationAuthorizationExceptionRulesService)(&c.common)
	c.DeviceAdministrationAuthorizationGlobalExceptionRules = (*DeviceAdministrationAuthorizationGlobalExceptionRulesService)(&c.common)
	c.DeviceAdministrationAuthorizationRules = (*DeviceAdministrationAuthorizationRulesService)(&c.common)
	c.DeviceAdministrationCommandSet = (*DeviceAdministrationCommandSetService)(&c.common)
	c.DeviceAdministrationConditions = (*DeviceAdministrationConditionsService)(&c.common)
	c.DeviceAdministrationDictionaryAttributesList = (*DeviceAdministrationDictionaryAttributesListService)(&c.common)
	c.DeviceAdministrationIDentityStores = (*DeviceAdministrationIDentityStoresService)(&c.common)
	c.DeviceAdministrationNetworkConditions = (*DeviceAdministrationNetworkConditionsService)(&c.common)
	c.DeviceAdministrationPolicySet = (*DeviceAdministrationPolicySetService)(&c.common)
	c.DeviceAdministrationProfiles = (*DeviceAdministrationProfilesService)(&c.common)
	c.DeviceAdministrationServiceNames = (*DeviceAdministrationServiceNamesService)(&c.common)
	c.DeviceAdministrationTimeDateConditions = (*DeviceAdministrationTimeDateConditionsService)(&c.common)
	c.DownloadableACL = (*DownloadableACLService)(&c.common)
	c.EgressMatrixCell = (*EgressMatrixCellService)(&c.common)
	c.EndpointCertificate = (*EndpointCertificateService)(&c.common)
	c.EndpointIDentityGroup = (*EndpointIDentityGroupService)(&c.common)
	c.ExternalRadiusServer = (*ExternalRadiusServerService)(&c.common)
	c.FilterPolicy = (*FilterPolicyService)(&c.common)
	c.GuestLocation = (*GuestLocationService)(&c.common)
	c.GuestSmtpNotificationConfiguration = (*GuestSmtpNotificationConfigurationService)(&c.common)
	c.GuestSSID = (*GuestSSIDService)(&c.common)
	c.GuestType = (*GuestTypeService)(&c.common)
	c.GuestUser = (*GuestUserService)(&c.common)
	c.HotspotPortal = (*HotspotPortalService)(&c.common)
	c.IPToSgtMapping = (*IPToSgtMappingService)(&c.common)
	c.IPToSgtMappingGroup = (*IPToSgtMappingGroupService)(&c.common)
	c.IDentityGroups = (*IDentityGroupsService)(&c.common)
	c.IDentitySequence = (*IDentitySequenceService)(&c.common)
	c.InternalUser = (*InternalUserService)(&c.common)
	c.Mdm = (*MdmService)(&c.common)
	c.Misc = (*MiscService)(&c.common)
	c.MyDevicePortal = (*MyDevicePortalService)(&c.common)
	c.NativeSupplicantProfile = (*NativeSupplicantProfileService)(&c.common)
	c.NetworkAccessAuthenticationRules = (*NetworkAccessAuthenticationRulesService)(&c.common)
	c.NetworkAccessAuthorizationExceptionRules = (*NetworkAccessAuthorizationExceptionRulesService)(&c.common)
	c.NetworkAccessAuthorizationGlobalExceptionRules = (*NetworkAccessAuthorizationGlobalExceptionRulesService)(&c.common)
	c.NetworkAccessAuthorizationRules = (*NetworkAccessAuthorizationRulesService)(&c.common)
	c.NetworkAccessConditions = (*NetworkAccessConditionsService)(&c.common)
	c.NetworkAccessDictionary = (*NetworkAccessDictionaryService)(&c.common)
	c.NetworkAccessDictionaryAttribute = (*NetworkAccessDictionaryAttributeService)(&c.common)
	c.NetworkAccessDictionaryAttributesList = (*NetworkAccessDictionaryAttributesListService)(&c.common)
	c.NetworkAccessIDentityStores = (*NetworkAccessIDentityStoresService)(&c.common)
	c.NetworkAccessNetworkConditions = (*NetworkAccessNetworkConditionsService)(&c.common)
	c.NetworkAccessPolicySet = (*NetworkAccessPolicySetService)(&c.common)
	c.NetworkAccessProfiles = (*NetworkAccessProfilesService)(&c.common)
	c.NetworkAccessSecurityGroups = (*NetworkAccessSecurityGroupsService)(&c.common)
	c.NetworkAccessServiceNames = (*NetworkAccessServiceNamesService)(&c.common)
	c.NetworkAccessTimeDateConditions = (*NetworkAccessTimeDateConditionsService)(&c.common)
	c.NetworkDevice = (*NetworkDeviceService)(&c.common)
	c.NetworkDeviceGroup = (*NetworkDeviceGroupService)(&c.common)
	c.NodeDeployment = (*NodeDeploymentService)(&c.common)
	c.NodeGroup = (*NodeGroupService)(&c.common)
	c.NodeDetails = (*NodeDetailsService)(&c.common)
	c.PanHa = (*PanHaService)(&c.common)
	c.PortalGlobalSetting = (*PortalGlobalSettingService)(&c.common)
	c.PortalTheme = (*PortalThemeService)(&c.common)
	c.Profiler = (*ProfilerService)(&c.common)
	c.ProfilerProfile = (*ProfilerProfileService)(&c.common)
	c.Provider = (*ProviderService)(&c.common)
	c.PsnNodeDetailsWithRadiusService = (*PsnNodeDetailsWithRadiusServiceService)(&c.common)
	c.PullDeploymentInfo = (*PullDeploymentInfoService)(&c.common)
	c.PxGridSettings = (*PxGridSettingsService)(&c.common)
	c.RadiusFailure = (*RadiusFailureService)(&c.common)
	c.RadiusServerSequence = (*RadiusServerSequenceService)(&c.common)
	c.RestidStore = (*RestidStoreService)(&c.common)
	c.ReplicationStatus = (*ReplicationStatusService)(&c.common)
	c.Repository = (*RepositoryService)(&c.common)
	c.SmsProvider = (*SmsProviderService)(&c.common)
	c.SxpConnections = (*SxpConnectionsService)(&c.common)
	c.SxpLocalBindings = (*SxpLocalBindingsService)(&c.common)
	c.SxpVpns = (*SxpVpnsService)(&c.common)
	c.SecurityGroupToVirtualNetwork = (*SecurityGroupToVirtualNetworkService)(&c.common)
	c.SecurityGroups = (*SecurityGroupsService)(&c.common)
	c.SecurityGroupsACLs = (*SecurityGroupsACLsService)(&c.common)
	c.SelfRegisteredPortal = (*SelfRegisteredPortalService)(&c.common)
	c.SessionDirectory = (*SessionDirectoryService)(&c.common)
	c.SponsorGroup = (*SponsorGroupService)(&c.common)
	c.SponsorGroupMember = (*SponsorGroupMemberService)(&c.common)
	c.SponsorPortal = (*SponsorPortalService)(&c.common)
	c.SponsoredGuestPortal = (*SponsoredGuestPortalService)(&c.common)
	c.SupportBundleDownload = (*SupportBundleDownloadService)(&c.common)
	c.SupportBundleStatus = (*SupportBundleStatusService)(&c.common)
	c.SupportBundleTriggerConfiguration = (*SupportBundleTriggerConfigurationService)(&c.common)
	c.SyncIseNode = (*SyncIseNodeService)(&c.common)
	c.SystemHealth = (*SystemHealthService)(&c.common)
	c.SystemCertificate = (*SystemCertificateService)(&c.common)
	c.TacacsCommandSets = (*TacacsCommandSetsService)(&c.common)
	c.TacacsExternalServers = (*TacacsExternalServersService)(&c.common)
	c.TacacsProfile = (*TacacsProfileService)(&c.common)
	c.TacacsServerSequence = (*TacacsServerSequenceService)(&c.common)
	c.TelemetryInformation = (*TelemetryInformationService)(&c.common)
	c.TrustSecConfiguration = (*TrustSecConfigurationService)(&c.common)
	c.TrustSecSxp = (*TrustSecSxpService)(&c.common)
	c.VersionAndPatch = (*VersionAndPatchService)(&c.common)
	c.VersionInfo = (*VersionInfoService)(&c.common)
	c.Endpoint = (*EndpointService)(&c.common)
	c.Portal = (*PortalService)(&c.common)
	c.PxGridNode = (*PxGridNodeService)(&c.common)
	c.Tasks = (*TasksService)(&c.common)

	client.SetBasicAuth(username, password)
	return c, nil
}

//NewClientWithOptions is the client with options passed with parameters
func NewClientWithOptions(baseURL string, username string, password string, debug string, sslVerify string, useAPIGateway string, useCSRFToken string) (*Client, error) {
	os.Setenv(ISE_BASE_URL, baseURL)
	os.Setenv(ISE_USERNAME, username)
	os.Setenv(ISE_PASSWORD, password)
	os.Setenv(ISE_DEBUG, debug)
	os.Setenv(ISE_SSL_VERIFY, sslVerify)
	os.Setenv(ISE_USE_API_GATEWAY, useAPIGateway)
	os.Setenv(ISE_USE_CSRF_TOKEN, useCSRFToken)
	return NewClient()
}
