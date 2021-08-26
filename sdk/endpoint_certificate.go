package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type EndpointCertificateService service

type ResponseEndpointCertificateGetVersion struct {
	VersionInfo *ResponseEndpointCertificateGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseEndpointCertificateGetVersionVersionInfo struct {
	CurrentServerVersion string                                                `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                                `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseEndpointCertificateGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseEndpointCertificateGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestEndpointCertificateCreateEndpointCertificate struct {
	ERSEndPointCert *RequestEndpointCertificateCreateEndpointCertificateERSEndPointCert `json:"ERSEndPointCert,omitempty"` //
}

type RequestEndpointCertificateCreateEndpointCertificateERSEndPointCert struct {
	CertTemplateName   string                                                                                `json:"certTemplateName,omitempty"`   // Name of an Internal CA template
	Format             string                                                                                `json:"format,omitempty"`             // Allowed values: - PKCS12, - PKCS12_CHAIN, - PKCS8, - PKCS8_CHAIN
	Password           string                                                                                `json:"password,omitempty"`           // Protects the private key. Must have more than 8 characters, less than 15 characters, at least one upper case letter, at least one lower case letter, at least one digit, and can only contain [A-Z][a-z][0-9]_#
	CertificateRequest *RequestEndpointCertificateCreateEndpointCertificateERSEndPointCertCertificateRequest `json:"certificateRequest,omitempty"` // Key value map. Must have CN and SAN entries
}

type RequestEndpointCertificateCreateEndpointCertificateERSEndPointCertCertificateRequest struct {
	San string `json:"san,omitempty"` // Valid MAC Address, delimited by '-'
	Cn  string `json:"cn,omitempty"`  // Matches the requester's User Name, unless the Requester is an ERS Admin. ERS Admins are allowed to create requests for any CN
}

//GetVersion Get endpoint certificate version information
/* This API helps to retrieve the version information related to the endpoint certificate.

 */
func (s *EndpointCertificateService) GetVersion() (*ResponseEndpointCertificateGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpointcert/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEndpointCertificateGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointCertificateGetVersion)
	return result, response, err

}

//CreateEndpointCertificate Create endpoint certificate
/* This API allows the client to create an endpoint certificate.

 */
func (s *EndpointCertificateService) CreateEndpointCertificate(requestEndpointCertificateCreateEndpointCertificate *RequestEndpointCertificateCreateEndpointCertificate) (FileDownload, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/endpointcert/certRequest"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEndpointCertificateCreateEndpointCertificate).
		SetError(&Error).
		Put(path)

	fdownload := FileDownload{}
	if err != nil {
		return fdownload, nil, err
	}

	if response.IsError() {
		return fdownload, response, fmt.Errorf("error with operation ExportTrustedCertificate")
	}

	fdownload.FileData = response.Body()
	headerVal := response.Header()["Content-Disposition"][0]
	fname := strings.SplitAfter(headerVal, "=")
	fdownload.FileName = strings.ReplaceAll(fname[1], "\"", "")

	getCSFRToken(response.Header())
	return fdownload, response, err

}
