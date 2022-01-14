package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type NbarAppService service

type GetNbarAppsQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sort       string   `url:"sort,omitempty"`       //sort type - asc or desc
	SortBy     string   `url:"sortBy,omitempty"`     //sort column by which objects needs to be sorted
	Filter     []string `url:"filter,omitempty"`     //<div> <style type="text/css" scoped> .apiServiceTable td, .apiServiceTable th { padding: 5px 10px !important; text-align: left; } </style> <span> <b>Simple filtering</b> should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the <i>"filterType=or"</i> query string parameter. Each resource Data model description should specify if an attribute is a filtered field. </span> <br /> <table class="apiServiceTable"> <thead> <tr> <th>OPERATOR</th> <th>DESCRIPTION</th> </tr> </thead> <tbody> <tr> <td>EQ</td> <td>Equals</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> </tr> <tr> <td>GT</td> <td>Greater Than</td> </tr> <tr> <td>LT</td> <td>Less Then</td> </tr> <tr> <td>STARTSW</td> <td>Starts With</td> </tr> <tr> <td>NSTARTSW</td> <td>Not Starts With</td> </tr> <tr> <td>ENDSW</td> <td>Ends With</td> </tr> <tr> <td>NENDSW</td> <td>Not Ends With</td> </tr> <tr> <td>CONTAINS</td> <td>Contains</td> </tr> <tr> <td>NCONTAINS</td> <td>Not Contains</td> </tr> </tbody> </table> </div>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseNbarAppGetNbarApps struct {
	Response *[]ResponseNbarAppGetNbarAppsResponse `json:"response,omitempty"` //
}

type ResponseNbarAppGetNbarAppsResponse struct {
	Description       string                                                 `json:"description,omitempty"`       //
	ID                string                                                 `json:"id,omitempty"`                //
	Name              string                                                 `json:"name,omitempty"`              //
	NetworkIDentities *[]ResponseNbarAppGetNbarAppsResponseNetworkIDentities `json:"networkIdentities,omitempty"` // Array of NIs
}

type ResponseNbarAppGetNbarAppsResponseNetworkIDentities struct {
	Ports    string `json:"ports,omitempty"`    //
	Protocol string `json:"protocol,omitempty"` //
}

type ResponseNbarAppCreateNbarApp struct {
	ID string `json:"id,omitempty"` // ID of the newly created object
}

type ResponseNbarAppGetNbarAppByID struct {
	Response *[]ResponseNbarAppGetNbarAppByIDResponse `json:"response,omitempty"` //
}

type ResponseNbarAppGetNbarAppByIDResponse struct {
	Description       string                                                    `json:"description,omitempty"`       //
	ID                string                                                    `json:"id,omitempty"`                //
	Name              string                                                    `json:"name,omitempty"`              //
	NetworkIDentities *[]ResponseNbarAppGetNbarAppByIDResponseNetworkIDentities `json:"networkIdentities,omitempty"` // Array of NIs
}

type ResponseNbarAppGetNbarAppByIDResponseNetworkIDentities struct {
	Ports    string `json:"ports,omitempty"`    //
	Protocol string `json:"protocol,omitempty"` //
}

type ResponseNbarAppUpdateNbarAppByID struct {
	Code    *int   `json:"code,omitempty"`    //
	Message string `json:"message,omitempty"` //
}

type ResponseNbarAppDeleteNbarAppByID struct {
	Code    *int   `json:"code,omitempty"`    //
	Message string `json:"message,omitempty"` //
}

type RequestNbarAppCreateNbarApp struct {
	Description       string                                          `json:"description,omitempty"`       //
	ID                string                                          `json:"id,omitempty"`                //
	Name              string                                          `json:"name,omitempty"`              //
	NetworkIDentities *[]RequestNbarAppCreateNbarAppNetworkIDentities `json:"networkIdentities,omitempty"` // Array of NIs
}

type RequestNbarAppCreateNbarAppNetworkIDentities struct {
	Ports    string `json:"ports,omitempty"`    //
	Protocol string `json:"protocol,omitempty"` //
}

type RequestNbarAppUpdateNbarAppByID struct {
	Description       string                                              `json:"description,omitempty"`       //
	ID                string                                              `json:"id,omitempty"`                //
	Name              string                                              `json:"name,omitempty"`              //
	NetworkIDentities *[]RequestNbarAppUpdateNbarAppByIDNetworkIDentities `json:"networkIdentities,omitempty"` // Array of NIs
}

type RequestNbarAppUpdateNbarAppByIDNetworkIDentities struct {
	Ports    string `json:"ports,omitempty"`    //
	Protocol string `json:"protocol,omitempty"` //
}

//GetNbarApps Get all NBAR Applications
/* Get all NBAR Applications

@param getNBARAppsQueryParams Filtering parameter
*/
func (s *NbarAppService) GetNbarApps(getNBARAppsQueryParams *GetNbarAppsQueryParams) (*ResponseNbarAppGetNbarApps, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/sgacl/nbarapp"

	queryString, _ := query.Values(getNBARAppsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNbarAppGetNbarApps{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNbarApps")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNbarAppGetNbarApps)
	return result, response, err

}

//GetNbarAppByID Get NBAR Application by id
/* Get NBAR Application by id

@param id id path parameter.
*/
func (s *NbarAppService) GetNbarAppByID(id string) (*ResponseNbarAppGetNbarAppByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/sgacl/nbarapp/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNbarAppGetNbarAppByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNbarAppById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNbarAppGetNbarAppByID)
	return result, response, err

}

//CreateNbarApp Create NBAR application
/* Create NBAR application

 */
func (s *NbarAppService) CreateNbarApp(requestNbarAppCreateNBARApp *RequestNbarAppCreateNbarApp) (*ResponseNbarAppCreateNbarApp, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/sgacl/nbarapp"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNbarAppCreateNBARApp).
		SetResult(&ResponseNbarAppCreateNbarApp{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNbarAppCreateNbarApp{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNbarApp")
	}

	result := response.Result().(*ResponseNbarAppCreateNbarApp)
	return result, response, err

}

//UpdateNbarAppByID Update NBAR Application
/* Update NBAR Application

@param id id path parameter.
*/
func (s *NbarAppService) UpdateNbarAppByID(id string, requestNbarAppUpdateNBARAppById *RequestNbarAppUpdateNbarAppByID) (*ResponseNbarAppUpdateNbarAppByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/sgacl/nbarapp/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNbarAppUpdateNBARAppById).
		SetResult(&ResponseNbarAppUpdateNbarAppByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNbarAppUpdateNbarAppByID{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNbarAppById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNbarAppUpdateNbarAppByID)
	return result, response, err

}

//DeleteNbarAppByID Delete NBAR Application
/* Delete NBAR Application

@param id id path parameter.
*/
func (s *NbarAppService) DeleteNbarAppByID(id string) (*ResponseNbarAppDeleteNbarAppByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/trustsec/sgacl/nbarapp/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNbarAppDeleteNbarAppByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteNbarAppById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNbarAppDeleteNbarAppByID)
	return result, response, err

}
