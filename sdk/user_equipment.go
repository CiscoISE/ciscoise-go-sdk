package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type UserEquipmentService service

type GetUserEquipmentsQueryParams struct {
	Page       int    `url:"page,omitempty"`       //Page number
	Size       int    `url:"size,omitempty"`       //Number of objects returned per page
	Filter     string `url:"filter,omitempty"`     //<div> <style type="text/css" scoped> .apiServiceTable td, .apiServiceTable th { padding: 5px 10px !important; text-align: left; } </style> <span> <b>Simple filtering</b> should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the <i>"filterType=or"</i> query string parameter. Each resource Data model description should specify if an attribute is a filtered field. </span> <br /> <table class="apiServiceTable"> <thead> <tr> <th>OPERATOR</th> <th>DESCRIPTION</th> </tr> </thead> <tbody> <tr> <td>EQ</td> <td>Equals</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> </tr> <tr> <td>GT</td> <td>Greater Than</td> </tr> <tr> <td>LT</td> <td>Less Then</td> </tr> <tr> <td>STARTSW</td> <td>Starts With</td> </tr> <tr> <td>NSTARTSW</td> <td>Not Starts With</td> </tr> <tr> <td>ENDSW</td> <td>Ends With</td> </tr> <tr> <td>NENDSW</td> <td>Not Ends With</td> </tr> <tr> <td>CONTAINS</td> <td>Contains</td> </tr> <tr> <td>NCONTAINS</td> <td>Not Contains</td> </tr> </tbody> </table> </div>
	FilterType string `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
	Sort       string `url:"sort,omitempty"`       //sort type - asc or desc
	SortBy     string `url:"sortBy,omitempty"`     //sort column by which objects needs to be sorted
}

type ResponseUserEquipmentGetUserEquipments struct {
	Response     *[]ResponseUserEquipmentGetUserEquipmentsResponse   `json:"response,omitempty"`     //
	NextPage     *ResponseUserEquipmentGetUserEquipmentsNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseUserEquipmentGetUserEquipmentsPreviousPage `json:"previousPage,omitempty"` //
	Version      string                                              `json:"version,omitempty"`      //
}

type ResponseUserEquipmentGetUserEquipmentsResponse struct {
	Description string                                              `json:"description,omitempty"` // Description for User Equipment
	DeviceGroup string                                              `json:"deviceGroup,omitempty"` // Device or Endpoint Group
	Imei        string                                              `json:"imei,omitempty"`        // IMEI for User Equipment
	CreateTime  string                                              `json:"createTime,omitempty"`  //
	UpdateTime  string                                              `json:"updateTime,omitempty"`  //
	ID          string                                              `json:"id,omitempty"`          //
	Link        *ResponseUserEquipmentGetUserEquipmentsResponseLink `json:"link,omitempty"`        //
}

type ResponseUserEquipmentGetUserEquipmentsResponseLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseUserEquipmentGetUserEquipmentsNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseUserEquipmentGetUserEquipmentsPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseUserEquipmentCreateUserEquipment struct {
	Response *ResponseUserEquipmentCreateUserEquipmentResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}

type ResponseUserEquipmentCreateUserEquipmentResponse struct {
	Description string                                                `json:"description,omitempty"` // Description for User Equipment
	DeviceGroup string                                                `json:"deviceGroup,omitempty"` // Device or Endpoint Group
	Imei        string                                                `json:"imei,omitempty"`        // IMEI for User Equipment
	CreateTime  string                                                `json:"createTime,omitempty"`  //
	UpdateTime  string                                                `json:"updateTime,omitempty"`  //
	ID          string                                                `json:"id,omitempty"`          //
	Link        *ResponseUserEquipmentCreateUserEquipmentResponseLink `json:"link,omitempty"`        //
}

type ResponseUserEquipmentCreateUserEquipmentResponseLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseUserEquipmentGetUserEquipmentByID struct {
	Response *ResponseUserEquipmentGetUserEquipmentByIDResponse `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  //
}

type ResponseUserEquipmentGetUserEquipmentByIDResponse struct {
	Description string                                                 `json:"description,omitempty"` // Description for User Equipment
	DeviceGroup string                                                 `json:"deviceGroup,omitempty"` // Device or Endpoint Group
	Imei        string                                                 `json:"imei,omitempty"`        // IMEI for User Equipment
	CreateTime  string                                                 `json:"createTime,omitempty"`  //
	UpdateTime  string                                                 `json:"updateTime,omitempty"`  //
	ID          string                                                 `json:"id,omitempty"`          //
	Link        *ResponseUserEquipmentGetUserEquipmentByIDResponseLink `json:"link,omitempty"`        //
}

type ResponseUserEquipmentGetUserEquipmentByIDResponseLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseUserEquipmentUpdateUserEquipment struct {
	Response *ResponseUserEquipmentUpdateUserEquipmentResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}

