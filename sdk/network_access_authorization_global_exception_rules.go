package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type NetworkAccessAuthorizationGlobalExceptionRulesService service

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRules struct {
	Response *[]ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponse `json:"response,omitempty"` //
	Version  string                                                                                                         `json:"version,omitempty"`  //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponse struct {
	Link          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseLink `json:"link,omitempty"`          //
	Profile       []string                                                                                                         `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                                                           `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRule struct {
	Condition *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                                     `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                                      `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                                    `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                                    `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                                      `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                                    `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleCondition struct {
	ConditionType       string                                                                                                                                       `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                                        `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                                       `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                                       `json:"id,omitempty"`                  //
	Name                string                                                                                                                                       `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                                       `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                                                       `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                                       `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                                       `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                                       `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                                     `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                                     `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionChildren struct {
	ConditionType string                                                                                                                                `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                                                 `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRule struct {
	Response *ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponse `json:"response,omitempty"` // Authorization rule for network access
	Version  string                                                                                                         `json:"version,omitempty"`  //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponse struct {
	Link          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseLink `json:"link,omitempty"`          //
	Profile       []string                                                                                                           `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                                                             `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRule struct {
	Condition *ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                                       `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                                        `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                                      `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                                      `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                                        `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                                      `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRuleCondition struct {
	ConditionType       string                                                                                                                                         `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                                          `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                                         `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                                         `json:"id,omitempty"`                  //
	Name                string                                                                                                                                         `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                                         `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                                                         `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                                         `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                                         `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                                         `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                                       `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                                       `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRuleConditionChildren struct {
	ConditionType string                                                                                                                                  `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                                                   `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesResetHitCountsNetworkAccessGlobalExceptions struct {
	Message string `json:"message,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByID struct {
	Response *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponse `json:"response,omitempty"` // Authorization rule for network access
	Version  string                                                                                                          `json:"version,omitempty"`  //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponse struct {
	Link          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseLink `json:"link,omitempty"`          //
	Profile       []string                                                                                                            `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                                                              `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRule struct {
	Condition *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                                        `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                                         `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                                       `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                                       `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                                         `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                                       `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleCondition struct {
	ConditionType       string                                                                                                                                          `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                                           `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                                          `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                                          `json:"id,omitempty"`                  //
	Name                string                                                                                                                                          `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                                          `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                                                          `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                                          `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                                          `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                                          `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                                        `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                                        `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionChildren struct {
	ConditionType string                                                                                                                                   `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                                                    `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByID struct {
	Response *ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponse `json:"response,omitempty"` // Authorization rule for network access
	Version  string                                                                                                             `json:"version,omitempty"`  //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponse struct {
	Link          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseLink `json:"link,omitempty"`          //
	Profile       []string                                                                                                               `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                                                                 `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRule struct {
	Condition *ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                                           `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                                            `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                                          `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                                          `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                                            `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                                          `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleCondition struct {
	ConditionType       string                                                                                                                                             `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                                              `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                                             `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                                             `json:"id,omitempty"`                  //
	Name                string                                                                                                                                             `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                                             `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                                                             `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                                             `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                                             `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                                             `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                                           `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                                           `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionChildren struct {
	ConditionType string                                                                                                                                      `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                                                       `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationGlobalExceptionRulesDeleteNetworkAccessPolicySetGlobalExceptionRuleByID struct {
	ID string `json:"id,omitempty"` //
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRule struct {
	Link          *RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleLink `json:"link,omitempty"`          //
	Profile       []string                                                                                                  `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                                                    `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRule struct {
	Condition *RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                              `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                               `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                             `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                             `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                               `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                             `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRuleCondition struct {
	ConditionType       string                                                                                                                                `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                                 `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                                `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                                `json:"id,omitempty"`                  //
	Name                string                                                                                                                                `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                                `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                                                `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                                `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                                `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                                `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                              `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                              `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRuleConditionChildren struct {
	ConditionType string                                                                                                                         `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                                          `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRuleRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByID struct {
	Link          *RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDLink `json:"link,omitempty"`          //
	Profile       []string                                                                                                      `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                                                        `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRule struct {
	Condition *RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                                  `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                                   `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                                 `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                                 `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                                   `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                                 `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRuleCondition struct {
	ConditionType       string                                                                                                                                    `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                                     `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                                    `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                                    `json:"id,omitempty"`                  //
	Name                string                                                                                                                                    `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                                    `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeValue      string                                                                                                                                    `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                                    `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                                    `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                                    `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                                  `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                                  `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRuleConditionChildren struct {
	ConditionType string                                                                                                                             `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                                              `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByIDRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

//GetNetworkAccessPolicySetGlobalExceptionRules Network Access - Get global execption rules.
/* Network Access Get global execption rules.

 */
func (s *NetworkAccessAuthorizationGlobalExceptionRulesService) GetNetworkAccessPolicySetGlobalExceptionRules() (*ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRules, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/global-exception"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessPolicySetGlobalExceptionRules")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRules)
	return result, response, err

}

//GetNetworkAccessPolicySetGlobalExceptionRuleByID Network Access - Get global exception rule attributes.
/* Network Access Get global exception rule attributes.

@param id id path parameter. Rule id
*/
func (s *NetworkAccessAuthorizationGlobalExceptionRulesService) GetNetworkAccessPolicySetGlobalExceptionRuleByID(id string) (*ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/global-exception/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessPolicySetGlobalExceptionRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByID)
	return result, response, err

}

//CreateNetworkAccessPolicySetGlobalExceptionRule Network Access - Create global exception authorization rule.
/* Network Access Create global exception authorization rule:

 Rule must include name and condition.

 Condition has hierarchical structure which define a set of conditions for which authoriztion policy rule could be match.

 Condition can be either reference to a stored Library condition, using model
ConditionReference

 or dynamically built conditions which are not stored in the conditions Library, using models
ConditionAttributes, ConditionAndBlock, ConditionOrBlock
.


*/
func (s *NetworkAccessAuthorizationGlobalExceptionRulesService) CreateNetworkAccessPolicySetGlobalExceptionRule(requestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRule *RequestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRule) (*ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRule, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/global-exception"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRule).
		SetResult(&ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRule{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRule{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkAccessPolicySetGlobalExceptionRule")
	}

	result := response.Result().(*ResponseNetworkAccessAuthorizationGlobalExceptionRulesCreateNetworkAccessPolicySetGlobalExceptionRule)
	return result, response, err

}

//ResetHitCountsNetworkAccessGlobalExceptions Network Access - Reset HitCount for Global Exceptions
/* Network Access Reset HitCount for Global Exceptions

 */
func (s *NetworkAccessAuthorizationGlobalExceptionRulesService) ResetHitCountsNetworkAccessGlobalExceptions() (*ResponseNetworkAccessAuthorizationGlobalExceptionRulesResetHitCountsNetworkAccessGlobalExceptions, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/global-exception/reset-hitcount"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessAuthorizationGlobalExceptionRulesResetHitCountsNetworkAccessGlobalExceptions{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessAuthorizationGlobalExceptionRulesResetHitCountsNetworkAccessGlobalExceptions{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ResetHitCountsNetworkAccessGlobalExceptions")
	}

	result := response.Result().(*ResponseNetworkAccessAuthorizationGlobalExceptionRulesResetHitCountsNetworkAccessGlobalExceptions)
	return result, response, err

}

//UpdateNetworkAccessPolicySetGlobalExceptionRuleByID Network Access - Update global exception authorization rule.
/* Network Access Update global exception authorization rule.

@param id id path parameter. Rule id
*/
func (s *NetworkAccessAuthorizationGlobalExceptionRulesService) UpdateNetworkAccessPolicySetGlobalExceptionRuleByID(id string, requestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleById *RequestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByID) (*ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/global-exception/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleById).
		SetResult(&ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByID{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkAccessPolicySetGlobalExceptionRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessAuthorizationGlobalExceptionRulesUpdateNetworkAccessPolicySetGlobalExceptionRuleByID)
	return result, response, err

}

//DeleteNetworkAccessPolicySetGlobalExceptionRuleByID Network Access - Delete global exception authorization rule.
/* Network Access Delete global exception authorization rule.

@param id id path parameter. Rule id
*/
func (s *NetworkAccessAuthorizationGlobalExceptionRulesService) DeleteNetworkAccessPolicySetGlobalExceptionRuleByID(id string) (*ResponseNetworkAccessAuthorizationGlobalExceptionRulesDeleteNetworkAccessPolicySetGlobalExceptionRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/global-exception/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessAuthorizationGlobalExceptionRulesDeleteNetworkAccessPolicySetGlobalExceptionRuleByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteNetworkAccessPolicySetGlobalExceptionRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessAuthorizationGlobalExceptionRulesDeleteNetworkAccessPolicySetGlobalExceptionRuleByID)
	return result, response, err

}
