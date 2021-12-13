package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type DeviceAdministrationNetworkConditionsService service

type ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditions struct {
	Response *[]ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsResponse `json:"response,omitempty"` //
	Version  string                                                                                  `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsResponse struct {
	ConditionType string                                                                                            `json:"conditionType,omitempty"` // This field determines the content of the conditions field
	Description   string                                                                                            `json:"description,omitempty"`   //
	ID            string                                                                                            `json:"id,omitempty"`            //
	Link          *ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsResponseLink         `json:"link,omitempty"`          //
	Name          string                                                                                            `json:"name,omitempty"`          // Network Condition name
	Conditions    *[]ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsResponseConditions `json:"conditions,omitempty"`    //
}

type ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsResponseConditions struct {
	CliDnisList     []string                                                                                            `json:"cliDnisList,omitempty"`     // <p>This field should contain a Caller ID (CLI), comma, and Called ID (DNIS).<br> Line format -  Caller ID (CLI), Called ID (DNIS)</p>
	ConditionType   string                                                                                              `json:"conditionType,omitempty"`   // This field determines the content of the conditions field
	Description     string                                                                                              `json:"description,omitempty"`     //
	ID              string                                                                                              `json:"id,omitempty"`              //
	IPAddrList      []string                                                                                            `json:"ipAddrList,omitempty"`      // <p>This field should contain IP-address-or-subnet,port number<br> IP address can be IPV4 format (n.n.n.n) or IPV6 format (n:n:n:n:n:n:n:n).<br> IP subnet can be IPV4 format (n.n.n.n/m) or IPV6 format (n:n:n:n:n:n:n:n/m).<br> Line format - IP Address or subnet,Port</p>
	Link            *ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsResponseConditionsLink `json:"link,omitempty"`            //
	MacAddrList     []string                                                                                            `json:"macAddrList,omitempty"`     // <p>This field should contain Endstation MAC address, comma, and Destination MAC addresses.<br> Each Max address must include twelve hexadecimal digits using formats nn:nn:nn:nn:nn:nn or nn-nn-nn-nn-nn-nn or nnnn.nnnn.nnnn or nnnnnnnnnnnn.<br> Line format - Endstation MAC,Destination MAC </p>
	Name            string                                                                                              `json:"name,omitempty"`            // Network Condition name
	DeviceGroupList []string                                                                                            `json:"deviceGroupList,omitempty"` // <p>This field should contain a tuple with NDG Root, comma, and an NDG (that it under the root).<br> Line format - NDG Root Name, NDG, Port</p>
	DeviceList      []string                                                                                            `json:"deviceList,omitempty"`      // <p>This field should contain Device-Name,port-number. The device name must be the same as the name field in a Network Device object.<br> Line format - Device Name,Port</p>
}

type ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsResponseConditionsLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkCondition struct {
	Response *ResponseDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionResponse `json:"response,omitempty"` // Unique network conditions to restrict access to the network
	Version  string                                                                                  `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionResponse struct {
	ConditionType string                                                                                              `json:"conditionType,omitempty"` // This field determines the content of the conditions field
	Description   string                                                                                              `json:"description,omitempty"`   //
	ID            string                                                                                              `json:"id,omitempty"`            //
	Link          *ResponseDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionResponseLink         `json:"link,omitempty"`          //
	Name          string                                                                                              `json:"name,omitempty"`          // Network Condition name
	Conditions    *[]ResponseDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionResponseConditions `json:"conditions,omitempty"`    //
}

type ResponseDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionResponseConditions struct {
	CliDnisList     []string                                                                                              `json:"cliDnisList,omitempty"`     // <p>This field should contain a Caller ID (CLI), comma, and Called ID (DNIS).<br> Line format -  Caller ID (CLI), Called ID (DNIS)</p>
	ConditionType   string                                                                                                `json:"conditionType,omitempty"`   // This field determines the content of the conditions field
	Description     string                                                                                                `json:"description,omitempty"`     //
	ID              string                                                                                                `json:"id,omitempty"`              //
	IPAddrList      []string                                                                                              `json:"ipAddrList,omitempty"`      // <p>This field should contain IP-address-or-subnet,port number<br> IP address can be IPV4 format (n.n.n.n) or IPV6 format (n:n:n:n:n:n:n:n).<br> IP subnet can be IPV4 format (n.n.n.n/m) or IPV6 format (n:n:n:n:n:n:n:n/m).<br> Line format - IP Address or subnet,Port</p>
	Link            *ResponseDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionResponseConditionsLink `json:"link,omitempty"`            //
	MacAddrList     []string                                                                                              `json:"macAddrList,omitempty"`     // <p>This field should contain Endstation MAC address, comma, and Destination MAC addresses.<br> Each Max address must include twelve hexadecimal digits using formats nn:nn:nn:nn:nn:nn or nn-nn-nn-nn-nn-nn or nnnn.nnnn.nnnn or nnnnnnnnnnnn.<br> Line format - Endstation MAC,Destination MAC </p>
	Name            string                                                                                                `json:"name,omitempty"`            // Network Condition name
	DeviceGroupList []string                                                                                              `json:"deviceGroupList,omitempty"` // <p>This field should contain a tuple with NDG Root, comma, and an NDG (that it under the root).<br> Line format - NDG Root Name, NDG, Port</p>
	DeviceList      []string                                                                                              `json:"deviceList,omitempty"`      // <p>This field should contain Device-Name,port-number. The device name must be the same as the name field in a Network Device object.<br> Line format - Device Name,Port</p>
}

type ResponseDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionResponseConditionsLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByID struct {
	Response *ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDResponse `json:"response,omitempty"` // Unique network conditions to restrict access to the network
	Version  string                                                                                   `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDResponse struct {
	ConditionType string                                                                                               `json:"conditionType,omitempty"` // This field determines the content of the conditions field
	Description   string                                                                                               `json:"description,omitempty"`   //
	ID            string                                                                                               `json:"id,omitempty"`            //
	Link          *ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDResponseLink         `json:"link,omitempty"`          //
	Name          string                                                                                               `json:"name,omitempty"`          // Network Condition name
	Conditions    *[]ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDResponseConditions `json:"conditions,omitempty"`    //
}

type ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDResponseConditions struct {
	CliDnisList     []string                                                                                               `json:"cliDnisList,omitempty"`     // <p>This field should contain a Caller ID (CLI), comma, and Called ID (DNIS).<br> Line format -  Caller ID (CLI), Called ID (DNIS)</p>
	ConditionType   string                                                                                                 `json:"conditionType,omitempty"`   // This field determines the content of the conditions field
	Description     string                                                                                                 `json:"description,omitempty"`     //
	ID              string                                                                                                 `json:"id,omitempty"`              //
	IPAddrList      []string                                                                                               `json:"ipAddrList,omitempty"`      // <p>This field should contain IP-address-or-subnet,port number<br> IP address can be IPV4 format (n.n.n.n) or IPV6 format (n:n:n:n:n:n:n:n).<br> IP subnet can be IPV4 format (n.n.n.n/m) or IPV6 format (n:n:n:n:n:n:n:n/m).<br> Line format - IP Address or subnet,Port</p>
	Link            *ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDResponseConditionsLink `json:"link,omitempty"`            //
	MacAddrList     []string                                                                                               `json:"macAddrList,omitempty"`     // <p>This field should contain Endstation MAC address, comma, and Destination MAC addresses.<br> Each Max address must include twelve hexadecimal digits using formats nn:nn:nn:nn:nn:nn or nn-nn-nn-nn-nn-nn or nnnn.nnnn.nnnn or nnnnnnnnnnnn.<br> Line format - Endstation MAC,Destination MAC </p>
	Name            string                                                                                                 `json:"name,omitempty"`            // Network Condition name
	DeviceGroupList []string                                                                                               `json:"deviceGroupList,omitempty"` // <p>This field should contain a tuple with NDG Root, comma, and an NDG (that it under the root).<br> Line format - NDG Root Name, NDG, Port</p>
	DeviceList      []string                                                                                               `json:"deviceList,omitempty"`      // <p>This field should contain Device-Name,port-number. The device name must be the same as the name field in a Network Device object.<br> Line format - Device Name,Port</p>
}

type ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDResponseConditionsLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByID struct {
	Response *ResponseDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDResponse `json:"response,omitempty"` // Unique network conditions to restrict access to the network
	Version  string                                                                                      `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDResponse struct {
	ConditionType string                                                                                                  `json:"conditionType,omitempty"` // This field determines the content of the conditions field
	Description   string                                                                                                  `json:"description,omitempty"`   //
	ID            string                                                                                                  `json:"id,omitempty"`            //
	Link          *ResponseDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDResponseLink         `json:"link,omitempty"`          //
	Name          string                                                                                                  `json:"name,omitempty"`          // Network Condition name
	Conditions    *[]ResponseDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDResponseConditions `json:"conditions,omitempty"`    //
}

type ResponseDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDResponseConditions struct {
	CliDnisList     []string                                                                                                  `json:"cliDnisList,omitempty"`     // <p>This field should contain a Caller ID (CLI), comma, and Called ID (DNIS).<br> Line format -  Caller ID (CLI), Called ID (DNIS)</p>
	ConditionType   string                                                                                                    `json:"conditionType,omitempty"`   // This field determines the content of the conditions field
	Description     string                                                                                                    `json:"description,omitempty"`     //
	ID              string                                                                                                    `json:"id,omitempty"`              //
	IPAddrList      []string                                                                                                  `json:"ipAddrList,omitempty"`      // <p>This field should contain IP-address-or-subnet,port number<br> IP address can be IPV4 format (n.n.n.n) or IPV6 format (n:n:n:n:n:n:n:n).<br> IP subnet can be IPV4 format (n.n.n.n/m) or IPV6 format (n:n:n:n:n:n:n:n/m).<br> Line format - IP Address or subnet,Port</p>
	Link            *ResponseDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDResponseConditionsLink `json:"link,omitempty"`            //
	MacAddrList     []string                                                                                                  `json:"macAddrList,omitempty"`     // <p>This field should contain Endstation MAC address, comma, and Destination MAC addresses.<br> Each Max address must include twelve hexadecimal digits using formats nn:nn:nn:nn:nn:nn or nn-nn-nn-nn-nn-nn or nnnn.nnnn.nnnn or nnnnnnnnnnnn.<br> Line format - Endstation MAC,Destination MAC </p>
	Name            string                                                                                                    `json:"name,omitempty"`            // Network Condition name
	DeviceGroupList []string                                                                                                  `json:"deviceGroupList,omitempty"` // <p>This field should contain a tuple with NDG Root, comma, and an NDG (that it under the root).<br> Line format - NDG Root Name, NDG, Port</p>
	DeviceList      []string                                                                                                  `json:"deviceList,omitempty"`      // <p>This field should contain Device-Name,port-number. The device name must be the same as the name field in a Network Device object.<br> Line format - Device Name,Port</p>
}

type ResponseDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDResponseConditionsLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationNetworkConditionsDeleteDeviceAdminNetworkConditionByID struct {
	ID string `json:"id,omitempty"` //
}

type RequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkCondition struct {
	ConditionType string                                                                                     `json:"conditionType,omitempty"` // This field determines the content of the conditions field
	Description   string                                                                                     `json:"description,omitempty"`   //
	ID            string                                                                                     `json:"id,omitempty"`            //
	Link          *RequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionLink         `json:"link,omitempty"`          //
	Name          string                                                                                     `json:"name,omitempty"`          // Network Condition name
	Conditions    *[]RequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionConditions `json:"conditions,omitempty"`    //
}

type RequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionConditions struct {
	CliDnisList     []string                                                                                     `json:"cliDnisList,omitempty"`     // <p>This field should contain a Caller ID (CLI), comma, and Called ID (DNIS).<br> Line format -  Caller ID (CLI), Called ID (DNIS)</p>
	ConditionType   string                                                                                       `json:"conditionType,omitempty"`   // This field determines the content of the conditions field
	Description     string                                                                                       `json:"description,omitempty"`     //
	ID              string                                                                                       `json:"id,omitempty"`              //
	IPAddrList      []string                                                                                     `json:"ipAddrList,omitempty"`      // <p>This field should contain IP-address-or-subnet,port number<br> IP address can be IPV4 format (n.n.n.n) or IPV6 format (n:n:n:n:n:n:n:n).<br> IP subnet can be IPV4 format (n.n.n.n/m) or IPV6 format (n:n:n:n:n:n:n:n/m).<br> Line format - IP Address or subnet,Port</p>
	Link            *RequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionConditionsLink `json:"link,omitempty"`            //
	MacAddrList     []string                                                                                     `json:"macAddrList,omitempty"`     // <p>This field should contain Endstation MAC address, comma, and Destination MAC addresses.<br> Each Max address must include twelve hexadecimal digits using formats nn:nn:nn:nn:nn:nn or nn-nn-nn-nn-nn-nn or nnnn.nnnn.nnnn or nnnnnnnnnnnn.<br> Line format - Endstation MAC,Destination MAC </p>
	Name            string                                                                                       `json:"name,omitempty"`            // Network Condition name
	DeviceGroupList []string                                                                                     `json:"deviceGroupList,omitempty"` // <p>This field should contain a tuple with NDG Root, comma, and an NDG (that it under the root).<br> Line format - NDG Root Name, NDG, Port</p>
	DeviceList      []string                                                                                     `json:"deviceList,omitempty"`      // <p>This field should contain Device-Name,port-number. The device name must be the same as the name field in a Network Device object.<br> Line format - Device Name,Port</p>
}

type RequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionConditionsLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByID struct {
	ConditionType string                                                                                         `json:"conditionType,omitempty"` // This field determines the content of the conditions field
	Description   string                                                                                         `json:"description,omitempty"`   //
	ID            string                                                                                         `json:"id,omitempty"`            //
	Link          *RequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDLink         `json:"link,omitempty"`          //
	Name          string                                                                                         `json:"name,omitempty"`          // Network Condition name
	Conditions    *[]RequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDConditions `json:"conditions,omitempty"`    //
}

type RequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDConditions struct {
	CliDnisList     []string                                                                                         `json:"cliDnisList,omitempty"`     // <p>This field should contain a Caller ID (CLI), comma, and Called ID (DNIS).<br> Line format -  Caller ID (CLI), Called ID (DNIS)</p>
	ConditionType   string                                                                                           `json:"conditionType,omitempty"`   // This field determines the content of the conditions field
	Description     string                                                                                           `json:"description,omitempty"`     //
	ID              string                                                                                           `json:"id,omitempty"`              //
	IPAddrList      []string                                                                                         `json:"ipAddrList,omitempty"`      // <p>This field should contain IP-address-or-subnet,port number<br> IP address can be IPV4 format (n.n.n.n) or IPV6 format (n:n:n:n:n:n:n:n).<br> IP subnet can be IPV4 format (n.n.n.n/m) or IPV6 format (n:n:n:n:n:n:n:n/m).<br> Line format - IP Address or subnet,Port</p>
	Link            *RequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDConditionsLink `json:"link,omitempty"`            //
	MacAddrList     []string                                                                                         `json:"macAddrList,omitempty"`     // <p>This field should contain Endstation MAC address, comma, and Destination MAC addresses.<br> Each Max address must include twelve hexadecimal digits using formats nn:nn:nn:nn:nn:nn or nn-nn-nn-nn-nn-nn or nnnn.nnnn.nnnn or nnnnnnnnnnnn.<br> Line format - Endstation MAC,Destination MAC </p>
	Name            string                                                                                           `json:"name,omitempty"`            // Network Condition name
	DeviceGroupList []string                                                                                         `json:"deviceGroupList,omitempty"` // <p>This field should contain a tuple with NDG Root, comma, and an NDG (that it under the root).<br> Line format - NDG Root Name, NDG, Port</p>
	DeviceList      []string                                                                                         `json:"deviceList,omitempty"`      // <p>This field should contain Device-Name,port-number. The device name must be the same as the name field in a Network Device object.<br> Line format - Device Name,Port</p>
}

type RequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDConditionsLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

//GetDeviceAdminNetworkConditions Device Admin - Returns a list of network conditions.
/* Device Admin Returns a list of network conditions.

 */
func (s *DeviceAdministrationNetworkConditionsService) GetDeviceAdminNetworkConditions() (*ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditions, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/network-condition"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminNetworkConditions")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditions)
	return result, response, err

}

