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

	params := &isegosdk.GetSystemCertificatesQueryParams{}

	fmt.Println("executing GetSystemCertificates...")
	result, _, err := Client.Certificates.GetSystemCertificates("ise", params)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := &isegosdk.RequestCertificatesExportSystemCertificate{
		Export: "CERTIFICATE",
	}

	for _, row := range result.Response {
		fmt.Println("---------------")
		data.ID = row.ID
		fmt.Println("executing ExportSystemCertificate... ", row.ID)
		cert, _, err := Client.Certificates.ExportSystemCertificate(data)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path := "/home/carguello/Desktop/" //path to save the certificate e.g., "/home/user/foo". Setting empty "" will save the certificate in the current directory.
		err = cert.SaveDownload(path)
		if err != nil {
			fmt.Println(err)
		}
	}
}
