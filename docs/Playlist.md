# \Playlist

All URIs are relative to https://api.w3stream.xyz/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteThumbnail**](Playlist.md#DeleteThumbnail) | **Delete** /playlists/{id}/thumbnail | Delete a playlist thumbnail
[**AddItem**](Playlist.md#AddItem) | **Post** /playlists/{id}/items | Add a video to a playlist
[**CreatePlaylist**](Playlist.md#CreatePlaylist) | **Post** /playlists/create | Create a new playlist
[**DeleteItem**](Playlist.md#DeleteItem) | **Delete** /playlists/{id}/items/{item_id} | Remove a video from a playlist
[**DeletePlaylist**](Playlist.md#DeletePlaylist) | **Delete** /playlists/{id} | Delete a playlist by ID
[**GetPlaylistById**](Playlist.md#GetPlaylistById) | **Get** /playlists/{id} | Get playlist by ID
[**GetPlaylistInfo**](Playlist.md#GetPlaylistInfo) | **Get** /playlists/{id}/player.json | Get a playlist public
[**GetPlaylists**](Playlist.md#GetPlaylists) | **Post** /playlists | Get user&#39;s playlists
[**MoveItems**](Playlist.md#MoveItems) | **Put** /playlists/{id}/items | Move a video within a playlist
[**UpdatePlaylist**](Playlist.md#UpdatePlaylist) | **Patch** /playlists/{id} | Update a playlist



## DeleteThumbnail

> DeleteThumbnail(id string) (*ResponseSuccess, error)

> DeleteThumbnailWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


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

    
    res, err := client.Playlist.DeleteThumbnail(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.DeleteThumbnail``: %v\n", err)
    }
    // response from `DeleteThumbnail`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.DeleteThumbnail`")
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


## AddItem

> AddItem(id string, request AddVideoToPlaylistRequest) (*ResponseSuccess, error)

> AddItemWithContext(ctx context.Context, id string, request AddVideoToPlaylistRequest) (*ResponseSuccess, error)


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
    request := *w3streamsdk.NewAddVideoToPlaylistRequest() // AddVideoToPlaylistRequest | Video details

    
    res, err := client.Playlist.AddItem(id, request)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.AddItem``: %v\n", err)
    }
    // response from `AddItem`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.AddItem`")
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
**request** | [**AddVideoToPlaylistRequest**](AddVideoToPlaylistRequest.md) | Video details | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlaylist

> CreatePlaylist(payload CreatePlaylistRequest) (*CreatePlaylistResponse, error)

> CreatePlaylistWithContext(ctx context.Context, payload CreatePlaylistRequest) (*CreatePlaylistResponse, error)


Create a new playlist



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
        
    payload := *w3streamsdk.NewCreatePlaylistRequest() // CreatePlaylistRequest | Create playlist request

    
    res, err := client.Playlist.CreatePlaylist(payload)

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
**payload** | [**CreatePlaylistRequest**](CreatePlaylistRequest.md) | Create playlist request | 

### Return type

[**CreatePlaylistResponse**](CreatePlaylistResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteItem

> DeleteItem(id string, itemId string) (*ResponseSuccess, error)

> DeleteItemWithContext(ctx context.Context, id string, itemId string) (*ResponseSuccess, error)


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

    
    res, err := client.Playlist.DeleteItem(id, itemId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.DeleteItem``: %v\n", err)
    }
    // response from `DeleteItem`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.DeleteItem`")
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


## DeletePlaylist

> DeletePlaylist(id string) (*ResponseSuccess, error)

> DeletePlaylistWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


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

    
    res, err := client.Playlist.DeletePlaylist(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.DeletePlaylist``: %v\n", err)
    }
    // response from `DeletePlaylist`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.DeletePlaylist`")
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
    req.SortBy("sortBy_example") // string | Sort field
    req.OrderBy("orderBy_example") // string | Sort order

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
**sortBy** | **string** | Sort field | 
**orderBy** | **string** | Sort order | 

### Return type

[**GetPlaylistByIdResponse**](GetPlaylistByIdResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaylistInfo

> GetPlaylistInfo(id string) (*PublicPlaylistObject, error)

> GetPlaylistInfoWithContext(ctx context.Context, id string) (*PublicPlaylistObject, error)


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

    
    res, err := client.Playlist.GetPlaylistInfo(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.GetPlaylistInfo``: %v\n", err)
    }
    // response from `GetPlaylistInfo`: PublicPlaylistObject
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.GetPlaylistInfo`")
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

> GetPlaylists(payload GetPlaylistListRequest) (*GetPlaylistListResponse, error)

> GetPlaylistsWithContext(ctx context.Context, payload GetPlaylistListRequest) (*GetPlaylistListResponse, error)


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
        
    payload := *w3streamsdk.NewGetPlaylistListRequest() // GetPlaylistListRequest | Get playlist list request

    
    res, err := client.Playlist.GetPlaylists(payload)

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
**payload** | [**GetPlaylistListRequest**](GetPlaylistListRequest.md) | Get playlist list request | 

### Return type

[**GetPlaylistListResponse**](GetPlaylistListResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveItems

> MoveItems(id string, payload MoveVideoInPlaylistRequest) (*ResponseSuccess, error)

> MoveItemsWithContext(ctx context.Context, id string, payload MoveVideoInPlaylistRequest) (*ResponseSuccess, error)


Move a video within a playlist



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
    payload := *w3streamsdk.NewMoveVideoInPlaylistRequest() // MoveVideoInPlaylistRequest | Move video details

    
    res, err := client.Playlist.MoveItems(id, payload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.MoveItems``: %v\n", err)
    }
    // response from `MoveItems`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.MoveItems`")
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
**payload** | [**MoveVideoInPlaylistRequest**](MoveVideoInPlaylistRequest.md) | Move video details | 

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
    file := os.NewFile(1234, "some_file") // *os.File | Thumbnail
    name := "name_example" // string | Playlist name
    tags := []string{"Inner_example"} // []string | Tags
    metadata := []string{"Inner_example"} // []string | Metadata

    
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
**file** | ***os.File** | Thumbnail | 
**name** | **string** | Playlist name | 
**tags** | **[]string** | Tags | 
**metadata** | **[]string** | Metadata | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

