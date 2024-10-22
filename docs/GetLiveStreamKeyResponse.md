# GetLiveStreamKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GetLiveStreamKeyData**](GetLiveStreamKeyData.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetLiveStreamKeyResponse

`func NewGetLiveStreamKeyResponse() *GetLiveStreamKeyResponse`

NewGetLiveStreamKeyResponse instantiates a new GetLiveStreamKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLiveStreamKeyResponseWithDefaults

`func NewGetLiveStreamKeyResponseWithDefaults() *GetLiveStreamKeyResponse`

NewGetLiveStreamKeyResponseWithDefaults instantiates a new GetLiveStreamKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetLiveStreamKeyResponse) GetData() GetLiveStreamKeyData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetLiveStreamKeyResponse) GetDataOk() (*GetLiveStreamKeyData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetLiveStreamKeyResponse) SetData(v GetLiveStreamKeyData)`

SetData sets Data field to given value.

### HasData

`func (o *GetLiveStreamKeyResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *GetLiveStreamKeyResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetLiveStreamKeyResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetLiveStreamKeyResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetLiveStreamKeyResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


