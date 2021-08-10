package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type NetworkAccessDictionaryService service

type ResponseNetworkAccessDictionaryGetNetworkAccessDictionaries struct {
	Response []ResponseNetworkAccessDictionaryGetNetworkAccessDictionariesResponse `json:"response,omitempty"` //
	Version  string                                                                `json:"version,omitempty"`  //
}

type ResponseNetworkAccessDictionaryGetNetworkAccessDictionariesResponse struct {
	Description        string                                                                  `json:"description,omitempty"`        // The description of the Dictionary
	DictionaryAttrType string                                                                  `json:"dictionaryAttrType,omitempty"` // The dictionary attribute type
	ID                 string                                                                  `json:"id,omitempty"`                 // Identifier for the dictionary
	Link               ResponseNetworkAccessDictionaryGetNetworkAccessDictionariesResponseLink `json:"link,omitempty"`               //
	Name               string                                                                  `json:"name,omitempty"`               // The dictionary name
	Version            string                                                                  `json:"version,omitempty"`            // The dictionary version
}

type ResponseNetworkAccessDictionaryGetNetworkAccessDictionariesResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessDictionaryCreateNetworkAccessDictionaries struct {
	Response ResponseNetworkAccessDictionaryCreateNetworkAccessDictionariesResponse `json:"response,omitempty"` // Dictionary POST format
	Version  string                                                                 `json:"version,omitempty"`  //
}

type ResponseNetworkAccessDictionaryCreateNetworkAccessDictionariesResponse struct {
	Description        string                                                                     `json:"description,omitempty"`        // The description of the Dictionary
	DictionaryAttrType string                                                                     `json:"dictionaryAttrType,omitempty"` // The dictionary attribute type
	ID                 string                                                                     `json:"id,omitempty"`                 // Identifier for the dictionary
	Link               ResponseNetworkAccessDictionaryCreateNetworkAccessDictionariesResponseLink `json:"link,omitempty"`               //
	Name               string                                                                     `json:"name,omitempty"`               // The dictionary name
	Version            string                                                                     `json:"version,omitempty"`            // The dictionary version
}

type ResponseNetworkAccessDictionaryCreateNetworkAccessDictionariesResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByName struct {
	Response ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByNameResponse `json:"response,omitempty"` // Dictionary POST format
	Version  string                                                                  `json:"version,omitempty"`  //
}

type ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByNameResponse struct {
	Description        string                                                                      `json:"description,omitempty"`        // The description of the Dictionary
	DictionaryAttrType string                                                                      `json:"dictionaryAttrType,omitempty"` // The dictionary attribute type
	ID                 string                                                                      `json:"id,omitempty"`                 // Identifier for the dictionary
	Link               ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByNameResponseLink `json:"link,omitempty"`               //
	Name               string                                                                      `json:"name,omitempty"`               // The dictionary name
	Version            string                                                                      `json:"version,omitempty"`            // The dictionary version
}

type ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByNameResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessDictionaryUpdateNetworkAccessDictionaryByName struct {
	Response ResponseNetworkAccessDictionaryUpdateNetworkAccessDictionaryByNameResponse `json:"response,omitempty"` // Dictionary POST format
	Version  string                                                                     `json:"version,omitempty"`  //
}

type ResponseNetworkAccessDictionaryUpdateNetworkAccessDictionaryByNameResponse struct {
	Description        string                                                                         `json:"description,omitempty"`        // The description of the Dictionary
	DictionaryAttrType string                                                                         `json:"dictionaryAttrType,omitempty"` // The dictionary attribute type
	ID                 string                                                                         `json:"id,omitempty"`                 // Identifier for the dictionary
	Link               ResponseNetworkAccessDictionaryUpdateNetworkAccessDictionaryByNameResponseLink `json:"link,omitempty"`               //
	Name               string                                                                         `json:"name,omitempty"`               // The dictionary name
	Version            string                                                                         `json:"version,omitempty"`            // The dictionary version
}

type ResponseNetworkAccessDictionaryUpdateNetworkAccessDictionaryByNameResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessDictionaryDeleteNetworkAccessDictionaryByName struct {
	Response ResponseNetworkAccessDictionaryDeleteNetworkAccessDictionaryByNameResponse `json:"response,omitempty"` // Dictionary POST format
	Version  string                                                                     `json:"version,omitempty"`  //
}

