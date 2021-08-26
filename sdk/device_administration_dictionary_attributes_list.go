package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type DeviceAdministrationDictionaryAttributesListService service

type ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthentication struct {
	Response *[]ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthenticationResponse `json:"response,omitempty"` //
	Version  string                                                                                                  `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthenticationResponse struct {
	AllowedValues  *[]ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthenticationResponseAllowedValues `json:"allowedValues,omitempty"`  // all of the allowed values for the dictionary attribute
	DataType       string                                                                                                               `json:"dataType,omitempty"`       // the data type for the dictionary attribute
	Description    string                                                                                                               `json:"description,omitempty"`    // The description of the Dictionary attribute
	DictionaryName string                                                                                                               `json:"dictionaryName,omitempty"` // the name of the dictionary which the dictionary attribute belongs to
	DirectionType  string                                                                                                               `json:"directionType,omitempty"`  // the direction for the useage of the dictionary attribute
	ID             string                                                                                                               `json:"id,omitempty"`             // Identifier for the dictionary attribute
	InternalName   string                                                                                                               `json:"internalName,omitempty"`   // the internal name of the dictionary attribute
	Name           string                                                                                                               `json:"name,omitempty"`           // The dictionary attribute's name
}

type ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthenticationResponseAllowedValues struct {
	IsDefault *bool  `json:"isDefault,omitempty"` // true if this key value is the default between the allowed values of the dictionary attribute
	Key       string `json:"key,omitempty"`       //
	Value     string `json:"value,omitempty"`     //
}

type ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthorization struct {
	Response *[]ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthorizationResponse `json:"response,omitempty"` //
	Version  string                                                                                                 `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthorizationResponse struct {
	AllowedValues  *[]ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthorizationResponseAllowedValues `json:"allowedValues,omitempty"`  // all of the allowed values for the dictionary attribute
	DataType       string                                                                                                              `json:"dataType,omitempty"`       // the data type for the dictionary attribute
	Description    string                                                                                                              `json:"description,omitempty"`    // The description of the Dictionary attribute
	DictionaryName string                                                                                                              `json:"dictionaryName,omitempty"` // the name of the dictionary which the dictionary attribute belongs to
	DirectionType  string                                                                                                              `json:"directionType,omitempty"`  // the direction for the useage of the dictionary attribute
	ID             string                                                                                                              `json:"id,omitempty"`             // Identifier for the dictionary attribute
	InternalName   string                                                                                                              `json:"internalName,omitempty"`   // the internal name of the dictionary attribute
	Name           string                                                                                                              `json:"name,omitempty"`           // The dictionary attribute's name
}

type ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthorizationResponseAllowedValues struct {
	IsDefault *bool  `json:"isDefault,omitempty"` // true if this key value is the default between the allowed values of the dictionary attribute
	Key       string `json:"key,omitempty"`       //
	Value     string `json:"value,omitempty"`     //
}

type ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesPolicySet struct {
	Response *[]ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesPolicySetResponse `json:"response,omitempty"` //
	Version  string                                                                                             `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesPolicySetResponse struct {
	AllowedValues  *[]ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesPolicySetResponseAllowedValues `json:"allowedValues,omitempty"`  // all of the allowed values for the dictionary attribute
	DataType       string                                                                                                          `json:"dataType,omitempty"`       // the data type for the dictionary attribute
	Description    string                                                                                                          `json:"description,omitempty"`    // The description of the Dictionary attribute
	DictionaryName string                                                                                                          `json:"dictionaryName,omitempty"` // the name of the dictionary which the dictionary attribute belongs to
	DirectionType  string                                                                                                          `json:"directionType,omitempty"`  // the direction for the useage of the dictionary attribute
	ID             string                                                                                                          `json:"id,omitempty"`             // Identifier for the dictionary attribute
	InternalName   string                                                                                                          `json:"internalName,omitempty"`   // the internal name of the dictionary attribute
	Name           string                                                                                                          `json:"name,omitempty"`           // The dictionary attribute's name
}

type ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesPolicySetResponseAllowedValues struct {
	IsDefault *bool  `json:"isDefault,omitempty"` // true if this key value is the default between the allowed values of the dictionary attribute
	Key       string `json:"key,omitempty"`       //
	Value     string `json:"value,omitempty"`     //
}

//GetDeviceAdminDictionariesAuthentication Network Access - Returns list of dictionary attributes for authentication.
/* Network Access Returns list of dictionary attributes for authentication.

 */
func (s *DeviceAdministrationDictionaryAttributesListService) GetDeviceAdminDictionariesAuthentication() (*ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthentication, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/dictionaries/authentication"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthentication{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminDictionariesAuthentication")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthentication)
	return result, response, err

}

//GetDeviceAdminDictionariesAuthorization Network Access - Returns list of dictionary attributes for authorization.
/* Network Access Returns list of dictionary attributes for authorization.

 */
func (s *DeviceAdministrationDictionaryAttributesListService) GetDeviceAdminDictionariesAuthorization() (*ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthorization, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/dictionaries/authorization"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthorization{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminDictionariesAuthorization")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthorization)
	return result, response, err

}

//GetDeviceAdminDictionariesPolicySet Network Access - Returns list of dictionary attributes for policyset.
/* Network Access Returns list of dictionary attributes for policyset.

 */
func (s *DeviceAdministrationDictionaryAttributesListService) GetDeviceAdminDictionariesPolicySet() (*ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesPolicySet, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/dictionaries/policyset"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesPolicySet{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminDictionariesPolicySet")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesPolicySet)
	return result, response, err

}
