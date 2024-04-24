package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type EndpointStopReplicationServiceService service

type ResponseEndpointStopReplicationServiceGetStopReplicationStatus struct {
	Response *ResponseEndpointStopReplicationServiceGetStopReplicationStatusResponse `json:"response,omitempty"` //
	Version  string                                                                  `json:"version,omitempty"`  //
}

type ResponseEndpointStopReplicationServiceGetStopReplicationStatusResponse struct {
	IsEnabled *bool `json:"isEnabled,omitempty"` //
}

type ResponseEndpointStopReplicationServiceSetStopReplicationService struct {
	Success *ResponseEndpointStopReplicationServiceSetStopReplicationServiceSuccess `json:"success,omitempty"` //
	Version string                                                                  `json:"version,omitempty"` //
}

type ResponseEndpointStopReplicationServiceSetStopReplicationServiceSuccess struct {
	Message string `json:"message,omitempty"` //
}

type RequestEndpointStopReplicationServiceSetStopReplicationService struct {
	IsEnabled *bool `json:"isEnabled,omitempty"` //
}

//GetStopReplicationStatus Retrieve the status of Endpoint stop replication Service
/* This API retrieves the status of Endpoint stop replication Service

 */
func (s *EndpointStopReplicationServiceService) GetStopReplicationStatus() (*ResponseEndpointStopReplicationServiceGetStopReplicationStatus, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/stop-replication"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEndpointStopReplicationServiceGetStopReplicationStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetStopReplicationStatus")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointStopReplicationServiceGetStopReplicationStatus)
	return result, response, err

}

//SetStopReplicationService Update the status of Endpoint stop replication Service
/* This API updates the status of Endpoint stop replication Service.

 */
func (s *EndpointStopReplicationServiceService) SetStopReplicationService(requestEndpointStopReplicationServiceSetStopReplicationService *RequestEndpointStopReplicationServiceSetStopReplicationService) (*ResponseEndpointStopReplicationServiceSetStopReplicationService, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/stop-replication"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEndpointStopReplicationServiceSetStopReplicationService).
		SetResult(&ResponseEndpointStopReplicationServiceSetStopReplicationService{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseEndpointStopReplicationServiceSetStopReplicationService{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation SetStopReplicationService")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseEndpointStopReplicationServiceSetStopReplicationService)
	return result, response, err

}
