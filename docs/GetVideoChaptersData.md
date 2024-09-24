# GetVideoChaptersData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**VideoChapters** | Pointer to [**[]VideoChapter**](VideoChapter.md) |  | [optional] 

## Methods

### NewGetVideoChaptersData

`func NewGetVideoChaptersData() *GetVideoChaptersData`

NewGetVideoChaptersData instantiates a new GetVideoChaptersData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVideoChaptersDataWithDefaults

`func NewGetVideoChaptersDataWithDefaults() *GetVideoChaptersData`

NewGetVideoChaptersDataWithDefaults instantiates a new GetVideoChaptersData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetVideoChaptersData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetVideoChaptersData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetVideoChaptersData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetVideoChaptersData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetVideoChapters

`func (o *GetVideoChaptersData) GetVideoChapters() []VideoChapter`

GetVideoChapters returns the VideoChapters field if non-nil, zero value otherwise.

### GetVideoChaptersOk

`func (o *GetVideoChaptersData) GetVideoChaptersOk() (*[]VideoChapter, bool)`

GetVideoChaptersOk returns a tuple with the VideoChapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChapters

`func (o *GetVideoChaptersData) SetVideoChapters(v []VideoChapter)`

SetVideoChapters sets VideoChapters field to given value.

### HasVideoChapters

`func (o *GetVideoChaptersData) HasVideoChapters() bool`

HasVideoChapters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


