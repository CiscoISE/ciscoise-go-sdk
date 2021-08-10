package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type PullDeploymentInfoService service

type ResponsePullDeploymentInfoGetDeploymentInfo struct {
	ERSDeploymentInfo ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfo `json:"ERSDeploymentInfo,omitempty"` //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfo struct {
	NetworkAccessInfo ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoNetworkAccessInfo `json:"networkAccessInfo,omitempty"` // networkAccessInfo Details
	ProfilerInfo      ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoProfilerInfo      `json:"profilerInfo,omitempty"`      // profilerInfo Details
	DeploymentInfo    ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoDeploymentInfo    `json:"deploymentInfo,omitempty"`    //
	NadInfo           ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoNadInfo           `json:"nadInfo,omitempty"`           //
	MdmInfo           ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoMdmInfo           `json:"mdmInfo,omitempty"`           //
	LicensesInfo      ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoLicensesInfo      `json:"licensesInfo,omitempty"`      //
	PostureInfo       ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoPostureInfo       `json:"postureInfo,omitempty"`       //
	KongInfo          ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoKongInfo          `json:"kongInfo,omitempty"`          //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoNetworkAccessInfo struct {
	DeploymentID    string                                                                                `json:"deploymentID,omitempty"`    // Deployment ID
	IsCsnEnabled    bool                                                                                  `json:"isCsnEnabled,omitempty"`    //
	NodeList        ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoNetworkAccessInfoNodeList `json:"nodeList,omitempty"`        //
	SdaVns          []interface{}                                                                         `json:"sdaVNs,omitempty"`          //
	TrustSecControl string                                                                                `json:"trustSecControl,omitempty"` //
	Radius3RdParty  []interface{}                                                                         `json:"radius3RdParty,omitempty"`  //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoNetworkAccessInfoNodeList struct {
	NodeAndScope []interface{} `json:"nodeAndScope,omitempty"` //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoProfilerInfo struct {
	DeploymentID string                                                                           `json:"deploymentID,omitempty"` // Deployment ID
	NodeList     ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoProfilerInfoNodeList `json:"nodeList,omitempty"`     // Deployment ID
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoProfilerInfoNodeList struct {
	Node []ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoProfilerInfoNodeListNode `json:"node,omitempty"` //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoProfilerInfoNodeListNode struct {
	OnlineSubscriptionEnabled bool        `json:"onlineSubscriptionEnabled,omitempty"` //
	LastAppliedFeedDateTime   string      `json:"lastAppliedFeedDateTime,omitempty"`   //
	Scope                     string      `json:"scope,omitempty"`                     //
	Profiles                  interface{} `json:"profiles,omitempty"`                  //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoDeploymentInfo struct {
	DeploymentID       string                                                                                         `json:"deploymentID,omitempty"`       //
	VersionHistoryInfo []ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoDeploymentInfoVersionHistoryInfo `json:"versionHistoryInfo,omitempty"` //
	NodeList           ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoDeploymentInfoNodeList             `json:"nodeList,omitempty"`           //
	Fipsstatus         string                                                                                         `json:"fipsstatus,omitempty"`         //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoDeploymentInfoVersionHistoryInfo struct {
	OpType      string `json:"opType,omitempty"`      //
	MainVersion string `json:"mainVersion,omitempty"` //
	EpochTime   int    `json:"epochTime,omitempty"`   //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoDeploymentInfoNodeList struct {
	NodeAndNodeCountAndCountInfo []ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoDeploymentInfoNodeListNodeAndNodeCountAndCountInfo `json:"nodeAndNodeCountAndCountInfo,omitempty"` //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoDeploymentInfoNodeListNodeAndNodeCountAndCountInfo struct {
	Name            string      `json:"name,omitempty"`            //
	Value           interface{} `json:"value,omitempty"`           // It varies type
	DeclaredType    string      `json:"declaredType,omitempty"`    //
	Scope           string      `json:"scope,omitempty"`           //
	Nil             bool        `json:"nil,omitempty"`             //
	GlobalScope     bool        `json:"globalScope,omitempty"`     //
	TypeSubstituted bool        `json:"typeSubstituted,omitempty"` //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoNadInfo struct {
	NodeList     ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoNadInfoNodeList     `json:"nodeList,omitempty"`     //
	NadcountInfo ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoNadInfoNadcountInfo `json:"nadcountInfo,omitempty"` //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoNadInfoNodeList struct {
	NodeAndScope []interface{} `json:"nodeAndScope,omitempty"` //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoNadInfoNadcountInfo struct {
	TotalActiveNADCount int `json:"totalActiveNADCount,omitempty"` //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoMdmInfo struct {
	ActiveMdmServersCount        string                                                                      `json:"activeMdmServersCount,omitempty"`        //
	ActiveDesktopMdmServersCount string                                                                      `json:"activeDesktopMdmServersCount,omitempty"` //
	ActiveMobileMdmServersCount  string                                                                      `json:"activeMobileMdmServersCount,omitempty"`  //
	DeploymentID                 string                                                                      `json:"deploymentID,omitempty"`                 //
	NodeList                     ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoMdmInfoNodeList `json:"nodeList,omitempty"`                     //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoMdmInfoNodeList struct {
	NodeAndScope []interface{} `json:"nodeAndScope,omitempty"` //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoLicensesInfo struct {
	DeploymentID string                                                                           `json:"deploymentID,omitempty"` //
	NodeList     ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoLicensesInfoNodeList `json:"nodeList,omitempty"`     //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoLicensesInfoNodeList struct {
	Node []interface{} `json:"node,omitempty"` //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoPostureInfo struct {
	Content []ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoPostureInfoContent `json:"content,omitempty"` //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoPostureInfoContent struct {
	Name            string      `json:"name,omitempty"`            //
	Value           interface{} `json:"value,omitempty"`           // It varies type
	DeclaredType    string      `json:"declaredType,omitempty"`    //
	Scope           string      `json:"scope,omitempty"`           //
	Nil             bool        `json:"nil,omitempty"`             //
	GlobalScope     bool        `json:"globalScope,omitempty"`     //
	TypeSubstituted bool        `json:"typeSubstituted,omitempty"` //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoKongInfo struct {
	DeploymentID string                                                                       `json:"deploymentID,omitempty"` //
	NodeList     ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoKongInfoNodeList `json:"nodeList,omitempty"`     //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoKongInfoNodeList struct {
	Node []ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoKongInfoNodeListNode `json:"node,omitempty"` //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoKongInfoNodeListNode struct {
	Sn      string                                                                                    `json:"sn,omitempty"`      //
	Service []ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoKongInfoNodeListNodeService `json:"service,omitempty"` //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoKongInfoNodeListNodeService struct {
	ServiceName string                                                                                         `json:"serviceName,omitempty"` //
	Route       []ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoKongInfoNodeListNodeServiceRoute `json:"route,omitempty"`       //
}

type ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoKongInfoNodeListNodeServiceRoute struct {
	RouteName    string      `json:"routeName,omitempty"`    //
	HTTPCount    interface{} `json:"httpCount,omitempty"`    //
	LatencyCount interface{} `json:"latencyCount,omitempty"` //
	LatencySum   interface{} `json:"latencySum,omitempty"`   //
}

type ResponsePullDeploymentInfoGetVersion struct {
	VersionInfo ResponsePullDeploymentInfoGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponsePullDeploymentInfoGetVersionVersionInfo struct {
	CurrentServerVersion string                                              `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                              `json:"supportedVersions,omitempty"`    //
	Link                 ResponsePullDeploymentInfoGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponsePullDeploymentInfoGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

//GetDeploymentInfo Pull deployment information
/* This API allows the client to pull the deployment information.

 */
func (s *PullDeploymentInfoService) GetDeploymentInfo() (*ResponsePullDeploymentInfoGetDeploymentInfo, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/deploymentinfo/getAllInfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePullDeploymentInfoGetDeploymentInfo{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeploymentInfo")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePullDeploymentInfoGetDeploymentInfo)
	return result, response, err

}

//GetVersion Get pull deployment info version information
/* This API helps to retrieve the version information related to the pull deployment info.

 */
func (s *PullDeploymentInfoService) GetVersion() (*ResponsePullDeploymentInfoGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/deploymentinfo/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePullDeploymentInfoGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponsePullDeploymentInfoGetVersion)
	return result, response, err

}
