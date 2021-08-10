package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type SystemCertificateService service

type ResponseSystemCertificateGetVersion struct {
	VersionInfo ResponseSystemCertificateGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseSystemCertificateGetVersionVersionInfo struct {
	CurrentServerVersion string                                             `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                             `json:"supportedVersions,omitempty"`    //
	Link                 ResponseSystemCertificateGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseSystemCertificateGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestSystemCertificateCreateSystemCertificate struct {
	ERSSystemCertificate *RequestSystemCertificateCreateSystemCertificateERSSystemCertificate `json:"ERSSystemCertificate,omitempty"` //
}

type RequestSystemCertificateCreateSystemCertificateERSSystemCertificate struct {
	NodeID           string                                                                               `json:"nodeId,omitempty"`           // NodeId of Cisco ISE application
	ErsLocalCertStub *RequestSystemCertificateCreateSystemCertificateERSSystemCertificateErsLocalCertStub `json:"ersLocalCertStub,omitempty"` // Inputs for certificate creation
}

type RequestSystemCertificateCreateSystemCertificateERSSystemCertificateErsLocalCertStub struct {
	FriendlyName              string                                                                                             `json:"friendlyName,omitempty"`              //
	ErsSubjectStub            *RequestSystemCertificateCreateSystemCertificateERSSystemCertificateErsLocalCertStubErsSubjectStub `json:"ersSubjectStub,omitempty"`            // Subject data of certificate
	KeyLength                 string                                                                                             `json:"keyLength,omitempty"`                 //
	XgridCertificate          string                                                                                             `json:"xgridCertificate,omitempty"`          //
	GroupTagDD                string                                                                                             `json:"groupTagDD,omitempty"`                //
	SamlCertificate           string                                                                                             `json:"samlCertificate,omitempty"`           //
	KeyType                   string                                                                                             `json:"keyType,omitempty"`                   //
	Digest                    string                                                                                             `json:"digest,omitempty"`                    //
	CertificatePolicies       string                                                                                             `json:"certificatePolicies,omitempty"`       //
	ExpirationTTL             *int                                                                                               `json:"expirationTTL,omitempty"`             //
	SelectedExpirationTTLUnit string                                                                                             `json:"selectedExpirationTTLUnit,omitempty"` //
	AllowWildcardCerts        string                                                                                             `json:"allowWildcardCerts,omitempty"`        //
	CertificateSanDNS         string                                                                                             `json:"certificateSanDns,omitempty"`         //
	CertificateSanIP          string                                                                                             `json:"certificateSanIp,omitempty"`          //
	CertificateSanURI         string                                                                                             `json:"certificateSanUri,omitempty"`         //
}

type RequestSystemCertificateCreateSystemCertificateERSSystemCertificateErsLocalCertStubErsSubjectStub struct {
	CommonName             string `json:"commonName,omitempty"`             //
	OrganizationalUnitName string `json:"organizationalUnitName,omitempty"` //
	OrganizationName       string `json:"organizationName,omitempty"`       //
	LocalityName           string `json:"localityName,omitempty"`           //
	StateOrProvinceName    string `json:"stateOrProvinceName,omitempty"`    //
	CountryName            string `json:"countryName,omitempty"`            //
}

//GetVersion Get system certificate version information
/* This API helps to retrieve the version information related to the system certificate.

 */
func (s *SystemCertificateService) GetVersion() (*ResponseSystemCertificateGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/systemcertificate/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSystemCertificateGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSystemCertificateGetVersion)
	return result, response, err

}

//CreateSystemCertificate Create system certificate
/* This API allows the client to create a system certificate.

 */
func (s *SystemCertificateService) CreateSystemCertificate(requestSystemCertificateCreateSystemCertificate *RequestSystemCertificateCreateSystemCertificate) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/systemcertificate"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSystemCertificateCreateSystemCertificate).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateSystemCertificate")
	}

	return response, err

}
