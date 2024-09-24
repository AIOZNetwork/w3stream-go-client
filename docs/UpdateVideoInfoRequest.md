# UpdateVideoInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the video | [optional] 
**IsPublic** | Pointer to **bool** | Video&#39;s publish status | [optional] 
**Metadata** | Pointer to [**[]Metadata**](Metadata.md) | Video&#39;s metadata | [optional] 
**PlayerId** | Pointer to **string** | Video player &#39;s id | [optional] 
**Tags** | Pointer to **[]string** | Video&#39;s tags | [optional] 
**Title** | Pointer to **string** | Title of the video | [optional] 

## Methods

### NewUpdateVideoInfoRequest

`func NewUpdateVideoInfoRequest() *UpdateVideoInfoRequest`

NewUpdateVideoInfoRequest instantiates a new UpdateVideoInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVideoInfoRequestWithDefaults

`func NewUpdateVideoInfoRequestWithDefaults() *UpdateVideoInfoRequest`

NewUpdateVideoInfoRequestWithDefaults instantiates a new UpdateVideoInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateVideoInfoRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateVideoInfoRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateVideoInfoRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateVideoInfoRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPublic

`func (o *UpdateVideoInfoRequest) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *UpdateVideoInfoRequest) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *UpdateVideoInfoRequest) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *UpdateVideoInfoRequest) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateVideoInfoRequest) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateVideoInfoRequest) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateVideoInfoRequest) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateVideoInfoRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPlayerId

`func (o *UpdateVideoInfoRequest) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *UpdateVideoInfoRequest) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *UpdateVideoInfoRequest) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.

### HasPlayerId

`func (o *UpdateVideoInfoRequest) HasPlayerId() bool`

HasPlayerId returns a boolean if a field has been set.

### GetTags

`func (o *UpdateVideoInfoRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateVideoInfoRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateVideoInfoRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateVideoInfoRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *UpdateVideoInfoRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateVideoInfoRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateVideoInfoRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateVideoInfoRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


