package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type LicensingService service

type ResponseLicensingGetConnectionType struct {
	ConnectionType string `json:"connectionType,omitempty"` //
	State          string `json:"state,omitempty"`          //
}

type ResponseLicensingGetEvalLicenseInfo struct {
	DaysRemaining *int `json:"daysRemaining,omitempty"` // Number of eval license days left
}

type ResponseLicensingGetFeatureToTierMapping []ResponseItemLicensingGetFeatureToTierMapping // Array of ResponseLicensingGetFeatureToTierMapping

type ResponseItemLicensingGetFeatureToTierMapping struct {
	FeatureName []string `json:"featureName,omitempty"` // <p>List of feature names</p>
	Tier        string   `json:"tier,omitempty"`        //
}

type ResponseLicensingGetRegistrationInfo struct {
	Response *ResponseLicensingGetRegistrationInfoResponse `json:"response,omitempty"` // Registration information format
	Version  string                                        `json:"version,omitempty"`  //
}

type ResponseLicensingGetRegistrationInfoResponse struct {
	ConnectionType    string   `json:"connectionType,omitempty"`    //
	RegistrationState string   `json:"registrationState,omitempty"` //
	SsmOnPremServer   string   `json:"ssmOnPremServer,omitempty"`   // If connection type is selected as SSM_ONPREM_SERVER, then  IP address or the hostname (or FQDN) of the SSM On-Prem server Host.
	Tier              []string `json:"tier,omitempty"`              //
}

// # Review unknown case

type ResponseLicensingGetSmartState struct {
	Response *ResponseLicensingGetSmartStateResponse `json:"response,omitempty"` // State of the smart license format
	Version  string                                  `json:"version,omitempty"`  //
}

type ResponseLicensingGetSmartStateResponse struct {
	ConnectionType string `json:"connectionType,omitempty"` //
	State          string `json:"state,omitempty"`          //
}

// # Review unknown case

type ResponseLicensingGetTierStateInfo []ResponseItemLicensingGetTierStateInfo // Array of ResponseLicensingGetTierStateInfo

type ResponseItemLicensingGetTierStateInfo struct {
	Compliance          string `json:"compliance,omitempty"`          //
	ConsumptionCounter  *int   `json:"consumptionCounter,omitempty"`  // Compliance counter for tier
	DaysOutOfCompliance string `json:"daysOutOfCompliance,omitempty"` // Number of days tier is out of compliance
	LastAuthorization   string `json:"lastAuthorization,omitempty"`   // Last date of authorization
	Name                string `json:"name,omitempty"`                //
	Status              string `json:"status,omitempty"`              //
}

type ResponseLicensingUpdateTierStateInfo struct {
	Response *[]ResponseLicensingUpdateTierStateInfoResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  //
}

type ResponseLicensingUpdateTierStateInfoResponse struct {
	Message string `json:"message,omitempty"` // Response message on successful change of license tier state.
	Name    string `json:"name,omitempty"`    //
	Status  string `json:"status,omitempty"`  //
}

type RequestLicensingCreateRegistrationInfo struct {
	ConnectionType   string   `json:"connectionType,omitempty"`   //
	RegistrationType string   `json:"registrationType,omitempty"` //
	SsmOnPremServer  string   `json:"ssmOnPremServer,omitempty"`  // If connection type is selected as SSM_ONPREM_SERVER, then  IP address or the hostname (or FQDN) of the SSM On-Prem server Host.
	Tier             []string `json:"tier,omitempty"`             //
	Token            string   `json:"token,omitempty"`            // token
}

type RequestLicensingUpdateTierStateInfo []RequestItemLicensingUpdateTierStateInfo // Array of RequestLicensingUpdateTierStateInfo

type RequestItemLicensingUpdateTierStateInfo struct {
	Name   string `json:"name,omitempty"`   //
	Status string `json:"status,omitempty"` //
}

//GetConnectionType License - Connection Type
/* Get connection type

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

//GetEvalLicenseInfo License - registration information
/* Get registration information

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

//GetFeatureToTierMapping License - feature to tier mapping
/* Get feature to tier mapping

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

//GetRegistrationInfo License - registration information
/* Get registration information

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

//GetSmartState License - smart license information
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

//CreateRegistrationInfo License - registration information
/* License Configure registration information.

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

//ConfigureSmartState License - smart state information
/* License Configure smart state information.

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
