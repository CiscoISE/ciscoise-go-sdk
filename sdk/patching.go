package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type PatchingService service

type ResponsePatchingListInstalledHotpatches struct {
	Response *[]ResponsePatchingListInstalledHotpatchesResponse `json:"response,omitempty"` //
}

type ResponsePatchingListInstalledHotpatchesResponse struct {
	HotpatchName string `json:"hotpatchName,omitempty"` //
	InstallDate  string `json:"installDate,omitempty"`  //
}

type ResponsePatchingInstallHotpatch struct {
	Response *ResponsePatchingInstallHotpatchResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}

type ResponsePatchingInstallHotpatchResponse struct {
	ID      string `json:"id,omitempty"`      // ID which can be used to track the status of install / rollback of patch and hotpatch.
	Message string `json:"message,omitempty"` //
}

type ResponsePatchingRollbackHotpatch struct {
	Response *ResponsePatchingRollbackHotpatchResponse `json:"response,omitempty"` //
	Version  string                                    `json:"version,omitempty"`  //
}

type ResponsePatchingRollbackHotpatchResponse struct {
	ID      string `json:"id,omitempty"`      // ID which can be used to track the status of install / rollback of patch and hotpatch.
	Message string `json:"message,omitempty"` //
}

type ResponsePatchingListInstalledPatches struct {
	IseVersion   string                                              `json:"iseVersion,omitempty"`   //
	PatchVersion *[]ResponsePatchingListInstalledPatchesPatchVersion `json:"patchVersion,omitempty"` //
}

type ResponsePatchingListInstalledPatchesPatchVersion struct {
	InstallDate string `json:"installDate,omitempty"` //
	PatchNumber *int   `json:"patchNumber,omitempty"` //
}

type ResponsePatchingInstallPatch struct {
	Response *ResponsePatchingInstallPatchResponse `json:"response,omitempty"` //
	Version  string                                `json:"version,omitempty"`  //
}

type ResponsePatchingInstallPatchResponse struct {
	ID      string `json:"id,omitempty"`      // ID which can be used to track the status of install / rollback of patch and hotpatch.
	Message string `json:"message,omitempty"` //
}

type ResponsePatchingRollbackPatch struct {
	Response *ResponsePatchingRollbackPatchResponse `json:"response,omitempty"` //
	Version  string                                 `json:"version,omitempty"`  //
}

type ResponsePatchingRollbackPatchResponse struct {
	ID      string `json:"id,omitempty"`      // ID which can be used to track the status of install / rollback of patch and hotpatch.
	Message string `json:"message,omitempty"` //
}

type RequestPatchingInstallHotpatch struct {
	HotpatchName   string `json:"hotpatchName,omitempty"`   //
	RepositoryName string `json:"repositoryName,omitempty"` //
}

type RequestPatchingRollbackHotpatch struct {
	HotpatchName   string `json:"hotpatchName,omitempty"`   //
	RepositoryName string `json:"repositoryName,omitempty"` //
}

type RequestPatchingInstallPatch struct {
	PatchName      string `json:"patchName,omitempty"`      //
	RepositoryName string `json:"repositoryName,omitempty"` //
}

type RequestPatchingRollbackPatch struct {
	PatchNumber *int `json:"patchNumber,omitempty"` //
}

//ListInstalledHotpatches List installed hot patches
/* List all the installed hot patches in the system.

 */
func (s *PatchingService) ListInstalledHotpatches() (*ResponsePatchingListInstalledHotpatches, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/hotpatch"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePatchingListInstalledHotpatches{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ListInstalledHotpatches")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePatchingListInstalledHotpatches)
	return result, response, err

}

//ListInstalledPatches List installed patches
/* List all the installed patches in the system, with the patch number for rollback.

 */
func (s *PatchingService) ListInstalledPatches() (*ResponsePatchingListInstalledPatches, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/patch"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePatchingListInstalledPatches{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ListInstalledPatches")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePatchingListInstalledPatches)
	return result, response, err

}

//InstallHotpatch Trigger hotpatch installation on Cisco ISE node.
/* Triggers hot patch installation on the Cisco ISE node. A task ID is returned which  can be used to monitor the progress of the hot patch installation process.  As hot patch installation triggers the Cisco ISE to restart, the task API becomes  unavailable for a certain period of time.

 */
func (s *PatchingService) InstallHotpatch(requestPatchingInstallHotpatch *RequestPatchingInstallHotpatch) (*ResponsePatchingInstallHotpatch, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/hotpatch/install"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestPatchingInstallHotpatch).
		SetResult(&ResponsePatchingInstallHotpatch{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation InstallHotpatch")
	}

	result := response.Result().(*ResponsePatchingInstallHotpatch)
	return result, response, err

}

//RollbackHotpatch Rollback hotpatch from the Cisco ISE node.
/* Triggers hot patch rollback on the Cisco ISE node. A task ID is returned which  can be used to monitor the progress of the hot patch rollback process.  As hot patch rollback triggers the Cisco ISE to restart, the task API becomes  unavailable for a certain period of time.

 */
func (s *PatchingService) RollbackHotpatch(requestPatchingRollbackHotpatch *RequestPatchingRollbackHotpatch) (*ResponsePatchingRollbackHotpatch, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/hotpatch/rollback"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestPatchingRollbackHotpatch).
		SetResult(&ResponsePatchingRollbackHotpatch{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RollbackHotpatch")
	}

	result := response.Result().(*ResponsePatchingRollbackHotpatch)
	return result, response, err

}

//InstallPatch Trigger patch installation on the Cisco ISE node.
/* Triggers patch installation on the Cisco ISE node. A task ID is returned which can be used to monitor the progress of the patch installation process. As the patch   installation triggers the Cisco ISE to restart, the task API becomes unavailable for  a certain period of time.

 */
func (s *PatchingService) InstallPatch(requestPatchingInstallPatch *RequestPatchingInstallPatch) (*ResponsePatchingInstallPatch, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/patch/install"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestPatchingInstallPatch).
		SetResult(&ResponsePatchingInstallPatch{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation InstallPatch")
	}

	result := response.Result().(*ResponsePatchingInstallPatch)
	return result, response, err

}

//RollbackPatch Removes patch from the Cisco ISE node.
/* Triggers patch rollback on the Cisco ISE node. A task ID is returned which can be used to monitor the progress of the patch rollback process. As the patch   rollback triggers the Cisco ISE to restart, the task API becomes unavailable for  a certain period of time.

 */
func (s *PatchingService) RollbackPatch(requestPatchingRollbackPatch *RequestPatchingRollbackPatch) (*ResponsePatchingRollbackPatch, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/patch/rollback"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestPatchingRollbackPatch).
		SetResult(&ResponsePatchingRollbackPatch{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RollbackPatch")
	}

	result := response.Result().(*ResponsePatchingRollbackPatch)
	return result, response, err

}
