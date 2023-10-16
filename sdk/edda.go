package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type EddaService service

type ResponseEddaGetConnectorConfig struct {
	Response *[]ResponseEddaGetConnectorConfigResponse `json:"response,omitempty"` //
	Version  string                                    `json:"version,omitempty"`  //
}

type ResponseEddaGetConnectorConfigResponse struct {
	Connector *ResponseEddaGetConnectorConfigResponseConnector `json:"connector,omitempty"` // connectorName
}

type ResponseEddaGetConnectorConfigResponseConnector struct {
	AdditionalProperties       *ResponseEddaGetConnectorConfigResponseConnectorAdditionalProperties `json:"additionalProperties,omitempty"`       //
	Attributes                 *ResponseEddaGetConnectorConfigResponseConnectorAttributes           `json:"attributes,omitempty"`                 // connectorName
	ConnectorName              string                                                               `json:"connectorName,omitempty"`              // connectorName
	ConnectorType              string                                                               `json:"connectorType,omitempty"`              // connector Type list
	DeltasyncSchedule          *ResponseEddaGetConnectorConfigResponseConnectorDeltasyncSchedule    `json:"deltasyncSchedule,omitempty"`          //
	Description                string                                                               `json:"description,omitempty"`                // description
	Enabled                    *bool                                                                `json:"enabled,omitempty"`                    //
	FullsyncSchedule           *ResponseEddaGetConnectorConfigResponseConnectorFullsyncSchedule     `json:"fullsyncSchedule,omitempty"`           //
	Protocol                   string                                                               `json:"protocol,omitempty"`                   // protocol
	SkipCertificateValidations *bool                                                                `json:"skipCertificateValidations,omitempty"` //
	URL                        *ResponseEddaGetConnectorConfigResponseConnectorURL                  `json:"url,omitempty"`                        //
}

type ResponseEddaGetConnectorConfigResponseConnectorAdditionalProperties interface{}

type ResponseEddaGetConnectorConfigResponseConnectorAttributes struct {
	AttributeMapping     *[]ResponseEddaGetConnectorConfigResponseConnectorAttributesAttributeMapping `json:"attributeMapping,omitempty"`     // <p>List of feature names</p>
	BulkUniqueIDentifier string                                                                       `json:"bulkUniqueIdentifier,omitempty"` // uniqueness to identify
	TopLevelObject       string                                                                       `json:"topLevelObject,omitempty"`       // root level of json
	UniqueIDentifier     string                                                                       `json:"uniqueIdentifier,omitempty"`     // uniqueness to identify
	VersionIDentifier    string                                                                       `json:"versionIdentifier,omitempty"`    // version uniqueness to identify
}

type ResponseEddaGetConnectorConfigResponseConnectorAttributesAttributeMapping struct {
	DictionaryAttribute string `json:"dictionaryAttribute,omitempty"` //
	IncludeInDictionary *bool  `json:"includeInDictionary,omitempty"` //
	JSONAttribute       string `json:"jsonAttribute,omitempty"`       //
}

type ResponseEddaGetConnectorConfigResponseConnectorDeltasyncSchedule struct {
	Interval     *int   `json:"interval,omitempty"`     // run at interval (hours)
	IntervalUnit string `json:"intervalUnit,omitempty"` // interval Units
	StartDate    string `json:"startDate,omitempty"`    // start date and Time
}

type ResponseEddaGetConnectorConfigResponseConnectorFullsyncSchedule struct {
	Interval     *int   `json:"interval,omitempty"`     // run at interval (hours)
	IntervalUnit string `json:"intervalUnit,omitempty"` // interval Units
	StartDate    string `json:"startDate,omitempty"`    // start date and Time
}

type ResponseEddaGetConnectorConfigResponseConnectorURL struct {
	AccessKey          string `json:"accessKey,omitempty"`          // accesskey
	AuthenticationType string `json:"authenticationType,omitempty"` // authentication Type list
	BulkURL            string `json:"bulkUrl,omitempty"`            // bulkUrl
	ClientID           string `json:"clientId,omitempty"`           // clientid
	ClientSecret       string `json:"clientSecret,omitempty"`       // clientsecret
	IncrementalURL     string `json:"incrementalUrl,omitempty"`     // incrementalUrl
	Password           string `json:"password,omitempty"`           // password
	RefreshToken       string `json:"refreshToken,omitempty"`       // refreshtoken
	TokenHeader        string `json:"tokenHeader,omitempty"`        // tokenHeader
	UserName           string `json:"userName,omitempty"`           // userName
}

// # Review unknown case

