package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type NodeDeploymentService service

type ResponseNodeDeploymentGetNodes struct {
	Response []ResponseNodeDeploymentGetNodesResponse `json:"response,omitempty"` //
}

type ResponseNodeDeploymentGetNodesResponse struct {
	Hostname    string   `json:"hostname,omitempty"`    //
	PersonaType []string `json:"personaType,omitempty"` //
	Roles       []string `json:"roles,omitempty"`       //
	Services    []string `json:"services,omitempty"`    //
	NodeStatus  string   `json:"nodeStatus,omitempty"`  //
}

type ResponseNodeDeploymentRegisterNode struct {
	Code      int    `json:"code,omitempty"`      //
	Message   string `json:"message,omitempty"`   //
	RootCause string `json:"rootCause,omitempty"` //
}

type ResponseNodeDeploymentPromoteNode struct {
	Code      int    `json:"code,omitempty"`      //
	Message   string `json:"message,omitempty"`   //
	RootCause string `json:"rootCause,omitempty"` //
}

type ResponseNodeDeploymentGetNodeDetails struct {
	Response ResponseNodeDeploymentGetNodeDetailsResponse `json:"response,omitempty"` //
}

type ResponseNodeDeploymentGetNodeDetailsResponse struct {
	Hostname               string                                                             `json:"hostname,omitempty"`               //
	Fqdn                   string                                                             `json:"fqdn,omitempty"`                   //
	IPAddress              string                                                             `json:"ipAddress,omitempty"`              //
	NodeType               string                                                             `json:"nodeType,omitempty"`               //
	Administration         ResponseNodeDeploymentGetNodeDetailsResponseAdministration         `json:"administration,omitempty"`         //
	GeneralSettings        ResponseNodeDeploymentGetNodeDetailsResponseGeneralSettings        `json:"generalSettings,omitempty"`        //
	ProfilingConfiguration ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfiguration `json:"profilingConfiguration,omitempty"` //
}

type ResponseNodeDeploymentGetNodeDetailsResponseAdministration struct {
	IsEnabled bool   `json:"isEnabled,omitempty"` //
	Role      string `json:"role,omitempty"`      //
}

type ResponseNodeDeploymentGetNodeDetailsResponseGeneralSettings struct {
	Monitoring ResponseNodeDeploymentGetNodeDetailsResponseGeneralSettingsMonitoring `json:"monitoring,omitempty"` //
}

type ResponseNodeDeploymentGetNodeDetailsResponseGeneralSettingsMonitoring struct {
	IsEnabled           bool                                                                               `json:"isEnabled,omitempty"`           //
	Role                string                                                                             `json:"role,omitempty"`                //
	OtherMonitoringNode string                                                                             `json:"otherMonitoringNode,omitempty"` //
	IsMntDedicated      bool                                                                               `json:"isMntDedicated,omitempty"`      //
	Policyservice       ResponseNodeDeploymentGetNodeDetailsResponseGeneralSettingsMonitoringPolicyservice `json:"policyservice,omitempty"`       //
	EnablePXGrid        bool                                                                               `json:"enablePXGrid,omitempty"`        //
}

type ResponseNodeDeploymentGetNodeDetailsResponseGeneralSettingsMonitoringPolicyservice struct {
	Enabled                      bool                                                                                             `json:"enabled,omitempty"`                      //
	SessionService               ResponseNodeDeploymentGetNodeDetailsResponseGeneralSettingsMonitoringPolicyserviceSessionService `json:"sessionService,omitempty"`               //
	EnableProfilingService       bool                                                                                             `json:"enableProfilingService,omitempty"`       //
	EnableNACService             bool                                                                                             `json:"enableNACService,omitempty"`             //
	Sxpservice                   ResponseNodeDeploymentGetNodeDetailsResponseGeneralSettingsMonitoringPolicyserviceSxpservice     `json:"sxpservice,omitempty"`                   //
	EnableDeviceAdminService     bool                                                                                             `json:"enableDeviceAdminService,omitempty"`     //
	EnablePassiveIDentityService bool                                                                                             `json:"enablePassiveIdentityService,omitempty"` //
}

