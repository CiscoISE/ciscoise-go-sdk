package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SgtRangeReservationService service

type GetSgtReservedRangesQueryParams struct {
	Page int `url:"page,omitempty"` //Page number
	Size int `url:"size,omitempty"` //Number of objects returned per page
}

type ResponseSgtRangeReservationGetSgtReservedRanges struct {
	Response *[]ResponseSgtRangeReservationGetSgtReservedRangesResponse `json:"response,omitempty"` //
	Version  string                                                     `json:"version,omitempty"`  //
}

type ResponseSgtRangeReservationGetSgtReservedRangesResponse struct {
	ClientID   string `json:"clientId,omitempty"`   // Unique ID of the given client
	ClientName string `json:"clientName,omitempty"` // Name of the given client
	EndIndex   *int   `json:"endIndex,omitempty"`   // End index of the reserved range
	StartIndex *int   `json:"startIndex,omitempty"` // Start index of the reserved range
}

type ResponseSgtRangeReservationReserveSgtRange struct {
	Response *ResponseSgtRangeReservationReserveSgtRangeResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  //
}

type ResponseSgtRangeReservationReserveSgtRangeResponse struct {
	ClientID   string `json:"clientId,omitempty"`   // Unique ID of the given client
	ClientName string `json:"clientName,omitempty"` // Name of the given client
	EndIndex   *int   `json:"endIndex,omitempty"`   // End index of the reserved range
	StartIndex *int   `json:"startIndex,omitempty"` // Start index of the reserved range
}

type ResponseSgtRangeReservationGetSgtReservedRange struct {
	Response *ResponseSgtRangeReservationGetSgtReservedRangeResponse `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  //
}

type ResponseSgtRangeReservationGetSgtReservedRangeResponse struct {
	ClientID   string `json:"clientId,omitempty"`   // Unique ID of the given client
	ClientName string `json:"clientName,omitempty"` // Name of the given client
	EndIndex   *int   `json:"endIndex,omitempty"`   // End index of the reserved range
	StartIndex *int   `json:"startIndex,omitempty"` // Start index of the reserved range
}

type ResponseSgtRangeReservationUpdateReservedRange struct {
	Success *ResponseSgtRangeReservationUpdateReservedRangeSuccess `json:"success,omitempty"` //
	Version string                                                 `json:"version,omitempty"` //
}

type ResponseSgtRangeReservationUpdateReservedRangeSuccess struct {
	Message string `json:"message,omitempty"` //
}

type ResponseSgtRangeReservationDeleteSgtReserveRange struct {
	Success *ResponseSgtRangeReservationDeleteSgtReserveRangeSuccess `json:"success,omitempty"` //
	Version string                                                   `json:"version,omitempty"` //
}

type ResponseSgtRangeReservationDeleteSgtReserveRangeSuccess struct {
	Message string `json:"message,omitempty"` //
}

type RequestSgtRangeReservationReserveSgtRange struct {
	ClientName   string `json:"clientName,omitempty"`   // Name of the given client
	NumberOfTags *int   `json:"numberOfTags,omitempty"` // Nummber of tags required to be reserved in ISE.
}

type RequestSgtRangeReservationUpdateReservedRange struct {
	ClientID   string `json:"clientId,omitempty"`   // Unique ID of the given client
	EndIndex   *int   `json:"endIndex,omitempty"`   // End index of the reserved range
	StartIndex *int   `json:"startIndex,omitempty"` // Start index of the reserved range
}

//GetSgtReservedRanges Get all the reserved Security Group tag ranges in ISE.
/* Get all the reserved Security Group tag ranges in ISE.

@param getSgtReservedRangesQueryParams Filtering parameter
*/
func (s *SgtRangeReservationService) GetSgtReservedRanges(getSgtReservedRangesQueryParams *GetSgtReservedRangesQueryParams) (*ResponseSgtRangeReservationGetSgtReservedRanges, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/sgt/reservation"

	queryString, _ := query.Values(getSgtReservedRangesQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSgtRangeReservationGetSgtReservedRanges{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSgtReservedRanges")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSgtRangeReservationGetSgtReservedRanges)
	return result, response, err

}

//GetSgtReservedRange Get the reserved range for the specific Client.
/* Get the reserved range of SGT for the specific client which is passed in the URL.

@param clientID clientID path parameter. Unique name for a Client
*/
func (s *SgtRangeReservationService) GetSgtReservedRange(clientID string) (*ResponseSgtRangeReservationGetSgtReservedRange, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/sgt/reservation/{clientID}"
	path = strings.Replace(path, "{clientID}", fmt.Sprintf("%v", clientID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSgtRangeReservationGetSgtReservedRange{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSgtReservedRange")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSgtRangeReservationGetSgtReservedRange)
	return result, response, err

}

//ReserveSgtRange Reserve range of SGT in ISE for the given Client.
/* Reserve given number of SGTs in a continuous range for the given Client.

 */
func (s *SgtRangeReservationService) ReserveSgtRange(requestSgtRangeReservationReserveSgtRange *RequestSgtRangeReservationReserveSgtRange) (*ResponseSgtRangeReservationReserveSgtRange, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/sgt/reservation/reserveRange"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSgtRangeReservationReserveSgtRange).
		SetResult(&ResponseSgtRangeReservationReserveSgtRange{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseSgtRangeReservationReserveSgtRange{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ReserveSgtRange")
	}

	result := response.Result().(*ResponseSgtRangeReservationReserveSgtRange)
	return result, response, err

}

//UpdateReservedRange Update the reserved ranges of a specific Client
/* Update the reserved ranges of a specific Client by giving the custom start and end index

@param clientID clientID path parameter. Unique name for a Client
*/
func (s *SgtRangeReservationService) UpdateReservedRange(clientID string, requestSgtRangeReservationUpdateReservedRange *RequestSgtRangeReservationUpdateReservedRange) (*ResponseSgtRangeReservationUpdateReservedRange, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/sgt/reservation/{clientID}"
	path = strings.Replace(path, "{clientID}", fmt.Sprintf("%v", clientID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSgtRangeReservationUpdateReservedRange).
		SetResult(&ResponseSgtRangeReservationUpdateReservedRange{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseSgtRangeReservationUpdateReservedRange{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateReservedRange")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSgtRangeReservationUpdateReservedRange)
	return result, response, err

}

//DeleteSgtReserveRange Delete the reserved range for the given Client
/* Delete the reserved range of SGT for the given Client

@param clientID clientID path parameter. Unique name for a Client
*/
func (s *SgtRangeReservationService) DeleteSgtReserveRange(clientID string) (*ResponseSgtRangeReservationDeleteSgtReserveRange, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/sgt/reservation/{clientID}"
	path = strings.Replace(path, "{clientID}", fmt.Sprintf("%v", clientID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSgtRangeReservationDeleteSgtReserveRange{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteSgtReserveRange")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseSgtRangeReservationDeleteSgtReserveRange)
	return result, response, err

}
