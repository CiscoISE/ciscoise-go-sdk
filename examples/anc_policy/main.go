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
	policyName := "policy1"

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	policyContent := &isegosdk.RequestAncPolicyCreateAncPolicyErsAncPolicy{
		Name:    policyName,
		Actions: []string{"QUARANTINE"},
	}

	policy := &isegosdk.RequestAncPolicyCreateAncPolicy{
		ErsAncPolicy: policyContent,
	}

	// New policy creation
	fmt.Println("executing CreateAncPolicy...")
	_, err = Client.AncPolicy.CreateAncPolicy(policy)
	if err != nil {
		fmt.Println(err)
	}

	// Searchs for all policies
	params := &isegosdk.GetAncPolicyQueryParams{
		Size: 10,
	}

	fmt.Println("executing GetAncPolicy...")
	pols, _, err := Client.AncPolicy.GetAncPolicy(params)
	if err != nil {
		fmt.Println(err)
	}

	for _, pol := range pols.SearchResult.Resources {
		fmt.Println("-------------")
		fmt.Printf("Policy ID: %s \n", pol.ID)
		fmt.Printf("Policy Name: %s \n", pol.Name)
		fmt.Println("-------------")
	}

	// Update policy by ID
	polData := &isegosdk.RequestAncPolicyUpdateAncPolicyByIDErsAncPolicy{
		ID:      policyName,
		Name:    policyName,
		Actions: []string{"PORTBOUNCE"},
	}
	polRequest := &isegosdk.RequestAncPolicyUpdateAncPolicyByID{
		ErsAncPolicy: polData,
	}

	fmt.Println("executing UpdateAncPolicyByID...")
	updateResult, _, err := Client.AncPolicy.UpdateAncPolicyByID(policyName, polRequest)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("changes")
	for _, row := range updateResult.UpdatedFieldsList.UpdatedField {
		fmt.Println("-------------")
		fmt.Println("Field:", row.Field)
		fmt.Println("OldValue:", row.OldValue)
		fmt.Println("NewValue:", row.NewValue)
		fmt.Println("-------------")
	}

	// Delete policy by ID
	fmt.Println("executing DeleteAncPolicyByID...")
	_, err = Client.AncPolicy.DeleteAncPolicyByID(policyName)
	if err != nil {
		fmt.Println(err)
	}
}
