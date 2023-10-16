package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type DeviceAdministrationAuthorizationExceptionRulesService service

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRules struct {
	Response *[]ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponse `json:"response,omitempty"` //
	Version  string                                                                                              `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponse struct {
	Commands []string                                                                                              `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseLink `json:"link,omitempty"`     //
	Profile  string                                                                                                `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRule struct {
	Condition *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                          `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                           `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                         `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                         `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                           `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                         `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRuleCondition struct {
	ConditionType       string                                                                                                                            `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                             `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                            `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                            `json:"id,omitempty"`                  //
	Name                string                                                                                                                            `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                            `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                            `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                            `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                            `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                            `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                            `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                          `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                          `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRuleConditionChildren struct {
	ConditionType  string                                                                                                                     `json:"conditionType,omitempty"`  // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate       *bool                                                                                                                      `json:"isNegate,omitempty"`       // Indicates whereas this condition is in negate mode
	Link           *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRuleConditionChildrenLink `json:"link,omitempty"`           //
	DictionaryName string                                                                                                                     `json:"dictionaryName,omitempty"` // Dictionary Name
	AttributeName  string                                                                                                                     `json:"attributeName,omitempty"`  // Atribute Name
	Operator       string                                                                                                                     `json:"operator,omitempty"`       // Operator
	AttributeValue string                                                                                                                     `json:"attributeValue,omitempty"` // Attibute Name
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRule struct {
	Response *ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponse `json:"response,omitempty"` // Authorization rule for device admin
	Version  string                                                                                              `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponse struct {
	Commands []string                                                                                                `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     *ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseLink `json:"link,omitempty"`     //
	Profile  string                                                                                                  `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     *ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRule struct {
	Condition *ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                            `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                             `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                           `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                           `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                             `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                           `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRuleCondition struct {
	ConditionType       string                                                                                                                              `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                               `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                              `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                              `json:"id,omitempty"`                  //
	Name                string                                                                                                                              `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                              `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                              `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                              `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                              `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                              `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                              `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                            `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                            `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRuleConditionChildren struct {
	ConditionType  string                                                                                                                       `json:"conditionType,omitempty"`  // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate       *bool                                                                                                                        `json:"isNegate,omitempty"`       // Indicates whereas this condition is in negate mode
	Link           *ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRuleConditionChildrenLink `json:"link,omitempty"`           //
	DictionaryName string                                                                                                                       `json:"dictionaryName,omitempty"` // Dictionary Name
	AttributeName  string                                                                                                                       `json:"attributeName,omitempty"`  // Atribute Name
	Operator       string                                                                                                                       `json:"operator,omitempty"`       // Operator
	AttributeValue string                                                                                                                       `json:"attributeValue,omitempty"` // Attibute Name
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesResetHitCountsDeviceAdminLocalExceptions struct {
	Message string `json:"message,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByID struct {
	Response *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponse `json:"response,omitempty"` // Authorization rule for device admin
	Version  string                                                                                               `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponse struct {
	Commands []string                                                                                                 `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseLink `json:"link,omitempty"`     //
	Profile  string                                                                                                   `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRule struct {
	Condition *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                             `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                              `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                            `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                            `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                              `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                            `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRuleCondition struct {
	ConditionType       string                                                                                                                               `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                                `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                               `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                               `json:"id,omitempty"`                  //
	Name                string                                                                                                                               `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                               `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                               `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                               `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                               `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                               `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                               `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                             `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                             `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRuleConditionChildren struct {
	ConditionType  string                                                                                                                        `json:"conditionType,omitempty"`  // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate       *bool                                                                                                                         `json:"isNegate,omitempty"`       // Indicates whereas this condition is in negate mode
	Link           *ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRuleConditionChildrenLink `json:"link,omitempty"`           //
	DictionaryName string                                                                                                                        `json:"dictionaryName,omitempty"` // Dictionary Name
	AttributeName  string                                                                                                                        `json:"attributeName,omitempty"`  // Atribute Name
	Operator       string                                                                                                                        `json:"operator,omitempty"`       // Operator
	AttributeValue string                                                                                                                        `json:"attributeValue,omitempty"` // Attibute Name
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByID struct {
	Response *ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponse `json:"response,omitempty"` // Authorization rule for device admin
	Version  string                                                                                                  `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponse struct {
	Commands []string                                                                                                    `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     *ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseLink `json:"link,omitempty"`     //
	Profile  string                                                                                                      `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     *ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRule struct {
	Condition *ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                                `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                                 `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                               `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                               `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                                 `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                               `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRuleCondition struct {
	ConditionType       string                                                                                                                                  `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                                   `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                                  `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                                  `json:"id,omitempty"`                  //
	Name                string                                                                                                                                  `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                                  `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                                  `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                                  `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                                  `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                                  `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                                  `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                                `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                                `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRuleConditionChildren struct {
	ConditionType  string                                                                                                                           `json:"conditionType,omitempty"`  // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate       *bool                                                                                                                            `json:"isNegate,omitempty"`       // Indicates whereas this condition is in negate mode
	Link           *ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRuleConditionChildrenLink `json:"link,omitempty"`           //
	DictionaryName string                                                                                                                           `json:"dictionaryName,omitempty"` // Dictionary Name
	AttributeName  string                                                                                                                           `json:"attributeName,omitempty"`  // Atribute Name
	Operator       string                                                                                                                           `json:"operator,omitempty"`       // Operator
	AttributeValue string                                                                                                                           `json:"attributeValue,omitempty"` // Attibute Name
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationExceptionRulesDeleteDeviceAdminLocalExceptionRuleByID struct {
	ID string `json:"id,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRule struct {
	Commands []string                                                                                       `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     *RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleLink `json:"link,omitempty"`     //
	Profile  string                                                                                         `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     *RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRule struct {
	Condition *RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                   `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                    `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                  `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                  `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                    `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                  `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleCondition struct {
	ConditionType       string                                                                                                                     `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                      `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                     `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                     `json:"id,omitempty"`                  //
	Name                string                                                                                                                     `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                     `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                     `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                     `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                     `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                     `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                     `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                   `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                   `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionChildren struct {
	ConditionType  string                                                                                                              `json:"conditionType,omitempty"`  // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate       *bool                                                                                                               `json:"isNegate,omitempty"`       // Indicates whereas this condition is in negate mode
	Link           *RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionChildrenLink `json:"link,omitempty"`           //
	DictionaryName string                                                                                                              `json:"dictionaryName,omitempty"` // Dictionary Name
	AttributeName  string                                                                                                              `json:"attributeName,omitempty"`  // Atribute Name
	Operator       string                                                                                                              `json:"operator,omitempty"`       // Operator
	AttributeValue string                                                                                                              `json:"attributeValue,omitempty"` // Attibute Name
}

type RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByID struct {
	Commands []string                                                                                           `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     *RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDLink `json:"link,omitempty"`     //
	Profile  string                                                                                             `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     *RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRule struct {
	Condition *RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                       `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                        `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                      `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                      `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                        `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                      `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleCondition struct {
	ConditionType       string                                                                                                                         `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                          `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                         `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                         `json:"id,omitempty"`                  //
	Name                string                                                                                                                         `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                         `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                         `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                         `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                         `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                         `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                         `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                       `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                       `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionChildren struct {
	ConditionType  string                                                                                                                  `json:"conditionType,omitempty"`  // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate       *bool                                                                                                                   `json:"isNegate,omitempty"`       // Indicates whereas this condition is in negate mode
	Link           *RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionChildrenLink `json:"link,omitempty"`           //
	DictionaryName string                                                                                                                  `json:"dictionaryName,omitempty"` // Dictionary Name
	AttributeName  string                                                                                                                  `json:"attributeName,omitempty"`  // Atribute Name
	Operator       string                                                                                                                  `json:"operator,omitempty"`       // Operator
	AttributeValue string                                                                                                                  `json:"attributeValue,omitempty"` // Attibute Name
}

type RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

//GetDeviceAdminLocalExceptionRules Device Admin - Get local exception rules.
/* Device Admin Get local exception rules.

@param policyID policyId path parameter. Policy id
*/
func (s *DeviceAdministrationAuthorizationExceptionRulesService) GetDeviceAdminLocalExceptionRules(policyID string) (*ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRules, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/exception"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminLocalExceptionRules")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRules)
	return result, response, err

}

//GetDeviceAdminLocalExceptionRuleByID Device Admin - Get local exception rule attributes.
/* Device Admin Get local exception rule attributes.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *DeviceAdministrationAuthorizationExceptionRulesService) GetDeviceAdminLocalExceptionRuleByID(policyID string, id string) (*ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/exception/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminLocalExceptionRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByID)
	return result, response, err

}

//CreateDeviceAdminLocalExceptionRule Device Admin - Create local authorization exception rule.
/* Device Admin Create local authorization exception rule.

@param policyID policyId path parameter. Policy id
*/
func (s *DeviceAdministrationAuthorizationExceptionRulesService) CreateDeviceAdminLocalExceptionRule(policyID string, requestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRule *RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRule) (*ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRule, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/exception"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRule).
		SetResult(&ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRule{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRule{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceAdminLocalExceptionRule")
	}

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRule)
	return result, response, err

}

//ResetHitCountsDeviceAdminLocalExceptions Device Admin - Reset HitCount for local exceptions
/* Device Admin Reset HitCount for local exceptions

@param policyID policyId path parameter. Policy id
*/
func (s *DeviceAdministrationAuthorizationExceptionRulesService) ResetHitCountsDeviceAdminLocalExceptions(policyID string) (*ResponseDeviceAdministrationAuthorizationExceptionRulesResetHitCountsDeviceAdminLocalExceptions, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/exception/reset-hitcount"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationAuthorizationExceptionRulesResetHitCountsDeviceAdminLocalExceptions{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseDeviceAdministrationAuthorizationExceptionRulesResetHitCountsDeviceAdminLocalExceptions{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ResetHitCountsDeviceAdminLocalExceptions")
	}

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationExceptionRulesResetHitCountsDeviceAdminLocalExceptions)
	return result, response, err

}

//UpdateDeviceAdminLocalExceptionRuleByID Device Admin - Update local exception rule.
/* Device Admin Update local exception rule.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *DeviceAdministrationAuthorizationExceptionRulesService) UpdateDeviceAdminLocalExceptionRuleByID(policyID string, id string, requestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleById *RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByID) (*ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/exception/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleById).
		SetResult(&ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByID{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceAdminLocalExceptionRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByID)
	return result, response, err

}

//DeleteDeviceAdminLocalExceptionRuleByID Device Admin - Delete local exception rule.
/* Device Admin Delete local exception rule.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *DeviceAdministrationAuthorizationExceptionRulesService) DeleteDeviceAdminLocalExceptionRuleByID(policyID string, id string) (*ResponseDeviceAdministrationAuthorizationExceptionRulesDeleteDeviceAdminLocalExceptionRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/exception/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationAuthorizationExceptionRulesDeleteDeviceAdminLocalExceptionRuleByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteDeviceAdminLocalExceptionRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationExceptionRulesDeleteDeviceAdminLocalExceptionRuleByID)
	return result, response, err

}
