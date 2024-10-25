# \LiveStream

All URIs are relative to https://api.w3stream.xyz/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLiveStreamKey**](LiveStream.md#CreateLiveStreamKey) | **Post** /live_streams | Create live stream key
[**CreateStreaming**](LiveStream.md#CreateStreaming) | **Post** /live_streams/{id}/streamings | Create a new live stream video
[**DeleteLiveStreamKey**](LiveStream.md#DeleteLiveStreamKey) | **Delete** /live_streams/{id} | Delete live stream key
[**DeleteLiveStreamVideo**](LiveStream.md#DeleteLiveStreamVideo) | **Delete** /live_streams/{id}/videos | Delete live stream video
[**GetLiveStreamKey**](LiveStream.md#GetLiveStreamKey) | **Get** /live_streams/{id} | Get live stream key
[**GetLiveStreamKeys**](LiveStream.md#GetLiveStreamKeys) | **Get** /live_streams | Get live stream key list
[**GetLiveStreamPlayerInfo**](LiveStream.md#GetLiveStreamPlayerInfo) | **Get** /live_streams/player/{id}/videos | Get live stream video public
[**GetLiveStreamVideo**](LiveStream.md#GetLiveStreamVideo) | **Get** /live_streams/{id}/videos | Get live stream video
[**GetLiveStreamVideos**](LiveStream.md#GetLiveStreamVideos) | **Post** /live_streams/{id}/videos | Get live stream videos
[**GetStreaming**](LiveStream.md#GetStreaming) | **Get** /live_streams/{id}/streamings/{stream_id} | Get live stream video streaming
[**GetStreamings**](LiveStream.md#GetStreamings) | **Get** /live_streams/{id}/streamings | Get live stream video streamings
[**UpdateLiveStreamKey**](LiveStream.md#UpdateLiveStreamKey) | **Put** /live_streams/{id} | Update live stream key
[**UpdateLiveStreamVideo**](LiveStream.md#UpdateLiveStreamVideo) | **Put** /live_streams/{id}/streamings | Update live stream video



## CreateLiveStreamKey

> CreateLiveStreamKey(input CreateLiveStreamKeyRequest) (*CreateLiveStreamKeyResponse, error)

> CreateLiveStreamKeyWithContext(ctx context.Context, input CreateLiveStreamKeyRequest) (*CreateLiveStreamKeyResponse, error)


Create live stream key



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
        
    input := *w3streamsdk.NewCreateLiveStreamKeyRequest() // CreateLiveStreamKeyRequest | CreateLiveStreamKeyRequest

    
    res, err := client.LiveStream.CreateLiveStreamKey(input)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.CreateLiveStreamKey``: %v\n", err)
    }
    // response from `CreateLiveStreamKey`: CreateLiveStreamKeyResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.CreateLiveStreamKey`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**input** | [**CreateLiveStreamKeyRequest**](CreateLiveStreamKeyRequest.md) | CreateLiveStreamKeyRequest | 

### Return type

[**CreateLiveStreamKeyResponse**](CreateLiveStreamKeyResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStreaming

> CreateStreaming(id string, input CreateStreamingRequest) (*CreateStreamingResponse, error)

> CreateStreamingWithContext(ctx context.Context, id string, input CreateStreamingRequest) (*CreateStreamingResponse, error)


Create a new live stream video



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
        
    id := "id_example" // string | Live stream key ID
    input := *w3streamsdk.NewCreateStreamingRequest() // CreateStreamingRequest | CreateStreamingRequest

    
    res, err := client.LiveStream.CreateStreaming(id, input)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.CreateStreaming``: %v\n", err)
    }
    // response from `CreateStreaming`: CreateStreamingResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.CreateStreaming`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream key ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**input** | [**CreateStreamingRequest**](CreateStreamingRequest.md) | CreateStreamingRequest | 

### Return type

[**CreateStreamingResponse**](CreateStreamingResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLiveStreamKey

> DeleteLiveStreamKey(id string) (*ResponseSuccess, error)

> DeleteLiveStreamKeyWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete live stream key



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
        
    id := "id_example" // string | Live stream key ID

    
    res, err := client.LiveStream.DeleteLiveStreamKey(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.DeleteLiveStreamKey``: %v\n", err)
    }
    // response from `DeleteLiveStreamKey`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.DeleteLiveStreamKey`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream key ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLiveStreamVideo

> DeleteLiveStreamVideo(id string) (*ResponseSuccess, error)

> DeleteLiveStreamVideoWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete live stream video



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
        
    id := "id_example" // string | Live stream video ID

    
    res, err := client.LiveStream.DeleteLiveStreamVideo(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.DeleteLiveStreamVideo``: %v\n", err)
    }
    // response from `DeleteLiveStreamVideo`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.DeleteLiveStreamVideo`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream video ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveStreamKey

> GetLiveStreamKey(id string) (*GetLiveStreamKeyResponse, error)

> GetLiveStreamKeyWithContext(ctx context.Context, id string) (*GetLiveStreamKeyResponse, error)


Get live stream key



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
        
    id := "id_example" // string | ID

    
    res, err := client.LiveStream.GetLiveStreamKey(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetLiveStreamKey``: %v\n", err)
    }
    // response from `GetLiveStreamKey`: GetLiveStreamKeyResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetLiveStreamKey`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetLiveStreamKeyResponse**](GetLiveStreamKeyResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveStreamKeys

> GetLiveStreamKeys(r LiveStreamApiGetLiveStreamKeysRequest) (*GetLiveStreamKeysListResponse, error)


> GetLiveStreamKeysWithContext(ctx context.Context, r LiveStreamApiGetLiveStreamKeysRequest) (*GetLiveStreamKeysListResponse, error)



Get live stream key list



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
    req := w3streamsdk.LiveStreamApiGetLiveStreamKeysRequest{}
    
    req.Search("search_example") // string | only support search by name
    req.SortBy("sortBy_example") // string | sort by (default to "created_at")
    req.OrderBy("orderBy_example") // string | allowed: asc, desc. Default: asc (default to "asc")
    req.Offset(int32(56)) // int32 | offset, allowed values greater than or equal to 0.
    req.Limit(int32(56)) // int32 | results per page.

    res, err := client.LiveStream.GetLiveStreamKeys(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetLiveStreamKeys``: %v\n", err)
    }
    // response from `GetLiveStreamKeys`: GetLiveStreamKeysListResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetLiveStreamKeys`")
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
**offset** | **int32** | offset, allowed values greater than or equal to 0. | 
**limit** | **int32** | results per page. | 

### Return type

[**GetLiveStreamKeysListResponse**](GetLiveStreamKeysListResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveStreamPlayerInfo

> GetLiveStreamPlayerInfo(id string) (*GetLiveStreamVideoPublicResponse, error)

> GetLiveStreamPlayerInfoWithContext(ctx context.Context, id string) (*GetLiveStreamVideoPublicResponse, error)


Get live stream video public



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
        
    id := "id_example" // string | Live stream key ID

    
    res, err := client.LiveStream.GetLiveStreamPlayerInfo(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetLiveStreamPlayerInfo``: %v\n", err)
    }
    // response from `GetLiveStreamPlayerInfo`: GetLiveStreamVideoPublicResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetLiveStreamPlayerInfo`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream key ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetLiveStreamVideoPublicResponse**](GetLiveStreamVideoPublicResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveStreamVideo

> GetLiveStreamVideo(id string) (*GetLiveStreamVideoResponse, error)

> GetLiveStreamVideoWithContext(ctx context.Context, id string) (*GetLiveStreamVideoResponse, error)


Get live stream video



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
        
    id := "id_example" // string | Live stream video ID

    
    res, err := client.LiveStream.GetLiveStreamVideo(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetLiveStreamVideo``: %v\n", err)
    }
    // response from `GetLiveStreamVideo`: GetLiveStreamVideoResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetLiveStreamVideo`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream video ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetLiveStreamVideoResponse**](GetLiveStreamVideoResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveStreamVideos

> GetLiveStreamVideos(id string, data GetLiveStreamVideosRequest) (*GetLiveStreamVideosResponse, error)

> GetLiveStreamVideosWithContext(ctx context.Context, id string, data GetLiveStreamVideosRequest) (*GetLiveStreamVideosResponse, error)


Get live stream videos



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
        
    id := "id_example" // string | Live stream key ID
    data := *w3streamsdk.NewGetLiveStreamVideosRequest() // GetLiveStreamVideosRequest | data

    
    res, err := client.LiveStream.GetLiveStreamVideos(id, data)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetLiveStreamVideos``: %v\n", err)
    }
    // response from `GetLiveStreamVideos`: GetLiveStreamVideosResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetLiveStreamVideos`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream key ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**data** | [**GetLiveStreamVideosRequest**](GetLiveStreamVideosRequest.md) | data | 

### Return type

[**GetLiveStreamVideosResponse**](GetLiveStreamVideosResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreaming

> GetStreaming(id string, streamId string) (*GetStreamingResponse, error)

> GetStreamingWithContext(ctx context.Context, id string, streamId string) (*GetStreamingResponse, error)


Get live stream video streaming



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
        
    id := "id_example" // string | Live stream key ID
    streamId := "streamId_example" // string | Stream ID

    
    res, err := client.LiveStream.GetStreaming(id, streamId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetStreaming``: %v\n", err)
    }
    // response from `GetStreaming`: GetStreamingResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetStreaming`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream key ID | 
**streamId** | **string** | Stream ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetStreamingResponse**](GetStreamingResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamings

> GetStreamings(id string) (*GetStreamingsResponse, error)

> GetStreamingsWithContext(ctx context.Context, id string) (*GetStreamingsResponse, error)


Get live stream video streamings



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
        
    id := "id_example" // string | Live stream key ID

    
    res, err := client.LiveStream.GetStreamings(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetStreamings``: %v\n", err)
    }
    // response from `GetStreamings`: GetStreamingsResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetStreamings`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream key ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetStreamingsResponse**](GetStreamingsResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLiveStreamKey

> UpdateLiveStreamKey(id string, input UpdateLiveStreamKeyRequest) (*UpdateLiveStreamKeyResponse, error)

> UpdateLiveStreamKeyWithContext(ctx context.Context, id string, input UpdateLiveStreamKeyRequest) (*UpdateLiveStreamKeyResponse, error)


Update live stream key



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
        
    id := "id_example" // string | Live stream key ID
    input := *w3streamsdk.NewUpdateLiveStreamKeyRequest() // UpdateLiveStreamKeyRequest | UpdateLiveStreamKeyRequest

    
    res, err := client.LiveStream.UpdateLiveStreamKey(id, input)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.UpdateLiveStreamKey``: %v\n", err)
    }
    // response from `UpdateLiveStreamKey`: UpdateLiveStreamKeyResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.UpdateLiveStreamKey`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream key ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**input** | [**UpdateLiveStreamKeyRequest**](UpdateLiveStreamKeyRequest.md) | UpdateLiveStreamKeyRequest | 

### Return type

[**UpdateLiveStreamKeyResponse**](UpdateLiveStreamKeyResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLiveStreamVideo

> UpdateLiveStreamVideo(id string, data UpdateLiveStreamVideoRequest) (*ResponseSuccess, error)

> UpdateLiveStreamVideoWithContext(ctx context.Context, id string, data UpdateLiveStreamVideoRequest) (*ResponseSuccess, error)


Update live stream video



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
        
    id := "id_example" // string | Live stream key ID
    data := *w3streamsdk.NewUpdateLiveStreamVideoRequest() // UpdateLiveStreamVideoRequest | data

    
    res, err := client.LiveStream.UpdateLiveStreamVideo(id, data)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.UpdateLiveStreamVideo``: %v\n", err)
    }
    // response from `UpdateLiveStreamVideo`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.UpdateLiveStreamVideo`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream key ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**data** | [**UpdateLiveStreamVideoRequest**](UpdateLiveStreamVideoRequest.md) | data | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

