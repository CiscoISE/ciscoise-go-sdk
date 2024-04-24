package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type LicensingService service

type ResponseLicensingGetSmartState struct {
	Response *ResponseLicensingGetSmartStateResponse `json:"response,omitempty"` // State of the smart license format
	Version  string                                  `json:"version,omitempty"`  //
}

type ResponseLicensingGetSmartStateResponse struct {
	ConnectionType string `json:"connectionType,omitempty"` // Type of connection for the registration
	State          string `json:"state,omitempty"`          // Current Smart Licensing State
}

// # Review unknown case

type ResponseLicensingGetRegistrationInfo struct {
	Response *ResponseLicensingGetRegistrationInfoResponse `json:"response,omitempty"` // Registration information format
	Version  string                                        `json:"version,omitempty"`  //
}

type ResponseLicensingGetRegistrationInfoResponse struct {
	Tier              []string `json:"tier,omitempty"`              //
	ConnectionType    string   `json:"connectionType,omitempty"`    // Type of connection for the registration
	RegistrationState string   `json:"registrationState,omitempty"` // Registered or Unregistered
	SsmOnPremServer   string   `json:"ssmOnPremServer,omitempty"`   // If connection type is selected as SSM_ONPREM_SERVER, then the IP address or the hostname (or FQDN) of the SSM On-Prem server host.
}

// # Review unknown case

type ResponseLicensingGetTierStateInfo []ResponseItemLicensingGetTierStateInfo // Array of ResponseLicensingGetTierStateInfo

type ResponseItemLicensingGetTierStateInfo struct {
	Name                string `json:"name,omitempty"`                // License tier list
	Status              string `json:"status,omitempty"`              // Current Smart Licensing State
	Compliance          string `json:"compliance,omitempty"`          // Tier compliance state
	ConsumptionCounter  *int   `json:"consumptionCounter,omitempty"`  // Compliance counter for tier
	DaysOutOfCompliance string `json:"daysOutOfCompliance,omitempty"` // Number of days, the tier is out of compliance
	LastAuthorization   string `json:"lastAuthorization,omitempty"`   // Last date of authorization
}

type ResponseLicensingUpdateTierStateInfo struct {
	Version  string                                          `json:"version,omitempty"`  //
	Response *[]ResponseLicensingUpdateTierStateInfoResponse `json:"response,omitempty"` //
}

type ResponseLicensingUpdateTierStateInfoResponse struct {
	Name    string `json:"name,omitempty"`    // License tier list
	Status  string `json:"status,omitempty"`  // Current Smart Licensing State
	Message string `json:"message,omitempty"` // Response message received on successful change of license tier state.
}

type ResponseLicensingGetEvalLicenseInfo struct {
	DaysRemaining *int `json:"daysRemaining,omitempty"` // Number of eval license days left
}

type ResponseLicensingGetConnectionType struct {
	ConnectionType string `json:"connectionType,omitempty"` // Type of connection for the registration
	State          string `json:"state,omitempty"`          // Current Smart Licensing State
}

type ResponseLicensingGetFeatureToTierMapping []ResponseItemLicensingGetFeatureToTierMapping // Array of ResponseLicensingGetFeatureToTierMapping

type ResponseItemLicensingGetFeatureToTierMapping struct {
	Tier        string   `json:"tier,omitempty"`        // License tier list
	FeatureName []string `json:"featureName,omitempty"` // <p>List of feature names</p>
}

type RequestLicensingCreateRegistrationInfo struct {
	Token            string   `json:"token,omitempty"`            // token
	Tier             []string `json:"tier,omitempty"`             //
	ConnectionType   string   `json:"connectionType,omitempty"`   // Type of connection for the registration
	RegistrationType string   `json:"registrationType,omitempty"` // Register, deregister, renew or update
	SsmOnPremServer  string   `json:"ssmOnPremServer,omitempty"`  // If the connection type is selected as SSM_ONPREM_SERVER, then this field consists of IP address or the hostname (or FQDN) of the SSM On-Prem server host.
}

