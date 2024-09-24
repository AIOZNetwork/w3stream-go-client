# GetUserWebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GetUserWebhookData**](GetUserWebhookData.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetUserWebhookResponse

`func NewGetUserWebhookResponse() *GetUserWebhookResponse`

NewGetUserWebhookResponse instantiates a new GetUserWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserWebhookResponseWithDefaults

`func NewGetUserWebhookResponseWithDefaults() *GetUserWebhookResponse`

NewGetUserWebhookResponseWithDefaults instantiates a new GetUserWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetUserWebhookResponse) GetData() GetUserWebhookData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUserWebhookResponse) GetDataOk() (*GetUserWebhookData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUserWebhookResponse) SetData(v GetUserWebhookData)`

SetData sets Data field to given value.

### HasData

`func (o *GetUserWebhookResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *GetUserWebhookResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetUserWebhookResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetUserWebhookResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetUserWebhookResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


