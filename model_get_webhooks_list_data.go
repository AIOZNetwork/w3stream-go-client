/*
 * W3STREAM API
 *
 * W3STREAM Service
 *
 * API version: 1.0
 * Contact: support@swagger.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package w3streamsdk

import (
//"encoding/json"
)

// GetWebhooksListData struct for GetWebhooksListData
type GetWebhooksListData struct {
	Total    *int32     `json:"total,omitempty"`
	Webhooks *[]Webhook `json:"webhooks,omitempty"`
}

// NewGetWebhooksListData instantiates a new GetWebhooksListData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWebhooksListData() *GetWebhooksListData {
	this := GetWebhooksListData{}
	return &this
}

// NewGetWebhooksListDataWithDefaults instantiates a new GetWebhooksListData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWebhooksListDataWithDefaults() *GetWebhooksListData {
	this := GetWebhooksListData{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetWebhooksListData) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWebhooksListData) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetWebhooksListData) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *GetWebhooksListData) SetTotal(v int32) {
	o.Total = &v
}

// GetWebhooks returns the Webhooks field value if set, zero value otherwise.
func (o *GetWebhooksListData) GetWebhooks() []Webhook {
	if o == nil || o.Webhooks == nil {
		var ret []Webhook
		return ret
	}
	return *o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWebhooksListData) GetWebhooksOk() (*[]Webhook, bool) {
	if o == nil || o.Webhooks == nil {
		return nil, false
	}
	return o.Webhooks, true
}

// HasWebhooks returns a boolean if a field has been set.
func (o *GetWebhooksListData) HasWebhooks() bool {
	if o != nil && o.Webhooks != nil {
		return true
	}

	return false
}

// SetWebhooks gets a reference to the given []Webhook and assigns it to the Webhooks field.
func (o *GetWebhooksListData) SetWebhooks(v []Webhook) {
	o.Webhooks = &v
}

type NullableGetWebhooksListData struct {
	value *GetWebhooksListData
	isSet bool
}

func (v NullableGetWebhooksListData) Get() *GetWebhooksListData {
	return v.value
}

func (v *NullableGetWebhooksListData) Set(val *GetWebhooksListData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWebhooksListData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWebhooksListData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWebhooksListData(val *GetWebhooksListData) *NullableGetWebhooksListData {
	return &NullableGetWebhooksListData{value: val, isSet: true}
}