package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SubscriberService service

type GetAllSubscribersQueryParams struct {
	Page       int    `url:"page,omitempty"`       //Page number
	Size       int    `url:"size,omitempty"`       //Number of objects returned per page
	Filter     string `url:"filter,omitempty"`     //<div> <style type="text/css" scoped> .apiServiceTable td, .apiServiceTable th { padding: 5px 10px !important; text-align: left; } </style> <span> <b>Simple filtering</b> should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the <i>"filterType=or"</i> query string parameter. Each resource Data model description should specify if an attribute is a filtered field. </span> <br /> <table class="apiServiceTable"> <thead> <tr> <th>OPERATOR</th> <th>DESCRIPTION</th> </tr> </thead> <tbody> <tr> <td>EQ</td> <td>Equals</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> </tr> <tr> <td>GT</td> <td>Greater Than</td> </tr> <tr> <td>LT</td> <td>Less Then</td> </tr> <tr> <td>STARTSW</td> <td>Starts With</td> </tr> <tr> <td>NSTARTSW</td> <td>Not Starts With</td> </tr> <tr> <td>ENDSW</td> <td>Ends With</td> </tr> <tr> <td>NENDSW</td> <td>Not Ends With</td> </tr> <tr> <td>CONTAINS</td> <td>Contains</td> </tr> <tr> <td>NCONTAINS</td> <td>Not Contains</td> </tr> </tbody> </table> </div>
	FilterType string `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
	Sort       string `url:"sort,omitempty"`       //sort type - asc or desc
	SortBy     string `url:"sortBy,omitempty"`     //sort column by which objects needs to be sorted
}

type ResponseSubscriberGetAllSubscribers struct {
	Response     *[]ResponseSubscriberGetAllSubscribersResponse   `json:"response,omitempty"`     //
	NextPage     *ResponseSubscriberGetAllSubscribersNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseSubscriberGetAllSubscribersPreviousPage `json:"previousPage,omitempty"` //
	Version      string                                           `json:"version,omitempty"`      //
}

type ResponseSubscriberGetAllSubscribersResponse struct {
	Enabled        *bool                                            `json:"enabled,omitempty"`        // Subscriber is enabled or not
	FriendlyName   string                                           `json:"friendlyName,omitempty"`   // Friendly name for the subscriber
	IDentityGroups string                                           `json:"identityGroups,omitempty"` // Identity Groups. If you add more than one identity group, they need to be comma separated.
	Imeis          string                                           `json:"imeis,omitempty"`          // IMEI to be attached to the subscriber
	Opc            string                                           `json:"opc,omitempty"`            // OPC
	Ki             string                                           `json:"ki,omitempty"`             // KI
	Imsi           string                                           `json:"imsi,omitempty"`           // IMSI for subscriber
	CreateTime     string                                           `json:"createTime,omitempty"`     //
	UpdateTime     string                                           `json:"updateTime,omitempty"`     //
	ID             string                                           `json:"id,omitempty"`             //
	Link           *ResponseSubscriberGetAllSubscribersResponseLink `json:"link,omitempty"`           //
}

type ResponseSubscriberGetAllSubscribersResponseLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSubscriberGetAllSubscribersNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSubscriberGetAllSubscribersPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSubscriberCreateSubscriber struct {
	Response *ResponseSubscriberCreateSubscriberResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  //
}

type ResponseSubscriberCreateSubscriberResponse struct {
	Enabled        *bool                                           `json:"enabled,omitempty"`        // Subscriber is enabled or not
	FriendlyName   string                                          `json:"friendlyName,omitempty"`   // Friendly name for the subscriber
	IDentityGroups string                                          `json:"identityGroups,omitempty"` // Identity Groups. If you add more than one identity group, they need to be comma separated.
	Imeis          string                                          `json:"imeis,omitempty"`          // IMEI to be attached to the subscriber
	Opc            string                                          `json:"opc,omitempty"`            // OPC
	Ki             string                                          `json:"ki,omitempty"`             // KI
	Imsi           string                                          `json:"imsi,omitempty"`           // IMSI for subscriber
	CreateTime     string                                          `json:"createTime,omitempty"`     //
	UpdateTime     string                                          `json:"updateTime,omitempty"`     //
	ID             string                                          `json:"id,omitempty"`             //
	Link           *ResponseSubscriberCreateSubscriberResponseLink `json:"link,omitempty"`           //
}

type ResponseSubscriberCreateSubscriberResponseLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSubscriberGetSubscriberByID struct {
	Response *ResponseSubscriberGetSubscriberByIDResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  //
}

type ResponseSubscriberGetSubscriberByIDResponse struct {
	Enabled        *bool                                            `json:"enabled,omitempty"`        // Subscriber is enabled or not
	FriendlyName   string                                           `json:"friendlyName,omitempty"`   // Friendly name for the subscriber
	IDentityGroups string                                           `json:"identityGroups,omitempty"` // Identity Groups. If you add more than one identity group, they need to be comma separated.
	Imeis          string                                           `json:"imeis,omitempty"`          // IMEI to be attached to the subscriber
	Opc            string                                           `json:"opc,omitempty"`            // OPC
	Ki             string                                           `json:"ki,omitempty"`             // KI
	Imsi           string                                           `json:"imsi,omitempty"`           // IMSI for subscriber
	CreateTime     string                                           `json:"createTime,omitempty"`     //
	UpdateTime     string                                           `json:"updateTime,omitempty"`     //
	ID             string                                           `json:"id,omitempty"`             //
	Link           *ResponseSubscriberGetSubscriberByIDResponseLink `json:"link,omitempty"`           //
}

type ResponseSubscriberGetSubscriberByIDResponseLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSubscriberUpdateSubscriber struct {
	Response *ResponseSubscriberUpdateSubscriberResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  //
}

type ResponseSubscriberUpdateSubscriberResponse struct {
	Enabled        *bool                                           `json:"enabled,omitempty"`        // Subscriber is enabled or not
	FriendlyName   string                                          `json:"friendlyName,omitempty"`   // Friendly name for the subscriber
	IDentityGroups string                                          `json:"identityGroups,omitempty"` // Identity Groups. If you add more than one identity group, they need to be comma separated.
	Imeis          string                                          `json:"imeis,omitempty"`          // IMEI to be attached to the subscriber
	Opc            string                                          `json:"opc,omitempty"`            // OPC
	Ki             string                                          `json:"ki,omitempty"`             // KI
	Imsi           string                                          `json:"imsi,omitempty"`           // IMSI for subscriber
	CreateTime     string                                          `json:"createTime,omitempty"`     //
	UpdateTime     string                                          `json:"updateTime,omitempty"`     //
	ID             string                                          `json:"id,omitempty"`             //
	Link           *ResponseSubscriberUpdateSubscriberResponseLink `json:"link,omitempty"`           //
}

type ResponseSubscriberUpdateSubscriberResponseLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSubscriberDeleteSubscriber struct {
	Message string `json:"message,omitempty"` //
}

type ResponseSubscriberGetSubscriberByIMSI struct {
	Response *ResponseSubscriberGetSubscriberByIMSIResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  //
}

type ResponseSubscriberGetSubscriberByIMSIResponse struct {
	Enabled        *bool                                              `json:"enabled,omitempty"`        // Subscriber is enabled or not
	FriendlyName   string                                             `json:"friendlyName,omitempty"`   // Friendly name for the subscriber
	IDentityGroups string                                             `json:"identityGroups,omitempty"` // Identity Groups. If you add more than one identity group, they need to be comma separated.
	Imeis          string                                             `json:"imeis,omitempty"`          // IMEI to be attached to the subscriber
	Opc            string                                             `json:"opc,omitempty"`            // OPC
	Ki             string                                             `json:"ki,omitempty"`             // KI
	Imsi           string                                             `json:"imsi,omitempty"`           // IMSI for subscriber
	CreateTime     string                                             `json:"createTime,omitempty"`     //
	UpdateTime     string                                             `json:"updateTime,omitempty"`     //
	ID             string                                             `json:"id,omitempty"`             //
	Link           *ResponseSubscriberGetSubscriberByIMSIResponseLink `json:"link,omitempty"`           //
}

type ResponseSubscriberGetSubscriberByIMSIResponseLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseSubscriberBulkSubscriberOperation struct {
	ID string `json:"id,omitempty"` //
}

type RequestSubscriberCreateSubscriber struct {
	Enabled        *bool  `json:"enabled,omitempty"`        // Subscriber is enabled or not
	FriendlyName   string `json:"friendlyName,omitempty"`   // Friendly name for the subscriber
	IDentityGroups string `json:"identityGroups,omitempty"` // Identity Groups. If you add more than one identity group, they need to be comma separated.
	Imeis          string `json:"imeis,omitempty"`          // IMEI to be attached to the subscriber
	Opc            string `json:"opc,omitempty"`            // OPC
	Ki             string `json:"ki,omitempty"`             // KI
	Imsi           string `json:"imsi,omitempty"`           // IMSI for subscriber
}

type RequestSubscriberUpdateSubscriber struct {
	Enabled        *bool  `json:"enabled,omitempty"`        // Subscriber is enabled or not
	FriendlyName   string `json:"friendlyName,omitempty"`   // Friendly name for the subscriber
	IDentityGroups string `json:"identityGroups,omitempty"` // Identity Groups. If you add more than one identity group, they need to be comma separated.
	Imeis          string `json:"imeis,omitempty"`          // IMEI to be attached to the subscriber
	Opc            string `json:"opc,omitempty"`            // OPC
	Ki             string `json:"ki,omitempty"`             // KI
}

type RequestSubscriberBulkSubscriberOperation struct {
	Operation string                                              `json:"operation,omitempty"` //
	ItemList  *[]RequestSubscriberBulkSubscriberOperationItemList `json:"ItemList,omitempty"`  //
}

type RequestSubscriberBulkSubscriberOperationItemList struct {
	ID             string `json:"id,omitempty"`             //
	Imsi           string `json:"imsi,omitempty"`           // IMSI for Subscriber
	Enabled        *bool  `json:"enabled,omitempty"`        // Subscriber is enabled or not
	FriendlyName   string `json:"friendlyName,omitempty"`   // Friendly name for the subscriber
	IDentityGroups string `json:"identityGroups,omitempty"` // Identity Groups. If you add more than one identity group, they need to be comma separated
	Imeis          string `json:"imeis,omitempty"`          // Comma separated IMEIs to be attached to the subscriber
	Opc            string `json:"opc,omitempty"`            // OPC
	Ki             string `json:"ki,omitempty"`             // KI
}

//GetAllSubscribers Get all subscribers
/* Get all subscribers

@param getAllSubscribersQueryParams Filtering parameter
*/
func (s *SubscriberService) GetAllSubscribers(getAllSubscribersQueryParams *GetAllSubscribersQueryParams) (*ResponseSubscriberGetAllSubscribers, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/fiveg/subscriber"

	queryString, _ := query.Values(getAllSubscribersQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSubscriberGetAllSubscribers{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAllSubscribers")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSubscriberGetAllSubscribers)
	return result, response, err

}

//GetSubscriberByID Get the subscriber for a given ID
/* Get the subscriber for a given ID

@param subscriberID subscriberId path parameter. Unique ID for a subscriber object
*/
func (s *SubscriberService) GetSubscriberByID(subscriberID string) (*ResponseSubscriberGetSubscriberByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/fiveg/subscriber/{subscriberId}"
	path = strings.Replace(path, "{subscriberId}", fmt.Sprintf("%v", subscriberID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSubscriberGetSubscriberByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSubscriberById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSubscriberGetSubscriberByID)
	return result, response, err

}

//GetSubscriberByIMSI Get a subscriber by IMSI
/* Get a subscriber by IMSI

@param imsi imsi path parameter. IMSI parameter
*/
func (s *SubscriberService) GetSubscriberByIMSI(imsi string) (*ResponseSubscriberGetSubscriberByIMSI, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/fiveg/subscriber/imsi/{imsi}"
	path = strings.Replace(path, "{imsi}", fmt.Sprintf("%v", imsi), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSubscriberGetSubscriberByIMSI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSubscriberByIMSI")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSubscriberGetSubscriberByIMSI)
	return result, response, err

}

//CreateSubscriber Create subscriber holding the IMEI
/* Create subscriber holding the IMEI

 */
func (s *SubscriberService) CreateSubscriber(requestSubscriberCreateSubscriber *RequestSubscriberCreateSubscriber) (*ResponseSubscriberCreateSubscriber, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/fiveg/subscriber"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSubscriberCreateSubscriber).
		SetResult(&ResponseSubscriberCreateSubscriber{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseSubscriberCreateSubscriber{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateSubscriber")
	}

	result := response.Result().(*ResponseSubscriberCreateSubscriber)
	return result, response, err

}

//BulkSubscriberOperation Create, update and delete subscribers in bulk
/* Create, update and delete subscribers in bulk

 */
func (s *SubscriberService) BulkSubscriberOperation(requestSubscriberBulkSubscriberOperation *RequestSubscriberBulkSubscriberOperation) (*ResponseSubscriberBulkSubscriberOperation, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/fiveg/subscriber/bulk"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSubscriberBulkSubscriberOperation).
		SetResult(&ResponseSubscriberBulkSubscriberOperation{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseSubscriberBulkSubscriberOperation{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BulkSubscriberOperation")
	}

	result := response.Result().(*ResponseSubscriberBulkSubscriberOperation)
	return result, response, err

}

//UpdateSubscriber Update a subscriber for a given ID and the request payload
/* Update a subscriber for a given ID and the request payload

@param subscriberID subscriberId path parameter. Unique ID for a subscriber object
*/
func (s *SubscriberService) UpdateSubscriber(subscriberID string, requestSubscriberUpdateSubscriber *RequestSubscriberUpdateSubscriber) (*ResponseSubscriberUpdateSubscriber, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/fiveg/subscriber/{subscriberId}"
	path = strings.Replace(path, "{subscriberId}", fmt.Sprintf("%v", subscriberID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSubscriberUpdateSubscriber).
		SetResult(&ResponseSubscriberUpdateSubscriber{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseSubscriberUpdateSubscriber{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSubscriber")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSubscriberUpdateSubscriber)
	return result, response, err

}

//DeleteSubscriber Delete the subscriber for a given ID
/* Delete the subscriber for a given ID

@param subscriberID subscriberId path parameter. Unique ID for a subscriber object
*/
func (s *SubscriberService) DeleteSubscriber(subscriberID string) (*ResponseSubscriberDeleteSubscriber, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/fiveg/subscriber/{subscriberId}"
	path = strings.Replace(path, "{subscriberId}", fmt.Sprintf("%v", subscriberID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSubscriberDeleteSubscriber{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteSubscriber")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSubscriberDeleteSubscriber)
	return result, response, err

}
