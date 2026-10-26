# UpdateSourceV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | An optional human-readable name to associate with the Source.  Config API note: equal to &#x60;displayName&#x60;. | [optional] 
**Enabled** | Pointer to **bool** | Enable to allow the Source to send data. | [optional] 
**Slug** | Pointer to **string** | The slug that identifies the Source.  Config API note: equal to &#x60;name&#x60;. | [optional] 
**Settings** | Pointer to **map[string]interface{}** | A key-value object that contains instance-specific settings for a Source.  The &#x60;options&#x60; field in the Source metadata defines the schema of this object. | [optional] 

## Methods

### NewUpdateSourceV1Input

`func NewUpdateSourceV1Input() *UpdateSourceV1Input`

NewUpdateSourceV1Input instantiates a new UpdateSourceV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSourceV1InputWithDefaults

`func NewUpdateSourceV1InputWithDefaults() *UpdateSourceV1Input`

NewUpdateSourceV1InputWithDefaults instantiates a new UpdateSourceV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSourceV1Input) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSourceV1Input) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSourceV1Input) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSourceV1Input) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateSourceV1Input) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateSourceV1Input) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateSourceV1Input) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateSourceV1Input) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSlug

`func (o *UpdateSourceV1Input) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *UpdateSourceV1Input) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *UpdateSourceV1Input) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *UpdateSourceV1Input) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSettings

`func (o *UpdateSourceV1Input) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateSourceV1Input) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateSourceV1Input) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UpdateSourceV1Input) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


