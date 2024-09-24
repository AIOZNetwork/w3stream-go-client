# GetTranscodeCostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GetTranscodeCostData**](GetTranscodeCostData.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetTranscodeCostResponse

`func NewGetTranscodeCostResponse() *GetTranscodeCostResponse`

NewGetTranscodeCostResponse instantiates a new GetTranscodeCostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTranscodeCostResponseWithDefaults

`func NewGetTranscodeCostResponseWithDefaults() *GetTranscodeCostResponse`

NewGetTranscodeCostResponseWithDefaults instantiates a new GetTranscodeCostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTranscodeCostResponse) GetData() GetTranscodeCostData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTranscodeCostResponse) GetDataOk() (*GetTranscodeCostData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTranscodeCostResponse) SetData(v GetTranscodeCostData)`

SetData sets Data field to given value.

### HasData

`func (o *GetTranscodeCostResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *GetTranscodeCostResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTranscodeCostResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTranscodeCostResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTranscodeCostResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


