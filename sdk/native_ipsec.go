package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type NativeIPsecService service

type GetIPsecEnabledNodesQueryParams struct {
	Page       int    `url:"page,omitempty"`       //Page number
	Size       int    `url:"size,omitempty"`       //Number of objects returned per page
	Filter     string `url:"filter,omitempty"`     //<div> <style type="text/css" scoped> .apiServiceTable td, .apiServiceTable th { padding: 5px 10px !important; text-align: left; } </style> <span> <b>Simple filtering</b> should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the <i>"filterType=or"</i> query string parameter. Each resource Data model description should specify if an attribute is a filtered field. </span> <br /> <br/> <table class="apiServiceTable"> <thead> <tr> <th>OPERATOR</th> <th>DESCRIPTION</th> <th>APPLICABLE ON FIELDS<th> </tr> </thead> <tbody> <tr> <td>EQ</td> <td>Equals</td> <td>authType</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> <td>authType</td> </tr> <tr> <td>EQ</td> <td>Equals</td> <td>hostName</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> <td>hostName</td> </tr> <tr> <td>EQ</td> <td>Equals</td> <td>nadIp</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> <td>nadIp</td> </tr> <tr> <td>EQ</td> <td>Equals</td> <td>status</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> <td>status</td> </tr> </tbody> </table> </div>
	FilterType string `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
	Sort       string `url:"sort,omitempty"`       //sort type - asc or desc
	SortBy     string `url:"sortBy,omitempty"`     //Sort column -  The IPsec enabled nodes are sorted based on the columns. This is applicable for the field - hostName.
}

type ResponseNativeIPsecGetIPsecEnabledNodes struct {
	NextPage     *ResponseNativeIPsecGetIPsecEnabledNodesNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseNativeIPsecGetIPsecEnabledNodesPreviousPage `json:"previousPage,omitempty"` //
	Response     *[]ResponseNativeIPsecGetIPsecEnabledNodesResponse   `json:"response,omitempty"`     //
	Version      string                                               `json:"version,omitempty"`      //
}

type ResponseNativeIPsecGetIPsecEnabledNodesNextPage struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNativeIPsecGetIPsecEnabledNodesPreviousPage struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNativeIPsecGetIPsecEnabledNodesResponse struct {
	AuthType               string `json:"authType,omitempty"`               // Authentication type for establishing connection
	CertID                 string `json:"certId,omitempty"`                 // ID of the certificate for establishing connection
	ConfigureVti           *bool  `json:"configureVti,omitempty"`           // Authentication type for establishing connection
	CreateTime             string `json:"createTime,omitempty"`             //
	EspAhProtocol          string `json:"espAhProtocol,omitempty"`          // Encryption protocol used for establishing connection
	HostName               string `json:"hostName,omitempty"`               // Hostname of the node
	ID                     string `json:"id,omitempty"`                     //
	Iface                  string `json:"iface,omitempty"`                  // Ethernet port of the node
	IkeReAuthTime          *int   `json:"ikeReAuthTime,omitempty"`          // IKE re-authentication time
	IkeVersion             string `json:"ikeVersion,omitempty"`             // IKE version
	LocalInternalIP        string `json:"localInternalIp,omitempty"`        // Local Tunnel IP address
	ModeOption             string `json:"modeOption,omitempty"`             // The Mode type used for establishing the connection
	NadIP                  string `json:"nadIp,omitempty"`                  // NAD IP address for establishing connection
	PhaseOneDHGroup        string `json:"phaseOneDHGroup,omitempty"`        // Phase-one DH group used for establishing connection
	PhaseOneEncryptionAlgo string `json:"phaseOneEncryptionAlgo,omitempty"` // Phase-one encryption algorithm used for establishing connection
	PhaseOneHashAlgo       string `json:"phaseOneHashAlgo,omitempty"`       // Phase-one hashing algorithm used for establishing connection
	PhaseOneLifeTime       *int   `json:"phaseOneLifeTime,omitempty"`       // Phase-one connection lifetime
	PhaseTwoDHGroup        string `json:"phaseTwoDHGroup,omitempty"`        // Phase-two DH group used for establishing connection
	PhaseTwoEncryptionAlgo string `json:"phaseTwoEncryptionAlgo,omitempty"` // Phase-two encryption algorithm used for establishing connection
	PhaseTwoHashAlgo       string `json:"phaseTwoHashAlgo,omitempty"`       // Phase-two hashing algorithm used for establishing connection
	PhaseTwoLifeTime       *int   `json:"phaseTwoLifeTime,omitempty"`       // Phase-two connection lifetime
	Psk                    string `json:"psk,omitempty"`                    // Pre-shared key used for establishing connection
	RemotePeerInternalIP   string `json:"remotePeerInternalIp,omitempty"`   // Remote Tunnel IP address
	Status                 string `json:"status,omitempty"`                 //
	UpdateTime             string `json:"updateTime,omitempty"`             //
}

type ResponseNativeIPsecUpdateIPsecConnectionConfig struct {
	Response *ResponseNativeIPsecUpdateIPsecConnectionConfigResponse `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  //
}

