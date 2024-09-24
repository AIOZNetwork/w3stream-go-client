# GetApiKeysData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeys** | Pointer to [**[]ApiKey**](ApiKey.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetApiKeysData

`func NewGetApiKeysData() *GetApiKeysData`

NewGetApiKeysData instantiates a new GetApiKeysData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiKeysDataWithDefaults

`func NewGetApiKeysDataWithDefaults() *GetApiKeysData`

NewGetApiKeysDataWithDefaults instantiates a new GetApiKeysData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeys

`func (o *GetApiKeysData) GetApiKeys() []ApiKey`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *GetApiKeysData) GetApiKeysOk() (*[]ApiKey, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *GetApiKeysData) SetApiKeys(v []ApiKey)`

SetApiKeys sets ApiKeys field to given value.

### HasApiKeys

`func (o *GetApiKeysData) HasApiKeys() bool`

HasApiKeys returns a boolean if a field has been set.

### GetTotal

`func (o *GetApiKeysData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetApiKeysData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetApiKeysData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetApiKeysData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


