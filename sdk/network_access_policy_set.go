package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type NetworkAccessPolicySetService service

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySets struct {
	Response *[]ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponse `json:"response,omitempty"` //
	Version  string                                                              `json:"version,omitempty"`  //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponse struct {
	Condition   *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseCondition `json:"condition,omitempty"`   //
	Default     *bool                                                                      `json:"default,omitempty"`     // Flag which indicates if this policy set is the default one
	Description string                                                                     `json:"description,omitempty"` // The description for the policy set
	HitCounts   *int                                                                       `json:"hitCounts,omitempty"`   // The amount of times the policy was matched
	ID          string                                                                     `json:"id,omitempty"`          // Identifier for the policy set
	IsProxy     *bool                                                                      `json:"isProxy,omitempty"`     // Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	Link        *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseLink      `json:"link,omitempty"`        //
	Name        string                                                                     `json:"name,omitempty"`        // Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank        *int                                                                       `json:"rank,omitempty"`        // The rank(priority) in relation to other policy set. Lower rank is higher priority.
	ServiceName string                                                                     `json:"serviceName,omitempty"` // Policy set service identifier - Allowed Protocols,Server Sequence..
	State       string                                                                     `json:"state,omitempty"`       // The state that the policy set is in. A disabled policy set cannot be matched.
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseCondition struct {
	ConditionType       string                                                                                        `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                         `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                        `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                        `json:"id,omitempty"`                  //
	Name                string                                                                                        `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                        `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                        `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                        `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                        `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                        `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                        `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                      `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                      `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionChildren struct {
	ConditionType string                                                                                 `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                  `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySet struct {
	Response *ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponse `json:"response,omitempty"` // Policy set structure
	Version  string                                                              `json:"version,omitempty"`  //
}

type ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponse struct {
	Condition   *ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseCondition `json:"condition,omitempty"`   //
	Default     *bool                                                                        `json:"default,omitempty"`     // Flag which indicates if this policy set is the default one
	Description string                                                                       `json:"description,omitempty"` // The description for the policy set
	HitCounts   *int                                                                         `json:"hitCounts,omitempty"`   // The amount of times the policy was matched
	ID          string                                                                       `json:"id,omitempty"`          // Identifier for the policy set
	IsProxy     *bool                                                                        `json:"isProxy,omitempty"`     // Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	Link        *ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseLink      `json:"link,omitempty"`        //
	Name        string                                                                       `json:"name,omitempty"`        // Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank        *int                                                                         `json:"rank,omitempty"`        // The rank(priority) in relation to other policy set. Lower rank is higher priority.
	ServiceName string                                                                       `json:"serviceName,omitempty"` // Policy set service identifier - Allowed Protocols,Server Sequence..
	State       string                                                                       `json:"state,omitempty"`       // The state that the policy set is in. A disabled policy set cannot be matched.
}

type ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseCondition struct {
	ConditionType       string                                                                                          `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                           `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                          `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                          `json:"id,omitempty"`                  //
	Name                string                                                                                          `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                          `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                          `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                          `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                          `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                          `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                          `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                        `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                        `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionChildren struct {
	ConditionType string                                                                                   `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                    `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessPolicySetResetHitCountsNetworkAccessPolicySets struct {
	Message string `json:"message,omitempty"` //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByID struct {
	Response *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponse `json:"response,omitempty"` // Policy set structure
	Version  string                                                               `json:"version,omitempty"`  //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponse struct {
	Condition   *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseCondition `json:"condition,omitempty"`   //
	Default     *bool                                                                         `json:"default,omitempty"`     // Flag which indicates if this policy set is the default one
	Description string                                                                        `json:"description,omitempty"` // The description for the policy set
	HitCounts   *int                                                                          `json:"hitCounts,omitempty"`   // The amount of times the policy was matched
	ID          string                                                                        `json:"id,omitempty"`          // Identifier for the policy set
	IsProxy     *bool                                                                         `json:"isProxy,omitempty"`     // Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	Link        *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseLink      `json:"link,omitempty"`        //
	Name        string                                                                        `json:"name,omitempty"`        // Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank        *int                                                                          `json:"rank,omitempty"`        // The rank(priority) in relation to other policy set. Lower rank is higher priority.
	ServiceName string                                                                        `json:"serviceName,omitempty"` // Policy set service identifier - Allowed Protocols,Server Sequence..
	State       string                                                                        `json:"state,omitempty"`       // The state that the policy set is in. A disabled policy set cannot be matched.
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseCondition struct {
	ConditionType       string                                                                                           `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                            `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                           `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                           `json:"id,omitempty"`                  //
	Name                string                                                                                           `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                           `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                           `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                           `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                           `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                           `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                           `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                         `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                         `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionChildren struct {
	ConditionType string                                                                                    `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                     `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByID struct {
	Response *ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponse `json:"response,omitempty"` // Policy set structure
	Version  string                                                                  `json:"version,omitempty"`  //
}

type ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponse struct {
	Condition   *ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseCondition `json:"condition,omitempty"`   //
	Default     *bool                                                                            `json:"default,omitempty"`     // Flag which indicates if this policy set is the default one
	Description string                                                                           `json:"description,omitempty"` // The description for the policy set
	HitCounts   *int                                                                             `json:"hitCounts,omitempty"`   // The amount of times the policy was matched
	ID          string                                                                           `json:"id,omitempty"`          // Identifier for the policy set
	IsProxy     *bool                                                                            `json:"isProxy,omitempty"`     // Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	Link        *ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseLink      `json:"link,omitempty"`        //
	Name        string                                                                           `json:"name,omitempty"`        // Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank        *int                                                                             `json:"rank,omitempty"`        // The rank(priority) in relation to other policy set. Lower rank is higher priority.
	ServiceName string                                                                           `json:"serviceName,omitempty"` // Policy set service identifier - Allowed Protocols,Server Sequence..
	State       string                                                                           `json:"state,omitempty"`       // The state that the policy set is in. A disabled policy set cannot be matched.
}

type ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseCondition struct {
	ConditionType       string                                                                                              `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                               `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                              `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                              `json:"id,omitempty"`                  //
	Name                string                                                                                              `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                              `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                              `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                              `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                              `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                              `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                              `json:"operator,omitempty"`            // Equality operator
	Children            *[]ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                            `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                            `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionChildren struct {
	ConditionType string                                                                                       `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                        `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionChildrenLink `json:"link,omitempty"`          //
}

type ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessPolicySetDeleteNetworkAccessPolicySetByID struct {
	ID string `json:"id,omitempty"` //
}

type RequestNetworkAccessPolicySetCreateNetworkAccessPolicySet struct {
	Condition   *RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetCondition `json:"condition,omitempty"`   //
	Default     *bool                                                               `json:"default,omitempty"`     // Flag which indicates if this policy set is the default one
	Description string                                                              `json:"description,omitempty"` // The description for the policy set
	HitCounts   *int                                                                `json:"hitCounts,omitempty"`   // The amount of times the policy was matched
	ID          string                                                              `json:"id,omitempty"`          // Identifier for the policy set
	IsProxy     *bool                                                               `json:"isProxy,omitempty"`     // Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	Link        *RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetLink      `json:"link,omitempty"`        //
	Name        string                                                              `json:"name,omitempty"`        // Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank        *int                                                                `json:"rank,omitempty"`        // The rank(priority) in relation to other policy set. Lower rank is higher priority.
	ServiceName string                                                              `json:"serviceName,omitempty"` // Policy set service identifier - Allowed Protocols,Server Sequence..
	State       string                                                              `json:"state,omitempty"`       // The state that the policy set is in. A disabled policy set cannot be matched.
}

type RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetCondition struct {
	ConditionType       string                                                                                 `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                  `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                 `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                 `json:"id,omitempty"`                  //
	Name                string                                                                                 `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                 `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                 `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                 `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                 `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                 `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                 `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                               `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                               `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildren struct {
	ConditionType string                                                                          `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                           `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByID struct {
	Condition   *RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDCondition `json:"condition,omitempty"`   //
	Default     *bool                                                                   `json:"default,omitempty"`     // Flag which indicates if this policy set is the default one
	Description string                                                                  `json:"description,omitempty"` // The description for the policy set
	HitCounts   *int                                                                    `json:"hitCounts,omitempty"`   // The amount of times the policy was matched
	ID          string                                                                  `json:"id,omitempty"`          // Identifier for the policy set
	IsProxy     *bool                                                                   `json:"isProxy,omitempty"`     // Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	Link        *RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDLink      `json:"link,omitempty"`        //
	Name        string                                                                  `json:"name,omitempty"`        // Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Rank        *int                                                                    `json:"rank,omitempty"`        // The rank(priority) in relation to other policy set. Lower rank is higher priority.
	ServiceName string                                                                  `json:"serviceName,omitempty"` // Policy set service identifier - Allowed Protocols,Server Sequence..
	State       string                                                                  `json:"state,omitempty"`       // The state that the policy set is in. A disabled policy set cannot be matched.
}

type RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDCondition struct {
	ConditionType       string                                                                                     `json:"conditionType,omitempty"`       // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate            *bool                                                                                      `json:"isNegate,omitempty"`            // Indicates whereas this condition is in negate mode
	Link                *RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionLink                `json:"link,omitempty"`                //
	Description         string                                                                                     `json:"description,omitempty"`         // Condition description
	ID                  string                                                                                     `json:"id,omitempty"`                  //
	Name                string                                                                                     `json:"name,omitempty"`                // Condition name
	AttributeName       string                                                                                     `json:"attributeName,omitempty"`       // Dictionary attribute name
	AttributeID         string                                                                                     `json:"attributeId,omitempty"`         // Dictionary attribute id (Optional), used for additional verification
	AttributeValue      string                                                                                     `json:"attributeValue,omitempty"`      // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName      string                                                                                     `json:"dictionaryName,omitempty"`      // Dictionary name
	DictionaryValue     string                                                                                     `json:"dictionaryValue,omitempty"`     // Dictionary value
	Operator            string                                                                                     `json:"operator,omitempty"`            // Equality operator
	Children            *[]RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildren          `json:"children,omitempty"`            // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
	DatesRange          *RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionDatesRange          `json:"datesRange,omitempty"`          // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	DatesRangeException *RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionDatesRangeException `json:"datesRangeException,omitempty"` // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
	HoursRange          *RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionHoursRange          `json:"hoursRange,omitempty"`          // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	HoursRangeException *RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionHoursRangeException `json:"hoursRangeException,omitempty"` // <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>
	WeekDays            []string                                                                                   `json:"weekDays,omitempty"`            // <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
	WeekDaysException   []string                                                                                   `json:"weekDaysException,omitempty"`   // <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>
}

type RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildren struct {
	ConditionType string                                                                              `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                               `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildrenLink `json:"link,omitempty"`          //
}

type RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildrenLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionDatesRange struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionDatesRangeException struct {
	EndDate   string `json:"endDate,omitempty"`   //
	StartDate string `json:"startDate,omitempty"` //
}

type RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionHoursRange struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionHoursRangeException struct {
	EndTime   string `json:"endTime,omitempty"`   //
	StartTime string `json:"startTime,omitempty"` //
}

type RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

//GetNetworkAccessPolicySets Network Access - List of policy sets.
/* Get all network access policy sets.

 */
func (s *NetworkAccessPolicySetService) GetNetworkAccessPolicySets() (*ResponseNetworkAccessPolicySetGetNetworkAccessPolicySets, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessPolicySetGetNetworkAccessPolicySets{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessPolicySets")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessPolicySetGetNetworkAccessPolicySets)
	return result, response, err

}

//GetNetworkAccessPolicySetByID Network Access - Get policy set attributes.
/* Network Access Get policy set attributes.

@param id id path parameter. Policy id
*/
func (s *NetworkAccessPolicySetService) GetNetworkAccessPolicySetByID(id string) (*ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAccessPolicySetById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByID)
	return result, response, err

}

//CreateNetworkAccessPolicySet Network Access - Create a new policy set.
/* Network Access Create a new policy set:



 Policy must include name , service identifier (either server sequence or allowed protocol) and a condition.


 Condition has hierarchical structure which define a set of condition for which policy could be match.


 Condition can be either reference to a stored Library condition, using model
ConditionReference
,
or, dynamically built conditions which are not stored in the conditions Library, using models
ConditionAttributes, ConditionAndBlock, ConditionOrBlock
.





*/
func (s *NetworkAccessPolicySetService) CreateNetworkAccessPolicySet(requestNetworkAccessPolicySetCreateNetworkAccessPolicySet *RequestNetworkAccessPolicySetCreateNetworkAccessPolicySet) (*ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySet, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessPolicySetCreateNetworkAccessPolicySet).
		SetResult(&ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySet{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkAccessPolicySet")
	}

	result := response.Result().(*ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySet)
	return result, response, err

}

//ResetHitCountsNetworkAccessPolicySets Network Access - Reset HitCount for PolicySets
/* Network Access Reset HitCount for PolicySets

 */
func (s *NetworkAccessPolicySetService) ResetHitCountsNetworkAccessPolicySets() (*ResponseNetworkAccessPolicySetResetHitCountsNetworkAccessPolicySets, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/reset-hitcount"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessPolicySetResetHitCountsNetworkAccessPolicySets{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ResetHitCountsNetworkAccessPolicySets")
	}

	result := response.Result().(*ResponseNetworkAccessPolicySetResetHitCountsNetworkAccessPolicySets)
	return result, response, err

}

//UpdateNetworkAccessPolicySetByID Network Access - Update a policy set.
/* Network Access Update a policy set.

@param id id path parameter. Policy id
*/
func (s *NetworkAccessPolicySetService) UpdateNetworkAccessPolicySetByID(id string, requestNetworkAccessPolicySetUpdateNetworkAccessPolicySetById *RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByID) (*ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkAccessPolicySetUpdateNetworkAccessPolicySetById).
		SetResult(&ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkAccessPolicySetById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByID)
	return result, response, err

}

//DeleteNetworkAccessPolicySetByID Network Access - Delete a policy set.
/* Network Access Delete a policy set.

@param id id path parameter. Policy id
*/
func (s *NetworkAccessPolicySetService) DeleteNetworkAccessPolicySetByID(id string) (*ResponseNetworkAccessPolicySetDeleteNetworkAccessPolicySetByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/policy/network-access/policy-set/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkAccessPolicySetDeleteNetworkAccessPolicySetByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteNetworkAccessPolicySetById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNetworkAccessPolicySetDeleteNetworkAccessPolicySetByID)
	return result, response, err

}
