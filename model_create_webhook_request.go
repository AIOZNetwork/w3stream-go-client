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

// CreateWebhookRequest struct for CreateWebhookRequest
type CreateWebhookRequest struct {
	EncodingFinished *bool   `json:"encoding_finished,omitempty"`
	EncodingStarted  *bool   `json:"encoding_started,omitempty"`
	FileReceived     *bool   `json:"file_received,omitempty"`
	Name             *string `json:"name,omitempty"`
	Url              *string `json:"url,omitempty"`
}

// NewCreateWebhookRequest instantiates a new CreateWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebhookRequest() *CreateWebhookRequest {
	this := CreateWebhookRequest{}
	return &this
}

// NewCreateWebhookRequestWithDefaults instantiates a new CreateWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebhookRequestWithDefaults() *CreateWebhookRequest {
	this := CreateWebhookRequest{}
	return &this
}

// GetEncodingFinished returns the EncodingFinished field value if set, zero value otherwise.
func (o *CreateWebhookRequest) GetEncodingFinished() bool {
	if o == nil || o.EncodingFinished == nil {
		var ret bool
		return ret
	}
	return *o.EncodingFinished
}

// GetEncodingFinishedOk returns a tuple with the EncodingFinished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetEncodingFinishedOk() (*bool, bool) {
	if o == nil || o.EncodingFinished == nil {
		return nil, false
	}
	return o.EncodingFinished, true
}

// HasEncodingFinished returns a boolean if a field has been set.
func (o *CreateWebhookRequest) HasEncodingFinished() bool {
	if o != nil && o.EncodingFinished != nil {
		return true
	}

	return false
}

// SetEncodingFinished gets a reference to the given bool and assigns it to the EncodingFinished field.
func (o *CreateWebhookRequest) SetEncodingFinished(v bool) {
	o.EncodingFinished = &v
}

// GetEncodingStarted returns the EncodingStarted field value if set, zero value otherwise.
func (o *CreateWebhookRequest) GetEncodingStarted() bool {
	if o == nil || o.EncodingStarted == nil {
		var ret bool
		return ret
	}
	return *o.EncodingStarted
}

// GetEncodingStartedOk returns a tuple with the EncodingStarted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetEncodingStartedOk() (*bool, bool) {
	if o == nil || o.EncodingStarted == nil {
		return nil, false
	}
	return o.EncodingStarted, true
}

// HasEncodingStarted returns a boolean if a field has been set.
func (o *CreateWebhookRequest) HasEncodingStarted() bool {
	if o != nil && o.EncodingStarted != nil {
		return true
	}

	return false
}

// SetEncodingStarted gets a reference to the given bool and assigns it to the EncodingStarted field.
func (o *CreateWebhookRequest) SetEncodingStarted(v bool) {
	o.EncodingStarted = &v
}

// GetFileReceived returns the FileReceived field value if set, zero value otherwise.
func (o *CreateWebhookRequest) GetFileReceived() bool {
	if o == nil || o.FileReceived == nil {
		var ret bool
		return ret
	}
	return *o.FileReceived
}

// GetFileReceivedOk returns a tuple with the FileReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetFileReceivedOk() (*bool, bool) {
	if o == nil || o.FileReceived == nil {
		return nil, false
	}
	return o.FileReceived, true
}

// HasFileReceived returns a boolean if a field has been set.
func (o *CreateWebhookRequest) HasFileReceived() bool {
	if o != nil && o.FileReceived != nil {
		return true
	}

	return false
}

// SetFileReceived gets a reference to the given bool and assigns it to the FileReceived field.
func (o *CreateWebhookRequest) SetFileReceived(v bool) {
	o.FileReceived = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateWebhookRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateWebhookRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateWebhookRequest) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CreateWebhookRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateWebhookRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CreateWebhookRequest) SetUrl(v string) {
	o.Url = &v
}

type NullableCreateWebhookRequest struct {
	value *CreateWebhookRequest
	isSet bool
}

func (v NullableCreateWebhookRequest) Get() *CreateWebhookRequest {
	return v.value
}

func (v *NullableCreateWebhookRequest) Set(val *CreateWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebhookRequest(val *CreateWebhookRequest) *NullableCreateWebhookRequest {
	return &NullableCreateWebhookRequest{value: val, isSet: true}
}