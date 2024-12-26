# \ApiKey

All URIs are relative to https://api.w3stream.xyz/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ApiKey.md#Create) | **Post** /api_keys | Create API key
[**Update**](ApiKey.md#Update) | **Patch** /api_keys/{id} | Rename API key
[**Delete**](ApiKey.md#Delete) | **Delete** /api_keys/{id} | Delete API key
[**List**](ApiKey.md#List) | **Get** /api_keys | Get list API keys



## Create

> Create(request CreateApiKeyRequest) (*CreateApiKeyResponse, error)

> CreateWithContext(ctx context.Context, request CreateApiKeyRequest) (*CreateApiKeyResponse, error)


Create API key



### Example

```go
package main

import (
    "context"
    "fmt"
    "encoding/json"
    "os"
    w3streamsdk "github.com/AIOZNetwork/w3stream-go-client"
)

func main() {
    // create a new client
    apiCreds := w3streamsdk.AuthCredentials{
		SecretKey: "YOUR_SECRET_KEY",
		PublicKey: "YOUR_PUBLIC_KEY",
    }
    client := w3streamsdk.ClientBuilder(apiCreds).Build()
        
    request := *w3streamsdk.NewCreateApiKeyRequest() // CreateApiKeyRequest | api key's data

    
    res, err := client.ApiKey.Create(request)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKey.Create``: %v\n", err)
    }
    // response from `Create`: CreateApiKeyResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `ApiKey.Create`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**CreateApiKeyRequest**](CreateApiKeyRequest.md) | api key&#39;s data | 

### Return type

[**CreateApiKeyResponse**](CreateApiKeyResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(id string, request RenameAPIKeyRequest) (*ResponseSuccess, error)

> UpdateWithContext(ctx context.Context, id string, request RenameAPIKeyRequest) (*ResponseSuccess, error)


Rename API key



### Example

```go
package main

import (
    "context"
    "fmt"
    "encoding/json"
    "os"
    w3streamsdk "github.com/AIOZNetwork/w3stream-go-client"
)

func main() {
    // create a new client
    apiCreds := w3streamsdk.AuthCredentials{
		SecretKey: "YOUR_SECRET_KEY",
		PublicKey: "YOUR_PUBLIC_KEY",
    }
    client := w3streamsdk.ClientBuilder(apiCreds).Build()
        
    id := "id_example" // string | api key id
    request := *w3streamsdk.NewRenameAPIKeyRequest() // RenameAPIKeyRequest | new api key name

    
    res, err := client.ApiKey.Update(id, request)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKey.Update``: %v\n", err)
    }
    // response from `Update`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `ApiKey.Update`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | api key id | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**RenameAPIKeyRequest**](RenameAPIKeyRequest.md) | new api key name | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(id string) (*ResponseSuccess, error)

> DeleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete API key



### Example

```go
package main

import (
    "context"
    "fmt"
    "encoding/json"
    "os"
    w3streamsdk "github.com/AIOZNetwork/w3stream-go-client"
)

func main() {
    // create a new client
    apiCreds := w3streamsdk.AuthCredentials{
		SecretKey: "YOUR_SECRET_KEY",
		PublicKey: "YOUR_PUBLIC_KEY",
    }
    client := w3streamsdk.ClientBuilder(apiCreds).Build()
        
    id := "id_example" // string | API key's ID

    
    res, err := client.ApiKey.Delete(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKey.Delete``: %v\n", err)
    }
    // response from `Delete`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `ApiKey.Delete`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | API key&#39;s ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(r ApiKeyApiListRequest) (*GetApiKeysResponse, error)


> ListWithContext(ctx context.Context, r ApiKeyApiListRequest) (*GetApiKeysResponse, error)



Get list API keys



### Example

```go
package main

import (
    "context"
    "fmt"
    "encoding/json"
    "os"
    w3streamsdk "github.com/AIOZNetwork/w3stream-go-client"
)

func main() {
    // create a new client
    apiCreds := w3streamsdk.AuthCredentials{
		SecretKey: "YOUR_SECRET_KEY",
		PublicKey: "YOUR_PUBLIC_KEY",
    }
    client := w3streamsdk.ClientBuilder(apiCreds).Build()
    req := w3streamsdk.ApiKeyApiListRequest{}
    
    req.Search("search_example") // string | only support search by name
    req.SortBy("sortBy_example") // string | sort by (default to "created_at")
    req.OrderBy("orderBy_example") // string | allowed: asc, desc. Default: asc (default to "asc")
    req.Offset(int32(56)) // int32 | offset, allowed values greater than or equal to 0. Default(0) (default to 0)
    req.Limit(int32(56)) // int32 | results per page. Allowed values 1-100, default is 25 (default to 25)

    res, err := client.ApiKey.List(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKey.List``: %v\n", err)
    }
    // response from `List`: GetApiKeysResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `ApiKey.List`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**search** | **string** | only support search by name | 
**sortBy** | **string** | sort by | [default to &quot;created_at&quot;]
**orderBy** | **string** | allowed: asc, desc. Default: asc | [default to &quot;asc&quot;]
**offset** | **int32** | offset, allowed values greater than or equal to 0. Default(0) | [default to 0]
**limit** | **int32** | results per page. Allowed values 1-100, default is 25 | [default to 25]

### Return type

[**GetApiKeysResponse**](GetApiKeysResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

