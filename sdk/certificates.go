package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type CertificatesService service

type GetCsrsQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sort       string   `url:"sort,omitempty"`       //sort type - asc or desc
	SortBy     string   `url:"sortBy,omitempty"`     //sort column by which objects needs to be sorted
	Filter     []string `url:"filter,omitempty"`     //<div> <style type="text/css" scoped> .apiServiceTable td, .apiServiceTable th { padding: 5px 10px !important; text-align: left; } </style> <span> <b>Simple filtering</b> should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the <i>"filterType=or"</i> query string parameter. Each resource Data model description should specify if an attribute is a filtered field. </span> <br /> <table class="apiServiceTable"> <thead> <tr> <th>OPERATOR</th> <th>DESCRIPTION</th> </tr> </thead> <tbody> <tr> <td>EQ</td> <td>Equals</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> </tr> <tr> <td>GT</td> <td>Greater Than</td> </tr> <tr> <td>LT</td> <td>Less Then</td> </tr> <tr> <td>STARTSW</td> <td>Starts With</td> </tr> <tr> <td>NSTARTSW</td> <td>Not Starts With</td> </tr> <tr> <td>ENDSW</td> <td>Ends With</td> </tr> <tr> <td>NENDSW</td> <td>Not Ends With</td> </tr> <tr> <td>CONTAINS</td> <td>Contains</td> </tr> <tr> <td>NCONTAINS</td> <td>Not Contains</td> </tr> </tbody> </table> </div>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type GetSystemCertificatesQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sort       string   `url:"sort,omitempty"`       //sort type - asc or desc
	SortBy     string   `url:"sortBy,omitempty"`     //sort column by which objects needs to be sorted
	Filter     []string `url:"filter,omitempty"`     //<div> <style type="text/css" scoped> .apiServiceTable td, .apiServiceTable th { padding: 5px 10px !important; text-align: left; } </style> <span> <b>Simple filtering</b> should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the <i>"filterType=or"</i> query string parameter. Each resource Data model description should specify if an attribute is a filtered field. </span> <br /> <table class="apiServiceTable"> <thead> <tr> <th>OPERATOR</th> <th>DESCRIPTION</th> </tr> </thead> <tbody> <tr> <td>EQ</td> <td>Equals</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> </tr> <tr> <td>GT</td> <td>Greater Than</td> </tr> <tr> <td>LT</td> <td>Less Then</td> </tr> <tr> <td>STARTSW</td> <td>Starts With</td> </tr> <tr> <td>NSTARTSW</td> <td>Not Starts With</td> </tr> <tr> <td>ENDSW</td> <td>Ends With</td> </tr> <tr> <td>NENDSW</td> <td>Not Ends With</td> </tr> <tr> <td>CONTAINS</td> <td>Contains</td> </tr> <tr> <td>NCONTAINS</td> <td>Not Contains</td> </tr> </tbody> </table> </div>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
}

type GetTrustedCertificatesQueryParams struct {
	Page       int      `url:"page,omitempty"`       //Page number
	Size       int      `url:"size,omitempty"`       //Number of objects returned per page
	Sort       string   `url:"sort,omitempty"`       //sort type - asc or desc
	SortBy     string   `url:"sortBy,omitempty"`     //sort column by which objects needs to be sorted
	Filter     []string `url:"filter,omitempty"`     //<div> <style type="text/css" scoped> .apiServiceTable td, .apiServiceTable th { padding: 5px 10px !important; text-align: left; } </style> <span> <b>Simple filtering</b> should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the <i>"filterType=or"</i> query string parameter. Each resource Data model description should specify if an attribute is a filtered field. </span> <br /> <table class="apiServiceTable"> <thead> <tr> <th>OPERATOR</th> <th>DESCRIPTION</th> </tr> </thead> <tbody> <tr> <td>EQ</td> <td>Equals</td> </tr> <tr> <td>NEQ</td> <td>Not Equals</td> </tr> <tr> <td>GT</td> <td>Greater Than</td> </tr> <tr> <td>LT</td> <td>Less Then</td> </tr> <tr> <td>STARTSW</td> <td>Starts With</td> </tr> <tr> <td>NSTARTSW</td> <td>Not Starts With</td> </tr> <tr> <td>ENDSW</td> <td>Ends With</td> </tr> <tr> <td>NENDSW</td> <td>Not Ends With</td> </tr> <tr> <td>CONTAINS</td> <td>Contains</td> </tr> <tr> <td>NCONTAINS</td> <td>Not Contains</td> </tr> </tbody> </table> </div>
	FilterType string   `url:"filterType,omitempty"` //The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
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
	HostName           string                                   `json:"hostName,omitempty"`           // Hostname or IP address of the ISE node.
	ID                 string                                   `json:"id,omitempty"`                 // ID of the certificate.
	KeySize            string                                   `json:"keySize,omitempty"`            // Size of the cryptographic key used.
	Link               *ResponseCertificatesGetCsrsResponseLink `json:"link,omitempty"`               //
	SignatureAlgorithm string                                   `json:"signatureAlgorithm,omitempty"` // Algorithm used for encrypting CSR
	Subject            string                                   `json:"subject,omitempty"`            // Subject of the certificate. Includes Common Name (CN), Organizational Unit (OU), etc.
	TimeStamp          string                                   `json:"timeStamp,omitempty"`          // Timestamp of the certificate generation.
	UsedFor            string                                   `json:"usedFor,omitempty"`            // Services for which the certificate is used for(for eg- MGMT, GENERIC).
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
	HostName           string                                      `json:"hostName,omitempty"`           // Hostname or IP address of the ISE node.
	ID                 string                                      `json:"id,omitempty"`                 // ID of the certificate.
	KeySize            string                                      `json:"keySize,omitempty"`            // Size of the cryptographic key used.
	Link               *ResponseCertificatesGetCsrByIDResponseLink `json:"link,omitempty"`               //
	SignatureAlgorithm string                                      `json:"signatureAlgorithm,omitempty"` // Algorithm used for encrypting CSR
	Subject            string                                      `json:"subject,omitempty"`            // Subject of the certificate. Includes Common Name (CN), Organizational Unit (OU), etc.
	TimeStamp          string                                      `json:"timeStamp,omitempty"`          // Timestamp of the certificate generation.
	UsedFor            string                                      `json:"usedFor,omitempty"`            // Services for which the certificate is used for(for eg- MGMT, GENERIC).
}

type ResponseCertificatesGetCsrByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesDeleteCsrByID struct {
	Response *ResponseCertificatesDeleteCsrByIDResponse `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  //
}

type ResponseCertificatesDeleteCsrByIDResponse struct {
	Message string `json:"message,omitempty"` //
}

type ResponseCertificatesRegenerateIseRootCa struct {
	Response *ResponseCertificatesRegenerateIseRootCaResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  //
}

type ResponseCertificatesRegenerateIseRootCaResponse struct {
	ID      string                                               `json:"id,omitempty"`      // Id which can be used to track status of ISE root CA chain regeneration
	Link    *ResponseCertificatesRegenerateIseRootCaResponseLink `json:"link,omitempty"`    //
	Message string                                               `json:"message,omitempty"` //
}

type ResponseCertificatesRegenerateIseRootCaResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesRenewCertificates struct {
	Response *ResponseCertificatesRenewCertificatesResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  //
}

type ResponseCertificatesRenewCertificatesResponse struct {
	ID      string                                             `json:"id,omitempty"`      // Id which can be used to track status of certificate regeneration
	Link    *ResponseCertificatesRenewCertificatesResponseLink `json:"link,omitempty"`    //
	Message string                                             `json:"message,omitempty"` //
}

type ResponseCertificatesRenewCertificatesResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesBindCsr struct {
	Response *ResponseCertificatesBindCsrResponse `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  //
}

type ResponseCertificatesBindCsrResponse struct {
	Message string `json:"message,omitempty"` //
	Status  string `json:"status,omitempty"`  // Response status after import
}

type ResponseCertificatesImportSystemCertificate struct {
	Response *ResponseCertificatesImportSystemCertificateResponse `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  //
}

type ResponseCertificatesImportSystemCertificateResponse struct {
	ID      string `json:"id,omitempty"`      // ID of the imported trust certificate
	Message string `json:"message,omitempty"` // Response message on import of system or trust certificate
	Status  string `json:"status,omitempty"`  // HTTP response status after import
}

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
	ExpirationDate            string                                                 `json:"expirationDate,omitempty"`            // The time and date past which the certificate is no longer valid
	FriendlyName              string                                                 `json:"friendlyName,omitempty"`              // Friendly name of system certificate
	GroupTag                  string                                                 `json:"groupTag,omitempty"`                  //
	ID                        string                                                 `json:"id,omitempty"`                        // ID of system certificate
	IssuedBy                  string                                                 `json:"issuedBy,omitempty"`                  // Common Name of the certificate issuer
	IssuedTo                  string                                                 `json:"issuedTo,omitempty"`                  // Common Name of the certificate subject
	KeySize                   *int                                                   `json:"keySize,omitempty"`                   // The length of key used for encrypting system certificate
	Link                      *ResponseCertificatesGetSystemCertificatesResponseLink `json:"link,omitempty"`                      //
	PortalsUsingTheTag        string                                                 `json:"portalsUsingTheTag,omitempty"`        //
	SelfSigned                *bool                                                  `json:"selfSigned,omitempty"`                //
	SerialNumberDecimalFormat string                                                 `json:"serialNumberDecimalFormat,omitempty"` // Used to uniquely identify the certificate within a CA's systems
	Sha256Fingerprint         string                                                 `json:"sha256Fingerprint,omitempty"`         //
	SignatureAlgorithm        string                                                 `json:"signatureAlgorithm,omitempty"`        //
	UsedBy                    string                                                 `json:"usedBy,omitempty"`                    //
	ValidFrom                 string                                                 `json:"validFrom,omitempty"`                 // The time and date on which the certificate was created, also known as the Not Before certificate attribute
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
	ExpirationDate            string                                                    `json:"expirationDate,omitempty"`            // The time and date past which the certificate is no longer valid
	FriendlyName              string                                                    `json:"friendlyName,omitempty"`              // Friendly name of system certificate
	GroupTag                  string                                                    `json:"groupTag,omitempty"`                  //
	ID                        string                                                    `json:"id,omitempty"`                        // ID of system certificate
	IssuedBy                  string                                                    `json:"issuedBy,omitempty"`                  // Common Name of the certificate issuer
	IssuedTo                  string                                                    `json:"issuedTo,omitempty"`                  // Common Name of the certificate subject
	KeySize                   *int                                                      `json:"keySize,omitempty"`                   // The length of key used for encrypting system certificate
	Link                      *ResponseCertificatesGetSystemCertificateByIDResponseLink `json:"link,omitempty"`                      //
	PortalsUsingTheTag        string                                                    `json:"portalsUsingTheTag,omitempty"`        //
	SelfSigned                *bool                                                     `json:"selfSigned,omitempty"`                //
	SerialNumberDecimalFormat string                                                    `json:"serialNumberDecimalFormat,omitempty"` // Used to uniquely identify the certificate within a CA's systems
	Sha256Fingerprint         string                                                    `json:"sha256Fingerprint,omitempty"`         //
	SignatureAlgorithm        string                                                    `json:"signatureAlgorithm,omitempty"`        //
	UsedBy                    string                                                    `json:"usedBy,omitempty"`                    //
	ValidFrom                 string                                                    `json:"validFrom,omitempty"`                 // The time and date on which the certificate was created, also known as the Not Before certificate attribute
}

type ResponseCertificatesGetSystemCertificateByIDResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

type ResponseCertificatesUpdateSystemCertificate struct {
	Response *ResponseCertificatesUpdateSystemCertificateResponse `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  //
}

type ResponseCertificatesUpdateSystemCertificateResponse struct {
	ID      string                                                   `json:"id,omitempty"`      // ID of the imported trust certificate
	Link    *ResponseCertificatesUpdateSystemCertificateResponseLink `json:"link,omitempty"`    //
	Message string                                                   `json:"message,omitempty"` // Response message on import of system or trust certificate
	Status  string                                                   `json:"status,omitempty"`  // HTTP response status after import
}

