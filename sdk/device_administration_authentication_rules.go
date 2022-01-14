package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type DeviceAdministrationAuthenticationRulesService service

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRules struct {
	Response *[]ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponse `json:"response,omitempty"` //
	Version  string                                                                                      `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponse struct {
	IDentitySourceName string                                                                                        `json:"identitySourceName,omitempty"` // Identity source name from the identity stores
	IfAuthFail         string                                                                                        `json:"ifAuthFail,omitempty"`         // Action to perform when authentication fails such as Bad credentials, disabled user and so on
	IfProcessFail      string                                                                                        `json:"ifProcessFail,omitempty"`      // Action to perform when ISE is uanble to access the identity database
	IfUserNotFound     string                                                                                        `json:"ifUserNotFound,omitempty"`     // Action to perform when user is not found in any of identity stores
	Link               *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseLink `json:"link,omitempty"`               //
	Rule               *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRule `json:"rule,omitempty"`               // Common attributes in rule authentication/authorization
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRule struct {
	Condition *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                  `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                   `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                 `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                 `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                   `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                 `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleCondition struct {
	ConditionType       string                                                                                                                    `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                     `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                    `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                    `json:"id,omitempty"`                  //
	Name                string                                                                                                                    `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                    `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                                    `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                    `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                    `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                    `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                  `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                  `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionChildren struct {
	ConditionType string                                                                                                             `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                              `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRule struct {
	Response *ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponse `json:"response,omitempty"` // Rule for authentication in Network Access/Device Admin
	Version  string                                                                                      `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponse struct {
	IDentitySourceName string                                                                                          `json:"identitySourceName,omitempty"` // Identity source name from the identity stores
	IfAuthFail         string                                                                                          `json:"ifAuthFail,omitempty"`         // Action to perform when authentication fails such as Bad credentials, disabled user and so on
	IfProcessFail      string                                                                                          `json:"ifProcessFail,omitempty"`      // Action to perform when ISE is uanble to access the identity database
	IfUserNotFound     string                                                                                          `json:"ifUserNotFound,omitempty"`     // Action to perform when user is not found in any of identity stores
	Link               *ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseLink `json:"link,omitempty"`               //
	Rule               *ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRule `json:"rule,omitempty"`               // Common attributes in rule authentication/authorization
}

type ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRule struct {
	Condition *ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                    `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                     `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                   `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                   `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                     `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                   `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRuleCondition struct {
	ConditionType       string                                                                                                                      `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                       `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                      `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                      `json:"id,omitempty"`                  //
	Name                string                                                                                                                      `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                      `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                                      `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                      `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                      `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                      `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                    `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                    `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRuleConditionChildren struct {
	ConditionType string                                                                                                               `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                                `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesResetHitCountsDeviceAdminAuthenticationRules struct {
	Message string `json:"message,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByID struct {
	Response *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponse `json:"response,omitempty"` // Rule for authentication in Network Access/Device Admin
	Version  string                                                                                       `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponse struct {
	IDentitySourceName string                                                                                           `json:"identitySourceName,omitempty"` // Identity source name from the identity stores
	IfAuthFail         string                                                                                           `json:"ifAuthFail,omitempty"`         // Action to perform when authentication fails such as Bad credentials, disabled user and so on
	IfProcessFail      string                                                                                           `json:"ifProcessFail,omitempty"`      // Action to perform when ISE is uanble to access the identity database
	IfUserNotFound     string                                                                                           `json:"ifUserNotFound,omitempty"`     // Action to perform when user is not found in any of identity stores
	Link               *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseLink `json:"link,omitempty"`               //
	Rule               *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRule `json:"rule,omitempty"`               // Common attributes in rule authentication/authorization
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRule struct {
	Condition *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                     `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                      `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                    `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                    `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                      `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                    `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleCondition struct {
	ConditionType       string                                                                                                                       `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                        `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                       `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                       `json:"id,omitempty"`                  //
	Name                string                                                                                                                       `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                       `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                                       `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                       `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                       `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                       `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                     `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                     `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionChildren struct {
	ConditionType string                                                                                                                `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                                 `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByID struct {
	Response *ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponse `json:"response,omitempty"` // Rule for authentication in Network Access/Device Admin
	Version  string                                                                                          `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponse struct {
	IDentitySourceName string                                                                                              `json:"identitySourceName,omitempty"` // Identity source name from the identity stores
	IfAuthFail         string                                                                                              `json:"ifAuthFail,omitempty"`         // Action to perform when authentication fails such as Bad credentials, disabled user and so on
	IfProcessFail      string                                                                                              `json:"ifProcessFail,omitempty"`      // Action to perform when ISE is uanble to access the identity database
	IfUserNotFound     string                                                                                              `json:"ifUserNotFound,omitempty"`     // Action to perform when user is not found in any of identity stores
	Link               *ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseLink `json:"link,omitempty"`               //
	Rule               *ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRule `json:"rule,omitempty"`               // Common attributes in rule authentication/authorization
}

type ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRule struct {
	Condition *ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                        `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                         `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                       `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                       `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                         `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                       `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRuleCondition struct {
	ConditionType       string                                                                                                                          `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                           `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                          `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                          `json:"id,omitempty"`                  //
	Name                string                                                                                                                          `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                          `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                                          `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                          `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                          `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                          `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                        `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                        `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRuleConditionChildren struct {
	ConditionType string                                                                                                                   `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                                    `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthenticationRulesDeleteDeviceAdminAuthenticationRuleByID struct {
	ID string `json:"id,omitempty"` //
}

type RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRule struct {
	IDentitySourceName string                                                                                 `json:"identitySourceName,omitempty"` // Identity source name from the identity stores
	IfAuthFail         string                                                                                 `json:"ifAuthFail,omitempty"`         // Action to perform when authentication fails such as Bad credentials, disabled user and so on
	IfProcessFail      string                                                                                 `json:"ifProcessFail,omitempty"`      // Action to perform when ISE is uanble to access the identity database
	IfUserNotFound     string                                                                                 `json:"ifUserNotFound,omitempty"`     // Action to perform when user is not found in any of identity stores
	Link               *RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleLink `json:"link,omitempty"`               //
	Rule               *RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRule `json:"rule,omitempty"`               // Common attributes in rule authentication/authorization
}

type RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRule struct {
	Condition *RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                           `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                            `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                          `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                          `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                            `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                          `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRuleCondition struct {
	ConditionType       string                                                                                                             `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                              `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                             `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                             `json:"id,omitempty"`                  //
	Name                string                                                                                                             `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                             `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                             `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                             `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                             `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                             `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                           `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                           `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRuleConditionChildren struct {
	ConditionType string                                                                                                      `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                       `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRuleRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByID struct {
	IDentitySourceName string                                                                                     `json:"identitySourceName,omitempty"` // Identity source name from the identity stores
	IfAuthFail         string                                                                                     `json:"ifAuthFail,omitempty"`         // Action to perform when authentication fails such as Bad credentials, disabled user and so on
	IfProcessFail      string                                                                                     `json:"ifProcessFail,omitempty"`      // Action to perform when ISE is uanble to access the identity database
	IfUserNotFound     string                                                                                     `json:"ifUserNotFound,omitempty"`     // Action to perform when user is not found in any of identity stores
	Link               *RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDLink `json:"link,omitempty"`               //
	Rule               *RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRule `json:"rule,omitempty"`               // Common attributes in rule authentication/authorization
}

type RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRule struct {
	Condition *RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                               `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                              `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                              `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                              `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRuleCondition struct {
	ConditionType       string                                                                                                                 `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                  `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                 `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                 `json:"id,omitempty"`                  //
	Name                string                                                                                                                 `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                 `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                                 `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                 `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                 `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                 `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                               `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                               `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRuleConditionChildren struct {
	ConditionType string                                                                                                          `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                           `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByIDRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

//GetDeviceAdminAuthenticationRules Device Admin - Get authentication rules.
/* Device Admin Get authentication rules.

@param policyID policyId path parameter. Policy id
*/
func (s *DeviceAdministrationAuthenticationRulesService) GetDeviceAdminAuthenticationRules(policyID string) (*ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRules, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/authentication"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminAuthenticationRules")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRules)
	return result, response, err

}

//GetDeviceAdminAuthenticationRuleByID Device Admin - Get rule attributes.
/* Device Admin Get rule attributes.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *DeviceAdministrationAuthenticationRulesService) GetDeviceAdminAuthenticationRuleByID(policyID string, id string) (*ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/authentication/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminAuthenticationRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByID)
	return result, response, err

}

//CreateDeviceAdminAuthenticationRule Device Admin - Create authentication rule.
/* Device Admin Create authentication rule:

 Rule must include name and condition.

 Condition has hierarchical structure which define a set of conditions for which authentication policy rule could be match.

 Condition can be either reference to a stored Library condition, using model
ConditionReference

 or dynamically built conditions which are not stored in the conditions Library, using models
ConditionAttributes, ConditionAndBlock, ConditionOrBlock
.


@param policyID policyId path parameter. Policy id
*/
func (s *DeviceAdministrationAuthenticationRulesService) CreateDeviceAdminAuthenticationRule(policyID string, requestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRule *RequestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRule) (*ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRule, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/authentication"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRule).
		SetResult(&ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRule{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRule{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceAdminAuthenticationRule")
	}

	result := response.Result().(*ResponseDeviceAdministrationAuthenticationRulesCreateDeviceAdminAuthenticationRule)
	return result, response, err

}

//ResetHitCountsDeviceAdminAuthenticationRules Device Admin - Reset HitCount for Authentication Rules
/* Device Admin Reset HitCount for Authentication Rules

@param policyID policyId path parameter. Policy id
*/
func (s *DeviceAdministrationAuthenticationRulesService) ResetHitCountsDeviceAdminAuthenticationRules(policyID string) (*ResponseDeviceAdministrationAuthenticationRulesResetHitCountsDeviceAdminAuthenticationRules, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/authentication/reset-hitcount"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationAuthenticationRulesResetHitCountsDeviceAdminAuthenticationRules{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseDeviceAdministrationAuthenticationRulesResetHitCountsDeviceAdminAuthenticationRules{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ResetHitCountsDeviceAdminAuthenticationRules")
	}

	result := response.Result().(*ResponseDeviceAdministrationAuthenticationRulesResetHitCountsDeviceAdminAuthenticationRules)
	return result, response, err

}

//UpdateDeviceAdminAuthenticationRuleByID Device Admin - - Update rule.
/* Device Admin Update rule.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *DeviceAdministrationAuthenticationRulesService) UpdateDeviceAdminAuthenticationRuleByID(policyID string, id string, requestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleById *RequestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByID) (*ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/authentication/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleById).
		SetResult(&ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByID{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceAdminAuthenticationRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationAuthenticationRulesUpdateDeviceAdminAuthenticationRuleByID)
	return result, response, err

}

//DeleteDeviceAdminAuthenticationRuleByID Device Admin - Delete rule.
/* Device Admin Delete rule.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *DeviceAdministrationAuthenticationRulesService) DeleteDeviceAdminAuthenticationRuleByID(policyID string, id string) (*ResponseDeviceAdministrationAuthenticationRulesDeleteDeviceAdminAuthenticationRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/authentication/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationAuthenticationRulesDeleteDeviceAdminAuthenticationRuleByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteDeviceAdminAuthenticationRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationAuthenticationRulesDeleteDeviceAdminAuthenticationRuleByID)
	return result, response, err

}
