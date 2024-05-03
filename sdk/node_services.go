package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type NodeServicesService service

type ResponseNodeServicesGetInterfaces struct {
	Response *[]ResponseNodeServicesGetInterfacesResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  //
}

type ResponseNodeServicesGetInterfacesResponse struct {
	Interface string `json:"interface,omitempty"` //
}

type ResponseNodeServicesGetSxpInterface struct {
	Response *ResponseNodeServicesGetSxpInterfaceResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  //
}

type ResponseNodeServicesGetSxpInterfaceResponse struct {
	Interface string `json:"interface,omitempty"` //
}

type ResponseNodeServicesSetSxpInterface struct {
	Success *ResponseNodeServicesSetSxpInterfaceSuccess `json:"success,omitempty"` //
	Version string                                      `json:"version,omitempty"` //
}

type ResponseNodeServicesSetSxpInterfaceSuccess struct {
	Message string `json:"message,omitempty"` //
}

type ResponseNodeServicesGetProfilerProbeConfig struct {
	Response *ResponseNodeServicesGetProfilerProbeConfigResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  //
}

type ResponseNodeServicesGetProfilerProbeConfigResponse struct {
	ActiveDirectory *ResponseNodeServicesGetProfilerProbeConfigResponseActiveDirectory `json:"activeDirectory,omitempty"` // The Active Directory probe queries the Active Directory for Windows information.
	Dhcp            *ResponseNodeServicesGetProfilerProbeConfigResponseDhcp            `json:"dhcp,omitempty"`            // The DHCP probe listens for DHCP packets from IP helpers.
	DhcpSpan        *ResponseNodeServicesGetProfilerProbeConfigResponseDhcpSpan        `json:"dhcpSpan,omitempty"`        // The DHCP SPAN probe collects DHCP packets.
	DNS             *ResponseNodeServicesGetProfilerProbeConfigResponseDNS             `json:"dns,omitempty"`             // The DNS probe performs a DNS lookup for the FQDN.
	HTTP            *ResponseNodeServicesGetProfilerProbeConfigResponseHTTP            `json:"http,omitempty"`            // The HTTP probe receives and parses HTTP packets.
	Netflow         *ResponseNodeServicesGetProfilerProbeConfigResponseNetflow         `json:"netflow,omitempty"`         // The NetFlow probe collects the NetFlow packets that are sent to it from routers.
	Nmap            *[]ResponseNodeServicesGetProfilerProbeConfigResponseNmap          `json:"nmap,omitempty"`            // The NMAP probe scans endpoints for open ports and OS.
	Pxgrid          *[]ResponseNodeServicesGetProfilerProbeConfigResponsePxgrid        `json:"pxgrid,omitempty"`          // The pxGrid probe fetches attributes of MAC address or IP address as a subscriber from the pxGrid queue.
	Radius          *[]ResponseNodeServicesGetProfilerProbeConfigResponseRadius        `json:"radius,omitempty"`          // The RADIUS probe collects RADIUS session attributes as well as CDP, LLDP, DHCP, HTTP, and MDM attributes from IOS Sensors.
	SNMPQuery       *ResponseNodeServicesGetProfilerProbeConfigResponseSNMPQuery       `json:"snmpQuery,omitempty"`       // The SNMP query probe collects details from network devices such as interface, CDP, LLDP, and ARP.
	SNMPTrap        *ResponseNodeServicesGetProfilerProbeConfigResponseSNMPTrap        `json:"snmpTrap,omitempty"`        // The SNMP trap probe receives linkup, linkdown, and MAC notification traps from network devices.
}

type ResponseNodeServicesGetProfilerProbeConfigResponseActiveDirectory struct {
	DaysBeforeRescan *int `json:"daysBeforeRescan,omitempty"` //
}

type ResponseNodeServicesGetProfilerProbeConfigResponseDhcp struct {
	Interfaces *[]ResponseNodeServicesGetProfilerProbeConfigResponseDhcpInterfaces `json:"interfaces,omitempty"` //
	Port       *int                                                                `json:"port,omitempty"`       //
}

type ResponseNodeServicesGetProfilerProbeConfigResponseDhcpInterfaces struct {
	Interface string `json:"interface,omitempty"` //
}