type ResponseNodeDeploymentGetNodeDetailsResponseGeneralSettingsMonitoringPolicyserviceSessionService struct {
	IsEnabled bool   `json:"isEnabled,omitempty"` //
	Nodegroup string `json:"nodegroup,omitempty"` //
}

type ResponseNodeDeploymentGetNodeDetailsResponseGeneralSettingsMonitoringPolicyserviceSxpservice struct {
	IsEnabled     bool   `json:"isEnabled,omitempty"`     //
	UserInterface string `json:"userInterface,omitempty"` //
}

type ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfiguration struct {
	Netflow         ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationNetflow         `json:"netflow,omitempty"`         //
	Dhcp            ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationDhcp            `json:"dhcp,omitempty"`            //
	DhcpSpan        ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationDhcpSpan        `json:"dhcpSpan,omitempty"`        //
	HTTP            ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationHTTP            `json:"http,omitempty"`            //
	Radius          ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationRadius          `json:"radius,omitempty"`          //
	Nmap            ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationNmap            `json:"nmap,omitempty"`            //
	DNS             ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationDNS             `json:"dns,omitempty"`             //
	SNMPQuery       ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationSNMPQuery       `json:"snmpQuery,omitempty"`       //
	SNMPTrap        ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationSNMPTrap        `json:"snmpTrap,omitempty"`        //
	ActiveDirectory ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationActiveDirectory `json:"activeDirectory,omitempty"` //
	Pxgrid          ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationPxgrid          `json:"pxgrid,omitempty"`          //
}

type ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationNetflow struct {
	Enabled     bool   `json:"enabled,omitempty"`     //
	Interface   string `json:"interface,omitempty"`   //
	Port        int    `json:"port,omitempty"`        //
	Description string `json:"description,omitempty"` //
}

type ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationDhcp struct {
	Enabled     bool   `json:"enabled,omitempty"`     //
	Interface   string `json:"interface,omitempty"`   //
	Port        int    `json:"port,omitempty"`        //
	Description string `json:"description,omitempty"` //
}

type ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationDhcpSpan struct {
	Enabled     bool   `json:"enabled,omitempty"`     //
	Interface   string `json:"interface,omitempty"`   //
	Description string `json:"description,omitempty"` //
}

type ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationHTTP struct {
	Enabled     bool   `json:"enabled,omitempty"`     //
	Interface   string `json:"interface,omitempty"`   //
	Description string `json:"description,omitempty"` //
}

type ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationRadius struct {
	Enabled     bool   `json:"enabled,omitempty"`     //
	Description string `json:"description,omitempty"` //
}

type ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationNmap struct {
	Enabled     bool   `json:"enabled,omitempty"`     //
	Description string `json:"description,omitempty"` //
}

type ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationDNS struct {
	Enabled     bool   `json:"enabled,omitempty"`     //
	Description string `json:"description,omitempty"` //
}

type ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationSNMPQuery struct {
	Enabled      bool   `json:"enabled,omitempty"`      //
	Description  string `json:"description,omitempty"`  //
	Retries      int    `json:"retries,omitempty"`      //
	Timeout      int    `json:"timeout,omitempty"`      //
	EventTimeout int    `json:"eventTimeout,omitempty"` //
}

type ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationSNMPTrap struct {
	LinkTrapQuery bool   `json:"linkTrapQuery,omitempty"` //
	MacTrapQuery  bool   `json:"macTrapQuery,omitempty"`  //
	Interface     string `json:"interface,omitempty"`     //
	Port          int    `json:"port,omitempty"`          //
	Description   string `json:"description,omitempty"`   //
}

type ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationActiveDirectory struct {
	Enabled          bool   `json:"enabled,omitempty"`          //
	DaysBeforeRescan int    `json:"daysBeforeRescan,omitempty"` //
	Description      string `json:"description,omitempty"`      //
}

type ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationPxgrid struct {
	Enabled     bool   `json:"enabled,omitempty"`     //
	Description string `json:"description,omitempty"` //
}

type ResponseNodeDeploymentUpdateNode struct {
	Code      int    `json:"code,omitempty"`      //
	Message   string `json:"message,omitempty"`   //
	RootCause string `json:"rootCause,omitempty"` //
}

type ResponseNodeDeploymentDeleteNode struct {
	Code      int    `json:"code,omitempty"`      //
	Message   string `json:"message,omitempty"`   //
	RootCause string `json:"rootCause,omitempty"` //
}

type RequestNodeDeploymentRegisterNode struct {
	Fdqn                 string                                                 `json:"fdqn,omitempty"`                 //
	UserName             string                                                 `json:"userName,omitempty"`             //
	Password             string                                                 `json:"password,omitempty"`             //
	Administration       *RequestNodeDeploymentRegisterNodeAdministration       `json:"administration,omitempty"`       //
	GeneralSettings      *RequestNodeDeploymentRegisterNodeGeneralSettings      `json:"generalSettings,omitempty"`      //
	ProfileConfiguration *RequestNodeDeploymentRegisterNodeProfileConfiguration `json:"profileConfiguration,omitempty"` //
}

type RequestNodeDeploymentRegisterNodeAdministration struct {
	IsEnabled *bool  `json:"isEnabled,omitempty"` //
	Role      string `json:"role,omitempty"`      //
}

type RequestNodeDeploymentRegisterNodeGeneralSettings struct {
	Monitoring *RequestNodeDeploymentRegisterNodeGeneralSettingsMonitoring `json:"monitoring,omitempty"` //
}

type RequestNodeDeploymentRegisterNodeGeneralSettingsMonitoring struct {
	IsEnabled           *bool                                                                    `json:"isEnabled,omitempty"`           //
	Role                string                                                                   `json:"role,omitempty"`                //
	OtherMonitoringNode string                                                                   `json:"otherMonitoringNode,omitempty"` //
	IsMntDedicated      *bool                                                                    `json:"isMntDedicated,omitempty"`      //
	Policyservice       *RequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyservice `json:"policyservice,omitempty"`       //
	EnablePXGrid        *bool                                                                    `json:"enablePXGrid,omitempty"`        //
}

type RequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyservice struct {
	Enabled                      *bool                                                                                  `json:"enabled,omitempty"`                      //
	SessionService               *RequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyserviceSessionService `json:"sessionService,omitempty"`               //
	EnableProfilingService       *bool                                                                                  `json:"enableProfilingService,omitempty"`       //
	EnableNACService             *bool                                                                                  `json:"enableNACService,omitempty"`             //
	Sxpservice                   *RequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyserviceSxpservice     `json:"sxpservice,omitempty"`                   //
	EnableDeviceAdminService     *bool                                                                                  `json:"enableDeviceAdminService,omitempty"`     //
	EnablePassiveIDentityService *bool                                                                                  `json:"enablePassiveIdentityService,omitempty"` //
}

type RequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyserviceSessionService struct {
	IsEnabled *bool  `json:"isEnabled,omitempty"` //
	Nodegroup string `json:"nodegroup,omitempty"` //
}

type RequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyserviceSxpservice struct {
	IsEnabled     *bool  `json:"isEnabled,omitempty"`     //
	UserInterface string `json:"userInterface,omitempty"` //
}

