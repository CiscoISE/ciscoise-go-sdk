package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type PxGridSettingsService service

type ResponsePxGridSettingsGetVersion struct {
	VersionInfo *ResponsePxGridSettingsGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponsePxGridSettingsGetVersionVersionInfo struct {
	CurrentServerVersion string                                           `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                           `json:"supportedVersions,omitempty"`    //
	Link                 *ResponsePxGridSettingsGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponsePxGridSettingsGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestPxGridSettingsAutoapprovePxGridSettings struct {
	PxgridSettings *RequestPxGridSettingsAutoapprovePxGridSettingsPxgridSettings `json:"PxgridSettings,omitempty"` //
}

type RequestPxGridSettingsAutoapprovePxGridSettingsPxgridSettings struct {
	AutoApproveCertBasedAccounts *bool `json:"autoApproveCertBasedAccounts,omitempty"` // Auto approve certificate based accounts when true
	AllowPasswordBasedAccounts   *bool `json:"allowPasswordBasedAccounts,omitempty"`   // Allow password based accounts when true
}

//GetVersion Get pxGrid settings version information
/* This API helps to retrieve the version information related to the pxGrid settings.

 */
func (s *PxGridSettingsService) GetVersion() (*ResponsePxGridSettingsGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/pxgridsettings/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePxGridSettingsGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePxGridSettingsGetVersion)
	return result, response, err

}

//AutoapprovePxGridSettings Auto approve pxGrid settings
/* This API allows the client to auto approve the pxGrid settings.

 */
func (s *PxGridSettingsService) AutoapprovePxGridSettings(requestPxGridSettingsAutoapprovePxGridSettings *RequestPxGridSettingsAutoapprovePxGridSettings) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/pxgridsettings/autoapprove"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestPxGridSettingsAutoapprovePxGridSettings).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation AutoapprovePxGridSettings")
	}

	getCSFRToken(response.Header())

	return response, err

}