type ResponseNodeServicesGetProfilerProbeConfigResponseDhcpSpan struct {
	Interfaces *[]ResponseNodeServicesGetProfilerProbeConfigResponseDhcpSpanInterfaces `json:"interfaces,omitempty"` //
}

type ResponseNodeServicesGetProfilerProbeConfigResponseDhcpSpanInterfaces struct {
	Interface string `json:"interface,omitempty"` //
}

type ResponseNodeServicesGetProfilerProbeConfigResponseDNS struct {
	Timeout *int `json:"timeout,omitempty"` //
}

type ResponseNodeServicesGetProfilerProbeConfigResponseHTTP struct {
	Interfaces *[]ResponseNodeServicesGetProfilerProbeConfigResponseHTTPInterfaces `json:"interfaces,omitempty"` //
}

type ResponseNodeServicesGetProfilerProbeConfigResponseHTTPInterfaces struct {
	Interface string `json:"interface,omitempty"` //
}

type ResponseNodeServicesGetProfilerProbeConfigResponseNetflow struct {
	Interfaces *[]ResponseNodeServicesGetProfilerProbeConfigResponseNetflowInterfaces `json:"interfaces,omitempty"` //
	Port       *int                                                                   `json:"port,omitempty"`       //
}

type ResponseNodeServicesGetProfilerProbeConfigResponseNetflowInterfaces struct {
	Interface string `json:"interface,omitempty"` //
}

type ResponseNodeServicesGetProfilerProbeConfigResponseNmap interface{}

type ResponseNodeServicesGetProfilerProbeConfigResponsePxgrid interface{}

type ResponseNodeServicesGetProfilerProbeConfigResponseRadius interface{}

type ResponseNodeServicesGetProfilerProbeConfigResponseSNMPQuery struct {
	EventTimeout *int `json:"eventTimeout,omitempty"` //
	Retries      *int `json:"retries,omitempty"`      //
	Timeout      *int `json:"timeout,omitempty"`      //
}

type ResponseNodeServicesGetProfilerProbeConfigResponseSNMPTrap struct {
	Interfaces    *[]ResponseNodeServicesGetProfilerProbeConfigResponseSNMPTrapInterfaces `json:"interfaces,omitempty"`    //
	LinkTrapQuery *bool                                                                   `json:"linkTrapQuery,omitempty"` //
	MacTrapQuery  *bool                                                                   `json:"macTrapQuery,omitempty"`  //
	Port          *int                                                                    `json:"port,omitempty"`          //
}

type ResponseNodeServicesGetProfilerProbeConfigResponseSNMPTrapInterfaces struct {
	Interface string `json:"interface,omitempty"` //
}

type ResponseNodeServicesSetProfilerProbeConfig struct {
	Success *ResponseNodeServicesSetProfilerProbeConfigSuccess `json:"success,omitempty"` //
	Version string                                             `json:"version,omitempty"` //
}

type ResponseNodeServicesSetProfilerProbeConfigSuccess struct {
	Message string `json:"message,omitempty"` //
}

type RequestNodeServicesSetSxpInterface struct {
	Interface string `json:"interface,omitempty"` //
}

