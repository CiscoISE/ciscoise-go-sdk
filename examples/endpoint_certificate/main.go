package main

import (
	"fmt"
	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"os"
)

func main() {
	Client, err := isegosdk.NewClientWithOptions("https://198.18.133.27",
		"admin", "C1sco12345",
		"false", "false",
		"false", "false")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cr := &isegosdk.RequestEndpointCertificateCreateEndpointCertificateERSEndPointCertCertificateRequest{
		San: "11-22-33-44-55-66",
		Cn:  "userName",
	}
	ctemp := &isegosdk.RequestEndpointCertificateCreateEndpointCertificateERSEndPointCert{
		CertTemplateName:   "CA_SERVICE_Certificate_Template",
		Format:             "PKCS8",
		Password:           "Password_123",
		CertificateRequest: cr,
	}

	endpointReqBody := &isegosdk.RequestEndpointCertificateCreateEndpointCertificate{
		ERSEndPointCert: ctemp,
	}

	fmt.Println("executing CreateEndpointCertificate...")
	file, _, err := Client.EndpointCertificate.CreateEndpointCertificate(endpointReqBody)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	path := "/tmp" //path to save the certificate e.g., "/home/user/foo". Setting empty "" will save the certificate in the current directory.
	err = file.SaveDownload(path)
	if err != nil {
		fmt.Println(err)
	}

}
