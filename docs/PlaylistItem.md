# PlaylistItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Next** | Pointer to [**PlaylistItem**](PlaylistItem.md) |  | [optional] 
**NextId** | Pointer to **string** |  | [optional] 
**PlaylistId** | Pointer to **string** |  | [optional] 
**Previous** | Pointer to [**PlaylistItem**](PlaylistItem.md) |  | [optional] 
**PreviousId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Video** | Pointer to [**PlaylistItemVideo**](PlaylistItemVideo.md) |  | [optional] 
**VideoId** | Pointer to **string** |  | [optional] 

## Methods

### NewPlaylistItem

`func NewPlaylistItem() *PlaylistItem`

NewPlaylistItem instantiates a new PlaylistItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistItemWithDefaults

`func NewPlaylistItemWithDefaults() *PlaylistItem`

NewPlaylistItemWithDefaults instantiates a new PlaylistItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PlaylistItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlaylistItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlaylistItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PlaylistItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *PlaylistItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlaylistItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlaylistItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlaylistItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNext

`func (o *PlaylistItem) GetNext() PlaylistItem`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PlaylistItem) GetNextOk() (*PlaylistItem, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PlaylistItem) SetNext(v PlaylistItem)`

SetNext sets Next field to given value.

### HasNext

`func (o *PlaylistItem) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetNextId

`func (o *PlaylistItem) GetNextId() string`

GetNextId returns the NextId field if non-nil, zero value otherwise.

### GetNextIdOk

`func (o *PlaylistItem) GetNextIdOk() (*string, bool)`

GetNextIdOk returns a tuple with the NextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextId

`func (o *PlaylistItem) SetNextId(v string)`

SetNextId sets NextId field to given value.

### HasNextId

`func (o *PlaylistItem) HasNextId() bool`

HasNextId returns a boolean if a field has been set.

### GetPlaylistId

`func (o *PlaylistItem) GetPlaylistId() string`

GetPlaylistId returns the PlaylistId field if non-nil, zero value otherwise.

### GetPlaylistIdOk

`func (o *PlaylistItem) GetPlaylistIdOk() (*string, bool)`

GetPlaylistIdOk returns a tuple with the PlaylistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistId

`func (o *PlaylistItem) SetPlaylistId(v string)`

SetPlaylistId sets PlaylistId field to given value.

### HasPlaylistId

`func (o *PlaylistItem) HasPlaylistId() bool`

HasPlaylistId returns a boolean if a field has been set.

### GetPrevious

`func (o *PlaylistItem) GetPrevious() PlaylistItem`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PlaylistItem) GetPreviousOk() (*PlaylistItem, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PlaylistItem) SetPrevious(v PlaylistItem)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PlaylistItem) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetPreviousId

`func (o *PlaylistItem) GetPreviousId() string`

GetPreviousId returns the PreviousId field if non-nil, zero value otherwise.

### GetPreviousIdOk

`func (o *PlaylistItem) GetPreviousIdOk() (*string, bool)`

GetPreviousIdOk returns a tuple with the PreviousId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousId

`func (o *PlaylistItem) SetPreviousId(v string)`

SetPreviousId sets PreviousId field to given value.

### HasPreviousId

`func (o *PlaylistItem) HasPreviousId() bool`

HasPreviousId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PlaylistItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlaylistItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlaylistItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PlaylistItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVideo

`func (o *PlaylistItem) GetVideo() PlaylistItemVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *PlaylistItem) GetVideoOk() (*PlaylistItemVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *PlaylistItem) SetVideo(v PlaylistItemVideo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *PlaylistItem) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetVideoId

`func (o *PlaylistItem) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *PlaylistItem) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *PlaylistItem) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *PlaylistItem) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


