package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type PxGridDirectService service

type ResponsePxGridDirectGetConnectorConfig struct {
	Response *[]ResponsePxGridDirectGetConnectorConfigResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}

type ResponsePxGridDirectGetConnectorConfigResponse struct {
	Connector *ResponsePxGridDirectGetConnectorConfigResponseConnector `json:"connector,omitempty"` // connectorName
}

type ResponsePxGridDirectGetConnectorConfigResponseConnector struct {
	AdditionalProperties       *ResponsePxGridDirectGetConnectorConfigResponseConnectorAdditionalProperties `json:"additionalProperties,omitempty"`       //
	Attributes                 *ResponsePxGridDirectGetConnectorConfigResponseConnectorAttributes           `json:"attributes,omitempty"`                 // connectorName
	ConnectorName              string                                                                       `json:"connectorName,omitempty"`              // connectorName
	ConnectorType              string                                                                       `json:"connectorType,omitempty"`              // connector Type list
	DeltasyncSchedule          *ResponsePxGridDirectGetConnectorConfigResponseConnectorDeltasyncSchedule    `json:"deltasyncSchedule,omitempty"`          //
	Description                string                                                                       `json:"description,omitempty"`                // description
	Enabled                    *bool                                                                        `json:"enabled,omitempty"`                    //
	FullsyncSchedule           *ResponsePxGridDirectGetConnectorConfigResponseConnectorFullsyncSchedule     `json:"fullsyncSchedule,omitempty"`           //
	Protocol                   string                                                                       `json:"protocol,omitempty"`                   // protocol
	SkipCertificateValidations *bool                                                                        `json:"skipCertificateValidations,omitempty"` //
	URL                        *ResponsePxGridDirectGetConnectorConfigResponseConnectorURL                  `json:"url,omitempty"`                        //
}

type ResponsePxGridDirectGetConnectorConfigResponseConnectorAdditionalProperties interface{}

type ResponsePxGridDirectGetConnectorConfigResponseConnectorAttributes struct {
	AttributeMapping      *[]ResponsePxGridDirectGetConnectorConfigResponseConnectorAttributesAttributeMapping `json:"attributeMapping,omitempty"`      // <p>List of feature names</p>
	CorrelationIDentifier string                                                                               `json:"correlationIdentifier,omitempty"` // uniqueness to identify
	TopLevelObject        string                                                                               `json:"topLevelObject,omitempty"`        // root level of json
	UniqueIDentifier      string                                                                               `json:"uniqueIdentifier,omitempty"`      // uniqueness to identify
	VersionIDentifier     string                                                                               `json:"versionIdentifier,omitempty"`     // version uniqueness to identify
}

type ResponsePxGridDirectGetConnectorConfigResponseConnectorAttributesAttributeMapping struct {
	DictionaryAttribute string `json:"dictionaryAttribute,omitempty"` //
	IncludeInDictionary *bool  `json:"includeInDictionary,omitempty"` //
	JSONAttribute       string `json:"jsonAttribute,omitempty"`       //
}

type ResponsePxGridDirectGetConnectorConfigResponseConnectorDeltasyncSchedule struct {
	Interval     *int   `json:"interval,omitempty"`     // run at interval (hours)
	IntervalUnit string `json:"intervalUnit,omitempty"` // interval Units
	StartDate    string `json:"startDate,omitempty"`    // start date and Time
}

type ResponsePxGridDirectGetConnectorConfigResponseConnectorFullsyncSchedule struct {
	Interval     *int   `json:"interval,omitempty"`     // run at interval (hours)
	IntervalUnit string `json:"intervalUnit,omitempty"` // interval Units
	StartDate    string `json:"startDate,omitempty"`    // start date and Time
}

type ResponsePxGridDirectGetConnectorConfigResponseConnectorURL struct {
	AuthenticationType string `json:"authenticationType,omitempty"` // authentication Type list
	BulkURL            string `json:"bulkUrl,omitempty"`            // bulkUrl
	IncrementalURL     string `json:"incrementalUrl,omitempty"`     // incrementalUrl
	Password           string `json:"password,omitempty"`           // password
	UserName           string `json:"userName,omitempty"`           // userName
}

