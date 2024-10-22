# GetLiveStreamKeysListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LiveStreamKeys** | Pointer to [**[]GetLiveStreamKeyData**](GetLiveStreamKeyData.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetLiveStreamKeysListData

`func NewGetLiveStreamKeysListData() *GetLiveStreamKeysListData`

NewGetLiveStreamKeysListData instantiates a new GetLiveStreamKeysListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLiveStreamKeysListDataWithDefaults

`func NewGetLiveStreamKeysListDataWithDefaults() *GetLiveStreamKeysListData`

NewGetLiveStreamKeysListDataWithDefaults instantiates a new GetLiveStreamKeysListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLiveStreamKeys

`func (o *GetLiveStreamKeysListData) GetLiveStreamKeys() []GetLiveStreamKeyData`

GetLiveStreamKeys returns the LiveStreamKeys field if non-nil, zero value otherwise.

### GetLiveStreamKeysOk

`func (o *GetLiveStreamKeysListData) GetLiveStreamKeysOk() (*[]GetLiveStreamKeyData, bool)`

GetLiveStreamKeysOk returns a tuple with the LiveStreamKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamKeys

`func (o *GetLiveStreamKeysListData) SetLiveStreamKeys(v []GetLiveStreamKeyData)`

SetLiveStreamKeys sets LiveStreamKeys field to given value.

### HasLiveStreamKeys

`func (o *GetLiveStreamKeysListData) HasLiveStreamKeys() bool`

HasLiveStreamKeys returns a boolean if a field has been set.

### GetTotal

`func (o *GetLiveStreamKeysListData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetLiveStreamKeysListData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetLiveStreamKeysListData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetLiveStreamKeysListData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


