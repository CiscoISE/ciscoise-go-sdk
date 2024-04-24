package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type DuoIDentitySyncService service

type ResponseDuoIDentitySyncGetIDentitysync struct {
	Response *[]ResponseDuoIDentitySyncGetIDentitysyncResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}

type ResponseDuoIDentitySyncGetIDentitysyncResponse struct {
	AdGroups    *int   `json:"adGroups,omitempty"`    // Number of Active Directory Groups
	Destination string `json:"destination,omitempty"` // Destination of Identitysync (Mfa Provider)
	LastSync    string `json:"lastSync,omitempty"`    // Time of the last Sync
	Name        string `json:"name,omitempty"`        // Name of the Identitysync configuration
	Source      string `json:"source,omitempty"`      // Source of the Identitysync (Active Directory)
	SyncStatus  string `json:"syncStatus,omitempty"`  //
}

// # Review unknown case
// # Review unknown case
// # Review unknown case
// # Review unknown case

type ResponseDuoIDentitySyncGetIDentitysyncBySyncName struct {
	Response *ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponse `json:"response,omitempty"` // Identitysync configuration information
	Version  string                                                    `json:"version,omitempty"`  //
}

type ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponse struct {
	IDentitySync *ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponseIDentitySync `json:"identitySync,omitempty"` // Identitysync configuration
}

type ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponseIDentitySync struct {
	AdGroups       *[]ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponseIDentitySyncAdGroups     `json:"adGroups,omitempty"`       //
	Configurations *ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponseIDentitySyncConfigurations `json:"configurations,omitempty"` //
	LastSync       string                                                                              `json:"lastSync,omitempty"`       // Time of the last Sync
	SyncName       string                                                                              `json:"syncName,omitempty"`       // Name of the Identitysync config
	SyncSchedule   *ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponseIDentitySyncSyncSchedule   `json:"syncSchedule,omitempty"`   //
	SyncStatus     string                                                                              `json:"syncStatus,omitempty"`     //
}

type ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponseIDentitySyncAdGroups struct {
	Name   string `json:"name,omitempty"`   // Active Directory Group ID
	Sid    string `json:"sid,omitempty"`    // Active Directory Group Security Identifier (SID)
	Source string `json:"source,omitempty"` // Source of the Active Directory Group
}

type ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponseIDentitySyncConfigurations struct {
	ActiveDirectories *[]ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponseIDentitySyncConfigurationsActiveDirectories `json:"activeDirectories,omitempty"` //
}

type ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponseIDentitySyncConfigurationsActiveDirectories struct {
	DirectoryID string `json:"directoryID,omitempty"` // Active Directory ID
	Domain      string `json:"domain,omitempty"`      // Active Directory domain name
	Name        string `json:"name,omitempty"`        // Name of the Active Directory
}

type ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponseIDentitySyncSyncSchedule struct {
	Interval      *int   `json:"interval,omitempty"`      // Frequency of the sync schedule
	IntervalUnit  string `json:"intervalUnit,omitempty"`  // Unit of the time interval
	SchedulerSync string `json:"schedulerSync,omitempty"` // Type of Sync Schedule - If "Recurring", please specify schedule details
	StartDate     string `json:"startDate,omitempty"`     // Start date and start time of the sync schedule
}

// # Review unknown case
// # Review unknown case

type RequestDuoIDentitySyncCreateIDentitysync struct {
	IDentitySync *RequestDuoIDentitySyncCreateIDentitysyncIDentitySync `json:"identitySync,omitempty"` // Identitysync config information format
}

type RequestDuoIDentitySyncCreateIDentitysyncIDentitySync struct {
	AdGroups       *[]RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncAdGroups     `json:"adGroups,omitempty"`       //
	Configurations *RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncConfigurations `json:"configurations,omitempty"` //
	LastSync       string                                                              `json:"lastSync,omitempty"`       // Time of the last Sync
	SyncName       string                                                              `json:"syncName,omitempty"`       // Name of the Identitysync configuration
	SyncSchedule   *RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncSyncSchedule   `json:"syncSchedule,omitempty"`   //
	SyncStatus     string                                                              `json:"syncStatus,omitempty"`     //
}

type RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncAdGroups struct {
	Name   string `json:"name,omitempty"`   // Active Directory Group ID
	Source string `json:"source,omitempty"` // Source of the Active Directory Group
}

type RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncConfigurations struct {
	ActiveDirectories *[]RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncConfigurationsActiveDirectories `json:"activeDirectories,omitempty"` //
}

type RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncConfigurationsActiveDirectories struct {
	DirectoryID string `json:"directoryID,omitempty"` // Active Directory ID
	Domain      string `json:"domain,omitempty"`      // Active Directory domain name
	Name        string `json:"name,omitempty"`        // Name of the Active Directory
}

type RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncSyncSchedule struct {
	Interval      *int   `json:"interval,omitempty"`      // Frequency of the sync schedule
	IntervalUnit  string `json:"intervalUnit,omitempty"`  // Unit of the time interval
	SchedulerSync string `json:"schedulerSync,omitempty"` // Type of Sync Schedule - If "Recurring", please specify schedule details
	StartDate     string `json:"startDate,omitempty"`     // Start date and start time of the sync schedule
}

type RequestDuoIDentitySyncUpdateStatus struct {
	ErrorList *[]RequestDuoIDentitySyncUpdateStatusErrorList `json:"errorList,omitempty"` //
	Status    string                                         `json:"status,omitempty"`    // status of sync
}

type RequestDuoIDentitySyncUpdateStatusErrorList struct {
	Reason string                                           `json:"reason,omitempty"` // Reason user failed sync
	User   *RequestDuoIDentitySyncUpdateStatusErrorListUser `json:"user,omitempty"`   // User to be synced to Duo
}

type RequestDuoIDentitySyncUpdateStatusErrorListUser struct {
	Directoryname string `json:"directoryname,omitempty"` // active directory that duo user is contained in
	Email         string `json:"email,omitempty"`         // email of Duo user
	Firstname     string `json:"firstname,omitempty"`     // first name of Duo user
	Groupname     string `json:"groupname,omitempty"`     // acrive directory group that duo user is contained in
	Lastname      string `json:"lastname,omitempty"`      // last name of duo user
	Notes         string `json:"notes,omitempty"`         // notes of Duo user
	Realname      string `json:"realname,omitempty"`      // realname of Duo user
	Status        string `json:"status,omitempty"`        // status of Duo user
	Username      string `json:"username,omitempty"`      // username of Duo user
}

type RequestDuoIDentitySyncUpdateIDentitysyncBySyncName struct {
	IDentitySync *RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySync `json:"identitySync,omitempty"` // Identitysync config information format
}

type RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySync struct {
	AdGroups       *[]RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncAdGroups     `json:"adGroups,omitempty"`       //
	Configurations *RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncConfigurations `json:"configurations,omitempty"` //
	LastSync       string                                                                        `json:"lastSync,omitempty"`       // Time of the last Sync
	SyncName       string                                                                        `json:"syncName,omitempty"`       // Name of the Identitysync configuration
	SyncSchedule   *RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncSyncSchedule   `json:"syncSchedule,omitempty"`   //
	SyncStatus     string                                                                        `json:"syncStatus,omitempty"`     //
}

type RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncAdGroups struct {
	Name   string `json:"name,omitempty"`   // Active Directory Group ID
	Source string `json:"source,omitempty"` // Source of the Active Directory Group
}

type RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncConfigurations struct {
	ActiveDirectories *[]RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncConfigurationsActiveDirectories `json:"activeDirectories,omitempty"` //
}

type RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncConfigurationsActiveDirectories struct {
	DirectoryID string `json:"directoryID,omitempty"` // Active Directory ID
	Domain      string `json:"domain,omitempty"`      // Active Directory domain name
	Name        string `json:"name,omitempty"`        // Name of the Active Directory
}

type RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncSyncSchedule struct {
	Interval      *int   `json:"interval,omitempty"`      // Frequency of the sync schedule
	IntervalUnit  string `json:"intervalUnit,omitempty"`  // Unit of the time interval
	SchedulerSync string `json:"schedulerSync,omitempty"` // Type of Sync Schedule - If "Recurring", please specify schedule details
	StartDate     string `json:"startDate,omitempty"`     // Start date and start time of the sync schedule
}

//GetIDentitysync Get the list of all Identitysync configurations
/* Duo-IdentitySync Get the list of all Identitysync configurations

 */
func (s *DuoIDentitySyncService) GetIDentitysync() (*ResponseDuoIDentitySyncGetIDentitysync, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-identitysync/identitysync"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDuoIDentitySyncGetIDentitysync{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIdentitysync")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDuoIDentitySyncGetIDentitysync)
	return result, response, err

}

