package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type VersionAndPatchService service

type ResponseVersionAndPatchGetIseVersionAndPatch struct {
	OperationResult *ResponseVersionAndPatchGetIseVersionAndPatchOperationResult `json:"OperationResult,omitempty"` //
}

type ResponseVersionAndPatchGetIseVersionAndPatchOperationResult struct {
	ResultValue *[]ResponseVersionAndPatchGetIseVersionAndPatchOperationResultResultValue `json:"resultValue,omitempty"` //
}

type ResponseVersionAndPatchGetIseVersionAndPatchOperationResultResultValue struct {
	Value string `json:"value,omitempty"` //
	Name  string `json:"name,omitempty"`  //
}

//GetIseVersionAndPatch Get Cisco ISE version and patch information
/* This API allows the client to get Cisco ISE version and patch information.

 */
func (s *VersionAndPatchService) GetIseVersionAndPatch() (*ResponseVersionAndPatchGetIseVersionAndPatch, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/op/systemconfig/iseversion"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseVersionAndPatchGetIseVersionAndPatch{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIseVersionAndPatch")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseVersionAndPatchGetIseVersionAndPatch)
	return result, response, err

}
