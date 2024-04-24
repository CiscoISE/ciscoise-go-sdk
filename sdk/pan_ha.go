package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type PanHaService service

type ResponsePanHaGetPanHaStatus struct {
	Response *ResponsePanHaGetPanHaStatusResponse `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  //
}

type ResponsePanHaGetPanHaStatusResponse struct {
	IsEnabled                *bool                                                        `json:"isEnabled,omitempty"`                //
	PrimaryHealthCheckNode   *ResponsePanHaGetPanHaStatusResponsePrimaryHealthCheckNode   `json:"primaryHealthCheckNode,omitempty"`   //
	SecondaryHealthCheckNode *ResponsePanHaGetPanHaStatusResponseSecondaryHealthCheckNode `json:"secondaryHealthCheckNode,omitempty"` //
	PollingInterval          *int                                                         `json:"pollingInterval,omitempty"`          // Administration nodes are checked after each interval. Seconds (30 - 300) <br> The default value is 120.
	FailedAttempts           *int                                                         `json:"failedAttempts,omitempty"`           // Failover occurs if the primary PAN is down for the specified number of failure polls. Count (2 - 60).<br> The default value is 5.
}

type ResponsePanHaGetPanHaStatusResponsePrimaryHealthCheckNode struct {
	Hostname string `json:"hostname,omitempty"` //
}

type ResponsePanHaGetPanHaStatusResponseSecondaryHealthCheckNode struct {
	Hostname string `json:"hostname,omitempty"` //
}

type ResponsePanHaUpdatePanHa struct {
	Success *ResponsePanHaUpdatePanHaSuccess `json:"success,omitempty"` //
	Version string                           `json:"version,omitempty"` //
}

type ResponsePanHaUpdatePanHaSuccess struct {
	Message string `json:"message,omitempty"` //
}

type RequestPanHaUpdatePanHa struct {
	IsEnabled                *bool                                            `json:"isEnabled,omitempty"`                //
	PrimaryHealthCheckNode   *RequestPanHaUpdatePanHaPrimaryHealthCheckNode   `json:"primaryHealthCheckNode,omitempty"`   //
	SecondaryHealthCheckNode *RequestPanHaUpdatePanHaSecondaryHealthCheckNode `json:"secondaryHealthCheckNode,omitempty"` //
	PollingInterval          *int                                             `json:"pollingInterval,omitempty"`          // Administration nodes are checked after each interval. Seconds (30 - 300) <br> The default value is 120.
	FailedAttempts           *int                                             `json:"failedAttempts,omitempty"`           // Failover occurs if the primary PAN is down for the specified number of failure polls. Count (2 - 60).<br> The default value is 5.
}

type RequestPanHaUpdatePanHaPrimaryHealthCheckNode struct {
	Hostname string `json:"hostname,omitempty"` //
}

type RequestPanHaUpdatePanHaSecondaryHealthCheckNode struct {
	Hostname string `json:"hostname,omitempty"` //
}

//GetPanHaStatus Get the current configuration of the PAN HA.
/* In a high availability configuration, the primary PAN is in active state. The secondary PAN (backup PAN) is in standby state, which means that it receives all the configuration updates from the primary PAN, but is not active in the Cisco ISE cluster. You can configure Cisco ISE to automatically promote the secondary PAN when the primary PAN becomes unavailable.

 */
func (s *PanHaService) GetPanHaStatus() (*ResponsePanHaGetPanHaStatus, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/pan-ha"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePanHaGetPanHaStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPanHaStatus")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePanHaGetPanHaStatus)
	return result, response, err

}

//UpdatePanHa Enable, update or disable PAN failover configuration.
/* To deploy the auto-failover feature, you must have at least three nodes, where two of the nodes assume the Administration persona, and one node acts as the health check node. A health check node is a non-administration node and can be a Policy Service, Monitoring, or pxGrid node, or any combination of these. If the PANs are in different data centers, you must have a health check node for each PAN.
All the fields are mandatory to enable PanHA.
Values of failedAttempts, pollingInterval, primaryHealthCheckNode, and secondaryHealthCheckNode are not considered when the isEnable value is "false" in the request body.

*/
func (s *PanHaService) UpdatePanHa(requestPanHaUpdatePanHA *RequestPanHaUpdatePanHa) (*ResponsePanHaUpdatePanHa, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/pan-ha"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestPanHaUpdatePanHA).
		SetResult(&ResponsePanHaUpdatePanHa{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponsePanHaUpdatePanHa{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdatePanHa")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePanHaUpdatePanHa)
	return result, response, err

}
