package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type NetworkAccessTimeDateConditionsService service

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditions struct {
	Response *[]ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponse `json:"response,omitempty"` //
	Version  string                                                                           `json:"version,omitempty"`  //
}

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponse struct {
	ConditionType       string                                                                                            `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                             `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                            `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                            `json:"id,omitempty"`                  //
	Name                string                                                                                            `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                            `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                            `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                            `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                            `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                            `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                          `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                          `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponseChildren struct {
	ConditionType string                                                                                     `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                      `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeCondition struct {
	Response *ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionResponse `json:"response,omitempty"` //
	Version  string                                                                           `json:"version,omitempty"`  //
}

type ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionResponse struct {
	ConditionType       string                                                                                              `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                               `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                              `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                              `json:"id,omitempty"`                  //
	Name                string                                                                                              `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                              `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                              `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                              `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                              `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                              `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                            `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                            `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionResponseChildren struct {
	ConditionType string                                                                                       `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                        `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByID struct {
	Response *ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponse `json:"response,omitempty"` //
	Version  string                                                                            `json:"version,omitempty"`  //
}

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponse struct {
	ConditionType       string                                                                                               `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                               `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                               `json:"id,omitempty"`                  //
	Name                string                                                                                               `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                               `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                               `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                               `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                               `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                               `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                             `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                             `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseChildren struct {
	ConditionType string                                                                                        `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                         `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByID struct {
	Response *ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDResponse `json:"response,omitempty"` //
	Version  string                                                                               `json:"version,omitempty"`  //
}

type ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDResponse struct {
	ConditionType       string                                                                                                  `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                   `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDResponseLink                `json:"link,omitempty"`                //
	Description         string                                                                                                  `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                  `json:"id,omitempty"`                  //
	Name                string                                                                                                  `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                  `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                  `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                  `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                  `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                  `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDResponseChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDResponseDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDResponseDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDResponseHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDResponseHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDResponseChildren struct {
	ConditionType string                                                                                           `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                            `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDResponseChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDResponseChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDResponseDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDResponseDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDResponseHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDResponseHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessTimeDateConditionsDeleteNetworkAccessTimeConditionByID struct {
	ID string `json:"id,omitempty"` //
}

type RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeCondition struct {
	ConditionType       string                                                                                     `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                      `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                     `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                     `json:"id,omitempty"`                  //
	Name                string                                                                                     `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                     `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                     `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                     `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                     `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                     `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                   `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                   `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionChildren struct {
	ConditionType string                                                                              `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                               `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByID struct {
	ConditionType       string                                                                                         `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                          `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDLink                `json:"link,omitempty"`                //
	Description         string                                                                                         `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                         `json:"id,omitempty"`                  //
	Name                string                                                                                         `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                         `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                         `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                         `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                         `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                         `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                       `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                       `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDChildren struct {
	ConditionType string                                                                                  `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                   `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDChildrenLink `json:"link,omitempty"`          //
}

type RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

//GetNetworkAccessTimeConditions Network Access - Returns a list of time and date conditions.
/* Network Access Returns a list of time and date conditions.

 */
func (s *NetworkAccessTimeDateConditionsService) GetNetworkAccessTimeConditions() (*ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditions, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/time-condition"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessTimeConditions")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditions)
	return result, response, err

}

//GetNetworkAccessTimeConditionByID Network Access - returns a network condition.
/* Network Access returns a network condition.

@param id id path parameter. Condition id
*/
func (s *NetworkAccessTimeDateConditionsService) GetNetworkAccessTimeConditionByID(id string) (*ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/time-condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessTimeConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByID)
	return result, response, err

}

//CreateNetworkAccessTimeCondition Network Access - Creates time/date condition.
/* Network Access Creates time/date condition

 */
func (s *NetworkAccessTimeDateConditionsService) CreateNetworkAccessTimeCondition(requestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeCondition *RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeCondition) (*ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeCondition, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/time-condition"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeCondition).
		SetResult(&ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeCondition{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeCondition{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkAccessTimeCondition")
	}

	result := response.Result().(*ResponseNetworkAccessTimeDateConditionsCreateNetworkAccessTimeCondition)
	return result, response, err

}

//UpdateNetworkAccessTimeConditionByID Network Access - Update network condition.
/* Network Access Update network condition

@param id id path parameter. Condition id
*/
func (s *NetworkAccessTimeDateConditionsService) UpdateNetworkAccessTimeConditionByID(id string, requestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionById *RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByID) (*ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/time-condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionById).
		SetResult(&ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByID{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkAccessTimeConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByID)
	return result, response, err

}

//DeleteNetworkAccessTimeConditionByID Network Access - Delete Time/Date condition.
/* Network Access Delete Time/Date condition.

@param id id path parameter. Condition id
*/
func (s *NetworkAccessTimeDateConditionsService) DeleteNetworkAccessTimeConditionByID(id string) (*ResponseNetworkAccessTimeDateConditionsDeleteNetworkAccessTimeConditionByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/time-condition/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessTimeDateConditionsDeleteNetworkAccessTimeConditionByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteNetworkAccessTimeConditionById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessTimeDateConditionsDeleteNetworkAccessTimeConditionByID)
	return result, response, err

}
