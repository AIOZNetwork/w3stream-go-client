# \Playlist

All URIs are relative to https://api.w3stream.xyz/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVideoToPlaylist**](Playlist.md#AddVideoToPlaylist) | **Post** /playlists/{id}/items | Add a video to a playlist
[**CreatePlaylist**](Playlist.md#CreatePlaylist) | **Post** /playlists/create | Create a playlist
[**DeletePlaylistById**](Playlist.md#DeletePlaylistById) | **Delete** /playlists/{id} | Delete a playlist by ID
[**DeletePlaylistThumbnail**](Playlist.md#DeletePlaylistThumbnail) | **Delete** /playlists/{id}/thumbnail | Delete a playlist thumbnail
[**GetPlaylistById**](Playlist.md#GetPlaylistById) | **Get** /playlists/{id} | Get playlist by ID
[**GetPlaylistPublicInfo**](Playlist.md#GetPlaylistPublicInfo) | **Get** /playlists/{id}/player.json | Get a playlist public
[**GetPlaylists**](Playlist.md#GetPlaylists) | **Post** /playlists | Get user&#39;s playlists
[**MoveVideoInPlaylist**](Playlist.md#MoveVideoInPlaylist) | **Put** /playlists/{id}/items | Move a video in a playlist
[**RemoveVideoFromPlaylist**](Playlist.md#RemoveVideoFromPlaylist) | **Delete** /playlists/{id}/items/{item_id} | Remove a video from a playlist
[**UpdatePlaylist**](Playlist.md#UpdatePlaylist) | **Patch** /playlists/{id} | Update a playlist



## AddVideoToPlaylist

> AddVideoToPlaylist(id string, payload AddVideoToPlaylistRequest) (*ResponseSuccess, error)

> AddVideoToPlaylistWithContext(ctx context.Context, id string, payload AddVideoToPlaylistRequest) (*ResponseSuccess, error)


Add a video to a playlist



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
        
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Playlist ID
    payload := *w3streamsdk.NewAddVideoToPlaylistRequest() // AddVideoToPlaylistRequest | Video details

    
    res, err := client.Playlist.AddVideoToPlaylist(id, payload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.AddVideoToPlaylist``: %v\n", err)
    }
    // response from `AddVideoToPlaylist`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.AddVideoToPlaylist`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Playlist ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**payload** | [**AddVideoToPlaylistRequest**](AddVideoToPlaylistRequest.md) | Video details | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlaylist

> CreatePlaylist(request CreatePlaylistRequest) (*CreatePlaylistResponse, error)

> CreatePlaylistWithContext(ctx context.Context, request CreatePlaylistRequest) (*CreatePlaylistResponse, error)


Create a playlist



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
        
    request := *w3streamsdk.NewCreatePlaylistRequest() // CreatePlaylistRequest | Playlist input

    
    res, err := client.Playlist.CreatePlaylist(request)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.CreatePlaylist``: %v\n", err)
    }
    // response from `CreatePlaylist`: CreatePlaylistResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.CreatePlaylist`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**CreatePlaylistRequest**](CreatePlaylistRequest.md) | Playlist input | 

### Return type

[**CreatePlaylistResponse**](CreatePlaylistResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlaylistById

> DeletePlaylistById(id string) (*ResponseSuccess, error)

> DeletePlaylistByIdWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete a playlist by ID



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
        
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Playlist ID

    
    res, err := client.Playlist.DeletePlaylistById(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.DeletePlaylistById``: %v\n", err)
    }
    // response from `DeletePlaylistById`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.DeletePlaylistById`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Playlist ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlaylistThumbnail

> DeletePlaylistThumbnail(id string) (*ResponseSuccess, error)

> DeletePlaylistThumbnailWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete a playlist thumbnail



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
        
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Playlist ID

    
    res, err := client.Playlist.DeletePlaylistThumbnail(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.DeletePlaylistThumbnail``: %v\n", err)
    }
    // response from `DeletePlaylistThumbnail`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.DeletePlaylistThumbnail`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Playlist ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaylistById

> GetPlaylistById(id string, r PlaylistApiGetPlaylistByIdRequest) (*GetPlaylistByIdResponse, error)


> GetPlaylistByIdWithContext(ctx context.Context, id string, r PlaylistApiGetPlaylistByIdRequest) (*GetPlaylistByIdResponse, error)



Get playlist by ID



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
    req := w3streamsdk.PlaylistApiGetPlaylistByIdRequest{}
    
    req.Id("id_example") // string | Playlist ID
    req.SortBy("sortBy_example") // string | sort by (default to "created_at")
    req.OrderBy("orderBy_example") // string | allowed: asc, desc. Default: asc (default to "asc")

    res, err := client.Playlist.GetPlaylistById(id string, req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.GetPlaylistById``: %v\n", err)
    }
    // response from `GetPlaylistById`: GetPlaylistByIdResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.GetPlaylistById`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Playlist ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sortBy** | **string** | sort by | [default to &quot;created_at&quot;]
**orderBy** | **string** | allowed: asc, desc. Default: asc | [default to &quot;asc&quot;]

### Return type

[**GetPlaylistByIdResponse**](GetPlaylistByIdResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaylistPublicInfo

> GetPlaylistPublicInfo(id string) (*PublicPlaylistObject, error)

> GetPlaylistPublicInfoWithContext(ctx context.Context, id string) (*PublicPlaylistObject, error)


Get a playlist public



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
        
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Playlist ID

    
    res, err := client.Playlist.GetPlaylistPublicInfo(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.GetPlaylistPublicInfo``: %v\n", err)
    }
    // response from `GetPlaylistPublicInfo`: PublicPlaylistObject
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.GetPlaylistPublicInfo`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Playlist ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PublicPlaylistObject**](PublicPlaylistObject.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaylists

> GetPlaylists(request GetPlaylistListRequest) (*GetPlaylistListResponse, error)

> GetPlaylistsWithContext(ctx context.Context, request GetPlaylistListRequest) (*GetPlaylistListResponse, error)


Get user's playlists



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
        
    request := *w3streamsdk.NewGetPlaylistListRequest() // GetPlaylistListRequest | Playlist filter

    
    res, err := client.Playlist.GetPlaylists(request)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.GetPlaylists``: %v\n", err)
    }
    // response from `GetPlaylists`: GetPlaylistListResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.GetPlaylists`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**GetPlaylistListRequest**](GetPlaylistListRequest.md) | Playlist filter | 

### Return type

[**GetPlaylistListResponse**](GetPlaylistListResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveVideoInPlaylist

> MoveVideoInPlaylist(id string, payload MoveVideoInPlaylistRequest) (*ResponseSuccess, error)

> MoveVideoInPlaylistWithContext(ctx context.Context, id string, payload MoveVideoInPlaylistRequest) (*ResponseSuccess, error)


Move a video in a playlist



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
        
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Playlist ID
    payload := *w3streamsdk.NewMoveVideoInPlaylistRequest() // MoveVideoInPlaylistRequest | Video details

    
    res, err := client.Playlist.MoveVideoInPlaylist(id, payload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.MoveVideoInPlaylist``: %v\n", err)
    }
    // response from `MoveVideoInPlaylist`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.MoveVideoInPlaylist`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Playlist ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**payload** | [**MoveVideoInPlaylistRequest**](MoveVideoInPlaylistRequest.md) | Video details | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveVideoFromPlaylist

> RemoveVideoFromPlaylist(id string, itemId string) (*ResponseSuccess, error)

> RemoveVideoFromPlaylistWithContext(ctx context.Context, id string, itemId string) (*ResponseSuccess, error)


Remove a video from a playlist



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
        
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Playlist ID
    itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Playlist Item ID

    
    res, err := client.Playlist.RemoveVideoFromPlaylist(id, itemId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.RemoveVideoFromPlaylist``: %v\n", err)
    }
    // response from `RemoveVideoFromPlaylist`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.RemoveVideoFromPlaylist`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Playlist ID | 
**itemId** | **string** | Playlist Item ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlaylist

> UpdatePlaylistFile(id string, file *os.File) (*ResponseSuccess, error)
> UpdatePlaylist(id string, fileName string, fileReader io.Reader)
> UpdatePlaylistFileWithContext(ctx context.Context, id string, file *os.File) (*ResponseSuccess, error)
> UpdatePlaylistWithContext(ctx context.Context, id string, fileName string, fileReader io.Reader)

Update a playlist



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
        
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Playlist ID
    file := os.NewFile(1234, "some_file") // *os.File | 
    metadata := []w3streamsdk.Metadata{*w3streamsdk.NewMetadata()} // []Metadata | 
    name := "name_example" // string | 
    tags := []string{"Inner_example"} // []string | 

    
    res, err := client.Playlist.UpdatePlaylistFile(id, file)

    // you can also use a Reader instead of a File:
    // we recommend using Reader instead!
    // client.Playlist.UpdatePlaylist(id, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.UpdatePlaylist``: %v\n", err)
    }
    // response from `UpdatePlaylist`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.UpdatePlaylist`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Playlist ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** |  | 
**metadata** | [**[]Metadata**](Metadata.md) |  | 
**name** | **string** |  | 
**tags** | **[]string** |  | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

