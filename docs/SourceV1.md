# SourceV1

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

### NewSourceV1

`func NewSourceV1(id string, slug string, metadata SourceMetadataV1, workspaceId string, enabled bool, writeKeys []string, labels []LabelV1, ) *SourceV1`

NewSourceV1 instantiates a new SourceV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceV1WithDefaults

`func NewSourceV1WithDefaults() *SourceV1`

NewSourceV1WithDefaults instantiates a new SourceV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SourceV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceV1) SetId(v string)`

SetId sets Id field to given value.


### GetSlug

`func (o *SourceV1) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *SourceV1) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *SourceV1) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *SourceV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceV1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourceV1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadata

`func (o *SourceV1) GetMetadata() SourceMetadataV1`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SourceV1) GetMetadataOk() (*SourceMetadataV1, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SourceV1) SetMetadata(v SourceMetadataV1)`

SetMetadata sets Metadata field to given value.


### GetWorkspaceId

`func (o *SourceV1) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *SourceV1) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *SourceV1) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetEnabled

`func (o *SourceV1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SourceV1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SourceV1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetWriteKeys

`func (o *SourceV1) GetWriteKeys() []string`

GetWriteKeys returns the WriteKeys field if non-nil, zero value otherwise.

### GetWriteKeysOk

`func (o *SourceV1) GetWriteKeysOk() (*[]string, bool)`

GetWriteKeysOk returns a tuple with the WriteKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteKeys

`func (o *SourceV1) SetWriteKeys(v []string)`

SetWriteKeys sets WriteKeys field to given value.


### GetSettings

`func (o *SourceV1) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SourceV1) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SourceV1) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SourceV1) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetLabels

`func (o *SourceV1) GetLabels() []LabelV1`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SourceV1) GetLabelsOk() (*[]LabelV1, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SourceV1) SetLabels(v []LabelV1)`

SetLabels sets Labels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


