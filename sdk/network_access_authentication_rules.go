package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type NetworkAccessAuthenticationRulesService service

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRules struct {
	Response *[]ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponse `json:"response,omitempty"` //
	Version  string                                                                                 `json:"version,omitempty"`  //
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponse struct {
	IDentitySourceName string                                                                                   `json:"identitySourceName,omitempty"` // Identity source name from the identity stores
	IfAuthFail         string                                                                                   `json:"ifAuthFail,omitempty"`         // Action to perform when authentication fails such as Bad credentials, disabled user and so on
	IfProcessFail      string                                                                                   `json:"ifProcessFail,omitempty"`      // Action to perform when ISE is uanble to access the identity database
	IfUserNotFound     string                                                                                   `json:"ifUserNotFound,omitempty"`     // Action to perform when user is not found in any of identity stores
	Link               *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseLink `json:"link,omitempty"`               //
	Rule               *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRule `json:"rule,omitempty"`               // Common attributes in rule authentication/authorization
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRule struct {
	Condition *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                             `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                              `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                            `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                            `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                              `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                            `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleCondition struct {
	ConditionType       string                                                                                                               `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                               `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                               `json:"id,omitempty"`                  //
	Name                string                                                                                                               `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                               `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                               `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                               `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                               `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                               `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                             `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                             `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionChildren struct {
	ConditionType string                                                                                                        `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                         `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRule struct {
	Response *ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponse `json:"response,omitempty"` // Rule for authentication in Network Access/Device Admin
	Version  string                                                                                 `json:"version,omitempty"`  //
}

type ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponse struct {
	IDentitySourceName string                                                                                     `json:"identitySourceName,omitempty"` // Identity source name from the identity stores
	IfAuthFail         string                                                                                     `json:"ifAuthFail,omitempty"`         // Action to perform when authentication fails such as Bad credentials, disabled user and so on
	IfProcessFail      string                                                                                     `json:"ifProcessFail,omitempty"`      // Action to perform when ISE is uanble to access the identity database
	IfUserNotFound     string                                                                                     `json:"ifUserNotFound,omitempty"`     // Action to perform when user is not found in any of identity stores
	Link               *ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseLink `json:"link,omitempty"`               //
	Rule               *ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRule `json:"rule,omitempty"`               // Common attributes in rule authentication/authorization
}

type ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRule struct {
	Condition *ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                               `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                              `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                              `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                              `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRuleCondition struct {
	ConditionType       string                                                                                                                 `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                  `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                 `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                 `json:"id,omitempty"`                  //
	Name                string                                                                                                                 `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                 `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                                 `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                 `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                 `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                 `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                               `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                               `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRuleConditionChildren struct {
	ConditionType string                                                                                                          `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                           `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesResetHitCountsNetworkAccessAuthenticationRules struct {
	Message string `json:"message,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByID struct {
	Response *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponse `json:"response,omitempty"` // Rule for authentication in Network Access/Device Admin
	Version  string                                                                                  `json:"version,omitempty"`  //
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponse struct {
	IDentitySourceName string                                                                                      `json:"identitySourceName,omitempty"` // Identity source name from the identity stores
	IfAuthFail         string                                                                                      `json:"ifAuthFail,omitempty"`         // Action to perform when authentication fails such as Bad credentials, disabled user and so on
	IfProcessFail      string                                                                                      `json:"ifProcessFail,omitempty"`      // Action to perform when ISE is uanble to access the identity database
	IfUserNotFound     string                                                                                      `json:"ifUserNotFound,omitempty"`     // Action to perform when user is not found in any of identity stores
	Link               *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseLink `json:"link,omitempty"`               //
	Rule               *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRule `json:"rule,omitempty"`               // Common attributes in rule authentication/authorization
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRule struct {
	Condition *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                 `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                               `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                               `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                 `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                               `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleCondition struct {
	ConditionType       string                                                                                                                  `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                   `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                  `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                  `json:"id,omitempty"`                  //
	Name                string                                                                                                                  `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                  `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                                  `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                  `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                  `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                  `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionChildren struct {
	ConditionType string                                                                                                           `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                            `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByID struct {
	Response *ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponse `json:"response,omitempty"` // Rule for authentication in Network Access/Device Admin
	Version  string                                                                                     `json:"version,omitempty"`  //
}

type ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponse struct {
	IDentitySourceName string                                                                                         `json:"identitySourceName,omitempty"` // Identity source name from the identity stores
	IfAuthFail         string                                                                                         `json:"ifAuthFail,omitempty"`         // Action to perform when authentication fails such as Bad credentials, disabled user and so on
	IfProcessFail      string                                                                                         `json:"ifProcessFail,omitempty"`      // Action to perform when ISE is uanble to access the identity database
	IfUserNotFound     string                                                                                         `json:"ifUserNotFound,omitempty"`     // Action to perform when user is not found in any of identity stores
	Link               *ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseLink `json:"link,omitempty"`               //
	Rule               *ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRule `json:"rule,omitempty"`               // Common attributes in rule authentication/authorization
}

type ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRule struct {
	Condition *ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                   `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                    `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                  `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                  `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                    `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                  `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRuleCondition struct {
	ConditionType       string                                                                                                                     `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                      `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                     `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                     `json:"id,omitempty"`                  //
	Name                string                                                                                                                     `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                     `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                                     `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                     `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                     `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                     `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                   `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                   `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRuleConditionChildren struct {
	ConditionType string                                                                                                              `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                               `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthenticationRulesDeleteNetworkAccessAuthenticationRuleByID struct {
	ID string `json:"id,omitempty"` //
}

type RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRule struct {
	IDentitySourceName string                                                                            `json:"identitySourceName,omitempty"` // Identity source name from the identity stores
	IfAuthFail         string                                                                            `json:"ifAuthFail,omitempty"`         // Action to perform when authentication fails such as Bad credentials, disabled user and so on
	IfProcessFail      string                                                                            `json:"ifProcessFail,omitempty"`      // Action to perform when ISE is uanble to access the identity database
	IfUserNotFound     string                                                                            `json:"ifUserNotFound,omitempty"`     // Action to perform when user is not found in any of identity stores
	Link               *RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleLink `json:"link,omitempty"`               //
	Rule               *RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRule `json:"rule,omitempty"`               // Common attributes in rule authentication/authorization
}

type RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRule struct {
	Condition *RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                      `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                       `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                     `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                     `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                       `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                     `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleCondition struct {
	ConditionType       string                                                                                                        `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                         `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                        `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                        `json:"id,omitempty"`                  //
	Name                string                                                                                                        `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                        `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                        `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                        `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                        `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                        `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                      `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                      `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionChildren struct {
	ConditionType string                                                                                                 `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                  `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByID struct {
	IDentitySourceName string                                                                                `json:"identitySourceName,omitempty"` // Identity source name from the identity stores
	IfAuthFail         string                                                                                `json:"ifAuthFail,omitempty"`         // Action to perform when authentication fails such as Bad credentials, disabled user and so on
	IfProcessFail      string                                                                                `json:"ifProcessFail,omitempty"`      // Action to perform when ISE is uanble to access the identity database
	IfUserNotFound     string                                                                                `json:"ifUserNotFound,omitempty"`     // Action to perform when user is not found in any of identity stores
	Link               *RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDLink `json:"link,omitempty"`               //
	Rule               *RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRule `json:"rule,omitempty"`               // Common attributes in rule authentication/authorization
}

type RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRule struct {
	Condition *RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                          `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                           `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                         `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                         `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                           `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                         `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleCondition struct {
	ConditionType       string                                                                                                            `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                             `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                            `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                            `json:"id,omitempty"`                  //
	Name                string                                                                                                            `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                            `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                            `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                            `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                            `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                            `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                          `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                          `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionChildren struct {
	ConditionType string                                                                                                     `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                      `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

//GetNetworkAccessAuthenticationRules Network Access - Get authentication rules.
/* Network Access Get authentication rules.

@param policyID policyId path parameter. Policy id
*/
func (s *NetworkAccessAuthenticationRulesService) GetNetworkAccessAuthenticationRules(policyID string) (*ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRules, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/authentication"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessAuthenticationRules")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRules)
	return result, response, err

}

//GetNetworkAccessAuthenticationRuleByID Network Access - Get rule attributes.
/* Network Access Get rule attributes.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *NetworkAccessAuthenticationRulesService) GetNetworkAccessAuthenticationRuleByID(policyID string, id string) (*ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/authentication/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessAuthenticationRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByID)
	return result, response, err

}

//CreateNetworkAccessAuthenticationRule Network Access - Create authentication rule.
/* Network Access Create authentication rule:

 Rule must include name and condition.

 Condition has hierarchical structure which define a set of conditions for which authentication policy rule could be match.

 Condition can be either reference to a stored Library condition, using model
ConditionReference

 or dynamically built conditions which are not stored in the conditions Library, using models
ConditionAttributes, ConditionAndBlock, ConditionOrBlock
.


@param policyID policyId path parameter. Policy id
*/
func (s *NetworkAccessAuthenticationRulesService) CreateNetworkAccessAuthenticationRule(policyID string, requestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRule *RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRule) (*ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRule, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/authentication"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRule).
		SetResult(&ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRule{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRule{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkAccessAuthenticationRule")
	}

	result := response.Result().(*ResponseNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRule)
	return result, response, err

}

//ResetHitCountsNetworkAccessAuthenticationRules Network Access - Reset HitCount for Authentication Rules
/* Network Access Reset HitCount for Authentication Rules

@param policyID policyId path parameter. Policy id
*/
func (s *NetworkAccessAuthenticationRulesService) ResetHitCountsNetworkAccessAuthenticationRules(policyID string) (*ResponseNetworkAccessAuthenticationRulesResetHitCountsNetworkAccessAuthenticationRules, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/authentication/reset-hitcount"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessAuthenticationRulesResetHitCountsNetworkAccessAuthenticationRules{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessAuthenticationRulesResetHitCountsNetworkAccessAuthenticationRules{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ResetHitCountsNetworkAccessAuthenticationRules")
	}

	result := response.Result().(*ResponseNetworkAccessAuthenticationRulesResetHitCountsNetworkAccessAuthenticationRules)
	return result, response, err

}

//UpdateNetworkAccessAuthenticationRuleByID Network Access - Update rule.
/* Network Access Update rule.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *NetworkAccessAuthenticationRulesService) UpdateNetworkAccessAuthenticationRuleByID(policyID string, id string, requestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleById *RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByID) (*ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/authentication/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleById).
		SetResult(&ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByID{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkAccessAuthenticationRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByID)
	return result, response, err

}

//DeleteNetworkAccessAuthenticationRuleByID Network Access - Delete rule.
/* Network Access Delete rule.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *NetworkAccessAuthenticationRulesService) DeleteNetworkAccessAuthenticationRuleByID(policyID string, id string) (*ResponseNetworkAccessAuthenticationRulesDeleteNetworkAccessAuthenticationRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/authentication/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessAuthenticationRulesDeleteNetworkAccessAuthenticationRuleByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteNetworkAccessAuthenticationRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessAuthenticationRulesDeleteNetworkAccessAuthenticationRuleByID)
	return result, response, err

}