// # Review unknown case

type ResponsePxGridDirectGetConnectorConfigByConnectorName struct {
	Response *ResponsePxGridDirectGetConnectorConfigByConnectorNameResponse `json:"response,omitempty"` // ConnectorConfig information format
	Version  string                                                         `json:"version,omitempty"`  //
}

type ResponsePxGridDirectGetConnectorConfigByConnectorNameResponse struct {
	Connector *ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnector `json:"connector,omitempty"` // connectorName
}

type ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnector struct {
	AdditionalProperties       *ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorAdditionalProperties `json:"additionalProperties,omitempty"`       //
	Attributes                 *ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorAttributes           `json:"attributes,omitempty"`                 // connectorName
	ConnectorName              string                                                                                      `json:"connectorName,omitempty"`              // connectorName
	ConnectorType              string                                                                                      `json:"connectorType,omitempty"`              // connector Type list
	DeltasyncSchedule          *ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorDeltasyncSchedule    `json:"deltasyncSchedule,omitempty"`          //
	Description                string                                                                                      `json:"description,omitempty"`                // description
	Enabled                    *bool                                                                                       `json:"enabled,omitempty"`                    //
	FullsyncSchedule           *ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorFullsyncSchedule     `json:"fullsyncSchedule,omitempty"`           //
	Protocol                   string                                                                                      `json:"protocol,omitempty"`                   // protocol
	SkipCertificateValidations *bool                                                                                       `json:"skipCertificateValidations,omitempty"` //
	URL                        *ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorURL                  `json:"url,omitempty"`                        //
}

type ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorAdditionalProperties interface{}

type ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorAttributes struct {
	AttributeMapping      *[]ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorAttributesAttributeMapping `json:"attributeMapping,omitempty"`      // <p>List of feature names</p>
	CorrelationIDentifier string                                                                                              `json:"correlationIdentifier,omitempty"` // uniqueness to identify
	TopLevelObject        string                                                                                              `json:"topLevelObject,omitempty"`        // root level of json
	UniqueIDentifier      string                                                                                              `json:"uniqueIdentifier,omitempty"`      // uniqueness to identify
	VersionIDentifier     string                                                                                              `json:"versionIdentifier,omitempty"`     // version uniqueness to identify
}

type ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorAttributesAttributeMapping struct {
	DictionaryAttribute string `json:"dictionaryAttribute,omitempty"` //
	IncludeInDictionary *bool  `json:"includeInDictionary,omitempty"` //
	JSONAttribute       string `json:"jsonAttribute,omitempty"`       //
}

type ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorDeltasyncSchedule struct {
	Interval     *int   `json:"interval,omitempty"`     // run at interval (hours)
	IntervalUnit string `json:"intervalUnit,omitempty"` // interval Units
	StartDate    string `json:"startDate,omitempty"`    // start date and Time
}

type ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorFullsyncSchedule struct {
	Interval     *int   `json:"interval,omitempty"`     // run at interval (hours)
	IntervalUnit string `json:"intervalUnit,omitempty"` // interval Units
	StartDate    string `json:"startDate,omitempty"`    // start date and Time
}

type ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorURL struct {
	AuthenticationType string `json:"authenticationType,omitempty"` // authentication Type list
	BulkURL            string `json:"bulkUrl,omitempty"`            // bulkUrl
	IncrementalURL     string `json:"incrementalUrl,omitempty"`     // incrementalUrl
	Password           string `json:"password,omitempty"`           // password
	UserName           string `json:"userName,omitempty"`           // userName
}

// # Review unknown case
// # Review unknown case

type ResponsePxGridDirectGetpxgridDirectDictionaryReferences struct {
	Response *ResponsePxGridDirectGetpxgridDirectDictionaryReferencesResponse `json:"response,omitempty"` //
	Version  string                                                           `json:"version,omitempty"`  //
}

type ResponsePxGridDirectGetpxgridDirectDictionaryReferencesResponse interface{}

type ResponsePxGridDirectGetConnectorConfigSyncNowStatus struct {
	Response *ResponsePxGridDirectGetConnectorConfigSyncNowStatusResponse `json:"response,omitempty"` // ConnectorConfig information format
	Version  string                                                       `json:"version,omitempty"`  //
}

