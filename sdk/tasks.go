package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type TasksService service

type ResponseTasksGetAllTaskStatus []ResponseItemTasksGetAllTaskStatus // Array of ResponseTasksGetAllTaskStatus
// ResponseTasksGetTaskStatus is deprecated and does not have valid alias
// type ResponseTasksGetTaskStatus = ResponseTasksGetAllTaskStatus

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

// ResponseItemTasksGetTaskStatus is deprecated, please use ResponseItemTasksGetAllTaskStatus
type ResponseItemTasksGetTaskStatus = ResponseItemTasksGetAllTaskStatus

type ResponseItemTasksGetAllTaskStatusDetailStatus interface{}

// ResponseItemTasksGetTaskStatusDetailStatus is deprecated, please use ResponseItemTasksGetAllTaskStatusDetailStatus
type ResponseItemTasksGetTaskStatusDetailStatus = ResponseItemTasksGetAllTaskStatusDetailStatus

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

// ResponseTasksGetTaskStatusByID is deprecated, please use ResponseTasksGetTaskStatus
type ResponseTasksGetTaskStatusByID = ResponseTasksGetTaskStatus

type ResponseTasksGetTaskStatusDetailStatus interface{}

// ResponseTasksGetTaskStatusByIDDetailStatus is deprecated, please use ResponseTasksGetTaskStatusDetailStatus
type ResponseTasksGetTaskStatusByIDDetailStatus = ResponseTasksGetTaskStatusDetailStatus

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

//Alias of GetTaskStatus Monitor task status
func (s *TasksService) GetTaskStatusByID(taskID string) (*ResponseTasksGetTaskStatus, *resty.Response, error) {
	return s.GetTaskStatus(taskID)
}
