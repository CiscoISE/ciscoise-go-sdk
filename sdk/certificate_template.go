package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type CertificateTemplateService service

type GetCertificateTemplateQueryParams struct {
	Page int `url:"page,omitempty"` //Page number
	Size int `url:"size,omitempty"` //Number of objects returned per page
}

type ResponseCertificateTemplateGetCertificateTemplateByName struct {
	ERSCertificateTemplate ResponseCertificateTemplateGetCertificateTemplateByNameERSCertificateTemplate `json:"ERSCertificateTemplate,omitempty"` //
}

type ResponseCertificateTemplateGetCertificateTemplateByNameERSCertificateTemplate struct {
	ID             string `json:"id,omitempty"`             //
	Name           string `json:"name,omitempty"`           //
	Description    string `json:"description,omitempty"`    //
	KeySize        int    `json:"keySize,omitempty"`        // Key Size of the Certificate Template
	ValidityPeriod int    `json:"validityPeriod,omitempty"` // Validity period of the Certificate Template: Valid Range 21 - 3652
	Raprofile      string `json:"raprofile,omitempty"`      // RA profile for the Certificate template
}

type ResponseCertificateTemplateGetCertificateTemplateByID struct {
	ERSCertificateTemplate ResponseCertificateTemplateGetCertificateTemplateByIDERSCertificateTemplate `json:"ERSCertificateTemplate,omitempty"` //
}

type ResponseCertificateTemplateGetCertificateTemplateByIDERSCertificateTemplate struct {
	ID             string `json:"id,omitempty"`             //
	Name           string `json:"name,omitempty"`           //
	Description    string `json:"description,omitempty"`    //
	KeySize        int    `json:"keySize,omitempty"`        // Key Size of the Certificate Template
	ValidityPeriod int    `json:"validityPeriod,omitempty"` // Validity period of the Certificate Template: Valid Range 21 - 3652
	Raprofile      string `json:"raprofile,omitempty"`      // RA profile for the Certificate template
}

type ResponseCertificateTemplateGetCertificateTemplate struct {
	SearchResult ResponseCertificateTemplateGetCertificateTemplateSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseCertificateTemplateGetCertificateTemplateSearchResult struct {
	Total        int                                                                       `json:"total,omitempty"`        //
	Resources    []ResponseCertificateTemplateGetCertificateTemplateSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponseCertificateTemplateGetCertificateTemplateSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponseCertificateTemplateGetCertificateTemplateSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseCertificateTemplateGetCertificateTemplateSearchResultResources struct {
	ID          string                                                                     `json:"id,omitempty"`          //
	Name        string                                                                     `json:"name,omitempty"`        //
	Description string                                                                     `json:"description,omitempty"` //
	Link        ResponseCertificateTemplateGetCertificateTemplateSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseCertificateTemplateGetCertificateTemplateSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificateTemplateGetCertificateTemplateSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificateTemplateGetCertificateTemplateSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificateTemplateGetVersion struct {
	VersionInfo ResponseCertificateTemplateGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseCertificateTemplateGetVersionVersionInfo struct {
	CurrentServerVersion string                                               `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                               `json:"supportedVersions,omitempty"`    //
	Link                 ResponseCertificateTemplateGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseCertificateTemplateGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

//GetCertificateTemplateByName Get certificate template by name
/* This API allows the client to get a certificate template by name.

@param name name path parameter.
*/
func (s *CertificateTemplateService) GetCertificateTemplateByName(name string) (*ResponseCertificateTemplateGetCertificateTemplateByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/certificatetemplate/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCertificateTemplateGetCertificateTemplateByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetCertificateTemplateByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificateTemplateGetCertificateTemplateByName)
	return result, response, err

}

//GetCertificateTemplateByID Get certificate template by ID
/* This API allows the client to get a certificate template by ID.

@param id id path parameter.
*/
func (s *CertificateTemplateService) GetCertificateTemplateByID(id string) (*ResponseCertificateTemplateGetCertificateTemplateByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/certificatetemplate/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCertificateTemplateGetCertificateTemplateByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetCertificateTemplateById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificateTemplateGetCertificateTemplateByID)
	return result, response, err

}

//GetCertificateTemplate Get all certificate templates
/* This API allows the client to get aall the certificate templates.

@param getCertificateTemplateQueryParams Filtering parameter
*/
func (s *CertificateTemplateService) GetCertificateTemplate(getCertificateTemplateQueryParams *GetCertificateTemplateQueryParams) (*ResponseCertificateTemplateGetCertificateTemplate, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/certificatetemplate/"

	queryString, _ := query.Values(getCertificateTemplateQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseCertificateTemplateGetCertificateTemplate{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetCertificateTemplate")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificateTemplateGetCertificateTemplate)
	return result, response, err

}

//GetVersion Get certificate template version information
/* This API helps to retrieve the version information related to the certificate template.

 */
func (s *CertificateTemplateService) GetVersion() (*ResponseCertificateTemplateGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/certificatetemplate/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCertificateTemplateGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificateTemplateGetVersion)
	return result, response, err

}
