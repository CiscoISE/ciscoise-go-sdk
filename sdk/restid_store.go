package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type RestidStoreService service

type GetRestIDStoreQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseRestidStoreGetRestIDStoreByName struct {
	ERSRestIDStore *ResponseRestidStoreGetRestIDStoreByNameERSRestIDStore `json:"ERSRestIDStore,omitempty"` //
}

type ResponseRestidStoreGetRestIDStoreByNameERSRestIDStore struct {
	ID                       string                                                                         `json:"id,omitempty"`                       //
	Name                     string                                                                         `json:"name,omitempty"`                     //
	Description              string                                                                         `json:"description,omitempty"`              //
	ErsRestIDStoreAttributes *ResponseRestidStoreGetRestIDStoreByNameERSRestIDStoreErsRestIDStoreAttributes `json:"ersRestIDStoreAttributes,omitempty"` //
	Link                     *ResponseRestidStoreGetRestIDStoreByNameERSRestIDStoreLink                     `json:"link,omitempty"`                     //
}

type ResponseRestidStoreGetRestIDStoreByNameERSRestIDStoreErsRestIDStoreAttributes struct {
	UsernameSuffix string                                                                                  `json:"usernameSuffix,omitempty"` // Suffix of the username domain
	RootURL        string                                                                                  `json:"rootUrl,omitempty"`        // url of the root of the RestIDStore
	Predefined     string                                                                                  `json:"predefined,omitempty"`     // The cloud provider connected to of the RestIDStore. Options are: - Azure, - Okta, -
	Headers        *[]ResponseRestidStoreGetRestIDStoreByNameERSRestIDStoreErsRestIDStoreAttributesHeaders `json:"headers,omitempty"`        //
}

type ResponseRestidStoreGetRestIDStoreByNameERSRestIDStoreErsRestIDStoreAttributesHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

