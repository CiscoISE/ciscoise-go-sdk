package isegosdk

import (
	"encoding/xml"

	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type MiscService service

type ResponseMiscGetActiveCount struct {
	Count *int `xml:"count,omitempty"` //
}

type ResponseMiscGetActiveList struct {
	NoOfActiveSession *int `xml:"noOfActiveSession,omitempty"` //
}

type ResponseMiscGetSessionAuthList struct {
	NoOfActiveSession *int `xml:"noOfActiveSession,omitempty"` //
}

type ResponseMiscGetPostureCount struct {
	Count *int `xml:"count,omitempty"` //
}

type ResponseMiscGetProfilerCount struct {
	Count *int `xml:"count,omitempty"` //
}

type ResponseMiscGetMntVersion struct {
	XMLName    xml.Name `xml:"product"`
	Name       string   `xml:"name,omitempty"`         //
	Version    string   `xml:"version,omitempty"`      //
	TypeOfNode *int     `xml:"type_of_node,omitempty"` //
}

//GetActiveCount ActiveCount
/* ActiveCount

 */
func (s *MiscService) GetActiveCount() (*ResponseMiscGetActiveCount, *resty.Response, error) {
	setHost(s.client, "_mnt")
	path := "/admin/API/mnt/Session/ActiveCount"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/xml").
		SetHeader("Accept", "application/xml").
		SetResult(&ResponseMiscGetActiveCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetActiveCount")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseMiscGetActiveCount)
	return result, response, err

}

//GetActiveList ActiveList
/* ActiveList

 */
func (s *MiscService) GetActiveList() (*ResponseMiscGetActiveList, *resty.Response, error) {
	setHost(s.client, "_mnt")
	path := "/admin/API/mnt/Session/ActiveList"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/xml").
		SetHeader("Accept", "application/xml").
		SetResult(&ResponseMiscGetActiveList{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetActiveList")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseMiscGetActiveList)
	return result, response, err

}

//GetSessionAuthList Session/AuthList
/* Session/AuthList

 */
func (s *MiscService) GetSessionAuthList() (*ResponseMiscGetSessionAuthList, *resty.Response, error) {
	setHost(s.client, "_mnt")
	path := "/admin/API/mnt/Session/AuthList/null/null"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/xml").
		SetHeader("Accept", "application/xml").
		SetResult(&ResponseMiscGetSessionAuthList{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSessionAuthList")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseMiscGetSessionAuthList)
	return result, response, err

}

//GetPostureCount PostureCount
/* PostureCount

 */
func (s *MiscService) GetPostureCount() (*ResponseMiscGetPostureCount, *resty.Response, error) {
	setHost(s.client, "_mnt")
	path := "/admin/API/mnt/Session/PostureCount"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/xml").
		SetHeader("Accept", "application/xml").
		SetResult(&ResponseMiscGetPostureCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPostureCount")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseMiscGetPostureCount)
	return result, response, err

}

//GetProfilerCount ProfilerCount
/* ProfilerCount

 */
func (s *MiscService) GetProfilerCount() (*ResponseMiscGetProfilerCount, *resty.Response, error) {
	setHost(s.client, "_mnt")
	path := "/admin/API/mnt/Session/ProfilerCount"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/xml").
		SetHeader("Accept", "application/xml").
		SetResult(&ResponseMiscGetProfilerCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetProfilerCount")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseMiscGetProfilerCount)
	return result, response, err

}

//GetSessionsByMac Sessions by MAC
/* Sessions by MAC

@param mac mac path parameter.
*/
func (s *MiscService) GetSessionsByMac(mac string) (*resty.Response, error) {
	setHost(s.client, "_mnt")
	path := "/admin/API/mnt/Session/MACAddress/{mac}"
	path = strings.Replace(path, "{mac}", fmt.Sprintf("%v", mac), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/xml").
		SetHeader("Accept", "application/xml").
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetSessionsByMac")
	}

	getCSFRToken(response.Header())

	return response, err

}

//GetSessionsByUsername Sessions by Username
/* Sessions by Username

@param username username path parameter.
*/
func (s *MiscService) GetSessionsByUsername(username string) (*resty.Response, error) {
	setHost(s.client, "_mnt")
	path := "/admin/API/mnt/Session/UserName/{username}"
	path = strings.Replace(path, "{username}", fmt.Sprintf("%v", username), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/xml").
		SetHeader("Accept", "application/xml").
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetSessionsByUsername")
	}

	getCSFRToken(response.Header())

	return response, err

}

//GetSessionsByNasIP Sessions by NAS IP
/* Sessions by NAS IP

@param nasipv4 nas_ipv4 path parameter.
*/
func (s *MiscService) GetSessionsByNasIP(nasipv4 string) (*resty.Response, error) {
	setHost(s.client, "_mnt")
	path := "/admin/API/mnt/Session/IPAddress/{nas_ipv4}"
	path = strings.Replace(path, "{nas_ipv4}", fmt.Sprintf("%v", nasipv4), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/xml").
		SetHeader("Accept", "application/xml").
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetSessionsByNasIp")
	}

	getCSFRToken(response.Header())

	return response, err

}

//GetSessionsByEndpointIP Sessions by Endpoint IP
/* Sessions by Endpoint IP

@param endpointipv4 endpoint_ipv4 path parameter.
*/
func (s *MiscService) GetSessionsByEndpointIP(endpointipv4 string) (*resty.Response, error) {
	setHost(s.client, "_mnt")
	path := "/admin/API/mnt/Session/EndPointIPAddress/{endpoint_ipv4}"
	path = strings.Replace(path, "{endpoint_ipv4}", fmt.Sprintf("%v", endpointipv4), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/xml").
		SetHeader("Accept", "application/xml").
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetSessionsByEndpointIp")
	}

	getCSFRToken(response.Header())

	return response, err

}

//GetSessionsBySessionID Sessions by SessionID
/* Sessions by SessionID

@param sessionTypeID session_id path parameter.
*/
func (s *MiscService) GetSessionsBySessionID(sessionTypeID string) (*resty.Response, error) {
	setHost(s.client, "_mnt")
	path := "/admin/API/mnt/Session/Active/SessionID/{session_id}/0"
	path = strings.Replace(path, "{session_id}", fmt.Sprintf("%v", sessionTypeID), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/xml").
		SetHeader("Accept", "application/xml").
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetSessionsBySessionId")
	}

	getCSFRToken(response.Header())

	return response, err

}

//GetMntVersion MNT Version
/* MNT Version

 */
func (s *MiscService) GetMntVersion() (*ResponseMiscGetMntVersion, *resty.Response, error) {
	setHost(s.client, "_mnt")
	path := "/admin/API/mnt/Version"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/xml").
		SetHeader("Accept", "application/xml").
		SetResult(&ResponseMiscGetMntVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetMntVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseMiscGetMntVersion)
	return result, response, err

}

//GetFailureReasons FailureReasons
/* FailureReasons

 */
func (s *MiscService) GetFailureReasons() (*resty.Response, error) {
	setHost(s.client, "_mnt")
	path := "/admin/API/mnt/FailureReasons"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/xml").
		SetHeader("Accept", "application/xml").
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetFailureReasons")
	}

	getCSFRToken(response.Header())

	return response, err

}

//GetAuthenticationStatusByMac AuthenticationStatus by MAC
/* AuthenticationStatus by MAC

@param MAC MAC path parameter.
@param SECONDS SECONDS path parameter.
@param RECORDS RECORDS path parameter.
*/
func (s *MiscService) GetAuthenticationStatusByMac(MAC string, SECONDS string, RECORDS string) (*resty.Response, error) {
	setHost(s.client, "_mnt")
	path := "/admin/API/mnt/AuthStatus/MACAddress/{MAC}/{SECONDS}/{RECORDS}/All"
	path = strings.Replace(path, "{MAC}", fmt.Sprintf("%v", MAC), -1)
	path = strings.Replace(path, "{SECONDS}", fmt.Sprintf("%v", SECONDS), -1)
	path = strings.Replace(path, "{RECORDS}", fmt.Sprintf("%v", RECORDS), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/xml").
		SetHeader("Accept", "application/xml").
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetAuthenticationStatusByMac")
	}

	getCSFRToken(response.Header())

	return response, err

}

//SessionReauthenticationByMac Session Reauthentication by MAC
/* Session Reauthentication by MAC

@param PSNNAME PSN_NAME path parameter.
@param ENDPOINTMAC ENDPOINT_MAC path parameter.
@param REAuthTYPE REAUTH_TYPE path parameter.
*/
func (s *MiscService) SessionReauthenticationByMac(PSNNAME string, ENDPOINTMAC string, REAuthTYPE string) (*resty.Response, error) {
	setHost(s.client, "_mnt")
	path := "/admin/API/mnt/CoA/Reauth/{PSN_NAME}/{ENDPOINT_MAC}/{REAUTH_TYPE}"
	path = strings.Replace(path, "{PSN_NAME}", fmt.Sprintf("%v", PSNNAME), -1)
	path = strings.Replace(path, "{ENDPOINT_MAC}", fmt.Sprintf("%v", ENDPOINTMAC), -1)
	path = strings.Replace(path, "{REAUTH_TYPE}", fmt.Sprintf("%v", REAuthTYPE), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/xml").
		SetHeader("Accept", "application/xml").
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation SessionReauthenticationByMac")
	}

	getCSFRToken(response.Header())

	return response, err

}

//SessionDisconnect Session Disconnect
/* Session Disconnect

@param ENDPOINTIP ENDPOINT_IP path parameter.
@param PSNNAME PSN_NAME path parameter.
@param MAC MAC path parameter.
@param DISCONNECTTYPE DISCONNECT_TYPE path parameter.
@param NASIPV4 NAS_IPV4 path parameter.
*/
func (s *MiscService) SessionDisconnect(ENDPOINTIP string, PSNNAME string, MAC string, DISCONNECTTYPE string, NASIPV4 string) (*resty.Response, error) {
	setHost(s.client, "_mnt")
	path := "/admin/API/mnt/CoA/Disconnect>/{PSN_NAME}/{MAC}/{DISCONNECT_TYPE}/{NAS_IPV4}/{{ENDPOINT_IP}}"
	path = strings.Replace(path, "{ENDPOINT_IP}", fmt.Sprintf("%v", ENDPOINTIP), -1)
	path = strings.Replace(path, "{PSN_NAME}", fmt.Sprintf("%v", PSNNAME), -1)
	path = strings.Replace(path, "{MAC}", fmt.Sprintf("%v", MAC), -1)
	path = strings.Replace(path, "{DISCONNECT_TYPE}", fmt.Sprintf("%v", DISCONNECTTYPE), -1)
	path = strings.Replace(path, "{NAS_IPV4}", fmt.Sprintf("%v", NASIPV4), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/xml").
		SetHeader("Accept", "application/xml").
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation SessionDisconnect")
	}

	getCSFRToken(response.Header())

	return response, err

}

//GetAccountStatusByMac AccountStatus by MAC
/* AccountStatus by MAC

@param mac mac path parameter.
@param duration duration path parameter.
*/
func (s *MiscService) GetAccountStatusByMac(mac string, duration string) (*resty.Response, error) {
	setHost(s.client, "_mnt")
	path := "/admin/API/mnt/AcctStatus/MACAddress/{mac}/{duration}"
	path = strings.Replace(path, "{mac}", fmt.Sprintf("%v", mac), -1)
	path = strings.Replace(path, "{duration}", fmt.Sprintf("%v", duration), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/xml").
		SetHeader("Accept", "application/xml").
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetAccountStatusByMac")
	}

	getCSFRToken(response.Header())

	return response, err

}

//DeleteAllSessions Delete All Sessions
/* Delete All Sessions

 */
func (s *MiscService) DeleteAllSessions() (*resty.Response, error) {
	setHost(s.client, "_mnt")
	path := "/admin/API/mnt/Session/Delete/All"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/xml").
		SetHeader("Accept", "application/xml").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteAllSessions")
	}

	getCSFRToken(response.Header())

	return response, err

}
