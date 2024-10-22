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

// LiveStreamVideoResponse struct for LiveStreamVideoResponse
type LiveStreamVideoResponse struct {
	Assets          *LiveStreamAssets `json:"assets,omitempty"`
	CreatedAt       *string           `json:"created_at,omitempty"`
	Duration        *int32            `json:"duration,omitempty"`
	Id              *string           `json:"id,omitempty"`
	LiveStreamKeyId *string           `json:"live_stream_key_id,omitempty"`
	Qualities       *[]string         `json:"qualities,omitempty"`
	Save            *bool             `json:"save,omitempty"`
	Status          *string           `json:"status,omitempty"`
	Title           *string           `json:"title,omitempty"`
	UpdatedAt       *string           `json:"updated_at,omitempty"`
	UserId          *string           `json:"user_id,omitempty"`
	Video           *VideoObject      `json:"video,omitempty"`
}

// NewLiveStreamVideoResponse instantiates a new LiveStreamVideoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveStreamVideoResponse() *LiveStreamVideoResponse {
	this := LiveStreamVideoResponse{}
	return &this
}

// NewLiveStreamVideoResponseWithDefaults instantiates a new LiveStreamVideoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveStreamVideoResponseWithDefaults() *LiveStreamVideoResponse {
	this := LiveStreamVideoResponse{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *LiveStreamVideoResponse) GetAssets() LiveStreamAssets {
	if o == nil || o.Assets == nil {
		var ret LiveStreamAssets
		return ret
	}
	return *o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStreamVideoResponse) GetAssetsOk() (*LiveStreamAssets, bool) {
	if o == nil || o.Assets == nil {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *LiveStreamVideoResponse) HasAssets() bool {
	if o != nil && o.Assets != nil {
		return true
	}

	return false
}

// SetAssets gets a reference to the given LiveStreamAssets and assigns it to the Assets field.
func (o *LiveStreamVideoResponse) SetAssets(v LiveStreamAssets) {
	o.Assets = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LiveStreamVideoResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStreamVideoResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LiveStreamVideoResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *LiveStreamVideoResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *LiveStreamVideoResponse) GetDuration() int32 {
	if o == nil || o.Duration == nil {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStreamVideoResponse) GetDurationOk() (*int32, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *LiveStreamVideoResponse) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *LiveStreamVideoResponse) SetDuration(v int32) {
	o.Duration = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LiveStreamVideoResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStreamVideoResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LiveStreamVideoResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LiveStreamVideoResponse) SetId(v string) {
	o.Id = &v
}

// GetLiveStreamKeyId returns the LiveStreamKeyId field value if set, zero value otherwise.
func (o *LiveStreamVideoResponse) GetLiveStreamKeyId() string {
	if o == nil || o.LiveStreamKeyId == nil {
		var ret string
		return ret
	}
	return *o.LiveStreamKeyId
}

// GetLiveStreamKeyIdOk returns a tuple with the LiveStreamKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStreamVideoResponse) GetLiveStreamKeyIdOk() (*string, bool) {
	if o == nil || o.LiveStreamKeyId == nil {
		return nil, false
	}
	return o.LiveStreamKeyId, true
}

// HasLiveStreamKeyId returns a boolean if a field has been set.
func (o *LiveStreamVideoResponse) HasLiveStreamKeyId() bool {
	if o != nil && o.LiveStreamKeyId != nil {
		return true
	}

	return false
}

// SetLiveStreamKeyId gets a reference to the given string and assigns it to the LiveStreamKeyId field.
func (o *LiveStreamVideoResponse) SetLiveStreamKeyId(v string) {
	o.LiveStreamKeyId = &v
}

// GetQualities returns the Qualities field value if set, zero value otherwise.
func (o *LiveStreamVideoResponse) GetQualities() []string {
	if o == nil || o.Qualities == nil {
		var ret []string
		return ret
	}
	return *o.Qualities
}

// GetQualitiesOk returns a tuple with the Qualities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStreamVideoResponse) GetQualitiesOk() (*[]string, bool) {
	if o == nil || o.Qualities == nil {
		return nil, false
	}
	return o.Qualities, true
}

// HasQualities returns a boolean if a field has been set.
func (o *LiveStreamVideoResponse) HasQualities() bool {
	if o != nil && o.Qualities != nil {
		return true
	}

	return false
}

// SetQualities gets a reference to the given []string and assigns it to the Qualities field.
func (o *LiveStreamVideoResponse) SetQualities(v []string) {
	o.Qualities = &v
}

// GetSave returns the Save field value if set, zero value otherwise.
func (o *LiveStreamVideoResponse) GetSave() bool {
	if o == nil || o.Save == nil {
		var ret bool
		return ret
	}
	return *o.Save
}

// GetSaveOk returns a tuple with the Save field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStreamVideoResponse) GetSaveOk() (*bool, bool) {
	if o == nil || o.Save == nil {
		return nil, false
	}
	return o.Save, true
}

// HasSave returns a boolean if a field has been set.
func (o *LiveStreamVideoResponse) HasSave() bool {
	if o != nil && o.Save != nil {
		return true
	}

	return false
}

// SetSave gets a reference to the given bool and assigns it to the Save field.
func (o *LiveStreamVideoResponse) SetSave(v bool) {
	o.Save = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LiveStreamVideoResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStreamVideoResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LiveStreamVideoResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LiveStreamVideoResponse) SetStatus(v string) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *LiveStreamVideoResponse) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStreamVideoResponse) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *LiveStreamVideoResponse) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *LiveStreamVideoResponse) SetTitle(v string) {
	o.Title = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *LiveStreamVideoResponse) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStreamVideoResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LiveStreamVideoResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *LiveStreamVideoResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *LiveStreamVideoResponse) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStreamVideoResponse) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *LiveStreamVideoResponse) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *LiveStreamVideoResponse) SetUserId(v string) {
	o.UserId = &v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *LiveStreamVideoResponse) GetVideo() VideoObject {
	if o == nil || o.Video == nil {
		var ret VideoObject
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStreamVideoResponse) GetVideoOk() (*VideoObject, bool) {
	if o == nil || o.Video == nil {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *LiveStreamVideoResponse) HasVideo() bool {
	if o != nil && o.Video != nil {
		return true
	}

	return false
}

// SetVideo gets a reference to the given VideoObject and assigns it to the Video field.
func (o *LiveStreamVideoResponse) SetVideo(v VideoObject) {
	o.Video = &v
}

type NullableLiveStreamVideoResponse struct {
	value *LiveStreamVideoResponse
	isSet bool
}

func (v NullableLiveStreamVideoResponse) Get() *LiveStreamVideoResponse {
	return v.value
}

func (v *NullableLiveStreamVideoResponse) Set(val *LiveStreamVideoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveStreamVideoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveStreamVideoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveStreamVideoResponse(val *LiveStreamVideoResponse) *NullableLiveStreamVideoResponse {
	return &NullableLiveStreamVideoResponse{value: val, isSet: true}
}