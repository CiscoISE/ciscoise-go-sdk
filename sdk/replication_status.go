package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type ReplicationStatusService service

type ResponseReplicationStatusGetNodeReplicationStatus struct {
	NodeStatus string `json:"NodeStatus,omitempty"` //
}

//GetNodeReplicationStatus API to get replication status of a node
/* Retrives replication status of a node

@param node node path parameter. ID of the existing node.
*/
func (s *ReplicationStatusService) GetNodeReplicationStatus(node string) (*ResponseReplicationStatusGetNodeReplicationStatus, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/replication-status/{node}"
	path = strings.Replace(path, "{node}", fmt.Sprintf("%v", node), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseReplicationStatusGetNodeReplicationStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNodeReplicationStatus")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseReplicationStatusGetNodeReplicationStatus)
	return result, response, err

}
