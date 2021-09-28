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

	params := &isegosdk.GetAuthorizationProfilesQueryParams{
		Size: 2,
	}

	fmt.Println("executing GetAuthorizationProfiles...")
	profiles, _, err := Client.AuthorizationProfile.GetAuthorizationProfiles(params)

	if err != nil {
		fmt.Println(err)
	}

	if profiles != nil && profiles.SearchResult != nil && profiles.SearchResult.Resources != nil {
		for _, row := range *profiles.SearchResult.Resources {
			fmt.Println("-------------")
			fmt.Println("ID: ", row.ID)
			fmt.Println("Name: ", row.Name)
			fmt.Println("Description: ", row.Description)
		}
	}
	fmt.Println("-------------")
	falseVal := false

	prof := &isegosdk.RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfile{
		Name:                      "Internet_Only",
		Description:               "Sample authorization profile using ciscoise-go-sdk",
		AccessType:                "ACCESS_ACCEPT",
		AuthzProfileType:          "SWITCH",
		TrackMovement:             &falseVal,
		AgentlessPosture:          &falseVal,
		ServiceTemplate:           &falseVal,
		EasywiredSessionCandidate: &falseVal,
		DaclName:                  "Internet_Only",
		ProfileName:               "Cisco",
	}

	reqProf := &isegosdk.RequestAuthorizationProfileCreateAuthorizationProfile{
		AuthorizationProfile: prof,
	}

	fmt.Println("executing CreateAuthorizationProfile...")
	_, err = Client.AuthorizationProfile.CreateAuthorizationProfile(reqProf)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("executing GetAuthorizationProfileByName...")
	profByID, _, err := Client.AuthorizationProfile.GetAuthorizationProfileByName("Internet_Only")
	if err != nil {
		fmt.Println(err)
	}
	if profByID != nil && profByID.AuthorizationProfile != nil {
		newProfileID := profByID.AuthorizationProfile.ID
		fmt.Println("-------------")
		fmt.Println("ID: ", profByID.AuthorizationProfile.ID)
		fmt.Println("Name: ", profByID.AuthorizationProfile.Name)
		fmt.Println("Description: ", profByID.AuthorizationProfile.Description)
		fmt.Println("AdvancedAttributes: ", profByID.AuthorizationProfile.AdvancedAttributes)
		fmt.Println("AccessType: ", profByID.AuthorizationProfile.AccessType)
		fmt.Println("AuthzProfileType: ", profByID.AuthorizationProfile.AuthzProfileType)
		fmt.Println("VLAN: ", profByID.AuthorizationProfile.VLAN)
		fmt.Println("Reauth: ", profByID.AuthorizationProfile.Reauth)
		fmt.Println("WebRedirection: ", profByID.AuthorizationProfile.WebRedirection)
		fmt.Println("TrackMovement: ", profByID.AuthorizationProfile.TrackMovement)
		fmt.Println("AgentlessPosture: ", profByID.AuthorizationProfile.AgentlessPosture)
		fmt.Println("ServiceTemplate: ", profByID.AuthorizationProfile.ServiceTemplate)
		fmt.Println("EasywiredSessionCandidate: ", profByID.AuthorizationProfile.EasywiredSessionCandidate)
		fmt.Println("DaclName: ", profByID.AuthorizationProfile.DaclName)
		fmt.Println("VoiceDomainPermission: ", profByID.AuthorizationProfile.VoiceDomainPermission)
		fmt.Println("Neat: ", profByID.AuthorizationProfile.Neat)
		fmt.Println("WebAuth: ", profByID.AuthorizationProfile.WebAuth)
		fmt.Println("ProfileName: ", profByID.AuthorizationProfile.ProfileName)
		fmt.Println("Link: ", profByID.AuthorizationProfile.Link)
		fmt.Println("-------------")

		fmt.Println("executing DeleteAuthorizationProfileByID...")
		_, err = Client.AuthorizationProfile.DeleteAuthorizationProfileByID(newProfileID)

		if err != nil {
			fmt.Println(err)
		}
	}
}