type ResponseUserEquipmentUpdateUserEquipmentResponse struct {
	Description string                                                `json:"description,omitempty"` // Description for User Equipment
	DeviceGroup string                                                `json:"deviceGroup,omitempty"` // Device or Endpoint Group
	Imei        string                                                `json:"imei,omitempty"`        // IMEI for User Equipment
	CreateTime  string                                                `json:"createTime,omitempty"`  //
	UpdateTime  string                                                `json:"updateTime,omitempty"`  //
	ID          string                                                `json:"id,omitempty"`          //
	Link        *ResponseUserEquipmentUpdateUserEquipmentResponseLink `json:"link,omitempty"`        //
}

type ResponseUserEquipmentUpdateUserEquipmentResponseLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseUserEquipmentDeleteUserEquipment struct {
	Message string `json:"message,omitempty"` //
}

type ResponseUserEquipmentGetUserEquipmentsBySubscriberID struct {
	Response *[]ResponseUserEquipmentGetUserEquipmentsBySubscriberIDResponse `json:"response,omitempty"` //
	Version  string                                                          `json:"version,omitempty"`  //
}

type ResponseUserEquipmentGetUserEquipmentsBySubscriberIDResponse struct {
	Description string                                                            `json:"description,omitempty"` // Description for User Equipment
	DeviceGroup string                                                            `json:"deviceGroup,omitempty"` // Device or Endpoint Group
	Imei        string                                                            `json:"imei,omitempty"`        // IMEI for User Equipment
	CreateTime  string                                                            `json:"createTime,omitempty"`  //
	UpdateTime  string                                                            `json:"updateTime,omitempty"`  //
	ID          string                                                            `json:"id,omitempty"`          //
	Link        *ResponseUserEquipmentGetUserEquipmentsBySubscriberIDResponseLink `json:"link,omitempty"`        //
}

