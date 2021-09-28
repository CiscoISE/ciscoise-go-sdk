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

	currentPage := 1
	params := &isegosdk.GetIDentityGroupsQueryParams{
		Size: 2,
	}

	for {
		fmt.Printf("\nexecuting GetIDentityGroups page[%d]...", currentPage)
		params.Page = currentPage
		result, _, err := Client.IDentityGroups.GetIDentityGroups(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if result != nil && result.SearchResult != nil && result.SearchResult.Resources != nil {
			for ix, row := range *result.SearchResult.Resources {
				fmt.Printf("\n\telement [%d], page [%d]", ix+1, currentPage)
				fmt.Println("\tID: ", row.ID)
				fmt.Println("\tName: ", row.Name)
				fmt.Println("\tDescription: ", row.Description)
				fmt.Println("\t********************")
			}
		}

		if result == nil || result.SearchResult == nil || result.SearchResult.NextPage == nil || result.SearchResult.NextPage.Rel != "next" {
			break
		}
		currentPage++
		fmt.Println("--------------------")
	}

}
