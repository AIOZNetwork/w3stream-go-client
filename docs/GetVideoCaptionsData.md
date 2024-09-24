# GetVideoCaptionsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**VideoCaptions** | Pointer to [**[]VideoCaption**](VideoCaption.md) |  | [optional] 

## Methods

### NewGetVideoCaptionsData

`func NewGetVideoCaptionsData() *GetVideoCaptionsData`

NewGetVideoCaptionsData instantiates a new GetVideoCaptionsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVideoCaptionsDataWithDefaults

`func NewGetVideoCaptionsDataWithDefaults() *GetVideoCaptionsData`

NewGetVideoCaptionsDataWithDefaults instantiates a new GetVideoCaptionsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetVideoCaptionsData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetVideoCaptionsData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetVideoCaptionsData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetVideoCaptionsData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetVideoCaptions

`func (o *GetVideoCaptionsData) GetVideoCaptions() []VideoCaption`

GetVideoCaptions returns the VideoCaptions field if non-nil, zero value otherwise.

### GetVideoCaptionsOk

`func (o *GetVideoCaptionsData) GetVideoCaptionsOk() (*[]VideoCaption, bool)`

GetVideoCaptionsOk returns a tuple with the VideoCaptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCaptions

`func (o *GetVideoCaptionsData) SetVideoCaptions(v []VideoCaption)`

SetVideoCaptions sets VideoCaptions field to given value.

### HasVideoCaptions

`func (o *GetVideoCaptionsData) HasVideoCaptions() bool`

HasVideoCaptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


