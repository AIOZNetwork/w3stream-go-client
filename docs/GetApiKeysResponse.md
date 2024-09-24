# GetApiKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GetApiKeysData**](GetApiKeysData.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetApiKeysResponse

`func NewGetApiKeysResponse() *GetApiKeysResponse`

NewGetApiKeysResponse instantiates a new GetApiKeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiKeysResponseWithDefaults

`func NewGetApiKeysResponseWithDefaults() *GetApiKeysResponse`

NewGetApiKeysResponseWithDefaults instantiates a new GetApiKeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetApiKeysResponse) GetData() GetApiKeysData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetApiKeysResponse) GetDataOk() (*GetApiKeysData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetApiKeysResponse) SetData(v GetApiKeysData)`

SetData sets Data field to given value.

### HasData

`func (o *GetApiKeysResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *GetApiKeysResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetApiKeysResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetApiKeysResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetApiKeysResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


