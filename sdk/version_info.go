package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type VersionInfoService service

type ResponseVersionInfoGetVersionInfo struct {
	VersionInfo *ResponseVersionInfoGetVersionInfoVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseVersionInfoGetVersionInfoVersionInfo struct {
	CurrentServerVersion string                                            `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                            `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseVersionInfoGetVersionInfoVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseVersionInfoGetVersionInfoVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

//GetVersionInfo Get the version info of the given resource
/* Get all VersionInfo

@param resource resource path parameter.
*/
func (s *VersionInfoService) GetVersionInfo(resource string) (*ResponseVersionInfoGetVersionInfo, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/{resource}/versioninfo"
	path = strings.Replace(path, "{resource}", fmt.Sprintf("%v", resource), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseVersionInfoGetVersionInfo{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersionInfo")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseVersionInfoGetVersionInfo)
	return result, response, err

}
