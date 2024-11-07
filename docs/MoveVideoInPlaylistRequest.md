# MoveVideoInPlaylistRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentId** | Pointer to **string** | CurrentId is the UUID of the playlist item (video) to be moved | [optional] 
**NextId** | Pointer to **string** | NextId is the UUID of the playlist item that should come after the moved item | [optional] 
**PreviousId** | Pointer to **string** | PreviousId is the UUID of the playlist item that should come before the moved item | [optional] 

## Methods

### NewMoveVideoInPlaylistRequest

`func NewMoveVideoInPlaylistRequest() *MoveVideoInPlaylistRequest`

NewMoveVideoInPlaylistRequest instantiates a new MoveVideoInPlaylistRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveVideoInPlaylistRequestWithDefaults

`func NewMoveVideoInPlaylistRequestWithDefaults() *MoveVideoInPlaylistRequest`

NewMoveVideoInPlaylistRequestWithDefaults instantiates a new MoveVideoInPlaylistRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentId

`func (o *MoveVideoInPlaylistRequest) GetCurrentId() string`

GetCurrentId returns the CurrentId field if non-nil, zero value otherwise.

### GetCurrentIdOk

`func (o *MoveVideoInPlaylistRequest) GetCurrentIdOk() (*string, bool)`

GetCurrentIdOk returns a tuple with the CurrentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentId

`func (o *MoveVideoInPlaylistRequest) SetCurrentId(v string)`

SetCurrentId sets CurrentId field to given value.

### HasCurrentId

`func (o *MoveVideoInPlaylistRequest) HasCurrentId() bool`

HasCurrentId returns a boolean if a field has been set.

### GetNextId

`func (o *MoveVideoInPlaylistRequest) GetNextId() string`

GetNextId returns the NextId field if non-nil, zero value otherwise.

### GetNextIdOk

`func (o *MoveVideoInPlaylistRequest) GetNextIdOk() (*string, bool)`

GetNextIdOk returns a tuple with the NextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextId

`func (o *MoveVideoInPlaylistRequest) SetNextId(v string)`

SetNextId sets NextId field to given value.

### HasNextId

`func (o *MoveVideoInPlaylistRequest) HasNextId() bool`

HasNextId returns a boolean if a field has been set.

### GetPreviousId

`func (o *MoveVideoInPlaylistRequest) GetPreviousId() string`

GetPreviousId returns the PreviousId field if non-nil, zero value otherwise.

### GetPreviousIdOk

`func (o *MoveVideoInPlaylistRequest) GetPreviousIdOk() (*string, bool)`

GetPreviousIdOk returns a tuple with the PreviousId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousId

`func (o *MoveVideoInPlaylistRequest) SetPreviousId(v string)`

SetPreviousId sets PreviousId field to given value.

### HasPreviousId

`func (o *MoveVideoInPlaylistRequest) HasPreviousId() bool`

HasPreviousId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


