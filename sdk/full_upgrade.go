package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type FullUpgradeService service

type GetPrecheckStatusQueryParams struct {
	PreCheckReportID string `url:"preCheckReportID,omitempty"` //preCheckReportID
	PreCheckID       string `url:"preCheckID,omitempty"`       //preCheckID
}

type ProceedStatusQueryParams struct {
	PreCheckReportID string `url:"preCheckReportID,omitempty"` //preCheckReportID
}

type StageStatusQueryParams struct {
	PreCheckReportID string `url:"preCheckReportID,omitempty"` //preCheckReportID
}

type ResponseFullUpgradeGetPrecheckStatus struct {
	IsValid          *bool                                            `json:"isValid,omitempty"`          //
	Nodecount        *int                                             `json:"nodecount,omitempty"`        //
	PreCheckReportID string                                           `json:"preCheckReportID,omitempty"` //
	PreChecks        *[]ResponseFullUpgradeGetPrecheckStatusPreChecks `json:"preChecks,omitempty"`        //
	Status           string                                           `json:"status,omitempty"`           //
}

type ResponseFullUpgradeGetPrecheckStatusPreChecks struct {
	CheckType      string                                                `json:"checkType,omitempty"`      //
	Displayname    string                                                `json:"displayname,omitempty"`    //
	ExecutionTime  *int                                                  `json:"executionTime,omitempty"`  //
	Message        string                                                `json:"message,omitempty"`        //
	Name           string                                                `json:"name,omitempty"`           //
	Nodes          *[]ResponseFullUpgradeGetPrecheckStatusPreChecksNodes `json:"nodes,omitempty"`          //
	OnFailure      string                                                `json:"onFailure,omitempty"`      //
	Percentage     *int                                                  `json:"percentage,omitempty"`     //
	RemediationMsg string                                                `json:"remediationMsg,omitempty"` //
	Status         string                                                `json:"status,omitempty"`         //
	SuccessMsg     string                                                `json:"successMsg,omitempty"`     //
	SuccessNodes   *int                                                  `json:"successNodes,omitempty"`   //
	UpdateTime     *int                                                  `json:"updateTime,omitempty"`     //
}

type ResponseFullUpgradeGetPrecheckStatusPreChecksNodes struct {
	Msg    string `json:"msg,omitempty"`    //
	Name   string `json:"name,omitempty"`   //
	Status string `json:"status,omitempty"` //
}

type ResponseFullUpgradeRunPreChecksInPPAN struct {
	Response *ResponseFullUpgradeRunPreChecksInPPANResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  //
}

type ResponseFullUpgradeRunPreChecksInPPANResponse struct {
	Message          string `json:"message,omitempty"`          //
	PreCheckReportID string `json:"preCheckReportID,omitempty"` // ID which can be used to track the status of upgrade task.
}

type ResponseFullUpgradeProceedStatus struct {
	Nodes                *[]ResponseFullUpgradeProceedStatusNodes `json:"nodes,omitempty"`                //
	Percentage           *int                                     `json:"percentage,omitempty"`           //
	RemainingUpgradeTime *int                                     `json:"remainingUpgradeTime,omitempty"` //
	Status               string                                   `json:"status,omitempty"`               //
}

type ResponseFullUpgradeProceedStatusNodes struct {
	DbStatus    string `json:"dbStatus,omitempty"`    //
	Message     string `json:"message,omitempty"`     //
	Node        string `json:"node,omitempty"`        //
	Percentage  *int   `json:"percentage,omitempty"`  //
	ProgressMsg string `json:"progressMsg,omitempty"` //
	Status      string `json:"status,omitempty"`      //
	UpgradeTime *int   `json:"upgradeTime,omitempty"` //
}

type ResponseFullUpgradeInitiateUpgradeOnPPAN struct {
	Response *ResponseFullUpgradeInitiateUpgradeOnPPANResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}

type ResponseFullUpgradeInitiateUpgradeOnPPANResponse struct {
	Message          string `json:"message,omitempty"`          //
	PreCheckReportID string `json:"preCheckReportID,omitempty"` // ID which can be used to track the status of upgrade task.
}

type ResponseFullUpgradeCancelStagingOnPPAN struct {
	Response *ResponseFullUpgradeCancelStagingOnPPANResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  //
}

type ResponseFullUpgradeCancelStagingOnPPANResponse struct {
	Message          string `json:"message,omitempty"`          //
	PreCheckReportID string `json:"preCheckReportID,omitempty"` // ID which can be used to track the status of upgrade task.
}

