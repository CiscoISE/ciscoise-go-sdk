package main

import (
	"fmt"
	"os"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"
)

func main() {
	Client, err := isegosdk.NewClientWithOptions("https://198.18.133.27",
		"admin", "907113",
		"true", "false",
		"false", "false")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cr := &isegosdk.RequestBackupAndRestoreConfigBackup{
		BackupEncryptionKey: "My3ncryptionkey",
		BackupName:          "myBackup",
		RepositoryName:      "myRepo",
	}
	fmt.Println("executing ConfigBackup...")
	st, _, err := Client.BackupAndRestore.ConfigBackup(cr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("St: %+v\n", st)
}