//CancelSync Cancel the Sync between the ActiveDirectory and the Mfa Provider
/* Cancel the sync between the Active Directory and the corresponding Mfa provider associated with this Identitysync config.

@param syncName syncName path parameter. Name of the Identitysync configuration used to cancel sync.
*/
func (s *DuoIDentitySyncService) CancelSync(syncName string) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-identitysync/identitysync/sync/cancel/{syncName}"
	path = strings.Replace(path, "{syncName}", fmt.Sprintf("%v", syncName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CancelSync")
	}

	getCSFRToken(response.Header())

	return response, err

}

//Sync Initiate the Sync between the ActiveDirectory and the Mfa Provider
/* Initiate the sync between the Active Directory and the corresponding Mfa provider associated with this Identitysync config.

@param syncName syncName path parameter. Name of the Identitysync configuration used to initiate sync.
*/
func (s *DuoIDentitySyncService) Sync(syncName string) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-identitysync/identitysync/sync/{syncName}"
	path = strings.Replace(path, "{syncName}", fmt.Sprintf("%v", syncName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation Sync")
	}

	getCSFRToken(response.Header())

	return response, err

}

//GetIDentitysyncBySyncName Get the specified Identitysync config
/* Duo-IdentitySync Get the Identitysync config specified in the syncName

@param syncName syncName path parameter. This name is used to update, delete or retrieve the specific Identitysync config.
*/
func (s *DuoIDentitySyncService) GetIDentitysyncBySyncName(syncName string) (*ResponseDuoIDentitySyncGetIDentitysyncBySyncName, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-identitysync/identitysync/{syncName}"
	path = strings.Replace(path, "{syncName}", fmt.Sprintf("%v", syncName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDuoIDentitySyncGetIDentitysyncBySyncName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIdentitysyncBySyncName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseDuoIDentitySyncGetIDentitysyncBySyncName)
	return result, response, err

}

//CreateIDentitysync Create a new IdentitySync configuration
/* Duo-IdentitySync Create a new IdentitySync configuration

 */
func (s *DuoIDentitySyncService) CreateIDentitysync(requestDuoIDentitySyncCreateIdentitysync *RequestDuoIDentitySyncCreateIDentitysync) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-identitysync/identitysync"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDuoIDentitySyncCreateIdentitysync).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateIdentitysync")
	}

	return response, err

}

//UpdateStatus Update Identity Sync Status
/* duo-identitysync update sync status.

@param syncName syncName path parameter. Sync connection to be updated
*/
func (s *DuoIDentitySyncService) UpdateStatus(syncName string, requestDuoIDentitySyncUpdateStatus *RequestDuoIDentitySyncUpdateStatus) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-identitysync/identitysync/status/{syncName}"
	path = strings.Replace(path, "{syncName}", fmt.Sprintf("%v", syncName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDuoIDentitySyncUpdateStatus).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateStatus")
	}

	getCSFRToken(response.Header())

	return response, err

}

//UpdateIDentitysyncBySyncName Update the specified Identitysync configuration
/* Duo-Identitysync Update the Identitysync configuration specified in the syncName.

@param syncName syncName path parameter. This name is used to update, delete or retrieve the specific Identitysync config.
*/
func (s *DuoIDentitySyncService) UpdateIDentitysyncBySyncName(syncName string, requestDuoIDentitySyncUpdateIdentitysyncBySyncName *RequestDuoIDentitySyncUpdateIDentitysyncBySyncName) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-identitysync/identitysync/{syncName}"
	path = strings.Replace(path, "{syncName}", fmt.Sprintf("%v", syncName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDuoIDentitySyncUpdateIdentitysyncBySyncName).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateIdentitysyncBySyncName")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteIDentitySyncBySyncName Delete the specified Identitysync configuration
/* Duo-Identitysync Delete the Identitysync configuration specified in the syncName.

@param syncName syncName path parameter. This name is used to update, delete or retrieve the specific Identitysync config.
*/
func (s *DuoIDentitySyncService) DeleteIDentitySyncBySyncName(syncName string) (*resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/duo-identitysync/identitysync/{syncName}"
	path = strings.Replace(path, "{syncName}", fmt.Sprintf("%v", syncName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteIdentitySyncBySyncName")
	}

	getCSFRToken(response.Header())

	return response, err

}
