# DestinationV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of this instance of a Destination.  Config API note: analogous to &#x60;name&#x60;. | 
**Name** | Pointer to **string** | The name of this instance of a Destination.  Config API note: equal to &#x60;displayName&#x60;. | [optional] 
**Enabled** | **bool** | Whether this instance of a Destination receives data. | 
**Metadata** | [**DestinationMetadataV1**](DestinationMetadataV1.md) |  | 
**SourceId** | **string** | The id of a Source connected to this instance of a Destination.  Config API note: analogous to &#x60;parent&#x60;. | 
**Settings** | **map[string]interface{}** | The collection of settings associated with a Destination.  Config API note: equal to &#x60;config&#x60;. | 

## Methods

### NewDestinationV1

`func NewDestinationV1(id string, enabled bool, metadata DestinationMetadataV1, sourceId string, settings map[string]interface{}, ) *DestinationV1`

NewDestinationV1 instantiates a new DestinationV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationV1WithDefaults

`func NewDestinationV1WithDefaults() *DestinationV1`

NewDestinationV1WithDefaults instantiates a new DestinationV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DestinationV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DestinationV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DestinationV1) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DestinationV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationV1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DestinationV1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *DestinationV1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DestinationV1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DestinationV1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMetadata

`func (o *DestinationV1) GetMetadata() DestinationMetadataV1`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DestinationV1) GetMetadataOk() (*DestinationMetadataV1, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DestinationV1) SetMetadata(v DestinationMetadataV1)`

SetMetadata sets Metadata field to given value.


### GetSourceId

`func (o *DestinationV1) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DestinationV1) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DestinationV1) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSettings

`func (o *DestinationV1) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *DestinationV1) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *DestinationV1) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


