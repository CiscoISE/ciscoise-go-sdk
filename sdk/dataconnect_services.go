package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type DataconnectServicesService service

type ResponseDataconnectServicesGetOdbcDetail struct {
	Response *ResponseDataconnectServicesGetOdbcDetailResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}

type ResponseDataconnectServicesGetOdbcDetailResponse struct {
	Hostname    string `json:"hostname,omitempty"`    //
	Port        *int   `json:"port,omitempty"`        //
	Servicename string `json:"servicename,omitempty"` //
	Username    string `json:"username,omitempty"`    //
}

type ResponseDataconnectServicesGetDataconnectService struct {
	Response *ResponseDataconnectServicesGetDataconnectServiceResponse `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  //
}

type ResponseDataconnectServicesGetDataconnectServiceResponse struct {
	IsEnabled             *bool  `json:"isEnabled,omitempty"`             //
	IsPasswordChanged     *bool  `json:"isPasswordChanged,omitempty"`     //
	PasswordExpiresInDays *int   `json:"passwordExpiresInDays,omitempty"` //
	PasswordExpiresOn     string `json:"passwordExpiresOn,omitempty"`     //
}

type ResponseDataconnectServicesUpdateDataconnectPassword struct {
	Success *ResponseDataconnectServicesUpdateDataconnectPasswordSuccess `json:"success,omitempty"` //
	Version string                                                       `json:"version,omitempty"` //
}

type ResponseDataconnectServicesUpdateDataconnectPasswordSuccess struct {
	Message string `json:"message,omitempty"` //
}

type ResponseDataconnectServicesUpdateDataconnectPasswordExpiry struct {
	Success *ResponseDataconnectServicesUpdateDataconnectPasswordExpirySuccess `json:"success,omitempty"` //
	Version string                                                             `json:"version,omitempty"` //
}

type ResponseDataconnectServicesUpdateDataconnectPasswordExpirySuccess struct {
	Message string `json:"message,omitempty"` //
}

type ResponseDataconnectServicesSetDataConnectService struct {
	Success *ResponseDataconnectServicesSetDataConnectServiceSuccess `json:"success,omitempty"` //
	Version string                                                   `json:"version,omitempty"` //
}

type ResponseDataconnectServicesSetDataConnectServiceSuccess struct {
	Message string `json:"message,omitempty"` //
}

type RequestDataconnectServicesUpdateDataconnectPassword struct {
	Password string `json:"password,omitempty"` // Password for dataconnect user. Password must satisfy the following criteria - Contains at least one uppercase letter [A-Z], Contains at least one lowercase letter [a-z], Contains at least one digit [0-9], Contains at least one special character #$%&*+,-.:;=?^_~, Has at least 12 characters, Has not more than 30 characters
}

type RequestDataconnectServicesUpdateDataconnectPasswordExpiry struct {
	PasswordExpiresInDays *int `json:"passwordExpiresInDays,omitempty"` //
}

type RequestDataconnectServicesSetDataConnectService struct {
	IsEnabled *bool `json:"isEnabled,omitempty"` //
}

//GetOdbcDetail Retrieve the Dataconnect ODBC details.
/* This API retrieves the Dataconnect ODBC details.

 */
func (s *DataconnectServicesService) GetOdbcDetail() (*ResponseDataconnectServicesGetOdbcDetail, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/mnt/data-connect/details"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDataconnectServicesGetOdbcDetail{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOdbcDetail")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDataconnectServicesGetOdbcDetail)
	return result, response, err

}

//GetDataconnectService Retrieve the status of the Dataconnect feature
/* This API retrieves the status of the Dataconnect feature.

 */
func (s *DataconnectServicesService) GetDataconnectService() (*ResponseDataconnectServicesGetDataconnectService, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/mnt/data-connect/settings"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDataconnectServicesGetDataconnectService{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDataconnectService")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDataconnectServicesGetDataconnectService)
	return result, response, err

}

//UpdateDataconnectPassword Update the Dataconnect user password.
/* This API updates the Dataconnect user password

 */
func (s *DataconnectServicesService) UpdateDataconnectPassword(requestDataconnectServicesUpdateDataconnectPassword *RequestDataconnectServicesUpdateDataconnectPassword) (*ResponseDataconnectServicesUpdateDataconnectPassword, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/mnt/data-connect/settings/password"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDataconnectServicesUpdateDataconnectPassword).
		SetResult(&ResponseDataconnectServicesUpdateDataconnectPassword{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseDataconnectServicesUpdateDataconnectPassword{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDataconnectPassword")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDataconnectServicesUpdateDataconnectPassword)
	return result, response, err

}

//UpdateDataconnectPasswordExpiry Updates the number of days of Dataconnect password expiry.
/* This API updates the number of days of Dataconnect password expiry

 */
func (s *DataconnectServicesService) UpdateDataconnectPasswordExpiry(requestDataconnectServicesUpdateDataconnectPasswordExpiry *RequestDataconnectServicesUpdateDataconnectPasswordExpiry) (*ResponseDataconnectServicesUpdateDataconnectPasswordExpiry, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/mnt/data-connect/settings/password/expiry"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDataconnectServicesUpdateDataconnectPasswordExpiry).
		SetResult(&ResponseDataconnectServicesUpdateDataconnectPasswordExpiry{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseDataconnectServicesUpdateDataconnectPasswordExpiry{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDataconnectPasswordExpiry")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDataconnectServicesUpdateDataconnectPasswordExpiry)
	return result, response, err

}

//SetDataConnectService Update the DataConnect feature status
/* This API updates the DataConnect feature status.

 */
func (s *DataconnectServicesService) SetDataConnectService(requestDataconnectServicesSetDataConnectService *RequestDataconnectServicesSetDataConnectService) (*ResponseDataconnectServicesSetDataConnectService, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/mnt/data-connect/settings/status"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDataconnectServicesSetDataConnectService).
		SetResult(&ResponseDataconnectServicesSetDataConnectService{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseDataconnectServicesSetDataConnectService{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation SetDataConnectService")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDataconnectServicesSetDataConnectService)
	return result, response, err

}
