package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type NodeDeploymentService service

type GetDeploymentNodesQueryParams struct {
	Filter     []string `url:"filter,omitempty"`     //<div> <style type="text/css" scoped> .apiServiceTable td, .apiServiceTable th { padding: 5px 10px !important; text-align: left; } </style> <span> <b>Simple filtering</b> is available through the filter query string parameter. The structure of a filter is a triplet of field operator and value, separated by dots. More than one filter can be sent. The logical operator common to all filter criteria is AND by default, and can be changed by using the <i>"filterType=or"</i> query string parameter. Each resource Data model description should specify if an attribute is a filtered field. </span> <br /> <table class="apiServiceTable"> <thead> <tr> <th>OPERATOR</th> <th>DESCRIPTION</th> </tr> </thead> <tbody> <tr> <td>EQ</td> <td>Equals</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> </tr> <tr> <td>STARTSW</td> <td>Starts With</td> </tr> <tr> <td>NSTARTSW</td> <td>Not Starts With</td> </tr> <tr> <td>ENDSW</td> <td>Ends With</td> </tr> <tr> <td>NENDSW</td> <td>Not Ends With</td> </tr> <tr> <td>CONTAINS</td> <td>Contains</td> </tr> <tr> <td>NCONTAINS</td> <td>Not Contains</td> </tr> </tbody> </table> </div>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to all filter criteria is AND by default, and can be changed by using this parameter.
}

type ResponseNodeDeploymentMakePrimary struct {
	Success *ResponseNodeDeploymentMakePrimarySuccess `json:"success,omitempty"` //
	Version string                                    `json:"version,omitempty"` //
}

type ResponseNodeDeploymentMakePrimarySuccess struct {
	Message string `json:"message,omitempty"` //
}

type ResponseNodeDeploymentMakeStandalone struct {
	Success *ResponseNodeDeploymentMakeStandaloneSuccess `json:"success,omitempty"` //
	Version string                                       `json:"version,omitempty"` //
}

type ResponseNodeDeploymentMakeStandaloneSuccess struct {
	Message string `json:"message,omitempty"` //
}

type ResponseNodeDeploymentGetDeploymentNodes struct {
	Response *[]ResponseNodeDeploymentGetDeploymentNodesResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  //
}

type ResponseNodeDeploymentGetDeploymentNodesResponse struct {
	Hostname   string   `json:"hostname,omitempty"`   //
	Fqdn       string   `json:"fqdn,omitempty"`       //
	IPAddress  string   `json:"ipAddress,omitempty"`  //
	Roles      []string `json:"roles,omitempty"`      //
	Services   []string `json:"services,omitempty"`   //
	Nodestatus string   `json:"nodeStatus,omitempty"` //
}

type ResponseNodeDeploymentRegisterNode struct {
	Success *ResponseNodeDeploymentRegisterNodeSuccess `json:"success,omitempty"` //
	Version string                                     `json:"version,omitempty"` //
}

type ResponseNodeDeploymentRegisterNodeSuccess struct {
	Message string `json:"message,omitempty"` //
}

