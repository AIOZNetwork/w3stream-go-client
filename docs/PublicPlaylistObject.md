# PublicPlaylistObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayerTheme** | Pointer to [**PlayerTheme**](PlayerTheme.md) |  | [optional] 
**Playlist** | Pointer to [**Playlist**](Playlist.md) |  | [optional] 

## Methods

### NewPublicPlaylistObject

`func NewPublicPlaylistObject() *PublicPlaylistObject`

NewPublicPlaylistObject instantiates a new PublicPlaylistObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPlaylistObjectWithDefaults

`func NewPublicPlaylistObjectWithDefaults() *PublicPlaylistObject`

NewPublicPlaylistObjectWithDefaults instantiates a new PublicPlaylistObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayerTheme

`func (o *PublicPlaylistObject) GetPlayerTheme() PlayerTheme`

GetPlayerTheme returns the PlayerTheme field if non-nil, zero value otherwise.

### GetPlayerThemeOk

`func (o *PublicPlaylistObject) GetPlayerThemeOk() (*PlayerTheme, bool)`

GetPlayerThemeOk returns a tuple with the PlayerTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerTheme

`func (o *PublicPlaylistObject) SetPlayerTheme(v PlayerTheme)`

SetPlayerTheme sets PlayerTheme field to given value.

### HasPlayerTheme

`func (o *PublicPlaylistObject) HasPlayerTheme() bool`

HasPlayerTheme returns a boolean if a field has been set.

### GetPlaylist

`func (o *PublicPlaylistObject) GetPlaylist() Playlist`

GetPlaylist returns the Playlist field if non-nil, zero value otherwise.

### GetPlaylistOk

`func (o *PublicPlaylistObject) GetPlaylistOk() (*Playlist, bool)`

GetPlaylistOk returns a tuple with the Playlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylist

`func (o *PublicPlaylistObject) SetPlaylist(v Playlist)`

SetPlaylist sets Playlist field to given value.

### HasPlaylist

`func (o *PublicPlaylistObject) HasPlaylist() bool`

HasPlaylist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