type ResponseEddaGetConnectorConfigByConnectorName struct {
	Response *ResponseEddaGetConnectorConfigByConnectorNameResponse `json:"response,omitempty"` // ConnectorConfig information format
	Version  string                                                 `json:"version,omitempty"`  //
}

type ResponseEddaGetConnectorConfigByConnectorNameResponse struct {
	Connector *ResponseEddaGetConnectorConfigByConnectorNameResponseConnector `json:"connector,omitempty"` // connectorName
}

type ResponseEddaGetConnectorConfigByConnectorNameResponseConnector struct {
	AdditionalProperties       *ResponseEddaGetConnectorConfigByConnectorNameResponseConnectorAdditionalProperties `json:"additionalProperties,omitempty"`       //
	Attributes                 *ResponseEddaGetConnectorConfigByConnectorNameResponseConnectorAttributes           `json:"attributes,omitempty"`                 // connectorName
	ConnectorName              string                                                                              `json:"connectorName,omitempty"`              // connectorName
	ConnectorType              string                                                                              `json:"connectorType,omitempty"`              // connector Type list
	DeltasyncSchedule          *ResponseEddaGetConnectorConfigByConnectorNameResponseConnectorDeltasyncSchedule    `json:"deltasyncSchedule,omitempty"`          //
	Description                string                                                                              `json:"description,omitempty"`                // description
	Enabled                    *bool                                                                               `json:"enabled,omitempty"`                    //
	FullsyncSchedule           *ResponseEddaGetConnectorConfigByConnectorNameResponseConnectorFullsyncSchedule     `json:"fullsyncSchedule,omitempty"`           //
	Protocol                   string                                                                              `json:"protocol,omitempty"`                   // protocol
	SkipCertificateValidations *bool                                                                               `json:"skipCertificateValidations,omitempty"` //
	URL                        *ResponseEddaGetConnectorConfigByConnectorNameResponseConnectorURL                  `json:"url,omitempty"`                        //
}

type ResponseEddaGetConnectorConfigByConnectorNameResponseConnectorAdditionalProperties interface{}

type ResponseEddaGetConnectorConfigByConnectorNameResponseConnectorAttributes struct {
	AttributeMapping     *[]ResponseEddaGetConnectorConfigByConnectorNameResponseConnectorAttributesAttributeMapping `json:"attributeMapping,omitempty"`     // <p>List of feature names</p>
	BulkUniqueIDentifier string                                                                                      `json:"bulkUniqueIdentifier,omitempty"` // uniqueness to identify
	TopLevelObject       string                                                                                      `json:"topLevelObject,omitempty"`       // root level of json
	UniqueIDentifier     string                                                                                      `json:"uniqueIdentifier,omitempty"`     // uniqueness to identify
	VersionIDentifier    string                                                                                      `json:"versionIdentifier,omitempty"`    // version uniqueness to identify
}

type ResponseEddaGetConnectorConfigByConnectorNameResponseConnectorAttributesAttributeMapping struct {
	DictionaryAttribute string `json:"dictionaryAttribute,omitempty"` //
	IncludeInDictionary *bool  `json:"includeInDictionary,omitempty"` //
	JSONAttribute       string `json:"jsonAttribute,omitempty"`       //
}

type ResponseEddaGetConnectorConfigByConnectorNameResponseConnectorDeltasyncSchedule struct {
	Interval     *int   `json:"interval,omitempty"`     // run at interval (hours)
	IntervalUnit string `json:"intervalUnit,omitempty"` // interval Units
	StartDate    string `json:"startDate,omitempty"`    // start date and Time
}

type ResponseEddaGetConnectorConfigByConnectorNameResponseConnectorFullsyncSchedule struct {
	Interval     *int   `json:"interval,omitempty"`     // run at interval (hours)
	IntervalUnit string `json:"intervalUnit,omitempty"` // interval Units
	StartDate    string `json:"startDate,omitempty"`    // start date and Time
}

type ResponseEddaGetConnectorConfigByConnectorNameResponseConnectorURL struct {
	AccessKey          string `json:"accessKey,omitempty"`          // accesskey
	AuthenticationType string `json:"authenticationType,omitempty"` // authentication Type list
	BulkURL            string `json:"bulkUrl,omitempty"`            // bulkUrl
	ClientID           string `json:"clientId,omitempty"`           // clientid
	ClientSecret       string `json:"clientSecret,omitempty"`       // clientsecret
	IncrementalURL     string `json:"incrementalUrl,omitempty"`     // incrementalUrl
	Password           string `json:"password,omitempty"`           // password
	RefreshToken       string `json:"refreshToken,omitempty"`       // refreshtoken
	TokenHeader        string `json:"tokenHeader,omitempty"`        // tokenHeader
	UserName           string `json:"userName,omitempty"`           // userName
}

