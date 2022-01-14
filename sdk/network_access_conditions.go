package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type NetworkAccessConditionsService service

type ResponseNetworkAccessConditionsGetNetworkAccessConditions struct {
	Response *[]ResponseNetworkAccessConditionsGetNetworkAccessConditionsResponse `json:"response,omitempty"` //
	Version  string                                                               `json:"version,omitempty"`  //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsResponse struct {
	ConditionType       string                                                                                `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                 `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessConditionsGetNetworkAccessConditionsResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                `json:"id,omitempty"`                  //
	Name                string                                                                                `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessConditionsGetNetworkAccessConditionsResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessConditionsGetNetworkAccessConditionsResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessConditionsGetNetworkAccessConditionsResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessConditionsGetNetworkAccessConditionsResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessConditionsGetNetworkAccessConditionsResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                              `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                              `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsResponseChildren struct {
	ConditionType string                                                                         `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                          `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessConditionsGetNetworkAccessConditionsResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsCreateNetworkAccessCondition struct {
	Response *ResponseNetworkAccessConditionsCreateNetworkAccessConditionResponse `json:"response,omitempty"` //
	Version  string                                                               `json:"version,omitempty"`  //
}

type ResponseNetworkAccessConditionsCreateNetworkAccessConditionResponse struct {
	ConditionType       string                                                                                  `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                   `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessConditionsCreateNetworkAccessConditionResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                  `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                  `json:"id,omitempty"`                  //
	Name                string                                                                                  `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                  `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                  `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                  `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                  `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                  `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessConditionsCreateNetworkAccessConditionResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessConditionsCreateNetworkAccessConditionResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessConditionsCreateNetworkAccessConditionResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessConditionsCreateNetworkAccessConditionResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessConditionsCreateNetworkAccessConditionResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessConditionsCreateNetworkAccessConditionResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsCreateNetworkAccessConditionResponseChildren struct {
	ConditionType string                                                                           `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                            `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessConditionsCreateNetworkAccessConditionResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessConditionsCreateNetworkAccessConditionResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsCreateNetworkAccessConditionResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsCreateNetworkAccessConditionResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsCreateNetworkAccessConditionResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsCreateNetworkAccessConditionResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRules struct {
	Response *[]ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRulesResponse `json:"response,omitempty"` //
	Version  string                                                                                     `json:"version,omitempty"`  //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRulesResponse struct {
	ConditionType       string                                                                                                      `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                       `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRulesResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                                      `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                      `json:"id,omitempty"`                  //
	Name                string                                                                                                      `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                      `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                      `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                      `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                      `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                      `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRulesResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRulesResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRulesResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRulesResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRulesResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                    `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                    `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRulesResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRulesResponseChildren struct {
	ConditionType string                                                                                               `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRulesResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRulesResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRulesResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRulesResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRulesResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRulesResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRules struct {
	Response *[]ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRulesResponse `json:"response,omitempty"` //
	Version  string                                                                                    `json:"version,omitempty"`  //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRulesResponse struct {
	ConditionType       string                                                                                                     `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                      `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRulesResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                                     `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                     `json:"id,omitempty"`                  //
	Name                string                                                                                                     `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                     `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                     `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                     `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                     `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                     `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRulesResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRulesResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRulesResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRulesResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRulesResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                   `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                   `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRulesResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRulesResponseChildren struct {
	ConditionType string                                                                                              `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                               `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRulesResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRulesResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRulesResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRulesResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRulesResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRulesResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByName struct {
	Response *ResponseNetworkAccessConditionsGetNetworkAccessConditionByNameResponse `json:"response,omitempty"` //
	Version  string                                                                  `json:"version,omitempty"`  //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByNameResponse struct {
	ConditionType       string                                                                                     `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                      `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessConditionsGetNetworkAccessConditionByNameResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                     `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                     `json:"id,omitempty"`                  //
	Name                string                                                                                     `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                     `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                     `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                     `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                     `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                     `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessConditionsGetNetworkAccessConditionByNameResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessConditionsGetNetworkAccessConditionByNameResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessConditionsGetNetworkAccessConditionByNameResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessConditionsGetNetworkAccessConditionByNameResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessConditionsGetNetworkAccessConditionByNameResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                   `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                   `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByNameResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByNameResponseChildren struct {
	ConditionType string                                                                              `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                               `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessConditionsGetNetworkAccessConditionByNameResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByNameResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByNameResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByNameResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByNameResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByNameResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByName struct {
	Response *ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByNameResponse `json:"response,omitempty"` //
	Version  string                                                                     `json:"version,omitempty"`  //
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByNameResponse struct {
	ConditionType       string                                                                                        `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                         `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByNameResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                        `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                        `json:"id,omitempty"`                  //
	Name                string                                                                                        `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                        `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                        `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                        `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                        `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                        `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByNameResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByNameResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByNameResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByNameResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByNameResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                      `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                      `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByNameResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByNameResponseChildren struct {
	ConditionType string                                                                                 `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                  `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByNameResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByNameResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByNameResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByNameResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByNameResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByNameResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsDeleteNetworkAccessConditionByName struct {
	ID string `json:"id,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySets struct {
	Response *[]ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySetsResponse `json:"response,omitempty"` //
	Version  string                                                                            `json:"version,omitempty"`  //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySetsResponse struct {
	ConditionType       string                                                                                             `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                              `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySetsResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                             `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                             `json:"id,omitempty"`                  //
	Name                string                                                                                             `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                             `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                             `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                             `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                             `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                             `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySetsResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySetsResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySetsResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySetsResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySetsResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                           `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                           `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySetsResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySetsResponseChildren struct {
	ConditionType string                                                                                      `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                       `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySetsResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySetsResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySetsResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySetsResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySetsResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySetsResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByID struct {
	Response *ResponseNetworkAccessConditionsGetNetworkAccessConditionByIDResponse `json:"response,omitempty"` //
	Version  string                                                                `json:"version,omitempty"`  //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByIDResponse struct {
	ConditionType       string                                                                                   `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                    `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessConditionsGetNetworkAccessConditionByIDResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                   `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                   `json:"id,omitempty"`                  //
	Name                string                                                                                   `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                   `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                   `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                   `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                   `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                   `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessConditionsGetNetworkAccessConditionByIDResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessConditionsGetNetworkAccessConditionByIDResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessConditionsGetNetworkAccessConditionByIDResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessConditionsGetNetworkAccessConditionByIDResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessConditionsGetNetworkAccessConditionByIDResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                 `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                 `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByIDResponseChildren struct {
	ConditionType string                                                                            `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                             `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessConditionsGetNetworkAccessConditionByIDResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByIDResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByIDResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByIDResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByIDResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsGetNetworkAccessConditionByIDResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByID struct {
	Response *ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByIDResponse `json:"response,omitempty"` //
	Version  string                                                                   `json:"version,omitempty"`  //
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByIDResponse struct {
	ConditionType       string                                                                                      `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                       `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByIDResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                      `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                      `json:"id,omitempty"`                  //
	Name                string                                                                                      `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                      `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                      `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                      `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                      `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                      `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByIDResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByIDResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByIDResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByIDResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByIDResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                    `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                    `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByIDResponseChildren struct {
	ConditionType string                                                                               `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByIDResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByIDResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByIDResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByIDResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByIDResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByIDResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessConditionsDeleteNetworkAccessConditionByID struct {
	ID string `json:"id,omitempty"` //
}

type RequestNetworkAccessConditionsCreateNetworkAccessCondition struct {
	ConditionType       string                                                                         `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                          `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestNetworkAccessConditionsCreateNetworkAccessConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                         `json:"description,omitempty"`         // Condition description
	ID                  string                                                                         `json:"id,omitempty"`                  //
	Name                string                                                                         `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                         `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                         `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                         `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                         `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                         `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestNetworkAccessConditionsCreateNetworkAccessConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestNetworkAccessConditionsCreateNetworkAccessConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestNetworkAccessConditionsCreateNetworkAccessConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestNetworkAccessConditionsCreateNetworkAccessConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestNetworkAccessConditionsCreateNetworkAccessConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                       `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                       `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestNetworkAccessConditionsCreateNetworkAccessConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessConditionsCreateNetworkAccessConditionChildren struct {
	ConditionType string                                                                  `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                   `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestNetworkAccessConditionsCreateNetworkAccessConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestNetworkAccessConditionsCreateNetworkAccessConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessConditionsCreateNetworkAccessConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessConditionsCreateNetworkAccessConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessConditionsCreateNetworkAccessConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessConditionsCreateNetworkAccessConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessConditionsUpdateNetworkAccessConditionByName struct {
	ConditionType       string                                                                               `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestNetworkAccessConditionsUpdateNetworkAccessConditionByNameLink                `json:"link,omitempty"`                //
	Description         string                                                                               `json:"description,omitempty"`         // Condition description
	ID                  string                                                                               `json:"id,omitempty"`                  //
	Name                string                                                                               `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                               `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                               `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                               `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                               `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                               `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestNetworkAccessConditionsUpdateNetworkAccessConditionByNameChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestNetworkAccessConditionsUpdateNetworkAccessConditionByNameDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestNetworkAccessConditionsUpdateNetworkAccessConditionByNameDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestNetworkAccessConditionsUpdateNetworkAccessConditionByNameHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestNetworkAccessConditionsUpdateNetworkAccessConditionByNameHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                             `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                             `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestNetworkAccessConditionsUpdateNetworkAccessConditionByNameLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessConditionsUpdateNetworkAccessConditionByNameChildren struct {
	ConditionType string                                                                        `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                         `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestNetworkAccessConditionsUpdateNetworkAccessConditionByNameChildrenLink `json:"link,omitempty"`          //
}

type RequestNetworkAccessConditionsUpdateNetworkAccessConditionByNameChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessConditionsUpdateNetworkAccessConditionByNameDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessConditionsUpdateNetworkAccessConditionByNameDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessConditionsUpdateNetworkAccessConditionByNameHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessConditionsUpdateNetworkAccessConditionByNameHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessConditionsUpdateNetworkAccessConditionByID struct {
	ConditionType       string                                                                             `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                              `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestNetworkAccessConditionsUpdateNetworkAccessConditionByIDLink                `json:"link,omitempty"`                //
	Description         string                                                                             `json:"description,omitempty"`         // Condition description
	ID                  string                                                                             `json:"id,omitempty"`                  //
	Name                string                                                                             `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                             `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                             `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                             `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                             `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                             `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestNetworkAccessConditionsUpdateNetworkAccessConditionByIDChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestNetworkAccessConditionsUpdateNetworkAccessConditionByIDDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestNetworkAccessConditionsUpdateNetworkAccessConditionByIDDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestNetworkAccessConditionsUpdateNetworkAccessConditionByIDHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestNetworkAccessConditionsUpdateNetworkAccessConditionByIDHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                           `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                           `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestNetworkAccessConditionsUpdateNetworkAccessConditionByIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessConditionsUpdateNetworkAccessConditionByIDChildren struct {
	ConditionType string                                                                      `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                       `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestNetworkAccessConditionsUpdateNetworkAccessConditionByIDChildrenLink `json:"link,omitempty"`          //
}

type RequestNetworkAccessConditionsUpdateNetworkAccessConditionByIDChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessConditionsUpdateNetworkAccessConditionByIDDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessConditionsUpdateNetworkAccessConditionByIDDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessConditionsUpdateNetworkAccessConditionByIDHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessConditionsUpdateNetworkAccessConditionByIDHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

//GetNetworkAccessConditions Network Access - Returns all library conditions
/* Network Access Returns all library conditions

 */
func (s *NetworkAccessConditionsService) GetNetworkAccessConditions() (*ResponseNetworkAccessConditionsGetNetworkAccessConditions, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/condition"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessConditionsGetNetworkAccessConditions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessConditions")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessConditionsGetNetworkAccessConditions)
	return result, response, err

}

//GetNetworkAccessConditionsForAuthenticationRules Network Access - Returns list of library conditions for Authentication rules scope.
/* Network Access Returns list of library conditions for Authentication rules scope.

 */
func (s *NetworkAccessConditionsService) GetNetworkAccessConditionsForAuthenticationRules() (*ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRules, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/condition/authentication"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessConditionsForAuthenticationRules")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthenticationRules)
	return result, response, err

}

//GetNetworkAccessConditionsForAuthorizationRules Network Access - Returns list of library conditions for Authorization rules scope.
/* Network Access Returns list of library conditions for Authorization rules scope.

 */
func (s *NetworkAccessConditionsService) GetNetworkAccessConditionsForAuthorizationRules() (*ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRules, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/condition/authorization"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessConditionsForAuthorizationRules")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessConditionsGetNetworkAccessConditionsForAuthorizationRules)
	return result, response, err

}

//GetNetworkAccessConditionByName Network Access - Returns a library condition.
/* Network Access Returns a library condition.

@param name name path parameter. Condition name
*/
func (s *NetworkAccessConditionsService) GetNetworkAccessConditionByName(name string) (*ResponseNetworkAccessConditionsGetNetworkAccessConditionByName, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/condition/condition-by-name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessConditionsGetNetworkAccessConditionByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessConditionByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessConditionsGetNetworkAccessConditionByName)
	return result, response, err

}

//GetNetworkAccessConditionsForPolicySets Network Access - Returns list of library conditions for PolicySet scope.
/* Network Access Returns list of library conditions for PolicySet scope.

 */
func (s *NetworkAccessConditionsService) GetNetworkAccessConditionsForPolicySets() (*ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySets, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/condition/policyset"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySets{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessConditionsForPolicySets")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessConditionsGetNetworkAccessConditionsForPolicySets)
	return result, response, err

}

//GetNetworkAccessConditionByID Network Access - Returns a library condition.
/* Network Access Returns a library condition.

@param id id path parameter. Condition id
*/
func (s *NetworkAccessConditionsService) GetNetworkAccessConditionByID(id string) (*ResponseNetworkAccessConditionsGetNetworkAccessConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessConditionsGetNetworkAccessConditionByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessConditionsGetNetworkAccessConditionByID)
	return result, response, err

}

//CreateNetworkAccessCondition Network Access - Creates a library condition.
/* Network Access Creates a library condition:

 Library Condition has hierarchical structure which define a set of condition for which authentication and authorization policy rules could be match.

 Condition can be compose from single dictionary attribute name and value using model
LibraryConditionAttributes
 , or from combination of dictionary attributes with logical operation of AND/OR between them, using models:
LibraryConditionAndBlock
 or
LibraryConditionOrBlock
.

 When using AND/OR blocks, the condition will include inner layers inside these blocks, these layers are built using the inner condition models:
ConditionAttributes
,
ConditionAndBlock
,
ConditionOrBlock
, that represent dynamically built Conditions which are not stored in the conditions Library, or using
ConditionReference
, which includes an ID to existing stored condition in the library.

 The LibraryCondition models can only be used in the outer-most layer (root of the condition) and must always include the condition name.

 When using one of the 3 inner condition models (
ConditionAttributes, ConditionAndBlock, ConditionOrBlock
), condition name cannot be included in the request, since these will not be stored in the conditions library, and used only as inner members of the root condition.

 When using
ConditionReference
 model in inner layers, the condition name is not required.

 ConditionReference objects can also include a reference ID to a condition of type
TimeAndDate
.



*/
func (s *NetworkAccessConditionsService) CreateNetworkAccessCondition(requestNetworkAccessConditionsCreateNetworkAccessCondition *RequestNetworkAccessConditionsCreateNetworkAccessCondition) (*ResponseNetworkAccessConditionsCreateNetworkAccessCondition, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/condition"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessConditionsCreateNetworkAccessCondition).
		SetResult(&ResponseNetworkAccessConditionsCreateNetworkAccessCondition{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessConditionsCreateNetworkAccessCondition{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkAccessCondition")
	}

	result := response.Result().(*ResponseNetworkAccessConditionsCreateNetworkAccessCondition)
	return result, response, err

}

//UpdateNetworkAccessConditionByName Network Access - Update library condition by condition name.
/* Network Access Update library condition using condition name.

@param name name path parameter. Condition name
*/
func (s *NetworkAccessConditionsService) UpdateNetworkAccessConditionByName(name string, requestNetworkAccessConditionsUpdateNetworkAccessConditionByName *RequestNetworkAccessConditionsUpdateNetworkAccessConditionByName) (*ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByName, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/condition/condition-by-name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessConditionsUpdateNetworkAccessConditionByName).
		SetResult(&ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByName{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByName{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkAccessConditionByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByName)
	return result, response, err

}

//UpdateNetworkAccessConditionByID Network Access - Update library condition.
/* Network Access Update library condition.

@param id id path parameter. Condition id
*/
func (s *NetworkAccessConditionsService) UpdateNetworkAccessConditionByID(id string, requestNetworkAccessConditionsUpdateNetworkAccessConditionById *RequestNetworkAccessConditionsUpdateNetworkAccessConditionByID) (*ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessConditionsUpdateNetworkAccessConditionById).
		SetResult(&ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByID{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkAccessConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessConditionsUpdateNetworkAccessConditionByID)
	return result, response, err

}

//DeleteNetworkAccessConditionByName Network Access - Delete a library condition by condition Name.
/* Network Access Delete a library condition using condition Name.

@param name name path parameter. Condition name
*/
func (s *NetworkAccessConditionsService) DeleteNetworkAccessConditionByName(name string) (*ResponseNetworkAccessConditionsDeleteNetworkAccessConditionByName, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/condition/condition-by-name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessConditionsDeleteNetworkAccessConditionByName{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteNetworkAccessConditionByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessConditionsDeleteNetworkAccessConditionByName)
	return result, response, err

}

//DeleteNetworkAccessConditionByID Network Access - Delete a library condition.
/* Network Access Delete a library condition.

@param id id path parameter. Condition id
*/
func (s *NetworkAccessConditionsService) DeleteNetworkAccessConditionByID(id string) (*ResponseNetworkAccessConditionsDeleteNetworkAccessConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessConditionsDeleteNetworkAccessConditionByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteNetworkAccessConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessConditionsDeleteNetworkAccessConditionByID)
	return result, response, err

}
