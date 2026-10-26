# CreateSourceV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** | The slug by which to identify the Source in the Segment app. | 
**Enabled** | **bool** | Enable to allow this Source to send data. Defaults to true. | 
**MetadataId** | **string** | The id of the Source metadata from which this instance of the Source derives.  All Source metadata is available under &#x60;/catalog/sources&#x60;. | 
**Settings** | Pointer to **map[string]interface{}** | A key-value object that contains instance-specific settings for a Source.  The &#x60;options&#x60; field in the Source metadata defines the schema of this object. | [optional] 
**DisconnectAllWarehouses** | Pointer to **bool** | Whether to disconnect all Warehouses from the Source. | [optional] 

## Methods

### NewCreateSourceV1Input

`func NewCreateSourceV1Input(slug string, enabled bool, metadataId string, ) *CreateSourceV1Input`

NewCreateSourceV1Input instantiates a new CreateSourceV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSourceV1InputWithDefaults

`func NewCreateSourceV1InputWithDefaults() *CreateSourceV1Input`

NewCreateSourceV1InputWithDefaults instantiates a new CreateSourceV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *CreateSourceV1Input) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CreateSourceV1Input) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CreateSourceV1Input) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetEnabled

`func (o *CreateSourceV1Input) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateSourceV1Input) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateSourceV1Input) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMetadataId

`func (o *CreateSourceV1Input) GetMetadataId() string`

GetMetadataId returns the MetadataId field if non-nil, zero value otherwise.

### GetMetadataIdOk

`func (o *CreateSourceV1Input) GetMetadataIdOk() (*string, bool)`

GetMetadataIdOk returns a tuple with the MetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataId

`func (o *CreateSourceV1Input) SetMetadataId(v string)`

SetMetadataId sets MetadataId field to given value.


### GetSettings

`func (o *CreateSourceV1Input) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateSourceV1Input) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateSourceV1Input) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CreateSourceV1Input) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetDisconnectAllWarehouses

`func (o *CreateSourceV1Input) GetDisconnectAllWarehouses() bool`

GetDisconnectAllWarehouses returns the DisconnectAllWarehouses field if non-nil, zero value otherwise.

### GetDisconnectAllWarehousesOk

`func (o *CreateSourceV1Input) GetDisconnectAllWarehousesOk() (*bool, bool)`

GetDisconnectAllWarehousesOk returns a tuple with the DisconnectAllWarehouses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectAllWarehouses

`func (o *CreateSourceV1Input) SetDisconnectAllWarehouses(v bool)`

SetDisconnectAllWarehouses sets DisconnectAllWarehouses field to given value.

### HasDisconnectAllWarehouses

`func (o *CreateSourceV1Input) HasDisconnectAllWarehouses() bool`

HasDisconnectAllWarehouses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