type ResponsePxGridDirectGetConnectorConfigSyncNowStatusResponse struct {
	Connector *ResponsePxGridDirectGetConnectorConfigSyncNowStatusResponseConnector `json:"connector,omitempty"` // connectorName
}

type ResponsePxGridDirectGetConnectorConfigSyncNowStatusResponseConnector struct {
	ConnectorName string `json:"connectorName,omitempty"` // connectorName
	SyncStatus    string `json:"syncStatus,omitempty"`    // description
}

// # Review unknown case

type ResponsePxGridDirectTestConnector struct {
	ConnectorName              string `json:"connectorName,omitempty"`              // connectorName
	Data                       string `json:"data,omitempty"`                       // Response data
	Error                      string `json:"error,omitempty"`                      // error
	SkipCertificateValidations *bool  `json:"skipCertificateValidations,omitempty"` // skipCertificateValidations
	Status                     string `json:"status,omitempty"`                     // status of the request
	UniqueID                   string `json:"uniqueID,omitempty"`                   // uniqueness to identify
}

type RequestPxGridDirectCreateConnectorConfig struct {
	Connector *RequestPxGridDirectCreateConnectorConfigConnector `json:"connector,omitempty"` // connectorName
}

type RequestPxGridDirectCreateConnectorConfigConnector struct {
	AdditionalProperties       *RequestPxGridDirectCreateConnectorConfigConnectorAdditionalProperties `json:"additionalProperties,omitempty"`       //
	Attributes                 *RequestPxGridDirectCreateConnectorConfigConnectorAttributes           `json:"attributes,omitempty"`                 // connectorName
	ConnectorName              string                                                                 `json:"connectorName,omitempty"`              // connectorName
	ConnectorType              string                                                                 `json:"connectorType,omitempty"`              // connector Type list
	DeltasyncSchedule          *RequestPxGridDirectCreateConnectorConfigConnectorDeltasyncSchedule    `json:"deltasyncSchedule,omitempty"`          //
	Description                string                                                                 `json:"description,omitempty"`                // description
	Enabled                    *bool                                                                  `json:"enabled,omitempty"`                    //
	FullsyncSchedule           *RequestPxGridDirectCreateConnectorConfigConnectorFullsyncSchedule     `json:"fullsyncSchedule,omitempty"`           //
	Protocol                   string                                                                 `json:"protocol,omitempty"`                   // protocol
	SkipCertificateValidations *bool                                                                  `json:"skipCertificateValidations,omitempty"` //
	URL                        *RequestPxGridDirectCreateConnectorConfigConnectorURL                  `json:"url,omitempty"`                        //
}

type RequestPxGridDirectCreateConnectorConfigConnectorAdditionalProperties interface{}

type RequestPxGridDirectCreateConnectorConfigConnectorAttributes struct {
	AttributeMapping      *[]RequestPxGridDirectCreateConnectorConfigConnectorAttributesAttributeMapping `json:"attributeMapping,omitempty"`      // <p>List of feature names</p>
	CorrelationIDentifier string                                                                         `json:"correlationIdentifier,omitempty"` // uniqueness to identify
	TopLevelObject        string                                                                         `json:"topLevelObject,omitempty"`        // root level of json
	UniqueIDentifier      string                                                                         `json:"uniqueIdentifier,omitempty"`      // uniqueness to identify
	VersionIDentifier     string                                                                         `json:"versionIdentifier,omitempty"`     // version uniqueness to identify
}

type RequestPxGridDirectCreateConnectorConfigConnectorAttributesAttributeMapping struct {
	DictionaryAttribute string `json:"dictionaryAttribute,omitempty"` //
	IncludeInDictionary *bool  `json:"includeInDictionary,omitempty"` //
	JSONAttribute       string `json:"jsonAttribute,omitempty"`       //
}

type RequestPxGridDirectCreateConnectorConfigConnectorDeltasyncSchedule struct {
	Interval     *int   `json:"interval,omitempty"`     // run at interval (hours)
	IntervalUnit string `json:"intervalUnit,omitempty"` // interval Units
	StartDate    string `json:"startDate,omitempty"`    // start date and Time
}