type ResponseCertificatesUpdateSystemCertificateResponseLink struct {
	Href string `json:"href,omitempty"` //
	Rel  string `json:"rel,omitempty"`  //
	Type string `json:"type,omitempty"` //
}

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
	AuthenticateBeforeCRLReceived  string                                                  `json:"authenticateBeforeCRLReceived,omitempty"`  // Switch to enable/disable authentication before receiving CRL
	AutomaticCRLUpdate             string                                                  `json:"automaticCRLUpdate,omitempty"`             // Switch to enable/disable automatic CRL update
	AutomaticCRLUpdatePeriod       string                                                  `json:"automaticCRLUpdatePeriod,omitempty"`       // Automatic CRL update period
	AutomaticCRLUpdateUnits        string                                                  `json:"automaticCRLUpdateUnits,omitempty"`        // Unit of time of automatic CRL update
	CrlDistributionURL             string                                                  `json:"crlDistributionUrl,omitempty"`             // CRL Distribution URL
	CrlDownloadFailureRetries      string                                                  `json:"crlDownloadFailureRetries,omitempty"`      // If CRL download fails, wait time before retry
	CrlDownloadFailureRetriesUnits string                                                  `json:"crlDownloadFailureRetriesUnits,omitempty"` // Unit of time before retry if CRL download fails
	Description                    string                                                  `json:"description,omitempty"`                    // Description of trust certificate
	DownloadCRL                    string                                                  `json:"downloadCRL,omitempty"`                    // Switch to enable/disable download of CRL
	EnableOCSpValidation           string                                                  `json:"enableOCSPValidation,omitempty"`           // Switch to enable/disable OCSP Validation
	EnableServerIDentityCheck      string                                                  `json:"enableServerIdentityCheck,omitempty"`      // Switch to enable/disable Server Identity Check
	ExpirationDate                 string                                                  `json:"expirationDate,omitempty"`                 // The time and date past which the certificate is no longer valid
	FriendlyName                   string                                                  `json:"friendlyName,omitempty"`                   // Friendly name of trust certificate
	ID                             string                                                  `json:"id,omitempty"`                             // ID of trust certificate
	IgnoreCRLExpiration            string                                                  `json:"ignoreCRLExpiration,omitempty"`            // Switch to enable/disable ignore CRL Expiration
	InternalCa                     *bool                                                   `json:"internalCA,omitempty"`                     //
	IssuedBy                       string                                                  `json:"issuedBy,omitempty"`                       // The entity that verified the information and signed the certificate
	IssuedTo                       string                                                  `json:"issuedTo,omitempty"`                       // Entity to which trust certificate is issued
	KeySize                        string                                                  `json:"keySize,omitempty"`                        // The length of key used for encrypting trust certificate
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

type ResponseCertificatesImportTrustCertificate struct {
	Response *ResponseCertificatesImportTrustCertificateResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  //
}

type ResponseCertificatesImportTrustCertificateResponse struct {
	ID      string `json:"id,omitempty"`      // ID of the imported trust certificate
	Message string `json:"message,omitempty"` // Response message on import of system or trust certificate
	Status  string `json:"status,omitempty"`  // HTTP response status after import
}

