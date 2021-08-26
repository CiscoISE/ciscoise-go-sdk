package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ClearThreatsAndVulnerabilitiesService service

type ResponseClearThreatsAndVulnerabilitiesGetVersion struct {
	VersionInfo *ResponseClearThreatsAndVulnerabilitiesGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseClearThreatsAndVulnerabilitiesGetVersionVersionInfo struct {
	CurrentServerVersion string                                                           `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                                           `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseClearThreatsAndVulnerabilitiesGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseClearThreatsAndVulnerabilitiesGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestClearThreatsAndVulnerabilitiesClearThreatsAndVulnerabilities struct {
	ERSIrfThreatContext *RequestClearThreatsAndVulnerabilitiesClearThreatsAndVulnerabilitiesERSIrfThreatContext `json:"ERSIrfThreatContext,omitempty"` //
}

type RequestClearThreatsAndVulnerabilitiesClearThreatsAndVulnerabilitiesERSIrfThreatContext struct {
	MacAddresses string `json:"macAddresses,omitempty"` //
}

//GetVersion Get clear threats and vulneribilities version information
/* This API helps to retrieve the version information related to the clear threats and vulneribilities.

 */
func (s *ClearThreatsAndVulnerabilitiesService) GetVersion() (*ResponseClearThreatsAndVulnerabilitiesGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/threat/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseClearThreatsAndVulnerabilitiesGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseClearThreatsAndVulnerabilitiesGetVersion)
	return result, response, err

}

//ClearThreatsAndVulnerabilities Delete ThreatContext and Threat events that are associated with given MacAddress
/* This API allows the client to delete the ThreatContext and Threat events that are associated with the given MAC Address.

 */
func (s *ClearThreatsAndVulnerabilitiesService) ClearThreatsAndVulnerabilities(requestClearThreatsAndVulnerabilitiesClearThreatsAndVulnerabilities *RequestClearThreatsAndVulnerabilitiesClearThreatsAndVulnerabilities) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/threat/clearThreatsAndVulneribilities"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestClearThreatsAndVulnerabilitiesClearThreatsAndVulnerabilities).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation ClearThreatsAndVulnerabilities")
	}

	getCSFRToken(response.Header())

	return response, err

}
