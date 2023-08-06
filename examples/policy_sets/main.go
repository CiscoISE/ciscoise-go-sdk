package main

import (
	"fmt"
	"os"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"
)

func main() {
	Client, err := isegosdk.NewClientWithOptions("https://10.48.190.181",
		"admin", "Cisco12345",
		"false", "false",
		"false", "false")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// GET NETWORK ACCESS POLICY SETS
	fmt.Println("executing GetNetworkAccessPolicySets...")
	policy_sets, _, err := Client.NetworkAccessPolicySet.GetNetworkAccessPolicySets()

	if err != nil {
		fmt.Println(err)
	}

	if policy_sets != nil && policy_sets.Response != nil {
		for _, row := range *policy_sets.Response {
			fmt.Println("-------------")
			fmt.Println("ID: ", row.ID)
			fmt.Println("Name: ", row.Name)
			fmt.Println("Description: ", row.Description)
			fmt.Println("Condition: ", row.Condition)
			if row.Condition != nil {
				fmt.Println("Children: ", row.Condition.Children)
			}
		}
	}
	fmt.Println("-------------")

	// GET NETWORK ACCESS POLICY SETS BY ID
	fmt.Println("executing GetNetworkAccessPolicySetsbyID...")
	policy_set, _, err := Client.NetworkAccessPolicySet.GetNetworkAccessPolicySetByID("13727e31-93d1-4fc6-9a7e-80f6998b1916")

	if err != nil {
		fmt.Println(err)
	}
	if policy_set != nil && policy_set.Response != nil {

		fmt.Println("-------------")
		fmt.Println("ID: ", policy_set.Response.ID)
		fmt.Println("Name: ", policy_set.Response.Name)
		fmt.Println("Description: ", policy_set.Response.Description)
		fmt.Println("Condition: ", policy_set.Response.Condition)
		if policy_set.Response.Condition != nil {
			fmt.Println("Children: ", policy_set.Response.Condition.Children)
		}
	}
	fmt.Println("-------------")

	falseVal := false

	policy_set1 := &isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySet{
		Name:        "Test_9983",
		Description: "Sample policy_set using ciscoise-go-sdk",
		ServiceName: "Default Network Access",
		State:       "enabled",
		IsProxy:     &falseVal,
		Condition: &isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetCondition{
			ConditionType: "ConditionAndBlock",
			IsNegate:      &falseVal,
			Children: &[]isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetCondition{
				{
					ConditionType: "ConditionReference",
					ID:            "49ec9c17-6987-43ef-9478-a42111c59b81",
				},
				{
					ConditionType: "ConditionReference",
					ID:            "1e9f02d0-13e7-4a3b-8fbc-fb38330d74b5",
				},
				{
					ConditionType: "ConditionAndBlock",
					IsNegate:      &falseVal,
					Children: &[]isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetCondition{
						{
							ConditionType: "ConditionReference",
							ID:            "49ec9c17-6987-43ef-9478-a42111c59b81",
						},
						{
							ConditionType: "ConditionReference",
							ID:            "1e9f02d0-13e7-4a3b-8fbc-fb38330d74b5",
						},
						{
							ConditionType: "ConditionAndBlock",
							IsNegate:      &falseVal,
							Children: &[]isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetCondition{
								{
									ConditionType: "ConditionReference",
									ID:            "49ec9c17-6987-43ef-9478-a42111c59b81",
								},
								{
									ConditionType: "ConditionReference",
									ID:            "1e9f02d0-13e7-4a3b-8fbc-fb38330d74b5",
								},
							},
						},
					},
				},
			},
		},
	}

	policy_set2 := &isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByID{
		Name:        "Test_02",
		Description: "Sample policy_set using ciscoise-go-sdk",
		ServiceName: "Default Network Access",
		State:       "enabled",
		IsProxy:     &falseVal,
		Condition: &isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDCondition{
			ConditionType: "ConditionAndBlock",
			IsNegate:      &falseVal,
			Children: &[]isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDCondition{
				{
					ConditionType: "ConditionReference",
					ID:            "49ec9c17-6987-43ef-9478-a42111c59b81",
				},
				{
					ConditionType: "ConditionReference",
					ID:            "1e9f02d0-13e7-4a3b-8fbc-fb38330d74b5",
				},
				{
					ConditionType: "ConditionAndBlock",
					IsNegate:      &falseVal,
					Children: &[]isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDCondition{
						{
							ConditionType: "ConditionReference",
							ID:            "49ec9c17-6987-43ef-9478-a42111c59b81",
						},
						{
							ConditionType: "ConditionReference",
							ID:            "1e9f02d0-13e7-4a3b-8fbc-fb38330d74b5",
						},
						{
							ConditionType: "ConditionAndBlock",
							IsNegate:      &falseVal,
							Children: &[]isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDCondition{
								{
									ConditionType: "ConditionReference",
									ID:            "49ec9c17-6987-43ef-9478-a42111c59b81",
								},
								{
									ConditionType: "ConditionReference",
									ID:            "1e9f02d0-13e7-4a3b-8fbc-fb38330d74b5",
								},
							},
						},
					},
				},
			},
		},
	}

	//CREATE NETWORK ACCESS POLICY SET
	fmt.Println("executing CreatePolicySet...")
	policy_set_create, _, err := Client.NetworkAccessPolicySet.CreateNetworkAccessPolicySet(policy_set1)
	if err != nil {
		fmt.Println(err)
	}

	if policy_set_create != nil {
		fmt.Println(policy_set_create.Response)
	}

	//UPDATE NEWORK ACCESS POLICY SET BY ID
	fmt.Println("executing UpdateNetworkAccessPolicySetByID...")
	policy_setByID, _, err := Client.NetworkAccessPolicySet.UpdateNetworkAccessPolicySetByID("575b7017-e53d-4c66-a20c-d6794822321e", policy_set2)
	if err != nil {
		fmt.Println(err)
	}

	if policy_setByID != nil {
		fmt.Println(policy_setByID.Response)
	}
}
