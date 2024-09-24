<!--<documentation_excluded>-->
<h1 align="center">W3Stream Go client</h1>

W3Stream is the video infrastructure for product builders. Lightning fast video APIs for integrating, scaling, and managing on-demand & low latency live streaming features in your app.

## Project description

W3Stream's Go client streamlines the coding process. Chunking files is handled for you, as is pagination and refreshing your tokens.

## Getting started

### Installation
```bash
go get github.com/AIOZNetwork/w3stream-go-client
```


### Code sample

For a more advanced usage you can checkout the rest of the documentation in the [docs directory](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs)

```golang

import (
	"encoding/json"
	"fmt"
	"os"

	w3streamSdk "github.com/AIOZNetwork/w3stream-go-client"
)

func main() {
    // Connect to production environment
    secretKey := "YOUR_SECRET_KEY" // Replace with your actual secret key
    publicKey := "YOUR_PUBLIC_KEY" // Replace with your public key
	apiCreds := w3streamsdk.AuthCredentials{
		PublicKey: publicKey,
		SecretKey: secretKey,
	}
    client := w3streamsdk.ClientBuilder(apiCreds).Build()

    // Create a video object
	title := "Sample Video Title"
	videoData := w3streamsdk.CreateVideoRequest{
		Title: &title,
	}
	createResult, err := client.Video.Create(videoData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating video: %v\n", err)
		return
	}

    videoId := createResult.Data.VideoId // Get the video ID from the response

    // Open the video file
    videoFile, err := os.Open("/path/to/video.mp4")
    if err != nil {
        fmt.Println("Error opening video file:", err)
        return
    }
    defer videoFile.Close() // Close the file after use

    fileInfo, err := videoFile.Stat()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error getting file info: %v\n", err)
        return
    }

    fileSize := fileInfo.Size()
    fileName := fileInfo.Name()

    // Option 1: Use client upload with videoId
	err = client.UploadVideo(context.Background(), *videoId, fileName, videoFile, fileSize)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error uploading video with client: %v\n", err)
		return
	}

    // Option 2: Upload parts yourself
    // This example is commented out as you already used option 1
	//_, err = client.Video.UploadPart(*videoId, nil, nil, "./path/to/video.mp4", videoFile, fileInfo.Size())
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error uploading video part: %v\n", err)
	//	return
	//}
	//
	//success, err := client.Video.UploadVideoComplete(*videoId)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error completing video upload: %v\n", err)
	//	return
	//}
	//
	//jsonString, err := json.MarshalIndent(success, "", "  ")
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error marshalling response: %v\n", err)
	//	return
	//}
	//fmt.Println(string(jsonString))
    fmt.Println("Video uploaded successfully!")
}
```

## Documentation

### API endpoints

All urls are relative to https://api.w3stream.xyz/api


#### Api key