type ResponseRestidStoreGetRestIDStoreByNameERSRestIDStoreLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseRestidStoreUpdateRestIDStoreByName struct {
	UpdatedFieldsList *ResponseRestidStoreUpdateRestIDStoreByNameUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseRestidStoreUpdateRestIDStoreByNameUpdatedFieldsList struct {
	UpdatedField *[]ResponseRestidStoreUpdateRestIDStoreByNameUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                     `json:"field,omitempty"`        //
	OldValue     string                                                                     `json:"oldValue,omitempty"`     //
	NewValue     string                                                                     `json:"newValue,omitempty"`     //
}

type ResponseRestidStoreUpdateRestIDStoreByNameUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseRestidStoreGetRestIDStoreByID struct {
	ERSRestIDStore *ResponseRestidStoreGetRestIDStoreByIDERSRestIDStore `json:"ERSRestIDStore,omitempty"` //
}

type ResponseRestidStoreGetRestIDStoreByIDERSRestIDStore struct {
	ID                       string                                                                       `json:"id,omitempty"`                       //
	Name                     string                                                                       `json:"name,omitempty"`                     //
	Description              string                                                                       `json:"description,omitempty"`              //
	ErsRestIDStoreAttributes *ResponseRestidStoreGetRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributes `json:"ersRestIDStoreAttributes,omitempty"` //
	Link                     *ResponseRestidStoreGetRestIDStoreByIDERSRestIDStoreLink                     `json:"link,omitempty"`                     //
}

type ResponseRestidStoreGetRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributes struct {
	UsernameSuffix string                                                                                `json:"usernameSuffix,omitempty"` // Suffix of the username domain
	RootURL        string                                                                                `json:"rootUrl,omitempty"`        // url of the root of the RestIDStore
	Predefined     string                                                                                `json:"predefined,omitempty"`     // The cloud provider connected to of the RestIDStore. Options are: - Azure, - Okta, -
	Headers        *[]ResponseRestidStoreGetRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributesHeaders `json:"headers,omitempty"`        //
}

type ResponseRestidStoreGetRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributesHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

type ResponseRestidStoreGetRestIDStoreByIDERSRestIDStoreLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseRestidStoreUpdateRestIDStoreByID struct {
	UpdatedFieldsList *ResponseRestidStoreUpdateRestIDStoreByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseRestidStoreUpdateRestIDStoreByIDUpdatedFieldsList struct {
	UpdatedField *[]ResponseRestidStoreUpdateRestIDStoreByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                   `json:"field,omitempty"`        //
	OldValue     string                                                                   `json:"oldValue,omitempty"`     //
	NewValue     string                                                                   `json:"newValue,omitempty"`     //
}

type ResponseRestidStoreUpdateRestIDStoreByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseRestidStoreGetRestIDStore struct {
	SearchResult *ResponseRestidStoreGetRestIDStoreSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseRestidStoreGetRestIDStoreSearchResult struct {
	Total        *int                                                       `json:"total,omitempty"`        //
	Resources    *[]ResponseRestidStoreGetRestIDStoreSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     *ResponseRestidStoreGetRestIDStoreSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseRestidStoreGetRestIDStoreSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseRestidStoreGetRestIDStoreSearchResultResources struct {
	ID          string                                                      `json:"id,omitempty"`          //
	Name        string                                                      `json:"name,omitempty"`        //
	Description string                                                      `json:"description,omitempty"` //
	Link        *ResponseRestidStoreGetRestIDStoreSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseRestidStoreGetRestIDStoreSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseRestidStoreGetRestIDStoreSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseRestidStoreGetRestIDStoreSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseRestidStoreGetVersion struct {
	VersionInfo *ResponseRestidStoreGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseRestidStoreGetVersionVersionInfo struct {
	CurrentServerVersion string                                        `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                        `json:"supportedVersions,omitempty"`    //
	Link                 *ResponseRestidStoreGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseRestidStoreGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestRestidStoreUpdateRestIDStoreByName struct {
	ERSRestIDStore *RequestRestidStoreUpdateRestIDStoreByNameERSRestIDStore `json:"ERSRestIDStore,omitempty"` //
}

type RequestRestidStoreUpdateRestIDStoreByNameERSRestIDStore struct {
	ID                       string                                                                           `json:"id,omitempty"`                       //
	Name                     string                                                                           `json:"name,omitempty"`                     //
	Description              string                                                                           `json:"description,omitempty"`              //
	ErsRestIDStoreAttributes *RequestRestidStoreUpdateRestIDStoreByNameERSRestIDStoreErsRestIDStoreAttributes `json:"ersRestIDStoreAttributes,omitempty"` //
}

type RequestRestidStoreUpdateRestIDStoreByNameERSRestIDStoreErsRestIDStoreAttributes struct {
	UsernameSuffix string                                                                                    `json:"usernameSuffix,omitempty"` // Suffix of the username domain
	RootURL        string                                                                                    `json:"rootUrl,omitempty"`        // url of the root of the RestIDStore
	Predefined     string                                                                                    `json:"predefined,omitempty"`     // The cloud provider connected to of the RestIDStore. Options are: - Azure, - Okta, -
	Headers        *[]RequestRestidStoreUpdateRestIDStoreByNameERSRestIDStoreErsRestIDStoreAttributesHeaders `json:"headers,omitempty"`        //
}

type RequestRestidStoreUpdateRestIDStoreByNameERSRestIDStoreErsRestIDStoreAttributesHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

type RequestRestidStoreUpdateRestIDStoreByID struct {
	ERSRestIDStore *RequestRestidStoreUpdateRestIDStoreByIDERSRestIDStore `json:"ERSRestIDStore,omitempty"` //
}

type RequestRestidStoreUpdateRestIDStoreByIDERSRestIDStore struct {
	ID                       string                                                                         `json:"id,omitempty"`                       //
	Name                     string                                                                         `json:"name,omitempty"`                     //
	Description              string                                                                         `json:"description,omitempty"`              //
	ErsRestIDStoreAttributes *RequestRestidStoreUpdateRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributes `json:"ersRestIDStoreAttributes,omitempty"` //
}

type RequestRestidStoreUpdateRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributes struct {
	UsernameSuffix string                                                                                  `json:"usernameSuffix,omitempty"` // Suffix of the username domain
	RootURL        string                                                                                  `json:"rootUrl,omitempty"`        // url of the root of the RestIDStore
	Predefined     string                                                                                  `json:"predefined,omitempty"`     // The cloud provider connected to of the RestIDStore. Options are: - Azure, - Okta, -
	Headers        *[]RequestRestidStoreUpdateRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributesHeaders `json:"headers,omitempty"`        //
}

type RequestRestidStoreUpdateRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributesHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

type RequestRestidStoreCreateRestIDStore struct {
	ERSRestIDStore *RequestRestidStoreCreateRestIDStoreERSRestIDStore `json:"ERSRestIDStore,omitempty"` //
}

type RequestRestidStoreCreateRestIDStoreERSRestIDStore struct {
	Name                     string                                                                     `json:"name,omitempty"`                     //
	Description              string                                                                     `json:"description,omitempty"`              //
	ErsRestIDStoreAttributes *RequestRestidStoreCreateRestIDStoreERSRestIDStoreErsRestIDStoreAttributes `json:"ersRestIDStoreAttributes,omitempty"` //
}

type RequestRestidStoreCreateRestIDStoreERSRestIDStoreErsRestIDStoreAttributes struct {
	UsernameSuffix string                                                                              `json:"usernameSuffix,omitempty"` // Suffix of the username domain
	RootURL        string                                                                              `json:"rootUrl,omitempty"`        // url of the root of the RestIDStore
	Predefined     string                                                                              `json:"predefined,omitempty"`     // The cloud provider connected to of the RestIDStore. Options are: - Azure, - Okta, -
	Headers        *[]RequestRestidStoreCreateRestIDStoreERSRestIDStoreErsRestIDStoreAttributesHeaders `json:"headers,omitempty"`        //
}

type RequestRestidStoreCreateRestIDStoreERSRestIDStoreErsRestIDStoreAttributesHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

//GetRestIDStoreByName Get REST ID store by name
/* This API allows the client to get a REST ID store by name.

@param name name path parameter.
*/
func (s *RestidStoreService) GetRestIDStoreByName(name string) (*ResponseRestidStoreGetRestIDStoreByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/restidstore/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseRestidStoreGetRestIDStoreByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetRestIdStoreByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseRestidStoreGetRestIDStoreByName)
	return result, response, err

}

//GetRestIDStoreByID Get REST ID store by ID
/* This API allows the client to get a REST ID store by ID.

@param id id path parameter.
*/
func (s *RestidStoreService) GetRestIDStoreByID(id string) (*ResponseRestidStoreGetRestIDStoreByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/restidstore/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseRestidStoreGetRestIDStoreByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetRestIdStoreById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseRestidStoreGetRestIDStoreByID)
	return result, response, err

}

//GetRestIDStore Get all REST ID stores
/* This API allows the client to get all the REST ID stores.

Filter:

[name]

To search resources by using
toDate
 column,follow the format:

DD-MON-YY (Example:13-SEP-18)


Day or Year:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13

Month:GET /ers/config/guestuser/?filter=toDate.CONTAINS.SEP

Date:GET /ers/config/guestuser/?filter=toDate.CONTAINS.13-SEP-18


Sorting:

[name]

@param getRestIdStoreQueryParams Filtering parameter
*/
func (s *RestidStoreService) GetRestIDStore(getRestIdStoreQueryParams *GetRestIDStoreQueryParams) (*ResponseRestidStoreGetRestIDStore, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/restidstore"

	queryString, _ := query.Values(getRestIdStoreQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseRestidStoreGetRestIDStore{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetRestIdStore")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseRestidStoreGetRestIDStore)
	return result, response, err

}

//GetVersion Get REST ID store version information
/* This API helps to retrieve the version information related to the REST ID store.

 */
func (s *RestidStoreService) GetVersion() (*ResponseRestidStoreGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/restidstore/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseRestidStoreGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseRestidStoreGetVersion)
	return result, response, err

}

//CreateRestIDStore Create REST ID store
/* This API creates a REST ID store.

 */
func (s *RestidStoreService) CreateRestIDStore(requestRestidStoreCreateRestIdStore *RequestRestidStoreCreateRestIDStore) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/restidstore"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestRestidStoreCreateRestIdStore).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateRestIdStore")
	}

	return response, err

}

//UpdateRestIDStoreByName Update REST ID store by name
/* This API allows the client to update a REST ID store by name.

@param name name path parameter.
*/
func (s *RestidStoreService) UpdateRestIDStoreByName(name string, requestRestidStoreUpdateRestIdStoreByName *RequestRestidStoreUpdateRestIDStoreByName) (*ResponseRestidStoreUpdateRestIDStoreByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/restidstore/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestRestidStoreUpdateRestIdStoreByName).
		SetResult(&ResponseRestidStoreUpdateRestIDStoreByName{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseRestidStoreUpdateRestIDStoreByName{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateRestIdStoreByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseRestidStoreUpdateRestIDStoreByName)
	return result, response, err

}

//UpdateRestIDStoreByID Update REST ID store
/* This API allows the client to update a REST ID store.

@param id id path parameter.
*/
func (s *RestidStoreService) UpdateRestIDStoreByID(id string, requestRestidStoreUpdateRestIdStoreById *RequestRestidStoreUpdateRestIDStoreByID) (*ResponseRestidStoreUpdateRestIDStoreByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/restidstore/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestRestidStoreUpdateRestIdStoreById).
		SetResult(&ResponseRestidStoreUpdateRestIDStoreByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseRestidStoreUpdateRestIDStoreByID{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateRestIdStoreById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseRestidStoreUpdateRestIDStoreByID)
	return result, response, err

}

//DeleteRestIDStoreByName Delete REST ID store by name
/* This API deletes a REST ID store by name.

@param name name path parameter.
*/
func (s *RestidStoreService) DeleteRestIDStoreByName(name string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/restidstore/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteRestIdStoreByName")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteRestIDStoreByID Delete REST ID store
/* This API deletes a REST ID store.

@param id id path parameter.
*/
func (s *RestidStoreService) DeleteRestIDStoreByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/restidstore/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteRestIdStoreById")
	}

	getCSFRToken(response.Header())

	return response, err

}