type ResponseNativeIPsecUpdateIPsecConnectionConfigResponse struct {
	AuthType               string                                                      `json:"authType,omitempty"`               // Authentication type for establishing connection
	CertID                 string                                                      `json:"certId,omitempty"`                 // ID of the certificate for establishing connection
	ConfigureVti           *bool                                                       `json:"configureVti,omitempty"`           // Authentication type for establishing connection
	CreateTime             string                                                      `json:"createTime,omitempty"`             //
	EspAhProtocol          string                                                      `json:"espAhProtocol,omitempty"`          // Encryption protocol used for establishing connection
	HostName               string                                                      `json:"hostName,omitempty"`               // Hostname of the node
	ID                     string                                                      `json:"id,omitempty"`                     //
	Iface                  string                                                      `json:"iface,omitempty"`                  // Ethernet port of the node
	IkeReAuthTime          *int                                                        `json:"ikeReAuthTime,omitempty"`          // IKE re-authentication time
	IkeVersion             string                                                      `json:"ikeVersion,omitempty"`             // IKE version
	Link                   *ResponseNativeIPsecUpdateIPsecConnectionConfigResponseLink `json:"link,omitempty"`                   //
	LocalInternalIP        string                                                      `json:"localInternalIp,omitempty"`        // Local Tunnel IP address
	ModeOption             string                                                      `json:"modeOption,omitempty"`             // The Mode type used for establishing the connection
	NadIP                  string                                                      `json:"nadIp,omitempty"`                  // NAD IP address for establishing connection
	PhaseOneDHGroup        string                                                      `json:"phaseOneDHGroup,omitempty"`        // Phase-one DH group used for establishing connection
	PhaseOneEncryptionAlgo string                                                      `json:"phaseOneEncryptionAlgo,omitempty"` // Phase-one encryption algorithm used for establishing connection
	PhaseOneHashAlgo       string                                                      `json:"phaseOneHashAlgo,omitempty"`       // Phase-one hashing algorithm used for establishing connection
	PhaseOneLifeTime       *int                                                        `json:"phaseOneLifeTime,omitempty"`       // Phase-one connection lifetime
	PhaseTwoDHGroup        string                                                      `json:"phaseTwoDHGroup,omitempty"`        // Phase-two DH group used for establishing connection
	PhaseTwoEncryptionAlgo string                                                      `json:"phaseTwoEncryptionAlgo,omitempty"` // Phase-two encryption algorithm used for establishing connection
	PhaseTwoHashAlgo       string                                                      `json:"phaseTwoHashAlgo,omitempty"`       // Phase-two hashing algorithm used for establishing connection
	PhaseTwoLifeTime       *int                                                        `json:"phaseTwoLifeTime,omitempty"`       // Phase-two connection lifetime
	Psk                    string                                                      `json:"psk,omitempty"`                    // Pre-shared key used for establishing connection
	RemotePeerInternalIP   string                                                      `json:"remotePeerInternalIp,omitempty"`   // Remote Tunnel IP address
	Status                 string                                                      `json:"status,omitempty"`                 //
	UpdateTime             string                                                      `json:"updateTime,omitempty"`             //
}

type ResponseNativeIPsecUpdateIPsecConnectionConfigResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNativeIPsecCreateIPsecConnection struct {
	Response *ResponseNativeIPsecCreateIPsecConnectionResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}

