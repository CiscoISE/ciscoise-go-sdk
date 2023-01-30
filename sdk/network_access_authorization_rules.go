package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type NetworkAccessAuthorizationRulesService service

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRules struct {
	Response *[]ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponse `json:"response,omitempty"` //
	Version  string                                                                               `json:"version,omitempty"`  //
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponse struct {
	Link          *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseLink `json:"link,omitempty"`          //
	Profile       []string                                                                               `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                                 `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRule struct {
	Condition *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                           `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                            `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                          `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                          `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                            `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                          `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRuleCondition struct {
	ConditionType       string                                                                                                             `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                              `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRuleConditionLink                `json:"link,omitempty"`                //
	DictionaryName      string                                                                                                             `json:"dictionaryName,omitempty"`      // Dictionary Name
	AttributeName       string                                                                                                             `json:"attributeName,omitempty"`       // Atribute Name
	Operator            string                                                                                                             `json:"operator,omitempty"`            // Operator
	AttributeValue      string                                                                                                             `json:"attributeValue,omitempty"`      // Attibute Name
	Description         string                                                                                                             `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                             `json:"id,omitempty"`                  //
	Name                string                                                                                                             `json:"name,omitempty"`                // Condition name
	DictionaryValue     string                                                                                                             `json:"dictionaryValue,omitempty"`     // Dictionary value
	Children            *[]ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                           `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                           `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRuleConditionChildren struct {
	ConditionType  string                                                                                                      `json:"conditionType,omitempty"`  // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate       *bool                                                                                                       `json:"isNegate,omitempty"`       // Indicates whereas this condition is in negate mode
	Link           *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRuleConditionChildrenLink `json:"link,omitempty"`           //
	DictionaryName string                                                                                                      `json:"dictionaryName,omitempty"` // Dictionary Name
	AttributeName  string                                                                                                      `json:"attributeName,omitempty"`  // Atribute Name
	Operator       string                                                                                                      `json:"operator,omitempty"`       // Operator
	AttributeValue string                                                                                                      `json:"attributeValue,omitempty"` // Attibute Name
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRulesResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRule struct {
	Response *ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponse `json:"response,omitempty"` // Authorization rule for network access
	Version  string                                                                               `json:"version,omitempty"`  //
}

type ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponse struct {
	Link          *ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseLink `json:"link,omitempty"`          //
	Profile       []string                                                                                 `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                                   `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRule struct {
	Condition *ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                             `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                              `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                            `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                            `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                              `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                            `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRuleCondition struct {
	ConditionType       string                                                                                                               `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRuleConditionLink                `json:"link,omitempty"`                //
	DictionaryName      string                                                                                                               `json:"dictionaryName,omitempty"`      // Dictionary Name
	AttributeName       string                                                                                                               `json:"attributeName,omitempty"`       // Atribute Name
	Operator            string                                                                                                               `json:"operator,omitempty"`            // Operator
	AttributeValue      string                                                                                                               `json:"attributeValue,omitempty"`      // Attibute Name
	Description         string                                                                                                               `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                               `json:"id,omitempty"`                  //
	Name                string                                                                                                               `json:"name,omitempty"`                // Condition name
	DictionaryValue     string                                                                                                               `json:"dictionaryValue,omitempty"`     // Dictionary value
	Children            *[]ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                             `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                             `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRuleConditionChildren struct {
	ConditionType  string                                                                                                        `json:"conditionType,omitempty"`  // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate       *bool                                                                                                         `json:"isNegate,omitempty"`       // Indicates whereas this condition is in negate mode
	Link           *ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRuleConditionChildrenLink `json:"link,omitempty"`           //
	DictionaryName string                                                                                                        `json:"dictionaryName,omitempty"` // Dictionary Name
	AttributeName  string                                                                                                        `json:"attributeName,omitempty"`  // Atribute Name
	Operator       string                                                                                                        `json:"operator,omitempty"`       // Operator
	AttributeValue string                                                                                                        `json:"attributeValue,omitempty"` // Attibute Name
	ID             string                                                                                                        `json:"id,omitempty"`             // id
}

type ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesResetHitCountsNetworkAccessAuthorizationRules struct {
	Message string `json:"message,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByID struct {
	Response *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponse `json:"response,omitempty"` // Authorization rule for network access
	Version  string                                                                                `json:"version,omitempty"`  //
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponse struct {
	Link          *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseLink `json:"link,omitempty"`          //
	Profile       []string                                                                                  `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                                    `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRule struct {
	Condition *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                              `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                               `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                             `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                             `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                               `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                             `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRuleCondition struct {
	ConditionType       string                                                                                                                `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                 `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRuleConditionLink                `json:"link,omitempty"`                //
	DictionaryName      string                                                                                                                `json:"dictionaryName,omitempty"`      // Dictionary Name
	AttributeName       string                                                                                                                `json:"attributeName,omitempty"`       // Atribute Name
	Operator            string                                                                                                                `json:"operator,omitempty"`            // Operator
	AttributeValue      string                                                                                                                `json:"attributeValue,omitempty"`      // Attibute Name
	Description         string                                                                                                                `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                `json:"id,omitempty"`                  //
	Name                string                                                                                                                `json:"name,omitempty"`                // Condition name
	DictionaryValue     string                                                                                                                `json:"dictionaryValue,omitempty"`     // Dictionary value
	Children            *[]ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                              `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                              `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRuleConditionChildren struct {
	ConditionType  string                                                                                                         `json:"conditionType,omitempty"`  // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate       *bool                                                                                                          `json:"isNegate,omitempty"`       // Indicates whereas this condition is in negate mode
	Link           *ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRuleConditionChildrenLink `json:"link,omitempty"`           //
	DictionaryName string                                                                                                         `json:"dictionaryName,omitempty"` // Dictionary Name
	AttributeName  string                                                                                                         `json:"attributeName,omitempty"`  // Atribute Name
	Operator       string                                                                                                         `json:"operator,omitempty"`       // Operator
	AttributeValue string                                                                                                         `json:"attributeValue,omitempty"` // Attibute Name
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByID struct {
	Response *ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponse `json:"response,omitempty"` // Authorization rule for network access
	Version  string                                                                                   `json:"version,omitempty"`  //
}

type ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponse struct {
	Link          *ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseLink `json:"link,omitempty"`          //
	Profile       []string                                                                                     `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                                       `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRule struct {
	Condition *ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                 `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                  `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                  `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRuleCondition struct {
	ConditionType       string                                                                                                                   `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                    `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRuleConditionLink                `json:"link,omitempty"`                //
	DictionaryName      string                                                                                                                   `json:"dictionaryName,omitempty"`      // Dictionary Name
	AttributeName       string                                                                                                                   `json:"attributeName,omitempty"`       // Atribute Name
	Operator            string                                                                                                                   `json:"operator,omitempty"`            // Operator
	AttributeValue      string                                                                                                                   `json:"attributeValue,omitempty"`      // Attibute Name
	Description         string                                                                                                                   `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                   `json:"id,omitempty"`                  //
	Name                string                                                                                                                   `json:"name,omitempty"`                // Condition name
	DictionaryValue     string                                                                                                                   `json:"dictionaryValue,omitempty"`     // Dictionary value
	Children            *[]ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                 `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                 `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRuleConditionChildren struct {
	ConditionType  string                                                                                                            `json:"conditionType,omitempty"`  // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate       *bool                                                                                                             `json:"isNegate,omitempty"`       // Indicates whereas this condition is in negate mode
	Link           *ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRuleConditionChildrenLink `json:"link,omitempty"`           //
	DictionaryName string                                                                                                            `json:"dictionaryName,omitempty"` // Dictionary Name
	AttributeName  string                                                                                                            `json:"attributeName,omitempty"`  // Atribute Name
	Operator       string                                                                                                            `json:"operator,omitempty"`       // Operator
	AttributeValue string                                                                                                            `json:"attributeValue,omitempty"` // Attibute Name
}

type ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationRulesDeleteNetworkAccessAuthorizationRuleByID struct {
	ID string `json:"id,omitempty"` //
}

type RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRule struct {
	Link          *RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleLink `json:"link,omitempty"`          //
	Profile       []string                                                                        `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                          `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRule struct {
	Condition *RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                    `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                     `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                   `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                   `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                     `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                   `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRuleCondition struct {
	ConditionType       string                                                                                                      `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                       `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRuleConditionLink                `json:"link,omitempty"`                //
	DictionaryName      string                                                                                                      `json:"dictionaryName,omitempty"`      // Dictionary Name
	AttributeName       string                                                                                                      `json:"attributeName,omitempty"`       // Atribute Name
	Operator            string                                                                                                      `json:"operator,omitempty"`            // Operator
	AttributeValue      string                                                                                                      `json:"attributeValue,omitempty"`      // Attibute Name
	Description         string                                                                                                      `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                      `json:"id,omitempty"`                  //
	Name                string                                                                                                      `json:"name,omitempty"`                // Condition name
	DictionaryValue     string                                                                                                      `json:"dictionaryValue,omitempty"`     // Dictionary value
	Children            *[]RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                    `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                    `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRuleConditionChildren struct {
	ConditionType  string                                                                                               `json:"conditionType,omitempty"`  // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate       *bool                                                                                                `json:"isNegate,omitempty"`       // Indicates whereas this condition is in negate mode
	Link           *RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRuleConditionChildrenLink `json:"link,omitempty"`           //
	DictionaryName string                                                                                               `json:"dictionaryName,omitempty"` // Dictionary Name
	AttributeName  string                                                                                               `json:"attributeName,omitempty"`  // Atribute Name
	Operator       string                                                                                               `json:"operator,omitempty"`       // Operator
	AttributeValue string                                                                                               `json:"attributeValue,omitempty"` // Attibute Name
	ID             string                                                                                               `json:"id,omitempty"`             // id
}

type RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRuleRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByID struct {
	Link          *RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDLink `json:"link,omitempty"`          //
	Profile       []string                                                                            `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                              `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRule struct {
	Condition *RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                        `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                         `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                       `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                       `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                         `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                       `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRuleCondition struct {
	ConditionType       string                                                                                                          `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                           `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRuleConditionLink                `json:"link,omitempty"`                //
	DictionaryName      string                                                                                                          `json:"dictionaryName,omitempty"`      // Dictionary Name
	AttributeName       string                                                                                                          `json:"attributeName,omitempty"`       // Atribute Name
	Operator            string                                                                                                          `json:"operator,omitempty"`            // Operator
	AttributeValue      string                                                                                                          `json:"attributeValue,omitempty"`      // Attibute Name
	Description         string                                                                                                          `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                          `json:"id,omitempty"`                  //
	Name                string                                                                                                          `json:"name,omitempty"`                // Condition name
	DictionaryValue     string                                                                                                          `json:"dictionaryValue,omitempty"`     // Dictionary value
	Children            *[]RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                        `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                        `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRuleConditionChildren struct {
	ConditionType  string                                                                                                   `json:"conditionType,omitempty"`  // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate       *bool                                                                                                    `json:"isNegate,omitempty"`       // Indicates whereas this condition is in negate mode
	Link           *RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRuleConditionChildrenLink `json:"link,omitempty"`           //
	DictionaryName string                                                                                                   `json:"dictionaryName,omitempty"` // Dictionary Name
	AttributeName  string                                                                                                   `json:"attributeName,omitempty"`  // Atribute Name
	Operator       string                                                                                                   `json:"operator,omitempty"`       // Operator
	AttributeValue string                                                                                                   `json:"attributeValue,omitempty"` // Attibute Name
}

type RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByIDRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

//GetNetworkAccessAuthorizationRules Network Access - Get authorization rules.
/* Network Access Get authorization rules.

@param policyID policyId path parameter. Policy id
*/
func (s *NetworkAccessAuthorizationRulesService) GetNetworkAccessAuthorizationRules(policyID string) (*ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRules, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/authorization"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessAuthorizationRules")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRules)
	return result, response, err

}

//GetNetworkAccessAuthorizationRuleByID Network Access - Get authorization rule attributes.
/* Network Access Get authorization rule attributes.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *NetworkAccessAuthorizationRulesService) GetNetworkAccessAuthorizationRuleByID(policyID string, id string) (*ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/authorization/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessAuthorizationRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByID)
	return result, response, err

}

//CreateNetworkAccessAuthorizationRule Network Access - Create authorization rule.
/* Network Access Create authorization rule:

 Rule must include name and condition.

 Condition has hierarchical structure which define a set of conditions for which authoriztion policy rule could be match.

 Condition can be either reference to a stored Library condition, using model
ConditionReference

 or dynamically built conditions which are not stored in the conditions Library, using models
ConditionAttributes, ConditionAndBlock, ConditionOrBlock
.


@param policyID policyId path parameter. Policy id
*/
func (s *NetworkAccessAuthorizationRulesService) CreateNetworkAccessAuthorizationRule(policyID string, requestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRule *RequestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRule) (*ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRule, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/authorization"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRule).
		SetResult(&ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRule{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRule{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkAccessAuthorizationRule")
	}

	result := response.Result().(*ResponseNetworkAccessAuthorizationRulesCreateNetworkAccessAuthorizationRule)
	return result, response, err

}

//ResetHitCountsNetworkAccessAuthorizationRules Network Access - Reset HitCount for Authorization Rules
/* Network Access Reset HitCount for Authorization Rules

@param policyID policyId path parameter. Policy id
*/
func (s *NetworkAccessAuthorizationRulesService) ResetHitCountsNetworkAccessAuthorizationRules(policyID string) (*ResponseNetworkAccessAuthorizationRulesResetHitCountsNetworkAccessAuthorizationRules, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/authorization/reset-hitcount"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessAuthorizationRulesResetHitCountsNetworkAccessAuthorizationRules{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessAuthorizationRulesResetHitCountsNetworkAccessAuthorizationRules{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ResetHitCountsNetworkAccessAuthorizationRules")
	}

	result := response.Result().(*ResponseNetworkAccessAuthorizationRulesResetHitCountsNetworkAccessAuthorizationRules)
	return result, response, err

}

//UpdateNetworkAccessAuthorizationRuleByID Network Access - Update authorization rule.
/* Network Access Update authorization rule.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *NetworkAccessAuthorizationRulesService) UpdateNetworkAccessAuthorizationRuleByID(policyID string, id string, requestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleById *RequestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByID) (*ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/authorization/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleById).
		SetResult(&ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByID{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkAccessAuthorizationRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessAuthorizationRulesUpdateNetworkAccessAuthorizationRuleByID)
	return result, response, err

}

//DeleteNetworkAccessAuthorizationRuleByID Network Access - Delete authorization rule.
/* Network Access Delete authorization rule.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *NetworkAccessAuthorizationRulesService) DeleteNetworkAccessAuthorizationRuleByID(policyID string, id string) (*ResponseNetworkAccessAuthorizationRulesDeleteNetworkAccessAuthorizationRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/authorization/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessAuthorizationRulesDeleteNetworkAccessAuthorizationRuleByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteNetworkAccessAuthorizationRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessAuthorizationRulesDeleteNetworkAccessAuthorizationRuleByID)
	return result, response, err

}