type RequestLicensingUpdateTierStateInfo []RequestItemLicensingUpdateTierStateInfo // Array of RequestLicensingUpdateTierStateInfo

type RequestItemLicensingUpdateTierStateInfo struct {
	Name   string `json:"name,omitempty"`   // License tier list
	Status string `json:"status,omitempty"` // Current Smart Licensing State
}

//GetSmartState Get smart license information
/* Get smart license information

 */
func (s *LicensingService) GetSmartState() (*ResponseLicensingGetSmartState, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/license/system/smart-state"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensingGetSmartState{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSmartState")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseLicensingGetSmartState)
	return result, response, err

}

//GetRegistrationInfo Get the registration information
/* Get the registration information

 */
func (s *LicensingService) GetRegistrationInfo() (*ResponseLicensingGetRegistrationInfo, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/license/system/register"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensingGetRegistrationInfo{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetRegistrationInfo")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseLicensingGetRegistrationInfo)
	return result, response, err

}

//GetTierStateInfo Get tier state information
/* Get tier state information

 */
func (s *LicensingService) GetTierStateInfo() (*ResponseLicensingGetTierStateInfo, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/license/system/tier-state"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensingGetTierStateInfo{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTierStateInfo")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseLicensingGetTierStateInfo)
	return result, response, err

}

//GetEvalLicenseInfo Get registration information related to the evaluation license
/* Get registration information related to the evaluation license

 */
func (s *LicensingService) GetEvalLicenseInfo() (*ResponseLicensingGetEvalLicenseInfo, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/license/system/eval-license"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensingGetEvalLicenseInfo{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetEvalLicenseInfo")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseLicensingGetEvalLicenseInfo)
	return result, response, err

}

//GetConnectionType Get the connection type
/* Get the connection type

 */
func (s *LicensingService) GetConnectionType() (*ResponseLicensingGetConnectionType, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/license/system/connection-type"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensingGetConnectionType{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetConnectionType")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseLicensingGetConnectionType)
	return result, response, err

}

//GetFeatureToTierMapping Get feature to tier mapping information
/* Get feature to tier mapping information

 */
func (s *LicensingService) GetFeatureToTierMapping() (*ResponseLicensingGetFeatureToTierMapping, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/license/system/feature-to-tier-mapping"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensingGetFeatureToTierMapping{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetFeatureToTierMapping")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseLicensingGetFeatureToTierMapping)
	return result, response, err

}

//ConfigureSmartState Configure smart state information
/* Configure smart state information.

 */
func (s *LicensingService) ConfigureSmartState() (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/license/system/smart-state"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation ConfigureSmartState")
	}

	return response, err

}

//CreateRegistrationInfo Configure registration information
/* Configure registration information.

 */
func (s *LicensingService) CreateRegistrationInfo(requestLicensingCreateRegistrationInfo *RequestLicensingCreateRegistrationInfo) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/license/system/register"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLicensingCreateRegistrationInfo).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateRegistrationInfo")
	}

	return response, err

}

//UpdateTierStateInfo Post tier state information
/* Applicable values for
name
 &
status
 parameters:


name:
 ESSENTIAL, ADVANTAGE, PREMIER, DEVICEADMIN

status:
 ENABLED, DISABLED


*/
func (s *LicensingService) UpdateTierStateInfo(requestLicensingUpdateTierStateInfo *RequestLicensingUpdateTierStateInfo) (*ResponseLicensingUpdateTierStateInfo, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/license/system/tier-state"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLicensingUpdateTierStateInfo).
		SetResult(&ResponseLicensingUpdateTierStateInfo{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseLicensingUpdateTierStateInfo{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateTierStateInfo")
	}

	result := response.Result().(*ResponseLicensingUpdateTierStateInfo)
	return result, response, err

}
