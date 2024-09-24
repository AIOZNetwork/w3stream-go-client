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

// VideoAssets struct for VideoAssets
type VideoAssets struct {
	HlsUrl       *string `json:"hls_url,omitempty"`
	Iframe       *string `json:"iframe,omitempty"`
	Mp4Url       *string `json:"mp4_url,omitempty"`
	PlayerUrl    *string `json:"player_url,omitempty"`
	SourceUrl    *string `json:"source_url,omitempty"`
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
}

// NewVideoAssets instantiates a new VideoAssets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoAssets() *VideoAssets {
	this := VideoAssets{}
	return &this
}

// NewVideoAssetsWithDefaults instantiates a new VideoAssets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoAssetsWithDefaults() *VideoAssets {
	this := VideoAssets{}
	return &this
}

// GetHlsUrl returns the HlsUrl field value if set, zero value otherwise.
func (o *VideoAssets) GetHlsUrl() string {
	if o == nil || o.HlsUrl == nil {
		var ret string
		return ret
	}
	return *o.HlsUrl
}

// GetHlsUrlOk returns a tuple with the HlsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoAssets) GetHlsUrlOk() (*string, bool) {
	if o == nil || o.HlsUrl == nil {
		return nil, false
	}
	return o.HlsUrl, true
}

// HasHlsUrl returns a boolean if a field has been set.
func (o *VideoAssets) HasHlsUrl() bool {
	if o != nil && o.HlsUrl != nil {
		return true
	}

	return false
}

// SetHlsUrl gets a reference to the given string and assigns it to the HlsUrl field.
func (o *VideoAssets) SetHlsUrl(v string) {
	o.HlsUrl = &v
}

// GetIframe returns the Iframe field value if set, zero value otherwise.
func (o *VideoAssets) GetIframe() string {
	if o == nil || o.Iframe == nil {
		var ret string
		return ret
	}
	return *o.Iframe
}

// GetIframeOk returns a tuple with the Iframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoAssets) GetIframeOk() (*string, bool) {
	if o == nil || o.Iframe == nil {
		return nil, false
	}
	return o.Iframe, true
}

// HasIframe returns a boolean if a field has been set.
func (o *VideoAssets) HasIframe() bool {
	if o != nil && o.Iframe != nil {
		return true
	}

	return false
}

// SetIframe gets a reference to the given string and assigns it to the Iframe field.
func (o *VideoAssets) SetIframe(v string) {
	o.Iframe = &v
}

// GetMp4Url returns the Mp4Url field value if set, zero value otherwise.
func (o *VideoAssets) GetMp4Url() string {
	if o == nil || o.Mp4Url == nil {
		var ret string
		return ret
	}
	return *o.Mp4Url
}

// GetMp4UrlOk returns a tuple with the Mp4Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoAssets) GetMp4UrlOk() (*string, bool) {
	if o == nil || o.Mp4Url == nil {
		return nil, false
	}
	return o.Mp4Url, true
}

// HasMp4Url returns a boolean if a field has been set.
func (o *VideoAssets) HasMp4Url() bool {
	if o != nil && o.Mp4Url != nil {
		return true
	}

	return false
}

// SetMp4Url gets a reference to the given string and assigns it to the Mp4Url field.
func (o *VideoAssets) SetMp4Url(v string) {
	o.Mp4Url = &v
}

// GetPlayerUrl returns the PlayerUrl field value if set, zero value otherwise.
func (o *VideoAssets) GetPlayerUrl() string {
	if o == nil || o.PlayerUrl == nil {
		var ret string
		return ret
	}
	return *o.PlayerUrl
}

// GetPlayerUrlOk returns a tuple with the PlayerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoAssets) GetPlayerUrlOk() (*string, bool) {
	if o == nil || o.PlayerUrl == nil {
		return nil, false
	}
	return o.PlayerUrl, true
}

// HasPlayerUrl returns a boolean if a field has been set.
func (o *VideoAssets) HasPlayerUrl() bool {
	if o != nil && o.PlayerUrl != nil {
		return true
	}

	return false
}

// SetPlayerUrl gets a reference to the given string and assigns it to the PlayerUrl field.
func (o *VideoAssets) SetPlayerUrl(v string) {
	o.PlayerUrl = &v
}

// GetSourceUrl returns the SourceUrl field value if set, zero value otherwise.
func (o *VideoAssets) GetSourceUrl() string {
	if o == nil || o.SourceUrl == nil {
		var ret string
		return ret
	}
	return *o.SourceUrl
}

// GetSourceUrlOk returns a tuple with the SourceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoAssets) GetSourceUrlOk() (*string, bool) {
	if o == nil || o.SourceUrl == nil {
		return nil, false
	}
	return o.SourceUrl, true
}

// HasSourceUrl returns a boolean if a field has been set.
func (o *VideoAssets) HasSourceUrl() bool {
	if o != nil && o.SourceUrl != nil {
		return true
	}

	return false
}

// SetSourceUrl gets a reference to the given string and assigns it to the SourceUrl field.
func (o *VideoAssets) SetSourceUrl(v string) {
	o.SourceUrl = &v
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise.
func (o *VideoAssets) GetThumbnailUrl() string {
	if o == nil || o.ThumbnailUrl == nil {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoAssets) GetThumbnailUrlOk() (*string, bool) {
	if o == nil || o.ThumbnailUrl == nil {
		return nil, false
	}
	return o.ThumbnailUrl, true
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *VideoAssets) HasThumbnailUrl() bool {
	if o != nil && o.ThumbnailUrl != nil {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given string and assigns it to the ThumbnailUrl field.
func (o *VideoAssets) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = &v
}

type NullableVideoAssets struct {
	value *VideoAssets
	isSet bool
}

func (v NullableVideoAssets) Get() *VideoAssets {
	return v.value
}

func (v *NullableVideoAssets) Set(val *VideoAssets) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoAssets) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoAssets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoAssets(val *VideoAssets) *NullableVideoAssets {
	return &NullableVideoAssets{value: val, isSet: true}
}