# CreateApiKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyName** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateApiKeyRequest

`func NewCreateApiKeyRequest() *CreateApiKeyRequest`

NewCreateApiKeyRequest instantiates a new CreateApiKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiKeyRequestWithDefaults

`func NewCreateApiKeyRequestWithDefaults() *CreateApiKeyRequest`

NewCreateApiKeyRequestWithDefaults instantiates a new CreateApiKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyName

`func (o *CreateApiKeyRequest) GetApiKeyName() string`

GetApiKeyName returns the ApiKeyName field if non-nil, zero value otherwise.

### GetApiKeyNameOk

`func (o *CreateApiKeyRequest) GetApiKeyNameOk() (*string, bool)`

GetApiKeyNameOk returns a tuple with the ApiKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyName

`func (o *CreateApiKeyRequest) SetApiKeyName(v string)`

SetApiKeyName sets ApiKeyName field to given value.

### HasApiKeyName

`func (o *CreateApiKeyRequest) HasApiKeyName() bool`

HasApiKeyName returns a boolean if a field has been set.

### GetTtl

`func (o *CreateApiKeyRequest) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CreateApiKeyRequest) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CreateApiKeyRequest) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CreateApiKeyRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *CreateApiKeyRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateApiKeyRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateApiKeyRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateApiKeyRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