// # Review unknown case
// # Review unknown case

type ResponseEddaGetEddaDictionaryReferences struct {
	Response *ResponseEddaGetEddaDictionaryReferencesResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  //
}

type ResponseEddaGetEddaDictionaryReferencesResponse interface{}

type ResponseEddaTestConnector struct {
	ConnectorName string `json:"connectorName,omitempty"` // connectorName
	Data          string `json:"data,omitempty"`          // Response data
	Error         string `json:"error,omitempty"`         // error
	Status        string `json:"status,omitempty"`        // status of the request
	UniqueID      string `json:"uniqueID,omitempty"`      // uniqueness to identify
}

type RequestEddaCreateConnectorConfig struct {
	Connector *RequestEddaCreateConnectorConfigConnector `json:"connector,omitempty"` // connectorName
}

type RequestEddaCreateConnectorConfigConnector struct {
	AdditionalProperties       *RequestEddaCreateConnectorConfigConnectorAdditionalProperties `json:"additionalProperties,omitempty"`       //
	Attributes                 *RequestEddaCreateConnectorConfigConnectorAttributes           `json:"attributes,omitempty"`                 // connectorName
	ConnectorName              string                                                         `json:"connectorName,omitempty"`              // connectorName
	ConnectorType              string                                                         `json:"connectorType,omitempty"`              // connector Type list
	DeltasyncSchedule          *RequestEddaCreateConnectorConfigConnectorDeltasyncSchedule    `json:"deltasyncSchedule,omitempty"`          //
	Description                string                                                         `json:"description,omitempty"`                // description
	Enabled                    *bool                                                          `json:"enabled,omitempty"`                    //
	FullsyncSchedule           *RequestEddaCreateConnectorConfigConnectorFullsyncSchedule     `json:"fullsyncSchedule,omitempty"`           //
	Protocol                   string                                                         `json:"protocol,omitempty"`                   // protocol
	SkipCertificateValidations *bool                                                          `json:"skipCertificateValidations,omitempty"` //
	URL                        *RequestEddaCreateConnectorConfigConnectorURL                  `json:"url,omitempty"`                        //
}

type RequestEddaCreateConnectorConfigConnectorAdditionalProperties interface{}

type RequestEddaCreateConnectorConfigConnectorAttributes struct {
	AttributeMapping     *[]RequestEddaCreateConnectorConfigConnectorAttributesAttributeMapping `json:"attributeMapping,omitempty"`     // <p>List of feature names</p>
	BulkUniqueIDentifier string                                                                 `json:"bulkUniqueIdentifier,omitempty"` // uniqueness to identify
	TopLevelObject       string                                                                 `json:"topLevelObject,omitempty"`       // root level of json
	UniqueIDentifier     string                                                                 `json:"uniqueIdentifier,omitempty"`     // uniqueness to identify
	VersionIDentifier    string                                                                 `json:"versionIdentifier,omitempty"`    // version uniqueness to identify
}

type RequestEddaCreateConnectorConfigConnectorAttributesAttributeMapping struct {
	DictionaryAttribute string `json:"dictionaryAttribute,omitempty"` //
	IncludeInDictionary *bool  `json:"includeInDictionary,omitempty"` //
	JSONAttribute       string `json:"jsonAttribute,omitempty"`       //
}

type RequestEddaCreateConnectorConfigConnectorDeltasyncSchedule struct {
	Interval     *int   `json:"interval,omitempty"`     // run at interval (hours)
	IntervalUnit string `json:"intervalUnit,omitempty"` // interval Units
	StartDate    string `json:"startDate,omitempty"`    // start date and Time
}

type RequestEddaCreateConnectorConfigConnectorFullsyncSchedule struct {
	Interval     *int   `json:"interval,omitempty"`     // run at interval (hours)
	IntervalUnit string `json:"intervalUnit,omitempty"` // interval Units
	StartDate    string `json:"startDate,omitempty"`    // start date and Time
}

type RequestEddaCreateConnectorConfigConnectorURL struct {
	AccessKey          string `json:"accessKey,omitempty"`          // accesskey
	AuthenticationType string `json:"authenticationType,omitempty"` // authentication Type list
	BulkURL            string `json:"bulkUrl,omitempty"`            // bulkUrl
	ClientID           string `json:"clientId,omitempty"`           // clientid
	ClientSecret       string `json:"clientSecret,omitempty"`       // clientsecret
	IncrementalURL     string `json:"incrementalUrl,omitempty"`     // incrementalUrl
	Password           string `json:"password,omitempty"`           // password
	RefreshToken       string `json:"refreshToken,omitempty"`       // refreshtoken
	TokenHeader        string `json:"tokenHeader,omitempty"`        // tokenHeader
	UserName           string `json:"userName,omitempty"`           // userName
}

