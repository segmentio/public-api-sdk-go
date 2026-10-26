# SourceAlpha

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the Source.  Config API note: analogous to &#x60;name&#x60;. | 
**Slug** | **string** | The slug used to identify the Source in the Segment app.  Config API note: equal to &#x60;name&#x60;. | 
**Name** | Pointer to **string** | The name of the Source.  Config API note: equal to &#x60;displayName&#x60;. | [optional] 
**Metadata** | [**SourceMetadataV1**](SourceMetadataV1.md) |  | 
**WorkspaceId** | **string** | The id of the Workspace that owns the Source.  Config API note: equal to &#x60;parent&#x60;. | 
**Enabled** | **bool** | Enable to receive data from the Source. | 
**WriteKeys** | **[]string** | The write keys used to send data from the Source. This field is left empty when the current token does not have the &#39;source admin&#39; permission. | 
**Settings** | Pointer to **map[string]interface{}** | A key-value object that contains instance-specific settings for a Source.  The &#x60;options&#x60; field in the Source metadata defines the schema of this object. | [optional] 
**Labels** | [**[]LabelV1**](LabelV1.md) | A list of labels applied to the Source. | 

## Methods

### NewSourceAlpha

`func NewSourceAlpha(id string, slug string, metadata SourceMetadataV1, workspaceId string, enabled bool, writeKeys []string, labels []LabelV1, ) *SourceAlpha`

NewSourceAlpha instantiates a new SourceAlpha object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAlphaWithDefaults

`func NewSourceAlphaWithDefaults() *SourceAlpha`

NewSourceAlphaWithDefaults instantiates a new SourceAlpha object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SourceAlpha) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceAlpha) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceAlpha) SetId(v string)`

SetId sets Id field to given value.


### GetSlug

`func (o *SourceAlpha) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *SourceAlpha) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *SourceAlpha) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *SourceAlpha) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceAlpha) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceAlpha) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourceAlpha) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadata

`func (o *SourceAlpha) GetMetadata() SourceMetadataV1`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SourceAlpha) GetMetadataOk() (*SourceMetadataV1, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SourceAlpha) SetMetadata(v SourceMetadataV1)`

SetMetadata sets Metadata field to given value.


### GetWorkspaceId

`func (o *SourceAlpha) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *SourceAlpha) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *SourceAlpha) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetEnabled

`func (o *SourceAlpha) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SourceAlpha) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SourceAlpha) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetWriteKeys

`func (o *SourceAlpha) GetWriteKeys() []string`

GetWriteKeys returns the WriteKeys field if non-nil, zero value otherwise.

### GetWriteKeysOk

`func (o *SourceAlpha) GetWriteKeysOk() (*[]string, bool)`

GetWriteKeysOk returns a tuple with the WriteKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteKeys

`func (o *SourceAlpha) SetWriteKeys(v []string)`

SetWriteKeys sets WriteKeys field to given value.


### GetSettings

`func (o *SourceAlpha) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SourceAlpha) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SourceAlpha) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SourceAlpha) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetLabels

`func (o *SourceAlpha) GetLabels() []LabelV1`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SourceAlpha) GetLabelsOk() (*[]LabelV1, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SourceAlpha) SetLabels(v []LabelV1)`

SetLabels sets Labels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


