# CreatePlayerThemeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controls** | Pointer to [**Controls**](Controls.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Theme** | Pointer to [**Theme**](Theme.md) |  | [optional] 

## Methods

### NewCreatePlayerThemeRequest

`func NewCreatePlayerThemeRequest() *CreatePlayerThemeRequest`

NewCreatePlayerThemeRequest instantiates a new CreatePlayerThemeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePlayerThemeRequestWithDefaults

`func NewCreatePlayerThemeRequestWithDefaults() *CreatePlayerThemeRequest`

NewCreatePlayerThemeRequestWithDefaults instantiates a new CreatePlayerThemeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControls

`func (o *CreatePlayerThemeRequest) GetControls() Controls`

GetControls returns the Controls field if non-nil, zero value otherwise.

### GetControlsOk

`func (o *CreatePlayerThemeRequest) GetControlsOk() (*Controls, bool)`

GetControlsOk returns a tuple with the Controls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControls

`func (o *CreatePlayerThemeRequest) SetControls(v Controls)`

SetControls sets Controls field to given value.

### HasControls

`func (o *CreatePlayerThemeRequest) HasControls() bool`

HasControls returns a boolean if a field has been set.

### GetName

`func (o *CreatePlayerThemeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePlayerThemeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePlayerThemeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatePlayerThemeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTheme

`func (o *CreatePlayerThemeRequest) GetTheme() Theme`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *CreatePlayerThemeRequest) GetThemeOk() (*Theme, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *CreatePlayerThemeRequest) SetTheme(v Theme)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *CreatePlayerThemeRequest) HasTheme() bool`

HasTheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


