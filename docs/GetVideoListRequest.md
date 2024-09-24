# GetVideoListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to [**[]Metadata**](Metadata.md) |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**OrderBy** | Pointer to **string** |  | [optional] 
**Search** | Pointer to **string** |  | [optional] 
**SortBy** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetVideoListRequest

`func NewGetVideoListRequest() *GetVideoListRequest`

NewGetVideoListRequest instantiates a new GetVideoListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVideoListRequestWithDefaults

`func NewGetVideoListRequestWithDefaults() *GetVideoListRequest`

NewGetVideoListRequestWithDefaults instantiates a new GetVideoListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *GetVideoListRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetVideoListRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetVideoListRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetVideoListRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMetadata

`func (o *GetVideoListRequest) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetVideoListRequest) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetVideoListRequest) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetVideoListRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOffset

`func (o *GetVideoListRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetVideoListRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetVideoListRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetVideoListRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderBy

`func (o *GetVideoListRequest) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *GetVideoListRequest) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *GetVideoListRequest) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *GetVideoListRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetSearch

`func (o *GetVideoListRequest) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *GetVideoListRequest) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *GetVideoListRequest) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *GetVideoListRequest) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetSortBy

`func (o *GetVideoListRequest) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *GetVideoListRequest) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *GetVideoListRequest) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *GetVideoListRequest) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetStatus

`func (o *GetVideoListRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVideoListRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVideoListRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetVideoListRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *GetVideoListRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetVideoListRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetVideoListRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetVideoListRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


