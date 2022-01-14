package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type TacacsProfileService service

type GetTacacsProfileQueryParams struct {
	Page int `url:"page,omitempty"` //Page number
	Size int `url:"size,omitempty"` //Number of objects returned per page
}

type ResponseTacacsProfileGetTacacsProfileByName struct {
	TacacsProfile *ResponseTacacsProfileGetTacacsProfileByNameTacacsProfile `json:"TacacsProfile,omitempty"` //
}

type ResponseTacacsProfileGetTacacsProfileByNameTacacsProfile struct {
	ID                string                                                                     `json:"id,omitempty"`                //
	Name              string                                                                     `json:"name,omitempty"`              //
	Description       string                                                                     `json:"description,omitempty"`       //
	SessionAttributes *ResponseTacacsProfileGetTacacsProfileByNameTacacsProfileSessionAttributes `json:"sessionAttributes,omitempty"` // Holds list of session attributes. View type for GUI is Shell by default
	Link              *ResponseTacacsProfileGetTacacsProfileByNameTacacsProfileLink              `json:"link,omitempty"`              //
}

type ResponseTacacsProfileGetTacacsProfileByNameTacacsProfileSessionAttributes struct {
	SessionAttributeList *[]ResponseTacacsProfileGetTacacsProfileByNameTacacsProfileSessionAttributesSessionAttributeList `json:"sessionAttributeList,omitempty"` //
}

type ResponseTacacsProfileGetTacacsProfileByNameTacacsProfileSessionAttributesSessionAttributeList struct {
	Type  string `json:"type,omitempty"`  // Allowed values: MANDATORY, OPTIONAL
	Name  string `json:"name,omitempty"`  //
	Value string `json:"value,omitempty"` //
}

type ResponseTacacsProfileGetTacacsProfileByNameTacacsProfileLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsProfileGetTacacsProfileByID struct {
	TacacsProfile *ResponseTacacsProfileGetTacacsProfileByIDTacacsProfile `json:"TacacsProfile,omitempty"` //
}

type ResponseTacacsProfileGetTacacsProfileByIDTacacsProfile struct {
	ID                string                                                                   `json:"id,omitempty"`                //
	Name              string                                                                   `json:"name,omitempty"`              //
	Description       string                                                                   `json:"description,omitempty"`       //
	SessionAttributes *ResponseTacacsProfileGetTacacsProfileByIDTacacsProfileSessionAttributes `json:"sessionAttributes,omitempty"` // Holds list of session attributes. View type for GUI is Shell by default
	Link              *ResponseTacacsProfileGetTacacsProfileByIDTacacsProfileLink              `json:"link,omitempty"`              //
}

type ResponseTacacsProfileGetTacacsProfileByIDTacacsProfileSessionAttributes struct {
	SessionAttributeList *[]ResponseTacacsProfileGetTacacsProfileByIDTacacsProfileSessionAttributesSessionAttributeList `json:"sessionAttributeList,omitempty"` //
}

type ResponseTacacsProfileGetTacacsProfileByIDTacacsProfileSessionAttributesSessionAttributeList struct {
	Type  string `json:"type,omitempty"`  // Allowed values: MANDATORY, OPTIONAL
	Name  string `json:"name,omitempty"`  //
	Value string `json:"value,omitempty"` //
}

