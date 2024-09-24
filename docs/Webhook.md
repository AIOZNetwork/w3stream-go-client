# \Webhook

All URIs are relative to https://api.w3stream.xyz/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](Webhook.md#Create) | **Post** /webhooks | Create webhook
[**Get**](Webhook.md#Get) | **Get** /webhooks/{id} | Get user&#39;s webhook by id
[**Update**](Webhook.md#Update) | **Patch** /webhooks/{id} | Update event webhook
[**Delete**](Webhook.md#Delete) | **Delete** /webhooks/{id} | Delete webhook
[**List**](Webhook.md#List) | **Get** /webhooks | Get list webhooks
[**Check**](Webhook.md#Check) | **Post** /webhooks/check/{id} | Check webhook by id



## Create

> Create(request CreateWebhookRequest) (*CreateWebhookResponse, error)

> CreateWithContext(ctx context.Context, request CreateWebhookRequest) (*CreateWebhookResponse, error)


Create webhook



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
        
    request := *w3streamsdk.NewCreateWebhookRequest() // CreateWebhookRequest | Create Webhook input

    
    res, err := client.Webhook.Create(request)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhook.Create``: %v\n", err)
    }
    // response from `Create`: CreateWebhookResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Webhook.Create`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**CreateWebhookRequest**](CreateWebhookRequest.md) | Create Webhook input | 

### Return type

[**CreateWebhookResponse**](CreateWebhookResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Get(id string) (*GetUserWebhookResponse, error)

> GetWithContext(ctx context.Context, id string) (*GetUserWebhookResponse, error)


Get user's webhook by id



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
        
    id := "id_example" // string | webhook's id

    
    res, err := client.Webhook.Get(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhook.Get``: %v\n", err)
    }
    // response from `Get`: GetUserWebhookResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Webhook.Get`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | webhook&#39;s id | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetUserWebhookResponse**](GetUserWebhookResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(id string, request UpdateWebhookRequest) (*ResponseSuccess, error)

> UpdateWithContext(ctx context.Context, id string, request UpdateWebhookRequest) (*ResponseSuccess, error)


Update event webhook



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
        
    id := "id_example" // string | webhook's id
    request := *w3streamsdk.NewUpdateWebhookRequest() // UpdateWebhookRequest | Update Webhook input, events example: video.encoding.quality.completed

    
    res, err := client.Webhook.Update(id, request)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhook.Update``: %v\n", err)
    }
    // response from `Update`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Webhook.Update`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | webhook&#39;s id | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**UpdateWebhookRequest**](UpdateWebhookRequest.md) | Update Webhook input, events example: video.encoding.quality.completed | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(id string) (*ResponseSuccess, error)

> DeleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete webhook



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
        
    id := "id_example" // string | Webhook ID

    
    res, err := client.Webhook.Delete(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhook.Delete``: %v\n", err)
    }
    // response from `Delete`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Webhook.Delete`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Webhook ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(r WebhookApiListRequest) (*GetWebhooksListResponse, error)


> ListWithContext(ctx context.Context, r WebhookApiListRequest) (*GetWebhooksListResponse, error)



Get list webhooks



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
    req := w3streamsdk.WebhookApiListRequest{}
    
    req.Search("search_example") // string | only support search by name
    req.SortBy("sortBy_example") // string | sort by (default to "created_at")
    req.OrderBy("orderBy_example") // string | allowed: asc, desc. Default: asc (default to "asc")
    req.Offset(int32(56)) // int32 | offset, allowed values greater than or equal to 0. Default(0) (default to 0)
    req.Limit(int32(56)) // int32 | results per page. Allowed values 1-100, default is 25 (default to 25)
    req.EncodingFinished(true) // bool | search by event encoding finished
    req.EncodingStarted(true) // bool | search by event encoding started
    req.FileReceived(true) // bool | search by event file received

    res, err := client.Webhook.List(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhook.List``: %v\n", err)
    }
    // response from `List`: GetWebhooksListResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Webhook.List`")
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
**encodingFinished** | **bool** | search by event encoding finished | 
**encodingStarted** | **bool** | search by event encoding started | 
**fileReceived** | **bool** | search by event file received | 

### Return type

[**GetWebhooksListResponse**](GetWebhooksListResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Check

> Check(id string) (*ResponseSuccess, error)

> CheckWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Check webhook by id



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
        
    id := "id_example" // string | webhook's id

    
    res, err := client.Webhook.Check(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhook.Check``: %v\n", err)
    }
    // response from `Check`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Webhook.Check`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | webhook&#39;s id | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