type RequestPxGridDirectCreateConnectorConfigConnectorFullsyncSchedule struct {
	Interval     *int   `json:"interval,omitempty"`     // run at interval (hours)
	IntervalUnit string `json:"intervalUnit,omitempty"` // interval Units
	StartDate    string `json:"startDate,omitempty"`    // start date and Time
}

type RequestPxGridDirectCreateConnectorConfigConnectorURL struct {
	AuthenticationType string `json:"authenticationType,omitempty"` // authentication Type list
	BulkURL            string `json:"bulkUrl,omitempty"`            // bulkUrl
	IncrementalURL     string `json:"incrementalUrl,omitempty"`     // incrementalUrl
	Password           string `json:"password,omitempty"`           // password
	UserName           string `json:"userName,omitempty"`           // userName
}

type RequestPxGridDirectUpdateConnectorConfigByConnectorName struct {
	Connector *RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnector `json:"connector,omitempty"` // connectorName
}

type RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnector struct {
	AdditionalProperties       *RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAdditionalProperties `json:"additionalProperties,omitempty"`       //
	Attributes                 *RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAttributes           `json:"attributes,omitempty"`                 // connectorName
	ConnectorName              string                                                                                `json:"connectorName,omitempty"`              // connectorName
	ConnectorType              string                                                                                `json:"connectorType,omitempty"`              // connector Type list
	DeltasyncSchedule          *RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorDeltasyncSchedule    `json:"deltasyncSchedule,omitempty"`          //
	Description                string                                                                                `json:"description,omitempty"`                // description
	Enabled                    *bool                                                                                 `json:"enabled,omitempty"`                    //
	FullsyncSchedule           *RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorFullsyncSchedule     `json:"fullsyncSchedule,omitempty"`           //
	Protocol                   string                                                                                `json:"protocol,omitempty"`                   // protocol
	SkipCertificateValidations *bool                                                                                 `json:"skipCertificateValidations,omitempty"` //
	URL                        *RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorURL                  `json:"url,omitempty"`                        //
}

type RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAdditionalProperties interface{}

type RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAttributes struct {
	AttributeMapping      *[]RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAttributesAttributeMapping `json:"attributeMapping,omitempty"`      // <p>List of feature names</p>
	CorrelationIDentifier string                                                                                        `json:"correlationIdentifier,omitempty"` // uniqueness to identify
	TopLevelObject        string                                                                                        `json:"topLevelObject,omitempty"`        // root level of json
	UniqueIDentifier      string                                                                                        `json:"uniqueIdentifier,omitempty"`      // uniqueness to identify
	VersionIDentifier     string                                                                                        `json:"versionIdentifier,omitempty"`     // version uniqueness to identify
}

type RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAttributesAttributeMapping struct {
	DictionaryAttribute string `json:"dictionaryAttribute,omitempty"` //
	IncludeInDictionary *bool  `json:"includeInDictionary,omitempty"` //
	JSONAttribute       string `json:"jsonAttribute,omitempty"`       //
}

type RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorDeltasyncSchedule struct {
	Interval     *int   `json:"interval,omitempty"`     // run at interval (hours)
	IntervalUnit string `json:"intervalUnit,omitempty"` // interval Units
	StartDate    string `json:"startDate,omitempty"`    // start date and Time
}

type RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorFullsyncSchedule struct {
	Interval     *int   `json:"interval,omitempty"`     // run at interval (hours)
	IntervalUnit string `json:"intervalUnit,omitempty"` // interval Units
	StartDate    string `json:"startDate,omitempty"`    // start date and Time
}

type RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorURL struct {
	AuthenticationType string `json:"authenticationType,omitempty"` // authentication Type list
	BulkURL            string `json:"bulkUrl,omitempty"`            // bulkUrl
	IncrementalURL     string `json:"incrementalUrl,omitempty"`     // incrementalUrl
	Password           string `json:"password,omitempty"`           // password
	UserName           string `json:"userName,omitempty"`           // userName
}

