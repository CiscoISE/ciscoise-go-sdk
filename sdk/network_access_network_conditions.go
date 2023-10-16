package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type NetworkAccessNetworkConditionsService service

type ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions struct {
	Response *[]ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsResponse `json:"response,omitempty"` //
	Version  string                                                                             `json:"version,omitempty"`  //
}

type ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsResponse struct {
	ConditionType   string                                                                               `json:"conditionType,omitempty"`   // This field determines the content of the conditions field
	Description     string                                                                               `json:"description,omitempty"`     //
	ID              string                                                                               `json:"id,omitempty"`              //
	Link            *ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsResponseLink `json:"link,omitempty"`            //
	Name            string                                                                               `json:"name,omitempty"`            // Network Condition name
	DeviceList      []string                                                                             `json:"deviceList,omitempty"`      // <p>This field should contain Device Name. The device name must be the same as the name field in a Network Device object. Line format - Device Name</p>
	CliDnisList     []string                                                                             `json:"cliDnisList,omitempty"`     // <p>This field should contain a Caller ID (CLI), comma, and Called ID (DNIS).<br> Line format -  Caller ID (CLI), Called ID (DNIS)</p>
	IPAddrList      []string                                                                             `json:"ipAddrList,omitempty"`      // <p>This field should contain IP address or subnet.<br> IP address can be IPV4 format (n.n.n.n) or IPV6 format (n:n:n:n:n:n:n:n).<br> IP subnet can be IPV4 format (n.n.n.n/m) or IPV6 format (n:n:n:n:n:n:n:n/m).<br> Line format - IP Address or subnet</p>
	MacAddrList     []string                                                                             `json:"macAddrList,omitempty"`     // <p>This field should contain Endstation MAC address, comma, and Destination MAC addresses.<br> Each Max address must include twelve hexadecimal digits using formats nn:nn:nn:nn:nn:nn or nn-nn-nn-nn-nn-nn or nnnn.nnnn.nnnn or nnnnnnnnnnnn.<br> Line format - Endstation MAC,Destination MAC </p>
	DeviceGroupList []string                                                                             `json:"deviceGroupList,omitempty"` // <p>This field should contain a tuple with NDG Root, comma, and an NDG (that it under the root).<br> Line format - NDG Root Name, NDG, Port</p>
}

type ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessNetworkConditionsCreateNetworkAccessNetworkCondition struct {
	Response *ResponseNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionResponse `json:"response,omitempty"` // Unique network conditions to restrict access to the network
	Version  string                                                                             `json:"version,omitempty"`  //
}

type ResponseNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionResponse struct {
	ConditionType   string                                                                                 `json:"conditionType,omitempty"`   // This field determines the content of the conditions field
	Description     string                                                                                 `json:"description,omitempty"`     //
	ID              string                                                                                 `json:"id,omitempty"`              //
	Link            *ResponseNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionResponseLink `json:"link,omitempty"`            //
	Name            string                                                                                 `json:"name,omitempty"`            // Network Condition name
	DeviceList      []string                                                                               `json:"deviceList,omitempty"`      // <p>This field should contain Device Name. The device name must be the same as the name field in a Network Device object. Line format - Device Name</p>
	CliDnisList     []string                                                                               `json:"cliDnisList,omitempty"`     // <p>This field should contain a Caller ID (CLI), comma, and Called ID (DNIS).<br> Line format -  Caller ID (CLI), Called ID (DNIS)</p>
	IPAddrList      []string                                                                               `json:"ipAddrList,omitempty"`      // <p>This field should contain IP address or subnet.<br> IP address can be IPV4 format (n.n.n.n) or IPV6 format (n:n:n:n:n:n:n:n).<br> IP subnet can be IPV4 format (n.n.n.n/m) or IPV6 format (n:n:n:n:n:n:n:n/m).<br> Line format - IP Address or subnet</p>
	MacAddrList     []string                                                                               `json:"macAddrList,omitempty"`     // <p>This field should contain Endstation MAC address, comma, and Destination MAC addresses.<br> Each Max address must include twelve hexadecimal digits using formats nn:nn:nn:nn:nn:nn or nn-nn-nn-nn-nn-nn or nnnn.nnnn.nnnn or nnnnnnnnnnnn.<br> Line format - Endstation MAC,Destination MAC </p>
	DeviceGroupList []string                                                                               `json:"deviceGroupList,omitempty"` // <p>This field should contain a tuple with NDG Root, comma, and an NDG (that it under the root).<br> Line format - NDG Root Name, NDG, Port</p>
}

type ResponseNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByID struct {
	Response *ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDResponse `json:"response,omitempty"` // Unique network conditions to restrict access to the network
	Version  string                                                                              `json:"version,omitempty"`  //
}

type ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDResponse struct {
	ConditionType   string                                                                                  `json:"conditionType,omitempty"`   // This field determines the content of the conditions field
	Description     string                                                                                  `json:"description,omitempty"`     //
	ID              string                                                                                  `json:"id,omitempty"`              //
	Link            *ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDResponseLink `json:"link,omitempty"`            //
	Name            string                                                                                  `json:"name,omitempty"`            // Network Condition name
	DeviceList      []string                                                                                `json:"deviceList,omitempty"`      // <p>This field should contain Device Name. The device name must be the same as the name field in a Network Device object. Line format - Device Name</p>
	CliDnisList     []string                                                                                `json:"cliDnisList,omitempty"`     // <p>This field should contain a Caller ID (CLI), comma, and Called ID (DNIS).<br> Line format -  Caller ID (CLI), Called ID (DNIS)</p>
	IPAddrList      []string                                                                                `json:"ipAddrList,omitempty"`      // <p>This field should contain IP address or subnet.<br> IP address can be IPV4 format (n.n.n.n) or IPV6 format (n:n:n:n:n:n:n:n).<br> IP subnet can be IPV4 format (n.n.n.n/m) or IPV6 format (n:n:n:n:n:n:n:n/m).<br> Line format - IP Address or subnet</p>
	MacAddrList     []string                                                                                `json:"macAddrList,omitempty"`     // <p>This field should contain Endstation MAC address, comma, and Destination MAC addresses.<br> Each Max address must include twelve hexadecimal digits using formats nn:nn:nn:nn:nn:nn or nn-nn-nn-nn-nn-nn or nnnn.nnnn.nnnn or nnnnnnnnnnnn.<br> Line format - Endstation MAC,Destination MAC </p>
	DeviceGroupList []string                                                                                `json:"deviceGroupList,omitempty"` // <p>This field should contain a tuple with NDG Root, comma, and an NDG (that it under the root).<br> Line format - NDG Root Name, NDG, Port</p>
}

type ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByID struct {
	Response *ResponseNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDResponse `json:"response,omitempty"` // Unique network conditions to restrict access to the network
	Version  string                                                                                 `json:"version,omitempty"`  //
}

type ResponseNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDResponse struct {
	ConditionType   string                                                                                     `json:"conditionType,omitempty"`   // This field determines the content of the conditions field
	Description     string                                                                                     `json:"description,omitempty"`     //
	ID              string                                                                                     `json:"id,omitempty"`              //
	Link            *ResponseNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDResponseLink `json:"link,omitempty"`            //
	Name            string                                                                                     `json:"name,omitempty"`            // Network Condition name
	DeviceList      []string                                                                                   `json:"deviceList,omitempty"`      // <p>This field should contain Device Name. The device name must be the same as the name field in a Network Device object. Line format - Device Name</p>
	CliDnisList     []string                                                                                   `json:"cliDnisList,omitempty"`     // <p>This field should contain a Caller ID (CLI), comma, and Called ID (DNIS).<br> Line format -  Caller ID (CLI), Called ID (DNIS)</p>
	IPAddrList      []string                                                                                   `json:"ipAddrList,omitempty"`      // <p>This field should contain IP address or subnet.<br> IP address can be IPV4 format (n.n.n.n) or IPV6 format (n:n:n:n:n:n:n:n).<br> IP subnet can be IPV4 format (n.n.n.n/m) or IPV6 format (n:n:n:n:n:n:n:n/m).<br> Line format - IP Address or subnet</p>
	MacAddrList     []string                                                                                   `json:"macAddrList,omitempty"`     // <p>This field should contain Endstation MAC address, comma, and Destination MAC addresses.<br> Each Max address must include twelve hexadecimal digits using formats nn:nn:nn:nn:nn:nn or nn-nn-nn-nn-nn-nn or nnnn.nnnn.nnnn or nnnnnnnnnnnn.<br> Line format - Endstation MAC,Destination MAC </p>
	DeviceGroupList []string                                                                                   `json:"deviceGroupList,omitempty"` // <p>This field should contain a tuple with NDG Root, comma, and an NDG (that it under the root).<br> Line format - NDG Root Name, NDG, Port</p>
}

type ResponseNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessNetworkConditionsDeleteNetworkAccessNetworkConditionByID struct {
	ID string `json:"id,omitempty"` //
}

type RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkCondition struct {
	ConditionType   string                                                                        `json:"conditionType,omitempty"`   // This field determines the content of the conditions field
	Description     string                                                                        `json:"description,omitempty"`     //
	ID              string                                                                        `json:"id,omitempty"`              //
	Link            *RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionLink `json:"link,omitempty"`            //
	Name            string                                                                        `json:"name,omitempty"`            // Network Condition name
	DeviceList      []string                                                                      `json:"deviceList,omitempty"`      // <p>This field should contain Device Name. The device name must be the same as the name field in a Network Device object. Line format - Device Name</p>
	CliDnisList     []string                                                                      `json:"cliDnisList,omitempty"`     // <p>This field should contain a Caller ID (CLI), comma, and Called ID (DNIS).<br> Line format -  Caller ID (CLI), Called ID (DNIS)</p>
	IPAddrList      []string                                                                      `json:"ipAddrList,omitempty"`      // <p>This field should contain IP address or subnet.<br> IP address can be IPV4 format (n.n.n.n) or IPV6 format (n:n:n:n:n:n:n:n).<br> IP subnet can be IPV4 format (n.n.n.n/m) or IPV6 format (n:n:n:n:n:n:n:n/m).<br> Line format - IP Address or subnet</p>
	MacAddrList     []string                                                                      `json:"macAddrList,omitempty"`     // <p>This field should contain Endstation MAC address, comma, and Destination MAC addresses.<br> Each Max address must include twelve hexadecimal digits using formats nn:nn:nn:nn:nn:nn or nn-nn-nn-nn-nn-nn or nnnn.nnnn.nnnn or nnnnnnnnnnnn.<br> Line format - Endstation MAC,Destination MAC </p>
	DeviceGroupList []string                                                                      `json:"deviceGroupList,omitempty"` // <p>This field should contain a tuple with NDG Root, comma, and an NDG (that it under the root).<br> Line format - NDG Root Name, NDG, Port</p>
}

type RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByID struct {
	ConditionType   string                                                                            `json:"conditionType,omitempty"`   // This field determines the content of the conditions field
	Description     string                                                                            `json:"description,omitempty"`     //
	ID              string                                                                            `json:"id,omitempty"`              //
	Link            *RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDLink `json:"link,omitempty"`            //
	Name            string                                                                            `json:"name,omitempty"`            // Network Condition name
	DeviceList      []string                                                                          `json:"deviceList,omitempty"`      // <p>This field should contain Device Name. The device name must be the same as the name field in a Network Device object. Line format - Device Name</p>
	CliDnisList     []string                                                                          `json:"cliDnisList,omitempty"`     // <p>This field should contain a Caller ID (CLI), comma, and Called ID (DNIS).<br> Line format -  Caller ID (CLI), Called ID (DNIS)</p>
	IPAddrList      []string                                                                          `json:"ipAddrList,omitempty"`      // <p>This field should contain IP address or subnet.<br> IP address can be IPV4 format (n.n.n.n) or IPV6 format (n:n:n:n:n:n:n:n).<br> IP subnet can be IPV4 format (n.n.n.n/m) or IPV6 format (n:n:n:n:n:n:n:n/m).<br> Line format - IP Address or subnet</p>
	MacAddrList     []string                                                                          `json:"macAddrList,omitempty"`     // <p>This field should contain Endstation MAC address, comma, and Destination MAC addresses.<br> Each Max address must include twelve hexadecimal digits using formats nn:nn:nn:nn:nn:nn or nn-nn-nn-nn-nn-nn or nnnn.nnnn.nnnn or nnnnnnnnnnnn.<br> Line format - Endstation MAC,Destination MAC </p>
	DeviceGroupList []string                                                                          `json:"deviceGroupList,omitempty"` // <p>This field should contain a tuple with NDG Root, comma, and an NDG (that it under the root).<br> Line format - NDG Root Name, NDG, Port</p>
}

type RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

//GetNetworkAccessNetworkConditions Network Access - Returns a list of network conditions.
/* Network Access Returns a list of network conditions.

 */
func (s *NetworkAccessNetworkConditionsService) GetNetworkAccessNetworkConditions() (*ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/network-condition"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessNetworkConditions")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions)
	return result, response, err

}

//GetNetworkAccessNetworkConditionByID Network Access - Returns a network condition.
/* Network Access Returns a network condition.

@param id id path parameter. Condition id
*/
func (s *NetworkAccessNetworkConditionsService) GetNetworkAccessNetworkConditionByID(id string) (*ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/network-condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessNetworkConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByID)
	return result, response, err

}

//CreateNetworkAccessNetworkCondition Network Access - Creates network condition.
/* Network Access Creates network condition.

 */
func (s *NetworkAccessNetworkConditionsService) CreateNetworkAccessNetworkCondition(requestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkCondition *RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkCondition) (*ResponseNetworkAccessNetworkConditionsCreateNetworkAccessNetworkCondition, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/network-condition"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkCondition).
		SetResult(&ResponseNetworkAccessNetworkConditionsCreateNetworkAccessNetworkCondition{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessNetworkConditionsCreateNetworkAccessNetworkCondition{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkAccessNetworkCondition")
	}

	result := response.Result().(*ResponseNetworkAccessNetworkConditionsCreateNetworkAccessNetworkCondition)
	return result, response, err

}

//UpdateNetworkAccessNetworkConditionByID Network Access - Update network condition.
/* Network Access Update network condition.

@param id id path parameter. Condition id
*/
func (s *NetworkAccessNetworkConditionsService) UpdateNetworkAccessNetworkConditionByID(id string, requestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionById *RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByID) (*ResponseNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/network-condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionById).
		SetResult(&ResponseNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByID{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkAccessNetworkConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByID)
	return result, response, err

}

//DeleteNetworkAccessNetworkConditionByID Network Access - Delete network condition.
/* Network Access Delete network condition.

@param id id path parameter. Condition id
*/
func (s *NetworkAccessNetworkConditionsService) DeleteNetworkAccessNetworkConditionByID(id string) (*ResponseNetworkAccessNetworkConditionsDeleteNetworkAccessNetworkConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/network-condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessNetworkConditionsDeleteNetworkAccessNetworkConditionByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteNetworkAccessNetworkConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessNetworkConditionsDeleteNetworkAccessNetworkConditionByID)
	return result, response, err

}
