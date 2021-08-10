package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type DeviceAdministrationAuthorizationRulesService service

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRules struct {
	Response []ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponse `json:"response,omitempty"` //
	Version  string                                                                                   `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponse struct {
	Commands []string                                                                                   `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseLink `json:"link,omitempty"`     //
	Profile  string                                                                                     `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseRule struct {
	Condition ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseRuleCondition `json:"condition,omitempty"` // <ul><li>Hierarchical structure which defines a set of conditions for which authentication or authorization policy rules could be matched.</li> <li>Logical operations(AND, OR) relationship between conditions are supported</li> <li>Each condition can have subconditions with relation to logical operations</li></ul>
	Default   bool                                                                                                `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts int                                                                                                 `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                              `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                              `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      int                                                                                                 `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                              `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseRuleCondition struct {
	ConditionType string                                                                                                  `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      bool                                                                                                    `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseRuleConditionLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRule struct {
	Response ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponse `json:"response,omitempty"` // Authorization rule for device admin
	Version  string                                                                                   `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponse struct {
	Commands []string                                                                                     `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseLink `json:"link,omitempty"`     //
	Profile  string                                                                                       `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRule struct {
	Condition ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRuleCondition `json:"condition,omitempty"` //
	Default   bool                                                                                                  `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts int                                                                                                   `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      int                                                                                                   `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRuleCondition struct {
	ConditionType       string                                                                                                                   `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            bool                                                                                                                     `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                   `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                   `json:"id,omitempty"`                  //
	Name                string                                                                                                                   `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                   `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                   `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                   `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                   `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                   `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                   `json:"operator,omitempty"`            // Equality operator
	Children            []ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                 `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                 `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRuleConditionChildren struct {
	ConditionType string                                                                                                            `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      bool                                                                                                              `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesResetHitCountsDeviceAdminAuthorizationRules struct {
	Message string `json:"message,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByID struct {
	Response ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponse `json:"response,omitempty"` // Authorization rule for device admin
	Version  string                                                                                    `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponse struct {
	Commands []string                                                                                      `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseLink `json:"link,omitempty"`     //
	Profile  string                                                                                        `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRule struct {
	Condition ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleCondition `json:"condition,omitempty"` //
	Default   bool                                                                                                   `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts int                                                                                                    `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                 `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                 `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      int                                                                                                    `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                 `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleCondition struct {
	ConditionType       string                                                                                                                    `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            bool                                                                                                                      `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                    `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                    `json:"id,omitempty"`                  //
	Name                string                                                                                                                    `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                    `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                    `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                    `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                    `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                    `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                    `json:"operator,omitempty"`            // Equality operator
	Children            []ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                  `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                  `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionChildren struct {
	ConditionType string                                                                                                             `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      bool                                                                                                               `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByID struct {
	Response ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponse `json:"response,omitempty"` // Authorization rule for device admin
	Version  string                                                                                       `json:"version,omitempty"`  //
}

type ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponse struct {
	Commands []string                                                                                         `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseLink `json:"link,omitempty"`     //
	Profile  string                                                                                           `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRule struct {
	Condition ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleCondition `json:"condition,omitempty"` //
	Default   bool                                                                                                      `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts int                                                                                                       `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                                    `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                                    `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      int                                                                                                       `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                                    `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleCondition struct {
	ConditionType       string                                                                                                                       `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            bool                                                                                                                         `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                                       `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                                       `json:"id,omitempty"`                  //
	Name                string                                                                                                                       `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                                       `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                                       `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                                       `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                                       `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                                       `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                                       `json:"operator,omitempty"`            // Equality operator
	Children            []ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                                     `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                                     `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionChildren struct {
	ConditionType string                                                                                                                `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      bool                                                                                                                  `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseDeviceAdministrationAuthorizationRulesDeleteDeviceAdminAuthorizationRuleByID struct {
	ID string `json:"id,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRule struct {
	Commands []string                                                                             `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     *RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleLink `json:"link,omitempty"`     //
	Profile  string                                                                               `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     *RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRule struct {
	Condition *RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                         `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                          `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                        `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                        `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                          `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                        `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleCondition struct {
	ConditionType       string                                                                                                           `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                            `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                           `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                           `json:"id,omitempty"`                  //
	Name                string                                                                                                           `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                           `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                           `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                           `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                           `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                           `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                           `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                         `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                         `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionChildren struct {
	ConditionType string                                                                                                    `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                     `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByID struct {
	Commands []string                                                                                 `json:"commands,omitempty"` // Command sets enforce the specified list of commands that can be executed by a device administrator
	Link     *RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDLink `json:"link,omitempty"`     //
	Profile  string                                                                                   `json:"profile,omitempty"`  // Device admin profiles control the initial login session of the device administrator
	Rule     *RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRule `json:"rule,omitempty"`     // Common attributes in rule authentication/authorization
}

type RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRule struct {
	Condition *RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleCondition `json:"condition,omitempty"` //
	Default   *bool                                                                                             `json:"default,omitempty"`   // Indicates if this rule is the default one
	HitCounts *int                                                                                              `json:"hitCounts,omitempty"` // The amount of times the rule was matched
	ID        string                                                                                            `json:"id,omitempty"`        // The identifier of the rule
	Name      string                                                                                            `json:"name,omitempty"`      // Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank      *int                                                                                              `json:"rank,omitempty"`      // The rank(priority) in relation to other rules. Lower rank is higher priority.
	State     string                                                                                            `json:"state,omitempty"`     // The state that the rule is in. A disabled rule cannot be matched.
}

type RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleCondition struct {
	ConditionType       string                                                                                                               `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                                                `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                                               `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                                               `json:"id,omitempty"`                  //
	Name                string                                                                                                               `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                                               `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                                               `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                                               `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                                               `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                                               `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                                               `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                                             `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                                             `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildren struct {
	ConditionType string                                                                                                        `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                                         `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

//GetDeviceAdminAuthorizationRules Device Admin - Get authorization rules.
/* Device Admin Get authorization rules.

@param policyID policyId path parameter. Policy id
*/
func (s *DeviceAdministrationAuthorizationRulesService) GetDeviceAdminAuthorizationRules(policyID string) (*ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRules, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/authorization"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminAuthorizationRules")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRules)
	return result, response, err

}

//GetDeviceAdminAuthorizationRuleByID Device Admin - Get authorization rule attributes.
/* Device Admin Get authorization rule attributes.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *DeviceAdministrationAuthorizationRulesService) GetDeviceAdminAuthorizationRuleByID(policyID string, id string) (*ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/authorization/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAdminAuthorizationRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByID)
	return result, response, err

}

//CreateDeviceAdminAuthorizationRule Device Admin - Create authorization rule.
/* Device Admin Create authorization rule.

@param policyID policyId path parameter. Policy id
*/
func (s *DeviceAdministrationAuthorizationRulesService) CreateDeviceAdminAuthorizationRule(policyID string, requestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRule *RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRule) (*ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRule, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/authorization"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRule).
		SetResult(&ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRule{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceAdminAuthorizationRule")
	}

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRule)
	return result, response, err

}

//ResetHitCountsDeviceAdminAuthorizationRules Device Admin - Reset HitCount for Authorization Rules
/* Device Admin Reset HitCount for Authorization Rules

@param policyID policyId path parameter. Policy id
*/
func (s *DeviceAdministrationAuthorizationRulesService) ResetHitCountsDeviceAdminAuthorizationRules(policyID string) (*ResponseDeviceAdministrationAuthorizationRulesResetHitCountsDeviceAdminAuthorizationRules, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/authorization/reset-hitcount"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationAuthorizationRulesResetHitCountsDeviceAdminAuthorizationRules{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ResetHitCountsDeviceAdminAuthorizationRules")
	}

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationRulesResetHitCountsDeviceAdminAuthorizationRules)
	return result, response, err

}

//UpdateDeviceAdminAuthorizationRuleByID Device Admin - Update authorization rule.
/* Device Admin Update authorization rule.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *DeviceAdministrationAuthorizationRulesService) UpdateDeviceAdminAuthorizationRuleByID(policyID string, id string, requestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleById *RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByID) (*ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/authorization/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleById).
		SetResult(&ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceAdminAuthorizationRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByID)
	return result, response, err

}

//DeleteDeviceAdminAuthorizationRuleByID Device Admin - Delete authorization rule.
/* Device Admin Delete authorization rule.

@param policyID policyId path parameter. Policy id
@param id id path parameter. Rule id
*/
func (s *DeviceAdministrationAuthorizationRulesService) DeleteDeviceAdminAuthorizationRuleByID(policyID string, id string) (*ResponseDeviceAdministrationAuthorizationRulesDeleteDeviceAdminAuthorizationRuleByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/device-admin/policy-set/{policyId}/authorization/{id}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceAdministrationAuthorizationRulesDeleteDeviceAdminAuthorizationRuleByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteDeviceAdminAuthorizationRuleById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDeviceAdministrationAuthorizationRulesDeleteDeviceAdminAuthorizationRuleByID)
	return result, response, err

}
