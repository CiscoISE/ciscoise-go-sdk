package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type PortalGlobalSettingService service

type GetPortalGlobalSettingsQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sortasc    string   `url:"sortasc,omitempty"`    //sort asc
	Sortdsc    string   `url:"sortdsc,omitempty"`    //sort desc
	Filter     []string `url:"filter,omitempty"`     //<br/> **Simple filtering** should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query string parameter. Each resource Data model description should specify if an attribute is a filtered field. <br/>                Operator    | Description <br/>               ------------|----------------- <br/>               EQ          | Equals <br/>               NEQ         | Not Equals <br/>               GT          | Greater Than <br/>               LT          | Less Then <br/>               STARTSW     | Starts With <br/>               NSTARTSW    | Not Starts With <br/>               ENDSW       | Ends With <br/>               NENDSW      | Not Ends With <br/>               CONTAINS	  | Contains <br/>               NCONTAINS	  | Not Contains <br/>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponsePortalGlobalSettingGetPortalGlobalSettingByID struct {
	PortalCustomizationGlobalSetting ResponsePortalGlobalSettingGetPortalGlobalSettingByIDPortalCustomizationGlobalSetting `json:"PortalCustomizationGlobalSetting,omitempty"` //
}

type ResponsePortalGlobalSettingGetPortalGlobalSettingByIDPortalCustomizationGlobalSetting struct {
	ID            string                                                                                    `json:"id,omitempty"`            //
	Customization string                                                                                    `json:"customization,omitempty"` // Allowed values: - HTML, - HTMLANDJAVASCRIPT
	Link          ResponsePortalGlobalSettingGetPortalGlobalSettingByIDPortalCustomizationGlobalSettingLink `json:"link,omitempty"`          //
}

type ResponsePortalGlobalSettingGetPortalGlobalSettingByIDPortalCustomizationGlobalSettingLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePortalGlobalSettingUpdatePortalGlobalSettingByID struct {
	UpdatedFieldsList ResponsePortalGlobalSettingUpdatePortalGlobalSettingByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponsePortalGlobalSettingUpdatePortalGlobalSettingByIDUpdatedFieldsList struct {
	UpdatedField ResponsePortalGlobalSettingUpdatePortalGlobalSettingByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                                `json:"field,omitempty"`        //
	OldValue     string                                                                                `json:"oldValue,omitempty"`     //
	NewValue     string                                                                                `json:"newValue,omitempty"`     //
}

type ResponsePortalGlobalSettingUpdatePortalGlobalSettingByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponsePortalGlobalSettingGetPortalGlobalSettings struct {
	SearchResult ResponsePortalGlobalSettingGetPortalGlobalSettingsSearchResult `json:"SearchResult,omitempty"` //
}

type ResponsePortalGlobalSettingGetPortalGlobalSettingsSearchResult struct {
	Total        int                                                                        `json:"total,omitempty"`        //
	Resources    []ResponsePortalGlobalSettingGetPortalGlobalSettingsSearchResultResources  `json:"resources,omitempty"`    //
	NextPage     ResponsePortalGlobalSettingGetPortalGlobalSettingsSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage ResponsePortalGlobalSettingGetPortalGlobalSettingsSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponsePortalGlobalSettingGetPortalGlobalSettingsSearchResultResources struct {
	ID   string                                                                      `json:"id,omitempty"`   //
	Link ResponsePortalGlobalSettingGetPortalGlobalSettingsSearchResultResourcesLink `json:"link,omitempty"` //
}

type ResponsePortalGlobalSettingGetPortalGlobalSettingsSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePortalGlobalSettingGetPortalGlobalSettingsSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePortalGlobalSettingGetPortalGlobalSettingsSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponsePortalGlobalSettingGetVersion struct {
	VersionInfo ResponsePortalGlobalSettingGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponsePortalGlobalSettingGetVersionVersionInfo struct {
	CurrentServerVersion string                                               `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                               `json:"supportedVersions,omitempty"`    //
	Link                 ResponsePortalGlobalSettingGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponsePortalGlobalSettingGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestPortalGlobalSettingUpdatePortalGlobalSettingByID struct {
	PortalCustomizationGlobalSetting *RequestPortalGlobalSettingUpdatePortalGlobalSettingByIDPortalCustomizationGlobalSetting `json:"PortalCustomizationGlobalSetting,omitempty"` //
}

type RequestPortalGlobalSettingUpdatePortalGlobalSettingByIDPortalCustomizationGlobalSetting struct {
	ID            string `json:"id,omitempty"`            //
	Customization string `json:"customization,omitempty"` // Allowed values: - HTML, - HTMLANDJAVASCRIPT
}

//GetPortalGlobalSettingByID Get portal global settings by id
/* This API allows the client to get the portal global settings by id.

@param id id path parameter.
*/
func (s *PortalGlobalSettingService) GetPortalGlobalSettingByID(id string) (*ResponsePortalGlobalSettingGetPortalGlobalSettingByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/portalglobalsetting/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePortalGlobalSettingGetPortalGlobalSettingByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPortalGlobalSettingById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePortalGlobalSettingGetPortalGlobalSettingByID)
	return result, response, err

}

//GetPortalGlobalSettings Get all portal global settings
/* This API allows the client to get all the portal global settings.

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

[name, description]

@param getPortalGlobalSettingsQueryParams Filtering parameter
*/
func (s *PortalGlobalSettingService) GetPortalGlobalSettings(getPortalGlobalSettingsQueryParams *GetPortalGlobalSettingsQueryParams) (*ResponsePortalGlobalSettingGetPortalGlobalSettings, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/portalglobalsetting"

	queryString, _ := query.Values(getPortalGlobalSettingsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponsePortalGlobalSettingGetPortalGlobalSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPortalGlobalSettings")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePortalGlobalSettingGetPortalGlobalSettings)
	return result, response, err

}

//GetVersion Get portal global setting version information
/* This API helps to retrieve the version information related to the portal global setting .

 */
func (s *PortalGlobalSettingService) GetVersion() (*ResponsePortalGlobalSettingGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/portalglobalsetting/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePortalGlobalSettingGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePortalGlobalSettingGetVersion)
	return result, response, err

}

//UpdatePortalGlobalSettingByID Update portal global settings by id
/* This API allows the client to update the portal global settings by id.

@param id id path parameter.
*/
func (s *PortalGlobalSettingService) UpdatePortalGlobalSettingByID(id string, requestPortalGlobalSettingUpdatePortalGlobalSettingById *RequestPortalGlobalSettingUpdatePortalGlobalSettingByID) (*ResponsePortalGlobalSettingUpdatePortalGlobalSettingByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/portalglobalsetting/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestPortalGlobalSettingUpdatePortalGlobalSettingById).
		SetResult(&ResponsePortalGlobalSettingUpdatePortalGlobalSettingByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdatePortalGlobalSettingById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePortalGlobalSettingUpdatePortalGlobalSettingByID)
	return result, response, err

}
