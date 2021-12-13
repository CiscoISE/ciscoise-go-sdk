package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type TelemetryService service

type ResponseTelemetryGetTransportGateway struct {
	Response *ResponseTelemetryGetTransportGatewayResponse `json:"response,omitempty"` // Transport Gateway acts a proxy for the communication between the ISE servers in your network and the Telemetry servers in case of air-gapped network.
	Version  string                                        `json:"version,omitempty"`  //
}

type ResponseTelemetryGetTransportGatewayResponse struct {
	EnableTransportGateway *bool  `json:"enableTransportGateway,omitempty"` // Indicates whether transport gateway is enabled or not.
	URL                    string `json:"url,omitempty"`                    // URL of transport gateway
}

type ResponseTelemetryUpdateTransportGateway struct {
	Response *ResponseTelemetryUpdateTransportGatewayResponse `json:"response,omitempty"` // Transport Gateway acts a proxy for the communication between the ISE servers in your network and the Telemetry servers in case of air-gapped network.
	Version  string                                           `json:"version,omitempty"`  //
}

type ResponseTelemetryUpdateTransportGatewayResponse struct {
	EnableTransportGateway *bool  `json:"enableTransportGateway,omitempty"` // Indicates whether transport gateway is enabled or not.
	URL                    string `json:"url,omitempty"`                    // URL of transport gateway
}

type RequestTelemetryUpdateTransportGateway struct {
	EnableTransportGateway *bool  `json:"enableTransportGateway,omitempty"` // Indicates whether transport gateway is enabled or not.
	URL                    string `json:"url,omitempty"`                    // URL of transport gateway
}

//GetTransportGateway Returns ISE transport gateway settings
/* Transport Gateway acts a proxy for the communication between the ISE servers in your network and the Telemetry servers in case of air-gapped network.

 */
func (s *TelemetryService) GetTransportGateway() (*ResponseTelemetryGetTransportGateway, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/system-settings/telemetry/transport-gateway"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTelemetryGetTransportGateway{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTransportGateway")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTelemetryGetTransportGateway)
	return result, response, err

}

//UpdateTransportGateway Configure ISE transport gateway settings
/* Transport Gateway acts a proxy for the communication between the ISE servers in your network and the Telemetry servers in case of air-gapped network.

 */
func (s *TelemetryService) UpdateTransportGateway(requestTelemetryUpdateTransportGateway *RequestTelemetryUpdateTransportGateway) (*ResponseTelemetryUpdateTransportGateway, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/system-settings/telemetry/transport-gateway"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTelemetryUpdateTransportGateway).
		SetResult(&ResponseTelemetryUpdateTransportGateway{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateTransportGateway")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTelemetryUpdateTransportGateway)
	return result, response, err

}
