# \Video

All URIs are relative to https://api.w3stream.xyz/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](Video.md#Create) | **Post** /videos/create | Create video object
[**Update**](Video.md#Update) | **Patch** /videos/{id} | update video info
[**Delete**](Video.md#Delete) | **Delete** /videos/{id} | Delete video
[**UploadThumbnail**](Video.md#UploadThumbnail) | **Post** /videos/{id}/thumbnail | Upload video thumbnail
[**CreateCaption**](Video.md#CreateCaption) | **Post** /videos/{id}/caption/{lan} | Create a new video caption
[**DeleteCaption**](Video.md#DeleteCaption) | **Delete** /videos/{id}/caption/{lan} | Delete a video caption
[**GetCaptions**](Video.md#GetCaptions) | **Get** /videos/{id}/captions | Get video captions
[**GetCost**](Video.md#GetCost) | **Get** /videos/cost | get video transcoding cost
[**GetDetail**](Video.md#GetDetail) | **Get** /videos/{id} | get video detail
[**GetVideoList**](Video.md#GetVideoList) | **Post** /videos | Get user videos list
[**SetCaption**](Video.md#SetCaption) | **Patch** /videos/{id}/caption/{lan} | Set default video caption
[**UploadPart**](Video.md#UploadPart) | **Post** /videos/{id}/part | Upload part of video
[**UploadVideoComplete**](Video.md#UploadVideoComplete) | **Get** /videos/{id}/complete | Get upload video when complete



## Create

> Create(request CreateVideoRequest) (*CreateVideoResponse, error)

> CreateWithContext(ctx context.Context, request CreateVideoRequest) (*CreateVideoResponse, error)


Create video object



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
        
    request := *w3streamsdk.NewCreateVideoRequest() // CreateVideoRequest | video's info

    
    res, err := client.Video.Create(request)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Video.Create``: %v\n", err)
    }
    // response from `Create`: CreateVideoResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Video.Create`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**CreateVideoRequest**](CreateVideoRequest.md) | video&#39;s info | 

### Return type

[**CreateVideoResponse**](CreateVideoResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(id string, input UpdateVideoInfoRequest) (*ResponseSuccess, error)

> UpdateWithContext(ctx context.Context, id string, input UpdateVideoInfoRequest) (*ResponseSuccess, error)


update video info

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
        
    id := "id_example" // string | video's id
    input := *w3streamsdk.NewUpdateVideoInfoRequest() // UpdateVideoInfoRequest | input

    
    res, err := client.Video.Update(id, input)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Video.Update``: %v\n", err)
    }
    // response from `Update`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Video.Update`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | video&#39;s id | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**input** | [**UpdateVideoInfoRequest**](UpdateVideoInfoRequest.md) | input | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(id string) (*ResponseSuccess, error)

> DeleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete video



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

    
    res, err := client.Video.Delete(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Video.Delete``: %v\n", err)
    }
    // response from `Delete`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Video.Delete`")
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

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadThumbnail

> UploadThumbnailFile(id string, file *os.File) (*ResponseSuccess, error)
> UploadThumbnail(id string, fileName string, fileReader io.Reader)
> UploadThumbnailFileWithContext(ctx context.Context, id string, file *os.File) (*ResponseSuccess, error)
> UploadThumbnailWithContext(ctx context.Context, id string, fileName string, fileReader io.Reader)

Upload video thumbnail

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
        
    id := "id_example" // string | video's id
    file := os.NewFile(1234, "some_file") // *os.File | file video to be uploaded

    
    res, err := client.Video.UploadThumbnailFile(id, file)

    // you can also use a Reader instead of a File:
    // we recommend using Reader instead!
    // client.Video.UploadThumbnail(id, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Video.UploadThumbnail``: %v\n", err)
    }
    // response from `UploadThumbnail`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Video.UploadThumbnail`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | video&#39;s id | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | file video to be uploaded | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCaption

> CreateCaptionFile(id string, lan string, file *os.File) (*CreateVideoCaptionResponse, error)
> CreateCaption(id string, lan string, fileName string, fileReader io.Reader)
> CreateCaptionFileWithContext(ctx context.Context, id string, lan string, file *os.File) (*CreateVideoCaptionResponse, error)
> CreateCaptionWithContext(ctx context.Context, id string, lan string, fileName string, fileReader io.Reader)

Create a new video caption



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

    
    res, err := client.Video.CreateCaptionFile(id, lan, file)

    // you can also use a Reader instead of a File:
    // we recommend using Reader instead!
    // client.Video.CreateCaption(id, lan, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Video.CreateCaption``: %v\n", err)
    }
    // response from `CreateCaption`: CreateVideoCaptionResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Video.CreateCaption`")
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

[**CreateVideoCaptionResponse**](CreateVideoCaptionResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCaption

> DeleteCaption(id string, lan string) (*ResponseSuccess, error)

> DeleteCaptionWithContext(ctx context.Context, id string, lan string) (*ResponseSuccess, error)


Delete a video caption



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

    
    res, err := client.Video.DeleteCaption(id, lan)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Video.DeleteCaption``: %v\n", err)
    }
    // response from `DeleteCaption`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Video.DeleteCaption`")
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


## GetCaptions

> GetCaptions(id string, r VideoApiGetCaptionsRequest) (*GetVideoCaptionsResponse, error)


> GetCaptionsWithContext(ctx context.Context, id string, r VideoApiGetCaptionsRequest) (*GetVideoCaptionsResponse, error)



Get video captions



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
    req := w3streamsdk.VideoApiGetCaptionsRequest{}
    
    req.Id("id_example") // string | Video ID
    req.Offset(int32(56)) // int32 | offset, allowed values greater than or equal to 0. Default(0) (default to 0)
    req.Limit(int32(56)) // int32 | results per page. Allowed values 1-100, default is 25 (default to 25)

    res, err := client.Video.GetCaptions(id string, req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Video.GetCaptions``: %v\n", err)
    }
    // response from `GetCaptions`: GetVideoCaptionsResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Video.GetCaptions`")
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

[**GetVideoCaptionsResponse**](GetVideoCaptionsResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCost

> GetCost(qualities string, duration float32) (*GetTranscodeCostResponse, error)

> GetCostWithContext(ctx context.Context, qualities string, duration float32) (*GetTranscodeCostResponse, error)


get video transcoding cost



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
        
    qualities := "qualities_example" // string | video's qualities
    duration := float32(8.14) // float32 | video's duration

    
    res, err := client.Video.GetCost(qualities, duration)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Video.GetCost``: %v\n", err)
    }
    // response from `GetCost`: GetTranscodeCostResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Video.GetCost`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**qualities** | **string** | video&#39;s qualities | 
**duration** | **float32** | video&#39;s duration | 

### Return type

[**GetTranscodeCostResponse**](GetTranscodeCostResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDetail

> GetDetail(id string) (*GetVideoDetailResponse, error)

> GetDetailWithContext(ctx context.Context, id string) (*GetVideoDetailResponse, error)


get video detail



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
        
    id := "id_example" // string | video's id

    
    res, err := client.Video.GetDetail(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Video.GetDetail``: %v\n", err)
    }
    // response from `GetDetail`: GetVideoDetailResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Video.GetDetail`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | video&#39;s id | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetVideoDetailResponse**](GetVideoDetailResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideoList

> GetVideoList(request GetVideoListRequest) (*GetVideoListResponse, error)

> GetVideoListWithContext(ctx context.Context, request GetVideoListRequest) (*GetVideoListResponse, error)


Get user videos list



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
        
    request := *w3streamsdk.NewGetVideoListRequest() // GetVideoListRequest | video's info

    
    res, err := client.Video.GetVideoList(request)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Video.GetVideoList``: %v\n", err)
    }
    // response from `GetVideoList`: GetVideoListResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Video.GetVideoList`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**GetVideoListRequest**](GetVideoListRequest.md) | video&#39;s info | 

### Return type

[**GetVideoListResponse**](GetVideoListResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCaption

> SetCaption(id string, lan string, isDefault SetDefaultCaptionRequest) (*ResponseSuccess, error)

> SetCaptionWithContext(ctx context.Context, id string, lan string, isDefault SetDefaultCaptionRequest) (*ResponseSuccess, error)


Set default video caption



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
    isDefault := *w3streamsdk.NewSetDefaultCaptionRequest() // SetDefaultCaptionRequest | Set Default Caption Request

    
    res, err := client.Video.SetCaption(id, lan, isDefault)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Video.SetCaption``: %v\n", err)
    }
    // response from `SetCaption`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Video.SetCaption`")
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
**isDefault** | [**SetDefaultCaptionRequest**](SetDefaultCaptionRequest.md) | Set Default Caption Request | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadPart

> UploadPartFile(id string, file *os.File) (*ResponseSuccess, error)
> UploadPart(id string, fileName string, fileReader io.Reader, fileSize int64)
> UploadPartFileWithContext(ctx context.Context, id string, file *os.File) (*ResponseSuccess, error)
> UploadPartWithContext(ctx context.Context, id string, fileName string, fileReader io.Reader, fileSize int64)

Upload part of video



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
        
    id := "id_example" // string | video's id
    file := os.NewFile(1234, "some_file") // *os.File | File video to be uploaded
    hash := "hash_example" // string | Md5 hash of part
    index := "index_example" // string | Index of the part

    
    res, err := client.Video.UploadPartFile(id, file)

    // you can also use a Reader instead of a File:
    // we recommend using Reader instead!
    // client.Video.UploadPart(id, fileName, fileReader, fileSize)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Video.UploadPart``: %v\n", err)
    }
    // response from `UploadPart`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Video.UploadPart`")
    fmt.Println(string(newJsonString))
}
```


### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | video&#39;s id | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | File video to be uploaded | 
**hash** | **string** | Md5 hash of part | 
**index** | **string** | Index of the part | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadVideoComplete

> UploadVideoComplete(id string) (*ResponseSuccess, error)

> UploadVideoCompleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Get upload video when complete



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
        
    id := "id_example" // string | video's id

    
    res, err := client.Video.UploadVideoComplete(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Video.UploadVideoComplete``: %v\n", err)
    }
    // response from `UploadVideoComplete`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Video.UploadVideoComplete`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | video&#39;s id | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