type RequestPxGridDirectSyncNowConnector struct {
	Connector *RequestPxGridDirectSyncNowConnectorConnector `json:"connector,omitempty"` // connectorName
}

type RequestPxGridDirectSyncNowConnectorConnector struct {
	SyncType      string `json:"SyncType,omitempty"`      // connector Type list
	ConnectorName string `json:"connectorName,omitempty"` // connectorName
	Description   string `json:"description,omitempty"`   // description
}

type RequestPxGridDirectTestConnector struct {
	AuthType                   string                                      `json:"authType,omitempty"`                   // authentication Type list
	AuthValues                 *RequestPxGridDirectTestConnectorAuthValues `json:"authValues,omitempty"`                 // Request  to test Connector
	ConnectorName              string                                      `json:"connectorName,omitempty"`              // connectorName
	ResponseParsing            string                                      `json:"responseParsing,omitempty"`            // uniqueness to identify
	SkipCertificateValidations *bool                                       `json:"skipCertificateValidations,omitempty"` // skipCertificateValidations
	UniqueID                   string                                      `json:"uniqueID,omitempty"`                   // uniqueness to identify
	URL                        string                                      `json:"url,omitempty"`                        // bulkUrl
}

type RequestPxGridDirectTestConnectorAuthValues struct {
	Password string `json:"password,omitempty"` // password
	UserName string `json:"userName,omitempty"` // userName
}

//GetConnectorConfig get All connectorConfig information
/* pxGrid Direct Get ALL connectorConfig information

 */