type ResponseCertificatesGetTrustedCertificateByID struct {
	Response *ResponseCertificatesGetTrustedCertificateByIDResponse `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  //
}

type ResponseCertificatesGetTrustedCertificateByIDResponse struct {
	AuthenticateBeforeCRLReceived  string                                                     `json:"authenticateBeforeCRLReceived,omitempty"`  // Switch to enable/disable authentication before receiving CRL
	AutomaticCRLUpdate             string                                                     `json:"automaticCRLUpdate,omitempty"`             // Switch to enable/disable automatic CRL update
	AutomaticCRLUpdatePeriod       string                                                     `json:"automaticCRLUpdatePeriod,omitempty"`       // Automatic CRL update period
	AutomaticCRLUpdateUnits        string                                                     `json:"automaticCRLUpdateUnits,omitempty"`        // Unit of time of automatic CRL update
	CrlDistributionURL             string                                                     `json:"crlDistributionUrl,omitempty"`             // CRL Distribution URL
	CrlDownloadFailureRetries      string                                                     `json:"crlDownloadFailureRetries,omitempty"`      // If CRL download fails, wait time before retry
	CrlDownloadFailureRetriesUnits string                                                     `json:"crlDownloadFailureRetriesUnits,omitempty"` // Unit of time before retry if CRL download fails
	Description                    string                                                     `json:"description,omitempty"`                    // Description of trust certificate
	DownloadCRL                    string                                                     `json:"downloadCRL,omitempty"`                    // Switch to enable/disable download of CRL
	EnableOCSpValidation           string                                                     `json:"enableOCSPValidation,omitempty"`           // Switch to enable/disable OCSP Validation
	EnableServerIDentityCheck      string                                                     `json:"enableServerIdentityCheck,omitempty"`      // Switch to enable/disable Server Identity Check
	ExpirationDate                 string                                                     `json:"expirationDate,omitempty"`                 // The time and date past which the certificate is no longer valid
	FriendlyName                   string                                                     `json:"friendlyName,omitempty"`                   // Friendly name of trust certificate
	ID                             string                                                     `json:"id,omitempty"`                             // ID of trust certificate
	IgnoreCRLExpiration            string                                                     `json:"ignoreCRLExpiration,omitempty"`            // Switch to enable/disable ignore CRL Expiration
	InternalCa                     *bool                                                      `json:"internalCA,omitempty"`                     //
	IsReferredInPolicy             *bool                                                      `json:"isReferredInPolicy,omitempty"`             //
	IssuedBy                       string                                                     `json:"issuedBy,omitempty"`                       // The entity that verified the information and signed the certificate
	IssuedTo                       string                                                     `json:"issuedTo,omitempty"`                       // Entity to which trust certificate is issued
	KeySize                        string                                                     `json:"keySize,omitempty"`                        // The length of key used for encrypting trust certificate
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
	RemoveExistingIseIntermediateCsr *bool `json:"removeExistingISEIntermediateCSR,omitempty"` // Setting this attribute to true will remove existing ISE Intermediate CSR
}

type RequestCertificatesRenewCertificates struct {
	CertType string `json:"certType,omitempty"` //
}

type RequestCertificatesBindCsr struct {
	Admin                            *bool  `json:"admin,omitempty"`                            //  Use certificate to authenticate the ISE Admin Portal
	AllowExtendedValidity            *bool  `json:"allowExtendedValidity,omitempty"`            // Allow import of certificates with validity greater than 398 days (required)
	AllowOutOfDateCert               *bool  `json:"allowOutOfDateCert,omitempty"`               // Allow out of date certificates (required)
	AllowReplacementOfCertificates   *bool  `json:"allowReplacementOfCertificates,omitempty"`   // Allow Replacement of certificates (required)
	AllowReplacementOfPortalGroupTag *bool  `json:"allowReplacementOfPortalGroupTag,omitempty"` // Allow Replacement of Portal Group Tag (required)
	Data                             string `json:"data,omitempty"`                             // Signed Certificate data (required)
	Eap                              *bool  `json:"eap,omitempty"`                              // Use certificate for EAP protocols that use SSL/TLS tunneling
	HostName                         string `json:"hostName,omitempty"`                         // Name of Host whose CSR ID has been provided (required)
	ID                               string `json:"id,omitempty"`                               // ID of the generated CSR (required)
	Ims                              *bool  `json:"ims,omitempty"`                              // Use certificate for the ISE Messaging Service
	Name                             string `json:"name,omitempty"`                             // Friendly Name of the certificate
	Portal                           *bool  `json:"portal,omitempty"`                           // Use for portal
	PortalGroupTag                   string `json:"portalGroupTag,omitempty"`                   // Set Group tag
	Pxgrid                           *bool  `json:"pxgrid,omitempty"`                           // Use certificate for the pxGrid Controller
	Radius                           *bool  `json:"radius,omitempty"`                           // Use certificate for the RADSec server
	Saml                             *bool  `json:"saml,omitempty"`                             // Use certificate for SAML Signing
	ValidateCertificateExtensions    *bool  `json:"validateCertificateExtensions,omitempty"`    // Validate Certificate Extensions
}

type RequestCertificatesExportSystemCertificate struct {
	Export   string `json:"export,omitempty"`   //
	ID       string `json:"id,omitempty"`       //
	Password string `json:"password,omitempty"` //
}

type RequestCertificatesImportSystemCertificate struct {
	Admin                            *bool  `json:"admin,omitempty"`                            // Use certificate to authenticate the ISE Admin Portal
	AllowExtendedValidity            *bool  `json:"allowExtendedValidity,omitempty"`            // Allow import of certificates with validity greater than 398 days (required)
	AllowOutOfDateCert               *bool  `json:"allowOutOfDateCert,omitempty"`               // Allow out of date certificates (required)
	AllowReplacementOfCertificates   *bool  `json:"allowReplacementOfCertificates,omitempty"`   // Allow Replacement of certificates (required)
	AllowReplacementOfPortalGroupTag *bool  `json:"allowReplacementOfPortalGroupTag,omitempty"` // Allow Replacement of Portal Group Tag (required)
	AllowSHA1Certificates            *bool  `json:"allowSHA1Certificates,omitempty"`            // Allow SHA1 based certificates (required)
	AllowWildCardCertificates        *bool  `json:"allowWildCardCertificates,omitempty"`        // Allow Wildcard Certificates
	Data                             string `json:"data,omitempty"`                             // Certificate Content (required)
	Eap                              *bool  `json:"eap,omitempty"`                              // Use certificate for EAP protocols that use SSL/TLS tunneling
	Ims                              *bool  `json:"ims,omitempty"`                              // Use certificate for the ISE Messaging Service
	Name                             string `json:"name,omitempty"`                             // Name of the certificate
	Password                         string `json:"password,omitempty"`                         // Certificate Password (required).
	Portal                           *bool  `json:"portal,omitempty"`                           // Use for portal
	PortalGroupTag                   string `json:"portalGroupTag,omitempty"`                   // Set Group tag
	PortalTagTransferForSameSubject  *bool  `json:"portalTagTransferForSameSubject,omitempty"`  // Allow overwriting the portal tag from matching certificate of same subject
	PrivateKeyData                   string `json:"privateKeyData,omitempty"`                   // Private Key data (required)
	Pxgrid                           *bool  `json:"pxgrid,omitempty"`                           // Use certificate for the pxGrid Controller
	Radius                           *bool  `json:"radius,omitempty"`                           // Use certificate for the RADSec server
	RoleTransferForSameSubject       *bool  `json:"roleTransferForSameSubject,omitempty"`       // Allow transfer of roles for certificate with matching subject
	Saml                             *bool  `json:"saml,omitempty"`                             // Use certificate for SAML Signing
	ValidateCertificateExtensions    *bool  `json:"validateCertificateExtensions,omitempty"`    // Validate Certificate Extensions
}

type RequestCertificatesUpdateSystemCertificate struct {
	Admin                            *bool  `json:"admin,omitempty"`                            // Use certificate to authenticate the ISE Admin Portal
	AllowReplacementOfPortalGroupTag *bool  `json:"allowReplacementOfPortalGroupTag,omitempty"` // Allow Replacement of Portal Group Tag (required)
	Description                      string `json:"description,omitempty"`                      // Description of System Certificate
	Eap                              *bool  `json:"eap,omitempty"`                              // Use certificate for EAP protocols that use SSL/TLS tunneling
	ExpirationTTLPeriod              *int   `json:"expirationTTLPeriod,omitempty"`              //
	ExpirationTTLUnits               string `json:"expirationTTLUnits,omitempty"`               //
	Ims                              *bool  `json:"ims,omitempty"`                              // Use certificate for the ISE Messaging Service
	Name                             string `json:"name,omitempty"`                             // Name of the certificate
	Portal                           *bool  `json:"portal,omitempty"`                           // Use for portal
	PortalGroupTag                   string `json:"portalGroupTag,omitempty"`                   // Set Group tag
	PortalTagTransferForSameSubject  *bool  `json:"portalTagTransferForSameSubject,omitempty"`  // Allow overwriting the portal tag from matching certificate of same subject
	Pxgrid                           *bool  `json:"pxgrid,omitempty"`                           // Use certificate for the pxGrid Controller
	Radius                           *bool  `json:"radius,omitempty"`                           // Use certificate for the RADSec server
	RenewSelfSignedCertificate       *bool  `json:"renewSelfSignedCertificate,omitempty"`       // Renew Self Signed Certificate
	RoleTransferForSameSubject       *bool  `json:"roleTransferForSameSubject,omitempty"`       // Allow transfer of roles for certificate with matching subject
	Saml                             *bool  `json:"saml,omitempty"`                             // Use certificate for SAML Signing
}

type RequestCertificatesImportTrustCertificate struct {
	AllowBasicConstraintCaFalse       *bool  `json:"allowBasicConstraintCAFalse,omitempty"`       // Allow Certificates with Basic Constraints CA Field as False (required)
	AllowOutOfDateCert                *bool  `json:"allowOutOfDateCert,omitempty"`                // Allow out of date certificates (required)
	AllowSHA1Certificates             *bool  `json:"allowSHA1Certificates,omitempty"`             // Allow SHA1 based certificates (required)
	Data                              string `json:"data,omitempty"`                              // Certificate content (required)
	Description                       string `json:"description,omitempty"`                       // Description of the certificate
	Name                              string `json:"name,omitempty"`                              // Name of the certificate
	TrustForCertificateBasedAdminAuth *bool  `json:"trustForCertificateBasedAdminAuth,omitempty"` // Trust for Certificate based Admin authentication
	TrustForCiscoServicesAuth         *bool  `json:"trustForCiscoServicesAuth,omitempty"`         // Trust for authentication of Cisco Services
	TrustForClientAuth                *bool  `json:"trustForClientAuth,omitempty"`                // Trust for client authentication and Syslog
	TrustForIseAuth                   *bool  `json:"trustForIseAuth,omitempty"`                   // Trust for authentication within ISE
	ValidateCertificateExtensions     *bool  `json:"validateCertificateExtensions,omitempty"`     // Validate trust certificate extension
}

type RequestCertificatesUpdateTrustedCertificate struct {
	AuthenticateBeforeCRLReceived     *bool  `json:"authenticateBeforeCRLReceived,omitempty"`     // Switch to enable/disable CRL Verification if CRL is not Received
	AutomaticCRLUpdate                *bool  `json:"automaticCRLUpdate,omitempty"`                // Switch to enable/disable automatic CRL update
	AutomaticCRLUpdatePeriod          *int   `json:"automaticCRLUpdatePeriod,omitempty"`          // Automatic CRL update period
	AutomaticCRLUpdateUnits           string `json:"automaticCRLUpdateUnits,omitempty"`           // Unit of time for automatic CRL update
	CrlDistributionURL                string `json:"crlDistributionUrl,omitempty"`                // CRL Distribution URL
	CrlDownloadFailureRetries         *int   `json:"crlDownloadFailureRetries,omitempty"`         // If CRL download fails, wait time before retry
	CrlDownloadFailureRetriesUnits    string `json:"crlDownloadFailureRetriesUnits,omitempty"`    // Unit of time before retry if CRL download fails
	Description                       string `json:"description,omitempty"`                       // Description for trust certificate
	DownloadCRL                       *bool  `json:"downloadCRL,omitempty"`                       // Switch to enable/disable download of CRL
	EnableOCSpValidation              *bool  `json:"enableOCSPValidation,omitempty"`              // Switch to enable/disable OCSP Validation
	EnableServerIDentityCheck         *bool  `json:"enableServerIdentityCheck,omitempty"`         // Switch to enable/disable verification if HTTPS or LDAP server certificate name fits the configured server URL
	IgnoreCRLExpiration               *bool  `json:"ignoreCRLExpiration,omitempty"`               // Switch to enable/disable ignore CRL Expiration
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
	TrustForIseAuth                   *bool  `json:"trustForIseAuth,omitempty"`                   // Trust for authentication within ISE
}

//GetCsrs Purpose of the API is to get all Certificate Signing Requests from PAN.
/*  This API supports Filtering, Sorting and Pagination.

Filtering and Sorting supported on below mentioned attributes:


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

//ExportCsr Purpose of the API is to export a CSR, given a CSR id and a hostname.
/* The response of this API carries a CSR corresponding to the requested ID

@param hostname hostname path parameter. The hostname to which the CSR belongs.
@param id id path parameter. The ID of the CSR to be exported.
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

//GetCsrByID Purpose of the API is to get Certificate Signing Request(CSR) by ID
/* This API displays details of a Certificate Signing Request of a particular node based on a given HostName and ID.

@param hostName hostName path parameter. Name of the host of which CSR's should be returned
@param id id path parameter. The ID of the Certificate Signing Request returned
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

//GetSystemCertificates Purpose of this API is to get all system certificates of a particular node
/*  This API supports Filtering, Sorting and Pagination.

Filtering and Sorting supported on below mentioned attributes:


friendlyName

issuedTo

issuedBy

validFrom


Supported Date Format: yyyy-MM-dd HH:mm:ss

Supported Operators: EQ, NEQ, GT and LT


expirationDate


Supported Date Format: yyyy-MM-dd HH:mm:ss

Supported Operators: EQ, NEQ, GT and LT




@param hostName hostName path parameter. Name of the host of which system certificates should be returned
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

//GetSystemCertificateByID Purpose of this API is to get system certificate of a particular node by Id
/* This API displays details of a System Certificate of a particular node based on a given HostName and ID.

@param hostName hostName path parameter. Name of the host of which system certificates should be returned
@param id id path parameter. The id of the system certificate
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

//GetTrustedCertificates Purpose of this API is to get list of all trusted certificates
/*  This API supports Filtering, Sorting and Pagination.

Filtering and Sorting supported on below mentioned attributes:


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

//ExportTrustedCertificate Purpose of the API is to export a trust certificate given a certificate id.
/* The response of this API carries a trusted certificate file mapped to the requested id

@param id id path parameter. The ID of the Trusted Certificate to be exported.
*/
func (s *CertificatesService) ExportTrustedCertificate(id string) (FileDownload, *resty.Response, error) {
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

//GetTrustedCertificateByID Purpose of this API is to get Trust Certificate By Id
/* This API can displays details of a Trust Certificate based on a given ID.

@param id id path parameter. The id of the trust certificate
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

//GenerateCsr Purpose of this API is to generate a Certificate Signing Request (CSR).
/*

Generate a certificate signing request for Multi-Use, Admin, EAP Authentication, RADIUS DTLS, PxGrid, SAML, Portal and IMS Services.

Following Parameters are present in POST request body




PARAMETER

DESCRIPTION

EXAMPLE





hostnames

List of ise node hostnames for which CSRs should be generated

[ise-host1, ise-host2]



allowWildCardCert

Allow use of WildCards in certificates

false



keyLength

Length of the Key used for CSR generation (required)

512



keyType

Type of key used for CSR generation either RSA or ECDSA(required)

RSA



digestType

Hash algorithm used for signing CSR(required)

SHA-256



usedFor

Certificate Usage(required)

MULTI-USE



subjectCommonName

Certificate common name(CN)(required)

$FQDN$



subjectOrgUnit

Certificate organizational unit(OU)

Engineering



subjectOrg

Certificate organization (O)

Cisco



subjectCity

Certificate city or locality (L)

San Jose


subjectState

Certificate state (ST)

California


subjectCountry

Certificate country ( C)

US



sanDNS

Array of SAN(Subject Alternative Name) DNS entries(optional)

[ise.example.com]


sanIP

Array of SAN IP entries(optional)

[1.1.1.1]


sanURI

Array of SAN URI entries(optional)

[https://1.1.1.1]



sanDir

Array of SAN DIR entries(optional)

[CN=AAA,DC=COM,C=IL]



portalGroupTag

Portal Group Tag when using certificate for PORTAL service

Default Portal Certificate Group




NOTE:
For allowWildCardCert to be false, the below mentioned parameter is mandatory:
hostnames

When Certificate is selected to be used for Portal Service, the below mentioned parameter is mandatory:
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
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GenerateCsr")
	}

	result := response.Result().(*ResponseCertificatesGenerateCsr)
	return result, response, err

}

//GenerateIntermediateCaCsr Purpose of this API is to generate a Intermediate CA Certificate Signing Request (CSR).
/* CSR Generation for Intermediate Certificates.

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
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GenerateIntermediateCaCsr")
	}

	result := response.Result().(*ResponseCertificatesGenerateIntermediateCaCsr)
	return result, response, err

}

//RegenerateIseRootCa Purpose of this API is to regenerate the entire internal CA certificate chain including the root CA on the Primary PAN and subordinate CAs on the PSNs(Applicable only for the internal CA service ).
/* This API will initiate regeneration of ISE root CA certificate chain. Response contains id which can be used to track the status.
Setting "removeExistingISEIntermediateCSR" to true will remove existing ISE Intermediate CSR

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
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RegenerateIseRootCa")
	}

	result := response.Result().(*ResponseCertificatesRegenerateIseRootCa)
	return result, response, err

}

//RenewCertificates Purpose of this API is to renew certificates of OCSP responder and ISE Messaging Service.
/* This API will initiate regeneration of certificates. Response contains id which can be used to track the status

 */
func (s *CertificatesService) RenewCertificates(requestCertificatesRenewCertificates *RequestCertificatesRenewCertificates) (*ResponseCertificatesRenewCertificates, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/renew-certificate"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCertificatesRenewCertificates).
		SetResult(&ResponseCertificatesRenewCertificates{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RenewCertificates")
	}

	result := response.Result().(*ResponseCertificatesRenewCertificates)
	return result, response, err

}

//BindCsr Purpose of the API is to Bind CA Signed Certificate.
/*

Bind CA Signed Certificate.

NOTE:
This API requires an existing Certificate Signing Request, and the root certificate must already be trusted.

NOTE:
The certificate may have a validity period longer than 398 days. It may be untrusted by many browsers.

NOTE:
Request Parameters accepting True and False as input can be replaced by 1 and 0 respectively.

Following Parameters are used in POST body




PARAMETER

DESCRIPTION

EXAMPLE





name

Friendly name of the certificate.

Signed Certificate



data

Plain-text contents of the certificate file. Every space needs to be replaced with newline escape sequence (\n) (required).

awk 'NF {sub(/\r/, ""); printf "%s\\n",$0;}' <<your .pem file>>



allowExtendedValidity

Allow the certificates greater than validity of 398 days (required)

false



allowOutOfDateCert

Allow out of date certificates (required)

false



allowReplacementOfCertificates

Allow Replacement of certificates (required)

false



allowReplacementOfPortalGroupTag

Allow Replacement of Portal Group Tag (required)

false


admin

Use certificate to authenticate the ISE Admin Portal

false


eap

Use certificate for EAP protocols that use SSL/TLS tunneling

false



radius

Use certificate for RADSec server

false



pxgrid

Use certificate for the pxGrid Controller

false



ims

Use certificate for the ISE Messaging Service

false



saml

Use certificate for SAML Signing

false



portal

Use certificate for portal

false



portalGroupTag

Portal Group Tag for using certificate with portal role

Default Portal Certificate Group



validateCertificateExtensions

Validate Certificate Extensions

false





Following Roles can be used in any combinations




ROLE

DEFAULT

WARNING





Admin

False

Enabling Admin role for this certificate will cause an application server restart on the selected node.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates



EAP Authentication

False

Only one system certificate can be used for EAP. Assigning EAP to this certificate will remove the assignment from another certificate.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates



RADIUS DTLS

False

Only one system certificate can be used for DTLS. Assigning DTLS to this certificate will remove the assignment from another certificate.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates



SAML

False

SAML cannot be used with other Usage. Enabling SAML will uncheck all other Usage.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates





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
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BindCsr")
	}

	result := response.Result().(*ResponseCertificatesBindCsr)
	return result, response, err

}

//ExportSystemCertificate Purpose of the API is to export a system certificate given a certificate id.
/*

Export System Certificate.

Following Parameters are used in POST body




PARAMETER

DESCRIPTION

EXAMPLE





id

ID of a System Certificate(required)

< SYSTEM_CERT_ID >



export

 One of the below option is required

"CERTIFICATE" :
Export only Certificate without Private Key

"CERTIFICATE_WITH_PRIVATE_KEY" :
Export both Certificate and Private Key(
"certificatePassword"
 is required)



CERTIFICATE_WITH_PRIVATE_KEY



password

Certificate Password (required if
"export" : CERTIFICATE_WITH_PRIVATE_KEY
)
Password Constraints:


Alphanumeric

Minimum of 8 Characters

Maximum of 100 Characters



Passw***





NOTE:
The response of this API carries a ZIP file containing the certificate and private key if
"export" : "CERTIFICATE_WITH_PRIVATE_KEY"
 in the request. If
"export" : "CERTIFICATE"
 in request body, the response carries a ZIP file containing only the certificate.



WARNING:
Exporting a private key is not a secure operation. It could lead to possible exposure of the private key.



*/
func (s *CertificatesService) ExportSystemCertificate(requestCertificatesExportSystemCertificate *RequestCertificatesExportSystemCertificate) (FileDownload, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/system-certificate/export"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCertificatesExportSystemCertificate).
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

//ImportSystemCertificate Purpose of the API is to import system certificate into ISE.
/*

Import an X509 certificate as a system certificate.

NOTE:
The certificate may have a validity period longer than 398 days. It may be untrusted by many browsers.

NOTE:
Request Parameters accepting True and False as input can be replaced by 1 and 0 respectively.

Following Parameters are used in POST body




PARAMETER

DESCRIPTION

EXAMPLE





name

Friendly name of the certificate.

System Certificate



password

Password of the certificate to be imported (required).

Passw***



data

Plain-text contents of the certificate file. Every space needs to be replaced with newline escape sequence (\n) (required).

awk 'NF {sub(/\r/, ""); printf "%s\\n",$0;}' <<your .pem file>>



privateKeyData

Plain-text contents of the private key file. Every space needs to be replaced with newline escape sequence (\n) (required).

awk 'NF {sub(/\r/, ""); printf "%s\\n",$0;}' <<your .pem file>>



allowOutOfDateCert

Allow out of date certificates (required)

false



allowSHA1Certificates

Allow SHA1 based certificates (required)

false



allowExtendedValidity

Allow the certificates greater than validity of 398 days (required)

false



roleTransferForSameSubject

Allow the transfer of roles to certificates with same subject

false



portalTagTransferForSameSubject

Acquire the group tag of the matching certificate

false



admin

Use certificate to authenticate the ISE Admin Portal

false



eap

Use certificate for EAP protocols that use SSL/TLS tunneling

false



radius

Use certificate for RADSec server

false



pxgrid

Use certificate for the pxGrid Controller

false



ims

Use certificate for the ISE Messaging Service

false



saml

Use certificate for SAML Signing

false



portal

Use certificate for portal

false



portalGroupTag

Portal Group Tag for using certificate with portal role

Default Portal Certificate Group



allowReplacementOfPortalGroupTag

Allow Replacement of Portal Group Tag (required)

false



allowWildCardCertificates

Allow use of WildCards in certificates

false



validateCertificateExtensions

Validate Certificate Extensions

false





Following Roles can be used in any combinations




ROLE

DEFAULT

WARNING





Admin

False

Enabling Admin role for this certificate will cause an application server restart on the selected node.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates



EAP Authentication

False

Only one system certificate can be used for EAP. Assigning EAP to this certificate will remove the assignment from another certificate.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates



RADIUS DTLS

False

Only one system certificate can be used for DTLS. Assigning DTLS to this certificate will remove the assignment from another certificate.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates



SAML

False

SAML cannot be used with other Usage. Enabling SAML will uncheck all other Usage.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates





*/
func (s *CertificatesService) ImportSystemCertificate(requestCertificatesImportSystemCertificate *RequestCertificatesImportSystemCertificate) (*ResponseCertificatesImportSystemCertificate, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/system-certificate/import"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCertificatesImportSystemCertificate).
		SetResult(&ResponseCertificatesImportSystemCertificate{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ImportSystemCertificate")
	}

	result := response.Result().(*ResponseCertificatesImportSystemCertificate)
	return result, response, err

}

//ImportTrustCertificate Purpose of the API is to add root certificate to the ISE truststore.
/*

Import an X509 certificate as a trust certificate.

NOTE:
Request Parameters accepting True and False as input can be replaced by 1 and 0 respectively.

Following Parameters are used in POST body




PARAMETER

DESCRIPTION

EXAMPLE





name

Friendly name of the certificate

Trust Certificate



description

Description of the certificate

Passw***



data

Plain-text contents of the certificate file. Every space needs to be replaced with newline escape sequence (\n) (required).

awk 'NF {sub(/\r/, ""); printf "%s\\n",$0;}' <<your .pem file>>



allowOutOfDateCert

Allow out of date certificates (required)

false



allowSHA1Certificates

Allow SHA1 based certificates (required)

false



trustForIseAuth

Trust for authentication within ISE

false



trustForClientAuth

Trust for client authentication and Syslog

false



trustForCertificateBasedAdminAuth

Trust for Certificate based Admin authentication

false



trustForCiscoServicesAuth

Trust for authentication of Cisco Services

false



validateCertificateExtensions

Validate extensions for trust certificate

false




NOTE:
If name is not set, a default name of the following format will be generated:
common-name#issuer#nnnnn

    where
"nnnnn"
 is a unique number. You can always change the friendly name later by editing the certificate.


    You must choose how this certificate will be trusted in ISE. The objective here is to distinguish between certificates that are used for trust within an ISE deployment and public certificates that are used to trust Cisco services. Typically, you will not want to use a given certificate for both purposes.




Trusted For

Usage





Authentication within ISE

Use
"trustForIseAuth":true
 if the certificate is used for trust within ISE, such as for secure communication between ISE nodes



Client authentication and Syslog

Use
"trustForClientAuth":true
 if the certificate is to be used for authentication of endpoints that contact ISE over the EAP protocol. Also check this box if certificate is used to trust a Syslog server. Make sure to have keyCertSign bit asserted under KeyUsage extension for this certificate.
Note:
 "" can be set true only if the "trustForIseAuth" has been set true.



Certificate based admin authentication

Use
"trustForCertificateBasedAdminAuth":true
 if the certificate is used for trust within ISE, such as for secure communication between ISE nodes
Note:
 "trustForCertificateBasedAdminAuth" can be set true only if "trustForIseAuth" and "trustForClientAuth" are true.



Authentication of Cisco Services

 Use
"trustForCiscoServicesAuth":true
 if the certificate is to be used for trusting external Cisco services, such as Feed Service.





*/
func (s *CertificatesService) ImportTrustCertificate(requestCertificatesImportTrustCertificate *RequestCertificatesImportTrustCertificate) (*ResponseCertificatesImportTrustCertificate, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/trusted-certificate/import"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCertificatesImportTrustCertificate).
		SetResult(&ResponseCertificatesImportTrustCertificate{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ImportTrustCertificate")
	}

	result := response.Result().(*ResponseCertificatesImportTrustCertificate)
	return result, response, err

}

//UpdateSystemCertificate Purpose of the API is to update data for existing system certificate.
/*

Update a System Certificate.

NOTE:
Renewing a certificate will cause an application server restart on the selected node.

NOTE:
Request Parameters accepting True and False as input can be replaced by 1 and 0 respectively.

Following Parameters are used in POST body




PARAMETER

DESCRIPTION

EXAMPLE





name

Friendly name of the certificate.

System Certificate



description
Description of the Certificate
Default Description


admin

Use certificate to authenticate the ISE Admin Portal

false



eap

Use certificate for EAP protocols that use SSL/TLS tunneling

false



radius

Use certificate for RADSec server

false



pxgrid

Use certificate for the pxGrid Controller

false



ims

Use certificate for the ISE Messaging Service

false



saml

Use certificate for SAML Signing

false



portal

Use certificate for portal

false



portalGroupTag

Portal Group Tag for using certificate with portal role

Default Portal Certificate Group



allowReplacementOfPortalGroupTag

Allow Replacement of Portal Group Tag (required)

false



roleTransferForSameSubject

Allow the transfer of roles to certificates with same subject

false



portalTagTransferForSameSubject

Acquire the group tag of the matching certificate

false



renewSelfSignedCertificate

Renew Self Signed Certificate

false



expirationTTLPeriod

Expiration Period

365



expirationTTLUnits

Expiration Units in one of the below formats

days / weeks / months / years



days





Following Roles can be used in any combinations




ROLE

DEFAULT

WARNING





Admin

False

Enabling Admin role for this certificate will cause an application server restart on the selected node.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates



EAP Authentication

False

Only one system certificate can be used for EAP. Assigning EAP to this certificate will remove the assignment from another certificate.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates



RADIUS DTLS

False

Only one system certificate can be used for DTLS. Assigning DTLS to this certificate will remove the assignment from another certificate.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates



SAML

False

SAML cannot be used with other Usage. Enabling SAML will uncheck all other Usage.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates





@param id id path parameter. The ID of the System Certificate to be updated
@param hostName hostName path parameter. Name of Host whose certificate needs to be updated
*/
func (s *CertificatesService) UpdateSystemCertificate(id string, hostName string, requestCertificatesUpdateSystemCertificate *RequestCertificatesUpdateSystemCertificate) (*ResponseCertificatesUpdateSystemCertificate, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/system-certificate/{hostName}/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{hostName}", fmt.Sprintf("%v", hostName), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCertificatesUpdateSystemCertificate).
		SetResult(&ResponseCertificatesUpdateSystemCertificate{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSystemCertificate")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificatesUpdateSystemCertificate)
	return result, response, err

}

//UpdateTrustedCertificate Purpose of the API is to update trust certificate already present in ISE trust store.
/*

Update a trusted certificate present in ISE trust store.

Following Parameters are used in PUT request body




PARAMETER

DESCRIPTION

EXAMPLE





name

Friendly name of the certificate(required)

Trust Certificate



status

Status of the certificate

Enabled



description

Description of the certificate

Certificate for secure connection to cisco.com



trustForIseAuth

Trust for authentication within ISE

false



trustForClientAuth

Trust for client authentication and Syslog

false



trustForCertificateBasedAdminAuth

Trust for Certificate based Admin authentication

false



trustForCiscoServicesAuth

Trust for authentication of Cisco Services

false



enableOCSPValidation

Switch to enable/disable OCSP Validation

false



selectedOCSPService

Name of selected OCSP Service

INTERNAL_OCSP_SERVICE



rejectIfNoStatusFromOCSP

Switch to reject certificate if there is no status from OCSP

false



rejectIfUnreachableFromOCSP

Switch to reject certificate if unreachable from OCSP

false



downloadCRL

Switch to enable/disable download of CRL

false



crlDistributionUrl

Certificate Revocation List Distribution URL




automaticCRLUpdate

Switch to enable/disable automatic CRL update

false



automaticCRLUpdatePeriod

Automatic CRL update period

5



automaticCRLUpdateUnits

Unit of time for automatic CRL update

Minutes



nonAutomaticCRLUpdatePeriod

Non automatic CRL update period

1



nonAutomaticCRLUpdateUnits

Unit of time of non automatic CRL update

Hours



crlDownloadFailureRetries

If CRL download fails, wait time before retry

10



crlDownloadFailureRetriesUnits

Unit of time before retry if CRL download fails

Minutes



enableServerIdentityCheck

Switch to enable/disable verification if HTTPS or LDAP server certificate name fits the configured server URL

false



authenticateBeforeCRLReceived

Switch to enable/disable CRL Verification if CRL is not Received

false



ignoreCRLExpiration

Switch to enable/disable ignore CRL Expiration

false








Trusted For

Usage





Authentication within ISE

Use
"trustForIseAuth":true
 if the certificate is used for trust within ISE, such as for secure communication between ISE nodes



Client authentication and Syslog

Use
"trustForClientAuth":true
 if the certificate is to be used for authentication of endpoints that contact ISE over the EAP protocol. Also check this box if certificate is used to trust a Syslog server. Make sure to have keyCertSign bit asserted under KeyUsage extension for this certificate.
Note:
 "trustForClientAuth" can be set true only if the "trustForIseAuth" has been set true.



Certificate based admin authentication

Use
"trustForCertificateBasedAdminAuth":true
 if the certificate is used for trust within ISE, such as for secure communication between ISE nodes
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



OCSP Service name

Use
"selectedOCSPService":"Name of OCSP Service"
 Name of the OCSP service against which the certificate should be validated
Note:
 "selectedOCSPService" value will on be used if "enableOCSPValidation" has been set true.



Reject the request if OCSP returns UNKNOWN status

Use
"rejectIfNoStatusFromOCSP":true
 to reject the certificate if the OCSP service returns UNKNOWN status
Note:
 "rejectIfNoStatusFromOCSP:true" can be used only if "enableOCSPValidation" has been set true.



Reject the request if OCSP Responder is unreachable

 Use
"rejectIfUnreachableFromOCSP":true
 to reject the certificate if the OCSP service is unreachable.
Note:
 "rejectIfUnreachableFromOCSP:true" can be used only if "enableOCSPValidation" has been set true.







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
 "crlDistributionUrl" value will only be used if "downloadCRL" has been set true.



Retrieve CRL time

Use
"automaticCRLUpdate":true and automaticCRLUpdatePeriod, automaticCRLUpdatePeriod
 to set the time before which CRL is automatically retrieved prior to expiration Use
"nonAutomaticCRLUpdatePeriod, nonAutomaticCRLUpdateUnits
 to set the time period for CRL retrieval in loop.

Note:
 All the above fields can be used only if "downloadCRL" has been set true.



If download fails

Use
"crlDownloadFailureRetries" and "crlDownloadFailureRetriesUnits"
 to set retry time period if CRL download fails
Note:
 "crlDownloadFailureRetries" and "crlDownloadFailureRetriesUnits" can be used only if "downloadCRL" has been set true.



Enable Server Identity Check

Use
"enableServerIdentityCheck":true
 to verify that HTTPS or LDAPS server certificate name fits the configured server URL
Note:
 "enableServerIdentityCheck:true" can be used only if "downloadCRL" has been set true.



Bypass CRL Verification if CRL is not Received

Use
"authenticateBeforeCRLReceived":true
 to bypass CRL Verification if CRL is not Received
Note:
 "authenticateBeforeCRLReceived:true" can be used only if "downloadCRL" has been set true.



Ignore that CRL is not yet valid or has expired

 Use
"ignoreCRLExpiration":true
 to ignore if CRL is not yet valid or expired
Note:
 "ignoreCRLExpiration:true" can be used only if "downloadCRL" has been set true.





Note:
boolean properties accept integers values as well, with 0 considered as false and other values being considered as true


@param id id path parameter. The id of the trust certificate
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
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateTrustedCertificate")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificatesUpdateTrustedCertificate)
	return result, response, err

}

//DeleteCsrByID Purpose of the API is to delete Certificate Signing Request(CSR) by ID
/* This API deletes a Certificate Signing Request of a particular node based on a given HostName and ID.

@param hostName hostName path parameter. Name of the host of which CSR's should be deleted
@param id id path parameter. The ID of the Certificate Signing Request to be deleted
*/
func (s *CertificatesService) DeleteCsrByID(hostName string, id string) (*ResponseCertificatesDeleteCsrByID, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/certs/certificate-signing-request/{hostName}/{id}"
	path = strings.Replace(path, "{hostName}", fmt.Sprintf("%v", hostName), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCertificatesDeleteCsrByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteCsrById")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseCertificatesDeleteCsrByID)
	return result, response, err

}

//DeleteSystemCertificateByID Purpose of the API is to delete System Certificate by ID and hostname
/* This API deletes a System Certificate of a particular node based on a given HostName and ID.

@param hostName hostName path parameter. Name of the host from which the System Certificate needs to be deleted
@param id id path parameter. The ID of the System Certificate to be deleted
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

//DeleteTrustedCertificateByID Purpose of the API is to delete Trusted Certificate by ID
/* This API deletes a Trust Certificate from Trusted Certificate Store based on a given ID.

@param id id path parameter. The ID of the Trusted Certificate to be deleted
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
