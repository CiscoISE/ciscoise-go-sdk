package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type NativeSupplicantProfileService service

type GetNativeSupplicantProfileQueryParams struct {
	Page int `url:"page,omitempty"` //Page number
	Size int `url:"size,omitempty"` //Number of objects returned per page
}

type ResponseNativeSupplicantProfileGetNativeSupplicantProfileByID struct {
	ERSNSpProfile *ResponseNativeSupplicantProfileGetNativeSupplicantProfileByIDERSNSpProfile `json:"ERSNSPProfile,omitempty"` //
}

type ResponseNativeSupplicantProfileGetNativeSupplicantProfileByIDERSNSpProfile struct {
	ID               string                                                                                        `json:"id,omitempty"`               //
	Name             string                                                                                        `json:"name,omitempty"`             //
	Description      string                                                                                        `json:"description,omitempty"`      //
	WirelessProfiles *[]ResponseNativeSupplicantProfileGetNativeSupplicantProfileByIDERSNSpProfileWirelessProfiles `json:"wirelessProfiles,omitempty"` //
	Link             *ResponseNativeSupplicantProfileGetNativeSupplicantProfileByIDERSNSpProfileLink               `json:"link,omitempty"`             //
}

type ResponseNativeSupplicantProfileGetNativeSupplicantProfileByIDERSNSpProfileWirelessProfiles struct {
	SSID                  string `json:"ssid,omitempty"`                  //
	AllowedProtocol       string `json:"allowedProtocol,omitempty"`       //
	CertificateTemplateID string `json:"certificateTemplateId,omitempty"` //
	ActionType            string `json:"actionType,omitempty"`            // Action type for WifiProfile. Allowed values: - ADD, - UPDATE, - DELETE (required for updating existing WirelessProfile)
	PreviousSSID          string `json:"previousSsid,omitempty"`          // Previous ssid for WifiProfile (required for updating existing WirelessProfile)
}

