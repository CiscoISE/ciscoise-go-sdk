package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type NodeGroupService service

type DeleteNodeGroupQueryParams struct {
	ForceDelete bool `url:"forceDelete,omitempty"` //Force delete the group even if the node group contains one or more nodes.
}

type ResponseNodeGroupGetNodeGroups struct {
	Response *[]ResponseNodeGroupGetNodeGroupsResponse `json:"response,omitempty"` //
	Version  string                                    `json:"version,omitempty"`  //
}

type ResponseNodeGroupGetNodeGroupsResponse struct {
	Description string                                          `json:"description,omitempty"` //
	MarCache    *ResponseNodeGroupGetNodeGroupsResponseMarCache `json:"marCache,omitempty"`    //
	Name        string                                          `json:"name,omitempty"`        //
}

type ResponseNodeGroupGetNodeGroupsResponseMarCache struct {
	QueryAttempts       *int `json:"query-attempts,omitempty"`       // The number of times Cisco ISE attempts to perform the cache entry query. (0 - 5).
	QueryTimeout        *int `json:"query-timeout,omitempty"`        // The time, in seconds, after which the cache entry query times out. (1 - 10).
	ReplicationAttempts *int `json:"replication-attempts,omitempty"` // The number of times Cisco ISE attempts to perform MAR cache entry replication. (0 - 5).
	ReplicationTimeout  *int `json:"replication-timeout,omitempty"`  // The time, in seconds, after which the cache entry replication times out. (1 - 10).
}

type ResponseNodeGroupCreateNodeGroup struct {
	Success *ResponseNodeGroupCreateNodeGroupSuccess `json:"success,omitempty"` //
	Version string                                   `json:"version,omitempty"` //
}

type ResponseNodeGroupCreateNodeGroupSuccess struct {
	Message string `json:"message,omitempty"` //
}

type ResponseNodeGroupGetNodeGroup struct {
	Response *ResponseNodeGroupGetNodeGroupResponse `json:"response,omitempty"` //
	Version  string                                 `json:"version,omitempty"`  //
}

type ResponseNodeGroupGetNodeGroupResponse struct {
	Description string                                         `json:"description,omitempty"` //
	MarCache    *ResponseNodeGroupGetNodeGroupResponseMarCache `json:"marCache,omitempty"`    //
	Name        string                                         `json:"name,omitempty"`        //
}

type ResponseNodeGroupGetNodeGroupResponseMarCache struct {
	QueryAttempts       *int `json:"query-attempts,omitempty"`       // The number of times Cisco ISE attempts to perform the cache entry query. (0 - 5).
	QueryTimeout        *int `json:"query-timeout,omitempty"`        // The time, in seconds, after which the cache entry query times out. (1 - 10).
	ReplicationAttempts *int `json:"replication-attempts,omitempty"` // The number of times Cisco ISE attempts to perform MAR cache entry replication. (0 - 5).
	ReplicationTimeout  *int `json:"replication-timeout,omitempty"`  // The time, in seconds, after which the cache entry replication times out. (1 - 10).
}

type ResponseNodeGroupUpdateNodeGroup struct {
	Success *ResponseNodeGroupUpdateNodeGroupSuccess `json:"success,omitempty"` //
	Version string                                   `json:"version,omitempty"` //
}

type ResponseNodeGroupUpdateNodeGroupSuccess struct {
	Message string `json:"message,omitempty"` //
}

type ResponseNodeGroupDeleteNodeGroup struct {
	Success *ResponseNodeGroupDeleteNodeGroupSuccess `json:"success,omitempty"` //
	Version string                                   `json:"version,omitempty"` //
}

type ResponseNodeGroupDeleteNodeGroupSuccess struct {
	Message string `json:"message,omitempty"` //
}

type ResponseNodeGroupAddNode struct {
	Success *ResponseNodeGroupAddNodeSuccess `json:"success,omitempty"` //
	Version string                           `json:"version,omitempty"` //
}

type ResponseNodeGroupAddNodeSuccess struct {
	Message string `json:"message,omitempty"` //
}

