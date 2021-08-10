package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type DeviceAdministrationTimeDateConditionsService service

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditions struct {
	Response []interface{} `json:"response,omitempty"` //
	Version  string        `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeCondition struct {
	ConditionType       string                                                                                          `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            bool                                                                                            `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionLink                `json:"link,omitempty"`                //
	DatesRange          ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	Description         string                                                                                          `json:"description,omitempty"`         // Condition description
	HoursRange          ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	ID                  string                                                                                          `json:"id,omitempty"`                  //
	Name                string                                                                                          `json:"name,omitempty"`                // Condition name
	WeekDays            []string                                                                                        `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                        `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
	Version             string                                                                                          `json:"version,omitempty"`             //
}

type ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByID struct {
	ConditionType       string                                                                                           `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            bool                                                                                             `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDLink                `json:"link,omitempty"`                //
	DatesRange          ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	Description         string                                                                                           `json:"description,omitempty"`         // Condition description
	HoursRange          ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	ID                  string                                                                                           `json:"id,omitempty"`                  //
	Name                string                                                                                           `json:"name,omitempty"`                // Condition name
	WeekDays            []string                                                                                         `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                         `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
	Version             string                                                                                           `json:"version,omitempty"`             //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByID struct {
	ConditionType       string                                                                                              `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            bool                                                                                                `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDLink                `json:"link,omitempty"`                //
	DatesRange          ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	Description         string                                                                                              `json:"description,omitempty"`         // Condition description
	HoursRange          ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	ID                  string                                                                                              `json:"id,omitempty"`                  //
	Name                string                                                                                              `json:"name,omitempty"`                // Condition name
	WeekDays            []string                                                                                            `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                            `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
	Version             string                                                                                              `json:"version,omitempty"`             //
}

type ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsDeleteDeviceAdminTimeConditionByID struct {
	ID string `json:"id,omitempty"` //
}

type RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeCondition struct {
	ConditionType       string                                                                                          `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                           `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                          `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                          `json:"id,omitempty"`                  //
	Name                string                                                                                          `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                          `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                          `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                          `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                          `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                          `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                          `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                        `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                        `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionChildren struct {
	ConditionType string                                                                                   `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                    `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByID struct {
	ConditionType       string                                                                                              `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                               `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDLink                `json:"link,omitempty"`                //
	Description         string                                                                                              `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                              `json:"id,omitempty"`                  //
	Name                string                                                                                              `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                              `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                              `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                              `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                              `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                              `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                              `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                            `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                            `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDChildren struct {
	ConditionType string                                                                                       `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                        `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDChildrenLink `json:"link,omitempty"`          //
}

type RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

//GetDeviceAdminTimeConditions Device Admin - Returns a list of time and date conditions.
/* Device Admin Returns a list of time and date conditions.

 */
func (s *DeviceAdministrationTimeDateConditionsService) GetDeviceAdminTimeConditions() (*ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditions, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/time-condition"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminTimeConditions")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditions)
	return result, response, err

}

//GetDeviceAdminTimeConditionByID Device Admin - Returns a network condition.
/* Device Admin Returns a network condition.

@param id id path parameter. Condition id
*/
func (s *DeviceAdministrationTimeDateConditionsService) GetDeviceAdminTimeConditionByID(id string) (*ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/time-condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminTimeConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByID)
	return result, response, err

}

//CreateDeviceAdminTimeCondition Device Admin - Creates time/date condition.
/* Device Admin Creates time/date condition.

 */
func (s *DeviceAdministrationTimeDateConditionsService) CreateDeviceAdminTimeCondition(requestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeCondition *RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeCondition) (*ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeCondition, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/time-condition"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeCondition).
		SetResult(&ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeCondition{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceAdminTimeCondition")
	}

	result := response.Result().(*ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeCondition)
	return result, response, err

}

//UpdateDeviceAdminTimeConditionByID Device Admin - Update network condition.
/* Device Admin Update network condition.

@param id id path parameter. Condition id
*/
func (s *DeviceAdministrationTimeDateConditionsService) UpdateDeviceAdminTimeConditionByID(id string, requestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionById *RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByID) (*ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/time-condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionById).
		SetResult(&ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceAdminTimeConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByID)
	return result, response, err

}

//DeleteDeviceAdminTimeConditionByID Device Admin - Delete Time/Date condition.
/* Device Admin Delete Time/Date condition.

@param id id path parameter. Condition id
*/
func (s *DeviceAdministrationTimeDateConditionsService) DeleteDeviceAdminTimeConditionByID(id string) (*ResponseDeviceAdministrationTimeDateConditionsDeleteDeviceAdminTimeConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/time-condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationTimeDateConditionsDeleteDeviceAdminTimeConditionByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteDeviceAdminTimeConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationTimeDateConditionsDeleteDeviceAdminTimeConditionByID)
	return result, response, err

}
