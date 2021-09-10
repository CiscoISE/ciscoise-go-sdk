package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type PanHaService service

type ResponsePanHaGetPanHaStatus struct {
	Response *[]ResponsePanHaGetPanHaStatusResponse `json:"response,omitempty"` //
}

type ResponsePanHaGetPanHaStatusResponse struct {
	IsEnabled                *bool  `json:"isEnabled,omitempty"`                //
	PrimaryHealthCheckNode   string `json:"primaryHealthCheckNode,omitempty"`   //
	SecondaryHealthCheckNode string `json:"secondaryHealthCheckNode,omitempty"` //
	PollingInterval          *int   `json:"pollingInterval,omitempty"`          //
	FailedAttempts           *int   `json:"failedAttempts,omitempty"`           //
}

type ResponsePanHaEnablePanHa struct {
	Code      *int   `json:"code,omitempty"`      //
	Message   string `json:"message,omitempty"`   //
	RootCause string `json:"rootCause,omitempty"` //
}

type ResponsePanHaDisablePanHa struct {
	Code      *int   `json:"code,omitempty"`      //
	Message   string `json:"message,omitempty"`   //
	RootCause string `json:"rootCause,omitempty"` //
}

type RequestPanHaEnablePanHa struct {
	Request *RequestPanHaEnablePanHaRequest `json:"request,omitempty"` //
}

type RequestPanHaEnablePanHaRequest struct {
	IsEnabled                *bool  `json:"isEnabled,omitempty"`                //
	PrimaryHealthCheckNode   string `json:"primaryHealthCheckNode,omitempty"`   //
	SecondaryHealthCheckNode string `json:"secondaryHealthCheckNode,omitempty"` //
	PollingInterval          *int   `json:"pollingInterval,omitempty"`          //
	FailedAttempts           *int   `json:"failedAttempts,omitempty"`           //
}

//GetPanHaStatus Get current status of the PAN HA.
/* In a high availability configuration, the Primary Administration Node (PAN) is in the active state. The Secondary PAN (backup PAN) is in the standby state, which means it receives all configuration updates from the Primary PAN, but is not active in the ISE network. You can configure ISE to automatically the promote the secondary PAN when the primary PAN becomes unavailable.

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

//EnablePanHa API to enable/update PAN failover..
/* To deploy the auto-failover feature, you must have at least three nodes, where two of the nodes assume the Administration persona, and one node acts as the health check node. A health check node is a non-administration node and can be a Policy Service, Monitoring, or pxGrid node, or a combination of these. If the PANs are in different data centers, you must have a health check node for each PAN.

 */
func (s *PanHaService) EnablePanHa(requestPanHaEnablePanHA *RequestPanHaEnablePanHa) (*ResponsePanHaEnablePanHa, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/pan-ha"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestPanHaEnablePanHA).
		SetResult(&ResponsePanHaEnablePanHa{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation EnablePanHa")
	}

	result := response.Result().(*ResponsePanHaEnablePanHa)
	return result, response, err

}

//DisablePanHa Disable PAN failover..
/* Disable the automatic PAN failover

 */
func (s *PanHaService) DisablePanHa() (*ResponsePanHaDisablePanHa, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/pan-ha"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePanHaDisablePanHa{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DisablePanHa")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePanHaDisablePanHa)
	return result, response, err

}