type ResponseNetworkAccessDictionaryDeleteNetworkAccessDictionaryByNameResponse struct {
	Description        string                                                                         `json:"description,omitempty"`        // The description of the Dictionary
	DictionaryAttrType string                                                                         `json:"dictionaryAttrType,omitempty"` // The dictionary attribute type
	ID                 string                                                                         `json:"id,omitempty"`                 // Identifier for the dictionary
	Link               ResponseNetworkAccessDictionaryDeleteNetworkAccessDictionaryByNameResponseLink `json:"link,omitempty"`               //
	Name               string                                                                         `json:"name,omitempty"`               // The dictionary name
	Version            string                                                                         `json:"version,omitempty"`            // The dictionary version
}

type ResponseNetworkAccessDictionaryDeleteNetworkAccessDictionaryByNameResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessDictionaryCreateNetworkAccessDictionaries struct {
	Description        string                                                             `json:"description,omitempty"`        // The description of the Dictionary
	DictionaryAttrType string                                                             `json:"dictionaryAttrType,omitempty"` // The dictionary attribute type
	ID                 string                                                             `json:"id,omitempty"`                 // Identifier for the dictionary
	Link               *RequestNetworkAccessDictionaryCreateNetworkAccessDictionariesLink `json:"link,omitempty"`               //
	Name               string                                                             `json:"name,omitempty"`               // The dictionary name
	Version            string                                                             `json:"version,omitempty"`            // The dictionary version
}

type RequestNetworkAccessDictionaryCreateNetworkAccessDictionariesLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByName struct {
	Description        string                                                                 `json:"description,omitempty"`        // The description of the Dictionary
	DictionaryAttrType string                                                                 `json:"dictionaryAttrType,omitempty"` // The dictionary attribute type
	ID                 string                                                                 `json:"id,omitempty"`                 // Identifier for the dictionary
	Link               *RequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByNameLink `json:"link,omitempty"`               //
	Name               string                                                                 `json:"name,omitempty"`               // The dictionary name
	Version            string                                                                 `json:"version,omitempty"`            // The dictionary version
}

type RequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByNameLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

//GetNetworkAccessDictionaries Network Access - Returns a list of Dictionaries.
/* Get all Dictionaries.

 */
func (s *NetworkAccessDictionaryService) GetNetworkAccessDictionaries() (*ResponseNetworkAccessDictionaryGetNetworkAccessDictionaries, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/dictionaries"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessDictionaryGetNetworkAccessDictionaries{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessDictionaries")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessDictionaryGetNetworkAccessDictionaries)
	return result, response, err

}

//GetNetworkAccessDictionaryByName GET a dictionary by name
/* GET a dictionary by name

@param name name path parameter. the dictionary name
*/
func (s *NetworkAccessDictionaryService) GetNetworkAccessDictionaryByName(name string) (*ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByName, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/dictionaries/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessDictionaryByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByName)
	return result, response, err

}

//CreateNetworkAccessDictionaries Network Access - Create a new Dictionary.
/* Network Access Create a new Dictionary.

 */
func (s *NetworkAccessDictionaryService) CreateNetworkAccessDictionaries(requestNetworkAccessDictionaryCreateNetworkAccessDictionaries *RequestNetworkAccessDictionaryCreateNetworkAccessDictionaries) (*ResponseNetworkAccessDictionaryCreateNetworkAccessDictionaries, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/dictionaries"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessDictionaryCreateNetworkAccessDictionaries).
		SetResult(&ResponseNetworkAccessDictionaryCreateNetworkAccessDictionaries{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkAccessDictionaries")
	}

	result := response.Result().(*ResponseNetworkAccessDictionaryCreateNetworkAccessDictionaries)
	return result, response, err

}

//UpdateNetworkAccessDictionaryByName Network Access - Update a Dictionary.
/* Network Access Update a Dictionary.

@param name name path parameter. the dictionary name
*/
func (s *NetworkAccessDictionaryService) UpdateNetworkAccessDictionaryByName(name string, requestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByName *RequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByName) (*ResponseNetworkAccessDictionaryUpdateNetworkAccessDictionaryByName, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/dictionaries/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByName).
		SetResult(&ResponseNetworkAccessDictionaryUpdateNetworkAccessDictionaryByName{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkAccessDictionaryByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessDictionaryUpdateNetworkAccessDictionaryByName)
	return result, response, err

}

//DeleteNetworkAccessDictionaryByName Network Access - Delete a Dictionary.
/* Network Access Delete a Dictionary.

@param name name path parameter. the dictionary name
*/
func (s *NetworkAccessDictionaryService) DeleteNetworkAccessDictionaryByName(name string) (*ResponseNetworkAccessDictionaryDeleteNetworkAccessDictionaryByName, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/dictionaries/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessDictionaryDeleteNetworkAccessDictionaryByName{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteNetworkAccessDictionaryByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessDictionaryDeleteNetworkAccessDictionaryByName)
	return result, response, err

}