type ResponseNodeGroupGetNodes struct {
	Response *[]ResponseNodeGroupGetNodesResponse `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  //
}

type ResponseNodeGroupGetNodesResponse struct {
	Hostname string `json:"hostname,omitempty"` //
}

type ResponseNodeGroupRemoveNode struct {
	Success *ResponseNodeGroupRemoveNodeSuccess `json:"success,omitempty"` //
	Version string                              `json:"version,omitempty"` //
}

type ResponseNodeGroupRemoveNodeSuccess struct {
	Message string `json:"message,omitempty"` //
}

type RequestNodeGroupCreateNodeGroup struct {
	Description string                                   `json:"description,omitempty"` //
	MarCache    *RequestNodeGroupCreateNodeGroupMarCache `json:"marCache,omitempty"`    //
	Name        string                                   `json:"name,omitempty"`        //
}

type RequestNodeGroupCreateNodeGroupMarCache struct {
	QueryAttempts       *int `json:"query-attempts,omitempty"`       // The number of times Cisco ISE attempts to perform the cache entry query. (0 - 5).
	QueryTimeout        *int `json:"query-timeout,omitempty"`        // The time, in seconds, after which the cache entry query times out. (1 - 10).
	ReplicationAttempts *int `json:"replication-attempts,omitempty"` // The number of times Cisco ISE attempts to perform MAR cache entry replication. (0 - 5).
	ReplicationTimeout  *int `json:"replication-timeout,omitempty"`  // The time, in seconds, after which the cache entry replication times out. (1 - 10).
}

type RequestNodeGroupUpdateNodeGroup struct {
	Description string                                   `json:"description,omitempty"` //
	MarCache    *RequestNodeGroupUpdateNodeGroupMarCache `json:"marCache,omitempty"`    //
	Name        string                                   `json:"name,omitempty"`        //
}

type RequestNodeGroupUpdateNodeGroupMarCache struct {
	QueryAttempts       *int `json:"query-attempts,omitempty"`       // The number of times Cisco ISE attempts to perform the cache entry query. (0 - 5).
	QueryTimeout        *int `json:"query-timeout,omitempty"`        // The time, in seconds, after which the cache entry query times out. (1 - 10).
	ReplicationAttempts *int `json:"replication-attempts,omitempty"` // The number of times Cisco ISE attempts to perform MAR cache entry replication. (0 - 5).
	ReplicationTimeout  *int `json:"replication-timeout,omitempty"`  // The time, in seconds, after which the cache entry replication times out. (1 - 10).
}

type RequestNodeGroupAddNode struct {
	Hostname string `json:"hostname,omitempty"` //
}

type RequestNodeGroupRemoveNode struct {
	Hostname string `json:"hostname,omitempty"` //
}

//GetNodeGroups Retrieve the list of all the node groups.
/* This API retrieves the details of all the node groups in the cluster.
Each node group retrieved consists of name, description and MAR cache details like query-attempts, query-timeout, replication-attempts, replication-timeout.

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

//GetNodeGroup Retrieve the details of a node group
/* This API retrieves the details of a node group in the cluster using a node group name.

@param nodeGroupName nodeGroupName path parameter. Name of the existing node group.
*/
func (s *NodeGroupService) GetNodeGroup(nodeGroupName string) (*ResponseNodeGroupGetNodeGroup, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node-group/{nodeGroupName}"
	path = strings.Replace(path, "{nodeGroupName}", fmt.Sprintf("%v", nodeGroupName), -1)

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

//GetNodes Retrieve the list of nodes in a given node group.
/* This API retrieves the list of nodes associated with a node group in the cluster with a given node group name.

@param nodeGroupName nodeGroupName path parameter. Name of the existing node group.
*/
func (s *NodeGroupService) GetNodes(nodeGroupName string) (*ResponseNodeGroupGetNodes, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node-group/{nodeGroupName}/node"
	path = strings.Replace(path, "{nodeGroupName}", fmt.Sprintf("%v", nodeGroupName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeGroupGetNodes{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNodes")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeGroupGetNodes)
	return result, response, err

}

//CreateNodeGroup Create a node group.
/*

This API creates a node group in the cluster.  A node group is a group of PSNs, where the PSNs maintain a heartbeat with each other. It is used primarily to terminate or transfer posture-pending sessions when a PSN in a local node group fails.  Node group members can communicate over TCP/7800.

The following parameters are used in the request body of the API:




PARAMETER

DESCRIPTION

EXAMPLE





name
* required

Name of the node group(
valid-range:
 1-100 characters)

{"name": "site1"}



description

Description of the node group (
valid-range:
 0-256 characters)

{"name": "site2", "description": "sample"}



query-attempts

The number of times Cisco ISE attempts to perform the cache entry query. (
valid-range:
 0 5,
default-value:
 1)

{"name": "site3","marCache": {"query-attempts": 1}}



query-timeout

The time, in seconds, after which a cache entry query times out. (
valid-range:
 1 10,
default-value:
 2) second(s)

{"name": "site4","marCache": {"query-timeout": 2}}


replication-attempts

The number of times Cisco ISE attempts to perform MAR cache entry replication. (
valid-range:
 0 5,
default-value:
 2)

{"name": "site5","marCache": {"replication-attempts": 2}}


replication-timeout

The time, in seconds, after which the cache entry replication times out. (
valid-range:
 1 10,
default-value:
 5) second(s)

{"name": "site6","marCache": {"replication-timeout": 5}}




NOTE 1:
: Node group name and description cannot contain any of the following characters: ! % ^ : ; , . ~ @ # & [ { ( | ) } ] ` > <  / \ " + = ?
NOTE 2:
: Parameter marCache stands for Machine Access Restriction (MAR) cache that provides an additional means of controlling authorization for Active Directory-authentication users. We can enable the marCache for a nodegroup by providing key "marCache" in json request. Additionally we may also provide any combination of parameters query-attempts, query-timeout, replication-attempts, replication-timeout in marCache object. If no value is specified for a particular parameter its default value will be recorded.If no marCache object is given, marCache will be considered as disabled.


*/
func (s *NodeGroupService) CreateNodeGroup(requestNodeGroupCreateNodeGroup *RequestNodeGroupCreateNodeGroup) (*ResponseNodeGroupCreateNodeGroup, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node-group"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNodeGroupCreateNodeGroup).
		SetResult(&ResponseNodeGroupCreateNodeGroup{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNodeGroupCreateNodeGroup{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNodeGroup")
	}

	result := response.Result().(*ResponseNodeGroupCreateNodeGroup)
	return result, response, err

}

//AddNode Add a node to a node group.
/*

This API adds a node to the node group in the cluster. When a node that belongs to a node group fails, another node in the same node group issues a Change of Authorization (CoA) for all the URL-redirected sessions on the failed node.

The following parameters are used in the request body of the API:




PARAMETER

DESCRIPTION

EXAMPLE





hostname
* required

Name of the host name

{"hostname": "isenode"}





@param nodeGroupName nodeGroupName path parameter. Name of the existing node group.
*/
func (s *NodeGroupService) AddNode(nodeGroupName string, requestNodeGroupAddNode *RequestNodeGroupAddNode) (*ResponseNodeGroupAddNode, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node-group/{nodeGroupName}/add-node"
	path = strings.Replace(path, "{nodeGroupName}", fmt.Sprintf("%v", nodeGroupName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNodeGroupAddNode).
		SetResult(&ResponseNodeGroupAddNode{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNodeGroupAddNode{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation AddNode")
	}

	result := response.Result().(*ResponseNodeGroupAddNode)
	return result, response, err

}

//RemoveNode Remove a node from a node group.
/*

Purpose of this API is to remove a node from a node group in the cluster. Removing node from the node group does not delete the node, but failover is no longer carried out if the node is not part any node group.

The following parameters are used in the request body of the API:




PARAMETER

DESCRIPTION

EXAMPLE





hostname
* required

Name of the host name

{"hostname": "isenode"}





@param nodeGroupName nodeGroupName path parameter. Name of the existing node group.
*/
func (s *NodeGroupService) RemoveNode(nodeGroupName string, requestNodeGroupRemoveNode *RequestNodeGroupRemoveNode) (*ResponseNodeGroupRemoveNode, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node-group/{nodeGroupName}/remove-node"
	path = strings.Replace(path, "{nodeGroupName}", fmt.Sprintf("%v", nodeGroupName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNodeGroupRemoveNode).
		SetResult(&ResponseNodeGroupRemoveNode{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNodeGroupRemoveNode{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RemoveNode")
	}

	result := response.Result().(*ResponseNodeGroupRemoveNode)
	return result, response, err

}

//UpdateNodeGroup Update an existing node group.
/*

Purpose of this API is to update an existing node group.

The following parameters are used in the request body of the API:




PARAMETER

DESCRIPTION

EXAMPLE





name
* required

 Name of the node group existing in ISE(
valid-range:
 1-100 characters)

NOTE :
name of an existing node group cannot be modified. "name" parameter should be identical to the node group name(nodeGroupName) provided in the path


{"name": "site1"}



description

 Description of the node group (
valid-range:
 0-256 characters)


{"name": "site2", "description": "sample"}



query-attempts

The number of times Cisco ISE attempts to perform the cache entry query. (
valid-range:
 0 5,
default-value:
 1)

{"name": "site3","marCache": {"query-attempts": 1}}



query-timeout

The time, in seconds, after which a cache entry query times out. (
valid-range:
 1 10,
default-value:
 2) second(s)

{"name": "site4","marCache": {"query-timeout": 2}}



replication-attempts

The number of times Cisco ISE attempts to perform MAR cache entry replication. (
valid-range:
 0 5,
default-value:
 2)

{"name": "site5","marCache": {"replication-attempts": 2}}



replication-timeout

The time, in seconds, after which the cache entry replication times out. (
valid-range:
 1 10,
default-value:
 5) second(s)

{"name": "site6","marCache": {"replication-timeout": 5}}




NOTE 1:
 Node group name and description cannot contain any of the following characters: ! % ^ : ; , . ~ @ # & [ { ( | ) } ] ` > <  / \ " + = ?
NOTE 2:
 Parameter marCache stands for Machine Access Restriction (MAR) cache that provides an additional means of controlling authorization for Active Directory-authentication users. We can enable the marCache for a nodegroup by providing key "marCache" in json request. Additionally we may also provide any combination of parameters query-attempts, query-timeout, replication-attempts, replication-timeout in marCache object. If no value is specified for a particular parameter its default value will be recorded. If no marCache object is given, marCache will be disabled.


@param nodeGroupName nodeGroupName path parameter. Name of the existing node group.
*/
func (s *NodeGroupService) UpdateNodeGroup(nodeGroupName string, requestNodeGroupUpdateNodeGroup *RequestNodeGroupUpdateNodeGroup) (*ResponseNodeGroupUpdateNodeGroup, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node-group/{nodeGroupName}"
	path = strings.Replace(path, "{nodeGroupName}", fmt.Sprintf("%v", nodeGroupName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNodeGroupUpdateNodeGroup).
		SetResult(&ResponseNodeGroupUpdateNodeGroup{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNodeGroupUpdateNodeGroup{}, response, nil
		}
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
/* Delete an existing node group in the cluster. Deleting the node group does not delete the nodes, but failover is no longer carried out among the nodes.

@param nodeGroupName nodeGroupName path parameter. Name of the existing node group.
@param deleteNodeGroupQueryParams Filtering parameter
*/
func (s *NodeGroupService) DeleteNodeGroup(nodeGroupName string, deleteNodeGroupQueryParams *DeleteNodeGroupQueryParams) (*ResponseNodeGroupDeleteNodeGroup, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node-group/{nodeGroupName}"
	path = strings.Replace(path, "{nodeGroupName}", fmt.Sprintf("%v", nodeGroupName), -1)

	queryString, _ := query.Values(deleteNodeGroupQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNodeGroupDeleteNodeGroup{}).
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
