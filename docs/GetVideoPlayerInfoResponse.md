# GetVideoPlayerInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**VideoAssets**](VideoAssets.md) |  | [optional] 
**Captions** | Pointer to [**[]VideoCaption**](VideoCaption.md) |  | [optional] 
**Chapters** | Pointer to [**[]VideoChapter**](VideoChapter.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsMp4** | Pointer to **bool** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to [**[]Metadata**](Metadata.md) |  | [optional] 
**PlayerTheme** | Pointer to [**PlayerTheme**](PlayerTheme.md) |  | [optional] 
**PlayerThemeId** | Pointer to **string** |  | [optional] 
**Qualities** | Pointer to [**[]QualityObject**](QualityObject.md) |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewGetVideoPlayerInfoResponse

`func NewGetVideoPlayerInfoResponse() *GetVideoPlayerInfoResponse`

NewGetVideoPlayerInfoResponse instantiates a new GetVideoPlayerInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVideoPlayerInfoResponseWithDefaults

`func NewGetVideoPlayerInfoResponseWithDefaults() *GetVideoPlayerInfoResponse`

NewGetVideoPlayerInfoResponseWithDefaults instantiates a new GetVideoPlayerInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *GetVideoPlayerInfoResponse) GetAssets() VideoAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *GetVideoPlayerInfoResponse) GetAssetsOk() (*VideoAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *GetVideoPlayerInfoResponse) SetAssets(v VideoAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *GetVideoPlayerInfoResponse) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetCaptions

`func (o *GetVideoPlayerInfoResponse) GetCaptions() []VideoCaption`

GetCaptions returns the Captions field if non-nil, zero value otherwise.

### GetCaptionsOk

`func (o *GetVideoPlayerInfoResponse) GetCaptionsOk() (*[]VideoCaption, bool)`

GetCaptionsOk returns a tuple with the Captions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptions

`func (o *GetVideoPlayerInfoResponse) SetCaptions(v []VideoCaption)`

SetCaptions sets Captions field to given value.

### HasCaptions

`func (o *GetVideoPlayerInfoResponse) HasCaptions() bool`

HasCaptions returns a boolean if a field has been set.

### GetChapters

`func (o *GetVideoPlayerInfoResponse) GetChapters() []VideoChapter`

GetChapters returns the Chapters field if non-nil, zero value otherwise.

### GetChaptersOk

`func (o *GetVideoPlayerInfoResponse) GetChaptersOk() (*[]VideoChapter, bool)`

GetChaptersOk returns a tuple with the Chapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapters

`func (o *GetVideoPlayerInfoResponse) SetChapters(v []VideoChapter)`

SetChapters sets Chapters field to given value.

### HasChapters

`func (o *GetVideoPlayerInfoResponse) HasChapters() bool`

HasChapters returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetVideoPlayerInfoResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetVideoPlayerInfoResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetVideoPlayerInfoResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetVideoPlayerInfoResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *GetVideoPlayerInfoResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetVideoPlayerInfoResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetVideoPlayerInfoResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetVideoPlayerInfoResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *GetVideoPlayerInfoResponse) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetVideoPlayerInfoResponse) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetVideoPlayerInfoResponse) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetVideoPlayerInfoResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetId

`func (o *GetVideoPlayerInfoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetVideoPlayerInfoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetVideoPlayerInfoResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetVideoPlayerInfoResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsMp4

`func (o *GetVideoPlayerInfoResponse) GetIsMp4() bool`

GetIsMp4 returns the IsMp4 field if non-nil, zero value otherwise.

### GetIsMp4Ok

`func (o *GetVideoPlayerInfoResponse) GetIsMp4Ok() (*bool, bool)`

GetIsMp4Ok returns a tuple with the IsMp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMp4

`func (o *GetVideoPlayerInfoResponse) SetIsMp4(v bool)`

SetIsMp4 sets IsMp4 field to given value.

### HasIsMp4

`func (o *GetVideoPlayerInfoResponse) HasIsMp4() bool`

HasIsMp4 returns a boolean if a field has been set.

### GetIsPublic

`func (o *GetVideoPlayerInfoResponse) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *GetVideoPlayerInfoResponse) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *GetVideoPlayerInfoResponse) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *GetVideoPlayerInfoResponse) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetMetadata

`func (o *GetVideoPlayerInfoResponse) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetVideoPlayerInfoResponse) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetVideoPlayerInfoResponse) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetVideoPlayerInfoResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPlayerTheme

`func (o *GetVideoPlayerInfoResponse) GetPlayerTheme() PlayerTheme`

GetPlayerTheme returns the PlayerTheme field if non-nil, zero value otherwise.

### GetPlayerThemeOk

`func (o *GetVideoPlayerInfoResponse) GetPlayerThemeOk() (*PlayerTheme, bool)`

GetPlayerThemeOk returns a tuple with the PlayerTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerTheme

`func (o *GetVideoPlayerInfoResponse) SetPlayerTheme(v PlayerTheme)`

SetPlayerTheme sets PlayerTheme field to given value.

### HasPlayerTheme

`func (o *GetVideoPlayerInfoResponse) HasPlayerTheme() bool`

HasPlayerTheme returns a boolean if a field has been set.

### GetPlayerThemeId

`func (o *GetVideoPlayerInfoResponse) GetPlayerThemeId() string`

GetPlayerThemeId returns the PlayerThemeId field if non-nil, zero value otherwise.

### GetPlayerThemeIdOk

`func (o *GetVideoPlayerInfoResponse) GetPlayerThemeIdOk() (*string, bool)`

GetPlayerThemeIdOk returns a tuple with the PlayerThemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerThemeId

`func (o *GetVideoPlayerInfoResponse) SetPlayerThemeId(v string)`

SetPlayerThemeId sets PlayerThemeId field to given value.

### HasPlayerThemeId

`func (o *GetVideoPlayerInfoResponse) HasPlayerThemeId() bool`

HasPlayerThemeId returns a boolean if a field has been set.

### GetQualities

`func (o *GetVideoPlayerInfoResponse) GetQualities() []QualityObject`

GetQualities returns the Qualities field if non-nil, zero value otherwise.

### GetQualitiesOk

`func (o *GetVideoPlayerInfoResponse) GetQualitiesOk() (*[]QualityObject, bool)`

GetQualitiesOk returns a tuple with the Qualities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualities

`func (o *GetVideoPlayerInfoResponse) SetQualities(v []QualityObject)`

SetQualities sets Qualities field to given value.

### HasQualities

`func (o *GetVideoPlayerInfoResponse) HasQualities() bool`

HasQualities returns a boolean if a field has been set.

### GetSize

`func (o *GetVideoPlayerInfoResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetVideoPlayerInfoResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetVideoPlayerInfoResponse) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetVideoPlayerInfoResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *GetVideoPlayerInfoResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVideoPlayerInfoResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVideoPlayerInfoResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetVideoPlayerInfoResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *GetVideoPlayerInfoResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetVideoPlayerInfoResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetVideoPlayerInfoResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetVideoPlayerInfoResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *GetVideoPlayerInfoResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetVideoPlayerInfoResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetVideoPlayerInfoResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetVideoPlayerInfoResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetVideoPlayerInfoResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetVideoPlayerInfoResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetVideoPlayerInfoResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetVideoPlayerInfoResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *GetVideoPlayerInfoResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetVideoPlayerInfoResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetVideoPlayerInfoResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GetVideoPlayerInfoResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