type RequestNodeServicesSetProfilerProbeConfig struct {
	ActiveDirectory *RequestNodeServicesSetProfilerProbeConfigActiveDirectory `json:"activeDirectory,omitempty"` // The Active Directory probe queries the Active Directory for Windows information.
	Dhcp            *RequestNodeServicesSetProfilerProbeConfigDhcp            `json:"dhcp,omitempty"`            // The DHCP probe listens for DHCP packets from IP helpers.
	DhcpSpan        *RequestNodeServicesSetProfilerProbeConfigDhcpSpan        `json:"dhcpSpan,omitempty"`        // The DHCP SPAN probe collects DHCP packets.
	DNS             *RequestNodeServicesSetProfilerProbeConfigDNS             `json:"dns,omitempty"`             // The DNS probe performs a DNS lookup for the FQDN.
	HTTP            *RequestNodeServicesSetProfilerProbeConfigHTTP            `json:"http,omitempty"`            // The HTTP probe receives and parses HTTP packets.
	Netflow         *RequestNodeServicesSetProfilerProbeConfigNetflow         `json:"netflow,omitempty"`         // The NetFlow probe collects the NetFlow packets that are sent to it from routers.
	Nmap            *[]RequestNodeServicesSetProfilerProbeConfigNmap          `json:"nmap,omitempty"`            // The NMAP probe scans endpoints for open ports and OS.
	Pxgrid          *[]RequestNodeServicesSetProfilerProbeConfigPxgrid        `json:"pxgrid,omitempty"`          // The pxGrid probe fetches attributes of MAC address or IP address as a subscriber from the pxGrid queue.
	Radius          *[]RequestNodeServicesSetProfilerProbeConfigRadius        `json:"radius,omitempty"`          // The RADIUS probe collects RADIUS session attributes as well as CDP, LLDP, DHCP, HTTP, and MDM attributes from IOS Sensors.
	SNMPQuery       *RequestNodeServicesSetProfilerProbeConfigSNMPQuery       `json:"snmpQuery,omitempty"`       // The SNMP query probe collects details from network devices such as interface, CDP, LLDP, and ARP.
	SNMPTrap        *RequestNodeServicesSetProfilerProbeConfigSNMPTrap        `json:"snmpTrap,omitempty"`        // The SNMP trap probe receives linkup, linkdown, and MAC notification traps from network devices.
}

type RequestNodeServicesSetProfilerProbeConfigActiveDirectory struct {
	DaysBeforeRescan *int `json:"daysBeforeRescan,omitempty"` //
}

type RequestNodeServicesSetProfilerProbeConfigDhcp struct {
	Interfaces *[]RequestNodeServicesSetProfilerProbeConfigDhcpInterfaces `json:"interfaces,omitempty"` //
	Port       *int                                                       `json:"port,omitempty"`       //
}

type RequestNodeServicesSetProfilerProbeConfigDhcpInterfaces struct {
	Interface string `json:"interface,omitempty"` //
}

type RequestNodeServicesSetProfilerProbeConfigDhcpSpan struct {
	Interfaces *[]RequestNodeServicesSetProfilerProbeConfigDhcpSpanInterfaces `json:"interfaces,omitempty"` //
}

type RequestNodeServicesSetProfilerProbeConfigDhcpSpanInterfaces struct {
	Interface string `json:"interface,omitempty"` //
}

type RequestNodeServicesSetProfilerProbeConfigDNS struct {
	Timeout *int `json:"timeout,omitempty"` //
}

type RequestNodeServicesSetProfilerProbeConfigHTTP struct {
	Interfaces *[]RequestNodeServicesSetProfilerProbeConfigHTTPInterfaces `json:"interfaces,omitempty"` //
}

type RequestNodeServicesSetProfilerProbeConfigHTTPInterfaces struct {
	Interface string `json:"interface,omitempty"` //
}

type RequestNodeServicesSetProfilerProbeConfigNetflow struct {
	Interfaces *[]RequestNodeServicesSetProfilerProbeConfigNetflowInterfaces `json:"interfaces,omitempty"` //
	Port       *int                                                          `json:"port,omitempty"`       //
}

type RequestNodeServicesSetProfilerProbeConfigNetflowInterfaces struct {
	Interface string `json:"interface,omitempty"` //
}

type RequestNodeServicesSetProfilerProbeConfigNmap interface{}

type RequestNodeServicesSetProfilerProbeConfigPxgrid interface{}

type RequestNodeServicesSetProfilerProbeConfigRadius interface{}

type RequestNodeServicesSetProfilerProbeConfigSNMPQuery struct {
	EventTimeout *int `json:"eventTimeout,omitempty"` //
	Retries      *int `json:"retries,omitempty"`      //
	Timeout      *int `json:"timeout,omitempty"`      //
}

type RequestNodeServicesSetProfilerProbeConfigSNMPTrap struct {
	Interfaces    *[]RequestNodeServicesSetProfilerProbeConfigSNMPTrapInterfaces `json:"interfaces,omitempty"`    //
	LinkTrapQuery *bool                                                          `json:"linkTrapQuery,omitempty"` //
	MacTrapQuery  *bool                                                          `json:"macTrapQuery,omitempty"`  //
	Port          *int                                                           `json:"port,omitempty"`          //
}

