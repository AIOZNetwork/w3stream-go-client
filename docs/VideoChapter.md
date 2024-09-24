# \VideoChapter

All URIs are relative to https://api.w3stream.xyz/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](VideoChapter.md#Create) | **Post** /videos/{id}/chapters/{lan} | Create a video chapter
[**Get**](VideoChapter.md#Get) | **Get** /videos/{id}/chapters | Get video chapters
[**Delete**](VideoChapter.md#Delete) | **Delete** /videos/{id}/chapters/{lan} | Delete a video chapter



## Create

> CreateFile(id string, lan string, file *os.File) (*CreateVideoChapterResponse, error)
> Create(id string, lan string, fileName string, fileReader io.Reader)
> CreateFileWithContext(ctx context.Context, id string, lan string, file *os.File) (*CreateVideoChapterResponse, error)
> CreateWithContext(ctx context.Context, id string, lan string, fileName string, fileReader io.Reader)

Create a video chapter



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
        
    id := "id_example" // string | Video ID
    lan := "lan_example" // string | Language
    file := os.NewFile(1234, "some_file") // *os.File | VTT File

    
    res, err := client.VideoChapter.CreateFile(id, lan, file)

    // you can also use a Reader instead of a File:
    // we recommend using Reader instead!
    // client.VideoChapter.Create(id, lan, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoChapter.Create``: %v\n", err)
    }
    // response from `Create`: CreateVideoChapterResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `VideoChapter.Create`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Video ID | 
**lan** | **string** | Language | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | VTT File | 

### Return type

[**CreateVideoChapterResponse**](CreateVideoChapterResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Get(id string, r VideoChapterApiGetRequest) (*GetVideoChaptersResponse, error)


> GetWithContext(ctx context.Context, id string, r VideoChapterApiGetRequest) (*GetVideoChaptersResponse, error)



Get video chapters



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
    req := w3streamsdk.VideoChapterApiGetRequest{}
    
    req.Id("id_example") // string | Video ID
    req.Offset(int32(56)) // int32 | offset, allowed values greater than or equal to 0. Default(0) (default to 0)
    req.Limit(int32(56)) // int32 | results per page. Allowed values 1-100, default is 25 (default to 25)

    res, err := client.VideoChapter.Get(id string, req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoChapter.Get``: %v\n", err)
    }
    // response from `Get`: GetVideoChaptersResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `VideoChapter.Get`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Video ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**offset** | **int32** | offset, allowed values greater than or equal to 0. Default(0) | [default to 0]
**limit** | **int32** | results per page. Allowed values 1-100, default is 25 | [default to 25]

### Return type

[**GetVideoChaptersResponse**](GetVideoChaptersResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(id string, lan string) (*ResponseSuccess, error)

> DeleteWithContext(ctx context.Context, id string, lan string) (*ResponseSuccess, error)


Delete a video chapter



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
        
    id := "id_example" // string | Video ID
    lan := "lan_example" // string | Language

    
    res, err := client.VideoChapter.Delete(id, lan)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoChapter.Delete``: %v\n", err)
    }
    // response from `Delete`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `VideoChapter.Delete`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Video ID | 
**lan** | **string** | Language | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