type ResponseNativeSupplicantProfileGetNativeSupplicantProfileByIDERSNSpProfileLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNativeSupplicantProfileUpdateNativeSupplicantProfileByID struct {
	UpdatedFieldsList *ResponseNativeSupplicantProfileUpdateNativeSupplicantProfileByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseNativeSupplicantProfileUpdateNativeSupplicantProfileByIDUpdatedFieldsList struct {
	UpdatedField *[]ResponseNativeSupplicantProfileUpdateNativeSupplicantProfileByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                                           `json:"field,omitempty"`        //
	OldValue     string                                                                                           `json:"oldValue,omitempty"`     //
	NewValue     string                                                                                           `json:"newValue,omitempty"`     //
}

type ResponseNativeSupplicantProfileUpdateNativeSupplicantProfileByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseNativeSupplicantProfileGetNativeSupplicantProfile struct {
	SearchResult *ResponseNativeSupplicantProfileGetNativeSupplicantProfileSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseNativeSupplicantProfileGetNativeSupplicantProfileSearchResult struct {
	Total        *int                                                                               `json:"total,omitempty"`        //
	Resources    *[]ResponseNativeSupplicantProfileGetNativeSupplicantProfileSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponseNativeSupplicantProfileGetNativeSupplicantProfileSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseNativeSupplicantProfileGetNativeSupplicantProfileSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseNativeSupplicantProfileGetNativeSupplicantProfileSearchResultResources struct {
	ID          string                                                                              `json:"id,omitempty"`          //
	Name        string                                                                              `json:"name,omitempty"`        //
	Description string                                                                              `json:"description,omitempty"` //
	Link        *ResponseNativeSupplicantProfileGetNativeSupplicantProfileSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseNativeSupplicantProfileGetNativeSupplicantProfileSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNativeSupplicantProfileGetNativeSupplicantProfileSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNativeSupplicantProfileGetNativeSupplicantProfileSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseNativeSupplicantProfileGetVersion struct {
	VersionInfo *ResponseNativeSupplicantProfileGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseNativeSupplicantProfileGetVersionVersionInfo struct {
	CurrentServerVersion string                                                    `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                                    `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseNativeSupplicantProfileGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseNativeSupplicantProfileGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestNativeSupplicantProfileUpdateNativeSupplicantProfileByID struct {
	ERSNSpProfile *RequestNativeSupplicantProfileUpdateNativeSupplicantProfileByIDERSNSpProfile `json:"ERSNSPProfile,omitempty"` //
}

type RequestNativeSupplicantProfileUpdateNativeSupplicantProfileByIDERSNSpProfile struct {
	ID               string                                                                                          `json:"id,omitempty"`               //
	Name             string                                                                                          `json:"name,omitempty"`             //
	Description      string                                                                                          `json:"description,omitempty"`      //
	WirelessProfiles *[]RequestNativeSupplicantProfileUpdateNativeSupplicantProfileByIDERSNSpProfileWirelessProfiles `json:"wirelessProfiles,omitempty"` //
}

type RequestNativeSupplicantProfileUpdateNativeSupplicantProfileByIDERSNSpProfileWirelessProfiles struct {
	SSID                  string `json:"ssid,omitempty"`                  //
	AllowedProtocol       string `json:"allowedProtocol,omitempty"`       //
	CertificateTemplateID string `json:"certificateTemplateId,omitempty"` //
	ActionType            string `json:"actionType,omitempty"`            // Action type for WifiProfile. Allowed values: - ADD, - UPDATE, - DELETE (required for updating existing WirelessProfile)
	PreviousSSID          string `json:"previousSsid,omitempty"`          // Previous ssid for WifiProfile (required for updating existing WirelessProfile)
}

//GetNativeSupplicantProfileByID Get native supplicant profile by ID
/* This API allows the client to get a native supplicant profile by ID.

@param id id path parameter.
*/
func (s *NativeSupplicantProfileService) GetNativeSupplicantProfileByID(id string) (*ResponseNativeSupplicantProfileGetNativeSupplicantProfileByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/nspprofile/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNativeSupplicantProfileGetNativeSupplicantProfileByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNativeSupplicantProfileById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNativeSupplicantProfileGetNativeSupplicantProfileByID)
	return result, response, err

}

//GetNativeSupplicantProfile Get all native supplicant profiles
/* This API allows the client to get all the native supplicant profiles.

@param getNativeSupplicantProfileQueryParams Filtering parameter
*/
func (s *NativeSupplicantProfileService) GetNativeSupplicantProfile(getNativeSupplicantProfileQueryParams *GetNativeSupplicantProfileQueryParams) (*ResponseNativeSupplicantProfileGetNativeSupplicantProfile, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/nspprofile"

	queryString, _ := query.Values(getNativeSupplicantProfileQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNativeSupplicantProfileGetNativeSupplicantProfile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNativeSupplicantProfile")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNativeSupplicantProfileGetNativeSupplicantProfile)
	return result, response, err

}

//GetVersion Get native supplicant profile version information
/* This API helps to retrieve the version information related to the native supplicant profile.

 */
func (s *NativeSupplicantProfileService) GetVersion() (*ResponseNativeSupplicantProfileGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/nspprofile/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNativeSupplicantProfileGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNativeSupplicantProfileGetVersion)
	return result, response, err

}

//UpdateNativeSupplicantProfileByID Update native supplicant profile
/* This API allows the client to update a native supplicant profile

@param id id path parameter.
*/
func (s *NativeSupplicantProfileService) UpdateNativeSupplicantProfileByID(id string, requestNativeSupplicantProfileUpdateNativeSupplicantProfileById *RequestNativeSupplicantProfileUpdateNativeSupplicantProfileByID) (*ResponseNativeSupplicantProfileUpdateNativeSupplicantProfileByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/nspprofile/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNativeSupplicantProfileUpdateNativeSupplicantProfileById).
		SetResult(&ResponseNativeSupplicantProfileUpdateNativeSupplicantProfileByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNativeSupplicantProfileById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNativeSupplicantProfileUpdateNativeSupplicantProfileByID)
	return result, response, err

}

//DeleteNativeSupplicantProfileByID Delete native supplicant profile
/* This API deletes a native supplicant profile.

@param id id path parameter.
*/
func (s *NativeSupplicantProfileService) DeleteNativeSupplicantProfileByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/nspprofile/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNativeSupplicantProfileById")
	}

	getCSFRToken(response.Header())

	return response, err

}
