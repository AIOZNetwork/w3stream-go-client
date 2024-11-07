# GetPlaylistListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GetPlaylistListData**](GetPlaylistListData.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetPlaylistListResponse

`func NewGetPlaylistListResponse() *GetPlaylistListResponse`

NewGetPlaylistListResponse instantiates a new GetPlaylistListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPlaylistListResponseWithDefaults

`func NewGetPlaylistListResponseWithDefaults() *GetPlaylistListResponse`

NewGetPlaylistListResponseWithDefaults instantiates a new GetPlaylistListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetPlaylistListResponse) GetData() GetPlaylistListData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPlaylistListResponse) GetDataOk() (*GetPlaylistListData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPlaylistListResponse) SetData(v GetPlaylistListData)`

SetData sets Data field to given value.

### HasData

`func (o *GetPlaylistListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *GetPlaylistListResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPlaylistListResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPlaylistListResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetPlaylistListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