type ResponseNativeIPsecCreateIPsecConnectionResponse struct {
	AuthType               string                                                `json:"authType,omitempty"`               // Authentication type for establishing connection
	CertID                 string                                                `json:"certId,omitempty"`                 // ID of the certificate for establishing connection
	ConfigureVti           *bool                                                 `json:"configureVti,omitempty"`           // Authentication type for establishing connection
	CreateTime             string                                                `json:"createTime,omitempty"`             //
	EspAhProtocol          string                                                `json:"espAhProtocol,omitempty"`          // Encryption protocol used for establishing connection
	HostName               string                                                `json:"hostName,omitempty"`               // Hostname of the node
	ID                     string                                                `json:"id,omitempty"`                     //
	Iface                  string                                                `json:"iface,omitempty"`                  // Ethernet port of the node
	IkeReAuthTime          *int                                                  `json:"ikeReAuthTime,omitempty"`          // IKE re-authentication time
	IkeVersion             string                                                `json:"ikeVersion,omitempty"`             // IKE version
	Link                   *ResponseNativeIPsecCreateIPsecConnectionResponseLink `json:"link,omitempty"`                   //
	LocalInternalIP        string                                                `json:"localInternalIp,omitempty"`        // Local Tunnel IP address
	ModeOption             string                                                `json:"modeOption,omitempty"`             // The Mode type used for establishing the connection
	NadIP                  string                                                `json:"nadIp,omitempty"`                  // NAD IP address for establishing connection
	PhaseOneDHGroup        string                                                `json:"phaseOneDHGroup,omitempty"`        // Phase-one DH group used for establishing connection
	PhaseOneEncryptionAlgo string                                                `json:"phaseOneEncryptionAlgo,omitempty"` // Phase-one encryption algorithm used for establishing connection
	PhaseOneHashAlgo       string                                                `json:"phaseOneHashAlgo,omitempty"`       // Phase-one hashing algorithm used for establishing connection
	PhaseOneLifeTime       *int                                                  `json:"phaseOneLifeTime,omitempty"`       // Phase-one connection lifetime
	PhaseTwoDHGroup        string                                                `json:"phaseTwoDHGroup,omitempty"`        // Phase-two DH group used for establishing connection
	PhaseTwoEncryptionAlgo string                                                `json:"phaseTwoEncryptionAlgo,omitempty"` // Phase-two encryption algorithm used for establishing connection
	PhaseTwoHashAlgo       string                                                `json:"phaseTwoHashAlgo,omitempty"`       // Phase-two hashing algorithm used for establishing connection
	PhaseTwoLifeTime       *int                                                  `json:"phaseTwoLifeTime,omitempty"`       // Phase-two connection lifetime
	Psk                    string                                                `json:"psk,omitempty"`                    // Pre-shared key used for establishing connection
	RemotePeerInternalIP   string                                                `json:"remotePeerInternalIp,omitempty"`   // Remote Tunnel IP address
	Status                 string                                                `json:"status,omitempty"`                 //
	UpdateTime             string                                                `json:"updateTime,omitempty"`             //
}

type ResponseNativeIPsecCreateIPsecConnectionResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNativeIPsecBulkIPSecOperation struct {
	ID string `json:"id,omitempty"` //
}

