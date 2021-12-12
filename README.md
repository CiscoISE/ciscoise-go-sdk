# ciscoise-go-sdk
ciscoise-go-sdk is a go client library for [Cisco Identity Services Engine (ISE) ](https://developer.cisco.com/identity-services-engine/)  

## Introduction
The ciscoise-go-sdk makes it easier to work with the Cisco Identity Services Engine (ISE) RESTFul APIs from Go.

It supports version 3.1.1, but it is backward compatible with other versions as long as those versions use the same URLs and options as version 3.1.1.

## Getting started

The first thing you need to do is to generate an API client. There are two options to do it:

    Parameters
    Environment variables

### Parameters

The client could be generated with the following parameters:

- `baseURL`: The base URL, FQDN or IP, of the ISE instance.
- `username`: The username for the API authentication and authorization.
- `password`: The password for the API authentication and authorization.
- `debug`: Boolean to enable debugging
- `sslVerify`: Boolean to enable or disable SSL certificate verification.
- `useAPIGateway`: Boolean to enable or disable API Gateway usage.
- `useCSRFToken`: Boolean to enable or disable CSRF token.

```go
    Client, err := isegosdk.NewClientWithOptions("https://198.18.133.27",
        "admin", "C1sco12345",
        "false", "false",
        "false", "false")

```

### Using environment variables

The client can be configured with the following environment variables:

- `ISE_BASE_URL`: The base URL, FQDN or IP, of the ISE instance.
- `ISE_USERNAME`: The username for the API authentication and authorization.
- `ISE_PASSWORD`: The password for the API authentication and authorization.
- `ISE_DEBUG`: Boolean to enable debugging
- `ISE_SSL_VERIFY`: Boolean to enable or disable SSL certificate verification.
- `ISE_USE_API_GATEWAY`: Boolean to enable or disable API Gateway usage.
- `ISE_USE_CSRF_TOKEN`: Boolean to enable or disable CSRF token.

```go
Client, err = isegosdk.NewClient()
if err != nil {
    fmt.Println(err)
    os.Exit(1)
}
devicesCount, _, err := Client.Devices.GetDeviceCount()
```





## Examples
The following section show how to create a new client, create a new ANC policy, list all policies and delete the policy that was created. 

```go
// New client definition
Client, err := isegosdk.NewClientWithOptions("https://198.18.133.27",
	"admin", "C1sco12345",
	"false", "false",
	"false", "false")

if err != nil {
	fmt.Println(err)
	os.Exit(1)
}
policyName := "policy1"

policyContent := &isegosdk.RequestAncPolicyCreateAncPolicyErsAncPolicy{
	Name:    policyName,
	Actions: []string{"QUARANTINE"},
}

policy := &isegosdk.RequestAncPolicyCreateAncPolicy{
	ErsAncPolicy: policyContent,
}

// New policy creation
_, err = Client.AncPolicy.CreateAncPolicy(policy)
if err != nil {
	fmt.Println(err)
}

params := &isegosdk.GetAncPolicyQueryParams{
	Size: 10,
}

// Searchs for all policies
pols, _, err := Client.AncPolicy.GetAncPolicy(params)
if err != nil {
	fmt.Println(err)
}
if pols != nil && pols.SearchResult != nil && pols.SearchResult.Resources != nil {
	for _, pol := range *pols.SearchResult.Resources {
		fmt.Printf("Policy ID: %s \n", pol.ID)
		fmt.Printf("Policy Name: %s \n", pol.Name)
	}
}

// Delete policy by ID
_, err = Client.AncPolicy.DeleteAncPolicyByID(policyName)
if err != nil {
	fmt.Println(err)
}
```

## Documentation
https://pkg.go.dev/github.com/CiscoISE/ciscoise-go-sdk

## Compatibility matrix

| SDK versions | Cisco ISE version supported |
|--------------|-----------------------------|
| 0.y.z        |  3.1.0                      |
| 1.0.z        |  3.1.0                      |
| 1.1.z        |  3.1.1                      |

## Changelog

All notable changes to this project will be documented in the [CHANGELOG](https://github.com/CiscoISE/ciscoise-go-sdk/blob/main/CHANGELOG.md) file.

The development team may make additional name changes as the library evolves with the Cisco DNA Center APIs.

## License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE) file.