type RequestNodeDeploymentRegisterNodeProfileConfiguration struct {
	Netflow         *RequestNodeDeploymentRegisterNodeProfileConfigurationNetflow         `json:"netflow,omitempty"`         //
	Dhcp            *RequestNodeDeploymentRegisterNodeProfileConfigurationDhcp            `json:"dhcp,omitempty"`            //
	DhcpSpan        *RequestNodeDeploymentRegisterNodeProfileConfigurationDhcpSpan        `json:"dhcpSpan,omitempty"`        //
	HTTP            *RequestNodeDeploymentRegisterNodeProfileConfigurationHTTP            `json:"http,omitempty"`            //
	Radius          *RequestNodeDeploymentRegisterNodeProfileConfigurationRadius          `json:"radius,omitempty"`          //
	Nmap            *RequestNodeDeploymentRegisterNodeProfileConfigurationNmap            `json:"nmap,omitempty"`            //
	DNS             *RequestNodeDeploymentRegisterNodeProfileConfigurationDNS             `json:"dns,omitempty"`             //
	SNMPQuery       *RequestNodeDeploymentRegisterNodeProfileConfigurationSNMPQuery       `json:"snmpQuery,omitempty"`       //
	SNMPTrap        *RequestNodeDeploymentRegisterNodeProfileConfigurationSNMPTrap        `json:"snmpTrap,omitempty"`        //
	ActiveDirectory *RequestNodeDeploymentRegisterNodeProfileConfigurationActiveDirectory `json:"activeDirectory,omitempty"` //
	Pxgrid          *RequestNodeDeploymentRegisterNodeProfileConfigurationPxgrid          `json:"pxgrid,omitempty"`          //
}

type RequestNodeDeploymentRegisterNodeProfileConfigurationNetflow struct {
	Enabled     *bool  `json:"enabled,omitempty"`     //
	Interface   string `json:"interface,omitempty"`   //
	Port        *int   `json:"port,omitempty"`        //
	Description string `json:"description,omitempty"` //
}

type RequestNodeDeploymentRegisterNodeProfileConfigurationDhcp struct {
	Enabled     *bool  `json:"enabled,omitempty"`     //
	Interface   string `json:"interface,omitempty"`   //
	Port        *int   `json:"port,omitempty"`        //
	Description string `json:"description,omitempty"` //
}

type RequestNodeDeploymentRegisterNodeProfileConfigurationDhcpSpan struct {
	Enabled     *bool  `json:"enabled,omitempty"`     //
	Interface   string `json:"interface,omitempty"`   //
	Description string `json:"description,omitempty"` //
}

type RequestNodeDeploymentRegisterNodeProfileConfigurationHTTP struct {
	Enabled     *bool  `json:"enabled,omitempty"`     //
	Interface   string `json:"interface,omitempty"`   //
	Description string `json:"description,omitempty"` //
}

type RequestNodeDeploymentRegisterNodeProfileConfigurationRadius struct {
	Enabled     *bool  `json:"enabled,omitempty"`     //
	Description string `json:"description,omitempty"` //
}

type RequestNodeDeploymentRegisterNodeProfileConfigurationNmap struct {
	Enabled     *bool  `json:"enabled,omitempty"`     //
	Description string `json:"description,omitempty"` //
}

type RequestNodeDeploymentRegisterNodeProfileConfigurationDNS struct {
	Enabled     *bool  `json:"enabled,omitempty"`     //
	Description string `json:"description,omitempty"` //
}

type RequestNodeDeploymentRegisterNodeProfileConfigurationSNMPQuery struct {
	Enabled      *bool  `json:"enabled,omitempty"`      //
	Description  string `json:"description,omitempty"`  //
	Retries      *int   `json:"retries,omitempty"`      //
	Timeout      *int   `json:"timeout,omitempty"`      //
	EventTimeout *int   `json:"eventTimeout,omitempty"` //
}

type RequestNodeDeploymentRegisterNodeProfileConfigurationSNMPTrap struct {
	LinkTrapQuery *bool  `json:"linkTrapQuery,omitempty"` //
	MacTrapQuery  *bool  `json:"macTrapQuery,omitempty"`  //
	Interface     string `json:"interface,omitempty"`     //
	Port          *int   `json:"port,omitempty"`          //
	Description   string `json:"description,omitempty"`   //
}