type RequestEddaUpdateConnectorConfigByConnectorName struct {
	Connector *RequestEddaUpdateConnectorConfigByConnectorNameConnector `json:"connector,omitempty"` // connectorName
}

type RequestEddaUpdateConnectorConfigByConnectorNameConnector struct {
	AdditionalProperties       *RequestEddaUpdateConnectorConfigByConnectorNameConnectorAdditionalProperties `json:"additionalProperties,omitempty"`       //
	Attributes                 *RequestEddaUpdateConnectorConfigByConnectorNameConnectorAttributes           `json:"attributes,omitempty"`                 // connectorName
	ConnectorName              string                                                                        `json:"connectorName,omitempty"`              // connectorName
	ConnectorType              string                                                                        `json:"connectorType,omitempty"`              // connector Type list
	DeltasyncSchedule          *RequestEddaUpdateConnectorConfigByConnectorNameConnectorDeltasyncSchedule    `json:"deltasyncSchedule,omitempty"`          //
	Description                string                                                                        `json:"description,omitempty"`                // description
	Enabled                    *bool                                                                         `json:"enabled,omitempty"`                    //
	FullsyncSchedule           *RequestEddaUpdateConnectorConfigByConnectorNameConnectorFullsyncSchedule     `json:"fullsyncSchedule,omitempty"`           //
	Protocol                   string                                                                        `json:"protocol,omitempty"`                   // protocol
	SkipCertificateValidations *bool                                                                         `json:"skipCertificateValidations,omitempty"` //
	URL                        *RequestEddaUpdateConnectorConfigByConnectorNameConnectorURL                  `json:"url,omitempty"`                        //
}

type RequestEddaUpdateConnectorConfigByConnectorNameConnectorAdditionalProperties interface{}

type RequestEddaUpdateConnectorConfigByConnectorNameConnectorAttributes struct {
	AttributeMapping     *[]RequestEddaUpdateConnectorConfigByConnectorNameConnectorAttributesAttributeMapping `json:"attributeMapping,omitempty"`     // <p>List of feature names</p>
	BulkUniqueIDentifier string                                                                                `json:"bulkUniqueIdentifier,omitempty"` // uniqueness to identify
	TopLevelObject       string                                                                                `json:"topLevelObject,omitempty"`       // root level of json
	UniqueIDentifier     string                                                                                `json:"uniqueIdentifier,omitempty"`     // uniqueness to identify
	VersionIDentifier    string                                                                                `json:"versionIdentifier,omitempty"`    // version uniqueness to identify
}

type RequestEddaUpdateConnectorConfigByConnectorNameConnectorAttributesAttributeMapping struct {
	DictionaryAttribute string `json:"dictionaryAttribute,omitempty"` //
	IncludeInDictionary *bool  `json:"includeInDictionary,omitempty"` //
	JSONAttribute       string `json:"jsonAttribute,omitempty"`       //
}

type RequestEddaUpdateConnectorConfigByConnectorNameConnectorDeltasyncSchedule struct {
	Interval     *int   `json:"interval,omitempty"`     // run at interval (hours)
	IntervalUnit string `json:"intervalUnit,omitempty"` // interval Units
	StartDate    string `json:"startDate,omitempty"`    // start date and Time
}

type RequestEddaUpdateConnectorConfigByConnectorNameConnectorFullsyncSchedule struct {
	Interval     *int   `json:"interval,omitempty"`     // run at interval (hours)
	IntervalUnit string `json:"intervalUnit,omitempty"` // interval Units
	StartDate    string `json:"startDate,omitempty"`    // start date and Time
}

type RequestEddaUpdateConnectorConfigByConnectorNameConnectorURL struct {
	AccessKey          string `json:"accessKey,omitempty"`          // accesskey
	AuthenticationType string `json:"authenticationType,omitempty"` // authentication Type list
	BulkURL            string `json:"bulkUrl,omitempty"`            // bulkUrl
	ClientID           string `json:"clientId,omitempty"`           // clientid
	ClientSecret       string `json:"clientSecret,omitempty"`       // clientsecret
	IncrementalURL     string `json:"incrementalUrl,omitempty"`     // incrementalUrl
	Password           string `json:"password,omitempty"`           // password
	RefreshToken       string `json:"refreshToken,omitempty"`       // refreshtoken
	TokenHeader        string `json:"tokenHeader,omitempty"`        // tokenHeader
	UserName           string `json:"userName,omitempty"`           // userName
}