type ResponseFullUpgradestageStatus struct {
	Response *[]ResponseFullUpgradestageStatusResponse `json:"response,omitempty"` //
}

type ResponseFullUpgradestageStatusResponse struct {
	DbStatus    string `json:"dbStatus,omitempty"`    //
	Message     string `json:"message,omitempty"`     //
	Node        string `json:"node,omitempty"`        //
	Percentage  *int   `json:"percentage,omitempty"`  //
	ProgressMsg string `json:"progressMsg,omitempty"` //
	Status      string `json:"status,omitempty"`      //
}

type ResponseFullUpgradeInitiateStagingOnPPAN struct {
	Response *ResponseFullUpgradeInitiateStagingOnPPANResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}

type ResponseFullUpgradeInitiateStagingOnPPANResponse struct {
	Message          string `json:"message,omitempty"`          //
	PreCheckReportID string `json:"preCheckReportID,omitempty"` // ID which can be used to track the status of upgrade task.
}

type ResponseFullUpgradeUpgradesummary struct {
	Nodes                *[]ResponseFullUpgradeUpgradesummaryNodes `json:"nodes,omitempty"`                //
	Percentage           *int                                      `json:"percentage,omitempty"`           //
	RemainingUpgradeTime *int                                      `json:"remainingUpgradeTime,omitempty"` //
	Status               string                                    `json:"status,omitempty"`               //
}

type ResponseFullUpgradeUpgradesummaryNodes struct {
	DbStatus    string `json:"dbStatus,omitempty"`    //
	Message     string `json:"message,omitempty"`     //
	Node        string `json:"node,omitempty"`        //
	Percentage  *int   `json:"percentage,omitempty"`  //
	ProgressMsg string `json:"progressMsg,omitempty"` //
	Status      string `json:"status,omitempty"`      //
	UpgradeTime *int   `json:"upgradeTime,omitempty"` //
}

type RequestFullUpgradeRunPreChecksInPPAN struct {
	BundleName       string   `json:"bundleName,omitempty"`       //
	Hostnames        []string `json:"hostnames,omitempty"`        //
	PatchBundleName  string   `json:"patchBundleName,omitempty"`  //
	PreCheckReportID string   `json:"preCheckReportID,omitempty"` //
	PreChecks        []string `json:"preChecks,omitempty"`        //
	ReTrigger        *bool    `json:"reTrigger,omitempty"`        //
	RepoName         string   `json:"repoName,omitempty"`         //
	UpgradeType      string   `json:"upgradeType,omitempty"`      //
}

type RequestFullUpgradeInitiateUpgradeOnPPAN struct {
	Hostnames        []string `json:"hostnames,omitempty"`        //
	PreCheckReportID string   `json:"preCheckReportID,omitempty"` //
	UpgradeType      string   `json:"upgradeType,omitempty"`      //
}

type RequestFullUpgradeCancelStagingOnPPAN struct {
	Hostnames        []string `json:"hostnames,omitempty"`        //
	PreCheckReportID string   `json:"preCheckReportID,omitempty"` //
	UpgradeType      string   `json:"upgradeType,omitempty"`      //
}

type RequestFullUpgradeInitiateStagingOnPPAN struct {
	Hostnames        []string `json:"hostnames,omitempty"`        //
	PreCheckReportID string   `json:"preCheckReportID,omitempty"` //
	ReTrigger        *bool    `json:"reTrigger,omitempty"`        //
	UpgradeType      string   `json:"upgradeType,omitempty"`      //
}

