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
	ConditionType string                                                                               `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionLink       `json:"link,omitempty"`          //
	Children      *[]ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionChildren `json:"children,omitempty"`      // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionChildren struct {
	Name            string `json:"name,omitempty"`            // Dictionary attribute name
	AttributeName   string `json:"attributeName,omitempty"`   // Dictionary attribute name
	AttributeID     string `json:"attributeId,omitempty"`     // Dictionary attribute id (Optional), used for additional verification
	AttributeValue  string `json:"attributeValue,omitempty"`  // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName  string `json:"dictionaryName,omitempty"`  // Dictionary name
	DictionaryValue string `json:"dictionaryValue,omitempty"` // Dictionary value
	Operator        string `json:"operator,omitempty"`        // Equality operator
	EndDate         string `json:"endDate,omitempty"`         //
	StartDate       string `json:"startDate,omitempty"`       // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
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
	ConditionType string                                                                                 `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                  `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionLink       `json:"link,omitempty"`          //
	Children      *[]ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionChildren `json:"children,omitempty"`      // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
}

type ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySetResponseConditionChildren struct {
	Name            string `json:"name,omitempty"`            // Dictionary attribute name
	AttributeName   string `json:"attributeName,omitempty"`   // Dictionary attribute name
	AttributeID     string `json:"attributeId,omitempty"`     // Dictionary attribute id (Optional), used for additional verification
	AttributeValue  string `json:"attributeValue,omitempty"`  // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName  string `json:"dictionaryName,omitempty"`  // Dictionary name
	DictionaryValue string `json:"dictionaryValue,omitempty"` // Dictionary value
	Operator        string `json:"operator,omitempty"`        // Equality operator
	EndDate         string `json:"endDate,omitempty"`         //
	StartDate       string `json:"startDate,omitempty"`       // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
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
	ConditionType string                                                                                  `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                   `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionLink       `json:"link,omitempty"`          //
	Children      *[]ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionChildren `json:"children,omitempty"`      // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionChildren struct {
	Name            string `json:"name,omitempty"`            // Dictionary attribute name
	AttributeName   string `json:"attributeName,omitempty"`   // Dictionary attribute name
	AttributeID     string `json:"attributeId,omitempty"`     // Dictionary attribute id (Optional), used for additional verification
	AttributeValue  string `json:"attributeValue,omitempty"`  // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName  string `json:"dictionaryName,omitempty"`  // Dictionary name
	DictionaryValue string `json:"dictionaryValue,omitempty"` // Dictionary value
	Operator        string `json:"operator,omitempty"`        // Equality operator
	EndDate         string `json:"endDate,omitempty"`         //
	StartDate       string `json:"startDate,omitempty"`       // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
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
	ConditionType string                                                                                     `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                                      `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionLink       `json:"link,omitempty"`          //
	Children      *[]ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionChildren `json:"children,omitempty"`      // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
}

type ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDResponseConditionChildren struct {
	Name            string `json:"name,omitempty"`            // Dictionary attribute name
	AttributeName   string `json:"attributeName,omitempty"`   // Dictionary attribute name
	AttributeID     string `json:"attributeId,omitempty"`     // Dictionary attribute id (Optional), used for additional verification
	AttributeValue  string `json:"attributeValue,omitempty"`  // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName  string `json:"dictionaryName,omitempty"`  // Dictionary name
	DictionaryValue string `json:"dictionaryValue,omitempty"` // Dictionary value
	Operator        string `json:"operator,omitempty"`        // Equality operator
	EndDate         string `json:"endDate,omitempty"`         //
	StartDate       string `json:"startDate,omitempty"`       // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
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
	ConditionType string                                                                        `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	ID            string                                                                        `json:"id,omitempty"`            // Condition Id
	Name          string                                                                        `json:"name,omitempty"`          // Condition name
	IsNegate      *bool                                                                         `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionLink       `json:"link,omitempty"`          //
	Children      *[]RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildren `json:"children,omitempty"`      // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
}

type RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildren struct {
	Name            string `json:"name,omitempty"`            // Dictionary attribute name
	AttributeName   string `json:"attributeName,omitempty"`   // Dictionary attribute name
	AttributeID     string `json:"attributeId,omitempty"`     // Dictionary attribute id (Optional), used for additional verification
	AttributeValue  string `json:"attributeValue,omitempty"`  // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName  string `json:"dictionaryName,omitempty"`  // Dictionary name
	DictionaryValue string `json:"dictionaryValue,omitempty"` // Dictionary value
	Operator        string `json:"operator,omitempty"`        // Equality operator
	EndDate         string `json:"endDate,omitempty"`         //
	StartDate       string `json:"startDate,omitempty"`       // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
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
	ConditionType string                                                                            `json:"conditionType,omitempty"` // <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
	IsNegate      *bool                                                                             `json:"isNegate,omitempty"`      // Indicates whereas this condition is in negate mode
	Link          *RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionLink       `json:"link,omitempty"`          //
	Children      *[]RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildren `json:"children,omitempty"`      // In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition
}

type RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildren struct {
	Name            string `json:"name,omitempty"`            // Dictionary attribute name
	AttributeName   string `json:"attributeName,omitempty"`   // Dictionary attribute name
	AttributeID     string `json:"attributeId,omitempty"`     // Dictionary attribute id (Optional), used for additional verification
	AttributeValue  string `json:"attributeValue,omitempty"`  // <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
	DictionaryName  string `json:"dictionaryName,omitempty"`  // Dictionary name
	DictionaryValue string `json:"dictionaryValue,omitempty"` // Dictionary value
	Operator        string `json:"operator,omitempty"`        // Equality operator
	EndDate         string `json:"endDate,omitempty"`         //
	StartDate       string `json:"startDate,omitempty"`       // <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>
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
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessPolicySetCreateNetworkAccessPolicySet{}, response, nil
		}
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
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessPolicySetResetHitCountsNetworkAccessPolicySets{}, response, nil
		}
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
		if err.Error() == emptyStringToJSONError {
			return &ResponseNetworkAccessPolicySetUpdateNetworkAccessPolicySetByID{}, response, nil
		}
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
