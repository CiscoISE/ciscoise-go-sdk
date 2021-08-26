package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type SupportBundleDownloadService service

type ResponseSupportBundleDownloadGetVersion struct {
	VersionInfo *ResponseSupportBundleDownloadGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseSupportBundleDownloadGetVersionVersionInfo struct {
	CurrentServerVersion string                                                  `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                                  `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseSupportBundleDownloadGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseSupportBundleDownloadGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestSupportBundleDownloadDownloadSupportBundle struct {
	ErsSupportBundleDownload *RequestSupportBundleDownloadDownloadSupportBundleErsSupportBundleDownload `json:"ErsSupportBundleDownload,omitempty"` //
}

type RequestSupportBundleDownloadDownloadSupportBundleErsSupportBundleDownload struct {
	FileName string `json:"fileName,omitempty"` //
}

//GetVersion Get support bundle download version information
/* This API helps to retrieve the version information related to the support bundle download.

 */
func (s *SupportBundleDownloadService) GetVersion() (*ResponseSupportBundleDownloadGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/supportbundledownload/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSupportBundleDownloadGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSupportBundleDownloadGetVersion)
	return result, response, err

}

//DownloadSupportBundle Upload support bundle
/* This API allows the client to upload a support bundle.

 */
func (s *SupportBundleDownloadService) DownloadSupportBundle(requestSupportBundleDownloadDownloadSupportBundle *RequestSupportBundleDownloadDownloadSupportBundle) (FileDownload, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/supportbundledownload"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSupportBundleDownloadDownloadSupportBundle).
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
