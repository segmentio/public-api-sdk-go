# UpdateSourceAlphaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | An optional human-readable name to associate with the Source.  Config API note: equal to &#x60;displayName&#x60;. | [optional] 
**Enabled** | Pointer to **bool** | Enable to allow the Source to send data. | [optional] 
**Slug** | Pointer to **string** | The slug that identifies the Source.  Config API note: equal to &#x60;name&#x60;. | [optional] 
**Settings** | Pointer to **map[string]interface{}** | A key-value object that contains instance-specific settings for a Source.  The &#x60;options&#x60; field in the Source metadata defines the schema of this object. | [optional] 

## Methods

### NewUpdateSourceAlphaInput

`func NewUpdateSourceAlphaInput() *UpdateSourceAlphaInput`

NewUpdateSourceAlphaInput instantiates a new UpdateSourceAlphaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSourceAlphaInputWithDefaults

`func NewUpdateSourceAlphaInputWithDefaults() *UpdateSourceAlphaInput`

NewUpdateSourceAlphaInputWithDefaults instantiates a new UpdateSourceAlphaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSourceAlphaInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSourceAlphaInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSourceAlphaInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSourceAlphaInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateSourceAlphaInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateSourceAlphaInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateSourceAlphaInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateSourceAlphaInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSlug

`func (o *UpdateSourceAlphaInput) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *UpdateSourceAlphaInput) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *UpdateSourceAlphaInput) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *UpdateSourceAlphaInput) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSettings

`func (o *UpdateSourceAlphaInput) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateSourceAlphaInput) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateSourceAlphaInput) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UpdateSourceAlphaInput) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


