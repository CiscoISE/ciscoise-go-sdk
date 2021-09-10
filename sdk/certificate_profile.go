package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type CertificateProfileService service

type GetCertificateProfileQueryParams struct {
	Page int `url:"page,omitempty"` //Page number
	Size int `url:"size,omitempty"` //Number of objects returned per page
}

type ResponseCertificateProfileGetCertificateProfileByName struct {
	CertificateProfile *ResponseCertificateProfileGetCertificateProfileByNameCertificateProfile `json:"CertificateProfile,omitempty"` //
}

type ResponseCertificateProfileGetCertificateProfileByNameCertificateProfile struct {
	ID                        string                                                                       `json:"id,omitempty"`                        //
	Name                      string                                                                       `json:"name,omitempty"`                      //
	Description               string                                                                       `json:"description,omitempty"`               //
	ExternalIDentityStoreName string                                                                       `json:"externalIdentityStoreName,omitempty"` //
	CertificateAttributeName  string                                                                       `json:"certificateAttributeName,omitempty"`  //
	AllowedAsUserName         *bool                                                                        `json:"allowedAsUserName,omitempty"`         //
	MatchMode                 string                                                                       `json:"matchMode,omitempty"`                 //
	UsernameFrom              string                                                                       `json:"usernameFrom,omitempty"`              //
	Link                      *ResponseCertificateProfileGetCertificateProfileByNameCertificateProfileLink `json:"link,omitempty"`                      //
}

type ResponseCertificateProfileGetCertificateProfileByNameCertificateProfileLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificateProfileGetCertificateProfileByID struct {
	CertificateProfile *ResponseCertificateProfileGetCertificateProfileByIDCertificateProfile `json:"CertificateProfile,omitempty"` //
}

type ResponseCertificateProfileGetCertificateProfileByIDCertificateProfile struct {
	ID                        string                                                                     `json:"id,omitempty"`                        //
	Name                      string                                                                     `json:"name,omitempty"`                      //
	Description               string                                                                     `json:"description,omitempty"`               //
	ExternalIDentityStoreName string                                                                     `json:"externalIdentityStoreName,omitempty"` //
	CertificateAttributeName  string                                                                     `json:"certificateAttributeName,omitempty"`  //
	AllowedAsUserName         *bool                                                                      `json:"allowedAsUserName,omitempty"`         //
	MatchMode                 string                                                                     `json:"matchMode,omitempty"`                 //
	UsernameFrom              string                                                                     `json:"usernameFrom,omitempty"`              //
	Link                      *ResponseCertificateProfileGetCertificateProfileByIDCertificateProfileLink `json:"link,omitempty"`                      //
}

