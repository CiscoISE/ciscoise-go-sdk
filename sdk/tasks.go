package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type TasksService service

type ResponseTasksGetTaskStatus struct {
	ID              string        `json:"id,omitempty"`              //
	ExecutionStatus string        `json:"executionStatus,omitempty"` //
	ModuleType      string        `json:"moduleType,omitempty"`      //
	StartTime       string        `json:"startTime,omitempty"`       //
	ResourcesCount  int           `json:"resourcesCount,omitempty"`  //
	SuccessCount    int           `json:"successCount,omitempty"`    //
	FailCount       int           `json:"failCount,omitempty"`       //
	DetailStatus    []interface{} `json:"detailStatus,omitempty"`    //
}

type ResponseTasksGetTaskStatusByID struct {
	ID              string        `json:"id,omitempty"`              //
	ExecutionStatus string        `json:"executionStatus,omitempty"` //
	ModuleType      string        `json:"moduleType,omitempty"`      //
	StartTime       string        `json:"startTime,omitempty"`       //
	ResourcesCount  int           `json:"resourcesCount,omitempty"`  //
	SuccessCount    int           `json:"successCount,omitempty"`    //
	FailCount       int           `json:"failCount,omitempty"`       //
	DetailStatus    []interface{} `json:"detailStatus,omitempty"`    //
}

//GetTaskStatus get all task status
/* get all task status

 */
func (s *TasksService) GetTaskStatus() (*[]ResponseTasksGetTaskStatus, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/task"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&[]ResponseTasksGetTaskStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTaskStatus")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*[]ResponseTasksGetTaskStatus)
	return result, response, err

}

//GetTaskStatusByID Monitor task status
/* Monitor task status

@param taskID taskId path parameter. The id of the task executed before
*/
func (s *TasksService) GetTaskStatusByID(taskID string) (*ResponseTasksGetTaskStatusByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/task/{taskId}"
	path = strings.Replace(path, "{taskId}", fmt.Sprintf("%v", taskID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTasksGetTaskStatusByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTaskStatusById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTasksGetTaskStatusByID)
	return result, response, err

}
