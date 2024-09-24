# UpdatePlayerThemeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to [**Asset**](Asset.md) |  | [optional] 
**Controls** | Pointer to [**Controls**](Controls.md) |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Theme** | Pointer to [**Theme**](Theme.md) |  | [optional] 

## Methods

### NewUpdatePlayerThemeRequest

`func NewUpdatePlayerThemeRequest() *UpdatePlayerThemeRequest`

NewUpdatePlayerThemeRequest instantiates a new UpdatePlayerThemeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePlayerThemeRequestWithDefaults

`func NewUpdatePlayerThemeRequestWithDefaults() *UpdatePlayerThemeRequest`

NewUpdatePlayerThemeRequestWithDefaults instantiates a new UpdatePlayerThemeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *UpdatePlayerThemeRequest) GetAsset() Asset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *UpdatePlayerThemeRequest) GetAssetOk() (*Asset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *UpdatePlayerThemeRequest) SetAsset(v Asset)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *UpdatePlayerThemeRequest) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetControls

`func (o *UpdatePlayerThemeRequest) GetControls() Controls`

GetControls returns the Controls field if non-nil, zero value otherwise.

### GetControlsOk

`func (o *UpdatePlayerThemeRequest) GetControlsOk() (*Controls, bool)`

GetControlsOk returns a tuple with the Controls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControls

`func (o *UpdatePlayerThemeRequest) SetControls(v Controls)`

SetControls sets Controls field to given value.

### HasControls

`func (o *UpdatePlayerThemeRequest) HasControls() bool`

HasControls returns a boolean if a field has been set.

### GetIsDefault

`func (o *UpdatePlayerThemeRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *UpdatePlayerThemeRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *UpdatePlayerThemeRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *UpdatePlayerThemeRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *UpdatePlayerThemeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePlayerThemeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePlayerThemeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePlayerThemeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTheme

`func (o *UpdatePlayerThemeRequest) GetTheme() Theme`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *UpdatePlayerThemeRequest) GetThemeOk() (*Theme, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *UpdatePlayerThemeRequest) SetTheme(v Theme)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *UpdatePlayerThemeRequest) HasTheme() bool`

HasTheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


