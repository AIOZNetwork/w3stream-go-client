# \Players

All URIs are relative to https://api.w3stream.xyz/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](Players.md#Create) | **Post** /players | Create a player theme
[**Get**](Players.md#Get) | **Get** /players/{id} | Get a player theme by ID
[**Update**](Players.md#Update) | **Patch** /players/{id} | Update a player theme by ID
[**Delete**](Players.md#Delete) | **Delete** /players/{id} | Delete a player theme by ID
[**List**](Players.md#List) | **Get** /players | List all player themes
[**UploadLogo**](Players.md#UploadLogo) | **Post** /players/{id}/logo | Upload a logo for a player theme by ID
[**DeleteLogo**](Players.md#DeleteLogo) | **Delete** /players/{id}/logo | Delete a logo for a player theme by ID
[**AddPlayer**](Players.md#AddPlayer) | **Post** /players/add-player | Add a player theme to a video
[**RemovePlayer**](Players.md#RemovePlayer) | **Post** /players/remove-player | Remove a player theme from a video



## Create

> Create(request CreatePlayerThemeRequest) (*CreatePlayerThemesResponse, error)

> CreateWithContext(ctx context.Context, request CreatePlayerThemeRequest) (*CreatePlayerThemesResponse, error)


Create a player theme



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
        
    request := *w3streamsdk.NewCreatePlayerThemeRequest() // CreatePlayerThemeRequest | Player theme input

    
    res, err := client.Players.Create(request)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.Create``: %v\n", err)
    }
    // response from `Create`: CreatePlayerThemesResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.Create`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**CreatePlayerThemeRequest**](CreatePlayerThemeRequest.md) | Player theme input | 

### Return type

[**CreatePlayerThemesResponse**](CreatePlayerThemesResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Get(id string) (*GetPlayerThemeByIdResponse, error)

> GetWithContext(ctx context.Context, id string) (*GetPlayerThemeByIdResponse, error)


Get a player theme by ID



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
        
    id := "id_example" // string | Player theme ID

    
    res, err := client.Players.Get(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.Get``: %v\n", err)
    }
    // response from `Get`: GetPlayerThemeByIdResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.Get`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Player theme ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetPlayerThemeByIdResponse**](GetPlayerThemeByIdResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(id string, input UpdatePlayerThemeRequest) (*UpdatePlayerThemeResponse, error)

> UpdateWithContext(ctx context.Context, id string, input UpdatePlayerThemeRequest) (*UpdatePlayerThemeResponse, error)


Update a player theme by ID



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
        
    id := "id_example" // string | Player theme ID
    input := *w3streamsdk.NewUpdatePlayerThemeRequest() // UpdatePlayerThemeRequest | Player theme input

    
    res, err := client.Players.Update(id, input)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.Update``: %v\n", err)
    }
    // response from `Update`: UpdatePlayerThemeResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.Update`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Player theme ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**input** | [**UpdatePlayerThemeRequest**](UpdatePlayerThemeRequest.md) | Player theme input | 

### Return type

[**UpdatePlayerThemeResponse**](UpdatePlayerThemeResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(id string) (*ResponseSuccess, error)

> DeleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete a player theme by ID



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
        
    id := "id_example" // string | Player theme ID

    
    res, err := client.Players.Delete(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.Delete``: %v\n", err)
    }
    // response from `Delete`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.Delete`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Player theme ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(r PlayersApiListRequest) (*GetPlayerThemeResponse, error)


> ListWithContext(ctx context.Context, r PlayersApiListRequest) (*GetPlayerThemeResponse, error)



List all player themes



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
    req := w3streamsdk.PlayersApiListRequest{}
    
    req.Search("search_example") // string | only support search by name
    req.SortBy("sortBy_example") // string | sort by (default to "created_at")
    req.OrderBy("orderBy_example") // string | allowed: asc, desc. Default: asc (default to "asc")
    req.Offset(int32(56)) // int32 | offset, allowed values greater than or equal to 0. Default(0) (default to 0)
    req.Limit(int32(56)) // int32 | results per page. Allowed values 1-100, default is 25 (default to 25)

    res, err := client.Players.List(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.List``: %v\n", err)
    }
    // response from `List`: GetPlayerThemeResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.List`")
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

[**GetPlayerThemeResponse**](GetPlayerThemeResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadLogo

> UploadLogoFile(id string, file *os.File, link string) (*UploadLogoByIdResponse, error)
> UploadLogo(id string, link string, fileName string, fileReader io.Reader)
> UploadLogoFileWithContext(ctx context.Context, id string, file *os.File, link string) (*UploadLogoByIdResponse, error)
> UploadLogoWithContext(ctx context.Context, id string, link string, fileName string, fileReader io.Reader)

Upload a logo for a player theme by ID



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
        
    id := "id_example" // string | Player theme ID
    file := os.NewFile(1234, "some_file") // *os.File | The uploaded file (JPG or PNG)
    link := "link_example" // string | The link to the logo (optional if a file is provided)

    
    res, err := client.Players.UploadLogoFile(id, file, link)

    // you can also use a Reader instead of a File:
    // we recommend using Reader instead!
    // client.Players.UploadLogo(id, link, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.UploadLogo``: %v\n", err)
    }
    // response from `UploadLogo`: UploadLogoByIdResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.UploadLogo`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Player theme ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | The uploaded file (JPG or PNG) | 
**link** | **string** | The link to the logo (optional if a file is provided) | 

### Return type

[**UploadLogoByIdResponse**](UploadLogoByIdResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogo

> DeleteLogo(id string) (*ResponseSuccess, error)

> DeleteLogoWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete a logo for a player theme by ID



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
        
    id := "id_example" // string | Player theme ID

    
    res, err := client.Players.DeleteLogo(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.DeleteLogo``: %v\n", err)
    }
    // response from `DeleteLogo`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.DeleteLogo`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Player theme ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddPlayer

> AddPlayer(request AddPlayerThemesToVideoRequest) (*ResponseSuccess, error)

> AddPlayerWithContext(ctx context.Context, request AddPlayerThemesToVideoRequest) (*ResponseSuccess, error)


Add a player theme to a video



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
        
    request := *w3streamsdk.NewAddPlayerThemesToVideoRequest() // AddPlayerThemesToVideoRequest | Add player theme to video request

    
    res, err := client.Players.AddPlayer(request)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.AddPlayer``: %v\n", err)
    }
    // response from `AddPlayer`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.AddPlayer`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**AddPlayerThemesToVideoRequest**](AddPlayerThemesToVideoRequest.md) | Add player theme to video request | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePlayer

> RemovePlayer(request RemovePlayerThemesFromVideoRequest) (*ResponseSuccess, error)

> RemovePlayerWithContext(ctx context.Context, request RemovePlayerThemesFromVideoRequest) (*ResponseSuccess, error)


Remove a player theme from a video



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
        
    request := *w3streamsdk.NewRemovePlayerThemesFromVideoRequest() // RemovePlayerThemesFromVideoRequest | Remove player theme from video request

    
    res, err := client.Players.RemovePlayer(request)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.RemovePlayer``: %v\n", err)
    }
    // response from `RemovePlayer`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.RemovePlayer`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**RemovePlayerThemesFromVideoRequest**](RemovePlayerThemesFromVideoRequest.md) | Remove player theme from video request | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

