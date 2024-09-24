# GetPlayerThemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GetPlayerThemeData**](GetPlayerThemeData.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetPlayerThemeResponse

`func NewGetPlayerThemeResponse() *GetPlayerThemeResponse`

NewGetPlayerThemeResponse instantiates a new GetPlayerThemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPlayerThemeResponseWithDefaults

`func NewGetPlayerThemeResponseWithDefaults() *GetPlayerThemeResponse`

NewGetPlayerThemeResponseWithDefaults instantiates a new GetPlayerThemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetPlayerThemeResponse) GetData() GetPlayerThemeData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPlayerThemeResponse) GetDataOk() (*GetPlayerThemeData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPlayerThemeResponse) SetData(v GetPlayerThemeData)`

SetData sets Data field to given value.

### HasData

`func (o *GetPlayerThemeResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *GetPlayerThemeResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPlayerThemeResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPlayerThemeResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetPlayerThemeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