//GetDeviceAdminNetworkConditionByID Device Admin - Returns a network condition.
/* Device Admin Returns a network condition.

@param id id path parameter. Condition id
*/
func (s *DeviceAdministrationNetworkConditionsService) GetDeviceAdminNetworkConditionByID(id string) (*ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/network-condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminNetworkConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByID)
	return result, response, err

}

//CreateDeviceAdminNetworkCondition Device Admin- Creates network condition.
/* Device AdminCreates network condition.

 */
func (s *DeviceAdministrationNetworkConditionsService) CreateDeviceAdminNetworkCondition(requestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkCondition *RequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkCondition) (*ResponseDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkCondition, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/network-condition"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkCondition).
		SetResult(&ResponseDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkCondition{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceAdminNetworkCondition")
	}

	result := response.Result().(*ResponseDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkCondition)
	return result, response, err

}

//UpdateDeviceAdminNetworkConditionByID Device Admin - Update network condition.
/* Device Admin Update network condition.

@param id id path parameter. Condition id
*/
func (s *DeviceAdministrationNetworkConditionsService) UpdateDeviceAdminNetworkConditionByID(id string, requestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionById *RequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByID) (*ResponseDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/network-condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionById).
		SetResult(&ResponseDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceAdminNetworkConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByID)
	return result, response, err

}

//DeleteDeviceAdminNetworkConditionByID Device Admin - Delete network condition.
/* Device Admin Delete network condition.

@param id id path parameter. Condition id
*/
func (s *DeviceAdministrationNetworkConditionsService) DeleteDeviceAdminNetworkConditionByID(id string) (*ResponseDeviceAdministrationNetworkConditionsDeleteDeviceAdminNetworkConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/network-condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationNetworkConditionsDeleteDeviceAdminNetworkConditionByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteDeviceAdminNetworkConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationNetworkConditionsDeleteDeviceAdminNetworkConditionByID)
	return result, response, err

}
