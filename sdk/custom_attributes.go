package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type CustomAttributesService service

type ResponseCustomAttributesList []ResponseItemCustomAttributesList // Array of ResponseCustomAttributesList

type ResponseItemCustomAttributesList struct {
	AttributeName string `json:"attributeName,omitempty"` //
	AttributeType string `json:"attributeType,omitempty"` //
}

// # Review unknown case
// # Review unknown case

type ResponseCustomAttributesGet struct {
	AttributeName string `json:"attributeName,omitempty"` //
	AttributeType string `json:"attributeType,omitempty"` //
}

// # Review unknown case

type RequestCustomAttributesCreateCustomAttribute struct {
	AttributeName string `json:"attributeName,omitempty"` //
	AttributeType string `json:"attributeType,omitempty"` //
}

type RequestCustomAttributesRename struct {
	CurrentName string `json:"currentName,omitempty"` //
	NewName     string `json:"newName,omitempty"`     //
}

//List Get all custom attributes
/* Get all custom attributes

 */
func (s *CustomAttributesService) List() (*ResponseCustomAttributesList, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/endpoint-custom-attribute"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCustomAttributesList{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation List")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCustomAttributesList)
	return result, response, err

}

//Get Get custom attribute by name
/* Get custom attribute by name

@param name name path parameter. Name of the custom attribute
*/
func (s *CustomAttributesService) Get(name string) (*ResponseCustomAttributesGet, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/endpoint-custom-attribute/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCustomAttributesGet{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation Get")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCustomAttributesGet)
	return result, response, err

}

//CreateCustomAttribute Create Custom Attribute
/* Create Custom Attribute

 */
func (s *CustomAttributesService) CreateCustomAttribute(requestCustomAttributesCreateCustomAttribute *RequestCustomAttributesCreateCustomAttribute) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/endpoint-custom-attribute"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCustomAttributesCreateCustomAttribute).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateCustomAttribute")
	}

	return response, err

}

//Rename rename custom attribute
/* rename custom attribute

 */
func (s *CustomAttributesService) Rename(requestCustomAttributesRename *RequestCustomAttributesRename) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/endpoint-custom-attribute/rename"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCustomAttributesRename).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation Rename")
	}

	return response, err

}

//Delete Delete custom attribute by name
/* Delete custom attribute by name

@param name name path parameter. The name of the custom attribute
*/
func (s *CustomAttributesService) Delete(name string) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/endpoint-custom-attribute/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

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
		return response, fmt.Errorf("error with operation Delete")
	}

	getCSFRToken(response.Header())

	return response, err

}
