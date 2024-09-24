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

// CreateApiKeyData struct for CreateApiKeyData
type CreateApiKeyData struct {
	ApiKey *ApiKey `json:"api_key,omitempty"`
}

// NewCreateApiKeyData instantiates a new CreateApiKeyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApiKeyData() *CreateApiKeyData {
	this := CreateApiKeyData{}
	return &this
}

// NewCreateApiKeyDataWithDefaults instantiates a new CreateApiKeyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApiKeyDataWithDefaults() *CreateApiKeyData {
	this := CreateApiKeyData{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *CreateApiKeyData) GetApiKey() ApiKey {
	if o == nil || o.ApiKey == nil {
		var ret ApiKey
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiKeyData) GetApiKeyOk() (*ApiKey, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *CreateApiKeyData) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given ApiKey and assigns it to the ApiKey field.
func (o *CreateApiKeyData) SetApiKey(v ApiKey) {
	o.ApiKey = &v
}

type NullableCreateApiKeyData struct {
	value *CreateApiKeyData
	isSet bool
}

func (v NullableCreateApiKeyData) Get() *CreateApiKeyData {
	return v.value
}

func (v *NullableCreateApiKeyData) Set(val *CreateApiKeyData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApiKeyData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApiKeyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApiKeyData(val *CreateApiKeyData) *NullableCreateApiKeyData {
	return &NullableCreateApiKeyData{value: val, isSet: true}
}