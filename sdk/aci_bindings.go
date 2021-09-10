package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type AciBindingsService service

type GetAciBindingsQueryParams struct {
	Page        int      `url:"page,omitempty"`        //Page number
	Size        int      `url:"size,omitempty"`        //Number of objects returned per page
	Sort        string   `url:"sort,omitempty"`        //sort type - asc or desc
	SortBy      string   `url:"sortBy,omitempty"`      //sort column by which objects needs to be sorted
	FilterBy    []string `url:"filterBy,omitempty"`    //
	FilterValue []string `url:"filterValue,omitempty"` //
}

type ResponseAciBindingsGetAciBindings struct {
	AciBindings *ResponseAciBindingsGetAciBindingsAciBindings `json:"ACIBindings,omitempty"` //
}

type ResponseAciBindingsGetAciBindingsAciBindings struct {
	ID          string `json:"id,omitempty"`          // Resource UUID value
	Name        string `json:"name,omitempty"`        // Resource Name
	Description string `json:"description,omitempty"` //
	IP          string `json:"ip,omitempty"`          // Binding IPv4 address. Each binding will be exclusively identified by its IP address and virtual network
	SgtValue    string `json:"sgtValue,omitempty"`    // Security Group Tag (SGT) value. The valid range for SGT values is 0-65534
	Vn          string `json:"vn,omitempty"`          // Virtual network. Each binding will be exclusively identified by its IP address and virtual network
	Psn         string `json:"psn,omitempty"`         // Cisco ISE Policy Service node (PSN) IP address
	LearnedFrom string `json:"learnedFrom,omitempty"` // Binding Source
	LearnedBy   string `json:"learnedBy,omitempty"`   // Binding Type
}

type ResponseAciBindingsGetVersion struct {
	VersionInfo *ResponseAciBindingsGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseAciBindingsGetVersionVersionInfo struct {
	CurrentServerVersion string                                        `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                        `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseAciBindingsGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseAciBindingsGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

//GetAciBindings Get all binding information
/* This API allows clients to retrieve all the bindings that were sent to Cisco ISE by ACI or received on ACI from Cisco ISE.The binding information will be identical to the information on ACI bindings page in the Cisco ISE UI. Filtering will be based on one attribute only, such as ip/sgt/vn/psn/learnedFrom/learnedBy with CONTAINS mode of search.

@param getAciBindingsQueryParams Filtering parameter
*/
func (s *AciBindingsService) GetAciBindings(getAciBindingsQueryParams *GetAciBindingsQueryParams) (*ResponseAciBindingsGetAciBindings, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/acibindings/getall"

	queryString, _ := query.Values(getAciBindingsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAciBindingsGetAciBindings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAciBindings")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAciBindingsGetAciBindings)
	return result, response, err

}

//GetVersion Get ACI bindings version information
/* This API helps to retrieve the version information related to the Cisco ACI bindings.

 */
func (s *AciBindingsService) GetVersion() (*ResponseAciBindingsGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/acibindings/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAciBindingsGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAciBindingsGetVersion)
	return result, response, err

}
