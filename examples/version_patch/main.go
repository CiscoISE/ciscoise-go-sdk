package main

import (
	"fmt"
	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"os"
)

func main() {
	// New client definition
	Client, err := isegosdk.NewClientWithOptions("https://198.18.133.27",
		"admin", "C1sco12345",
		"false", "false",
		"false", "false")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Executing GetIseVersionAndPatch ...")
	version, _, err := Client.VersionAndPatch.GetIseVersionAndPatch()

	if err != nil {
		fmt.Println(err)
	}

	for _, row := range version.OperationResult.ResultValue {
		fmt.Println("-------------")
		fmt.Println("Name: ", row.Name)
		fmt.Println("Value: ", row.Value)
		fmt.Println("-------------")
	}
}
