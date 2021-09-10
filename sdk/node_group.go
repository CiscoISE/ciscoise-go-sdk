package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type NodeGroupService service

type ResponseNodeGroupGetNodeGroups struct {
	Response *[]ResponseNodeGroupGetNodeGroupsResponse `json:"response,omitempty"` //
}

type ResponseNodeGroupGetNodeGroupsResponse struct {
	Name        string                                          `json:"name,omitempty"`        //
	Description string                                          `json:"description,omitempty"` //
	MarCache    *ResponseNodeGroupGetNodeGroupsResponseMarCache `json:"mar-cache,omitempty"`   //
}

type ResponseNodeGroupGetNodeGroupsResponseMarCache struct {
	Enabled             *bool `json:"enabled,omitempty"`              //
	ReplicationTimeout  *int  `json:"replication-timeout,omitempty"`  //
	ReplicationAttempts *int  `json:"replication-attempts,omitempty"` //
	QueryTimeout        *int  `json:"query-timeout,omitempty"`        //
	QueryAttempts       *int  `json:"query-attempts,omitempty"`       //
}

type ResponseNodeGroupGetNodeGroup struct {
	Name        string                                 `json:"name,omitempty"`        //
	Description string                                 `json:"description,omitempty"` //
	MarCache    *ResponseNodeGroupGetNodeGroupMarCache `json:"mar-cache,omitempty"`   //
}

type ResponseNodeGroupGetNodeGroupMarCache struct {
	Enabled             *bool `json:"enabled,omitempty"`              //
	ReplicationTimeout  *int  `json:"replication-timeout,omitempty"`  //
	ReplicationAttempts *int  `json:"replication-attempts,omitempty"` //
	QueryTimeout        *int  `json:"query-timeout,omitempty"`        //
	QueryAttempts       *int  `json:"query-attempts,omitempty"`       //
}

type ResponseNodeGroupCreateNodeGroup struct {
	Code      *int   `json:"code,omitempty"`      //
	Message   string `json:"message,omitempty"`   //
	RootCause string `json:"rootCause,omitempty"` //
}

type ResponseNodeGroupUpdateNodeGroup struct {
	Code      *int   `json:"code,omitempty"`      //
	Message   string `json:"message,omitempty"`   //
	RootCause string `json:"rootCause,omitempty"` //
}

type ResponseNodeGroupDeleteNodeGroup struct {
	Code      *int   `json:"code,omitempty"`      //
	Message   string `json:"message,omitempty"`   //
	RootCause string `json:"rootCause,omitempty"` //
}

type RequestNodeGroupCreateNodeGroup struct {
	Description string                                   `json:"description,omitempty"` //
	MarCache    *RequestNodeGroupCreateNodeGroupMarCache `json:"mar-cache,omitempty"`   //
}

type RequestNodeGroupCreateNodeGroupMarCache struct {
	Enabled             *bool `json:"enabled,omitempty"`              //
	ReplicationTimeout  *int  `json:"replication-timeout,omitempty"`  //
	ReplicationAttempts *int  `json:"replication-attempts,omitempty"` //
	QueryTimeout        *int  `json:"query-timeout,omitempty"`        //
	QueryAttempts       *int  `json:"query-attempts,omitempty"`       //
}

type RequestNodeGroupUpdateNodeGroup struct {
	Description string                                   `json:"description,omitempty"` //
	MarCache    *RequestNodeGroupUpdateNodeGroupMarCache `json:"mar-cache,omitempty"`   //
}

type RequestNodeGroupUpdateNodeGroupMarCache struct {
	Enabled             *bool `json:"enabled,omitempty"`              //
	ReplicationTimeout  *int  `json:"replication-timeout,omitempty"`  //
	ReplicationAttempts *int  `json:"replication-attempts,omitempty"` //
	QueryTimeout        *int  `json:"query-timeout,omitempty"`        //
	QueryAttempts       *int  `json:"query-attempts,omitempty"`       //
}

//GetNodeGroups Get the list of all node group.
/* Get details of all the node groups in the cluster. To detect node failure and to reset all URL-redirected sessions on the failed node, two or more Policy Service nodes can be placed in the same node group. When a node that belongs to a node group fails, another node in the same node group issues a Change of Authorization (CoA) for all URL-redirected sessions on the failed node.

 */
func (s *NodeGroupService) GetNodeGroups() (*ResponseNodeGroupGetNodeGroups, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node-group"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeGroupGetNodeGroups{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNodeGroups")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeGroupGetNodeGroups)
	return result, response, err

}

//GetNodeGroup Get details of a node group
/* Get details of a node group in the cluster.

@param nodegroupname node-group-name path parameter. ID of the existing node group.
*/
func (s *NodeGroupService) GetNodeGroup(nodegroupname string) (*ResponseNodeGroupGetNodeGroup, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node-group/{node-group-name}"
	path = strings.Replace(path, "{node-group-name}", fmt.Sprintf("%v", nodegroupname), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeGroupGetNodeGroup{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNodeGroup")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeGroupGetNodeGroup)
	return result, response, err

}

//CreateNodeGroup Creates a node group.
/* Developers need to create node group in the system.Node Group is a group of PSNs, mainly used for terminating posture pending sessions when a PSN in local node group fails.Node group members can communicate over TCP/7800.

@param nodegroupname node-group-name path parameter. ID of the existing node group.
*/
func (s *NodeGroupService) CreateNodeGroup(nodegroupname string, requestNodeGroupCreateNodeGroup *RequestNodeGroupCreateNodeGroup) (*ResponseNodeGroupCreateNodeGroup, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node-group/{node-group-name}"
	path = strings.Replace(path, "{node-group-name}", fmt.Sprintf("%v", nodegroupname), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNodeGroupCreateNodeGroup).
		SetResult(&ResponseNodeGroupCreateNodeGroup{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNodeGroup")
	}

	result := response.Result().(*ResponseNodeGroupCreateNodeGroup)
	return result, response, err

}

//UpdateNodeGroup Edits an existing node group.
/* API updates an existing node group in the system.

@param nodegroupname node-group-name path parameter. ID of the existing node group.
*/
func (s *NodeGroupService) UpdateNodeGroup(nodegroupname string, requestNodeGroupUpdateNodeGroup *RequestNodeGroupUpdateNodeGroup) (*ResponseNodeGroupUpdateNodeGroup, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node-group/{node-group-name}"
	path = strings.Replace(path, "{node-group-name}", fmt.Sprintf("%v", nodegroupname), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNodeGroupUpdateNodeGroup).
		SetResult(&ResponseNodeGroupUpdateNodeGroup{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNodeGroup")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeGroupUpdateNodeGroup)
	return result, response, err

}

//DeleteNodeGroup Delete a node group.
/* Developers need to delete node group in the system.

@param nodegroupname node-group-name path parameter. ID of the existing node group.
*/
func (s *NodeGroupService) DeleteNodeGroup(nodegroupname string) (*ResponseNodeGroupDeleteNodeGroup, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node-group/{node-group-name}"
	path = strings.Replace(path, "{node-group-name}", fmt.Sprintf("%v", nodegroupname), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeGroupDeleteNodeGroup{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteNodeGroup")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeGroupDeleteNodeGroup)
	return result, response, err

}
