# CreateVideoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**VideoObject**](VideoObject.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateVideoResponse

`func NewCreateVideoResponse() *CreateVideoResponse`

NewCreateVideoResponse instantiates a new CreateVideoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVideoResponseWithDefaults

`func NewCreateVideoResponseWithDefaults() *CreateVideoResponse`

NewCreateVideoResponseWithDefaults instantiates a new CreateVideoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateVideoResponse) GetData() VideoObject`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateVideoResponse) GetDataOk() (*VideoObject, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateVideoResponse) SetData(v VideoObject)`

SetData sets Data field to given value.

### HasData

`func (o *CreateVideoResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *CreateVideoResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateVideoResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateVideoResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateVideoResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