type ResponseTacacsProfileGetTacacsProfileByIDTacacsProfileLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsProfileUpdateTacacsProfileByID struct {
	UpdatedFieldsList *ResponseTacacsProfileUpdateTacacsProfileByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseTacacsProfileUpdateTacacsProfileByIDUpdatedFieldsList struct {
	UpdatedField *[]ResponseTacacsProfileUpdateTacacsProfileByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                       `json:"field,omitempty"`        //
	OldValue     string                                                                       `json:"oldValue,omitempty"`     //
	NewValue     string                                                                       `json:"newValue,omitempty"`     //
}

type ResponseTacacsProfileUpdateTacacsProfileByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseTacacsProfileGetTacacsProfile struct {
	SearchResult *ResponseTacacsProfileGetTacacsProfileSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseTacacsProfileGetTacacsProfileSearchResult struct {
	Total        *int                                                           `json:"total,omitempty"`        //
	Resources    *[]ResponseTacacsProfileGetTacacsProfileSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponseTacacsProfileGetTacacsProfileSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseTacacsProfileGetTacacsProfileSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseTacacsProfileGetTacacsProfileSearchResultResources struct {
	ID          string                                                          `json:"id,omitempty"`          //
	Name        string                                                          `json:"name,omitempty"`        //
	Description string                                                          `json:"description,omitempty"` //
	Link        *ResponseTacacsProfileGetTacacsProfileSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseTacacsProfileGetTacacsProfileSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsProfileGetTacacsProfileSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsProfileGetTacacsProfileSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseTacacsProfileGetVersion struct {
	VersionInfo *ResponseTacacsProfileGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseTacacsProfileGetVersionVersionInfo struct {
	CurrentServerVersion string                                          `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                          `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseTacacsProfileGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseTacacsProfileGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestTacacsProfileUpdateTacacsProfileByID struct {
	TacacsProfile *RequestTacacsProfileUpdateTacacsProfileByIDTacacsProfile `json:"TacacsProfile,omitempty"` //
}

type RequestTacacsProfileUpdateTacacsProfileByIDTacacsProfile struct {
	ID                string                                                                     `json:"id,omitempty"`                //
	Name              string                                                                     `json:"name,omitempty"`              //
	Description       string                                                                     `json:"description,omitempty"`       //
	SessionAttributes *RequestTacacsProfileUpdateTacacsProfileByIDTacacsProfileSessionAttributes `json:"sessionAttributes,omitempty"` // Holds list of session attributes. View type for GUI is Shell by default
}

type RequestTacacsProfileUpdateTacacsProfileByIDTacacsProfileSessionAttributes struct {
	SessionAttributeList *[]RequestTacacsProfileUpdateTacacsProfileByIDTacacsProfileSessionAttributesSessionAttributeList `json:"sessionAttributeList,omitempty"` //
}

type RequestTacacsProfileUpdateTacacsProfileByIDTacacsProfileSessionAttributesSessionAttributeList struct {
	Type  string `json:"type,omitempty"`  // Allowed values: MANDATORY, OPTIONAL
	Name  string `json:"name,omitempty"`  //
	Value string `json:"value,omitempty"` //
}

type RequestTacacsProfileCreateTacacsProfile struct {
	TacacsProfile *RequestTacacsProfileCreateTacacsProfileTacacsProfile `json:"TacacsProfile,omitempty"` //
}

type RequestTacacsProfileCreateTacacsProfileTacacsProfile struct {
	Name              string                                                                 `json:"name,omitempty"`              //
	Description       string                                                                 `json:"description,omitempty"`       //
	SessionAttributes *RequestTacacsProfileCreateTacacsProfileTacacsProfileSessionAttributes `json:"sessionAttributes,omitempty"` // Holds list of session attributes. View type for GUI is Shell by default
}

type RequestTacacsProfileCreateTacacsProfileTacacsProfileSessionAttributes struct {
	SessionAttributeList *[]RequestTacacsProfileCreateTacacsProfileTacacsProfileSessionAttributesSessionAttributeList `json:"sessionAttributeList,omitempty"` //
}

type RequestTacacsProfileCreateTacacsProfileTacacsProfileSessionAttributesSessionAttributeList struct {
	Type  string `json:"type,omitempty"`  // Allowed values: MANDATORY, OPTIONAL
	Name  string `json:"name,omitempty"`  //
	Value string `json:"value,omitempty"` //
}

//GetTacacsProfileByName Get TACACS profile by name
/* This API allows the client to get a TACACS profile by name.

@param name name path parameter.
*/
func (s *TacacsProfileService) GetTacacsProfileByName(name string) (*ResponseTacacsProfileGetTacacsProfileByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsprofile/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTacacsProfileGetTacacsProfileByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTacacsProfileByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsProfileGetTacacsProfileByName)
	return result, response, err

}

//GetTacacsProfileByID Get TACACS profile by ID
/* This API allows the client to get a TACACS profile by ID.

@param id id path parameter.
*/
func (s *TacacsProfileService) GetTacacsProfileByID(id string) (*ResponseTacacsProfileGetTacacsProfileByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsprofile/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTacacsProfileGetTacacsProfileByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTacacsProfileById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsProfileGetTacacsProfileByID)
	return result, response, err

}

//GetTacacsProfile Get all TACACS eprofiles
/* This API allows the client to get all the TACACS profiles.

@param getTACACSProfileQueryParams Filtering parameter
*/
func (s *TacacsProfileService) GetTacacsProfile(getTACACSProfileQueryParams *GetTacacsProfileQueryParams) (*ResponseTacacsProfileGetTacacsProfile, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsprofile"

	queryString, _ := query.Values(getTACACSProfileQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTacacsProfileGetTacacsProfile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTacacsProfile")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsProfileGetTacacsProfile)
	return result, response, err

}

//GetVersion Get TACACS profile version information
/* This API helps to retrieve the version information related to the TACACS profile.

 */
func (s *TacacsProfileService) GetVersion() (*ResponseTacacsProfileGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsprofile/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTacacsProfileGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsProfileGetVersion)
	return result, response, err

}

//CreateTacacsProfile Create TACACS profile
/* This API creates a TACACS profile.

 */
func (s *TacacsProfileService) CreateTacacsProfile(requestTacacsProfileCreateTACACSProfile *RequestTacacsProfileCreateTacacsProfile) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsprofile"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTacacsProfileCreateTACACSProfile).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateTacacsProfile")
	}

	return response, err

}

//UpdateTacacsProfileByID Update TACACS profile
/* This API allows the client to update a TACACS profile.

@param id id path parameter.
*/
func (s *TacacsProfileService) UpdateTacacsProfileByID(id string, requestTacacsProfileUpdateTACACSProfileById *RequestTacacsProfileUpdateTacacsProfileByID) (*ResponseTacacsProfileUpdateTacacsProfileByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsprofile/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTacacsProfileUpdateTACACSProfileById).
		SetResult(&ResponseTacacsProfileUpdateTacacsProfileByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseTacacsProfileUpdateTacacsProfileByID{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateTacacsProfileById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTacacsProfileUpdateTacacsProfileByID)
	return result, response, err

}

//DeleteTacacsProfileByID Delete TACACS profile
/* This API deletes a TACACS profile.

@param id id path parameter.
*/
func (s *TacacsProfileService) DeleteTacacsProfileByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/tacacsprofile/{id}"
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
		return response, fmt.Errorf("error with operation DeleteTacacsProfileById")
	}

	getCSFRToken(response.Header())

	return response, err

}
