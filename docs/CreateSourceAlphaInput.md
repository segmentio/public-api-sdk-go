# CreateSourceAlphaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** | The slug by which to identify the Source in the Segment app. | 
**Enabled** | **bool** | Enable to allow this Source to send data. Defaults to true. | 
**Name** | Pointer to **string** | An optional human-readable name for this Source. | [optional] 
**MetadataId** | **string** | The id of the Source metadata from which this instance of the Source derives.  All Source metadata is available under &#x60;/catalog/sources&#x60;. | 
**Settings** | Pointer to **map[string]interface{}** | A key-value object that contains instance-specific settings for a Source.  The &#x60;options&#x60; field in the Source metadata defines the schema of this object. | [optional] 

## Methods

### NewCreateSourceAlphaInput

`func NewCreateSourceAlphaInput(slug string, enabled bool, metadataId string, ) *CreateSourceAlphaInput`

NewCreateSourceAlphaInput instantiates a new CreateSourceAlphaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSourceAlphaInputWithDefaults

`func NewCreateSourceAlphaInputWithDefaults() *CreateSourceAlphaInput`

NewCreateSourceAlphaInputWithDefaults instantiates a new CreateSourceAlphaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *CreateSourceAlphaInput) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CreateSourceAlphaInput) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CreateSourceAlphaInput) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetEnabled

`func (o *CreateSourceAlphaInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateSourceAlphaInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateSourceAlphaInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetName

`func (o *CreateSourceAlphaInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSourceAlphaInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSourceAlphaInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateSourceAlphaInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadataId

`func (o *CreateSourceAlphaInput) GetMetadataId() string`

GetMetadataId returns the MetadataId field if non-nil, zero value otherwise.

### GetMetadataIdOk

`func (o *CreateSourceAlphaInput) GetMetadataIdOk() (*string, bool)`

GetMetadataIdOk returns a tuple with the MetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataId

`func (o *CreateSourceAlphaInput) SetMetadataId(v string)`

SetMetadataId sets MetadataId field to given value.


### GetSettings

`func (o *CreateSourceAlphaInput) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateSourceAlphaInput) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateSourceAlphaInput) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CreateSourceAlphaInput) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