##### Retrieve an instance of the Api key:
```go
secretKey := "YOUR_SECRET_KEY" // Replace with your actual secret key
publicKey := "YOUR_PUBLIC_KEY" // Replace with your public key
apiCreds := w3streamsdk.AuthCredentials{
	PublicKey: publicKey,
	SecretKey: secretKey,
}
client := w3streamsdk.ClientBuilder(apiCreds).Build()
apiKeyApi := client.ApiKey
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/ApiKey.md#Create) | **Post** `/api_keys` | Create API key
[**Update**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/ApiKey.md#Update) | **Patch** `/api_keys/{id}` | Rename API key
[**Delete**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/ApiKey.md#Delete) | **Delete** `/api_keys/{id}` | Delete API key
[**List**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/ApiKey.md#List) | **Get** `/api_keys` | Get list API keys


#### Players


##### Retrieve an instance of the players API:
```go
secretKey := "YOUR_SECRET_KEY" // Replace with your actual secret key
publicKey := "YOUR_PUBLIC_KEY" // Replace with your public key
apiCreds := w3streamsdk.AuthCredentials{
	PublicKey: publicKey,
	SecretKey: secretKey,
}
client := w3streamsdk.ClientBuilder(apiCreds).Build()
playersApi := client.Players
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Players.md#Create) | **Post** `/players` | Create a player theme
[**Get**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Players.md#Get) | **Get** `/players/{id}` | Get a player theme by ID
[**Update**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Players.md#Update) | **Patch** `/players/{id}` | Update a player theme by ID
[**Delete**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Players.md#Delete) | **Delete** `/players/{id}` | Delete a player theme by ID
[**List**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Players.md#List) | **Get** `/players` | List all player themes
[**UploadLogo**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Players.md#UploadLogo) | **Post** `/players/{id}/logo` | Upload a logo for a player theme by ID
[**DeleteLogo**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Players.md#DeleteLogo) | **Delete** `/players/{id}/logo` | Delete a logo for a player theme by ID
[**AddPlayer**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Players.md#AddPlayer) | **Post** `/players/add-player` | Add a player theme to a video
[**RemovePlayer**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Players.md#RemovePlayer) | **Post** `/players/remove-player` | Remove a player theme from a video


#### Video


##### Retrieve an instance of the video API:
```go
secretKey := "YOUR_SECRET_KEY" // Replace with your actual secret key
publicKey := "YOUR_PUBLIC_KEY" // Replace with your public key
apiCreds := w3streamsdk.AuthCredentials{
	PublicKey: publicKey,
	SecretKey: secretKey,
}
client := w3streamsdk.ClientBuilder(apiCreds).Build()
videoApi := client.Video
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Video.md#Create) | **Post** `/videos/create` | Create video object
[**Update**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Video.md#Update) | **Patch** `/videos/{id}` | update video info
[**Delete**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Video.md#Delete) | **Delete** `/videos/{id}` | Delete video
[**UploadThumbnail**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Video.md#UploadThumbnail) | **Post** `/videos/{id}/thumbnail` | Upload video thumbnail
[**CreateCaption**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Video.md#CreateCaption) | **Post** `/videos/{id}/captions/{lan}` | Create a new video caption
[**DeleteCaption**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Video.md#DeleteCaption) | **Delete** `/videos/{id}/captions/{lan}` | Delete a video caption
[**GetCaptions**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Video.md#GetCaptions) | **Get** `/videos/{id}/captions` | Get video captions
[**GetCost**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Video.md#GetCost) | **Get** `/videos/cost` | get video transcoding cost
[**GetDetail**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Video.md#GetDetail) | **Get** `/videos/{id}` | get video detail
[**GetVideoList**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Video.md#GetVideoList) | **Post** `/videos` | Get user videos list
[**SetCaption**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Video.md#SetCaption) | **Patch** `/videos/{id}/captions/{lan}` | Set default video caption
[**UploadPart**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Video.md#UploadPart) | **Post** `/videos/{id}/part` | Upload part of video
[**UploadVideoComplete**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Video.md#UploadVideoComplete) | **Get** `/videos/{id}/complete` | Get upload video when complete


#### Video chapter


##### Retrieve an instance of the video chapter API:
```go
secretKey := "YOUR_SECRET_KEY" // Replace with your actual secret key
publicKey := "YOUR_PUBLIC_KEY" // Replace with your public key
apiCreds := w3streamsdk.AuthCredentials{
	PublicKey: publicKey,
	SecretKey: secretKey,
}
client := w3streamsdk.ClientBuilder(apiCreds).Build()
videoChapterApi := client.VideoChapter
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/VideoChapter.md#Create) | **Post** `/videos/{id}/chapters/{lan}` | Create a video chapter
[**Get**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/VideoChapter.md#Get) | **Get** `/videos/{id}/chapters` | Get video chapters
[**Delete**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/VideoChapter.md#Delete) | **Delete** `/videos/{id}/chapters/{lan}` | Delete a video chapter


#### Watermark


##### Retrieve an instance of the watermark API:
```go
secretKey := "YOUR_SECRET_KEY" // Replace with your actual secret key
publicKey := "YOUR_PUBLIC_KEY" // Replace with your public key
apiCreds := w3streamsdk.AuthCredentials{
	PublicKey: publicKey,
	SecretKey: secretKey,
}
client := w3streamsdk.ClientBuilder(apiCreds).Build()
watermarkApi := client.Watermark
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Upload**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Watermark.md#Upload) | **Post** `/watermarks` | Create a new watermark
[**Delete**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Watermark.md#Delete) | **Delete** `/watermarks/{id}` | Delete a watermark by ID
[**List**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Watermark.md#List) | **Get** `/watermarks` | List all watermarks


#### Webhook


##### Retrieve an instance of the webhook API:
```go
secretKey := "YOUR_SECRET_KEY" // Replace with your actual secret key
publicKey := "YOUR_PUBLIC_KEY" // Replace with your public key
apiCreds := w3streamsdk.AuthCredentials{
	PublicKey: publicKey,
	SecretKey: secretKey,
}
client := w3streamsdk.ClientBuilder(apiCreds).Build()
webhookApi := client.Webhook
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Webhook.md#Create) | **Post** `/webhooks` | Create webhook
[**Get**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Webhook.md#Get) | **Get** `/webhooks/{id}` | Get user&#39;s webhook by id
[**Update**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Webhook.md#Update) | **Patch** `/webhooks/{id}` | Update event webhook
[**Delete**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Webhook.md#Delete) | **Delete** `/webhooks/{id}` | Delete webhook
[**List**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Webhook.md#List) | **Get** `/webhooks` | Get list webhooks
[**Check**](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Webhook.md#Check) | **Post** `/webhooks/check/{id}` | Check webhook by id




### Models

 - [AddPlayerThemesToVideoRequest](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/AddPlayerThemesToVideoRequest.md)
 - [ApiKey](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/ApiKey.md)
 - [Asset](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Asset.md)
 - [Controls](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Controls.md)
 - [CreateApiKeyData](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/CreateApiKeyData.md)
 - [CreateApiKeyRequest](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/CreateApiKeyRequest.md)
 - [CreateApiKeyResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/CreateApiKeyResponse.md)
 - [CreatePlayerThemeRequest](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/CreatePlayerThemeRequest.md)
 - [CreatePlayerThemesData](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/CreatePlayerThemesData.md)
 - [CreatePlayerThemesResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/CreatePlayerThemesResponse.md)
 - [CreateVideoCaptionData](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/CreateVideoCaptionData.md)
 - [CreateVideoCaptionResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/CreateVideoCaptionResponse.md)
 - [CreateVideoChapterData](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/CreateVideoChapterData.md)
 - [CreateVideoChapterResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/CreateVideoChapterResponse.md)
 - [CreateVideoRequest](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/CreateVideoRequest.md)
 - [CreateVideoResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/CreateVideoResponse.md)
 - [CreateWatermarkData](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/CreateWatermarkData.md)
 - [CreateWatermarkResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/CreateWatermarkResponse.md)
 - [CreateWebhookData](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/CreateWebhookData.md)
 - [CreateWebhookRequest](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/CreateWebhookRequest.md)
 - [CreateWebhookResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/CreateWebhookResponse.md)
 - [GetAllWatermarkData](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetAllWatermarkData.md)
 - [GetAllWatermarkResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetAllWatermarkResponse.md)
 - [GetApiKeysData](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetApiKeysData.md)
 - [GetApiKeysResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetApiKeysResponse.md)
 - [GetPlayerThemeByIdData](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetPlayerThemeByIdData.md)
 - [GetPlayerThemeByIdResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetPlayerThemeByIdResponse.md)
 - [GetPlayerThemeData](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetPlayerThemeData.md)
 - [GetPlayerThemeResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetPlayerThemeResponse.md)
 - [GetTranscodeCostData](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetTranscodeCostData.md)
 - [GetTranscodeCostResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetTranscodeCostResponse.md)
 - [GetUserWebhookData](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetUserWebhookData.md)
 - [GetUserWebhookResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetUserWebhookResponse.md)
 - [GetVideoCaptionsData](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetVideoCaptionsData.md)
 - [GetVideoCaptionsResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetVideoCaptionsResponse.md)
 - [GetVideoChaptersData](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetVideoChaptersData.md)
 - [GetVideoChaptersResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetVideoChaptersResponse.md)
 - [GetVideoDetailResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetVideoDetailResponse.md)
 - [GetVideoListData](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetVideoListData.md)
 - [GetVideoListRequest](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetVideoListRequest.md)
 - [GetVideoListResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetVideoListResponse.md)
 - [GetWebhooksListData](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetWebhooksListData.md)
 - [GetWebhooksListResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/GetWebhooksListResponse.md)
 - [Metadata](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Metadata.md)
 - [PlayerTheme](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/PlayerTheme.md)
 - [QualityObject](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/QualityObject.md)
 - [RemovePlayerThemesFromVideoRequest](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/RemovePlayerThemesFromVideoRequest.md)
 - [RenameAPIKeyRequest](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/RenameAPIKeyRequest.md)
 - [ResponseError](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/ResponseError.md)
 - [ResponseSuccess](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/ResponseSuccess.md)
 - [SetDefaultCaptionRequest](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/SetDefaultCaptionRequest.md)
 - [Theme](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Theme.md)
 - [UpdatePlayerThemeRequest](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/UpdatePlayerThemeRequest.md)
 - [UpdatePlayerThemeResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/UpdatePlayerThemeResponse.md)
 - [UpdateVideoInfoRequest](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/UpdateVideoInfoRequest.md)
 - [UpdateWebhookRequest](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/UpdateWebhookRequest.md)
 - [UploadLogoByIdResponse](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/UploadLogoByIdResponse.md)
 - [Video](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Video.md)
 - [VideoAssets](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/VideoAssets.md)
 - [VideoCaption](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/VideoCaption.md)
 - [VideoChapter](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/VideoChapter.md)
 - [VideoWatermark](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/VideoWatermark.md)
 - [Watermark](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Watermark.md)
 - [Webhook](https://github.com/AIOZNetwork/w3stream-go-client/blob/main/docs/Webhook.md)



## Have you gotten use from this API client?

Please take a moment to leave a star on the client ⭐

This helps other users to find the clients and also helps us understand which clients are most popular. Thank you!