type ResponseNodeDeploymentGetNodeDetails struct {
	Response *ResponseNodeDeploymentGetNodeDetailsResponse `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  //
}

type ResponseNodeDeploymentGetNodeDetailsResponse struct {
	Hostname   string   `json:"hostname,omitempty"`   //
	Fqdn       string   `json:"fqdn,omitempty"`       //
	IPAddress  string   `json:"ipAddress,omitempty"`  //
	Roles      []string `json:"roles,omitempty"`      //
	Services   []string `json:"services,omitempty"`   //
	Nodestatus string   `json:"nodeStatus,omitempty"` //
}

type ResponseNodeDeploymentUpdateDeploymentNode struct {
	Success *ResponseNodeDeploymentUpdateDeploymentNodeSuccess `json:"success,omitempty"` //
	Version string                                             `json:"version,omitempty"` //
}

// ResponseNodeDeploymentUpdateNode is deprecated, please use ResponseNodeDeploymentUpdateDeploymentNode
type ResponseNodeDeploymentUpdateNode = ResponseNodeDeploymentUpdateDeploymentNode

type ResponseNodeDeploymentUpdateDeploymentNodeSuccess struct {
	Message string `json:"message,omitempty"` //
}

// ResponseNodeDeploymentUpdateNodeSuccess is deprecated, please use ResponseNodeDeploymentUpdateDeploymentNodeSuccess
type ResponseNodeDeploymentUpdateNodeSuccess = ResponseNodeDeploymentUpdateDeploymentNodeSuccess

type ResponseNodeDeploymentDeleteDeploymentNode struct {
	Success *ResponseNodeDeploymentDeleteDeploymentNodeSuccess `json:"success,omitempty"` //
	Version string                                             `json:"version,omitempty"` //
}

// ResponseNodeDeploymentDeleteNode is deprecated, please use ResponseNodeDeploymentDeleteDeploymentNode
type ResponseNodeDeploymentDeleteNode = ResponseNodeDeploymentDeleteDeploymentNode

type ResponseNodeDeploymentDeleteDeploymentNodeSuccess struct {
	Message string `json:"message,omitempty"` //
}

// ResponseNodeDeploymentDeleteNodeSuccess is deprecated, please use ResponseNodeDeploymentDeleteDeploymentNodeSuccess
type ResponseNodeDeploymentDeleteNodeSuccess = ResponseNodeDeploymentDeleteDeploymentNodeSuccess

type ResponseNodeDeploymentSyncNode struct {
	Response *ResponseNodeDeploymentSyncNodeResponse `json:"response,omitempty"` //
	Version  string                                  `json:"version,omitempty"`  //
}

type ResponseNodeDeploymentSyncNodeResponse struct {
	ID      string `json:"id,omitempty"`      // ID which can be used to track the status of the deployment task using the Task Service API.
	Message string `json:"message,omitempty"` //
}

type ResponseNodeDeploymentPromoteNode struct {
	Success *ResponseNodeDeploymentPromoteNodeSuccess `json:"success,omitempty"` //
	Version string                                    `json:"version,omitempty"` //
}

type ResponseNodeDeploymentPromoteNodeSuccess struct {
	Message string `json:"message,omitempty"` //
}

type RequestNodeDeploymentRegisterNode struct {
	Fqdn            string   `json:"fqdn,omitempty"`            //
	UserName        string   `json:"userName,omitempty"`        //
	Password        string   `json:"password,omitempty"`        //
	AllowCertImport *bool    `json:"allowCertImport,omitempty"` // Consent to import the self-signed certificate of the registering node.
	Roles           []string `json:"roles,omitempty"`           //
	Services        []string `json:"services,omitempty"`        //
}

type RequestNodeDeploymentUpdateDeploymentNode struct {
	Roles    []string `json:"roles,omitempty"`    //
	Services []string `json:"services,omitempty"` //
}

// RequestNodeDeploymentUpdateNode is deprecated, please use RequestNodeDeploymentUpdateDeploymentNode
type RequestNodeDeploymentUpdateNode = RequestNodeDeploymentUpdateDeploymentNode

//GetDeploymentNodes Retrieve the list of all the nodes that are deployed in the cluster.
/* The API lists all the nodes that are deployed in the cluster.

Returns basic information about each of the deployed nodes in the cluster like hostname, status, roles, and services.

Supports filtering on FQDN, hostname, IP address, roles, services and node status.


@param getDeploymentNodesQueryParams Filtering parameter
*/
func (s *NodeDeploymentService) GetDeploymentNodes(getDeploymentNodesQueryParams *GetDeploymentNodesQueryParams) (*ResponseNodeDeploymentGetDeploymentNodes, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node"

	queryString, _ := query.Values(getDeploymentNodesQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNodeDeploymentGetDeploymentNodes{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeploymentNodes")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeDeploymentGetDeploymentNodes)
	return result, response, err

}

//GetNodeDetails Retrieve details of a deployed node.
/* This API retrieves detailed information of the deployed node.


@param hostname hostname path parameter. Hostname of the deployed node.
*/
func (s *NodeDeploymentService) GetNodeDetails(hostname string) (*ResponseNodeDeploymentGetNodeDetails, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node/{hostname}"
	path = strings.Replace(path, "{hostname}", fmt.Sprintf("%v", hostname), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeDeploymentGetNodeDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNodeDetails")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeDeploymentGetNodeDetails)
	return result, response, err

}

//MakePrimary Promote a standalone node to a primary PAN.
/* This API promotes the standalone node on which the API is invoked to the primary Policy Administration node (PAN).


 */
func (s *NodeDeploymentService) MakePrimary() (*ResponseNodeDeploymentMakePrimary, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/primary"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeDeploymentMakePrimary{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNodeDeploymentMakePrimary{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MakePrimary")
	}

	result := response.Result().(*ResponseNodeDeploymentMakePrimary)
	return result, response, err

}

//MakeStandalone Change a primary PAN to a standalone node.
/* This API changes the primary PAN in a single node cluster on which the API is invoked, to a standalone node.


 */
func (s *NodeDeploymentService) MakeStandalone() (*ResponseNodeDeploymentMakeStandalone, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/standalone"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeDeploymentMakeStandalone{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNodeDeploymentMakeStandalone{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MakeStandalone")
	}

	result := response.Result().(*ResponseNodeDeploymentMakeStandalone)
	return result, response, err

}

//RegisterNode Register a standalone node to the cluster
/* This API registers a Cisco ISE node to form a multi-node deployment.


 */
func (s *NodeDeploymentService) RegisterNode(requestNodeDeploymentRegisterNode *RequestNodeDeploymentRegisterNode) (*ResponseNodeDeploymentRegisterNode, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNodeDeploymentRegisterNode).
		SetResult(&ResponseNodeDeploymentRegisterNode{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNodeDeploymentRegisterNode{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RegisterNode")
	}

	result := response.Result().(*ResponseNodeDeploymentRegisterNode)
	return result, response, err

}

//SyncNode Trigger manual synchronization of the node.
/* Performing a manual synchronization involves a reload of the target node, but not the primary PAN.

@param hostname hostname path parameter. Hostname of the node.
*/
func (s *NodeDeploymentService) SyncNode(hostname string) (*ResponseNodeDeploymentSyncNode, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/sync-node/{hostname}"
	path = strings.Replace(path, "{hostname}", fmt.Sprintf("%v", hostname), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeDeploymentSyncNode{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNodeDeploymentSyncNode{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation SyncNode")
	}

	result := response.Result().(*ResponseNodeDeploymentSyncNode)
	return result, response, err

}

//PromoteNode Promote the secondary PAN in a multi-node cluster to the role of primary PAN.
/* Execute this API in the secondary PAN in the cluster to promote the node to primary PAN.  Ensure that the API Gateway setting is enabled in the secondary PAN.


 */
func (s *NodeDeploymentService) PromoteNode() (*ResponseNodeDeploymentPromoteNode, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/promote"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeDeploymentPromoteNode{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNodeDeploymentPromoteNode{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation PromoteNode")
	}

	result := response.Result().(*ResponseNodeDeploymentPromoteNode)
	return result, response, err

}

//UpdateDeploymentNode Replace the existing configuration of the Cisco ISE node with the configuration provided.
/* This API updates the configuration of the Cisco ISE node with the configuration provided.


@param hostname hostname path parameter. Hostname of the deployed node.
*/
func (s *NodeDeploymentService) UpdateDeploymentNode(hostname string, requestNodeDeploymentUpdateDeploymentNode *RequestNodeDeploymentUpdateDeploymentNode) (*ResponseNodeDeploymentUpdateDeploymentNode, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node/{hostname}"
	path = strings.Replace(path, "{hostname}", fmt.Sprintf("%v", hostname), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNodeDeploymentUpdateDeploymentNode).
		SetResult(&ResponseNodeDeploymentUpdateDeploymentNode{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseNodeDeploymentUpdateDeploymentNode{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeploymentNode")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeDeploymentUpdateDeploymentNode)
	return result, response, err

}

// Alias of UpdateDeploymentNode Replace the existing configuration of the Cisco ISE node with the configuration provided.
func (s *NodeDeploymentService) UpdateNode(hostname string, requestNodeDeploymentUpdateDeploymentNode *RequestNodeDeploymentUpdateDeploymentNode) (*ResponseNodeDeploymentUpdateDeploymentNode, *resty.Response, error) {
	return s.UpdateDeploymentNode(hostname, requestNodeDeploymentUpdateDeploymentNode)
}

//DeleteDeploymentNode Remove a deployed node from a cluster.
/* The deregistered node becomes a standalone Cisco ISE node.

It retains the last configuration that it received from the primary PAN and assumes the default roles and services of a standalone node.


@param hostname hostname path parameter. The hostname of the node in the deployment to be deregistered.
*/
func (s *NodeDeploymentService) DeleteDeploymentNode(hostname string) (*ResponseNodeDeploymentDeleteDeploymentNode, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/node/{hostname}"
	path = strings.Replace(path, "{hostname}", fmt.Sprintf("%v", hostname), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNodeDeploymentDeleteDeploymentNode{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteDeploymentNode")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseNodeDeploymentDeleteDeploymentNode)
	return result, response, err

}

// Alias of DeleteDeploymentNode Remove a deployed node from a cluster.
func (s *NodeDeploymentService) DeleteNode(hostname string) (*ResponseNodeDeploymentDeleteDeploymentNode, *resty.Response, error) {
	return s.DeleteDeploymentNode(hostname)
}
