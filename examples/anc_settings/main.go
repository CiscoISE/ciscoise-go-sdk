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

	fmt.Println("executing GetAciSettings...")
	ACIbstt, _, err := Client.AciSettings.GetAciSettings()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("-----------------------")
	fmt.Println("ID: ", ACIbstt.AciSettings.ID)
	fmt.Println("EnableAci: ", ACIbstt.AciSettings.EnableAci)
	fmt.Println("IPAddressHostName: ", ACIbstt.AciSettings.IPAddressHostName)
	fmt.Println("AdminName: ", ACIbstt.AciSettings.AdminName)
	fmt.Println("AdminPassword: ", ACIbstt.AciSettings.AdminPassword)
	fmt.Println("Aciipaddress: ", ACIbstt.AciSettings.Aciipaddress)
	fmt.Println("AciuserName: ", ACIbstt.AciSettings.AciuserName)
	fmt.Println("Acipassword: ", ACIbstt.AciSettings.Acipassword)
	fmt.Println("TenantName: ", ACIbstt.AciSettings.TenantName)
	fmt.Println("L3RouteNetwork: ", ACIbstt.AciSettings.L3RouteNetwork)
	fmt.Println("SuffixToEpg: ", ACIbstt.AciSettings.SuffixToEpg)
	fmt.Println("SuffixToSgt: ", ACIbstt.AciSettings.SuffixToSgt)
	fmt.Println("AllSxpDomain: ", ACIbstt.AciSettings.AllSxpDomain)
	fmt.Println("SpecificSxpDomain: ", ACIbstt.AciSettings.SpecificSxpDomain)
	fmt.Println("SpecifixSxpDomainList: ", ACIbstt.AciSettings.SpecifixSxpDomainList)
	fmt.Println("EnableDataPlane: ", ACIbstt.AciSettings.EnableDataPlane)
	fmt.Println("UntaggedPacketIepgName: ", ACIbstt.AciSettings.UntaggedPacketIepgName)
	fmt.Println("DefaultSgtName: ", ACIbstt.AciSettings.DefaultSgtName)
	fmt.Println("EnableElementsLimit: ", ACIbstt.AciSettings.EnableElementsLimit)
	fmt.Println("MaxNumIepgFromAci: ", ACIbstt.AciSettings.MaxNumIepgFromAci)
	fmt.Println("Aci50: ", ACIbstt.AciSettings.Aci50)
	fmt.Println("Aci51: ", ACIbstt.AciSettings.Aci51)
}