type ResponseNativeIPsecGetIPSecCertificates struct {
	Response *[]ResponseNativeIPsecGetIPSecCertificatesResponse `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  //
}

type ResponseNativeIPsecGetIPSecCertificatesResponse struct {
	FriendlyName string `json:"friendlyName,omitempty"` // Friendly name of system certificate
	ID           string `json:"id,omitempty"`           // ID of system certificate
}

type ResponseNativeIPsecDisableIPsecConnection struct {
	Message string `json:"message,omitempty"` //
}

type ResponseNativeIPsecEnableIPsecConnection struct {
	Message string `json:"message,omitempty"` //
}

type ResponseNativeIPsecGetIPsecNode struct {
	Response *ResponseNativeIPsecGetIPsecNodeResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}

type ResponseNativeIPsecGetIPsecNodeResponse struct {
	AuthType               string                                       `json:"authType,omitempty"`               // Authentication type for establishing connection
	CertID                 string                                       `json:"certId,omitempty"`                 // ID of the certificate for establishing connection
	ConfigureVti           *bool                                        `json:"configureVti,omitempty"`           // Authentication type for establishing connection
	CreateTime             string                                       `json:"createTime,omitempty"`             //
	EspAhProtocol          string                                       `json:"espAhProtocol,omitempty"`          // Encryption protocol used for establishing connection
	HostName               string                                       `json:"hostName,omitempty"`               // Hostname of the node
	ID                     string                                       `json:"id,omitempty"`                     //
	Iface                  string                                       `json:"iface,omitempty"`                  // Ethernet port of the node
	IkeReAuthTime          *int                                         `json:"ikeReAuthTime,omitempty"`          // IKE re-authentication time
	IkeVersion             string                                       `json:"ikeVersion,omitempty"`             // IKE version
	Link                   *ResponseNativeIPsecGetIPsecNodeResponseLink `json:"link,omitempty"`                   //
	LocalInternalIP        string                                       `json:"localInternalIp,omitempty"`        // Local Tunnel IP address
	ModeOption             string                                       `json:"modeOption,omitempty"`             // The Mode type used for establishing the connection
	NadIP                  string                                       `json:"nadIp,omitempty"`                  // NAD IP address for establishing connection
	PhaseOneDHGroup        string                                       `json:"phaseOneDHGroup,omitempty"`        // Phase-one DH group used for establishing connection
	PhaseOneEncryptionAlgo string                                       `json:"phaseOneEncryptionAlgo,omitempty"` // Phase-one encryption algorithm used for establishing connection
	PhaseOneHashAlgo       string                                       `json:"phaseOneHashAlgo,omitempty"`       // Phase-one hashing algorithm used for establishing connection
	PhaseOneLifeTime       *int                                         `json:"phaseOneLifeTime,omitempty"`       // Phase-one connection lifetime
	PhaseTwoDHGroup        string                                       `json:"phaseTwoDHGroup,omitempty"`        // Phase-two DH group used for establishing connection
	PhaseTwoEncryptionAlgo string                                       `json:"phaseTwoEncryptionAlgo,omitempty"` // Phase-two encryption algorithm used for establishing connection
	PhaseTwoHashAlgo       string                                       `json:"phaseTwoHashAlgo,omitempty"`       // Phase-two hashing algorithm used for establishing connection
	PhaseTwoLifeTime       *int                                         `json:"phaseTwoLifeTime,omitempty"`       // Phase-two connection lifetime
	Psk                    string                                       `json:"psk,omitempty"`                    // Pre-shared key used for establishing connection
	RemotePeerInternalIP   string                                       `json:"remotePeerInternalIp,omitempty"`   // Remote Tunnel IP address
	Status                 string                                       `json:"status,omitempty"`                 //
	UpdateTime             string                                       `json:"updateTime,omitempty"`             //
}

type ResponseNativeIPsecGetIPsecNodeResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNativeIPsecRemoveIPsecConnection struct {
	Message string `json:"message,omitempty"` //
}

type RequestNativeIPsecUpdateIPsecConnectionConfig struct {
	AuthType               string `json:"authType,omitempty"`               // Authentication type for establishing connection
	CertID                 string `json:"certId,omitempty"`                 // ID of the certificate for establishing connection
	ConfigureVti           *bool  `json:"configureVti,omitempty"`           // Authentication type for establishing connection
	EspAhProtocol          string `json:"espAhProtocol,omitempty"`          // Encryption protocol used for establishing connection
	HostName               string `json:"hostName,omitempty"`               // Hostname of the node
	Iface                  string `json:"iface,omitempty"`                  // Ethernet port of the node
	IkeReAuthTime          *int   `json:"ikeReAuthTime,omitempty"`          // IKE re-authentication time
	IkeVersion             string `json:"ikeVersion,omitempty"`             // IKE version
	LocalInternalIP        string `json:"localInternalIp,omitempty"`        // Local Tunnel IP address
	ModeOption             string `json:"modeOption,omitempty"`             // The Mode type used for establishing the connection
	NadIP                  string `json:"nadIp,omitempty"`                  // NAD IP address for establishing connection
	PhaseOneDHGroup        string `json:"phaseOneDHGroup,omitempty"`        // Phase-one DH group used for establishing connection
	PhaseOneEncryptionAlgo string `json:"phaseOneEncryptionAlgo,omitempty"` // Phase-one encryption algorithm used for establishing connection
	PhaseOneHashAlgo       string `json:"phaseOneHashAlgo,omitempty"`       // Phase-one hashing algorithm used for establishing connection
	PhaseOneLifeTime       *int   `json:"phaseOneLifeTime,omitempty"`       // Phase-one connection lifetime
	PhaseTwoDHGroup        string `json:"phaseTwoDHGroup,omitempty"`        // Phase-two DH group used for establishing connection
	PhaseTwoEncryptionAlgo string `json:"phaseTwoEncryptionAlgo,omitempty"` // Phase-two encryption algorithm used for establishing connection
	PhaseTwoHashAlgo       string `json:"phaseTwoHashAlgo,omitempty"`       // Phase-two hashing algorithm used for establishing connection
	PhaseTwoLifeTime       *int   `json:"phaseTwoLifeTime,omitempty"`       // Phase-two connection lifetime
	Psk                    string `json:"psk,omitempty"`                    // Pre-shared key used for establishing connection
	RemotePeerInternalIP   string `json:"remotePeerInternalIp,omitempty"`   // Remote Tunnel IP address
}

type RequestNativeIPsecCreateIPsecConnection struct {
	AuthType               string `json:"authType,omitempty"`               // Authentication type for establishing connection
	CertID                 string `json:"certId,omitempty"`                 // ID of the certificate for establishing connection
	ConfigureVti           *bool  `json:"configureVti,omitempty"`           // Authentication type for establishing connection
	EspAhProtocol          string `json:"espAhProtocol,omitempty"`          // Encryption protocol used for establishing connection
	HostName               string `json:"hostName,omitempty"`               // Hostname of the node
	Iface                  string `json:"iface,omitempty"`                  // Ethernet port of the node
	IkeReAuthTime          *int   `json:"ikeReAuthTime,omitempty"`          // IKE re-authentication time
	IkeVersion             string `json:"ikeVersion,omitempty"`             // IKE version
	LocalInternalIP        string `json:"localInternalIp,omitempty"`        // Local Tunnel IP address
	ModeOption             string `json:"modeOption,omitempty"`             // The Mode type used for establishing the connection
	NadIP                  string `json:"nadIp,omitempty"`                  // NAD IP address for establishing connection
	PhaseOneDHGroup        string `json:"phaseOneDHGroup,omitempty"`        // Phase-one DH group used for establishing connection
	PhaseOneEncryptionAlgo string `json:"phaseOneEncryptionAlgo,omitempty"` // Phase-one encryption algorithm used for establishing connection
	PhaseOneHashAlgo       string `json:"phaseOneHashAlgo,omitempty"`       // Phase-one hashing algorithm used for establishing connection
	PhaseOneLifeTime       *int   `json:"phaseOneLifeTime,omitempty"`       // Phase-one connection lifetime
	PhaseTwoDHGroup        string `json:"phaseTwoDHGroup,omitempty"`        // Phase-two DH group used for establishing connection
	PhaseTwoEncryptionAlgo string `json:"phaseTwoEncryptionAlgo,omitempty"` // Phase-two encryption algorithm used for establishing connection
	PhaseTwoHashAlgo       string `json:"phaseTwoHashAlgo,omitempty"`       // Phase-two hashing algorithm used for establishing connection
	PhaseTwoLifeTime       *int   `json:"phaseTwoLifeTime,omitempty"`       // Phase-two connection lifetime
	Psk                    string `json:"psk,omitempty"`                    // Pre-shared key used for establishing connection
	RemotePeerInternalIP   string `json:"remotePeerInternalIp,omitempty"`   // Remote Tunnel IP address
}

type RequestNativeIPsecBulkIPSecOperation struct {
	ItemList  *[]RequestNativeIPsecBulkIPSecOperationItemList `json:"ItemList,omitempty"`  //
	Operation string                                          `json:"operation,omitempty"` //
}

type RequestNativeIPsecBulkIPSecOperationItemList struct {
	AuthType               string `json:"authType,omitempty"`               // Authentication type for establishing connection
	CertID                 string `json:"certId,omitempty"`                 // ID of the certificate for establishing connection
	ConfigureVti           *bool  `json:"configureVti,omitempty"`           // Authentication type for establishing connection
	EspAhProtocol          string `json:"espAhProtocol,omitempty"`          // Encryption protocol used for establishing connection
	HostName               string `json:"hostName,omitempty"`               // Hostname of the node
	Iface                  string `json:"iface,omitempty"`                  // Ethernet port of the node
	IkeReAuthTime          *int   `json:"ikeReAuthTime,omitempty"`          // IKE re-authentication time
	IkeVersion             string `json:"ikeVersion,omitempty"`             // IKE version
	LocalInternalIP        string `json:"localInternalIp,omitempty"`        // Local Tunnel IP address
	ModeOption             string `json:"modeOption,omitempty"`             // The Mode type used for establishing the connection
	NadIP                  string `json:"nadIp,omitempty"`                  // NAD IP address for establishing connection
	PhaseOneDHGroup        string `json:"phaseOneDHGroup,omitempty"`        // Phase-one DH group used for establishing connection
	PhaseOneEncryptionAlgo string `json:"phaseOneEncryptionAlgo,omitempty"` // Phase-one encryption algorithm used for establishing connection
	PhaseOneHashAlgo       string `json:"phaseOneHashAlgo,omitempty"`       // Phase-one hashing algorithm used for establishing connection
	PhaseOneLifeTime       *int   `json:"phaseOneLifeTime,omitempty"`       // Phase-one connection lifetime
	PhaseTwoDHGroup        string `json:"phaseTwoDHGroup,omitempty"`        // Phase-two DH group used for establishing connection
	PhaseTwoEncryptionAlgo string `json:"phaseTwoEncryptionAlgo,omitempty"` // Phase-two encryption algorithm used for establishing connection
	PhaseTwoHashAlgo       string `json:"phaseTwoHashAlgo,omitempty"`       // Phase-two hashing algorithm used for establishing connection
	PhaseTwoLifeTime       *int   `json:"phaseTwoLifeTime,omitempty"`       // Phase-two connection lifetime
	Psk                    string `json:"psk,omitempty"`                    // Pre-shared key used for establishing connection
	RemotePeerInternalIP   string `json:"remotePeerInternalIp,omitempty"`   // Remote Tunnel IP address
}

//GetIPsecEnabledNodes Get all IPsec enabled nodes
/*  Returns all the IPsec enabled nodes with configuration details.

 This API supports filtering, sorting and pagination.

The attributes that are suppported for filtering are:


hostName

nadIp

status

authType


The attribute that is suppported for sorting is:


hostName



@param getIpsecEnabledNodesQueryParams Filtering parameter
*/
func (s *NativeIPsecService) GetIPsecEnabledNodes(getIpsecEnabledNodesQueryParams *GetIPsecEnabledNodesQueryParams) (*ResponseNativeIPsecGetIPsecEnabledNodes, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/ipsec"

	queryString, _ := query.Values(getIpsecEnabledNodesQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNativeIPsecGetIPsecEnabledNodes{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIpsecEnabledNodes")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNativeIPsecGetIPsecEnabledNodes)
	return result, response, err

}

//GetIPSecCertificates Get all IPsec related certificates
/*  Returns all the certificates for IPsec role.


 */
func (s *NativeIPsecService) GetIPSecCertificates() (*ResponseNativeIPsecGetIPSecCertificates, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/ipsec/certificates"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNativeIPsecGetIPSecCertificates{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIpSecCertificates")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNativeIPsecGetIPSecCertificates)
	return result, response, err

}

//GetIPsecNode Get the IPsec connection details for a given node with the hostname and the NAD IP
/* Returns the IPsec configuration details of a given node with the hostname and the NAD IP.

@param hostName hostName path parameter. Hostname of the deployed node.
@param nadIP nadIp path parameter. IP address of the NAD.
*/
func (s *NativeIPsecService) GetIPsecNode(hostName string, nadIP string) (*ResponseNativeIPsecGetIPsecNode, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/ipsec/{hostName}/{nadIp}"
	path = strings.Replace(path, "{hostName}", fmt.Sprintf("%v", hostName), -1)
	path = strings.Replace(path, "{nadIp}", fmt.Sprintf("%v", nadIP), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNativeIPsecGetIPsecNode{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIpsecNode")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNativeIPsecGetIPsecNode)
	return result, response, err

}

//CreateIPsecConnection Create an IPsec connection on a node
/*
Creates an IPsec connection.
 The following parameters are present in the POST request body:




PARAMETER

DESCRIPTION

EXAMPLE





hostName
*required

Hostname of the node for which IPsec should be enabled

"hostName": "ise-host1"



iface
*required

Ethernet port used for establishing connection

"iface": "0"



psk
*required

Pre-shared key used for establishing connection.

"psk": "psk12345"



authType
*required

Pre-shared key used for establishing connection.

"authType": "psk"



configureVti

Used For VTI Configurations

"configureVti": "false"



remotePeerInternalIp

VTI Internal IP of the NAD

"remotePeerInternalIp": "1.2.3.1"



localInternalIp

IP address assigned to the VTI interface so this would be the internal ip

"localInternalIp": "1.1.3.1"



certId
*required

ID of the certificate for establishing connection.

"certId": "21323243545433"



phaseOneEncryptionAlgo
*required

Phase-one encryption algorithm used for establishing connection.

"phaseOneEncryptionAlgo": "aes"



phaseTwoEncryptionAlgo
*required

Phase-two encryption algorithm used for establishing connection.

"phaseTwoEncryptionAlgo": "aes"



espAhProtocol
*required

Encryption protocol used for establishing connection.

"espAhProtocol": "ah"



phaseOneHashAlgo
*required

Phase-one hashing algorithm used for establishing connection.

"phaseOneHashAlgo": "sha"



phaseTwoHashAlgo
*required

Phase-two hashing algorithm used for establishing connection.

"phaseTwoHashAlgo": "sha"



phaseOneDHGroup
*required

Phase-one DH group used for establishing connection.

"phaseOneDHGroup": "GROUP1"



phaseTwoDHGroup

Phase-two DH group used for establishing connection.

"phaseTwoDHGroup": "GROUP1"



phaseOneLifeTime

DH Phase-one connection lifetime.

"phaseOneLifeTime": 14400



phaseTwoLifeTime

DH Phase-two connection lifetime.

"phaseTwoLifeTime": 14400



ikeVersion
*required

IKE version.

"ikeVersion": "1"



ikeReAuthTime

IKE re-authentication time.

"ikeReAuthTime": 86400



nadIp
*required

NAD IP for establishing the connection.

"nadIp": "1.1.1.1"



modeOption
*required

The Mode type used for establishing the connection.

"modeOption": "tunnel"




NOTE:

psk
field is mandatory if authType=psk
certId
field is mandatory if authType=x509

If FIPS mode is on.:


Cannot choose DES or 3DES for Phase-one and Phase-two Encryption algorithms.

PSK length must be 14 characters or more.

DH Groups 1, 2, and 5 cannot be chosen for Phase-one and Phase-two fields.




*/
func (s *NativeIPsecService) CreateIPsecConnection(requestNativeIPsecCreateIpsecConnection *RequestNativeIPsecCreateIPsecConnection) (*ResponseNativeIPsecCreateIPsecConnection, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/ipsec"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNativeIPsecCreateIpsecConnection).
		SetResult(&ResponseNativeIPsecCreateIPsecConnection{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNativeIPsecCreateIPsecConnection{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateIpsecConnection")
	}

	result := response.Result().(*ResponseNativeIPsecCreateIPsecConnection)
	return result, response, err

}

//BulkIPSecOperation Create, update, disable, enable and remove IPsec connections in bulk
/* Create, update, disable, enable and remove IPsec connections in bulk

 */
func (s *NativeIPsecService) BulkIPSecOperation(requestNativeIPsecBulkIPSecOperation *RequestNativeIPsecBulkIPSecOperation) (*ResponseNativeIPsecBulkIPSecOperation, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/ipsec/bulk"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNativeIPsecBulkIPSecOperation).
		SetResult(&ResponseNativeIPsecBulkIPSecOperation{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNativeIPsecBulkIPSecOperation{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BulkIpSecOperation")
	}

	result := response.Result().(*ResponseNativeIPsecBulkIPSecOperation)
	return result, response, err

}

//UpdateIPsecConnectionConfig Update the configuration of an existing IPsec connection on a node
/*
Updates the configuration of existing IPsec connection.
 The following parameters are present in the PUT request body:




PARAMETER

DESCRIPTION

EXAMPLE





id
*required

ID of the existing IPsec configuration.

"id": "7c9484cf-0ebc-47ad-a9ef-bc12729ed73b"



iface
*required

Ethernet port used for establishing connection

"iface": "0"



psk
*required

Pre-shared key used for establishing connection.

"psk": "psk12345"



authType
*required

Pre-shared key used for establishing connection.

"authType": "psk"



configureVti

Used For VTI Configurations

"configureVti": "false"



remotePeerInternalIp

VTI Internal IP of the NAD

"remotePeerInternalIp": "1.2.3.1"



localInternalIp

IP address assigned to the VTI interface so this would be the internal ip

"localInternalIp": "1.1.3.1"



certId
*required

ID of the certificate for establishing connection.

"certId": "21323243545433"



phaseOneEncryptionAlgo
*required

Phase-one encryption algorithm used for establishing connection.

"phaseOneEncryptionAlgo": "aes"



phaseTwoEncryptionAlgo
*required

Phase-two encryption algorithm used for establishing connection.

"phaseTwoEncryptionAlgo": "aes"



espAhProtocol
*required

Encryption protocol used for establishing connection.

"espAhProtocol": "ah"



phaseOneHashAlgo
*required

Phase-one hashing algorithm used for establishing connection.

"phaseOneHashAlgo": "sha"



phaseTwoHashAlgo
*required

Phase-two hashing algorithm used for establishing connection.

"phaseTwoHashAlgo": "sha"



phaseOneDHGroup
*required

Phase-one DH group used for establishing connection.

"phaseOneDHGroup": "GROUP1"



phaseTwoDHGroup

Phase-two DH group used for establishing connection.

"phaseTwoDHGroup": "GROUP1"



phaseOneLifeTime

DH Phase-one connection lifetime.

"phaseOneLifeTime": 14400



phaseTwoLifeTime

DH Phase-two connection lifetime.

"phaseTwoLifeTime": 14400



ikeVersion
*required

IKE version.

"ikeVersion": "1"



ikeReAuthTime

IKE re-authentication time.

"ikeReAuthTime": 86400



nadIp
*required

NAD IP for establishing connection.

"nadIp": "1.1.1.1"



modeOption
*required

The Mode type used for establishing the connection.

"modeOption": "tunnel"




NOTE:

psk
field is mandatory if authType=psk
certId
field is mandatory if authType=x509

If FIPS mode is on.:


Cannot choose DES or 3DES for Phase-one and Phase-two Encryption algorithms.

PSK length must be 14 characters or more.

DH Groups 1, 2, and 5 cannot be chosen for Phase-one and Phase-two fields.




*/
func (s *NativeIPsecService) UpdateIPsecConnectionConfig(requestNativeIPsecUpdateIpsecConnectionConfig *RequestNativeIPsecUpdateIPsecConnectionConfig) (*ResponseNativeIPsecUpdateIPsecConnectionConfig, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/ipsec"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNativeIPsecUpdateIpsecConnectionConfig).
		SetResult(&ResponseNativeIPsecUpdateIPsecConnectionConfig{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNativeIPsecUpdateIPsecConnectionConfig{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateIpsecConnectionConfig")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNativeIPsecUpdateIPsecConnectionConfig)
	return result, response, err

}

//DisableIPsecConnection Disable the IPsec connection on a node for a given hostname and NAD IP
/* Disables an enabled IPsec node connection.

@param hostName hostName path parameter. Hostname of the deployed node.
@param nadIP nadIp path parameter. IP address of the NAD.
*/
func (s *NativeIPsecService) DisableIPsecConnection(hostName string, nadIP string) (*ResponseNativeIPsecDisableIPsecConnection, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/ipsec/disable/{hostName}/{nadIp}"
	path = strings.Replace(path, "{hostName}", fmt.Sprintf("%v", hostName), -1)
	path = strings.Replace(path, "{nadIp}", fmt.Sprintf("%v", nadIP), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNativeIPsecDisableIPsecConnection{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNativeIPsecDisableIPsecConnection{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DisableIpsecConnection")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNativeIPsecDisableIPsecConnection)
	return result, response, err

}

//EnableIPsecConnection Enable the IPsec connection on a node for a given hostname and NAD IP
/* Enables an disabled IPsec node connection.

@param hostName hostName path parameter. Hostname of the deployed node.
@param nadIP nadIp path parameter. IP address of the NAD.
*/
func (s *NativeIPsecService) EnableIPsecConnection(hostName string, nadIP string) (*ResponseNativeIPsecEnableIPsecConnection, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/ipsec/enable/{hostName}/{nadIp}"
	path = strings.Replace(path, "{hostName}", fmt.Sprintf("%v", hostName), -1)
	path = strings.Replace(path, "{nadIp}", fmt.Sprintf("%v", nadIP), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNativeIPsecEnableIPsecConnection{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNativeIPsecEnableIPsecConnection{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation EnableIpsecConnection")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNativeIPsecEnableIPsecConnection)
	return result, response, err

}

//RemoveIPsecConnection Remove the IPsec connection on a node for a given hostname and NAD IP
/* Removes an enabled IPsec node connection.

@param hostName hostName path parameter. Hostname of the deployed node.
@param nadIP nadIp path parameter. IP address of the NAD.
*/
func (s *NativeIPsecService) RemoveIPsecConnection(hostName string, nadIP string) (*ResponseNativeIPsecRemoveIPsecConnection, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/ipsec/{hostName}/{nadIp}"
	path = strings.Replace(path, "{hostName}", fmt.Sprintf("%v", hostName), -1)
	path = strings.Replace(path, "{nadIp}", fmt.Sprintf("%v", nadIP), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNativeIPsecRemoveIPsecConnection{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RemoveIpsecConnection")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNativeIPsecRemoveIPsecConnection)
	return result, response, err

}
