# GetVideoCaptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GetVideoCaptionsData**](GetVideoCaptionsData.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetVideoCaptionsResponse

`func NewGetVideoCaptionsResponse() *GetVideoCaptionsResponse`

NewGetVideoCaptionsResponse instantiates a new GetVideoCaptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVideoCaptionsResponseWithDefaults

`func NewGetVideoCaptionsResponseWithDefaults() *GetVideoCaptionsResponse`

NewGetVideoCaptionsResponseWithDefaults instantiates a new GetVideoCaptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetVideoCaptionsResponse) GetData() GetVideoCaptionsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetVideoCaptionsResponse) GetDataOk() (*GetVideoCaptionsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetVideoCaptionsResponse) SetData(v GetVideoCaptionsData)`

SetData sets Data field to given value.

### HasData

`func (o *GetVideoCaptionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *GetVideoCaptionsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVideoCaptionsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVideoCaptionsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetVideoCaptionsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