//GetPrecheckStatus gets status of pre-checks
/* get the latest precheck report

@param getPrecheckStatusQueryParams Filtering parameter
*/
func (s *FullUpgradeService) GetPrecheckStatus(getPrecheckStatusQueryParams *GetPrecheckStatusQueryParams) (*ResponseFullUpgradeGetPrecheckStatus, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/upgrade/prepare/get-status"

	queryString, _ := query.Values(getPrecheckStatusQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseFullUpgradeGetPrecheckStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPrecheckStatus")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseFullUpgradeGetPrecheckStatus)
	return result, response, err

}

//ProceedStatus get the status of upgrade proceed process
/* get the status of upgrade process for the requested nodes

@param proceedStatusQueryParams Filtering parameter
*/
func (s *FullUpgradeService) ProceedStatus(proceedStatusQueryParams *ProceedStatusQueryParams) (*ResponseFullUpgradeProceedStatus, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/upgrade/proceed/get-status"

	queryString, _ := query.Values(proceedStatusQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseFullUpgradeProceedStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ProceedStatus")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseFullUpgradeProceedStatus)
	return result, response, err

}

//StageStatus get the status a particular upgrade stage process
/* get the status of upgrade stage process for the requested nodes

@param stageStatusQueryParams Filtering parameter
*/
func (s *FullUpgradeService) StageStatus(stageStatusQueryParams *StageStatusQueryParams) (*ResponseFullUpgradestageStatus, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/upgrade/stage/get-status"

	queryString, _ := query.Values(stageStatusQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseFullUpgradestageStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation StageStatus")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseFullUpgradestageStatus)
	return result, response, err

}

//Upgradesummary get the summary of upgrade process
/* get the summary of upgrade process

 */
func (s *FullUpgradeService) Upgradesummary() (*ResponseFullUpgradeUpgradesummary, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/upgrade/summary/get-status"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseFullUpgradeUpgradesummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation Upgradesummary")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseFullUpgradeUpgradesummary)
	return result, response, err

}

//RunPreChecksInPPAN API's to initiate pre-checks execution on PPAN
/* API's purpose is to initiate pre-checks execution on PPAN

 */
func (s *FullUpgradeService) RunPreChecksInPPAN(requestFullUpgradeRunPreChecksInPPAN *RequestFullUpgradeRunPreChecksInPPAN) (*ResponseFullUpgradeRunPreChecksInPPAN, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/upgrade/prepare/pre-checks"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFullUpgradeRunPreChecksInPPAN).
		SetResult(&ResponseFullUpgradeRunPreChecksInPPAN{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseFullUpgradeRunPreChecksInPPAN{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RunPreChecksInPPAN")
	}

	result := response.Result().(*ResponseFullUpgradeRunPreChecksInPPAN)
	return result, response, err

}

//InitiateUpgradeOnPPAN API's purpose would be to orchestrate upgrade execution on PPAN
/* API's purpose would be to orchestrate upgrade execution on PPAN

 */
func (s *FullUpgradeService) InitiateUpgradeOnPPAN(requestFullUpgradeInitiateUpgradeOnPPAN *RequestFullUpgradeInitiateUpgradeOnPPAN) (*ResponseFullUpgradeInitiateUpgradeOnPPAN, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/upgrade/proceed/initiate-upgrade"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFullUpgradeInitiateUpgradeOnPPAN).
		SetResult(&ResponseFullUpgradeInitiateUpgradeOnPPAN{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseFullUpgradeInitiateUpgradeOnPPAN{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation InitiateUpgradeOnPPAN")
	}

	result := response.Result().(*ResponseFullUpgradeInitiateUpgradeOnPPAN)
	return result, response, err

}

//CancelStagingOnPPAN API to cancel staging process from PPAN
/* API to cancel staging process of specified nodes from PPAN

 */
func (s *FullUpgradeService) CancelStagingOnPPAN(requestFullUpgradeCancelStagingOnPPAN *RequestFullUpgradeCancelStagingOnPPAN) (*ResponseFullUpgradeCancelStagingOnPPAN, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/upgrade/stage/cancel-stage"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFullUpgradeCancelStagingOnPPAN).
		SetResult(&ResponseFullUpgradeCancelStagingOnPPAN{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseFullUpgradeCancelStagingOnPPAN{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CancelStagingOnPPAN")
	}

	result := response.Result().(*ResponseFullUpgradeCancelStagingOnPPAN)
	return result, response, err

}

//InitiateStagingOnPPAN API to initiate staging orcheastration from PPAN
/* API to initiate staging orcheastration from PPAN

 */
func (s *FullUpgradeService) InitiateStagingOnPPAN(requestFullUpgradeInitiateStagingOnPPAN *RequestFullUpgradeInitiateStagingOnPPAN) (*ResponseFullUpgradeInitiateStagingOnPPAN, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/upgrade/stage/start-stage"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFullUpgradeInitiateStagingOnPPAN).
		SetResult(&ResponseFullUpgradeInitiateStagingOnPPAN{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseFullUpgradeInitiateStagingOnPPAN{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation InitiateStagingOnPPAN")
	}

	result := response.Result().(*ResponseFullUpgradeInitiateStagingOnPPAN)
	return result, response, err

}
