package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type NetworkAccessAuthorizationExceptionRulesService service

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRules struct {
	Response *[]ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponse `json:"response,omitempty"` //
	Version  string                                                                                         `json:"version,omitempty"`  //
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponse struct {
	Link          *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseLink `json:"link,omitempty"`          //
	Profile       []string                                                                                         `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                                           `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRule struct {
	Condition *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                     `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                      `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                    `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                    `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                      `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                    `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleCondition struct {
	ConditionType       string                                                                                                                       `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                        `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                       `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                       `json:"id,omitempty"`                  //
	Name                string                                                                                                                       `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                       `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                       `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                       `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                       `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                       `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                       `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                     `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                     `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionChildren struct {
	ConditionType string                                                                                                                `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                                 `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRule struct {
	Response *ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponse `json:"response,omitempty"` // Authorization rule for network access
	Version  string                                                                                         `json:"version,omitempty"`  //
}

type ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponse struct {
	Link          *ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseLink `json:"link,omitempty"`          //
	Profile       []string                                                                                           `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                                             `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRule struct {
	Condition *ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                       `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                        `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                      `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                      `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                        `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                      `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRuleCondition struct {
	ConditionType       string                                                                                                                         `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                          `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                         `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                         `json:"id,omitempty"`                  //
	Name                string                                                                                                                         `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                         `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                         `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                         `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                         `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                         `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                         `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                       `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                       `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRuleConditionChildren struct {
	ConditionType string                                                                                                                  `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                                   `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesResetHitCountsNetworkAccessLocalExceptions struct {
	Message string `json:"message,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByID struct {
	Response *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponse `json:"response,omitempty"` // Authorization rule for network access
	Version  string                                                                                          `json:"version,omitempty"`  //
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponse struct {
	Link          *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseLink `json:"link,omitempty"`          //
	Profile       []string                                                                                            `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                                              `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRule struct {
	Condition *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                        `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                         `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                       `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                       `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                         `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                       `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleCondition struct {
	ConditionType       string                                                                                                                          `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                           `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                          `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                          `json:"id,omitempty"`                  //
	Name                string                                                                                                                          `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                          `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                          `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                          `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                          `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                          `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                          `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                        `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                        `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionChildren struct {
	ConditionType string                                                                                                                   `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                                    `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByID struct {
	Response *ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponse `json:"response,omitempty"` // Authorization rule for network access
	Version  string                                                                                             `json:"version,omitempty"`  //
}

type ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponse struct {
	Link          *ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseLink `json:"link,omitempty"`          //
	Profile       []string                                                                                               `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                                                 `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRule struct {
	Condition *ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                           `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                            `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                          `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                          `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                            `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                          `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRuleCondition struct {
	ConditionType       string                                                                                                                             `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                              `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                             `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                             `json:"id,omitempty"`                  //
	Name                string                                                                                                                             `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                             `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                             `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                             `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                             `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                             `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                             `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                           `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                           `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRuleConditionChildren struct {
	ConditionType string                                                                                                                      `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                                       `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessAuthorizationExceptionRulesDeleteNetworkAccessLocalExceptionRuleByID struct {
	ID string `json:"id,omitempty"` //
}

type RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRule struct {
	Link          *RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleLink `json:"link,omitempty"`          //
	Profile       []string                                                                                  `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                                    `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRule struct {
	Condition *RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                              `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                               `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                             `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                             `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                               `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                             `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleCondition struct {
	ConditionType       string                                                                                                                `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                 `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                `json:"id,omitempty"`                  //
	Name                string                                                                                                                `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                              `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                              `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionChildren struct {
	ConditionType string                                                                                                         `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                          `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByID struct {
	Link          *RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDLink `json:"link,omitempty"`          //
	Profile       []string                                                                                      `json:"profile,omitempty"`       // The authorization profile/s
	Rule          *RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRule `json:"rule,omitempty"`          // Common attributes in rule authentication/authorization
	SecurityGroup string                                                                                        `json:"securityGroup,omitempty"` // Security group used in authorization policies
}

type RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRule struct {
	Condition *RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                  `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                   `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                 `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                 `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                   `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                 `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleCondition struct {
	ConditionType       string                                                                                                                    `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                     `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                    `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                    `json:"id,omitempty"`                  //
	Name                string                                                                                                                    `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                    `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                    `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                    `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                    `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                    `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                    `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                  `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                  `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionChildren struct {
	ConditionType string                                                                                                             `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                              `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

//GetNetworkAccessLocalExceptionRules Network Access - Get local exception rules.
/* Network Access Get local exception rules.

@param policyID policyId path parameter. Policy id
*/
func (s *NetworkAccessAuthorizationExceptionRulesService) GetNetworkAccessLocalExceptionRules(policyID string) (*ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRules, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/exception"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessLocalExceptionRules")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRules)
	return result, response, err

}

//GetNetworkAccessLocalExceptionRuleByID Network Access - Get local exception rule attributes.
/* Network Access Get local exception rule attributes.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *NetworkAccessAuthorizationExceptionRulesService) GetNetworkAccessLocalExceptionRuleByID(policyID string, id string) (*ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/exception/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessLocalExceptionRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByID)
	return result, response, err

}

//CreateNetworkAccessLocalExceptionRule Network Access - Create local authorization exception rule.
/* Network Access Create local authorization exception rule:



 Rule must include name and condition.


 Condition has hierarchical structure which define a set of conditions for which authoriztion policy rule could be match.


 Condition can be either reference to a stored Library condition, using model
ConditionReference


or dynamically built conditions which are not stored in the conditions Library, using models
ConditionAttributes, ConditionAndBlock, ConditionOrBlock
.




@param policyID policyId path parameter. Policy id
*/
func (s *NetworkAccessAuthorizationExceptionRulesService) CreateNetworkAccessLocalExceptionRule(policyID string, requestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRule *RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRule) (*ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRule, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/exception"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRule).
		SetResult(&ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRule{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRule{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkAccessLocalExceptionRule")
	}

	result := response.Result().(*ResponseNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRule)
	return result, response, err

}

//ResetHitCountsNetworkAccessLocalExceptions Network Access - Reset HitCount for local exceptions
/* Network Access Reset HitCount for local exceptions

@param policyID policyId path parameter. Policy id
*/
func (s *NetworkAccessAuthorizationExceptionRulesService) ResetHitCountsNetworkAccessLocalExceptions(policyID string) (*ResponseNetworkAccessAuthorizationExceptionRulesResetHitCountsNetworkAccessLocalExceptions, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/exception/reset-hitcount"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessAuthorizationExceptionRulesResetHitCountsNetworkAccessLocalExceptions{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessAuthorizationExceptionRulesResetHitCountsNetworkAccessLocalExceptions{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ResetHitCountsNetworkAccessLocalExceptions")
	}

	result := response.Result().(*ResponseNetworkAccessAuthorizationExceptionRulesResetHitCountsNetworkAccessLocalExceptions)
	return result, response, err

}

//UpdateNetworkAccessLocalExceptionRuleByID Network Access - Update local exception rule.
/* Network Access Update local exception rule.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *NetworkAccessAuthorizationExceptionRulesService) UpdateNetworkAccessLocalExceptionRuleByID(policyID string, id string, requestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleById *RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByID) (*ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/exception/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleById).
		SetResult(&ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByID{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkAccessLocalExceptionRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByID)
	return result, response, err

}

//DeleteNetworkAccessLocalExceptionRuleByID Network Access - Delete local exception rule.
/* Network Access Delete local exception rule.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *NetworkAccessAuthorizationExceptionRulesService) DeleteNetworkAccessLocalExceptionRuleByID(policyID string, id string) (*ResponseNetworkAccessAuthorizationExceptionRulesDeleteNetworkAccessLocalExceptionRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{policyId}/exception/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessAuthorizationExceptionRulesDeleteNetworkAccessLocalExceptionRuleByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteNetworkAccessLocalExceptionRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessAuthorizationExceptionRulesDeleteNetworkAccessLocalExceptionRuleByID)
	return result, response, err

}
