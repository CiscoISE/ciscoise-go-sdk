package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type BackupAndRestoreService service

type ResponseBackupAndRestoreConfigBackup struct {
	Response *ResponseBackupAndRestoreConfigBackupResponse `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  //
}

type ResponseBackupAndRestoreConfigBackupResponse struct {
	ID      string                                            `json:"id,omitempty"`      // Id which can be used to track the status of backup / restore of config DB.
	Message string                                            `json:"message,omitempty"` //
	Link    *ResponseBackupAndRestoreConfigBackupResponseLink `json:"link,omitempty"`    //
}

type ResponseBackupAndRestoreConfigBackupResponseLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseBackupAndRestoreCancelBackup struct {
	Response *ResponseBackupAndRestoreCancelBackupResponse `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  //
}

type ResponseBackupAndRestoreCancelBackupResponse struct {
	Action          string `json:"action,omitempty"`          // Whether backup job or restore job
	Details         string `json:"details,omitempty"`         // Details of the job
	Error           string `json:"error,omitempty"`           // Error message. False in case of no error
	HostName        string `json:"hostName,omitempty"`        // Hostname where the job has executed
	InitiatedFrom   string `json:"initiatedFrom,omitempty"`   // Whether the job was triggered from CLI / ADMIN UI / OPEN API
	JustComplete    string `json:"justComplete,omitempty"`    // Whether the job completed now. Possible values - Yes, No
	Message         string `json:"message,omitempty"`         // detail message in case of exception
	Name            string `json:"name,omitempty"`            // Backup name given at the time of scheduling the job.
	PercentComplete string `json:"percentComplete,omitempty"` // shows the percent completion of the job. Possible value range  1 - 100
	Repository      string `json:"repository,omitempty"`      // configured repository selected for the backup / restore job
	Scheduled       string `json:"scheduled,omitempty"`       // last action performed with respect to backup job. Possible values- BACKUP , CANCEL
	StartDate       string `json:"startDate,omitempty"`       // Start date of the backup job
	Status          string `json:"status,omitempty"`          // status of the job - in_progress, complete
	Type            string `json:"type,omitempty"`            // Whether the job is for OPERATIONAL_DB or CONFIGURATION_DB
}

type ResponseBackupAndRestoreGetLastConfigBackupStatus struct {
	Response *ResponseBackupAndRestoreGetLastConfigBackupStatusResponse `json:"response,omitempty"` //
	Version  string                                                     `json:"version,omitempty"`  //
}

type ResponseBackupAndRestoreGetLastConfigBackupStatusResponse struct {
	Action          string `json:"action,omitempty"`          // Whether backup job or restore job
	Details         string `json:"details,omitempty"`         // Details of the job
	Error           string `json:"error,omitempty"`           // Error message. False in case of no error
	HostName        string `json:"hostName,omitempty"`        // Hostname where the job has executed
	InitiatedFrom   string `json:"initiatedFrom,omitempty"`   // Whether the job was triggered from CLI / ADMIN UI / OPEN API
	JustComplete    string `json:"justComplete,omitempty"`    // Whether the job completed now. Possible values - Yes, No
	Message         string `json:"message,omitempty"`         // detail message in case of exception
	Name            string `json:"name,omitempty"`            // Backup name given at the time of scheduling the job.
	PercentComplete string `json:"percentComplete,omitempty"` // shows the percent completion of the job. Possible value range  1 - 100
	Repository      string `json:"repository,omitempty"`      // configured repository selected for the backup / restore job
	Scheduled       string `json:"scheduled,omitempty"`       // last action performed with respect to backup job. Possible values- BACKUP , CANCEL
	StartDate       string `json:"startDate,omitempty"`       // Start date of the backup job
	Status          string `json:"status,omitempty"`          // status of the job - in_progress, complete
	Type            string `json:"type,omitempty"`            // Whether the job is for OPERATIONAL_DB or CONFIGURATION_DB
}