type RequestEddaTestConnector struct {
	AuthType        string                              `json:"authType,omitempty"`        // authentication Type list
	AuthValues      *RequestEddaTestConnectorAuthValues `json:"authValues,omitempty"`      // Request  to test Connector
	ConnectorName   string                              `json:"connectorName,omitempty"`   // connectorName
	ResponseParsing string                              `json:"responseParsing,omitempty"` // uniqueness to identify
	UniqueID        string                              `json:"uniqueID,omitempty"`        // uniqueness to identify
	URL             string                              `json:"url,omitempty"`             // bulkUrl
}

type RequestEddaTestConnectorAuthValues struct {
	Password string `json:"password,omitempty"` // password
	UserName string `json:"userName,omitempty"` // userName
}

//GetConnectorConfig get All connectorConfig information
/* EDDA Get ALL connectorConfig information

 */
func (s *EddaService) GetConnectorConfig() (*ResponseEddaGetConnectorConfig, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/edda/connector-config"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEddaGetConnectorConfig{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetConnectorConfig")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEddaGetConnectorConfig)
	return result, response, err

}

//GetConnectorConfigByConnectorName get connectorConfig information
/* EDDA Get connectorConfig information based on ConnectorName

@param connectorName connectorName path parameter. update or delete or retrieve the connector config.
*/
func (s *EddaService) GetConnectorConfigByConnectorName(connectorName string) (*ResponseEddaGetConnectorConfigByConnectorName, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/edda/connector-config/{connectorName}"
	path = strings.Replace(path, "{connectorName}", fmt.Sprintf("%v", connectorName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEddaGetConnectorConfigByConnectorName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetConnectorConfigByConnectorName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEddaGetConnectorConfigByConnectorName)
	return result, response, err

}

//GetEddaDictionaryReferences Get a map of references to EDDA dictionaries
/* EDDA Get a map of references to EDDA dictionaries

 */
func (s *EddaService) GetEddaDictionaryReferences() (*ResponseEddaGetEddaDictionaryReferences, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/edda/dictionary-references"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEddaGetEddaDictionaryReferences{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetEddaDictionaryReferences")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEddaGetEddaDictionaryReferences)
	return result, response, err

}

//CreateConnectorConfig Configure connectorconfig information
/* EDDA Configure connectorconfig information.

 */
func (s *EddaService) CreateConnectorConfig(requestEddaCreateConnectorConfig *RequestEddaCreateConnectorConfig) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/edda/connector-config"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEddaCreateConnectorConfig).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateConnectorConfig")
	}

	return response, err

}

//TestConnector test the Connector
/* EDDA test the Connector.

 */
func (s *EddaService) TestConnector(requestEddaTestConnector *RequestEddaTestConnector) (*ResponseEddaTestConnector, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/edda/test-connector"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEddaTestConnector).
		SetResult(&ResponseEddaTestConnector{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseEddaTestConnector{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation TestConnector")
	}

	result := response.Result().(*ResponseEddaTestConnector)
	return result, response, err

}

//UpdateConnectorConfigByConnectorName update connectoConfig information
/* EDDA update Configure connectorConfig information based on ConnectorName.

@param connectorName connectorName path parameter. update or delete or retrieve the connector config.
*/
func (s *EddaService) UpdateConnectorConfigByConnectorName(connectorName string, requestEddaUpdateConnectorConfigByConnectorName *RequestEddaUpdateConnectorConfigByConnectorName) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/edda/connector-config/{connectorName}"
	path = strings.Replace(path, "{connectorName}", fmt.Sprintf("%v", connectorName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEddaUpdateConnectorConfigByConnectorName).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateConnectorConfigByConnectorName")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteConnectorConfigByConnectorName Delete connectorConfig information
/* EDDA Delete Configure connectorConfig information based on ConnectorName.

@param connectorName connectorName path parameter. update or delete or retrieve the connector config.
*/
func (s *EddaService) DeleteConnectorConfigByConnectorName(connectorName string) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/edda/connector-config/{connectorName}"
	path = strings.Replace(path, "{connectorName}", fmt.Sprintf("%v", connectorName), -1)

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
		return response, fmt.Errorf("error with operation DeleteConnectorConfigByConnectorName")
	}

	getCSFRToken(response.Header())

	return response, err

}
