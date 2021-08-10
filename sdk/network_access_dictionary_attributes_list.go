package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type NetworkAccessDictionaryAttributesListService service

type ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesAuthentication struct {
	Response []ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesAuthenticationResponse `json:"response,omitempty"` //
	Version  string                                                                                            `json:"version,omitempty"`  //
}

type ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesAuthenticationResponse struct {
	AllowedValues  []ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesAuthenticationResponseAllowedValues `json:"allowedValues,omitempty"`  // all of the allowed values for the dictionary attribute
	DataType       string                                                                                                         `json:"dataType,omitempty"`       // the data type for the dictionary attribute
	Description    string                                                                                                         `json:"description,omitempty"`    // The description of the Dictionary attribute
	DictionaryName string                                                                                                         `json:"dictionaryName,omitempty"` // the name of the dictionary which the dictionary attribute belongs to
	DirectionType  string                                                                                                         `json:"directionType,omitempty"`  // the direction for the useage of the dictionary attribute
	ID             string                                                                                                         `json:"id,omitempty"`             // Identifier for the dictionary attribute
	InternalName   string                                                                                                         `json:"internalName,omitempty"`   // the internal name of the dictionary attribute
	Name           string                                                                                                         `json:"name,omitempty"`           // The dictionary attribute's name
}

type ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesAuthenticationResponseAllowedValues struct {
	IsDefault bool   `json:"isDefault,omitempty"` // true if this key value is the default between the allowed values of the dictionary attribute
	Key       string `json:"key,omitempty"`       //
	Value     string `json:"value,omitempty"`     //
}

type ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesAuthorization struct {
	Response []ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesAuthorizationResponse `json:"response,omitempty"` //
	Version  string                                                                                           `json:"version,omitempty"`  //
}

type ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesAuthorizationResponse struct {
	AllowedValues  []ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesAuthorizationResponseAllowedValues `json:"allowedValues,omitempty"`  // all of the allowed values for the dictionary attribute
	DataType       string                                                                                                        `json:"dataType,omitempty"`       // the data type for the dictionary attribute
	Description    string                                                                                                        `json:"description,omitempty"`    // The description of the Dictionary attribute
	DictionaryName string                                                                                                        `json:"dictionaryName,omitempty"` // the name of the dictionary which the dictionary attribute belongs to
	DirectionType  string                                                                                                        `json:"directionType,omitempty"`  // the direction for the useage of the dictionary attribute
	ID             string                                                                                                        `json:"id,omitempty"`             // Identifier for the dictionary attribute
	InternalName   string                                                                                                        `json:"internalName,omitempty"`   // the internal name of the dictionary attribute
	Name           string                                                                                                        `json:"name,omitempty"`           // The dictionary attribute's name
}

type ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesAuthorizationResponseAllowedValues struct {
	IsDefault bool   `json:"isDefault,omitempty"` // true if this key value is the default between the allowed values of the dictionary attribute
	Key       string `json:"key,omitempty"`       //
	Value     string `json:"value,omitempty"`     //
}

type ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesPolicySet struct {
	Response []ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesPolicySetResponse `json:"response,omitempty"` //
	Version  string                                                                                       `json:"version,omitempty"`  //
}

type ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesPolicySetResponse struct {
	AllowedValues  []ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesPolicySetResponseAllowedValues `json:"allowedValues,omitempty"`  // all of the allowed values for the dictionary attribute
	DataType       string                                                                                                    `json:"dataType,omitempty"`       // the data type for the dictionary attribute
	Description    string                                                                                                    `json:"description,omitempty"`    // The description of the Dictionary attribute
	DictionaryName string                                                                                                    `json:"dictionaryName,omitempty"` // the name of the dictionary which the dictionary attribute belongs to
	DirectionType  string                                                                                                    `json:"directionType,omitempty"`  // the direction for the useage of the dictionary attribute
	ID             string                                                                                                    `json:"id,omitempty"`             // Identifier for the dictionary attribute
	InternalName   string                                                                                                    `json:"internalName,omitempty"`   // the internal name of the dictionary attribute
	Name           string                                                                                                    `json:"name,omitempty"`           // The dictionary attribute's name
}

type ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesPolicySetResponseAllowedValues struct {
	IsDefault bool   `json:"isDefault,omitempty"` // true if this key value is the default between the allowed values of the dictionary attribute
	Key       string `json:"key,omitempty"`       //
	Value     string `json:"value,omitempty"`     //
}

//GetNetworkAccessDictionariesAuthentication Network Access - Returns list of dictionary attributes for authentication.
/* Network Access Returns list of dictionary attributes for authentication.

 */
func (s *NetworkAccessDictionaryAttributesListService) GetNetworkAccessDictionariesAuthentication() (*ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesAuthentication, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/dictionaries/authentication"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesAuthentication{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessDictionariesAuthentication")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesAuthentication)
	return result, response, err

}

//GetNetworkAccessDictionariesAuthorization Network Access - Returns list of dictionary attributes for authorization.
/* Network Access Returns list of dictionary attributes for authorization.

 */
func (s *NetworkAccessDictionaryAttributesListService) GetNetworkAccessDictionariesAuthorization() (*ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesAuthorization, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/dictionaries/authorization"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesAuthorization{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessDictionariesAuthorization")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesAuthorization)
	return result, response, err

}

//GetNetworkAccessDictionariesPolicySet Network Access - Returns list of dictionary attributes for policyset.
/* Network Access Returns list of dictionary attributes for policyset.

 */
func (s *NetworkAccessDictionaryAttributesListService) GetNetworkAccessDictionariesPolicySet() (*ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesPolicySet, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/dictionaries/policyset"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesPolicySet{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessDictionariesPolicySet")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesPolicySet)
	return result, response, err

}