type ResponseUserEquipmentGetUserEquipmentsBySubscriberIDResponseLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseUserEquipmentGetUserEquipmentByIMEI struct {
	Response *ResponseUserEquipmentGetUserEquipmentByIMEIResponse `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  //
}

type ResponseUserEquipmentGetUserEquipmentByIMEIResponse struct {
	Description string                                                   `json:"description,omitempty"` // Description for User Equipment
	DeviceGroup string                                                   `json:"deviceGroup,omitempty"` // Device or Endpoint Group
	Imei        string                                                   `json:"imei,omitempty"`        // IMEI for User Equipment
	CreateTime  string                                                   `json:"createTime,omitempty"`  //
	UpdateTime  string                                                   `json:"updateTime,omitempty"`  //
	ID          string                                                   `json:"id,omitempty"`          //
	Link        *ResponseUserEquipmentGetUserEquipmentByIMEIResponseLink `json:"link,omitempty"`        //
}

type ResponseUserEquipmentGetUserEquipmentByIMEIResponseLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseUserEquipmentCreateUserEquipmentsFromCSV struct {
	ID string `json:"id,omitempty"` //
}

type ResponseUserEquipmentBulkUserEquipmentOperation struct {
	ID string `json:"id,omitempty"` //
}

type RequestUserEquipmentCreateUserEquipment struct {
	Description string `json:"description,omitempty"` // Description for User Equipment
	DeviceGroup string `json:"deviceGroup,omitempty"` // Device or Endpoint Group
	Imei        string `json:"imei,omitempty"`        // IMEI for User Equipment
}

type RequestUserEquipmentUpdateUserEquipment struct {
	Description string `json:"description,omitempty"` // Description for User Equipment
	DeviceGroup string `json:"deviceGroup,omitempty"` // Device or Endpoint Group
}

type RequestUserEquipmentBulkUserEquipmentOperation struct {
	Operation string                                                    `json:"operation,omitempty"` //
	ItemList  *[]RequestUserEquipmentBulkUserEquipmentOperationItemList `json:"ItemList,omitempty"`  //
}

type RequestUserEquipmentBulkUserEquipmentOperationItemList struct {
	Imei        string `json:"imei,omitempty"`        //
	Description string `json:"description,omitempty"` //
	ID          string `json:"id,omitempty"`          //
	DeviceGroup string `json:"deviceGroup,omitempty"` // Device or Endpoint Group
}

//GetUserEquipments Get user equipments
/* Get user equipments

@param getUserEquipmentsQueryParams Filtering parameter
*/
func (s *UserEquipmentService) GetUserEquipments(getUserEquipmentsQueryParams *GetUserEquipmentsQueryParams) (*ResponseUserEquipmentGetUserEquipments, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/fiveg/user-equipment"

	queryString, _ := query.Values(getUserEquipmentsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseUserEquipmentGetUserEquipments{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetUserEquipments")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseUserEquipmentGetUserEquipments)
	return result, response, err

}

//GetUserEquipmentByID Get the user equipment for a given ID
/* Get the user equipment for a given ID

@param userEquipmentID userEquipmentId path parameter. Unique ID for a user equipment object
*/
func (s *UserEquipmentService) GetUserEquipmentByID(userEquipmentID string) (*ResponseUserEquipmentGetUserEquipmentByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/fiveg/user-equipment/{userEquipmentId}"
	path = strings.Replace(path, "{userEquipmentId}", fmt.Sprintf("%v", userEquipmentID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserEquipmentGetUserEquipmentByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetUserEquipmentById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseUserEquipmentGetUserEquipmentByID)
	return result, response, err

}

//GetUserEquipmentsBySubscriberID Get user equipments associated with a subscriber GUID
/* Get user equipments associated with a subscriber GUID

@param subscriberID subscriberId path parameter. Unique ID for a subscriber object
*/
func (s *UserEquipmentService) GetUserEquipmentsBySubscriberID(subscriberID string) (*ResponseUserEquipmentGetUserEquipmentsBySubscriberID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/fiveg/user-equipment/subscriber/{subscriberId}"
	path = strings.Replace(path, "{subscriberId}", fmt.Sprintf("%v", subscriberID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserEquipmentGetUserEquipmentsBySubscriberID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetUserEquipmentsBySubscriberId")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseUserEquipmentGetUserEquipmentsBySubscriberID)
	return result, response, err

}

//GetUserEquipmentByIMEI Get a user equipment based on the IMEI
/* Get a user equipment based on the IMEI

@param imei imei path parameter. IMEI for the user equipment object
*/
func (s *UserEquipmentService) GetUserEquipmentByIMEI(imei string) (*ResponseUserEquipmentGetUserEquipmentByIMEI, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/fiveg/user-equipment/imei/{imei}"
	path = strings.Replace(path, "{imei}", fmt.Sprintf("%v", imei), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserEquipmentGetUserEquipmentByIMEI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetUserEquipmentByIMEI")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseUserEquipmentGetUserEquipmentByIMEI)
	return result, response, err

}

//CreateUserEquipment Create a user equipment
/* Create a user equipment

 */
func (s *UserEquipmentService) CreateUserEquipment(requestUserEquipmentCreateUserEquipment *RequestUserEquipmentCreateUserEquipment) (*ResponseUserEquipmentCreateUserEquipment, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/fiveg/user-equipment"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestUserEquipmentCreateUserEquipment).
		SetResult(&ResponseUserEquipmentCreateUserEquipment{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseUserEquipmentCreateUserEquipment{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateUserEquipment")
	}

	result := response.Result().(*ResponseUserEquipmentCreateUserEquipment)
	return result, response, err

}

//CreateUserEquipmentsFromCSV Create user equipments from a CSV file
/* Create user equipments from a CSV file

 */
func (s *UserEquipmentService) CreateUserEquipmentsFromCSV() (*ResponseUserEquipmentCreateUserEquipmentsFromCSV, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/fiveg/user-equipment/csv"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserEquipmentCreateUserEquipmentsFromCSV{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseUserEquipmentCreateUserEquipmentsFromCSV{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateUserEquipmentsFromCSV")
	}

	result := response.Result().(*ResponseUserEquipmentCreateUserEquipmentsFromCSV)
	return result, response, err

}

//BulkUserEquipmentOperation Create, update and delete multiple user equipments
/* Create, update and delete multiple user equipments

 */
func (s *UserEquipmentService) BulkUserEquipmentOperation(requestUserEquipmentBulkUserEquipmentOperation *RequestUserEquipmentBulkUserEquipmentOperation) (*ResponseUserEquipmentBulkUserEquipmentOperation, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/fiveg/user-equipment/bulk"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestUserEquipmentBulkUserEquipmentOperation).
		SetResult(&ResponseUserEquipmentBulkUserEquipmentOperation{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseUserEquipmentBulkUserEquipmentOperation{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BulkUserEquipmentOperation")
	}

	result := response.Result().(*ResponseUserEquipmentBulkUserEquipmentOperation)
	return result, response, err

}

//UpdateUserEquipment Update the user equipment for a given ID and request payload
/* Update the user equipment for a given ID and request payload

@param userEquipmentID userEquipmentId path parameter. Unique ID for a user equipment object
*/
func (s *UserEquipmentService) UpdateUserEquipment(userEquipmentID string, requestUserEquipmentUpdateUserEquipment *RequestUserEquipmentUpdateUserEquipment) (*ResponseUserEquipmentUpdateUserEquipment, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/fiveg/user-equipment/{userEquipmentId}"
	path = strings.Replace(path, "{userEquipmentId}", fmt.Sprintf("%v", userEquipmentID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestUserEquipmentUpdateUserEquipment).
		SetResult(&ResponseUserEquipmentUpdateUserEquipment{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseUserEquipmentUpdateUserEquipment{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateUserEquipment")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseUserEquipmentUpdateUserEquipment)
	return result, response, err

}

//DeleteUserEquipment Delete the user equipment for a given ID
/* Delete the user equipment for a given ID

@param userEquipmentID userEquipmentId path parameter. Unique ID for a user equipment object
*/
func (s *UserEquipmentService) DeleteUserEquipment(userEquipmentID string) (*ResponseUserEquipmentDeleteUserEquipment, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/fiveg/user-equipment/{userEquipmentId}"
	path = strings.Replace(path, "{userEquipmentId}", fmt.Sprintf("%v", userEquipmentID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserEquipmentDeleteUserEquipment{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteUserEquipment")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseUserEquipmentDeleteUserEquipment)
	return result, response, err

}
