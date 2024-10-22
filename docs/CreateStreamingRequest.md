# CreateStreamingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qualities** | Pointer to **[]string** | Qualities of the video (default: 1080p, 720p,  360p, allow:2160p, 1440p, 1080p, 720p,  360p, 240p, 144p) | [optional] 
**Save** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateStreamingRequest

`func NewCreateStreamingRequest() *CreateStreamingRequest`

NewCreateStreamingRequest instantiates a new CreateStreamingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStreamingRequestWithDefaults

`func NewCreateStreamingRequestWithDefaults() *CreateStreamingRequest`

NewCreateStreamingRequestWithDefaults instantiates a new CreateStreamingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQualities

`func (o *CreateStreamingRequest) GetQualities() []string`

GetQualities returns the Qualities field if non-nil, zero value otherwise.

### GetQualitiesOk

`func (o *CreateStreamingRequest) GetQualitiesOk() (*[]string, bool)`

GetQualitiesOk returns a tuple with the Qualities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualities

`func (o *CreateStreamingRequest) SetQualities(v []string)`

SetQualities sets Qualities field to given value.

### HasQualities

`func (o *CreateStreamingRequest) HasQualities() bool`

HasQualities returns a boolean if a field has been set.

### GetSave

`func (o *CreateStreamingRequest) GetSave() bool`

GetSave returns the Save field if non-nil, zero value otherwise.

### GetSaveOk

`func (o *CreateStreamingRequest) GetSaveOk() (*bool, bool)`

GetSaveOk returns a tuple with the Save field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSave

`func (o *CreateStreamingRequest) SetSave(v bool)`

SetSave sets Save field to given value.

### HasSave

`func (o *CreateStreamingRequest) HasSave() bool`

HasSave returns a boolean if a field has been set.

### GetTitle

`func (o *CreateStreamingRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateStreamingRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateStreamingRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateStreamingRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


