package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type CertificatesService service

type GetCsrsQueryParams struct {
	Page       int    `url:"page,omitempty"`       //Page number
	Size       int    `url:"size,omitempty"`       //Number of objects returned per page
	Sort       string `url:"sort,omitempty"`       //sort type - asc or desc
	SortBy     string `url:"sortBy,omitempty"`     //sort column by which objects needs to be sorted
	Filter     string `url:"filter,omitempty"`     //<div> <style type="text/css" scoped> .apiServiceTable td, .apiServiceTable th { padding: 5px 10px !important; text-align: left; } </style> <span> <b>Simple filtering</b> should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the <i>"filterType=or"</i> query string parameter. Each resource Data model description should specify if an attribute is a filtered field. </span> <br /> <table class="apiServiceTable"> <thead> <tr> <th>OPERATOR</th> <th>DESCRIPTION</th> </tr> </thead> <tbody> <tr> <td>EQ</td> <td>Equals</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> </tr> <tr> <td>GT</td> <td>Greater Than</td> </tr> <tr> <td>LT</td> <td>Less Then</td> </tr> <tr> <td>STARTSW</td> <td>Starts With</td> </tr> <tr> <td>NSTARTSW</td> <td>Not Starts With</td> </tr> <tr> <td>ENDSW</td> <td>Ends With</td> </tr> <tr> <td>NENDSW</td> <td>Not Ends With</td> </tr> <tr> <td>CONTAINS</td> <td>Contains</td> </tr> <tr> <td>NCONTAINS</td> <td>Not Contains</td> </tr> </tbody> </table> </div>
	FilterType string `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type GetSystemCertificatesQueryParams struct {
	Page       int    `url:"page,omitempty"`       //Page number
	Size       int    `url:"size,omitempty"`       //Number of objects returned per page
	Sort       string `url:"sort,omitempty"`       //sort type - asc or desc
	SortBy     string `url:"sortBy,omitempty"`     //sort column by which objects needs to be sorted
	Filter     string `url:"filter,omitempty"`     //<div> <style type="text/css" scoped> .apiServiceTable td, .apiServiceTable th { padding: 5px 10px !important; text-align: left; } </style> <span> <b>Simple filtering</b> should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the <i>"filterType=or"</i> query string parameter. Each resource Data model description should specify if an attribute is a filtered field. </span> <br /> <table class="apiServiceTable"> <thead> <tr> <th>OPERATOR</th> <th>DESCRIPTION</th> </tr> </thead> <tbody> <tr> <td>EQ</td> <td>Equals</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> </tr> <tr> <td>GT</td> <td>Greater Than</td> </tr> <tr> <td>LT</td> <td>Less Then</td> </tr> <tr> <td>STARTSW</td> <td>Starts With</td> </tr> <tr> <td>NSTARTSW</td> <td>Not Starts With</td> </tr> <tr> <td>ENDSW</td> <td>Ends With</td> </tr> <tr> <td>NENDSW</td> <td>Not Ends With</td> </tr> <tr> <td>CONTAINS</td> <td>Contains</td> </tr> <tr> <td>NCONTAINS</td> <td>Not Contains</td> </tr> </tbody> </table> </div>
	FilterType string `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type GetTrustedCertificatesQueryParams struct {
	Page       int    `url:"page,omitempty"`       //Page number
	Size       int    `url:"size,omitempty"`       //Number of objects returned per page
	Sort       string `url:"sort,omitempty"`       //sort type - asc or desc
	SortBy     string `url:"sortBy,omitempty"`     //sort column by which objects needs to be sorted
	Filter     string `url:"filter,omitempty"`     //<div> <style type="text/css" scoped> .apiServiceTable td, .apiServiceTable th { padding: 5px 10px !important; text-align: left; } </style> <span> <b>Simple filtering</b> should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the <i>"filterType=or"</i> query string parameter. Each resource Data model description should specify if an attribute is a filtered field. </span> <br /> <table class="apiServiceTable"> <thead> <tr> <th>OPERATOR</th> <th>DESCRIPTION</th> </tr> </thead> <tbody> <tr> <td>EQ</td> <td>Equals</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> </tr> <tr> <td>GT</td> <td>Greater Than</td> </tr> <tr> <td>LT</td> <td>Less Then</td> </tr> <tr> <td>STARTSW</td> <td>Starts With</td> </tr> <tr> <td>NSTARTSW</td> <td>Not Starts With</td> </tr> <tr> <td>ENDSW</td> <td>Ends With</td> </tr> <tr> <td>NENDSW</td> <td>Not Ends With</td> </tr> <tr> <td>CONTAINS</td> <td>Contains</td> </tr> <tr> <td>NCONTAINS</td> <td>Not Contains</td> </tr> </tbody> </table> </div>
	FilterType string `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type ResponseCertificatesGetCsrs struct {
	NextPage     *ResponseCertificatesGetCsrsNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseCertificatesGetCsrsPreviousPage `json:"previousPage,omitempty"` //
	Response     *[]ResponseCertificatesGetCsrsResponse   `json:"response,omitempty"`     //
	Version      string                                   `json:"version,omitempty"`      //
}

type ResponseCertificatesGetCsrsNextPage struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesGetCsrsPreviousPage struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesGetCsrsResponse struct {
	FriendlyName       string                                   `json:"friendlyName,omitempty"`       // Friendly name of the certificate.
	GroupTag           string                                   `json:"groupTag,omitempty"`           // GroupTag of the certificate.
	HostName           string                                   `json:"hostName,omitempty"`           // Hostname or IP address of the Cisco ISE node.
	ID                 string                                   `json:"id,omitempty"`                 // ID of the certificate.
	KeySize            string                                   `json:"keySize,omitempty"`            // Size of the cryptographic key used.
	Link               *ResponseCertificatesGetCsrsResponseLink `json:"link,omitempty"`               //
	SanNames           string                                   `json:"sanNames,omitempty"`           // String representation of subject alternative names.
	SignatureAlgorithm string                                   `json:"signatureAlgorithm,omitempty"` // Algorithm used for encrypting CSR
	Subject            string                                   `json:"subject,omitempty"`            // Subject of the certificate. Includes Common Name (CN), Organizational Unit (OU), etc.
	TimeStamp          string                                   `json:"timeStamp,omitempty"`          // Timestamp of the certificate generation.
	UsedFor            string                                   `json:"usedFor,omitempty"`            // Services for which the certificate is used for (for eg- MGMT, GENERIC).
}

type ResponseCertificatesGetCsrsResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesGenerateCsr struct {
	Response *[]ResponseCertificatesGenerateCsrResponse `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  //
}

type ResponseCertificatesGenerateCsrResponse struct {
	ID      string                                       `json:"id,omitempty"`      // ID of the generated CSR
	Link    *ResponseCertificatesGenerateCsrResponseLink `json:"link,omitempty"`    //
	Message string                                       `json:"message,omitempty"` // Response message on generation of CSR
}

type ResponseCertificatesGenerateCsrResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesGenerateIntermediateCaCsr struct {
	Response *ResponseCertificatesGenerateIntermediateCaCsrResponse `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  //
}

type ResponseCertificatesGenerateIntermediateCaCsrResponse struct {
	ID      string                                                     `json:"id,omitempty"`      // ID of the generated CSR
	Link    *ResponseCertificatesGenerateIntermediateCaCsrResponseLink `json:"link,omitempty"`    //
	Message string                                                     `json:"message,omitempty"` // Response message on generation of CSR
}

type ResponseCertificatesGenerateIntermediateCaCsrResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesGetCsrByID struct {
	Response *ResponseCertificatesGetCsrByIDResponse `json:"response,omitempty"` //
	Version  string                                  `json:"version,omitempty"`  //
}

type ResponseCertificatesGetCsrByIDResponse struct {
	CsrContents        string                                      `json:"csrContents,omitempty"`        // Contents of the certificate file.
	FriendlyName       string                                      `json:"friendlyName,omitempty"`       // Friendly name of the certificate.
	GroupTag           string                                      `json:"groupTag,omitempty"`           // GroupTag of the certificate.
	HostName           string                                      `json:"hostName,omitempty"`           // Hostname or IP address of the Cisco ISE node.
	ID                 string                                      `json:"id,omitempty"`                 // ID of the certificate.
	KeySize            string                                      `json:"keySize,omitempty"`            // Size of the cryptographic key used.
	Link               *ResponseCertificatesGetCsrByIDResponseLink `json:"link,omitempty"`               //
	SanNames           string                                      `json:"sanNames,omitempty"`           // String representation of subject alternative names.
	SignatureAlgorithm string                                      `json:"signatureAlgorithm,omitempty"` // Algorithm used for encrypting CSR
	Subject            string                                      `json:"subject,omitempty"`            // Subject of the certificate. Includes Common Name (CN), Organizational Unit (OU), etc.
	TimeStamp          string                                      `json:"timeStamp,omitempty"`          // Timestamp of the certificate generation.
	UsedFor            string                                      `json:"usedFor,omitempty"`            // Services for which the certificate is used for (for eg- MGMT, GENERIC).
}

type ResponseCertificatesGetCsrByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesDeleteCsr struct {
	Response *ResponseCertificatesDeleteCsrResponse `json:"response,omitempty"` //
	Version  string                                 `json:"version,omitempty"`  //
}

// ResponseCertificatesDeleteCsrByID is deprecated, please use ResponseCertificatesDeleteCsr
type ResponseCertificatesDeleteCsrByID = ResponseCertificatesDeleteCsr

type ResponseCertificatesDeleteCsrResponse struct {
	Message string `json:"message,omitempty"` //
}

// ResponseCertificatesDeleteCsrByIDResponse is deprecated, please use ResponseCertificatesDeleteCsrResponse
type ResponseCertificatesDeleteCsrByIDResponse = ResponseCertificatesDeleteCsrResponse

type ResponseCertificatesRegenerateIseRootCa struct {
	Response *ResponseCertificatesRegenerateIseRootCaResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  //
}

type ResponseCertificatesRegenerateIseRootCaResponse struct {
	ID      string                                               `json:"id,omitempty"`      // ID which can be used to track the status of Cisco ISE root CA chain regeneration
	Link    *ResponseCertificatesRegenerateIseRootCaResponseLink `json:"link,omitempty"`    //
	Message string                                               `json:"message,omitempty"` //
}

type ResponseCertificatesRegenerateIseRootCaResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesRenewCerts struct {
	Response *ResponseCertificatesRenewCertsResponse `json:"response,omitempty"` //
	Version  string                                  `json:"version,omitempty"`  //
}

// ResponseCertificatesRenewCertificates is deprecated, please use ResponseCertificatesRenewCerts
type ResponseCertificatesRenewCertificates = ResponseCertificatesRenewCerts

type ResponseCertificatesRenewCertsResponse struct {
	ID      string                                      `json:"id,omitempty"`      // ID which can be used to track the status of certificate regeneration
	Link    *ResponseCertificatesRenewCertsResponseLink `json:"link,omitempty"`    //
	Message string                                      `json:"message,omitempty"` //
}

// ResponseCertificatesRenewCertificatesResponse is deprecated, please use ResponseCertificatesRenewCertsResponse
type ResponseCertificatesRenewCertificatesResponse = ResponseCertificatesRenewCertsResponse

type ResponseCertificatesRenewCertsResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

// ResponseCertificatesRenewCertificatesResponseLink is deprecated, please use ResponseCertificatesRenewCertsResponseLink
type ResponseCertificatesRenewCertificatesResponseLink = ResponseCertificatesRenewCertsResponseLink

type ResponseCertificatesBindCsr struct {
	Response *ResponseCertificatesBindCsrResponse `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  //
}

type ResponseCertificatesBindCsrResponse struct {
	Message string `json:"message,omitempty"` //
	Status  string `json:"status,omitempty"`  // Response status after import
}

type ResponseCertificatesGenerateSelfSignedCertificate struct {
	Response *ResponseCertificatesGenerateSelfSignedCertificateResponse `json:"response,omitempty"` //
	Version  string                                                     `json:"version,omitempty"`  //
}

type ResponseCertificatesGenerateSelfSignedCertificateResponse struct {
	ID      string `json:"id,omitempty"`      // ID of the generated self-signed system certificate
	Message string `json:"message,omitempty"` // Response message on generation of self-signed system certificate
	Status  string `json:"status,omitempty"`  // HTTP response status after import
}

type ResponseCertificatesImportSystemCert struct {
	Response *ResponseCertificatesImportSystemCertResponse `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  //
}

// ResponseCertificatesImportSystemCertificate is deprecated, please use ResponseCertificatesImportSystemCert
type ResponseCertificatesImportSystemCertificate = ResponseCertificatesImportSystemCert

type ResponseCertificatesImportSystemCertResponse struct {
	ID      string `json:"id,omitempty"`      // ID of the imported trust certificate
	Message string `json:"message,omitempty"` // Response message on import of system or trust certificate
	Status  string `json:"status,omitempty"`  // HTTP response status after import
}

// ResponseCertificatesImportSystemCertificateResponse is deprecated, please use ResponseCertificatesImportSystemCertResponse
type ResponseCertificatesImportSystemCertificateResponse = ResponseCertificatesImportSystemCertResponse

type ResponseCertificatesGetSystemCertificates struct {
	NextPage     *ResponseCertificatesGetSystemCertificatesNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseCertificatesGetSystemCertificatesPreviousPage `json:"previousPage,omitempty"` //
	Response     *[]ResponseCertificatesGetSystemCertificatesResponse   `json:"response,omitempty"`     //
	Version      string                                                 `json:"version,omitempty"`      //
}

type ResponseCertificatesGetSystemCertificatesNextPage struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesGetSystemCertificatesPreviousPage struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesGetSystemCertificatesResponse struct {
	ExpirationDate            string                                                 `json:"expirationDate,omitempty"`            // Time and date past which the certificate is no longer valid
	FriendlyName              string                                                 `json:"friendlyName,omitempty"`              // Friendly name of system certificate
	GroupTag                  string                                                 `json:"groupTag,omitempty"`                  //
	ID                        string                                                 `json:"id,omitempty"`                        // ID of system certificate
	IssuedBy                  string                                                 `json:"issuedBy,omitempty"`                  // Common Name of the certificate issuer
	IssuedTo                  string                                                 `json:"issuedTo,omitempty"`                  // Common Name of the certificate subject
	KeySize                   *int                                                   `json:"keySize,omitempty"`                   // Length of the key used for encrypting system certificate
	Link                      *ResponseCertificatesGetSystemCertificatesResponseLink `json:"link,omitempty"`                      //
	PortalsUsingTheTag        string                                                 `json:"portalsUsingTheTag,omitempty"`        //
	SelfSigned                *bool                                                  `json:"selfSigned,omitempty"`                //
	SerialNumberDecimalFormat string                                                 `json:"serialNumberDecimalFormat,omitempty"` // Used to uniquely identify the certificate within a CA's systems
	Sha256Fingerprint         string                                                 `json:"sha256Fingerprint,omitempty"`         //
	SignatureAlgorithm        string                                                 `json:"signatureAlgorithm,omitempty"`        //
	UsedBy                    string                                                 `json:"usedBy,omitempty"`                    //
	ValidFrom                 string                                                 `json:"validFrom,omitempty"`                 // Time and date on which the certificate was created, also known as the Not Before certificate attribute
}

type ResponseCertificatesGetSystemCertificatesResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesGetSystemCertificateByID struct {
	Response *ResponseCertificatesGetSystemCertificateByIDResponse `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  //
}

type ResponseCertificatesGetSystemCertificateByIDResponse struct {
	ExpirationDate            string                                                    `json:"expirationDate,omitempty"`            // Time and date past which the certificate is no longer valid
	FriendlyName              string                                                    `json:"friendlyName,omitempty"`              // Friendly name of system certificate
	GroupTag                  string                                                    `json:"groupTag,omitempty"`                  //
	ID                        string                                                    `json:"id,omitempty"`                        // ID of system certificate
	IssuedBy                  string                                                    `json:"issuedBy,omitempty"`                  // Common Name of the certificate issuer
	IssuedTo                  string                                                    `json:"issuedTo,omitempty"`                  // Common Name of the certificate subject
	KeySize                   *int                                                      `json:"keySize,omitempty"`                   // Length of the key used for encrypting system certificate
	Link                      *ResponseCertificatesGetSystemCertificateByIDResponseLink `json:"link,omitempty"`                      //
	PortalsUsingTheTag        string                                                    `json:"portalsUsingTheTag,omitempty"`        //
	SelfSigned                *bool                                                     `json:"selfSigned,omitempty"`                //
	SerialNumberDecimalFormat string                                                    `json:"serialNumberDecimalFormat,omitempty"` // Used to uniquely identify the certificate within a CA's systems
	Sha256Fingerprint         string                                                    `json:"sha256Fingerprint,omitempty"`         //
	SignatureAlgorithm        string                                                    `json:"signatureAlgorithm,omitempty"`        //
	UsedBy                    string                                                    `json:"usedBy,omitempty"`                    //
	ValidFrom                 string                                                    `json:"validFrom,omitempty"`                 // Time and date on which the certificate was created, also known as the Not Before certificate attribute
}

type ResponseCertificatesGetSystemCertificateByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesUpdateSystemCert struct {
	Response *ResponseCertificatesUpdateSystemCertResponse `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  //
}

// ResponseCertificatesUpdateSystemCertificate is deprecated, please use ResponseCertificatesUpdateSystemCert
type ResponseCertificatesUpdateSystemCertificate = ResponseCertificatesUpdateSystemCert

type ResponseCertificatesUpdateSystemCertResponse struct {
	ID      string                                            `json:"id,omitempty"`      // ID of the imported trust certificate
	Link    *ResponseCertificatesUpdateSystemCertResponseLink `json:"link,omitempty"`    //
	Message string                                            `json:"message,omitempty"` // Response message on import of system or trust certificate
	Status  string                                            `json:"status,omitempty"`  // HTTP response status after import
}

// ResponseCertificatesUpdateSystemCertificateResponse is deprecated, please use ResponseCertificatesUpdateSystemCertResponse
type ResponseCertificatesUpdateSystemCertificateResponse = ResponseCertificatesUpdateSystemCertResponse

type ResponseCertificatesUpdateSystemCertResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

// ResponseCertificatesUpdateSystemCertificateResponseLink is deprecated, please use ResponseCertificatesUpdateSystemCertResponseLink
type ResponseCertificatesUpdateSystemCertificateResponseLink = ResponseCertificatesUpdateSystemCertResponseLink

type ResponseCertificatesDeleteSystemCertificateByID struct {
	Response *ResponseCertificatesDeleteSystemCertificateByIDResponse `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  //
}

type ResponseCertificatesDeleteSystemCertificateByIDResponse struct {
	Message string `json:"message,omitempty"` //
}

type ResponseCertificatesGetTrustedCertificates struct {
	NextPage     *ResponseCertificatesGetTrustedCertificatesNextPage     `json:"nextPage,omitempty"`     //
	PreviousPage *ResponseCertificatesGetTrustedCertificatesPreviousPage `json:"previousPage,omitempty"` //
	Response     *[]ResponseCertificatesGetTrustedCertificatesResponse   `json:"response,omitempty"`     //
	Version      string                                                  `json:"version,omitempty"`      //
}

type ResponseCertificatesGetTrustedCertificatesNextPage struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesGetTrustedCertificatesPreviousPage struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesGetTrustedCertificatesResponse struct {
	AuthenticateBeforeCRLReceived  string                                                  `json:"authenticateBeforeCRLReceived,omitempty"`  // Switch to enable or disable authentication before receiving CRL
	AutomaticCRLUpdate             string                                                  `json:"automaticCRLUpdate,omitempty"`             // Switch to enable or disable automatic CRL update
	AutomaticCRLUpdatePeriod       string                                                  `json:"automaticCRLUpdatePeriod,omitempty"`       // Automatic CRL update period
	AutomaticCRLUpdateUnits        string                                                  `json:"automaticCRLUpdateUnits,omitempty"`        // Unit of time of automatic CRL update
	CrlDistributionURL             string                                                  `json:"crlDistributionUrl,omitempty"`             // CRL Distribution URL
	CrlDownloadFailureRetries      string                                                  `json:"crlDownloadFailureRetries,omitempty"`      // If CRL download fails, wait time before retry
	CrlDownloadFailureRetriesUnits string                                                  `json:"crlDownloadFailureRetriesUnits,omitempty"` // Unit of time before retry if CRL download fails
	Description                    string                                                  `json:"description,omitempty"`                    // Description of trust certificate
	DownloadCRL                    string                                                  `json:"downloadCRL,omitempty"`                    // Switch to enable or disable download of CRL
	EnableOCSpValidation           string                                                  `json:"enableOCSPValidation,omitempty"`           // Switch to enable or disable OCSP Validation
	EnableServerIDentityCheck      string                                                  `json:"enableServerIdentityCheck,omitempty"`      // Switch to enable or disable Server Identity Check
	ExpirationDate                 string                                                  `json:"expirationDate,omitempty"`                 // The time and date past which the certificate is no longer valid
	FriendlyName                   string                                                  `json:"friendlyName,omitempty"`                   // Friendly name of trust certificate
	ID                             string                                                  `json:"id,omitempty"`                             // ID of trust certificate
	IgnoreCRLExpiration            string                                                  `json:"ignoreCRLExpiration,omitempty"`            // Switch to enable or disable ignore CRL Expiration
	InternalCa                     *bool                                                   `json:"internalCA,omitempty"`                     //
	IssuedBy                       string                                                  `json:"issuedBy,omitempty"`                       // The entity that verified the information and signed the certificate
	IssuedTo                       string                                                  `json:"issuedTo,omitempty"`                       // Entity to which trust certificate is issued
	KeySize                        string                                                  `json:"keySize,omitempty"`                        // Length of the key used for encrypting trust certificate
	Link                           *ResponseCertificatesGetTrustedCertificatesResponseLink `json:"link,omitempty"`                           //
	NonAutomaticCRLUpdatePeriod    string                                                  `json:"nonAutomaticCRLUpdatePeriod,omitempty"`    // Non automatic CRL update period
	NonAutomaticCRLUpdateUnits     string                                                  `json:"nonAutomaticCRLUpdateUnits,omitempty"`     // Unit of time of non automatic CRL update
	RejectIfNoStatusFromOCSP       string                                                  `json:"rejectIfNoStatusFromOCSP,omitempty"`       // Switch to reject certificate if there is no status from OCSP
	RejectIfUnreachableFromOCSP    string                                                  `json:"rejectIfUnreachableFromOCSP,omitempty"`    // Switch to reject certificate if unreachable from OCSP
	SelectedOCSpService            string                                                  `json:"selectedOCSPService,omitempty"`            // Name of selected OCSP Service
	SerialNumberDecimalFormat      string                                                  `json:"serialNumberDecimalFormat,omitempty"`      // Used to uniquely identify the certificate within a CA's systems
	Sha256Fingerprint              string                                                  `json:"sha256Fingerprint,omitempty"`              //
	SignatureAlgorithm             string                                                  `json:"signatureAlgorithm,omitempty"`             // Algorithm used for encrypting trust certificate
	Status                         string                                                  `json:"status,omitempty"`                         //
	Subject                        string                                                  `json:"subject,omitempty"`                        // The Subject or entity with which public key of trust certificate is associated
	TrustedFor                     string                                                  `json:"trustedFor,omitempty"`                     // Different services for which the certificated is trusted
	ValidFrom                      string                                                  `json:"validFrom,omitempty"`                      // The earliest time and date on which the certificate is valid
}

type ResponseCertificatesGetTrustedCertificatesResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesImportTrustCert struct {
	Response *ResponseCertificatesImportTrustCertResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  //
}

// ResponseCertificatesImportTrustCertificate is deprecated, please use ResponseCertificatesImportTrustCert
type ResponseCertificatesImportTrustCertificate = ResponseCertificatesImportTrustCert

type ResponseCertificatesImportTrustCertResponse struct {
	ID      string `json:"id,omitempty"`      // ID of the imported trust certificate
	Message string `json:"message,omitempty"` // Response message on import of system or trust certificate
	Status  string `json:"status,omitempty"`  // HTTP response status after import
}

// ResponseCertificatesImportTrustCertificateResponse is deprecated, please use ResponseCertificatesImportTrustCertResponse
type ResponseCertificatesImportTrustCertificateResponse = ResponseCertificatesImportTrustCertResponse

type ResponseCertificatesGetTrustedCertificateByID struct {
	Response *ResponseCertificatesGetTrustedCertificateByIDResponse `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  //
}

type ResponseCertificatesGetTrustedCertificateByIDResponse struct {
	AuthenticateBeforeCRLReceived  string                                                     `json:"authenticateBeforeCRLReceived,omitempty"`  // Switch to enable or disable authentication before receiving CRL
	AutomaticCRLUpdate             string                                                     `json:"automaticCRLUpdate,omitempty"`             // Switch to enable or disable automatic CRL update
	AutomaticCRLUpdatePeriod       string                                                     `json:"automaticCRLUpdatePeriod,omitempty"`       // Automatic CRL update period
	AutomaticCRLUpdateUnits        string                                                     `json:"automaticCRLUpdateUnits,omitempty"`        // Unit of time of automatic CRL update
	CrlDistributionURL             string                                                     `json:"crlDistributionUrl,omitempty"`             // CRL Distribution URL
	CrlDownloadFailureRetries      string                                                     `json:"crlDownloadFailureRetries,omitempty"`      // If CRL download fails, wait time before retry
	CrlDownloadFailureRetriesUnits string                                                     `json:"crlDownloadFailureRetriesUnits,omitempty"` // Unit of time before retry if CRL download fails
	Description                    string                                                     `json:"description,omitempty"`                    // Description of trust certificate
	DownloadCRL                    string                                                     `json:"downloadCRL,omitempty"`                    // Switch to enable or disable download of CRL
	EnableOCSpValidation           string                                                     `json:"enableOCSPValidation,omitempty"`           // Switch to enable or disable OCSP Validation
	EnableServerIDentityCheck      string                                                     `json:"enableServerIdentityCheck,omitempty"`      // Switch to enable or disable Server Identity Check
	ExpirationDate                 string                                                     `json:"expirationDate,omitempty"`                 // The time and date past which the certificate is no longer valid
	FriendlyName                   string                                                     `json:"friendlyName,omitempty"`                   // Friendly name of trust certificate
	ID                             string                                                     `json:"id,omitempty"`                             // ID of trust certificate
	IgnoreCRLExpiration            string                                                     `json:"ignoreCRLExpiration,omitempty"`            // Switch to enable or disable ignore CRL Expiration
	InternalCa                     *bool                                                      `json:"internalCA,omitempty"`                     //
	IsReferredInPolicy             *bool                                                      `json:"isReferredInPolicy,omitempty"`             //
	IssuedBy                       string                                                     `json:"issuedBy,omitempty"`                       // The entity that verified the information and signed the certificate
	IssuedTo                       string                                                     `json:"issuedTo,omitempty"`                       // Entity to which trust certificate is issued
	KeySize                        string                                                     `json:"keySize,omitempty"`                        // Length of the key used for encrypting trust certificate
	Link                           *ResponseCertificatesGetTrustedCertificateByIDResponseLink `json:"link,omitempty"`                           //
	NonAutomaticCRLUpdatePeriod    string                                                     `json:"nonAutomaticCRLUpdatePeriod,omitempty"`    // Non automatic CRL update period
	NonAutomaticCRLUpdateUnits     string                                                     `json:"nonAutomaticCRLUpdateUnits,omitempty"`     // Unit of time of non automatic CRL update
	RejectIfNoStatusFromOCSP       string                                                     `json:"rejectIfNoStatusFromOCSP,omitempty"`       // Switch to reject certificate if there is no status from OCSP
	RejectIfUnreachableFromOCSP    string                                                     `json:"rejectIfUnreachableFromOCSP,omitempty"`    // Switch to reject certificate if unreachable from OCSP
	SelectedOCSpService            string                                                     `json:"selectedOCSPService,omitempty"`            // Name of selected OCSP Service
	SerialNumberDecimalFormat      string                                                     `json:"serialNumberDecimalFormat,omitempty"`      // Used to uniquely identify the certificate within a CA's systems
	Sha256Fingerprint              string                                                     `json:"sha256Fingerprint,omitempty"`              //
	SignatureAlgorithm             string                                                     `json:"signatureAlgorithm,omitempty"`             // Algorithm used for encrypting trust certificate
	Status                         string                                                     `json:"status,omitempty"`                         //
	Subject                        string                                                     `json:"subject,omitempty"`                        // The Subject or entity with which public key of trust certificate is associated
	TrustedFor                     string                                                     `json:"trustedFor,omitempty"`                     // Different services for which the certificated is trusted
	ValidFrom                      string                                                     `json:"validFrom,omitempty"`                      // The earliest time and date on which the certificate is valid
}

type ResponseCertificatesGetTrustedCertificateByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesUpdateTrustedCertificate struct {
	Response *ResponseCertificatesUpdateTrustedCertificateResponse `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  //
}

type ResponseCertificatesUpdateTrustedCertificateResponse struct {
	ID      string                                                    `json:"id,omitempty"`      // ID of the trust certificate
	Link    *ResponseCertificatesUpdateTrustedCertificateResponseLink `json:"link,omitempty"`    //
	Message string                                                    `json:"message,omitempty"` // Response message on successful update of trust certificate
}

type ResponseCertificatesUpdateTrustedCertificateResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesDeleteTrustedCertificateByID struct {
	Response *ResponseCertificatesDeleteTrustedCertificateByIDResponse `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  //
}

type ResponseCertificatesDeleteTrustedCertificateByIDResponse struct {
	Message string `json:"message,omitempty"` //
}

type RequestCertificatesGenerateCsr struct {
	AllowWildCardCert   *bool    `json:"allowWildCardCert,omitempty"`   //
	CertificatePolicies string   `json:"certificatePolicies,omitempty"` //
	DigestType          string   `json:"digestType,omitempty"`          //
	Hostnames           []string `json:"hostnames,omitempty"`           //
	KeyLength           string   `json:"keyLength,omitempty"`           //
	KeyType             string   `json:"keyType,omitempty"`             //
	PortalGroupTag      string   `json:"portalGroupTag,omitempty"`      //
	SanDNS              []string `json:"sanDNS,omitempty"`              //
	SanDir              []string `json:"sanDir,omitempty"`              //
	SanIP               []string `json:"sanIP,omitempty"`               //
	SanURI              []string `json:"sanURI,omitempty"`              //
	SubjectCity         string   `json:"subjectCity,omitempty"`         //
	SubjectCommonName   string   `json:"subjectCommonName,omitempty"`   //
	SubjectCountry      string   `json:"subjectCountry,omitempty"`      //
	SubjectOrg          string   `json:"subjectOrg,omitempty"`          //
	SubjectOrgUnit      string   `json:"subjectOrgUnit,omitempty"`      //
	SubjectState        string   `json:"subjectState,omitempty"`        //
	UsedFor             string   `json:"usedFor,omitempty"`             //
}

type RequestCertificatesRegenerateIseRootCa struct {
	RemoveExistingIseIntermediateCsr *bool `json:"removeExistingISEIntermediateCSR,omitempty"` // Setting this attribute to true removes existing Cisco ISE Intermediate CSR
}

type RequestCertificatesRenewCerts struct {
	CertType string `json:"certType,omitempty"` //
}

// RequestCertificatesRenewCertificates is deprecated, please use RequestCertificatesRenewCerts
type RequestCertificatesRenewCertificates = RequestCertificatesRenewCerts

type RequestCertificatesBindCsr struct {
	Admin                            *bool  `json:"admin,omitempty"`                            //  Use certificate to authenticate the Cisco ISE Admin Portal
	AllowExtendedValidity            *bool  `json:"allowExtendedValidity,omitempty"`            // Allow import of certificates with validity greater than 398 days
	AllowOutOfDateCert               *bool  `json:"allowOutOfDateCert,omitempty"`               // Allow out of date certificates
	AllowReplacementOfCertificates   *bool  `json:"allowReplacementOfCertificates,omitempty"`   // Allow Replacement of certificates
	AllowReplacementOfPortalGroupTag *bool  `json:"allowReplacementOfPortalGroupTag,omitempty"` // Allow Replacement of Portal Group Tag
	Data                             string `json:"data,omitempty"`                             // Signed certificate data
	Eap                              *bool  `json:"eap,omitempty"`                              // Use certificate for EAP protocols that use SSL/TLS tunneling
	HostName                         string `json:"hostName,omitempty"`                         // Name of Host whose CSR ID has been provided
	ID                               string `json:"id,omitempty"`                               // ID of the generated CSR
	Ims                              *bool  `json:"ims,omitempty"`                              // Use certificate for the Cisco ISE Messaging Service
	Name                             string `json:"name,omitempty"`                             // Friendly Name of the certificate
	Portal                           *bool  `json:"portal,omitempty"`                           // Use for portal
	PortalGroupTag                   string `json:"portalGroupTag,omitempty"`                   // Set Group tag
	Pxgrid                           *bool  `json:"pxgrid,omitempty"`                           // Use certificate for the pxGrid Controller
	Radius                           *bool  `json:"radius,omitempty"`                           // Use certificate for the RADSec server
	Saml                             *bool  `json:"saml,omitempty"`                             // Use certificate for SAML Signing
	ValidateCertificateExtensions    *bool  `json:"validateCertificateExtensions,omitempty"`    // Validate Certificate Extensions
}

type RequestCertificatesExportSystemCert struct {
	Export   string `json:"export,omitempty"`   //
	HostName string `json:"hostName,omitempty"` // Hostname of the Cisco ISE node in which self-signed certificate should be generated.
	ID       string `json:"id,omitempty"`       //
	Password string `json:"password,omitempty"` //
}

// RequestCertificatesExportSystemCertificate is deprecated, please use RequestCertificatesExportSystemCert
type RequestCertificatesExportSystemCertificate = RequestCertificatesExportSystemCert

type RequestCertificatesGenerateSelfSignedCertificate struct {
	Admin                                *bool    `json:"admin,omitempty"`                                // Use certificate to authenticate the Cisco ISE Admin Portal
	AllowExtendedValidity                *bool    `json:"allowExtendedValidity,omitempty"`                // Allow generation of self-signed certificate with validity greater than 398 days
	AllowPortalTagTransferForSameSubject *bool    `json:"allowPortalTagTransferForSameSubject,omitempty"` // Allow overwriting the portal tag from matching certificate of same subject
	AllowReplacementOfCertificates       *bool    `json:"allowReplacementOfCertificates,omitempty"`       // Allow Replacement of certificates
	AllowReplacementOfPortalGroupTag     *bool    `json:"allowReplacementOfPortalGroupTag,omitempty"`     // Allow Replacement of Portal Group Tag
	AllowRoleTransferForSameSubject      *bool    `json:"allowRoleTransferForSameSubject,omitempty"`      // Allow transfer of roles for certificate with matching subject
	AllowSanDNSBadName                   *bool    `json:"allowSanDnsBadName,omitempty"`                   // Allow usage of SAN DNS Bad name
	AllowSanDNSNonResolvable             *bool    `json:"allowSanDnsNonResolvable,omitempty"`             // Allow use of non resolvable Common Name or SAN Values
	AllowWildCardCertificates            *bool    `json:"allowWildCardCertificates,omitempty"`            // Allow Wildcard Certificates
	CertificatePolicies                  string   `json:"certificatePolicies,omitempty"`                  // Certificate Policies
	DigestType                           string   `json:"digestType,omitempty"`                           // Digest to sign with
	Eap                                  *bool    `json:"eap,omitempty"`                                  // Use certificate for EAP protocols that use SSL/TLS tunneling
	ExpirationTTL                        *int     `json:"expirationTTL,omitempty"`                        // Certificate expiration value
	ExpirationTTLUnit                    string   `json:"expirationTTLUnit,omitempty"`                    // Certificate expiration unit
	HostName                             string   `json:"hostName,omitempty"`                             // Hostname of the Cisco ISE node in which self-signed certificate should be generated.
	KeyLength                            string   `json:"keyLength,omitempty"`                            // Bit size of public key
	KeyType                              string   `json:"keyType,omitempty"`                              // Algorithm to use for certificate public key creation
	Name                                 string   `json:"name,omitempty"`                                 // Friendly name of the certificate.
	Portal                               *bool    `json:"portal,omitempty"`                               // Use for portal
	PortalGroupTag                       string   `json:"portalGroupTag,omitempty"`                       // Set Group tag
	Pxgrid                               *bool    `json:"pxgrid,omitempty"`                               // Use certificate for the pxGrid Controller
	Radius                               *bool    `json:"radius,omitempty"`                               // Use certificate for the RADSec server
	Saml                                 *bool    `json:"saml,omitempty"`                                 // Use certificate for SAML Signing
	SanDNS                               []string `json:"sanDNS,omitempty"`                               // Array of SAN (Subject Alternative Name) DNS entries
	SanIP                                []string `json:"sanIP,omitempty"`                                // Array of SAN IP entries
	SanURI                               []string `json:"sanURI,omitempty"`                               // Array of SAN URI entries
	SubjectCity                          string   `json:"subjectCity,omitempty"`                          // Certificate city or locality (L)
	SubjectCommonName                    string   `json:"subjectCommonName,omitempty"`                    // Certificate common name (CN)
	SubjectCountry                       string   `json:"subjectCountry,omitempty"`                       // Certificate country (C)
	SubjectOrg                           string   `json:"subjectOrg,omitempty"`                           // Certificate organization (O)
	SubjectOrgUnit                       string   `json:"subjectOrgUnit,omitempty"`                       // Certificate organizational unit (OU)
	SubjectState                         string   `json:"subjectState,omitempty"`                         // Certificate state (ST)
}

type RequestCertificatesImportSystemCert struct {
	Admin                                *bool  `json:"admin,omitempty"`                                // Use certificate to authenticate the Cisco ISE Admin Portal
	AllowExtendedValidity                *bool  `json:"allowExtendedValidity,omitempty"`                // Allow import of certificates with validity greater than 398 days
	AllowOutOfDateCert                   *bool  `json:"allowOutOfDateCert,omitempty"`                   // Allow out of date certificates
	AllowPortalTagTransferForSameSubject *bool  `json:"allowPortalTagTransferForSameSubject,omitempty"` // Allow overwriting the portal tag from matching certificate of same subject
	AllowReplacementOfCertificates       *bool  `json:"allowReplacementOfCertificates,omitempty"`       // Allow Replacement of certificates
	AllowReplacementOfPortalGroupTag     *bool  `json:"allowReplacementOfPortalGroupTag,omitempty"`     // Allow Replacement of Portal Group Tag
	AllowRoleTransferForSameSubject      *bool  `json:"allowRoleTransferForSameSubject,omitempty"`      // Allow transfer of roles for certificate with matching subject
	AllowSHA1Certificates                *bool  `json:"allowSHA1Certificates,omitempty"`                // Allow SHA1 based certificates
	AllowWildCardCertificates            *bool  `json:"allowWildCardCertificates,omitempty"`            // Allow Wildcard certificates
	Data                                 string `json:"data,omitempty"`                                 // Certificate Content
	Eap                                  *bool  `json:"eap,omitempty"`                                  // Use certificate for EAP protocols that use SSL/TLS tunneling
	Ims                                  *bool  `json:"ims,omitempty"`                                  // Use certificate for the Cisco ISE Messaging Service
	Name                                 string `json:"name,omitempty"`                                 // Name of the certificate
	Password                             string `json:"password,omitempty"`                             // Certificate Password .
	Portal                               *bool  `json:"portal,omitempty"`                               // Use for portal
	PortalGroupTag                       string `json:"portalGroupTag,omitempty"`                       // Set Group tag
	PrivateKeyData                       string `json:"privateKeyData,omitempty"`                       // Private Key data
	Pxgrid                               *bool  `json:"pxgrid,omitempty"`                               // Use certificate for the pxGrid Controller
	Radius                               *bool  `json:"radius,omitempty"`                               // Use certificate for the RADSec server
	Saml                                 *bool  `json:"saml,omitempty"`                                 // Use certificate for SAML Signing
	ValidateCertificateExtensions        *bool  `json:"validateCertificateExtensions,omitempty"`        // Validate certificate extensions
}

// RequestCertificatesImportSystemCertificate is deprecated, please use RequestCertificatesImportSystemCert
type RequestCertificatesImportSystemCertificate = RequestCertificatesImportSystemCert

type RequestCertificatesUpdateSystemCert struct {
	Admin                                *bool  `json:"admin,omitempty"`                                // Use certificate to authenticate the Cisco ISE Admin Portal
	AllowPortalTagTransferForSameSubject *bool  `json:"allowPortalTagTransferForSameSubject,omitempty"` // Allow overwriting the portal tag from matching certificate of same subject
	AllowReplacementOfPortalGroupTag     *bool  `json:"allowReplacementOfPortalGroupTag,omitempty"`     // Allow Replacement of Portal Group Tag
	AllowRoleTransferForSameSubject      *bool  `json:"allowRoleTransferForSameSubject,omitempty"`      // Allow transfer of roles for certificate with matching subject
	Description                          string `json:"description,omitempty"`                          // Description of System Certificate
	Eap                                  *bool  `json:"eap,omitempty"`                                  // Use certificate for EAP protocols that use SSL/TLS tunneling
	ExpirationTTLPeriod                  *int   `json:"expirationTTLPeriod,omitempty"`                  //
	ExpirationTTLUnits                   string `json:"expirationTTLUnits,omitempty"`                   //
	Ims                                  *bool  `json:"ims,omitempty"`                                  // Use certificate for the Cisco ISE Messaging Service
	Name                                 string `json:"name,omitempty"`                                 // Name of the certificate
	Portal                               *bool  `json:"portal,omitempty"`                               // Use for portal
	PortalGroupTag                       string `json:"portalGroupTag,omitempty"`                       // Set Group tag
	Pxgrid                               *bool  `json:"pxgrid,omitempty"`                               // Use certificate for the pxGrid Controller
	Radius                               *bool  `json:"radius,omitempty"`                               // Use certificate for the RADSec server
	RenewSelfSignedCertificate           *bool  `json:"renewSelfSignedCertificate,omitempty"`           // Renew Self-signed Certificate
	Saml                                 *bool  `json:"saml,omitempty"`                                 // Use certificate for SAML Signing
}

// RequestCertificatesUpdateSystemCertificate is deprecated, please use RequestCertificatesUpdateSystemCert
type RequestCertificatesUpdateSystemCertificate = RequestCertificatesUpdateSystemCert

type RequestCertificatesImportTrustCert struct {
	AllowBasicConstraintCaFalse       *bool  `json:"allowBasicConstraintCAFalse,omitempty"`       // Allow certificates with Basic Constraints CA Field as False
	AllowOutOfDateCert                *bool  `json:"allowOutOfDateCert,omitempty"`                // Allow out of date certificates
	AllowSHA1Certificates             *bool  `json:"allowSHA1Certificates,omitempty"`             // Allow SHA1 based certificates
	Data                              string `json:"data,omitempty"`                              // Certificate content
	Description                       string `json:"description,omitempty"`                       // Description of the certificate
	Name                              string `json:"name,omitempty"`                              // Name of the certificate
	TrustForCertificateBasedAdminAuth *bool  `json:"trustForCertificateBasedAdminAuth,omitempty"` // Trust for Certificate based Admin authentication
	TrustForCiscoServicesAuth         *bool  `json:"trustForCiscoServicesAuth,omitempty"`         // Trust for authentication of Cisco Services
	TrustForClientAuth                *bool  `json:"trustForClientAuth,omitempty"`                // Trust for client authentication and Syslog
	TrustForIseAuth                   *bool  `json:"trustForIseAuth,omitempty"`                   // Trust for authentication within Cisco ISE
	ValidateCertificateExtensions     *bool  `json:"validateCertificateExtensions,omitempty"`     // Validate trust certificate extension
}

// RequestCertificatesImportTrustCertificate is deprecated, please use RequestCertificatesImportTrustCert
type RequestCertificatesImportTrustCertificate = RequestCertificatesImportTrustCert

type RequestCertificatesUpdateTrustedCertificate struct {
	AuthenticateBeforeCRLReceived     *bool  `json:"authenticateBeforeCRLReceived,omitempty"`     // Switch to enable or disable CRL verification if CRL is not received
	AutomaticCRLUpdate                *bool  `json:"automaticCRLUpdate,omitempty"`                // Switch to enable or disable automatic CRL update
	AutomaticCRLUpdatePeriod          *int   `json:"automaticCRLUpdatePeriod,omitempty"`          // Automatic CRL update period
	AutomaticCRLUpdateUnits           string `json:"automaticCRLUpdateUnits,omitempty"`           // Unit of time for automatic CRL update
	CrlDistributionURL                string `json:"crlDistributionUrl,omitempty"`                // CRL Distribution URL
	CrlDownloadFailureRetries         *int   `json:"crlDownloadFailureRetries,omitempty"`         // If CRL download fails, wait time before retry
	CrlDownloadFailureRetriesUnits    string `json:"crlDownloadFailureRetriesUnits,omitempty"`    // Unit of time before retry if CRL download fails
	Description                       string `json:"description,omitempty"`                       // Description for trust certificate
	DownloadCRL                       *bool  `json:"downloadCRL,omitempty"`                       // Switch to enable or disable download of CRL
	EnableOCSpValidation              *bool  `json:"enableOCSPValidation,omitempty"`              // Switch to enable or disable OCSP Validation
	EnableServerIDentityCheck         *bool  `json:"enableServerIdentityCheck,omitempty"`         // Switch to enable or disable verification if HTTPS or LDAP server certificate name fits the configured server URL
	IgnoreCRLExpiration               *bool  `json:"ignoreCRLExpiration,omitempty"`               // Switch to enable or disable ignore CRL expiration
	Name                              string `json:"name,omitempty"`                              // Friendly name of the certificate
	NonAutomaticCRLUpdatePeriod       *int   `json:"nonAutomaticCRLUpdatePeriod,omitempty"`       // Non automatic CRL update period
	NonAutomaticCRLUpdateUnits        string `json:"nonAutomaticCRLUpdateUnits,omitempty"`        // Unit of time of non automatic CRL update
	RejectIfNoStatusFromOCSP          *bool  `json:"rejectIfNoStatusFromOCSP,omitempty"`          // Switch to reject certificate if there is no status from OCSP
	RejectIfUnreachableFromOCSP       *bool  `json:"rejectIfUnreachableFromOCSP,omitempty"`       // Switch to reject certificate if unreachable from OCSP
	SelectedOCSpService               string `json:"selectedOCSPService,omitempty"`               // Name of selected OCSP Service
	Status                            string `json:"status,omitempty"`                            //
	TrustForCertificateBasedAdminAuth *bool  `json:"trustForCertificateBasedAdminAuth,omitempty"` // Trust for Certificate based Admin authentication
	TrustForCiscoServicesAuth         *bool  `json:"trustForCiscoServicesAuth,omitempty"`         // Trust for authentication of Cisco Services
	TrustForClientAuth                *bool  `json:"trustForClientAuth,omitempty"`                // Trust for client authentication and Syslog
	TrustForIseAuth                   *bool  `json:"trustForIseAuth,omitempty"`                   // Trust for authentication within Cisco ISE
}

//GetCsrs Get all Certificate Signing Requests from PAN
/*  This API supports filtering, sorting and pagination.

Filtering and sorting are supported for the following attributes:


friendlyName

subject

timeStamp


Supported Date Format: yyyy-MM-dd HH:mm:ss.SSS

Supported Operators: EQ, NEQ, GT and LT




@param getCSRsQueryParams Filtering parameter
*/
func (s *CertificatesService) GetCsrs(getCSRsQueryParams *GetCsrsQueryParams) (*ResponseCertificatesGetCsrs, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/certificate-signing-request"

	queryString, _ := query.Values(getCSRsQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseCertificatesGetCsrs{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetCsrs")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificatesGetCsrs)
	return result, response, err

}

//ExportCsr Export a CSR for a given CSR ID and hostname
/* Response of this API carries a CSR corresponding to the requested ID.

@param hostname hostname path parameter. Hostname to which the CSR belongs.
@param id id path parameter. ID of the CSR to be exported.
*/
func (s *CertificatesService) ExportCsr(hostname string, id string) (FileDownload, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/certificate-signing-request/export/{hostname}/{id}"
	path = strings.Replace(path, "{hostname}", fmt.Sprintf("%v", hostname), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Get(path)

	fdownload := FileDownload{}
	if err != nil {
		return fdownload, nil, err
	}

	if response.IsError() {
		return fdownload, response, fmt.Errorf("error with operation ExportTrustedCertificate")
	}

	fdownload.FileData = response.Body()
	headerVal := response.Header()["Content-Disposition"][0]
	fname := strings.SplitAfter(headerVal, "=")
	fdownload.FileName = strings.ReplaceAll(fname[1], "\"", "")

	getCSFRToken(response.Header())
	return fdownload, response, err

}

//GetCsrByID Get the certificate signing request for a given ID
/* This API displays details of a certificate signing request of a particular node for a given hostname and ID.

@param hostName hostName path parameter. Name of the host of which CSR's should be returned
@param id id path parameter. ID of the Certificate Signing Request returned
*/
func (s *CertificatesService) GetCsrByID(hostName string, id string) (*ResponseCertificatesGetCsrByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/certificate-signing-request/{hostName}/{id}"
	path = strings.Replace(path, "{hostName}", fmt.Sprintf("%v", hostName), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCertificatesGetCsrByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetCsrById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificatesGetCsrByID)
	return result, response, err

}

//GetSystemCertificates Get all system certificates of a particular node
/*  This API supports filtering, sorting and pagination.

Filtering and sorting supported for the following attributes:


friendlyName

issuedTo

issuedBy

validFrom


Supported Date Format: yyyy-MM-dd HH:mm:ss

Supported Operators: EQ, NEQ, GT and LT


expirationDate


Supported Date Format: yyyy-MM-dd HH:mm:ss

Supported Operators: EQ, NEQ, GT and LT




@param hostName hostName path parameter. Name of the host for which the system certificates should be returned
@param getSystemCertificatesQueryParams Filtering parameter
*/
func (s *CertificatesService) GetSystemCertificates(hostName string, getSystemCertificatesQueryParams *GetSystemCertificatesQueryParams) (*ResponseCertificatesGetSystemCertificates, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/system-certificate/{hostName}"
	path = strings.Replace(path, "{hostName}", fmt.Sprintf("%v", hostName), -1)

	queryString, _ := query.Values(getSystemCertificatesQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseCertificatesGetSystemCertificates{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSystemCertificates")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificatesGetSystemCertificates)
	return result, response, err

}

//GetSystemCertificateByID Get system certificate of a particular node by ID
/* This API provides details of a system certificate of a particular node based on given hostname and ID.

@param hostName hostName path parameter. Name of the host of which system certificates should be returned
@param id id path parameter. ID of the system certificate
*/
func (s *CertificatesService) GetSystemCertificateByID(hostName string, id string) (*ResponseCertificatesGetSystemCertificateByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/system-certificate/{hostName}/{id}"
	path = strings.Replace(path, "{hostName}", fmt.Sprintf("%v", hostName), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCertificatesGetSystemCertificateByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSystemCertificateById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificatesGetSystemCertificateByID)
	return result, response, err

}

//GetTrustedCertificates Get list of all trusted certificates
/*  This API supports Filtering, Sorting and Pagination.

Filtering and Sorting are supported for the following attributes:


friendlyName

subject

issuedTo

issuedBy

validFrom


Supported Date Format: yyyy-MM-dd HH:mm:ss

Supported Operators: EQ, NEQ, GT and LT


expirationDate


Supported Date Format: yyyy-MM-dd HH:mm:ss

Supported Operators: EQ, NEQ, GT and LT


status


Allowed values: enabled, disabled

Supported Operators: EQ, NEQ




Note:
 ISE internal CA certificates will not be exported.

@param getTrustedCertificatesQueryParams Filtering parameter
*/
func (s *CertificatesService) GetTrustedCertificates(getTrustedCertificatesQueryParams *GetTrustedCertificatesQueryParams) (*ResponseCertificatesGetTrustedCertificates, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/trusted-certificate"

	queryString, _ := query.Values(getTrustedCertificatesQueryParams)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseCertificatesGetTrustedCertificates{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTrustedCertificates")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificatesGetTrustedCertificates)
	return result, response, err

}

//ExportTrustedCert Export a trust certificate by a given certificate ID
/* The response of this API carries a trusted certificate file mapped to the requested ID.

@param id id path parameter. ID of the Trusted Certificate to be exported.
*/
func (s *CertificatesService) ExportTrustedCert(id string) (FileDownload, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/trusted-certificate/export/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Get(path)

	fdownload := FileDownload{}
	if err != nil {
		return fdownload, nil, err
	}

	if response.IsError() {
		return fdownload, response, fmt.Errorf("error with operation ExportTrustedCertificate")
	}

	fdownload.FileData = response.Body()
	headerVal := response.Header()["Content-Disposition"][0]
	fname := strings.SplitAfter(headerVal, "=")
	fdownload.FileName = strings.ReplaceAll(fname[1], "\"", "")

	getCSFRToken(response.Header())
	return fdownload, response, err

}

// Alias of ExportTrustedCert Export a trust certificate by a given certificate ID
func (s *CertificatesService) ExportTrustedCertificate(id string) (FileDownload, *resty.Response, error) {
	return s.ExportTrustedCert(id)
}

//GetTrustedCertificateByID Get Trust Certificate By ID
/* This API can displays details of a Trust Certificate based on a given ID.

@param id id path parameter. ID of the trust certificate
*/
func (s *CertificatesService) GetTrustedCertificateByID(id string) (*ResponseCertificatesGetTrustedCertificateByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/trusted-certificate/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCertificatesGetTrustedCertificateByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTrustedCertificateById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificatesGetTrustedCertificateByID)
	return result, response, err

}

//GenerateCsr Generate a Certificate Signing Request (CSR)
/*
Generate a certificate signing request for Multi-Use, Admin, EAP Authentication, RADIUS DTLS, PxGrid, SAML, Portal and IMS Services.
 Following parameters are present in the POST request body




PARAMETER

DESCRIPTION

EXAMPLE





hostnames

List of Cisco ISE node hostnames for which CSRs should be generated

"hostnames": ["ise-host1", "ise-host2"]



allowWildCardCert

Allow use of wildCards in certificates

"allowWildCardCert": false



keyLength
*required

Length of the key used for CSR generation.

"keyLength": "512"



keyType
*required

Type of key used for CSR generation either RSA or ECDSA.

"keyType": "RSA"



digestType
*required

Hash algorithm used for signing CSR.

"digestType": "SHA-256"



usedFor
*required

Certificate usage.

"usedFor": "MULTI-USE"



certificatePolicies

Certificate policy OID or list of OIDs that the certificate should conform to. Use comma or space to separate the OIDs.

"certificatePolicies": "Certificate Policies"



subjectCommonName
*required

Certificate common name (CN).

"subjectCommonName": "$FQDN$"



subjectOrgUnit

Certificate organizational unit (OU).

"subjectOrgUnit": "Engineering"



subjectOrg

Certificate organization (O).

"subjectOrg": "Cisco"



subjectCity

Certificate city or locality (L).

"subjectCity": "San Jose"


subjectState

Certificate state (ST).

"subjectState": "California"


subjectCountry

Certificate country (C).

"subjectCountry": "US"



sanDNS

Array of SAN (Subject Alternative Name) DNS entries (optional).

"sanDNS": ["ise.example.com"]


sanIP

Array of SAN IP entries (optional).

"sanIP": ["1.1.1.1"]


sanURI

Array of SAN URI entries (optional).

"sanURI": ["https://1.1.1.1"]



sanDir

Array of SAN DIR entries (optional).

"sanDir": ["CN=AAA,DC=COM,C=IL"]



portalGroupTag

Portal Group Tag when using certificate for PORTAL service

"portalGroupTag": "Default Portal Certificate Group"




NOTE:
For
allowWildCardCert
 to be false, the following parameter is mandatory:
hostnames

When certificate is selected to be used for Portal Service, the following parameter is mandatory:
portalGroupTag



*/
func (s *CertificatesService) GenerateCsr(requestCertificatesGenerateCSR *RequestCertificatesGenerateCsr) (*ResponseCertificatesGenerateCsr, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/certificate-signing-request"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCertificatesGenerateCSR).
		SetResult(&ResponseCertificatesGenerateCsr{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseCertificatesGenerateCsr{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GenerateCsr")
	}

	result := response.Result().(*ResponseCertificatesGenerateCsr)
	return result, response, err

}

//GenerateIntermediateCaCsr Generate an intermediate CA CSR (certificate signing request)
/* CSR generation for Intermediate Certificates.

 */
func (s *CertificatesService) GenerateIntermediateCaCsr() (*ResponseCertificatesGenerateIntermediateCaCsr, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/certificate-signing-request/intermediate-ca"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCertificatesGenerateIntermediateCaCsr{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseCertificatesGenerateIntermediateCaCsr{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GenerateIntermediateCaCsr")
	}

	result := response.Result().(*ResponseCertificatesGenerateIntermediateCaCsr)
	return result, response, err

}

//RegenerateIseRootCa Regenerate entire internal CA certificate chain including root CA on the primary PAN and subordinate CAs on the PSNs (Applicable only for internal CA service)
/* This API initiates regeneration of Cisco ISE root CA certificate chain. The response contains an ID which can be used to track the status.
Setting "removeExistingISEIntermediateCSR" to true removes existing Cisco ISE Intermediate CSR.

*/
func (s *CertificatesService) RegenerateIseRootCa(requestCertificatesRegenerateISERootCA *RequestCertificatesRegenerateIseRootCa) (*ResponseCertificatesRegenerateIseRootCa, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/ise-root-ca/regenerate"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCertificatesRegenerateISERootCA).
		SetResult(&ResponseCertificatesRegenerateIseRootCa{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseCertificatesRegenerateIseRootCa{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RegenerateIseRootCa")
	}

	result := response.Result().(*ResponseCertificatesRegenerateIseRootCa)
	return result, response, err

}

//RenewCerts Renew certificates of OCSP responder and Cisco ISE Messaging Service
/* This API initiates regeneration of certificates. The response contains an ID which can be used to track the status.

 */
func (s *CertificatesService) RenewCerts(requestCertificatesRenewCerts *RequestCertificatesRenewCerts) (*ResponseCertificatesRenewCerts, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/renew-certificate"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCertificatesRenewCerts).
		SetResult(&ResponseCertificatesRenewCerts{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseCertificatesRenewCerts{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RenewCerts")
	}

	result := response.Result().(*ResponseCertificatesRenewCerts)
	return result, response, err

}

// Alias of RenewCerts Renew certificates of OCSP responder and Cisco ISE Messaging Service
func (s *CertificatesService) RenewCertificates(requestCertificatesRenewCerts *RequestCertificatesRenewCerts) (*ResponseCertificatesRenewCerts, *resty.Response, error) {
	return s.RenewCerts(requestCertificatesRenewCerts)
}

//BindCsr Bind CA Signed Certificate
/*
Bind CA Signed Certificate.

NOTE:
This API requires an existing certificate signing request, and the root certificate must already be trusted.

NOTE:
The certificate may have a validity period greater than 398 days. It may be untrusted by many browsers.

NOTE:
Request parameters accepting True and False as input can be replaced by 1 and 0 respectively.

Following parameters are used in the POST body




PARAMETER

DESCRIPTION

EXAMPLE





name

Friendly name of the certificate.

"name": "CA Signed Certificate"



data
*required

 Plain-text contents of the certificate file. Every space needs to be replaced with a newline escape sequence (\n).
 Use the command
awk 'NF {sub(/\r/, ""); printf "%s\\n",$0;}' <<your .pem file>>
 to extract data from certificate file.

"data": "Plain-text contents of the certificate file."



allowExtendedValidity
*required

Allow the certificates with validity greater than 398 days.

"allowExtendedValidity": true



allowOutOfDateCert
*required

 Allow out of date certificates.
SECURITY ALERT:
We recommend to set the parameter the parameter
allowOutOfDateCert
 as
false
 to avoid binding of expired certificates (not secure).

"allowOutOfDateCert": true



allowReplacementOfCertificates
*required

Allow Replacement of certificates.

"allowReplacementOfCertificates": true



allowReplacementOfPortalGroupTag
*required

Allow Replacement of Portal Group Tag.

"allowReplacementOfPortalGroupTag": true


admin

Use certificate to authenticate the Cisco ISE Admin Portal

"admin": false


eap

Use certificate for EAP protocols that use SSL/TLS tunneling

"eap": false



radius

Use certificate for RADSec server

"radius": false



pxgrid

Use certificate for the pxGrid Controller

"pxgrid": false



ims

Use certificate for the Cisco ISE Messaging Service

"ims": false



saml

Use certificate for SAML Signing

"saml": false



portal

Use certificate for portal

"portal": false



portalGroupTag

Portal Group Tag for using certificate with portal role

"portalGroupTag": "Default Portal Certificate Group"



validateCertificateExtensions

Validate Certificate Extensions

"validateCertificateExtensions": false





Following roles can be used in any combinations




ROLE

DEFAULT

WARNING





Admin

False

Enabling admin role for this certificate causes an application server restart on the selected node.
Note:
 Make sure that the required certificate chain is imported under Trusted Certificates.



EAP Authentication

False

Only one system certificate can be used for EAP. Assigning EAP to this certificate removes the assignment from another certificate.
Note:
 Make sure that the required certificate chain is imported under Trusted Certificates.



RADIUS DTLS

False

Only one system certificate can be used for DTLS. Assigning DTLS to this certificate removes the assignment from another certificate.
Note:
 Make sure that the required certificate chain is imported under Trusted Certificates



SAML

False

SAML cannot be used with other Usage. Enabling SAML unchecks all other Usage.
Note:
 Make sure that the required certificate chain is imported under Trusted Certificates.





*/
func (s *CertificatesService) BindCsr(requestCertificatesBindCSR *RequestCertificatesBindCsr) (*ResponseCertificatesBindCsr, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/signed-certificate/bind"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCertificatesBindCSR).
		SetResult(&ResponseCertificatesBindCsr{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseCertificatesBindCsr{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BindCsr")
	}

	result := response.Result().(*ResponseCertificatesBindCsr)
	return result, response, err

}

//ExportSystemCert Export a system certificate with a given a certificate ID
/*
Export System Certificate.
 Following parameters are used in the POST body



PARAMETER

DESCRIPTION

EXAMPLE





id
*required

ID of a System Certificate.

"id": "CERT-ID"



hostName
*required

Name of the host for which the system certificate should be exported

"hostName": "ise-node-001"



export

 One of the following options is required:

"CERTIFICATE" :
Export only certificate without private key

"CERTIFICATE_WITH_PRIVATE_KEY" :
Export both certificate and private key (
"certificatePassword"
 is required).



"export": "CERTIFICATE_WITH_PRIVATE_KEY"



password
*required

Certificate password (required if
"export" : CERTIFICATE_WITH_PRIVATE_KEY
).
Password constraints:


Alphanumeric

Minimum of 8 Characters

Maximum of 100 Characters



"password": "certificate password"




NOTE:
The response of this API carries a ZIP file containing the certificate and private key if  the request contains
"export" : "CERTIFICATE_WITH_PRIVATE_KEY"
. If the request body contains
"export" : "CERTIFICATE"
, the response carries a ZIP file containing only the certificate.

WARNING:
Exporting a private key is not a secure operation. It could lead to possible exposure of the private key.


*/
func (s *CertificatesService) ExportSystemCert(requestCertificatesExportSystemCert *RequestCertificatesExportSystemCert) (FileDownload, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/system-certificate/export"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCertificatesExportSystemCert).
		SetError(&Error).
		Post(path)

	fdownload := FileDownload{}
	if err != nil {
		return fdownload, nil, err
	}

	if response.IsError() {
		return fdownload, response, fmt.Errorf("error with operation ExportTrustedCertificate")
	}

	fdownload.FileData = response.Body()
	headerVal := response.Header()["Content-Disposition"][0]
	fname := strings.SplitAfter(headerVal, "=")
	fdownload.FileName = strings.ReplaceAll(fname[1], "\"", "")

	getCSFRToken(response.Header())
	return fdownload, response, err

}

// Alias of ExportSystemCert Export a system certificate with a given a certificate ID
func (s *CertificatesService) ExportSystemCertificate(requestCertificatesExportSystemCert *RequestCertificatesExportSystemCert) (FileDownload, *resty.Response, error) {
	return s.ExportSystemCert(requestCertificatesExportSystemCert)
}

//GenerateSelfSignedCertificate Generate self-signed certificate in Cisco ISE
/*
Generate Self-signed Certificate

NOTE:
The certificate may have a validity period greater than 398 days. It may be untrusted by many browsers.

NOTE:
Request parameters accepting True and False as input can be replaced by 1 and 0 respectively.

NOTE:
Wildcard certificate and SAML certificate can be generated only on the primary PAN or a standalone node.


Following parameters are used in the POST body




PARAMETER

DESCRIPTION

EXAMPLE





hostName
*required

Hostname or FQDN of the node in which the certificate needs to be created.

"hostName": "ise-node-001"



name

Friendly name of the certificate.

"name": "Self-signed System Certificate"



subjectCommonName

 Certificate common name (CN)

NOTE:
CN is Mandatory if SAN not configured.
>Subject can contain a multi-valued CN. For multi-valued RDNs, follow the format "CN=value1, CN=value2"


"subjectCommonName": "$FQDN$"



subjectOrgUnit

 Certificate organizational unit (OU)

NOTE:
Subject can contain a multi-valued OU. For multi-valued RDNs, follow the format "OU=value1, OU=value2"

"subjectOrgUnit": "Engineering"



subjectOrg

 Certificate organization (O)

NOTE:
Subject can contain multi-valued O fields. For multi-valued RDNs, follow the format "O=value1, O=value2"

"subjectOrg": "Cisco"



subjectCity

Certificate city or locality (L)

"subjectCity": "San Jose"



subjectState

Certificate state (ST)

"subjectState": "California"



subjectCountry

Certificate country (C)

"subjectCountry": "US"



sanDNS

Array of SAN (Subject Alternative Name) DNS entries

"sanDNS": ["ise.example.com"]



sanIP

Array of SAN IP address entries

"sanIP": ["1.1.1.1"]



sanURI

Array of SAN URI entries

"sanURI": ["https://1.1.1.1"]



keyType
*required

Algorithm to use for certificate public key creation.

"keyType": "RSA"



keyLength
*required

Bit size of the public key.

"keyLength": "4096"



digestType
*required

Digest to sign with.

"digestType": "SHA-384"



certificatePolicies

Certificate policy OID or list of OIDs that the certificate should conform to. Use comma or space to separate the OIDs.

"certificatePolicies": "Certificate Policies"



expirationTTL
*required

 Certificate expiration value.

NOTE:
Expiration TTL should be within Unix time limit

"expirationTTL": 2



expirationTTLUnit
*required

Certificate expiration unit.

"expirationTTLUnit": "years"



admin

Use certificate to authenticate the Cisco ISE Admin Portal

"admin": false



eap

Use certificate for EAP protocols that use SSL/TLS tunneling

"eap": false



radius

Use certificate for RADSec server

"radius": false



pxgrid

Use certificate for the pxGrid controller

"pxgrid": false



saml

Use certificate for SAML Signing

"saml": false



portal

Use certificate for portal

"portal": false



portalGroupTag

Portal Group Tag for using certificate with portal role

"portalGroupTag": "Default Portal Certificate Group"



allowReplacementOfPortalGroupTag
*required

Allow Replacement of Portal Group Tag.

"allowReplacementOfPortalGroupTag": true



allowWildCardCertificates

Allow use of WildCards in certificates

"allowWildCardCertificates": false



allowReplacementOfCertificates
*required

Allow replacement of certificates.

"allowReplacementOfCertificates": true



allowExtendedValidity
*required

Allow generation of self-signed certificate with validity greater than 398 days.

"allowExtendedValidity": true



allowRoleTransferForSameSubject
*required

Allow the transfer of roles to certificates with same subject.
 If the matching certificate on Cisco ISE has either admin or portal role and if the request has admin or portal role selected along with
allowRoleTransferForSameSubject
 parameter as true, a self-signed certificate would be generated with both admin and portal role enabled.

"allowRoleTransferForSameSubject": true



allowPortalTagTransferForSameSubject
*required

Acquire the group tag of the matching certificate. If the request portal groug tag is different from the group tag of the matching certificate (If matching certificate in Cisco ISE has portal role enabled), a self-signed certificate would be generated by acquiring the group tag of the matching certificate if the
allowPortalTagTransferForSameSubject
 parameter is true.

"allowPortalTagTransferForSameSubject": true



allowSanDnsBadName
*required

 Allow generation of self-signed certificates with bad common name & SAN values such as "example.org.","invalid.","test.","localhost" and so on.
SECURITY ALERT:
We recommend to set the parameter
allowSanDnsBadName
 as
false
 to avoid generation of certificates with bad Common Name & SAN Values which are not secure.

"allowSanDnsBadName": true



allowSanDnsNonResolvable
*required

Allow generation of self-signed certificate with non resolvable Common Name or SAN Values .

"allowSanDnsNonResolvable": true








ROLE

DEFAULT

WARNING





Admin

False

Enabling Admin role for this certificate causes an application server restart on the selected node.



EAP Authentication

False

Only one system certificate can be used for EAP. Assigning EAP to this certificate removes the assignment from another certificate.



RADIUS DTLS

False

Only one system certificate can be used for DTLS. Assigning DTLS to this certificate removes the assignment from another certificate.



SAML

False

SAML cannot be used with other Usage.





*/
func (s *CertificatesService) GenerateSelfSignedCertificate(requestCertificatesGenerateSelfSignedCertificate *RequestCertificatesGenerateSelfSignedCertificate) (*ResponseCertificatesGenerateSelfSignedCertificate, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/system-certificate/generate-selfsigned-certificate"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCertificatesGenerateSelfSignedCertificate).
		SetResult(&ResponseCertificatesGenerateSelfSignedCertificate{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseCertificatesGenerateSelfSignedCertificate{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GenerateSelfSignedCertificate")
	}

	result := response.Result().(*ResponseCertificatesGenerateSelfSignedCertificate)
	return result, response, err

}

//ImportSystemCert Import system certificate in Cisco ISE
/*
Import an X509 certificate as a system certificate.

NOTE:


This API is used to import a certificate on a specific node mentioned in the server section of the URL. To import a certificate on a secondary node, execute this RESTApi directly on a secondary node by specifying the server name of the secondary node in the URL
 Example: "<https://secondary-ise-node>/api/v1/certs/system-certificate/import"

The certificate may have a validity period of more than 398 days. It may be untrusted by many browsers.

Request parameters accepting True and False as input can be replaced by 1 and 0 respectively.


Following parameters are used in the POST body




PARAMETER

DESCRIPTION

EXAMPLE





name

Friendly name of the certificate.

"name": "System certificate"



password
*required

Password of the certificate to be imported.

"password": "certificate password"



data
*required

 Plain-text contents of the certificate file. Every space needs to be replaced with a newline escape sequence (\n).
 Use the command
awk 'NF {sub(/\r/, ""); printf "%s\\n",$0;}' <<your .pem file>>
 to extract data from the certificate file.

"data": "Plain-text contents of the certificate file."



privateKeyData
*required

 Plain-text contents of the private key file. Every space needs to be replaced with a newline escape sequence (\n). Use the command
awk 'NF {sub(/\r/, ""); printf "%s\\n",$0;}' <<your .pem file>>
 to extract privateKeyData from private key file.

"data": "Plain-text contents of the private key file."



allowOutOfDateCert
*required

 Allow out of date certificates .
SECURITY ALERT:
We recommend to set the parameter
allowOutOfDateCert
 as
false
 to avoid the import of expired certificates (not Secure).

"allowOutOfDateCert": true



allowSHA1certificates
*required

 Allow import of certificate with signature that uses the SHA-1 hashing algorithm and is considered less secure .
SECURITY ALERT:
We recommend to set the parameter
allowSHA1certificates
as
false
 to avoid the import of SHA1 based certificates (less secure).

"allowSHA1certificates": true



allowExtendedValidity
*required

Allow the certificates greater than validity of 398 days.

"allowExtendedValidity": true



allowRoleTransferForSameSubject

password
*required

Allow the transfer of roles to certificates with the same subject

"allowRoleTransferForSameSubject": true



allowPortalTagTransferForSameSubject

password
*required

Acquire the group tag of the matching certificate

"allowPortalTagTransferForSameSubject": true



admin

Use the certificate to authenticate the Cisco ISE admin portal

"admin": false



eap

Use the certificate for EAP protocols that use SSL/TLS tunneling

"eap": false



radius

Use the certificate for RADSec server

"radius": false



pxgrid

Use the certificate for the pxGrid Controller

"pxgrid": false



ims

Use the certificate for the Cisco ISE messaging service

"ims": false



saml

Use the certificate for SAML Signing

"saml": false



portal

Use the certificate for portal

"portal": false



portalGroupTag

Portal Group Tag for using certificate with portal role

"portalGroupTag": "Default Portal certificate Group"



allowReplacementOfPortalGroupTag
*required

Allow Replacement of Portal Group Tag .

"allowReplacementOfPortalGroupTag": true



allowWildCardcertificates

Allow use of wildcards in certificates

"allowWildCardcertificates": false



validatecertificateExtensions

Validate certificate extensions

"validatecertificateExtensions": false





Following roles can be used in any combinations




ROLE

DEFAULT

WARNING





Admin

False

Enabling Admin role for this certificate causes an application server restart on the selected node.
Note:
 Make sure the required certificate chain is imported under Trusted Certificates



EAP Authentication

False

Only one system certificate can be used for EAP. Assigning EAP to this certificate removes the assignment from another certificate.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates



RADIUS DTLS

False

Only one system certificate can be used for DTLS. Assigning DTLS to this certificate removes the assignment from another certificate.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates



SAML

False

SAML cannot be used with other Usage. Enabling SAML unchecks all other Usage.
Note:
 Make sure the required certificate chain is imported under Trusted Certificates





*/
func (s *CertificatesService) ImportSystemCert(requestCertificatesImportSystemCert *RequestCertificatesImportSystemCert) (*ResponseCertificatesImportSystemCert, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/system-certificate/import"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCertificatesImportSystemCert).
		SetResult(&ResponseCertificatesImportSystemCert{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseCertificatesImportSystemCert{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ImportSystemCert")
	}

	result := response.Result().(*ResponseCertificatesImportSystemCert)
	return result, response, err

}

// Alias of ImportSystemCert Import system certificate in Cisco ISE
func (s *CertificatesService) ImportSystemCertificate(requestCertificatesImportSystemCert *RequestCertificatesImportSystemCert) (*ResponseCertificatesImportSystemCert, *resty.Response, error) {
	return s.ImportSystemCert(requestCertificatesImportSystemCert)
}

//ImportTrustCert Add root certificate to the Cisco ISE truststore
/*
Import an X509 certificate as a trust certificate

NOTE:
Request parameters accepting True and False as input can be replaced by 1< and 0< respectively.
 Following parameters are used in the POST body:




PARAMETER

DESCRIPTION

EXAMPLE





name

Friendly name of the certificate

"name": "Trust Certificate"



description

Description of the certificate

"description": "Imported Trust Certificate"



data
*required

 Plain-text contents of the certificate file. Every space needs to be replaced with a newline escape sequence (\n).
 Use the command
awk 'NF {sub(/\r/, ""); printf "%s\\n",$0;}' <<your .pem file>>
 to extract data from the certificate file.

"data": "Plain-text contents of the certificate file."



allowOutOfDateCert
*required

 Allow out of date certificates.
SECURITY ALERT:
We recommend to set the parameter
allowOutOfDateCert
 as
false
 to avoid the import of expired certificates (not secure).

"allowOutOfDateCert": true



allowSHA1Certificates
*required

 Allow import of certificate with signature that uses SHA-1 hashing algorithm and is considered less secure.
SECURITY ALERT:
We recommend to set the parameter
allowSHA1Certificates
 as
false
 to avoid the import of SHA1 based certificates (less secure).

"allowSHA1Certificates": true



allowBasicConstraintCAFalse
*required

 Allow certificates with Basic Constraints CA Field as False.
SECURITY ALERT:
We recommend to set the parameter
allowBasicConstraintCAFalse
 as
false
 to avoid the import of certificates with
Basic Constraints CA Field
 set as
False
 (not Secure).

"allowBasicConstraintCAFalse": true



trustForIseAuth

Trust for authentication within Cisco ISE

"trustForIseAuth": false



trustForClientAuth

Trust for client authentication and syslog

"trustForClientAuth": false



trustForCertificateBasedAdminAuth

Trust for certificate based admin authentication

"trustForCertificateBasedAdminAuth": false



trustForCiscoServicesAuth

Trust for authentication of Cisco services

"trustForCiscoServicesAuth": false



validateCertificateExtensions

Validate extensions for trust certificate

"validateCertificateExtensions": false




NOTE
: If
name
 is not set, a default name with the following format is used where
nnnnn
 is a unique number:
common-name#issuer#nnnnn
 You can always change the friendly name later by editing the certificate.


You must choose how this certificate is trusted in Cisco ISE. The objective here is to distinguish between certificates that are used for trust within a Cisco ISE deployment and public certificates that are used to trust Cisco services. We recommend not using a given certificate for both purposes.




Trusted For

Usage





Authentication within Cisco ISE

Use
"trustForIseAuth":true
 if the certificate is used for trust within Cisco ISE, such as for secure communication between Cisco ISE nodes



Client authentication and Syslog

Use
"trustForClientAuth":true
 if the certificate is to be used for authentication of endpoints that contact Cisco ISE over the EAP protocol. This is also used if the certificate is used to trust a Syslog server. Make sure to have keyCertSign bit asserted under KeyUsage extension for this certificate.
Note:
 "" can be set true only if the "trustForIseAuth" has been set true.



Certificate based admin authentication

Use
"trustForCertificateBasedAdminAuth":true
 if the certificate is used for trust within Cisco ISE, such as for secure communication between Cisco ISE nodes
Note:
trustForCertificateBasedAdminAuth
 can be set true only if both
trustForIseAuth
 and
trustForClientAuth
 are true.



Authentication of Cisco Services

 Use
"trustForCiscoServicesAuth":true
 if the certificate is to be used for trusting external Cisco services, such as Feed Service.





*/
func (s *CertificatesService) ImportTrustCert(requestCertificatesImportTrustCert *RequestCertificatesImportTrustCert) (*ResponseCertificatesImportTrustCert, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/trusted-certificate/import"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCertificatesImportTrustCert).
		SetResult(&ResponseCertificatesImportTrustCert{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseCertificatesImportTrustCert{}, response, nil
		}
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ImportTrustCert")
	}

	result := response.Result().(*ResponseCertificatesImportTrustCert)
	return result, response, err

}

// Alias of ImportTrustCert Add root certificate to the Cisco ISE truststore
func (s *CertificatesService) ImportTrustCertificate(requestCertificatesImportTrustCert *RequestCertificatesImportTrustCert) (*ResponseCertificatesImportTrustCert, *resty.Response, error) {
	return s.ImportTrustCert(requestCertificatesImportTrustCert)
}

//UpdateSystemCert Update data for existing system certificate
/*
Update a System Certificate.

NOTE:
Renewing a certificate causes an application server restart on the selected node.

NOTE:
Request parameters accepting True and False as input can be replaced by 1 and 0 respectively.

Following parameters are used in the POST body




PARAMETER

DESCRIPTION

EXAMPLE





name

Friendly name of the certificate.

"name": "System Certificate"



description

Description of the certificate

"description": "Description of certificate"



admin

Use certificate to authenticate the Cisco ISE Admin Portal

"admin": false



eap

Use certificate for EAP protocols that use SSL/TLS tunneling

"eap": false



radius

Use certificate for RADSec server

"radius": false



pxgrid

Use certificate for the pxGrid Controller

"pxgrid": false



ims

Use certificate for the Cisco ISE Messaging Service

"ims": false



saml

Use certificate for SAML Signing

"saml": false



portal

Use certificate for portal

"portal": false



portalGroupTag

Portal Group Tag for using certificate with portal role

"portalGroupTag": "Default Portal Certificate Group"



allowReplacementOfPortalGroupTag
*required

Allow Replacement of Portal Group Tag.

"allowReplacementOfPortalGroupTag": true



allowRoleTransferForSameSubject
*required

Allow transfer of roles to certificates with same subject.

"allowRoleTransferForSameSubject": true



allowPortalTagTransferForSameSubject
*required

Acquire group tag of the matching certificate.

"allowPortalTagTransferForSameSubject": true



renewSelfSignedCertificate

Renew Self-signed Certificate

"renewSelfSignedCertificate": false



expirationTTLPeriod

Expiration Period

"expirationTTLPeriod": 365



expirationTTLUnits

Expiration Units in one of the below formats

days / weeks / months / years



"expirationTTLUnits": "days"





Following roles can be used in any combinations




ROLE

DEFAULT

WARNING





Admin

False

Enabling Admin role for this certificate causes an application server restart on the selected node.
Note:
 Make sure that the required certificate chain is imported under Trusted Certificates



EAP Authentication

False

Only one system certificate can be used for EAP. Assigning EAP to this certificate removes the assignment from another certificate.
Note:
 Make sure that the required certificate chain is imported under Trusted Certificates



RADIUS DTLS

False

Only one system certificate can be used for DTLS. Assigning DTLS to this certificate removes the assignment from another certificate.
Note:
 Make sure that the required certificate chain is imported under Trusted Certificates



SAML

False

SAML cannot be used with other usage. Enabling SAML unchecks all other usage.
Note:
 Make sure that the required certificate chain is imported under Trusted Certificates





@param id id path parameter. ID of the System Certificate to be updated
@param hostName hostName path parameter. Name of host whose certificate needs to be updated
*/
func (s *CertificatesService) UpdateSystemCert(id string, hostName string, requestCertificatesUpdateSystemCert *RequestCertificatesUpdateSystemCert) (*ResponseCertificatesUpdateSystemCert, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/system-certificate/{hostName}/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{hostName}", fmt.Sprintf("%v", hostName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCertificatesUpdateSystemCert).
		SetResult(&ResponseCertificatesUpdateSystemCert{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseCertificatesUpdateSystemCert{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSystemCert")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificatesUpdateSystemCert)
	return result, response, err

}

// Alias of UpdateSystemCert Update data for existing system certificate
func (s *CertificatesService) UpdateSystemCertificate(id string, hostName string, requestCertificatesUpdateSystemCert *RequestCertificatesUpdateSystemCert) (*ResponseCertificatesUpdateSystemCert, *resty.Response, error) {
	return s.UpdateSystemCert(id, hostName, requestCertificatesUpdateSystemCert)
}

//UpdateTrustedCertificate Update the trust certificate already present in the Cisco ISE trust store
/*
Update a trusted certificate present in Cisco ISE trust store.
 The following parameters are used in the PUT request body



PARAMETER

DESCRIPTION

EXAMPLE





name
*required

Friendly name of the certificate.

"name": "Trust Certificate"



status

Status of the certificate

"status": "Enabled"



description

Description of the certificate

"description": "Certificate for secure connection to cisco.com"



trustForIseAuth

Trust for authentication within Cisco ISE

"trustForIseAuth": false



trustForClientAuth

Trust for client authentication and Syslog

"trustForClientAuth": false



trustForCertificateBasedAdminAuth

Trust for certificate based Admin authentication

"trustForCertificateBasedAdminAuth": false



trustForCiscoServicesAuth

Trust for authentication of Cisco Services

"trustForCiscoServicesAuth": false



enableOCSPValidation

Switch to enable or disable OCSP Validation

"enableOCSPValidation": false



selectedOCSPService

Name of selected OCSP Service

"selectedOCSPService": "INTERNAL_OCSP_SERVICE"



rejectIfNoStatusFromOCSP

Switch to reject certificate if there is no status from OCSP

"rejectIfNoStatusFromOCSP": false



rejectIfUnreachableFromOCSP

Switch to reject certificate if unreachable from OCSP

"rejectIfUnreachableFromOCSP": false



downloadCRL

Switch to enable or disable download of CRL

"downloadCRL": false



crlDistributionUrl

Certificate Revocation List Distribution URL

"crlDistributionUrl": "CRL distribution URL"



automaticCRLUpdate

Switch to enable or disable automatic CRL update

"automaticCRLUpdate": false



automaticCRLUpdatePeriod

Automatic CRL update period

"automaticCRLUpdatePeriod": 5



automaticCRLUpdateUnits

Unit of time for automatic CRL update

"automaticCRLUpdateUnits": "Minutes"



nonAutomaticCRLUpdatePeriod

Non automatic CRL update period

"nonAutomaticCRLUpdatePeriod": 1



nonAutomaticCRLUpdateUnits

Unit of time of non automatic CRL update

"nonAutomaticCRLUpdateUnits": "Hours"



crlDownloadFailureRetries

If CRL download fails, wait time before retry

"crlDownloadFailureRetries": 10



crlDownloadFailureRetriesUnits

Unit of time before retry if CRL download fails

"crlDownloadFailureRetriesUnits": "Minutes"



enableServerIdentityCheck

Switch to enable or disable verification if HTTPS or LDAP server certificate name fits the configured server URL

"enableServerIdentityCheck": false



authenticateBeforeCRLReceived

Switch to enable or disable CRL Verification if CRL is not Received

"authenticateBeforeCRLReceived": false



ignoreCRLExpiration

Switch to enable or disable ignore CRL Expiration

"ignoreCRLExpiration": false








Trusted For

Usage





Authentication within Cisco ISE

Use
"trustForIseAuth":true
 if the certificate is used for trust within Cisco ISE, such as for secure communication between Cisco ISE nodes



Client authentication and Syslog

Use
"trustForClientAuth":true
 if the certificate is to be used for authentication of endpoints that contact Cisco ISE over the EAP protocol. Also check this box if certificate is used to trust a Syslog server. Make sure to have keyCertSign bit asserted under KeyUsage extension for this certificate.
Note:
 "trustForClientAuth" can be set true only if "trustForIseAuth" has been set true.



Certificate based admin authentication

Use
"trustForCertificateBasedAdminAuth":true
 if the certificate is used for trust within Cisco ISE, such as for secure communication between Cisco ISE nodes
Note:
 "trustForCertificateBasedAdminAuth" can be set true only if "trustForIseAuth" and "trustForClientAuth" are true.



Authentication of Cisco Services

 Use
"trustForCiscoServicesAuth":true
 if the certificate is to be used for trusting external Cisco services, such as Feed Service.







OCSP Configuration

Usage





Validation against OCSP service

Use
"enableOCSPValidation":true
 to validate the certificate against OCSP service mentioned in the field
selectedOCSPService
.



OCSP Service name

Use
"selectedOCSPService":"Name of OCSP Service"
 to mention the OCSP service name against which the certificate should be validated.
Note:

selectedOCSPService
 value is used if
enableOCSPValidation
 has been set true.



Reject the request if OCSP returns UNKNOWN status

Use
"rejectIfNoStatusFromOCSP":true
 to reject the certificate if the OCSP service returns UNKNOWN status.
Note:

"rejectIfNoStatusFromOCSP":true
 can be used only if the parameter
enableOCSPValidation
 has been set true.



Reject the request if OCSP Responder is unreachable

 Use
"rejectIfUnreachableFromOCSP":true
 to reject the certificate if the OCSP service is unreachable.
Note:

"rejectIfUnreachableFromOCSP":true
 can be used only if
enableOCSPValidation
 has been set true.







Certificate Revocation List Configuration

Usage





Validation against CRL

Use
"downloadCRL":true
 to validate the certificate against CRL downloaded from URL mentioned in the field
crlDistributionUrl



CRL distribution url

Use
"crlDistributionUrl"
 to specify the URL from where the CRL should be downloaded
Note:
 "crlDistributionUrl" value is used if "downloadCRL" has been set true.



Retrieve CRL time

Use
"automaticCRLUpdate":true
,
automaticCRLUpdatePeriod
, and
automaticCRLUpdatePeriod
 to set the time before which CRL is automatically retrieved prior to expiration Use
nonAutomaticCRLUpdatePeriod
 and
nonAutomaticCRLUpdateUnits
 to set the time period for CRL retrieval in loop.

Note:
 All the above fields can be used only if
"downloadCRL"
 has been set true.



If download fails

Use
"crlDownloadFailureRetries" and "crlDownloadFailureRetriesUnits"
 to set retry time period if CRL download fails
Note:
crlDownloadFailureRetries
 and
crlDownloadFailureRetriesUnits
 can be used only if
downloadCRL
 has been set true.



Enable Server Identity Check

Use
"enableServerIdentityCheck":true
 to verify that HTTPS or LDAPS server certificate name fits the configured server URL
Note:
"enableServerIdentityCheck":true
 can be used only if
downloadCRL
 has been set true.



Bypass CRL Verification if CRL is not Received

Use
"authenticateBeforeCRLReceived":true
 to bypass CRL Verification if CRL is not Received
Note:
"authenticateBeforeCRLReceived":true can be used only if
downloadCRL
 has been set true.



Ignore that CRL is not yet valid or has expired

 Use
"ignoreCRLExpiration":true
 to ignore if CRL is not yet valid or expired
Note:
"ignoreCRLExpiration":true
 can be used only if
downloadCRL
 has been set true.





Note:
boolean properties accept integers values as well, with 0 considered as false and other values being considered as true

@param id id path parameter. ID of the trust certificate
*/
func (s *CertificatesService) UpdateTrustedCertificate(id string, requestCertificatesUpdateTrustedCertificate *RequestCertificatesUpdateTrustedCertificate) (*ResponseCertificatesUpdateTrustedCertificate, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/trusted-certificate/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCertificatesUpdateTrustedCertificate).
		SetResult(&ResponseCertificatesUpdateTrustedCertificate{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		if err.Error() == emptyStringToJSONError {
			return &ResponseCertificatesUpdateTrustedCertificate{}, response, nil
		}
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateTrustedCertificate")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificatesUpdateTrustedCertificate)
	return result, response, err

}

//DeleteCsr Delete the certificate signing request for a given ID
/* This API deletes the certificate signing request of a particular node based on a given hostname and ID.

@param hostName hostName path parameter. Name of the host of which CSR's should be deleted
@param id id path parameter. ID of the Certificate Signing Request to be deleted
*/
func (s *CertificatesService) DeleteCsr(hostName string, id string) (*ResponseCertificatesDeleteCsr, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/certificate-signing-request/{hostName}/{id}"
	path = strings.Replace(path, "{hostName}", fmt.Sprintf("%v", hostName), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCertificatesDeleteCsr{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteCsr")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificatesDeleteCsr)
	return result, response, err

}

// Alias of DeleteCsr Delete the certificate signing request for a given ID
func (s *CertificatesService) DeleteCsrByID(hostName string, id string) (*ResponseCertificatesDeleteCsr, *resty.Response, error) {
	return s.DeleteCsr(hostName, id)
}

//DeleteSystemCertificateByID Delete System Certificate by ID and hostname
/* This API deletes a system certificate of a particular node based on the given hostname and ID.

@param hostName hostName path parameter. Name of the host from which system certificate needs to be deleted
@param id id path parameter. ID of the system certificate to be deleted
*/
func (s *CertificatesService) DeleteSystemCertificateByID(hostName string, id string) (*ResponseCertificatesDeleteSystemCertificateByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/system-certificate/{hostName}/{id}"
	path = strings.Replace(path, "{hostName}", fmt.Sprintf("%v", hostName), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCertificatesDeleteSystemCertificateByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteSystemCertificateById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificatesDeleteSystemCertificateByID)
	return result, response, err

}

//DeleteTrustedCertificateByID Delete a trusted certificate based on a given ID
/* This API deletes a Trust Certificate from Trusted Certificate Store based on a given ID.

@param id id path parameter. ID of the Trusted Certificate to be deleted
*/
func (s *CertificatesService) DeleteTrustedCertificateByID(id string) (*ResponseCertificatesDeleteTrustedCertificateByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/trusted-certificate/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCertificatesDeleteTrustedCertificateByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteTrustedCertificateById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificatesDeleteTrustedCertificateByID)
	return result, response, err

}