type RequestNodeDeploymentRegisterNodeProfileConfigurationActiveDirectory struct {
	Enabled          *bool  `json:"enabled,omitempty"`          //
	DaysBeforeRescan *int   `json:"daysBeforeRescan,omitempty"` //
	Description      string `json:"description,omitempty"`      //
}

type RequestNodeDeploymentRegisterNodeProfileConfigurationPxgrid struct {
	Enabled     *bool  `json:"enabled,omitempty"`     //
	Description string `json:"description,omitempty"` //
}

type RequestNodeDeploymentPromoteNode struct {
	PromotionType string `json:"promotionType,omitempty"` //
}

type RequestNodeDeploymentUpdateNode struct {
	Response *RequestNodeDeploymentUpdateNodeResponse `json:"response,omitempty"` //
}

type RequestNodeDeploymentUpdateNodeResponse struct {
	GeneralSettings      *RequestNodeDeploymentUpdateNodeResponseGeneralSettings      `json:"generalSettings,omitempty"`      //
	ProfileConfiguration *RequestNodeDeploymentUpdateNodeResponseProfileConfiguration `json:"profileConfiguration,omitempty"` //
}

type RequestNodeDeploymentUpdateNodeResponseGeneralSettings struct {
	Monitoring *RequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoring `json:"monitoring,omitempty"` //
}

type RequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoring struct {
	IsEnabled           *bool                                                                          `json:"isEnabled,omitempty"`           //
	Role                string                                                                         `json:"role,omitempty"`                //
	OtherMonitoringNode string                                                                         `json:"otherMonitoringNode,omitempty"` //
	IsMntDedicated      *bool                                                                          `json:"isMntDedicated,omitempty"`      //
	Policyservice       *RequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyservice `json:"policyservice,omitempty"`       //
	EnablePXGrid        *bool                                                                          `json:"enablePXGrid,omitempty"`        //
}

type RequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyservice struct {
	Enabled                      *bool                                                                                        `json:"enabled,omitempty"`                      //
	SessionService               *RequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyserviceSessionService `json:"sessionService,omitempty"`               //
	EnableProfilingService       *bool                                                                                        `json:"enableProfilingService,omitempty"`       //
	EnableNACService             *bool                                                                                        `json:"enableNACService,omitempty"`             //
	Sxpservice                   *RequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyserviceSxpservice     `json:"sxpservice,omitempty"`                   //
	EnableDeviceAdminService     *bool                                                                                        `json:"enableDeviceAdminService,omitempty"`     //
	EnablePassiveIDentityService *bool                                                                                        `json:"enablePassiveIdentityService,omitempty"` //
}

type RequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyserviceSessionService struct {
	IsEnabled *bool  `json:"isEnabled,omitempty"` //
	Nodegroup string `json:"nodegroup,omitempty"` //
}

type RequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyserviceSxpservice struct {
	IsEnabled     *bool  `json:"isEnabled,omitempty"`     //
	UserInterface string `json:"userInterface,omitempty"` //
}

