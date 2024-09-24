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

// GetPlayerThemeResponse struct for GetPlayerThemeResponse
type GetPlayerThemeResponse struct {
	Data   *GetPlayerThemeData `json:"data,omitempty"`
	Status *string             `json:"status,omitempty"`
}

// NewGetPlayerThemeResponse instantiates a new GetPlayerThemeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPlayerThemeResponse() *GetPlayerThemeResponse {
	this := GetPlayerThemeResponse{}
	return &this
}

// NewGetPlayerThemeResponseWithDefaults instantiates a new GetPlayerThemeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPlayerThemeResponseWithDefaults() *GetPlayerThemeResponse {
	this := GetPlayerThemeResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetPlayerThemeResponse) GetData() GetPlayerThemeData {
	if o == nil || o.Data == nil {
		var ret GetPlayerThemeData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPlayerThemeResponse) GetDataOk() (*GetPlayerThemeData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetPlayerThemeResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetPlayerThemeData and assigns it to the Data field.
func (o *GetPlayerThemeResponse) SetData(v GetPlayerThemeData) {
	o.Data = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetPlayerThemeResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPlayerThemeResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetPlayerThemeResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetPlayerThemeResponse) SetStatus(v string) {
	o.Status = &v
}

type NullableGetPlayerThemeResponse struct {
	value *GetPlayerThemeResponse
	isSet bool
}

func (v NullableGetPlayerThemeResponse) Get() *GetPlayerThemeResponse {
	return v.value
}

func (v *NullableGetPlayerThemeResponse) Set(val *GetPlayerThemeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPlayerThemeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPlayerThemeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPlayerThemeResponse(val *GetPlayerThemeResponse) *NullableGetPlayerThemeResponse {
	return &NullableGetPlayerThemeResponse{value: val, isSet: true}
}