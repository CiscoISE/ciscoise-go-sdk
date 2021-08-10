package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type AciSettingsService service

type ResponseAciSettingsGetAciSettings struct {
	AciSettings ResponseAciSettingsGetAciSettingsAciSettings `json:"AciSettings,omitempty"` //
}

type ResponseAciSettingsGetAciSettingsAciSettings struct {
	ID                     string   `json:"id,omitempty"`                     // Resource UUID value
	EnableAci              bool     `json:"enableAci,omitempty"`              // Enable ACI Integration
	IPAddressHostName      string   `json:"ipAddressHostName,omitempty"`      // ACI Cluster IP Address / Host name
	AdminName              string   `json:"adminName,omitempty"`              // ACI Cluster Admin name
	AdminPassword          string   `json:"adminPassword,omitempty"`          // ACI Cluster Admin password
	Aciipaddress           string   `json:"aciipaddress,omitempty"`           // ACI Domain manager Ip Address.
	AciuserName            string   `json:"aciuserName,omitempty"`            // ACI Domain manager Username.
	Acipassword            string   `json:"acipassword,omitempty"`            // ACI Domain manager Password.
	TenantName             string   `json:"tenantName,omitempty"`             //
	L3RouteNetwork         string   `json:"l3RouteNetwork,omitempty"`         //
	SuffixToEpg            string   `json:"suffixToEpg,omitempty"`            //
	SuffixToSgt            string   `json:"suffixToSgt,omitempty"`            //
	AllSxpDomain           bool     `json:"allSxpDomain,omitempty"`           //
	SpecificSxpDomain      bool     `json:"specificSxpDomain,omitempty"`      //
	SpecifixSxpDomainList  []string `json:"specifixSxpDomainList,omitempty"`  //
	EnableDataPlane        bool     `json:"enableDataPlane,omitempty"`        //
	UntaggedPacketIepgName string   `json:"untaggedPacketIepgName,omitempty"` //
	DefaultSgtName         string   `json:"defaultSgtName,omitempty"`         //
	EnableElementsLimit    bool     `json:"enableElementsLimit,omitempty"`    //
	MaxNumIepgFromAci      int      `json:"maxNumIepgFromAci,omitempty"`      //
	MaxNumSgtToAci         int      `json:"maxNumSgtToAci,omitempty"`         //
	Aci50                  bool     `json:"aci50,omitempty"`                  // Enable 5.0 ACI Version
	Aci51                  bool     `json:"aci51,omitempty"`                  // Enable 5.1 ACI Version
}

type ResponseAciSettingsTestAciConnectivity struct {
	AciTestConnectionResult ResponseAciSettingsTestAciConnectivityAciTestConnectionResult `json:"ACITestConnectionResult,omitempty"` //
}

type ResponseAciSettingsTestAciConnectivityAciTestConnectionResult struct {
	Result bool `json:"result,omitempty"` //
}

