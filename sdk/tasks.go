package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type TasksService service

type ResponseTasksGetAllTaskStatus []ResponseItemTasksGetAllTaskStatus // Array of ResponseTasksGetAllTaskStatus

type ResponseItemTasksGetAllTaskStatus struct {
	DetailStatus    *[]ResponseItemTasksGetAllTaskStatusDetailStatus `json:"detailStatus,omitempty"`    //
	ExecutionStatus string                                           `json:"executionStatus,omitempty"` //
	FailCount       *int                                             `json:"failCount,omitempty"`       //
	ID              string                                           `json:"id,omitempty"`              //
	ModuleType      string                                           `json:"moduleType,omitempty"`      //
	ResourcesCount  *int                                             `json:"resourcesCount,omitempty"`  //
	StartTime       string                                           `json:"startTime,omitempty"`       //
	SuccessCount    *int                                             `json:"successCount,omitempty"`    //
}

type ResponseItemTasksGetAllTaskStatusDetailStatus interface{}

type ResponseTasksGetTaskStatus struct {
	DetailStatus    *[]ResponseTasksGetTaskStatusDetailStatus `json:"detailStatus,omitempty"`    //
	ExecutionStatus string                                    `json:"executionStatus,omitempty"` //
	FailCount       *int                                      `json:"failCount,omitempty"`       //
	ID              string                                    `json:"id,omitempty"`              //
	ModuleType      string                                    `json:"moduleType,omitempty"`      //
	ResourcesCount  *int                                      `json:"resourcesCount,omitempty"`  //
	StartTime       string                                    `json:"startTime,omitempty"`       //
	SuccessCount    *int                                      `json:"successCount,omitempty"`    //
}

type ResponseTasksGetTaskStatusDetailStatus interface{}

//GetAllTaskStatus get all task status
/* get all task status

 */
func (s *TasksService) GetAllTaskStatus() (*ResponseTasksGetAllTaskStatus, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/task"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTasksGetAllTaskStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAllTaskStatus")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTasksGetAllTaskStatus)
	return result, response, err

}

//GetTaskStatus Monitor task status
/* Monitor task status

@param taskID taskId path parameter. The id of the task executed before
*/
func (s *TasksService) GetTaskStatus(taskID string) (*ResponseTasksGetTaskStatus, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/task/{taskId}"
	path = strings.Replace(path, "{taskId}", fmt.Sprintf("%v", taskID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTasksGetTaskStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTaskStatus")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseTasksGetTaskStatus)
	return result, response, err

}
