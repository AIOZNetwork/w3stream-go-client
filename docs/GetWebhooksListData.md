# GetWebhooksListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Webhooks** | Pointer to [**[]Webhook**](Webhook.md) |  | [optional] 

## Methods

### NewGetWebhooksListData

`func NewGetWebhooksListData() *GetWebhooksListData`

NewGetWebhooksListData instantiates a new GetWebhooksListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWebhooksListDataWithDefaults

`func NewGetWebhooksListDataWithDefaults() *GetWebhooksListData`

NewGetWebhooksListDataWithDefaults instantiates a new GetWebhooksListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetWebhooksListData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetWebhooksListData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetWebhooksListData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetWebhooksListData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetWebhooks

`func (o *GetWebhooksListData) GetWebhooks() []Webhook`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *GetWebhooksListData) GetWebhooksOk() (*[]Webhook, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *GetWebhooksListData) SetWebhooks(v []Webhook)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *GetWebhooksListData) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


