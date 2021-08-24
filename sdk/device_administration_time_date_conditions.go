package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type DeviceAdministrationTimeDateConditionsService service

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditions struct {
	Response []ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponse `json:"response,omitempty"` //
	Version  string                                                                               `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponse struct {
	ConditionType       string                                                                                                `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            bool                                                                                                  `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                                `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                `json:"id,omitempty"`                  //
	Name                string                                                                                                `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                `json:"operator,omitempty"`            // Equality operator
	Children            []ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                              `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                              `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseChildren struct {
	ConditionType string                                                                                         `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      bool                                                                                           `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeCondition struct {
	Response ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionResponse `json:"response,omitempty"` //
	Version  string                                                                               `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionResponse struct {
	ConditionType       string                                                                                                  `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            bool                                                                                                    `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                                  `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                  `json:"id,omitempty"`                  //
	Name                string                                                                                                  `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                  `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                  `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                  `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                  `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                  `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                  `json:"operator,omitempty"`            // Equality operator
	Children            []ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionResponseChildren struct {
	ConditionType string                                                                                           `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      bool                                                                                             `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByID struct {
	Response ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponse `json:"response,omitempty"` //
	Version  string                                                                                `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponse struct {
	ConditionType       string                                                                                                   `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            bool                                                                                                     `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                                   `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                   `json:"id,omitempty"`                  //
	Name                string                                                                                                   `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                   `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                   `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                   `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                   `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                   `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                   `json:"operator,omitempty"`            // Equality operator
	Children            []ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                 `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                 `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseChildren struct {
	ConditionType string                                                                                            `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      bool                                                                                              `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByID struct {
	Response ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDResponse `json:"response,omitempty"` //
	Version  string                                                                                   `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDResponse struct {
	ConditionType       string                                                                                                      `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            bool                                                                                                        `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                                      `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                      `json:"id,omitempty"`                  //
	Name                string                                                                                                      `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                      `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                      `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                      `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                      `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                      `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                      `json:"operator,omitempty"`            // Equality operator
	Children            []ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                    `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                    `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDResponseChildren struct {
	ConditionType string                                                                                               `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      bool                                                                                                 `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDResponseHoursRangeException struct {
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