type RequestNodeDeploymentUpdateNodeResponseProfileConfiguration struct {
	Netflow         *RequestNodeDeploymentUpdateNodeResponseProfileConfigurationNetflow         `json:"netflow,omitempty"`         //
	Dhcp            *RequestNodeDeploymentUpdateNodeResponseProfileConfigurationDhcp            `json:"dhcp,omitempty"`            //
	DhcpSpan        *RequestNodeDeploymentUpdateNodeResponseProfileConfigurationDhcpSpan        `json:"dhcpSpan,omitempty"`        //
	HTTP            *RequestNodeDeploymentUpdateNodeResponseProfileConfigurationHTTP            `json:"http,omitempty"`            //
	Radius          *RequestNodeDeploymentUpdateNodeResponseProfileConfigurationRadius          `json:"radius,omitempty"`          //
	Nmap            *RequestNodeDeploymentUpdateNodeResponseProfileConfigurationNmap            `json:"nmap,omitempty"`            //
	DNS             *RequestNodeDeploymentUpdateNodeResponseProfileConfigurationDNS             `json:"dns,omitempty"`             //
	SNMPQuery       *RequestNodeDeploymentUpdateNodeResponseProfileConfigurationSNMPQuery       `json:"snmpQuery,omitempty"`       //
	SNMPTrap        *RequestNodeDeploymentUpdateNodeResponseProfileConfigurationSNMPTrap        `json:"snmpTrap,omitempty"`        //
	ActiveDirectory *RequestNodeDeploymentUpdateNodeResponseProfileConfigurationActiveDirectory `json:"activeDirectory,omitempty"` //
	Pxgrid          *RequestNodeDeploymentUpdateNodeResponseProfileConfigurationPxgrid          `json:"pxgrid,omitempty"`          //
}

type RequestNodeDeploymentUpdateNodeResponseProfileConfigurationNetflow struct {
	Enabled     *bool  `json:"enabled,omitempty"`     //
	Interface   string `json:"interface,omitempty"`   //
	Port        *int   `json:"port,omitempty"`        //
	Description string `json:"description,omitempty"` //
}

type RequestNodeDeploymentUpdateNodeResponseProfileConfigurationDhcp struct {
	Enabled     *bool  `json:"enabled,omitempty"`     //
	Interface   string `json:"interface,omitempty"`   //
	Port        *int   `json:"port,omitempty"`        //
	Description string `json:"description,omitempty"` //
}

type RequestNodeDeploymentUpdateNodeResponseProfileConfigurationDhcpSpan struct {
	Enabled     *bool  `json:"enabled,omitempty"`     //
	Interface   string `json:"interface,omitempty"`   //
	Description string `json:"description,omitempty"` //
}

type RequestNodeDeploymentUpdateNodeResponseProfileConfigurationHTTP struct {
	Enabled     *bool  `json:"enabled,omitempty"`     //
	Interface   string `json:"interface,omitempty"`   //
	Description string `json:"description,omitempty"` //
}

type RequestNodeDeploymentUpdateNodeResponseProfileConfigurationRadius struct {
	Enabled     *bool  `json:"enabled,omitempty"`     //
	Description string `json:"description,omitempty"` //
}

type RequestNodeDeploymentUpdateNodeResponseProfileConfigurationNmap struct {
	Enabled     *bool  `json:"enabled,omitempty"`     //
	Description string `json:"description,omitempty"` //
}

type RequestNodeDeploymentUpdateNodeResponseProfileConfigurationDNS struct {
	Enabled     *bool  `json:"enabled,omitempty"`     //
	Description string `json:"description,omitempty"` //
}

type RequestNodeDeploymentUpdateNodeResponseProfileConfigurationSNMPQuery struct {
	Enabled      *bool  `json:"enabled,omitempty"`      //
	Description  string `json:"description,omitempty"`  //
	Retries      *int   `json:"retries,omitempty"`      //
	Timeout      *int   `json:"timeout,omitempty"`      //
	EventTimeout *int   `json:"eventTimeout,omitempty"` //
}

type RequestNodeDeploymentUpdateNodeResponseProfileConfigurationSNMPTrap struct {
	LinkTrapQuery *bool  `json:"linkTrapQuery,omitempty"` //
	MacTrapQuery  *bool  `json:"macTrapQuery,omitempty"`  //
	Interface     string `json:"interface,omitempty"`     //
	Port          *int   `json:"port,omitempty"`          //
	Description   string `json:"description,omitempty"`   //
}

type RequestNodeDeploymentUpdateNodeResponseProfileConfigurationActiveDirectory struct {
	Enabled          *bool  `json:"enabled,omitempty"`          //
	DaysBeforeRescan *int   `json:"daysBeforeRescan,omitempty"` //
	Description      string `json:"description,omitempty"`      //
}

