package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type DeviceAdministrationPolicySetService service

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySets struct {
	Response *[]ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponse `json:"response,omitempty"` //
	Version  string                                                                   `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponse struct {
	Condition   *ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseCondition `json:"condition,omitempty"`   //
	Default     *bool                                                                           `json:"default,omitempty"`     // Flag which indicates if this policy set is the default one
	Description string                                                                          `json:"description,omitempty"` // The description for the policy set
	HitCounts   *int                                                                            `json:"hitCounts,omitempty"`   // The amount of times the policy was matched
	ID          string                                                                          `json:"id,omitempty"`          // Identifier for the policy set
	IsProxy     *bool                                                                           `json:"isProxy,omitempty"`     // Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	Link        *ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseLink      `json:"link,omitempty"`        //
	Name        string                                                                          `json:"name,omitempty"`        // Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank        *int                                                                            `json:"rank,omitempty"`        // The rank(priority) in relation to other policy set. Lower rank is higher priority.
	ServiceName string                                                                          `json:"serviceName,omitempty"` // Policy set service identifier - Allowed Protocols,Server Sequence..
	State       string                                                                          `json:"state,omitempty"`       // The state that the policy set is in. A disabled policy set cannot be matched.
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseCondition struct {
	ConditionType       string                                                                                             `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                              `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                             `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                             `json:"id,omitempty"`                  //
	Name                string                                                                                             `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                             `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                             `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                             `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                             `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                             `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                           `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                           `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionChildren struct {
	ConditionType string                                                                                      `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                       `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySet struct {
	Response *ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponse `json:"response,omitempty"` // Policy set structure
	Version  string                                                                   `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponse struct {
	Condition   *ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseCondition `json:"condition,omitempty"`   //
	Default     *bool                                                                             `json:"default,omitempty"`     // Flag which indicates if this policy set is the default one
	Description string                                                                            `json:"description,omitempty"` // The description for the policy set
	HitCounts   *int                                                                              `json:"hitCounts,omitempty"`   // The amount of times the policy was matched
	ID          string                                                                            `json:"id,omitempty"`          // Identifier for the policy set
	IsProxy     *bool                                                                             `json:"isProxy,omitempty"`     // Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	Link        *ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseLink      `json:"link,omitempty"`        //
	Name        string                                                                            `json:"name,omitempty"`        // Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank        *int                                                                              `json:"rank,omitempty"`        // The rank(priority) in relation to other policy set. Lower rank is higher priority.
	ServiceName string                                                                            `json:"serviceName,omitempty"` // Policy set service identifier - Allowed Protocols,Server Sequence..
	State       string                                                                            `json:"state,omitempty"`       // The state that the policy set is in. A disabled policy set cannot be matched.
}

type ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseCondition struct {
	ConditionType       string                                                                                               `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                               `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                               `json:"id,omitempty"`                  //
	Name                string                                                                                               `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                               `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                               `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                               `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                               `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                               `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                             `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                             `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseConditionChildren struct {
	ConditionType string                                                                                        `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                         `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySetResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetResetHitCountsDeviceAdminPolicySets struct {
	Message string `json:"message,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByID struct {
	Response *ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponse `json:"response,omitempty"` // Policy set structure
	Version  string                                                                    `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponse struct {
	Condition   *ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseCondition `json:"condition,omitempty"`   //
	Default     *bool                                                                              `json:"default,omitempty"`     // Flag which indicates if this policy set is the default one
	Description string                                                                             `json:"description,omitempty"` // The description for the policy set
	HitCounts   *int                                                                               `json:"hitCounts,omitempty"`   // The amount of times the policy was matched
	ID          string                                                                             `json:"id,omitempty"`          // Identifier for the policy set
	IsProxy     *bool                                                                              `json:"isProxy,omitempty"`     // Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	Link        *ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseLink      `json:"link,omitempty"`        //
	Name        string                                                                             `json:"name,omitempty"`        // Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank        *int                                                                               `json:"rank,omitempty"`        // The rank(priority) in relation to other policy set. Lower rank is higher priority.
	ServiceName string                                                                             `json:"serviceName,omitempty"` // Policy set service identifier - Allowed Protocols,Server Sequence..
	State       string                                                                             `json:"state,omitempty"`       // The state that the policy set is in. A disabled policy set cannot be matched.
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseCondition struct {
	ConditionType       string                                                                                                `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                 `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                `json:"id,omitempty"`                  //
	Name                string                                                                                                `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                              `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                              `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionChildren struct {
	ConditionType string                                                                                         `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                          `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByID struct {
	Response *ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponse `json:"response,omitempty"` // Policy set structure
	Version  string                                                                       `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponse struct {
	Condition   *ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseCondition `json:"condition,omitempty"`   //
	Default     *bool                                                                                 `json:"default,omitempty"`     // Flag which indicates if this policy set is the default one
	Description string                                                                                `json:"description,omitempty"` // The description for the policy set
	HitCounts   *int                                                                                  `json:"hitCounts,omitempty"`   // The amount of times the policy was matched
	ID          string                                                                                `json:"id,omitempty"`          // Identifier for the policy set
	IsProxy     *bool                                                                                 `json:"isProxy,omitempty"`     // Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	Link        *ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseLink      `json:"link,omitempty"`        //
	Name        string                                                                                `json:"name,omitempty"`        // Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank        *int                                                                                  `json:"rank,omitempty"`        // The rank(priority) in relation to other policy set. Lower rank is higher priority.
	ServiceName string                                                                                `json:"serviceName,omitempty"` // Policy set service identifier - Allowed Protocols,Server Sequence..
	State       string                                                                                `json:"state,omitempty"`       // The state that the policy set is in. A disabled policy set cannot be matched.
}

type ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseCondition struct {
	ConditionType       string                                                                                                   `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                    `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                   `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                   `json:"id,omitempty"`                  //
	Name                string                                                                                                   `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                   `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                   `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                   `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                   `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                   `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                 `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                 `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseConditionChildren struct {
	ConditionType string                                                                                            `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                             `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationPolicySetDeleteDeviceAdminPolicySetByID struct {
	ID string `json:"id,omitempty"` //
}

type RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySet struct {
	Condition   *RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetCondition `json:"condition,omitempty"`   //
	Default     *bool                                                                    `json:"default,omitempty"`     // Flag which indicates if this policy set is the default one
	Description string                                                                   `json:"description,omitempty"` // The description for the policy set
	HitCounts   *int                                                                     `json:"hitCounts,omitempty"`   // The amount of times the policy was matched
	ID          string                                                                   `json:"id,omitempty"`          // Identifier for the policy set
	IsProxy     *bool                                                                    `json:"isProxy,omitempty"`     // Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	Link        *RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetLink      `json:"link,omitempty"`        //
	Name        string                                                                   `json:"name,omitempty"`        // Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank        *int                                                                     `json:"rank,omitempty"`        // The rank(priority) in relation to other policy set. Lower rank is higher priority.
	ServiceName string                                                                   `json:"serviceName,omitempty"` // Policy set service identifier - Allowed Protocols,Server Sequence..
	State       string                                                                   `json:"state,omitempty"`       // The state that the policy set is in. A disabled policy set cannot be matched.
}

type RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetCondition struct {
	ConditionType       string                                                                                      `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                       `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                      `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                      `json:"id,omitempty"`                  //
	Name                string                                                                                      `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                      `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                      `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                      `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                      `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                      `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                    `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                    `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetConditionChildren struct {
	ConditionType string                                                                               `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySetLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByID struct {
	Condition   *RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDCondition `json:"condition,omitempty"`   //
	Default     *bool                                                                        `json:"default,omitempty"`     // Flag which indicates if this policy set is the default one
	Description string                                                                       `json:"description,omitempty"` // The description for the policy set
	HitCounts   *int                                                                         `json:"hitCounts,omitempty"`   // The amount of times the policy was matched
	ID          string                                                                       `json:"id,omitempty"`          // Identifier for the policy set
	IsProxy     *bool                                                                        `json:"isProxy,omitempty"`     // Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	Link        *RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDLink      `json:"link,omitempty"`        //
	Name        string                                                                       `json:"name,omitempty"`        // Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank        *int                                                                         `json:"rank,omitempty"`        // The rank(priority) in relation to other policy set. Lower rank is higher priority.
	ServiceName string                                                                       `json:"serviceName,omitempty"` // Policy set service identifier - Allowed Protocols,Server Sequence..
	State       string                                                                       `json:"state,omitempty"`       // The state that the policy set is in. A disabled policy set cannot be matched.
}

type RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDCondition struct {
	ConditionType       string                                                                                          `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                           `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                          `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                          `json:"id,omitempty"`                  //
	Name                string                                                                                          `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                          `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                          `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                          `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                          `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                          `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                        `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                        `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDConditionChildren struct {
	ConditionType string                                                                                   `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                    `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

//GetDeviceAdminPolicySets Device Admin - List of policy sets.
/* Device Admin List of policy sets.

 */
func (s *DeviceAdministrationPolicySetService) GetDeviceAdminPolicySets() (*ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySets, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySets{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminPolicySets")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySets)
	return result, response, err

}

//GetDeviceAdminPolicySetByID Device Admin - Get policy set attributes.
/* Device Admin Get policy set attributes.

@param id id path parameter. Policy id
*/
func (s *DeviceAdministrationPolicySetService) GetDeviceAdminPolicySetByID(id string) (*ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminPolicySetById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByID)
	return result, response, err

}

//CreateDeviceAdminPolicySet Device Admin - Create a new policy set.
/* Device Admin Create a new policy set:

 Policy must include name , service identifier (either server sequence or allowed protocol) and a condition.

 Condition has hierarchical structure which define a set of condition for which policy could be match.

 Condition can be either reference to a stored Library condition, using model
ConditionReference
, or, dynamically built conditions which are not stored in the conditions Library, using models
ConditionAttributes, ConditionAndBlock, ConditionOrBlock
.



*/
func (s *DeviceAdministrationPolicySetService) CreateDeviceAdminPolicySet(requestDeviceAdministrationPolicySetCreateDeviceAdminPolicySet *RequestDeviceAdministrationPolicySetCreateDeviceAdminPolicySet) (*ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySet, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceAdministrationPolicySetCreateDeviceAdminPolicySet).
		SetResult(&ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySet{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySet{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceAdminPolicySet")
	}

	result := response.Result().(*ResponseDeviceAdministrationPolicySetCreateDeviceAdminPolicySet)
	return result, response, err

}

//ResetHitCountsDeviceAdminPolicySets Device Admin - Reset HitCount for PolicySets
/* Device Admin Reset HitCount for PolicySets

 */
func (s *DeviceAdministrationPolicySetService) ResetHitCountsDeviceAdminPolicySets() (*ResponseDeviceAdministrationPolicySetResetHitCountsDeviceAdminPolicySets, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/reset-hitcount"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationPolicySetResetHitCountsDeviceAdminPolicySets{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseDeviceAdministrationPolicySetResetHitCountsDeviceAdminPolicySets{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ResetHitCountsDeviceAdminPolicySets")
	}

	result := response.Result().(*ResponseDeviceAdministrationPolicySetResetHitCountsDeviceAdminPolicySets)
	return result, response, err

}

//UpdateDeviceAdminPolicySetByID Device Admin - Update a policy set.
/* Device Admin Update a policy set.

@param id id path parameter. Policy id
*/
func (s *DeviceAdministrationPolicySetService) UpdateDeviceAdminPolicySetByID(id string, requestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetById *RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByID) (*ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetById).
		SetResult(&ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByID{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceAdminPolicySetById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByID)
	return result, response, err

}

//DeleteDeviceAdminPolicySetByID Device Admin - Delete a policy set.
/* Device Admin Delete a policy set.

@param id id path parameter. Policy id
*/
func (s *DeviceAdministrationPolicySetService) DeleteDeviceAdminPolicySetByID(id string) (*ResponseDeviceAdministrationPolicySetDeleteDeviceAdminPolicySetByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationPolicySetDeleteDeviceAdminPolicySetByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteDeviceAdminPolicySetById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationPolicySetDeleteDeviceAdminPolicySetByID)
	return result, response, err

}
