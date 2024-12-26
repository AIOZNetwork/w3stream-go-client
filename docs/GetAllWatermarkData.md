# GetAllWatermarkData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Watermarks** | Pointer to [**[]Watermark**](Watermark.md) |  | [optional] 

## Methods

### NewGetAllWatermarkData

`func NewGetAllWatermarkData() *GetAllWatermarkData`

NewGetAllWatermarkData instantiates a new GetAllWatermarkData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllWatermarkDataWithDefaults

`func NewGetAllWatermarkDataWithDefaults() *GetAllWatermarkData`

NewGetAllWatermarkDataWithDefaults instantiates a new GetAllWatermarkData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetAllWatermarkData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetAllWatermarkData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetAllWatermarkData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetAllWatermarkData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetWatermarks

`func (o *GetAllWatermarkData) GetWatermarks() []Watermark`

GetWatermarks returns the Watermarks field if non-nil, zero value otherwise.

### GetWatermarksOk

`func (o *GetAllWatermarkData) GetWatermarksOk() (*[]Watermark, bool)`

GetWatermarksOk returns a tuple with the Watermarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarks

`func (o *GetAllWatermarkData) SetWatermarks(v []Watermark)`

SetWatermarks sets Watermarks field to given value.

### HasWatermarks

`func (o *GetAllWatermarkData) HasWatermarks() bool`

HasWatermarks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


