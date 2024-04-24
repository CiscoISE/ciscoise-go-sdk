package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ConfigurationService service

type ResponseConfigurationGetConfiguration struct {
	EnableEPO *bool `json:"enableEPO,omitempty"` // To enable/disable LSD ownership settings
	EnableRCM *bool `json:"enableRCM,omitempty"` // To enable/disable random mac(RCM) settings. Please note that this flag will be set to false if enableEPO flag is false
}

type ResponseConfigurationUpdateConfiguration struct {
	EnableEPO *bool `json:"enableEPO,omitempty"` // To enable/disable LSD ownership settings
	EnableRCM *bool `json:"enableRCM,omitempty"` // To enable/disable random mac(RCM) settings. Please note that this flag will be set to false if enableEPO flag is false
}

type RequestConfigurationUpdateConfiguration struct {
	EnableEPO *bool `json:"enableEPO,omitempty"` // To enable/disable LSD ownership settings
	EnableRCM *bool `json:"enableRCM,omitempty"` // To enable/disable random mac(RCM) settings. Please note that this flag will be set to false if enableEPO flag is false
}

//GetConfiguration Retrieve configuration information for LSD settings page
/* Retrieve configuration information for LSD settings page

 */
func (s *ConfigurationService) GetConfiguration() (*ResponseConfigurationGetConfiguration, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/lsd/updateLsdSettings"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationGetConfiguration{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetConfiguration")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseConfigurationGetConfiguration)
	return result, response, err

}

//UpdateConfiguration Update LSD enable/disable settings for endppint ownership and rcm (random changing mac)
/* API to enable/disable LSD settings

 */
func (s *ConfigurationService) UpdateConfiguration(requestConfigurationUpdateConfiguration *RequestConfigurationUpdateConfiguration) (*ResponseConfigurationUpdateConfiguration, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/lsd/updateLsdSettings"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationUpdateConfiguration).
		SetResult(&ResponseConfigurationUpdateConfiguration{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseConfigurationUpdateConfiguration{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateConfiguration")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseConfigurationUpdateConfiguration)
	return result, response, err

}
