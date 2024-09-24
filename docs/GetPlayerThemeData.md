# GetPlayerThemeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayerThemes** | Pointer to [**[]PlayerTheme**](PlayerTheme.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetPlayerThemeData

`func NewGetPlayerThemeData() *GetPlayerThemeData`

NewGetPlayerThemeData instantiates a new GetPlayerThemeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPlayerThemeDataWithDefaults

`func NewGetPlayerThemeDataWithDefaults() *GetPlayerThemeData`

NewGetPlayerThemeDataWithDefaults instantiates a new GetPlayerThemeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayerThemes

`func (o *GetPlayerThemeData) GetPlayerThemes() []PlayerTheme`

GetPlayerThemes returns the PlayerThemes field if non-nil, zero value otherwise.

### GetPlayerThemesOk

`func (o *GetPlayerThemeData) GetPlayerThemesOk() (*[]PlayerTheme, bool)`

GetPlayerThemesOk returns a tuple with the PlayerThemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerThemes

`func (o *GetPlayerThemeData) SetPlayerThemes(v []PlayerTheme)`

SetPlayerThemes sets PlayerThemes field to given value.

### HasPlayerThemes

`func (o *GetPlayerThemeData) HasPlayerThemes() bool`

HasPlayerThemes returns a boolean if a field has been set.

### GetTotal

`func (o *GetPlayerThemeData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetPlayerThemeData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetPlayerThemeData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetPlayerThemeData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


