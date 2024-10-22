# LiveStreamVideoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**LiveStreamAssets**](LiveStreamAssets.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LiveStreamKeyId** | Pointer to **string** |  | [optional] 
**Qualities** | Pointer to **[]string** |  | [optional] 
**Save** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Video** | Pointer to [**VideoObject**](VideoObject.md) |  | [optional] 

## Methods

### NewLiveStreamVideoResponse

`func NewLiveStreamVideoResponse() *LiveStreamVideoResponse`

NewLiveStreamVideoResponse instantiates a new LiveStreamVideoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamVideoResponseWithDefaults

`func NewLiveStreamVideoResponseWithDefaults() *LiveStreamVideoResponse`

NewLiveStreamVideoResponseWithDefaults instantiates a new LiveStreamVideoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *LiveStreamVideoResponse) GetAssets() LiveStreamAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *LiveStreamVideoResponse) GetAssetsOk() (*LiveStreamAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *LiveStreamVideoResponse) SetAssets(v LiveStreamAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *LiveStreamVideoResponse) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LiveStreamVideoResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LiveStreamVideoResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LiveStreamVideoResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LiveStreamVideoResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDuration

`func (o *LiveStreamVideoResponse) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *LiveStreamVideoResponse) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *LiveStreamVideoResponse) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *LiveStreamVideoResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetId

`func (o *LiveStreamVideoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LiveStreamVideoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LiveStreamVideoResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LiveStreamVideoResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLiveStreamKeyId

`func (o *LiveStreamVideoResponse) GetLiveStreamKeyId() string`

GetLiveStreamKeyId returns the LiveStreamKeyId field if non-nil, zero value otherwise.

### GetLiveStreamKeyIdOk

`func (o *LiveStreamVideoResponse) GetLiveStreamKeyIdOk() (*string, bool)`

GetLiveStreamKeyIdOk returns a tuple with the LiveStreamKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamKeyId

`func (o *LiveStreamVideoResponse) SetLiveStreamKeyId(v string)`

SetLiveStreamKeyId sets LiveStreamKeyId field to given value.

### HasLiveStreamKeyId

`func (o *LiveStreamVideoResponse) HasLiveStreamKeyId() bool`

HasLiveStreamKeyId returns a boolean if a field has been set.

### GetQualities

`func (o *LiveStreamVideoResponse) GetQualities() []string`

GetQualities returns the Qualities field if non-nil, zero value otherwise.

### GetQualitiesOk

`func (o *LiveStreamVideoResponse) GetQualitiesOk() (*[]string, bool)`

GetQualitiesOk returns a tuple with the Qualities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualities

`func (o *LiveStreamVideoResponse) SetQualities(v []string)`

SetQualities sets Qualities field to given value.

### HasQualities

`func (o *LiveStreamVideoResponse) HasQualities() bool`

HasQualities returns a boolean if a field has been set.

### GetSave

`func (o *LiveStreamVideoResponse) GetSave() bool`

GetSave returns the Save field if non-nil, zero value otherwise.

### GetSaveOk

`func (o *LiveStreamVideoResponse) GetSaveOk() (*bool, bool)`

GetSaveOk returns a tuple with the Save field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSave

`func (o *LiveStreamVideoResponse) SetSave(v bool)`

SetSave sets Save field to given value.

### HasSave

`func (o *LiveStreamVideoResponse) HasSave() bool`

HasSave returns a boolean if a field has been set.

### GetStatus

`func (o *LiveStreamVideoResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LiveStreamVideoResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LiveStreamVideoResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LiveStreamVideoResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *LiveStreamVideoResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LiveStreamVideoResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LiveStreamVideoResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LiveStreamVideoResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LiveStreamVideoResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LiveStreamVideoResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LiveStreamVideoResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LiveStreamVideoResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *LiveStreamVideoResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *LiveStreamVideoResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *LiveStreamVideoResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *LiveStreamVideoResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVideo

`func (o *LiveStreamVideoResponse) GetVideo() VideoObject`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *LiveStreamVideoResponse) GetVideoOk() (*VideoObject, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *LiveStreamVideoResponse) SetVideo(v VideoObject)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *LiveStreamVideoResponse) HasVideo() bool`

HasVideo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


