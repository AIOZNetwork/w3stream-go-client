/*
 * VMS API
 *
 * VMS Service
 *
 * API version: 1.0
 * Contact: support@swagger.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package w3streamsdk

import (
//"encoding/json"
)

// GetVideoPlayerInfoResponse struct for GetVideoPlayerInfoResponse
type GetVideoPlayerInfoResponse struct {
	Assets        *VideoAssets     `json:"assets,omitempty"`
	Captions      *[]VideoCaption  `json:"captions,omitempty"`
	Chapters      *[]VideoChapter  `json:"chapters,omitempty"`
	CreatedAt     *string          `json:"created_at,omitempty"`
	Description   *string          `json:"description,omitempty"`
	Duration      *float32         `json:"duration,omitempty"`
	Id            *string          `json:"id,omitempty"`
	IsMp4         *bool            `json:"is_mp4,omitempty"`
	IsPublic      *bool            `json:"is_public,omitempty"`
	Metadata      *[]Metadata      `json:"metadata,omitempty"`
	PlayerTheme   *PlayerTheme     `json:"player_theme,omitempty"`
	PlayerThemeId *string          `json:"player_theme_id,omitempty"`
	Qualities     *[]QualityObject `json:"qualities,omitempty"`
	Size          *int32           `json:"size,omitempty"`
	Status        *string          `json:"status,omitempty"`
	Tags          *[]string        `json:"tags,omitempty"`
	Title         *string          `json:"title,omitempty"`
	UpdatedAt     *string          `json:"updated_at,omitempty"`
	UserId        *string          `json:"user_id,omitempty"`
}

// NewGetVideoPlayerInfoResponse instantiates a new GetVideoPlayerInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVideoPlayerInfoResponse() *GetVideoPlayerInfoResponse {
	this := GetVideoPlayerInfoResponse{}
	return &this
}

// NewGetVideoPlayerInfoResponseWithDefaults instantiates a new GetVideoPlayerInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVideoPlayerInfoResponseWithDefaults() *GetVideoPlayerInfoResponse {
	this := GetVideoPlayerInfoResponse{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetAssets() VideoAssets {
	if o == nil || o.Assets == nil {
		var ret VideoAssets
		return ret
	}
	return *o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetAssetsOk() (*VideoAssets, bool) {
	if o == nil || o.Assets == nil {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasAssets() bool {
	if o != nil && o.Assets != nil {
		return true
	}

	return false
}

// SetAssets gets a reference to the given VideoAssets and assigns it to the Assets field.
func (o *GetVideoPlayerInfoResponse) SetAssets(v VideoAssets) {
	o.Assets = &v
}

// GetCaptions returns the Captions field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetCaptions() []VideoCaption {
	if o == nil || o.Captions == nil {
		var ret []VideoCaption
		return ret
	}
	return *o.Captions
}

// GetCaptionsOk returns a tuple with the Captions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetCaptionsOk() (*[]VideoCaption, bool) {
	if o == nil || o.Captions == nil {
		return nil, false
	}
	return o.Captions, true
}

// HasCaptions returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasCaptions() bool {
	if o != nil && o.Captions != nil {
		return true
	}

	return false
}

// SetCaptions gets a reference to the given []VideoCaption and assigns it to the Captions field.
func (o *GetVideoPlayerInfoResponse) SetCaptions(v []VideoCaption) {
	o.Captions = &v
}

// GetChapters returns the Chapters field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetChapters() []VideoChapter {
	if o == nil || o.Chapters == nil {
		var ret []VideoChapter
		return ret
	}
	return *o.Chapters
}

// GetChaptersOk returns a tuple with the Chapters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetChaptersOk() (*[]VideoChapter, bool) {
	if o == nil || o.Chapters == nil {
		return nil, false
	}
	return o.Chapters, true
}

// HasChapters returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasChapters() bool {
	if o != nil && o.Chapters != nil {
		return true
	}

	return false
}

// SetChapters gets a reference to the given []VideoChapter and assigns it to the Chapters field.
func (o *GetVideoPlayerInfoResponse) SetChapters(v []VideoChapter) {
	o.Chapters = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GetVideoPlayerInfoResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetVideoPlayerInfoResponse) SetDescription(v string) {
	o.Description = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetDuration() float32 {
	if o == nil || o.Duration == nil {
		var ret float32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetDurationOk() (*float32, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given float32 and assigns it to the Duration field.
func (o *GetVideoPlayerInfoResponse) SetDuration(v float32) {
	o.Duration = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetVideoPlayerInfoResponse) SetId(v string) {
	o.Id = &v
}

// GetIsMp4 returns the IsMp4 field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetIsMp4() bool {
	if o == nil || o.IsMp4 == nil {
		var ret bool
		return ret
	}
	return *o.IsMp4
}

// GetIsMp4Ok returns a tuple with the IsMp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetIsMp4Ok() (*bool, bool) {
	if o == nil || o.IsMp4 == nil {
		return nil, false
	}
	return o.IsMp4, true
}

// HasIsMp4 returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasIsMp4() bool {
	if o != nil && o.IsMp4 != nil {
		return true
	}

	return false
}

// SetIsMp4 gets a reference to the given bool and assigns it to the IsMp4 field.
func (o *GetVideoPlayerInfoResponse) SetIsMp4(v bool) {
	o.IsMp4 = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetIsPublic() bool {
	if o == nil || o.IsPublic == nil {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetIsPublicOk() (*bool, bool) {
	if o == nil || o.IsPublic == nil {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasIsPublic() bool {
	if o != nil && o.IsPublic != nil {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *GetVideoPlayerInfoResponse) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetMetadata() []Metadata {
	if o == nil || o.Metadata == nil {
		var ret []Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetMetadataOk() (*[]Metadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []Metadata and assigns it to the Metadata field.
func (o *GetVideoPlayerInfoResponse) SetMetadata(v []Metadata) {
	o.Metadata = &v
}

// GetPlayerTheme returns the PlayerTheme field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetPlayerTheme() PlayerTheme {
	if o == nil || o.PlayerTheme == nil {
		var ret PlayerTheme
		return ret
	}
	return *o.PlayerTheme
}

// GetPlayerThemeOk returns a tuple with the PlayerTheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetPlayerThemeOk() (*PlayerTheme, bool) {
	if o == nil || o.PlayerTheme == nil {
		return nil, false
	}
	return o.PlayerTheme, true
}

// HasPlayerTheme returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasPlayerTheme() bool {
	if o != nil && o.PlayerTheme != nil {
		return true
	}

	return false
}

// SetPlayerTheme gets a reference to the given PlayerTheme and assigns it to the PlayerTheme field.
func (o *GetVideoPlayerInfoResponse) SetPlayerTheme(v PlayerTheme) {
	o.PlayerTheme = &v
}

// GetPlayerThemeId returns the PlayerThemeId field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetPlayerThemeId() string {
	if o == nil || o.PlayerThemeId == nil {
		var ret string
		return ret
	}
	return *o.PlayerThemeId
}

// GetPlayerThemeIdOk returns a tuple with the PlayerThemeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetPlayerThemeIdOk() (*string, bool) {
	if o == nil || o.PlayerThemeId == nil {
		return nil, false
	}
	return o.PlayerThemeId, true
}

// HasPlayerThemeId returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasPlayerThemeId() bool {
	if o != nil && o.PlayerThemeId != nil {
		return true
	}

	return false
}

// SetPlayerThemeId gets a reference to the given string and assigns it to the PlayerThemeId field.
func (o *GetVideoPlayerInfoResponse) SetPlayerThemeId(v string) {
	o.PlayerThemeId = &v
}

// GetQualities returns the Qualities field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetQualities() []QualityObject {
	if o == nil || o.Qualities == nil {
		var ret []QualityObject
		return ret
	}
	return *o.Qualities
}

// GetQualitiesOk returns a tuple with the Qualities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetQualitiesOk() (*[]QualityObject, bool) {
	if o == nil || o.Qualities == nil {
		return nil, false
	}
	return o.Qualities, true
}

// HasQualities returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasQualities() bool {
	if o != nil && o.Qualities != nil {
		return true
	}

	return false
}

// SetQualities gets a reference to the given []QualityObject and assigns it to the Qualities field.
func (o *GetVideoPlayerInfoResponse) SetQualities(v []QualityObject) {
	o.Qualities = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *GetVideoPlayerInfoResponse) SetSize(v int32) {
	o.Size = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetVideoPlayerInfoResponse) SetStatus(v string) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GetVideoPlayerInfoResponse) SetTags(v []string) {
	o.Tags = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetVideoPlayerInfoResponse) SetTitle(v string) {
	o.Title = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GetVideoPlayerInfoResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *GetVideoPlayerInfoResponse) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoPlayerInfoResponse) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *GetVideoPlayerInfoResponse) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *GetVideoPlayerInfoResponse) SetUserId(v string) {
	o.UserId = &v
}

type NullableGetVideoPlayerInfoResponse struct {
	value *GetVideoPlayerInfoResponse
	isSet bool
}

func (v NullableGetVideoPlayerInfoResponse) Get() *GetVideoPlayerInfoResponse {
	return v.value
}

func (v *NullableGetVideoPlayerInfoResponse) Set(val *GetVideoPlayerInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVideoPlayerInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVideoPlayerInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVideoPlayerInfoResponse(val *GetVideoPlayerInfoResponse) *NullableGetVideoPlayerInfoResponse {
	return &NullableGetVideoPlayerInfoResponse{value: val, isSet: true}
}