func (s *PxGridDirectService) GetConnectorConfig() (*ResponsePxGridDirectGetConnectorConfig, *resty.Response, error) {
	setHost(s.client, "_px_grid_direct")
	path := "/api/v1/pxgrid-direct/connector-config"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePxGridDirectGetConnectorConfig{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetConnectorConfig")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePxGridDirectGetConnectorConfig)
	return result, response, err

}

//GetConnectorConfigByConnectorName get connectorConfig information
/* pxGrid Direct Get connectorConfig information based on ConnectorName

@param connectorName connectorName path parameter. update or delete or retrieve the connector config.
*/
func (s *PxGridDirectService) GetConnectorConfigByConnectorName(connectorName string) (*ResponsePxGridDirectGetConnectorConfigByConnectorName, *resty.Response, error) {
	setHost(s.client, "_px_grid_direct")
	path := "/api/v1/pxgrid-direct/connector-config/{connectorName}"
	path = strings.Replace(path, "{connectorName}", fmt.Sprintf("%v", connectorName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePxGridDirectGetConnectorConfigByConnectorName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetConnectorConfigByConnectorName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePxGridDirectGetConnectorConfigByConnectorName)
	return result, response, err

}

//GetpxgridDirectDictionaryReferences Get a map of references to pxgrid-direct dictionaries
/* pxGrid Direct Get a map of references to pxgrid-direct dictionaries

 */
func (s *PxGridDirectService) GetpxgridDirectDictionaryReferences() (*ResponsePxGridDirectGetpxgridDirectDictionaryReferences, *resty.Response, error) {
	setHost(s.client, "_px_grid_direct")
	path := "/api/v1/pxgrid-direct/dictionary-references"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePxGridDirectGetpxgridDirectDictionaryReferences{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetpxgridDirectDictionaryReferences")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePxGridDirectGetpxgridDirectDictionaryReferences)
	return result, response, err

}

//GetConnectorConfigSyncNowStatus sync Now Status for  Connector information
/* This API is used to get the status for SyncNow Status

It returns the sync status as


syncstatus


QUEUED ,means its in  QUEUED state

SUBMITTED ,means its in  Submited to fetch the data

INPROGRESS ,means its inprogress of fetching and saving in ISE

ERRORED ,means some internal error while fetching and saving in ISE further debugging logs will help

COMPLETED ,means its COMPLETED of fetching and saving in ISE

SCH_INPROGRESS ,means its inprogress for schedule time fetch and saving in ISE

SCH_SUBMITTED ,means its submitted for schedule time fetch and will start to saving in ISE

CANCELED ,means its cancelled if any of ISE Service start when its in middle of QUEUED/SUBMITTED/INPROGRESS


connectorName



@param connectorName connectorName path parameter. retrieve the connector syncnow status.
*/
func (s *PxGridDirectService) GetConnectorConfigSyncNowStatus(connectorName string) (*ResponsePxGridDirectGetConnectorConfigSyncNowStatus, *resty.Response, error) {
	setHost(s.client, "_px_grid_direct")
	path := "/api/v1/pxgrid-direct/syncNowStatus/{connectorName}"
	path = strings.Replace(path, "{connectorName}", fmt.Sprintf("%v", connectorName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePxGridDirectGetConnectorConfigSyncNowStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetConnectorConfigSyncNowStatus")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePxGridDirectGetConnectorConfigSyncNowStatus)
	return result, response, err

}

//CreateConnectorConfig Configure connectorconfig information
/* pxGrid Direct Configure connectorconfig information.

 */
func (s *PxGridDirectService) CreateConnectorConfig(requestPxGridDirectCreateConnectorConfig *RequestPxGridDirectCreateConnectorConfig) (*resty.Response, error) {
	setHost(s.client, "_px_grid_direct")
	path := "/api/v1/pxgrid-direct/connector-config"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestPxGridDirectCreateConnectorConfig).
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

//SyncNowConnector syncnow for the Connector
/*
This syncNow is used on demand on a URLFetch Type connector only
 Following parameters are present in the POST request body




PARAMETER

DESCRIPTION

EXAMPLE





SyncType
*required

its for FULL or INCREMENTAL

"SyncType": "FULL or INCREMENTAL"



connectorName
*required

Name of the Connector for only URLFetcher type

name of Connector



description

Decription of the Connector

"length": "256 character"




NOTE:
For
Use syncNowStatus api to get the status of the connector



*/
func (s *PxGridDirectService) SyncNowConnector(requestPxGridDirectSyncNowConnector *RequestPxGridDirectSyncNowConnector) (*resty.Response, error) {
	setHost(s.client, "_px_grid_direct")
	path := "/api/v1/pxgrid-direct/syncnow"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestPxGridDirectSyncNowConnector).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation SyncNowConnector")
	}

	return response, err

}

//TestConnector test the Connector
/* pxGrid Direct test the Connector.

 */
func (s *PxGridDirectService) TestConnector(requestPxGridDirectTestConnector *RequestPxGridDirectTestConnector) (*ResponsePxGridDirectTestConnector, *resty.Response, error) {
	setHost(s.client, "_px_grid_direct")
	path := "/api/v1/pxgrid-direct/test-connector"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestPxGridDirectTestConnector).
		SetResult(&ResponsePxGridDirectTestConnector{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponsePxGridDirectTestConnector{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation TestConnector")
	}

	result := response.Result().(*ResponsePxGridDirectTestConnector)
	return result, response, err

}

//UpdateConnectorConfigByConnectorName update connectoConfig information
/* pxGrid Direct update Configure connectorConfig information based on ConnectorName.

@param connectorName connectorName path parameter. update or delete or retrieve the connector config.
*/
func (s *PxGridDirectService) UpdateConnectorConfigByConnectorName(connectorName string, requestPxGridDirectUpdateConnectorConfigByConnectorName *RequestPxGridDirectUpdateConnectorConfigByConnectorName) (*resty.Response, error) {
	setHost(s.client, "_px_grid_direct")
	path := "/api/v1/pxgrid-direct/connector-config/{connectorName}"
	path = strings.Replace(path, "{connectorName}", fmt.Sprintf("%v", connectorName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestPxGridDirectUpdateConnectorConfigByConnectorName).
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
/* pxGrid Direct Delete Configure connectorConfig information based on ConnectorName.

@param connectorName connectorName path parameter. update or delete or retrieve the connector config.
*/
func (s *PxGridDirectService) DeleteConnectorConfigByConnectorName(connectorName string) (*resty.Response, error) {
	setHost(s.client, "_px_grid_direct")
	path := "/api/v1/pxgrid-direct/connector-config/{connectorName}"
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
