# \Watermark

All URIs are relative to https://api.w3stream.xyz/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**Upload**](Watermark.md#Upload) | **Post** /watermarks | Create a new watermark
[**Delete**](Watermark.md#Delete) | **Delete** /watermarks/{id} | Delete a watermark by ID
[**List**](Watermark.md#List) | **Get** /watermarks | List all watermarks



## Upload

> UploadFile(file *os.File) (*CreateWatermarkResponse, error)
> Upload(fileName string, fileReader io.Reader)
> UploadFileWithContext(ctx context.Context, file *os.File) (*CreateWatermarkResponse, error)
> UploadWithContext(ctx context.Context, fileName string, fileReader io.Reader)

Create a new watermark



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
        
    file := os.NewFile(1234, "some_file") // *os.File | Watermark image file

    
    res, err := client.Watermark.UploadFile(file)

    // you can also use a Reader instead of a File:
    // we recommend using Reader instead!
    // client.Watermark.Upload(fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Watermark.Upload``: %v\n", err)
    }
    // response from `Upload`: CreateWatermarkResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Watermark.Upload`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | Watermark image file | 

### Return type

[**CreateWatermarkResponse**](CreateWatermarkResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(id string) (*GetWebhooksListResponse, error)

> DeleteWithContext(ctx context.Context, id string) (*GetWebhooksListResponse, error)


Delete a watermark by ID



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
        
    id := "id_example" // string | Watermark ID

    
    res, err := client.Watermark.Delete(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Watermark.Delete``: %v\n", err)
    }
    // response from `Delete`: GetWebhooksListResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Watermark.Delete`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Watermark ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetWebhooksListResponse**](GetWebhooksListResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(r WatermarkApiListRequest) (*GetAllWatermarkResponse, error)


> ListWithContext(ctx context.Context, r WatermarkApiListRequest) (*GetAllWatermarkResponse, error)



List all watermarks



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
    req := w3streamsdk.WatermarkApiListRequest{}
    
    req.SortBy("sortBy_example") // string | sort by (default to "created_at")
    req.OrderBy("orderBy_example") // string | allowed: asc, desc. Default: asc (default to "asc")
    req.Offset(int32(56)) // int32 | offset, allowed values greater than or equal to 0. Default(0) (default to 0)
    req.Limit(int32(56)) // int32 | results per page. Allowed values 1-100, default is 25 (default to 25)

    res, err := client.Watermark.List(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Watermark.List``: %v\n", err)
    }
    // response from `List`: GetAllWatermarkResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Watermark.List`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sortBy** | **string** | sort by | [default to &quot;created_at&quot;]
**orderBy** | **string** | allowed: asc, desc. Default: asc | [default to &quot;asc&quot;]
**offset** | **int32** | offset, allowed values greater than or equal to 0. Default(0) | [default to 0]
**limit** | **int32** | results per page. Allowed values 1-100, default is 25 | [default to 25]

### Return type

[**GetAllWatermarkResponse**](GetAllWatermarkResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

