package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type SessionDirectoryService service

//GetSessions 🚧 getSessions
/* 🚧 getSessions

 */
func (s *SessionDirectoryService) GetSessions() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/ise/session/getSessions"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation GetSessions")
	}

	return response, err

}

//GetSessionsForRecovery 🚧 getSessionsForRecovery
/* 🚧 getSessionsForRecovery

 */
func (s *SessionDirectoryService) GetSessionsForRecovery() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/ise/session/getSessionsForRecovery"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation GetSessionsForRecovery")
	}

	return response, err

}

//GetSessionByIPAddress 🚧 getSessionByIpAddress
/* 🚧 getSessionByIpAddress

 */
func (s *SessionDirectoryService) GetSessionByIPAddress() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/ise/session/getSessionByIpAddress"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation GetSessionByIpAddress")
	}

	return response, err

}

//GetSessionByMacAddress 🚧 getSessionByMacAddress
/* 🚧 getSessionByMacAddress

 */
func (s *SessionDirectoryService) GetSessionByMacAddress() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/ise/session/getSessionByMacAddress"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation GetSessionByMacAddress")
	}

	return response, err

}

//GetUserGroups 🚧 getUserGroups
/* 🚧 getUserGroups

 */
func (s *SessionDirectoryService) GetUserGroups() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/ise/session/getUserGroups"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation GetUserGroups")
	}

	return response, err

}

//GetUserGroupByUserName 🚧 getUserGroupByUserName
/* 🚧 getUserGroupByUserName

 */
func (s *SessionDirectoryService) GetUserGroupByUserName() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/ise/session/getUserGroupByUserName"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation GetUserGroupByUserName")
	}

	return response, err

}
