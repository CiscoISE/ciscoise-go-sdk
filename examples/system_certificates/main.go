package main

import (
	"fmt"
	"os"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"
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

	params := &isegosdk.GetSystemCertificatesQueryParams{}

	fmt.Println("executing GetSystemCertificates...")
	result, _, err := Client.Certificates.GetSystemCertificates("ise", params)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := &isegosdk.RequestCertificatesExportSystemCert{
		Export: "CERTIFICATE",
	}

	if result != nil && result.Response != nil {
		for _, row := range *result.Response {
			fmt.Println("---------------")
			data.ID = row.ID
			fmt.Println("executing ExportSystemCertificate... ", row.ID)
			cert, _, err := Client.Certificates.ExportSystemCert(data)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			path := "/tmp" //path to save the certificate e.g., "/home/user/foo". Setting empty "" will save the certificate in the current directory.
			err = cert.SaveDownload(path)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
