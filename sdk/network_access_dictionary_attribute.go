package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type NetworkAccessDictionaryAttributeService service

type ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryName struct {
	Response *[]ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryNameResponse `json:"response,omitempty"` //
	Version  string                                                                                                  `json:"version,omitempty"`  //
}

type ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryNameResponse struct {
	AllowedValues  *[]ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryNameResponseAllowedValues `json:"allowedValues,omitempty"`  // all of the allowed values for the dictionary attribute
	DataType       string                                                                                                               `json:"dataType,omitempty"`       // the data type for the dictionary attribute
	Description    string                                                                                                               `json:"description,omitempty"`    // The description of the Dictionary attribute
	DictionaryName string                                                                                                               `json:"dictionaryName,omitempty"` // the name of the dictionary which the dictionary attribute belongs to
	DirectionType  string                                                                                                               `json:"directionType,omitempty"`  // the direction for the useage of the dictionary attribute
	ID             string                                                                                                               `json:"id,omitempty"`             // Identifier for the dictionary attribute
	InternalName   string                                                                                                               `json:"internalName,omitempty"`   // the internal name of the dictionary attribute
	Name           string                                                                                                               `json:"name,omitempty"`           // The dictionary attribute's name
}

type ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryNameResponseAllowedValues struct {
	IsDefault *bool  `json:"isDefault,omitempty"` // true if this key value is the default between the allowed values of the dictionary attribute
	Key       string `json:"key,omitempty"`       //
	Value     string `json:"value,omitempty"`     //
}

type ResponseNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttribute struct {
	Response *ResponseNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttributeResponse `json:"response,omitempty"` // Dictionary Attribute format
	Version  string                                                                                  `json:"version,omitempty"`  //
}

type ResponseNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttributeResponse struct {
	AllowedValues  *[]ResponseNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttributeResponseAllowedValues `json:"allowedValues,omitempty"`  // all of the allowed values for the dictionary attribute
	DataType       string                                                                                                 `json:"dataType,omitempty"`       // the data type for the dictionary attribute
	Description    string                                                                                                 `json:"description,omitempty"`    // The description of the Dictionary attribute
	DictionaryName string                                                                                                 `json:"dictionaryName,omitempty"` // the name of the dictionary which the dictionary attribute belongs to
	DirectionType  string                                                                                                 `json:"directionType,omitempty"`  // the direction for the useage of the dictionary attribute
	ID             string                                                                                                 `json:"id,omitempty"`             // Identifier for the dictionary attribute
	InternalName   string                                                                                                 `json:"internalName,omitempty"`   // the internal name of the dictionary attribute
	Name           string                                                                                                 `json:"name,omitempty"`           // The dictionary attribute's name
}

type ResponseNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttributeResponseAllowedValues struct {
	IsDefault *bool  `json:"isDefault,omitempty"` // true if this key value is the default between the allowed values of the dictionary attribute
	Key       string `json:"key,omitempty"`       //
	Value     string `json:"value,omitempty"`     //
}

type ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByName struct {
	Response *ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByNameResponse `json:"response,omitempty"` // Dictionary Attribute format
	Version  string                                                                                     `json:"version,omitempty"`  //
}

type ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByNameResponse struct {
	AllowedValues  *[]ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByNameResponseAllowedValues `json:"allowedValues,omitempty"`  // all of the allowed values for the dictionary attribute
	DataType       string                                                                                                    `json:"dataType,omitempty"`       // the data type for the dictionary attribute
	Description    string                                                                                                    `json:"description,omitempty"`    // The description of the Dictionary attribute
	DictionaryName string                                                                                                    `json:"dictionaryName,omitempty"` // the name of the dictionary which the dictionary attribute belongs to
	DirectionType  string                                                                                                    `json:"directionType,omitempty"`  // the direction for the useage of the dictionary attribute
	ID             string                                                                                                    `json:"id,omitempty"`             // Identifier for the dictionary attribute
	InternalName   string                                                                                                    `json:"internalName,omitempty"`   // the internal name of the dictionary attribute
	Name           string                                                                                                    `json:"name,omitempty"`           // The dictionary attribute's name
}

type ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByNameResponseAllowedValues struct {
	IsDefault *bool  `json:"isDefault,omitempty"` // true if this key value is the default between the allowed values of the dictionary attribute
	Key       string `json:"key,omitempty"`       //
	Value     string `json:"value,omitempty"`     //
}

type ResponseNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByName struct {
	Response *ResponseNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByNameResponse `json:"response,omitempty"` // Dictionary Attribute format
	Version  string                                                                                        `json:"version,omitempty"`  //
}

type ResponseNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByNameResponse struct {
	AllowedValues  *[]ResponseNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByNameResponseAllowedValues `json:"allowedValues,omitempty"`  // all of the allowed values for the dictionary attribute
	DataType       string                                                                                                       `json:"dataType,omitempty"`       // the data type for the dictionary attribute
	Description    string                                                                                                       `json:"description,omitempty"`    // The description of the Dictionary attribute
	DictionaryName string                                                                                                       `json:"dictionaryName,omitempty"` // the name of the dictionary which the dictionary attribute belongs to
	DirectionType  string                                                                                                       `json:"directionType,omitempty"`  // the direction for the useage of the dictionary attribute
	ID             string                                                                                                       `json:"id,omitempty"`             // Identifier for the dictionary attribute
	InternalName   string                                                                                                       `json:"internalName,omitempty"`   // the internal name of the dictionary attribute
	Name           string                                                                                                       `json:"name,omitempty"`           // The dictionary attribute's name
}

type ResponseNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByNameResponseAllowedValues struct {
	IsDefault *bool  `json:"isDefault,omitempty"` // true if this key value is the default between the allowed values of the dictionary attribute
	Key       string `json:"key,omitempty"`       //
	Value     string `json:"value,omitempty"`     //
}

type ResponseNetworkAccessDictionaryAttributeDeleteNetworkAccessDictionaryAttributeByName struct {
	Response *ResponseNetworkAccessDictionaryAttributeDeleteNetworkAccessDictionaryAttributeByNameResponse `json:"response,omitempty"` // Dictionary Attribute format
	Version  string                                                                                        `json:"version,omitempty"`  //
}

type ResponseNetworkAccessDictionaryAttributeDeleteNetworkAccessDictionaryAttributeByNameResponse struct {
	AllowedValues  *[]ResponseNetworkAccessDictionaryAttributeDeleteNetworkAccessDictionaryAttributeByNameResponseAllowedValues `json:"allowedValues,omitempty"`  // all of the allowed values for the dictionary attribute
	DataType       string                                                                                                       `json:"dataType,omitempty"`       // the data type for the dictionary attribute
	Description    string                                                                                                       `json:"description,omitempty"`    // The description of the Dictionary attribute
	DictionaryName string                                                                                                       `json:"dictionaryName,omitempty"` // the name of the dictionary which the dictionary attribute belongs to
	DirectionType  string                                                                                                       `json:"directionType,omitempty"`  // the direction for the useage of the dictionary attribute
	ID             string                                                                                                       `json:"id,omitempty"`             // Identifier for the dictionary attribute
	InternalName   string                                                                                                       `json:"internalName,omitempty"`   // the internal name of the dictionary attribute
	Name           string                                                                                                       `json:"name,omitempty"`           // The dictionary attribute's name
}

type ResponseNetworkAccessDictionaryAttributeDeleteNetworkAccessDictionaryAttributeByNameResponseAllowedValues struct {
	IsDefault *bool  `json:"isDefault,omitempty"` // true if this key value is the default between the allowed values of the dictionary attribute
	Key       string `json:"key,omitempty"`       //
	Value     string `json:"value,omitempty"`     //
}

type RequestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttribute struct {
	AllowedValues  *[]RequestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttributeAllowedValues `json:"allowedValues,omitempty"`  // all of the allowed values for the dictionary attribute
	DataType       string                                                                                        `json:"dataType,omitempty"`       // the data type for the dictionary attribute
	Description    string                                                                                        `json:"description,omitempty"`    // The description of the Dictionary attribute
	DictionaryName string                                                                                        `json:"dictionaryName,omitempty"` // the name of the dictionary which the dictionary attribute belongs to
	DirectionType  string                                                                                        `json:"directionType,omitempty"`  // the direction for the useage of the dictionary attribute
	ID             string                                                                                        `json:"id,omitempty"`             // Identifier for the dictionary attribute
	InternalName   string                                                                                        `json:"internalName,omitempty"`   // the internal name of the dictionary attribute
	Name           string                                                                                        `json:"name,omitempty"`           // The dictionary attribute's name
}

type RequestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttributeAllowedValues struct {
	IsDefault *bool  `json:"isDefault,omitempty"` // true if this key value is the default between the allowed values of the dictionary attribute
	Key       string `json:"key,omitempty"`       //
	Value     string `json:"value,omitempty"`     //
}

type RequestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByName struct {
	AllowedValues  *[]RequestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByNameAllowedValues `json:"allowedValues,omitempty"`  // all of the allowed values for the dictionary attribute
	DataType       string                                                                                              `json:"dataType,omitempty"`       // the data type for the dictionary attribute
	Description    string                                                                                              `json:"description,omitempty"`    // The description of the Dictionary attribute
	DictionaryName string                                                                                              `json:"dictionaryName,omitempty"` // the name of the dictionary which the dictionary attribute belongs to
	DirectionType  string                                                                                              `json:"directionType,omitempty"`  // the direction for the useage of the dictionary attribute
	ID             string                                                                                              `json:"id,omitempty"`             // Identifier for the dictionary attribute
	InternalName   string                                                                                              `json:"internalName,omitempty"`   // the internal name of the dictionary attribute
	Name           string                                                                                              `json:"name,omitempty"`           // The dictionary attribute's name
}

type RequestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByNameAllowedValues struct {
	IsDefault *bool  `json:"isDefault,omitempty"` // true if this key value is the default between the allowed values of the dictionary attribute
	Key       string `json:"key,omitempty"`       //
	Value     string `json:"value,omitempty"`     //
}

//GetNetworkAccessDictionaryAttributesByDictionaryName Returns a list of Dictionary Attributes for an existing Dictionary.
/* Returns a list of Dictionary Attributes for an existing Dictionary.

@param dictionaryName dictionaryName path parameter. the name of the dictionary the dictionary attribute belongs to
*/
func (s *NetworkAccessDictionaryAttributeService) GetNetworkAccessDictionaryAttributesByDictionaryName(dictionaryName string) (*ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryName, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/dictionaries/{dictionaryName}/attribute"
	path = strings.Replace(path, "{dictionaryName}", fmt.Sprintf("%v", dictionaryName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessDictionaryAttributesByDictionaryName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryName)
	return result, response, err

}

//GetNetworkAccessDictionaryAttributeByName Get a Dictionary Attribute.
/* Get a Dictionary Attribute.

@param name name path parameter. the dictionary attribute name
@param dictionaryName dictionaryName path parameter. the name of the dictionary the dictionary attribute belongs to
*/
func (s *NetworkAccessDictionaryAttributeService) GetNetworkAccessDictionaryAttributeByName(name string, dictionaryName string) (*ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByName, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/dictionaries/{dictionaryName}/attribute/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)
	path = strings.Replace(path, "{dictionaryName}", fmt.Sprintf("%v", dictionaryName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessDictionaryAttributeByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByName)
	return result, response, err

}

//CreateNetworkAccessDictionaryAttribute Create a new Dictionary Attribute for an existing Dictionary.
/* Create a new Dictionary Attribute for an existing Dictionary.

@param dictionaryName dictionaryName path parameter. the name of the dictionary the dictionary attribute belongs to
*/
func (s *NetworkAccessDictionaryAttributeService) CreateNetworkAccessDictionaryAttribute(dictionaryName string, requestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttribute *RequestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttribute) (*ResponseNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttribute, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/dictionaries/{dictionaryName}/attribute"
	path = strings.Replace(path, "{dictionaryName}", fmt.Sprintf("%v", dictionaryName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttribute).
		SetResult(&ResponseNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttribute{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkAccessDictionaryAttribute")
	}

	result := response.Result().(*ResponseNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttribute)
	return result, response, err

}

//UpdateNetworkAccessDictionaryAttributeByName Update a Dictionary Attribute.
/* Update a Dictionary Attribute

@param name name path parameter. the dictionary attribute name
@param dictionaryName dictionaryName path parameter. the name of the dictionary the dictionary attribute belongs to
*/
func (s *NetworkAccessDictionaryAttributeService) UpdateNetworkAccessDictionaryAttributeByName(name string, dictionaryName string, requestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByName *RequestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByName) (*ResponseNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByName, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/dictionaries/{dictionaryName}/attribute/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)
	path = strings.Replace(path, "{dictionaryName}", fmt.Sprintf("%v", dictionaryName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByName).
		SetResult(&ResponseNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByName{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkAccessDictionaryAttributeByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByName)
	return result, response, err

}

//DeleteNetworkAccessDictionaryAttributeByName Delete a Dictionary Attribute.
/* Delete a Dictionary Attribute.

@param name name path parameter. the dictionary attribute name
@param dictionaryName dictionaryName path parameter. the name of the dictionary the dictionary attribute belongs to
*/
func (s *NetworkAccessDictionaryAttributeService) DeleteNetworkAccessDictionaryAttributeByName(name string, dictionaryName string) (*ResponseNetworkAccessDictionaryAttributeDeleteNetworkAccessDictionaryAttributeByName, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/dictionaries/{dictionaryName}/attribute/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)
	path = strings.Replace(path, "{dictionaryName}", fmt.Sprintf("%v", dictionaryName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessDictionaryAttributeDeleteNetworkAccessDictionaryAttributeByName{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteNetworkAccessDictionaryAttributeByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessDictionaryAttributeDeleteNetworkAccessDictionaryAttributeByName)
	return result, response, err

}
