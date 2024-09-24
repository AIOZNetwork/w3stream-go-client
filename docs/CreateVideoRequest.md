# CreateVideoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the video | [optional] 
**IsPublic** | Pointer to **bool** | // Is panoramic video IsPanoramic *bool &#x60;json:\&quot;is_panoramic\&quot; form:\&quot;is_panoramic\&quot;&#x60; Is public video | [optional] 
**Metadata** | Pointer to [**[]Metadata**](Metadata.md) | Metadata of the video (key-value pair, max: 50 items, key max length: 255, value max length: 255) | [optional] 
**Qualities** | Pointer to **[]string** | Qualities of the video (default: 1080p, 720p,  360p, allow:2160p, 1440p, 1080p, 720p,  360p, 240p, 144p) | [optional] 
**Tags** | Pointer to **[]string** | Tags of the video (max: 50 items, max length: 255) | [optional] 
**Title** | Pointer to **string** | Title of the video | [optional] 
**Watermark** | Pointer to [**VideoWatermark**](VideoWatermark.md) | Video thumbnailConfig | [optional] 

## Methods

### NewCreateVideoRequest

`func NewCreateVideoRequest() *CreateVideoRequest`

NewCreateVideoRequest instantiates a new CreateVideoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVideoRequestWithDefaults

`func NewCreateVideoRequestWithDefaults() *CreateVideoRequest`

NewCreateVideoRequestWithDefaults instantiates a new CreateVideoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateVideoRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVideoRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVideoRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVideoRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPublic

`func (o *CreateVideoRequest) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *CreateVideoRequest) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *CreateVideoRequest) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *CreateVideoRequest) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateVideoRequest) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateVideoRequest) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateVideoRequest) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateVideoRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetQualities

`func (o *CreateVideoRequest) GetQualities() []string`

GetQualities returns the Qualities field if non-nil, zero value otherwise.

### GetQualitiesOk

`func (o *CreateVideoRequest) GetQualitiesOk() (*[]string, bool)`

GetQualitiesOk returns a tuple with the Qualities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualities

`func (o *CreateVideoRequest) SetQualities(v []string)`

SetQualities sets Qualities field to given value.

### HasQualities

`func (o *CreateVideoRequest) HasQualities() bool`

HasQualities returns a boolean if a field has been set.

### GetTags

`func (o *CreateVideoRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateVideoRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateVideoRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateVideoRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *CreateVideoRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateVideoRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateVideoRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateVideoRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWatermark

`func (o *CreateVideoRequest) GetWatermark() VideoWatermark`

GetWatermark returns the Watermark field if non-nil, zero value otherwise.

### GetWatermarkOk

`func (o *CreateVideoRequest) GetWatermarkOk() (*VideoWatermark, bool)`

GetWatermarkOk returns a tuple with the Watermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermark

`func (o *CreateVideoRequest) SetWatermark(v VideoWatermark)`

SetWatermark sets Watermark field to given value.

### HasWatermark

`func (o *CreateVideoRequest) HasWatermark() bool`

HasWatermark returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


