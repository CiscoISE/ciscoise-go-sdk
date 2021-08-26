package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type DeviceAdministrationConditionsService service

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditions struct {
	Response *[]ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponse `json:"response,omitempty"` //
	Version  string                                                                    `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponse struct {
	ConditionType       string                                                                                     `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                      `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                     `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                     `json:"id,omitempty"`                  //
	Name                string                                                                                     `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                     `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                     `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                     `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                     `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                     `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                     `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                   `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                   `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseChildren struct {
	ConditionType string                                                                              `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                               `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsCreateDeviceAdminCondition struct {
	Response *ResponseDeviceAdministrationConditionsCreateDeviceAdminConditionResponse `json:"response,omitempty"` //
	Version  string                                                                    `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationConditionsCreateDeviceAdminConditionResponse struct {
	ConditionType       string                                                                                       `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                        `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationConditionsCreateDeviceAdminConditionResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                       `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                       `json:"id,omitempty"`                  //
	Name                string                                                                                       `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                       `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                       `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                       `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                       `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                       `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                       `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationConditionsCreateDeviceAdminConditionResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationConditionsCreateDeviceAdminConditionResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationConditionsCreateDeviceAdminConditionResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationConditionsCreateDeviceAdminConditionResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationConditionsCreateDeviceAdminConditionResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                     `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                     `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationConditionsCreateDeviceAdminConditionResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsCreateDeviceAdminConditionResponseChildren struct {
	ConditionType string                                                                                `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                 `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseDeviceAdministrationConditionsCreateDeviceAdminConditionResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationConditionsCreateDeviceAdminConditionResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsCreateDeviceAdminConditionResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsCreateDeviceAdminConditionResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsCreateDeviceAdminConditionResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsCreateDeviceAdminConditionResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRules struct {
	Response *[]ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRulesResponse `json:"response,omitempty"` //
	Version  string                                                                                          `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRulesResponse struct {
	ConditionType       string                                                                                                           `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                            `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRulesResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                                           `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                           `json:"id,omitempty"`                  //
	Name                string                                                                                                           `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                           `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                           `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                           `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                           `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                           `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                           `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRulesResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRulesResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRulesResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRulesResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRulesResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                         `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                         `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRulesResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRulesResponseChildren struct {
	ConditionType string                                                                                                    `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                     `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRulesResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRulesResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRulesResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRulesResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRulesResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRulesResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRules struct {
	Response *[]ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponse `json:"response,omitempty"` //
	Version  string                                                                                         `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponse struct {
	ConditionType       string                                                                                                          `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                           `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                                          `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                          `json:"id,omitempty"`                  //
	Name                string                                                                                                          `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                          `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                          `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                          `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                          `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                          `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                          `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                        `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                        `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseChildren struct {
	ConditionType string                                                                                                   `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                    `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByName struct {
	Response *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponse `json:"response,omitempty"` //
	Version  string                                                                       `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponse struct {
	ConditionType       string                                                                                          `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                           `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                          `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                          `json:"id,omitempty"`                  //
	Name                string                                                                                          `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                          `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                          `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                          `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                          `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                          `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                          `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                        `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                        `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseChildren struct {
	ConditionType string                                                                                   `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                    `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByName struct {
	Response *ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameResponse `json:"response,omitempty"` //
	Version  string                                                                          `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameResponse struct {
	ConditionType       string                                                                                             `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                              `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                             `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                             `json:"id,omitempty"`                  //
	Name                string                                                                                             `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                             `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                             `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                             `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                             `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                             `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                             `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                           `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                           `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameResponseChildren struct {
	ConditionType string                                                                                      `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                       `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsDeleteDeviceAdminConditionByName struct {
	ID string `json:"id,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySets struct {
	Response *[]ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponse `json:"response,omitempty"` //
	Version  string                                                                                 `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponse struct {
	ConditionType       string                                                                                                  `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                   `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                                  `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                  `json:"id,omitempty"`                  //
	Name                string                                                                                                  `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                  `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                  `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                  `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                  `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                  `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                  `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseChildren struct {
	ConditionType string                                                                                           `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                            `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByID struct {
	Response *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponse `json:"response,omitempty"` //
	Version  string                                                                     `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponse struct {
	ConditionType       string                                                                                        `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                         `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                        `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                        `json:"id,omitempty"`                  //
	Name                string                                                                                        `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                        `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                        `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                        `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                        `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                        `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                        `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                      `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                      `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseChildren struct {
	ConditionType string                                                                                 `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                  `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByID struct {
	Response *ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDResponse `json:"response,omitempty"` //
	Version  string                                                                        `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDResponse struct {
	ConditionType       string                                                                                           `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                            `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                           `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                           `json:"id,omitempty"`                  //
	Name                string                                                                                           `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                           `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                           `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                           `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                           `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                           `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                           `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                         `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                         `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDResponseChildren struct {
	ConditionType string                                                                                    `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                     `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationConditionsDeleteDeviceAdminConditionByID struct {
	ID string `json:"id,omitempty"` //
}

type RequestDeviceAdministrationConditionsCreateDeviceAdminCondition struct {
	ConditionType       string                                                                              `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                               `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestDeviceAdministrationConditionsCreateDeviceAdminConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                              `json:"description,omitempty"`         // Condition description
	ID                  string                                                                              `json:"id,omitempty"`                  //
	Name                string                                                                              `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                              `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                              `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                              `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                              `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                              `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                              `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                            `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                            `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestDeviceAdministrationConditionsCreateDeviceAdminConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildren struct {
	ConditionType string                                                                       `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                        `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByName struct {
	ConditionType       string                                                                                    `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                     `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameLink                `json:"link,omitempty"`                //
	Description         string                                                                                    `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                    `json:"id,omitempty"`                  //
	Name                string                                                                                    `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                    `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                    `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                    `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                    `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                    `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                    `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                  `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                  `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameChildren struct {
	ConditionType string                                                                             `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                              `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameChildrenLink `json:"link,omitempty"`          //
}

type RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByNameHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByID struct {
	ConditionType       string                                                                                  `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                   `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDLink                `json:"link,omitempty"`                //
	Description         string                                                                                  `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                  `json:"id,omitempty"`                  //
	Name                string                                                                                  `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                  `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                  `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                  `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                  `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                  `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                  `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildren struct {
	ConditionType string                                                                           `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                            `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildrenLink `json:"link,omitempty"`          //
}

type RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

//GetDeviceAdminConditions Device Admin - Returns list of library conditions.
/* Device Admin Returns list of library conditions.

 */
func (s *DeviceAdministrationConditionsService) GetDeviceAdminConditions() (*ResponseDeviceAdministrationConditionsGetDeviceAdminConditions, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/condition"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationConditionsGetDeviceAdminConditions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminConditions")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationConditionsGetDeviceAdminConditions)
	return result, response, err

}

//GetDeviceAdminConditionsForAuthenticationRules Device Admin - Returns list of library conditions for authentication rules.
/* Device Admin Returns list of library conditions for authentication rules.

 */
func (s *DeviceAdministrationConditionsService) GetDeviceAdminConditionsForAuthenticationRules() (*ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRules, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/condition/authentication"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminConditionsForAuthenticationRules")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthenticationRules)
	return result, response, err

}

//GetDeviceAdminConditionsForAuthorizationRules Device Admin - Returns list of library conditions for authorization rules.
/* Device Admin Returns list of library conditions for authorization rules.

 */
func (s *DeviceAdministrationConditionsService) GetDeviceAdminConditionsForAuthorizationRules() (*ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRules, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/condition/authorization"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminConditionsForAuthorizationRules")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRules)
	return result, response, err

}

//GetDeviceAdminConditionByName Device Admin - Returns a library condition.
/* Device Admin Returns a library condition.

@param name name path parameter. Condition name
*/
func (s *DeviceAdministrationConditionsService) GetDeviceAdminConditionByName(name string) (*ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByName, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/condition/condition-by-name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminConditionByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByName)
	return result, response, err

}

//GetDeviceAdminConditionsForPolicySets Device Admin - Returns list of library conditions for policy sets.
/* Device Admin Returns list of library conditions for policy sets.

 */
func (s *DeviceAdministrationConditionsService) GetDeviceAdminConditionsForPolicySets() (*ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySets, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/condition/policyset"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySets{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminConditionsForPolicySets")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySets)
	return result, response, err

}

//GetDeviceAdminConditionByID Device Admin - Returns a library condition.
/* Device Admin Returns a library condition.

@param id id path parameter. Condition id
*/
func (s *DeviceAdministrationConditionsService) GetDeviceAdminConditionByID(id string) (*ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByID)
	return result, response, err

}

//CreateDeviceAdminCondition Device Admin - Creates a library condition.
/* Device Admin Creates a library condition.

 */
func (s *DeviceAdministrationConditionsService) CreateDeviceAdminCondition(requestDeviceAdministrationConditionsCreateDeviceAdminCondition *RequestDeviceAdministrationConditionsCreateDeviceAdminCondition) (*ResponseDeviceAdministrationConditionsCreateDeviceAdminCondition, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/condition"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceAdministrationConditionsCreateDeviceAdminCondition).
		SetResult(&ResponseDeviceAdministrationConditionsCreateDeviceAdminCondition{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceAdminCondition")
	}

	result := response.Result().(*ResponseDeviceAdministrationConditionsCreateDeviceAdminCondition)
	return result, response, err

}

//UpdateDeviceAdminConditionByName Device Admin - Update library condition by condition name.
/* Device Admin Update library condition using condition name.

@param name name path parameter. Condition name
*/
func (s *DeviceAdministrationConditionsService) UpdateDeviceAdminConditionByName(name string, requestDeviceAdministrationConditionsUpdateDeviceAdminConditionByName *RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByName) (*ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByName, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/condition/condition-by-name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceAdministrationConditionsUpdateDeviceAdminConditionByName).
		SetResult(&ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByName{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceAdminConditionByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByName)
	return result, response, err

}

//UpdateDeviceAdminConditionByID Device Admin - Update library condition.
/* Device Admin Update library condition.

@param id id path parameter. Condition id
*/
func (s *DeviceAdministrationConditionsService) UpdateDeviceAdminConditionByID(id string, requestDeviceAdministrationConditionsUpdateDeviceAdminConditionById *RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByID) (*ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceAdministrationConditionsUpdateDeviceAdminConditionById).
		SetResult(&ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceAdminConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationConditionsUpdateDeviceAdminConditionByID)
	return result, response, err

}

//DeleteDeviceAdminConditionByName Device Admin - Delete a library condition by condition Name.
/* NDevice Admin Delete a library condition using condition Name.

@param name name path parameter. Condition name
*/
func (s *DeviceAdministrationConditionsService) DeleteDeviceAdminConditionByName(name string) (*ResponseDeviceAdministrationConditionsDeleteDeviceAdminConditionByName, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/condition/condition-by-name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationConditionsDeleteDeviceAdminConditionByName{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteDeviceAdminConditionByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationConditionsDeleteDeviceAdminConditionByName)
	return result, response, err

}

//DeleteDeviceAdminConditionByID Device Admin - Delete a library condition.
/* Device Admin Delete a library condition.

@param id id path parameter. Condition id
*/
func (s *DeviceAdministrationConditionsService) DeleteDeviceAdminConditionByID(id string) (*ResponseDeviceAdministrationConditionsDeleteDeviceAdminConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationConditionsDeleteDeviceAdminConditionByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteDeviceAdminConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationConditionsDeleteDeviceAdminConditionByID)
	return result, response, err

}
