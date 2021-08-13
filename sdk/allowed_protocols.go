package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type AllowedProtocolsService service

type GetAllowedProtocolsQueryParams struct {
	Page int `url:"page,omitempty"` //Page number
	Size int `url:"size,omitempty"` //Number of objects returned per page
}

type ResponseAllowedProtocolsGetAllowedProtocolByName struct {
	AllowedProtocols ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocols `json:"AllowedProtocols,omitempty"` //
}

type ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocols struct {
	ID                        string                                                                  `json:"id,omitempty"`                        // Resource UUID, Mandatory for update
	Name                      string                                                                  `json:"name,omitempty"`                      // Resource Name
	Description               string                                                                  `json:"description,omitempty"`               //
	EapTls                    ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsEapTls  `json:"eapTls,omitempty"`                    // The eapTls is required only if allowEapTls is true, otherwise it must be ignored. The object eapTls contains the settings for EAP TLS protocol
	Peap                      ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsPeap    `json:"peap,omitempty"`                      //
	EapFast                   ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsEapFast `json:"eapFast,omitempty"`                   // The eapFast is required only if allowEapFast is true, otherwise it must be ignored. The object eapFast contains the settings for EAP FAST protocol
	EapTtls                   ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsEapTtls `json:"eapTtls,omitempty"`                   // The eapTtls is required only if allowEapTtls is true, otherwise it must be ignored. The object eapTtls contains the settings for EAP TTLS protocol
	Teap                      ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsTeap    `json:"teap,omitempty"`                      // The teap is required only if allowTeap is true, otherwise it must be ignored. The object teap contains the settings for TEAP protocol
	ProcessHostLookup         bool                                                                    `json:"processHostLookup,omitempty"`         //
	AllowPapAscii             bool                                                                    `json:"allowPapAscii,omitempty"`             //
	AllowChap                 bool                                                                    `json:"allowChap,omitempty"`                 //
	AllowMsChapV1             bool                                                                    `json:"allowMsChapV1,omitempty"`             //
	AllowMsChapV2             bool                                                                    `json:"allowMsChapV2,omitempty"`             //
	AllowEapMd5               bool                                                                    `json:"allowEapMd5,omitempty"`               //
	AllowLeap                 bool                                                                    `json:"allowLeap,omitempty"`                 //
	AllowEapTls               bool                                                                    `json:"allowEapTls,omitempty"`               //
	AllowEapTtls              bool                                                                    `json:"allowEapTtls,omitempty"`              //
	AllowEapFast              bool                                                                    `json:"allowEapFast,omitempty"`              //
	AllowPeap                 bool                                                                    `json:"allowPeap,omitempty"`                 //
	AllowTeap                 bool                                                                    `json:"allowTeap,omitempty"`                 //
	AllowPreferredEapProtocol bool                                                                    `json:"allowPreferredEapProtocol,omitempty"` //
	PreferredEapProtocol      string                                                                  `json:"preferredEapProtocol,omitempty"`      // The preferredEapProtocol is required only if allowPreferredEapProtocol is true, otherwise it must be ignored. Allowed Values:  - EAP_FAST, - PEAP, - LEAP, - EAP_MD5, - EAP_TLS, - EAP_TTLS, - TEAP
	EapTlsLBit                bool                                                                    `json:"eapTlsLBit,omitempty"`                //
	AllowWeakCiphersForEap    bool                                                                    `json:"allowWeakCiphersForEap,omitempty"`    //
	RequireMessageAuth        bool                                                                    `json:"requireMessageAuth,omitempty"`        //
	Link                      ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsLink    `json:"link,omitempty"`                      //
}

type ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsEapTls struct {
	AllowEapTlsAuthOfExpiredCerts      bool   `json:"allowEapTlsAuthOfExpiredCerts,omitempty"`      //
	EapTlsEnableStatelessSessionResume bool   `json:"eapTlsEnableStatelessSessionResume,omitempty"` //
	EapTlsSessionTicketTtl             int    `json:"eapTlsSessionTicketTtl,omitempty"`             // Time to live. The eapTlsSessionTicketTtl is required only if eapTlsEnableStatelessSessionResume is true, otherwise it must be ignored
	EapTlsSessionTicketTtlUnits        string `json:"eapTlsSessionTicketTtlUnits,omitempty"`        // Time to live time units. The eapTlsSessionTicketTtlUnits is required only if eapTlsEnableStatelessSessionResume is true, otherwise it must be ignored. Allowed Values: - SECONDS, - MINUTES, - HOURS, - DAYS, - WEEKS
	EapTlsSessionTicketPrecentage      int    `json:"eapTlsSessionTicketPrecentage,omitempty"`      // The eapTlsSessionTicketPrecentage is required only if eapTlsEnableStatelessSessionResume is true, otherwise it must be ignored
}

type ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsPeap struct {
	AllowPeapEapMsChapV2                 bool `json:"allowPeapEapMsChapV2,omitempty"`                 //
	AllowPeapEapMsChapV2PwdChange        bool `json:"allowPeapEapMsChapV2PwdChange,omitempty"`        // The allowPeapEapMsChapV2PwdChange is required only if allowPeapEapMsChapV2 is true, otherwise it must be ignored
	AllowPeapEapMsChapV2PwdChangeRetries int  `json:"allowPeapEapMsChapV2PwdChangeRetries,omitempty"` // The allowPeapEapMsChapV2PwdChangeRetries is required only if allowPeapEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3
	AllowPeapEapGtc                      bool `json:"allowPeapEapGtc,omitempty"`                      //
	AllowPeapEapGtcPwdChange             bool `json:"allowPeapEapGtcPwdChange,omitempty"`             // The allowPeapEapGtcPwdChange is required only if allowPeapEapGtc is true, otherwise it must be ignored
	AllowPeapEapGtcPwdChangeRetries      int  `json:"allowPeapEapGtcPwdChangeRetries,omitempty"`      // The allowPeapEapGtcPwdChangeRetries is required only if allowPeapEapGtc is true, otherwise it must be ignored. Valid range is 0-3
	AllowPeapEapTls                      bool `json:"allowPeapEapTls,omitempty"`                      //
	AllowPeapEapTlsAuthOfExpiredCerts    bool `json:"allowPeapEapTlsAuthOfExpiredCerts,omitempty"`    // The allowPeapEapTlsAuthOfExpiredCerts is required only if allowPeapEapTls is true, otherwise it must be ignored
	RequireCryptobinding                 bool `json:"requireCryptobinding,omitempty"`                 //
	AllowPeapV0                          bool `json:"allowPeapV0,omitempty"`                          //
}

type ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsEapFast struct {
	AllowEapFastEapMsChapV2                                        bool   `json:"allowEapFastEapMsChapV2,omitempty"`                                        //
	AllowEapFastEapMsChapV2PwdChange                               bool   `json:"allowEapFastEapMsChapV2PwdChange,omitempty"`                               // The allowEapFastEapMsChapV2PwdChange is required only if allowEapFastEapMsChapV2 is true, otherwise it must be ignored
	AllowEapFastEapMsChapV2PwdChangeRetries                        int    `json:"allowEapFastEapMsChapV2PwdChangeRetries,omitempty"`                        // The allowEapFastEapMsChapV2PwdChangeRetries is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3
	AllowEapFastEapGtc                                             bool   `json:"allowEapFastEapGtc,omitempty"`                                             //
	AllowEapFastEapGtcPwdChange                                    bool   `json:"allowEapFastEapGtcPwdChange,omitempty"`                                    // The allowEapFastEapGtcPwdChange is required only if allowEapFastEapGtc is true, otherwise it must be ignored
	AllowEapFastEapGtcPwdChangeRetries                             int    `json:"allowEapFastEapGtcPwdChangeRetries,omitempty"`                             // The allowEapFastEapGtcPwdChangeRetries is required only if allowEapFastEapGtc is true, otherwise it must be ignored. Valid range is 0-3
	AllowEapFastEapTls                                             bool   `json:"allowEapFastEapTls,omitempty"`                                             //
	AllowEapFastEapTlsAuthOfExpiredCerts                           bool   `json:"allowEapFastEapTlsAuthOfExpiredCerts,omitempty"`                           // The allowEapFastEapTlsAuthOfExpiredCerts is required only if allowEapFastEapTls is true, otherwise it must be ignored
	EapFastUsePacs                                                 bool   `json:"eapFastUsePacs,omitempty"`                                                 //
	EapFastUsePacsTunnelPacTtl                                     int    `json:"eapFastUsePacsTunnelPacTtl,omitempty"`                                     // The eapFastUsePacsTunnelPacTtl is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsTunnelPacTtlUnits                                string `json:"eapFastUsePacsTunnelPacTtlUnits,omitempty"`                                // The eapFastUsePacsTunnelPacTtlUnits is required only if eapFastUsePacs is true, otherwise it must be ignored. Allowed Values: - SECONDS, - MINUTES, - HOURS, - DAYS, - WEEKS
	EapFastUsePacsUseProactivePacUpdatePrecentage                  int    `json:"eapFastUsePacsUseProactivePacUpdatePrecentage,omitempty"`                  // The eapFastUsePacsUseProactivePacUpdatePrecentage is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsAllowAnonymProvisioning                          bool   `json:"eapFastUsePacsAllowAnonymProvisioning,omitempty"`                          // The eapFastUsePacsAllowAnonymProvisioning is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsAllowAuthenProvisioning                          bool   `json:"eapFastUsePacsAllowAuthenProvisioning,omitempty"`                          // The eapFastUsePacsAllowAuthenProvisioning is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning bool   `json:"eapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning,omitempty"` // The eapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning is required only if eapFastUsePacsAllowAuthenProvisioning is true, otherwise it must be ignored
	EapFastUsePacsAcceptClientCert                                 bool   `json:"eapFastUsePacsAcceptClientCert,omitempty"`                                 // The eapFastUsePacsAcceptClientCert is required only if eapFastUsePacsAllowAuthenProvisioning is true, otherwise it must be ignored
	EapFastUsePacsMachinePacTtl                                    int    `json:"eapFastUsePacsMachinePacTtl,omitempty"`                                    // The eapFastUsePacsMachinePacTtl is required only if eapFastUsePacsAllowMachineAuthentication is true, otherwise it must be ignored
	EapFastUsePacsMachinePacTtlUnits                               string `json:"eapFastUsePacsMachinePacTtlUnits,omitempty"`                               // The eapFastUsePacsMachinePacTtlUnits is required only if eapFastUsePacsAllowMachineAuthentication is true, otherwise it must be ignored. Allowed Values: - SECONDS, - MINUTES, - HOURS, - DAYS, - WEEKS
	EapFastUsePacsAllowMachineAuthentication                       bool   `json:"eapFastUsePacsAllowMachineAuthentication,omitempty"`                       //
	EapFastUsePacsStatelessSessionResume                           bool   `json:"eapFastUsePacsStatelessSessionResume,omitempty"`                           // The eapFastUsePacsStatelessSessionResume is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsAuthorizationPacTtl                              int    `json:"eapFastUsePacsAuthorizationPacTtl,omitempty"`                              // The eapFastUsePacsAuthorizationPacTtl is required only if eapFastUsePacsStatelessSessionResume is true, otherwise it must be ignored
	EapFastUsePacsAuthorizationPacTtlUnits                         string `json:"eapFastUsePacsAuthorizationPacTtlUnits,omitempty"`                         // The eapFastUsePacsAuthorizationPacTtlUnits is required only if eapFastUsePacsStatelessSessionResume is true, otherwise it must be ignored. Allowed Values: - SECONDS, - MINUTES, - HOURS, - DAYS, - WEEKS
	EapFastDontUsePacsAcceptClientCert                             bool   `json:"eapFastDontUsePacsAcceptClientCert,omitempty"`                             // The eapFastDontUsePacsAcceptClientCert is required only if eapFastUsePacs is FALSE, otherwise it must be ignored
	EapFastDontUsePacsAllowMachineAuthentication                   bool   `json:"eapFastDontUsePacsAllowMachineAuthentication,omitempty"`                   // The eapFastDontUsePacsAllowMachineAuthentication is required only if eapFastUsePacs is FALSE, otherwise it must be ignored
	EapFastEnableEApChaining                                       bool   `json:"eapFastEnableEAPChaining,omitempty"`                                       //
}

type ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsEapTtls struct {
	EapTtlsPapAscii                    bool `json:"eapTtlsPapAscii,omitempty"`                    //
	EapTtlsChap                        bool `json:"eapTtlsChap,omitempty"`                        //
	EapTtlsMsChapV1                    bool `json:"eapTtlsMsChapV1,omitempty"`                    //
	EapTtlsMsChapV2                    bool `json:"eapTtlsMsChapV2,omitempty"`                    //
	EapTtlsEapMd5                      bool `json:"eapTtlsEapMd5,omitempty"`                      //
	EapTtlsEapMsChapV2                 bool `json:"eapTtlsEapMsChapV2,omitempty"`                 //
	EapTtlsEapMsChapV2PwdChange        bool `json:"eapTtlsEapMsChapV2PwdChange,omitempty"`        // The eapTtlsEapMsChapV2PwdChange is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored
	EapTtlsEapMsChapV2PwdChangeRetries int  `json:"eapTtlsEapMsChapV2PwdChangeRetries,omitempty"` // The eapTtlsEapMsChapV2PwdChangeRetries is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3
}

type ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsTeap struct {
	AllowTeapEapMsChapV2                 bool `json:"allowTeapEapMsChapV2,omitempty"`                 //
	AllowTeapEapMsChapV2PwdChange        bool `json:"allowTeapEapMsChapV2PwdChange,omitempty"`        // The allowTeapEapMsChapV2PwdChange is required only if allowTeapEapMsChapV2 is true, otherwise it must be ignored
	AllowTeapEapMsChapV2PwdChangeRetries int  `json:"allowTeapEapMsChapV2PwdChangeRetries,omitempty"` // The allowTeapEapMsChapV2PwdChangeRetries is required only if allowTeapEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3
	AllowTeapEapTls                      bool `json:"allowTeapEapTls,omitempty"`                      //
	AllowTeapEapTlsAuthOfExpiredCerts    bool `json:"allowTeapEapTlsAuthOfExpiredCerts,omitempty"`    // The allowTeapEapTlsAuthOfExpiredCerts is required only if allowTeapEapTls is true, otherwise it must be ignored
	AcceptClientCertDuringTunnelEst      bool `json:"acceptClientCertDuringTunnelEst,omitempty"`      //
	EnableEapChaining                    bool `json:"enableEapChaining,omitempty"`                    //
	AllowDowngradeMsk                    bool `json:"allowDowngradeMsk,omitempty"`                    //
}

type ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAllowedProtocolsGetAllowedProtocolByID struct {
	AllowedProtocols ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocols `json:"AllowedProtocols,omitempty"` //
}

type ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocols struct {
	ID                        string                                                                `json:"id,omitempty"`                        // Resource UUID, Mandatory for update
	Name                      string                                                                `json:"name,omitempty"`                      // Resource Name
	Description               string                                                                `json:"description,omitempty"`               //
	EapTls                    ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsEapTls  `json:"eapTls,omitempty"`                    // The eapTls is required only if allowEapTls is true, otherwise it must be ignored. The object eapTls contains the settings for EAP TLS protocol
	Peap                      ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsPeap    `json:"peap,omitempty"`                      //
	EapFast                   ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsEapFast `json:"eapFast,omitempty"`                   // The eapFast is required only if allowEapFast is true, otherwise it must be ignored. The object eapFast contains the settings for EAP FAST protocol
	EapTtls                   ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsEapTtls `json:"eapTtls,omitempty"`                   // The eapTtls is required only if allowEapTtls is true, otherwise it must be ignored. The object eapTtls contains the settings for EAP TTLS protocol
	Teap                      ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsTeap    `json:"teap,omitempty"`                      // The teap is required only if allowTeap is true, otherwise it must be ignored. The object teap contains the settings for TEAP protocol
	ProcessHostLookup         bool                                                                  `json:"processHostLookup,omitempty"`         //
	AllowPapAscii             bool                                                                  `json:"allowPapAscii,omitempty"`             //
	AllowChap                 bool                                                                  `json:"allowChap,omitempty"`                 //
	AllowMsChapV1             bool                                                                  `json:"allowMsChapV1,omitempty"`             //
	AllowMsChapV2             bool                                                                  `json:"allowMsChapV2,omitempty"`             //
	AllowEapMd5               bool                                                                  `json:"allowEapMd5,omitempty"`               //
	AllowLeap                 bool                                                                  `json:"allowLeap,omitempty"`                 //
	AllowEapTls               bool                                                                  `json:"allowEapTls,omitempty"`               //
	AllowEapTtls              bool                                                                  `json:"allowEapTtls,omitempty"`              //
	AllowEapFast              bool                                                                  `json:"allowEapFast,omitempty"`              //
	AllowPeap                 bool                                                                  `json:"allowPeap,omitempty"`                 //
	AllowTeap                 bool                                                                  `json:"allowTeap,omitempty"`                 //
	AllowPreferredEapProtocol bool                                                                  `json:"allowPreferredEapProtocol,omitempty"` //
	PreferredEapProtocol      string                                                                `json:"preferredEapProtocol,omitempty"`      // The preferredEapProtocol is required only if allowPreferredEapProtocol is true, otherwise it must be ignored. Allowed Values:  - EAP_FAST, - PEAP, - LEAP, - EAP_MD5, - EAP_TLS, - EAP_TTLS, - TEAP
	EapTlsLBit                bool                                                                  `json:"eapTlsLBit,omitempty"`                //
	AllowWeakCiphersForEap    bool                                                                  `json:"allowWeakCiphersForEap,omitempty"`    //
	RequireMessageAuth        bool                                                                  `json:"requireMessageAuth,omitempty"`        //
	Link                      ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsLink    `json:"link,omitempty"`                      //
}

type ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsEapTls struct {
	AllowEapTlsAuthOfExpiredCerts      bool   `json:"allowEapTlsAuthOfExpiredCerts,omitempty"`      //
	EapTlsEnableStatelessSessionResume bool   `json:"eapTlsEnableStatelessSessionResume,omitempty"` //
	EapTlsSessionTicketTtl             int    `json:"eapTlsSessionTicketTtl,omitempty"`             // Time to live. The eapTlsSessionTicketTtl is required only if eapTlsEnableStatelessSessionResume is true, otherwise it must be ignored
	EapTlsSessionTicketTtlUnits        string `json:"eapTlsSessionTicketTtlUnits,omitempty"`        // Time to live time units. The eapTlsSessionTicketTtlUnits is required only if eapTlsEnableStatelessSessionResume is true, otherwise it must be ignored. Allowed Values: - SECONDS, - MINUTES, - HOURS, - DAYS, - WEEKS
	EapTlsSessionTicketPrecentage      int    `json:"eapTlsSessionTicketPrecentage,omitempty"`      // The eapTlsSessionTicketPrecentage is required only if eapTlsEnableStatelessSessionResume is true, otherwise it must be ignored
}

type ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsPeap struct {
	AllowPeapEapMsChapV2                 bool `json:"allowPeapEapMsChapV2,omitempty"`                 //
	AllowPeapEapMsChapV2PwdChange        bool `json:"allowPeapEapMsChapV2PwdChange,omitempty"`        // The allowPeapEapMsChapV2PwdChange is required only if allowPeapEapMsChapV2 is true, otherwise it must be ignored
	AllowPeapEapMsChapV2PwdChangeRetries int  `json:"allowPeapEapMsChapV2PwdChangeRetries,omitempty"` // The allowPeapEapMsChapV2PwdChangeRetries is required only if allowPeapEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3
	AllowPeapEapGtc                      bool `json:"allowPeapEapGtc,omitempty"`                      //
	AllowPeapEapGtcPwdChange             bool `json:"allowPeapEapGtcPwdChange,omitempty"`             // The allowPeapEapGtcPwdChange is required only if allowPeapEapGtc is true, otherwise it must be ignored
	AllowPeapEapGtcPwdChangeRetries      int  `json:"allowPeapEapGtcPwdChangeRetries,omitempty"`      // The allowPeapEapGtcPwdChangeRetries is required only if allowPeapEapGtc is true, otherwise it must be ignored. Valid range is 0-3
	AllowPeapEapTls                      bool `json:"allowPeapEapTls,omitempty"`                      //
	AllowPeapEapTlsAuthOfExpiredCerts    bool `json:"allowPeapEapTlsAuthOfExpiredCerts,omitempty"`    // The allowPeapEapTlsAuthOfExpiredCerts is required only if allowPeapEapTls is true, otherwise it must be ignored
	RequireCryptobinding                 bool `json:"requireCryptobinding,omitempty"`                 //
	AllowPeapV0                          bool `json:"allowPeapV0,omitempty"`                          //
}

type ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsEapFast struct {
	AllowEapFastEapMsChapV2                                        bool   `json:"allowEapFastEapMsChapV2,omitempty"`                                        //
	AllowEapFastEapMsChapV2PwdChange                               bool   `json:"allowEapFastEapMsChapV2PwdChange,omitempty"`                               // The allowEapFastEapMsChapV2PwdChange is required only if allowEapFastEapMsChapV2 is true, otherwise it must be ignored
	AllowEapFastEapMsChapV2PwdChangeRetries                        int    `json:"allowEapFastEapMsChapV2PwdChangeRetries,omitempty"`                        // The allowEapFastEapMsChapV2PwdChangeRetries is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3
	AllowEapFastEapGtc                                             bool   `json:"allowEapFastEapGtc,omitempty"`                                             //
	AllowEapFastEapGtcPwdChange                                    bool   `json:"allowEapFastEapGtcPwdChange,omitempty"`                                    // The allowEapFastEapGtcPwdChange is required only if allowEapFastEapGtc is true, otherwise it must be ignored
	AllowEapFastEapGtcPwdChangeRetries                             int    `json:"allowEapFastEapGtcPwdChangeRetries,omitempty"`                             // The allowEapFastEapGtcPwdChangeRetries is required only if allowEapFastEapGtc is true, otherwise it must be ignored. Valid range is 0-3
	AllowEapFastEapTls                                             bool   `json:"allowEapFastEapTls,omitempty"`                                             //
	AllowEapFastEapTlsAuthOfExpiredCerts                           bool   `json:"allowEapFastEapTlsAuthOfExpiredCerts,omitempty"`                           // The allowEapFastEapTlsAuthOfExpiredCerts is required only if allowEapFastEapTls is true, otherwise it must be ignored
	EapFastUsePacs                                                 bool   `json:"eapFastUsePacs,omitempty"`                                                 //
	EapFastUsePacsTunnelPacTtl                                     int    `json:"eapFastUsePacsTunnelPacTtl,omitempty"`                                     // The eapFastUsePacsTunnelPacTtl is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsTunnelPacTtlUnits                                string `json:"eapFastUsePacsTunnelPacTtlUnits,omitempty"`                                // The eapFastUsePacsTunnelPacTtlUnits is required only if eapFastUsePacs is true, otherwise it must be ignored. Allowed Values: - SECONDS, - MINUTES, - HOURS, - DAYS, - WEEKS
	EapFastUsePacsUseProactivePacUpdatePrecentage                  int    `json:"eapFastUsePacsUseProactivePacUpdatePrecentage,omitempty"`                  // The eapFastUsePacsUseProactivePacUpdatePrecentage is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsAllowAnonymProvisioning                          bool   `json:"eapFastUsePacsAllowAnonymProvisioning,omitempty"`                          // The eapFastUsePacsAllowAnonymProvisioning is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsAllowAuthenProvisioning                          bool   `json:"eapFastUsePacsAllowAuthenProvisioning,omitempty"`                          // The eapFastUsePacsAllowAuthenProvisioning is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning bool   `json:"eapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning,omitempty"` // The eapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning is required only if eapFastUsePacsAllowAuthenProvisioning is true, otherwise it must be ignored
	EapFastUsePacsAcceptClientCert                                 bool   `json:"eapFastUsePacsAcceptClientCert,omitempty"`                                 // The eapFastUsePacsAcceptClientCert is required only if eapFastUsePacsAllowAuthenProvisioning is true, otherwise it must be ignored
	EapFastUsePacsMachinePacTtl                                    int    `json:"eapFastUsePacsMachinePacTtl,omitempty"`                                    // The eapFastUsePacsMachinePacTtl is required only if eapFastUsePacsAllowMachineAuthentication is true, otherwise it must be ignored
	EapFastUsePacsMachinePacTtlUnits                               string `json:"eapFastUsePacsMachinePacTtlUnits,omitempty"`                               // The eapFastUsePacsMachinePacTtlUnits is required only if eapFastUsePacsAllowMachineAuthentication is true, otherwise it must be ignored. Allowed Values: - SECONDS, - MINUTES, - HOURS, - DAYS, - WEEKS
	EapFastUsePacsAllowMachineAuthentication                       bool   `json:"eapFastUsePacsAllowMachineAuthentication,omitempty"`                       //
	EapFastUsePacsStatelessSessionResume                           bool   `json:"eapFastUsePacsStatelessSessionResume,omitempty"`                           // The eapFastUsePacsStatelessSessionResume is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsAuthorizationPacTtl                              int    `json:"eapFastUsePacsAuthorizationPacTtl,omitempty"`                              // The eapFastUsePacsAuthorizationPacTtl is required only if eapFastUsePacsStatelessSessionResume is true, otherwise it must be ignored
	EapFastUsePacsAuthorizationPacTtlUnits                         string `json:"eapFastUsePacsAuthorizationPacTtlUnits,omitempty"`                         // The eapFastUsePacsAuthorizationPacTtlUnits is required only if eapFastUsePacsStatelessSessionResume is true, otherwise it must be ignored. Allowed Values: - SECONDS, - MINUTES, - HOURS, - DAYS, - WEEKS
	EapFastDontUsePacsAcceptClientCert                             bool   `json:"eapFastDontUsePacsAcceptClientCert,omitempty"`                             // The eapFastDontUsePacsAcceptClientCert is required only if eapFastUsePacs is FALSE, otherwise it must be ignored
	EapFastDontUsePacsAllowMachineAuthentication                   bool   `json:"eapFastDontUsePacsAllowMachineAuthentication,omitempty"`                   // The eapFastDontUsePacsAllowMachineAuthentication is required only if eapFastUsePacs is FALSE, otherwise it must be ignored
	EapFastEnableEApChaining                                       bool   `json:"eapFastEnableEAPChaining,omitempty"`                                       //
}

type ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsEapTtls struct {
	EapTtlsPapAscii                    bool `json:"eapTtlsPapAscii,omitempty"`                    //
	EapTtlsChap                        bool `json:"eapTtlsChap,omitempty"`                        //
	EapTtlsMsChapV1                    bool `json:"eapTtlsMsChapV1,omitempty"`                    //
	EapTtlsMsChapV2                    bool `json:"eapTtlsMsChapV2,omitempty"`                    //
	EapTtlsEapMd5                      bool `json:"eapTtlsEapMd5,omitempty"`                      //
	EapTtlsEapMsChapV2                 bool `json:"eapTtlsEapMsChapV2,omitempty"`                 //
	EapTtlsEapMsChapV2PwdChange        bool `json:"eapTtlsEapMsChapV2PwdChange,omitempty"`        // The eapTtlsEapMsChapV2PwdChange is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored
	EapTtlsEapMsChapV2PwdChangeRetries int  `json:"eapTtlsEapMsChapV2PwdChangeRetries,omitempty"` // The eapTtlsEapMsChapV2PwdChangeRetries is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3
}

type ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsTeap struct {
	AllowTeapEapMsChapV2                 bool `json:"allowTeapEapMsChapV2,omitempty"`                 //
	AllowTeapEapMsChapV2PwdChange        bool `json:"allowTeapEapMsChapV2PwdChange,omitempty"`        // The allowTeapEapMsChapV2PwdChange is required only if allowTeapEapMsChapV2 is true, otherwise it must be ignored
	AllowTeapEapMsChapV2PwdChangeRetries int  `json:"allowTeapEapMsChapV2PwdChangeRetries,omitempty"` // The allowTeapEapMsChapV2PwdChangeRetries is required only if allowTeapEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3
	AllowTeapEapTls                      bool `json:"allowTeapEapTls,omitempty"`                      //
	AllowTeapEapTlsAuthOfExpiredCerts    bool `json:"allowTeapEapTlsAuthOfExpiredCerts,omitempty"`    // The allowTeapEapTlsAuthOfExpiredCerts is required only if allowTeapEapTls is true, otherwise it must be ignored
	AcceptClientCertDuringTunnelEst      bool `json:"acceptClientCertDuringTunnelEst,omitempty"`      //
	EnableEapChaining                    bool `json:"enableEapChaining,omitempty"`                    //
	AllowDowngradeMsk                    bool `json:"allowDowngradeMsk,omitempty"`                    //
}

type ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAllowedProtocolsUpdateAllowedProtocolByID struct {
	UpdatedFieldsList ResponseAllowedProtocolsUpdateAllowedProtocolByIDUpdatedFieldsList `json:"UpdatedFieldsList,omitempty"` //
}

type ResponseAllowedProtocolsUpdateAllowedProtocolByIDUpdatedFieldsList struct {
	UpdatedField ResponseAllowedProtocolsUpdateAllowedProtocolByIDUpdatedFieldsListUpdatedField `json:"updatedField,omitempty"` //
	Field        string                                                                         `json:"field,omitempty"`        //
	OldValue     string                                                                         `json:"oldValue,omitempty"`     //
	NewValue     string                                                                         `json:"newValue,omitempty"`     //
}

type ResponseAllowedProtocolsUpdateAllowedProtocolByIDUpdatedFieldsListUpdatedField struct {
	Field    string `json:"field,omitempty"`    //
	OldValue string `json:"oldValue,omitempty"` //
	NewValue string `json:"newValue,omitempty"` //
}

type ResponseAllowedProtocolsGetAllowedProtocols struct {
	SearchResult ResponseAllowedProtocolsGetAllowedProtocolsSearchResult `json:"SearchResult,omitempty"` //
}

type ResponseAllowedProtocolsGetAllowedProtocolsSearchResult struct {
	Total        int                                                                   `json:"total,omitempty"`        //
	Resources    []ResponseAllowedProtocolsGetAllowedProtocolsSearchResultResources    `json:"resources,omitempty"`    //
	NextPage     []ResponseAllowedProtocolsGetAllowedProtocolsSearchResultNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage []ResponseAllowedProtocolsGetAllowedProtocolsSearchResultPreviousPage `json:"previousPage,omitempty"` //
}

type ResponseAllowedProtocolsGetAllowedProtocolsSearchResultResources struct {
	ID          string                                                               `json:"id,omitempty"`          //
	Name        string                                                               `json:"name,omitempty"`        //
	Description string                                                               `json:"description,omitempty"` //
	Link        ResponseAllowedProtocolsGetAllowedProtocolsSearchResultResourcesLink `json:"link,omitempty"`        //
}

type ResponseAllowedProtocolsGetAllowedProtocolsSearchResultResourcesLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAllowedProtocolsGetAllowedProtocolsSearchResultNextPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAllowedProtocolsGetAllowedProtocolsSearchResultPreviousPage struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type ResponseAllowedProtocolsGetVersion struct {
	VersionInfo ResponseAllowedProtocolsGetVersionVersionInfo `json:"VersionInfo,omitempty"` //
}

type ResponseAllowedProtocolsGetVersionVersionInfo struct {
	CurrentServerVersion string                                            `json:"currentServerVersion,omitempty"` //
	SupportedVersions    string                                            `json:"supportedVersions,omitempty"`    //
	Link                 ResponseAllowedProtocolsGetVersionVersionInfoLink `json:"link,omitempty"`                 //
}

type ResponseAllowedProtocolsGetVersionVersionInfoLink struct {
	Rel  string `json:"rel,omitempty"`  //
	Href string `json:"href,omitempty"` //
	Type string `json:"type,omitempty"` //
}

type RequestAllowedProtocolsUpdateAllowedProtocolByID struct {
	AllowedProtocols *RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocols `json:"AllowedProtocols,omitempty"` //
}

type RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocols struct {
	ID                        string                                                                   `json:"id,omitempty"`                        // Resource UUID, Mandatory for update
	Name                      string                                                                   `json:"name,omitempty"`                      // Resource Name
	Description               string                                                                   `json:"description,omitempty"`               //
	EapTls                    *RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTls  `json:"eapTls,omitempty"`                    // The eapTls is required only if allowEapTls is true, otherwise it must be ignored. The object eapTls contains the settings for EAP TLS protocol
	Peap                      *RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsPeap    `json:"peap,omitempty"`                      //
	EapFast                   *RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapFast `json:"eapFast,omitempty"`                   // The eapFast is required only if allowEapFast is true, otherwise it must be ignored. The object eapFast contains the settings for EAP FAST protocol
	EapTtls                   *RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTtls `json:"eapTtls,omitempty"`                   // The eapTtls is required only if allowEapTtls is true, otherwise it must be ignored. The object eapTtls contains the settings for EAP TTLS protocol
	Teap                      *RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsTeap    `json:"teap,omitempty"`                      // The teap is required only if allowTeap is true, otherwise it must be ignored. The object teap contains the settings for TEAP protocol
	ProcessHostLookup         *bool                                                                    `json:"processHostLookup,omitempty"`         //
	AllowPapAscii             *bool                                                                    `json:"allowPapAscii,omitempty"`             //
	AllowChap                 *bool                                                                    `json:"allowChap,omitempty"`                 //
	AllowMsChapV1             *bool                                                                    `json:"allowMsChapV1,omitempty"`             //
	AllowMsChapV2             *bool                                                                    `json:"allowMsChapV2,omitempty"`             //
	AllowEapMd5               *bool                                                                    `json:"allowEapMd5,omitempty"`               //
	AllowLeap                 *bool                                                                    `json:"allowLeap,omitempty"`                 //
	AllowEapTls               *bool                                                                    `json:"allowEapTls,omitempty"`               //
	AllowEapTtls              *bool                                                                    `json:"allowEapTtls,omitempty"`              //
	AllowEapFast              *bool                                                                    `json:"allowEapFast,omitempty"`              //
	AllowPeap                 *bool                                                                    `json:"allowPeap,omitempty"`                 //
	AllowTeap                 *bool                                                                    `json:"allowTeap,omitempty"`                 //
	AllowPreferredEapProtocol *bool                                                                    `json:"allowPreferredEapProtocol,omitempty"` //
	PreferredEapProtocol      string                                                                   `json:"preferredEapProtocol,omitempty"`      // The preferredEapProtocol is required only if allowPreferredEapProtocol is true, otherwise it must be ignored. Allowed Values:  - EAP_FAST, - PEAP, - LEAP, - EAP_MD5, - EAP_TLS, - EAP_TTLS, - TEAP
	EapTlsLBit                *bool                                                                    `json:"eapTlsLBit,omitempty"`                //
	AllowWeakCiphersForEap    *bool                                                                    `json:"allowWeakCiphersForEap,omitempty"`    //
	RequireMessageAuth        *bool                                                                    `json:"requireMessageAuth,omitempty"`        //
}

type RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTls struct {
	AllowEapTlsAuthOfExpiredCerts      *bool  `json:"allowEapTlsAuthOfExpiredCerts,omitempty"`      //
	EapTlsEnableStatelessSessionResume *bool  `json:"eapTlsEnableStatelessSessionResume,omitempty"` //
	EapTlsSessionTicketTtl             *int   `json:"eapTlsSessionTicketTtl,omitempty"`             // Time to live. The eapTlsSessionTicketTtl is required only if eapTlsEnableStatelessSessionResume is true, otherwise it must be ignored
	EapTlsSessionTicketTtlUnits        string `json:"eapTlsSessionTicketTtlUnits,omitempty"`        // Time to live time units. The eapTlsSessionTicketTtlUnits is required only if eapTlsEnableStatelessSessionResume is true, otherwise it must be ignored. Allowed Values: - SECONDS, - MINUTES, - HOURS, - DAYS, - WEEKS
	EapTlsSessionTicketPrecentage      *int   `json:"eapTlsSessionTicketPrecentage,omitempty"`      // The eapTlsSessionTicketPrecentage is required only if eapTlsEnableStatelessSessionResume is true, otherwise it must be ignored
}

type RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsPeap struct {
	AllowPeapEapMsChapV2                 *bool `json:"allowPeapEapMsChapV2,omitempty"`                 //
	AllowPeapEapMsChapV2PwdChange        *bool `json:"allowPeapEapMsChapV2PwdChange,omitempty"`        // The allowPeapEapMsChapV2PwdChange is required only if allowPeapEapMsChapV2 is true, otherwise it must be ignored
	AllowPeapEapMsChapV2PwdChangeRetries *int  `json:"allowPeapEapMsChapV2PwdChangeRetries,omitempty"` // The allowPeapEapMsChapV2PwdChangeRetries is required only if allowPeapEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3
	AllowPeapEapGtc                      *bool `json:"allowPeapEapGtc,omitempty"`                      //
	AllowPeapEapGtcPwdChange             *bool `json:"allowPeapEapGtcPwdChange,omitempty"`             // The allowPeapEapGtcPwdChange is required only if allowPeapEapGtc is true, otherwise it must be ignored
	AllowPeapEapGtcPwdChangeRetries      *int  `json:"allowPeapEapGtcPwdChangeRetries,omitempty"`      // The allowPeapEapGtcPwdChangeRetries is required only if allowPeapEapGtc is true, otherwise it must be ignored. Valid range is 0-3
	AllowPeapEapTls                      *bool `json:"allowPeapEapTls,omitempty"`                      //
	AllowPeapEapTlsAuthOfExpiredCerts    *bool `json:"allowPeapEapTlsAuthOfExpiredCerts,omitempty"`    // The allowPeapEapTlsAuthOfExpiredCerts is required only if allowPeapEapTls is true, otherwise it must be ignored
	RequireCryptobinding                 *bool `json:"requireCryptobinding,omitempty"`                 //
	AllowPeapV0                          *bool `json:"allowPeapV0,omitempty"`                          //
}

type RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapFast struct {
	AllowEapFastEapMsChapV2                                        *bool  `json:"allowEapFastEapMsChapV2,omitempty"`                                        //
	AllowEapFastEapMsChapV2PwdChange                               *bool  `json:"allowEapFastEapMsChapV2PwdChange,omitempty"`                               // The allowEapFastEapMsChapV2PwdChange is required only if allowEapFastEapMsChapV2 is true, otherwise it must be ignored
	AllowEapFastEapMsChapV2PwdChangeRetries                        *int   `json:"allowEapFastEapMsChapV2PwdChangeRetries,omitempty"`                        // The allowEapFastEapMsChapV2PwdChangeRetries is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3
	AllowEapFastEapGtc                                             *bool  `json:"allowEapFastEapGtc,omitempty"`                                             //
	AllowEapFastEapGtcPwdChange                                    *bool  `json:"allowEapFastEapGtcPwdChange,omitempty"`                                    // The allowEapFastEapGtcPwdChange is required only if allowEapFastEapGtc is true, otherwise it must be ignored
	AllowEapFastEapGtcPwdChangeRetries                             *int   `json:"allowEapFastEapGtcPwdChangeRetries,omitempty"`                             // The allowEapFastEapGtcPwdChangeRetries is required only if allowEapFastEapGtc is true, otherwise it must be ignored. Valid range is 0-3
	AllowEapFastEapTls                                             *bool  `json:"allowEapFastEapTls,omitempty"`                                             //
	AllowEapFastEapTlsAuthOfExpiredCerts                           *bool  `json:"allowEapFastEapTlsAuthOfExpiredCerts,omitempty"`                           // The allowEapFastEapTlsAuthOfExpiredCerts is required only if allowEapFastEapTls is true, otherwise it must be ignored
	EapFastUsePacs                                                 *bool  `json:"eapFastUsePacs,omitempty"`                                                 //
	EapFastUsePacsTunnelPacTtl                                     *int   `json:"eapFastUsePacsTunnelPacTtl,omitempty"`                                     // The eapFastUsePacsTunnelPacTtl is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsTunnelPacTtlUnits                                string `json:"eapFastUsePacsTunnelPacTtlUnits,omitempty"`                                // The eapFastUsePacsTunnelPacTtlUnits is required only if eapFastUsePacs is true, otherwise it must be ignored. Allowed Values: - SECONDS, - MINUTES, - HOURS, - DAYS, - WEEKS
	EapFastUsePacsUseProactivePacUpdatePrecentage                  *int   `json:"eapFastUsePacsUseProactivePacUpdatePrecentage,omitempty"`                  // The eapFastUsePacsUseProactivePacUpdatePrecentage is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsAllowAnonymProvisioning                          *bool  `json:"eapFastUsePacsAllowAnonymProvisioning,omitempty"`                          // The eapFastUsePacsAllowAnonymProvisioning is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsAllowAuthenProvisioning                          *bool  `json:"eapFastUsePacsAllowAuthenProvisioning,omitempty"`                          // The eapFastUsePacsAllowAuthenProvisioning is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning *bool  `json:"eapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning,omitempty"` // The eapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning is required only if eapFastUsePacsAllowAuthenProvisioning is true, otherwise it must be ignored
	EapFastUsePacsAcceptClientCert                                 *bool  `json:"eapFastUsePacsAcceptClientCert,omitempty"`                                 // The eapFastUsePacsAcceptClientCert is required only if eapFastUsePacsAllowAuthenProvisioning is true, otherwise it must be ignored
	EapFastUsePacsMachinePacTtl                                    *int   `json:"eapFastUsePacsMachinePacTtl,omitempty"`                                    // The eapFastUsePacsMachinePacTtl is required only if eapFastUsePacsAllowMachineAuthentication is true, otherwise it must be ignored
	EapFastUsePacsMachinePacTtlUnits                               string `json:"eapFastUsePacsMachinePacTtlUnits,omitempty"`                               // The eapFastUsePacsMachinePacTtlUnits is required only if eapFastUsePacsAllowMachineAuthentication is true, otherwise it must be ignored. Allowed Values: - SECONDS, - MINUTES, - HOURS, - DAYS, - WEEKS
	EapFastUsePacsAllowMachineAuthentication                       *bool  `json:"eapFastUsePacsAllowMachineAuthentication,omitempty"`                       //
	EapFastUsePacsStatelessSessionResume                           *bool  `json:"eapFastUsePacsStatelessSessionResume,omitempty"`                           // The eapFastUsePacsStatelessSessionResume is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsAuthorizationPacTtl                              *int   `json:"eapFastUsePacsAuthorizationPacTtl,omitempty"`                              // The eapFastUsePacsAuthorizationPacTtl is required only if eapFastUsePacsStatelessSessionResume is true, otherwise it must be ignored
	EapFastUsePacsAuthorizationPacTtlUnits                         string `json:"eapFastUsePacsAuthorizationPacTtlUnits,omitempty"`                         // The eapFastUsePacsAuthorizationPacTtlUnits is required only if eapFastUsePacsStatelessSessionResume is true, otherwise it must be ignored. Allowed Values: - SECONDS, - MINUTES, - HOURS, - DAYS, - WEEKS
	EapFastDontUsePacsAcceptClientCert                             *bool  `json:"eapFastDontUsePacsAcceptClientCert,omitempty"`                             // The eapFastDontUsePacsAcceptClientCert is required only if eapFastUsePacs is FALSE, otherwise it must be ignored
	EapFastDontUsePacsAllowMachineAuthentication                   *bool  `json:"eapFastDontUsePacsAllowMachineAuthentication,omitempty"`                   // The eapFastDontUsePacsAllowMachineAuthentication is required only if eapFastUsePacs is FALSE, otherwise it must be ignored
	EapFastEnableEApChaining                                       *bool  `json:"eapFastEnableEAPChaining,omitempty"`                                       //
}

type RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTtls struct {
	EapTtlsPapAscii                    *bool `json:"eapTtlsPapAscii,omitempty"`                    //
	EapTtlsChap                        *bool `json:"eapTtlsChap,omitempty"`                        //
	EapTtlsMsChapV1                    *bool `json:"eapTtlsMsChapV1,omitempty"`                    //
	EapTtlsMsChapV2                    *bool `json:"eapTtlsMsChapV2,omitempty"`                    //
	EapTtlsEapMd5                      *bool `json:"eapTtlsEapMd5,omitempty"`                      //
	EapTtlsEapMsChapV2                 *bool `json:"eapTtlsEapMsChapV2,omitempty"`                 //
	EapTtlsEapMsChapV2PwdChange        *bool `json:"eapTtlsEapMsChapV2PwdChange,omitempty"`        // The eapTtlsEapMsChapV2PwdChange is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored
	EapTtlsEapMsChapV2PwdChangeRetries *int  `json:"eapTtlsEapMsChapV2PwdChangeRetries,omitempty"` // The eapTtlsEapMsChapV2PwdChangeRetries is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3
}

type RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsTeap struct {
	AllowTeapEapMsChapV2                 *bool `json:"allowTeapEapMsChapV2,omitempty"`                 //
	AllowTeapEapMsChapV2PwdChange        *bool `json:"allowTeapEapMsChapV2PwdChange,omitempty"`        // The allowTeapEapMsChapV2PwdChange is required only if allowTeapEapMsChapV2 is true, otherwise it must be ignored
	AllowTeapEapMsChapV2PwdChangeRetries *int  `json:"allowTeapEapMsChapV2PwdChangeRetries,omitempty"` // The allowTeapEapMsChapV2PwdChangeRetries is required only if allowTeapEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3
	AllowTeapEapTls                      *bool `json:"allowTeapEapTls,omitempty"`                      //
	AllowTeapEapTlsAuthOfExpiredCerts    *bool `json:"allowTeapEapTlsAuthOfExpiredCerts,omitempty"`    // The allowTeapEapTlsAuthOfExpiredCerts is required only if allowTeapEapTls is true, otherwise it must be ignored
	AcceptClientCertDuringTunnelEst      *bool `json:"acceptClientCertDuringTunnelEst,omitempty"`      //
	EnableEapChaining                    *bool `json:"enableEapChaining,omitempty"`                    //
	AllowDowngradeMsk                    *bool `json:"allowDowngradeMsk,omitempty"`                    //
}

type RequestAllowedProtocolsCreateAllowedProtocol struct {
	AllowedProtocols *RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocols `json:"AllowedProtocols,omitempty"` //
}

type RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocols struct {
	Name                      string                                                               `json:"name,omitempty"`                      // Resource Name
	Description               string                                                               `json:"description,omitempty"`               //
	EapTls                    *RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTls  `json:"eapTls,omitempty"`                    // The eapTls is required only if allowEapTls is true, otherwise it must be ignored. The object eapTls contains the settings for EAP TLS protocol
	Peap                      *RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsPeap    `json:"peap,omitempty"`                      //
	EapFast                   *RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapFast `json:"eapFast,omitempty"`                   // The eapFast is required only if allowEapFast is true, otherwise it must be ignored. The object eapFast contains the settings for EAP FAST protocol
	EapTtls                   *RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTtls `json:"eapTtls,omitempty"`                   // The eapTtls is required only if allowEapTtls is true, otherwise it must be ignored. The object eapTtls contains the settings for EAP TTLS protocol
	Teap                      *RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsTeap    `json:"teap,omitempty"`                      // The teap is required only if allowTeap is true, otherwise it must be ignored. The object teap contains the settings for TEAP protocol
	ProcessHostLookup         *bool                                                                `json:"processHostLookup,omitempty"`         //
	AllowPapAscii             *bool                                                                `json:"allowPapAscii,omitempty"`             //
	AllowChap                 *bool                                                                `json:"allowChap,omitempty"`                 //
	AllowMsChapV1             *bool                                                                `json:"allowMsChapV1,omitempty"`             //
	AllowMsChapV2             *bool                                                                `json:"allowMsChapV2,omitempty"`             //
	AllowEapMd5               *bool                                                                `json:"allowEapMd5,omitempty"`               //
	AllowLeap                 *bool                                                                `json:"allowLeap,omitempty"`                 //
	AllowEapTls               *bool                                                                `json:"allowEapTls,omitempty"`               //
	AllowEapTtls              *bool                                                                `json:"allowEapTtls,omitempty"`              //
	AllowEapFast              *bool                                                                `json:"allowEapFast,omitempty"`              //
	AllowPeap                 *bool                                                                `json:"allowPeap,omitempty"`                 //
	AllowTeap                 *bool                                                                `json:"allowTeap,omitempty"`                 //
	AllowPreferredEapProtocol *bool                                                                `json:"allowPreferredEapProtocol,omitempty"` //
	PreferredEapProtocol      string                                                               `json:"preferredEapProtocol,omitempty"`      // The preferredEapProtocol is required only if allowPreferredEapProtocol is true, otherwise it must be ignored. Allowed Values:  - EAP_FAST, - PEAP, - LEAP, - EAP_MD5, - EAP_TLS, - EAP_TTLS, - TEAP
	EapTlsLBit                *bool                                                                `json:"eapTlsLBit,omitempty"`                //
	AllowWeakCiphersForEap    *bool                                                                `json:"allowWeakCiphersForEap,omitempty"`    //
	RequireMessageAuth        *bool                                                                `json:"requireMessageAuth,omitempty"`        //
}

type RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTls struct {
	AllowEapTlsAuthOfExpiredCerts      *bool  `json:"allowEapTlsAuthOfExpiredCerts,omitempty"`      //
	EapTlsEnableStatelessSessionResume *bool  `json:"eapTlsEnableStatelessSessionResume,omitempty"` //
	EapTlsSessionTicketTtl             *int   `json:"eapTlsSessionTicketTtl,omitempty"`             // Time to live. The eapTlsSessionTicketTtl is required only if eapTlsEnableStatelessSessionResume is true, otherwise it must be ignored
	EapTlsSessionTicketTtlUnits        string `json:"eapTlsSessionTicketTtlUnits,omitempty"`        // Time to live time units. The eapTlsSessionTicketTtlUnits is required only if eapTlsEnableStatelessSessionResume is true, otherwise it must be ignored. Allowed Values: - SECONDS, - MINUTES, - HOURS, - DAYS, - WEEKS
	EapTlsSessionTicketPrecentage      *int   `json:"eapTlsSessionTicketPrecentage,omitempty"`      // The eapTlsSessionTicketPrecentage is required only if eapTlsEnableStatelessSessionResume is true, otherwise it must be ignored
}

type RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsPeap struct {
	AllowPeapEapMsChapV2                 *bool `json:"allowPeapEapMsChapV2,omitempty"`                 //
	AllowPeapEapMsChapV2PwdChange        *bool `json:"allowPeapEapMsChapV2PwdChange,omitempty"`        // The allowPeapEapMsChapV2PwdChange is required only if allowPeapEapMsChapV2 is true, otherwise it must be ignored
	AllowPeapEapMsChapV2PwdChangeRetries *int  `json:"allowPeapEapMsChapV2PwdChangeRetries,omitempty"` // The allowPeapEapMsChapV2PwdChangeRetries is required only if allowPeapEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3
	AllowPeapEapGtc                      *bool `json:"allowPeapEapGtc,omitempty"`                      //
	AllowPeapEapGtcPwdChange             *bool `json:"allowPeapEapGtcPwdChange,omitempty"`             // The allowPeapEapGtcPwdChange is required only if allowPeapEapGtc is true, otherwise it must be ignored
	AllowPeapEapGtcPwdChangeRetries      *int  `json:"allowPeapEapGtcPwdChangeRetries,omitempty"`      // The allowPeapEapGtcPwdChangeRetries is required only if allowPeapEapGtc is true, otherwise it must be ignored. Valid range is 0-3
	AllowPeapEapTls                      *bool `json:"allowPeapEapTls,omitempty"`                      //
	AllowPeapEapTlsAuthOfExpiredCerts    *bool `json:"allowPeapEapTlsAuthOfExpiredCerts,omitempty"`    // The allowPeapEapTlsAuthOfExpiredCerts is required only if allowPeapEapTls is true, otherwise it must be ignored
	RequireCryptobinding                 *bool `json:"requireCryptobinding,omitempty"`                 //
	AllowPeapV0                          *bool `json:"allowPeapV0,omitempty"`                          //
}

type RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapFast struct {
	AllowEapFastEapMsChapV2                                        *bool  `json:"allowEapFastEapMsChapV2,omitempty"`                                        //
	AllowEapFastEapMsChapV2PwdChange                               *bool  `json:"allowEapFastEapMsChapV2PwdChange,omitempty"`                               // The allowEapFastEapMsChapV2PwdChange is required only if allowEapFastEapMsChapV2 is true, otherwise it must be ignored
	AllowEapFastEapMsChapV2PwdChangeRetries                        *int   `json:"allowEapFastEapMsChapV2PwdChangeRetries,omitempty"`                        // The allowEapFastEapMsChapV2PwdChangeRetries is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3
	AllowEapFastEapGtc                                             *bool  `json:"allowEapFastEapGtc,omitempty"`                                             //
	AllowEapFastEapGtcPwdChange                                    *bool  `json:"allowEapFastEapGtcPwdChange,omitempty"`                                    // The allowEapFastEapGtcPwdChange is required only if allowEapFastEapGtc is true, otherwise it must be ignored
	AllowEapFastEapGtcPwdChangeRetries                             *int   `json:"allowEapFastEapGtcPwdChangeRetries,omitempty"`                             // The allowEapFastEapGtcPwdChangeRetries is required only if allowEapFastEapGtc is true, otherwise it must be ignored. Valid range is 0-3
	AllowEapFastEapTls                                             *bool  `json:"allowEapFastEapTls,omitempty"`                                             //
	AllowEapFastEapTlsAuthOfExpiredCerts                           *bool  `json:"allowEapFastEapTlsAuthOfExpiredCerts,omitempty"`                           // The allowEapFastEapTlsAuthOfExpiredCerts is required only if allowEapFastEapTls is true, otherwise it must be ignored
	EapFastUsePacs                                                 *bool  `json:"eapFastUsePacs,omitempty"`                                                 //
	EapFastUsePacsTunnelPacTtl                                     *int   `json:"eapFastUsePacsTunnelPacTtl,omitempty"`                                     // The eapFastUsePacsTunnelPacTtl is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsTunnelPacTtlUnits                                string `json:"eapFastUsePacsTunnelPacTtlUnits,omitempty"`                                // The eapFastUsePacsTunnelPacTtlUnits is required only if eapFastUsePacs is true, otherwise it must be ignored. Allowed Values: - SECONDS, - MINUTES, - HOURS, - DAYS, - WEEKS
	EapFastUsePacsUseProactivePacUpdatePrecentage                  *int   `json:"eapFastUsePacsUseProactivePacUpdatePrecentage,omitempty"`                  // The eapFastUsePacsUseProactivePacUpdatePrecentage is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsAllowAnonymProvisioning                          *bool  `json:"eapFastUsePacsAllowAnonymProvisioning,omitempty"`                          // The eapFastUsePacsAllowAnonymProvisioning is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsAllowAuthenProvisioning                          *bool  `json:"eapFastUsePacsAllowAuthenProvisioning,omitempty"`                          // The eapFastUsePacsAllowAuthenProvisioning is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning *bool  `json:"eapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning,omitempty"` // The eapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning is required only if eapFastUsePacsAllowAuthenProvisioning is true, otherwise it must be ignored
	EapFastUsePacsAcceptClientCert                                 *bool  `json:"eapFastUsePacsAcceptClientCert,omitempty"`                                 // The eapFastUsePacsAcceptClientCert is required only if eapFastUsePacsAllowAuthenProvisioning is true, otherwise it must be ignored
	EapFastUsePacsMachinePacTtl                                    *int   `json:"eapFastUsePacsMachinePacTtl,omitempty"`                                    // The eapFastUsePacsMachinePacTtl is required only if eapFastUsePacsAllowMachineAuthentication is true, otherwise it must be ignored
	EapFastUsePacsMachinePacTtlUnits                               string `json:"eapFastUsePacsMachinePacTtlUnits,omitempty"`                               // The eapFastUsePacsMachinePacTtlUnits is required only if eapFastUsePacsAllowMachineAuthentication is true, otherwise it must be ignored. Allowed Values: - SECONDS, - MINUTES, - HOURS, - DAYS, - WEEKS
	EapFastUsePacsAllowMachineAuthentication                       *bool  `json:"eapFastUsePacsAllowMachineAuthentication,omitempty"`                       //
	EapFastUsePacsStatelessSessionResume                           *bool  `json:"eapFastUsePacsStatelessSessionResume,omitempty"`                           // The eapFastUsePacsStatelessSessionResume is required only if eapFastUsePacs is true, otherwise it must be ignored
	EapFastUsePacsAuthorizationPacTtl                              *int   `json:"eapFastUsePacsAuthorizationPacTtl,omitempty"`                              // The eapFastUsePacsAuthorizationPacTtl is required only if eapFastUsePacsStatelessSessionResume is true, otherwise it must be ignored
	EapFastUsePacsAuthorizationPacTtlUnits                         string `json:"eapFastUsePacsAuthorizationPacTtlUnits,omitempty"`                         // The eapFastUsePacsAuthorizationPacTtlUnits is required only if eapFastUsePacsStatelessSessionResume is true, otherwise it must be ignored. Allowed Values: - SECONDS, - MINUTES, - HOURS, - DAYS, - WEEKS
	EapFastDontUsePacsAcceptClientCert                             *bool  `json:"eapFastDontUsePacsAcceptClientCert,omitempty"`                             // The eapFastDontUsePacsAcceptClientCert is required only if eapFastUsePacs is FALSE, otherwise it must be ignored
	EapFastDontUsePacsAllowMachineAuthentication                   *bool  `json:"eapFastDontUsePacsAllowMachineAuthentication,omitempty"`                   // The eapFastDontUsePacsAllowMachineAuthentication is required only if eapFastUsePacs is FALSE, otherwise it must be ignored
	EapFastEnableEApChaining                                       *bool  `json:"eapFastEnableEAPChaining,omitempty"`                                       //
}

type RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTtls struct {
	EapTtlsPapAscii                    *bool `json:"eapTtlsPapAscii,omitempty"`                    //
	EapTtlsChap                        *bool `json:"eapTtlsChap,omitempty"`                        //
	EapTtlsMsChapV1                    *bool `json:"eapTtlsMsChapV1,omitempty"`                    //
	EapTtlsMsChapV2                    *bool `json:"eapTtlsMsChapV2,omitempty"`                    //
	EapTtlsEapMd5                      *bool `json:"eapTtlsEapMd5,omitempty"`                      //
	EapTtlsEapMsChapV2                 *bool `json:"eapTtlsEapMsChapV2,omitempty"`                 //
	EapTtlsEapMsChapV2PwdChange        *bool `json:"eapTtlsEapMsChapV2PwdChange,omitempty"`        // The eapTtlsEapMsChapV2PwdChange is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored
	EapTtlsEapMsChapV2PwdChangeRetries *int  `json:"eapTtlsEapMsChapV2PwdChangeRetries,omitempty"` // The eapTtlsEapMsChapV2PwdChangeRetries is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3
}

type RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsTeap struct {
	AllowTeapEapMsChapV2                 *bool `json:"allowTeapEapMsChapV2,omitempty"`                 //
	AllowTeapEapMsChapV2PwdChange        *bool `json:"allowTeapEapMsChapV2PwdChange,omitempty"`        // The allowTeapEapMsChapV2PwdChange is required only if allowTeapEapMsChapV2 is true, otherwise it must be ignored
	AllowTeapEapMsChapV2PwdChangeRetries *int  `json:"allowTeapEapMsChapV2PwdChangeRetries,omitempty"` // The allowTeapEapMsChapV2PwdChangeRetries is required only if allowTeapEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3
	AllowTeapEapTls                      *bool `json:"allowTeapEapTls,omitempty"`                      //
	AllowTeapEapTlsAuthOfExpiredCerts    *bool `json:"allowTeapEapTlsAuthOfExpiredCerts,omitempty"`    // The allowTeapEapTlsAuthOfExpiredCerts is required only if allowTeapEapTls is true, otherwise it must be ignored
	AcceptClientCertDuringTunnelEst      *bool `json:"acceptClientCertDuringTunnelEst,omitempty"`      //
	EnableEapChaining                    *bool `json:"enableEapChaining,omitempty"`                    //
	AllowDowngradeMsk                    *bool `json:"allowDowngradeMsk,omitempty"`                    //
}

//GetAllowedProtocolByName Get allowed protocol by name
/* This API allows the client to get an allowed protocol by name.

@param name name path parameter.
*/
func (s *AllowedProtocolsService) GetAllowedProtocolByName(name string) (*ResponseAllowedProtocolsGetAllowedProtocolByName, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/allowedprotocols/name/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAllowedProtocolsGetAllowedProtocolByName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAllowedProtocolByName")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAllowedProtocolsGetAllowedProtocolByName)
	return result, response, err

}

//GetAllowedProtocolByID Get allowed protocol by ID
/* This API allows the client to get an allowed protocol by ID.

@param id id path parameter.
*/
func (s *AllowedProtocolsService) GetAllowedProtocolByID(id string) (*ResponseAllowedProtocolsGetAllowedProtocolByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/allowedprotocols/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAllowedProtocolsGetAllowedProtocolByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAllowedProtocolById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAllowedProtocolsGetAllowedProtocolByID)
	return result, response, err

}

//GetAllowedProtocols Get all allowed protocols
/* This API allows the client to get all the allowed protocols.

@param getAllowedProtocolsQueryParams Filtering parameter
*/
func (s *AllowedProtocolsService) GetAllowedProtocols(getAllowedProtocolsQueryParams *GetAllowedProtocolsQueryParams) (*ResponseAllowedProtocolsGetAllowedProtocols, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/allowedprotocols"

	queryString, _ := query.Values(getAllowedProtocolsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAllowedProtocolsGetAllowedProtocols{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAllowedProtocols")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAllowedProtocolsGetAllowedProtocols)
	return result, response, err

}

//GetVersion Get allowed protocols version information
/* This API helps to retrieve the version information related to the allowed protocols.

 */
func (s *AllowedProtocolsService) GetVersion() (*ResponseAllowedProtocolsGetVersion, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/allowedprotocols/versioninfo"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAllowedProtocolsGetVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVersion")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAllowedProtocolsGetVersion)
	return result, response, err

}

//CreateAllowedProtocol Create allowed protocol
/* This API creates an allowed protocol.

 */
func (s *AllowedProtocolsService) CreateAllowedProtocol(requestAllowedProtocolsCreateAllowedProtocol *RequestAllowedProtocolsCreateAllowedProtocol) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/allowedprotocols"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAllowedProtocolsCreateAllowedProtocol).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateAllowedProtocol")
	}

	return response, err

}

//UpdateAllowedProtocolByID Update allowed protocol
/* This API allows the client to update an allowed protocol.

@param id id path parameter.
*/
func (s *AllowedProtocolsService) UpdateAllowedProtocolByID(id string, requestAllowedProtocolsUpdateAllowedProtocolById *RequestAllowedProtocolsUpdateAllowedProtocolByID) (*ResponseAllowedProtocolsUpdateAllowedProtocolByID, *resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/allowedprotocols/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAllowedProtocolsUpdateAllowedProtocolById).
		SetResult(&ResponseAllowedProtocolsUpdateAllowedProtocolByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateAllowedProtocolById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseAllowedProtocolsUpdateAllowedProtocolByID)
	return result, response, err

}

//DeleteAllowedProtocolByID Delete allowed protocol
/* This API deletes an allowed protocol.

@param id id path parameter.
*/
func (s *AllowedProtocolsService) DeleteAllowedProtocolByID(id string) (*resty.Response, error) {
	setHost(s.client, "_ers")
	path := "/ers/config/allowedprotocols/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

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
		return response, fmt.Errorf("error with operation DeleteAllowedProtocolById")
	}

	getCSFRToken(response.Header())

	return response, err

}