type RequestNodeServicesSetProfilerProbeConfigSNMPTrapInterfaces struct {
	Interface string `json:"interface,omitempty"` //
}

//GetInterfaces Get the list of interfaces on a node in a cluster.
/* This API retrieves the list of interfaces on a node in a cluster.

@param hostname hostname path parameter. Hostname of the node.
*/
func (s *NodeServicesService) GetInterfaces(hostname string) (*ResponseNodeServicesGetInterfaces, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/node/{hostname}/interface"
	path = strings.Replace(path, "{hostname}", fmt.Sprintf("%v", hostname), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeServicesGetInterfaces{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetInterfaces")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeServicesGetInterfaces)
	return result, response, err

}

//GetSxpInterface Get the interface configured for SXP.
/* This API retrieves the SXP interface.

@param hostname hostname path parameter. Hostname of the node.
*/
func (s *NodeServicesService) GetSxpInterface(hostname string) (*ResponseNodeServicesGetSxpInterface, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/node/{hostname}/sxp-interface"
	path = strings.Replace(path, "{hostname}", fmt.Sprintf("%v", hostname), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeServicesGetSxpInterface{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSxpInterface")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeServicesGetSxpInterface)
	return result, response, err

}

//GetProfilerProbeConfig Retrieve the profiler probe configuration of a PSN.
/* This API retrieves the profiler probe configuration of a PSN.

@param hostname hostname path parameter. Hostname of the node.
*/
func (s *NodeServicesService) GetProfilerProbeConfig(hostname string) (*ResponseNodeServicesGetProfilerProbeConfig, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/profile/{hostname}"
	path = strings.Replace(path, "{hostname}", fmt.Sprintf("%v", hostname), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeServicesGetProfilerProbeConfig{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetProfilerProbeConfig")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeServicesGetProfilerProbeConfig)
	return result, response, err

}

//SetSxpInterface Configure the interface for use with SXP.
/* This API configures the SXP interface.

@param hostname hostname path parameter. Hostname of the node.
*/
func (s *NodeServicesService) SetSxpInterface(hostname string, requestNodeServicesSetSxpInterface *RequestNodeServicesSetSxpInterface) (*ResponseNodeServicesSetSxpInterface, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/node/{hostname}/sxp-interface"
	path = strings.Replace(path, "{hostname}", fmt.Sprintf("%v", hostname), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNodeServicesSetSxpInterface).
		SetResult(&ResponseNodeServicesSetSxpInterface{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNodeServicesSetSxpInterface{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation SetSxpInterface")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeServicesSetSxpInterface)
	return result, response, err

}

//SetProfilerProbeConfig Update the profiler probe configuration of a PSN.
/* This API updates the profiler probe configuration of a PSN.
Set probe value as
null
 to disable probe.
Ex: Below payload will disable NMAP, PxGrid and SNMPTRAP probes
{
  "activeDirectory": { "daysBeforeRescan": 1 },
  "dhcp": { "interfaces": "[{"interface":"GigabitEthernet 0"}]", "port": 0 },
  "dhcpSpan": { "interfaces": "[{"interface":"GigabitEthernet 0"}]" },
  "dns": { "timeout": 2 },
  "http": { "interfaces": "[{"interface":"GigabitEthernet 0"}]" },
  "netflow": { "interfaces": "[{"interface":"GigabitEthernet 0"}]", "port": 0 },
  "nmap":
null
,
  "pxgrid":
null
,
  "radius": [],
  "snmpQuery": { "eventTimeout": 30, "retries": 2, "timeout": 1000 },
  "snmpTrap":
null

}


@param hostname hostname path parameter. Hostname of the node.
*/
func (s *NodeServicesService) SetProfilerProbeConfig(hostname string, requestNodeServicesSetProfilerProbeConfig *RequestNodeServicesSetProfilerProbeConfig) (*ResponseNodeServicesSetProfilerProbeConfig, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/profile/{hostname}"
	path = strings.Replace(path, "{hostname}", fmt.Sprintf("%v", hostname), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNodeServicesSetProfilerProbeConfig).
		SetResult(&ResponseNodeServicesSetProfilerProbeConfig{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNodeServicesSetProfilerProbeConfig{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation SetProfilerProbeConfig")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeServicesSetProfilerProbeConfig)
	return result, response, err

}
