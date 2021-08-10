package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type SyncIseNodeService service

type ResponseSyncIseNodeSyncNode struct {
	Code      int    `json:"code,omitempty"`      //
	Message   string `json:"message,omitempty"`   //
	RootCause string `json:"rootCause,omitempty"` //
}

type RequestSyncIseNodeSyncNode struct {
	Hostname string `json:"hostname,omitempty"` //
}

//SyncNode API to trigger manual sync of the node this API is invoked on
/* Performing a manual sync will involve a reload of the target node, but not the primary PAN. There might be situations where if the node has been out of sync for a long time, it may not be possible to recover via a manual sync.

 */
func (s *SyncIseNodeService) SyncNode(requestSyncIseNodeSyncNode *RequestSyncIseNodeSyncNode) (*ResponseSyncIseNodeSyncNode, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/deployment/sync-node"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSyncIseNodeSyncNode).
		SetResult(&ResponseSyncIseNodeSyncNode{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation SyncNode")
	}

	result := response.Result().(*ResponseSyncIseNodeSyncNode)
	return result, response, err

}