type ResponseAciSettingsUpdateAciSettingsByID struct {
	UpdatedFieldsList ResponseAciSettingsUpdateAciSettingsByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseAciSettingsUpdateAciSettingsByIDUpdatedFieldsList struct {
	UpdatedField ResponseAciSettingsUpdateAciSettingsByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                `json:"field,omitempty"`        //
	OldValue     string                                                                `json:"oldValue,omitempty"`     //
	NewValue     string                                                                `json:"newValue,omitempty"`     //
}

type ResponseAciSettingsUpdateAciSettingsByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseAciSettingsGetVersion struct {
	VersionInfo ResponseAciSettingsGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseAciSettingsGetVersionVersionInfo struct {
	CurrentServerVersion string                                       `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                       `json:"supportedVersions,omitempty"`    //
	Link                 ResponseAciSettingsGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseAciSettingsGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestAciSettingsUpdateAciSettingsByID struct {
	AciSettings *RequestAciSettingsUpdateAciSettingsByIDAciSettings `json:"AciSettings,omitempty"` //
}

type RequestAciSettingsUpdateAciSettingsByIDAciSettings struct {
	ID                     string   `json:"id,omitempty"`                     // Resource UUID value
	EnableAci              *bool    `json:"enableAci,omitempty"`              // Enable ACI Integration
	IPAddressHostName      string   `json:"ipAddressHostName,omitempty"`      // ACI Cluster IP Address / Host name
	AdminName              string   `json:"adminName,omitempty"`              // ACI Cluster Admin name
	AdminPassword          string   `json:"adminPassword,omitempty"`          // ACI Cluster Admin password
	Aciipaddress           string   `json:"aciipaddress,omitempty"`           // ACI Domain manager Ip Address.
	AciuserName            string   `json:"aciuserName,omitempty"`            // ACI Domain manager Username.
	Acipassword            string   `json:"acipassword,omitempty"`            // ACI Domain manager Password.
	TenantName             string   `json:"tenantName,omitempty"`             //
	L3RouteNetwork         string   `json:"l3RouteNetwork,omitempty"`         //
	SuffixToEpg            string   `json:"suffixToEpg,omitempty"`            //
	SuffixToSgt            string   `json:"suffixToSgt,omitempty"`            //
	AllSxpDomain           *bool    `json:"allSxpDomain,omitempty"`           //
	SpecificSxpDomain      *bool    `json:"specificSxpDomain,omitempty"`      //
	SpecifixSxpDomainList  []string `json:"specifixSxpDomainList,omitempty"`  //
	EnableDataPlane        *bool    `json:"enableDataPlane,omitempty"`        //
	UntaggedPacketIepgName string   `json:"untaggedPacketIepgName,omitempty"` //
	DefaultSgtName         string   `json:"defaultSgtName,omitempty"`         //
	EnableElementsLimit    *bool    `json:"enableElementsLimit,omitempty"`    //
	MaxNumIepgFromAci      *int     `json:"maxNumIepgFromAci,omitempty"`      //
	MaxNumSgtToAci         *int     `json:"maxNumSgtToAci,omitempty"`         //
	Aci50                  *bool    `json:"aci50,omitempty"`                  // Enable 5.0 ACI Version
	Aci51                  *bool    `json:"aci51,omitempty"`                  // Enable 5.1 ACI Version
}

//GetAciSettings Get all ACI Information
/* This API allows the client to get ACI Settings.

 */
func (s *AciSettingsService) GetAciSettings() (*ResponseAciSettingsGetAciSettings, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/acisettings"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAciSettingsGetAciSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAciSettings")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAciSettingsGetAciSettings)
	return result, response, err

}

//GetVersion Get ACI settings version information
/* This API helps to retrieve the version information related to the Cisco ACI settings.

 */
func (s *AciSettingsService) GetVersion() (*ResponseAciSettingsGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/acisettings/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAciSettingsGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAciSettingsGetVersion)
	return result, response, err

}

//TestAciConnectivity Test ACI Domain Manager connection
/* This API allows the client to test ACI Domain Manager connection.

 */
func (s *AciSettingsService) TestAciConnectivity() (*ResponseAciSettingsTestAciConnectivity, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/acisettings/testACIConnectivity"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAciSettingsTestAciConnectivity{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation TestAciConnectivity")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAciSettingsTestAciConnectivity)
	return result, response, err

}

//UpdateAciSettingsByID Update ACI settings
/* This API allows the client to update ACI settings.

@param id id path parameter.
*/
func (s *AciSettingsService) UpdateAciSettingsByID(id string, requestAciSettingsUpdateAciSettingsById *RequestAciSettingsUpdateAciSettingsByID) (*ResponseAciSettingsUpdateAciSettingsByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/acisettings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAciSettingsUpdateAciSettingsById).
		SetResult(&ResponseAciSettingsUpdateAciSettingsByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateAciSettingsById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAciSettingsUpdateAciSettingsByID)
	return result, response, err

}
