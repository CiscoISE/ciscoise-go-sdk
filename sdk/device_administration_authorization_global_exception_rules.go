package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type DeviceAdministrationAuthorizationGlobalExceptionRulesService service

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRules struct {
	Response []ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponse `json:"response,omitempty"` //
	Version  string                                                                                                             `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponse struct {
	Commands []string                                                                                                             `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseLink `json:"link,omitempty"`     //
	Profile  string                                                                                                               `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseRule struct {
	Condition ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseRuleCondition `json:"condition,omitempty"` // <ul><li>Hierarchical structure which defines a set of conditions for which authentication or authorization policy rules could be matched.</li> <li>Logical operations(AND, OR) relationship between conditions are supported</li> <li>Each condition can have subconditions with relation to logical operations</li></ul>
	Default   bool                                                                                                                          `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts int                                                                                                                           `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                                        `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                                        `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      int                                                                                                                           `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                                        `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseRuleCondition struct {
	ConditionType string                                                                                                                            `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      bool                                                                                                                              `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseRuleConditionLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalException struct {
	Response ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponse `json:"response,omitempty"` // Authorization rule for device admin
	Version  string                                                                                                         `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponse struct {
	Commands []string                                                                                                           `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseLink `json:"link,omitempty"`     //
	Profile  string                                                                                                             `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRule struct {
	Condition ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRuleCondition `json:"condition,omitempty"` //
	Default   bool                                                                                                                        `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts int                                                                                                                         `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                                      `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                                      `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      int                                                                                                                         `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                                      `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRuleCondition struct {
	ConditionType       string                                                                                                                                         `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            bool                                                                                                                                           `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                                         `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                                         `json:"id,omitempty"`                  //
	Name                string                                                                                                                                         `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                                         `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                                         `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                                         `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                                         `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                                         `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                                         `json:"operator,omitempty"`            // Equality operator
	Children            []ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                                       `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                                       `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRuleConditionChildren struct {
	ConditionType string                                                                                                                                  `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      bool                                                                                                                                    `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesResetHitCountsDeviceAdminGlobalExceptions struct {
	Message string `json:"message,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleID struct {
	Response ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponse `json:"response,omitempty"` // Authorization rule for device admin
	Version  string                                                                                                              `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponse struct {
	Commands []string                                                                                                                `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseLink `json:"link,omitempty"`     //
	Profile  string                                                                                                                  `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRule struct {
	Condition ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleCondition `json:"condition,omitempty"` //
	Default   bool                                                                                                                             `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts int                                                                                                                              `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                                           `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                                           `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      int                                                                                                                              `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                                           `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleCondition struct {
	ConditionType       string                                                                                                                                              `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            bool                                                                                                                                                `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                                              `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                                              `json:"id,omitempty"`                  //
	Name                string                                                                                                                                              `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                                              `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                                              `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                                              `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                                              `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                                              `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                                              `json:"operator,omitempty"`            // Equality operator
	Children            []ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                                            `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                                            `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionChildren struct {
	ConditionType string                                                                                                                                       `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      bool                                                                                                                                         `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleID struct {
	Response ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponse `json:"response,omitempty"` // Authorization rule for device admin
	Version  string                                                                                                                 `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponse struct {
	Commands []string                                                                                                                   `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseLink `json:"link,omitempty"`     //
	Profile  string                                                                                                                     `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRule struct {
	Condition ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleCondition `json:"condition,omitempty"` //
	Default   bool                                                                                                                                `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts int                                                                                                                                 `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                                              `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                                              `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      int                                                                                                                                 `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                                              `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleCondition struct {
	ConditionType       string                                                                                                                                                 `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            bool                                                                                                                                                   `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                                                 `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                                                 `json:"id,omitempty"`                  //
	Name                string                                                                                                                                                 `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                                                 `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                                                 `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                                                 `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                                                 `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                                                 `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                                                 `json:"operator,omitempty"`            // Equality operator
	Children            []ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                                               `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                                               `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionChildren struct {
	ConditionType string                                                                                                                                          `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      bool                                                                                                                                            `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesDeleteDeviceAdminPolicySetGlobalExceptionByRuleID struct {
	ID string `json:"id,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalException struct {
	Commands []string                                                                                                   `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionLink `json:"link,omitempty"`     //
	Profile  string                                                                                                     `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRule struct {
	Condition *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                               `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                                `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                              `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                              `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                                `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                              `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleCondition struct {
	ConditionType       string                                                                                                                                 `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                                  `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                                 `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                                 `json:"id,omitempty"`                  //
	Name                string                                                                                                                                 `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                                 `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                                 `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                                 `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                                 `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                                 `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                                 `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                               `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                               `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionChildren struct {
	ConditionType string                                                                                                                          `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                                           `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleID struct {
	Commands []string                                                                                                           `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDLink `json:"link,omitempty"`     //
	Profile  string                                                                                                             `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRule struct {
	Condition *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                                                       `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                                                        `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                                      `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                                      `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                                                        `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                                      `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleCondition struct {
	ConditionType       string                                                                                                                                         `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                                          `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                                         `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                                         `json:"id,omitempty"`                  //
	Name                string                                                                                                                                         `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                                         `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                                         `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                                         `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                                         `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                                         `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                                         `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                                       `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                                       `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionChildren struct {
	ConditionType string                                                                                                                                  `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                                                   `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

//GetDeviceAdminPolicySetGlobalExceptionRules Device Admin - Get global execption rules.
/* Device Admin Get global execption rules.

 */
func (s *DeviceAdministrationAuthorizationGlobalExceptionRulesService) GetDeviceAdminPolicySetGlobalExceptionRules() (*ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRules, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/global-exception"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminPolicySetGlobalExceptionRules")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRules)
	return result, response, err

}

//GetDeviceAdminPolicySetGlobalExceptionByRuleID Device Admin - Get global exception rule attributes.
/* Device Admin Get global exception rule attribute

@param id id path parameter. Rule id
*/
func (s *DeviceAdministrationAuthorizationGlobalExceptionRulesService) GetDeviceAdminPolicySetGlobalExceptionByRuleID(id string) (*ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/global-exception/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminPolicySetGlobalExceptionByRuleId")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleID)
	return result, response, err

}

//CreateDeviceAdminPolicySetGlobalException Device Admin - Create global exception authorization rule.
/* Device Admin Create global exception authorization rule.

 */
func (s *DeviceAdministrationAuthorizationGlobalExceptionRulesService) CreateDeviceAdminPolicySetGlobalException(requestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalException *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalException) (*ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalException, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/global-exception"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalException).
		SetResult(&ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalException{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceAdminPolicySetGlobalException")
	}

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalException)
	return result, response, err

}

//ResetHitCountsDeviceAdminGlobalExceptions Device Admin - Reset HitCount for Global Exceptions
/* Device Admin Reset HitCount for Global Exceptions

 */
func (s *DeviceAdministrationAuthorizationGlobalExceptionRulesService) ResetHitCountsDeviceAdminGlobalExceptions() (*ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesResetHitCountsDeviceAdminGlobalExceptions, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/global-exception/reset-hitcount"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesResetHitCountsDeviceAdminGlobalExceptions{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ResetHitCountsDeviceAdminGlobalExceptions")
	}

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesResetHitCountsDeviceAdminGlobalExceptions)
	return result, response, err

}

//UpdateDeviceAdminPolicySetGlobalExceptionByRuleID Device Admin - Update global exception authorization rule.
/* Device Admin Update global exception authorization rule.

@param id id path parameter. Rule id
*/
func (s *DeviceAdministrationAuthorizationGlobalExceptionRulesService) UpdateDeviceAdminPolicySetGlobalExceptionByRuleID(id string, requestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleId *RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleID) (*ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/global-exception/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleId).
		SetResult(&ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceAdminPolicySetGlobalExceptionByRuleId")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleID)
	return result, response, err

}

//DeleteDeviceAdminPolicySetGlobalExceptionByRuleID Device Admin - Delete global exception authorization rule.
/* Device Admin Delete global exception authorization rule.

@param id id path parameter. Rule id
*/
func (s *DeviceAdministrationAuthorizationGlobalExceptionRulesService) DeleteDeviceAdminPolicySetGlobalExceptionByRuleID(id string) (*ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesDeleteDeviceAdminPolicySetGlobalExceptionByRuleID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/global-exception/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesDeleteDeviceAdminPolicySetGlobalExceptionByRuleID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteDeviceAdminPolicySetGlobalExceptionByRuleId")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesDeleteDeviceAdminPolicySetGlobalExceptionByRuleID)
	return result, response, err

}
