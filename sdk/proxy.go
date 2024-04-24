package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ProxyService service

type ResponseProxyGetProxyConnection struct {
	Response *ResponseProxyGetProxyConnectionResponse `json:"response,omitempty"` // ISE Proxy connection settings
	Version  string                                   `json:"version,omitempty"`  //
}

type ResponseProxyGetProxyConnectionResponse struct {
	BypassHosts      string `json:"bypassHosts,omitempty"`      // Bypass hosts for the proxy connection
	Fqdn             string `json:"fqdn,omitempty"`             // proxy IP address or DNS-resolvable host name
	Password         string `json:"password,omitempty"`         // Password for the proxy connection
	PasswordRequired *bool  `json:"passwordRequired,omitempty"` // Indicates whether password configuration is required for Proxy.
	Port             *int   `json:"port,omitempty"`             // Port for proxy connection. should be between 1 and 65535
	UserName         string `json:"userName,omitempty"`         // User name for the proxy connection
}

type ResponseProxyUpdateProxyConnection struct {
	Response *ResponseProxyUpdateProxyConnectionResponse `json:"response,omitempty"` // ISE Proxy connection settings
	Version  string                                      `json:"version,omitempty"`  //
}

type ResponseProxyUpdateProxyConnectionResponse struct {
	BypassHosts      string `json:"bypassHosts,omitempty"`      // Bypass hosts for the proxy connection
	Fqdn             string `json:"fqdn,omitempty"`             // proxy IP address or DNS-resolvable host name
	Password         string `json:"password,omitempty"`         // Password for the proxy connection
	PasswordRequired *bool  `json:"passwordRequired,omitempty"` // Indicates whether password configuration is required for Proxy.
	Port             *int   `json:"port,omitempty"`             // Port for proxy connection. should be between 1 and 65535
	UserName         string `json:"userName,omitempty"`         // User name for the proxy connection
}

type RequestProxyUpdateProxyConnection struct {
	BypassHosts      string `json:"bypassHosts,omitempty"`      // Bypass hosts for the proxy connection
	Fqdn             string `json:"fqdn,omitempty"`             // proxy IP address or DNS-resolvable host name
	Password         string `json:"password,omitempty"`         // Password for the proxy connection
	PasswordRequired *bool  `json:"passwordRequired,omitempty"` // Indicates whether password configuration is required for Proxy.
	Port             *int   `json:"port,omitempty"`             // Port for proxy connection. should be between 1 and 65535
	UserName         string `json:"userName,omitempty"`         // User name for the proxy connection
}

//GetProxyConnection Return ISE proxy connection settings
/* The following functionalities are impacted by the proxy settings:


Partner Mobile Management

Endpoint Profiler Feed Service Update

Endpoint Posture Update

Endpoint Posture Agent Resources Download

CRL (Certificate Revocation List) Download

SMS Message Transmission

Social Login

Rest Auth Service Azure AD

pxGrid Cloud



*/
func (s *ProxyService) GetProxyConnection() (*ResponseProxyGetProxyConnection, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/system-settings/proxy"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseProxyGetProxyConnection{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetProxyConnection")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseProxyGetProxyConnection)
	return result, response, err

}

//UpdateProxyConnection Configure ISE proxy connection settings
/* The following functionalities are impacted by the proxy settings:


Partner Mobile Management

Endpoint Profiler Feed Service Update

Endpoint Posture Update

Endpoint Posture Agent Resources Download

CRL (Certificate Revocation List) Download

SMS Message Transmission

Social Login

Rest Auth Service Azure AD

pxGrid Cloud



*/
func (s *ProxyService) UpdateProxyConnection(requestProxyUpdateProxyConnection *RequestProxyUpdateProxyConnection) (*ResponseProxyUpdateProxyConnection, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/system-settings/proxy"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestProxyUpdateProxyConnection).
		SetResult(&ResponseProxyUpdateProxyConnection{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseProxyUpdateProxyConnection{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateProxyConnection")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseProxyUpdateProxyConnection)
	return result, response, err

}
