# CreateLiveStreamKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**LiveStreamKeyData**](LiveStreamKeyData.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateLiveStreamKeyResponse

`func NewCreateLiveStreamKeyResponse() *CreateLiveStreamKeyResponse`

NewCreateLiveStreamKeyResponse instantiates a new CreateLiveStreamKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLiveStreamKeyResponseWithDefaults

`func NewCreateLiveStreamKeyResponseWithDefaults() *CreateLiveStreamKeyResponse`

NewCreateLiveStreamKeyResponseWithDefaults instantiates a new CreateLiveStreamKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateLiveStreamKeyResponse) GetData() LiveStreamKeyData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateLiveStreamKeyResponse) GetDataOk() (*LiveStreamKeyData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateLiveStreamKeyResponse) SetData(v LiveStreamKeyData)`

SetData sets Data field to given value.

### HasData

`func (o *CreateLiveStreamKeyResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *CreateLiveStreamKeyResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateLiveStreamKeyResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateLiveStreamKeyResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateLiveStreamKeyResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


