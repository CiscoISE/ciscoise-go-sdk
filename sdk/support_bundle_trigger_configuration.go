package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type SupportBundleTriggerConfigurationService service

type ResponseSupportBundleTriggerConfigurationGetVersion struct {
	VersionInfo *ResponseSupportBundleTriggerConfigurationGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseSupportBundleTriggerConfigurationGetVersionVersionInfo struct {
	CurrentServerVersion string                                                              `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                                              `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseSupportBundleTriggerConfigurationGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseSupportBundleTriggerConfigurationGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestSupportBundleTriggerConfigurationCreateSupportBundle struct {
	SupportBundle *RequestSupportBundleTriggerConfigurationCreateSupportBundleSupportBundle `json:"SupportBundle,omitempty"` //
}

type RequestSupportBundleTriggerConfigurationCreateSupportBundleSupportBundle struct {
	Name                        string                                                                                               `json:"name,omitempty"`                        // Resource Name
	Description                 string                                                                                               `json:"description,omitempty"`                 //
	HostName                    string                                                                                               `json:"hostName,omitempty"`                    // This parameter is hostName only, xxxx of xxxx.yyy.zz
	SupportBundleIncludeOptions *RequestSupportBundleTriggerConfigurationCreateSupportBundleSupportBundleSupportBundleIncludeOptions `json:"supportBundleIncludeOptions,omitempty"` //
}

type RequestSupportBundleTriggerConfigurationCreateSupportBundleSupportBundleSupportBundleIncludeOptions struct {
	IncludeConfigDB   *bool  `json:"includeConfigDB,omitempty"`   // Set to include Config DB in Support Bundle
	IncludeDebugLogs  *bool  `json:"includeDebugLogs,omitempty"`  // Set to include Debug logs in Support Bundle
	IncludeLocalLogs  *bool  `json:"includeLocalLogs,omitempty"`  // Set to include Local logs in Support Bundle
	IncludeCoreFiles  *bool  `json:"includeCoreFiles,omitempty"`  // Set to include Core files in Support Bundle
	MntLogs           *bool  `json:"mntLogs,omitempty"`           // Set to include Monitoring and troublshooting logs in Support Bundle
	IncludeSystemLogs *bool  `json:"includeSystemLogs,omitempty"` // Set to include System logs in Support Bundle
	PolicyXml         *bool  `json:"policyXml,omitempty"`         // Set to include Policy XML in Support Bundle
	FromDate          string `json:"fromDate,omitempty"`          // Date from where support bundle should include the logs
	ToDate            string `json:"toDate,omitempty"`            // Date upto where support bundle should include the logs
}

//GetVersion Get support bundle trigger configuration version information
/* This API helps to retrieve the version information related to the support bundle trigger configuration.

 */
func (s *SupportBundleTriggerConfigurationService) GetVersion() (*ResponseSupportBundleTriggerConfigurationGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/supportbundle/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSupportBundleTriggerConfigurationGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSupportBundleTriggerConfigurationGetVersion)
	return result, response, err

}

//CreateSupportBundle Create support bundle trigger configuration
/* This API allows the client to create a support bundle trigger configuration.

 */
func (s *SupportBundleTriggerConfigurationService) CreateSupportBundle(requestSupportBundleTriggerConfigurationCreateSupportBundle *RequestSupportBundleTriggerConfigurationCreateSupportBundle) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/supportbundle"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSupportBundleTriggerConfigurationCreateSupportBundle).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateSupportBundle")
	}

	return response, err

}