type ResponseBackupAndRestoreRestoreConfigBackup struct {
	Response *ResponseBackupAndRestoreRestoreConfigBackupResponse `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  //
}

type ResponseBackupAndRestoreRestoreConfigBackupResponse struct {
	ID      string                                                   `json:"id,omitempty"`      // Id which can be used to track the status of backup / restore of config DB.
	Message string                                                   `json:"message,omitempty"` //
	Link    *ResponseBackupAndRestoreRestoreConfigBackupResponseLink `json:"link,omitempty"`    //
}

type ResponseBackupAndRestoreRestoreConfigBackupResponseLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseBackupAndRestoreUpdateScheduledConfigBackup struct {
	Response *ResponseBackupAndRestoreUpdateScheduledConfigBackupResponse `json:"response,omitempty"` //
	Version  string                                                       `json:"version,omitempty"`  //
}

type ResponseBackupAndRestoreUpdateScheduledConfigBackupResponse struct {
	Message string                                                           `json:"message,omitempty"` // Response message on successful scheduling the backup job.
	Link    *ResponseBackupAndRestoreUpdateScheduledConfigBackupResponseLink `json:"link,omitempty"`    //
}

type ResponseBackupAndRestoreUpdateScheduledConfigBackupResponseLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseBackupAndRestoreCreateScheduledConfigBackup struct {
	Response *ResponseBackupAndRestoreCreateScheduledConfigBackupResponse `json:"response,omitempty"` //
	Version  string                                                       `json:"version,omitempty"`  //
}

type ResponseBackupAndRestoreCreateScheduledConfigBackupResponse struct {
	Message string                                                           `json:"message,omitempty"` // Response message on successful scheduling the backup job.
	Link    *ResponseBackupAndRestoreCreateScheduledConfigBackupResponseLink `json:"link,omitempty"`    //
}

type ResponseBackupAndRestoreCreateScheduledConfigBackupResponseLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestBackupAndRestoreConfigBackup struct {
	BackupEncryptionKey string `json:"backupEncryptionKey,omitempty"` // The encyption key for the backed up file. Encryption key must satisfy the following criteria - Contains at least one uppercase letter [A-Z], Contains at least one lowercase letter [a-z], Contains at least one digit [0-9], Contain only [A-Z][a-z][0-9]_#, Has at least 8 characters, Has not more than 15 characters, Must not contain 'CcIiSsCco', Must not begin with
	BackupName          string `json:"backupName,omitempty"`          // The backup file will get saved with this name.
	RepositoryName      string `json:"repositoryName,omitempty"`      // Name of the configured repository where the generated backup file will get copied.
}

type RequestBackupAndRestoreRestoreConfigBackup struct {
	BackupEncryptionKey string `json:"backupEncryptionKey,omitempty"` // The encryption key which was provided at the time of taking backup.
	RepositoryName      string `json:"repositoryName,omitempty"`      // Name of the configred repository where the backup file exists.
	RestoreFile         string `json:"restoreFile,omitempty"`         // Name of the backup file to be restored on ISE node.
	RestoreIncludeAdeos string `json:"restoreIncludeAdeos,omitempty"` // Determines whether the ADE-OS configure is restored. Possible values true, false
}

type RequestBackupAndRestoreUpdateScheduledConfigBackup struct {
	BackupDescription   string `json:"backupDescription,omitempty"`   // Description of the backup.
	BackupEncryptionKey string `json:"backupEncryptionKey,omitempty"` // The encyption key for the backed up file. Encryption key must satisfy the following criteria - Contains at least one uppercase letter [A-Z], Contains at least one lowercase letter [a-z], Contains at least one digit [0-9], Contain only [A-Z][a-z][0-9]_#, Has at least 8 characters, Has not more than 15 characters, Must not contain 'CcIiSsCco', Must not begin with
	BackupName          string `json:"backupName,omitempty"`          // The backup file will get saved with this name.
	EndDate             string `json:"endDate,omitempty"`             // End date of the scheduled backup job. Allowed format MM/DD/YYYY. End date is not required in case of ONE_TIME frequency.
	Frequency           string `json:"frequency,omitempty"`           // Frequency with which the backup will get scheduled in the ISE node.
	MonthDay            string `json:"monthDay,omitempty"`            // Day of month you want backup to be performed on when scheduled frequency is MONTHLY. Allowed values - from 1 to 28.
	RepositoryName      string `json:"repositoryName,omitempty"`      // Name of the configured repository where the generated backup file will get copied.
	StartDate           string `json:"startDate,omitempty"`           // Start date for scheduling the backup job. Allowed format MM/DD/YYYY.
	Status              string `json:"status,omitempty"`              // Enable or disable scheduled backup.
	Time                string `json:"time,omitempty"`                // Time at which backup job get scheduled. example- 12:00 AM
	WeekDay             string `json:"weekDay,omitempty"`             // Day of week you want backup to be performed on when scheduled frequency is WEEKLY
}

type RequestBackupAndRestoreCreateScheduledConfigBackup struct {
	BackupDescription   string `json:"backupDescription,omitempty"`   // Description of the backup.
	BackupEncryptionKey string `json:"backupEncryptionKey,omitempty"` // The encyption key for the backed up file. Encryption key must satisfy the following criteria - Contains at least one uppercase letter [A-Z], Contains at least one lowercase letter [a-z], Contains at least one digit [0-9], Contain only [A-Z][a-z][0-9]_#, Has at least 8 characters, Has not more than 15 characters, Must not contain 'CcIiSsCco', Must not begin with
	BackupName          string `json:"backupName,omitempty"`          // The backup file will get saved with this name.
	EndDate             string `json:"endDate,omitempty"`             // End date of the scheduled backup job. Allowed format MM/DD/YYYY. End date is not required in case of ONE_TIME frequency.
	Frequency           string `json:"frequency,omitempty"`           // Frequency with which the backup will get scheduled in the ISE node.
	MonthDay            string `json:"monthDay,omitempty"`            // Day of month you want backup to be performed on when scheduled frequency is MONTHLY. Allowed values - from 1 to 28.
	RepositoryName      string `json:"repositoryName,omitempty"`      // Name of the configured repository where the generated backup file will get copied.
	StartDate           string `json:"startDate,omitempty"`           // Start date for scheduling the backup job. Allowed format MM/DD/YYYY.
	Status              string `json:"status,omitempty"`              // Enable or disable scheduled backup.
	Time                string `json:"time,omitempty"`                // Time at which backup job get scheduled. example- 12:00 AM
	WeekDay             string `json:"weekDay,omitempty"`             // Day of week you want backup to be performed on when scheduled frequency is WEEKLY
}

//GetLastConfigBackupStatus Gives the last backup status
/* Gives the last backup status

 */
func (s *BackupAndRestoreService) GetLastConfigBackupStatus() (*ResponseBackupAndRestoreGetLastConfigBackupStatus, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/backup-restore/config/last-backup-status"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseBackupAndRestoreGetLastConfigBackupStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetLastConfigBackupStatus")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseBackupAndRestoreGetLastConfigBackupStatus)
	return result, response, err

}

//ConfigBackup Take the config DB backup now by providing the name of the backup,repository name and encryption key. The API returns the task ID. Use the Task Service status API to get the status of the backup job
/* Triggers on demand configuration backup on the ISE node. The API returns the task ID. Use the Task Service status API to get the status of the backup job.

 */
func (s *BackupAndRestoreService) ConfigBackup(requestBackupAndRestoreConfigBackup *RequestBackupAndRestoreConfigBackup) (*ResponseBackupAndRestoreConfigBackup, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/backup-restore/config/backup"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestBackupAndRestoreConfigBackup).
		SetResult(&ResponseBackupAndRestoreConfigBackup{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ConfigBackup")
	}

	result := response.Result().(*ResponseBackupAndRestoreConfigBackup)
	return result, response, err

}

//CancelBackup Cancel the running backup
/* Cancels the backup job running on the node.

 */
func (s *BackupAndRestoreService) CancelBackup() (*ResponseBackupAndRestoreCancelBackup, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/backup-restore/config/cancel-backup"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseBackupAndRestoreCancelBackup{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CancelBackup")
	}

	result := response.Result().(*ResponseBackupAndRestoreCancelBackup)
	return result, response, err

}

//RestoreConfigBackup Restore a config DB backup by giving the name of the backup file, repository name and encryption key. The API returns the task ID. Use the Task Service status API to get the status of the restore job
/* Triggers a configuration DB restore job on the ISE node. The API returns the task ID. Use the Task Service status API to get the status of the backup job

 */
func (s *BackupAndRestoreService) RestoreConfigBackup(requestBackupAndRestoreRestoreConfigBackup *RequestBackupAndRestoreRestoreConfigBackup) (*ResponseBackupAndRestoreRestoreConfigBackup, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/backup-restore/config/restore"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestBackupAndRestoreRestoreConfigBackup).
		SetResult(&ResponseBackupAndRestoreRestoreConfigBackup{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RestoreConfigBackup")
	}

	result := response.Result().(*ResponseBackupAndRestoreRestoreConfigBackup)
	return result, response, err

}

//CreateScheduledConfigBackup Schedules the configuration backup on the ISE node as per the input parameters.
/* Schedules the configuration backup on the ISE node as per the input parameters. This API helps in creating the schedule for the first time.

 */
func (s *BackupAndRestoreService) CreateScheduledConfigBackup(requestBackupAndRestoreCreateScheduledConfigBackup *RequestBackupAndRestoreCreateScheduledConfigBackup) (*ResponseBackupAndRestoreCreateScheduledConfigBackup, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/backup-restore/config/schedule-config-backup"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestBackupAndRestoreCreateScheduledConfigBackup).
		SetResult(&ResponseBackupAndRestoreCreateScheduledConfigBackup{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateScheduledConfigBackup")
	}

	result := response.Result().(*ResponseBackupAndRestoreCreateScheduledConfigBackup)
	return result, response, err

}

//UpdateScheduledConfigBackup Update scheduled configuration backup
/* Update the Schedule of the configuration backup on the ISE node as per the input parameters. This API only helps in editing the schedule.

 */
func (s *BackupAndRestoreService) UpdateScheduledConfigBackup(requestBackupAndRestoreUpdateScheduledConfigBackup *RequestBackupAndRestoreUpdateScheduledConfigBackup) (*ResponseBackupAndRestoreUpdateScheduledConfigBackup, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/backup-restore/config/schedule-config-backup"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestBackupAndRestoreUpdateScheduledConfigBackup).
		SetResult(&ResponseBackupAndRestoreUpdateScheduledConfigBackup{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateScheduledConfigBackup")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseBackupAndRestoreUpdateScheduledConfigBackup)
	return result, response, err

}
