# LiveStreamVideosResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Videos** | Pointer to [**[]LiveStreamVideoResponse**](LiveStreamVideoResponse.md) |  | [optional] 

## Methods

### NewLiveStreamVideosResponse

`func NewLiveStreamVideosResponse() *LiveStreamVideosResponse`

NewLiveStreamVideosResponse instantiates a new LiveStreamVideosResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamVideosResponseWithDefaults

`func NewLiveStreamVideosResponseWithDefaults() *LiveStreamVideosResponse`

NewLiveStreamVideosResponseWithDefaults instantiates a new LiveStreamVideosResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *LiveStreamVideosResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LiveStreamVideosResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *LiveStreamVideosResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *LiveStreamVideosResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetVideos

`func (o *LiveStreamVideosResponse) GetVideos() []LiveStreamVideoResponse`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *LiveStreamVideosResponse) GetVideosOk() (*[]LiveStreamVideoResponse, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *LiveStreamVideosResponse) SetVideos(v []LiveStreamVideoResponse)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *LiveStreamVideosResponse) HasVideos() bool`

HasVideos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