type RequestNodeDeploymentUpdateNodeResponseProfileConfigurationPxgrid struct {
	Enabled     *bool  `json:"enabled,omitempty"`     //
	Description string `json:"description,omitempty"` //
}

//GetNodes Retreives the list of all deployed nodes in the cluster.
/* Discovers all deployment nodes in the cluster.
It provides basic information about each of deployed nodes in the cluster like Hostname, personas, status, roles and services.


*/
func (s *NodeDeploymentService) GetNodes() (*ResponseNodeDeploymentGetNodes, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node/"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeDeploymentGetNodes{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNodes")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeDeploymentGetNodes)
	return result, response, err

}

//GetNodeDetails Retreives details of a deployed node.
/* It provides detailed information of the deployed node in the cluster.


@param hostname hostname path parameter. ID of the existing deployed node.
*/
func (s *NodeDeploymentService) GetNodeDetails(hostname string) (*ResponseNodeDeploymentGetNodeDetails, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node/{hostname}"
	path = strings.Replace(path, "{hostname}", fmt.Sprintf("%v", hostname), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeDeploymentGetNodeDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNodeDetails")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeDeploymentGetNodeDetails)
	return result, response, err

}

//RegisterNode Registers a standalone node in the cluster
/* Register ISE node to form a multi-node deployment


 */
func (s *NodeDeploymentService) RegisterNode(requestNodeDeploymentRegisterNode *RequestNodeDeploymentRegisterNode) (*ResponseNodeDeploymentRegisterNode, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node/"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNodeDeploymentRegisterNode).
		SetResult(&ResponseNodeDeploymentRegisterNode{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RegisterNode")
	}

	result := response.Result().(*ResponseNodeDeploymentRegisterNode)
	return result, response, err

}

//PromoteNode Promotes a secondary or standalone node to primary node
/* Changes the cluster setting by promoting a node to primary when exceuted on standalone or secondary node.
It could also be used to convert a deployment node to standalone node.


*/
func (s *NodeDeploymentService) PromoteNode(requestNodeDeploymentPromoteNode *RequestNodeDeploymentPromoteNode) (*ResponseNodeDeploymentPromoteNode, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node-promotion/"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNodeDeploymentPromoteNode).
		SetResult(&ResponseNodeDeploymentPromoteNode{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation PromoteNode")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeDeploymentPromoteNode)
	return result, response, err

}

//UpdateNode Replaces the existing configuration of the ISE node with the one provided.
/* Updates the deployed ISE node with the information provided


@param hostname hostname path parameter. ID of the existing deployed node.
*/
func (s *NodeDeploymentService) UpdateNode(hostname string, requestNodeDeploymentUpdateNode *RequestNodeDeploymentUpdateNode) (*ResponseNodeDeploymentUpdateNode, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node/{hostname}"
	path = strings.Replace(path, "{hostname}", fmt.Sprintf("%v", hostname), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNodeDeploymentUpdateNode).
		SetResult(&ResponseNodeDeploymentUpdateNode{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNode")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeDeploymentUpdateNode)
	return result, response, err

}

//DeleteNode Remove a deployed node from a cluster.
/* The de-register ednode becomes a standalone Cisco ISE node.
It retains the last configuration that it received rom the PrimaryPAN and assumes the default personas of a standalone node
that are Administration, PolicyService, and Monitoring.


@param hostname hostname path parameter. node name of the existing deployed node.
*/
func (s *NodeDeploymentService) DeleteNode(hostname string) (*ResponseNodeDeploymentDeleteNode, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node/{hostname}"
	path = strings.Replace(path, "{hostname}", fmt.Sprintf("%v", hostname), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeDeploymentDeleteNode{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteNode")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeDeploymentDeleteNode)
	return result, response, err

}
