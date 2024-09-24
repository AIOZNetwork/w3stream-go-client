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

// Webhook struct for Webhook
type Webhook struct {
	CreatedAt        *string `json:"created_at,omitempty"`
	EncodingFinished *bool   `json:"encoding_finished,omitempty"`
	EncodingStarted  *bool   `json:"encoding_started,omitempty"`
	FileReceived     *bool   `json:"file_received,omitempty"`
	Id               *string `json:"id,omitempty"`
	LastTriggeredAt  *string `json:"last_triggered_at,omitempty"`
	Name             *string `json:"name,omitempty"`
	UpdatedAt        *string `json:"updated_at,omitempty"`
	Url              *string `json:"url,omitempty"`
	UserId           *string `json:"user_id,omitempty"`
}

// NewWebhook instantiates a new Webhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhook() *Webhook {
	this := Webhook{}
	return &this
}

// NewWebhookWithDefaults instantiates a new Webhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookWithDefaults() *Webhook {
	this := Webhook{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Webhook) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Webhook) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Webhook) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetEncodingFinished returns the EncodingFinished field value if set, zero value otherwise.
func (o *Webhook) GetEncodingFinished() bool {
	if o == nil || o.EncodingFinished == nil {
		var ret bool
		return ret
	}
	return *o.EncodingFinished
}

// GetEncodingFinishedOk returns a tuple with the EncodingFinished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetEncodingFinishedOk() (*bool, bool) {
	if o == nil || o.EncodingFinished == nil {
		return nil, false
	}
	return o.EncodingFinished, true
}

// HasEncodingFinished returns a boolean if a field has been set.
func (o *Webhook) HasEncodingFinished() bool {
	if o != nil && o.EncodingFinished != nil {
		return true
	}

	return false
}

// SetEncodingFinished gets a reference to the given bool and assigns it to the EncodingFinished field.
func (o *Webhook) SetEncodingFinished(v bool) {
	o.EncodingFinished = &v
}

// GetEncodingStarted returns the EncodingStarted field value if set, zero value otherwise.
func (o *Webhook) GetEncodingStarted() bool {
	if o == nil || o.EncodingStarted == nil {
		var ret bool
		return ret
	}
	return *o.EncodingStarted
}

// GetEncodingStartedOk returns a tuple with the EncodingStarted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetEncodingStartedOk() (*bool, bool) {
	if o == nil || o.EncodingStarted == nil {
		return nil, false
	}
	return o.EncodingStarted, true
}

// HasEncodingStarted returns a boolean if a field has been set.
func (o *Webhook) HasEncodingStarted() bool {
	if o != nil && o.EncodingStarted != nil {
		return true
	}

	return false
}

// SetEncodingStarted gets a reference to the given bool and assigns it to the EncodingStarted field.
func (o *Webhook) SetEncodingStarted(v bool) {
	o.EncodingStarted = &v
}

// GetFileReceived returns the FileReceived field value if set, zero value otherwise.
func (o *Webhook) GetFileReceived() bool {
	if o == nil || o.FileReceived == nil {
		var ret bool
		return ret
	}
	return *o.FileReceived
}

// GetFileReceivedOk returns a tuple with the FileReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetFileReceivedOk() (*bool, bool) {
	if o == nil || o.FileReceived == nil {
		return nil, false
	}
	return o.FileReceived, true
}

// HasFileReceived returns a boolean if a field has been set.
func (o *Webhook) HasFileReceived() bool {
	if o != nil && o.FileReceived != nil {
		return true
	}

	return false
}

// SetFileReceived gets a reference to the given bool and assigns it to the FileReceived field.
func (o *Webhook) SetFileReceived(v bool) {
	o.FileReceived = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Webhook) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Webhook) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Webhook) SetId(v string) {
	o.Id = &v
}

// GetLastTriggeredAt returns the LastTriggeredAt field value if set, zero value otherwise.
func (o *Webhook) GetLastTriggeredAt() string {
	if o == nil || o.LastTriggeredAt == nil {
		var ret string
		return ret
	}
	return *o.LastTriggeredAt
}

// GetLastTriggeredAtOk returns a tuple with the LastTriggeredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetLastTriggeredAtOk() (*string, bool) {
	if o == nil || o.LastTriggeredAt == nil {
		return nil, false
	}
	return o.LastTriggeredAt, true
}

// HasLastTriggeredAt returns a boolean if a field has been set.
func (o *Webhook) HasLastTriggeredAt() bool {
	if o != nil && o.LastTriggeredAt != nil {
		return true
	}

	return false
}

// SetLastTriggeredAt gets a reference to the given string and assigns it to the LastTriggeredAt field.
func (o *Webhook) SetLastTriggeredAt(v string) {
	o.LastTriggeredAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Webhook) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Webhook) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Webhook) SetName(v string) {
	o.Name = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Webhook) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Webhook) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Webhook) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Webhook) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Webhook) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Webhook) SetUrl(v string) {
	o.Url = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Webhook) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Webhook) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Webhook) SetUserId(v string) {
	o.UserId = &v
}

type NullableWebhook struct {
	value *Webhook
	isSet bool
}

func (v NullableWebhook) Get() *Webhook {
	return v.value
}

func (v *NullableWebhook) Set(val *Webhook) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhook(val *Webhook) *NullableWebhook {
	return &NullableWebhook{value: val, isSet: true}
}