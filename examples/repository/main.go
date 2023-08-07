package main

import (
	"fmt"
	"os"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"
)

func main() {
	// New client definition
	Client, err := isegosdk.NewClientWithOptions("https://10.48.35.230",
		"admin", "Devnet.12345",
		"false", "false",
		"false", "false")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	trueval := true
	fmt.Println("executing CreateRepository...")
	rep := &isegosdk.RequestRepositoryCreateRepository{
		Name:       "RepoName",
		Protocol:   "SFTP",
		ServerName: "demo",
		Path:       "/pathdemo",
		EnablePki:  &trueval,
		UserName:   "foo",
		Password:   "123456",
	}

	_, _, err = Client.Repository.CreateRepository(rep)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("executing repository  GetRepositories...")
	reGeAll, _, err := Client.Repository.GetRepositories()
	if err != nil {
		fmt.Println(err)
	}

	if reGeAll != nil && reGeAll.Response != nil {
		fmt.Println("Repository list...")
		fmt.Println("Version: ", reGeAll.Version)
		for _, row := range *reGeAll.Response {
			fmt.Println("-------------")
			fmt.Println("Name: ", row.Name)
			fmt.Println("Protocol: ", row.Protocol)
			fmt.Println("Path: ", row.Path)
			fmt.Println("ServerName: ", row.ServerName)
			fmt.Println("UserName: ", row.UserName)
			fmt.Println("EnablePki: ", row.EnablePki)
			fmt.Println("-------------")
		}
	}

	fmt.Println("executing  UpdateRepository...")
	repUp := &isegosdk.RequestRepositoryUpdateRepository{
		Name:       "RepoName",
		Protocol:   "SFTP",
		ServerName: "demo2",
		Path:       "/pathdemo2",
		EnablePki:  &trueval,
		UserName:   "newFoo",
		Password:   "654321",
	}

	_, _, err = Client.Repository.UpdateRepository("RepoName", repUp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("executing GetRepository...")
	repOne, _, err := Client.Repository.GetRepository("RepoName")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(repOne)

	fmt.Println("executing repository  DeleteRepository...")
	_, _, err = Client.Repository.DeleteRepository("RepoName")
	if err != nil {
		fmt.Println(err)
	}

}