type ResponseCertificateProfileGetCertificateProfileByIDCertificateProfileLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificateProfileUpdateCertificateProfileByID struct {
	UpdatedFieldsList *ResponseCertificateProfileUpdateCertificateProfileByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseCertificateProfileUpdateCertificateProfileByIDUpdatedFieldsList struct {
	UpdatedField *[]ResponseCertificateProfileUpdateCertificateProfileByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
}

type ResponseCertificateProfileUpdateCertificateProfileByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseCertificateProfileGetCertificateProfile struct {
	SearchResult *ResponseCertificateProfileGetCertificateProfileSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseCertificateProfileGetCertificateProfileSearchResult struct {
	Total        *int                                                                     `json:"total,omitempty"`        //
	Resources    *[]ResponseCertificateProfileGetCertificateProfileSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponseCertificateProfileGetCertificateProfileSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseCertificateProfileGetCertificateProfileSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseCertificateProfileGetCertificateProfileSearchResultResources struct {
	ID          string                                                                    `json:"id,omitempty"`          //
	Name        string                                                                    `json:"name,omitempty"`        //
	Description string                                                                    `json:"description,omitempty"` //
	Link        *ResponseCertificateProfileGetCertificateProfileSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseCertificateProfileGetCertificateProfileSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificateProfileGetCertificateProfileSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificateProfileGetCertificateProfileSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificateProfileGetVersion struct {
	VersionInfo *ResponseCertificateProfileGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseCertificateProfileGetVersionVersionInfo struct {
	CurrentServerVersion string                                               `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                               `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseCertificateProfileGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseCertificateProfileGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestCertificateProfileUpdateCertificateProfileByID struct {
	CertificateProfile *RequestCertificateProfileUpdateCertificateProfileByIDCertificateProfile `json:"CertificateProfile,omitempty"` //
}

type RequestCertificateProfileUpdateCertificateProfileByIDCertificateProfile struct {
	ID                        string `json:"id,omitempty"`                        //
	Name                      string `json:"name,omitempty"`                      //
	Description               string `json:"description,omitempty"`               //
	ExternalIDentityStoreName string `json:"externalIdentityStoreName,omitempty"` // Referred IDStore name for the Certificate Profile or [not applicable] in case no identity store is chosen
	CertificateAttributeName  string `json:"certificateAttributeName,omitempty"`  // Attribute name of the Certificate Profile - used only when CERTIFICATE is chosen in usernameFrom. Allowed values: - SUBJECT_COMMON_NAME - SUBJECT_ALTERNATIVE_NAME - SUBJECT_SERIAL_NUMBER - SUBJECT - SUBJECT_ALTERNATIVE_NAME_OTHER_NAME - SUBJECT_ALTERNATIVE_NAME_EMAIL - SUBJECT_ALTERNATIVE_NAME_DNS. - Additional internal value ALL_SUBJECT_AND_ALTERNATIVE_NAMES is used automatically when usernameFrom=UPN
	AllowedAsUserName         *bool  `json:"allowedAsUserName,omitempty"`         //
	MatchMode                 string `json:"matchMode,omitempty"`                 // Match mode of the Certificate Profile. Allowed values: - NEVER - RESOLVE_IDENTITY_AMBIGUITY - BINARY_COMPARISON
	UsernameFrom              string `json:"usernameFrom,omitempty"`              // The attribute in the certificate where the user name should be taken from. Allowed values: - CERTIFICATE (for a specific attribute as defined in certificateAttributeName) - UPN (for using any Subject or Alternative Name Attributes in the Certificate - an option only in AD)
}

type RequestCertificateProfileCreateCertificateProfile struct {
	CertificateProfile *RequestCertificateProfileCreateCertificateProfileCertificateProfile `json:"CertificateProfile,omitempty"` //
}

type RequestCertificateProfileCreateCertificateProfileCertificateProfile struct {
	ID                        string `json:"id,omitempty"`                        //
	Name                      string `json:"name,omitempty"`                      //
	Description               string `json:"description,omitempty"`               //
	ExternalIDentityStoreName string `json:"externalIdentityStoreName,omitempty"` // Referred IDStore name for the Certificate Profile or [not applicable] in case no identity store is chosen
	CertificateAttributeName  string `json:"certificateAttributeName,omitempty"`  // Attribute name of the Certificate Profile - used only when CERTIFICATE is chosen in usernameFrom. Allowed values: - SUBJECT_COMMON_NAME - SUBJECT_ALTERNATIVE_NAME - SUBJECT_SERIAL_NUMBER - SUBJECT - SUBJECT_ALTERNATIVE_NAME_OTHER_NAME - SUBJECT_ALTERNATIVE_NAME_EMAIL - SUBJECT_ALTERNATIVE_NAME_DNS. - Additional internal value ALL_SUBJECT_AND_ALTERNATIVE_NAMES is used automatically when usernameFrom=UPN
	AllowedAsUserName         *bool  `json:"allowedAsUserName,omitempty"`         //
	MatchMode                 string `json:"matchMode,omitempty"`                 // Match mode of the Certificate Profile. Allowed values: - NEVER - RESOLVE_IDENTITY_AMBIGUITY - BINARY_COMPARISON
	UsernameFrom              string `json:"usernameFrom,omitempty"`              // The attribute in the certificate where the user name should be taken from. Allowed values: - CERTIFICATE (for a specific attribute as defined in certificateAttributeName) - UPN (for using any Subject or Alternative Name Attributes in the Certificate - an option only in AD)
}

//GetCertificateProfileByName Get certificate profile by name
/* This API allows the client to get a certificate profile by name.

@param name name path parameter.
*/
func (s *CertificateProfileService) GetCertificateProfileByName(name string) (*ResponseCertificateProfileGetCertificateProfileByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/certificateprofile/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCertificateProfileGetCertificateProfileByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetCertificateProfileByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificateProfileGetCertificateProfileByName)
	return result, response, err

}

//GetCertificateProfileByID Get certificate profile by ID
/* This API allows the client to get a certificate profile by ID.

@param id id path parameter.
*/
func (s *CertificateProfileService) GetCertificateProfileByID(id string) (*ResponseCertificateProfileGetCertificateProfileByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/certificateprofile/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCertificateProfileGetCertificateProfileByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetCertificateProfileById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificateProfileGetCertificateProfileByID)
	return result, response, err

}

//GetCertificateProfile Get all certificate profiles
/* This API allows the client to get all the certificate profiles.

@param getCertificateProfileQueryParams Filtering parameter
*/
func (s *CertificateProfileService) GetCertificateProfile(getCertificateProfileQueryParams *GetCertificateProfileQueryParams) (*ResponseCertificateProfileGetCertificateProfile, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/certificateprofile"

	queryString, _ := query.Values(getCertificateProfileQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseCertificateProfileGetCertificateProfile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetCertificateProfile")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificateProfileGetCertificateProfile)
	return result, response, err

}

//GetVersion Get certificate profile version information
/* This API helps to retrieve the version information related to the certificate profile.

 */
func (s *CertificateProfileService) GetVersion() (*ResponseCertificateProfileGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/certificateprofile/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCertificateProfileGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificateProfileGetVersion)
	return result, response, err

}

//CreateCertificateProfile Create certificate profile
/* This API allows the client to create a certificate profile.

 */
func (s *CertificateProfileService) CreateCertificateProfile(requestCertificateProfileCreateCertificateProfile *RequestCertificateProfileCreateCertificateProfile) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/certificateprofile"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCertificateProfileCreateCertificateProfile).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateCertificateProfile")
	}

	return response, err

}

//UpdateCertificateProfileByID Update certificate profile
/* This API allows the client to update a certificate profile.

@param id id path parameter.
*/
func (s *CertificateProfileService) UpdateCertificateProfileByID(id string, requestCertificateProfileUpdateCertificateProfileById *RequestCertificateProfileUpdateCertificateProfileByID) (*ResponseCertificateProfileUpdateCertificateProfileByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/certificateprofile/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCertificateProfileUpdateCertificateProfileById).
		SetResult(&ResponseCertificateProfileUpdateCertificateProfileByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateCertificateProfileById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificateProfileUpdateCertificateProfileByID)
	return result, response, err

}
